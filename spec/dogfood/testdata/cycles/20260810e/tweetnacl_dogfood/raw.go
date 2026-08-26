// Code generated for linux/amd64 by 'ccgo --package-name=df_tweetnacl -o spec/dogfood/cycles/20260810e/tweetnacl_dogfood/raw.go -I spec/dogfood/cycles/20260810e/tweetnacl_dogfood spec/dogfood/cycles/20260810e/tweetnacl_dogfood/src.c', DO NOT EDIT.

//go:build linux && amd64

package df_tweetnacl

import (
	"math"
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

const INT16_MAX = 0x7fff
const INT32_MAX = 0x7fffffff
const INT64_MAX = 0x7fffffffffffffff
const INT8_MAX = 0x7f
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT64_MAX"
const INTPTR_MIN = "INT64_MIN"
const INT_FAST16_MAX = "INT32_MAX"
const INT_FAST16_MIN = "INT32_MIN"
const INT_FAST32_MAX = "INT32_MAX"
const INT_FAST32_MIN = "INT32_MIN"
const INT_FAST64_MAX = "INT64_MAX"
const INT_FAST64_MIN = "INT64_MIN"
const INT_FAST8_MAX = "INT8_MAX"
const INT_FAST8_MIN = "INT8_MIN"
const INT_LEAST16_MAX = "INT16_MAX"
const INT_LEAST16_MIN = "INT16_MIN"
const INT_LEAST32_MAX = "INT32_MAX"
const INT_LEAST32_MIN = "INT32_MIN"
const INT_LEAST64_MAX = "INT64_MAX"
const INT_LEAST64_MIN = "INT64_MIN"
const INT_LEAST8_MAX = "INT8_MAX"
const INT_LEAST8_MIN = "INT8_MIN"
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const UINT16_MAX = 0xffff
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 0xff
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT64_MAX"
const UINT_FAST16_MAX = "UINT32_MAX"
const UINT_FAST32_MAX = "UINT32_MAX"
const UINT_FAST64_MAX = "UINT64_MAX"
const UINT_FAST8_MAX = "UINT8_MAX"
const UINT_LEAST16_MAX = "UINT16_MAX"
const UINT_LEAST32_MAX = "UINT32_MAX"
const UINT_LEAST64_MAX = "UINT64_MAX"
const UINT_LEAST8_MAX = "UINT8_MAX"
const WINT_MAX = "UINT32_MAX"
const WINT_MIN = 0
const _GNU_SOURCE = 1
const _LP64 = 1
const _STDC_PREDEF_H = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CET__ = 3
const __CHAR_BIT__ = 8
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 3
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 13
const __GXX_ABI_VERSION = 1018
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const __LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const __LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __LITTLE_ENDIAN = 1234
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP_STRONG__ = 3
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __USE_TIME_BITS64 = 1
const __VERSION__ = "13.3.0"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const crypto_auth = "crypto_auth_hmacsha512256"
const crypto_auth_BYTES = "crypto_auth_hmacsha512256_BYTES"
const crypto_auth_IMPLEMENTATION = "crypto_auth_hmacsha512256_IMPLEMENTATION"
const crypto_auth_KEYBYTES = "crypto_auth_hmacsha512256_KEYBYTES"
const crypto_auth_PRIMITIVE = "hmacsha512256"
const crypto_auth_VERSION = "crypto_auth_hmacsha512256_VERSION"
const crypto_auth_hmacsha512256 = "crypto_auth_hmacsha512256_tweet"
const crypto_auth_hmacsha512256_BYTES = "crypto_auth_hmacsha512256_tweet_BYTES"
const crypto_auth_hmacsha512256_IMPLEMENTATION = "crypto_auth/hmacsha512256/tweet"
const crypto_auth_hmacsha512256_KEYBYTES = "crypto_auth_hmacsha512256_tweet_KEYBYTES"
const crypto_auth_hmacsha512256_VERSION = "crypto_auth_hmacsha512256_tweet_VERSION"
const crypto_auth_hmacsha512256_tweet_BYTES = 32
const crypto_auth_hmacsha512256_tweet_KEYBYTES = 32
const crypto_auth_hmacsha512256_tweet_VERSION = "-"
const crypto_auth_hmacsha512256_verify = "crypto_auth_hmacsha512256_tweet_verify"
const crypto_auth_verify = "crypto_auth_hmacsha512256_verify"
const crypto_box = "crypto_box_curve25519xsalsa20poly1305"
const crypto_box_BEFORENMBYTES = "crypto_box_curve25519xsalsa20poly1305_BEFORENMBYTES"
const crypto_box_BOXZEROBYTES = "crypto_box_curve25519xsalsa20poly1305_BOXZEROBYTES"
const crypto_box_IMPLEMENTATION = "crypto_box_curve25519xsalsa20poly1305_IMPLEMENTATION"
const crypto_box_NONCEBYTES = "crypto_box_curve25519xsalsa20poly1305_NONCEBYTES"
const crypto_box_PRIMITIVE = "curve25519xsalsa20poly1305"
const crypto_box_PUBLICKEYBYTES = "crypto_box_curve25519xsalsa20poly1305_PUBLICKEYBYTES"
const crypto_box_SECRETKEYBYTES = "crypto_box_curve25519xsalsa20poly1305_SECRETKEYBYTES"
const crypto_box_VERSION = "crypto_box_curve25519xsalsa20poly1305_VERSION"
const crypto_box_ZEROBYTES = "crypto_box_curve25519xsalsa20poly1305_ZEROBYTES"
const crypto_box_afternm = "crypto_box_curve25519xsalsa20poly1305_afternm"
const crypto_box_beforenm = "crypto_box_curve25519xsalsa20poly1305_beforenm"
const crypto_box_curve25519xsalsa20poly1305 = "crypto_box_curve25519xsalsa20poly1305_tweet"
const crypto_box_curve25519xsalsa20poly1305_BEFORENMBYTES = "crypto_box_curve25519xsalsa20poly1305_tweet_BEFORENMBYTES"
const crypto_box_curve25519xsalsa20poly1305_BOXZEROBYTES = "crypto_box_curve25519xsalsa20poly1305_tweet_BOXZEROBYTES"
const crypto_box_curve25519xsalsa20poly1305_IMPLEMENTATION = "crypto_box/curve25519xsalsa20poly1305/tweet"
const crypto_box_curve25519xsalsa20poly1305_NONCEBYTES = "crypto_box_curve25519xsalsa20poly1305_tweet_NONCEBYTES"
const crypto_box_curve25519xsalsa20poly1305_PUBLICKEYBYTES = "crypto_box_curve25519xsalsa20poly1305_tweet_PUBLICKEYBYTES"
const crypto_box_curve25519xsalsa20poly1305_SECRETKEYBYTES = "crypto_box_curve25519xsalsa20poly1305_tweet_SECRETKEYBYTES"
const crypto_box_curve25519xsalsa20poly1305_VERSION = "crypto_box_curve25519xsalsa20poly1305_tweet_VERSION"
const crypto_box_curve25519xsalsa20poly1305_ZEROBYTES = "crypto_box_curve25519xsalsa20poly1305_tweet_ZEROBYTES"
const crypto_box_curve25519xsalsa20poly1305_afternm = "crypto_box_curve25519xsalsa20poly1305_tweet_afternm"
const crypto_box_curve25519xsalsa20poly1305_beforenm = "crypto_box_curve25519xsalsa20poly1305_tweet_beforenm"
const crypto_box_curve25519xsalsa20poly1305_keypair = "crypto_box_curve25519xsalsa20poly1305_tweet_keypair"
const crypto_box_curve25519xsalsa20poly1305_open = "crypto_box_curve25519xsalsa20poly1305_tweet_open"
const crypto_box_curve25519xsalsa20poly1305_open_afternm = "crypto_box_curve25519xsalsa20poly1305_tweet_open_afternm"
const crypto_box_curve25519xsalsa20poly1305_tweet_BEFORENMBYTES = 32
const crypto_box_curve25519xsalsa20poly1305_tweet_BOXZEROBYTES = 16
const crypto_box_curve25519xsalsa20poly1305_tweet_NONCEBYTES = 24
const crypto_box_curve25519xsalsa20poly1305_tweet_PUBLICKEYBYTES = 32
const crypto_box_curve25519xsalsa20poly1305_tweet_SECRETKEYBYTES = 32
const crypto_box_curve25519xsalsa20poly1305_tweet_VERSION = "-"
const crypto_box_curve25519xsalsa20poly1305_tweet_ZEROBYTES = 32
const crypto_box_keypair = "crypto_box_curve25519xsalsa20poly1305_keypair"
const crypto_box_open = "crypto_box_curve25519xsalsa20poly1305_open"
const crypto_box_open_afternm = "crypto_box_curve25519xsalsa20poly1305_open_afternm"
const crypto_core = "crypto_core_salsa20"
const crypto_core_CONSTBYTES = "crypto_core_salsa20_CONSTBYTES"
const crypto_core_IMPLEMENTATION = "crypto_core_salsa20_IMPLEMENTATION"
const crypto_core_INPUTBYTES = "crypto_core_salsa20_INPUTBYTES"
const crypto_core_KEYBYTES = "crypto_core_salsa20_KEYBYTES"
const crypto_core_OUTPUTBYTES = "crypto_core_salsa20_OUTPUTBYTES"
const crypto_core_PRIMITIVE = "salsa20"
const crypto_core_VERSION = "crypto_core_salsa20_VERSION"
const crypto_core_hsalsa20 = "crypto_core_hsalsa20_tweet"
const crypto_core_hsalsa20_CONSTBYTES = "crypto_core_hsalsa20_tweet_CONSTBYTES"
const crypto_core_hsalsa20_IMPLEMENTATION = "crypto_core/hsalsa20/tweet"
const crypto_core_hsalsa20_INPUTBYTES = "crypto_core_hsalsa20_tweet_INPUTBYTES"
const crypto_core_hsalsa20_KEYBYTES = "crypto_core_hsalsa20_tweet_KEYBYTES"
const crypto_core_hsalsa20_OUTPUTBYTES = "crypto_core_hsalsa20_tweet_OUTPUTBYTES"
const crypto_core_hsalsa20_VERSION = "crypto_core_hsalsa20_tweet_VERSION"
const crypto_core_hsalsa20_tweet_CONSTBYTES = 16
const crypto_core_hsalsa20_tweet_INPUTBYTES = 16
const crypto_core_hsalsa20_tweet_KEYBYTES = 32
const crypto_core_hsalsa20_tweet_OUTPUTBYTES = 32
const crypto_core_hsalsa20_tweet_VERSION = "-"
const crypto_core_salsa20 = "crypto_core_salsa20_tweet"
const crypto_core_salsa20_CONSTBYTES = "crypto_core_salsa20_tweet_CONSTBYTES"
const crypto_core_salsa20_IMPLEMENTATION = "crypto_core/salsa20/tweet"
const crypto_core_salsa20_INPUTBYTES = "crypto_core_salsa20_tweet_INPUTBYTES"
const crypto_core_salsa20_KEYBYTES = "crypto_core_salsa20_tweet_KEYBYTES"
const crypto_core_salsa20_OUTPUTBYTES = "crypto_core_salsa20_tweet_OUTPUTBYTES"
const crypto_core_salsa20_VERSION = "crypto_core_salsa20_tweet_VERSION"
const crypto_core_salsa20_tweet_CONSTBYTES = 16
const crypto_core_salsa20_tweet_INPUTBYTES = 16
const crypto_core_salsa20_tweet_KEYBYTES = 32
const crypto_core_salsa20_tweet_OUTPUTBYTES = 64
const crypto_core_salsa20_tweet_VERSION = "-"
const crypto_hash = "crypto_hash_sha512"
const crypto_hash_BYTES = "crypto_hash_sha512_BYTES"
const crypto_hash_IMPLEMENTATION = "crypto_hash_sha512_IMPLEMENTATION"
const crypto_hash_PRIMITIVE = "sha512"
const crypto_hash_VERSION = "crypto_hash_sha512_VERSION"
const crypto_hash_sha256 = "crypto_hash_sha256_tweet"
const crypto_hash_sha256_BYTES = "crypto_hash_sha256_tweet_BYTES"
const crypto_hash_sha256_IMPLEMENTATION = "crypto_hash/sha256/tweet"
const crypto_hash_sha256_VERSION = "crypto_hash_sha256_tweet_VERSION"
const crypto_hash_sha256_tweet_BYTES = 32
const crypto_hash_sha256_tweet_VERSION = "-"
const crypto_hash_sha512 = "crypto_hash_sha512_tweet"
const crypto_hash_sha512_BYTES = "crypto_hash_sha512_tweet_BYTES"
const crypto_hash_sha512_IMPLEMENTATION = "crypto_hash/sha512/tweet"
const crypto_hash_sha512_VERSION = "crypto_hash_sha512_tweet_VERSION"
const crypto_hash_sha512_tweet_BYTES = 64
const crypto_hash_sha512_tweet_VERSION = "-"
const crypto_hashblocks = "crypto_hashblocks_sha512"
const crypto_hashblocks_BLOCKBYTES = "crypto_hashblocks_sha512_BLOCKBYTES"
const crypto_hashblocks_IMPLEMENTATION = "crypto_hashblocks_sha512_IMPLEMENTATION"
const crypto_hashblocks_PRIMITIVE = "sha512"
const crypto_hashblocks_STATEBYTES = "crypto_hashblocks_sha512_STATEBYTES"
const crypto_hashblocks_VERSION = "crypto_hashblocks_sha512_VERSION"
const crypto_hashblocks_sha256 = "crypto_hashblocks_sha256_tweet"
const crypto_hashblocks_sha256_BLOCKBYTES = "crypto_hashblocks_sha256_tweet_BLOCKBYTES"
const crypto_hashblocks_sha256_IMPLEMENTATION = "crypto_hashblocks/sha256/tweet"
const crypto_hashblocks_sha256_STATEBYTES = "crypto_hashblocks_sha256_tweet_STATEBYTES"
const crypto_hashblocks_sha256_VERSION = "crypto_hashblocks_sha256_tweet_VERSION"
const crypto_hashblocks_sha256_tweet_BLOCKBYTES = 64
const crypto_hashblocks_sha256_tweet_STATEBYTES = 32
const crypto_hashblocks_sha256_tweet_VERSION = "-"
const crypto_hashblocks_sha512 = "crypto_hashblocks_sha512_tweet"
const crypto_hashblocks_sha512_BLOCKBYTES = "crypto_hashblocks_sha512_tweet_BLOCKBYTES"
const crypto_hashblocks_sha512_IMPLEMENTATION = "crypto_hashblocks/sha512/tweet"
const crypto_hashblocks_sha512_STATEBYTES = "crypto_hashblocks_sha512_tweet_STATEBYTES"
const crypto_hashblocks_sha512_VERSION = "crypto_hashblocks_sha512_tweet_VERSION"
const crypto_hashblocks_sha512_tweet_BLOCKBYTES = 128
const crypto_hashblocks_sha512_tweet_STATEBYTES = 64
const crypto_hashblocks_sha512_tweet_VERSION = "-"
const crypto_onetimeauth = "crypto_onetimeauth_poly1305"
const crypto_onetimeauth_BYTES = "crypto_onetimeauth_poly1305_BYTES"
const crypto_onetimeauth_IMPLEMENTATION = "crypto_onetimeauth_poly1305_IMPLEMENTATION"
const crypto_onetimeauth_KEYBYTES = "crypto_onetimeauth_poly1305_KEYBYTES"
const crypto_onetimeauth_PRIMITIVE = "poly1305"
const crypto_onetimeauth_VERSION = "crypto_onetimeauth_poly1305_VERSION"
const crypto_onetimeauth_poly1305 = "crypto_onetimeauth_poly1305_tweet"
const crypto_onetimeauth_poly1305_BYTES = "crypto_onetimeauth_poly1305_tweet_BYTES"
const crypto_onetimeauth_poly1305_IMPLEMENTATION = "crypto_onetimeauth/poly1305/tweet"
const crypto_onetimeauth_poly1305_KEYBYTES = "crypto_onetimeauth_poly1305_tweet_KEYBYTES"
const crypto_onetimeauth_poly1305_VERSION = "crypto_onetimeauth_poly1305_tweet_VERSION"
const crypto_onetimeauth_poly1305_tweet_BYTES = 16
const crypto_onetimeauth_poly1305_tweet_KEYBYTES = 32
const crypto_onetimeauth_poly1305_tweet_VERSION = "-"
const crypto_onetimeauth_poly1305_verify = "crypto_onetimeauth_poly1305_tweet_verify"
const crypto_onetimeauth_verify = "crypto_onetimeauth_poly1305_verify"
const crypto_scalarmult = "crypto_scalarmult_curve25519"
const crypto_scalarmult_BYTES = "crypto_scalarmult_curve25519_BYTES"
const crypto_scalarmult_IMPLEMENTATION = "crypto_scalarmult_curve25519_IMPLEMENTATION"
const crypto_scalarmult_PRIMITIVE = "curve25519"
const crypto_scalarmult_SCALARBYTES = "crypto_scalarmult_curve25519_SCALARBYTES"
const crypto_scalarmult_VERSION = "crypto_scalarmult_curve25519_VERSION"
const crypto_scalarmult_base = "crypto_scalarmult_curve25519_base"
const crypto_scalarmult_curve25519 = "crypto_scalarmult_curve25519_tweet"
const crypto_scalarmult_curve25519_BYTES = "crypto_scalarmult_curve25519_tweet_BYTES"
const crypto_scalarmult_curve25519_IMPLEMENTATION = "crypto_scalarmult/curve25519/tweet"
const crypto_scalarmult_curve25519_SCALARBYTES = "crypto_scalarmult_curve25519_tweet_SCALARBYTES"
const crypto_scalarmult_curve25519_VERSION = "crypto_scalarmult_curve25519_tweet_VERSION"
const crypto_scalarmult_curve25519_base = "crypto_scalarmult_curve25519_tweet_base"
const crypto_scalarmult_curve25519_tweet_BYTES = 32
const crypto_scalarmult_curve25519_tweet_SCALARBYTES = 32
const crypto_scalarmult_curve25519_tweet_VERSION = "-"
const crypto_secretbox = "crypto_secretbox_xsalsa20poly1305"
const crypto_secretbox_BOXZEROBYTES = "crypto_secretbox_xsalsa20poly1305_BOXZEROBYTES"
const crypto_secretbox_IMPLEMENTATION = "crypto_secretbox_xsalsa20poly1305_IMPLEMENTATION"
const crypto_secretbox_KEYBYTES = "crypto_secretbox_xsalsa20poly1305_KEYBYTES"
const crypto_secretbox_NONCEBYTES = "crypto_secretbox_xsalsa20poly1305_NONCEBYTES"
const crypto_secretbox_PRIMITIVE = "xsalsa20poly1305"
const crypto_secretbox_VERSION = "crypto_secretbox_xsalsa20poly1305_VERSION"
const crypto_secretbox_ZEROBYTES = "crypto_secretbox_xsalsa20poly1305_ZEROBYTES"
const crypto_secretbox_open = "crypto_secretbox_xsalsa20poly1305_open"
const crypto_secretbox_xsalsa20poly1305 = "crypto_secretbox_xsalsa20poly1305_tweet"
const crypto_secretbox_xsalsa20poly1305_BOXZEROBYTES = "crypto_secretbox_xsalsa20poly1305_tweet_BOXZEROBYTES"
const crypto_secretbox_xsalsa20poly1305_IMPLEMENTATION = "crypto_secretbox/xsalsa20poly1305/tweet"
const crypto_secretbox_xsalsa20poly1305_KEYBYTES = "crypto_secretbox_xsalsa20poly1305_tweet_KEYBYTES"
const crypto_secretbox_xsalsa20poly1305_NONCEBYTES = "crypto_secretbox_xsalsa20poly1305_tweet_NONCEBYTES"
const crypto_secretbox_xsalsa20poly1305_VERSION = "crypto_secretbox_xsalsa20poly1305_tweet_VERSION"
const crypto_secretbox_xsalsa20poly1305_ZEROBYTES = "crypto_secretbox_xsalsa20poly1305_tweet_ZEROBYTES"
const crypto_secretbox_xsalsa20poly1305_open = "crypto_secretbox_xsalsa20poly1305_tweet_open"
const crypto_secretbox_xsalsa20poly1305_tweet_BOXZEROBYTES = 16
const crypto_secretbox_xsalsa20poly1305_tweet_KEYBYTES = 32
const crypto_secretbox_xsalsa20poly1305_tweet_NONCEBYTES = 24
const crypto_secretbox_xsalsa20poly1305_tweet_VERSION = "-"
const crypto_secretbox_xsalsa20poly1305_tweet_ZEROBYTES = 32
const crypto_sign = "crypto_sign_ed25519"
const crypto_sign_BYTES = "crypto_sign_ed25519_BYTES"
const crypto_sign_IMPLEMENTATION = "crypto_sign_ed25519_IMPLEMENTATION"
const crypto_sign_PRIMITIVE = "ed25519"
const crypto_sign_PUBLICKEYBYTES = "crypto_sign_ed25519_PUBLICKEYBYTES"
const crypto_sign_SECRETKEYBYTES = "crypto_sign_ed25519_SECRETKEYBYTES"
const crypto_sign_VERSION = "crypto_sign_ed25519_VERSION"
const crypto_sign_ed25519 = "crypto_sign_ed25519_tweet"
const crypto_sign_ed25519_BYTES = "crypto_sign_ed25519_tweet_BYTES"
const crypto_sign_ed25519_IMPLEMENTATION = "crypto_sign/ed25519/tweet"
const crypto_sign_ed25519_PUBLICKEYBYTES = "crypto_sign_ed25519_tweet_PUBLICKEYBYTES"
const crypto_sign_ed25519_SECRETKEYBYTES = "crypto_sign_ed25519_tweet_SECRETKEYBYTES"
const crypto_sign_ed25519_VERSION = "crypto_sign_ed25519_tweet_VERSION"
const crypto_sign_ed25519_keypair = "crypto_sign_ed25519_tweet_keypair"
const crypto_sign_ed25519_open = "crypto_sign_ed25519_tweet_open"
const crypto_sign_ed25519_tweet_BYTES = 64
const crypto_sign_ed25519_tweet_PUBLICKEYBYTES = 32
const crypto_sign_ed25519_tweet_SECRETKEYBYTES = 64
const crypto_sign_ed25519_tweet_VERSION = "-"
const crypto_sign_keypair = "crypto_sign_ed25519_keypair"
const crypto_sign_open = "crypto_sign_ed25519_open"
const crypto_stream = "crypto_stream_xsalsa20"
const crypto_stream_IMPLEMENTATION = "crypto_stream_xsalsa20_IMPLEMENTATION"
const crypto_stream_KEYBYTES = "crypto_stream_xsalsa20_KEYBYTES"
const crypto_stream_NONCEBYTES = "crypto_stream_xsalsa20_NONCEBYTES"
const crypto_stream_PRIMITIVE = "xsalsa20"
const crypto_stream_VERSION = "crypto_stream_xsalsa20_VERSION"
const crypto_stream_salsa20 = "crypto_stream_salsa20_tweet"
const crypto_stream_salsa20_IMPLEMENTATION = "crypto_stream/salsa20/tweet"
const crypto_stream_salsa20_KEYBYTES = "crypto_stream_salsa20_tweet_KEYBYTES"
const crypto_stream_salsa20_NONCEBYTES = "crypto_stream_salsa20_tweet_NONCEBYTES"
const crypto_stream_salsa20_VERSION = "crypto_stream_salsa20_tweet_VERSION"
const crypto_stream_salsa20_tweet_KEYBYTES = 32
const crypto_stream_salsa20_tweet_NONCEBYTES = 8
const crypto_stream_salsa20_tweet_VERSION = "-"
const crypto_stream_salsa20_xor = "crypto_stream_salsa20_tweet_xor"
const crypto_stream_xor = "crypto_stream_xsalsa20_xor"
const crypto_stream_xsalsa20 = "crypto_stream_xsalsa20_tweet"
const crypto_stream_xsalsa20_IMPLEMENTATION = "crypto_stream/xsalsa20/tweet"
const crypto_stream_xsalsa20_KEYBYTES = "crypto_stream_xsalsa20_tweet_KEYBYTES"
const crypto_stream_xsalsa20_NONCEBYTES = "crypto_stream_xsalsa20_tweet_NONCEBYTES"
const crypto_stream_xsalsa20_VERSION = "crypto_stream_xsalsa20_tweet_VERSION"
const crypto_stream_xsalsa20_tweet_KEYBYTES = 32
const crypto_stream_xsalsa20_tweet_NONCEBYTES = 24
const crypto_stream_xsalsa20_tweet_VERSION = "-"
const crypto_stream_xsalsa20_xor = "crypto_stream_xsalsa20_tweet_xor"
const crypto_verify = "crypto_verify_16"
const crypto_verify_16 = "crypto_verify_16_tweet"
const crypto_verify_16_BYTES = "crypto_verify_16_tweet_BYTES"
const crypto_verify_16_IMPLEMENTATION = "crypto_verify/16/tweet"
const crypto_verify_16_VERSION = "crypto_verify_16_tweet_VERSION"
const crypto_verify_16_tweet_BYTES = 16
const crypto_verify_16_tweet_VERSION = "-"
const crypto_verify_32 = "crypto_verify_32_tweet"
const crypto_verify_32_BYTES = "crypto_verify_32_tweet_BYTES"
const crypto_verify_32_IMPLEMENTATION = "crypto_verify/32/tweet"
const crypto_verify_32_VERSION = "crypto_verify_32_tweet_VERSION"
const crypto_verify_32_tweet_BYTES = 32
const crypto_verify_32_tweet_VERSION = "-"
const crypto_verify_BYTES = "crypto_verify_16_BYTES"
const crypto_verify_IMPLEMENTATION = "crypto_verify_16_IMPLEMENTATION"
const crypto_verify_PRIMITIVE = "16"
const crypto_verify_VERSION = "crypto_verify_16_VERSION"
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type uintptr_t = uint64

type intptr_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type intmax_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type uintmax_t = uint64

type int_fast8_t = int8

type int_fast64_t = int64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_fast8_t = uint8

type uint_fast64_t = uint64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast16_t = int32

type int_fast32_t = int32

type uint_fast16_t = uint32

type uint_fast32_t = uint32

// C documentation
//
//	/* dogfood stub — non crypto RNG */
func randombytes(tls *libc.TLS, p uintptr, n uint64) {
	var i uint64
	_ = i
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		**(**uint8)(__ccgo_up(p + uintptr(i))) = uint8(i*libc.Uint64FromUint32(17) + libc.Uint64FromUint32(41))
		goto _1
	_1:
		;
		i = i + 1
	}
}

