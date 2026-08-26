// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"testing"
)

func TestTopology_DeterministicRouting(t *testing.T) {
	nodes := []NodeConfig{
		{
			NodeID:          "redhost",
			Addr:            "213.32.71.129:4455",
			CertFingerprint: "sha256:d8a9f1b2c3d4e5f60718293a4b5c6d7e8f90123456789abcdef0123456789abc",
			LeaderShards:    []ShardRange{{Start: 0, End: 7}},
			FollowerShards:  []ShardRange{{Start: 8, End: 15}},
		},
		{
			NodeID:          "redbo",
			Addr:            "127.0.0.1:4456",
			CertFingerprint: "sha256:a1b2c3d4e5f60718293a4b5c6d7e8f90123456789abcdef0123456789abcdef0",
			LeaderShards:    []ShardRange{{Start: 8, End: 15}},
			FollowerShards:  []ShardRange{{Start: 0, End: 7}},
		},
	}

	// Sur redhost : Shards 0..7 sont locaux, 8..15 sont distants (redbo)
	topoHost := NewClusterTopology("horos55-bus", 16, "redhost", nodes)

	for s := 0; s <= 7; s++ {
		leaderID, addr, isLocal, err := topoHost.RouteShard(s)
		if err != nil {
			t.Fatalf("shard %d route failed: %v", s, err)
		}
		if leaderID != "redhost" || addr != "213.32.71.129:4455" || !isLocal {
			t.Fatalf("shard %d wrong route: got %s (%s, local=%v)", s, leaderID, addr, isLocal)
		}
		if topoHost.IsFollowerForShard(s) {
			t.Fatalf("redhost should not be follower for shard %d", s)
		}
	}

	for s := 8; s <= 15; s++ {
		leaderID, addr, isLocal, err := topoHost.RouteShard(s)
		if err != nil {
			t.Fatalf("shard %d route failed: %v", s, err)
		}
		if leaderID != "redbo" || addr != "127.0.0.1:4456" || isLocal {
			t.Fatalf("shard %d wrong route: got %s (%s, local=%v)", s, leaderID, addr, isLocal)
		}
		if !topoHost.IsFollowerForShard(s) {
			t.Fatalf("redhost should be follower for shard %d", s)
		}
	}

	// Hors limites
	if _, _, _, err := topoHost.RouteShard(16); err == nil {
		t.Fatalf("shard 16 should return out of bounds error")
	}
}

func TestTopology_ZeroAlloc(t *testing.T) {
	nodes := []NodeConfig{
		{
			NodeID:       "redhost",
			Addr:         "213.32.71.129:4455",
			LeaderShards: []ShardRange{{Start: 0, End: 7}},
		},
	}
	topo := NewClusterTopology("horos55-bus", 8, "redhost", nodes)

	allocs := testing.AllocsPerRun(1000, func() {
		_, _, _, _ = topo.RouteShard(3)
	})

	if allocs != 0 {
		t.Fatalf("RouteShard allocated %f allocs/op, want 0", allocs)
	}
}
