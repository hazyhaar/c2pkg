// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2client

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/hazyhaar/c2pkg/blake3archtsim"
	"github.com/hazyhaar/c2pkg/c2uuidv7"
)

const (
	TenantMACContext    = "c2cluster/tenant-mac/v1"
	serverDomainContext = "c2cluster/tenant-server/v1"
	clientDomainContext = "c2cluster/tenant-client/v1"
)

// DeriveTenantMAC dérive cryptographiquement la graine d'identité du tenant
// à partir de la clé maîtresse, de l'identifiant numérique de tenant et de l'époque active.
func DeriveTenantMAC(master [32]byte, tenant uint16, epoch c2uuidv7.UUID) [32]byte {
	var mat [50]byte
	copy(mat[:32], master[:])
	binary.LittleEndian.PutUint16(mat[32:34], tenant)
	copy(mat[34:], epoch[:])
	return blake3archtsim.DeriveKey(TenantMACContext, mat[:])
}

// GenerateTenantTLSConfig génère une configuration TLS 1.3 souveraine (Zero-CA)
// avec séparation stricte des domaines cryptographiques client et serveur.
// Les clés Ed25519 dérivées respectent strictement la norme RFC 8410.
func GenerateTenantTLSConfig(tenantSeed [32]byte, tenant uint16, isServer bool) (*tls.Config, error) {
	serverSeed := blake3archtsim.DeriveKey(serverDomainContext, tenantSeed[:])
	clientSeed := blake3archtsim.DeriveKey(clientDomainContext, tenantSeed[:])

	serverPriv := ed25519.NewKeyFromSeed(serverSeed[:])
	serverPub := serverPriv.Public().(ed25519.PublicKey)

	clientPriv := ed25519.NewKeyFromSeed(clientSeed[:])
	clientPub := clientPriv.Public().(ed25519.PublicKey)

	var localPriv ed25519.PrivateKey
	var localPub ed25519.PublicKey
	var expectedPeerPub ed25519.PublicKey
	var extKeyUsage x509.ExtKeyUsage

	if isServer {
		localPriv = serverPriv
		localPub = serverPub
		expectedPeerPub = clientPub
		extKeyUsage = x509.ExtKeyUsageServerAuth
	} else {
		localPriv = clientPriv
		localPub = clientPub
		expectedPeerPub = serverPub
		extKeyUsage = x509.ExtKeyUsageClientAuth
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	roleName := "client"
	if isServer {
		roleName = "server"
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   fmt.Sprintf("tenant-%04x-%s", tenant, roleName),
			Organization: []string{"c2client"},
		},
		NotBefore:             time.Now().Add(-1 * time.Hour),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature, // RFC 8410 strict
		ExtKeyUsage:           []x509.ExtKeyUsage{extKeyUsage},
		BasicConstraintsValid: true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, localPub, localPriv)
	if err != nil {
		return nil, err
	}

	cert := tls.Certificate{
		Certificate: [][]byte{certDER},
		PrivateKey:  localPriv,
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS13,
		NextProtos:   []string{"c2postnet-v1"},
	}

	if isServer {
		tlsConfig.ClientAuth = tls.RequireAnyClientCert
		tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
			if len(rawCerts) == 0 {
				return errors.New("c2client: missing peer certificate")
			}
			peerCert, err := x509.ParseCertificate(rawCerts[0])
			if err != nil {
				return fmt.Errorf("c2client: parse peer certificate: %w", err)
			}
			peerPub, ok := peerCert.PublicKey.(ed25519.PublicKey)
			if !ok {
				return errors.New("c2client: invalid peer public key type (must be ed25519)")
			}
			if !bytes.Equal(peerPub, expectedPeerPub) {
				return errors.New("c2client: peer public key mismatch (unauthorized tenant or epoch)")
			}
			return nil
		}
	} else {
		tlsConfig.ClientSessionCache = tls.NewLRUClientSessionCache(64)
		tlsConfig.InsecureSkipVerify = true
		tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
			if len(rawCerts) == 0 {
				return errors.New("c2client: missing peer certificate")
			}
			peerCert, err := x509.ParseCertificate(rawCerts[0])
			if err != nil {
				return fmt.Errorf("c2client: parse peer certificate: %w", err)
			}
			peerPub, ok := peerCert.PublicKey.(ed25519.PublicKey)
			if !ok {
				return errors.New("c2client: invalid peer public key type (must be ed25519)")
			}
			if !bytes.Equal(peerPub, expectedPeerPub) {
				return errors.New("c2client: peer public key mismatch (unauthorized tenant or epoch)")
			}
			return nil
		}
	}

	return tlsConfig, nil
}
