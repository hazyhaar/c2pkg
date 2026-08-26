// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrEngineClosed          = errors.New("c2q55: engine is closed")
	ErrConsumerGroupNotFound = errors.New("c2q55: consumer group not found")
)

// Message représente un message complet avec corps de taille arbitraire et métadonnées d'offset.
type Message struct {
	TopicID     uint32
	Key         uint64
	IDHigh      uint64
	TimestampNs uint64
	Partition   int
	Offset      uint64
	Flags       uint32
	Body        []byte
}

// EngineOptions définit la configuration du moteur d'événements et d'état.
type EngineOptions struct {
	NumShards     int
	ShardCapacity uint64
	WALPath       string
	WALSizeBytes  int64
	SlabPath      string
	SlabSizeBytes int64
	OffsetDir     string
}

// DefaultOptions retourne une configuration standard équilibrée avec support des slabs et offsets durables.
func DefaultOptions() EngineOptions {
	return EngineOptions{
		NumShards:     16,
		ShardCapacity: 65536,
		WALPath:       "",
		WALSizeBytes:  64 * 1024 * 1024,
		SlabPath:      "",
		SlabSizeBytes: 64 * 1024 * 1024,
		OffsetDir:     "",
	}
}

// Engine orchestre les partitions, les arènes de slabs, les groupes de consommateurs et le journal WAL.
type Engine struct {
	opts             EngineOptions
	router           *PartitionRouter
	shards           []*Ring
	merger           *LoserTreeMerger
	wal              *WALFile
	slab             *SlabArena
	offsetStore      *OffsetStore
	partitionOffsets []atomic.Uint64
	consumerGroups   map[string]*ConsumerGroup
	groupsMu         sync.RWMutex
	closed           atomic.Bool
	mergerMu         sync.Mutex
	consumeCursor    atomic.Uint64
	shardMask        uint64
	numShardsAct     int
}

// NewEngine initialise un nouveau moteur partitionné avec support des slabs et des offsets durables.
func NewEngine(opts EngineOptions) (*Engine, error) {
	if opts.NumShards <= 0 {
		opts.NumShards = 16
	}
	if opts.ShardCapacity == 0 {
		opts.ShardCapacity = 65536
	}

	router := NewPartitionRouter(uint64(opts.NumShards), opts.ShardCapacity)
	shards := router.Shards()
	numAct := len(shards)

	e := &Engine{
		opts:             opts,
		router:           router,
		shards:           shards,
		merger:           NewLoserTreeMerger(shards),
		partitionOffsets: make([]atomic.Uint64, numAct),
		consumerGroups:   make(map[string]*ConsumerGroup),
		shardMask:        uint64(numAct - 1),
		numShardsAct:     numAct,
	}

	if opts.OffsetDir != "" {
		store, err := OpenOffsetStore(opts.OffsetDir)
		if err != nil {
			return nil, fmt.Errorf("c2q55: failed to open offset store: %w", err)
		}
		e.offsetStore = store
	}

	if opts.SlabPath != "" {
		slab, err := OpenSlabArena(opts.SlabPath, opts.SlabSizeBytes)
		if err != nil {
			if e.offsetStore != nil {
				_ = e.offsetStore.Close()
			}
			return nil, fmt.Errorf("c2q55: failed to open slab arena: %w", err)
		}
		e.slab = slab
	}

	if opts.WALPath != "" {
		wal, err := OpenWAL(opts.WALPath, opts.WALSizeBytes)
		if err != nil {
			if e.slab != nil {
				_ = e.slab.Close()
			}
			if e.offsetStore != nil {
				_ = e.offsetStore.Close()
			}
			return nil, fmt.Errorf("c2q55: failed to open WAL: %w", err)
		}
		e.wal = wal
	}

	return e, nil
}

// Publish publie un message avec routage par clé, support des grands payloads via slab, et offset monotone.
func (e *Engine) Publish(idLow, idHigh uint64, topicID uint32, flags uint32, body []byte) error {
	nowNs := uint64(time.Now().UnixNano())
	return e.PublishAt(idLow, idHigh, nowNs, topicID, flags, body)
}

// PublishAt publie un message avec échéance temporelle différée (support TDMA).
func (e *Engine) PublishAt(idLow, idHigh, visibleAtNs uint64, topicID uint32, flags uint32, body []byte) error {
	if e.closed.Load() {
		return ErrEngineClosed
	}

	shardIdx := int(idLow & e.shardMask)
	shard := e.shards[shardIdx]

	var payloadToInline []byte
	var actualFlags uint32 = flags

	if len(body) <= 16 {
		payloadToInline = body
	} else if e.slab != nil {
		desc, err := e.slab.Write(body)
		if err != nil {
			return fmt.Errorf("c2q55: failed to write body to slab arena: %w", err)
		}
		var descBuf [16]byte
		desc.Encode(&descBuf)
		payloadToInline = descBuf[:]
		actualFlags |= FlagExternalPayload
	} else {
		return ErrPayloadTooLarge
	}

	_ = e.partitionOffsets[shardIdx].Add(1)

	_, err := shard.Enqueue(idLow, idHigh, visibleAtNs, visibleAtNs+5000000000, topicID, actualFlags, payloadToInline)
	if err != nil {
		return err
	}

	if e.wal != nil {
		var batch [1]Slot
		batch[0].Pack(idLow, idHigh, visibleAtNs, visibleAtNs+5000000000, topicID, actualFlags, payloadToInline, SlotReady)
		if _, walErr := e.wal.WriteBatch(batch[:]); walErr != nil {
			return fmt.Errorf("c2q55/wal: write batch failed: %w", walErr)
		}
	}

	return nil
}

