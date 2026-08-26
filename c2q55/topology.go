// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrShardNotAssigned = errors.New("c2q55/topology: shard not assigned in topology")
	ErrNotLeader        = errors.New("c2q55/topology: local node is not leader for this shard")
)

// ShardRange définit une plage inclusive de shards [Start, End].
type ShardRange struct {
	Start int
	End   int
}

// NodeConfig définit la configuration d'un nœud dans la topologie statique.
type NodeConfig struct {
	NodeID          string
	Addr            string
	CertFingerprint string
	LeaderShards    []ShardRange
	FollowerShards  []ShardRange
}

// ClusterTopology représente la table de routage ARCHTIME d'un cluster c2q55.
type ClusterTopology struct {
	mu           sync.RWMutex
	ClusterName  string
	TotalShards  int
	LocalNodeID  string
	nodes        map[string]*NodeConfig
	shardLeaders []string // Indexé par ShardID -> NodeID du leader
}

// NewClusterTopology construit une topologie à partir d'une liste statique de configurations de nœuds.
func NewClusterTopology(clusterName string, totalShards int, localNodeID string, nodes []NodeConfig) *ClusterTopology {
	ct := &ClusterTopology{
		ClusterName:  clusterName,
		TotalShards:  totalShards,
		LocalNodeID:  localNodeID,
		nodes:        make(map[string]*NodeConfig),
		shardLeaders: make([]string, totalShards),
	}

	for i := range nodes {
		n := &nodes[i]
		ct.nodes[n.NodeID] = n

		for _, r := range n.LeaderShards {
			for s := r.Start; s <= r.End && s < totalShards; s++ {
				ct.shardLeaders[s] = n.NodeID
			}
		}
	}

	return ct
}

// RouteShard retourne le nœud leader et son adresse pour un shardID donné.
// Retourne isLocal=true si le nœud local est le leader de ce shard.
func (ct *ClusterTopology) RouteShard(shardID int) (leaderID, leaderAddr string, isLocal bool, err error) {
	if ct == nil {
		return "", "", false, ErrShardNotAssigned
	}
	if shardID < 0 || shardID >= ct.TotalShards {
		return "", "", false, fmt.Errorf("%w: shard %d out of bounds [0, %d)", ErrShardNotAssigned, shardID, ct.TotalShards)
	}

	ct.mu.RLock()
	defer ct.mu.RUnlock()

	leaderID = ct.shardLeaders[shardID]
	if leaderID == "" {
		return "", "", false, fmt.Errorf("%w: shard %d has no configured leader", ErrShardNotAssigned, shardID)
	}

	cfg, exists := ct.nodes[leaderID]
	if !exists {
		return leaderID, "", (leaderID == ct.LocalNodeID), fmt.Errorf("%w: leader node %s has no address config", ErrShardNotAssigned, leaderID)
	}

	return leaderID, cfg.Addr, (leaderID == ct.LocalNodeID), nil
}

// IsFollowerForShard vérifie si le nœud local est configuré comme follower/réplicat pour ce shard.
func (ct *ClusterTopology) IsFollowerForShard(shardID int) bool {
	if ct == nil {
		return false
	}
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	cfg, exists := ct.nodes[ct.LocalNodeID]
	if !exists {
		return false
	}

	for _, r := range cfg.FollowerShards {
		if shardID >= r.Start && shardID <= r.End {
			return true
		}
	}
	return false
}
