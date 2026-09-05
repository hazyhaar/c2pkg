// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2client

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/hazyhaar/c2pkg/blake3archtsim"
	"github.com/hazyhaar/c2pkg/c2uuidv7"
	"github.com/quic-go/quic-go"
	"golang.org/x/sync/singleflight"
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

	dialGroup singleflight.Group
	bufPool   sync.Pool
}

// NewClient instancie un client c2client configuré avec l'identité mTLS du tenant.
// NewClientWithSeed instancie un client c2client à partir d'une graine de tenant pré-dérivée,
// garantissant l'étanchéité cryptographique sans nécessiter l'accès à la clé maîtresse du cluster.
func NewClientWithSeed(tenantSeed [32]byte, tenant uint16, epoch c2uuidv7.UUID, nodes []string) (*Client, error) {
	if len(nodes) == 0 {
		return nil, errors.New("c2client: at least one node address required")
	}

	tlsConf, err := GenerateTenantTLSConfig(tenantSeed, tenant, false)
	if err != nil {
		return nil, fmt.Errorf("c2client: generate tls config: %w", err)
	}

	quicConf := &quic.Config{
		Allow0RTT:             true,
		EnableDatagrams:       true,
		MaxIdleTimeout:        30 * time.Second,
		KeepAlivePeriod:       10 * time.Second,
		MaxIncomingStreams:    10000,
		MaxIncomingUniStreams: 10000,
	}

	c := &Client{
		tenant:   tenant,
		epoch:    epoch,
		tlsConf:  tlsConf,
		quicConf: quicConf,
		conns:    make(map[string]*quic.Conn),
		bufPool: sync.Pool{
			New: func() any {
				b := make([]byte, 64*1024)
				return &b
			},
		},
	}

	// Répartition déterministe des 1024 shards sur les nœuds du cluster
	for i := 0; i < NumShards; i++ {
		c.nodeAddrs[i] = nodes[i%len(nodes)]
	}

	return c, nil
}

// NewClient instancie un client c2client configuré avec l'identité mTLS du tenant.
// nodes représente la liste ordonnée des adresses UDP/QUIC (ex: "10.0.0.1:8155") du cluster.
func NewClient(master [32]byte, tenant uint16, epoch c2uuidv7.UUID, nodes []string) (*Client, error) {
	tenantSeed := DeriveTenantMAC(master, tenant, epoch)
	c, err := NewClientWithSeed(tenantSeed, tenant, epoch, nodes)
	if err != nil {
		return nil, err
	}
	c.master = master
	return c, nil
}

// Route détermine le Shard et l'adresse réseau du nœud cible à partir de la clé.
func (c *Client) Route(key []byte) (shard uint16, addr string) {
	shard = RouteShard(key)
	addr = c.nodeAddrs[shard]
	return shard, addr
}

// getConn récupère une connexion existante du pool ou établit une nouvelle session QUIC en 0-RTT sans bloquer les autres cibles.
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
			// Nettoyage immédiat d'une connexion morte détectée opportunistement
			c.connsMu.RUnlock()
			c.connsMu.Lock()
			if c.conns[addr] == conn {
				delete(c.conns, addr)
			}
			c.connsMu.Unlock()
		default:
			c.connsMu.RUnlock()
			return conn, nil
		}
	} else {
		c.connsMu.RUnlock()
	}

	// Déduplication des tentatives de connexion concurrentes sur la même adresse
	res, err, _ := c.dialGroup.Do(addr, func() (interface{}, error) {
		// Double vérification après acquisition du ticket de dial
		c.connsMu.RLock()
		if c.closed {
			c.connsMu.RUnlock()
			return nil, errors.New("c2client: client closed")
		}
		if existing, ok := c.conns[addr]; ok && existing != nil {
			select {
			case <-existing.Context().Done():
			default:
				c.connsMu.RUnlock()
				return existing, nil
			}
		}
		c.connsMu.RUnlock()

		// Établissement de session rapide 0-RTT sans état (hors verrou global)
		newConn, dialErr := quic.DialAddrEarly(ctx, addr, c.tlsConf, c.quicConf)
		if dialErr != nil {
			return nil, fmt.Errorf("c2client: dial 0-rtt %s: %w", addr, dialErr)
		}

		c.connsMu.Lock()
		defer c.connsMu.Unlock()
		if c.closed {
			_ = newConn.CloseWithError(0, "client closed")
			return nil, errors.New("c2client: client closed")
		}
		
		if existing, ok := c.conns[addr]; ok && existing != nil {
			select {
			case <-existing.Context().Done():
			default:
				_ = newConn.CloseWithError(0, "replaced")
				return existing, nil
			}
		}

		c.conns[addr] = newConn
		return newConn, nil
	})

	if err != nil {
		return nil, err
	}
	return res.(*quic.Conn), nil
}

// Put exécute une mutation transactionnelle sur le nœud hébergeant le shard ciblé.
func (c *Client) Put(ctx context.Context, key, val []byte) error {
	shard, addr := c.Route(key)
	conn, err := c.getConn(ctx, addr)
	if err != nil {
		return err
	}

	// Barrière anti-rejeu RFC 9001 : interdiction du 0-RTT sur les mutations non-idempotentes
	select {
	case <-conn.HandshakeComplete():
	case <-ctx.Done():
		return ctx.Err()
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
	if shard >= NumShards && shard != ShardBroadcast {
		shard = shard % NumShards
	}
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