type u8 = uint8

type u32 = uint64

type u64 = uint64

type i64 = int64

type gf = [16]i64

var _0 [16]u8
var _9 = [32]u8{
	0: uint8(9),
}

var gf0 gf
var gf1 = gf{
	0: int64(1),
}
var _121665 = gf{
	0: int64(0xDB41),
	1: int64(1),
}
var D = gf{
	0:  int64(0x78a3),
	1:  int64(0x1359),
	2:  int64(0x4dca),
	3:  int64(0x75eb),
	4:  int64(0xd8ab),
	5:  int64(0x4141),
	6:  int64(0x0a4d),
	7:  int64(0x0070),
	8:  int64(0xe898),
	9:  int64(0x7779),
	10: int64(0x4079),
	11: int64(0x8cc7),
	12: int64(0xfe73),
	13: int64(0x2b6f),
	14: int64(0x6cee),
	15: int64(0x5203),
}
var D2 = gf{
	0:  int64(0xf159),
	1:  int64(0x26b2),
	2:  int64(0x9b94),
	3:  int64(0xebd6),
	4:  int64(0xb156),
	5:  int64(0x8283),
	6:  int64(0x149a),
	7:  int64(0x00e0),
	8:  int64(0xd130),
	9:  int64(0xeef3),
	10: int64(0x80f2),
	11: int64(0x198e),
	12: int64(0xfce7),
	13: int64(0x56df),
	14: int64(0xd9dc),
	15: int64(0x2406),
}
var X = gf{
	0:  int64(0xd51a),
	1:  int64(0x8f25),
	2:  int64(0x2d60),
	3:  int64(0xc956),
	4:  int64(0xa7b2),
	5:  int64(0x9525),
	6:  int64(0xc760),
	7:  int64(0x692c),
	8:  int64(0xdc5c),
	9:  int64(0xfdd6),
	10: int64(0xe231),
	11: int64(0xc0a4),
	12: int64(0x53fe),
	13: int64(0xcd6e),
	14: int64(0x36d3),
	15: int64(0x2169),
}
var Y = gf{
	0:  int64(0x6658),
	1:  int64(0x6666),
	2:  int64(0x6666),
	3:  int64(0x6666),
	4:  int64(0x6666),
	5:  int64(0x6666),
	6:  int64(0x6666),
	7:  int64(0x6666),
	8:  int64(0x6666),
	9:  int64(0x6666),
	10: int64(0x6666),
	11: int64(0x6666),
	12: int64(0x6666),
	13: int64(0x6666),
	14: int64(0x6666),
	15: int64(0x6666),
}
var I = gf{
	0:  int64(0xa0b0),
	1:  int64(0x4a0e),
	2:  int64(0x1b27),
	3:  int64(0xc4ee),
	4:  int64(0xe478),
	5:  int64(0xad2f),
	6:  int64(0x1806),
	7:  int64(0x2f43),
	8:  int64(0xd7a7),
	9:  int64(0x3dfb),
	10: int64(0x0099),
	11: int64(0x2b4d),
	12: int64(0xdf0b),
	13: int64(0x4fc1),
	14: int64(0x2480),
	15: int64(0x2b83),
}

