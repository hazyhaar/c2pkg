package c2q55

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go"
)

const (
	QUICFrameMagic uint32 = 0x43325155 // "C2QU"
	QUICHeaderSize int    = 32         // 32 octets d'en-tête réseau binaire fixe
	QUICALPN              = "c2q55/1"  // ALPN TLS 1.3 dédié
	CmdPublish     uint8  = 1
	CmdAck         uint8  = 2
	CmdSync        uint8  = 3
	CmdFetch       uint8  = 4
	CmdHeartbeat   uint8  = 5
)

var (
	ErrInvalidFrameMagic = errors.New("c2q55/quic: invalid frame magic")
	ErrFrameCRCMismatch  = errors.New("c2q55/quic: frame CRC32-C mismatch")
)

// QUICHeader définit l'en-tête binaire réseau de 32 octets aligné sur stream QUIC.
type QUICHeader struct {
	Magic       uint32 // 0..3   - 0x43325155
	Cmd         uint8  // 4      - Commande
	Flags       uint8  // 5      - Options
	PartitionID uint16 // 6..7   - Identifiant de partition
	LSN         uint64 // 8..15  - Numéro de séquence log
	Offset      uint64 // 16..23 - Décalage de partition
	PayloadLen  uint32 // 24..27 - Taille des données
	CRC32C      uint32 // 28..31 - Somme Castagnoli matérielle sur la charge utile
}

// Encode sérialise l'en-tête en 32 octets.
func (h *QUICHeader) Encode(dst *[32]byte) {
	binary.LittleEndian.PutUint32(dst[0:4], h.Magic)
	dst[4] = h.Cmd
	dst[5] = h.Flags
	binary.LittleEndian.PutUint16(dst[6:8], h.PartitionID)
	binary.LittleEndian.PutUint64(dst[8:16], h.LSN)
	binary.LittleEndian.PutUint64(dst[16:24], h.Offset)
	binary.LittleEndian.PutUint32(dst[24:28], h.PayloadLen)
	binary.LittleEndian.PutUint32(dst[28:32], h.CRC32C)
}

// DecodeQUICHeader lit l'en-tête depuis 32 octets.
func DecodeQUICHeader(src []byte) (QUICHeader, error) {
	if len(src) < QUICHeaderSize {
		return QUICHeader{}, errors.New("c2q55/quic: header buffer too small")
	}

	magic := binary.LittleEndian.Uint32(src[0:4])
	if magic != QUICFrameMagic {
		return QUICHeader{}, ErrInvalidFrameMagic
	}

	return QUICHeader{
		Magic:       magic,
		Cmd:         src[4],
		Flags:       src[5],
		PartitionID: binary.LittleEndian.Uint16(src[6:8]),
		LSN:         binary.LittleEndian.Uint64(src[8:16]),
		Offset:      binary.LittleEndian.Uint64(src[16:24]),
		PayloadLen:  binary.LittleEndian.Uint32(src[24:28]),
		CRC32C:      binary.LittleEndian.Uint32(src[28:32]),
	}, nil
}

// QUICTransportServer gère la réception de flux réseau partitionnés via streams quic-go sous mTLS 1.3.
type QUICTransportServer struct {
	listener *quic.Listener
	engine   *Engine
	table    *CompactTable
	closed   atomic.Bool
	received atomic.Uint64
	bytesRx  atomic.Uint64
	wg       sync.WaitGroup
}

// ListenQUICTransportWithTLS démarre un serveur de transport réseau QUIC mTLS 1.3 strict.
func ListenQUICTransportWithTLS(bindAddr string, tlsConf *tls.Config, engine *Engine, table *CompactTable) (*QUICTransportServer, error) {
	quicConf := &quic.Config{
		MaxIdleTimeout:                 30 * time.Second,
		InitialStreamReceiveWindow:     16 * 1024 * 1024,
		InitialConnectionReceiveWindow: 64 * 1024 * 1024,
		MaxIncomingStreams:             10000,
	}

	listener, err := quic.ListenAddr(bindAddr, tlsConf, quicConf)
	if err != nil {
		return nil, fmt.Errorf("c2q55/quic: quic listen failed: %w", err)
	}

	s := &QUICTransportServer{
		listener: listener,
		engine:   engine,
		table:    table,
	}

	s.wg.Add(1)
	go s.acceptLoop()

	return s, nil
}

// ListenQUICTransport démarre un serveur configuré avec CA par défaut en mTLS strict.
func ListenQUICTransport(bindAddr string, engine *Engine, table *CompactTable) (*QUICTransportServer, error) {
	ca := GetDefaultTestCA()

	nodeCert, err := ca.IssueNodeCertificate("node-server", []net.IP{net.ParseIP("127.0.0.1")}, []string{"localhost", "node-server"})
	if err != nil {
		return nil, err
	}

	tlsConf := ca.ServerTLSConfig(nodeCert, nil) // RequireAndVerifyClientCert actif
	return ListenQUICTransportWithTLS(bindAddr, tlsConf, engine, table)
}

