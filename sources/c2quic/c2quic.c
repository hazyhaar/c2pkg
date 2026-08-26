#include "c2quic.h"

const uint32_t c2quic_vint_len_tab[4] = {1u, 2u, 4u, 8u};

uint64_t c2quic_select64(uint64_t cond, uint64_t a, uint64_t b) {
    uint64_t mask = (uint64_t)(-(int64_t)(cond != 0));
    return (a & mask) | (b & ~mask);
}

uint32_t c2quic_vint_nbytes(uint64_t value) {
    uint32_t p = 0;
    p += (uint32_t)(value >= 64ull);
    p += (uint32_t)(value >= 16384ull);
    p += (uint32_t)(value >= 1073741824ull);
    return c2quic_vint_len_tab[p];
}

uint32_t c2quic_vint_decode(const uint8_t *buf, uint32_t len, uint32_t off, c2quic_vint_t *out) {
    uint32_t first;
    uint32_t prefix;
    uint32_t n;
    uint32_t i;
    if (out == 0) {
        return C2QUIC_INVALID;
    }
    out->value = 0;
    out->nbytes = 0;
    out->status = C2QUIC_TRUNC;
    if (buf == 0 || off >= len) {
        return C2QUIC_TRUNC;
    }
    first = (uint32_t)buf[off];
    prefix = first >> 6;
    n = c2quic_vint_len_tab[prefix];
    if (len - off < n) {
        return C2QUIC_TRUNC;
    }
    out->value = (uint64_t)(buf[off] & 0x3Fu);
    for (i = 1; i < n; i++) {
        out->value = (out->value << 8) | (uint64_t)buf[off + i];
    }
    out->nbytes = n;
    out->status = C2QUIC_OK;
    return C2QUIC_OK;
}

uint32_t c2quic_vint_encode(uint64_t value, uint8_t *buf, uint32_t len, uint32_t off, c2quic_vint_t *out) {
    uint32_t n;
    uint32_t prefix;
    uint32_t i;
    uint64_t tagged;
    uint32_t shift;
    if (out == 0) {
        return C2QUIC_INVALID;
    }
    out->value = value;
    out->nbytes = 0;
    out->status = C2QUIC_OVERFLOW;
    if (value >= 4611686018427387904ull) {
        return C2QUIC_OVERFLOW;
    }
    n = c2quic_vint_nbytes(value);
    prefix = 0;
    prefix = (uint32_t)c2quic_select64(n == 2, 1, prefix);
    prefix = (uint32_t)c2quic_select64(n == 4, 2, prefix);
    prefix = (uint32_t)c2quic_select64(n == 8, 3, prefix);
    if (buf == 0) {
        out->status = C2QUIC_TRUNC;
        return C2QUIC_TRUNC;
    }
    if (off > len) {
        out->status = C2QUIC_TRUNC;
        return C2QUIC_TRUNC;
    }
    if (len - off < n) {
        out->status = C2QUIC_TRUNC;
        return C2QUIC_TRUNC;
    }
    shift = (n << 3) - 2;
    tagged = value | ((uint64_t)prefix << shift);
    for (i = 0; i < n; i++) {
        uint32_t s = (n - 1u - i) << 3;
        buf[off + i] = (uint8_t)((tagged >> s) & 0xFFull);
    }
    out->nbytes = n;
    out->status = C2QUIC_OK;
    return C2QUIC_OK;
}

void c2quic_vint_decode_batch8_scalar(const uint8_t *buf, uint32_t len, c2quic_batch8_t *b) {
    int i;
    c2quic_vint_t r;
    if (b == 0) {
        return;
    }
    for (i = 0; i < 8; i++) {
        r.value = 0;
        r.nbytes = 0;
        r.status = C2QUIC_INVALID;
        c2quic_vint_decode(buf, len, b->offs[i], &r);
        b->vals[i] = r.value;
        b->nbytes[i] = r.nbytes;
        b->status[i] = r.status;
    }
}

void c2quic_vint_decode_batch8(const uint8_t *buf, uint32_t len, c2quic_batch8_t *b) {
    c2quic_vint_decode_batch8_scalar(buf, len, b);
}

