//go:build ignore

#ifndef C2JSON_H
#define C2JSON_H

#include <stdint.h>

#define C2JSON_ERR     0
#define C2JSON_OBJECT  1
#define C2JSON_ARRAY   2
#define C2JSON_STRING  3
#define C2JSON_NUMBER  4
#define C2JSON_TRUE    5
#define C2JSON_FALSE   6
#define C2JSON_NULL    7

int c2json_validate(const uint8_t *json, int len);
int c2json_find_key(const uint8_t *json, int len, const uint8_t *key, int key_len, int *val_start, int *val_len, int *val_type);
int c2json_skip_value(const uint8_t *json, int len, int pos, int *end_pos, int *val_type);
int c2json_get_string(const uint8_t *json, int len, const uint8_t *key, int key_len, uint8_t *out, int out_max);
int c2json_get_int64(const uint8_t *json, int len, const uint8_t *key, int key_len, int64_t *out);
int c2json_get_bool(const uint8_t *json, int len, const uint8_t *key, int key_len, int *out);

#endif /* C2JSON_H */
