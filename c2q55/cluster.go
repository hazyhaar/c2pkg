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
	ErrQuorumNotReached = errors.New("c2q55/cluster: quorum consensus not reached")
)

// ClusterPeer représente un nœud pair du parc horos55.
type ClusterPeer struct {
	NodeID      string
	Addr        string
	LastSeenNs  atomic.Uint64
	Alive       atomic.Bool
	LeaderEpoch atomic.Uint64
}

// ClusterNode orchestre la vivacité, les battements de cœur, et l'élection de leader par époque.
type ClusterNode struct {
	nodeID          string
	peers           map[string]*ClusterPeer
	partitions      map[int]*PartitionState
	deathTimeoutNs  uint64
	mu              sync.RWMutex
	closed          atomic.Bool
}

// NewClusterNode initialise un gestionnaire de cluster sur un nœud.
func NewClusterNode(nodeID string, peerAddrs map[string]string, deathTimeout time.Duration) *ClusterNode {
	peers := make(map[string]*ClusterPeer)
	nowNs := uint64(time.Now().UnixNano())

	for id, addr := range peerAddrs {
		if id != nodeID {
			p := &ClusterPeer{
				NodeID: id,
				Addr:   addr,
			}
			p.LastSeenNs.Store(nowNs)
			p.Alive.Store(true)
			peers[id] = p
		}
	}

	return &ClusterNode{
		nodeID:         nodeID,
		peers:          peers,
		partitions:     make(map[int]*PartitionState),
		deathTimeoutNs: uint64(deathTimeout.Nanoseconds()),
	}
}

// RegisterPartition enregistre une partition sous la surveillance du cluster.
func (cn *ClusterNode) RegisterPartition(partitionID int, role PartitionRole, initialLeader string, initialEpoch uint64) *PartitionState {
	cn.mu.Lock()
	defer cn.mu.Unlock()

	ps := NewPartitionState(partitionID, role, initialLeader, initialEpoch)
	cn.partitions[partitionID] = ps
	return ps
}

// RecordHeartbeat enregistre la vivacité d'un nœud pair distant.
func (cn *ClusterNode) RecordHeartbeat(peerID string, epoch uint64) {
	cn.mu.RLock()
	peer, exists := cn.peers[peerID]
	cn.mu.RUnlock()

	if exists {
		nowNs := uint64(time.Now().UnixNano())
		peer.LastSeenNs.Store(nowNs)
		peer.Alive.Store(true)
		peer.LeaderEpoch.Store(epoch)
	}
}

// CheckLiveness évalue la vivacité des pairs et déclare les nœuds morts si le délai est dépassé.
func (cn *ClusterNode) CheckLiveness(nowNs uint64) []string {
	cn.mu.RLock()
	defer cn.mu.RUnlock()

	deadNodes := make([]string, 0)
	for id, peer := range cn.peers {
		lastSeen := peer.LastSeenNs.Load()
		if nowNs-lastSeen > cn.deathTimeoutNs {
			if peer.Alive.Swap(false) {
				deadNodes = append(deadNodes, id)
			}
		}
	}
	return deadNodes
}

// ElectLeaderIfLeaderDead tente une élection de quorum si le leader actuel est déclaré mort.
func (cn *ClusterNode) ElectLeaderIfLeaderDead(partitionID int, nowNs uint64) (bool, uint64, error) {
	cn.mu.Lock()
	defer cn.mu.Unlock()

	ps, exists := cn.partitions[partitionID]
	if !exists {
		return false, 0, errors.New("c2q55/cluster: partition not found")
	}

	if ps.Role == RoleLeader {
		return false, ps.LeaderEpoch.Load(), nil // Déjà leader
	}

	// Vérifier si le leader actuel est mort
	currentLeader := ps.LeaderID
	if peer, ok := cn.peers[currentLeader]; ok {
		lastSeen := peer.LastSeenNs.Load()
		if nowNs-lastSeen <= cn.deathTimeoutNs {
			return false, ps.LeaderEpoch.Load(), nil // Le leader est encore vivant
		}
	}

	// Calcul du quorum vivant parmi l'ensemble du cluster (nœud courant inclus)
	totalNodes := len(cn.peers) + 1
	quorumNeeded := (totalNodes / 2) + 1
	aliveCount := 1 // Le nœud courant

	for _, peer := range cn.peers {
		if peer.Alive.Load() {
			aliveCount++
		}
	}

	if aliveCount < quorumNeeded {
		return false, ps.LeaderEpoch.Load(), fmt.Errorf("%w: alive=%d, needed=%d", ErrQuorumNotReached, aliveCount, quorumNeeded)
	}

	// Promotion locale et incrément de l'époque
	newEpoch := ps.PromoteLeader(cn.nodeID)
	return true, newEpoch, nil
}
