#ifndef C2QUIC_H
#define C2QUIC_H

#include <stdint.h>
#include <stddef.h>

#define C2QUIC_OK        0u
#define C2QUIC_TRUNC     1u
#define C2QUIC_OVERFLOW  2u
#define C2QUIC_INVALID   3u
#define C2QUIC_FULL      4u

#define C2QUIC_BATCH         8
#define C2QUIC_MAX_FRAMES    64
#define C2QUIC_HP_SAMPLE_LEN 16
#define C2QUIC_HP_MASK_LEN   5

#define C2QUIC_FT_PADDING      0x00u
#define C2QUIC_FT_PING         0x01u
#define C2QUIC_FT_ACK          0x02u
#define C2QUIC_FT_ACK_ECN      0x03u
#define C2QUIC_FT_RESET_STREAM 0x04u
#define C2QUIC_FT_DATAGRAM     0x30u
#define C2QUIC_FT_DATAGRAM_LEN 0x31u

#define C2QUIC_FF_FIN     0x01u
#define C2QUIC_FF_HAS_LEN 0x02u
#define C2QUIC_FF_HAS_OFF 0x04u
#define C2QUIC_FF_ECN     0x08u

typedef struct c2quic_vint {
    uint64_t value;
    uint32_t nbytes;
    uint32_t status;
} c2quic_vint_t;

typedef struct c2quic_cursor {
    uint32_t off;
    uint32_t status;
    uint64_t value;
    uint32_t nbytes;
    uint32_t pad;
} c2quic_cursor_t;

typedef struct c2quic_batch8 {
    uint32_t offs[8];
    uint64_t vals[8];
    uint32_t nbytes[8];
    uint32_t status[8];
} c2quic_batch8_t;

typedef struct c2quic_hp_io {
    uint8_t  sample[16];
    uint8_t  mask[5];
    uint8_t  pad[3];
    uint32_t pn_len;
    uint64_t pn;
} c2quic_hp_io_t;

typedef struct c2quic_frame {
    uint32_t typ;
    uint32_t flags;
    uint64_t id;
    uint64_t offset;
    uint64_t length;
    uint64_t extra;
    uint32_t payload_off;
    uint32_t payload_len;
} c2quic_frame_t;

typedef struct c2quic_frame_slab {
    c2quic_frame_t frames[64];
    uint32_t n;
    uint32_t status;
} c2quic_frame_slab_t;

extern const uint32_t c2quic_vint_len_tab[4];

uint32_t c2quic_vint_nbytes(uint64_t value);
uint32_t c2quic_vint_decode(const uint8_t *buf, uint32_t len, uint32_t off, c2quic_vint_t *out);
uint32_t c2quic_vint_encode(uint64_t value, uint8_t *buf, uint32_t len, uint32_t off, c2quic_vint_t *out);
void c2quic_vint_decode_batch8_scalar(const uint8_t *buf, uint32_t len, c2quic_batch8_t *b);
void c2quic_vint_decode_batch8(const uint8_t *buf, uint32_t len, c2quic_batch8_t *b);

uint32_t c2quic_hp_sample(const uint8_t *pkt, uint32_t len, uint32_t pn_off, c2quic_hp_io_t *io);
uint32_t c2quic_hp_apply(uint8_t *pkt, uint32_t len, uint32_t pn_off, uint32_t pn_len, c2quic_hp_io_t *io);
uint32_t c2quic_hp_remove(uint8_t *pkt, uint32_t len, uint32_t pn_off, c2quic_hp_io_t *io);
uint32_t c2quic_hp_read_pn(const uint8_t *pkt, uint32_t len, uint32_t pn_off, uint32_t pn_len, c2quic_hp_io_t *io);

uint32_t c2quic_unpack_frames(const uint8_t *buf, uint32_t len, c2quic_frame_slab_t *slab);

#endif
