package c2q55

import (
	"testing"
)

var (
	ClusterBenchSinkBool bool
	ClusterBenchSinkStr  string
	ClusterBenchSinkU64  uint64
)

// 1. Benchmark ACL Check
func BenchmarkCluster_ACLCheck(b *testing.B) {
	mgr := NewACLManager()
	mgr.AddRule("node-a", "telemetry.sensors", PrivilegePublish)
	mgr.AddRule("node-b", "telemetry.*", PrivilegeConsume)
	mgr.AddRule("admin-node", "*", PrivilegeAdmin)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok := mgr.CheckAccess("node-a", "telemetry.sensors", PrivilegePublish)
		ClusterBenchSinkBool = ok
	}
}

// 2. Benchmark Topology Route
func BenchmarkCluster_TopologyRoute(b *testing.B) {
	nodes := []NodeConfig{
		{
			NodeID:          "redhost",
			Addr:            "213.32.71.129:4455",
			LeaderShards:    []ShardRange{{Start: 0, End: 7}},
			FollowerShards:  []ShardRange{{Start: 8, End: 15}},
		},
		{
			NodeID:          "redbo",
			Addr:            "127.0.0.1:4456",
			LeaderShards:    []ShardRange{{Start: 8, End: 15}},
			FollowerShards:  []ShardRange{{Start: 0, End: 7}},
		},
	}
	topo := NewClusterTopology("horos55-bus", 16, "redhost", nodes)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		leaderID, _, isLocal, _ := topo.RouteShard(i & 15)
		ClusterBenchSinkStr = leaderID
		ClusterBenchSinkBool = isLocal
	}
}

// 3. Benchmark Fetch RPC Encode/Decode
func BenchmarkCluster_FetchRPCRoundtrip(b *testing.B) {
	req := FetchRequestHeader{
		LeaderEpoch: 10,
		ShardID:     3,
		FetchOffset: 65536,
		MaxBytes:    32768,
	}
	var buf [32]byte
	var decoded FetchRequestHeader

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req.Encode(buf[:])
		_ = DecodeFetchRequest(buf[:], &decoded)
		ClusterBenchSinkU64 ^= decoded.FetchOffset
	}
}
