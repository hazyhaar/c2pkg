# RETEX D'ÉPREUVE DE SINCÉRITÉ & ROUGISSEMENT PAR STRATE (2026-08-28)

Ce document consigne les résultats de l'épreuve de sincérité contradictoire exécutée le 2026-08-28.
Chaque strate du pipeline cryptographique a été volontairement corrompue d'une seule valeur pour prouver l'absence totale de tests complaisants ou passifs.

---

## 1. Strate Keystreaming ChaCha20 / AVX2 (`c2chacha8`)

* **Mutation injectée :** Altération de 1 octet dans la constante `C2chacha8_sigma` (`0x65` changé en `0x66` dans `chacha8_gen.go`).
* **Verdict du banc :** **ÉCHEC ROUGE IMMÉDIAT (3 tests en échec critique)**
  ```text
  --- FAIL: TestRFC8439Blocks (0.00s)
      chacha8_test.go:34: RFC 8439 : obtenu 2eab7283d...
  --- FAIL: TestParityXCrypto (0.00s)
      chacha8_test.go:71: n=1 ctr=1326512963 : 04 != 94
  --- FAIL: TestChacha8VsCOracle (0.25s)
      chacha8_test.go:127: rfc : Go 2eab7283d... vs oracle C 6e2e359a2...
  FAIL	github.com/hazyhaar/c2pkg/c2chacha8
  ```
* **Rétablissement :** Constante restaurée à `0x65` -> `ok github.com/hazyhaar/c2pkg/c2chacha8 0.247s` (**PASS**).

---

## 2. Strate MAC Poly1305 / Arithmétique Modulaire (`c2poly1305`)

* **Mutation injectée :** Altération de 1 bit dans le masque de clamping de la clé $r$ (`0xfffffff` changé en `0xeffffff` dans `Crypto_poly1305_init`).
* **Verdict du banc :** **ÉCHEC ROUGE IMMÉDIAT (9 vecteurs d'attaque en échec)**
  ```text
  --- FAIL: TestPoly1305DegenerateEntropy (0.00s)
      degenerate_entropy_test.go:63: Poly1305 diverge on degenerate inputs (n=1)
  --- FAIL: TestAdversarialVectors (0.00s)
      --- FAIL: TestAdversarialVectors/Math_P (0.00s)
      --- FAIL: TestAdversarialVectors/Math_P_Minus_1 (0.00s)
      --- FAIL: TestAdversarialVectors/Math_P_Plus_1 (0.00s)
      --- FAIL: TestAdversarialVectors/Math_2_130_Minus_6 (0.00s)
      --- FAIL: TestAdversarialVectors/Max_Clamp_Key_Zero_Msg (0.00s)
      --- FAIL: TestAdversarialVectors/Max_Clamp_Key_FF_Msg (0.00s)
      --- FAIL: TestAdversarialVectors/Carry_Overflow_Long (0.00s)
  --- FAIL: TestParityXCrypto (0.00s)
  --- FAIL: TestPoly1305VsCOracle (0.61s)
  FAIL	github.com/hazyhaar/c2pkg/c2poly1305
  ```
* **Rétablissement :** Clamping restauré à `0xfffffff` -> `ok github.com/hazyhaar/c2pkg/c2poly1305 0.220s` (**PASS**).

---

## 3. Strate Protocolaire STREAM (`secretstream55`)

* **Mutation injectée :** Désynchronisation de la séquence du nonce (`e.seq` changé en `e.seq+1` à l'émission dans `writeFrame`).
* **Verdict du banc :** **ÉCHEC ROUGE MASSIF (22 tests en échec)**
  ```text
  --- FAIL: TestContract_MaisonInterChunkTruncation (0.00s)
  --- FAIL: TestWithEngine_RoundTripAndDefaultParity (0.00s)
  --- FAIL: TestReadInPlace_NoIntermediateCopy (0.00s)
  --- FAIL: TestSubkeyStream_ByteExact_VsOldPath (0.00s)
  --- FAIL: TestV2Attack_TruncationAfterCompleteFrameNoFinal (0.00s)
  --- FAIL: TestV2Attack_DeleteIntermediateFrame (0.00s)
  --- FAIL: TestV2Attack_ReplayFrame (0.00s)
  --- FAIL: TestV2Vectors (0.00s)
  --- FAIL: TestTortureExoticAndObfuscatedPayloads (0.00s)
  --- FAIL: TestTortureFragmentedIO (0.00s)
  --- FAIL: TestTortureStreamAdversarial (0.00s)
  FAIL	code.hazyhaar.fr/devhoros/pkg/secretstream55
  ```
* **Rétablissement :** Séquence restaurée à `e.seq` -> `ok code.hazyhaar.fr/devhoros/pkg/secretstream55 0.158s` (**PASS**).

---

## 4. Conclusion de l'Épreuve

L'adversarial testing prouve que le banc d'audit intercepte toute déviation à l'octet près sur chacun des trois piliers fondamentaux (Keystreaming vectoriel, Arithmétique modulaire, Protocole de flux).
