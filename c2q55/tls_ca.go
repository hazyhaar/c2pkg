package c2q55

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"sync"
	"time"
)

var (
	ErrUntrustedPeer    = errors.New("c2q55/tls: peer certificate not trusted by cluster CA")
	ErrUnauthorizedNode = errors.New("c2q55/tls: peer node ID not in cluster allowlist")
)

// ClusterCA représente l'autorité de certification racine du parc horos55.
type ClusterCA struct {
	Cert    *x509.Certificate
	Key     *ecdsa.PrivateKey
	CertDER []byte
	Pool    *x509.CertPool
}

var (
	defaultTestCA     *ClusterCA
	defaultTestCAOnce sync.Once
)

// GetDefaultTestCA retourne une CA de test commune et déterministe pour le loopback.
func GetDefaultTestCA() *ClusterCA {
	defaultTestCAOnce.Do(func() {
		ca, err := NewClusterCA()
		if err != nil {
			panic(err)
		}
		defaultTestCA = ca
	})
	return defaultTestCA
}

// NewClusterCA initialise une nouvelle autorité racine pour le cluster.
func NewClusterCA() (*ClusterCA, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("c2q55/ca: generate key failed: %w", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"HOROS55 Sovereign Infrastructure CA"},
			CommonName:   "horos55.cluster.ca",
		},
		NotBefore:             time.Now().Add(-1 * time.Hour),
		NotAfter:              time.Now().Add(24 * 365 * 10 * time.Hour), // 10 ans
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return nil, fmt.Errorf("c2q55/ca: create cert failed: %w", err)
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AddCert(cert)

	return &ClusterCA{
		Cert:    cert,
		Key:     key,
		CertDER: certDER,
		Pool:    pool,
	}, nil
}

// IssueNodeCertificate délivre un certificat signé par la CA pour un nœud spécifique.
func (ca *ClusterCA) IssueNodeCertificate(nodeID string, ips []net.IP, dnsNames []string) (tls.Certificate, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return tls.Certificate{}, err
	}

	serialNumber, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))

	dnsList := append([]string{nodeID, "localhost"}, dnsNames...)
	ipList := append([]net.IP{net.ParseIP("127.0.0.1"), net.IPv6loopback}, ips...)

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"HOROS55 Cluster Member"},
			CommonName:   nodeID,
		},
		NotBefore:             time.Now().Add(-1 * time.Hour),
		NotAfter:              time.Now().Add(24 * 365 * 2 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IPAddresses:           ipList,
		DNSNames:              dnsList,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, ca.Cert, &key.PublicKey, ca.Key)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("c2q55/ca: sign cert failed: %w", err)
	}

	return tls.Certificate{
		Certificate: [][]byte{certDER, ca.CertDER},
		PrivateKey:  key,
	}, nil
}

// ServerTLSConfig configure le serveur avec mTLS strict (ClientAuth RequireAndVerifyClientCert).
func (ca *ClusterCA) ServerTLSConfig(nodeCert tls.Certificate, allowlist []string) *tls.Config {
	allowMap := make(map[string]bool)
	for _, id := range allowlist {
		allowMap[id] = true
	}

	return &tls.Config{
		Certificates: []tls.Certificate{nodeCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    ca.Pool,
		NextProtos:   []string{QUICALPN},
		MinVersion:   tls.VersionTLS13,
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			if len(rawCerts) == 0 {
				return ErrUntrustedPeer
			}
			peerCert, err := x509.ParseCertificate(rawCerts[0])
			if err != nil {
				return err
			}
			if len(allowlist) > 0 && !allowMap[peerCert.Subject.CommonName] {
				return fmt.Errorf("%w: node %q is not authorized", ErrUnauthorizedNode, peerCert.Subject.CommonName)
			}
			return nil
		},
	}
}

// ClientTLSConfig configure le client avec mTLS strict (InsecureSkipVerify: false obligatoire).
func (ca *ClusterCA) ClientTLSConfig(nodeCert tls.Certificate, serverExpectedName string) *tls.Config {
	return &tls.Config{
		Certificates:       []tls.Certificate{nodeCert},
		RootCAs:            ca.Pool,
		ServerName:         serverExpectedName,
		InsecureSkipVerify: false, // INTERDICTION DE SKIP
		NextProtos:         []string{QUICALPN},
		MinVersion:         tls.VersionTLS13,
	}
}

// SaveCertAndKeyPEM sauvegarde un certificat et sa clé privée au format PEM sur disque.
func SaveCertAndKeyPEM(certDER []byte, key *ecdsa.PrivateKey, certPath, keyPath string) error {
	certFile, err := os.OpenFile(certPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer certFile.Close()

	if err := pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER}); err != nil {
		return err
	}

	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	keyFile, err := os.OpenFile(keyPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	return pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})
}

// LoadNodeCertFromPEM charge une paire certificat / clé privée PEM depuis le disque.
func LoadNodeCertFromPEM(certPath, keyPath string) (tls.Certificate, error) {
	return tls.LoadX509KeyPair(certPath, keyPath)
}

// LoadClusterCAFromPEM charge un certificat racine CA PEM pour initialiser le CertPool.
func LoadClusterCAFromPEM(caCertPath string) (*x509.CertPool, error) {
	caData, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, fmt.Errorf("c2q55/ca: read ca cert failed: %w", err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caData) {
		return nil, errors.New("c2q55/ca: failed to parse ca certificate PEM")
	}

	return pool, nil
}
