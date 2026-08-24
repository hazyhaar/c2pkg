package c2q55

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	ErrNoPartitionAssigned = errors.New("c2q55/group: no partitions assigned to this consumer")
	ErrOffsetOutOfRange    = errors.New("c2q55/group: requested offset is out of range")
)

// ConsumerGroup coordonne plusieurs consommateurs concurrents sur les partitions avec persistance d'offsets.
type ConsumerGroup struct {
	name             string
	numPartitions    int
	committedOffsets []atomic.Uint64
	activeConsumers  map[string][]int // consumerID -> liste des partitions assignées
	store            *OffsetStore
	mu               sync.RWMutex
}

// NewConsumerGroup initialise un groupe de consommateurs nommé avec persistance d'offsets optionnelle.
func NewConsumerGroup(name string, numPartitions int, store *OffsetStore) *ConsumerGroup {
	if numPartitions <= 0 {
		numPartitions = 16
	}

	g := &ConsumerGroup{
		name:             name,
		numPartitions:    numPartitions,
		committedOffsets: make([]atomic.Uint64, numPartitions),
		activeConsumers:  make(map[string][]int),
		store:            store,
	}

	if store != nil {
		if loaded, err := store.LoadOffsets(name, numPartitions); err == nil {
			for p, off := range loaded {
				g.committedOffsets[p].Store(off)
			}
		}
	}

	return g
}

// RegisterConsumer enregistre un consommateur et rééquilibre les partitions (Rebalance).
func (g *ConsumerGroup) RegisterConsumer(consumerID string) []int {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.activeConsumers[consumerID] = nil
	g.rebalanceLocked()
	return g.activeConsumers[consumerID]
}

// UnregisterConsumer retire un consommateur et rééquilibre les partitions.
func (g *ConsumerGroup) UnregisterConsumer(consumerID string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	delete(g.activeConsumers, consumerID)
	g.rebalanceLocked()
}

// rebalanceLocked distribue les partitions uniformément selon un algorithme Range/Round-Robin.
func (g *ConsumerGroup) rebalanceLocked() {
	numConsumers := len(g.activeConsumers)
	if numConsumers == 0 {
		return
	}

	consumerList := make([]string, 0, numConsumers)
	for id := range g.activeConsumers {
		consumerList = append(consumerList, id)
		g.activeConsumers[id] = nil
	}

	for p := 0; p < g.numPartitions; p++ {
		targetConsumer := consumerList[p%numConsumers]
		g.activeConsumers[targetConsumer] = append(g.activeConsumers[targetConsumer], p)
	}
}

// CommitOffset valide le traitement d'un message sur une partition (sémantique At-least-once) et persiste sur disque.
func (g *ConsumerGroup) CommitOffset(partition int, offset uint64) error {
	if partition < 0 || partition >= g.numPartitions {
		return errors.New("c2q55/group: invalid partition index")
	}

	for {
		cur := g.committedOffsets[partition].Load()
		if offset <= cur {
			return nil
		}
		if g.committedOffsets[partition].CompareAndSwap(cur, offset) {
			if g.store != nil {
				return g.store.CommitOffset(g.name, partition, offset)
			}
			return nil
		}
	}
}

// GetCommittedOffset retourne le dernier offset validé sur une partition.
func (g *ConsumerGroup) GetCommittedOffset(partition int) uint64 {
	if partition < 0 || partition >= g.numPartitions {
		return 0
	}
	return g.committedOffsets[partition].Load()
}

// AssignedPartitions retourne la liste des partitions allouées à un consommateur.
func (g *ConsumerGroup) AssignedPartitions(consumerID string) []int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	partitions := g.activeConsumers[consumerID]
	res := make([]int, len(partitions))
	copy(res, partitions)
	return res
}

// Name retourne le nom du groupe.
func (g *ConsumerGroup) Name() string {
	return g.name
}

// String retourne une synthèse textuelle du groupe.
func (g *ConsumerGroup) String() string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return fmt.Sprintf("ConsumerGroup[name=%s, partitions=%d, consumers=%d]", g.name, g.numPartitions, len(g.activeConsumers))
}
