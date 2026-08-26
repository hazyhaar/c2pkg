/* sgoiter dogfood wrapper — disable arch asm / SIMD */
#define YYJSON_DISABLE_READER 0
#define YYJSON_DISABLE_WRITER 0
#define YYJSON_DISABLE_FAST_FP_CONV 1
/* force portable paths if present */
#define YYJSON_DISABLE_UTILS 0
#include "yyjson.c"
