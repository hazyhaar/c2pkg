#include "c2_cell.h"
#include "c2_grid.h"
#include <stddef.h>

typedef struct {
    int state;
    int cursor_x;
    int cursor_y;
    int cur_fg;
    int cur_bg;
    int cur_attr;
    int width;
    int height;
    int prm;
    int nparams;
    int params[4];
    uint8_t pending_utf8[4];
    int pending_len;
} c2_vt_parser_t;

void c2_vt_init(c2_vt_parser_t *p, int width, int height) {
    p->state = 0;
    p->cursor_x = 0;
    p->cursor_y = 0;
    p->cur_fg = 0;
    p->cur_bg = 0;
    p->cur_attr = 0;
    p->width = width;
    p->height = height;
    p->prm = 0;
    p->nparams = 0;
    p->pending_len = 0;
}

void c2_vt_put(c2_vt_parser_t *p, c2_cell_t *cells, uint32_t r) {
    int w = p->width;
    int h = p->height;
    int idx;
    if (w <= 0 || h <= 0) {
        return;
    }
    if (p->cursor_x >= w) {
        p->cursor_x = 0;
        p->cursor_y = p->cursor_y + 1;
    }
    if (p->cursor_y >= h) {
        p->cursor_y = h - 1;
    }
    idx = p->cursor_y * w + p->cursor_x;
    cells[idx].rune = r;
    cells[idx].fg = (uint8_t)p->cur_fg;
    cells[idx].bg = (uint8_t)p->cur_bg;
    cells[idx].flags = (uint8_t)p->cur_attr;
    cells[idx].width = 1;
    p->cursor_x = p->cursor_x + 1;
}

void c2_vt_add_digit(c2_vt_parser_t *p, uint8_t b) {
    int d = (int)b;
    d = d - 48;
    int v = p->prm;
    v = v * 10;
    v = v + d;
    if (v > 9999) {
        v = 9999;
    }
    p->prm = v;
}

void c2_vt_finish_prm(c2_vt_parser_t *p) {
    if (p->nparams < 4) {
        p->params[p->nparams] = p->prm;
        p->nparams = p->nparams + 1;
    }
    p->prm = 0;
}

int c2_vt_param(c2_vt_parser_t *p, int i, int def) {
    if (i < p->nparams && p->params[i] != 0) {
        return p->params[i];
    }
    return def;
}

void c2_vt_sgr(c2_vt_parser_t *p) {
    int i;
    for (i = 0; i < p->nparams; i++) {
        int v = p->params[i];
        if (v == 0) {
            p->cur_fg = 0;
            p->cur_bg = 0;
            p->cur_attr = 0;
        }
        if (v == 1) {
            p->cur_attr = p->cur_attr | 1;
        }
        if (v == 4) {
            p->cur_attr = p->cur_attr | 4;
        }
        if (v == 7) {
            p->cur_attr = p->cur_attr | 8;
        }
        if (v >= 30 && v <= 37) {
            p->cur_fg = v - 30;
        }
        if (v >= 40 && v <= 47) {
            p->cur_bg = v - 40;
        }
    }
}

void c2_vt_clear_span(c2_cell_t *cells, int i0, int i1, uint8_t fg, uint8_t bg) {
    int i;
    for (i = i0; i < i1; i++) {
        cells[i].rune = 32;
        cells[i].fg = fg;
        cells[i].bg = bg;
        cells[i].flags = 0;
        cells[i].width = 1;
    }
}

