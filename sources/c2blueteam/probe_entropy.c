/*
 * probe_entropy.c — Calculateur d'entropie Shannon universel ARCHTIME Q8.8
 * Zéro FPU, zéro allocation, précision continue de 1 octet à 4 Go.
 */

#include "c2blueteam.h"
#include <string.h>

/*
 * Table ARCHTIME 1 : f(c) = round(256 * c * log2(c)) pour c de 0 à 256.
 * Élimine la multiplication au runtime dans la boucle des 256 fréquences.
 */
static const uint32_t c_log2_c_table[257] = {
           0,        0,      512,     1217,     2048,     2972,     3971,     5031,
        6144,     7304,     8504,     9742,    11013,    12315,    13646,    15002,
       16384,    17789,    19215,    20662,    22128,    23613,    25116,    26635,
       28170,    29721,    31286,    32866,    34459,    36066,    37685,    39317,
       40960,    42615,    44281,    45958,    47646,    49344,    51052,    52769,
       54497,    56233,    57978,    59732,    61495,    63266,    65045,    66833,
       68628,    70431,    72241,    74059,    75884,    77716,    79556,    81402,
       83254,    85114,    86979,    88851,    90730,    92614,    94505,    96402,
       98304,   100212,   102126,   104046,   105971,   107901,   109837,   111778,
      113724,   115675,   117632,   119593,   121560,   123531,   125507,   127488,
      129473,   131463,   133458,   135457,   137460,   139468,   141481,   143497,
      145518,   147543,   149572,   151606,   153643,   155684,   157730,   159779,
      161832,   163889,   165950,   168014,   170083,   172155,   174230,   176310,
      178393,   180479,   182569,   184662,   186759,   188859,   190963,   193070,
      195180,   197294,   199411,   201531,   203655,   205781,   207911,   210044,
      212180,   214319,   216461,   218606,   220754,   222905,   225059,   227216,
      229376,   231539,   233704,   235873,   238044,   240218,   242395,   244575,
      246757,   248942,   251130,   253320,   255514,   257709,   259908,   262109,
      264312,   266518,   268727,   270938,   273152,   275368,   277587,   279808,
      282031,   284257,   286486,   288717,   290950,   293185,   295423,   297664,
      299906,   302151,   304398,   306648,   308899,   311153,   313410,   315668,
      317929,   320192,   322457,   324724,   326993,   329265,   331538,   333814,
      336092,   338372,   340654,   342938,   345225,   347513,   349803,   352096,
      354390,   356686,   358985,   361285,   363587,   365892,   368198,   370506,
      372816,   375128,   377442,   379758,   382076,   384395,   386717,   389040,
      391365,   393693,   396021,   398352,   400685,   403019,   405355,   407693,
      410033,   412375,   414718,   417063,   419410,   421758,   424109,   426461,
      428814,   431170,   433527,   435886,   438246,   440608,   442972,   445338,
      447705,   450074,   452444,   454816,   457190,   459565,   461942,   464321,
      466701,   469083,   471466,   473851,   476238,   478626,   481015,   483407,
      485799,   488194,   490589,   492987,   495386,   497786,   500188,   502591,
      504996,   507402,   509810,   512220,   514630,   517043,   519456,   521871,
      524288
};

