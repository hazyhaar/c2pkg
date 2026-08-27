//go:build ignore

#include "c2json.h"

static inline int is_ws(uint8_t c) {
    if (c == ' ') return 1;
    if (c == '\t') return 1;
    if (c == '\n') return 1;
    if (c == '\r') return 1;
    return 0;
}

static inline int is_num_char(uint8_t c) {
    if (c >= '0' && c <= '9') return 1;
    if (c == '.') return 1;
    if (c == 'e' || c == 'E') return 1;
    if (c == '+' || c == '-') return 1;
    return 0;
}

static inline int skip_ws(const uint8_t *json, int len, int pos) {
    while (pos < len) {
        if (is_ws(json[pos]) == 0) {
            break;
        }
        pos++;
    }
    return pos;
}

int c2json_skip_value(const uint8_t *json, int len, int pos, int *end_pos, int *val_type) {
    int depth = 0;
    int in_str = 0;
    int num_start = 0;
    uint8_t c = 0;
    uint8_t ch = 0;

    pos = skip_ws(json, len, pos);
    if (pos >= len) {
        return 0;
    }

    c = json[pos];
    if (c == '"') {
        pos++;
        while (pos < len) {
            if (json[pos] == '\\') {
                pos += 2;
            } else if (json[pos] == '"') {
                pos++;
                *end_pos = pos;
                *val_type = C2JSON_STRING;
                return 1;
            } else {
                pos++;
            }
        }
        return 0;
    }

    if (c == '{') {
        depth = 1;
        in_str = 0;
        pos++;
        while (pos < len && depth > 0) {
            ch = json[pos];
            if (in_str != 0) {
                if (ch == '\\') {
                    pos += 2;
                    continue;
                } else if (ch == '"') {
                    in_str = 0;
                }
            } else {
                if (ch == '"') {
                    in_str = 1;
                } else if (ch == '{') {
                    depth++;
                } else if (ch == '}') {
                    depth--;
                }
            }
            pos++;
        }
        if (depth == 0) {
            *end_pos = pos;
            *val_type = C2JSON_OBJECT;
            return 1;
        }
        return 0;
    }

    if (c == '[') {
        depth = 1;
        in_str = 0;
        pos++;
        while (pos < len && depth > 0) {
            ch = json[pos];
            if (in_str != 0) {
                if (ch == '\\') {
                    pos += 2;
                    continue;
                } else if (ch == '"') {
                    in_str = 0;
                }
            } else {
                if (ch == '"') {
                    in_str = 1;
                } else if (ch == '[') {
                    depth++;
                } else if (ch == ']') {
                    depth--;
                }
            }
            pos++;
        }
        if (depth == 0) {
            *end_pos = pos;
            *val_type = C2JSON_ARRAY;
            return 1;
        }
        return 0;
    }

    if (c == 't' && pos + 3 < len && json[pos+1] == 'r' && json[pos+2] == 'u' && json[pos+3] == 'e') {
        *end_pos = pos + 4;
        *val_type = C2JSON_TRUE;
        return 1;
    }

    if (c == 'f' && pos + 4 < len && json[pos+1] == 'a' && json[pos+2] == 'l' && json[pos+3] == 's' && json[pos+4] == 'e') {
        *end_pos = pos + 5;
        *val_type = C2JSON_FALSE;
        return 1;
    }

    if (c == 'n' && pos + 3 < len && json[pos+1] == 'u' && json[pos+2] == 'l' && json[pos+3] == 'l') {
        *end_pos = pos + 4;
        *val_type = C2JSON_NULL;
        return 1;
    }

    if (c == '-' || (c >= '0' && c <= '9')) {
        num_start = pos;
        if (c == '-') pos++;
        while (pos < len) {
            if (is_num_char(json[pos]) == 0) {
                break;
            }
            pos++;
        }
        if (pos > num_start) {
            *end_pos = pos;
            *val_type = C2JSON_NUMBER;
            return 1;
        }
        return 0;
    }

    return 0;
}

int c2json_validate(const uint8_t *json, int len) {
    int end_pos = 0;
    int val_type = 0;

    if (!json || len <= 0) return 0;
    if (c2json_skip_value(json, len, 0, &end_pos, &val_type) == 0) {
        return 0;
    }
    end_pos = skip_ws(json, len, end_pos);
    if (end_pos == len) {
        return 1;
    }
    return 0;
}

