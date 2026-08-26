#include "c2_cell.h"

typedef struct {
    int x;
    int y;
    int length;
} c2_span_t;

int c2_chunk_dirty4(const uint64_t *front, const uint64_t *back) {
    int m = 0;
    if (front[0] != back[0]) {
        m = m + 1;
    }
    if (front[1] != back[1]) {
        m = m + 2;
    }
    if (front[2] != back[2]) {
        m = m + 4;
    }
    if (front[3] != back[3]) {
        m = m + 8;
    }
    return m;
}

int c2_cells_equal_at(const c2_cell_t *front, const c2_cell_t *back, int i) {
    if (front[i].rune != back[i].rune) return 0;
    if (front[i].fg != back[i].fg) return 0;
    if (front[i].bg != back[i].bg) return 0;
    if (front[i].flags != back[i].flags) return 0;
    if (front[i].width != back[i].width) return 0;
    return 1;
}

int c2_diff_grid_scalar(const c2_cell_t *front, const c2_cell_t *back, int count, int stride, int width, c2_span_t *spans, int max_spans) {
    int span_count = 0;
    int in_span = 0;
    int span_start_x = 0;
    int span_y = 0;
    int span_len = 0;
    int overflow = 0;
    int i;

    if (width <= 0) {
        width = stride;
    }

    for (i = 0; i < count; i++) {
        int x = i % stride;
        int y = i / stride;
        int diff;
        if (x >= width) {
            if (in_span) {
                if (span_count < max_spans) {
                    spans[span_count].x = span_start_x;
                    spans[span_count].y = span_y;
                    spans[span_count].length = span_len;
                    span_count++;
                } else {
                    overflow = 1;
                }
                in_span = 0;
                span_len = 0;
            }
        } else {
            diff = 0;
            if (front[i].rune != back[i].rune) {
                diff = 1;
            } else if (front[i].fg != back[i].fg) {
                diff = 1;
            } else if (front[i].bg != back[i].bg) {
                diff = 1;
            } else if (front[i].flags != back[i].flags) {
                diff = 1;
            } else if (front[i].width != back[i].width) {
                diff = 1;
            }

            if (diff) {
                if (!in_span) {
                    in_span = 1;
                    span_start_x = x;
                    span_y = y;
                    span_len = 1;
                } else {
                    if (y == span_y && x == span_start_x + span_len) {
                        span_len++;
                    } else {
                        if (span_count < max_spans) {
                            spans[span_count].x = span_start_x;
                            spans[span_count].y = span_y;
                            spans[span_count].length = span_len;
                            span_count++;
                        } else {
                            overflow = 1;
                        }
                        span_start_x = x;
                        span_y = y;
                        span_len = 1;
                    }
                }
            } else {
                if (in_span) {
                    if (span_count < max_spans) {
                        spans[span_count].x = span_start_x;
                        spans[span_count].y = span_y;
                        spans[span_count].length = span_len;
                        span_count++;
                    } else {
                        overflow = 1;
                    }
                    in_span = 0;
                    span_len = 0;
                }
            }
        }
    }

    if (in_span) {
        if (span_count < max_spans) {
            spans[span_count].x = span_start_x;
            spans[span_count].y = span_y;
            spans[span_count].length = span_len;
            span_count++;
        } else {
            overflow = 1;
        }
    }

    if (overflow) {
        return -1;
    }
    return span_count;
}
