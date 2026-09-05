// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2client

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"sync"

	"github.com/hazyhaar/c2pkg/blake3archtsim"
	"github.com/hazyhaar/c2pkg/c2uuidv7"
	"github.com/quic-go/quic-go"
)

// NumShards définit le nombre canonique de fragments virtuels (1024)
// pour le partitionnement déterministe haute dispersion.
const NumShards = 1024

// RouteShard calcule de manière déterministe le numéro de fragment (0 à 1023)
// associé à une clé par hachage SIMD Blake3 ultra-rapide (10 bits de poids fort).
func RouteShard(key []byte) uint16 {
	sum := blake3archtsim.Sum256(key)
	return binary.BigEndian.Uint16(sum[:2]) >> 6
}

// Client fournit un client distribué Agent-First avec routage déterministe,
// pool de sessions QUIC multi-nœuds et reconnexion 0-RTT sans état.
type Client struct {
	master    [32]byte
	tenant    uint16
	epoch     c2uuidv7.UUID
	tlsConf   *tls.Config
	quicConf  *quic.Config
	nodeAddrs [NumShards]string

	connsMu sync.RWMutex
	conns   map[string]*quic.Conn
	closed  bool
}

// NewClient instancie un client c2client configuré avec l'identité mTLS du tenant.
// nodes représente la liste ordonnée des adresses UDP/QUIC (ex: "10.0.0.1:8155") du cluster.
func NewClient(master [32]byte, tenant uint16, epoch c2uuidv7.UUID, nodes []string) (*Client, error) {
	if len(nodes) == 0 {
		return nil, errors.New("c2client: at least one node address required")
	}

	tlsConf, err := GenerateTenantTLSConfig(master, tenant, epoch, false)
	if err != nil {
		return nil, fmt.Errorf("c2client: generate tls config: %w", err)
	}

	quicConf := &quic.Config{
		Allow0RTT:             true,
		EnableDatagrams:       true,
		MaxIdleTimeout:        0,
		KeepAlivePeriod:       0,
		MaxIncomingStreams:    10000,
		MaxIncomingUniStreams: 10000,
	}

	c := &Client{
		master:   master,
		tenant:   tenant,
		epoch:    epoch,
		tlsConf:  tlsConf,
		quicConf: quicConf,
		conns:    make(map[string]*quic.Conn),
	}

	// Répartition déterministe des 1024 shards sur les nœuds du cluster
	for i := 0; i < NumShards; i++ {
		c.nodeAddrs[i] = nodes[i%len(nodes)]
	}

	return c, nil
}

// Route détermine le Shard et l'adresse réseau du nœud cible à partir de la clé.
func (c *Client) Route(key []byte) (shard uint16, addr string) {
	shard = RouteShard(key)
	addr = c.nodeAddrs[shard]
	return shard, addr
}

// getConn récupère une connexion existante du pool ou établit une nouvelle session QUIC en 0-RTT.
func (c *Client) getConn(ctx context.Context, addr string) (*quic.Conn, error) {
	c.connsMu.RLock()
	if c.closed {
		c.connsMu.RUnlock()
		return nil, errors.New("c2client: client closed")
	}
	conn, ok := c.conns[addr]
	if ok && conn != nil {
		select {
		case <-conn.Context().Done():
			// Connexion terminée, nécessité de reconnexion
		default:
			c.connsMu.RUnlock()
			return conn, nil
		}
	}
	c.connsMu.RUnlock()

	c.connsMu.Lock()
	defer c.connsMu.Unlock()

	if c.closed {
		return nil, errors.New("c2client: client closed")
	}

	if conn, ok := c.conns[addr]; ok && conn != nil {
		select {
		case <-conn.Context().Done():
		default:
			return conn, nil
		}
	}

	// Établissement de session rapide 0-RTT sans état
	newConn, err := quic.DialAddrEarly(ctx, addr, c.tlsConf, c.quicConf)
	if err != nil {
		return nil, fmt.Errorf("c2client: dial 0-rtt %s: %w", addr, err)
	}

	c.conns[addr] = newConn
	return newConn, nil
}

// Put exécute une mutation transactionnelle sur le nœud hébergeant le shard ciblé.
func (c *Client) Put(ctx context.Context, key, val []byte) error {
	shard, addr := c.Route(key)
	conn, err := c.getConn(ctx, addr)
	if err != nil {
		return err
	}

	stream, err := conn.OpenStreamSync(ctx)
	if err != nil {
		return fmt.Errorf("c2client: open stream: %w", err)
	}
	defer stream.Close()

	reqHdr := Header{
		OpCode: OpPut,
		Tenant: c.tenant,
		Shard:  shard,
	}
	if err := WriteFrame(stream, reqHdr, key, val); err != nil {
		return fmt.Errorf("c2client: send put frame: %w", err)
	}

	respHdr, err := DecodeHeader(stream)
	if err != nil {
		return fmt.Errorf("c2client: decode response: %w", err)
	}

	if respHdr.OpCode == OpOK {
		return nil
	}

	_, errMsg, _ := ReadFramePayload(stream, respHdr)
	return fmt.Errorf("c2client: put error: %s", string(errMsg))
}