uint32_t c2quic_hp_sample(const uint8_t *pkt, uint32_t len, uint32_t pn_off, c2quic_hp_io_t *io) {
    uint32_t src;
    uint32_t i;
    if (pkt == 0 || io == 0) {
        return C2QUIC_INVALID;
    }
    if (pn_off > 0xFFFFFFFBu) {
        return C2QUIC_TRUNC;
    }
    src = pn_off + 4u;
    if (len < src || len - src < 16u) {
        return C2QUIC_TRUNC;
    }
    for (i = 0; i < 16; i++) {
        io->sample[i] = pkt[src + i];
    }
    return C2QUIC_OK;
}

uint32_t c2quic_hp_apply(uint8_t *pkt, uint32_t len, uint32_t pn_off, uint32_t pn_len, c2quic_hp_io_t *io) {
    uint32_t bits;
    uint32_t i;
    if (pkt == 0 || io == 0 || len == 0) {
        return C2QUIC_INVALID;
    }
    if (pn_len < 1u || pn_len > 4u) {
        return C2QUIC_INVALID;
    }
    if (pn_off >= len || len - pn_off < pn_len) {
        return C2QUIC_TRUNC;
    }
    bits = (uint32_t)c2quic_select64((pkt[0] & 0x80u) != 0, 0x0Fu, 0x1Fu);
    pkt[0] = (uint8_t)(pkt[0] ^ (io->mask[0] & (uint8_t)bits));
    for (i = 0; i < pn_len; i++) {
        pkt[pn_off + i] = (uint8_t)(pkt[pn_off + i] ^ io->mask[1u + i]);
    }
    return C2QUIC_OK;
}

uint32_t c2quic_hp_remove(uint8_t *pkt, uint32_t len, uint32_t pn_off, c2quic_hp_io_t *io) {
    uint32_t bits;
    uint32_t pn_len;
    uint32_t i;
    uint8_t first;
    if (pkt == 0 || io == 0 || len == 0) {
        return C2QUIC_INVALID;
    }
    bits = (uint32_t)c2quic_select64((pkt[0] & 0x80u) != 0, 0x0Fu, 0x1Fu);
    first = (uint8_t)(pkt[0] ^ (io->mask[0] & (uint8_t)bits));
    pn_len = (uint32_t)((first & 0x03u) + 1u);
    if (pn_off >= len || len - pn_off < pn_len) {
        return C2QUIC_TRUNC;
    }
    pkt[0] = first;
    for (i = 0; i < pn_len; i++) {
        pkt[pn_off + i] = (uint8_t)(pkt[pn_off + i] ^ io->mask[1u + i]);
    }
    io->pn_len = pn_len;
    return C2QUIC_OK;
}

uint32_t c2quic_hp_read_pn(const uint8_t *pkt, uint32_t len, uint32_t pn_off, uint32_t pn_len, c2quic_hp_io_t *io) {
    uint64_t v;
    uint32_t i;
    if (io == 0) {
        return C2QUIC_INVALID;
    }
    io->pn = 0;
    if (pkt == 0 || pn_len < 1u || pn_len > 4u) {
        return C2QUIC_INVALID;
    }
    if (pn_off > len || len - pn_off < pn_len) {
        return C2QUIC_TRUNC;
    }
    v = 0;
    for (i = 0; i < pn_len; i++) {
        v = (v << 8) | (uint64_t)pkt[pn_off + i];
    }
    io->pn = v;
    return C2QUIC_OK;
}

uint32_t c2quic_eat_vint(const uint8_t *buf, uint32_t len, c2quic_cursor_t *cur) {
    c2quic_vint_t r;
    uint32_t st;
    if (cur == 0) {
        return C2QUIC_INVALID;
    }
    st = c2quic_vint_decode(buf, len, cur->off, &r);
    cur->status = st;
    cur->value = r.value;
    cur->nbytes = r.nbytes;
    if (st != C2QUIC_OK) {
        return st;
    }
    cur->off = cur->off + r.nbytes;
    return C2QUIC_OK;
}