static const uint8_t log2_mantissa_q8[256] = {
       0,    1,    3,    4,    6,    7,    9,   10,   11,   13,   14,   16,   17,   18,   20,   21,
      22,   24,   25,   26,   28,   29,   30,   32,   33,   34,   36,   37,   38,   40,   41,   42,
      44,   45,   46,   47,   49,   50,   51,   52,   54,   55,   56,   57,   59,   60,   61,   62,
      63,   65,   66,   67,   68,   69,   71,   72,   73,   74,   75,   77,   78,   79,   80,   81,
      82,   84,   85,   86,   87,   88,   89,   90,   92,   93,   94,   95,   96,   97,   98,   99,
     100,  102,  103,  104,  105,  106,  107,  108,  109,  110,  111,  112,  113,  114,  116,  117,
     118,  119,  120,  121,  122,  123,  124,  125,  126,  127,  128,  129,  130,  131,  132,  133,
     134,  135,  136,  137,  138,  139,  140,  141,  142,  143,  144,  145,  146,  147,  148,  149,
     150,  151,  152,  153,  154,  155,  155,  156,  157,  158,  159,  160,  161,  162,  163,  164,
     165,  166,  167,  168,  169,  169,  170,  171,  172,  173,  174,  175,  176,  177,  178,  178,
     179,  180,  181,  182,  183,  184,  185,  185,  186,  187,  188,  189,  190,  191,  192,  192,
     193,  194,  195,  196,  197,  198,  198,  199,  200,  201,  202,  203,  203,  204,  205,  206,
     207,  208,  208,  209,  210,  211,  212,  212,  213,  214,  215,  216,  216,  217,  218,  219,
     220,  220,  221,  222,  223,  224,  224,  225,  226,  227,  228,  228,  229,  230,  231,  231,
     232,  233,  234,  234,  235,  236,  237,  238,  238,  239,  240,  241,  241,  242,  243,  244,
     244,  245,  246,  247,  247,  248,  249,  249,  250,  251,  252,  252,  253,  254,  255,  255
};

static const uint8_t archtime_char_class[256] = {
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x21, 0x21, 0x08, 0x08, 0x21, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x21, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x05, 0x01, 0x05, 0x41, 0x05,
    0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x01, 0x01, 0x01, 0x05, 0x01, 0x01,
    0x01, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
    0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x01, 0x01, 0x01, 0x01, 0x05,
    0x01, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
    0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x01, 0x01, 0x01, 0x01, 0x08,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
    0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10
};

static inline uint32_t arch_log2_q8_8(uint64_t n) {
    if (n <= 1) return 0;
    if (n <= 256) {
        uint32_t val = c_log2_c_table[n];
        uint32_t denom = (uint32_t)n;
        return val / denom;
    }
    int lz = __builtin_clzll(n);
    int msb = 63 - lz;
    uint32_t mantissa_idx = (uint32_t)((n << (lz + 1)) >> 56);
    uint32_t mant = (uint32_t)log2_mantissa_q8[mantissa_idx];
    return (uint32_t)((msb << 8) + mant);
}

uint32_t c2bt_calc_entropy_8_8(const uint8_t *data, size_t len) {
    if (data == NULL || len == 0) {
        return 0;
    }

    uint32_t freq[256];
    memset(freq, 0, sizeof(freq));

    size_t i = 0;
    for (; i + 7 < len; i += 8) {
        freq[data[i]]++;
        freq[data[i+1]]++;
        freq[data[i+2]]++;
        freq[data[i+3]]++;
        freq[data[i+4]]++;
        freq[data[i+5]]++;
        freq[data[i+6]]++;
        freq[data[i+7]]++;
    }
    for (; i < len; i++) {
        freq[data[i]]++;
    }

    uint64_t sum_c_log_c = 0;
    for (int j = 0; j < 256; j++) {
        uint32_t c = freq[j];
        if (c > 0) {
            if (c <= 256) {
                sum_c_log_c += c_log2_c_table[c];
            } else {
                sum_c_log_c += (uint64_t)c * (uint64_t)arch_log2_q8_8(c);
            }
        }
    }

    uint32_t log2_len = arch_log2_q8_8((uint64_t)len);
    uint32_t avg_c_log_c = (uint32_t)(sum_c_log_c / len);

    if (log2_len >= avg_c_log_c) {
        return log2_len - avg_c_log_c;
    }
    return 0;
}