func L32(tls *libc.TLS, x u32, c int32) (r u32) {
	return x<<c | x&uint64(0xffffffff)>>(int32(32)-c)
}

func ld32(tls *libc.TLS, x uintptr) (r u32) {
	var u u32
	_ = u
	u = uint64(**(**u8)(__ccgo_up(x + 3)))
	u = u<<libc.Int32FromInt32(8) | uint64(**(**u8)(__ccgo_up(x + 2)))
	u = u<<libc.Int32FromInt32(8) | uint64(**(**u8)(__ccgo_up(x + 1)))
	return u<<libc.Int32FromInt32(8) | uint64(**(**u8)(__ccgo_up(x)))
}

func dl64(tls *libc.TLS, x uintptr) (r u64) {
	var i, u u64
	_, _ = i, u
	u = uint64(0)
	i = uint64(0)
	for {
		if !(i < uint64(8)) {
			break
		}
		u = u<<libc.Int32FromInt32(8) | uint64(**(**u8)(__ccgo_up(x + uintptr(i))))
		goto _1
	_1:
		;
		i = i + 1
	}
	return u
}

func st32(tls *libc.TLS, x uintptr, u u32) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(4)) {
			break
		}
		**(**u8)(__ccgo_up(x + uintptr(i))) = uint8(u)
		u = u >> uint64(8)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func ts64(tls *libc.TLS, x uintptr, u u64) {
	var i int32
	_ = i
	i = int32(7)
	for {
		if !(i >= 0) {
			break
		}
		**(**u8)(__ccgo_up(x + uintptr(i))) = uint8(u)
		u = u >> uint64(8)
		goto _1
	_1:
		;
		i = i - 1
	}
}

func vn(tls *libc.TLS, x uintptr, y uintptr, n int32) (r int32) {
	var d, i u32
	_, _ = d, i
	d = uint64(0)
	i = uint64(0)
	for {
		if !(i < libc.Uint64FromInt32(n)) {
			break
		}
		d = d | libc.Uint64FromInt32(libc.Int32FromUint8(**(**u8)(__ccgo_up(x + uintptr(i))))^libc.Int32FromUint8(**(**u8)(__ccgo_up(y + uintptr(i)))))
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.Int32FromUint64(uint64(1)&((d-uint64(1))>>int32(8)) - uint64(1))
}

func crypto_verify_16_tweet(tls *libc.TLS, x uintptr, y uintptr) (r int32) {
	return vn(tls, x, y, int32(16))
}

func crypto_verify_32_tweet(tls *libc.TLS, x uintptr, y uintptr) (r int32) {
	return vn(tls, x, y, int32(32))
}

func core(tls *libc.TLS, out uintptr, in uintptr, k uintptr, c uintptr, h int32) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var i, j, m int32
	var w, y [16]u32
	var _ /* t at bp+128 */ [4]u32
	var _ /* x at bp+0 */ [16]u32
	_, _, _, _, _ = i, j, m, w, y
	i = 0
	for {
		if !(i < int32(4)) {
			break
		}
		(**(**[16]u32)(__ccgo_up(bp)))[int32(5)*i] = ld32(tls, c+uintptr(int32(4)*i))
		(**(**[16]u32)(__ccgo_up(bp)))[int32(1)+i] = ld32(tls, k+uintptr(int32(4)*i))
		(**(**[16]u32)(__ccgo_up(bp)))[int32(6)+i] = ld32(tls, in+uintptr(int32(4)*i))
		(**(**[16]u32)(__ccgo_up(bp)))[int32(11)+i] = ld32(tls, k+uintptr(16)+uintptr(int32(4)*i))
		goto _1
	_1:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		y[i] = (**(**[16]u32)(__ccgo_up(bp)))[i]
		goto _2
	_2:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int32(20)) {
			break
		}
		j = 0
		for {
			if !(j < int32(4)) {
				break
			}
			m = 0
			for {
				if !(m < int32(4)) {
					break
				}
				(**(**[4]u32)(__ccgo_up(bp + 128)))[m] = (**(**[16]u32)(__ccgo_up(bp)))[(int32(5)*j+int32(4)*m)%int32(16)]
				goto _5
			_5:
				;
				m = m + 1
			}
			**(**u32)(__ccgo_up(bp + 128 + 1*8)) ^= L32(tls, (**(**[4]u32)(__ccgo_up(bp + 128)))[0]+(**(**[4]u32)(__ccgo_up(bp + 128)))[int32(3)], int32(7))
			**(**u32)(__ccgo_up(bp + 128 + 2*8)) ^= L32(tls, (**(**[4]u32)(__ccgo_up(bp + 128)))[int32(1)]+(**(**[4]u32)(__ccgo_up(bp + 128)))[0], int32(9))
			**(**u32)(__ccgo_up(bp + 128 + 3*8)) ^= L32(tls, (**(**[4]u32)(__ccgo_up(bp + 128)))[int32(2)]+(**(**[4]u32)(__ccgo_up(bp + 128)))[int32(1)], int32(13))
			**(**u32)(__ccgo_up(bp + 128)) ^= L32(tls, (**(**[4]u32)(__ccgo_up(bp + 128)))[int32(3)]+(**(**[4]u32)(__ccgo_up(bp + 128)))[int32(2)], int32(18))
			m = 0
			for {
				if !(m < int32(4)) {
					break
				}
				w[int32(4)*j+(j+m)%int32(4)] = (**(**[4]u32)(__ccgo_up(bp + 128)))[m]
				goto _6
			_6:
				;
				m = m + 1
			}
			goto _4
		_4:
			;
			j = j + 1
		}
		m = 0
		for {
			if !(m < int32(16)) {
				break
			}
			(**(**[16]u32)(__ccgo_up(bp)))[m] = w[m]
			goto _7
		_7:
			;
			m = m + 1
		}
		goto _3
	_3:
		;
		i = i + 1
	}
	if h != 0 {
		i = 0
		for {
			if !(i < int32(16)) {
				break
			}
			**(**u32)(__ccgo_up(bp + uintptr(i)*8)) += y[i]
			goto _8
		_8:
			;
			i = i + 1
		}
		i = 0
		for {
			if !(i < int32(4)) {
				break
			}
			**(**u32)(__ccgo_up(bp + uintptr(int32(5)*i)*8)) -= ld32(tls, c+uintptr(int32(4)*i))
			**(**u32)(__ccgo_up(bp + uintptr(int32(6)+i)*8)) -= ld32(tls, in+uintptr(int32(4)*i))
			goto _9
		_9:
			;
			i = i + 1
		}
		i = 0
		for {
			if !(i < int32(4)) {
				break
			}
			st32(tls, out+uintptr(int32(4)*i), (**(**[16]u32)(__ccgo_up(bp)))[int32(5)*i])
			st32(tls, out+uintptr(16)+uintptr(int32(4)*i), (**(**[16]u32)(__ccgo_up(bp)))[int32(6)+i])
			goto _10
		_10:
			;
			i = i + 1
		}
	} else {
		i = 0
		for {
			if !(i < int32(16)) {
				break
			}
			st32(tls, out+uintptr(int32(4)*i), (**(**[16]u32)(__ccgo_up(bp)))[i]+y[i])
			goto _11
		_11:
			;
			i = i + 1
		}
	}
}

