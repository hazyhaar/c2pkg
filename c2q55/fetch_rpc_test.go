// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"testing"
	"time"
)

func TestFetchRPC_BinaryHeaderRoundtrip(t *testing.T) {
	req := FetchRequestHeader{
		LeaderEpoch: 42,
		ShardID:     7,
		FetchOffset: 1048576,
		MaxBytes:    65536,
	}

	var buf [32]byte
	req.Encode(buf[:])

	var decodedReq FetchRequestHeader
	if err := DecodeFetchRequest(buf[:], &decodedReq); err != nil {
		t.Fatalf("DecodeFetchRequest failed: %v", err)
	}

	if decodedReq.LeaderEpoch != req.LeaderEpoch ||
		decodedReq.ShardID != req.ShardID ||
		decodedReq.FetchOffset != req.FetchOffset ||
		decodedReq.MaxBytes != req.MaxBytes {
		t.Fatalf("DecodeFetchRequest mismatch: got %+v, want %+v", decodedReq, req)
	}

	resp := FetchResponseHeader{
		LeaderEpoch:   42,
		HighWatermark: 1050000,
		ShardID:       7,
		PayloadBytes:  4096,
		ErrorCode:     FetchErrOK,
	}

	var respBuf [32]byte
	resp.Encode(respBuf[:])

	var decodedResp FetchResponseHeader
	if err := DecodeFetchResponse(respBuf[:], &decodedResp); err != nil {
		t.Fatalf("DecodeFetchResponse failed: %v", err)
	}

	if decodedResp.LeaderEpoch != resp.LeaderEpoch ||
		decodedResp.HighWatermark != resp.HighWatermark ||
		decodedResp.ShardID != resp.ShardID ||
		decodedResp.PayloadBytes != resp.PayloadBytes ||
		decodedResp.ErrorCode != resp.ErrorCode {
		t.Fatalf("DecodeFetchResponse mismatch: got %+v, want %+v", decodedResp, resp)
	}
}

func TestFetchRPC_ZeroAlloc(t *testing.T) {
	req := FetchRequestHeader{LeaderEpoch: 1, ShardID: 2, FetchOffset: 3, MaxBytes: 4}
	var buf [32]byte
	var decoded FetchRequestHeader

	allocs := testing.AllocsPerRun(1000, func() {
		req.Encode(buf[:])
		_ = DecodeFetchRequest(buf[:], &decoded)
	})

	if allocs != 0 {
		t.Fatalf("FetchRequest roundtrip allocated %f allocs/op, want 0", allocs)
	}
}

func TestFollowerReplicationWorker_Lifecycle(t *testing.T) {
	nodes := []NodeConfig{
		{
			NodeID:          "redhost",
			Addr:            "127.0.0.1:4455",
			LeaderShards:    []ShardRange{{Start: 0, End: 3}},
			FollowerShards:  []ShardRange{{Start: 4, End: 7}},
		},
	}
	topo := NewClusterTopology("test-bus", 8, "redhost", nodes)
	worker := NewFollowerReplicationWorker("redhost", nil, topo, 5*time.Millisecond)

	worker.Start()
	time.Sleep(30 * time.Millisecond)
	worker.Stop()

	if worker.FetchedCount() == 0 {
		t.Fatalf("worker should have executed synchronization cycles")
	}
}