int c2json_find_key(const uint8_t *json, int len, const uint8_t *key, int key_len, int *val_start, int *val_len, int *val_type) {
    int pos = 0;
    int k_start = 0;
    int k_len = 0;
    int v_start = 0;
    int v_end = 0;
    int v_type = 0;
    int match = 0;
    int i = 0;

    if (!json || len <= 0 || !key || key_len <= 0) return 0;

    pos = skip_ws(json, len, 0);
    if (pos >= len || json[pos] != '{') return 0;
    pos++;

    while (pos < len) {
        pos = skip_ws(json, len, pos);
        if (pos >= len) return 0;
        if (json[pos] == '}') return 0;

        if (json[pos] != '"') return 0;
        pos++;
        k_start = pos;
        while (pos < len && json[pos] != '"') {
            if (json[pos] == '\\') pos += 2;
            else pos++;
        }
        if (pos >= len) return 0;
        k_len = pos - k_start;
        pos++;

        pos = skip_ws(json, len, pos);
        if (pos >= len || json[pos] != ':') return 0;
        pos++;

        pos = skip_ws(json, len, pos);
        v_start = pos;
        v_end = 0;
        v_type = 0;
        if (c2json_skip_value(json, len, pos, &v_end, &v_type) == 0) {
            return 0;
        }

        if (k_len == key_len) {
            match = 1;
            for (i = 0; i < key_len; i++) {
                if (json[k_start + i] != key[i]) {
                    match = 0;
                    break;
                }
            }
            if (match != 0) {
                *val_start = v_start;
                *val_len = v_end - v_start;
                *val_type = v_type;
                return 1;
            }
        }

        pos = skip_ws(json, len, v_end);
        if (pos >= len) return 0;
        if (json[pos] == ',') {
            pos++;
        } else if (json[pos] == '}') {
            return 0;
        } else {
            return 0;
        }
    }

    return 0;
}

int c2json_get_string(const uint8_t *json, int len, const uint8_t *key, int key_len, uint8_t *out, int out_max) {
    int v_start = 0, v_len = 0, v_type = 0;
    int in_pos = 0, in_end = 0, out_pos = 0;
    uint8_t c = 0, esc = 0;

    if (c2json_find_key(json, len, key, key_len, &v_start, &v_len, &v_type) == 0) {
        return -1;
    }
    if (v_type != C2JSON_STRING || v_len < 2) {
        return -1;
    }

    in_pos = v_start + 1;
    in_end = v_start + v_len - 1;
    out_pos = 0;

    while (in_pos < in_end && out_pos < out_max) {
        c = json[in_pos++];
        if (c == '\\' && in_pos < in_end) {
            esc = json[in_pos++];
            if (esc == '"') out[out_pos++] = '"';
            else if (esc == '\\') out[out_pos++] = '\\';
            else if (esc == 'n') out[out_pos++] = '\n';
            else if (esc == 't') out[out_pos++] = '\t';
            else if (esc == 'r') out[out_pos++] = '\r';
            else if (esc == '/') out[out_pos++] = '/';
            else out[out_pos++] = esc;
        } else {
            out[out_pos++] = c;
        }
    }
    return out_pos;
}

int c2json_get_int64(const uint8_t *json, int len, const uint8_t *key, int key_len, int64_t *out) {
    int v_start = 0, v_len = 0, v_type = 0;
    int pos = 0, end = 0;
    int64_t sign = 1;
    int64_t val = 0;

    if (c2json_find_key(json, len, key, key_len, &v_start, &v_len, &v_type) == 0) {
        return 0;
    }
    if (v_type != C2JSON_NUMBER || v_len <= 0) {
        return 0;
    }

    pos = v_start;
    end = v_start + v_len;
    sign = 1;
    if (json[pos] == '-') {
        sign = -1;
        pos++;
    }

    val = 0;
    while (pos < end && json[pos] >= '0' && json[pos] <= '9') {
        val = val * 10 + (json[pos] - '0');
        pos++;
    }

    *out = val * sign;
    return 1;
}

int c2json_get_bool(const uint8_t *json, int len, const uint8_t *key, int key_len, int *out) {
    int v_start = 0, v_len = 0, v_type = 0;
    if (c2json_find_key(json, len, key, key_len, &v_start, &v_len, &v_type) == 0) {
        return 0;
    }
    if (v_type == C2JSON_TRUE) {
        *out = 1;
        return 1;
    }
    if (v_type == C2JSON_FALSE) {
        *out = 0;
        return 1;
    }
    return 0;
}