func crypto_core_salsa20_tweet(tls *libc.TLS, out uintptr, in uintptr, k uintptr, c uintptr) (r int32) {
	core(tls, out, in, k, c, 0)
	return 0
}

func crypto_core_hsalsa20_tweet(tls *libc.TLS, out uintptr, in uintptr, k uintptr, c uintptr) (r int32) {
	core(tls, out, in, k, c, int32(1))
	return 0
}

var sigma = [16]u8{'e', 'x', 'p', 'a', 'n', 'd', ' ', '3', '2', '-', 'b', 'y', 't', 'e', ' ', 'k'}

func crypto_stream_salsa20_tweet_xor(tls *libc.TLS, c uintptr, m uintptr, b u64, n uintptr, k uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var i, u u32
	var v4 int32
	var _ /* x at bp+16 */ [64]u8
	var _ /* z at bp+0 */ [16]u8
	_, _, _ = i, u, v4
	if !(b != 0) {
		return 0
	}
	i = uint64(0)
	for {
		if !(i < uint64(16)) {
			break
		}
		(**(**[16]u8)(__ccgo_up(bp)))[i] = uint8(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	i = uint64(0)
	for {
		if !(i < uint64(8)) {
			break
		}
		(**(**[16]u8)(__ccgo_up(bp)))[i] = **(**u8)(__ccgo_up(n + uintptr(i)))
		goto _2
	_2:
		;
		i = i + 1
	}
	for b >= uint64(64) {
		crypto_core_salsa20_tweet(tls, bp+16, bp, k, uintptr(unsafe.Pointer(&sigma)))
		i = uint64(0)
		for {
			if !(i < uint64(64)) {
				break
			}
			if m != 0 {
				v4 = libc.Int32FromUint8(**(**u8)(__ccgo_up(m + uintptr(i))))
			} else {
				v4 = 0
			}
			**(**u8)(__ccgo_up(c + uintptr(i))) = libc.Uint8FromInt32(v4 ^ libc.Int32FromUint8((**(**[64]u8)(__ccgo_up(bp + 16)))[i]))
			goto _3
		_3:
			;
			i = i + 1
		}
		u = uint64(1)
		i = uint64(8)
		for {
			if !(i < uint64(16)) {
				break
			}
			u = u + uint64((**(**[16]u8)(__ccgo_up(bp)))[i])
			(**(**[16]u8)(__ccgo_up(bp)))[i] = uint8(u)
			u = u >> uint64(8)
			goto _5
		_5:
			;
			i = i + 1
		}
		b = b - uint64(64)
		c = c + uintptr(64)
		if m != 0 {
			m = m + uintptr(64)
		}
	}
	if b != 0 {
		crypto_core_salsa20_tweet(tls, bp+16, bp, k, uintptr(unsafe.Pointer(&sigma)))
		i = uint64(0)
		for {
			if !(i < b) {
				break
			}
			if m != 0 {
				v4 = libc.Int32FromUint8(**(**u8)(__ccgo_up(m + uintptr(i))))
			} else {
				v4 = 0
			}
			**(**u8)(__ccgo_up(c + uintptr(i))) = libc.Uint8FromInt32(v4 ^ libc.Int32FromUint8((**(**[64]u8)(__ccgo_up(bp + 16)))[i]))
			goto _6
		_6:
			;
			i = i + 1
		}
	}
	return 0
}

func crypto_stream_salsa20_tweet(tls *libc.TLS, c uintptr, d u64, n uintptr, k uintptr) (r int32) {
	return crypto_stream_salsa20_tweet_xor(tls, c, uintptr(0), d, n, k)
}

func crypto_stream_xsalsa20_tweet(tls *libc.TLS, c uintptr, d u64, n uintptr, k uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* s at bp+0 */ [32]u8
	crypto_core_hsalsa20_tweet(tls, bp, n, k, uintptr(unsafe.Pointer(&sigma)))
	return crypto_stream_salsa20_tweet(tls, c, d, n+uintptr(16), bp)
}

func crypto_stream_xsalsa20_tweet_xor(tls *libc.TLS, c uintptr, m uintptr, d u64, n uintptr, k uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* s at bp+0 */ [32]u8
	crypto_core_hsalsa20_tweet(tls, bp, n, k, uintptr(unsafe.Pointer(&sigma)))
	return crypto_stream_salsa20_tweet_xor(tls, c, m, d, n+uintptr(16), bp)
}

func add1305(tls *libc.TLS, h uintptr, c uintptr) {
	var j, u u32
	_, _ = j, u
	u = uint64(0)
	j = uint64(0)
	for {
		if !(j < uint64(17)) {
			break
		}
		u = u + (**(**u32)(__ccgo_up(h + uintptr(j)*8)) + **(**u32)(__ccgo_up(c + uintptr(j)*8)))
		**(**u32)(__ccgo_up(h + uintptr(j)*8)) = u & uint64(255)
		u = u >> uint64(8)
		goto _1
	_1:
		;
		j = j + 1
	}
}

var minusp = [17]u32{
	0:  uint64(5),
	16: uint64(252),
}

func crypto_onetimeauth_poly1305_tweet(tls *libc.TLS, out uintptr, m uintptr, n u64, k uintptr) (r int32) {
	bp := tls.Alloc(544)
	defer tls.Free(544)
	var g [17]u32
	var i, j, s, u, v2 u32
	var v8 uint64
	var _ /* c at bp+408 */ [17]u32
	var _ /* h at bp+272 */ [17]u32
	var _ /* r at bp+136 */ [17]u32
	var _ /* x at bp+0 */ [17]u32
	_, _, _, _, _, _, _ = g, i, j, s, u, v2, v8
	j = uint64(0)
	for {
		if !(j < uint64(17)) {
			break
		}
		v2 = libc.Uint64FromInt32(0)
		(**(**[17]u32)(__ccgo_up(bp + 272)))[j] = v2
		(**(**[17]u32)(__ccgo_up(bp + 136)))[j] = v2
		goto _1
	_1:
		;
		j = j + 1
	}
	j = uint64(0)
	for {
		if !(j < uint64(16)) {
			break
		}
		(**(**[17]u32)(__ccgo_up(bp + 136)))[j] = uint64(**(**u8)(__ccgo_up(k + uintptr(j))))
		goto _3
	_3:
		;
		j = j + 1
	}
	**(**u32)(__ccgo_up(bp + 136 + 3*8)) &= uint64(15)
	**(**u32)(__ccgo_up(bp + 136 + 4*8)) &= uint64(252)
	**(**u32)(__ccgo_up(bp + 136 + 7*8)) &= uint64(15)
	**(**u32)(__ccgo_up(bp + 136 + 8*8)) &= uint64(252)
	**(**u32)(__ccgo_up(bp + 136 + 11*8)) &= uint64(15)
	**(**u32)(__ccgo_up(bp + 136 + 12*8)) &= uint64(252)
	**(**u32)(__ccgo_up(bp + 136 + 15*8)) &= uint64(15)
	for n > uint64(0) {
		j = uint64(0)
		for {
			if !(j < uint64(17)) {
				break
			}
			(**(**[17]u32)(__ccgo_up(bp + 408)))[j] = uint64(0)
			goto _4
		_4:
			;
			j = j + 1
		}
		j = uint64(0)
		for {
			if !(j < uint64(16) && j < n) {
				break
			}
			(**(**[17]u32)(__ccgo_up(bp + 408)))[j] = uint64(**(**u8)(__ccgo_up(m + uintptr(j))))
			goto _5
		_5:
			;
			j = j + 1
		}
		(**(**[17]u32)(__ccgo_up(bp + 408)))[j] = uint64(1)
		m = m + uintptr(j)
		n = n - j
		add1305(tls, bp+272, bp+408)
		i = uint64(0)
		for {
			if !(i < uint64(17)) {
				break
			}
			(**(**[17]u32)(__ccgo_up(bp)))[i] = uint64(0)
			j = uint64(0)
			for {
				if !(j < uint64(17)) {
					break
				}
				if j <= i {
					v8 = (**(**[17]u32)(__ccgo_up(bp + 136)))[i-j]
				} else {
					v8 = uint64(320) * (**(**[17]u32)(__ccgo_up(bp + 136)))[i+uint64(17)-j]
				}
				**(**u32)(__ccgo_up(bp + uintptr(i)*8)) += (**(**[17]u32)(__ccgo_up(bp + 272)))[j] * v8
				goto _7
			_7:
				;
				j = j + 1
			}
			goto _6
		_6:
			;
			i = i + 1
		}
		i = uint64(0)
		for {
			if !(i < uint64(17)) {
				break
			}
			(**(**[17]u32)(__ccgo_up(bp + 272)))[i] = (**(**[17]u32)(__ccgo_up(bp)))[i]
			goto _9
		_9:
			;
			i = i + 1
		}
		u = uint64(0)
		j = uint64(0)
		for {
			if !(j < uint64(16)) {
				break
			}
			u = u + (**(**[17]u32)(__ccgo_up(bp + 272)))[j]
			(**(**[17]u32)(__ccgo_up(bp + 272)))[j] = u & uint64(255)
			u = u >> uint64(8)
			goto _10
		_10:
			;
			j = j + 1
		}
		u = u + (**(**[17]u32)(__ccgo_up(bp + 272)))[int32(16)]
		(**(**[17]u32)(__ccgo_up(bp + 272)))[int32(16)] = u & uint64(3)
		u = uint64(5) * (u >> libc.Int32FromInt32(2))
		j = uint64(0)
		for {
			if !(j < uint64(16)) {
				break
			}
			u = u + (**(**[17]u32)(__ccgo_up(bp + 272)))[j]
			(**(**[17]u32)(__ccgo_up(bp + 272)))[j] = u & uint64(255)
			u = u >> uint64(8)
			goto _11
		_11:
			;
			j = j + 1
		}
		u = u + (**(**[17]u32)(__ccgo_up(bp + 272)))[int32(16)]
		(**(**[17]u32)(__ccgo_up(bp + 272)))[int32(16)] = u
	}
	j = uint64(0)
	for {
		if !(j < uint64(17)) {
			break
		}
		g[j] = (**(**[17]u32)(__ccgo_up(bp + 272)))[j]
		goto _12
	_12:
		;
		j = j + 1
	}
	add1305(tls, bp+272, uintptr(unsafe.Pointer(&minusp)))
	s = -((**(**[17]u32)(__ccgo_up(bp + 272)))[int32(16)] >> libc.Int32FromInt32(7))
	j = uint64(0)
	for {
		if !(j < uint64(17)) {
			break
		}
		**(**u32)(__ccgo_up(bp + 272 + uintptr(j)*8)) ^= s & (g[j] ^ (**(**[17]u32)(__ccgo_up(bp + 272)))[j])
		goto _13
	_13:
		;
		j = j + 1
	}
	j = uint64(0)
	for {
		if !(j < uint64(16)) {
			break
		}
		(**(**[17]u32)(__ccgo_up(bp + 408)))[j] = uint64(**(**u8)(__ccgo_up(k + uintptr(j+uint64(16)))))
		goto _14
	_14:
		;
		j = j + 1
	}
	(**(**[17]u32)(__ccgo_up(bp + 408)))[int32(16)] = uint64(0)
	add1305(tls, bp+272, bp+408)
	j = uint64(0)
	for {
		if !(j < uint64(16)) {
			break
		}
		**(**u8)(__ccgo_up(out + uintptr(j))) = uint8((**(**[17]u32)(__ccgo_up(bp + 272)))[j])
		goto _15
	_15:
		;
		j = j + 1
	}
	return 0
}

func crypto_onetimeauth_poly1305_tweet_verify(tls *libc.TLS, h uintptr, m uintptr, n u64, k uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* x at bp+0 */ [16]u8
	crypto_onetimeauth_poly1305_tweet(tls, bp, m, n, k)
	return crypto_verify_16_tweet(tls, h, bp)
}

func crypto_secretbox_xsalsa20poly1305_tweet(tls *libc.TLS, c uintptr, m uintptr, d u64, n uintptr, k uintptr) (r int32) {
	var i int32
	_ = i
	if d < uint64(32) {
		return -int32(1)
	}
	crypto_stream_xsalsa20_tweet_xor(tls, c, m, d, n, k)
	crypto_onetimeauth_poly1305_tweet(tls, c+uintptr(16), c+uintptr(32), d-uint64(32), c)
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		**(**u8)(__ccgo_up(c + uintptr(i))) = uint8(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	return 0
}

func crypto_secretbox_xsalsa20poly1305_tweet_open(tls *libc.TLS, m uintptr, c uintptr, d u64, n uintptr, k uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var _ /* x at bp+0 */ [32]u8
	_ = i
	if d < uint64(32) {
		return -int32(1)
	}
	crypto_stream_xsalsa20_tweet(tls, bp, uint64(32), n, k)
	if crypto_onetimeauth_poly1305_tweet_verify(tls, c+uintptr(16), c+uintptr(32), d-uint64(32), bp) != 0 {
		return -int32(1)
	}
	crypto_stream_xsalsa20_tweet_xor(tls, m, c, d, n, k)
	i = 0
	for {
		if !(i < int32(32)) {
			break
		}
		**(**u8)(__ccgo_up(m + uintptr(i))) = uint8(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	return 0
}

func set25519(tls *libc.TLS, r uintptr, a uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		**(**i64)(__ccgo_up(r + uintptr(i)*8)) = **(**i64)(__ccgo_up(a + uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func car25519(tls *libc.TLS, o uintptr) {
	var c i64
	var i int32
	_, _ = c, i
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		**(**i64)(__ccgo_up(o + uintptr(i)*8)) += libc.Int64FromInt64(1) << libc.Int32FromInt32(16)
		c = **(**i64)(__ccgo_up(o + uintptr(i)*8)) >> int32(16)
		**(**i64)(__ccgo_up(o + uintptr((i+int32(1))*libc.BoolInt32(i < int32(15)))*8)) += c - int64(1) + int64(37)*(c-int64(1))*libc.BoolInt64(i == libc.Int32FromInt32(15))
		**(**i64)(__ccgo_up(o + uintptr(i)*8)) -= c << int32(16)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func sel25519(tls *libc.TLS, p uintptr, q uintptr, b int32) {
	var c, i, t i64
	_, _, _ = c, i, t
	c = int64(^(b - libc.Int32FromInt32(1)))
	i = 0
	for {
		if !(i < int64(16)) {
			break
		}
		t = c & (**(**i64)(__ccgo_up(p + uintptr(i)*8)) ^ **(**i64)(__ccgo_up(q + uintptr(i)*8)))
		**(**i64)(__ccgo_up(p + uintptr(i)*8)) ^= t
		**(**i64)(__ccgo_up(q + uintptr(i)*8)) ^= t
		goto _1
	_1:
		;
		i = i + 1
	}
}

func pack25519(tls *libc.TLS, o uintptr, n uintptr) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var b, i, j int32
	var _ /* m at bp+0 */ gf
	var _ /* t at bp+128 */ gf
	_, _, _ = b, i, j
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		(**(**gf)(__ccgo_up(bp + 128)))[i] = **(**i64)(__ccgo_up(n + uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
	car25519(tls, bp+128)
	car25519(tls, bp+128)
	car25519(tls, bp+128)
	j = 0
	for {
		if !(j < int32(2)) {
			break
		}
		(**(**gf)(__ccgo_up(bp)))[0] = (**(**gf)(__ccgo_up(bp + 128)))[0] - int64(0xffed)
		i = int32(1)
		for {
			if !(i < int32(15)) {
				break
			}
			(**(**gf)(__ccgo_up(bp)))[i] = (**(**gf)(__ccgo_up(bp + 128)))[i] - int64(0xffff) - (**(**gf)(__ccgo_up(bp)))[i-int32(1)]>>libc.Int32FromInt32(16)&int64(1)
			**(**i64)(__ccgo_up(bp + uintptr(i-int32(1))*8)) &= int64(0xffff)
			goto _3
		_3:
			;
			i = i + 1
		}
		(**(**gf)(__ccgo_up(bp)))[int32(15)] = (**(**gf)(__ccgo_up(bp + 128)))[int32(15)] - int64(0x7fff) - (**(**gf)(__ccgo_up(bp)))[int32(14)]>>libc.Int32FromInt32(16)&int64(1)
		b = int32((**(**gf)(__ccgo_up(bp)))[int32(15)] >> libc.Int32FromInt32(16) & int64(1))
		**(**i64)(__ccgo_up(bp + 14*8)) &= int64(0xffff)
		sel25519(tls, bp+128, bp, int32(1)-b)
		goto _2
	_2:
		;
		j = j + 1
	}
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		**(**u8)(__ccgo_up(o + uintptr(int32(2)*i))) = libc.Uint8FromInt64((**(**gf)(__ccgo_up(bp + 128)))[i] & int64(0xff))
		**(**u8)(__ccgo_up(o + uintptr(int32(2)*i+int32(1)))) = libc.Uint8FromInt64((**(**gf)(__ccgo_up(bp + 128)))[i] >> int32(8))
		goto _4
	_4:
		;
		i = i + 1
	}
}

func neq25519(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* c at bp+0 */ [32]u8
	var _ /* d at bp+32 */ [32]u8
	pack25519(tls, bp, a)
	pack25519(tls, bp+32, b)
	return crypto_verify_32_tweet(tls, bp, bp+32)
}

func par25519(tls *libc.TLS, a uintptr) (r u8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* d at bp+0 */ [32]u8
	pack25519(tls, bp, a)
	return libc.Uint8FromInt32(libc.Int32FromUint8((**(**[32]u8)(__ccgo_up(bp)))[0]) & int32(1))
}

func unpack25519(tls *libc.TLS, o uintptr, n uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		**(**i64)(__ccgo_up(o + uintptr(i)*8)) = libc.Int64FromUint8(**(**u8)(__ccgo_up(n + uintptr(int32(2)*i)))) + libc.Int64FromUint8(**(**u8)(__ccgo_up(n + uintptr(int32(2)*i+int32(1)))))<<libc.Int32FromInt32(8)
		goto _1
	_1:
		;
		i = i + 1
	}
	**(**i64)(__ccgo_up(o + 15*8)) &= int64(0x7fff)
}

func A(tls *libc.TLS, o uintptr, a uintptr, b uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		**(**i64)(__ccgo_up(o + uintptr(i)*8)) = **(**i64)(__ccgo_up(a + uintptr(i)*8)) + **(**i64)(__ccgo_up(b + uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Z(tls *libc.TLS, o uintptr, a uintptr, b uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(16)) {
			break
		}
		**(**i64)(__ccgo_up(o + uintptr(i)*8)) = **(**i64)(__ccgo_up(a + uintptr(i)*8)) - **(**i64)(__ccgo_up(b + uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func M(tls *libc.TLS, o uintptr, a uintptr, b uintptr) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var i, j i64
	var _ /* t at bp+0 */ [31]i64
	_, _ = i, j
	i = 0
	for {
		if !(i < int64(31)) {
			break
		}
		(**(**[31]i64)(__ccgo_up(bp)))[i] = 0
		goto _1
	_1:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int64(16)) {
			break
		}
		j = 0
		for {
			if !(j < int64(16)) {
				break
			}
			**(**i64)(__ccgo_up(bp + uintptr(i+j)*8)) += **(**i64)(__ccgo_up(a + uintptr(i)*8)) * **(**i64)(__ccgo_up(b + uintptr(j)*8))
			goto _3
		_3:
			;
			j = j + 1
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int64(15)) {
			break
		}
		**(**i64)(__ccgo_up(bp + uintptr(i)*8)) += int64(38) * (**(**[31]i64)(__ccgo_up(bp)))[i+int64(16)]
		goto _4
	_4:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int64(16)) {
			break
		}
		**(**i64)(__ccgo_up(o + uintptr(i)*8)) = (**(**[31]i64)(__ccgo_up(bp)))[i]
		goto _5
	_5:
		;
		i = i + 1
	}
	car25519(tls, o)
	car25519(tls, o)
}

func S(tls *libc.TLS, o uintptr, a uintptr) {
	M(tls, o, a, a)
}

func inv25519(tls *libc.TLS, o uintptr, i uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var a int32
	var _ /* c at bp+0 */ gf
	_ = a
	a = 0
	for {
		if !(a < int32(16)) {
			break
		}
		(**(**gf)(__ccgo_up(bp)))[a] = **(**i64)(__ccgo_up(i + uintptr(a)*8))
		goto _1
	_1:
		;
		a = a + 1
	}
	a = int32(253)
	for {
		if !(a >= 0) {
			break
		}
		S(tls, bp, bp)
		if a != int32(2) && a != int32(4) {
			M(tls, bp, bp, i)
		}
		goto _2
	_2:
		;
		a = a - 1
	}
	a = 0
	for {
		if !(a < int32(16)) {
			break
		}
		**(**i64)(__ccgo_up(o + uintptr(a)*8)) = (**(**gf)(__ccgo_up(bp)))[a]
		goto _3
	_3:
		;
		a = a + 1
	}
}

func pow2523(tls *libc.TLS, o uintptr, i uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var a int32
	var _ /* c at bp+0 */ gf
	_ = a
	a = 0
	for {
		if !(a < int32(16)) {
			break
		}
		(**(**gf)(__ccgo_up(bp)))[a] = **(**i64)(__ccgo_up(i + uintptr(a)*8))
		goto _1
	_1:
		;
		a = a + 1
	}
	a = int32(250)
	for {
		if !(a >= 0) {
			break
		}
		S(tls, bp, bp)
		if a != int32(1) {
			M(tls, bp, bp, i)
		}
		goto _2
	_2:
		;
		a = a - 1
	}
	a = 0
	for {
		if !(a < int32(16)) {
			break
		}
		**(**i64)(__ccgo_up(o + uintptr(a)*8)) = (**(**gf)(__ccgo_up(bp)))[a]
		goto _3
	_3:
		;
		a = a + 1
	}
}

func crypto_scalarmult_curve25519_tweet(tls *libc.TLS, q uintptr, n uintptr, p uintptr) (r1 int32) {
	bp := tls.Alloc(1440)
	defer tls.Free(1440)
	var i, r, v4, v5 i64
	var v2 uintptr
	var _ /* a at bp+672 */ gf
	var _ /* b at bp+800 */ gf
	var _ /* c at bp+928 */ gf
	var _ /* d at bp+1056 */ gf
	var _ /* e at bp+1184 */ gf
	var _ /* f at bp+1312 */ gf
	var _ /* x at bp+32 */ [80]i64
	var _ /* z at bp+0 */ [32]u8
	_, _, _, _, _ = i, r, v2, v4, v5
	i = 0
	for {
		if !(i < int64(31)) {
			break
		}
		(**(**[32]u8)(__ccgo_up(bp)))[i] = **(**u8)(__ccgo_up(n + uintptr(i)))
		goto _1
	_1:
		;
		i = i + 1
	}
	(**(**[32]u8)(__ccgo_up(bp)))[int32(31)] = libc.Uint8FromInt32(libc.Int32FromUint8(**(**u8)(__ccgo_up(n + 31)))&int32(127) | int32(64))
	v2 = bp
	*(*u8)(unsafe.Pointer(v2)) = u8(int32(*(*u8)(unsafe.Pointer(v2))) & libc.Int32FromInt32(248))
	unpack25519(tls, bp+32, p)
	i = 0
	for {
		if !(i < int64(16)) {
			break
		}
		(**(**gf)(__ccgo_up(bp + 800)))[i] = (**(**[80]i64)(__ccgo_up(bp + 32)))[i]
		v5 = libc.Int64FromInt32(0)
		(**(**gf)(__ccgo_up(bp + 928)))[i] = v5
		v4 = v5
		(**(**gf)(__ccgo_up(bp + 672)))[i] = v4
		(**(**gf)(__ccgo_up(bp + 1056)))[i] = v4
		goto _3
	_3:
		;
		i = i + 1
	}
	v4 = libc.Int64FromInt32(1)
	(**(**gf)(__ccgo_up(bp + 1056)))[0] = v4
	(**(**gf)(__ccgo_up(bp + 672)))[0] = v4
	i = int64(254)
	for {
		if !(i >= 0) {
			break
		}
		r = int64(libc.Int32FromUint8((**(**[32]u8)(__ccgo_up(bp)))[i>>int32(3)]) >> (i & int64(7)) & int32(1))
		sel25519(tls, bp+672, bp+800, int32(r))
		sel25519(tls, bp+928, bp+1056, int32(r))
		A(tls, bp+1184, bp+672, bp+928)
		Z(tls, bp+672, bp+672, bp+928)
		A(tls, bp+928, bp+800, bp+1056)
		Z(tls, bp+800, bp+800, bp+1056)
		S(tls, bp+1056, bp+1184)
		S(tls, bp+1312, bp+672)
		M(tls, bp+672, bp+928, bp+672)
		M(tls, bp+928, bp+800, bp+1184)
		A(tls, bp+1184, bp+672, bp+928)
		Z(tls, bp+672, bp+672, bp+928)
		S(tls, bp+800, bp+672)
		Z(tls, bp+928, bp+1056, bp+1312)
		M(tls, bp+672, bp+928, uintptr(unsafe.Pointer(&_121665)))
		A(tls, bp+672, bp+672, bp+1056)
		M(tls, bp+928, bp+928, bp+672)
		M(tls, bp+672, bp+1056, bp+1312)
		M(tls, bp+1056, bp+800, bp+32)
		S(tls, bp+800, bp+1184)
		sel25519(tls, bp+672, bp+800, int32(r))
		sel25519(tls, bp+928, bp+1056, int32(r))
		goto _7
	_7:
		;
		i = i - 1
	}
	i = 0
	for {
		if !(i < int64(16)) {
			break
		}
		(**(**[80]i64)(__ccgo_up(bp + 32)))[i+int64(16)] = (**(**gf)(__ccgo_up(bp + 672)))[i]
		(**(**[80]i64)(__ccgo_up(bp + 32)))[i+int64(32)] = (**(**gf)(__ccgo_up(bp + 928)))[i]
		(**(**[80]i64)(__ccgo_up(bp + 32)))[i+int64(48)] = (**(**gf)(__ccgo_up(bp + 800)))[i]
		(**(**[80]i64)(__ccgo_up(bp + 32)))[i+int64(64)] = (**(**gf)(__ccgo_up(bp + 1056)))[i]
		goto _8
	_8:
		;
		i = i + 1
	}
	inv25519(tls, bp+32+uintptr(32)*8, bp+32+uintptr(32)*8)
	M(tls, bp+32+uintptr(16)*8, bp+32+uintptr(16)*8, bp+32+uintptr(32)*8)
	pack25519(tls, q, bp+32+uintptr(16)*8)
	return 0
}

func crypto_scalarmult_curve25519_tweet_base(tls *libc.TLS, q uintptr, n uintptr) (r int32) {
	return crypto_scalarmult_curve25519_tweet(tls, q, n, uintptr(unsafe.Pointer(&_9)))
}

func crypto_box_curve25519xsalsa20poly1305_tweet_keypair(tls *libc.TLS, y uintptr, x uintptr) (r int32) {
	randombytes(tls, x, uint64(32))
	return crypto_scalarmult_curve25519_tweet_base(tls, y, x)
}

func crypto_box_curve25519xsalsa20poly1305_tweet_beforenm(tls *libc.TLS, k uintptr, y uintptr, x uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* s at bp+0 */ [32]u8
	crypto_scalarmult_curve25519_tweet(tls, bp, x, y)
	return crypto_core_hsalsa20_tweet(tls, k, uintptr(unsafe.Pointer(&_0)), bp, uintptr(unsafe.Pointer(&sigma)))
}

func crypto_box_curve25519xsalsa20poly1305_tweet_afternm(tls *libc.TLS, c uintptr, m uintptr, d u64, n uintptr, k uintptr) (r int32) {
	return crypto_secretbox_xsalsa20poly1305_tweet(tls, c, m, d, n, k)
}

func crypto_box_curve25519xsalsa20poly1305_tweet_open_afternm(tls *libc.TLS, m uintptr, c uintptr, d u64, n uintptr, k uintptr) (r int32) {
	return crypto_secretbox_xsalsa20poly1305_tweet_open(tls, m, c, d, n, k)
}

func crypto_box_curve25519xsalsa20poly1305_tweet(tls *libc.TLS, c uintptr, m uintptr, d u64, n uintptr, y uintptr, x uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* k at bp+0 */ [32]u8
	crypto_box_curve25519xsalsa20poly1305_tweet_beforenm(tls, bp, y, x)
	return crypto_box_curve25519xsalsa20poly1305_tweet_afternm(tls, c, m, d, n, bp)
}

func crypto_box_curve25519xsalsa20poly1305_tweet_open(tls *libc.TLS, m uintptr, c uintptr, d u64, n uintptr, y uintptr, x uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* k at bp+0 */ [32]u8
	crypto_box_curve25519xsalsa20poly1305_tweet_beforenm(tls, bp, y, x)
	return crypto_box_curve25519xsalsa20poly1305_tweet_open_afternm(tls, m, c, d, n, bp)
}

func R(tls *libc.TLS, x u64, c int32) (r u64) {
	return x>>c | x<<(libc.Int32FromInt32(64)-c)
}

func Ch(tls *libc.TLS, x u64, y u64, z u64) (r u64) {
	return x&y ^ ^x&z
}

func Maj(tls *libc.TLS, x u64, y u64, z u64) (r u64) {
	return x&y ^ x&z ^ y&z
}

func Sigma0(tls *libc.TLS, x u64) (r u64) {
	return R(tls, x, int32(28)) ^ R(tls, x, int32(34)) ^ R(tls, x, int32(39))
}

func Sigma1(tls *libc.TLS, x u64) (r u64) {
	return R(tls, x, int32(14)) ^ R(tls, x, int32(18)) ^ R(tls, x, int32(41))
}

func sigma0(tls *libc.TLS, x u64) (r u64) {
	return R(tls, x, int32(1)) ^ R(tls, x, int32(8)) ^ x>>libc.Int32FromInt32(7)
}

func sigma1(tls *libc.TLS, x u64) (r u64) {
	return R(tls, x, int32(19)) ^ R(tls, x, int32(61)) ^ x>>libc.Int32FromInt32(6)
}

var K = [80]u64{
	0:  uint64(0x428a2f98d728ae22),
	1:  uint64(0x7137449123ef65cd),
	2:  uint64(0xb5c0fbcfec4d3b2f),
	3:  uint64(0xe9b5dba58189dbbc),
	4:  uint64(0x3956c25bf348b538),
	5:  uint64(0x59f111f1b605d019),
	6:  uint64(0x923f82a4af194f9b),
	7:  uint64(0xab1c5ed5da6d8118),
	8:  uint64(0xd807aa98a3030242),
	9:  uint64(0x12835b0145706fbe),
	10: uint64(0x243185be4ee4b28c),
	11: uint64(0x550c7dc3d5ffb4e2),
	12: uint64(0x72be5d74f27b896f),
	13: uint64(0x80deb1fe3b1696b1),
	14: uint64(0x9bdc06a725c71235),
	15: uint64(0xc19bf174cf692694),
	16: uint64(0xe49b69c19ef14ad2),
	17: uint64(0xefbe4786384f25e3),
	18: uint64(0x0fc19dc68b8cd5b5),
	19: uint64(0x240ca1cc77ac9c65),
	20: uint64(0x2de92c6f592b0275),
	21: uint64(0x4a7484aa6ea6e483),
	22: uint64(0x5cb0a9dcbd41fbd4),
	23: uint64(0x76f988da831153b5),
	24: uint64(0x983e5152ee66dfab),
	25: uint64(0xa831c66d2db43210),
	26: uint64(0xb00327c898fb213f),
	27: uint64(0xbf597fc7beef0ee4),
	28: uint64(0xc6e00bf33da88fc2),
	29: uint64(0xd5a79147930aa725),
	30: uint64(0x06ca6351e003826f),
	31: uint64(0x142929670a0e6e70),
	32: uint64(0x27b70a8546d22ffc),
	33: uint64(0x2e1b21385c26c926),
	34: uint64(0x4d2c6dfc5ac42aed),
	35: uint64(0x53380d139d95b3df),
	36: uint64(0x650a73548baf63de),
	37: uint64(0x766a0abb3c77b2a8),
	38: uint64(0x81c2c92e47edaee6),
	39: uint64(0x92722c851482353b),
	40: uint64(0xa2bfe8a14cf10364),
	41: uint64(0xa81a664bbc423001),
	42: uint64(0xc24b8b70d0f89791),
	43: uint64(0xc76c51a30654be30),
	44: uint64(0xd192e819d6ef5218),
	45: uint64(0xd69906245565a910),
	46: uint64(0xf40e35855771202a),
	47: uint64(0x106aa07032bbd1b8),
	48: uint64(0x19a4c116b8d2d0c8),
	49: uint64(0x1e376c085141ab53),
	50: uint64(0x2748774cdf8eeb99),
	51: uint64(0x34b0bcb5e19b48a8),
	52: uint64(0x391c0cb3c5c95a63),
	53: uint64(0x4ed8aa4ae3418acb),
	54: uint64(0x5b9cca4f7763e373),
	55: uint64(0x682e6ff3d6b2b8a3),
	56: uint64(0x748f82ee5defb2fc),
	57: uint64(0x78a5636f43172f60),
	58: uint64(0x84c87814a1f0ab72),
	59: uint64(0x8cc702081a6439ec),
	60: uint64(0x90befffa23631e28),
	61: uint64(0xa4506cebde82bde9),
	62: uint64(0xbef9a3f7b2c67915),
	63: uint64(0xc67178f2e372532b),
	64: uint64(0xca273eceea26619c),
	65: uint64(0xd186b8c721c0c207),
	66: uint64(0xeada7dd6cde0eb1e),
	67: uint64(0xf57d4f7fee6ed178),
	68: uint64(0x06f067aa72176fba),
	69: uint64(0x0a637dc5a2c898a6),
	70: uint64(0x113f9804bef90dae),
	71: uint64(0x1b710b35131c471b),
	72: uint64(0x28db77f523047d84),
	73: uint64(0x32caab7b40c72493),
	74: uint64(0x3c9ebe0a15c9bebc),
	75: uint64(0x431d67c49c100d4c),
	76: uint64(0x4cc5d4becb3e42b6),
	77: uint64(0x597f299cfc657e2a),
	78: uint64(0x5fcb6fab3ad6faec),
	79: uint64(0x6c44198c4a475817),
}

func crypto_hashblocks_sha512_tweet(tls *libc.TLS, x uintptr, m uintptr, n u64) (r int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var i, j int32
	var t, v2 u64
	var z [8]u64
	var _ /* a at bp+64 */ [8]u64
	var _ /* b at bp+0 */ [8]u64
	var _ /* w at bp+128 */ [16]u64
	_, _, _, _, _ = i, j, t, z, v2
	i = 0
	for {
		if !(i < int32(8)) {
			break
		}
		v2 = dl64(tls, x+uintptr(int32(8)*i))
		(**(**[8]u64)(__ccgo_up(bp + 64)))[i] = v2
		z[i] = v2
		goto _1
	_1:
		;
		i = i + 1
	}
	for n >= uint64(128) {
		i = 0
		for {
			if !(i < int32(16)) {
				break
			}
			(**(**[16]u64)(__ccgo_up(bp + 128)))[i] = dl64(tls, m+uintptr(int32(8)*i))
			goto _3
		_3:
			;
			i = i + 1
		}
		i = 0
		for {
			if !(i < int32(80)) {
				break
			}
			j = 0
			for {
				if !(j < int32(8)) {
					break
				}
				(**(**[8]u64)(__ccgo_up(bp)))[j] = (**(**[8]u64)(__ccgo_up(bp + 64)))[j]
				goto _5
			_5:
				;
				j = j + 1
			}
			t = (**(**[8]u64)(__ccgo_up(bp + 64)))[int32(7)] + Sigma1(tls, (**(**[8]u64)(__ccgo_up(bp + 64)))[int32(4)]) + Ch(tls, (**(**[8]u64)(__ccgo_up(bp + 64)))[int32(4)], (**(**[8]u64)(__ccgo_up(bp + 64)))[int32(5)], (**(**[8]u64)(__ccgo_up(bp + 64)))[int32(6)]) + K[i] + (**(**[16]u64)(__ccgo_up(bp + 128)))[i%int32(16)]
			(**(**[8]u64)(__ccgo_up(bp)))[int32(7)] = t + Sigma0(tls, (**(**[8]u64)(__ccgo_up(bp + 64)))[0]) + Maj(tls, (**(**[8]u64)(__ccgo_up(bp + 64)))[0], (**(**[8]u64)(__ccgo_up(bp + 64)))[int32(1)], (**(**[8]u64)(__ccgo_up(bp + 64)))[int32(2)])
			**(**u64)(__ccgo_up(bp + 3*8)) += t
			j = 0
			for {
				if !(j < int32(8)) {
					break
				}
				(**(**[8]u64)(__ccgo_up(bp + 64)))[(j+int32(1))%int32(8)] = (**(**[8]u64)(__ccgo_up(bp)))[j]
				goto _6
			_6:
				;
				j = j + 1
			}
			if i%int32(16) == int32(15) {
				j = 0
				for {
					if !(j < int32(16)) {
						break
					}
					**(**u64)(__ccgo_up(bp + 128 + uintptr(j)*8)) += (**(**[16]u64)(__ccgo_up(bp + 128)))[(j+int32(9))%int32(16)] + sigma0(tls, (**(**[16]u64)(__ccgo_up(bp + 128)))[(j+int32(1))%int32(16)]) + sigma1(tls, (**(**[16]u64)(__ccgo_up(bp + 128)))[(j+int32(14))%int32(16)])
					goto _7
				_7:
					;
					j = j + 1
				}
			}
			goto _4
		_4:
			;
			i = i + 1
		}
		i = 0
		for {
			if !(i < int32(8)) {
				break
			}
			**(**u64)(__ccgo_up(bp + 64 + uintptr(i)*8)) += z[i]
			z[i] = (**(**[8]u64)(__ccgo_up(bp + 64)))[i]
			goto _8
		_8:
			;
			i = i + 1
		}
		m = m + uintptr(128)
		n = n - uint64(128)
	}
	i = 0
	for {
		if !(i < int32(8)) {
			break
		}
		ts64(tls, x+uintptr(int32(8)*i), z[i])
		goto _9
	_9:
		;
		i = i + 1
	}
	return libc.Int32FromUint64(n)
}

var iv = [64]u8{
	0:  uint8(0x6a),
	1:  uint8(0x09),
	2:  uint8(0xe6),
	3:  uint8(0x67),
	4:  uint8(0xf3),
	5:  uint8(0xbc),
	6:  uint8(0xc9),
	7:  uint8(0x08),
	8:  uint8(0xbb),
	9:  uint8(0x67),
	10: uint8(0xae),
	11: uint8(0x85),
	12: uint8(0x84),
	13: uint8(0xca),
	14: uint8(0xa7),
	15: uint8(0x3b),
	16: uint8(0x3c),
	17: uint8(0x6e),
	18: uint8(0xf3),
	19: uint8(0x72),
	20: uint8(0xfe),
	21: uint8(0x94),
	22: uint8(0xf8),
	23: uint8(0x2b),
	24: uint8(0xa5),
	25: uint8(0x4f),
	26: uint8(0xf5),
	27: uint8(0x3a),
	28: uint8(0x5f),
	29: uint8(0x1d),
	30: uint8(0x36),
	31: uint8(0xf1),
	32: uint8(0x51),
	33: uint8(0x0e),
	34: uint8(0x52),
	35: uint8(0x7f),
	36: uint8(0xad),
	37: uint8(0xe6),
	38: uint8(0x82),
	39: uint8(0xd1),
	40: uint8(0x9b),
	41: uint8(0x05),
	42: uint8(0x68),
	43: uint8(0x8c),
	44: uint8(0x2b),
	45: uint8(0x3e),
	46: uint8(0x6c),
	47: uint8(0x1f),
	48: uint8(0x1f),
	49: uint8(0x83),
	50: uint8(0xd9),
	51: uint8(0xab),
	52: uint8(0xfb),
	53: uint8(0x41),
	54: uint8(0xbd),
	55: uint8(0x6b),
	56: uint8(0x5b),
	57: uint8(0xe0),
	58: uint8(0xcd),
	59: uint8(0x19),
	60: uint8(0x13),
	61: uint8(0x7e),
	62: uint8(0x21),
	63: uint8(0x79),
}

func crypto_hash_sha512_tweet(tls *libc.TLS, out uintptr, m uintptr, n u64) (r int32) {
	bp := tls.Alloc(320)
	defer tls.Free(320)
	var b, i u64
	var _ /* h at bp+0 */ [64]u8
	var _ /* x at bp+64 */ [256]u8
	_, _ = b, i
	b = n
	i = uint64(0)
	for {
		if !(i < uint64(64)) {
			break
		}
		(**(**[64]u8)(__ccgo_up(bp)))[i] = iv[i]
		goto _1
	_1:
		;
		i = i + 1
	}
	crypto_hashblocks_sha512_tweet(tls, bp, m, n)
	m = m + uintptr(n)
	n = n & uint64(127)
	m = m - uintptr(n)
	i = uint64(0)
	for {
		if !(i < uint64(256)) {
			break
		}
		(**(**[256]u8)(__ccgo_up(bp + 64)))[i] = uint8(0)
		goto _2
	_2:
		;
		i = i + 1
	}
	i = uint64(0)
	for {
		if !(i < n) {
			break
		}
		(**(**[256]u8)(__ccgo_up(bp + 64)))[i] = **(**u8)(__ccgo_up(m + uintptr(i)))
		goto _3
	_3:
		;
		i = i + 1
	}
	(**(**[256]u8)(__ccgo_up(bp + 64)))[n] = uint8(128)
	n = libc.Uint64FromInt32(int32(256) - int32(128)*libc.BoolInt32(n < uint64(112)))
	(**(**[256]u8)(__ccgo_up(bp + 64)))[n-uint64(9)] = uint8(b >> int32(61))
	ts64(tls, bp+64+uintptr(n)-uintptr(8), b<<int32(3))
	crypto_hashblocks_sha512_tweet(tls, bp, bp+64, n)
	i = uint64(0)
	for {
		if !(i < uint64(64)) {
			break
		}
		**(**u8)(__ccgo_up(out + uintptr(i))) = (**(**[64]u8)(__ccgo_up(bp)))[i]
		goto _4
	_4:
		;
		i = i + 1
	}
	return 0
}

func add(tls *libc.TLS, p uintptr, q uintptr) {
	bp := tls.Alloc(1152)
	defer tls.Free(1152)
	var _ /* a at bp+0 */ gf
	var _ /* b at bp+128 */ gf
	var _ /* c at bp+256 */ gf
	var _ /* d at bp+384 */ gf
	var _ /* e at bp+640 */ gf
	var _ /* f at bp+768 */ gf
	var _ /* g at bp+896 */ gf
	var _ /* h at bp+1024 */ gf
	var _ /* t at bp+512 */ gf
	Z(tls, bp, p+1*128, p)
	Z(tls, bp+512, q+1*128, q)
	M(tls, bp, bp, bp+512)
	A(tls, bp+128, p, p+1*128)
	A(tls, bp+512, q, q+1*128)
	M(tls, bp+128, bp+128, bp+512)
	M(tls, bp+256, p+3*128, q+3*128)
	M(tls, bp+256, bp+256, uintptr(unsafe.Pointer(&D2)))
	M(tls, bp+384, p+2*128, q+2*128)
	A(tls, bp+384, bp+384, bp+384)
	Z(tls, bp+640, bp+128, bp)
	Z(tls, bp+768, bp+384, bp+256)
	A(tls, bp+896, bp+384, bp+256)
	A(tls, bp+1024, bp+128, bp)
	M(tls, p, bp+640, bp+768)
	M(tls, p+1*128, bp+1024, bp+896)
	M(tls, p+2*128, bp+896, bp+768)
	M(tls, p+3*128, bp+640, bp+1024)
}

func cswap(tls *libc.TLS, p uintptr, q uintptr, b u8) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(4)) {
			break
		}
		sel25519(tls, p+uintptr(i)*128, q+uintptr(i)*128, libc.Int32FromUint8(b))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func pack(tls *libc.TLS, r uintptr, p uintptr) {
	bp := tls.Alloc(384)
	defer tls.Free(384)
	var v1 uintptr
	var _ /* tx at bp+0 */ gf
	var _ /* ty at bp+128 */ gf
	var _ /* zi at bp+256 */ gf
	_ = v1
	inv25519(tls, bp+256, p+2*128)
	M(tls, bp, p, bp+256)
	M(tls, bp+128, p+1*128, bp+256)
	pack25519(tls, r, bp+128)
	v1 = r + 31
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) ^ libc.Int32FromUint8(par25519(tls, bp))<<libc.Int32FromInt32(7))
}

func scalarmult(tls *libc.TLS, p uintptr, q uintptr, s uintptr) {
	var b u8
	var i int32
	_, _ = b, i
	set25519(tls, p, uintptr(unsafe.Pointer(&gf0)))
	set25519(tls, p+1*128, uintptr(unsafe.Pointer(&gf1)))
	set25519(tls, p+2*128, uintptr(unsafe.Pointer(&gf1)))
	set25519(tls, p+3*128, uintptr(unsafe.Pointer(&gf0)))
	i = int32(255)
	for {
		if !(i >= 0) {
			break
		}
		b = libc.Uint8FromInt32(libc.Int32FromUint8(**(**u8)(__ccgo_up(s + uintptr(i/int32(8))))) >> (i & int32(7)) & int32(1))
		cswap(tls, p, q, b)
		add(tls, q, p)
		add(tls, p, p)
		cswap(tls, p, q, b)
		goto _1
	_1:
		;
		i = i - 1
	}
}

func scalarbase(tls *libc.TLS, p uintptr, s uintptr) {
	bp := tls.Alloc(512)
	defer tls.Free(512)
	var _ /* q at bp+0 */ [4]gf
	set25519(tls, bp, uintptr(unsafe.Pointer(&X)))
	set25519(tls, bp+1*128, uintptr(unsafe.Pointer(&Y)))
	set25519(tls, bp+2*128, uintptr(unsafe.Pointer(&gf1)))
	M(tls, bp+3*128, uintptr(unsafe.Pointer(&X)), uintptr(unsafe.Pointer(&Y)))
	scalarmult(tls, p, bp, s)
}

func crypto_sign_ed25519_tweet_keypair(tls *libc.TLS, pk uintptr, sk uintptr) (r int32) {
	bp := tls.Alloc(576)
	defer tls.Free(576)
	var i int32
	var v1 uintptr
	var _ /* d at bp+0 */ [64]u8
	var _ /* p at bp+64 */ [4]gf
	_, _ = i, v1
	randombytes(tls, sk, uint64(32))
	crypto_hash_sha512_tweet(tls, bp, sk, uint64(32))
	v1 = bp
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) & libc.Int32FromInt32(248))
	v1 = bp + 31
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) & libc.Int32FromInt32(127))
	v1 = bp + 31
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) | libc.Int32FromInt32(64))
	scalarbase(tls, bp+64, bp)
	pack(tls, pk, bp+64)
	i = 0
	for {
		if !(i < int32(32)) {
			break
		}
		**(**u8)(__ccgo_up(sk + uintptr(int32(32)+i))) = **(**u8)(__ccgo_up(pk + uintptr(i)))
		goto _4
	_4:
		;
		i = i + 1
	}
	return 0
}

var L = [32]u64{
	0:  uint64(0xed),
	1:  uint64(0xd3),
	2:  uint64(0xf5),
	3:  uint64(0x5c),
	4:  uint64(0x1a),
	5:  uint64(0x63),
	6:  uint64(0x12),
	7:  uint64(0x58),
	8:  uint64(0xd6),
	9:  uint64(0x9c),
	10: uint64(0xf7),
	11: uint64(0xa2),
	12: uint64(0xde),
	13: uint64(0xf9),
	14: uint64(0xde),
	15: uint64(0x14),
	31: uint64(0x10),
}

func modL(tls *libc.TLS, r uintptr, x uintptr) {
	var carry, i, j i64
	var v3 uintptr
	_, _, _, _ = carry, i, j, v3
	i = int64(63)
	for {
		if !(i >= int64(32)) {
			break
		}
		carry = 0
		j = i - int64(32)
		for {
			if !(j < i-int64(12)) {
				break
			}
			v3 = x + uintptr(j)*8
			*(*i64)(unsafe.Pointer(v3)) = i64(uint64(*(*i64)(unsafe.Pointer(v3))) + (libc.Uint64FromInt64(carry) - libc.Uint64FromInt64(libc.Int64FromInt32(16)***(**i64)(__ccgo_up(x + uintptr(i)*8)))*L[j-(i-int64(32))]))
			carry = (**(**i64)(__ccgo_up(x + uintptr(j)*8)) + int64(128)) >> int32(8)
			**(**i64)(__ccgo_up(x + uintptr(j)*8)) -= carry << int32(8)
			goto _2
		_2:
			;
			j = j + 1
		}
		**(**i64)(__ccgo_up(x + uintptr(j)*8)) += carry
		**(**i64)(__ccgo_up(x + uintptr(i)*8)) = 0
		goto _1
	_1:
		;
		i = i - 1
	}
	carry = 0
	j = 0
	for {
		if !(j < int64(32)) {
			break
		}
		v3 = x + uintptr(j)*8
		*(*i64)(unsafe.Pointer(v3)) = i64(uint64(*(*i64)(unsafe.Pointer(v3))) + (libc.Uint64FromInt64(carry) - libc.Uint64FromInt64(**(**i64)(__ccgo_up(x + 31*8))>>libc.Int32FromInt32(4))*L[j]))
		carry = **(**i64)(__ccgo_up(x + uintptr(j)*8)) >> int32(8)
		**(**i64)(__ccgo_up(x + uintptr(j)*8)) &= int64(255)
		goto _4
	_4:
		;
		j = j + 1
	}
	j = 0
	for {
		if !(j < int64(32)) {
			break
		}
		v3 = x + uintptr(j)*8
		*(*i64)(unsafe.Pointer(v3)) = i64(uint64(*(*i64)(unsafe.Pointer(v3))) - libc.Uint64FromInt64(carry)*L[j])
		goto _6
	_6:
		;
		j = j + 1
	}
	i = 0
	for {
		if !(i < int64(32)) {
			break
		}
		**(**i64)(__ccgo_up(x + uintptr(i+int64(1))*8)) += **(**i64)(__ccgo_up(x + uintptr(i)*8)) >> int32(8)
		**(**u8)(__ccgo_up(r + uintptr(i))) = libc.Uint8FromInt64(**(**i64)(__ccgo_up(x + uintptr(i)*8)) & int64(255))
		goto _8
	_8:
		;
		i = i + 1
	}
}

func reduce(tls *libc.TLS, r uintptr) {
	bp := tls.Alloc(512)
	defer tls.Free(512)
	var i i64
	var _ /* x at bp+0 */ [64]i64
	_ = i
	i = 0
	for {
		if !(i < int64(64)) {
			break
		}
		(**(**[64]i64)(__ccgo_up(bp)))[i] = libc.Int64FromUint64(uint64(**(**u8)(__ccgo_up(r + uintptr(i)))))
		goto _1
	_1:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int64(64)) {
			break
		}
		**(**u8)(__ccgo_up(r + uintptr(i))) = uint8(0)
		goto _2
	_2:
		;
		i = i + 1
	}
	modL(tls, r, bp)
}

