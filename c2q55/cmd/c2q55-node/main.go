// SPDX-License-Identifier: Apache-2.0 OR MIT

package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/hazyhaar/c2pkg/c2q55"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	mode := os.Args[1]
	switch mode {
	case "pki-init":
		runPKIInit(os.Args[2:])
	case "listen":
		runServer(os.Args[2:])
	case "publish":
		runPublisher(os.Args[2:])
	case "consume":
		runConsumer(os.Args[2:])
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: c2q55-node <pki-init|listen|publish|consume> [options]")
	fmt.Println("  pki-init --out-dir /path/to/pki")
	fmt.Println("  listen   --addr 213.32.71.129:8443 [--ca-cert ca.crt --node-cert node.crt --node-key node.key --allow REDBO]")
	fmt.Println("  publish  --target 213.32.71.129:8443 --server-name redhost [--ca-cert ca.crt --node-cert node.crt --node-key node.key] --msg 'hello'")
	fmt.Println("  consume  --group workers --id worker-1 [--wal /path --slab /path --offset-dir /path]")
}

func runPKIInit(args []string) {
	fs := flag.NewFlagSet("pki-init", flag.ExitOnError)
	outDir := fs.String("out-dir", "/tmp/c2q55-pki", "Output directory for cluster PKI")
	_ = fs.Parse(args)

	if err := os.MkdirAll(*outDir, 0700); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create PKI directory: %v\n", err)
		os.Exit(1)
	}

	ca, err := c2q55.NewClusterCA()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Cluster CA: %v\n", err)
		os.Exit(1)
	}

	caCertPath := filepath.Join(*outDir, "ca.crt")
	caKeyPath := filepath.Join(*outDir, "ca.key")
	if err := c2q55.SaveCertAndKeyPEM(ca.CertDER, ca.Key, caCertPath, caKeyPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save CA PEM: %v\n", err)
		os.Exit(1)
	}

	// 1. Certificat pour redhost (213.32.71.129)
	redhostCert, err := ca.IssueNodeCertificate("redhost", []net.IP{net.ParseIP("213.32.71.129"), net.ParseIP("127.0.0.1")}, []string{"redhost", "localhost"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to issue redhost cert: %v\n", err)
		os.Exit(1)
	}
	if err := c2q55.SaveCertAndKeyPEM(redhostCert.Certificate[0], redhostCert.PrivateKey.(*ecdsa.PrivateKey), filepath.Join(*outDir, "redhost.crt"), filepath.Join(*outDir, "redhost.key")); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save redhost cert PEM: %v\n", err)
		os.Exit(1)
	}

	// 2. Certificat pour REDBO (Back-office)
	redboCert, err := ca.IssueNodeCertificate("REDBO", []net.IP{net.ParseIP("127.0.0.1")}, []string{"REDBO", "localhost"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to issue REDBO cert: %v\n", err)
		os.Exit(1)
	}
	if err := c2q55.SaveCertAndKeyPEM(redboCert.Certificate[0], redboCert.PrivateKey.(*ecdsa.PrivateKey), filepath.Join(*outDir, "redbo.crt"), filepath.Join(*outDir, "redbo.key")); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save REDBO cert PEM: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[c2q55-node] Sovereign Cluster PKI generated successfully in %s\n", *outDir)
	fmt.Printf("  CA Certificate: %s\n", caCertPath)
	fmt.Printf("  Node Certs: redhost (213.32.71.129), REDBO\n")
}

func runServer(args []string) {
	fs := flag.NewFlagSet("listen", flag.ExitOnError)
	addr := fs.String("addr", "127.0.0.1:8443", "Bind address for QUIC listener")
	walPath := fs.String("wal", "", "Optional WAL file path")
	slabPath := fs.String("slab", "", "Optional Slab arena path")
	offsetDir := fs.String("offset-dir", "", "Optional Offset store directory")
	caCertPath := fs.String("ca-cert", "", "Path to root Cluster CA certificate PEM")
	nodeCertPath := fs.String("node-cert", "", "Path to node certificate PEM")
	nodeKeyPath := fs.String("node-key", "", "Path to node private key PEM")
	allowNodes := fs.String("allow", "", "Comma-separated allowlist of peer node IDs")
	_ = fs.Parse(args)

	opts := c2q55.DefaultOptions()
	opts.WALPath = *walPath
	opts.SlabPath = *slabPath
	opts.OffsetDir = *offsetDir

	eng, err := c2q55.NewEngine(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start engine: %v\n", err)
		os.Exit(1)
	}
	defer eng.Close()

	table := c2q55.NewCompactTable(opts.NumShards, eng.Slab())

	var srv *c2q55.QUICTransportServer

	if *caCertPath != "" && *nodeCertPath != "" && *nodeKeyPath != "" {
		pool, err := c2q55.LoadClusterCAFromPEM(*caCertPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load CA pool: %v\n", err)
			os.Exit(1)
		}
		cert, err := c2q55.LoadNodeCertFromPEM(*nodeCertPath, *nodeKeyPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load node cert/key: %v\n", err)
			os.Exit(1)
		}

		var allowList []string
		if *allowNodes != "" {
			allowList = strings.Split(*allowNodes, ",")
		}

		ca := &c2q55.ClusterCA{Pool: pool}
		tlsConf := ca.ServerTLSConfig(cert, allowList)
		srv, err = c2q55.ListenQUICTransportWithTLS(*addr, tlsConf, eng, table)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to start QUIC server with custom TLS: %v\n", err)
			os.Exit(1)
		}
	} else {
		srv, err = c2q55.ListenQUICTransport(*addr, eng, table)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to start QUIC server: %v\n", err)
			os.Exit(1)
		}
	}
	defer srv.Close()

	fmt.Printf("[c2q55-node] QUIC TLS 1.3 Sovereign Node listening on %s (mTLS Strict)\n", srv.Addr())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("\n[c2q55-node] Shutting down gracefully...")
}

