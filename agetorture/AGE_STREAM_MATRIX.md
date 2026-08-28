# Matrice d'épreuve formelle : age/stream (Protocole STREAM)

## 1. Contexte & Doctrine
Le protocole STREAM, tel que spécifié pour le format `age`, découpe un flux clair en segments de $64 \text{ KiB}$ chiffrés indépendamment via ChaCha20-Poly1305. 
La nonce de chaque segment ($12$ octets) est construite ainsi :
- 11 octets : compteur de segment (grand-boutien, commençant à 0).
- 1 octet : drapeau terminal (`0x01` pour le dernier segment, `0x00` pour les autres).

La faille de troncature consiste, pour un attaquant, à interrompre le flux prématurément. Si l'automate de déchiffrement ne vérifie pas la présence stricte du drapeau `0x01` lors de l'atteinte de l'EOF, un fichier tronqué peut être validé à tort.

## 2. Vecteurs d'Attaque & Épreuves de Torture

### Épreuve A : Coupure Abrupte (Truncation Flaw)
- **Scénario :** Délivrer les segments $0 \dots N-1$ (avec drapeau `0x00`), puis signaler un EOF prématuré sans jamais fournir le segment terminal (drapeau `0x01`).
- **Attente :** L'automate de déchiffrement doit lever une erreur d'intégrité (ex. `ErrTruncated` ou `io.ErrUnexpectedEOF`). La validation d'un tel flux est une compromission critique.

### Épreuve B : Rejeu et Permutation
- **Scénario :** Intercepter le segment $k$ et le rejouer à la position $j$ ($k \neq j$).
- **Attente :** Échec systématique de l'authentification Poly1305. Le compteur de segment encodé dans la nonce garantit la résistance au rejeu et à la réorganisation.

### Épreuve C : Extension Illégale Post-Clôture
- **Scénario :** Fournir un flux valide se terminant par un segment marqué `0x01`, puis injecter un segment supplémentaire (forgé ou rejoué) après cette clôture logique.
- **Attente :** L'automate de déchiffrement doit être verrouillé sitôt le drapeau `0x01` consommé. Tout appel subséquent à `Read()` doit retourner `EOF` ou une erreur stricte de violation de protocole, interdisant la désérialisation du reliquat.

### Épreuve D : Balayage Continu des Longueurs Critiques
- **Scénario :** Soumettre des charges utiles de tailles encadrant les frontières des blocs ChaCha20 et Poly1305, notamment $N \times 64 \pm 1$ ($63, 64, 65, 127, 128, 129$ octets).
- **Attente :** Validation absolue du déchiffrement partiel et de la clôture sur ces cas limites pour garantir l'absence de fuite hors limites (Out-Of-Bounds) dans les implémentations SIMD.
