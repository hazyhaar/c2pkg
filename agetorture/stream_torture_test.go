package agetorture

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"golang.org/x/crypto/chacha20poly1305"
)

const chunkSize = 64 * 1024

var ErrTruncated = errors.New("stream truncated without terminal flag")

// ageStreamReader implémente un automate de déchiffrement STREAM (spécifique à age).
type ageStreamReader struct {
	r      io.Reader
	key    []byte
	nonce  [11]byte
	aead   cipher.AEAD
	buf    []byte
	c      uint64 // Compteur de segment (age utilise un entier grand-boutien)
	closed bool
}

func newAgeStreamReader(r io.Reader, key []byte, nonce [11]byte) (*ageStreamReader, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}
	return &ageStreamReader{
		r:     r,
		key:   key,
		nonce: nonce,
		aead:  aead,
	}, nil
}

func (s *ageStreamReader) Read(p []byte) (int, error) {
	if len(s.buf) > 0 {
		n := copy(p, s.buf)
		s.buf = s.buf[n:]
		return n, nil
	}
	if s.closed {
		return 0, io.EOF
	}

	// Lecture d'un segment: jusqu'à chunkSize + 16 (MAC)
	cipherBuf := make([]byte, chunkSize+16)
	n, readErr := io.ReadFull(s.r, cipherBuf)
	if readErr != nil && readErr != io.EOF && readErr != io.ErrUnexpectedEOF {
		return 0, readErr
	}
	if n == 0 {
		return 0, ErrTruncated
	}
	cipherBuf = cipherBuf[:n]

	// Construction de la nonce: 11 octets base + 1 octet drapeau
	var chunkNonce [12]byte
	binary.BigEndian.PutUint64(chunkNonce[3:11], s.c)

	chunkNonce[11] = 0x00
	plain, decErr := s.aead.Open(nil, chunkNonce[:], cipherBuf, nil)
	if decErr != nil {
		// Échec avec 0x00, on tente 0x01
		chunkNonce[11] = 0x01
		plain, decErr = s.aead.Open(nil, chunkNonce[:], cipherBuf, nil)
		if decErr != nil {
			return 0, decErr
		}
		s.closed = true
	} else if len(plain) < chunkSize {
		// Un segment non terminal doit faire exactement chunkSize
		return 0, errors.New("invalid non-terminal chunk size")
	}

	// Si l'erreur de lecture n'était pas nil et que le segment n'est pas terminal
	if (readErr == io.EOF || readErr == io.ErrUnexpectedEOF) && !s.closed {
		return 0, ErrTruncated
	}

	s.c++
	s.buf = plain
	
	nCopied := copy(p, s.buf)
	s.buf = s.buf[nCopied:]
	return nCopied, nil
}

// Outil de chiffrement pour les tests
func sealChunk(key []byte, counter uint64, last bool, plain []byte) []byte {
	aead, _ := chacha20poly1305.New(key)
	var nonce [12]byte
	binary.BigEndian.PutUint64(nonce[3:11], counter)
	if last {
		nonce[11] = 0x01
	}
	return aead.Seal(nil, nonce[:], plain, nil)
}

func TestStreamAgeTruncation(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	var stream bytes.Buffer
	// Segment 0 non terminal
	stream.Write(sealChunk(key, 0, false, make([]byte, chunkSize)))
	// Segment 1 non terminal
	stream.Write(sealChunk(key, 1, false, make([]byte, chunkSize)))
	// COUPURE ABRUPTE: on ne fournit pas le segment terminal (0x01)

	r, _ := newAgeStreamReader(&stream, key, [11]byte{})
	buf := make([]byte, chunkSize)
	
	_, err := r.Read(buf) // Segment 0
	if err != nil {
		t.Fatalf("Erreur inattendue au segment 0 : %v", err)
	}
	_, err = r.Read(buf) // Segment 1
	if err != nil {
		t.Fatalf("Erreur inattendue au segment 1 : %v", err)
	}
	
	_, err = r.Read(buf) // Fin prématurée
	if err != ErrTruncated {
		t.Fatalf("Attendu ErrTruncated, reçu: %v", err)
	}
}

func TestStreamAgeReplay(t *testing.T) {
	key := make([]byte, 32)
	
	var stream bytes.Buffer
	chunk0 := sealChunk(key, 0, false, make([]byte, chunkSize))
	stream.Write(chunk0)
	// REJEU: on réinjecte le chunk 0 en position 1
	stream.Write(chunk0)
	
	r, _ := newAgeStreamReader(&stream, key, [11]byte{})
	buf := make([]byte, chunkSize)
	
	_, err := r.Read(buf) // Segment 0
	if err != nil {
		t.Fatalf("Erreur inattendue : %v", err)
	}
	_, err = r.Read(buf) // Rejeu en pos 1
	if err == nil {
		t.Fatal("Le rejeu a été accepté, violation d'intégrité")
	}
}

func TestStreamAgePostClosureExtension(t *testing.T) {
	key := make([]byte, 32)
	
	var stream bytes.Buffer
	// chunk terminal plein (pour éviter que l'automate ne fusionne les lectures sur un EOF anticipé)
	stream.Write(sealChunk(key, 0, true, make([]byte, chunkSize)))
	// EXTENSION ILLÉGALE: ajout d'un chunk après la clôture
	stream.Write(sealChunk(key, 1, true, []byte("post-closure")))
	
	r, _ := newAgeStreamReader(&stream, key, [11]byte{})
	buf := make([]byte, chunkSize)
	
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatalf("Erreur inattendue: %v", err)
	}
	if n != chunkSize {
		t.Fatalf("Taille lue invalide")
	}
	
	// Tentative de lecture après clôture
	n, err = r.Read(buf)
	if err != io.EOF {
		t.Fatalf("Attendu EOF pur après clôture, reçu: %v (n=%d)", err, n)
	}
}

func TestStreamAgeBoundaries(t *testing.T) {
	key := make([]byte, 32)
	
	// Frontières ChaCha20/Poly1305 (N*64 ± 1)
	sizes := []int{63, 64, 65, 127, 128, 129}
	
	for _, sz := range sizes {
		var stream bytes.Buffer
		stream.Write(sealChunk(key, 0, false, make([]byte, chunkSize)))
		stream.Write(sealChunk(key, 1, true, make([]byte, sz)))
		
		r, _ := newAgeStreamReader(&stream, key, [11]byte{})
		out, err := io.ReadAll(r)
		
		if err != nil {
			t.Fatalf("Échec au balayage continu (sz=%d): %v", sz, err)
		}
		if len(out) != chunkSize+sz {
			t.Fatalf("Longueur déchiffrée incorrecte (sz=%d): attendu %d, reçu %d", sz, chunkSize+sz, len(out))
		}
	}
}