func crypto_sign_ed25519_tweet(tls *libc.TLS, sm uintptr, smlen uintptr, m uintptr, n u64, sk uintptr) (r int32) {
	bp := tls.Alloc(1216)
	defer tls.Free(1216)
	var i, j i64
	var v1 uintptr
	var _ /* d at bp+0 */ [64]u8
	var _ /* h at bp+64 */ [64]u8
	var _ /* p at bp+704 */ [4]gf
	var _ /* r at bp+128 */ [64]u8
	var _ /* x at bp+192 */ [64]i64
	_, _, _ = i, j, v1
	crypto_hash_sha512_tweet(tls, bp, sk, uint64(32))
	v1 = bp
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) & libc.Int32FromInt32(248))
	v1 = bp + 31
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) & libc.Int32FromInt32(127))
	v1 = bp + 31
	*(*u8)(unsafe.Pointer(v1)) = u8(int32(*(*u8)(unsafe.Pointer(v1))) | libc.Int32FromInt32(64))
	**(**u64)(__ccgo_up(smlen)) = n + uint64(64)
	i = 0
	for {
		if !(libc.Uint64FromInt64(i) < n) {
			break
		}
		**(**u8)(__ccgo_up(sm + uintptr(int64(64)+i))) = **(**u8)(__ccgo_up(m + uintptr(i)))
		goto _4
	_4:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int64(32)) {
			break
		}
		**(**u8)(__ccgo_up(sm + uintptr(int64(32)+i))) = (**(**[64]u8)(__ccgo_up(bp)))[int64(32)+i]
		goto _5
	_5:
		;
		i = i + 1
	}
	crypto_hash_sha512_tweet(tls, bp+128, sm+uintptr(32), n+uint64(32))
	reduce(tls, bp+128)
	scalarbase(tls, bp+704, bp+128)
	pack(tls, sm, bp+704)
	i = 0
	for {
		if !(i < int64(32)) {
			break
		}
		**(**u8)(__ccgo_up(sm + uintptr(i+int64(32)))) = **(**u8)(__ccgo_up(sk + uintptr(i+int64(32))))
		goto _6
	_6:
		;
		i = i + 1
	}
	crypto_hash_sha512_tweet(tls, bp+64, sm, n+uint64(64))
	reduce(tls, bp+64)
	i = 0
	for {
		if !(i < int64(64)) {
			break
		}
		(**(**[64]i64)(__ccgo_up(bp + 192)))[i] = 0
		goto _7
	_7:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int64(32)) {
			break
		}
		(**(**[64]i64)(__ccgo_up(bp + 192)))[i] = libc.Int64FromUint64(uint64((**(**[64]u8)(__ccgo_up(bp + 128)))[i]))
		goto _8
	_8:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int64(32)) {
			break
		}
		j = 0
		for {
			if !(j < int64(32)) {
				break
			}
			v1 = bp + 192 + uintptr(i+j)*8
			*(*i64)(unsafe.Pointer(v1)) = i64(uint64(*(*i64)(unsafe.Pointer(v1))) + uint64((**(**[64]u8)(__ccgo_up(bp + 64)))[i])*uint64((**(**[64]u8)(__ccgo_up(bp)))[j]))
			goto _10
		_10:
			;
			j = j + 1
		}
		goto _9
	_9:
		;
		i = i + 1
	}
	modL(tls, sm+uintptr(32), bp+192)
	return 0
}