int c2bt_profile_payload(const uint8_t *data, size_t len, c2bt_entropy_profile_t *out_profile) {
    if (out_profile == NULL) {
        return -1;
    }
    memset(out_profile, 0, sizeof(*out_profile));
    out_profile->len = len;
    out_profile->payload_class = C2BT_PAYLOAD_CLASS_UNKNOWN;

    if (data == NULL || len == 0) {
        return 0;
    }

    uint32_t freq[256];
    memset(freq, 0, sizeof(freq));
    uint32_t combined_mask = 0;

    size_t i = 0;
    for (; i + 7 < len; i += 8) {
        uint8_t b0 = data[i];
        uint8_t b1 = data[i+1];
        uint8_t b2 = data[i+2];
        uint8_t b3 = data[i+3];
        uint8_t b4 = data[i+4];
        uint8_t b5 = data[i+5];
        uint8_t b6 = data[i+6];
        uint8_t b7 = data[i+7];
        freq[b0]++;
        freq[b1]++;
        freq[b2]++;
        freq[b3]++;
        freq[b4]++;
        freq[b5]++;
        freq[b6]++;
        freq[b7]++;
        combined_mask |= (uint32_t)(archtime_char_class[b0] | archtime_char_class[b1] |
                                    archtime_char_class[b2] | archtime_char_class[b3] |
                                    archtime_char_class[b4] | archtime_char_class[b5] |
                                    archtime_char_class[b6] | archtime_char_class[b7]);
    }
    for (; i < len; i++) {
        uint8_t b = data[i];
        freq[b]++;
        combined_mask |= (uint32_t)archtime_char_class[b];
    }

    out_profile->char_mask = combined_mask;

    uint64_t sum_c_log_c = 0;
    uint16_t distinct_count = 0;
    size_t cnt_hex = 0;
    size_t cnt_b64 = 0;
    size_t cnt_non_hex_b64 = 0;
    size_t cnt_printable = 0;
    size_t cnt_whitespace = 0;
    size_t cnt_prose_punct = 0;
    size_t cnt_dot = 0;
    size_t cnt_control = 0;
    size_t cnt_high = 0;

    /* Analyse des fréquences et comptage par catégorie d'alphabet (256 itérations fixes) */
    for (int i = 0; i < 256; i++) {
        uint32_t c = freq[i];
        if (c > 0) {
            distinct_count++;
            if (c <= 256) {
                sum_c_log_c += c_log2_c_table[c];
            } else {
                sum_c_log_c += (uint64_t)c * (uint64_t)arch_log2_q8_8(c);
            }

            uint8_t cls = archtime_char_class[i];
            if (cls & C2BT_CHAR_IS_HEX) {
                cnt_hex += c;
            }
            if (cls & C2BT_CHAR_IS_BASE64) {
                cnt_b64 += c;
                if (!(cls & C2BT_CHAR_IS_HEX)) {
                    cnt_non_hex_b64 += c;
                }
            }
            if (cls & C2BT_CHAR_IS_DOT) {
                cnt_dot += c;
            }
            if (cls & C2BT_CHAR_IS_PRINTABLE) {
                cnt_printable += c;
                if (!(cls & C2BT_CHAR_IS_BASE64) && !(cls & C2BT_CHAR_IS_WHITESPACE) && !(cls & C2BT_CHAR_IS_DOT)) {
                    cnt_prose_punct += c;
                }
            }
            if (cls & C2BT_CHAR_IS_WHITESPACE) {
                cnt_whitespace += c;
            }
            if (cls & C2BT_CHAR_IS_CONTROL) {
                cnt_control += c;
            }
            if (cls & C2BT_CHAR_IS_HIGH_BYTE) {
                cnt_high += c;
            }
        }
    }

    out_profile->distinct_count = distinct_count;

    uint32_t log2_len = arch_log2_q8_8((uint64_t)len);
    uint32_t avg_c_log_c = (uint32_t)(sum_c_log_c / len);
    uint32_t entropy_q8 = 0;
    if (log2_len >= avg_c_log_c) {
        entropy_q8 = log2_len - avg_c_log_c;
    }
    out_profile->entropy_q8 = entropy_q8;

    /* Décision déterministe de classification conjointe sans faux positif */
    if (len < 8) {
        if (cnt_printable == len) {
            out_profile->payload_class = C2BT_PAYLOAD_CLASS_PROSE;
        } else {
            out_profile->payload_class = C2BT_PAYLOAD_CLASS_UNKNOWN;
        }
        return 0;
    }

    /* 1. Chiffrement binaire ou compression dense */
    if (len >= 256) {
        if (entropy_q8 >= 1920 || (entropy_q8 >= 1792 && cnt_high > 0 && cnt_control > 0 && distinct_count >= 80)) {
            out_profile->payload_class = C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED;
            return 0;
        }
    } else if (len >= 16) {
        /* Proposition 1 (Rapport 08) : Régime court (16..255 octets) — Dispersion binaire >= 55% et octets de poids fort */
        if (cnt_high > 0 && ((size_t)distinct_count * 100) >= (len * 55) && cnt_control > 0) {
            out_profile->payload_class = C2BT_PAYLOAD_CLASS_CRYPTO_COMPRESSED;
            return 0;
        }
    }

    /* 2. Obfuscation hexadécimale (Base16) — Proposition 3 : Alphabet d'abord */
    if (cnt_control == 0 && cnt_high == 0 && cnt_prose_punct == 0 && cnt_non_hex_b64 == 0 && cnt_dot == 0) {
        if ((cnt_hex + cnt_whitespace) == len && cnt_hex >= (len * 90 / 100)) {
            if (distinct_count >= 4 || (len >= 16 && distinct_count >= 1)) {
                out_profile->payload_class = C2BT_PAYLOAD_CLASS_HEX;
                return 0;
            }
        }
    }

    /* 3. Jeton d'authentification structuré (JWT : Base64URL sans espace séparé par 1 à 3 points) */
    if (cnt_control == 0 && cnt_high == 0 && cnt_prose_punct == 0 && cnt_whitespace == 0 && cnt_dot >= 1 && cnt_dot <= 3 && len >= 32) {
        if ((cnt_b64 + cnt_dot) == len) {
            size_t min_distinct = (len < 64 ? len : 64) / 2;
            if (distinct_count >= min_distinct) {
                out_profile->payload_class = C2BT_PAYLOAD_CLASS_JWT;
                return 0;
            }
        }
    }

    /* 4. Charge utile encodée en Base64 / Base64URL — Proposition 2 : Plancher 48 octets et variété proportionnelle */
    if (cnt_control == 0 && cnt_high == 0 && cnt_prose_punct == 0 && cnt_dot == 0 && cnt_non_hex_b64 > 0 && len >= 48) {
        if ((cnt_b64 + cnt_whitespace) == len && cnt_b64 >= (len * 90 / 100)) {
            size_t min_distinct = (len < 64) ? (len / 2) : 32;
            if (distinct_count >= min_distinct && entropy_q8 >= 1150 && entropy_q8 <= 1550) {
                out_profile->payload_class = C2BT_PAYLOAD_CLASS_BASE64;
                return 0;
            }
        }
    }

    /* 5. Texte naturel / Prose / Documentation / Code source */
    if (cnt_control == 0 && (cnt_printable + cnt_high) == len) {
        if (entropy_q8 < 1600 || len < 48) {
            out_profile->payload_class = C2BT_PAYLOAD_CLASS_PROSE;
            return 0;
        }
    }

    /* 6. Inconnu / données binaires à basse entropie non structurées */
    out_profile->payload_class = C2BT_PAYLOAD_CLASS_UNKNOWN;
    return 0;
}

uint16_t c2bt_classify_payload(const uint8_t *data, size_t len, uint32_t *out_entropy_q8, uint32_t *out_char_mask) {
    c2bt_entropy_profile_t prof;
    c2bt_profile_payload(data, len, &prof);
    if (out_entropy_q8 != NULL) {
        *out_entropy_q8 = prof.entropy_q8;
    }
    if (out_char_mask != NULL) {
        *out_char_mask = prof.char_mask;
    }
    return prof.payload_class;
}