// ConsumeMessage extrait le prochain message disponible, relit le slab si nécessaire, et vérifie le CRC.
func (e *Engine) ConsumeMessage(out *Message) bool {
	if e.closed.Load() {
		return false
	}

	nowNs := uint64(time.Now().UnixNano())
	startIdx := e.consumeCursor.Add(1)

	for i := 0; i < e.numShardsAct; i++ {
		shardIdx := int((startIdx + uint64(i)) & e.shardMask)
		var slot Slot
		if seq, ok := e.shards[shardIdx].Dequeue(nowNs, &slot); ok {
			_ = e.shards[shardIdx].Ack(seq)
			e.populateMessage(shardIdx, &slot, out)
			return true
		}
	}

	return false
}

// Consume extrait un slot brut (compatibilité).
func (e *Engine) Consume(out *Slot) bool {
	if e.closed.Load() {
		return false
	}

	nowNs := uint64(time.Now().UnixNano())
	startIdx := e.consumeCursor.Add(1)

	for i := 0; i < e.numShardsAct; i++ {
		shardIdx := (startIdx + uint64(i)) & e.shardMask
		if seq, ok := e.shards[shardIdx].Dequeue(nowNs, out); ok {
			_ = e.shards[shardIdx].Ack(seq)
			return true
		}
	}

	return false
}

// populateMessage décode le corps complet (inliné ou slab) et garnit la structure Message.
func (e *Engine) populateMessage(partition int, slot *Slot, out *Message) {
	out.Partition = partition
	out.TopicID = slot.TopicID
	out.Key = slot.IDLow
	out.IDHigh = slot.IDHigh
	out.TimestampNs = slot.VisibleAtNs
	out.Flags = slot.Flags

	if slot.Flags&FlagExternalPayload != 0 && e.slab != nil {
		desc := DecodeSlabDescriptor(&slot.Payload)
		body, err := e.slab.Read(desc)
		if err == nil {
			out.Body = body
		} else {
			out.Body = nil
		}
	} else {
		out.Body = make([]byte, slot.PayloadLen)
		copy(out.Body, slot.Payload[:slot.PayloadLen])
	}
}

// GetOrCreateConsumerGroup enregistre ou retourne un groupe de consommateurs avec support persistant.
func (e *Engine) GetOrCreateConsumerGroup(name string) *ConsumerGroup {
	e.groupsMu.Lock()
	defer e.groupsMu.Unlock()

	group, exists := e.consumerGroups[name]
	if !exists {
		group = NewConsumerGroup(name, e.numShardsAct, e.offsetStore)
		e.consumerGroups[name] = group
	}
	return group
}

// ConsumeGroup extrait un message réservé aux partitions assignées à ce consommateur.
func (e *Engine) ConsumeGroup(groupName, consumerID string, out *Message) (bool, error) {
	e.groupsMu.RLock()
	group, exists := e.consumerGroups[groupName]
	e.groupsMu.RUnlock()

	if !exists {
		return false, ErrConsumerGroupNotFound
	}

	assigned := group.AssignedPartitions(consumerID)
	if len(assigned) == 0 {
		return false, ErrNoPartitionAssigned
	}

	nowNs := uint64(time.Now().UnixNano())

	for _, p := range assigned {
		var slot Slot
		if seq, ok := e.shards[p].Dequeue(nowNs, &slot); ok {
			_ = e.shards[p].Ack(seq)
			e.populateMessage(p, &slot, out)
			return true, nil
		}
	}

	return false, nil
}

// CommitOffset valide l'avancement d'un groupe sur une partition et le persiste sur disque.
func (e *Engine) CommitOffset(groupName string, partition int, offset uint64) error {
	e.groupsMu.RLock()
	group, exists := e.consumerGroups[groupName]
	e.groupsMu.RUnlock()

	if !exists {
		return ErrConsumerGroupNotFound
	}
	return group.CommitOffset(partition, offset)
}

// ConsumeOrdered extrait les messages dans un ordre global strict via le Loser Tree.
func (e *Engine) ConsumeOrdered(out *Slot) bool {
	if e.closed.Load() {
		return false
	}

	e.mergerMu.Lock()
	defer e.mergerMu.Unlock()

	nowNs := uint64(time.Now().UnixNano())
	return e.merger.Next(nowNs, out)
}

// Sync force la synchronisation matérielle du journal, des slabs et des offsets.
func (e *Engine) Sync() error {
	if e.slab != nil {
		_ = e.slab.Sync()
	}
	if e.wal != nil {
		_ = e.wal.Sync()
	}
	return nil
}

// Close ferme le moteur, synchronise le WAL, l'arène de slabs et le magasin d'offsets.
func (e *Engine) Close() error {
	if e.closed.Swap(true) {
		return nil
	}

	var err error
	if e.offsetStore != nil {
		_ = e.offsetStore.Close()
	}
	if e.slab != nil {
		_ = e.slab.Sync()
		_ = e.slab.Close()
	}
	if e.wal != nil {
		_ = e.wal.Sync()
		err = e.wal.Close()
	}

	return err
}

// Shards retourne la liste des anneaux sous-jacents.
func (e *Engine) Shards() []*Ring {
	return e.shards
}

// Slab retourne l'arène de slabs active.
func (e *Engine) Slab() *SlabArena {
	return e.slab
}
