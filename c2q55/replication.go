// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	ErrNotPartitionLeader  = errors.New("c2q55/replication: node is not leader for this partition")
	ErrFencedLeaderEpoch    = errors.New("c2q55/replication: write rejected due to stale leader epoch (split-brain fenced)")
	ErrOffsetAheadOfLeader  = errors.New("c2q55/replication: requested fetch offset is ahead of leader log end")
)

// PartitionRole définit le rôle d'un nœud pour une partition donnée.
type PartitionRole uint8

const (
	RoleLeader   PartitionRole = 1
	RoleFollower PartitionRole = 2
	RoleObserver PartitionRole = 3
)

// PartitionState maintient l'époque de leadership, les décalages de réplication et le High-Watermark.
type PartitionState struct {
	PartitionID    int
	Role           PartitionRole
	LeaderEpoch    atomic.Uint64
	LeaderID       string
	LogEndOffset   atomic.Uint64
	HighWatermark  atomic.Uint64
	ReplicaOffsets map[string]uint64
	mu             sync.RWMutex
}

// NewPartitionState initialise l'état distribué d'une partition.
func NewPartitionState(partitionID int, role PartitionRole, initialLeader string, initialEpoch uint64) *PartitionState {
	ps := &PartitionState{
		PartitionID:    partitionID,
		Role:           role,
		LeaderID:       initialLeader,
		ReplicaOffsets: make(map[string]uint64),
	}
	ps.LeaderEpoch.Store(initialEpoch)
	return ps
}

// ProposeWrite valide une écriture sur le Leader avec contrôle strict de l'époque.
func (ps *PartitionState) ProposeWrite(incomingEpoch uint64) (uint64, error) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.Role != RoleLeader {
		return 0, ErrNotPartitionLeader
	}

	curEpoch := ps.LeaderEpoch.Load()
	if incomingEpoch < curEpoch {
		return 0, fmt.Errorf("%w: incoming epoch %d < current epoch %d", ErrFencedLeaderEpoch, incomingEpoch, curEpoch)
	}

	offset := ps.LogEndOffset.Add(1)
	return offset, nil
}

// UpdateFollowerProgress met à jour le décalage confirmé par un follower et recalcule le High-Watermark.
func (ps *PartitionState) UpdateFollowerProgress(replicaID string, confirmedOffset uint64, quorumSize int) uint64 {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.Role != RoleLeader {
		return ps.HighWatermark.Load()
	}

	ps.ReplicaOffsets[replicaID] = confirmedOffset

	// Calcul du quorum : tri des offsets confirmés pour extraire la majorité
	offsets := make([]uint64, 0, len(ps.ReplicaOffsets)+1)
	offsets = append(offsets, ps.LogEndOffset.Load()) // Le leader lui-même
	for _, off := range ps.ReplicaOffsets {
		offsets = append(offsets, off)
	}

	// Tri à bulles / sélection simple (petit nombre de répliques <= 7)
	for i := 0; i < len(offsets); i++ {
		for j := i + 1; j < len(offsets); j++ {
			if offsets[i] < offsets[j] {
				offsets[i], offsets[j] = offsets[j], offsets[i]
			}
		}
	}

	// La majorité est située à l'index quorumSize - 1
	if quorumSize <= len(offsets) {
		majorityOffset := offsets[quorumSize-1]
		curHW := ps.HighWatermark.Load()
		if majorityOffset > curHW {
			ps.HighWatermark.Store(majorityOffset)
		}
	}

	return ps.HighWatermark.Load()
}

// PromoteLeader promeut le nœud en nouveau Leader en incrémentant strictement l'époque.
func (ps *PartitionState) PromoteLeader(nodeID string) uint64 {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	newEpoch := ps.LeaderEpoch.Add(1)
	ps.Role = RoleLeader
	ps.LeaderID = nodeID
	ps.ReplicaOffsets = make(map[string]uint64)
	return newEpoch
}

// StepDown rétrograde un ancien leader en Follower suite à la découverte d'une époque supérieure.
func (ps *PartitionState) StepDown(newLeader string, newEpoch uint64) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.Role = RoleFollower
	ps.LeaderID = newLeader
	ps.LeaderEpoch.Store(newEpoch)
}

// ReplicationRecord représente la trame de transmission continue d'un lot de messages répliqués.
type ReplicationRecord struct {
	LeaderEpoch   uint64
	PartitionID   uint16
	OffsetStart   uint64
	OffsetEnd     uint64
	HighWatermark uint64
	SlotCount     uint32
}

// EncodeRecord sérialise les métadonnées de réplication (40 octets).
func (r *ReplicationRecord) Encode(dst *[40]byte) {
	binary.LittleEndian.PutUint64(dst[0:8], r.LeaderEpoch)
	binary.LittleEndian.PutUint16(dst[8:10], r.PartitionID)
	binary.LittleEndian.PutUint64(dst[10:18], r.OffsetStart)
	binary.LittleEndian.PutUint64(dst[18:26], r.OffsetEnd)
	binary.LittleEndian.PutUint64(dst[26:34], r.HighWatermark)
	binary.LittleEndian.PutUint32(dst[34:38], r.SlotCount)
	binary.LittleEndian.PutUint16(dst[38:40], 0) // Padding
}

// DecodeRecord désérialise les métadonnées de réplication.
func DecodeReplicationRecord(src []byte) (ReplicationRecord, error) {
	if len(src) < 40 {
		return ReplicationRecord{}, errors.New("c2q55/replication: buffer too small")
	}

	return ReplicationRecord{
		LeaderEpoch:   binary.LittleEndian.Uint64(src[0:8]),
		PartitionID:   binary.LittleEndian.Uint16(src[8:10]),
		OffsetStart:   binary.LittleEndian.Uint64(src[10:18]),
		OffsetEnd:     binary.LittleEndian.Uint64(src[18:26]),
		HighWatermark: binary.LittleEndian.Uint64(src[26:34]),
		SlotCount:     binary.LittleEndian.Uint32(src[34:38]),
	}, nil
}