func (s *QUICTransportServer) acceptLoop() {
	defer s.wg.Done()

	for !s.closed.Load() {
		conn, err := s.listener.Accept(context.Background())
		if err != nil {
			if s.closed.Load() {
				break
			}
			continue
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *QUICTransportServer) handleConnection(conn quic.Connection) {
	defer s.wg.Done()

	for !s.closed.Load() {
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			break
		}

		s.wg.Add(1)
		go s.handleStream(stream)
	}
}

func (s *QUICTransportServer) handleStream(stream quic.Stream) {
	defer s.wg.Done()
	defer stream.Close()

	var hdrBuf [32]byte
	if _, err := io.ReadFull(stream, hdrBuf[:]); err != nil {
		return
	}

	hdr, err := DecodeQUICHeader(hdrBuf[:])
	if err != nil {
		return
	}

	payload := make([]byte, hdr.PayloadLen)
	if _, err := io.ReadFull(stream, payload); err != nil {
		return
	}

	if HardwareCRC32C(payload) != hdr.CRC32C {
		return
	}

	s.received.Add(1)
	s.bytesRx.Add(uint64(QUICHeaderSize + int(hdr.PayloadLen)))

	switch hdr.Cmd {
	case CmdPublish:
		if s.engine != nil && len(payload) >= 16 {
			idLow := binary.LittleEndian.Uint64(payload[0:8])
			idHigh := binary.LittleEndian.Uint64(payload[8:16])
			body := payload[16:]
			_ = s.engine.Publish(idLow, idHigh, uint32(hdr.PartitionID), 0, body)
		}
		if s.table != nil && len(payload) >= 8 {
			key := binary.LittleEndian.Uint64(payload[0:8])
			_ = s.table.Set(key, payload[8:], 0)
		}

		var ackHdr [32]byte
		ack := QUICHeader{
			Magic:       QUICFrameMagic,
			Cmd:         CmdAck,
			PartitionID: hdr.PartitionID,
			LSN:         hdr.LSN,
			Offset:      hdr.Offset,
		}
		ack.Encode(&ackHdr)
		_, _ = stream.Write(ackHdr[:])
	}
}

// Addr retourne l'adresse d'écoute réseau.
func (s *QUICTransportServer) Addr() string {
	return s.listener.Addr().String()
}

// Received retourne le nombre de messages reçus et validés.
func (s *QUICTransportServer) Received() uint64 {
	return s.received.Load()
}

// Close arrête proprement l'écouteur QUIC.
func (s *QUICTransportServer) Close() error {
	if s.closed.Swap(true) {
		return nil
	}
	_ = s.listener.Close()
	s.wg.Wait()
	return nil
}

// QUICTransportClient assure l'émission vers un serveur QUIC/TLS 1.3 distant.
type QUICTransportClient struct {
	conn  quic.Connection
	txLSN atomic.Uint64
}

// DialQUICTransportWithTLS connecte un client QUIC avec configuration TLS 1.3 explicite.
func DialQUICTransportWithTLS(targetAddr string, tlsConf *tls.Config) (*QUICTransportClient, error) {
	quicConf := &quic.Config{
		MaxIdleTimeout:                 30 * time.Second,
		InitialStreamReceiveWindow:     16 * 1024 * 1024,
		InitialConnectionReceiveWindow: 64 * 1024 * 1024,
	}

	conn, err := quic.DialAddr(context.Background(), targetAddr, tlsConf, quicConf)
	if err != nil {
		return nil, fmt.Errorf("c2q55/quic: dial failed: %w", err)
	}

	return &QUICTransportClient{
		conn: conn,
	}, nil
}

// DialQUICTransport connecte un client en utilisant la CA de cluster standard (mTLS strict).
func DialQUICTransport(targetAddr string) (*QUICTransportClient, error) {
	ca := GetDefaultTestCA()

	nodeCert, err := ca.IssueNodeCertificate("node-client", []net.IP{net.ParseIP("127.0.0.1")}, []string{"localhost", "node-client"})
	if err != nil {
		return nil, err
	}

	tlsConf := ca.ClientTLSConfig(nodeCert, "node-server")
	return DialQUICTransportWithTLS(targetAddr, tlsConf)
}

// SendPublish envoie un message (jusqu'à plusieurs Mo) via un stream QUIC dédié avec accusé de réception synchrone.
func (c *QUICTransportClient) SendPublish(partitionID uint16, idLow, idHigh uint64, body []byte, waitForAck bool) error {
	stream, err := c.conn.OpenStreamSync(context.Background())
	if err != nil {
		return fmt.Errorf("c2q55/quic: open stream failed: %w", err)
	}
	defer stream.Close()

	payloadLen := 16 + len(body)
	payload := make([]byte, payloadLen)

	binary.LittleEndian.PutUint64(payload[0:8], idLow)
	binary.LittleEndian.PutUint64(payload[8:16], idHigh)
	copy(payload[16:], body)

	crc := HardwareCRC32C(payload)
	lsn := c.txLSN.Add(1)

	hdr := QUICHeader{
		Magic:       QUICFrameMagic,
		Cmd:         CmdPublish,
		PartitionID: partitionID,
		LSN:         lsn,
		PayloadLen:  uint32(payloadLen),
		CRC32C:      crc,
	}

	var hdrBytes [32]byte
	hdr.Encode(&hdrBytes)

	if _, err := stream.Write(hdrBytes[:]); err != nil {
		return err
	}
	if _, err := stream.Write(payload); err != nil {
		return err
	}

	if waitForAck {
		var ackBuf [32]byte
		if _, err := io.ReadFull(stream, ackBuf[:]); err != nil {
			return fmt.Errorf("c2q55/quic: read ack failed: %w", err)
		}
		ackHdr, err := DecodeQUICHeader(ackBuf[:])
		if err != nil {
			return err
		}
		if ackHdr.Cmd != CmdAck || ackHdr.LSN != lsn {
			return errors.New("c2q55/quic: invalid ack")
		}
	}

	return nil
}

// Close ferme la session QUIC cliente.
func (c *QUICTransportClient) Close() error {
	return c.conn.CloseWithError(0, "client closed")
}
