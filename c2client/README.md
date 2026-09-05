# c2client — Client Réseau Distribué, Sharding Déterministe & QUIC 0-RTT

`c2client` fournit un client réseau distribué haute performance conçu pour les architectures temps réel à forte charge (systèmes distribués, trading, services agentiques).

Il élimine structurellement les goulets d'étranglement des architectures classiques (blocage en tête de ligne HTTP/2, cache stampede, saturation de liaisons centralisées, latence de poignée de main TLS).

---

## 1. Caractéristiques Principales

* **Partitionnement Déterministe sur 1024 Fragments Virtuels (`NumShards = 1024`) :**
  Calcul instantané du nœud cible par hachage SIMD Blake3 en pur Go (`RouteShard(key)`). Dispersion uniforme et affinité de cache processeur (L1/L2 hits maximisés).
* **Transport UDP / QUIC & Multiplexage Non-Bloquant :**
  Flux QUIC indépendants éliminant le phénomène de *Head-of-Line Blocking* inhérent aux connexions TCP/HTTP.
* **Poignée de Main mTLS Souveraine TLS 1.3 (Zero-CA) :**
  Clés et certificats éphémères Ed25519 (RFC 8410) dérivés mathématiquement via Blake3 KDF à partir de l'identité du tenant et de l'époque active. Aucune autorité de certification centrale externe requise.
* **Reconnexion 0-RTT sans État :**
  Établissement immédiat des sessions clientes sans surcoût d'aller-retour réseau (`quic.DialAddrEarly`).
* **Cadrage Binaire Zéro-Allocation :**
  Entiers à longueur variable Vint RFC 9000 encodés et décodés sans allocation mémoire sur le tas via `c2quic`.

---

## 2. Exemple d'Utilisation

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hazyhaar/c2pkg/c2client"
	"github.com/hazyhaar/c2pkg/c2uuidv7"
)

func main() {
	masterKey := [32]byte{ /* 32 octets secrets */ }
	tenantID := uint16(1)
	epoch := c2uuidv7.New()

	// Liste des adresses UDP QUIC du cluster
	nodes := []string{
		"10.0.0.1:8155",
		"10.0.0.2:8155",
		"10.0.0.3:8155",
	}

	client, err := c2client.NewClient(masterKey, tenantID, epoch, nodes)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 1. Calcul du shard local ou routage
	key := []byte("ticker:binance:btc-usdt")
	shard, addr := client.Route(key)
	fmt.Printf("Clé routée sur shard %d (nœud %s)\n", shard, addr)

	// 2. Écriture atomique
	val := []byte(`{"price": 64500.0, "qty": 1.25}`)
	if err := client.Put(ctx, key, val); err != nil {
		log.Fatalf("erreur Put: %v", err)
	}

	// 3. Lecture directe
	data, err := client.Get(ctx, key)
	if err != nil {
		log.Fatalf("erreur Get: %v", err)
	}
	fmt.Printf("Donnée récupérée: %s\n", string(data))

	// 4. Requêtage par préfixe distribué sur le cluster
	entries, err := client.Query([]byte("ticker:")).Collect(ctx)
	if err != nil {
		log.Fatalf("erreur Query: %v", err)
	}
	for _, entry := range entries {
		fmt.Printf("-> %s : %s\n", string(entry.Key), string(entry.Val))
	}
}
```

---

## 3. Formule de Partitionnement Déterministe

Le partitionnement direct s'effectue sans dépendance réseau :

```go
shard := c2client.RouteShard([]byte("ma-cle")) // Retourne 0..1023
```

Formule sous-jacente :
```go
sum := blake3archtsim.Sum256(key)
return binary.BigEndian.Uint16(sum[:2]) >> 6
```

---

## 4. Qualification & Tests

```bash
go test -race -count=1 ./c2client
```

---

## 5. Licence

Distribué sous double licence **Apache 2.0** et **MIT**.