void c2_vt_csi(c2_vt_parser_t *p, c2_cell_t *cells, uint8_t b) {
    int n;
    int w;
    int h;
    int idx;
    c2_vt_finish_prm(p);
    w = p->width;
    h = p->height;
    if (b == 'H' || b == 'f') {
        int row = c2_vt_param(p, 0, 1);
        int col = c2_vt_param(p, 1, 1);
        p->cursor_y = row - 1;
        p->cursor_x = col - 1;
        if (p->cursor_y < 0) {
            p->cursor_y = 0;
        }
        if (p->cursor_x < 0) {
            p->cursor_x = 0;
        }
        if (p->height > 0 && p->cursor_y >= p->height) {
            p->cursor_y = p->height - 1;
        }
        if (p->width > 0 && p->cursor_x >= p->width) {
            p->cursor_x = p->width - 1;
        }
    }
    if (b == 'A') {
        n = c2_vt_param(p, 0, 1);
        p->cursor_y = p->cursor_y - n;
        if (p->cursor_y < 0) {
            p->cursor_y = 0;
        }
    }
    if (b == 'B') {
        n = c2_vt_param(p, 0, 1);
        p->cursor_y = p->cursor_y + n;
        if (p->height > 0 && p->cursor_y >= p->height) {
            p->cursor_y = p->height - 1;
        }
    }
    if (b == 'C') {
        n = c2_vt_param(p, 0, 1);
        p->cursor_x = p->cursor_x + n;
        if (p->width > 0 && p->cursor_x >= p->width) {
            p->cursor_x = p->width - 1;
        }
    }
    if (b == 'D') {
        n = c2_vt_param(p, 0, 1);
        p->cursor_x = p->cursor_x - n;
        if (p->cursor_x < 0) {
            p->cursor_x = 0;
        }
    }
    if (b == 'm') {
        c2_vt_sgr(p);
    }
    if (b == 'J') {
        n = c2_vt_param(p, 0, 0);
        idx = p->cursor_y * w + p->cursor_x;
        if (n == 0) {
            c2_vt_clear_span(cells, idx, w * h, p->cur_fg, p->cur_bg);
        }
        if (n == 1) {
            c2_vt_clear_span(cells, 0, idx + 1, p->cur_fg, p->cur_bg);
        }
        if (n == 2) {
            c2_vt_clear_span(cells, 0, w * h, p->cur_fg, p->cur_bg);
        }
    }
    if (b == 'K') {
        n = c2_vt_param(p, 0, 0);
        idx = p->cursor_y * w;
        if (n == 0) {
            c2_vt_clear_span(cells, idx + p->cursor_x, idx + w, p->cur_fg, p->cur_bg);
        }
        if (n == 1) {
            c2_vt_clear_span(cells, idx, idx + p->cursor_x + 1, p->cur_fg, p->cur_bg);
        }
        if (n == 2) {
            c2_vt_clear_span(cells, idx, idx + w, p->cur_fg, p->cur_bg);
        }
    }
    p->state = 0;
}

int c2_vt_feed_byte(c2_vt_parser_t *p, c2_cell_t *cells, uint8_t b) {
    int z = 0;
    if (p->height < 0) {
        cells[z].rune = 0;
    }
    switch (p->state) {
    case 0:
        if (b == 0x1b) {
            p->state = 1;
            return 0;
        }
        if (b == 10) {
            p->cursor_y = p->cursor_y + 1;
            p->cursor_x = 0;
            return 1;
        }
        if (b == 13) {
            p->cursor_x = 0;
            return 1;
        }
        if (b >= 32 && b <= 0x7e) {
            c2_vt_put(p, cells, (uint32_t)b);
            return 1;
        }
        break;
    case 1:
        if (b == 0x5b) {
            p->state = 2;
            p->prm = 0;
            p->nparams = 0;
            return 0;
        }
        if (b == 0x5d) {
            p->state = 3;
            return 0;
        }
        p->state = 0;
        break;
    case 2:
        if (b >= 48 && b <= 57) {
            c2_vt_add_digit(p, b);
            return 0;
        }
        if (b == ';') {
            c2_vt_finish_prm(p);
            return 0;
        }
        if (b >= 64 && b <= 0x7e) {
            c2_vt_csi(p, cells, b);
            return 0;
        }
        break;
    case 3:
        if (b == 7 || b == 27) {
            p->state = 0;
            return 0;
        }
        break;
    }
    return 0;
}