func unpackneg(tls *libc.TLS, r uintptr, p uintptr) (r1 int32) {
	bp := tls.Alloc(896)
	defer tls.Free(896)
	var _ /* chk at bp+128 */ gf
	var _ /* den at bp+384 */ gf
	var _ /* den2 at bp+512 */ gf
	var _ /* den4 at bp+640 */ gf
	var _ /* den6 at bp+768 */ gf
	var _ /* num at bp+256 */ gf
	var _ /* t at bp+0 */ gf
	set25519(tls, r+2*128, uintptr(unsafe.Pointer(&gf1)))
	unpack25519(tls, r+1*128, p)
	S(tls, bp+256, r+1*128)
	M(tls, bp+384, bp+256, uintptr(unsafe.Pointer(&D)))
	Z(tls, bp+256, bp+256, r+2*128)
	A(tls, bp+384, r+2*128, bp+384)
	S(tls, bp+512, bp+384)
	S(tls, bp+640, bp+512)
	M(tls, bp+768, bp+640, bp+512)
	M(tls, bp, bp+768, bp+256)
	M(tls, bp, bp, bp+384)
	pow2523(tls, bp, bp)
	M(tls, bp, bp, bp+256)
	M(tls, bp, bp, bp+384)
	M(tls, bp, bp, bp+384)
	M(tls, r, bp, bp+384)
	S(tls, bp+128, r)
	M(tls, bp+128, bp+128, bp+384)
	if neq25519(tls, bp+128, bp+256) != 0 {
		M(tls, r, r, uintptr(unsafe.Pointer(&I)))
	}
	S(tls, bp+128, r)
	M(tls, bp+128, bp+128, bp+384)
	if neq25519(tls, bp+128, bp+256) != 0 {
		return -int32(1)
	}
	if libc.Int32FromUint8(par25519(tls, r)) == libc.Int32FromUint8(**(**u8)(__ccgo_up(p + 31)))>>int32(7) {
		Z(tls, r, uintptr(unsafe.Pointer(&gf0)), r)
	}
	M(tls, r+3*128, r, r+1*128)
	return 0
}

