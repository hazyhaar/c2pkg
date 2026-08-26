// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"math/bits"
)

// PartitionRouter distribue les écritures sur N shards indépendants sans contention.
type PartitionRouter struct {
	shards    []*Ring
	shardMask uint64
	numShards uint64
}

// NewPartitionRouter initialise un routeur multi-shards partitionné par UUIDv7.
func NewPartitionRouter(numShards uint64, shardCapacity uint64) *PartitionRouter {
	if numShards < 1 {
		numShards = 1
	}
	shift := bits.Len64(numShards - 1)
	numShards = uint64(1) << shift

	shards := make([]*Ring, numShards)
	for i := range shards {
		shards[i] = NewRing(shardCapacity)
	}

	return &PartitionRouter{
		shards:    shards,
		shardMask: numShards - 1,
		numShards: numShards,
	}
}

// ShardForID retourne l'anneau assigné à un identifiant UUIDv7.
func (p *PartitionRouter) ShardForID(idLow uint64) *Ring {
	idx := idLow & p.shardMask
	return p.shards[idx]
}

// Route est un alias canonique pour ShardForID.
func (p *PartitionRouter) Route(idLow uint64) *Ring {
	return p.ShardForID(idLow)
}

// Shards retourne l'ensemble des anneaux sous-jacents.
func (p *PartitionRouter) Shards() []*Ring {
	return p.shards
}

// NumShards retourne le nombre total effectif de shards (puissance de deux).
func (p *PartitionRouter) NumShards() uint64 {
	return p.numShards
}
