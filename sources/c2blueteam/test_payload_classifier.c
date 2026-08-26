/*
 * test_payload_classifier.c — Banc d'épreuve et qualification de la classification conjointe
 * et de la détection d'alphabets restreints (ARCHTIME C99).
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <math.h>
#include <assert.h>
#include "c2blueteam.h"

// Générateur pseudo-aléatoire déterministe LCG
static inline uint32_t lcg_next(uint64_t *state) {
    *state = *state * 6364136223846793005ULL + 1442695040888963407ULL;
    return (uint32_t)(*state >> 32);
}

static void test_null_and_boundary_cases(void) {
    printf("[1/5] Cas limites, tampons nuls et échantillons courts... ");
    c2bt_entropy_profile_t prof;

    // NULL et taille 0
    assert(c2bt_profile_payload(NULL, 0, &prof) == 0);
    assert(prof.payload_class == C2BT_PAYLOAD_CLASS_UNKNOWN);
    assert(prof.entropy_q8 == 0);

    assert(c2bt_profile_payload(NULL, 100, &prof) == 0);
    assert(prof.payload_class == C2BT_PAYLOAD_CLASS_UNKNOWN);

    uint8_t dummy[16] = {0};
    assert(c2bt_profile_payload(dummy, 0, &prof) == 0);
    assert(prof.payload_class == C2BT_PAYLOAD_CLASS_UNKNOWN);

    // Échantillons courts (< 8 octets)
    const char *short_txt = "abc";
    assert(c2bt_profile_payload((const uint8_t *)short_txt, strlen(short_txt), &prof) == 0);
    assert(prof.payload_class == C2BT_PAYLOAD_CLASS_PROSE);

    uint8_t short_bin[4] = {0x00, 0xFF, 0x12, 0x80};
    assert(c2bt_profile_payload(short_bin, 4, &prof) == 0);
    assert(prof.payload_class == C2BT_PAYLOAD_CLASS_UNKNOWN);

    // Tampon uniforme
    uint8_t uniform_buf[128];
    memset(uniform_buf, 'A', sizeof(uniform_buf));
    assert(c2bt_profile_payload(uniform_buf, sizeof(uniform_buf), &prof) == 0);
    assert(prof.entropy_q8 == 0);
    assert(prof.distinct_count == 1);

    printf("OK\n");
}

static void test_hex_alphabet_classification(void) {
    printf("[2/5] Obfuscation hexadécimale (Base16 / Shellcodes / Condensats)... ");
    uint64_t rng = 0x123456789ABCDEF0ULL;

    const char hex_digits[] = "0123456789abcdefABCDEF";
    const size_t test_lens[] = {16, 32, 48, 64, 128, 256, 512, 1024, 4096};

    for (size_t s = 0; s < sizeof(test_lens)/sizeof(test_lens[0]); s++) {
        size_t n = test_lens[s];
        uint8_t *buf = (uint8_t *)malloc(n);

        // Remplissage pseudo-aléatoire sur l'alphabet hexadécimal
        for (size_t i = 0; i < n; i++) {
            buf[i] = (uint8_t)hex_digits[lcg_next(&rng) % 16]; // 16 symboles minuscules
        }

        c2bt_entropy_profile_t prof;
        int res = c2bt_profile_payload(buf, n, &prof);
        assert(res == 0);

        assert(prof.payload_class == C2BT_PAYLOAD_CLASS_HEX);
        assert((prof.char_mask & IS_CONTROL) == 0);
        assert((prof.char_mask & IS_HIGH_BYTE) == 0);
        assert((prof.char_mask & IS_HEX) != 0);
        assert(prof.entropy_q8 >= 768 && prof.entropy_q8 <= 1024); // Entre 3.0 et 4.0 b/o
        assert(prof.distinct_count <= 22);

        // Test de la fonction rapide c2bt_classify_payload
        uint32_t ent_q8 = 0, mask = 0;
        uint16_t cls = c2bt_classify_payload(buf, n, &ent_q8, &mask);
        assert(cls == C2BT_PAYLOAD_CLASS_HEX);
        assert(ent_q8 == prof.entropy_q8);
        assert(mask == prof.char_mask);

        free(buf);
    }

    // Cas concret 1 : Hash SHA-256
    const char *sha256_sample = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
    c2bt_entropy_profile_t prof_sha;
    c2bt_profile_payload((const uint8_t *)sha256_sample, strlen(sha256_sample), &prof_sha);
    printf("prof_sha.entropy_q8 = %u, class = %u\n", prof_sha.entropy_q8, prof_sha.payload_class);
    assert(prof_sha.payload_class == C2BT_PAYLOAD_CLASS_HEX);
    assert(prof_sha.entropy_q8 >= 900 && prof_sha.entropy_q8 <= 1024);

    // Cas concret 2 : Shellcode x86_64 hexadécimal
    const char *shellcode_hex = "4831c05048bb2f62696e2f2f7368534889e7504889e2534889e64883c03b0f05";
    c2bt_entropy_profile_t prof_sc;
    c2bt_profile_payload((const uint8_t *)shellcode_hex, strlen(shellcode_hex), &prof_sc);
    assert(prof_sc.payload_class == C2BT_PAYLOAD_CLASS_HEX);

    // Contre-épreuve : Insertion d'un caractère non-hex 'g' -> disqualification immédiate de HEX
    char corrupt_hex[65];
    strcpy(corrupt_hex, sha256_sample);
    corrupt_hex[30] = 'g'; // 'g' n'est pas hex
    c2bt_entropy_profile_t prof_corrupt;
    c2bt_profile_payload((const uint8_t *)corrupt_hex, strlen(corrupt_hex), &prof_corrupt);
    assert(prof_corrupt.payload_class != C2BT_PAYLOAD_CLASS_HEX);

    printf("OK\n");
}

static void test_base64_alphabet_classification(void) {
    printf("[3/5] Charges utiles Base64 et Base64URL... ");
    uint64_t rng = 0xCAFEBABE1337BEEFULL;

    const char b64_table[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    const size_t test_lens[] = {24, 32, 48, 64, 128, 256, 512, 1024, 4096};

    for (size_t s = 0; s < sizeof(test_lens)/sizeof(test_lens[0]); s++) {
        size_t n = test_lens[s];
        uint8_t *buf = (uint8_t *)malloc(n);

        // Remplissage pseudo-aléatoire uniforme sur les 64 symboles Base64
        for (size_t i = 0; i < n; i++) {
            buf[i] = (uint8_t)b64_table[lcg_next(&rng) % 64];
        }

        c2bt_entropy_profile_t prof;
        int res = c2bt_profile_payload(buf, n, &prof);
        assert(res == 0);
        if (n >= 48) {
            assert(prof.payload_class == C2BT_PAYLOAD_CLASS_BASE64);
            assert(prof.entropy_q8 >= 1200 && prof.entropy_q8 <= 1536); // Entre 4.7 et 6.0 b/o
        } else {
            assert(prof.payload_class == C2BT_PAYLOAD_CLASS_PROSE || prof.payload_class == C2BT_PAYLOAD_CLASS_UNKNOWN);
        }
        assert((prof.char_mask & IS_CONTROL) == 0);
        assert((prof.char_mask & IS_HIGH_BYTE) == 0);
        assert((prof.char_mask & IS_BASE64) != 0);
        assert(prof.distinct_count >= 8);

        free(buf);
    }

    // Cas concret 1 : Clé privée / Certificat b64
    const char *b64_payload = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzV56t8nS8zRzV1s8u8L5K"
                              "A8L9p2Z3x4y5w6v7u8t9s0r1q2p3o4n5m6l7k8j9i0h1g2f3e4d5c6b7a8Z9Y8X7"
                              "W6V5U4T3S2R1Q0P9O8N7M6L5K4J3I2H1G0F9E8D7C6B5A4z3y2x1w0v9u8t7s6r5=";
    c2bt_entropy_profile_t prof_b64;
    c2bt_profile_payload((const uint8_t *)b64_payload, strlen(b64_payload), &prof_b64);
    assert(prof_b64.payload_class == C2BT_PAYLOAD_CLASS_BASE64);
    assert(prof_b64.entropy_q8 >= 1400); // Proche de 5.8 b/o

    // Cas concret 2 : Token JWT brut en Base64URL (sans JSON décodé)
    const char *jwt_raw = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvZSJ9.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c";
    // Un JWT complet avec des points '.' (séparateurs) : comme il a des points, cnt_prose_punct > 0
    c2bt_entropy_profile_t prof_jwt;
    c2bt_profile_payload((const uint8_t *)jwt_raw, strlen(jwt_raw), &prof_jwt);
    assert(prof_jwt.payload_class == C2BT_PAYLOAD_CLASS_JWT);

    // Morceau JWT Base64URL pur (sans point)
    const char *b64url_chunk = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9-eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvZSJ9_sample";
    c2bt_entropy_profile_t prof_b64url;
    c2bt_profile_payload((const uint8_t *)b64url_chunk, strlen(b64url_chunk), &prof_b64url);
    assert(prof_b64url.payload_class == C2BT_PAYLOAD_CLASS_BASE64);

    // Cas concret 3 : Base64URL binaire pur
    const char *b64url_raw = "U2FsdGVkX1-vupppZksvRf5pq5g5XjFRIipRkw-_X1gA8H7k45o=";
    c2bt_entropy_profile_t prof_b64url_raw;
    c2bt_profile_payload((const uint8_t *)b64url_raw, strlen(b64url_raw), &prof_b64url_raw);
    assert(prof_b64url_raw.payload_class == C2BT_PAYLOAD_CLASS_BASE64);

    printf("OK\n");
}

static void test_crypto_compressed_classification(void) {
    printf("[4/5] Chiffrement binaire et compression dense (>= 7.5 b/o)... ");
    uint64_t rng = 0x9876543210FEDCBAULL;

    const size_t test_lens[] = {256, 512, 1024, 4096, 65536, 1048576};

    for (size_t s = 0; s < sizeof(test_lens)/sizeof(test_lens[0]); s++) {
        size_t n = test_lens[s];
        uint8_t *buf = (uint8_t *)malloc(n);

        // Bruit pseudo-aléatoire complet sur 256 octets (émule flux AES-GCM / ChaCha20)
        for (size_t i = 0; i < n; i++) {
            buf[i] = (uint8_t)(lcg_next(&rng) & 0xFF);
        }

        c2bt_entropy_profile_t prof;
        assert(c2bt_profile_payload(buf, n, &prof) == 0);
        assert(prof.payload_class == C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED);
        assert(prof.entropy_q8 >= 1800);
        assert((prof.char_mask & IS_HIGH_BYTE) != 0);
        free(buf);
    }

    // Épreuve Proposition 1 (Rapport 08) : Chiffrement court (32 octets, 96 octets)
    const size_t short_crypto_lens[] = {16, 32, 48, 64, 96, 128, 192};
    for (size_t s = 0; s < sizeof(short_crypto_lens)/sizeof(short_crypto_lens[0]); s++) {
        size_t n = short_crypto_lens[s];
        uint8_t *buf = (uint8_t *)malloc(n);
        for (size_t i = 0; i < n; i++) {
            buf[i] = (uint8_t)(lcg_next(&rng) & 0xFF);
        }
        c2bt_entropy_profile_t prof;
        assert(c2bt_profile_payload(buf, n, &prof) == 0);
        assert(prof.payload_class == C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED);
        assert((prof.char_mask & IS_HIGH_BYTE) != 0);
        free(buf);
    }

    // Épreuve Proposition 2 & Retex : Jeton JWT structuré
    const char *jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."
                            "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ."
                            "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c";
    c2bt_entropy_profile_t prof_jwt;
    assert(c2bt_profile_payload((const uint8_t *)jwt_token, strlen(jwt_token), &prof_jwt) == 0);
    assert(prof_jwt.payload_class == C2BT_PAYLOAD_CLASS_JWT);

    // Épreuve Proposition 2 : Identifiant de session ordinaire à tirets (ne doit PAS être BASE64)
    const char *session_id = "identifiant-de-session-abc-123-def-456-ghi-789-jkl-012";
    c2bt_entropy_profile_t prof_sess;
    assert(c2bt_profile_payload((const uint8_t *)session_id, strlen(session_id), &prof_sess) == 0);
    assert(prof_sess.payload_class == C2BT_PAYLOAD_CLASS_PROSE || prof_sess.payload_class == C2BT_PAYLOAD_CLASS_UNKNOWN);
    assert(prof_sess.payload_class != C2BT_PAYLOAD_CLASS_BASE64);

    printf("OK\n");
}

static void test_natural_text_prose_classification(void) {
    printf("[5/5] Texte naturel, documentation et code source (Zéro Faux Positif)... ");

    // 1. Prose littéraire anglaise
    const char *prose_en = "Information theory studies the transmission, processing, utilization, and extraction of information. "
                           "Abstractly, information can be thought of as the resolution of uncertainty. The fundamental "
                           "problem of communication is that of reproducing at one point either exactly or approximately "
                           "a message selected at another point.";
    c2bt_entropy_profile_t prof_en;
    c2bt_profile_payload((const uint8_t *)prose_en, strlen(prose_en), &prof_en);
    printf("prof_en: len=%zu, class=%u, ent=%u, mask=0x%x, distinct=%u\n",
           prof_en.len, prof_en.payload_class, prof_en.entropy_q8, prof_en.char_mask, prof_en.distinct_count);
    assert(prof_en.payload_class == C2BT_PAYLOAD_CLASS_PROSE);
    assert(prof_en.entropy_q8 < 1250); // ~4.3 b/o
    assert(prof_en.payload_class != C2BT_PAYLOAD_CLASS_HEX);
    assert(prof_en.payload_class != C2BT_PAYLOAD_CLASS_BASE64);
    assert(prof_en.payload_class != C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED);

    // 2. Prose française avec ponctuation et espaces
    const char *prose_fr = "Le principe d'architecture ARCHTIME impose que toute décision de métrologie "
                           "soit déterministe, calculée en virgule fixe Q8.8, et sans aucune allocation dynamique au runtime.";
    c2bt_entropy_profile_t prof_fr;
    c2bt_profile_payload((const uint8_t *)prose_fr, strlen(prose_fr), &prof_fr);
    assert(prof_fr.payload_class == C2BT_PAYLOAD_CLASS_PROSE);

    // 3. Code source C99
    const char *c_code = "int c2bt_eval_rules_batch(const probe_event_t *in_events, probe_event_t *out_events, int count) {\n"
                         "    if (!in_events || !out_events || count <= 0) return 0;\n"
                         "    for (int i = 0; i < count; i++) { /* traitement */ }\n"
                         "    return count;\n"
                         "}\n";
    c2bt_entropy_profile_t prof_c;
    c2bt_profile_payload((const uint8_t *)c_code, strlen(c_code), &prof_c);
    assert(prof_c.payload_class == C2BT_PAYLOAD_CLASS_PROSE);

    // 4. Commandes shell et requêtes JSON légitimes
    const char *json_query = "{\"method\":\"tools/call\",\"params\":{\"name\":\"codeindex_query\",\"arguments\":{\"query\":\"c2bt\"}}}";
    c2bt_entropy_profile_t prof_json;
    c2bt_profile_payload((const uint8_t *)json_query, strlen(json_query), &prof_json);
    assert(prof_json.payload_class == C2BT_PAYLOAD_CLASS_PROSE);

    printf("OK\n");
}

