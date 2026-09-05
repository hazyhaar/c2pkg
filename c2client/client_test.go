// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2client

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/hazyhaar/c2pkg/c2uuidv7"
	"github.com/quic-go/quic-go"
)

func TestRouteShardDistribution(t *testing.T) {
	counts := make(map[uint16]int)
	total := 50000

	for i := 0; i < total; i++ {
		key := []byte(fmt.Sprintf("order:crypto:usdt:pair:%d", i))
		shard := RouteShard(key)
		if shard >= NumShards {
			t.Fatalf("shard hors limites: %d >= %d", shard, NumShards)
		}
		counts[shard]++

		// Test du déterminisme
		if replay := RouteShard(key); replay != shard {
			t.Fatalf("non-déterminisme sur clé %s: initial %d != replay %d", key, shard, replay)
		}
	}

	// Couverture attendue : les 1024 shards doivent être touchés avec 50 000 clés
	if len(counts) < 950 {
		t.Fatalf("dispersion insuffisante: seulement %d/1024 shards couverts", len(counts))
	}
}

func TestFrameRoundtrip(t *testing.T) {
	var buf bytes.Buffer

	hdr := Header{
		OpCode: OpPut,
		Flags:  0x01,
		Tenant: 0x42,
		Shard:  512,
		Extra:  999,
	}
	key := []byte("btc-usd-orderbook")
	val := []byte(`{"bid": 64230.50, "ask": 64231.00, "ts": 1725530000}`)

	if err := WriteFrame(&buf, hdr, key, val); err != nil {
		t.Fatalf("WriteFrame failed: %v", err)
	}

	decodedHdr, err := DecodeHeader(&buf)
	if err != nil {
		t.Fatalf("DecodeHeader failed: %v", err)
	}

	if decodedHdr.Magic != MagicProto {
		t.Fatalf("magic mismatch: got 0x%04X, want 0x%04X", decodedHdr.Magic, MagicProto)
	}
	if decodedHdr.OpCode != hdr.OpCode {
		t.Fatalf("opCode mismatch: got %d, want %d", decodedHdr.OpCode, hdr.OpCode)
	}
	if decodedHdr.Tenant != hdr.Tenant {
		t.Fatalf("tenant mismatch: got %d, want %d", decodedHdr.Tenant, hdr.Tenant)
	}
	if decodedHdr.Shard != hdr.Shard {
		t.Fatalf("shard mismatch: got %d, want %d", decodedHdr.Shard, hdr.Shard)
	}

	decodedKey, decodedVal, err := ReadFramePayload(&buf, decodedHdr)
	if err != nil {
		t.Fatalf("ReadFramePayload failed: %v", err)
	}

	if !bytes.Equal(decodedKey, key) {
		t.Fatalf("key mismatch: got %q, want %q", decodedKey, key)
	}
	if !bytes.Equal(decodedVal, val) {
		t.Fatalf("val mismatch: got %q, want %q", decodedVal, val)
	}
}