uint32_t c2quic_unpack_frames(const uint8_t *buf, uint32_t len, c2quic_frame_slab_t *slab) {
    c2quic_cursor_t cur;
    uint32_t n;
    uint32_t typ;
    uint32_t st;
    uint64_t range_count;
    uint64_t i;
    if (slab == 0) {
        return C2QUIC_INVALID;
    }
    slab->n = 0;
    slab->status = C2QUIC_OK;
    if (buf == 0) {
        slab->status = C2QUIC_INVALID;
        return C2QUIC_INVALID;
    }
    cur.off = 0;
    cur.status = 0;
    cur.value = 0;
    cur.nbytes = 0;
    cur.pad = 0;
    n = 0;
    while (cur.off < len) {
        if (buf[cur.off] == 0) {
            cur.off = cur.off + 1u;
            continue;
        }
        st = c2quic_eat_vint(buf, len, &cur);
        if (st != C2QUIC_OK) {
            slab->status = st;
            return st;
        }
        if (cur.value > 0xFFFFFFFFull) {
            slab->status = C2QUIC_INVALID;
            return C2QUIC_INVALID;
        }
        typ = (uint32_t)cur.value;
        if (n >= 64u) {
            slab->n = n;
            slab->status = C2QUIC_FULL;
            return C2QUIC_FULL;
        }
        slab->frames[n].typ = typ;
        slab->frames[n].flags = 0;
        slab->frames[n].id = 0;
        slab->frames[n].offset = 0;
        slab->frames[n].length = 0;
        slab->frames[n].extra = 0;
        slab->frames[n].payload_off = 0;
        slab->frames[n].payload_len = 0;
        if (typ == C2QUIC_FT_PING) {
            n = n + 1u;
            continue;
        }
        if (typ >= 0x08u && typ <= 0x0Fu) {
            slab->frames[n].flags = typ & 0x07u;
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            slab->frames[n].id = cur.value;
            if ((typ & C2QUIC_FF_HAS_OFF) != 0) {
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
                slab->frames[n].offset = cur.value;
            }
            if ((typ & C2QUIC_FF_HAS_LEN) != 0) {
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
                slab->frames[n].length = cur.value;
            } else {
                slab->frames[n].length = (uint64_t)(len - cur.off);
            }
            if (slab->frames[n].length > (uint64_t)(len - cur.off)) {
                slab->status = C2QUIC_TRUNC;
                return C2QUIC_TRUNC;
            }
            slab->frames[n].payload_off = cur.off;
            slab->frames[n].payload_len = (uint32_t)slab->frames[n].length;
            cur.off = cur.off + (uint32_t)slab->frames[n].length;
            n = n + 1u;
            continue;
        }
        if (typ == C2QUIC_FT_ACK || typ == C2QUIC_FT_ACK_ECN) {
            if (typ == C2QUIC_FT_ACK_ECN) {
                slab->frames[n].flags = C2QUIC_FF_ECN;
            }
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            slab->frames[n].id = cur.value;
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            slab->frames[n].offset = cur.value;
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            range_count = cur.value;
            slab->frames[n].length = range_count;
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            slab->frames[n].extra = cur.value;
            slab->frames[n].payload_off = cur.off;
            i = 0;
            while (i < range_count) {
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
                i = i + 1u;
            }
            if (typ == C2QUIC_FT_ACK_ECN) {
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
            }
            slab->frames[n].payload_len = cur.off - slab->frames[n].payload_off;
            n = n + 1u;
            continue;
        }
        if (typ == C2QUIC_FT_RESET_STREAM) {
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            slab->frames[n].id = cur.value;
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            slab->frames[n].extra = cur.value;
            st = c2quic_eat_vint(buf, len, &cur);
            if (st != C2QUIC_OK) {
                slab->status = st;
                return st;
            }
            slab->frames[n].length = cur.value;
            n = n + 1u;
            continue;
        }
        if (typ == C2QUIC_FT_DATAGRAM || typ == C2QUIC_FT_DATAGRAM_LEN) {
            if (typ == C2QUIC_FT_DATAGRAM_LEN) {
                slab->frames[n].flags = C2QUIC_FF_HAS_LEN;
                st = c2quic_eat_vint(buf, len, &cur);
                if (st != C2QUIC_OK) {
                    slab->status = st;
                    return st;
                }
                slab->frames[n].length = cur.value;
            } else {
                slab->frames[n].length = (uint64_t)(len - cur.off);
            }
            if (slab->frames[n].length > (uint64_t)(len - cur.off)) {
                slab->status = C2QUIC_TRUNC;
                return C2QUIC_TRUNC;
            }
            slab->frames[n].payload_off = cur.off;
            slab->frames[n].payload_len = (uint32_t)slab->frames[n].length;
            cur.off = cur.off + (uint32_t)slab->frames[n].length;
            n = n + 1u;
            continue;
        }
        slab->status = C2QUIC_INVALID;
        return C2QUIC_INVALID;
    }
    slab->n = n;
    return C2QUIC_OK;
}
