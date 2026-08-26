#ifndef C2_GRID_H
#define C2_GRID_H

#include "c2_cell.h"

void c2_grid_clear(c2_cell_t *cells, int n, uint8_t fg, uint8_t bg) {
    int i;
    for (i = 0; i < n; i++) {
        cells[i].rune = 32;
        cells[i].fg = fg;
        cells[i].bg = bg;
        cells[i].flags = 0;
        cells[i].width = 1;
    }
}

void c2_grid_clear_row(c2_cell_t *cells, int y, int stride, int width, uint8_t fg, uint8_t bg) {
    int x;
    int base = y * stride;
    for (x = 0; x < width; x++) {
        int i = base + x;
        cells[i].rune = 32;
        cells[i].fg = fg;
        cells[i].bg = bg;
        cells[i].flags = 0;
        cells[i].width = 1;
    }
}

void c2_grid_scroll_up(c2_cell_t *cells, int width, int height, int stride, int nlines) {
    int y;
    int x;
    int lim;
    int srcy;
    int dst;
    int src;
    if (nlines < 1) {
        return;
    }
    if (nlines > height) {
        nlines = height;
    }
    lim = height - nlines;
    for (y = 0; y < lim; y++) {
        srcy = y + nlines;
        dst = y * stride;
        src = srcy * stride;
        for (x = 0; x < width; x++) {
            cells[dst + x].rune = cells[src + x].rune;
            cells[dst + x].fg = cells[src + x].fg;
            cells[dst + x].bg = cells[src + x].bg;
            cells[dst + x].flags = cells[src + x].flags;
            cells[dst + x].width = cells[src + x].width;
        }
    }
    for (y = lim; y < height; y++) {
        c2_grid_clear_row(cells, y, stride, width, 0, 0);
    }
}

#endif
