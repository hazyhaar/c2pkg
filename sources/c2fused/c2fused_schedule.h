/* Code generated from sgoiter/spec/fused/fused_schedule.cue. DO NOT EDIT. */
#ifndef C2FUSED_SCHEDULE_H
#define C2FUSED_SCHEDULE_H

#include <stdint.h>

#define C2FUSED_N_QR 80
#define C2FUSED_N_POLY 32

/* slot[q] = index du bloc Poly (0..31) absorbé après le quart q, ou -1. */
static const int8_t c2fused_poly_slot[C2FUSED_N_QR] = {
    -1,  0, -1, -1,  1, -1,  2, -1,
    -1,  3, -1,  4, -1, -1,  5, -1,
     6, -1, -1,  7, -1,  8, -1, -1,
     9, -1, 10, -1, -1, 11, -1, 12,
    -1, -1, 13, -1, 14, -1, -1, 15,
    -1, 16, -1, -1, 17, -1, 18, -1,
    -1, 19, -1, 20, -1, -1, 21, -1,
    22, -1, -1, 23, -1, 24, -1, -1,
    25, -1, 26, -1, -1, 27, -1, 28,
    -1, -1, 29, -1, 30, -1, -1, 31
};

#endif
