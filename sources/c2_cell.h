#ifndef C2_CELL_H
#define C2_CELL_H

#include <stdint.h>

typedef struct {
    uint32_t rune;
    uint8_t fg;
    uint8_t bg;
    uint8_t flags;
    uint8_t width;
} c2_cell_t;

#endif
