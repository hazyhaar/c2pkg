package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/hazyhaar/c2pkg/c2q55"
)

// ServerMetricsSnapshot capture les métriques système et moteur sur 1 seconde.
type ServerMetricsSnapshot struct {
	TimestampSec      int64   `json:"timestamp_sec"`
	MsgRxTotal        uint64  `json:"msg_rx_total"`
	MsgRxPerSec       uint64  `json:"msg_rx_per_sec"`
	BytesRxPerSecMB   float64 `json:"bytes_rx_mb_per_sec"`
	CPUUserPct        float64 `json:"cpu_user_pct"`
	CPUSysPct         float64 `json:"cpu_sys_pct"`
	CPUSoftIRQPct     float64 `json:"cpu_softirq_pct"`
	GCPauseNsMax      uint64  `json:"gc_pause_ns_max"`
	GCPauseNsAvg      uint64  `json:"gc_pause_ns_avg"`
	AllocHeapMB       float64 `json:"alloc_heap_mb"`
	NumGC             uint32  `json:"num_gc"`
	EngineDrops       uint64  `json:"engine_drops"`
}

// StageResult résume les performances côté client pour un palier de 10s.
type StageResult struct {
	StageName       string  `json:"stage_name"`
	TargetRateMsgS  int     `json:"target_rate_msg_s"`
	ActualRateMsgS  float64 `json:"actual_rate_msg_s"`
	ActualThroughMB float64 `json:"actual_through_mb_s"`
	TotalSent       uint64  `json:"total_sent"`
	TotalAcked      uint64  `json:"total_acked"`
	TotalErrors     uint64  `json:"total_errors"`
	LatP50Ms        float64 `json:"lat_p50_ms"`
	LatP90Ms        float64 `json:"lat_p90_ms"`
	LatP99Ms        float64 `json:"lat_p99_ms"`
	LatP999Ms       float64 `json:"lat_p999_ms"`
	LatMaxMs        float64 `json:"lat_max_ms"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	mode := os.Args[1]
	switch mode {
	case "server":
		runServerProbe(os.Args[2:])
	case "client":
		runClientLoad(os.Args[2:])
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: c2q55-stratified-bench <server|client> [options]")
	fmt.Println("  server --addr 0.0.0.0:8443 --ca-cert ca.crt --node-cert node.crt --node-key node.key --allow REDBO,redbo --wal /tmp/wal.log --slab /tmp/data.slab")
	fmt.Println("  client --target 213.32.71.129:8443 --server-name redhost --ca-cert ca.crt --node-cert node.crt --node-key node.key")
}

func runServerProbe(args []string) {
	fs := flag.NewFlagSet("server", flag.ExitOnError)
	addr := fs.String("addr", "0.0.0.0:8443", "Bind address for QUIC server")
	caCertPath := fs.String("ca-cert", "", "Path to root Cluster CA PEM")
	nodeCertPath := fs.String("node-cert", "", "Path to node certificate PEM")
	nodeKeyPath := fs.String("node-key", "", "Path to node key PEM")
	allowNodes := fs.String("allow", "REDBO,redbo", "Comma-separated allowlist of peer node IDs")
	walPath := fs.String("wal", "/tmp/c2q55-bench-wal.log", "WAL path")
	slabPath := fs.String("slab", "/tmp/c2q55-bench-data.slab", "Slab arena path")
	offsetDir := fs.String("offset-dir", "/tmp/c2q55-bench-offsets", "Offset dir")
	metricsOut := fs.String("metrics-out", "/tmp/c2q55-server-metrics.jsonl", "JSONL output path for 1s metrics")
	_ = fs.Parse(args)

	opts := c2q55.DefaultOptions()
	opts.WALPath = *walPath
	opts.SlabPath = *slabPath
	opts.OffsetDir = *offsetDir
	opts.SlabSizeBytes = 512 * 1024 * 1024 // 512 Mo pour les paliers lourds

	eng, err := c2q55.NewEngine(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start engine: %v\n", err)
		os.Exit(1)
	}
	defer eng.Close()

	table := c2q55.NewCompactTable(opts.NumShards, eng.Slab())

	pool, err := c2q55.LoadClusterCAFromPEM(*caCertPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load CA pool: %v\n", err)
		os.Exit(1)
	}
	cert, err := c2q55.LoadNodeCertFromPEM(*nodeCertPath, *nodeKeyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load node cert: %v\n", err)
		os.Exit(1)
	}

	ca := &c2q55.ClusterCA{Pool: pool}
	allowList := strings.Split(*allowNodes, ",")
	tlsConf := ca.ServerTLSConfig(cert, allowList)

	srv, err := c2q55.ListenQUICTransportWithTLS(*addr, tlsConf, eng, table)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start QUIC server: %v\n", err)
		os.Exit(1)
	}
	defer srv.Close()

	fmt.Printf("[SERVER PROBE] Listening on %s (mTLS Strict TLS 1.3)\n", srv.Addr())
	fmt.Printf("[SERVER PROBE] Logging stratified metrics every 1s to %s\n", *metricsOut)

	outFile, err := os.OpenFile(*metricsOut, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open metrics file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	stopCh := make(chan struct{})
	go probeServerLoop(srv, eng, outFile, stopCh)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("\n[SERVER PROBE] Stopping server...")
	close(stopCh)
}

func probeServerLoop(srv *c2q55.QUICTransportServer, eng *c2q55.Engine, outFile *os.File, stopCh chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastRx uint64
	var lastMemStats runtime.MemStats
	var lastCPUUser, lastCPUSys, lastCPUSoftIRQ, lastCPUTotal uint64

	readCPUStats := func() (user, sys, softirq, total uint64) {
		data, err := os.ReadFile("/proc/stat")
		if err != nil {
			return
		}
		var cpuStr string
		var u, n, s, idle, iow, irq, sirq uint64
		_, _ = fmt.Sscanf(string(data), "%s %d %d %d %d %d %d %d", &cpuStr, &u, &n, &s, &idle, &iow, &irq, &sirq)
		user = u + n
		sys = s
		softirq = sirq
		total = u + n + s + idle + iow + irq + sirq
		return
	}

	lastCPUUser, lastCPUSys, lastCPUSoftIRQ, lastCPUTotal = readCPUStats()
	runtime.ReadMemStats(&lastMemStats)

	for {
		select {
		case <-stopCh:
			return
		case t := <-ticker.C:
			curRx := srv.Received()
			diffRx := curRx - lastRx
			lastRx = curRx

			// CPU Diff
			u, s, sirq, tot := readCPUStats()
			diffTot := float64(tot - lastCPUTotal)
			if diffTot <= 0 {
				diffTot = 1
			}
			cpuUserPct := float64(u-lastCPUUser) / diffTot * 100.0
			cpuSysPct := float64(s-lastCPUSys) / diffTot * 100.0
			cpuSoftIRQPct := float64(sirq-lastCPUSoftIRQ) / diffTot * 100.0

			lastCPUUser, lastCPUSys, lastCPUSoftIRQ, lastCPUTotal = u, s, sirq, tot

			// GC Stats
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			numGCDiff := m.NumGC - lastMemStats.NumGC
			var gcPauseMax, gcPauseSum uint64
			if numGCDiff > 0 {
				for i := uint32(0); i < numGCDiff && i < 256; i++ {
					idx := (m.NumGC - 1 - i) % 256
					p := m.PauseNs[idx]
					if p > gcPauseMax {
						gcPauseMax = p
					}
					gcPauseSum += p
				}
			}
			gcPauseAvg := uint64(0)
			if numGCDiff > 0 {
				gcPauseAvg = gcPauseSum / uint64(numGCDiff)
			}
			lastMemStats = m

			// Drops
			var totalDrops uint64
			for _, shard := range eng.Shards() {
				totalDrops += shard.Drops()
			}

			snap := ServerMetricsSnapshot{
				TimestampSec:    t.Unix(),
				MsgRxTotal:      curRx,
				MsgRxPerSec:     diffRx,
				BytesRxPerSecMB: float64(diffRx*4096) / (1024 * 1024),
				CPUUserPct:      cpuUserPct,
				CPUSysPct:       cpuSysPct,
				CPUSoftIRQPct:   cpuSoftIRQPct,
				GCPauseNsMax:    gcPauseMax,
				GCPauseNsAvg:    gcPauseAvg,
				AllocHeapMB:     float64(m.Alloc) / (1024 * 1024),
				NumGC:           numGCDiff,
				EngineDrops:     totalDrops,
			}

			data, _ := json.Marshal(snap)
			_, _ = outFile.WriteString(string(data) + "\n")
			_ = outFile.Sync()

			fmt.Printf("[PROBE 1s] Rx: %6d msg/s (%.1f MB/s) | CPU: usr=%.1f%% sys=%.1f%% sirq=%.1f%% | GC: %d (max=%.1fµs) | Drops: %d\n",
				diffRx, snap.BytesRxPerSecMB, cpuUserPct, cpuSysPct, cpuSoftIRQPct, numGCDiff, float64(gcPauseMax)/1000.0, totalDrops)
		}
	}
}

func runClientLoad(args []string) {
	fs := flag.NewFlagSet("client", flag.ExitOnError)
	target := fs.String("target", "213.32.71.129:8443", "Target server address")
	serverName := fs.String("server-name", "redhost", "Expected server SAN")
	caCertPath := fs.String("ca-cert", "", "CA cert path")
	nodeCertPath := fs.String("node-cert", "", "Node cert path")
	nodeKeyPath := fs.String("node-key", "", "Node key path")
	reportOut := fs.String("report-out", "/tmp/c2q55-client-report.json", "JSON output report path")
	_ = fs.Parse(args)

	pool, err := c2q55.LoadClusterCAFromPEM(*caCertPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load CA: %v\n", err)
		os.Exit(1)
	}
	cert, err := c2q55.LoadNodeCertFromPEM(*nodeCertPath, *nodeKeyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load node cert: %v\n", err)
		os.Exit(1)
	}

	tlsConf := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            pool,
		ServerName:         *serverName,
		InsecureSkipVerify: false,
		NextProtos:         []string{c2q55.QUICALPN},
		MinVersion:         tls.VersionTLS13,
	}

	stages := []struct {
		Name       string
		Duration   time.Duration
		TargetRate int // 0 = illimité
		Workers    int
		Mix1MBPct  int
	}{
		{Name: "Palier 1 : Régime Nominal (1 000 msg/s, 4 Ko)", Duration: 10 * time.Second, TargetRate: 1000, Workers: 8, Mix1MBPct: 0},
		{Name: "Palier 2 : Charge Moyenne (5 000 msg/s, 4 Ko)", Duration: 10 * time.Second, TargetRate: 5000, Workers: 16, Mix1MBPct: 0},
		{Name: "Palier 3 : Haute Charge (10 000 msg/s, 4 Ko)", Duration: 10 * time.Second, TargetRate: 10000, Workers: 32, Mix1MBPct: 0},
		{Name: "Palier 4 : Saturation Extrême (Illimité, 4 Ko)", Duration: 10 * time.Second, TargetRate: 0, Workers: 64, Mix1MBPct: 0},
		{Name: "Palier 5 : Charges Mixtes (80% 4 Ko + 20% 1 Mo)", Duration: 10 * time.Second, TargetRate: 2000, Workers: 16, Mix1MBPct: 20},
	}

	fmt.Println("================================================================================")
	fmt.Println("   BANC DE CHARGE STRATIFIÉ WAN C2Q55 (mTLS 1.3 / QUIC / PALIER 10s)")
	fmt.Printf("   Cible : %s (Serveur : %s) | Nœud Client : %s\n", *target, *serverName, "redbo")
	fmt.Println("================================================================================")

	var results []StageResult

	body4K := make([]byte, 4096)
	for i := range body4K {
		body4K[i] = byte(i % 256)
	}

	body1M := make([]byte, 1024*1024)
	for i := range body1M {
		body1M[i] = byte(i % 256)
	}

	for stageIdx, st := range stages {
		fmt.Printf("\n>>> DÉMARRAGE DU %s (Durée: %v, Workers: %d)\n", st.Name, st.Duration, st.Workers)

		clients := make([]*c2q55.QUICTransportClient, st.Workers)
		for w := 0; w < st.Workers; w++ {
			c, dialErr := c2q55.DialQUICTransportWithTLS(*target, tlsConf)
			if dialErr != nil {
				fmt.Fprintf(os.Stderr, "Worker %d failed to connect: %v\n", w, dialErr)
				os.Exit(1)
			}
			clients[w] = c
		}

		var sentCount, ackCount, errCount atomic.Uint64
		var totalBytes atomic.Uint64
		latencies := make([][]float64, st.Workers)
		var wg sync.WaitGroup

		startStage := time.Now()
		deadline := startStage.Add(st.Duration)

		intervalNs := int64(0)
		if st.TargetRate > 0 {
			intervalNs = int64(time.Second) * int64(st.Workers) / int64(st.TargetRate)
		}

		for w := 0; w < st.Workers; w++ {
			wg.Add(1)
			go func(workerID int, cl *c2q55.QUICTransportClient) {
				defer wg.Done()
				workerLats := make([]float64, 0, 50000)

				nextTick := time.Now().UnixNano()
				msgSeq := uint64(workerID * 1000000)

				for time.Now().Before(deadline) {
					if intervalNs > 0 {
						now := time.Now().UnixNano()
						if now < nextTick {
							time.Sleep(time.Duration(nextTick - now))
						}
						nextTick += intervalNs
					}

					var payload []byte
					if st.Mix1MBPct > 0 && rand.Intn(100) < st.Mix1MBPct {
						payload = body1M
					} else {
						payload = body4K
					}

					msgSeq++
					t0 := time.Now()
					sentCount.Add(1)

					sendErr := cl.SendPublish(uint16(workerID%16), msgSeq, uint64(t0.UnixNano()), payload, true)
					latMs := float64(time.Since(t0).Microseconds()) / 1000.0

					if sendErr != nil {
						errCount.Add(1)
					} else {
						ackCount.Add(1)
						totalBytes.Add(uint64(len(payload) + 32))
						workerLats = append(workerLats, latMs)
					}
				}
				latencies[workerID] = workerLats
			}(w, clients[w])
		}

		wg.Wait()
		elapsed := time.Since(startStage)

		for _, cl := range clients {
			_ = cl.Close()
		}

		allLats := make([]float64, 0, ackCount.Load())
		for _, l := range latencies {
			allLats = append(allLats, l...)
		}
		sort.Float64s(allLats)

		p50, p90, p99, p999, maxLat := 0.0, 0.0, 0.0, 0.0, 0.0
		nLats := len(allLats)
		if nLats > 0 {
			p50 = allLats[int(float64(nLats)*0.50)]
			p90 = allLats[int(float64(nLats)*0.90)]
			p99 = allLats[int(float64(nLats)*0.99)]
			p999 = allLats[int(float64(nLats)*0.999)]
			maxLat = allLats[nLats-1]
		}

		actualRate := float64(ackCount.Load()) / elapsed.Seconds()
		actualMB := float64(totalBytes.Load()) / (1024 * 1024) / elapsed.Seconds()

		res := StageResult{
			StageName:       st.Name,
			TargetRateMsgS:  st.TargetRate,
			ActualRateMsgS:  actualRate,
			ActualThroughMB: actualMB,
			TotalSent:       sentCount.Load(),
			TotalAcked:      ackCount.Load(),
			TotalErrors:     errCount.Load(),
			LatP50Ms:        p50,
			LatP90Ms:        p90,
			LatP99Ms:        p99,
			LatP999Ms:       p999,
			LatMaxMs:        maxLat,
		}
		results = append(results, res)

		fmt.Printf("   --- BILAN %s ---\n", st.Name)
		fmt.Printf("   Débit Réel    : %.2f msg/s (%.2f MB/s)\n", actualRate, actualMB)
		fmt.Printf("   Messages      : %d envoyés | %d acquittés | %d erreurs\n", sentCount.Load(), ackCount.Load(), errCount.Load())
		fmt.Printf("   Latences RTT  : p50=%.2fms | p90=%.2fms | p99=%.2fms | p99.9=%.2fms | max=%.2fms\n", p50, p90, p99, p999, maxLat)

		if stageIdx < len(stages)-1 {
			time.Sleep(2 * time.Second)
		}
	}

	reportBytes, _ := json.MarshalIndent(results, "", "  ")
	_ = os.WriteFile(*reportOut, reportBytes, 0644)
	fmt.Printf("\n[CLIENT LOAD] Rapport complet sauvegardé dans %s\n", *reportOut)
}