func TestTenantTLSHandshakeRejection(t *testing.T) {
	masterA := [32]byte{1, 2, 3}
	masterB := [32]byte{4, 5, 6}
	epoch := c2uuidv7.New()

	serverTLS, err := GenerateTenantTLSConfig(masterA, 1, epoch, true)
	if err != nil {
		t.Fatalf("GenerateTenantTLSConfig server: %v", err)
	}

	// Client avec master key différente
	clientTLSBadMaster, err := GenerateTenantTLSConfig(masterB, 1, epoch, false)
	if err != nil {
		t.Fatalf("GenerateTenantTLSConfig client bad master: %v", err)
	}

	// Client avec tenant différent
	clientTLSBadTenant, err := GenerateTenantTLSConfig(masterA, 2, epoch, false)
	if err != nil {
		t.Fatalf("GenerateTenantTLSConfig client bad tenant: %v", err)
	}

	// Client avec époque différente
	clientTLSBadEpoch, err := GenerateTenantTLSConfig(masterA, 1, c2uuidv7.New(), false)
	if err != nil {
		t.Fatalf("GenerateTenantTLSConfig client bad epoch: %v", err)
	}

	// Client légitime
	clientTLSValid, err := GenerateTenantTLSConfig(masterA, 1, epoch, false)
	if err != nil {
		t.Fatalf("GenerateTenantTLSConfig client valid: %v", err)
	}

	quicConf := &quic.Config{Allow0RTT: true}

	// Serveur d'écoute local
	listener, err := quic.ListenAddr("127.0.0.1:0", serverTLS, quicConf)
	if err != nil {
		t.Fatalf("quic.ListenAddr: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Gestion des connexions côté serveur
	go func() {
		for {
			conn, err := listener.Accept(context.Background())
			if err != nil {
				return
			}
			go func(c *quic.Conn) {
				<-c.Context().Done()
			}(conn)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 1. Rejet Bad Master
	if _, err := quic.DialAddrEarly(ctx, addr, clientTLSBadMaster, quicConf); err == nil {
		t.Fatal("connexion réussie attendue rejetée (bad master)")
	}

	// 2. Rejet Bad Tenant
	if _, err := quic.DialAddrEarly(ctx, addr, clientTLSBadTenant, quicConf); err == nil {
		t.Fatal("connexion réussie attendue rejetée (bad tenant)")
	}

	// 3. Rejet Bad Epoch
	if _, err := quic.DialAddrEarly(ctx, addr, clientTLSBadEpoch, quicConf); err == nil {
		t.Fatal("connexion réussie attendue rejetée (bad epoch)")
	}

	// 4. Succès Pair Valide
	connValid, err := quic.DialAddrEarly(ctx, addr, clientTLSValid, quicConf)
	if err != nil {
		t.Fatalf("échec poignée de main légitime: %v", err)
	}
	_ = connValid.CloseWithError(0, "ok")
}

func TestClientServerQUICRoundtrip(t *testing.T) {
	master := [32]byte{0xDE, 0xAD, 0xBE, 0xEF}
	tenant := uint16(0x07)
	epoch := c2uuidv7.New()

	serverTLS, err := GenerateTenantTLSConfig(master, tenant, epoch, true)
	if err != nil {
		t.Fatalf("GenerateTenantTLSConfig server: %v", err)
	}

	quicConf := &quic.Config{
		Allow0RTT:             true,
		MaxIncomingStreams:    1000,
		MaxIncomingUniStreams: 1000,
	}

	listener, err := quic.ListenAddr("127.0.0.1:0", serverTLS, quicConf)
	if err != nil {
		t.Fatalf("quic.ListenAddr: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Mémoire mock côté serveur
	store := make(map[string][]byte)
	var storeMu sync.RWMutex

	ctxSrv, cancelSrv := context.WithCancel(context.Background())
	defer cancelSrv()

	go func() {
		for {
			conn, err := listener.Accept(ctxSrv)
			if err != nil {
				return
			}
			go func(c *quic.Conn) {
				for {
					stream, err := c.AcceptStream(ctxSrv)
					if err != nil {
						return
					}
					go func(st *quic.Stream) {
						defer st.Close()
						hdr, err := DecodeHeader(st)
						if err != nil {
							return
						}
						key, val, err := ReadFramePayload(st, hdr)
						if err != nil {
							return
						}

						switch hdr.OpCode {
						case OpPut:
							storeMu.Lock()
							store[string(key)] = append([]byte(nil), val...)
							storeMu.Unlock()
							_ = WriteFrame(st, Header{OpCode: OpOK, Tenant: tenant, Shard: hdr.Shard}, nil, nil)

						case OpGet:
							storeMu.RLock()
							stored, ok := store[string(key)]
							storeMu.RUnlock()
							if ok {
								_ = WriteFrame(st, Header{OpCode: OpOK, Tenant: tenant, Shard: hdr.Shard}, nil, stored)
							} else {
								_ = WriteFrame(st, Header{OpCode: OpErr, Tenant: tenant, Shard: hdr.Shard}, nil, []byte("key not found"))
							}

						case OpQuery:
							storeMu.RLock()
							for k, v := range store {
								if bytes.HasPrefix([]byte(k), key) {
									_ = WriteFrame(st, Header{OpCode: OpEntry, Tenant: tenant, Shard: hdr.Shard}, []byte(k), v)
								}
							}
							storeMu.RUnlock()
							_ = WriteFrame(st, Header{OpCode: OpEnd, Tenant: tenant, Shard: hdr.Shard}, nil, nil)
						}
					}(stream)
				}
			}(conn)
		}
	}()

	client, err := NewClient(master, tenant, epoch, []string{addr})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 1. Put
	k1 := []byte("pool:eth-usdt:rate")
	v1 := []byte("2640.50")
	if err := client.Put(ctx, k1, v1); err != nil {
		t.Fatalf("Put k1 failed: %v", err)
	}

	// 2. Get
	retrieved, err := client.Get(ctx, k1)
	if err != nil {
		t.Fatalf("Get k1 failed: %v", err)
	}
	if !bytes.Equal(retrieved, v1) {
		t.Fatalf("valeur inattendue: got %s, want %s", string(retrieved), string(v1))
	}

	// 3. Put second key
	k2 := []byte("pool:sol-usdt:rate")
	v2 := []byte("145.20")
	if err := client.Put(ctx, k2, v2); err != nil {
		t.Fatalf("Put k2 failed: %v", err)
	}

	// 4. Query Prefix
	entries, err := client.Query([]byte("pool:")).Collect(ctx)
	if err != nil {
		t.Fatalf("Query prefix failed: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("nombre d'entrées inattendu: got %d, want 2", len(entries))
	}
}