func crypto_sign_ed25519_tweet_open(tls *libc.TLS, m uintptr, mlen uintptr, sm uintptr, n u64, pk uintptr) (r int32) {
	bp := tls.Alloc(1120)
	defer tls.Free(1120)
	var i int32
	var _ /* h at bp+32 */ [64]u8
	var _ /* p at bp+96 */ [4]gf
	var _ /* q at bp+608 */ [4]gf
	var _ /* t at bp+0 */ [32]u8
	_ = i
	**(**u64)(__ccgo_up(mlen)) = libc.Uint64FromInt32(-libc.Int32FromInt32(1))
	if n < uint64(64) {
		return -int32(1)
	}
	if unpackneg(tls, bp+608, pk) != 0 {
		return -int32(1)
	}
	i = 0
	for {
		if !(libc.Uint64FromInt32(i) < n) {
			break
		}
		**(**u8)(__ccgo_up(m + uintptr(i))) = **(**u8)(__ccgo_up(sm + uintptr(i)))
		goto _1
	_1:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < int32(32)) {
			break
		}
		**(**u8)(__ccgo_up(m + uintptr(i+int32(32)))) = **(**u8)(__ccgo_up(pk + uintptr(i)))
		goto _2
	_2:
		;
		i = i + 1
	}
	crypto_hash_sha512_tweet(tls, bp+32, m, n)
	reduce(tls, bp+32)
	scalarmult(tls, bp+96, bp+608, bp+32)
	scalarbase(tls, bp+608, sm+uintptr(32))
	add(tls, bp+96, bp+608)
	pack(tls, bp, bp+96)
	n = n - uint64(64)
	if crypto_verify_32_tweet(tls, sm, bp) != 0 {
		i = 0
		for {
			if !(libc.Uint64FromInt32(i) < n) {
				break
			}
			**(**u8)(__ccgo_up(m + uintptr(i))) = uint8(0)
			goto _3
		_3:
			;
			i = i + 1
		}
		return -int32(1)
	}
	i = 0
	for {
		if !(libc.Uint64FromInt32(i) < n) {
			break
		}
		**(**u8)(__ccgo_up(m + uintptr(i))) = **(**u8)(__ccgo_up(sm + uintptr(i+int32(64))))
		goto _4
	_4:
		;
		i = i + 1
	}
	**(**u64)(__ccgo_up(mlen)) = n
	return 0
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}
