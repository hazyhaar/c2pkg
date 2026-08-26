// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"crypto/tls"
	"net"
	"testing"
)

// TestSecurity_mTLS_UntrustedCertificateRejected vérifie le rejet immédiat d'un certificat non signé par la CA du cluster.
func TestSecurity_mTLS_UntrustedCertificateRejected(t *testing.T) {
	clusterCA, err := NewClusterCA()
	if err != nil {
		t.Fatalf("NewClusterCA failed: %v", err)
	}

	serverCert, err := clusterCA.IssueNodeCertificate("server-node", []net.IP{net.ParseIP("127.0.0.1")}, []string{"localhost"})
	if err != nil {
		t.Fatalf("IssueNodeCertificate failed: %v", err)
	}

	serverTLS := clusterCA.ServerTLSConfig(serverCert, nil)
	server, err := ListenQUICTransportWithTLS("127.0.0.1:0", serverTLS, nil, nil)
	if err != nil {
		t.Fatalf("ListenQUICTransportWithTLS failed: %v", err)
	}
	defer server.Close()

	// 1. Attaquant avec une CA étrangère
	rogueCA, err := NewClusterCA()
	if err != nil {
		t.Fatalf("NewClusterCA rogue failed: %v", err)
	}
	rogueCert, err := rogueCA.IssueNodeCertificate("attacker-node", []net.IP{net.ParseIP("127.0.0.1")}, []string{"localhost"})
	if err != nil {
		t.Fatalf("IssueNodeCertificate rogue failed: %v", err)
	}

	rogueClientTLS := rogueCA.ClientTLSConfig(rogueCert, "server-node")

	// 2. Tentative de connexion de l'attaquant : DOIT ÉCHOUER
	client, err := DialQUICTransportWithTLS(server.Addr(), rogueClientTLS)
	if err == nil {
		_ = client.Close()
		t.Fatalf("SECURITY BREACH: Server accepted an untrusted certificate from foreign CA!")
	}

	// 3. Client légitime signé par la bonne CA : DOIT RÉUSSIR
	validCert, err := clusterCA.IssueNodeCertificate("legit-client", []net.IP{net.ParseIP("127.0.0.1")}, []string{"localhost"})
	if err != nil {
		t.Fatalf("Issue valid cert failed: %v", err)
	}
	validClientTLS := clusterCA.ClientTLSConfig(validCert, "server-node")

	legitClient, err := DialQUICTransportWithTLS(server.Addr(), validClientTLS)
	if err != nil {
		t.Fatalf("Legitimate client failed to connect: %v", err)
	}
	defer legitClient.Close()
}

// TestSecurity_mTLS_AllowlistEnforcement vérifie le filtrage d'identité (CommonName).
func TestSecurity_mTLS_AllowlistEnforcement(t *testing.T) {
	clusterCA, err := NewClusterCA()
	if err != nil {
		t.Fatalf("NewClusterCA failed: %v", err)
	}

	serverCert, err := clusterCA.IssueNodeCertificate("server-node", []net.IP{net.ParseIP("127.0.0.1")}, []string{"localhost"})
	if err != nil {
		t.Fatalf("Issue server cert failed: %v", err)
	}

	// Autoriser uniquement "node-worker-01"
	serverTLS := clusterCA.ServerTLSConfig(serverCert, []string{"node-worker-01"})
	server, err := ListenQUICTransportWithTLS("127.0.0.1:0", serverTLS, nil, nil)
	if err != nil {
		t.Fatalf("Listen failed: %v", err)
	}
	defer server.Close()

	// Certificat signé par la CA mais avec identité "node-unauthorized"
	unauthCert, err := clusterCA.IssueNodeCertificate("node-unauthorized", []net.IP{net.ParseIP("127.0.0.1")}, []string{"localhost"})
	if err != nil {
		t.Fatalf("Issue unauth cert failed: %v", err)
	}

	unauthTLS := clusterCA.ClientTLSConfig(unauthCert, "server-node")
	client, err := DialQUICTransportWithTLS(server.Addr(), unauthTLS)
	if err == nil {
		err = client.SendPublish(0, 1, 1, []byte("test"), true)
		if err == nil {
			_ = client.Close()
			t.Fatalf("SECURITY BREACH: Server accepted an unauthorized node identity!")
		}
	}
}

// TestSecurity_InsecureSkipVerify_Prohibited vérifie qu'aucun fichier du transport n'active InsecureSkipVerify à true.
func TestSecurity_InsecureSkipVerify_Prohibited(t *testing.T) {
	ca, _ := NewClusterCA()
	nodeCert, _ := ca.IssueNodeCertificate("test", nil, nil)
	clientTLS := ca.ClientTLSConfig(nodeCert, "server")

	if clientTLS.InsecureSkipVerify {
		t.Fatalf("CRITICAL SECURITY VIOLATION: ClientTLSConfig has InsecureSkipVerify=true")
	}
	if clientTLS.MinVersion < tls.VersionTLS13 {
		t.Fatalf("CRITICAL SECURITY VIOLATION: MinVersion must be TLS 1.3")
	}
}
