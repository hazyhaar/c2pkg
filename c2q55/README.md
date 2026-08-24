# c2q55 — Moteur de Messagerie & Table d'État Souveraine Haute Performance

> **Projet HPM55 de référence :** `01a02fe2-9a1d-7302-b584-dfaf50be36d0`  
> **Doctrine :** C23 / CGO=0 / Pure Go 1.27 / TLS 1.3 / QUIC / ARCHTIME

---

## 1. Vue d'Ensemble & Piliers d'Architecture

`c2q55` est un moteur unifié combinant un journal d'événements ordonné (sémantique Kafka) et une table d'état compactée (sémantique Redis) sur un transport réseau chiffré QUIC TLS 1.3.

1. **Noyau de Créneaux (Slots) 64 octets :**
   - Alignement strict sur la taille de ligne de cache L1D ($64\text{ octets}$).
   - Zéro allocation mémoire sur le chemin chaud d'ingestion/consommation ($0\text{ B/op}$).
   - Inlinage de 16 octets pour les micro-messages ; déport transparent dans l'arène de slabs pour les corps volumineux.
2. **Arène de Slabs Contiguë & Ligne d'Eau Basse :**
   - Stockage mmapé contigu pour charges jusqu'à 16 Mo.
   - Protection absolue contre l'écrasement de corps vivants via repère d'eau basse (`lowWatermark`).
   - Somme de contrôle matérielle Castagnoli CRC32-C (débit SSE4.2 / ARMv8 > 34 Go/s).
   - Copie défensive isolée à la lecture.
3. **Consensus & Topologie Distribuée Multi-Nœuds :**
   - Consensus par quorum majoritaire ($\lfloor N/2 \rfloor + 1$).
   - Protection absolue contre le split-brain par époque de leadership strictement croissante (`LeaderEpoch` fencing).
   - Calcul dynamique du `HighWatermark` sur confirmation des followers synchronisés (ISR).
4. **Sécurité Réseau mTLS 1.3 Stricte :**
   - Autorité de certification racine souveraine (`ClusterCA`).
   - `ClientAuth: tls.RequireAndVerifyClientCert` obligatoire, `InsecureSkipVerify: false` strict.
   - Filtrage d'accès par liste blanche d'identités (`allowlist`).
5. **Durabilité des Décalages (`OffsetStore`) :**
   - Stockage mmapé synchronisé par `fdatasync`.
   - Reprise déterministe au redémarrage d'instance.

---

## 2. Commandes du Binaire CLI (`bin/c2q55-node`)

### Initialisation de la PKI Souveraine
```bash
c2q55-node pki-init --out-dir /etc/c2q55/pki
```

### Démarrage d'un Nœud Serveur QUIC (mTLS Strict)
```bash
c2q55-node listen \
  --addr 213.32.71.129:8443 \
  --ca-cert /etc/c2q55/pki/ca.crt \
  --node-cert /etc/c2q55/pki/redhost.crt \
  --node-key /etc/c2q55/pki/redhost.key \
  --allow REDBO \
  --wal /var/lib/c2q55/journal.wal \
  --slab /var/lib/c2q55/data.slab \
  --offset-dir /var/lib/c2q55/offsets
```

### Émission de Messages (Client / Producteur)
```bash
c2q55-node publish \
  --target 213.32.71.129:8443 \
  --server-name redhost \
  --ca-cert /etc/c2q55/pki/ca.crt \
  --node-cert /etc/c2q55/pki/redbo.crt \
  --node-key /etc/c2q55/pki/redbo.key \
  --topic 1 \
  --key 100 \
  --msg '{"event":"telemetry","status":"ok"}' \
  --count 10000
```

### Consommation et Validation d'Offsets
```bash
c2q55-node consume \
  --group telemetry-workers \
  --id worker-1 \
  --offset-dir /var/lib/c2q55/offsets
```

---

## 3. Matrice de Validation des Tests

```bash
# Vérification statique
go vet ./pkg/c2q55/...

# Suite complète sous détecteur de courses (0 data race)
go test -v -race -count=1 ./pkg/c2q55/...
```
