// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2client

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/hazyhaar/c2pkg/c2quic"
)

const (
	MagicProto     uint16 = 0xC250 // "C2P"
	ShardBroadcast uint16 = 0xFFFF // Indique une requête à diffuser sur l'ensemble des shards

	OpGet       uint8 = 0x01
	OpPut       uint8 = 0x02
	OpDel       uint8 = 0x03
	OpQuery     uint8 = 0x04
	OpStreamVal uint8 = 0x05
	OpOK        uint8 = 0x80
	OpErr       uint8 = 0x81
	OpEntry     uint8 = 0x82
	OpEnd       uint8 = 0x83

	MaxKeyLen uint32 = 65535           // Borne uint16 exacte (0xFFFF)
	MaxValLen uint32 = 16 * 1024 * 1024 // 16 Mo
)

var (
	ErrInvalidMagic  = errors.New("c2client: invalid protocol magic")
	ErrFrameTooLarge = errors.New("c2client: frame exceeds maximum capacity")
	ErrShortRead     = errors.New("c2client: short read in frame payload")
)

// Entry représente une paire clé-valeur transmise sur le réseau.
type Entry struct {
	Key []byte
	Val []byte
}

// Header décrit l'en-tête d'une trame applicative encodée avec les entiers Vint
// RFC 9000 du moteur matériel c2quic.
type Header struct {
	Magic  uint16
	OpCode uint8
	Flags  uint8
	Tenant uint16
	Shard  uint16
	KeyLen uint16
	ValLen uint32
	Extra  uint64
}

// EncodeHeader sérialise l'en-tête sur le flux en utilisant un unique buffer de pile
// contigu [48]byte et les Vints RFC 9000 sans allocation sur le tas.
func EncodeHeader(w io.Writer, h Header) error {
	var hdrBuf [48]byte
	binary.LittleEndian.PutUint16(hdrBuf[:2], MagicProto)
	n := 2

	var nn int
	var err error

	if nn, err = putVint(hdrBuf[n:], uint64(h.OpCode)); err != nil {
		return err
	}
	n += nn
	if nn, err = putVint(hdrBuf[n:], uint64(h.Flags)); err != nil {
		return err
	}
	n += nn
	if nn, err = putVint(hdrBuf[n:], uint64(h.Tenant)); err != nil {
		return err
	}
	n += nn
	if nn, err = putVint(hdrBuf[n:], uint64(h.Shard)); err != nil {
		return err
	}
	n += nn
	if nn, err = putVint(hdrBuf[n:], uint64(h.KeyLen)); err != nil {
		return err
	}
	n += nn
	if nn, err = putVint(hdrBuf[n:], uint64(h.ValLen)); err != nil {
		return err
	}
	n += nn
	if nn, err = putVint(hdrBuf[n:], h.Extra); err != nil {
		return err
	}
	n += nn

	_, err = w.Write(hdrBuf[:n])
	return err
}

// DecodeHeader lit et décode l'en-tête de trame via les primitives RFC 9000 de c2quic.
func DecodeHeader(r io.Reader) (Header, error) {
	var magicBuf [2]byte
	if _, err := io.ReadFull(r, magicBuf[:]); err != nil {
		return Header{}, err
	}
	magic := binary.LittleEndian.Uint16(magicBuf[:])
	if magic != MagicProto {
		return Header{}, fmt.Errorf("%w: got 0x%04X want 0x%04X", ErrInvalidMagic, magic, MagicProto)
	}

	opCode, err := readVint(r)
	if err != nil {
		return Header{}, err
	}
	if opCode > 0xFF {
		return Header{}, fmt.Errorf("%w: opCode out of range %d", ErrFrameTooLarge, opCode)
	}

	flags, err := readVint(r)
	if err != nil {
		return Header{}, err
	}
	if flags > 0xFF {
		return Header{}, fmt.Errorf("%w: flags out of range %d", ErrFrameTooLarge, flags)
	}

	tenant, err := readVint(r)
	if err != nil {
		return Header{}, err
	}
	if tenant > 0xFFFF {
		return Header{}, fmt.Errorf("%w: tenant out of range %d", ErrFrameTooLarge, tenant)
	}

	shard, err := readVint(r)
	if err != nil {
		return Header{}, err
	}
	if shard > 0xFFFF {
		return Header{}, fmt.Errorf("%w: shard out of range %d", ErrFrameTooLarge, shard)
	}

	keyLen, err := readVint(r)
	if err != nil {
		return Header{}, err
	}
	if keyLen > uint64(MaxKeyLen) {
		return Header{}, fmt.Errorf("%w: keyLen %d exceeds maximum %d", ErrFrameTooLarge, keyLen, MaxKeyLen)
	}

	valLen, err := readVint(r)
	if err != nil {
		return Header{}, err
	}
	if valLen > uint64(MaxValLen) {
		return Header{}, fmt.Errorf("%w: valLen %d exceeds maximum %d", ErrFrameTooLarge, valLen, MaxValLen)
	}

	extra, err := readVint(r)
	if err != nil {
		return Header{}, err
	}

	return Header{
		Magic:  magic,
		OpCode: uint8(opCode),
		Flags:  uint8(flags),
		Tenant: uint16(tenant),
		Shard:  uint16(shard),
		KeyLen: uint16(keyLen),
		ValLen: uint32(valLen),
		Extra:  extra,
	}, nil
}

