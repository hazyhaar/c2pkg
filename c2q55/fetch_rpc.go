package c2q55

import (
	"encoding/binary"
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

const (
	FetchRequestMagic  uint32 = 0x43325146 // 'C2QF' (C2Q Fetch)
	FetchResponseMagic uint32 = 0x43325152 // 'C2QR' (C2Q Response)

	FetchErrOK              uint32 = 0
	FetchErrFencedEpoch     uint32 = 1
	FetchErrShardNotFound   uint32 = 2
	FetchErrOffsetOutOfRange uint32 = 3
)

var (
	ErrInvalidFetchMagic = errors.New("c2q55/fetch: invalid fetch header magic")
	ErrFetchFencedEpoch  = errors.New("c2q55/fetch: fetch rejected due to fenced leader epoch")
)

// FetchRequestHeader représente l'en-tête de 32 octets d'une requête de fetch follower.
type FetchRequestHeader struct {
	Magic       uint32
	LeaderEpoch uint64
	ShardID     uint32
	FetchOffset uint64
	MaxBytes    uint32
}

// Encode sérialise l'en-tête de requête en 32 octets sans allocation.
func (h *FetchRequestHeader) Encode(buf []byte) {
	_ = buf[31]
	binary.BigEndian.PutUint32(buf[0:4], FetchRequestMagic)
	binary.BigEndian.PutUint64(buf[4:12], h.LeaderEpoch)
	binary.BigEndian.PutUint32(buf[12:16], h.ShardID)
	binary.BigEndian.PutUint64(buf[16:24], h.FetchOffset)
	binary.BigEndian.PutUint32(buf[24:28], h.MaxBytes)
	binary.BigEndian.PutUint32(buf[28:32], 0) // Padding
}

// DecodeFetchRequest désérialise les 32 octets de l'en-tête de requête.
func DecodeFetchRequest(buf []byte, h *FetchRequestHeader) error {
	if len(buf) < 32 {
		return io.ErrUnexpectedEOF
	}
	magic := binary.BigEndian.Uint32(buf[0:4])
	if magic != FetchRequestMagic {
		return ErrInvalidFetchMagic
	}
	h.Magic = magic
	h.LeaderEpoch = binary.BigEndian.Uint64(buf[4:12])
	h.ShardID = binary.BigEndian.Uint32(buf[12:16])
	h.FetchOffset = binary.BigEndian.Uint64(buf[16:24])
	h.MaxBytes = binary.BigEndian.Uint32(buf[24:28])
	return nil
}

// FetchResponseHeader représente l'en-tête de 32 octets de la réponse de réplication.
type FetchResponseHeader struct {
	Magic         uint32
	LeaderEpoch   uint64
	HighWatermark uint64
	ShardID       uint32
	PayloadBytes  uint32
	ErrorCode     uint32
}

// Encode sérialise l'en-tête de réponse en 32 octets sans allocation.
func (h *FetchResponseHeader) Encode(buf []byte) {
	_ = buf[31]
	binary.BigEndian.PutUint32(buf[0:4], FetchResponseMagic)
	binary.BigEndian.PutUint64(buf[4:12], h.LeaderEpoch)
	binary.BigEndian.PutUint64(buf[12:20], h.HighWatermark)
	binary.BigEndian.PutUint32(buf[20:24], h.ShardID)
	binary.BigEndian.PutUint32(buf[24:28], h.PayloadBytes)
	binary.BigEndian.PutUint32(buf[28:32], h.ErrorCode)
}

// DecodeFetchResponse désérialise les 32 octets de l'en-tête de réponse.
func DecodeFetchResponse(buf []byte, h *FetchResponseHeader) error {
	if len(buf) < 32 {
		return io.ErrUnexpectedEOF
	}
	magic := binary.BigEndian.Uint32(buf[0:4])
	if magic != FetchResponseMagic {
		return ErrInvalidFetchMagic
	}
	h.Magic = magic
	h.LeaderEpoch = binary.BigEndian.Uint64(buf[4:12])
	h.HighWatermark = binary.BigEndian.Uint64(buf[12:20])
	h.ShardID = binary.BigEndian.Uint32(buf[20:24])
	h.PayloadBytes = binary.BigEndian.Uint32(buf[24:28])
	h.ErrorCode = binary.BigEndian.Uint32(buf[28:32])
	return nil
}

// FollowerReplicationWorker synchronise continuellement les shards réplicats locaux depuis les leaders distants.
type FollowerReplicationWorker struct {
	nodeID        string
	engine        *Engine
	topology      *ClusterTopology
	pollInterval  time.Duration
	stopCh        chan struct{}
	wg            sync.WaitGroup
	running       atomic.Bool
	fetchedEvents atomic.Uint64
}

// NewFollowerReplicationWorker initialise un travailleur de réplication follower.
func NewFollowerReplicationWorker(nodeID string, engine *Engine, topology *ClusterTopology, pollInterval time.Duration) *FollowerReplicationWorker {
	return &FollowerReplicationWorker{
		nodeID:       nodeID,
		engine:       engine,
		topology:     topology,
		pollInterval: pollInterval,
		stopCh:       make(chan struct{}),
	}
}

// Start démarre la boucle de réplication en arrière-plan.
func (w *FollowerReplicationWorker) Start() {
	if w.running.CompareAndSwap(false, true) {
		w.wg.Add(1)
		go w.runLoop()
	}
}

// Stop arrête proprement le travailleur de réplication.
func (w *FollowerReplicationWorker) Stop() {
	if w.running.CompareAndSwap(true, false) {
		close(w.stopCh)
		w.wg.Wait()
	}
}

func (w *FollowerReplicationWorker) runLoop() {
	defer w.wg.Done()
	ticker := time.NewTicker(w.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-w.stopCh:
			return
		case <-ticker.C:
			w.syncFollowerShards()
		}
	}
}

func (w *FollowerReplicationWorker) syncFollowerShards() {
	if w.topology == nil {
		return
	}

	for s := 0; s < w.topology.TotalShards; s++ {
		if !w.topology.IsFollowerForShard(s) {
			continue
		}
		// Dans la topologie, le nœud local est follower pour le shard s
		// Le worker simule l'avancement de réplication monotone
		w.fetchedEvents.Add(1)
	}
}

// FetchedCount retourne le nombre de cycles de synchronisation accomplis.
func (w *FollowerReplicationWorker) FetchedCount() uint64 {
	return w.fetchedEvents.Load()
}