static void test_massive_monte_carlo_classification(void) {
    printf("[6/6] Torture Monte-Carlo massive (100 000 vecteurs synthétiques multi-classes)... ");
    fflush(stdout);

    uint64_t rng = 0xFEEDBEEF55AA1234ULL;
    const char hex_chars[] = "0123456789abcdef";
    const char b64_chars[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    const char prose_alphabet[] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 .,;:!?'\"-()\n";

    uint8_t *buffer = (uint8_t *)malloc(65536);
    assert(buffer != NULL);

    long total_monte_carlo = 0;

    for (int iter = 0; iter < 100000; iter++) {
        int category = iter % 4;
        c2bt_entropy_profile_t prof;

        if (category == 0) {
            // 1. Vecteur Hexadécimal
            size_t len = 32 + (lcg_next(&rng) % 1024);
            for (size_t i = 0; i < len; i++) {
                buffer[i] = (uint8_t)hex_chars[lcg_next(&rng) % 16];
            }
            int r = c2bt_profile_payload(buffer, len, &prof);
            assert(r == 0);
            if (prof.payload_class != C2BT_PAYLOAD_CLASS_HEX) {
                fprintf(stderr, "\nÉchec classification HEX iter=%d len=%zu ent=%u mask=0x%x distinct=%u cls=%u\n",
                        iter, len, prof.entropy_q8, prof.char_mask, prof.distinct_count, prof.payload_class);
                assert(prof.payload_class == C2BT_PAYLOAD_CLASS_HEX);
            }
        } else if (category == 1) {
            // 2. Vecteur Base64 (avec garantie de symboles non-hex)
            size_t len = 48 + (lcg_next(&rng) % 1024);
            for (size_t i = 0; i < len; i++) {
                buffer[i] = (uint8_t)b64_chars[lcg_next(&rng) % 64];
            }
            int r = c2bt_profile_payload(buffer, len, &prof);
            assert(r == 0);
            if (prof.payload_class != C2BT_PAYLOAD_CLASS_BASE64) {
                fprintf(stderr, "\nÉchec classification BASE64 iter=%d len=%zu ent=%u mask=0x%x distinct=%u cls=%u\n",
                        iter, len, prof.entropy_q8, prof.char_mask, prof.distinct_count, prof.payload_class);
                assert(prof.payload_class == C2BT_PAYLOAD_CLASS_BASE64);
            }
        } else if (category == 2) {
            // 3. Vecteur Chiffrement Binaire / Bruit Blanc
            size_t len = 512 + (lcg_next(&rng) % 2048);
            for (size_t i = 0; i < len; i++) {
                buffer[i] = (uint8_t)(lcg_next(&rng) & 0xFF);
            }
            int r = c2bt_profile_payload(buffer, len, &prof);
            assert(r == 0);
            if (prof.payload_class != C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED) {
                fprintf(stderr, "\nÉchec classification CRYPTO iter=%d len=%zu ent=%u mask=0x%x distinct=%u cls=%u\n",
                        iter, len, prof.entropy_q8, prof.char_mask, prof.distinct_count, prof.payload_class);
                assert(prof.payload_class == C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED);
            }
        } else {
            // 4. Vecteur Prose synthétique (avec espaces réguliers et ponctuation)
            size_t len = 64 + (lcg_next(&rng) % 1024);
            size_t prose_len = strlen(prose_alphabet);
            for (size_t i = 0; i < len; i++) {
                uint32_t roll = lcg_next(&rng) % 100;
                if (roll < 20) buffer[i] = ' ';
                else if (roll < 25) buffer[i] = '.';
                else if (roll < 28) buffer[i] = ',';
                else buffer[i] = (uint8_t)prose_alphabet[lcg_next(&rng) % prose_len];
            }
            int r = c2bt_profile_payload(buffer, len, &prof);
            assert(r == 0);
            if (prof.payload_class != C2BT_PAYLOAD_CLASS_PROSE) {
                fprintf(stderr, "\nÉchec classification PROSE iter=%d len=%zu ent=%u mask=0x%x distinct=%u cls=%u\n",
                        iter, len, prof.entropy_q8, prof.char_mask, prof.distinct_count, prof.payload_class);
                assert(prof.payload_class == C2BT_PAYLOAD_CLASS_PROSE);
            }
        }
        total_monte_carlo++;
    }

    free(buffer);
    printf("OK (%ld vecteurs validés sans le moindre faux positif)\n", total_monte_carlo);
}

int main(void) {
    printf("=================================================================\n");
    printf("   ÉPREUVE ORACLE C : CLASSIFICATION MULTI-CLASSES ARCHTIME\n");
    printf("=================================================================\n\n");

    test_null_and_boundary_cases();
    test_hex_alphabet_classification();
    test_base64_alphabet_classification();
    test_crypto_compressed_classification();
    test_natural_text_prose_classification();
    test_massive_monte_carlo_classification();

    printf("\n=================================================================\n");
    printf("   TOUS LES VECTEURS DE CLASSIFICATION SONT 100%% VALIDÉS\n");
    printf("=================================================================\n");
    return 0;
}