// WriteFrame écrit atomiquement l'en-tête Vint c2quic et les charges utiles associées.
func WriteFrame(w io.Writer, h Header, key, val []byte) error {
	if len(key) > int(MaxKeyLen) {
		return fmt.Errorf("%w: key exceeds maximum %d bytes", ErrFrameTooLarge, MaxKeyLen)
	}
	if len(val) > int(MaxValLen) {
		return fmt.Errorf("%w: val exceeds maximum %d bytes", ErrFrameTooLarge, MaxValLen)
	}
	h.KeyLen = uint16(len(key))
	h.ValLen = uint32(len(val))
	if err := EncodeHeader(w, h); err != nil {
		return err
	}
	if len(key) > 0 {
		if _, err := w.Write(key); err != nil {
			return err
		}
	}
	if len(val) > 0 {
		if _, err := w.Write(val); err != nil {
			return err
		}
	}
	return nil
}

// ReadFramePayload lit les tranches clé et valeur selon les longueurs spécifiées par l'en-tête.
func ReadFramePayload(r io.Reader, h Header) (key, val []byte, err error) {
	if uint32(h.KeyLen) > MaxKeyLen || h.ValLen > MaxValLen {
		return nil, nil, ErrFrameTooLarge
	}
	if h.KeyLen > 0 {
		key = make([]byte, h.KeyLen)
		if _, err := io.ReadFull(r, key); err != nil {
			return nil, nil, err
		}
	}
	if h.ValLen > 0 {
		val = make([]byte, h.ValLen)
		if _, err := io.ReadFull(r, val); err != nil {
			return nil, nil, err
		}
	}
	return key, val, nil
}

// ReadFramePayloadInto lit la charge utile directement dans les tampons pré-alloués fournis
// par l'appelant, éliminant toute allocation sur le tas sur le chemin chaud.
func ReadFramePayloadInto(r io.Reader, h Header, keyBuf, valBuf []byte) (keyLen, valLen int, err error) {
	if uint32(h.KeyLen) > MaxKeyLen || h.ValLen > MaxValLen {
		return 0, 0, ErrFrameTooLarge
	}
	if int(h.KeyLen) > cap(keyBuf) || int(h.ValLen) > cap(valBuf) {
		return 0, 0, ErrShortRead
	}
	if h.KeyLen > 0 {
		if _, err := io.ReadFull(r, keyBuf[:h.KeyLen]); err != nil {
			return 0, 0, err
		}
	}
	if h.ValLen > 0 {
		if _, err := io.ReadFull(r, valBuf[:h.ValLen]); err != nil {
			return int(h.KeyLen), 0, err
		}
	}
	return int(h.KeyLen), int(h.ValLen), nil
}

func readVint(r io.Reader) (uint64, error) {
	var first [1]byte
	if _, err := io.ReadFull(r, first[:]); err != nil {
		return 0, err
	}
	n := 1 << (first[0] >> 6)
	var buf [8]byte
	buf[0] = first[0]
	if n > 1 {
		if _, err := io.ReadFull(r, buf[1:n]); err != nil {
			return 0, err
		}
	}
	var out c2quic.C2quic_vint_t
	if st := c2quic.C2quic_vint_decode(buf[:n], uint32(n), 0, &out); st != c2quic.OK {
		return 0, fmt.Errorf("c2client: vint decode failed: st=%d", st)
	}
	return out.Value, nil
}


func putVint(buf []byte, val uint64) (int, error) {
	if val > (1<<62 - 1) {
		return 0, fmt.Errorf("c2client: vint encode failed: value %d exceeds 62-bit RFC 9000 limit", val)
	}
	var out c2quic.C2quic_vint_t
	if st := c2quic.C2quic_vint_encode(val, buf, uint32(len(buf)), 0, &out); st != c2quic.OK {
		return 0, fmt.Errorf("c2client: vint encode failed: st=%d", st)
	}
	return int(out.Nbytes), nil
}