func runPublisher(args []string) {
	fs := flag.NewFlagSet("publish", flag.ExitOnError)
	target := fs.String("target", "127.0.0.1:8443", "Target QUIC server address")
	serverName := fs.String("server-name", "node-server", "Expected server SAN/CN")
	caCertPath := fs.String("ca-cert", "", "Path to root Cluster CA certificate PEM")
	nodeCertPath := fs.String("node-cert", "", "Path to node certificate PEM")
	nodeKeyPath := fs.String("node-key", "", "Path to node private key PEM")
	topic := fs.Uint("topic", 1, "Topic ID")
	key := fs.Uint64("key", 1, "Partitioning Key")
	msg := fs.String("msg", "c2q55-telemetry-event", "Message body")
	msgFile := fs.String("msg-file", "", "Path to file containing message payload")
	count := fs.Int("count", 1, "Number of messages to publish")
	_ = fs.Parse(args)

	var payloadBytes []byte
	if *msgFile != "" {
		data, err := os.ReadFile(*msgFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read msg-file: %v\n", err)
			os.Exit(1)
		}
		payloadBytes = data
	} else {
		payloadBytes = []byte(*msg)
	}

	var client *c2q55.QUICTransportClient
	var err error

	if *caCertPath != "" && *nodeCertPath != "" && *nodeKeyPath != "" {
		pool, loadErr := c2q55.LoadClusterCAFromPEM(*caCertPath)
		if loadErr != nil {
			fmt.Fprintf(os.Stderr, "Failed to load CA pool: %v\n", loadErr)
			os.Exit(1)
		}
		cert, loadErr := c2q55.LoadNodeCertFromPEM(*nodeCertPath, *nodeKeyPath)
		if loadErr != nil {
			fmt.Fprintf(os.Stderr, "Failed to load node cert/key: %v\n", loadErr)
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
		client, err = c2q55.DialQUICTransportWithTLS(*target, tlsConf)
	} else {
		client, err = c2q55.DialQUICTransport(*target)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to QUIC server: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	start := time.Now()
	for i := 0; i < *count; i++ {
		k := *key + uint64(i)
		err := client.SendPublish(uint16(*topic), k, uint64(time.Now().UnixNano()), payloadBytes, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Publish failed at %d: %v\n", i, err)
			os.Exit(1)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("[c2q55-node] Published %d messages in %v (%.2f msg/s)\n", *count, elapsed, float64(*count)/elapsed.Seconds())
}

func runConsumer(args []string) {
	fs := flag.NewFlagSet("consume", flag.ExitOnError)
	walPath := fs.String("wal", "", "Optional WAL file path")
	slabPath := fs.String("slab", "", "Optional Slab arena path")
	offsetDir := fs.String("offset-dir", "", "Optional Offset store directory")
	groupName := fs.String("group", "default-group", "Consumer group name")
	consumerID := fs.String("id", "worker-1", "Consumer ID")
	_ = fs.Parse(args)

	opts := c2q55.DefaultOptions()
	opts.WALPath = *walPath
	opts.SlabPath = *slabPath
	opts.OffsetDir = *offsetDir

	eng, err := c2q55.NewEngine(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open engine: %v\n", err)
		os.Exit(1)
	}
	defer eng.Close()

	group := eng.GetOrCreateConsumerGroup(*groupName)
	assigned := group.RegisterConsumer(*consumerID)
	fmt.Printf("[c2q55-node] Consumer %q registered in group %q (assigned partitions: %v)\n", *consumerID, *groupName, assigned)

	var msg c2q55.Message
	consumedCount := 0

	for {
		if ok, _ := eng.ConsumeGroup(*groupName, *consumerID, &msg); ok {
			consumedCount++
			_ = group.CommitOffset(msg.Partition, msg.Offset)
			fmt.Printf("[msg #%d] partition=%d offset=%d key=%d len=%d body=%q\n",
				consumedCount, msg.Partition, msg.Offset, msg.Key, len(msg.Body), string(msg.Body))
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
}