// Get extrait une valeur par clé depuis le nœud hébergeant le shard ciblé.
func (c *Client) Get(ctx context.Context, key []byte) ([]byte, error) {
	shard, addr := c.Route(key)
	conn, err := c.getConn(ctx, addr)
	if err != nil {
		return nil, err
	}

	stream, err := conn.OpenStreamSync(ctx)
	if err != nil {
		return nil, fmt.Errorf("c2client: open stream: %w", err)
	}
	defer stream.Close()

	reqHdr := Header{
		OpCode: OpGet,
		Tenant: c.tenant,
		Shard:  shard,
	}
	if err := WriteFrame(stream, reqHdr, key, nil); err != nil {
		return nil, fmt.Errorf("c2client: send get frame: %w", err)
	}

	respHdr, err := DecodeHeader(stream)
	if err != nil {
		return nil, fmt.Errorf("c2client: decode response: %w", err)
	}

	_, val, err := ReadFramePayload(stream, respHdr)
	if err != nil {
		return nil, fmt.Errorf("c2client: read response payload: %w", err)
	}

	if respHdr.OpCode == OpOK {
		return val, nil
	}

	return nil, fmt.Errorf("c2client: get error: %s", string(val))
}

// Query initialise une requête déclarative distribuée sur l'ensemble du cluster.
func (c *Client) Query(prefix []byte) *RemoteQuery {
	return &RemoteQuery{
		client: c,
		shard:  ShardBroadcast,
		prefix: prefix,
		limit:  0,
	}
}

// QueryShard initialise une requête déclarative ciblée sur un shard précis.
func (c *Client) QueryShard(shard uint16, prefix []byte) *RemoteQuery {
	return &RemoteQuery{
		client: c,
		shard:  shard,
		prefix: prefix,
		limit:  0,
	}
}

// RemoteQuery décrit une requête déclarative en cours de construction.
type RemoteQuery struct {
	client *Client
	shard  uint16
	prefix []byte
	limit  int
}

// Limit borne le nombre maximal d'enregistrements retournés.
func (rq *RemoteQuery) Limit(limit int) *RemoteQuery {
	rq.limit = limit
	return rq
}

// Collect exécute la requête et matérialise l'ensemble des paires clé-valeur correspondantes.
func (rq *RemoteQuery) Collect(ctx context.Context) ([]Entry, error) {
	var entries []Entry
	err := rq.Scan(ctx, func(k, v []byte) bool {
		entries = append(entries, Entry{
			Key: append([]byte(nil), k...),
			Val: append([]byte(nil), v...),
		})
		if rq.limit > 0 && len(entries) >= rq.limit {
			return false
		}
		return true
	})
	return entries, err
}

// Scan exécute la requête en streaming direct sur les streams QUIC du cluster.
func (rq *RemoteQuery) Scan(ctx context.Context, fn func(k, v []byte) bool) error {
	var targetNodes []string
	if rq.shard != ShardBroadcast {
		targetNodes = []string{rq.client.nodeAddrs[rq.shard]}
	} else {
		seen := make(map[string]struct{})
		for _, addr := range rq.client.nodeAddrs {
			if _, ok := seen[addr]; !ok {
				seen[addr] = struct{}{}
				targetNodes = append(targetNodes, addr)
			}
		}
	}

	for _, addr := range targetNodes {
		conn, err := rq.client.getConn(ctx, addr)
		if err != nil {
			return err
		}

		stream, err := conn.OpenStreamSync(ctx)
		if err != nil {
			return fmt.Errorf("c2client: open stream: %w", err)
		}

		reqHdr := Header{
			OpCode: OpQuery,
			Tenant: rq.client.tenant,
			Shard:  rq.shard,
			Extra:  uint64(rq.limit),
		}
		if err := WriteFrame(stream, reqHdr, rq.prefix, nil); err != nil {
			_ = stream.Close()
			return fmt.Errorf("c2client: send query frame: %w", err)
		}

		done := false
		for {
			respHdr, err := DecodeHeader(stream)
			if err != nil {
				_ = stream.Close()
				return fmt.Errorf("c2client: decode query entry header: %w", err)
			}

			if respHdr.OpCode == OpEnd {
				break
			}

			if respHdr.OpCode == OpErr {
				_, errMsg, _ := ReadFramePayload(stream, respHdr)
				_ = stream.Close()
				return fmt.Errorf("c2client: query error: %s", string(errMsg))
			}

			if respHdr.OpCode != OpEntry {
				_ = stream.Close()
				return fmt.Errorf("c2client: unexpected opcode: 0x%02X", respHdr.OpCode)
			}

			k, v, err := ReadFramePayload(stream, respHdr)
			if err != nil {
				_ = stream.Close()
				return fmt.Errorf("c2client: read query payload: %w", err)
			}

			if !fn(k, v) {
				done = true
				break
			}
		}
		_ = stream.Close()
		if done {
			break
		}
	}

	return nil
}

// Close termine proprement les connexions QUIC actives du pool.
func (c *Client) Close() error {
	c.connsMu.Lock()
	defer c.connsMu.Unlock()

	c.closed = true
	var firstErr error
	for _, conn := range c.conns {
		if conn != nil {
			if err := conn.CloseWithError(0, "client closed"); err != nil && firstErr == nil {
				firstErr = err
			}
		}
	}
	c.conns = nil
	return firstErr
}
