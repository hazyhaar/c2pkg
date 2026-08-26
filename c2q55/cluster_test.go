// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"errors"
	"testing"
	"time"
)

// TestCluster_FailureDetectionAndQuorumElection teste la détection de mort et l'élection de leader par quorum (3 nœuds).
func TestCluster_FailureDetectionAndQuorumElection(t *testing.T) {
	peerAddrs := map[string]string{
		"node-A": "127.0.0.1:8441",
		"node-B": "127.0.0.1:8442",
		"node-C": "127.0.0.1:8443",
	}

	const deathTimeout = 50 * time.Millisecond

	nodeA := NewClusterNode("node-A", peerAddrs, deathTimeout)
	nodeB := NewClusterNode("node-B", peerAddrs, deathTimeout)
	nodeC := NewClusterNode("node-C", peerAddrs, deathTimeout)

	// Partition 0 initialisée avec Node A comme Leader (Epoch 1)
	nodeA.RegisterPartition(0, RoleLeader, "node-A", 1)
	nodeB.RegisterPartition(0, RoleFollower, "node-A", 1)
	nodeC.RegisterPartition(0, RoleFollower, "node-A", 1)

	// Simuler la vivacité normale
	nodeB.RecordHeartbeat("node-A", 1)
	nodeC.RecordHeartbeat("node-A", 1)

	// "Mort" du Node A (Node C continue d'émettre ses heartbeats vers B)
	time.Sleep(60 * time.Millisecond)
	nodeB.RecordHeartbeat("node-C", 1)
	time.Sleep(25 * time.Millisecond)
	nowDead := uint64(time.Now().UnixNano())

	// Node B constate la mort de A uniquement
	deadOnB := nodeB.CheckLiveness(nowDead)
	if len(deadOnB) != 1 || deadOnB[0] != "node-A" {
		t.Fatalf("Node B failed to detect death of node-A: %v", deadOnB)
	}

	electedOnB, newEpochB, err := nodeB.ElectLeaderIfLeaderDead(0, nowDead)
	if err != nil || !electedOnB {
		t.Fatalf("Node B failed to get elected as leader: err=%v", err)
	}
	if newEpochB != 2 {
		t.Fatalf("Leader epoch on Node B must be 2, got: %d", newEpochB)
	}

	// Node C constate également l'époque 2 et se synchronise
	nodeC.RecordHeartbeat("node-B", 2)
	psC := nodeC.partitions[0]
	psC.StepDown("node-B", 2)

	if psC.LeaderID != "node-B" || psC.LeaderEpoch.Load() != 2 {
		t.Fatalf("Node C failed to step down to new leader node-B with epoch 2")
	}
}

// TestCluster_SplitBrainFencingByLeaderEpoch teste la protection absolue contre le cerveau divisé (Split-Brain).
func TestCluster_SplitBrainFencingByLeaderEpoch(t *testing.T) {
	// Partition 0 sur le Leader B (Epoch 2 après failover)
	psB := NewPartitionState(0, RoleLeader, "node-B", 2)

	// 1. Écriture légitime avec l'Époque 2 en cours
	offset, err := psB.ProposeWrite(2)
	if err != nil || offset != 1 {
		t.Fatalf("Legitimate write with Epoch 2 failed: %v", err)
	}

	// 2. L'ancien Leader A (revenu d'un split-brain) tente une écriture avec l'ancienne Époque 1
	_, errFenced := psB.ProposeWrite(1)
	if errFenced == nil || !errors.Is(errFenced, ErrFencedLeaderEpoch) {
		t.Fatalf("CRITICAL SPLIT-BRAIN FLAW: Server accepted stale write from old leader! err: %v", errFenced)
	}

	// 3. Réception d'une écriture avec époque future (ex: Époque 3)
	psB.StepDown("node-C", 3)
	if psB.Role != RoleFollower || psB.LeaderEpoch.Load() != 3 {
		t.Fatalf("Node B failed to step down on higher epoch")
	}
}

// TestCluster_HighWatermarkQuorumReplication teste la progression stricte du High-Watermark sur quorum.
func TestCluster_HighWatermarkQuorumReplication(t *testing.T) {
	// Cluster de 3 nœuds (Quorum = 2)
	psLeader := NewPartitionState(0, RoleLeader, "node-A", 1)

	// 100 messages produits sur le Leader
	for i := 0; i < 100; i++ {
		_, _ = psLeader.ProposeWrite(1)
	}
	if psLeader.LogEndOffset.Load() != 100 {
		t.Fatalf("LogEndOffset mismatch: got %d, want 100", psLeader.LogEndOffset.Load())
	}

	// Initialement, aucun follower n'a confirmé : HW = 0
	if psLeader.HighWatermark.Load() != 0 {
		t.Fatalf("Initial HighWatermark must be 0, got: %d", psLeader.HighWatermark.Load())
	}

	// Follower 1 confirme jusqu'à l'offset 60 (Quorum 2/3 atteint jusqu'à 60 : Leader(100) + Follower1(60))
	hw := psLeader.UpdateFollowerProgress("node-B", 60, 2)
	if hw != 60 {
		t.Fatalf("HighWatermark mismatch after Follower 1: got %d, want 60", hw)
	}

	// Follower 2 confirme jusqu'à l'offset 90 (Quorum 2/3 atteint jusqu'à 90 : Leader(100) + Follower2(90))
	hw = psLeader.UpdateFollowerProgress("node-C", 90, 2)
	if hw != 90 {
		t.Fatalf("HighWatermark mismatch after Follower 2: got %d, want 90", hw)
	}

	// Follower 1 rattrape jusqu'à 100 (Quorum 3/3 à 100)
	hw = psLeader.UpdateFollowerProgress("node-B", 100, 2)
	if hw != 100 {
		t.Fatalf("HighWatermark mismatch after full catchup: got %d, want 100", hw)
	}
}
