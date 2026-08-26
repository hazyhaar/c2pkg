// Code generated for linux/amd64 by 'ccgo --package-name=df_libinj -o spec/dogfood/cycles/20260810f/libinjection_sqli/raw.go -I spec/dogfood/cycles/20260810f/libinjection_sqli spec/dogfood/cycles/20260810f/libinjection_sqli/src.c', DO NOT EDIT.

//go:build linux && amd64

package df_libinj

import (
	"math"
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

const BUFSIZ = 1024
const CHAR_DOUBLE = 34
const CHAR_NULL = 0
const CHAR_SINGLE = 39
const CHAR_TICK = 96
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const FALSE = 0
const FILENAME_MAX = 4096
const FOPEN_MAX = 1000
const LIBINJECTION_SQLI_MAX_TOKENS = 5
const LIBINJECTION_VERSION = "3.9.2"
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const TMP_MAX = 10000
const TRUE = 1
const WNOHANG = 1
const WUNTRACED = 2
const _GNU_SOURCE = 1
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
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
const __inline = "inline"
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const alloca1 = "__builtin_alloca"
const linux = 1
const static_assert = "_Static_assert"
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type size_t = uint64

type locale_t = uintptr

type wchar_t = int32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type ssize_t = int64

type off_t = int64

type va_list = uintptr

type __isoc_va_list = uintptr

type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]int8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type max_align_t = struct {
	F__ll int64
	F__ld float64
}

type ptrdiff_t = int64

type sqli_flags = int32

const FLAG_NONE = 0
const FLAG_QUOTE_NONE = 1
const FLAG_QUOTE_SINGLE = 2
const FLAG_QUOTE_DOUBLE = 4
const FLAG_SQL_ANSI = 8
const FLAG_SQL_MYSQL = 16

type lookup_type = int32

const LOOKUP_WORD = 1
const LOOKUP_TYPE = 2
const LOOKUP_OPERATOR = 3
const LOOKUP_FINGERPRINT = 4

type libinjection_sqli_token = struct {
	Fpos       size_t
	Flen1      size_t
	Fcount     int32
	Ftype1     int8
	Fstr_open  int8
	Fstr_close int8
	Fval       [32]int8
}

type stoken_t = struct {
	Fpos       size_t
	Flen1      size_t
	Fcount     int32
	Ftype1     int8
	Fstr_open  int8
	Fstr_close int8
	Fval       [32]int8
}

type libinjection_sqli_state = struct {
	Fs                  uintptr
	Fslen               size_t
	Flookup             ptr_lookup_fn
	Fuserdata           uintptr
	Fflags              int32
	Fpos                size_t
	Ftokenvec           [8]libinjection_sqli_token
	Fcurrent            uintptr
	Ffingerprint        [8]int8
	Freason             int32
	Fstats_comment_ddw  int32
	Fstats_comment_ddx  int32
	Fstats_comment_c    int32
	Fstats_comment_hash int32
	Fstats_folds        int32
	Fstats_tokens       int32
}

type ptr_lookup_fn = uintptr

type sfilter = struct {
	Fs                  uintptr
	Fslen               size_t
	Flookup             ptr_lookup_fn
	Fuserdata           uintptr
	Fflags              int32
	Fpos                size_t
	Ftokenvec           [8]libinjection_sqli_token
	Fcurrent            uintptr
	Ffingerprint        [8]int8
	Freason             int32
	Fstats_comment_ddw  int32
	Fstats_comment_ddx  int32
	Fstats_comment_c    int32
	Fstats_comment_hash int32
	Fstats_folds        int32
	Fstats_tokens       int32
}

type keyword_t = struct {
	Fword  uintptr
	Ftype1 int8
}

type pt2Function = uintptr

var char_parse_map = [256]pt2Function{}

func init() {
	p := unsafe.Pointer(&char_parse_map)
	*(*uintptr)(unsafe.Add(p, 0)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 8)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 40)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 48)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 64)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 112)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 120)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 128)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 136)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 160)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 168)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 176)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 224)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 232)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 240)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 248)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 256)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 264)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 272)) = __ccgo_fp(parse_string)
	*(*uintptr)(unsafe.Add(p, 280)) = __ccgo_fp(parse_hash)
	*(*uintptr)(unsafe.Add(p, 288)) = __ccgo_fp(parse_money)
	*(*uintptr)(unsafe.Add(p, 296)) = __ccgo_fp(parse_operator1)
	*(*uintptr)(unsafe.Add(p, 304)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 312)) = __ccgo_fp(parse_string)
	*(*uintptr)(unsafe.Add(p, 320)) = __ccgo_fp(parse_char)
	*(*uintptr)(unsafe.Add(p, 328)) = __ccgo_fp(parse_char)
	*(*uintptr)(unsafe.Add(p, 336)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 344)) = __ccgo_fp(parse_operator1)
	*(*uintptr)(unsafe.Add(p, 352)) = __ccgo_fp(parse_char)
	*(*uintptr)(unsafe.Add(p, 360)) = __ccgo_fp(parse_dash)
	*(*uintptr)(unsafe.Add(p, 368)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 376)) = __ccgo_fp(parse_slash)
	*(*uintptr)(unsafe.Add(p, 384)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 392)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 400)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 408)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 416)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 424)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 432)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 440)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 448)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 456)) = __ccgo_fp(parse_number)
	*(*uintptr)(unsafe.Add(p, 464)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 472)) = __ccgo_fp(parse_char)
	*(*uintptr)(unsafe.Add(p, 480)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 488)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 496)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 504)) = __ccgo_fp(parse_other)
	*(*uintptr)(unsafe.Add(p, 512)) = __ccgo_fp(parse_var)
	*(*uintptr)(unsafe.Add(p, 520)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 528)) = __ccgo_fp(parse_bstring)
	*(*uintptr)(unsafe.Add(p, 536)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 544)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 552)) = __ccgo_fp(parse_estring)
	*(*uintptr)(unsafe.Add(p, 560)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 568)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 576)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 584)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 592)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 600)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 608)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 616)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 624)) = __ccgo_fp(parse_nqstring)
	*(*uintptr)(unsafe.Add(p, 632)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 640)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 648)) = __ccgo_fp(parse_qstring)
	*(*uintptr)(unsafe.Add(p, 656)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 664)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 672)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 680)) = __ccgo_fp(parse_ustring)
	*(*uintptr)(unsafe.Add(p, 688)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 696)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 704)) = __ccgo_fp(parse_xstring)
	*(*uintptr)(unsafe.Add(p, 712)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 720)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 728)) = __ccgo_fp(parse_bword)
	*(*uintptr)(unsafe.Add(p, 736)) = __ccgo_fp(parse_backslash)
	*(*uintptr)(unsafe.Add(p, 744)) = __ccgo_fp(parse_other)
	*(*uintptr)(unsafe.Add(p, 752)) = __ccgo_fp(parse_operator1)
	*(*uintptr)(unsafe.Add(p, 760)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 768)) = __ccgo_fp(parse_tick)
	*(*uintptr)(unsafe.Add(p, 776)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 784)) = __ccgo_fp(parse_bstring)
	*(*uintptr)(unsafe.Add(p, 792)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 800)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 808)) = __ccgo_fp(parse_estring)
	*(*uintptr)(unsafe.Add(p, 816)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 824)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 832)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 840)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 848)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 856)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 864)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 872)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 880)) = __ccgo_fp(parse_nqstring)
	*(*uintptr)(unsafe.Add(p, 888)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 896)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 904)) = __ccgo_fp(parse_qstring)
	*(*uintptr)(unsafe.Add(p, 912)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 920)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 928)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 936)) = __ccgo_fp(parse_ustring)
	*(*uintptr)(unsafe.Add(p, 944)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 952)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 960)) = __ccgo_fp(parse_xstring)
	*(*uintptr)(unsafe.Add(p, 968)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 976)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 984)) = __ccgo_fp(parse_char)
	*(*uintptr)(unsafe.Add(p, 992)) = __ccgo_fp(parse_operator2)
	*(*uintptr)(unsafe.Add(p, 1000)) = __ccgo_fp(parse_char)
	*(*uintptr)(unsafe.Add(p, 1008)) = __ccgo_fp(parse_operator1)
	*(*uintptr)(unsafe.Add(p, 1016)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 1024)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1032)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1040)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1048)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1056)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1064)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1072)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1080)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1088)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1096)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1104)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1112)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1120)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1128)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1136)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1144)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1152)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1160)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1168)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1176)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1184)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1192)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1200)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1208)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1216)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1224)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1232)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1240)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1248)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1256)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1264)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1272)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1280)) = __ccgo_fp(parse_white)
	*(*uintptr)(unsafe.Add(p, 1288)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1296)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1304)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1312)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1320)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1328)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1336)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1344)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1352)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1360)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1368)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1376)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1384)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1392)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1400)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1408)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1416)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1424)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1432)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1440)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1448)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1456)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1464)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1472)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1480)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1488)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1496)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1504)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1512)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1520)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1528)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1536)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1544)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1552)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1560)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1568)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1576)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1584)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1592)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1600)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1608)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1616)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1624)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1632)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1640)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1648)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1656)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1664)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1672)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1680)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1688)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1696)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1704)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1712)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1720)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1728)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1736)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1744)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1752)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1760)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1768)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1776)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1784)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1792)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1800)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1808)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1816)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1824)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1832)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1840)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1848)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1856)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1864)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1872)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1880)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1888)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1896)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1904)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1912)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1920)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1928)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1936)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1944)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1952)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1960)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1968)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1976)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1984)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 1992)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 2000)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 2008)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 2016)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 2024)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 2032)) = __ccgo_fp(parse_word)
	*(*uintptr)(unsafe.Add(p, 2040)) = __ccgo_fp(parse_word)
}

var sql_keywords = [9352]keyword_t{
	0: {
		Fword:  __ccgo_ts,
		Ftype1: int8('o'),
	},
	1: {
		Fword:  __ccgo_ts + 3,
		Ftype1: int8('o'),
	},
	2: {
		Fword:  __ccgo_ts + 6,
		Ftype1: int8('o'),
	},
	3: {
		Fword:  __ccgo_ts + 9,
		Ftype1: int8('o'),
	},
	4: {
		Fword:  __ccgo_ts + 12,
		Ftype1: int8('o'),
	},
	5: {
		Fword:  __ccgo_ts + 15,
		Ftype1: int8('&'),
	},
	6: {
		Fword:  __ccgo_ts + 18,
		Ftype1: int8('o'),
	},
	7: {
		Fword:  __ccgo_ts + 21,
		Ftype1: int8('o'),
	},
	8: {
		Fword:  __ccgo_ts + 24,
		Ftype1: int8('o'),
	},
	9: {
		Fword:  __ccgo_ts + 27,
		Ftype1: int8('o'),
	},
	10: {
		Fword:  __ccgo_ts + 30,
		Ftype1: int8('o'),
	},
	11: {
		Fword:  __ccgo_ts + 33,
		Ftype1: int8('F'),
	},
	12: {
		Fword:  __ccgo_ts + 40,
		Ftype1: int8('F'),
	},
	13: {
		Fword:  __ccgo_ts + 47,
		Ftype1: int8('F'),
	},
	14: {
		Fword:  __ccgo_ts + 54,
		Ftype1: int8('F'),
	},
	15: {
		Fword:  __ccgo_ts + 61,
		Ftype1: int8('F'),
	},
	16: {
		Fword:  __ccgo_ts + 68,
		Ftype1: int8('F'),
	},
	17: {
		Fword:  __ccgo_ts + 75,
		Ftype1: int8('F'),
	},
	18: {
		Fword:  __ccgo_ts + 82,
		Ftype1: int8('F'),
	},
	19: {
		Fword:  __ccgo_ts + 89,
		Ftype1: int8('F'),
	},
	20: {
		Fword:  __ccgo_ts + 96,
		Ftype1: int8('F'),
	},
	21: {
		Fword:  __ccgo_ts + 103,
		Ftype1: int8('F'),
	},
	22: {
		Fword:  __ccgo_ts + 110,
		Ftype1: int8('F'),
	},
	23: {
		Fword:  __ccgo_ts + 117,
		Ftype1: int8('F'),
	},
	24: {
		Fword:  __ccgo_ts + 124,
		Ftype1: int8('F'),
	},
	25: {
		Fword:  __ccgo_ts + 131,
		Ftype1: int8('F'),
	},
	26: {
		Fword:  __ccgo_ts + 138,
		Ftype1: int8('F'),
	},
	27: {
		Fword:  __ccgo_ts + 145,
		Ftype1: int8('F'),
	},
	28: {
		Fword:  __ccgo_ts + 152,
		Ftype1: int8('F'),
	},
	29: {
		Fword:  __ccgo_ts + 159,
		Ftype1: int8('F'),
	},
	30: {
		Fword:  __ccgo_ts + 166,
		Ftype1: int8('F'),
	},
	31: {
		Fword:  __ccgo_ts + 173,
		Ftype1: int8('F'),
	},
	32: {
		Fword:  __ccgo_ts + 180,
		Ftype1: int8('F'),
	},
	33: {
		Fword:  __ccgo_ts + 187,
		Ftype1: int8('F'),
	},
	34: {
		Fword:  __ccgo_ts + 194,
		Ftype1: int8('F'),
	},
	35: {
		Fword:  __ccgo_ts + 201,
		Ftype1: int8('F'),
	},
	36: {
		Fword:  __ccgo_ts + 208,
		Ftype1: int8('F'),
	},
	37: {
		Fword:  __ccgo_ts + 215,
		Ftype1: int8('F'),
	},
	38: {
		Fword:  __ccgo_ts + 222,
		Ftype1: int8('F'),
	},
	39: {
		Fword:  __ccgo_ts + 229,
		Ftype1: int8('F'),
	},
	40: {
		Fword:  __ccgo_ts + 236,
		Ftype1: int8('F'),
	},
	41: {
		Fword:  __ccgo_ts + 243,
		Ftype1: int8('F'),
	},
	42: {
		Fword:  __ccgo_ts + 250,
		Ftype1: int8('F'),
	},
	43: {
		Fword:  __ccgo_ts + 257,
		Ftype1: int8('F'),
	},
	44: {
		Fword:  __ccgo_ts + 264,
		Ftype1: int8('F'),
	},
	45: {
		Fword:  __ccgo_ts + 271,
		Ftype1: int8('F'),
	},
	46: {
		Fword:  __ccgo_ts + 278,
		Ftype1: int8('F'),
	},
	47: {
		Fword:  __ccgo_ts + 285,
		Ftype1: int8('F'),
	},
	48: {
		Fword:  __ccgo_ts + 292,
		Ftype1: int8('F'),
	},
	49: {
		Fword:  __ccgo_ts + 299,
		Ftype1: int8('F'),
	},
	50: {
		Fword:  __ccgo_ts + 306,
		Ftype1: int8('F'),
	},
	51: {
		Fword:  __ccgo_ts + 313,
		Ftype1: int8('F'),
	},
	52: {
		Fword:  __ccgo_ts + 320,
		Ftype1: int8('F'),
	},
	53: {
		Fword:  __ccgo_ts + 327,
		Ftype1: int8('F'),
	},
	54: {
		Fword:  __ccgo_ts + 334,
		Ftype1: int8('F'),
	},
	55: {
		Fword:  __ccgo_ts + 341,
		Ftype1: int8('F'),
	},
	56: {
		Fword:  __ccgo_ts + 348,
		Ftype1: int8('F'),
	},
	57: {
		Fword:  __ccgo_ts + 355,
		Ftype1: int8('F'),
	},
	58: {
		Fword:  __ccgo_ts + 362,
		Ftype1: int8('F'),
	},
	59: {
		Fword:  __ccgo_ts + 369,
		Ftype1: int8('F'),
	},
	60: {
		Fword:  __ccgo_ts + 376,
		Ftype1: int8('F'),
	},
	61: {
		Fword:  __ccgo_ts + 383,
		Ftype1: int8('F'),
	},
	62: {
		Fword:  __ccgo_ts + 390,
		Ftype1: int8('F'),
	},
	63: {
		Fword:  __ccgo_ts + 397,
		Ftype1: int8('F'),
	},
	64: {
		Fword:  __ccgo_ts + 404,
		Ftype1: int8('F'),
	},
	65: {
		Fword:  __ccgo_ts + 411,
		Ftype1: int8('F'),
	},
	66: {
		Fword:  __ccgo_ts + 418,
		Ftype1: int8('F'),
	},
	67: {
		Fword:  __ccgo_ts + 425,
		Ftype1: int8('F'),
	},
	68: {
		Fword:  __ccgo_ts + 432,
		Ftype1: int8('F'),
	},
	69: {
		Fword:  __ccgo_ts + 439,
		Ftype1: int8('F'),
	},
	70: {
		Fword:  __ccgo_ts + 446,
		Ftype1: int8('F'),
	},
	71: {
		Fword:  __ccgo_ts + 453,
		Ftype1: int8('F'),
	},
	72: {
		Fword:  __ccgo_ts + 460,
		Ftype1: int8('F'),
	},
	73: {
		Fword:  __ccgo_ts + 467,
		Ftype1: int8('F'),
	},
	74: {
		Fword:  __ccgo_ts + 474,
		Ftype1: int8('F'),
	},
	75: {
		Fword:  __ccgo_ts + 481,
		Ftype1: int8('F'),
	},
	76: {
		Fword:  __ccgo_ts + 488,
		Ftype1: int8('F'),
	},
	77: {
		Fword:  __ccgo_ts + 495,
		Ftype1: int8('F'),
	},
	78: {
		Fword:  __ccgo_ts + 502,
		Ftype1: int8('F'),
	},
	79: {
		Fword:  __ccgo_ts + 509,
		Ftype1: int8('F'),
	},
	80: {
		Fword:  __ccgo_ts + 516,
		Ftype1: int8('F'),
	},
	81: {
		Fword:  __ccgo_ts + 523,
		Ftype1: int8('F'),
	},
	82: {
		Fword:  __ccgo_ts + 530,
		Ftype1: int8('F'),
	},
	83: {
		Fword:  __ccgo_ts + 537,
		Ftype1: int8('F'),
	},
	84: {
		Fword:  __ccgo_ts + 544,
		Ftype1: int8('F'),
	},
	85: {
		Fword:  __ccgo_ts + 551,
		Ftype1: int8('F'),
	},
	86: {
		Fword:  __ccgo_ts + 558,
		Ftype1: int8('F'),
	},
	87: {
		Fword:  __ccgo_ts + 565,
		Ftype1: int8('F'),
	},
	88: {
		Fword:  __ccgo_ts + 572,
		Ftype1: int8('F'),
	},
	89: {
		Fword:  __ccgo_ts + 579,
		Ftype1: int8('F'),
	},
	90: {
		Fword:  __ccgo_ts + 586,
		Ftype1: int8('F'),
	},
	91: {
		Fword:  __ccgo_ts + 593,
		Ftype1: int8('F'),
	},
	92: {
		Fword:  __ccgo_ts + 600,
		Ftype1: int8('F'),
	},
	93: {
		Fword:  __ccgo_ts + 607,
		Ftype1: int8('F'),
	},
	94: {
		Fword:  __ccgo_ts + 614,
		Ftype1: int8('F'),
	},
	95: {
		Fword:  __ccgo_ts + 621,
		Ftype1: int8('F'),
	},
	96: {
		Fword:  __ccgo_ts + 628,
		Ftype1: int8('F'),
	},
	97: {
		Fword:  __ccgo_ts + 635,
		Ftype1: int8('F'),
	},
	98: {
		Fword:  __ccgo_ts + 642,
		Ftype1: int8('F'),
	},
	99: {
		Fword:  __ccgo_ts + 649,
		Ftype1: int8('F'),
	},
	100: {
		Fword:  __ccgo_ts + 656,
		Ftype1: int8('F'),
	},
	101: {
		Fword:  __ccgo_ts + 663,
		Ftype1: int8('F'),
	},
	102: {
		Fword:  __ccgo_ts + 670,
		Ftype1: int8('F'),
	},
	103: {
		Fword:  __ccgo_ts + 677,
		Ftype1: int8('F'),
	},
	104: {
		Fword:  __ccgo_ts + 684,
		Ftype1: int8('F'),
	},
	105: {
		Fword:  __ccgo_ts + 691,
		Ftype1: int8('F'),
	},
	106: {
		Fword:  __ccgo_ts + 698,
		Ftype1: int8('F'),
	},
	107: {
		Fword:  __ccgo_ts + 705,
		Ftype1: int8('F'),
	},
	108: {
		Fword:  __ccgo_ts + 712,
		Ftype1: int8('F'),
	},
	109: {
		Fword:  __ccgo_ts + 719,
		Ftype1: int8('F'),
	},
	110: {
		Fword:  __ccgo_ts + 726,
		Ftype1: int8('F'),
	},
	111: {
		Fword:  __ccgo_ts + 733,
		Ftype1: int8('F'),
	},
	112: {
		Fword:  __ccgo_ts + 740,
		Ftype1: int8('F'),
	},
	113: {
		Fword:  __ccgo_ts + 747,
		Ftype1: int8('F'),
	},
	114: {
		Fword:  __ccgo_ts + 754,
		Ftype1: int8('F'),
	},
	115: {
		Fword:  __ccgo_ts + 761,
		Ftype1: int8('F'),
	},
	116: {
		Fword:  __ccgo_ts + 768,
		Ftype1: int8('F'),
	},
	117: {
		Fword:  __ccgo_ts + 775,
		Ftype1: int8('F'),
	},
	118: {
		Fword:  __ccgo_ts + 782,
		Ftype1: int8('F'),
	},
	119: {
		Fword:  __ccgo_ts + 789,
		Ftype1: int8('F'),
	},
	120: {
		Fword:  __ccgo_ts + 796,
		Ftype1: int8('F'),
	},
	121: {
		Fword:  __ccgo_ts + 803,
		Ftype1: int8('F'),
	},
	122: {
		Fword:  __ccgo_ts + 810,
		Ftype1: int8('F'),
	},
	123: {
		Fword:  __ccgo_ts + 817,
		Ftype1: int8('F'),
	},
	124: {
		Fword:  __ccgo_ts + 824,
		Ftype1: int8('F'),
	},
	125: {
		Fword:  __ccgo_ts + 831,
		Ftype1: int8('F'),
	},
	126: {
		Fword:  __ccgo_ts + 838,
		Ftype1: int8('F'),
	},
	127: {
		Fword:  __ccgo_ts + 845,
		Ftype1: int8('F'),
	},
	128: {
		Fword:  __ccgo_ts + 852,
		Ftype1: int8('F'),
	},
	129: {
		Fword:  __ccgo_ts + 859,
		Ftype1: int8('F'),
	},
	130: {
		Fword:  __ccgo_ts + 866,
		Ftype1: int8('F'),
	},
	131: {
		Fword:  __ccgo_ts + 873,
		Ftype1: int8('F'),
	},
	132: {
		Fword:  __ccgo_ts + 880,
		Ftype1: int8('F'),
	},
	133: {
		Fword:  __ccgo_ts + 887,
		Ftype1: int8('F'),
	},
	134: {
		Fword:  __ccgo_ts + 894,
		Ftype1: int8('F'),
	},
	135: {
		Fword:  __ccgo_ts + 901,
		Ftype1: int8('F'),
	},
	136: {
		Fword:  __ccgo_ts + 908,
		Ftype1: int8('F'),
	},
	137: {
		Fword:  __ccgo_ts + 915,
		Ftype1: int8('F'),
	},
	138: {
		Fword:  __ccgo_ts + 922,
		Ftype1: int8('F'),
	},
	139: {
		Fword:  __ccgo_ts + 929,
		Ftype1: int8('F'),
	},
	140: {
		Fword:  __ccgo_ts + 936,
		Ftype1: int8('F'),
	},
	141: {
		Fword:  __ccgo_ts + 943,
		Ftype1: int8('F'),
	},
	142: {
		Fword:  __ccgo_ts + 950,
		Ftype1: int8('F'),
	},
	143: {
		Fword:  __ccgo_ts + 957,
		Ftype1: int8('F'),
	},
	144: {
		Fword:  __ccgo_ts + 964,
		Ftype1: int8('F'),
	},
	145: {
		Fword:  __ccgo_ts + 971,
		Ftype1: int8('F'),
	},
	146: {
		Fword:  __ccgo_ts + 978,
		Ftype1: int8('F'),
	},
	147: {
		Fword:  __ccgo_ts + 985,
		Ftype1: int8('F'),
	},
	148: {
		Fword:  __ccgo_ts + 992,
		Ftype1: int8('F'),
	},
	149: {
		Fword:  __ccgo_ts + 999,
		Ftype1: int8('F'),
	},
	150: {
		Fword:  __ccgo_ts + 1006,
		Ftype1: int8('F'),
	},
	151: {
		Fword:  __ccgo_ts + 1013,
		Ftype1: int8('F'),
	},
	152: {
		Fword:  __ccgo_ts + 1020,
		Ftype1: int8('F'),
	},
	153: {
		Fword:  __ccgo_ts + 1027,
		Ftype1: int8('F'),
	},
	154: {
		Fword:  __ccgo_ts + 1034,
		Ftype1: int8('F'),
	},
	155: {
		Fword:  __ccgo_ts + 1041,
		Ftype1: int8('F'),
	},
	156: {
		Fword:  __ccgo_ts + 1048,
		Ftype1: int8('F'),
	},
	157: {
		Fword:  __ccgo_ts + 1055,
		Ftype1: int8('F'),
	},
	158: {
		Fword:  __ccgo_ts + 1062,
		Ftype1: int8('F'),
	},
	159: {
		Fword:  __ccgo_ts + 1069,
		Ftype1: int8('F'),
	},
	160: {
		Fword:  __ccgo_ts + 1076,
		Ftype1: int8('F'),
	},
	161: {
		Fword:  __ccgo_ts + 1083,
		Ftype1: int8('F'),
	},
	162: {
		Fword:  __ccgo_ts + 1090,
		Ftype1: int8('F'),
	},
	163: {
		Fword:  __ccgo_ts + 1097,
		Ftype1: int8('F'),
	},
	164: {
		Fword:  __ccgo_ts + 1104,
		Ftype1: int8('F'),
	},
	165: {
		Fword:  __ccgo_ts + 1111,
		Ftype1: int8('F'),
	},
	166: {
		Fword:  __ccgo_ts + 1118,
		Ftype1: int8('F'),
	},
	167: {
		Fword:  __ccgo_ts + 1125,
		Ftype1: int8('F'),
	},
	168: {
		Fword:  __ccgo_ts + 1132,
		Ftype1: int8('F'),
	},
	169: {
		Fword:  __ccgo_ts + 1139,
		Ftype1: int8('F'),
	},
	170: {
		Fword:  __ccgo_ts + 1146,
		Ftype1: int8('F'),
	},
	171: {
		Fword:  __ccgo_ts + 1153,
		Ftype1: int8('F'),
	},
	172: {
		Fword:  __ccgo_ts + 1160,
		Ftype1: int8('F'),
	},
	173: {
		Fword:  __ccgo_ts + 1167,
		Ftype1: int8('F'),
	},
	174: {
		Fword:  __ccgo_ts + 1174,
		Ftype1: int8('F'),
	},
	175: {
		Fword:  __ccgo_ts + 1181,
		Ftype1: int8('F'),
	},
	176: {
		Fword:  __ccgo_ts + 1188,
		Ftype1: int8('F'),
	},
	177: {
		Fword:  __ccgo_ts + 1195,
		Ftype1: int8('F'),
	},
	178: {
		Fword:  __ccgo_ts + 1202,
		Ftype1: int8('F'),
	},
	179: {
		Fword:  __ccgo_ts + 1209,
		Ftype1: int8('F'),
	},
	180: {
		Fword:  __ccgo_ts + 1216,
		Ftype1: int8('F'),
	},
	181: {
		Fword:  __ccgo_ts + 1223,
		Ftype1: int8('F'),
	},
	182: {
		Fword:  __ccgo_ts + 1230,
		Ftype1: int8('F'),
	},
	183: {
		Fword:  __ccgo_ts + 1237,
		Ftype1: int8('F'),
	},
	184: {
		Fword:  __ccgo_ts + 1244,
		Ftype1: int8('F'),
	},
	185: {
		Fword:  __ccgo_ts + 1251,
		Ftype1: int8('F'),
	},
	186: {
		Fword:  __ccgo_ts + 1258,
		Ftype1: int8('F'),
	},
	187: {
		Fword:  __ccgo_ts + 1265,
		Ftype1: int8('F'),
	},
	188: {
		Fword:  __ccgo_ts + 1272,
		Ftype1: int8('F'),
	},
	189: {
		Fword:  __ccgo_ts + 1279,
		Ftype1: int8('F'),
	},
	190: {
		Fword:  __ccgo_ts + 1286,
		Ftype1: int8('F'),
	},
	191: {
		Fword:  __ccgo_ts + 1293,
		Ftype1: int8('F'),
	},
	192: {
		Fword:  __ccgo_ts + 1300,
		Ftype1: int8('F'),
	},
	193: {
		Fword:  __ccgo_ts + 1307,
		Ftype1: int8('F'),
	},
	194: {
		Fword:  __ccgo_ts + 1314,
		Ftype1: int8('F'),
	},
	195: {
		Fword:  __ccgo_ts + 1321,
		Ftype1: int8('F'),
	},
	196: {
		Fword:  __ccgo_ts + 1328,
		Ftype1: int8('F'),
	},
	197: {
		Fword:  __ccgo_ts + 1335,
		Ftype1: int8('F'),
	},
	198: {
		Fword:  __ccgo_ts + 1340,
		Ftype1: int8('F'),
	},
	199: {
		Fword:  __ccgo_ts + 1347,
		Ftype1: int8('F'),
	},
	200: {
		Fword:  __ccgo_ts + 1354,
		Ftype1: int8('F'),
	},
	201: {
		Fword:  __ccgo_ts + 1361,
		Ftype1: int8('F'),
	},
	202: {
		Fword:  __ccgo_ts + 1368,
		Ftype1: int8('F'),
	},
	203: {
		Fword:  __ccgo_ts + 1375,
		Ftype1: int8('F'),
	},
	204: {
		Fword:  __ccgo_ts + 1382,
		Ftype1: int8('F'),
	},
	205: {
		Fword:  __ccgo_ts + 1389,
		Ftype1: int8('F'),
	},
	206: {
		Fword:  __ccgo_ts + 1396,
		Ftype1: int8('F'),
	},
	207: {
		Fword:  __ccgo_ts + 1403,
		Ftype1: int8('F'),
	},
	208: {
		Fword:  __ccgo_ts + 1410,
		Ftype1: int8('F'),
	},
	209: {
		Fword:  __ccgo_ts + 1416,
		Ftype1: int8('F'),
	},
	210: {
		Fword:  __ccgo_ts + 1423,
		Ftype1: int8('F'),
	},
	211: {
		Fword:  __ccgo_ts + 1430,
		Ftype1: int8('F'),
	},
	212: {
		Fword:  __ccgo_ts + 1437,
		Ftype1: int8('F'),
	},
	213: {
		Fword:  __ccgo_ts + 1444,
		Ftype1: int8('F'),
	},
	214: {
		Fword:  __ccgo_ts + 1451,
		Ftype1: int8('F'),
	},
	215: {
		Fword:  __ccgo_ts + 1458,
		Ftype1: int8('F'),
	},
	216: {
		Fword:  __ccgo_ts + 1465,
		Ftype1: int8('F'),
	},
	217: {
		Fword:  __ccgo_ts + 1472,
		Ftype1: int8('F'),
	},
	218: {
		Fword:  __ccgo_ts + 1479,
		Ftype1: int8('F'),
	},
	219: {
		Fword:  __ccgo_ts + 1485,
		Ftype1: int8('F'),
	},
	220: {
		Fword:  __ccgo_ts + 1492,
		Ftype1: int8('F'),
	},
	221: {
		Fword:  __ccgo_ts + 1499,
		Ftype1: int8('F'),
	},
	222: {
		Fword:  __ccgo_ts + 1506,
		Ftype1: int8('F'),
	},
	223: {
		Fword:  __ccgo_ts + 1513,
		Ftype1: int8('F'),
	},
	224: {
		Fword:  __ccgo_ts + 1520,
		Ftype1: int8('F'),
	},
	225: {
		Fword:  __ccgo_ts + 1527,
		Ftype1: int8('F'),
	},
	226: {
		Fword:  __ccgo_ts + 1534,
		Ftype1: int8('F'),
	},
	227: {
		Fword:  __ccgo_ts + 1541,
		Ftype1: int8('F'),
	},
	228: {
		Fword:  __ccgo_ts + 1548,
		Ftype1: int8('F'),
	},
	229: {
		Fword:  __ccgo_ts + 1555,
		Ftype1: int8('F'),
	},
	230: {
		Fword:  __ccgo_ts + 1562,
		Ftype1: int8('F'),
	},
	231: {
		Fword:  __ccgo_ts + 1569,
		Ftype1: int8('F'),
	},
	232: {
		Fword:  __ccgo_ts + 1576,
		Ftype1: int8('F'),
	},
	233: {
		Fword:  __ccgo_ts + 1583,
		Ftype1: int8('F'),
	},
	234: {
		Fword:  __ccgo_ts + 1589,
		Ftype1: int8('F'),
	},
	235: {
		Fword:  __ccgo_ts + 1596,
		Ftype1: int8('F'),
	},
	236: {
		Fword:  __ccgo_ts + 1603,
		Ftype1: int8('F'),
	},
	237: {
		Fword:  __ccgo_ts + 1610,
		Ftype1: int8('F'),
	},
	238: {
		Fword:  __ccgo_ts + 1617,
		Ftype1: int8('F'),
	},
	239: {
		Fword:  __ccgo_ts + 1624,
		Ftype1: int8('F'),
	},
	240: {
		Fword:  __ccgo_ts + 1631,
		Ftype1: int8('F'),
	},
	241: {
		Fword:  __ccgo_ts + 1638,
		Ftype1: int8('F'),
	},
	242: {
		Fword:  __ccgo_ts + 1645,
		Ftype1: int8('F'),
	},
	243: {
		Fword:  __ccgo_ts + 1652,
		Ftype1: int8('F'),
	},
	244: {
		Fword:  __ccgo_ts + 1659,
		Ftype1: int8('F'),
	},
	245: {
		Fword:  __ccgo_ts + 1665,
		Ftype1: int8('F'),
	},
	246: {
		Fword:  __ccgo_ts + 1672,
		Ftype1: int8('F'),
	},
	247: {
		Fword:  __ccgo_ts + 1679,
		Ftype1: int8('F'),
	},
	248: {
		Fword:  __ccgo_ts + 1686,
		Ftype1: int8('F'),
	},
	249: {
		Fword:  __ccgo_ts + 1693,
		Ftype1: int8('F'),
	},
	250: {
		Fword:  __ccgo_ts + 1700,
		Ftype1: int8('F'),
	},
	251: {
		Fword:  __ccgo_ts + 1707,
		Ftype1: int8('F'),
	},
	252: {
		Fword:  __ccgo_ts + 1714,
		Ftype1: int8('F'),
	},
	253: {
		Fword:  __ccgo_ts + 1721,
		Ftype1: int8('F'),
	},
	254: {
		Fword:  __ccgo_ts + 1728,
		Ftype1: int8('F'),
	},
	255: {
		Fword:  __ccgo_ts + 1735,
		Ftype1: int8('F'),
	},
	256: {
		Fword:  __ccgo_ts + 1742,
		Ftype1: int8('F'),
	},
	257: {
		Fword:  __ccgo_ts + 1749,
		Ftype1: int8('F'),
	},
	258: {
		Fword:  __ccgo_ts + 1755,
		Ftype1: int8('F'),
	},
	259: {
		Fword:  __ccgo_ts + 1762,
		Ftype1: int8('F'),
	},
	260: {
		Fword:  __ccgo_ts + 1769,
		Ftype1: int8('F'),
	},
	261: {
		Fword:  __ccgo_ts + 1776,
		Ftype1: int8('F'),
	},
	262: {
		Fword:  __ccgo_ts + 1783,
		Ftype1: int8('F'),
	},
	263: {
		Fword:  __ccgo_ts + 1789,
		Ftype1: int8('F'),
	},
	264: {
		Fword:  __ccgo_ts + 1796,
		Ftype1: int8('F'),
	},
	265: {
		Fword:  __ccgo_ts + 1803,
		Ftype1: int8('F'),
	},
	266: {
		Fword:  __ccgo_ts + 1810,
		Ftype1: int8('F'),
	},
	267: {
		Fword:  __ccgo_ts + 1817,
		Ftype1: int8('F'),
	},
	268: {
		Fword:  __ccgo_ts + 1824,
		Ftype1: int8('F'),
	},
	269: {
		Fword:  __ccgo_ts + 1830,
		Ftype1: int8('F'),
	},
	270: {
		Fword:  __ccgo_ts + 1837,
		Ftype1: int8('F'),
	},
	271: {
		Fword:  __ccgo_ts + 1844,
		Ftype1: int8('F'),
	},
	272: {
		Fword:  __ccgo_ts + 1851,
		Ftype1: int8('F'),
	},
	273: {
		Fword:  __ccgo_ts + 1858,
		Ftype1: int8('F'),
	},
	274: {
		Fword:  __ccgo_ts + 1865,
		Ftype1: int8('F'),
	},
	275: {
		Fword:  __ccgo_ts + 1872,
		Ftype1: int8('F'),
	},
	276: {
		Fword:  __ccgo_ts + 1879,
		Ftype1: int8('F'),
	},
	277: {
		Fword:  __ccgo_ts + 1886,
		Ftype1: int8('F'),
	},
	278: {
		Fword:  __ccgo_ts + 1893,
		Ftype1: int8('F'),
	},
	279: {
		Fword:  __ccgo_ts + 1900,
		Ftype1: int8('F'),
	},
	280: {
		Fword:  __ccgo_ts + 1907,
		Ftype1: int8('F'),
	},
	281: {
		Fword:  __ccgo_ts + 1914,
		Ftype1: int8('F'),
	},
	282: {
		Fword:  __ccgo_ts + 1921,
		Ftype1: int8('F'),
	},
	283: {
		Fword:  __ccgo_ts + 1928,
		Ftype1: int8('F'),
	},
	284: {
		Fword:  __ccgo_ts + 1935,
		Ftype1: int8('F'),
	},
	285: {
		Fword:  __ccgo_ts + 1942,
		Ftype1: int8('F'),
	},
	286: {
		Fword:  __ccgo_ts + 1949,
		Ftype1: int8('F'),
	},
	287: {
		Fword:  __ccgo_ts + 1956,
		Ftype1: int8('F'),
	},
	288: {
		Fword:  __ccgo_ts + 1963,
		Ftype1: int8('F'),
	},
	289: {
		Fword:  __ccgo_ts + 1970,
		Ftype1: int8('F'),
	},
	290: {
		Fword:  __ccgo_ts + 1977,
		Ftype1: int8('F'),
	},
	291: {
		Fword:  __ccgo_ts + 1984,
		Ftype1: int8('F'),
	},
	292: {
		Fword:  __ccgo_ts + 1991,
		Ftype1: int8('F'),
	},
	293: {
		Fword:  __ccgo_ts + 1997,
		Ftype1: int8('F'),
	},
	294: {
		Fword:  __ccgo_ts + 2004,
		Ftype1: int8('F'),
	},
	295: {
		Fword:  __ccgo_ts + 2011,
		Ftype1: int8('F'),
	},
	296: {
		Fword:  __ccgo_ts + 2018,
		Ftype1: int8('F'),
	},
	297: {
		Fword:  __ccgo_ts + 2025,
		Ftype1: int8('F'),
	},
	298: {
		Fword:  __ccgo_ts + 2032,
		Ftype1: int8('F'),
	},
	299: {
		Fword:  __ccgo_ts + 2039,
		Ftype1: int8('F'),
	},
	300: {
		Fword:  __ccgo_ts + 2046,
		Ftype1: int8('F'),
	},
	301: {
		Fword:  __ccgo_ts + 2053,
		Ftype1: int8('F'),
	},
	302: {
		Fword:  __ccgo_ts + 2060,
		Ftype1: int8('F'),
	},
	303: {
		Fword:  __ccgo_ts + 2067,
		Ftype1: int8('F'),
	},
	304: {
		Fword:  __ccgo_ts + 2074,
		Ftype1: int8('F'),
	},
	305: {
		Fword:  __ccgo_ts + 2081,
		Ftype1: int8('F'),
	},
	306: {
		Fword:  __ccgo_ts + 2088,
		Ftype1: int8('F'),
	},
	307: {
		Fword:  __ccgo_ts + 2095,
		Ftype1: int8('F'),
	},
	308: {
		Fword:  __ccgo_ts + 2102,
		Ftype1: int8('F'),
	},
	309: {
		Fword:  __ccgo_ts + 2109,
		Ftype1: int8('F'),
	},
	310: {
		Fword:  __ccgo_ts + 2116,
		Ftype1: int8('F'),
	},
	311: {
		Fword:  __ccgo_ts + 2123,
		Ftype1: int8('F'),
	},
	312: {
		Fword:  __ccgo_ts + 2130,
		Ftype1: int8('F'),
	},
	313: {
		Fword:  __ccgo_ts + 2137,
		Ftype1: int8('F'),
	},
	314: {
		Fword:  __ccgo_ts + 2144,
		Ftype1: int8('F'),
	},
	315: {
		Fword:  __ccgo_ts + 2150,
		Ftype1: int8('F'),
	},
	316: {
		Fword:  __ccgo_ts + 2157,
		Ftype1: int8('F'),
	},
	317: {
		Fword:  __ccgo_ts + 2164,
		Ftype1: int8('F'),
	},
	318: {
		Fword:  __ccgo_ts + 2171,
		Ftype1: int8('F'),
	},
	319: {
		Fword:  __ccgo_ts + 2178,
		Ftype1: int8('F'),
	},
	320: {
		Fword:  __ccgo_ts + 2185,
		Ftype1: int8('F'),
	},
	321: {
		Fword:  __ccgo_ts + 2192,
		Ftype1: int8('F'),
	},
	322: {
		Fword:  __ccgo_ts + 2199,
		Ftype1: int8('F'),
	},
	323: {
		Fword:  __ccgo_ts + 2206,
		Ftype1: int8('F'),
	},
	324: {
		Fword:  __ccgo_ts + 2213,
		Ftype1: int8('F'),
	},
	325: {
		Fword:  __ccgo_ts + 2219,
		Ftype1: int8('F'),
	},
	326: {
		Fword:  __ccgo_ts + 2226,
		Ftype1: int8('F'),
	},
	327: {
		Fword:  __ccgo_ts + 2233,
		Ftype1: int8('F'),
	},
	328: {
		Fword:  __ccgo_ts + 2240,
		Ftype1: int8('F'),
	},
	329: {
		Fword:  __ccgo_ts + 2247,
		Ftype1: int8('F'),
	},
	330: {
		Fword:  __ccgo_ts + 2254,
		Ftype1: int8('F'),
	},
	331: {
		Fword:  __ccgo_ts + 2261,
		Ftype1: int8('F'),
	},
	332: {
		Fword:  __ccgo_ts + 2268,
		Ftype1: int8('F'),
	},
	333: {
		Fword:  __ccgo_ts + 2275,
		Ftype1: int8('F'),
	},
	334: {
		Fword:  __ccgo_ts + 2282,
		Ftype1: int8('F'),
	},
	335: {
		Fword:  __ccgo_ts + 2289,
		Ftype1: int8('F'),
	},
	336: {
		Fword:  __ccgo_ts + 2296,
		Ftype1: int8('F'),
	},
	337: {
		Fword:  __ccgo_ts + 2303,
		Ftype1: int8('F'),
	},
	338: {
		Fword:  __ccgo_ts + 2310,
		Ftype1: int8('F'),
	},
	339: {
		Fword:  __ccgo_ts + 2316,
		Ftype1: int8('F'),
	},
	340: {
		Fword:  __ccgo_ts + 2323,
		Ftype1: int8('F'),
	},
	341: {
		Fword:  __ccgo_ts + 2330,
		Ftype1: int8('F'),
	},
	342: {
		Fword:  __ccgo_ts + 2337,
		Ftype1: int8('F'),
	},
	343: {
		Fword:  __ccgo_ts + 2344,
		Ftype1: int8('F'),
	},
	344: {
		Fword:  __ccgo_ts + 2349,
		Ftype1: int8('F'),
	},
	345: {
		Fword:  __ccgo_ts + 2356,
		Ftype1: int8('F'),
	},
	346: {
		Fword:  __ccgo_ts + 2363,
		Ftype1: int8('F'),
	},
	347: {
		Fword:  __ccgo_ts + 2370,
		Ftype1: int8('F'),
	},
	348: {
		Fword:  __ccgo_ts + 2377,
		Ftype1: int8('F'),
	},
	349: {
		Fword:  __ccgo_ts + 2384,
		Ftype1: int8('F'),
	},
	350: {
		Fword:  __ccgo_ts + 2391,
		Ftype1: int8('F'),
	},
	351: {
		Fword:  __ccgo_ts + 2398,
		Ftype1: int8('F'),
	},
	352: {
		Fword:  __ccgo_ts + 2405,
		Ftype1: int8('F'),
	},
	353: {
		Fword:  __ccgo_ts + 2412,
		Ftype1: int8('F'),
	},
	354: {
		Fword:  __ccgo_ts + 2419,
		Ftype1: int8('F'),
	},
	355: {
		Fword:  __ccgo_ts + 2425,
		Ftype1: int8('F'),
	},
	356: {
		Fword:  __ccgo_ts + 2432,
		Ftype1: int8('F'),
	},
	357: {
		Fword:  __ccgo_ts + 2439,
		Ftype1: int8('F'),
	},
	358: {
		Fword:  __ccgo_ts + 2445,
		Ftype1: int8('F'),
	},
	359: {
		Fword:  __ccgo_ts + 2452,
		Ftype1: int8('F'),
	},
	360: {
		Fword:  __ccgo_ts + 2459,
		Ftype1: int8('F'),
	},
	361: {
		Fword:  __ccgo_ts + 2466,
		Ftype1: int8('F'),
	},
	362: {
		Fword:  __ccgo_ts + 2473,
		Ftype1: int8('F'),
	},
	363: {
		Fword:  __ccgo_ts + 2480,
		Ftype1: int8('F'),
	},
	364: {
		Fword:  __ccgo_ts + 2487,
		Ftype1: int8('F'),
	},
	365: {
		Fword:  __ccgo_ts + 2494,
		Ftype1: int8('F'),
	},
	366: {
		Fword:  __ccgo_ts + 2501,
		Ftype1: int8('F'),
	},
	367: {
		Fword:  __ccgo_ts + 2508,
		Ftype1: int8('F'),
	},
	368: {
		Fword:  __ccgo_ts + 2514,
		Ftype1: int8('F'),
	},
	369: {
		Fword:  __ccgo_ts + 2521,
		Ftype1: int8('F'),
	},
	370: {
		Fword:  __ccgo_ts + 2528,
		Ftype1: int8('F'),
	},
	371: {
		Fword:  __ccgo_ts + 2535,
		Ftype1: int8('F'),
	},
	372: {
		Fword:  __ccgo_ts + 2542,
		Ftype1: int8('F'),
	},
	373: {
		Fword:  __ccgo_ts + 2549,
		Ftype1: int8('F'),
	},
	374: {
		Fword:  __ccgo_ts + 2556,
		Ftype1: int8('F'),
	},
	375: {
		Fword:  __ccgo_ts + 2563,
		Ftype1: int8('F'),
	},
	376: {
		Fword:  __ccgo_ts + 2570,
		Ftype1: int8('F'),
	},
	377: {
		Fword:  __ccgo_ts + 2577,
		Ftype1: int8('F'),
	},
	378: {
		Fword:  __ccgo_ts + 2584,
		Ftype1: int8('F'),
	},
	379: {
		Fword:  __ccgo_ts + 2591,
		Ftype1: int8('F'),
	},
	380: {
		Fword:  __ccgo_ts + 2598,
		Ftype1: int8('F'),
	},
	381: {
		Fword:  __ccgo_ts + 2605,
		Ftype1: int8('F'),
	},
	382: {
		Fword:  __ccgo_ts + 2612,
		Ftype1: int8('F'),
	},
	383: {
		Fword:  __ccgo_ts + 2619,
		Ftype1: int8('F'),
	},
	384: {
		Fword:  __ccgo_ts + 2626,
		Ftype1: int8('F'),
	},
	385: {
		Fword:  __ccgo_ts + 2632,
		Ftype1: int8('F'),
	},
	386: {
		Fword:  __ccgo_ts + 2639,
		Ftype1: int8('F'),
	},
	387: {
		Fword:  __ccgo_ts + 2646,
		Ftype1: int8('F'),
	},
	388: {
		Fword:  __ccgo_ts + 2653,
		Ftype1: int8('F'),
	},
	389: {
		Fword:  __ccgo_ts + 2660,
		Ftype1: int8('F'),
	},
	390: {
		Fword:  __ccgo_ts + 2666,
		Ftype1: int8('F'),
	},
	391: {
		Fword:  __ccgo_ts + 2673,
		Ftype1: int8('F'),
	},
	392: {
		Fword:  __ccgo_ts + 2680,
		Ftype1: int8('F'),
	},
	393: {
		Fword:  __ccgo_ts + 2687,
		Ftype1: int8('F'),
	},
	394: {
		Fword:  __ccgo_ts + 2692,
		Ftype1: int8('F'),
	},
	395: {
		Fword:  __ccgo_ts + 2699,
		Ftype1: int8('F'),
	},
	396: {
		Fword:  __ccgo_ts + 2706,
		Ftype1: int8('F'),
	},
	397: {
		Fword:  __ccgo_ts + 2713,
		Ftype1: int8('F'),
	},
	398: {
		Fword:  __ccgo_ts + 2720,
		Ftype1: int8('F'),
	},
	399: {
		Fword:  __ccgo_ts + 2727,
		Ftype1: int8('F'),
	},
	400: {
		Fword:  __ccgo_ts + 2734,
		Ftype1: int8('F'),
	},
	401: {
		Fword:  __ccgo_ts + 2741,
		Ftype1: int8('F'),
	},
	402: {
		Fword:  __ccgo_ts + 2748,
		Ftype1: int8('F'),
	},
	403: {
		Fword:  __ccgo_ts + 2755,
		Ftype1: int8('F'),
	},
	404: {
		Fword:  __ccgo_ts + 2762,
		Ftype1: int8('F'),
	},
	405: {
		Fword:  __ccgo_ts + 2768,
		Ftype1: int8('F'),
	},
	406: {
		Fword:  __ccgo_ts + 2775,
		Ftype1: int8('F'),
	},
	407: {
		Fword:  __ccgo_ts + 2782,
		Ftype1: int8('F'),
	},
	408: {
		Fword:  __ccgo_ts + 2789,
		Ftype1: int8('F'),
	},
	409: {
		Fword:  __ccgo_ts + 2796,
		Ftype1: int8('F'),
	},
	410: {
		Fword:  __ccgo_ts + 2803,
		Ftype1: int8('F'),
	},
	411: {
		Fword:  __ccgo_ts + 2810,
		Ftype1: int8('F'),
	},
	412: {
		Fword:  __ccgo_ts + 2817,
		Ftype1: int8('F'),
	},
	413: {
		Fword:  __ccgo_ts + 2824,
		Ftype1: int8('F'),
	},
	414: {
		Fword:  __ccgo_ts + 2831,
		Ftype1: int8('F'),
	},
	415: {
		Fword:  __ccgo_ts + 2837,
		Ftype1: int8('F'),
	},
	416: {
		Fword:  __ccgo_ts + 2844,
		Ftype1: int8('F'),
	},
	417: {
		Fword:  __ccgo_ts + 2851,
		Ftype1: int8('F'),
	},
	418: {
		Fword:  __ccgo_ts + 2858,
		Ftype1: int8('F'),
	},
	419: {
		Fword:  __ccgo_ts + 2865,
		Ftype1: int8('F'),
	},
	420: {
		Fword:  __ccgo_ts + 2872,
		Ftype1: int8('F'),
	},
	421: {
		Fword:  __ccgo_ts + 2879,
		Ftype1: int8('F'),
	},
	422: {
		Fword:  __ccgo_ts + 2886,
		Ftype1: int8('F'),
	},
	423: {
		Fword:  __ccgo_ts + 2893,
		Ftype1: int8('F'),
	},
	424: {
		Fword:  __ccgo_ts + 2900,
		Ftype1: int8('F'),
	},
	425: {
		Fword:  __ccgo_ts + 2907,
		Ftype1: int8('F'),
	},
	426: {
		Fword:  __ccgo_ts + 2914,
		Ftype1: int8('F'),
	},
	427: {
		Fword:  __ccgo_ts + 2921,
		Ftype1: int8('F'),
	},
	428: {
		Fword:  __ccgo_ts + 2927,
		Ftype1: int8('F'),
	},
	429: {
		Fword:  __ccgo_ts + 2934,
		Ftype1: int8('F'),
	},
	430: {
		Fword:  __ccgo_ts + 2941,
		Ftype1: int8('F'),
	},
	431: {
		Fword:  __ccgo_ts + 2948,
		Ftype1: int8('F'),
	},
	432: {
		Fword:  __ccgo_ts + 2955,
		Ftype1: int8('F'),
	},
	433: {
		Fword:  __ccgo_ts + 2961,
		Ftype1: int8('F'),
	},
	434: {
		Fword:  __ccgo_ts + 2968,
		Ftype1: int8('F'),
	},
	435: {
		Fword:  __ccgo_ts + 2975,
		Ftype1: int8('F'),
	},
	436: {
		Fword:  __ccgo_ts + 2982,
		Ftype1: int8('F'),
	},
	437: {
		Fword:  __ccgo_ts + 2989,
		Ftype1: int8('F'),
	},
	438: {
		Fword:  __ccgo_ts + 2996,
		Ftype1: int8('F'),
	},
	439: {
		Fword:  __ccgo_ts + 3003,
		Ftype1: int8('F'),
	},
	440: {
		Fword:  __ccgo_ts + 3010,
		Ftype1: int8('F'),
	},
	441: {
		Fword:  __ccgo_ts + 3017,
		Ftype1: int8('F'),
	},
	442: {
		Fword:  __ccgo_ts + 3024,
		Ftype1: int8('F'),
	},
	443: {
		Fword:  __ccgo_ts + 3031,
		Ftype1: int8('F'),
	},
	444: {
		Fword:  __ccgo_ts + 3038,
		Ftype1: int8('F'),
	},
	445: {
		Fword:  __ccgo_ts + 3045,
		Ftype1: int8('F'),
	},
	446: {
		Fword:  __ccgo_ts + 3052,
		Ftype1: int8('F'),
	},
	447: {
		Fword:  __ccgo_ts + 3059,
		Ftype1: int8('F'),
	},
	448: {
		Fword:  __ccgo_ts + 3066,
		Ftype1: int8('F'),
	},
	449: {
		Fword:  __ccgo_ts + 3072,
		Ftype1: int8('F'),
	},
	450: {
		Fword:  __ccgo_ts + 3079,
		Ftype1: int8('F'),
	},
	451: {
		Fword:  __ccgo_ts + 3086,
		Ftype1: int8('F'),
	},
	452: {
		Fword:  __ccgo_ts + 3093,
		Ftype1: int8('F'),
	},
	453: {
		Fword:  __ccgo_ts + 3100,
		Ftype1: int8('F'),
	},
	454: {
		Fword:  __ccgo_ts + 3107,
		Ftype1: int8('F'),
	},
	455: {
		Fword:  __ccgo_ts + 3114,
		Ftype1: int8('F'),
	},
	456: {
		Fword:  __ccgo_ts + 3121,
		Ftype1: int8('F'),
	},
	457: {
		Fword:  __ccgo_ts + 3128,
		Ftype1: int8('F'),
	},
	458: {
		Fword:  __ccgo_ts + 3135,
		Ftype1: int8('F'),
	},
	459: {
		Fword:  __ccgo_ts + 3141,
		Ftype1: int8('F'),
	},
	460: {
		Fword:  __ccgo_ts + 3148,
		Ftype1: int8('F'),
	},
	461: {
		Fword:  __ccgo_ts + 3155,
		Ftype1: int8('F'),
	},
	462: {
		Fword:  __ccgo_ts + 3162,
		Ftype1: int8('F'),
	},
	463: {
		Fword:  __ccgo_ts + 3169,
		Ftype1: int8('F'),
	},
	464: {
		Fword:  __ccgo_ts + 3176,
		Ftype1: int8('F'),
	},
	465: {
		Fword:  __ccgo_ts + 3183,
		Ftype1: int8('F'),
	},
	466: {
		Fword:  __ccgo_ts + 3190,
		Ftype1: int8('F'),
	},
	467: {
		Fword:  __ccgo_ts + 3197,
		Ftype1: int8('F'),
	},
	468: {
		Fword:  __ccgo_ts + 3203,
		Ftype1: int8('F'),
	},
	469: {
		Fword:  __ccgo_ts + 3210,
		Ftype1: int8('F'),
	},
	470: {
		Fword:  __ccgo_ts + 3217,
		Ftype1: int8('F'),
	},
	471: {
		Fword:  __ccgo_ts + 3224,
		Ftype1: int8('F'),
	},
	472: {
		Fword:  __ccgo_ts + 3231,
		Ftype1: int8('F'),
	},
	473: {
		Fword:  __ccgo_ts + 3238,
		Ftype1: int8('F'),
	},
	474: {
		Fword:  __ccgo_ts + 3245,
		Ftype1: int8('F'),
	},
	475: {
		Fword:  __ccgo_ts + 3252,
		Ftype1: int8('F'),
	},
	476: {
		Fword:  __ccgo_ts + 3259,
		Ftype1: int8('F'),
	},
	477: {
		Fword:  __ccgo_ts + 3265,
		Ftype1: int8('F'),
	},
	478: {
		Fword:  __ccgo_ts + 3272,
		Ftype1: int8('F'),
	},
	479: {
		Fword:  __ccgo_ts + 3279,
		Ftype1: int8('F'),
	},
	480: {
		Fword:  __ccgo_ts + 3286,
		Ftype1: int8('F'),
	},
	481: {
		Fword:  __ccgo_ts + 3293,
		Ftype1: int8('F'),
	},
	482: {
		Fword:  __ccgo_ts + 3300,
		Ftype1: int8('F'),
	},
	483: {
		Fword:  __ccgo_ts + 3307,
		Ftype1: int8('F'),
	},
	484: {
		Fword:  __ccgo_ts + 3314,
		Ftype1: int8('F'),
	},
	485: {
		Fword:  __ccgo_ts + 3321,
		Ftype1: int8('F'),
	},
	486: {
		Fword:  __ccgo_ts + 3328,
		Ftype1: int8('F'),
	},
	487: {
		Fword:  __ccgo_ts + 3335,
		Ftype1: int8('F'),
	},
	488: {
		Fword:  __ccgo_ts + 3342,
		Ftype1: int8('F'),
	},
	489: {
		Fword:  __ccgo_ts + 3349,
		Ftype1: int8('F'),
	},
	490: {
		Fword:  __ccgo_ts + 3356,
		Ftype1: int8('F'),
	},
	491: {
		Fword:  __ccgo_ts + 3363,
		Ftype1: int8('F'),
	},
	492: {
		Fword:  __ccgo_ts + 3370,
		Ftype1: int8('F'),
	},
	493: {
		Fword:  __ccgo_ts + 3377,
		Ftype1: int8('F'),
	},
	494: {
		Fword:  __ccgo_ts + 3384,
		Ftype1: int8('F'),
	},
	495: {
		Fword:  __ccgo_ts + 3391,
		Ftype1: int8('F'),
	},
	496: {
		Fword:  __ccgo_ts + 3398,
		Ftype1: int8('F'),
	},
	497: {
		Fword:  __ccgo_ts + 3405,
		Ftype1: int8('F'),
	},
	498: {
		Fword:  __ccgo_ts + 3412,
		Ftype1: int8('F'),
	},
	499: {
		Fword:  __ccgo_ts + 3419,
		Ftype1: int8('F'),
	},
	500: {
		Fword:  __ccgo_ts + 3426,
		Ftype1: int8('F'),
	},
	501: {
		Fword:  __ccgo_ts + 3433,
		Ftype1: int8('F'),
	},
	502: {
		Fword:  __ccgo_ts + 3440,
		Ftype1: int8('F'),
	},
	503: {
		Fword:  __ccgo_ts + 3447,
		Ftype1: int8('F'),
	},
	504: {
		Fword:  __ccgo_ts + 3454,
		Ftype1: int8('F'),
	},
	505: {
		Fword:  __ccgo_ts + 3461,
		Ftype1: int8('F'),
	},
	506: {
		Fword:  __ccgo_ts + 3468,
		Ftype1: int8('F'),
	},
	507: {
		Fword:  __ccgo_ts + 3475,
		Ftype1: int8('F'),
	},
	508: {
		Fword:  __ccgo_ts + 3482,
		Ftype1: int8('F'),
	},
	509: {
		Fword:  __ccgo_ts + 3489,
		Ftype1: int8('F'),
	},
	510: {
		Fword:  __ccgo_ts + 3496,
		Ftype1: int8('F'),
	},
	511: {
		Fword:  __ccgo_ts + 3503,
		Ftype1: int8('F'),
	},
	512: {
		Fword:  __ccgo_ts + 3509,
		Ftype1: int8('F'),
	},
	513: {
		Fword:  __ccgo_ts + 3516,
		Ftype1: int8('F'),
	},
	514: {
		Fword:  __ccgo_ts + 3523,
		Ftype1: int8('F'),
	},
	515: {
		Fword:  __ccgo_ts + 3530,
		Ftype1: int8('F'),
	},
	516: {
		Fword:  __ccgo_ts + 3537,
		Ftype1: int8('F'),
	},
	517: {
		Fword:  __ccgo_ts + 3544,
		Ftype1: int8('F'),
	},
	518: {
		Fword:  __ccgo_ts + 3551,
		Ftype1: int8('F'),
	},
	519: {
		Fword:  __ccgo_ts + 3558,
		Ftype1: int8('F'),
	},
	520: {
		Fword:  __ccgo_ts + 3565,
		Ftype1: int8('F'),
	},
	521: {
		Fword:  __ccgo_ts + 3571,
		Ftype1: int8('F'),
	},
	522: {
		Fword:  __ccgo_ts + 3578,
		Ftype1: int8('F'),
	},
	523: {
		Fword:  __ccgo_ts + 3585,
		Ftype1: int8('F'),
	},
	524: {
		Fword:  __ccgo_ts + 3592,
		Ftype1: int8('F'),
	},
	525: {
		Fword:  __ccgo_ts + 3599,
		Ftype1: int8('F'),
	},
	526: {
		Fword:  __ccgo_ts + 3606,
		Ftype1: int8('F'),
	},
	527: {
		Fword:  __ccgo_ts + 3613,
		Ftype1: int8('F'),
	},
	528: {
		Fword:  __ccgo_ts + 3619,
		Ftype1: int8('F'),
	},
	529: {
		Fword:  __ccgo_ts + 3626,
		Ftype1: int8('F'),
	},
	530: {
		Fword:  __ccgo_ts + 3633,
		Ftype1: int8('F'),
	},
	531: {
		Fword:  __ccgo_ts + 3640,
		Ftype1: int8('F'),
	},
	532: {
		Fword:  __ccgo_ts + 3647,
		Ftype1: int8('F'),
	},
	533: {
		Fword:  __ccgo_ts + 3654,
		Ftype1: int8('F'),
	},
	534: {
		Fword:  __ccgo_ts + 3661,
		Ftype1: int8('F'),
	},
	535: {
		Fword:  __ccgo_ts + 3667,
		Ftype1: int8('F'),
	},
	536: {
		Fword:  __ccgo_ts + 3674,
		Ftype1: int8('F'),
	},
	537: {
		Fword:  __ccgo_ts + 3681,
		Ftype1: int8('F'),
	},
	538: {
		Fword:  __ccgo_ts + 3688,
		Ftype1: int8('F'),
	},
	539: {
		Fword:  __ccgo_ts + 3695,
		Ftype1: int8('F'),
	},
	540: {
		Fword:  __ccgo_ts + 3702,
		Ftype1: int8('F'),
	},
	541: {
		Fword:  __ccgo_ts + 3709,
		Ftype1: int8('F'),
	},
	542: {
		Fword:  __ccgo_ts + 3714,
		Ftype1: int8('F'),
	},
	543: {
		Fword:  __ccgo_ts + 3721,
		Ftype1: int8('F'),
	},
	544: {
		Fword:  __ccgo_ts + 3728,
		Ftype1: int8('F'),
	},
	545: {
		Fword:  __ccgo_ts + 3735,
		Ftype1: int8('F'),
	},
	546: {
		Fword:  __ccgo_ts + 3742,
		Ftype1: int8('F'),
	},
	547: {
		Fword:  __ccgo_ts + 3749,
		Ftype1: int8('F'),
	},
	548: {
		Fword:  __ccgo_ts + 3756,
		Ftype1: int8('F'),
	},
	549: {
		Fword:  __ccgo_ts + 3763,
		Ftype1: int8('F'),
	},
	550: {
		Fword:  __ccgo_ts + 3770,
		Ftype1: int8('F'),
	},
	551: {
		Fword:  __ccgo_ts + 3777,
		Ftype1: int8('F'),
	},
	552: {
		Fword:  __ccgo_ts + 3784,
		Ftype1: int8('F'),
	},
	553: {
		Fword:  __ccgo_ts + 3791,
		Ftype1: int8('F'),
	},
	554: {
		Fword:  __ccgo_ts + 3798,
		Ftype1: int8('F'),
	},
	555: {
		Fword:  __ccgo_ts + 3805,
		Ftype1: int8('F'),
	},
	556: {
		Fword:  __ccgo_ts + 3812,
		Ftype1: int8('F'),
	},
	557: {
		Fword:  __ccgo_ts + 3819,
		Ftype1: int8('F'),
	},
	558: {
		Fword:  __ccgo_ts + 3826,
		Ftype1: int8('F'),
	},
	559: {
		Fword:  __ccgo_ts + 3833,
		Ftype1: int8('F'),
	},
	560: {
		Fword:  __ccgo_ts + 3840,
		Ftype1: int8('F'),
	},
	561: {
		Fword:  __ccgo_ts + 3847,
		Ftype1: int8('F'),
	},
	562: {
		Fword:  __ccgo_ts + 3854,
		Ftype1: int8('F'),
	},
	563: {
		Fword:  __ccgo_ts + 3861,
		Ftype1: int8('F'),
	},
	564: {
		Fword:  __ccgo_ts + 3868,
		Ftype1: int8('F'),
	},
	565: {
		Fword:  __ccgo_ts + 3875,
		Ftype1: int8('F'),
	},
	566: {
		Fword:  __ccgo_ts + 3882,
		Ftype1: int8('F'),
	},
	567: {
		Fword:  __ccgo_ts + 3889,
		Ftype1: int8('F'),
	},
	568: {
		Fword:  __ccgo_ts + 3896,
		Ftype1: int8('F'),
	},
	569: {
		Fword:  __ccgo_ts + 3903,
		Ftype1: int8('F'),
	},
	570: {
		Fword:  __ccgo_ts + 3910,
		Ftype1: int8('F'),
	},
	571: {
		Fword:  __ccgo_ts + 3917,
		Ftype1: int8('F'),
	},
	572: {
		Fword:  __ccgo_ts + 3924,
		Ftype1: int8('F'),
	},
	573: {
		Fword:  __ccgo_ts + 3931,
		Ftype1: int8('F'),
	},
	574: {
		Fword:  __ccgo_ts + 3938,
		Ftype1: int8('F'),
	},
	575: {
		Fword:  __ccgo_ts + 3945,
		Ftype1: int8('F'),
	},
	576: {
		Fword:  __ccgo_ts + 3952,
		Ftype1: int8('F'),
	},
	577: {
		Fword:  __ccgo_ts + 3959,
		Ftype1: int8('F'),
	},
	578: {
		Fword:  __ccgo_ts + 3966,
		Ftype1: int8('F'),
	},
	579: {
		Fword:  __ccgo_ts + 3973,
		Ftype1: int8('F'),
	},
	580: {
		Fword:  __ccgo_ts + 3980,
		Ftype1: int8('F'),
	},
	581: {
		Fword:  __ccgo_ts + 3987,
		Ftype1: int8('F'),
	},
	582: {
		Fword:  __ccgo_ts + 3994,
		Ftype1: int8('F'),
	},
	583: {
		Fword:  __ccgo_ts + 4001,
		Ftype1: int8('F'),
	},
	584: {
		Fword:  __ccgo_ts + 4008,
		Ftype1: int8('F'),
	},
	585: {
		Fword:  __ccgo_ts + 4015,
		Ftype1: int8('F'),
	},
	586: {
		Fword:  __ccgo_ts + 4022,
		Ftype1: int8('F'),
	},
	587: {
		Fword:  __ccgo_ts + 4029,
		Ftype1: int8('F'),
	},
	588: {
		Fword:  __ccgo_ts + 4036,
		Ftype1: int8('F'),
	},
	589: {
		Fword:  __ccgo_ts + 4043,
		Ftype1: int8('F'),
	},
	590: {
		Fword:  __ccgo_ts + 4050,
		Ftype1: int8('F'),
	},
	591: {
		Fword:  __ccgo_ts + 4057,
		Ftype1: int8('F'),
	},
	592: {
		Fword:  __ccgo_ts + 4064,
		Ftype1: int8('F'),
	},
	593: {
		Fword:  __ccgo_ts + 4071,
		Ftype1: int8('F'),
	},
	594: {
		Fword:  __ccgo_ts + 4078,
		Ftype1: int8('F'),
	},
	595: {
		Fword:  __ccgo_ts + 4085,
		Ftype1: int8('F'),
	},
	596: {
		Fword:  __ccgo_ts + 4092,
		Ftype1: int8('F'),
	},
	597: {
		Fword:  __ccgo_ts + 4099,
		Ftype1: int8('F'),
	},
	598: {
		Fword:  __ccgo_ts + 4106,
		Ftype1: int8('F'),
	},
	599: {
		Fword:  __ccgo_ts + 4113,
		Ftype1: int8('F'),
	},
	600: {
		Fword:  __ccgo_ts + 4120,
		Ftype1: int8('F'),
	},
	601: {
		Fword:  __ccgo_ts + 4127,
		Ftype1: int8('F'),
	},
	602: {
		Fword:  __ccgo_ts + 4134,
		Ftype1: int8('F'),
	},
	603: {
		Fword:  __ccgo_ts + 4141,
		Ftype1: int8('F'),
	},
	604: {
		Fword:  __ccgo_ts + 4148,
		Ftype1: int8('F'),
	},
	605: {
		Fword:  __ccgo_ts + 4155,
		Ftype1: int8('F'),
	},
	606: {
		Fword:  __ccgo_ts + 4162,
		Ftype1: int8('F'),
	},
	607: {
		Fword:  __ccgo_ts + 4169,
		Ftype1: int8('F'),
	},
	608: {
		Fword:  __ccgo_ts + 4175,
		Ftype1: int8('F'),
	},
	609: {
		Fword:  __ccgo_ts + 4182,
		Ftype1: int8('F'),
	},
	610: {
		Fword:  __ccgo_ts + 4189,
		Ftype1: int8('F'),
	},
	611: {
		Fword:  __ccgo_ts + 4196,
		Ftype1: int8('F'),
	},
	612: {
		Fword:  __ccgo_ts + 4203,
		Ftype1: int8('F'),
	},
	613: {
		Fword:  __ccgo_ts + 4210,
		Ftype1: int8('F'),
	},
	614: {
		Fword:  __ccgo_ts + 4217,
		Ftype1: int8('F'),
	},
	615: {
		Fword:  __ccgo_ts + 4224,
		Ftype1: int8('F'),
	},
	616: {
		Fword:  __ccgo_ts + 4231,
		Ftype1: int8('F'),
	},
	617: {
		Fword:  __ccgo_ts + 4238,
		Ftype1: int8('F'),
	},
	618: {
		Fword:  __ccgo_ts + 4245,
		Ftype1: int8('F'),
	},
	619: {
		Fword:  __ccgo_ts + 4252,
		Ftype1: int8('F'),
	},
	620: {
		Fword:  __ccgo_ts + 4259,
		Ftype1: int8('F'),
	},
	621: {
		Fword:  __ccgo_ts + 4266,
		Ftype1: int8('F'),
	},
	622: {
		Fword:  __ccgo_ts + 4273,
		Ftype1: int8('F'),
	},
	623: {
		Fword:  __ccgo_ts + 4280,
		Ftype1: int8('F'),
	},
	624: {
		Fword:  __ccgo_ts + 4286,
		Ftype1: int8('F'),
	},
	625: {
		Fword:  __ccgo_ts + 4293,
		Ftype1: int8('F'),
	},
	626: {
		Fword:  __ccgo_ts + 4300,
		Ftype1: int8('F'),
	},
	627: {
		Fword:  __ccgo_ts + 4307,
		Ftype1: int8('F'),
	},
	628: {
		Fword:  __ccgo_ts + 4314,
		Ftype1: int8('F'),
	},
	629: {
		Fword:  __ccgo_ts + 4321,
		Ftype1: int8('F'),
	},
	630: {
		Fword:  __ccgo_ts + 4328,
		Ftype1: int8('F'),
	},
	631: {
		Fword:  __ccgo_ts + 4335,
		Ftype1: int8('F'),
	},
	632: {
		Fword:  __ccgo_ts + 4341,
		Ftype1: int8('F'),
	},
	633: {
		Fword:  __ccgo_ts + 4348,
		Ftype1: int8('F'),
	},
	634: {
		Fword:  __ccgo_ts + 4355,
		Ftype1: int8('F'),
	},
	635: {
		Fword:  __ccgo_ts + 4362,
		Ftype1: int8('F'),
	},
	636: {
		Fword:  __ccgo_ts + 4369,
		Ftype1: int8('F'),
	},
	637: {
		Fword:  __ccgo_ts + 4376,
		Ftype1: int8('F'),
	},
	638: {
		Fword:  __ccgo_ts + 4383,
		Ftype1: int8('F'),
	},
	639: {
		Fword:  __ccgo_ts + 4390,
		Ftype1: int8('F'),
	},
	640: {
		Fword:  __ccgo_ts + 4397,
		Ftype1: int8('F'),
	},
	641: {
		Fword:  __ccgo_ts + 4404,
		Ftype1: int8('F'),
	},
	642: {
		Fword:  __ccgo_ts + 4411,
		Ftype1: int8('F'),
	},
	643: {
		Fword:  __ccgo_ts + 4418,
		Ftype1: int8('F'),
	},
	644: {
		Fword:  __ccgo_ts + 4425,
		Ftype1: int8('F'),
	},
	645: {
		Fword:  __ccgo_ts + 4432,
		Ftype1: int8('F'),
	},
	646: {
		Fword:  __ccgo_ts + 4439,
		Ftype1: int8('F'),
	},
	647: {
		Fword:  __ccgo_ts + 4446,
		Ftype1: int8('F'),
	},
	648: {
		Fword:  __ccgo_ts + 4453,
		Ftype1: int8('F'),
	},
	649: {
		Fword:  __ccgo_ts + 4460,
		Ftype1: int8('F'),
	},
	650: {
		Fword:  __ccgo_ts + 4467,
		Ftype1: int8('F'),
	},
	651: {
		Fword:  __ccgo_ts + 4474,
		Ftype1: int8('F'),
	},
	652: {
		Fword:  __ccgo_ts + 4481,
		Ftype1: int8('F'),
	},
	653: {
		Fword:  __ccgo_ts + 4488,
		Ftype1: int8('F'),
	},
	654: {
		Fword:  __ccgo_ts + 4495,
		Ftype1: int8('F'),
	},
	655: {
		Fword:  __ccgo_ts + 4502,
		Ftype1: int8('F'),
	},
	656: {
		Fword:  __ccgo_ts + 4509,
		Ftype1: int8('F'),
	},
	657: {
		Fword:  __ccgo_ts + 4516,
		Ftype1: int8('F'),
	},
	658: {
		Fword:  __ccgo_ts + 4523,
		Ftype1: int8('F'),
	},
	659: {
		Fword:  __ccgo_ts + 4530,
		Ftype1: int8('F'),
	},
	660: {
		Fword:  __ccgo_ts + 4537,
		Ftype1: int8('F'),
	},
	661: {
		Fword:  __ccgo_ts + 4544,
		Ftype1: int8('F'),
	},
	662: {
		Fword:  __ccgo_ts + 4551,
		Ftype1: int8('F'),
	},
	663: {
		Fword:  __ccgo_ts + 4558,
		Ftype1: int8('F'),
	},
	664: {
		Fword:  __ccgo_ts + 4565,
		Ftype1: int8('F'),
	},
	665: {
		Fword:  __ccgo_ts + 4572,
		Ftype1: int8('F'),
	},
	666: {
		Fword:  __ccgo_ts + 4579,
		Ftype1: int8('F'),
	},
	667: {
		Fword:  __ccgo_ts + 4586,
		Ftype1: int8('F'),
	},
	668: {
		Fword:  __ccgo_ts + 4593,
		Ftype1: int8('F'),
	},
	669: {
		Fword:  __ccgo_ts + 4600,
		Ftype1: int8('F'),
	},
	670: {
		Fword:  __ccgo_ts + 4607,
		Ftype1: int8('F'),
	},
	671: {
		Fword:  __ccgo_ts + 4614,
		Ftype1: int8('F'),
	},
	672: {
		Fword:  __ccgo_ts + 4621,
		Ftype1: int8('F'),
	},
	673: {
		Fword:  __ccgo_ts + 4628,
		Ftype1: int8('F'),
	},
	674: {
		Fword:  __ccgo_ts + 4635,
		Ftype1: int8('F'),
	},
	675: {
		Fword:  __ccgo_ts + 4642,
		Ftype1: int8('F'),
	},
	676: {
		Fword:  __ccgo_ts + 4649,
		Ftype1: int8('F'),
	},
	677: {
		Fword:  __ccgo_ts + 4656,
		Ftype1: int8('F'),
	},
	678: {
		Fword:  __ccgo_ts + 4663,
		Ftype1: int8('F'),
	},
	679: {
		Fword:  __ccgo_ts + 4670,
		Ftype1: int8('F'),
	},
	680: {
		Fword:  __ccgo_ts + 4677,
		Ftype1: int8('F'),
	},
	681: {
		Fword:  __ccgo_ts + 4684,
		Ftype1: int8('F'),
	},
	682: {
		Fword:  __ccgo_ts + 4691,
		Ftype1: int8('F'),
	},
	683: {
		Fword:  __ccgo_ts + 4698,
		Ftype1: int8('F'),
	},
	684: {
		Fword:  __ccgo_ts + 4705,
		Ftype1: int8('F'),
	},
	685: {
		Fword:  __ccgo_ts + 4712,
		Ftype1: int8('F'),
	},
	686: {
		Fword:  __ccgo_ts + 4719,
		Ftype1: int8('F'),
	},
	687: {
		Fword:  __ccgo_ts + 4726,
		Ftype1: int8('F'),
	},
	688: {
		Fword:  __ccgo_ts + 4733,
		Ftype1: int8('F'),
	},
	689: {
		Fword:  __ccgo_ts + 4740,
		Ftype1: int8('F'),
	},
	690: {
		Fword:  __ccgo_ts + 4747,
		Ftype1: int8('F'),
	},
	691: {
		Fword:  __ccgo_ts + 4754,
		Ftype1: int8('F'),
	},
	692: {
		Fword:  __ccgo_ts + 4761,
		Ftype1: int8('F'),
	},
	693: {
		Fword:  __ccgo_ts + 4768,
		Ftype1: int8('F'),
	},
	694: {
		Fword:  __ccgo_ts + 4775,
		Ftype1: int8('F'),
	},
	695: {
		Fword:  __ccgo_ts + 4782,
		Ftype1: int8('F'),
	},
	696: {
		Fword:  __ccgo_ts + 4789,
		Ftype1: int8('F'),
	},
	697: {
		Fword:  __ccgo_ts + 4796,
		Ftype1: int8('F'),
	},
	698: {
		Fword:  __ccgo_ts + 4803,
		Ftype1: int8('F'),
	},
	699: {
		Fword:  __ccgo_ts + 4810,
		Ftype1: int8('F'),
	},
	700: {
		Fword:  __ccgo_ts + 4817,
		Ftype1: int8('F'),
	},
	701: {
		Fword:  __ccgo_ts + 4824,
		Ftype1: int8('F'),
	},
	702: {
		Fword:  __ccgo_ts + 4831,
		Ftype1: int8('F'),
	},
	703: {
		Fword:  __ccgo_ts + 4838,
		Ftype1: int8('F'),
	},
	704: {
		Fword:  __ccgo_ts + 4845,
		Ftype1: int8('F'),
	},
	705: {
		Fword:  __ccgo_ts + 4852,
		Ftype1: int8('F'),
	},
	706: {
		Fword:  __ccgo_ts + 4859,
		Ftype1: int8('F'),
	},
	707: {
		Fword:  __ccgo_ts + 4866,
		Ftype1: int8('F'),
	},
	708: {
		Fword:  __ccgo_ts + 4873,
		Ftype1: int8('F'),
	},
	709: {
		Fword:  __ccgo_ts + 4880,
		Ftype1: int8('F'),
	},
	710: {
		Fword:  __ccgo_ts + 4887,
		Ftype1: int8('F'),
	},
	711: {
		Fword:  __ccgo_ts + 4894,
		Ftype1: int8('F'),
	},
	712: {
		Fword:  __ccgo_ts + 4901,
		Ftype1: int8('F'),
	},
	713: {
		Fword:  __ccgo_ts + 4908,
		Ftype1: int8('F'),
	},
	714: {
		Fword:  __ccgo_ts + 4915,
		Ftype1: int8('F'),
	},
	715: {
		Fword:  __ccgo_ts + 4922,
		Ftype1: int8('F'),
	},
	716: {
		Fword:  __ccgo_ts + 4929,
		Ftype1: int8('F'),
	},
	717: {
		Fword:  __ccgo_ts + 4936,
		Ftype1: int8('F'),
	},
	718: {
		Fword:  __ccgo_ts + 4943,
		Ftype1: int8('F'),
	},
	719: {
		Fword:  __ccgo_ts + 4950,
		Ftype1: int8('F'),
	},
	720: {
		Fword:  __ccgo_ts + 4957,
		Ftype1: int8('F'),
	},
	721: {
		Fword:  __ccgo_ts + 4964,
		Ftype1: int8('F'),
	},
	722: {
		Fword:  __ccgo_ts + 4971,
		Ftype1: int8('F'),
	},
	723: {
		Fword:  __ccgo_ts + 4978,
		Ftype1: int8('F'),
	},
	724: {
		Fword:  __ccgo_ts + 4985,
		Ftype1: int8('F'),
	},
	725: {
		Fword:  __ccgo_ts + 4992,
		Ftype1: int8('F'),
	},
	726: {
		Fword:  __ccgo_ts + 4998,
		Ftype1: int8('F'),
	},
	727: {
		Fword:  __ccgo_ts + 5005,
		Ftype1: int8('F'),
	},
	728: {
		Fword:  __ccgo_ts + 5012,
		Ftype1: int8('F'),
	},
	729: {
		Fword:  __ccgo_ts + 5019,
		Ftype1: int8('F'),
	},
	730: {
		Fword:  __ccgo_ts + 5026,
		Ftype1: int8('F'),
	},
	731: {
		Fword:  __ccgo_ts + 5033,
		Ftype1: int8('F'),
	},
	732: {
		Fword:  __ccgo_ts + 5040,
		Ftype1: int8('F'),
	},
	733: {
		Fword:  __ccgo_ts + 5047,
		Ftype1: int8('F'),
	},
	734: {
		Fword:  __ccgo_ts + 5054,
		Ftype1: int8('F'),
	},
	735: {
		Fword:  __ccgo_ts + 5061,
		Ftype1: int8('F'),
	},
	736: {
		Fword:  __ccgo_ts + 5068,
		Ftype1: int8('F'),
	},
	737: {
		Fword:  __ccgo_ts + 5075,
		Ftype1: int8('F'),
	},
	738: {
		Fword:  __ccgo_ts + 5082,
		Ftype1: int8('F'),
	},
	739: {
		Fword:  __ccgo_ts + 5089,
		Ftype1: int8('F'),
	},
	740: {
		Fword:  __ccgo_ts + 5096,
		Ftype1: int8('F'),
	},
	741: {
		Fword:  __ccgo_ts + 5103,
		Ftype1: int8('F'),
	},
	742: {
		Fword:  __ccgo_ts + 5110,
		Ftype1: int8('F'),
	},
	743: {
		Fword:  __ccgo_ts + 5117,
		Ftype1: int8('F'),
	},
	744: {
		Fword:  __ccgo_ts + 5124,
		Ftype1: int8('F'),
	},
	745: {
		Fword:  __ccgo_ts + 5131,
		Ftype1: int8('F'),
	},
	746: {
		Fword:  __ccgo_ts + 5138,
		Ftype1: int8('F'),
	},
	747: {
		Fword:  __ccgo_ts + 5145,
		Ftype1: int8('F'),
	},
	748: {
		Fword:  __ccgo_ts + 5152,
		Ftype1: int8('F'),
	},
	749: {
		Fword:  __ccgo_ts + 5159,
		Ftype1: int8('F'),
	},
	750: {
		Fword:  __ccgo_ts + 5166,
		Ftype1: int8('F'),
	},
	751: {
		Fword:  __ccgo_ts + 5173,
		Ftype1: int8('F'),
	},
	752: {
		Fword:  __ccgo_ts + 5180,
		Ftype1: int8('F'),
	},
	753: {
		Fword:  __ccgo_ts + 5187,
		Ftype1: int8('F'),
	},
	754: {
		Fword:  __ccgo_ts + 5194,
		Ftype1: int8('F'),
	},
	755: {
		Fword:  __ccgo_ts + 5201,
		Ftype1: int8('F'),
	},
	756: {
		Fword:  __ccgo_ts + 5208,
		Ftype1: int8('F'),
	},
	757: {
		Fword:  __ccgo_ts + 5215,
		Ftype1: int8('F'),
	},
	758: {
		Fword:  __ccgo_ts + 5222,
		Ftype1: int8('F'),
	},
	759: {
		Fword:  __ccgo_ts + 5229,
		Ftype1: int8('F'),
	},
	760: {
		Fword:  __ccgo_ts + 5236,
		Ftype1: int8('F'),
	},
	761: {
		Fword:  __ccgo_ts + 5243,
		Ftype1: int8('F'),
	},
	762: {
		Fword:  __ccgo_ts + 5250,
		Ftype1: int8('F'),
	},
	763: {
		Fword:  __ccgo_ts + 5257,
		Ftype1: int8('F'),
	},
	764: {
		Fword:  __ccgo_ts + 5264,
		Ftype1: int8('F'),
	},
	765: {
		Fword:  __ccgo_ts + 5271,
		Ftype1: int8('F'),
	},
	766: {
		Fword:  __ccgo_ts + 5278,
		Ftype1: int8('F'),
	},
	767: {
		Fword:  __ccgo_ts + 5285,
		Ftype1: int8('F'),
	},
	768: {
		Fword:  __ccgo_ts + 5292,
		Ftype1: int8('F'),
	},
	769: {
		Fword:  __ccgo_ts + 5299,
		Ftype1: int8('F'),
	},
	770: {
		Fword:  __ccgo_ts + 5306,
		Ftype1: int8('F'),
	},
	771: {
		Fword:  __ccgo_ts + 5313,
		Ftype1: int8('F'),
	},
	772: {
		Fword:  __ccgo_ts + 5320,
		Ftype1: int8('F'),
	},
	773: {
		Fword:  __ccgo_ts + 5327,
		Ftype1: int8('F'),
	},
	774: {
		Fword:  __ccgo_ts + 5334,
		Ftype1: int8('F'),
	},
	775: {
		Fword:  __ccgo_ts + 5341,
		Ftype1: int8('F'),
	},
	776: {
		Fword:  __ccgo_ts + 5348,
		Ftype1: int8('F'),
	},
	777: {
		Fword:  __ccgo_ts + 5355,
		Ftype1: int8('F'),
	},
	778: {
		Fword:  __ccgo_ts + 5362,
		Ftype1: int8('F'),
	},
	779: {
		Fword:  __ccgo_ts + 5369,
		Ftype1: int8('F'),
	},
	780: {
		Fword:  __ccgo_ts + 5376,
		Ftype1: int8('F'),
	},
	781: {
		Fword:  __ccgo_ts + 5383,
		Ftype1: int8('F'),
	},
	782: {
		Fword:  __ccgo_ts + 5390,
		Ftype1: int8('F'),
	},
	783: {
		Fword:  __ccgo_ts + 5397,
		Ftype1: int8('F'),
	},
	784: {
		Fword:  __ccgo_ts + 5404,
		Ftype1: int8('F'),
	},
	785: {
		Fword:  __ccgo_ts + 5411,
		Ftype1: int8('F'),
	},
	786: {
		Fword:  __ccgo_ts + 5418,
		Ftype1: int8('F'),
	},
	787: {
		Fword:  __ccgo_ts + 5425,
		Ftype1: int8('F'),
	},
	788: {
		Fword:  __ccgo_ts + 5432,
		Ftype1: int8('F'),
	},
	789: {
		Fword:  __ccgo_ts + 5439,
		Ftype1: int8('F'),
	},
	790: {
		Fword:  __ccgo_ts + 5446,
		Ftype1: int8('F'),
	},
	791: {
		Fword:  __ccgo_ts + 5453,
		Ftype1: int8('F'),
	},
	792: {
		Fword:  __ccgo_ts + 5460,
		Ftype1: int8('F'),
	},
	793: {
		Fword:  __ccgo_ts + 5467,
		Ftype1: int8('F'),
	},
	794: {
		Fword:  __ccgo_ts + 5474,
		Ftype1: int8('F'),
	},
	795: {
		Fword:  __ccgo_ts + 5481,
		Ftype1: int8('F'),
	},
	796: {
		Fword:  __ccgo_ts + 5488,
		Ftype1: int8('F'),
	},
	797: {
		Fword:  __ccgo_ts + 5495,
		Ftype1: int8('F'),
	},
	798: {
		Fword:  __ccgo_ts + 5502,
		Ftype1: int8('F'),
	},
	799: {
		Fword:  __ccgo_ts + 5509,
		Ftype1: int8('F'),
	},
	800: {
		Fword:  __ccgo_ts + 5516,
		Ftype1: int8('F'),
	},
	801: {
		Fword:  __ccgo_ts + 5523,
		Ftype1: int8('F'),
	},
	802: {
		Fword:  __ccgo_ts + 5530,
		Ftype1: int8('F'),
	},
	803: {
		Fword:  __ccgo_ts + 5537,
		Ftype1: int8('F'),
	},
	804: {
		Fword:  __ccgo_ts + 5544,
		Ftype1: int8('F'),
	},
	805: {
		Fword:  __ccgo_ts + 5551,
		Ftype1: int8('F'),
	},
	806: {
		Fword:  __ccgo_ts + 5556,
		Ftype1: int8('F'),
	},
	807: {
		Fword:  __ccgo_ts + 5563,
		Ftype1: int8('F'),
	},
	808: {
		Fword:  __ccgo_ts + 5570,
		Ftype1: int8('F'),
	},
	809: {
		Fword:  __ccgo_ts + 5577,
		Ftype1: int8('F'),
	},
	810: {
		Fword:  __ccgo_ts + 5584,
		Ftype1: int8('F'),
	},
	811: {
		Fword:  __ccgo_ts + 5591,
		Ftype1: int8('F'),
	},
	812: {
		Fword:  __ccgo_ts + 5598,
		Ftype1: int8('F'),
	},
	813: {
		Fword:  __ccgo_ts + 5605,
		Ftype1: int8('F'),
	},
	814: {
		Fword:  __ccgo_ts + 5612,
		Ftype1: int8('F'),
	},
	815: {
		Fword:  __ccgo_ts + 5618,
		Ftype1: int8('F'),
	},
	816: {
		Fword:  __ccgo_ts + 5625,
		Ftype1: int8('F'),
	},
	817: {
		Fword:  __ccgo_ts + 5632,
		Ftype1: int8('F'),
	},
	818: {
		Fword:  __ccgo_ts + 5639,
		Ftype1: int8('F'),
	},
	819: {
		Fword:  __ccgo_ts + 5646,
		Ftype1: int8('F'),
	},
	820: {
		Fword:  __ccgo_ts + 5653,
		Ftype1: int8('F'),
	},
	821: {
		Fword:  __ccgo_ts + 5660,
		Ftype1: int8('F'),
	},
	822: {
		Fword:  __ccgo_ts + 5667,
		Ftype1: int8('F'),
	},
	823: {
		Fword:  __ccgo_ts + 5673,
		Ftype1: int8('F'),
	},
	824: {
		Fword:  __ccgo_ts + 5680,
		Ftype1: int8('F'),
	},
	825: {
		Fword:  __ccgo_ts + 5687,
		Ftype1: int8('F'),
	},
	826: {
		Fword:  __ccgo_ts + 5694,
		Ftype1: int8('F'),
	},
	827: {
		Fword:  __ccgo_ts + 5701,
		Ftype1: int8('F'),
	},
	828: {
		Fword:  __ccgo_ts + 5708,
		Ftype1: int8('F'),
	},
	829: {
		Fword:  __ccgo_ts + 5715,
		Ftype1: int8('F'),
	},
	830: {
		Fword:  __ccgo_ts + 5722,
		Ftype1: int8('F'),
	},
	831: {
		Fword:  __ccgo_ts + 5729,
		Ftype1: int8('F'),
	},
	832: {
		Fword:  __ccgo_ts + 5736,
		Ftype1: int8('F'),
	},
	833: {
		Fword:  __ccgo_ts + 5743,
		Ftype1: int8('F'),
	},
	834: {
		Fword:  __ccgo_ts + 5750,
		Ftype1: int8('F'),
	},
	835: {
		Fword:  __ccgo_ts + 5757,
		Ftype1: int8('F'),
	},
	836: {
		Fword:  __ccgo_ts + 5764,
		Ftype1: int8('F'),
	},
	837: {
		Fword:  __ccgo_ts + 5771,
		Ftype1: int8('F'),
	},
	838: {
		Fword:  __ccgo_ts + 5778,
		Ftype1: int8('F'),
	},
	839: {
		Fword:  __ccgo_ts + 5785,
		Ftype1: int8('F'),
	},
	840: {
		Fword:  __ccgo_ts + 5792,
		Ftype1: int8('F'),
	},
	841: {
		Fword:  __ccgo_ts + 5799,
		Ftype1: int8('F'),
	},
	842: {
		Fword:  __ccgo_ts + 5806,
		Ftype1: int8('F'),
	},
	843: {
		Fword:  __ccgo_ts + 5813,
		Ftype1: int8('F'),
	},
	844: {
		Fword:  __ccgo_ts + 5820,
		Ftype1: int8('F'),
	},
	845: {
		Fword:  __ccgo_ts + 5827,
		Ftype1: int8('F'),
	},
	846: {
		Fword:  __ccgo_ts + 5834,
		Ftype1: int8('F'),
	},
	847: {
		Fword:  __ccgo_ts + 5841,
		Ftype1: int8('F'),
	},
	848: {
		Fword:  __ccgo_ts + 5846,
		Ftype1: int8('F'),
	},
	849: {
		Fword:  __ccgo_ts + 5853,
		Ftype1: int8('F'),
	},
	850: {
		Fword:  __ccgo_ts + 5860,
		Ftype1: int8('F'),
	},
	851: {
		Fword:  __ccgo_ts + 5867,
		Ftype1: int8('F'),
	},
	852: {
		Fword:  __ccgo_ts + 5874,
		Ftype1: int8('F'),
	},
	853: {
		Fword:  __ccgo_ts + 5881,
		Ftype1: int8('F'),
	},
	854: {
		Fword:  __ccgo_ts + 5888,
		Ftype1: int8('F'),
	},
	855: {
		Fword:  __ccgo_ts + 5895,
		Ftype1: int8('F'),
	},
	856: {
		Fword:  __ccgo_ts + 5902,
		Ftype1: int8('F'),
	},
	857: {
		Fword:  __ccgo_ts + 5908,
		Ftype1: int8('F'),
	},
	858: {
		Fword:  __ccgo_ts + 5915,
		Ftype1: int8('F'),
	},
	859: {
		Fword:  __ccgo_ts + 5922,
		Ftype1: int8('F'),
	},
	860: {
		Fword:  __ccgo_ts + 5929,
		Ftype1: int8('F'),
	},
	861: {
		Fword:  __ccgo_ts + 5936,
		Ftype1: int8('F'),
	},
	862: {
		Fword:  __ccgo_ts + 5943,
		Ftype1: int8('F'),
	},
	863: {
		Fword:  __ccgo_ts + 5950,
		Ftype1: int8('F'),
	},
	864: {
		Fword:  __ccgo_ts + 5957,
		Ftype1: int8('F'),
	},
	865: {
		Fword:  __ccgo_ts + 5963,
		Ftype1: int8('F'),
	},
	866: {
		Fword:  __ccgo_ts + 5970,
		Ftype1: int8('F'),
	},
	867: {
		Fword:  __ccgo_ts + 5977,
		Ftype1: int8('F'),
	},
	868: {
		Fword:  __ccgo_ts + 5984,
		Ftype1: int8('F'),
	},
	869: {
		Fword:  __ccgo_ts + 5991,
		Ftype1: int8('F'),
	},
	870: {
		Fword:  __ccgo_ts + 5998,
		Ftype1: int8('F'),
	},
	871: {
		Fword:  __ccgo_ts + 6005,
		Ftype1: int8('F'),
	},
	872: {
		Fword:  __ccgo_ts + 6012,
		Ftype1: int8('F'),
	},
	873: {
		Fword:  __ccgo_ts + 6019,
		Ftype1: int8('F'),
	},
	874: {
		Fword:  __ccgo_ts + 6026,
		Ftype1: int8('F'),
	},
	875: {
		Fword:  __ccgo_ts + 6033,
		Ftype1: int8('F'),
	},
	876: {
		Fword:  __ccgo_ts + 6040,
		Ftype1: int8('F'),
	},
	877: {
		Fword:  __ccgo_ts + 6047,
		Ftype1: int8('F'),
	},
	878: {
		Fword:  __ccgo_ts + 6052,
		Ftype1: int8('F'),
	},
	879: {
		Fword:  __ccgo_ts + 6059,
		Ftype1: int8('F'),
	},
	880: {
		Fword:  __ccgo_ts + 6066,
		Ftype1: int8('F'),
	},
	881: {
		Fword:  __ccgo_ts + 6073,
		Ftype1: int8('F'),
	},
	882: {
		Fword:  __ccgo_ts + 6080,
		Ftype1: int8('F'),
	},
	883: {
		Fword:  __ccgo_ts + 6087,
		Ftype1: int8('F'),
	},
	884: {
		Fword:  __ccgo_ts + 6094,
		Ftype1: int8('F'),
	},
	885: {
		Fword:  __ccgo_ts + 6101,
		Ftype1: int8('F'),
	},
	886: {
		Fword:  __ccgo_ts + 6108,
		Ftype1: int8('F'),
	},
	887: {
		Fword:  __ccgo_ts + 6114,
		Ftype1: int8('F'),
	},
	888: {
		Fword:  __ccgo_ts + 6121,
		Ftype1: int8('F'),
	},
	889: {
		Fword:  __ccgo_ts + 6128,
		Ftype1: int8('F'),
	},
	890: {
		Fword:  __ccgo_ts + 6135,
		Ftype1: int8('F'),
	},
	891: {
		Fword:  __ccgo_ts + 6142,
		Ftype1: int8('F'),
	},
	892: {
		Fword:  __ccgo_ts + 6149,
		Ftype1: int8('F'),
	},
	893: {
		Fword:  __ccgo_ts + 6156,
		Ftype1: int8('F'),
	},
	894: {
		Fword:  __ccgo_ts + 6163,
		Ftype1: int8('F'),
	},
	895: {
		Fword:  __ccgo_ts + 6169,
		Ftype1: int8('F'),
	},
	896: {
		Fword:  __ccgo_ts + 6176,
		Ftype1: int8('F'),
	},
	897: {
		Fword:  __ccgo_ts + 6183,
		Ftype1: int8('F'),
	},
	898: {
		Fword:  __ccgo_ts + 6190,
		Ftype1: int8('F'),
	},
	899: {
		Fword:  __ccgo_ts + 6197,
		Ftype1: int8('F'),
	},
	900: {
		Fword:  __ccgo_ts + 6204,
		Ftype1: int8('F'),
	},
	901: {
		Fword:  __ccgo_ts + 6211,
		Ftype1: int8('F'),
	},
	902: {
		Fword:  __ccgo_ts + 6218,
		Ftype1: int8('F'),
	},
	903: {
		Fword:  __ccgo_ts + 6225,
		Ftype1: int8('F'),
	},
	904: {
		Fword:  __ccgo_ts + 6232,
		Ftype1: int8('F'),
	},
	905: {
		Fword:  __ccgo_ts + 6239,
		Ftype1: int8('F'),
	},
	906: {
		Fword:  __ccgo_ts + 6246,
		Ftype1: int8('F'),
	},
	907: {
		Fword:  __ccgo_ts + 6253,
		Ftype1: int8('F'),
	},
	908: {
		Fword:  __ccgo_ts + 6260,
		Ftype1: int8('F'),
	},
	909: {
		Fword:  __ccgo_ts + 6267,
		Ftype1: int8('F'),
	},
	910: {
		Fword:  __ccgo_ts + 6272,
		Ftype1: int8('F'),
	},
	911: {
		Fword:  __ccgo_ts + 6279,
		Ftype1: int8('F'),
	},
	912: {
		Fword:  __ccgo_ts + 6286,
		Ftype1: int8('F'),
	},
	913: {
		Fword:  __ccgo_ts + 6293,
		Ftype1: int8('F'),
	},
	914: {
		Fword:  __ccgo_ts + 6300,
		Ftype1: int8('F'),
	},
	915: {
		Fword:  __ccgo_ts + 6307,
		Ftype1: int8('F'),
	},
	916: {
		Fword:  __ccgo_ts + 6314,
		Ftype1: int8('F'),
	},
	917: {
		Fword:  __ccgo_ts + 6321,
		Ftype1: int8('F'),
	},
	918: {
		Fword:  __ccgo_ts + 6328,
		Ftype1: int8('F'),
	},
	919: {
		Fword:  __ccgo_ts + 6334,
		Ftype1: int8('F'),
	},
	920: {
		Fword:  __ccgo_ts + 6341,
		Ftype1: int8('F'),
	},
	921: {
		Fword:  __ccgo_ts + 6348,
		Ftype1: int8('F'),
	},
	922: {
		Fword:  __ccgo_ts + 6355,
		Ftype1: int8('F'),
	},
	923: {
		Fword:  __ccgo_ts + 6362,
		Ftype1: int8('F'),
	},
	924: {
		Fword:  __ccgo_ts + 6369,
		Ftype1: int8('F'),
	},
	925: {
		Fword:  __ccgo_ts + 6376,
		Ftype1: int8('F'),
	},
	926: {
		Fword:  __ccgo_ts + 6383,
		Ftype1: int8('F'),
	},
	927: {
		Fword:  __ccgo_ts + 6389,
		Ftype1: int8('F'),
	},
	928: {
		Fword:  __ccgo_ts + 6396,
		Ftype1: int8('F'),
	},
	929: {
		Fword:  __ccgo_ts + 6403,
		Ftype1: int8('F'),
	},
	930: {
		Fword:  __ccgo_ts + 6410,
		Ftype1: int8('F'),
	},
	931: {
		Fword:  __ccgo_ts + 6417,
		Ftype1: int8('F'),
	},
	932: {
		Fword:  __ccgo_ts + 6424,
		Ftype1: int8('F'),
	},
	933: {
		Fword:  __ccgo_ts + 6431,
		Ftype1: int8('F'),
	},
	934: {
		Fword:  __ccgo_ts + 6438,
		Ftype1: int8('F'),
	},
	935: {
		Fword:  __ccgo_ts + 6445,
		Ftype1: int8('F'),
	},
	936: {
		Fword:  __ccgo_ts + 6452,
		Ftype1: int8('F'),
	},
	937: {
		Fword:  __ccgo_ts + 6459,
		Ftype1: int8('F'),
	},
	938: {
		Fword:  __ccgo_ts + 6466,
		Ftype1: int8('F'),
	},
	939: {
		Fword:  __ccgo_ts + 6470,
		Ftype1: int8('F'),
	},
	940: {
		Fword:  __ccgo_ts + 6477,
		Ftype1: int8('F'),
	},
	941: {
		Fword:  __ccgo_ts + 6484,
		Ftype1: int8('F'),
	},
	942: {
		Fword:  __ccgo_ts + 6491,
		Ftype1: int8('F'),
	},
	943: {
		Fword:  __ccgo_ts + 6498,
		Ftype1: int8('F'),
	},
	944: {
		Fword:  __ccgo_ts + 6505,
		Ftype1: int8('F'),
	},
	945: {
		Fword:  __ccgo_ts + 6512,
		Ftype1: int8('F'),
	},
	946: {
		Fword:  __ccgo_ts + 6519,
		Ftype1: int8('F'),
	},
	947: {
		Fword:  __ccgo_ts + 6526,
		Ftype1: int8('F'),
	},
	948: {
		Fword:  __ccgo_ts + 6533,
		Ftype1: int8('F'),
	},
	949: {
		Fword:  __ccgo_ts + 6540,
		Ftype1: int8('F'),
	},
	950: {
		Fword:  __ccgo_ts + 6546,
		Ftype1: int8('F'),
	},
	951: {
		Fword:  __ccgo_ts + 6553,
		Ftype1: int8('F'),
	},
	952: {
		Fword:  __ccgo_ts + 6560,
		Ftype1: int8('F'),
	},
	953: {
		Fword:  __ccgo_ts + 6567,
		Ftype1: int8('F'),
	},
	954: {
		Fword:  __ccgo_ts + 6574,
		Ftype1: int8('F'),
	},
	955: {
		Fword:  __ccgo_ts + 6581,
		Ftype1: int8('F'),
	},
	956: {
		Fword:  __ccgo_ts + 6588,
		Ftype1: int8('F'),
	},
	957: {
		Fword:  __ccgo_ts + 6595,
		Ftype1: int8('F'),
	},
	958: {
		Fword:  __ccgo_ts + 6602,
		Ftype1: int8('F'),
	},
	959: {
		Fword:  __ccgo_ts + 6609,
		Ftype1: int8('F'),
	},
	960: {
		Fword:  __ccgo_ts + 6616,
		Ftype1: int8('F'),
	},
	961: {
		Fword:  __ccgo_ts + 6623,
		Ftype1: int8('F'),
	},
	962: {
		Fword:  __ccgo_ts + 6630,
		Ftype1: int8('F'),
	},
	963: {
		Fword:  __ccgo_ts + 6637,
		Ftype1: int8('F'),
	},
	964: {
		Fword:  __ccgo_ts + 6644,
		Ftype1: int8('F'),
	},
	965: {
		Fword:  __ccgo_ts + 6651,
		Ftype1: int8('F'),
	},
	966: {
		Fword:  __ccgo_ts + 6658,
		Ftype1: int8('F'),
	},
	967: {
		Fword:  __ccgo_ts + 6665,
		Ftype1: int8('F'),
	},
	968: {
		Fword:  __ccgo_ts + 6672,
		Ftype1: int8('F'),
	},
	969: {
		Fword:  __ccgo_ts + 6679,
		Ftype1: int8('F'),
	},
	970: {
		Fword:  __ccgo_ts + 6686,
		Ftype1: int8('F'),
	},
	971: {
		Fword:  __ccgo_ts + 6693,
		Ftype1: int8('F'),
	},
	972: {
		Fword:  __ccgo_ts + 6700,
		Ftype1: int8('F'),
	},
	973: {
		Fword:  __ccgo_ts + 6707,
		Ftype1: int8('F'),
	},
	974: {
		Fword:  __ccgo_ts + 6714,
		Ftype1: int8('F'),
	},
	975: {
		Fword:  __ccgo_ts + 6721,
		Ftype1: int8('F'),
	},
	976: {
		Fword:  __ccgo_ts + 6728,
		Ftype1: int8('F'),
	},
	977: {
		Fword:  __ccgo_ts + 6735,
		Ftype1: int8('F'),
	},
	978: {
		Fword:  __ccgo_ts + 6742,
		Ftype1: int8('F'),
	},
	979: {
		Fword:  __ccgo_ts + 6749,
		Ftype1: int8('F'),
	},
	980: {
		Fword:  __ccgo_ts + 6756,
		Ftype1: int8('F'),
	},
	981: {
		Fword:  __ccgo_ts + 6763,
		Ftype1: int8('F'),
	},
	982: {
		Fword:  __ccgo_ts + 6770,
		Ftype1: int8('F'),
	},
	983: {
		Fword:  __ccgo_ts + 6777,
		Ftype1: int8('F'),
	},
	984: {
		Fword:  __ccgo_ts + 6784,
		Ftype1: int8('F'),
	},
	985: {
		Fword:  __ccgo_ts + 6791,
		Ftype1: int8('F'),
	},
	986: {
		Fword:  __ccgo_ts + 6798,
		Ftype1: int8('F'),
	},
	987: {
		Fword:  __ccgo_ts + 6805,
		Ftype1: int8('F'),
	},
	988: {
		Fword:  __ccgo_ts + 6812,
		Ftype1: int8('F'),
	},
	989: {
		Fword:  __ccgo_ts + 6819,
		Ftype1: int8('F'),
	},
	990: {
		Fword:  __ccgo_ts + 6826,
		Ftype1: int8('F'),
	},
	991: {
		Fword:  __ccgo_ts + 6833,
		Ftype1: int8('F'),
	},
	992: {
		Fword:  __ccgo_ts + 6840,
		Ftype1: int8('F'),
	},
	993: {
		Fword:  __ccgo_ts + 6847,
		Ftype1: int8('F'),
	},
	994: {
		Fword:  __ccgo_ts + 6854,
		Ftype1: int8('F'),
	},
	995: {
		Fword:  __ccgo_ts + 6861,
		Ftype1: int8('F'),
	},
	996: {
		Fword:  __ccgo_ts + 6868,
		Ftype1: int8('F'),
	},
	997: {
		Fword:  __ccgo_ts + 6875,
		Ftype1: int8('F'),
	},
	998: {
		Fword:  __ccgo_ts + 6882,
		Ftype1: int8('F'),
	},
	999: {
		Fword:  __ccgo_ts + 6889,
		Ftype1: int8('F'),
	},
	1000: {
		Fword:  __ccgo_ts + 6896,
		Ftype1: int8('F'),
	},
	1001: {
		Fword:  __ccgo_ts + 6903,
		Ftype1: int8('F'),
	},
	1002: {
		Fword:  __ccgo_ts + 6910,
		Ftype1: int8('F'),
	},
	1003: {
		Fword:  __ccgo_ts + 6916,
		Ftype1: int8('F'),
	},
	1004: {
		Fword:  __ccgo_ts + 6923,
		Ftype1: int8('F'),
	},
	1005: {
		Fword:  __ccgo_ts + 6930,
		Ftype1: int8('F'),
	},
	1006: {
		Fword:  __ccgo_ts + 6937,
		Ftype1: int8('F'),
	},
	1007: {
		Fword:  __ccgo_ts + 6944,
		Ftype1: int8('F'),
	},
	1008: {
		Fword:  __ccgo_ts + 6951,
		Ftype1: int8('F'),
	},
	1009: {
		Fword:  __ccgo_ts + 6958,
		Ftype1: int8('F'),
	},
	1010: {
		Fword:  __ccgo_ts + 6965,
		Ftype1: int8('F'),
	},
	1011: {
		Fword:  __ccgo_ts + 6972,
		Ftype1: int8('F'),
	},
	1012: {
		Fword:  __ccgo_ts + 6979,
		Ftype1: int8('F'),
	},
	1013: {
		Fword:  __ccgo_ts + 6986,
		Ftype1: int8('F'),
	},
	1014: {
		Fword:  __ccgo_ts + 6993,
		Ftype1: int8('F'),
	},
	1015: {
		Fword:  __ccgo_ts + 7000,
		Ftype1: int8('F'),
	},
	1016: {
		Fword:  __ccgo_ts + 7007,
		Ftype1: int8('F'),
	},
	1017: {
		Fword:  __ccgo_ts + 7014,
		Ftype1: int8('F'),
	},
	1018: {
		Fword:  __ccgo_ts + 7020,
		Ftype1: int8('F'),
	},
	1019: {
		Fword:  __ccgo_ts + 7027,
		Ftype1: int8('F'),
	},
	1020: {
		Fword:  __ccgo_ts + 7034,
		Ftype1: int8('F'),
	},
	1021: {
		Fword:  __ccgo_ts + 7041,
		Ftype1: int8('F'),
	},
	1022: {
		Fword:  __ccgo_ts + 7048,
		Ftype1: int8('F'),
	},
	1023: {
		Fword:  __ccgo_ts + 7055,
		Ftype1: int8('F'),
	},
	1024: {
		Fword:  __ccgo_ts + 7062,
		Ftype1: int8('F'),
	},
	1025: {
		Fword:  __ccgo_ts + 7069,
		Ftype1: int8('F'),
	},
	1026: {
		Fword:  __ccgo_ts + 7076,
		Ftype1: int8('F'),
	},
	1027: {
		Fword:  __ccgo_ts + 7083,
		Ftype1: int8('F'),
	},
	1028: {
		Fword:  __ccgo_ts + 7090,
		Ftype1: int8('F'),
	},
	1029: {
		Fword:  __ccgo_ts + 7097,
		Ftype1: int8('F'),
	},
	1030: {
		Fword:  __ccgo_ts + 7104,
		Ftype1: int8('F'),
	},
	1031: {
		Fword:  __ccgo_ts + 7111,
		Ftype1: int8('F'),
	},
	1032: {
		Fword:  __ccgo_ts + 7118,
		Ftype1: int8('F'),
	},
	1033: {
		Fword:  __ccgo_ts + 7125,
		Ftype1: int8('F'),
	},
	1034: {
		Fword:  __ccgo_ts + 7132,
		Ftype1: int8('F'),
	},
	1035: {
		Fword:  __ccgo_ts + 7139,
		Ftype1: int8('F'),
	},
	1036: {
		Fword:  __ccgo_ts + 7146,
		Ftype1: int8('F'),
	},
	1037: {
		Fword:  __ccgo_ts + 7153,
		Ftype1: int8('F'),
	},
	1038: {
		Fword:  __ccgo_ts + 7160,
		Ftype1: int8('F'),
	},
	1039: {
		Fword:  __ccgo_ts + 7167,
		Ftype1: int8('F'),
	},
	1040: {
		Fword:  __ccgo_ts + 7174,
		Ftype1: int8('F'),
	},
	1041: {
		Fword:  __ccgo_ts + 7181,
		Ftype1: int8('F'),
	},
	1042: {
		Fword:  __ccgo_ts + 7188,
		Ftype1: int8('F'),
	},
	1043: {
		Fword:  __ccgo_ts + 7195,
		Ftype1: int8('F'),
	},
	1044: {
		Fword:  __ccgo_ts + 7202,
		Ftype1: int8('F'),
	},
	1045: {
		Fword:  __ccgo_ts + 7209,
		Ftype1: int8('F'),
	},
	1046: {
		Fword:  __ccgo_ts + 7216,
		Ftype1: int8('F'),
	},
	1047: {
		Fword:  __ccgo_ts + 7223,
		Ftype1: int8('F'),
	},
	1048: {
		Fword:  __ccgo_ts + 7230,
		Ftype1: int8('F'),
	},
	1049: {
		Fword:  __ccgo_ts + 7237,
		Ftype1: int8('F'),
	},
	1050: {
		Fword:  __ccgo_ts + 7243,
		Ftype1: int8('F'),
	},
	1051: {
		Fword:  __ccgo_ts + 7250,
		Ftype1: int8('F'),
	},
	1052: {
		Fword:  __ccgo_ts + 7257,
		Ftype1: int8('F'),
	},
	1053: {
		Fword:  __ccgo_ts + 7264,
		Ftype1: int8('F'),
	},
	1054: {
		Fword:  __ccgo_ts + 7271,
		Ftype1: int8('F'),
	},
	1055: {
		Fword:  __ccgo_ts + 7278,
		Ftype1: int8('F'),
	},
	1056: {
		Fword:  __ccgo_ts + 7285,
		Ftype1: int8('F'),
	},
	1057: {
		Fword:  __ccgo_ts + 7292,
		Ftype1: int8('F'),
	},
	1058: {
		Fword:  __ccgo_ts + 7299,
		Ftype1: int8('F'),
	},
	1059: {
		Fword:  __ccgo_ts + 7306,
		Ftype1: int8('F'),
	},
	1060: {
		Fword:  __ccgo_ts + 7313,
		Ftype1: int8('F'),
	},
	1061: {
		Fword:  __ccgo_ts + 7320,
		Ftype1: int8('F'),
	},
	1062: {
		Fword:  __ccgo_ts + 7327,
		Ftype1: int8('F'),
	},
	1063: {
		Fword:  __ccgo_ts + 7334,
		Ftype1: int8('F'),
	},
	1064: {
		Fword:  __ccgo_ts + 7341,
		Ftype1: int8('F'),
	},
	1065: {
		Fword:  __ccgo_ts + 7348,
		Ftype1: int8('F'),
	},
	1066: {
		Fword:  __ccgo_ts + 7355,
		Ftype1: int8('F'),
	},
	1067: {
		Fword:  __ccgo_ts + 7362,
		Ftype1: int8('F'),
	},
	1068: {
		Fword:  __ccgo_ts + 7369,
		Ftype1: int8('F'),
	},
	1069: {
		Fword:  __ccgo_ts + 7376,
		Ftype1: int8('F'),
	},
	1070: {
		Fword:  __ccgo_ts + 7383,
		Ftype1: int8('F'),
	},
	1071: {
		Fword:  __ccgo_ts + 7390,
		Ftype1: int8('F'),
	},
	1072: {
		Fword:  __ccgo_ts + 7397,
		Ftype1: int8('F'),
	},
	1073: {
		Fword:  __ccgo_ts + 7404,
		Ftype1: int8('F'),
	},
	1074: {
		Fword:  __ccgo_ts + 7411,
		Ftype1: int8('F'),
	},
	1075: {
		Fword:  __ccgo_ts + 7418,
		Ftype1: int8('F'),
	},
	1076: {
		Fword:  __ccgo_ts + 7425,
		Ftype1: int8('F'),
	},
	1077: {
		Fword:  __ccgo_ts + 7432,
		Ftype1: int8('F'),
	},
	1078: {
		Fword:  __ccgo_ts + 7439,
		Ftype1: int8('F'),
	},
	1079: {
		Fword:  __ccgo_ts + 7446,
		Ftype1: int8('F'),
	},
	1080: {
		Fword:  __ccgo_ts + 7453,
		Ftype1: int8('F'),
	},
	1081: {
		Fword:  __ccgo_ts + 7460,
		Ftype1: int8('F'),
	},
	1082: {
		Fword:  __ccgo_ts + 7467,
		Ftype1: int8('F'),
	},
	1083: {
		Fword:  __ccgo_ts + 7474,
		Ftype1: int8('F'),
	},
	1084: {
		Fword:  __ccgo_ts + 7481,
		Ftype1: int8('F'),
	},
	1085: {
		Fword:  __ccgo_ts + 7488,
		Ftype1: int8('F'),
	},
	1086: {
		Fword:  __ccgo_ts + 7495,
		Ftype1: int8('F'),
	},
	1087: {
		Fword:  __ccgo_ts + 7502,
		Ftype1: int8('F'),
	},
	1088: {
		Fword:  __ccgo_ts + 7509,
		Ftype1: int8('F'),
	},
	1089: {
		Fword:  __ccgo_ts + 7516,
		Ftype1: int8('F'),
	},
	1090: {
		Fword:  __ccgo_ts + 7523,
		Ftype1: int8('F'),
	},
	1091: {
		Fword:  __ccgo_ts + 7530,
		Ftype1: int8('F'),
	},
	1092: {
		Fword:  __ccgo_ts + 7537,
		Ftype1: int8('F'),
	},
	1093: {
		Fword:  __ccgo_ts + 7544,
		Ftype1: int8('F'),
	},
	1094: {
		Fword:  __ccgo_ts + 7551,
		Ftype1: int8('F'),
	},
	1095: {
		Fword:  __ccgo_ts + 7558,
		Ftype1: int8('F'),
	},
	1096: {
		Fword:  __ccgo_ts + 7565,
		Ftype1: int8('F'),
	},
	1097: {
		Fword:  __ccgo_ts + 7572,
		Ftype1: int8('F'),
	},
	1098: {
		Fword:  __ccgo_ts + 7579,
		Ftype1: int8('F'),
	},
	1099: {
		Fword:  __ccgo_ts + 7586,
		Ftype1: int8('F'),
	},
	1100: {
		Fword:  __ccgo_ts + 7593,
		Ftype1: int8('F'),
	},
	1101: {
		Fword:  __ccgo_ts + 7600,
		Ftype1: int8('F'),
	},
	1102: {
		Fword:  __ccgo_ts + 7607,
		Ftype1: int8('F'),
	},
	1103: {
		Fword:  __ccgo_ts + 7614,
		Ftype1: int8('F'),
	},
	1104: {
		Fword:  __ccgo_ts + 7621,
		Ftype1: int8('F'),
	},
	1105: {
		Fword:  __ccgo_ts + 7628,
		Ftype1: int8('F'),
	},
	1106: {
		Fword:  __ccgo_ts + 7635,
		Ftype1: int8('F'),
	},
	1107: {
		Fword:  __ccgo_ts + 7642,
		Ftype1: int8('F'),
	},
	1108: {
		Fword:  __ccgo_ts + 7649,
		Ftype1: int8('F'),
	},
	1109: {
		Fword:  __ccgo_ts + 7656,
		Ftype1: int8('F'),
	},
	1110: {
		Fword:  __ccgo_ts + 7663,
		Ftype1: int8('F'),
	},
	1111: {
		Fword:  __ccgo_ts + 7670,
		Ftype1: int8('F'),
	},
	1112: {
		Fword:  __ccgo_ts + 7677,
		Ftype1: int8('F'),
	},
	1113: {
		Fword:  __ccgo_ts + 7684,
		Ftype1: int8('F'),
	},
	1114: {
		Fword:  __ccgo_ts + 7691,
		Ftype1: int8('F'),
	},
	1115: {
		Fword:  __ccgo_ts + 7698,
		Ftype1: int8('F'),
	},
	1116: {
		Fword:  __ccgo_ts + 7705,
		Ftype1: int8('F'),
	},
	1117: {
		Fword:  __ccgo_ts + 7712,
		Ftype1: int8('F'),
	},
	1118: {
		Fword:  __ccgo_ts + 7719,
		Ftype1: int8('F'),
	},
	1119: {
		Fword:  __ccgo_ts + 7726,
		Ftype1: int8('F'),
	},
	1120: {
		Fword:  __ccgo_ts + 7731,
		Ftype1: int8('F'),
	},
	1121: {
		Fword:  __ccgo_ts + 7738,
		Ftype1: int8('F'),
	},
	1122: {
		Fword:  __ccgo_ts + 7745,
		Ftype1: int8('F'),
	},
	1123: {
		Fword:  __ccgo_ts + 7752,
		Ftype1: int8('F'),
	},
	1124: {
		Fword:  __ccgo_ts + 7759,
		Ftype1: int8('F'),
	},
	1125: {
		Fword:  __ccgo_ts + 7766,
		Ftype1: int8('F'),
	},
	1126: {
		Fword:  __ccgo_ts + 7773,
		Ftype1: int8('F'),
	},
	1127: {
		Fword:  __ccgo_ts + 7779,
		Ftype1: int8('F'),
	},
	1128: {
		Fword:  __ccgo_ts + 7786,
		Ftype1: int8('F'),
	},
	1129: {
		Fword:  __ccgo_ts + 7793,
		Ftype1: int8('F'),
	},
	1130: {
		Fword:  __ccgo_ts + 7800,
		Ftype1: int8('F'),
	},
	1131: {
		Fword:  __ccgo_ts + 7807,
		Ftype1: int8('F'),
	},
	1132: {
		Fword:  __ccgo_ts + 7814,
		Ftype1: int8('F'),
	},
	1133: {
		Fword:  __ccgo_ts + 7821,
		Ftype1: int8('F'),
	},
	1134: {
		Fword:  __ccgo_ts + 7828,
		Ftype1: int8('F'),
	},
	1135: {
		Fword:  __ccgo_ts + 7835,
		Ftype1: int8('F'),
	},
	1136: {
		Fword:  __ccgo_ts + 7842,
		Ftype1: int8('F'),
	},
	1137: {
		Fword:  __ccgo_ts + 7848,
		Ftype1: int8('F'),
	},
	1138: {
		Fword:  __ccgo_ts + 7855,
		Ftype1: int8('F'),
	},
	1139: {
		Fword:  __ccgo_ts + 7862,
		Ftype1: int8('F'),
	},
	1140: {
		Fword:  __ccgo_ts + 7869,
		Ftype1: int8('F'),
	},
	1141: {
		Fword:  __ccgo_ts + 7876,
		Ftype1: int8('F'),
	},
	1142: {
		Fword:  __ccgo_ts + 7883,
		Ftype1: int8('F'),
	},
	1143: {
		Fword:  __ccgo_ts + 7890,
		Ftype1: int8('F'),
	},
	1144: {
		Fword:  __ccgo_ts + 7897,
		Ftype1: int8('F'),
	},
	1145: {
		Fword:  __ccgo_ts + 7904,
		Ftype1: int8('F'),
	},
	1146: {
		Fword:  __ccgo_ts + 7911,
		Ftype1: int8('F'),
	},
	1147: {
		Fword:  __ccgo_ts + 7918,
		Ftype1: int8('F'),
	},
	1148: {
		Fword:  __ccgo_ts + 7925,
		Ftype1: int8('F'),
	},
	1149: {
		Fword:  __ccgo_ts + 7932,
		Ftype1: int8('F'),
	},
	1150: {
		Fword:  __ccgo_ts + 7939,
		Ftype1: int8('F'),
	},
	1151: {
		Fword:  __ccgo_ts + 7946,
		Ftype1: int8('F'),
	},
	1152: {
		Fword:  __ccgo_ts + 7953,
		Ftype1: int8('F'),
	},
	1153: {
		Fword:  __ccgo_ts + 7960,
		Ftype1: int8('F'),
	},
	1154: {
		Fword:  __ccgo_ts + 7967,
		Ftype1: int8('F'),
	},
	1155: {
		Fword:  __ccgo_ts + 7974,
		Ftype1: int8('F'),
	},
	1156: {
		Fword:  __ccgo_ts + 7981,
		Ftype1: int8('F'),
	},
	1157: {
		Fword:  __ccgo_ts + 7986,
		Ftype1: int8('F'),
	},
	1158: {
		Fword:  __ccgo_ts + 7993,
		Ftype1: int8('F'),
	},
	1159: {
		Fword:  __ccgo_ts + 8000,
		Ftype1: int8('F'),
	},
	1160: {
		Fword:  __ccgo_ts + 8007,
		Ftype1: int8('F'),
	},
	1161: {
		Fword:  __ccgo_ts + 8014,
		Ftype1: int8('F'),
	},
	1162: {
		Fword:  __ccgo_ts + 8021,
		Ftype1: int8('F'),
	},
	1163: {
		Fword:  __ccgo_ts + 8028,
		Ftype1: int8('F'),
	},
	1164: {
		Fword:  __ccgo_ts + 8034,
		Ftype1: int8('F'),
	},
	1165: {
		Fword:  __ccgo_ts + 8041,
		Ftype1: int8('F'),
	},
	1166: {
		Fword:  __ccgo_ts + 8048,
		Ftype1: int8('F'),
	},
	1167: {
		Fword:  __ccgo_ts + 8055,
		Ftype1: int8('F'),
	},
	1168: {
		Fword:  __ccgo_ts + 8062,
		Ftype1: int8('F'),
	},
	1169: {
		Fword:  __ccgo_ts + 8069,
		Ftype1: int8('F'),
	},
	1170: {
		Fword:  __ccgo_ts + 8076,
		Ftype1: int8('F'),
	},
	1171: {
		Fword:  __ccgo_ts + 8083,
		Ftype1: int8('F'),
	},
	1172: {
		Fword:  __ccgo_ts + 8090,
		Ftype1: int8('F'),
	},
	1173: {
		Fword:  __ccgo_ts + 8097,
		Ftype1: int8('F'),
	},
	1174: {
		Fword:  __ccgo_ts + 8103,
		Ftype1: int8('F'),
	},
	1175: {
		Fword:  __ccgo_ts + 8110,
		Ftype1: int8('F'),
	},
	1176: {
		Fword:  __ccgo_ts + 8117,
		Ftype1: int8('F'),
	},
	1177: {
		Fword:  __ccgo_ts + 8124,
		Ftype1: int8('F'),
	},
	1178: {
		Fword:  __ccgo_ts + 8131,
		Ftype1: int8('F'),
	},
	1179: {
		Fword:  __ccgo_ts + 8138,
		Ftype1: int8('F'),
	},
	1180: {
		Fword:  __ccgo_ts + 8145,
		Ftype1: int8('F'),
	},
	1181: {
		Fword:  __ccgo_ts + 8152,
		Ftype1: int8('F'),
	},
	1182: {
		Fword:  __ccgo_ts + 8159,
		Ftype1: int8('F'),
	},
	1183: {
		Fword:  __ccgo_ts + 8164,
		Ftype1: int8('F'),
	},
	1184: {
		Fword:  __ccgo_ts + 8171,
		Ftype1: int8('F'),
	},
	1185: {
		Fword:  __ccgo_ts + 8178,
		Ftype1: int8('F'),
	},
	1186: {
		Fword:  __ccgo_ts + 8185,
		Ftype1: int8('F'),
	},
	1187: {
		Fword:  __ccgo_ts + 8192,
		Ftype1: int8('F'),
	},
	1188: {
		Fword:  __ccgo_ts + 8199,
		Ftype1: int8('F'),
	},
	1189: {
		Fword:  __ccgo_ts + 8206,
		Ftype1: int8('F'),
	},
	1190: {
		Fword:  __ccgo_ts + 8212,
		Ftype1: int8('F'),
	},
	1191: {
		Fword:  __ccgo_ts + 8219,
		Ftype1: int8('F'),
	},
	1192: {
		Fword:  __ccgo_ts + 8226,
		Ftype1: int8('F'),
	},
	1193: {
		Fword:  __ccgo_ts + 8233,
		Ftype1: int8('F'),
	},
	1194: {
		Fword:  __ccgo_ts + 8240,
		Ftype1: int8('F'),
	},
	1195: {
		Fword:  __ccgo_ts + 8247,
		Ftype1: int8('F'),
	},
	1196: {
		Fword:  __ccgo_ts + 8254,
		Ftype1: int8('F'),
	},
	1197: {
		Fword:  __ccgo_ts + 8261,
		Ftype1: int8('F'),
	},
	1198: {
		Fword:  __ccgo_ts + 8268,
		Ftype1: int8('F'),
	},
	1199: {
		Fword:  __ccgo_ts + 8275,
		Ftype1: int8('F'),
	},
	1200: {
		Fword:  __ccgo_ts + 8281,
		Ftype1: int8('F'),
	},
	1201: {
		Fword:  __ccgo_ts + 8288,
		Ftype1: int8('F'),
	},
	1202: {
		Fword:  __ccgo_ts + 8295,
		Ftype1: int8('F'),
	},
	1203: {
		Fword:  __ccgo_ts + 8302,
		Ftype1: int8('F'),
	},
	1204: {
		Fword:  __ccgo_ts + 8309,
		Ftype1: int8('F'),
	},
	1205: {
		Fword:  __ccgo_ts + 8316,
		Ftype1: int8('F'),
	},
	1206: {
		Fword:  __ccgo_ts + 8323,
		Ftype1: int8('F'),
	},
	1207: {
		Fword:  __ccgo_ts + 8330,
		Ftype1: int8('F'),
	},
	1208: {
		Fword:  __ccgo_ts + 8337,
		Ftype1: int8('F'),
	},
	1209: {
		Fword:  __ccgo_ts + 8344,
		Ftype1: int8('F'),
	},
	1210: {
		Fword:  __ccgo_ts + 8351,
		Ftype1: int8('F'),
	},
	1211: {
		Fword:  __ccgo_ts + 8358,
		Ftype1: int8('F'),
	},
	1212: {
		Fword:  __ccgo_ts + 8365,
		Ftype1: int8('F'),
	},
	1213: {
		Fword:  __ccgo_ts + 8372,
		Ftype1: int8('F'),
	},
	1214: {
		Fword:  __ccgo_ts + 8379,
		Ftype1: int8('F'),
	},
	1215: {
		Fword:  __ccgo_ts + 8386,
		Ftype1: int8('F'),
	},
	1216: {
		Fword:  __ccgo_ts + 8393,
		Ftype1: int8('F'),
	},
	1217: {
		Fword:  __ccgo_ts + 8400,
		Ftype1: int8('F'),
	},
	1218: {
		Fword:  __ccgo_ts + 8407,
		Ftype1: int8('F'),
	},
	1219: {
		Fword:  __ccgo_ts + 8414,
		Ftype1: int8('F'),
	},
	1220: {
		Fword:  __ccgo_ts + 8421,
		Ftype1: int8('F'),
	},
	1221: {
		Fword:  __ccgo_ts + 8428,
		Ftype1: int8('F'),
	},
	1222: {
		Fword:  __ccgo_ts + 8435,
		Ftype1: int8('F'),
	},
	1223: {
		Fword:  __ccgo_ts + 8440,
		Ftype1: int8('F'),
	},
	1224: {
		Fword:  __ccgo_ts + 8447,
		Ftype1: int8('F'),
	},
	1225: {
		Fword:  __ccgo_ts + 8454,
		Ftype1: int8('F'),
	},
	1226: {
		Fword:  __ccgo_ts + 8461,
		Ftype1: int8('F'),
	},
	1227: {
		Fword:  __ccgo_ts + 8468,
		Ftype1: int8('F'),
	},
	1228: {
		Fword:  __ccgo_ts + 8475,
		Ftype1: int8('F'),
	},
	1229: {
		Fword:  __ccgo_ts + 8482,
		Ftype1: int8('F'),
	},
	1230: {
		Fword:  __ccgo_ts + 8488,
		Ftype1: int8('F'),
	},
	1231: {
		Fword:  __ccgo_ts + 8495,
		Ftype1: int8('F'),
	},
	1232: {
		Fword:  __ccgo_ts + 8502,
		Ftype1: int8('F'),
	},
	1233: {
		Fword:  __ccgo_ts + 8509,
		Ftype1: int8('F'),
	},
	1234: {
		Fword:  __ccgo_ts + 8516,
		Ftype1: int8('F'),
	},
	1235: {
		Fword:  __ccgo_ts + 8523,
		Ftype1: int8('F'),
	},
	1236: {
		Fword:  __ccgo_ts + 8530,
		Ftype1: int8('F'),
	},
	1237: {
		Fword:  __ccgo_ts + 8537,
		Ftype1: int8('F'),
	},
	1238: {
		Fword:  __ccgo_ts + 8544,
		Ftype1: int8('F'),
	},
	1239: {
		Fword:  __ccgo_ts + 8551,
		Ftype1: int8('F'),
	},
	1240: {
		Fword:  __ccgo_ts + 8557,
		Ftype1: int8('F'),
	},
	1241: {
		Fword:  __ccgo_ts + 8564,
		Ftype1: int8('F'),
	},
	1242: {
		Fword:  __ccgo_ts + 8571,
		Ftype1: int8('F'),
	},
	1243: {
		Fword:  __ccgo_ts + 8578,
		Ftype1: int8('F'),
	},
	1244: {
		Fword:  __ccgo_ts + 8585,
		Ftype1: int8('F'),
	},
	1245: {
		Fword:  __ccgo_ts + 8592,
		Ftype1: int8('F'),
	},
	1246: {
		Fword:  __ccgo_ts + 8599,
		Ftype1: int8('F'),
	},
	1247: {
		Fword:  __ccgo_ts + 8606,
		Ftype1: int8('F'),
	},
	1248: {
		Fword:  __ccgo_ts + 8613,
		Ftype1: int8('F'),
	},
	1249: {
		Fword:  __ccgo_ts + 8620,
		Ftype1: int8('F'),
	},
	1250: {
		Fword:  __ccgo_ts + 8627,
		Ftype1: int8('F'),
	},
	1251: {
		Fword:  __ccgo_ts + 8634,
		Ftype1: int8('F'),
	},
	1252: {
		Fword:  __ccgo_ts + 8641,
		Ftype1: int8('F'),
	},
	1253: {
		Fword:  __ccgo_ts + 8648,
		Ftype1: int8('F'),
	},
	1254: {
		Fword:  __ccgo_ts + 8655,
		Ftype1: int8('F'),
	},
	1255: {
		Fword:  __ccgo_ts + 8662,
		Ftype1: int8('F'),
	},
	1256: {
		Fword:  __ccgo_ts + 8669,
		Ftype1: int8('F'),
	},
	1257: {
		Fword:  __ccgo_ts + 8676,
		Ftype1: int8('F'),
	},
	1258: {
		Fword:  __ccgo_ts + 8683,
		Ftype1: int8('F'),
	},
	1259: {
		Fword:  __ccgo_ts + 8690,
		Ftype1: int8('F'),
	},
	1260: {
		Fword:  __ccgo_ts + 8697,
		Ftype1: int8('F'),
	},
	1261: {
		Fword:  __ccgo_ts + 8704,
		Ftype1: int8('F'),
	},
	1262: {
		Fword:  __ccgo_ts + 8711,
		Ftype1: int8('F'),
	},
	1263: {
		Fword:  __ccgo_ts + 8718,
		Ftype1: int8('F'),
	},
	1264: {
		Fword:  __ccgo_ts + 8725,
		Ftype1: int8('F'),
	},
	1265: {
		Fword:  __ccgo_ts + 8732,
		Ftype1: int8('F'),
	},
	1266: {
		Fword:  __ccgo_ts + 8739,
		Ftype1: int8('F'),
	},
	1267: {
		Fword:  __ccgo_ts + 8746,
		Ftype1: int8('F'),
	},
	1268: {
		Fword:  __ccgo_ts + 8753,
		Ftype1: int8('F'),
	},
	1269: {
		Fword:  __ccgo_ts + 8760,
		Ftype1: int8('F'),
	},
	1270: {
		Fword:  __ccgo_ts + 8767,
		Ftype1: int8('F'),
	},
	1271: {
		Fword:  __ccgo_ts + 8774,
		Ftype1: int8('F'),
	},
	1272: {
		Fword:  __ccgo_ts + 8781,
		Ftype1: int8('F'),
	},
	1273: {
		Fword:  __ccgo_ts + 8788,
		Ftype1: int8('F'),
	},
	1274: {
		Fword:  __ccgo_ts + 8795,
		Ftype1: int8('F'),
	},
	1275: {
		Fword:  __ccgo_ts + 8802,
		Ftype1: int8('F'),
	},
	1276: {
		Fword:  __ccgo_ts + 8809,
		Ftype1: int8('F'),
	},
	1277: {
		Fword:  __ccgo_ts + 8815,
		Ftype1: int8('F'),
	},
	1278: {
		Fword:  __ccgo_ts + 8822,
		Ftype1: int8('F'),
	},
	1279: {
		Fword:  __ccgo_ts + 8829,
		Ftype1: int8('F'),
	},
	1280: {
		Fword:  __ccgo_ts + 8836,
		Ftype1: int8('F'),
	},
	1281: {
		Fword:  __ccgo_ts + 8843,
		Ftype1: int8('F'),
	},
	1282: {
		Fword:  __ccgo_ts + 8850,
		Ftype1: int8('F'),
	},
	1283: {
		Fword:  __ccgo_ts + 8857,
		Ftype1: int8('F'),
	},
	1284: {
		Fword:  __ccgo_ts + 8864,
		Ftype1: int8('F'),
	},
	1285: {
		Fword:  __ccgo_ts + 8871,
		Ftype1: int8('F'),
	},
	1286: {
		Fword:  __ccgo_ts + 8878,
		Ftype1: int8('F'),
	},
	1287: {
		Fword:  __ccgo_ts + 8885,
		Ftype1: int8('F'),
	},
	1288: {
		Fword:  __ccgo_ts + 8892,
		Ftype1: int8('F'),
	},
	1289: {
		Fword:  __ccgo_ts + 8899,
		Ftype1: int8('F'),
	},
	1290: {
		Fword:  __ccgo_ts + 8906,
		Ftype1: int8('F'),
	},
	1291: {
		Fword:  __ccgo_ts + 8913,
		Ftype1: int8('F'),
	},
	1292: {
		Fword:  __ccgo_ts + 8920,
		Ftype1: int8('F'),
	},
	1293: {
		Fword:  __ccgo_ts + 8927,
		Ftype1: int8('F'),
	},
	1294: {
		Fword:  __ccgo_ts + 8934,
		Ftype1: int8('F'),
	},
	1295: {
		Fword:  __ccgo_ts + 8941,
		Ftype1: int8('F'),
	},
	1296: {
		Fword:  __ccgo_ts + 8948,
		Ftype1: int8('F'),
	},
	1297: {
		Fword:  __ccgo_ts + 8955,
		Ftype1: int8('F'),
	},
	1298: {
		Fword:  __ccgo_ts + 8962,
		Ftype1: int8('F'),
	},
	1299: {
		Fword:  __ccgo_ts + 8969,
		Ftype1: int8('F'),
	},
	1300: {
		Fword:  __ccgo_ts + 8976,
		Ftype1: int8('F'),
	},
	1301: {
		Fword:  __ccgo_ts + 8983,
		Ftype1: int8('F'),
	},
	1302: {
		Fword:  __ccgo_ts + 8990,
		Ftype1: int8('F'),
	},
	1303: {
		Fword:  __ccgo_ts + 8997,
		Ftype1: int8('F'),
	},
	1304: {
		Fword:  __ccgo_ts + 9004,
		Ftype1: int8('F'),
	},
	1305: {
		Fword:  __ccgo_ts + 9011,
		Ftype1: int8('F'),
	},
	1306: {
		Fword:  __ccgo_ts + 9018,
		Ftype1: int8('F'),
	},
	1307: {
		Fword:  __ccgo_ts + 9025,
		Ftype1: int8('F'),
	},
	1308: {
		Fword:  __ccgo_ts + 9032,
		Ftype1: int8('F'),
	},
	1309: {
		Fword:  __ccgo_ts + 9039,
		Ftype1: int8('F'),
	},
	1310: {
		Fword:  __ccgo_ts + 9046,
		Ftype1: int8('F'),
	},
	1311: {
		Fword:  __ccgo_ts + 9053,
		Ftype1: int8('F'),
	},
	1312: {
		Fword:  __ccgo_ts + 9060,
		Ftype1: int8('F'),
	},
	1313: {
		Fword:  __ccgo_ts + 9067,
		Ftype1: int8('F'),
	},
	1314: {
		Fword:  __ccgo_ts + 9074,
		Ftype1: int8('F'),
	},
	1315: {
		Fword:  __ccgo_ts + 9081,
		Ftype1: int8('F'),
	},
	1316: {
		Fword:  __ccgo_ts + 9088,
		Ftype1: int8('F'),
	},
	1317: {
		Fword:  __ccgo_ts + 9095,
		Ftype1: int8('F'),
	},
	1318: {
		Fword:  __ccgo_ts + 9102,
		Ftype1: int8('F'),
	},
	1319: {
		Fword:  __ccgo_ts + 9109,
		Ftype1: int8('F'),
	},
	1320: {
		Fword:  __ccgo_ts + 9116,
		Ftype1: int8('F'),
	},
	1321: {
		Fword:  __ccgo_ts + 9123,
		Ftype1: int8('F'),
	},
	1322: {
		Fword:  __ccgo_ts + 9130,
		Ftype1: int8('F'),
	},
	1323: {
		Fword:  __ccgo_ts + 9137,
		Ftype1: int8('F'),
	},
	1324: {
		Fword:  __ccgo_ts + 9144,
		Ftype1: int8('F'),
	},
	1325: {
		Fword:  __ccgo_ts + 9151,
		Ftype1: int8('F'),
	},
	1326: {
		Fword:  __ccgo_ts + 9158,
		Ftype1: int8('F'),
	},
	1327: {
		Fword:  __ccgo_ts + 9165,
		Ftype1: int8('F'),
	},
	1328: {
		Fword:  __ccgo_ts + 9172,
		Ftype1: int8('F'),
	},
	1329: {
		Fword:  __ccgo_ts + 9179,
		Ftype1: int8('F'),
	},
	1330: {
		Fword:  __ccgo_ts + 9186,
		Ftype1: int8('F'),
	},
	1331: {
		Fword:  __ccgo_ts + 9193,
		Ftype1: int8('F'),
	},
	1332: {
		Fword:  __ccgo_ts + 9200,
		Ftype1: int8('F'),
	},
	1333: {
		Fword:  __ccgo_ts + 9207,
		Ftype1: int8('F'),
	},
	1334: {
		Fword:  __ccgo_ts + 9214,
		Ftype1: int8('F'),
	},
	1335: {
		Fword:  __ccgo_ts + 9221,
		Ftype1: int8('F'),
	},
	1336: {
		Fword:  __ccgo_ts + 9228,
		Ftype1: int8('F'),
	},
	1337: {
		Fword:  __ccgo_ts + 9235,
		Ftype1: int8('F'),
	},
	1338: {
		Fword:  __ccgo_ts + 9242,
		Ftype1: int8('F'),
	},
	1339: {
		Fword:  __ccgo_ts + 9249,
		Ftype1: int8('F'),
	},
	1340: {
		Fword:  __ccgo_ts + 9256,
		Ftype1: int8('F'),
	},
	1341: {
		Fword:  __ccgo_ts + 9263,
		Ftype1: int8('F'),
	},
	1342: {
		Fword:  __ccgo_ts + 9270,
		Ftype1: int8('F'),
	},
	1343: {
		Fword:  __ccgo_ts + 9277,
		Ftype1: int8('F'),
	},
	1344: {
		Fword:  __ccgo_ts + 9284,
		Ftype1: int8('F'),
	},
	1345: {
		Fword:  __ccgo_ts + 9291,
		Ftype1: int8('F'),
	},
	1346: {
		Fword:  __ccgo_ts + 9298,
		Ftype1: int8('F'),
	},
	1347: {
		Fword:  __ccgo_ts + 9305,
		Ftype1: int8('F'),
	},
	1348: {
		Fword:  __ccgo_ts + 9312,
		Ftype1: int8('F'),
	},
	1349: {
		Fword:  __ccgo_ts + 9319,
		Ftype1: int8('F'),
	},
	1350: {
		Fword:  __ccgo_ts + 9326,
		Ftype1: int8('F'),
	},
	1351: {
		Fword:  __ccgo_ts + 9333,
		Ftype1: int8('F'),
	},
	1352: {
		Fword:  __ccgo_ts + 9340,
		Ftype1: int8('F'),
	},
	1353: {
		Fword:  __ccgo_ts + 9347,
		Ftype1: int8('F'),
	},
	1354: {
		Fword:  __ccgo_ts + 9354,
		Ftype1: int8('F'),
	},
	1355: {
		Fword:  __ccgo_ts + 9361,
		Ftype1: int8('F'),
	},
	1356: {
		Fword:  __ccgo_ts + 9368,
		Ftype1: int8('F'),
	},
	1357: {
		Fword:  __ccgo_ts + 9375,
		Ftype1: int8('F'),
	},
	1358: {
		Fword:  __ccgo_ts + 9382,
		Ftype1: int8('F'),
	},
	1359: {
		Fword:  __ccgo_ts + 9389,
		Ftype1: int8('F'),
	},
	1360: {
		Fword:  __ccgo_ts + 9396,
		Ftype1: int8('F'),
	},
	1361: {
		Fword:  __ccgo_ts + 9403,
		Ftype1: int8('F'),
	},
	1362: {
		Fword:  __ccgo_ts + 9410,
		Ftype1: int8('F'),
	},
	1363: {
		Fword:  __ccgo_ts + 9417,
		Ftype1: int8('F'),
	},
	1364: {
		Fword:  __ccgo_ts + 9424,
		Ftype1: int8('F'),
	},
	1365: {
		Fword:  __ccgo_ts + 9431,
		Ftype1: int8('F'),
	},
	1366: {
		Fword:  __ccgo_ts + 9438,
		Ftype1: int8('F'),
	},
	1367: {
		Fword:  __ccgo_ts + 9445,
		Ftype1: int8('F'),
	},
	1368: {
		Fword:  __ccgo_ts + 9452,
		Ftype1: int8('F'),
	},
	1369: {
		Fword:  __ccgo_ts + 9459,
		Ftype1: int8('F'),
	},
	1370: {
		Fword:  __ccgo_ts + 9466,
		Ftype1: int8('F'),
	},
	1371: {
		Fword:  __ccgo_ts + 9473,
		Ftype1: int8('F'),
	},
	1372: {
		Fword:  __ccgo_ts + 9480,
		Ftype1: int8('F'),
	},
	1373: {
		Fword:  __ccgo_ts + 9487,
		Ftype1: int8('F'),
	},
	1374: {
		Fword:  __ccgo_ts + 9494,
		Ftype1: int8('F'),
	},
	1375: {
		Fword:  __ccgo_ts + 9500,
		Ftype1: int8('F'),
	},
	1376: {
		Fword:  __ccgo_ts + 9507,
		Ftype1: int8('F'),
	},
	1377: {
		Fword:  __ccgo_ts + 9514,
		Ftype1: int8('F'),
	},
	1378: {
		Fword:  __ccgo_ts + 9521,
		Ftype1: int8('F'),
	},
	1379: {
		Fword:  __ccgo_ts + 9528,
		Ftype1: int8('F'),
	},
	1380: {
		Fword:  __ccgo_ts + 9535,
		Ftype1: int8('F'),
	},
	1381: {
		Fword:  __ccgo_ts + 9542,
		Ftype1: int8('F'),
	},
	1382: {
		Fword:  __ccgo_ts + 9549,
		Ftype1: int8('F'),
	},
	1383: {
		Fword:  __ccgo_ts + 9556,
		Ftype1: int8('F'),
	},
	1384: {
		Fword:  __ccgo_ts + 9563,
		Ftype1: int8('F'),
	},
	1385: {
		Fword:  __ccgo_ts + 9570,
		Ftype1: int8('F'),
	},
	1386: {
		Fword:  __ccgo_ts + 9577,
		Ftype1: int8('F'),
	},
	1387: {
		Fword:  __ccgo_ts + 9584,
		Ftype1: int8('F'),
	},
	1388: {
		Fword:  __ccgo_ts + 9591,
		Ftype1: int8('F'),
	},
	1389: {
		Fword:  __ccgo_ts + 9598,
		Ftype1: int8('F'),
	},
	1390: {
		Fword:  __ccgo_ts + 9605,
		Ftype1: int8('F'),
	},
	1391: {
		Fword:  __ccgo_ts + 9612,
		Ftype1: int8('F'),
	},
	1392: {
		Fword:  __ccgo_ts + 9618,
		Ftype1: int8('F'),
	},
	1393: {
		Fword:  __ccgo_ts + 9625,
		Ftype1: int8('F'),
	},
	1394: {
		Fword:  __ccgo_ts + 9632,
		Ftype1: int8('F'),
	},
	1395: {
		Fword:  __ccgo_ts + 9639,
		Ftype1: int8('F'),
	},
	1396: {
		Fword:  __ccgo_ts + 9646,
		Ftype1: int8('F'),
	},
	1397: {
		Fword:  __ccgo_ts + 9653,
		Ftype1: int8('F'),
	},
	1398: {
		Fword:  __ccgo_ts + 9660,
		Ftype1: int8('F'),
	},
	1399: {
		Fword:  __ccgo_ts + 9667,
		Ftype1: int8('F'),
	},
	1400: {
		Fword:  __ccgo_ts + 9674,
		Ftype1: int8('F'),
	},
	1401: {
		Fword:  __ccgo_ts + 9681,
		Ftype1: int8('F'),
	},
	1402: {
		Fword:  __ccgo_ts + 9688,
		Ftype1: int8('F'),
	},
	1403: {
		Fword:  __ccgo_ts + 9695,
		Ftype1: int8('F'),
	},
	1404: {
		Fword:  __ccgo_ts + 9702,
		Ftype1: int8('F'),
	},
	1405: {
		Fword:  __ccgo_ts + 9709,
		Ftype1: int8('F'),
	},
	1406: {
		Fword:  __ccgo_ts + 9716,
		Ftype1: int8('F'),
	},
	1407: {
		Fword:  __ccgo_ts + 9723,
		Ftype1: int8('F'),
	},
	1408: {
		Fword:  __ccgo_ts + 9730,
		Ftype1: int8('F'),
	},
	1409: {
		Fword:  __ccgo_ts + 9737,
		Ftype1: int8('F'),
	},
	1410: {
		Fword:  __ccgo_ts + 9744,
		Ftype1: int8('F'),
	},
	1411: {
		Fword:  __ccgo_ts + 9751,
		Ftype1: int8('F'),
	},
	1412: {
		Fword:  __ccgo_ts + 9758,
		Ftype1: int8('F'),
	},
	1413: {
		Fword:  __ccgo_ts + 9765,
		Ftype1: int8('F'),
	},
	1414: {
		Fword:  __ccgo_ts + 9772,
		Ftype1: int8('F'),
	},
	1415: {
		Fword:  __ccgo_ts + 9779,
		Ftype1: int8('F'),
	},
	1416: {
		Fword:  __ccgo_ts + 9786,
		Ftype1: int8('F'),
	},
	1417: {
		Fword:  __ccgo_ts + 9793,
		Ftype1: int8('F'),
	},
	1418: {
		Fword:  __ccgo_ts + 9800,
		Ftype1: int8('F'),
	},
	1419: {
		Fword:  __ccgo_ts + 9807,
		Ftype1: int8('F'),
	},
	1420: {
		Fword:  __ccgo_ts + 9813,
		Ftype1: int8('F'),
	},
	1421: {
		Fword:  __ccgo_ts + 9820,
		Ftype1: int8('F'),
	},
	1422: {
		Fword:  __ccgo_ts + 9827,
		Ftype1: int8('F'),
	},
	1423: {
		Fword:  __ccgo_ts + 9834,
		Ftype1: int8('F'),
	},
	1424: {
		Fword:  __ccgo_ts + 9841,
		Ftype1: int8('F'),
	},
	1425: {
		Fword:  __ccgo_ts + 9848,
		Ftype1: int8('F'),
	},
	1426: {
		Fword:  __ccgo_ts + 9855,
		Ftype1: int8('F'),
	},
	1427: {
		Fword:  __ccgo_ts + 9862,
		Ftype1: int8('F'),
	},
	1428: {
		Fword:  __ccgo_ts + 9869,
		Ftype1: int8('F'),
	},
	1429: {
		Fword:  __ccgo_ts + 9876,
		Ftype1: int8('F'),
	},
	1430: {
		Fword:  __ccgo_ts + 9883,
		Ftype1: int8('F'),
	},
	1431: {
		Fword:  __ccgo_ts + 9890,
		Ftype1: int8('F'),
	},
	1432: {
		Fword:  __ccgo_ts + 9897,
		Ftype1: int8('F'),
	},
	1433: {
		Fword:  __ccgo_ts + 9904,
		Ftype1: int8('F'),
	},
	1434: {
		Fword:  __ccgo_ts + 9911,
		Ftype1: int8('F'),
	},
	1435: {
		Fword:  __ccgo_ts + 9918,
		Ftype1: int8('F'),
	},
	1436: {
		Fword:  __ccgo_ts + 9925,
		Ftype1: int8('F'),
	},
	1437: {
		Fword:  __ccgo_ts + 9932,
		Ftype1: int8('F'),
	},
	1438: {
		Fword:  __ccgo_ts + 9939,
		Ftype1: int8('F'),
	},
	1439: {
		Fword:  __ccgo_ts + 9946,
		Ftype1: int8('F'),
	},
	1440: {
		Fword:  __ccgo_ts + 9951,
		Ftype1: int8('F'),
	},
	1441: {
		Fword:  __ccgo_ts + 9958,
		Ftype1: int8('F'),
	},
	1442: {
		Fword:  __ccgo_ts + 9965,
		Ftype1: int8('F'),
	},
	1443: {
		Fword:  __ccgo_ts + 9972,
		Ftype1: int8('F'),
	},
	1444: {
		Fword:  __ccgo_ts + 9979,
		Ftype1: int8('F'),
	},
	1445: {
		Fword:  __ccgo_ts + 9986,
		Ftype1: int8('F'),
	},
	1446: {
		Fword:  __ccgo_ts + 9993,
		Ftype1: int8('F'),
	},
	1447: {
		Fword:  __ccgo_ts + 10000,
		Ftype1: int8('F'),
	},
	1448: {
		Fword:  __ccgo_ts + 10007,
		Ftype1: int8('F'),
	},
	1449: {
		Fword:  __ccgo_ts + 10014,
		Ftype1: int8('F'),
	},
	1450: {
		Fword:  __ccgo_ts + 10021,
		Ftype1: int8('F'),
	},
	1451: {
		Fword:  __ccgo_ts + 10028,
		Ftype1: int8('F'),
	},
	1452: {
		Fword:  __ccgo_ts + 10035,
		Ftype1: int8('F'),
	},
	1453: {
		Fword:  __ccgo_ts + 10042,
		Ftype1: int8('F'),
	},
	1454: {
		Fword:  __ccgo_ts + 10049,
		Ftype1: int8('F'),
	},
	1455: {
		Fword:  __ccgo_ts + 10056,
		Ftype1: int8('F'),
	},
	1456: {
		Fword:  __ccgo_ts + 10063,
		Ftype1: int8('F'),
	},
	1457: {
		Fword:  __ccgo_ts + 10070,
		Ftype1: int8('F'),
	},
	1458: {
		Fword:  __ccgo_ts + 10077,
		Ftype1: int8('F'),
	},
	1459: {
		Fword:  __ccgo_ts + 10084,
		Ftype1: int8('F'),
	},
	1460: {
		Fword:  __ccgo_ts + 10091,
		Ftype1: int8('F'),
	},
	1461: {
		Fword:  __ccgo_ts + 10098,
		Ftype1: int8('F'),
	},
	1462: {
		Fword:  __ccgo_ts + 10105,
		Ftype1: int8('F'),
	},
	1463: {
		Fword:  __ccgo_ts + 10112,
		Ftype1: int8('F'),
	},
	1464: {
		Fword:  __ccgo_ts + 10118,
		Ftype1: int8('F'),
	},
	1465: {
		Fword:  __ccgo_ts + 10125,
		Ftype1: int8('F'),
	},
	1466: {
		Fword:  __ccgo_ts + 10132,
		Ftype1: int8('F'),
	},
	1467: {
		Fword:  __ccgo_ts + 10139,
		Ftype1: int8('F'),
	},
	1468: {
		Fword:  __ccgo_ts + 10146,
		Ftype1: int8('F'),
	},
	1469: {
		Fword:  __ccgo_ts + 10153,
		Ftype1: int8('F'),
	},
	1470: {
		Fword:  __ccgo_ts + 10160,
		Ftype1: int8('F'),
	},
	1471: {
		Fword:  __ccgo_ts + 10167,
		Ftype1: int8('F'),
	},
	1472: {
		Fword:  __ccgo_ts + 10174,
		Ftype1: int8('F'),
	},
	1473: {
		Fword:  __ccgo_ts + 10181,
		Ftype1: int8('F'),
	},
	1474: {
		Fword:  __ccgo_ts + 10188,
		Ftype1: int8('F'),
	},
	1475: {
		Fword:  __ccgo_ts + 10195,
		Ftype1: int8('F'),
	},
	1476: {
		Fword:  __ccgo_ts + 10202,
		Ftype1: int8('F'),
	},
	1477: {
		Fword:  __ccgo_ts + 10209,
		Ftype1: int8('F'),
	},
	1478: {
		Fword:  __ccgo_ts + 10216,
		Ftype1: int8('F'),
	},
	1479: {
		Fword:  __ccgo_ts + 10223,
		Ftype1: int8('F'),
	},
	1480: {
		Fword:  __ccgo_ts + 10230,
		Ftype1: int8('F'),
	},
	1481: {
		Fword:  __ccgo_ts + 10236,
		Ftype1: int8('F'),
	},
	1482: {
		Fword:  __ccgo_ts + 10243,
		Ftype1: int8('F'),
	},
	1483: {
		Fword:  __ccgo_ts + 10250,
		Ftype1: int8('F'),
	},
	1484: {
		Fword:  __ccgo_ts + 10257,
		Ftype1: int8('F'),
	},
	1485: {
		Fword:  __ccgo_ts + 10264,
		Ftype1: int8('F'),
	},
	1486: {
		Fword:  __ccgo_ts + 10271,
		Ftype1: int8('F'),
	},
	1487: {
		Fword:  __ccgo_ts + 10278,
		Ftype1: int8('F'),
	},
	1488: {
		Fword:  __ccgo_ts + 10285,
		Ftype1: int8('F'),
	},
	1489: {
		Fword:  __ccgo_ts + 10292,
		Ftype1: int8('F'),
	},
	1490: {
		Fword:  __ccgo_ts + 10299,
		Ftype1: int8('F'),
	},
	1491: {
		Fword:  __ccgo_ts + 10306,
		Ftype1: int8('F'),
	},
	1492: {
		Fword:  __ccgo_ts + 10313,
		Ftype1: int8('F'),
	},
	1493: {
		Fword:  __ccgo_ts + 10320,
		Ftype1: int8('F'),
	},
	1494: {
		Fword:  __ccgo_ts + 10327,
		Ftype1: int8('F'),
	},
	1495: {
		Fword:  __ccgo_ts + 10334,
		Ftype1: int8('F'),
	},
	1496: {
		Fword:  __ccgo_ts + 10341,
		Ftype1: int8('F'),
	},
	1497: {
		Fword:  __ccgo_ts + 10348,
		Ftype1: int8('F'),
	},
	1498: {
		Fword:  __ccgo_ts + 10355,
		Ftype1: int8('F'),
	},
	1499: {
		Fword:  __ccgo_ts + 10362,
		Ftype1: int8('F'),
	},
	1500: {
		Fword:  __ccgo_ts + 10369,
		Ftype1: int8('F'),
	},
	1501: {
		Fword:  __ccgo_ts + 10376,
		Ftype1: int8('F'),
	},
	1502: {
		Fword:  __ccgo_ts + 10383,
		Ftype1: int8('F'),
	},
	1503: {
		Fword:  __ccgo_ts + 10390,
		Ftype1: int8('F'),
	},
	1504: {
		Fword:  __ccgo_ts + 10397,
		Ftype1: int8('F'),
	},
	1505: {
		Fword:  __ccgo_ts + 10404,
		Ftype1: int8('F'),
	},
	1506: {
		Fword:  __ccgo_ts + 10411,
		Ftype1: int8('F'),
	},
	1507: {
		Fword:  __ccgo_ts + 10418,
		Ftype1: int8('F'),
	},
	1508: {
		Fword:  __ccgo_ts + 10425,
		Ftype1: int8('F'),
	},
	1509: {
		Fword:  __ccgo_ts + 10432,
		Ftype1: int8('F'),
	},
	1510: {
		Fword:  __ccgo_ts + 10439,
		Ftype1: int8('F'),
	},
	1511: {
		Fword:  __ccgo_ts + 10446,
		Ftype1: int8('F'),
	},
	1512: {
		Fword:  __ccgo_ts + 10453,
		Ftype1: int8('F'),
	},
	1513: {
		Fword:  __ccgo_ts + 10460,
		Ftype1: int8('F'),
	},
	1514: {
		Fword:  __ccgo_ts + 10467,
		Ftype1: int8('F'),
	},
	1515: {
		Fword:  __ccgo_ts + 10474,
		Ftype1: int8('F'),
	},
	1516: {
		Fword:  __ccgo_ts + 10481,
		Ftype1: int8('F'),
	},
	1517: {
		Fword:  __ccgo_ts + 10488,
		Ftype1: int8('F'),
	},
	1518: {
		Fword:  __ccgo_ts + 10495,
		Ftype1: int8('F'),
	},
	1519: {
		Fword:  __ccgo_ts + 10502,
		Ftype1: int8('F'),
	},
	1520: {
		Fword:  __ccgo_ts + 10508,
		Ftype1: int8('F'),
	},
	1521: {
		Fword:  __ccgo_ts + 10515,
		Ftype1: int8('F'),
	},
	1522: {
		Fword:  __ccgo_ts + 10522,
		Ftype1: int8('F'),
	},
	1523: {
		Fword:  __ccgo_ts + 10529,
		Ftype1: int8('F'),
	},
	1524: {
		Fword:  __ccgo_ts + 10536,
		Ftype1: int8('F'),
	},
	1525: {
		Fword:  __ccgo_ts + 10543,
		Ftype1: int8('F'),
	},
	1526: {
		Fword:  __ccgo_ts + 10550,
		Ftype1: int8('F'),
	},
	1527: {
		Fword:  __ccgo_ts + 10557,
		Ftype1: int8('F'),
	},
	1528: {
		Fword:  __ccgo_ts + 10564,
		Ftype1: int8('F'),
	},
	1529: {
		Fword:  __ccgo_ts + 10571,
		Ftype1: int8('F'),
	},
	1530: {
		Fword:  __ccgo_ts + 10578,
		Ftype1: int8('F'),
	},
	1531: {
		Fword:  __ccgo_ts + 10585,
		Ftype1: int8('F'),
	},
	1532: {
		Fword:  __ccgo_ts + 10592,
		Ftype1: int8('F'),
	},
	1533: {
		Fword:  __ccgo_ts + 10599,
		Ftype1: int8('F'),
	},
	1534: {
		Fword:  __ccgo_ts + 10606,
		Ftype1: int8('F'),
	},
	1535: {
		Fword:  __ccgo_ts + 10613,
		Ftype1: int8('F'),
	},
	1536: {
		Fword:  __ccgo_ts + 10620,
		Ftype1: int8('F'),
	},
	1537: {
		Fword:  __ccgo_ts + 10627,
		Ftype1: int8('F'),
	},
	1538: {
		Fword:  __ccgo_ts + 10633,
		Ftype1: int8('F'),
	},
	1539: {
		Fword:  __ccgo_ts + 10640,
		Ftype1: int8('F'),
	},
	1540: {
		Fword:  __ccgo_ts + 10647,
		Ftype1: int8('F'),
	},
	1541: {
		Fword:  __ccgo_ts + 10654,
		Ftype1: int8('F'),
	},
	1542: {
		Fword:  __ccgo_ts + 10659,
		Ftype1: int8('F'),
	},
	1543: {
		Fword:  __ccgo_ts + 10665,
		Ftype1: int8('F'),
	},
	1544: {
		Fword:  __ccgo_ts + 10672,
		Ftype1: int8('F'),
	},
	1545: {
		Fword:  __ccgo_ts + 10678,
		Ftype1: int8('F'),
	},
	1546: {
		Fword:  __ccgo_ts + 10685,
		Ftype1: int8('F'),
	},
	1547: {
		Fword:  __ccgo_ts + 10692,
		Ftype1: int8('F'),
	},
	1548: {
		Fword:  __ccgo_ts + 10699,
		Ftype1: int8('F'),
	},
	1549: {
		Fword:  __ccgo_ts + 10706,
		Ftype1: int8('F'),
	},
	1550: {
		Fword:  __ccgo_ts + 10713,
		Ftype1: int8('F'),
	},
	1551: {
		Fword:  __ccgo_ts + 10720,
		Ftype1: int8('F'),
	},
	1552: {
		Fword:  __ccgo_ts + 10727,
		Ftype1: int8('F'),
	},
	1553: {
		Fword:  __ccgo_ts + 10734,
		Ftype1: int8('F'),
	},
	1554: {
		Fword:  __ccgo_ts + 10741,
		Ftype1: int8('F'),
	},
	1555: {
		Fword:  __ccgo_ts + 10748,
		Ftype1: int8('F'),
	},
	1556: {
		Fword:  __ccgo_ts + 10755,
		Ftype1: int8('F'),
	},
	1557: {
		Fword:  __ccgo_ts + 10762,
		Ftype1: int8('F'),
	},
	1558: {
		Fword:  __ccgo_ts + 10769,
		Ftype1: int8('F'),
	},
	1559: {
		Fword:  __ccgo_ts + 10776,
		Ftype1: int8('F'),
	},
	1560: {
		Fword:  __ccgo_ts + 10783,
		Ftype1: int8('F'),
	},
	1561: {
		Fword:  __ccgo_ts + 10790,
		Ftype1: int8('F'),
	},
	1562: {
		Fword:  __ccgo_ts + 10797,
		Ftype1: int8('F'),
	},
	1563: {
		Fword:  __ccgo_ts + 10804,
		Ftype1: int8('F'),
	},
	1564: {
		Fword:  __ccgo_ts + 10811,
		Ftype1: int8('F'),
	},
	1565: {
		Fword:  __ccgo_ts + 10818,
		Ftype1: int8('F'),
	},
	1566: {
		Fword:  __ccgo_ts + 10825,
		Ftype1: int8('F'),
	},
	1567: {
		Fword:  __ccgo_ts + 10832,
		Ftype1: int8('F'),
	},
	1568: {
		Fword:  __ccgo_ts + 10839,
		Ftype1: int8('F'),
	},
	1569: {
		Fword:  __ccgo_ts + 10846,
		Ftype1: int8('F'),
	},
	1570: {
		Fword:  __ccgo_ts + 10853,
		Ftype1: int8('F'),
	},
	1571: {
		Fword:  __ccgo_ts + 10860,
		Ftype1: int8('F'),
	},
	1572: {
		Fword:  __ccgo_ts + 10867,
		Ftype1: int8('F'),
	},
	1573: {
		Fword:  __ccgo_ts + 10874,
		Ftype1: int8('F'),
	},
	1574: {
		Fword:  __ccgo_ts + 10881,
		Ftype1: int8('F'),
	},
	1575: {
		Fword:  __ccgo_ts + 10888,
		Ftype1: int8('F'),
	},
	1576: {
		Fword:  __ccgo_ts + 10895,
		Ftype1: int8('F'),
	},
	1577: {
		Fword:  __ccgo_ts + 10902,
		Ftype1: int8('F'),
	},
	1578: {
		Fword:  __ccgo_ts + 10909,
		Ftype1: int8('F'),
	},
	1579: {
		Fword:  __ccgo_ts + 10916,
		Ftype1: int8('F'),
	},
	1580: {
		Fword:  __ccgo_ts + 10923,
		Ftype1: int8('F'),
	},
	1581: {
		Fword:  __ccgo_ts + 10930,
		Ftype1: int8('F'),
	},
	1582: {
		Fword:  __ccgo_ts + 10937,
		Ftype1: int8('F'),
	},
	1583: {
		Fword:  __ccgo_ts + 10944,
		Ftype1: int8('F'),
	},
	1584: {
		Fword:  __ccgo_ts + 10951,
		Ftype1: int8('F'),
	},
	1585: {
		Fword:  __ccgo_ts + 10958,
		Ftype1: int8('F'),
	},
	1586: {
		Fword:  __ccgo_ts + 10965,
		Ftype1: int8('F'),
	},
	1587: {
		Fword:  __ccgo_ts + 10972,
		Ftype1: int8('F'),
	},
	1588: {
		Fword:  __ccgo_ts + 10979,
		Ftype1: int8('F'),
	},
	1589: {
		Fword:  __ccgo_ts + 10986,
		Ftype1: int8('F'),
	},
	1590: {
		Fword:  __ccgo_ts + 10993,
		Ftype1: int8('F'),
	},
	1591: {
		Fword:  __ccgo_ts + 11000,
		Ftype1: int8('F'),
	},
	1592: {
		Fword:  __ccgo_ts + 11007,
		Ftype1: int8('F'),
	},
	1593: {
		Fword:  __ccgo_ts + 11014,
		Ftype1: int8('F'),
	},
	1594: {
		Fword:  __ccgo_ts + 11021,
		Ftype1: int8('F'),
	},
	1595: {
		Fword:  __ccgo_ts + 11028,
		Ftype1: int8('F'),
	},
	1596: {
		Fword:  __ccgo_ts + 11035,
		Ftype1: int8('F'),
	},
	1597: {
		Fword:  __ccgo_ts + 11042,
		Ftype1: int8('F'),
	},
	1598: {
		Fword:  __ccgo_ts + 11049,
		Ftype1: int8('F'),
	},
	1599: {
		Fword:  __ccgo_ts + 11056,
		Ftype1: int8('F'),
	},
	1600: {
		Fword:  __ccgo_ts + 11063,
		Ftype1: int8('F'),
	},
	1601: {
		Fword:  __ccgo_ts + 11070,
		Ftype1: int8('F'),
	},
	1602: {
		Fword:  __ccgo_ts + 11077,
		Ftype1: int8('F'),
	},
	1603: {
		Fword:  __ccgo_ts + 11084,
		Ftype1: int8('F'),
	},
	1604: {
		Fword:  __ccgo_ts + 11091,
		Ftype1: int8('F'),
	},
	1605: {
		Fword:  __ccgo_ts + 11098,
		Ftype1: int8('F'),
	},
	1606: {
		Fword:  __ccgo_ts + 11105,
		Ftype1: int8('F'),
	},
	1607: {
		Fword:  __ccgo_ts + 11112,
		Ftype1: int8('F'),
	},
	1608: {
		Fword:  __ccgo_ts + 11119,
		Ftype1: int8('F'),
	},
	1609: {
		Fword:  __ccgo_ts + 11126,
		Ftype1: int8('F'),
	},
	1610: {
		Fword:  __ccgo_ts + 11133,
		Ftype1: int8('F'),
	},
	1611: {
		Fword:  __ccgo_ts + 11140,
		Ftype1: int8('F'),
	},
	1612: {
		Fword:  __ccgo_ts + 11147,
		Ftype1: int8('F'),
	},
	1613: {
		Fword:  __ccgo_ts + 11154,
		Ftype1: int8('F'),
	},
	1614: {
		Fword:  __ccgo_ts + 11161,
		Ftype1: int8('F'),
	},
	1615: {
		Fword:  __ccgo_ts + 11168,
		Ftype1: int8('F'),
	},
	1616: {
		Fword:  __ccgo_ts + 11175,
		Ftype1: int8('F'),
	},
	1617: {
		Fword:  __ccgo_ts + 11182,
		Ftype1: int8('F'),
	},
	1618: {
		Fword:  __ccgo_ts + 11189,
		Ftype1: int8('F'),
	},
	1619: {
		Fword:  __ccgo_ts + 11196,
		Ftype1: int8('F'),
	},
	1620: {
		Fword:  __ccgo_ts + 11203,
		Ftype1: int8('F'),
	},
	1621: {
		Fword:  __ccgo_ts + 11210,
		Ftype1: int8('F'),
	},
	1622: {
		Fword:  __ccgo_ts + 11217,
		Ftype1: int8('F'),
	},
	1623: {
		Fword:  __ccgo_ts + 11224,
		Ftype1: int8('F'),
	},
	1624: {
		Fword:  __ccgo_ts + 11231,
		Ftype1: int8('F'),
	},
	1625: {
		Fword:  __ccgo_ts + 11238,
		Ftype1: int8('F'),
	},
	1626: {
		Fword:  __ccgo_ts + 11245,
		Ftype1: int8('F'),
	},
	1627: {
		Fword:  __ccgo_ts + 11252,
		Ftype1: int8('F'),
	},
	1628: {
		Fword:  __ccgo_ts + 11259,
		Ftype1: int8('F'),
	},
	1629: {
		Fword:  __ccgo_ts + 11266,
		Ftype1: int8('F'),
	},
	1630: {
		Fword:  __ccgo_ts + 11273,
		Ftype1: int8('F'),
	},
	1631: {
		Fword:  __ccgo_ts + 11280,
		Ftype1: int8('F'),
	},
	1632: {
		Fword:  __ccgo_ts + 11287,
		Ftype1: int8('F'),
	},
	1633: {
		Fword:  __ccgo_ts + 11294,
		Ftype1: int8('F'),
	},
	1634: {
		Fword:  __ccgo_ts + 11301,
		Ftype1: int8('F'),
	},
	1635: {
		Fword:  __ccgo_ts + 11308,
		Ftype1: int8('F'),
	},
	1636: {
		Fword:  __ccgo_ts + 11315,
		Ftype1: int8('F'),
	},
	1637: {
		Fword:  __ccgo_ts + 11322,
		Ftype1: int8('F'),
	},
	1638: {
		Fword:  __ccgo_ts + 11329,
		Ftype1: int8('F'),
	},
	1639: {
		Fword:  __ccgo_ts + 11333,
		Ftype1: int8('F'),
	},
	1640: {
		Fword:  __ccgo_ts + 11340,
		Ftype1: int8('F'),
	},
	1641: {
		Fword:  __ccgo_ts + 11347,
		Ftype1: int8('F'),
	},
	1642: {
		Fword:  __ccgo_ts + 11354,
		Ftype1: int8('F'),
	},
	1643: {
		Fword:  __ccgo_ts + 11361,
		Ftype1: int8('F'),
	},
	1644: {
		Fword:  __ccgo_ts + 11368,
		Ftype1: int8('F'),
	},
	1645: {
		Fword:  __ccgo_ts + 11375,
		Ftype1: int8('F'),
	},
	1646: {
		Fword:  __ccgo_ts + 11382,
		Ftype1: int8('F'),
	},
	1647: {
		Fword:  __ccgo_ts + 11389,
		Ftype1: int8('F'),
	},
	1648: {
		Fword:  __ccgo_ts + 11396,
		Ftype1: int8('F'),
	},
	1649: {
		Fword:  __ccgo_ts + 11403,
		Ftype1: int8('F'),
	},
	1650: {
		Fword:  __ccgo_ts + 11410,
		Ftype1: int8('F'),
	},
	1651: {
		Fword:  __ccgo_ts + 11417,
		Ftype1: int8('F'),
	},
	1652: {
		Fword:  __ccgo_ts + 11424,
		Ftype1: int8('F'),
	},
	1653: {
		Fword:  __ccgo_ts + 11431,
		Ftype1: int8('F'),
	},
	1654: {
		Fword:  __ccgo_ts + 11438,
		Ftype1: int8('F'),
	},
	1655: {
		Fword:  __ccgo_ts + 11445,
		Ftype1: int8('F'),
	},
	1656: {
		Fword:  __ccgo_ts + 11452,
		Ftype1: int8('F'),
	},
	1657: {
		Fword:  __ccgo_ts + 11459,
		Ftype1: int8('F'),
	},
	1658: {
		Fword:  __ccgo_ts + 11465,
		Ftype1: int8('F'),
	},
	1659: {
		Fword:  __ccgo_ts + 11472,
		Ftype1: int8('F'),
	},
	1660: {
		Fword:  __ccgo_ts + 11479,
		Ftype1: int8('F'),
	},
	1661: {
		Fword:  __ccgo_ts + 11486,
		Ftype1: int8('F'),
	},
	1662: {
		Fword:  __ccgo_ts + 11493,
		Ftype1: int8('F'),
	},
	1663: {
		Fword:  __ccgo_ts + 11498,
		Ftype1: int8('F'),
	},
	1664: {
		Fword:  __ccgo_ts + 11504,
		Ftype1: int8('F'),
	},
	1665: {
		Fword:  __ccgo_ts + 11509,
		Ftype1: int8('F'),
	},
	1666: {
		Fword:  __ccgo_ts + 11514,
		Ftype1: int8('F'),
	},
	1667: {
		Fword:  __ccgo_ts + 11521,
		Ftype1: int8('F'),
	},
	1668: {
		Fword:  __ccgo_ts + 11528,
		Ftype1: int8('F'),
	},
	1669: {
		Fword:  __ccgo_ts + 11535,
		Ftype1: int8('F'),
	},
	1670: {
		Fword:  __ccgo_ts + 11542,
		Ftype1: int8('F'),
	},
	1671: {
		Fword:  __ccgo_ts + 11549,
		Ftype1: int8('F'),
	},
	1672: {
		Fword:  __ccgo_ts + 11556,
		Ftype1: int8('F'),
	},
	1673: {
		Fword:  __ccgo_ts + 11563,
		Ftype1: int8('F'),
	},
	1674: {
		Fword:  __ccgo_ts + 11569,
		Ftype1: int8('F'),
	},
	1675: {
		Fword:  __ccgo_ts + 11576,
		Ftype1: int8('F'),
	},
	1676: {
		Fword:  __ccgo_ts + 11583,
		Ftype1: int8('F'),
	},
	1677: {
		Fword:  __ccgo_ts + 11590,
		Ftype1: int8('F'),
	},
	1678: {
		Fword:  __ccgo_ts + 11597,
		Ftype1: int8('F'),
	},
	1679: {
		Fword:  __ccgo_ts + 11604,
		Ftype1: int8('F'),
	},
	1680: {
		Fword:  __ccgo_ts + 11611,
		Ftype1: int8('F'),
	},
	1681: {
		Fword:  __ccgo_ts + 11618,
		Ftype1: int8('F'),
	},
	1682: {
		Fword:  __ccgo_ts + 11625,
		Ftype1: int8('F'),
	},
	1683: {
		Fword:  __ccgo_ts + 11632,
		Ftype1: int8('F'),
	},
	1684: {
		Fword:  __ccgo_ts + 11639,
		Ftype1: int8('F'),
	},
	1685: {
		Fword:  __ccgo_ts + 11646,
		Ftype1: int8('F'),
	},
	1686: {
		Fword:  __ccgo_ts + 11653,
		Ftype1: int8('F'),
	},
	1687: {
		Fword:  __ccgo_ts + 11660,
		Ftype1: int8('F'),
	},
	1688: {
		Fword:  __ccgo_ts + 11667,
		Ftype1: int8('F'),
	},
	1689: {
		Fword:  __ccgo_ts + 11673,
		Ftype1: int8('F'),
	},
	1690: {
		Fword:  __ccgo_ts + 11680,
		Ftype1: int8('F'),
	},
	1691: {
		Fword:  __ccgo_ts + 11686,
		Ftype1: int8('F'),
	},
	1692: {
		Fword:  __ccgo_ts + 11692,
		Ftype1: int8('F'),
	},
	1693: {
		Fword:  __ccgo_ts + 11699,
		Ftype1: int8('F'),
	},
	1694: {
		Fword:  __ccgo_ts + 11706,
		Ftype1: int8('F'),
	},
	1695: {
		Fword:  __ccgo_ts + 11713,
		Ftype1: int8('F'),
	},
	1696: {
		Fword:  __ccgo_ts + 11720,
		Ftype1: int8('F'),
	},
	1697: {
		Fword:  __ccgo_ts + 11726,
		Ftype1: int8('F'),
	},
	1698: {
		Fword:  __ccgo_ts + 11733,
		Ftype1: int8('F'),
	},
	1699: {
		Fword:  __ccgo_ts + 11740,
		Ftype1: int8('F'),
	},
	1700: {
		Fword:  __ccgo_ts + 11747,
		Ftype1: int8('F'),
	},
	1701: {
		Fword:  __ccgo_ts + 11754,
		Ftype1: int8('F'),
	},
	1702: {
		Fword:  __ccgo_ts + 11761,
		Ftype1: int8('F'),
	},
	1703: {
		Fword:  __ccgo_ts + 11768,
		Ftype1: int8('F'),
	},
	1704: {
		Fword:  __ccgo_ts + 11775,
		Ftype1: int8('F'),
	},
	1705: {
		Fword:  __ccgo_ts + 11782,
		Ftype1: int8('F'),
	},
	1706: {
		Fword:  __ccgo_ts + 11789,
		Ftype1: int8('F'),
	},
	1707: {
		Fword:  __ccgo_ts + 11795,
		Ftype1: int8('F'),
	},
	1708: {
		Fword:  __ccgo_ts + 11802,
		Ftype1: int8('F'),
	},
	1709: {
		Fword:  __ccgo_ts + 11809,
		Ftype1: int8('F'),
	},
	1710: {
		Fword:  __ccgo_ts + 11816,
		Ftype1: int8('F'),
	},
	1711: {
		Fword:  __ccgo_ts + 11823,
		Ftype1: int8('F'),
	},
	1712: {
		Fword:  __ccgo_ts + 11830,
		Ftype1: int8('F'),
	},
	1713: {
		Fword:  __ccgo_ts + 11837,
		Ftype1: int8('F'),
	},
	1714: {
		Fword:  __ccgo_ts + 11844,
		Ftype1: int8('F'),
	},
	1715: {
		Fword:  __ccgo_ts + 11851,
		Ftype1: int8('F'),
	},
	1716: {
		Fword:  __ccgo_ts + 11858,
		Ftype1: int8('F'),
	},
	1717: {
		Fword:  __ccgo_ts + 11865,
		Ftype1: int8('F'),
	},
	1718: {
		Fword:  __ccgo_ts + 11872,
		Ftype1: int8('F'),
	},
	1719: {
		Fword:  __ccgo_ts + 11879,
		Ftype1: int8('F'),
	},
	1720: {
		Fword:  __ccgo_ts + 11886,
		Ftype1: int8('F'),
	},
	1721: {
		Fword:  __ccgo_ts + 11893,
		Ftype1: int8('F'),
	},
	1722: {
		Fword:  __ccgo_ts + 11900,
		Ftype1: int8('F'),
	},
	1723: {
		Fword:  __ccgo_ts + 11907,
		Ftype1: int8('F'),
	},
	1724: {
		Fword:  __ccgo_ts + 11913,
		Ftype1: int8('F'),
	},
	1725: {
		Fword:  __ccgo_ts + 11920,
		Ftype1: int8('F'),
	},
	1726: {
		Fword:  __ccgo_ts + 11927,
		Ftype1: int8('F'),
	},
	1727: {
		Fword:  __ccgo_ts + 11934,
		Ftype1: int8('F'),
	},
	1728: {
		Fword:  __ccgo_ts + 11941,
		Ftype1: int8('F'),
	},
	1729: {
		Fword:  __ccgo_ts + 11948,
		Ftype1: int8('F'),
	},
	1730: {
		Fword:  __ccgo_ts + 11955,
		Ftype1: int8('F'),
	},
	1731: {
		Fword:  __ccgo_ts + 11962,
		Ftype1: int8('F'),
	},
	1732: {
		Fword:  __ccgo_ts + 11969,
		Ftype1: int8('F'),
	},
	1733: {
		Fword:  __ccgo_ts + 11976,
		Ftype1: int8('F'),
	},
	1734: {
		Fword:  __ccgo_ts + 11983,
		Ftype1: int8('F'),
	},
	1735: {
		Fword:  __ccgo_ts + 11990,
		Ftype1: int8('F'),
	},
	1736: {
		Fword:  __ccgo_ts + 11997,
		Ftype1: int8('F'),
	},
	1737: {
		Fword:  __ccgo_ts + 12004,
		Ftype1: int8('F'),
	},
	1738: {
		Fword:  __ccgo_ts + 12010,
		Ftype1: int8('F'),
	},
	1739: {
		Fword:  __ccgo_ts + 12017,
		Ftype1: int8('F'),
	},
	1740: {
		Fword:  __ccgo_ts + 12024,
		Ftype1: int8('F'),
	},
	1741: {
		Fword:  __ccgo_ts + 12031,
		Ftype1: int8('F'),
	},
	1742: {
		Fword:  __ccgo_ts + 12038,
		Ftype1: int8('F'),
	},
	1743: {
		Fword:  __ccgo_ts + 12045,
		Ftype1: int8('F'),
	},
	1744: {
		Fword:  __ccgo_ts + 12052,
		Ftype1: int8('F'),
	},
	1745: {
		Fword:  __ccgo_ts + 12059,
		Ftype1: int8('F'),
	},
	1746: {
		Fword:  __ccgo_ts + 12066,
		Ftype1: int8('F'),
	},
	1747: {
		Fword:  __ccgo_ts + 12073,
		Ftype1: int8('F'),
	},
	1748: {
		Fword:  __ccgo_ts + 12080,
		Ftype1: int8('F'),
	},
	1749: {
		Fword:  __ccgo_ts + 12087,
		Ftype1: int8('F'),
	},
	1750: {
		Fword:  __ccgo_ts + 12094,
		Ftype1: int8('F'),
	},
	1751: {
		Fword:  __ccgo_ts + 12101,
		Ftype1: int8('F'),
	},
	1752: {
		Fword:  __ccgo_ts + 12108,
		Ftype1: int8('F'),
	},
	1753: {
		Fword:  __ccgo_ts + 12115,
		Ftype1: int8('F'),
	},
	1754: {
		Fword:  __ccgo_ts + 12122,
		Ftype1: int8('F'),
	},
	1755: {
		Fword:  __ccgo_ts + 12129,
		Ftype1: int8('F'),
	},
	1756: {
		Fword:  __ccgo_ts + 12136,
		Ftype1: int8('F'),
	},
	1757: {
		Fword:  __ccgo_ts + 12143,
		Ftype1: int8('F'),
	},
	1758: {
		Fword:  __ccgo_ts + 12150,
		Ftype1: int8('F'),
	},
	1759: {
		Fword:  __ccgo_ts + 12157,
		Ftype1: int8('F'),
	},
	1760: {
		Fword:  __ccgo_ts + 12164,
		Ftype1: int8('F'),
	},
	1761: {
		Fword:  __ccgo_ts + 12171,
		Ftype1: int8('F'),
	},
	1762: {
		Fword:  __ccgo_ts + 12178,
		Ftype1: int8('F'),
	},
	1763: {
		Fword:  __ccgo_ts + 12185,
		Ftype1: int8('F'),
	},
	1764: {
		Fword:  __ccgo_ts + 12192,
		Ftype1: int8('F'),
	},
	1765: {
		Fword:  __ccgo_ts + 12199,
		Ftype1: int8('F'),
	},
	1766: {
		Fword:  __ccgo_ts + 12205,
		Ftype1: int8('F'),
	},
	1767: {
		Fword:  __ccgo_ts + 12212,
		Ftype1: int8('F'),
	},
	1768: {
		Fword:  __ccgo_ts + 12219,
		Ftype1: int8('F'),
	},
	1769: {
		Fword:  __ccgo_ts + 12226,
		Ftype1: int8('F'),
	},
	1770: {
		Fword:  __ccgo_ts + 12233,
		Ftype1: int8('F'),
	},
	1771: {
		Fword:  __ccgo_ts + 12240,
		Ftype1: int8('F'),
	},
	1772: {
		Fword:  __ccgo_ts + 12247,
		Ftype1: int8('F'),
	},
	1773: {
		Fword:  __ccgo_ts + 12254,
		Ftype1: int8('F'),
	},
	1774: {
		Fword:  __ccgo_ts + 12261,
		Ftype1: int8('F'),
	},
	1775: {
		Fword:  __ccgo_ts + 12268,
		Ftype1: int8('F'),
	},
	1776: {
		Fword:  __ccgo_ts + 12275,
		Ftype1: int8('F'),
	},
	1777: {
		Fword:  __ccgo_ts + 12282,
		Ftype1: int8('F'),
	},
	1778: {
		Fword:  __ccgo_ts + 12289,
		Ftype1: int8('F'),
	},
	1779: {
		Fword:  __ccgo_ts + 12296,
		Ftype1: int8('F'),
	},
	1780: {
		Fword:  __ccgo_ts + 12303,
		Ftype1: int8('F'),
	},
	1781: {
		Fword:  __ccgo_ts + 12309,
		Ftype1: int8('F'),
	},
	1782: {
		Fword:  __ccgo_ts + 12316,
		Ftype1: int8('F'),
	},
	1783: {
		Fword:  __ccgo_ts + 12323,
		Ftype1: int8('F'),
	},
	1784: {
		Fword:  __ccgo_ts + 12330,
		Ftype1: int8('F'),
	},
	1785: {
		Fword:  __ccgo_ts + 12337,
		Ftype1: int8('F'),
	},
	1786: {
		Fword:  __ccgo_ts + 12344,
		Ftype1: int8('F'),
	},
	1787: {
		Fword:  __ccgo_ts + 12351,
		Ftype1: int8('F'),
	},
	1788: {
		Fword:  __ccgo_ts + 12358,
		Ftype1: int8('F'),
	},
	1789: {
		Fword:  __ccgo_ts + 12365,
		Ftype1: int8('F'),
	},
	1790: {
		Fword:  __ccgo_ts + 12372,
		Ftype1: int8('F'),
	},
	1791: {
		Fword:  __ccgo_ts + 12379,
		Ftype1: int8('F'),
	},
	1792: {
		Fword:  __ccgo_ts + 12386,
		Ftype1: int8('F'),
	},
	1793: {
		Fword:  __ccgo_ts + 12393,
		Ftype1: int8('F'),
	},
	1794: {
		Fword:  __ccgo_ts + 12400,
		Ftype1: int8('F'),
	},
	1795: {
		Fword:  __ccgo_ts + 12407,
		Ftype1: int8('F'),
	},
	1796: {
		Fword:  __ccgo_ts + 12414,
		Ftype1: int8('F'),
	},
	1797: {
		Fword:  __ccgo_ts + 12421,
		Ftype1: int8('F'),
	},
	1798: {
		Fword:  __ccgo_ts + 12428,
		Ftype1: int8('F'),
	},
	1799: {
		Fword:  __ccgo_ts + 12435,
		Ftype1: int8('F'),
	},
	1800: {
		Fword:  __ccgo_ts + 12442,
		Ftype1: int8('F'),
	},
	1801: {
		Fword:  __ccgo_ts + 12449,
		Ftype1: int8('F'),
	},
	1802: {
		Fword:  __ccgo_ts + 12454,
		Ftype1: int8('F'),
	},
	1803: {
		Fword:  __ccgo_ts + 12460,
		Ftype1: int8('F'),
	},
	1804: {
		Fword:  __ccgo_ts + 12467,
		Ftype1: int8('F'),
	},
	1805: {
		Fword:  __ccgo_ts + 12473,
		Ftype1: int8('F'),
	},
	1806: {
		Fword:  __ccgo_ts + 12480,
		Ftype1: int8('F'),
	},
	1807: {
		Fword:  __ccgo_ts + 12487,
		Ftype1: int8('F'),
	},
	1808: {
		Fword:  __ccgo_ts + 12494,
		Ftype1: int8('F'),
	},
	1809: {
		Fword:  __ccgo_ts + 12501,
		Ftype1: int8('F'),
	},
	1810: {
		Fword:  __ccgo_ts + 12508,
		Ftype1: int8('F'),
	},
	1811: {
		Fword:  __ccgo_ts + 12515,
		Ftype1: int8('F'),
	},
	1812: {
		Fword:  __ccgo_ts + 12521,
		Ftype1: int8('F'),
	},
	1813: {
		Fword:  __ccgo_ts + 12528,
		Ftype1: int8('F'),
	},
	1814: {
		Fword:  __ccgo_ts + 12535,
		Ftype1: int8('F'),
	},
	1815: {
		Fword:  __ccgo_ts + 12542,
		Ftype1: int8('F'),
	},
	1816: {
		Fword:  __ccgo_ts + 12549,
		Ftype1: int8('F'),
	},
	1817: {
		Fword:  __ccgo_ts + 12556,
		Ftype1: int8('F'),
	},
	1818: {
		Fword:  __ccgo_ts + 12563,
		Ftype1: int8('F'),
	},
	1819: {
		Fword:  __ccgo_ts + 12570,
		Ftype1: int8('F'),
	},
	1820: {
		Fword:  __ccgo_ts + 12577,
		Ftype1: int8('F'),
	},
	1821: {
		Fword:  __ccgo_ts + 12584,
		Ftype1: int8('F'),
	},
	1822: {
		Fword:  __ccgo_ts + 12591,
		Ftype1: int8('F'),
	},
	1823: {
		Fword:  __ccgo_ts + 12598,
		Ftype1: int8('F'),
	},
	1824: {
		Fword:  __ccgo_ts + 12605,
		Ftype1: int8('F'),
	},
	1825: {
		Fword:  __ccgo_ts + 12612,
		Ftype1: int8('F'),
	},
	1826: {
		Fword:  __ccgo_ts + 12619,
		Ftype1: int8('F'),
	},
	1827: {
		Fword:  __ccgo_ts + 12626,
		Ftype1: int8('F'),
	},
	1828: {
		Fword:  __ccgo_ts + 12633,
		Ftype1: int8('F'),
	},
	1829: {
		Fword:  __ccgo_ts + 12640,
		Ftype1: int8('F'),
	},
	1830: {
		Fword:  __ccgo_ts + 12647,
		Ftype1: int8('F'),
	},
	1831: {
		Fword:  __ccgo_ts + 12654,
		Ftype1: int8('F'),
	},
	1832: {
		Fword:  __ccgo_ts + 12661,
		Ftype1: int8('F'),
	},
	1833: {
		Fword:  __ccgo_ts + 12668,
		Ftype1: int8('F'),
	},
	1834: {
		Fword:  __ccgo_ts + 12675,
		Ftype1: int8('F'),
	},
	1835: {
		Fword:  __ccgo_ts + 12682,
		Ftype1: int8('F'),
	},
	1836: {
		Fword:  __ccgo_ts + 12689,
		Ftype1: int8('F'),
	},
	1837: {
		Fword:  __ccgo_ts + 12696,
		Ftype1: int8('F'),
	},
	1838: {
		Fword:  __ccgo_ts + 12703,
		Ftype1: int8('F'),
	},
	1839: {
		Fword:  __ccgo_ts + 12710,
		Ftype1: int8('F'),
	},
	1840: {
		Fword:  __ccgo_ts + 12717,
		Ftype1: int8('F'),
	},
	1841: {
		Fword:  __ccgo_ts + 12724,
		Ftype1: int8('F'),
	},
	1842: {
		Fword:  __ccgo_ts + 12731,
		Ftype1: int8('F'),
	},
	1843: {
		Fword:  __ccgo_ts + 12738,
		Ftype1: int8('F'),
	},
	1844: {
		Fword:  __ccgo_ts + 12745,
		Ftype1: int8('F'),
	},
	1845: {
		Fword:  __ccgo_ts + 12752,
		Ftype1: int8('F'),
	},
	1846: {
		Fword:  __ccgo_ts + 12759,
		Ftype1: int8('F'),
	},
	1847: {
		Fword:  __ccgo_ts + 12766,
		Ftype1: int8('F'),
	},
	1848: {
		Fword:  __ccgo_ts + 12773,
		Ftype1: int8('F'),
	},
	1849: {
		Fword:  __ccgo_ts + 12780,
		Ftype1: int8('F'),
	},
	1850: {
		Fword:  __ccgo_ts + 12787,
		Ftype1: int8('F'),
	},
	1851: {
		Fword:  __ccgo_ts + 12794,
		Ftype1: int8('F'),
	},
	1852: {
		Fword:  __ccgo_ts + 12801,
		Ftype1: int8('F'),
	},
	1853: {
		Fword:  __ccgo_ts + 12808,
		Ftype1: int8('F'),
	},
	1854: {
		Fword:  __ccgo_ts + 12815,
		Ftype1: int8('F'),
	},
	1855: {
		Fword:  __ccgo_ts + 12822,
		Ftype1: int8('F'),
	},
	1856: {
		Fword:  __ccgo_ts + 12829,
		Ftype1: int8('F'),
	},
	1857: {
		Fword:  __ccgo_ts + 12836,
		Ftype1: int8('F'),
	},
	1858: {
		Fword:  __ccgo_ts + 12843,
		Ftype1: int8('F'),
	},
	1859: {
		Fword:  __ccgo_ts + 12850,
		Ftype1: int8('F'),
	},
	1860: {
		Fword:  __ccgo_ts + 12857,
		Ftype1: int8('F'),
	},
	1861: {
		Fword:  __ccgo_ts + 12864,
		Ftype1: int8('F'),
	},
	1862: {
		Fword:  __ccgo_ts + 12871,
		Ftype1: int8('F'),
	},
	1863: {
		Fword:  __ccgo_ts + 12878,
		Ftype1: int8('F'),
	},
	1864: {
		Fword:  __ccgo_ts + 12885,
		Ftype1: int8('F'),
	},
	1865: {
		Fword:  __ccgo_ts + 12892,
		Ftype1: int8('F'),
	},
	1866: {
		Fword:  __ccgo_ts + 12899,
		Ftype1: int8('F'),
	},
	1867: {
		Fword:  __ccgo_ts + 12906,
		Ftype1: int8('F'),
	},
	1868: {
		Fword:  __ccgo_ts + 12913,
		Ftype1: int8('F'),
	},
	1869: {
		Fword:  __ccgo_ts + 12920,
		Ftype1: int8('F'),
	},
	1870: {
		Fword:  __ccgo_ts + 12927,
		Ftype1: int8('F'),
	},
	1871: {
		Fword:  __ccgo_ts + 12934,
		Ftype1: int8('F'),
	},
	1872: {
		Fword:  __ccgo_ts + 12941,
		Ftype1: int8('F'),
	},
	1873: {
		Fword:  __ccgo_ts + 12948,
		Ftype1: int8('F'),
	},
	1874: {
		Fword:  __ccgo_ts + 12955,
		Ftype1: int8('F'),
	},
	1875: {
		Fword:  __ccgo_ts + 12962,
		Ftype1: int8('F'),
	},
	1876: {
		Fword:  __ccgo_ts + 12969,
		Ftype1: int8('F'),
	},
	1877: {
		Fword:  __ccgo_ts + 12976,
		Ftype1: int8('F'),
	},
	1878: {
		Fword:  __ccgo_ts + 12983,
		Ftype1: int8('F'),
	},
	1879: {
		Fword:  __ccgo_ts + 12990,
		Ftype1: int8('F'),
	},
	1880: {
		Fword:  __ccgo_ts + 12997,
		Ftype1: int8('F'),
	},
	1881: {
		Fword:  __ccgo_ts + 13004,
		Ftype1: int8('F'),
	},
	1882: {
		Fword:  __ccgo_ts + 13011,
		Ftype1: int8('F'),
	},
	1883: {
		Fword:  __ccgo_ts + 13018,
		Ftype1: int8('F'),
	},
	1884: {
		Fword:  __ccgo_ts + 13025,
		Ftype1: int8('F'),
	},
	1885: {
		Fword:  __ccgo_ts + 13032,
		Ftype1: int8('F'),
	},
	1886: {
		Fword:  __ccgo_ts + 13039,
		Ftype1: int8('F'),
	},
	1887: {
		Fword:  __ccgo_ts + 13046,
		Ftype1: int8('F'),
	},
	1888: {
		Fword:  __ccgo_ts + 13053,
		Ftype1: int8('F'),
	},
	1889: {
		Fword:  __ccgo_ts + 13060,
		Ftype1: int8('F'),
	},
	1890: {
		Fword:  __ccgo_ts + 13067,
		Ftype1: int8('F'),
	},
	1891: {
		Fword:  __ccgo_ts + 13074,
		Ftype1: int8('F'),
	},
	1892: {
		Fword:  __ccgo_ts + 13081,
		Ftype1: int8('F'),
	},
	1893: {
		Fword:  __ccgo_ts + 13088,
		Ftype1: int8('F'),
	},
	1894: {
		Fword:  __ccgo_ts + 13095,
		Ftype1: int8('F'),
	},
	1895: {
		Fword:  __ccgo_ts + 13102,
		Ftype1: int8('F'),
	},
	1896: {
		Fword:  __ccgo_ts + 13109,
		Ftype1: int8('F'),
	},
	1897: {
		Fword:  __ccgo_ts + 13116,
		Ftype1: int8('F'),
	},
	1898: {
		Fword:  __ccgo_ts + 13123,
		Ftype1: int8('F'),
	},
	1899: {
		Fword:  __ccgo_ts + 13130,
		Ftype1: int8('F'),
	},
	1900: {
		Fword:  __ccgo_ts + 13137,
		Ftype1: int8('F'),
	},
	1901: {
		Fword:  __ccgo_ts + 13144,
		Ftype1: int8('F'),
	},
	1902: {
		Fword:  __ccgo_ts + 13151,
		Ftype1: int8('F'),
	},
	1903: {
		Fword:  __ccgo_ts + 13158,
		Ftype1: int8('F'),
	},
	1904: {
		Fword:  __ccgo_ts + 13165,
		Ftype1: int8('F'),
	},
	1905: {
		Fword:  __ccgo_ts + 13172,
		Ftype1: int8('F'),
	},
	1906: {
		Fword:  __ccgo_ts + 13179,
		Ftype1: int8('F'),
	},
	1907: {
		Fword:  __ccgo_ts + 13186,
		Ftype1: int8('F'),
	},
	1908: {
		Fword:  __ccgo_ts + 13193,
		Ftype1: int8('F'),
	},
	1909: {
		Fword:  __ccgo_ts + 13200,
		Ftype1: int8('F'),
	},
	1910: {
		Fword:  __ccgo_ts + 13207,
		Ftype1: int8('F'),
	},
	1911: {
		Fword:  __ccgo_ts + 13214,
		Ftype1: int8('F'),
	},
	1912: {
		Fword:  __ccgo_ts + 13221,
		Ftype1: int8('F'),
	},
	1913: {
		Fword:  __ccgo_ts + 13228,
		Ftype1: int8('F'),
	},
	1914: {
		Fword:  __ccgo_ts + 13235,
		Ftype1: int8('F'),
	},
	1915: {
		Fword:  __ccgo_ts + 13242,
		Ftype1: int8('F'),
	},
	1916: {
		Fword:  __ccgo_ts + 13249,
		Ftype1: int8('F'),
	},
	1917: {
		Fword:  __ccgo_ts + 13256,
		Ftype1: int8('F'),
	},
	1918: {
		Fword:  __ccgo_ts + 13263,
		Ftype1: int8('F'),
	},
	1919: {
		Fword:  __ccgo_ts + 13270,
		Ftype1: int8('F'),
	},
	1920: {
		Fword:  __ccgo_ts + 13277,
		Ftype1: int8('F'),
	},
	1921: {
		Fword:  __ccgo_ts + 13284,
		Ftype1: int8('F'),
	},
	1922: {
		Fword:  __ccgo_ts + 13291,
		Ftype1: int8('F'),
	},
	1923: {
		Fword:  __ccgo_ts + 13298,
		Ftype1: int8('F'),
	},
	1924: {
		Fword:  __ccgo_ts + 13305,
		Ftype1: int8('F'),
	},
	1925: {
		Fword:  __ccgo_ts + 13312,
		Ftype1: int8('F'),
	},
	1926: {
		Fword:  __ccgo_ts + 13319,
		Ftype1: int8('F'),
	},
	1927: {
		Fword:  __ccgo_ts + 13326,
		Ftype1: int8('F'),
	},
	1928: {
		Fword:  __ccgo_ts + 13333,
		Ftype1: int8('F'),
	},
	1929: {
		Fword:  __ccgo_ts + 13340,
		Ftype1: int8('F'),
	},
	1930: {
		Fword:  __ccgo_ts + 13347,
		Ftype1: int8('F'),
	},
	1931: {
		Fword:  __ccgo_ts + 13354,
		Ftype1: int8('F'),
	},
	1932: {
		Fword:  __ccgo_ts + 13361,
		Ftype1: int8('F'),
	},
	1933: {
		Fword:  __ccgo_ts + 13368,
		Ftype1: int8('F'),
	},
	1934: {
		Fword:  __ccgo_ts + 13375,
		Ftype1: int8('F'),
	},
	1935: {
		Fword:  __ccgo_ts + 13382,
		Ftype1: int8('F'),
	},
	1936: {
		Fword:  __ccgo_ts + 13389,
		Ftype1: int8('F'),
	},
	1937: {
		Fword:  __ccgo_ts + 13396,
		Ftype1: int8('F'),
	},
	1938: {
		Fword:  __ccgo_ts + 13403,
		Ftype1: int8('F'),
	},
	1939: {
		Fword:  __ccgo_ts + 13410,
		Ftype1: int8('F'),
	},
	1940: {
		Fword:  __ccgo_ts + 13417,
		Ftype1: int8('F'),
	},
	1941: {
		Fword:  __ccgo_ts + 13424,
		Ftype1: int8('F'),
	},
	1942: {
		Fword:  __ccgo_ts + 13431,
		Ftype1: int8('F'),
	},
	1943: {
		Fword:  __ccgo_ts + 13438,
		Ftype1: int8('F'),
	},
	1944: {
		Fword:  __ccgo_ts + 13445,
		Ftype1: int8('F'),
	},
	1945: {
		Fword:  __ccgo_ts + 13452,
		Ftype1: int8('F'),
	},
	1946: {
		Fword:  __ccgo_ts + 13459,
		Ftype1: int8('F'),
	},
	1947: {
		Fword:  __ccgo_ts + 13466,
		Ftype1: int8('F'),
	},
	1948: {
		Fword:  __ccgo_ts + 13473,
		Ftype1: int8('F'),
	},
	1949: {
		Fword:  __ccgo_ts + 13480,
		Ftype1: int8('F'),
	},
	1950: {
		Fword:  __ccgo_ts + 13487,
		Ftype1: int8('F'),
	},
	1951: {
		Fword:  __ccgo_ts + 13494,
		Ftype1: int8('F'),
	},
	1952: {
		Fword:  __ccgo_ts + 13501,
		Ftype1: int8('F'),
	},
	1953: {
		Fword:  __ccgo_ts + 13508,
		Ftype1: int8('F'),
	},
	1954: {
		Fword:  __ccgo_ts + 13515,
		Ftype1: int8('F'),
	},
	1955: {
		Fword:  __ccgo_ts + 13522,
		Ftype1: int8('F'),
	},
	1956: {
		Fword:  __ccgo_ts + 13529,
		Ftype1: int8('F'),
	},
	1957: {
		Fword:  __ccgo_ts + 13536,
		Ftype1: int8('F'),
	},
	1958: {
		Fword:  __ccgo_ts + 13543,
		Ftype1: int8('F'),
	},
	1959: {
		Fword:  __ccgo_ts + 13550,
		Ftype1: int8('F'),
	},
	1960: {
		Fword:  __ccgo_ts + 13557,
		Ftype1: int8('F'),
	},
	1961: {
		Fword:  __ccgo_ts + 13564,
		Ftype1: int8('F'),
	},
	1962: {
		Fword:  __ccgo_ts + 13571,
		Ftype1: int8('F'),
	},
	1963: {
		Fword:  __ccgo_ts + 13578,
		Ftype1: int8('F'),
	},
	1964: {
		Fword:  __ccgo_ts + 13585,
		Ftype1: int8('F'),
	},
	1965: {
		Fword:  __ccgo_ts + 13592,
		Ftype1: int8('F'),
	},
	1966: {
		Fword:  __ccgo_ts + 13599,
		Ftype1: int8('F'),
	},
	1967: {
		Fword:  __ccgo_ts + 13606,
		Ftype1: int8('F'),
	},
	1968: {
		Fword:  __ccgo_ts + 13613,
		Ftype1: int8('F'),
	},
	1969: {
		Fword:  __ccgo_ts + 13620,
		Ftype1: int8('F'),
	},
	1970: {
		Fword:  __ccgo_ts + 13627,
		Ftype1: int8('F'),
	},
	1971: {
		Fword:  __ccgo_ts + 13634,
		Ftype1: int8('F'),
	},
	1972: {
		Fword:  __ccgo_ts + 13641,
		Ftype1: int8('F'),
	},
	1973: {
		Fword:  __ccgo_ts + 13648,
		Ftype1: int8('F'),
	},
	1974: {
		Fword:  __ccgo_ts + 13655,
		Ftype1: int8('F'),
	},
	1975: {
		Fword:  __ccgo_ts + 13662,
		Ftype1: int8('F'),
	},
	1976: {
		Fword:  __ccgo_ts + 13669,
		Ftype1: int8('F'),
	},
	1977: {
		Fword:  __ccgo_ts + 13676,
		Ftype1: int8('F'),
	},
	1978: {
		Fword:  __ccgo_ts + 13683,
		Ftype1: int8('F'),
	},
	1979: {
		Fword:  __ccgo_ts + 13690,
		Ftype1: int8('F'),
	},
	1980: {
		Fword:  __ccgo_ts + 13697,
		Ftype1: int8('F'),
	},
	1981: {
		Fword:  __ccgo_ts + 13704,
		Ftype1: int8('F'),
	},
	1982: {
		Fword:  __ccgo_ts + 13709,
		Ftype1: int8('F'),
	},
	1983: {
		Fword:  __ccgo_ts + 13716,
		Ftype1: int8('F'),
	},
	1984: {
		Fword:  __ccgo_ts + 13723,
		Ftype1: int8('F'),
	},
	1985: {
		Fword:  __ccgo_ts + 13730,
		Ftype1: int8('F'),
	},
	1986: {
		Fword:  __ccgo_ts + 13737,
		Ftype1: int8('F'),
	},
	1987: {
		Fword:  __ccgo_ts + 13744,
		Ftype1: int8('F'),
	},
	1988: {
		Fword:  __ccgo_ts + 13751,
		Ftype1: int8('F'),
	},
	1989: {
		Fword:  __ccgo_ts + 13757,
		Ftype1: int8('F'),
	},
	1990: {
		Fword:  __ccgo_ts + 13764,
		Ftype1: int8('F'),
	},
	1991: {
		Fword:  __ccgo_ts + 13771,
		Ftype1: int8('F'),
	},
	1992: {
		Fword:  __ccgo_ts + 13778,
		Ftype1: int8('F'),
	},
	1993: {
		Fword:  __ccgo_ts + 13785,
		Ftype1: int8('F'),
	},
	1994: {
		Fword:  __ccgo_ts + 13791,
		Ftype1: int8('F'),
	},
	1995: {
		Fword:  __ccgo_ts + 13798,
		Ftype1: int8('F'),
	},
	1996: {
		Fword:  __ccgo_ts + 13805,
		Ftype1: int8('F'),
	},
	1997: {
		Fword:  __ccgo_ts + 13812,
		Ftype1: int8('F'),
	},
	1998: {
		Fword:  __ccgo_ts + 13819,
		Ftype1: int8('F'),
	},
	1999: {
		Fword:  __ccgo_ts + 13826,
		Ftype1: int8('F'),
	},
	2000: {
		Fword:  __ccgo_ts + 13833,
		Ftype1: int8('F'),
	},
	2001: {
		Fword:  __ccgo_ts + 13840,
		Ftype1: int8('F'),
	},
	2002: {
		Fword:  __ccgo_ts + 13847,
		Ftype1: int8('F'),
	},
	2003: {
		Fword:  __ccgo_ts + 13854,
		Ftype1: int8('F'),
	},
	2004: {
		Fword:  __ccgo_ts + 13861,
		Ftype1: int8('F'),
	},
	2005: {
		Fword:  __ccgo_ts + 13868,
		Ftype1: int8('F'),
	},
	2006: {
		Fword:  __ccgo_ts + 13875,
		Ftype1: int8('F'),
	},
	2007: {
		Fword:  __ccgo_ts + 13882,
		Ftype1: int8('F'),
	},
	2008: {
		Fword:  __ccgo_ts + 13889,
		Ftype1: int8('F'),
	},
	2009: {
		Fword:  __ccgo_ts + 13896,
		Ftype1: int8('F'),
	},
	2010: {
		Fword:  __ccgo_ts + 13903,
		Ftype1: int8('F'),
	},
	2011: {
		Fword:  __ccgo_ts + 13910,
		Ftype1: int8('F'),
	},
	2012: {
		Fword:  __ccgo_ts + 13917,
		Ftype1: int8('F'),
	},
	2013: {
		Fword:  __ccgo_ts + 13924,
		Ftype1: int8('F'),
	},
	2014: {
		Fword:  __ccgo_ts + 13931,
		Ftype1: int8('F'),
	},
	2015: {
		Fword:  __ccgo_ts + 13938,
		Ftype1: int8('F'),
	},
	2016: {
		Fword:  __ccgo_ts + 13945,
		Ftype1: int8('F'),
	},
	2017: {
		Fword:  __ccgo_ts + 13952,
		Ftype1: int8('F'),
	},
	2018: {
		Fword:  __ccgo_ts + 13959,
		Ftype1: int8('F'),
	},
	2019: {
		Fword:  __ccgo_ts + 13966,
		Ftype1: int8('F'),
	},
	2020: {
		Fword:  __ccgo_ts + 13973,
		Ftype1: int8('F'),
	},
	2021: {
		Fword:  __ccgo_ts + 13980,
		Ftype1: int8('F'),
	},
	2022: {
		Fword:  __ccgo_ts + 13987,
		Ftype1: int8('F'),
	},
	2023: {
		Fword:  __ccgo_ts + 13994,
		Ftype1: int8('F'),
	},
	2024: {
		Fword:  __ccgo_ts + 14001,
		Ftype1: int8('F'),
	},
	2025: {
		Fword:  __ccgo_ts + 14008,
		Ftype1: int8('F'),
	},
	2026: {
		Fword:  __ccgo_ts + 14015,
		Ftype1: int8('F'),
	},
	2027: {
		Fword:  __ccgo_ts + 14022,
		Ftype1: int8('F'),
	},
	2028: {
		Fword:  __ccgo_ts + 14029,
		Ftype1: int8('F'),
	},
	2029: {
		Fword:  __ccgo_ts + 14036,
		Ftype1: int8('F'),
	},
	2030: {
		Fword:  __ccgo_ts + 14043,
		Ftype1: int8('F'),
	},
	2031: {
		Fword:  __ccgo_ts + 14050,
		Ftype1: int8('F'),
	},
	2032: {
		Fword:  __ccgo_ts + 14057,
		Ftype1: int8('F'),
	},
	2033: {
		Fword:  __ccgo_ts + 14064,
		Ftype1: int8('F'),
	},
	2034: {
		Fword:  __ccgo_ts + 14071,
		Ftype1: int8('F'),
	},
	2035: {
		Fword:  __ccgo_ts + 14078,
		Ftype1: int8('F'),
	},
	2036: {
		Fword:  __ccgo_ts + 14085,
		Ftype1: int8('F'),
	},
	2037: {
		Fword:  __ccgo_ts + 14092,
		Ftype1: int8('F'),
	},
	2038: {
		Fword:  __ccgo_ts + 14099,
		Ftype1: int8('F'),
	},
	2039: {
		Fword:  __ccgo_ts + 14106,
		Ftype1: int8('F'),
	},
	2040: {
		Fword:  __ccgo_ts + 14112,
		Ftype1: int8('F'),
	},
	2041: {
		Fword:  __ccgo_ts + 14119,
		Ftype1: int8('F'),
	},
	2042: {
		Fword:  __ccgo_ts + 14126,
		Ftype1: int8('F'),
	},
	2043: {
		Fword:  __ccgo_ts + 14133,
		Ftype1: int8('F'),
	},
	2044: {
		Fword:  __ccgo_ts + 14140,
		Ftype1: int8('F'),
	},
	2045: {
		Fword:  __ccgo_ts + 14147,
		Ftype1: int8('F'),
	},
	2046: {
		Fword:  __ccgo_ts + 14154,
		Ftype1: int8('F'),
	},
	2047: {
		Fword:  __ccgo_ts + 14161,
		Ftype1: int8('F'),
	},
	2048: {
		Fword:  __ccgo_ts + 14168,
		Ftype1: int8('F'),
	},
	2049: {
		Fword:  __ccgo_ts + 14175,
		Ftype1: int8('F'),
	},
	2050: {
		Fword:  __ccgo_ts + 14182,
		Ftype1: int8('F'),
	},
	2051: {
		Fword:  __ccgo_ts + 14189,
		Ftype1: int8('F'),
	},
	2052: {
		Fword:  __ccgo_ts + 14196,
		Ftype1: int8('F'),
	},
	2053: {
		Fword:  __ccgo_ts + 14202,
		Ftype1: int8('F'),
	},
	2054: {
		Fword:  __ccgo_ts + 14209,
		Ftype1: int8('F'),
	},
	2055: {
		Fword:  __ccgo_ts + 14215,
		Ftype1: int8('F'),
	},
	2056: {
		Fword:  __ccgo_ts + 14222,
		Ftype1: int8('F'),
	},
	2057: {
		Fword:  __ccgo_ts + 14229,
		Ftype1: int8('F'),
	},
	2058: {
		Fword:  __ccgo_ts + 14236,
		Ftype1: int8('F'),
	},
	2059: {
		Fword:  __ccgo_ts + 14243,
		Ftype1: int8('F'),
	},
	2060: {
		Fword:  __ccgo_ts + 14250,
		Ftype1: int8('F'),
	},
	2061: {
		Fword:  __ccgo_ts + 14257,
		Ftype1: int8('F'),
	},
	2062: {
		Fword:  __ccgo_ts + 14264,
		Ftype1: int8('F'),
	},
	2063: {
		Fword:  __ccgo_ts + 14271,
		Ftype1: int8('F'),
	},
	2064: {
		Fword:  __ccgo_ts + 14278,
		Ftype1: int8('F'),
	},
	2065: {
		Fword:  __ccgo_ts + 14285,
		Ftype1: int8('F'),
	},
	2066: {
		Fword:  __ccgo_ts + 14292,
		Ftype1: int8('F'),
	},
	2067: {
		Fword:  __ccgo_ts + 14299,
		Ftype1: int8('F'),
	},
	2068: {
		Fword:  __ccgo_ts + 14306,
		Ftype1: int8('F'),
	},
	2069: {
		Fword:  __ccgo_ts + 14313,
		Ftype1: int8('F'),
	},
	2070: {
		Fword:  __ccgo_ts + 14320,
		Ftype1: int8('F'),
	},
	2071: {
		Fword:  __ccgo_ts + 14327,
		Ftype1: int8('F'),
	},
	2072: {
		Fword:  __ccgo_ts + 14334,
		Ftype1: int8('F'),
	},
	2073: {
		Fword:  __ccgo_ts + 14341,
		Ftype1: int8('F'),
	},
	2074: {
		Fword:  __ccgo_ts + 14348,
		Ftype1: int8('F'),
	},
	2075: {
		Fword:  __ccgo_ts + 14355,
		Ftype1: int8('F'),
	},
	2076: {
		Fword:  __ccgo_ts + 14362,
		Ftype1: int8('F'),
	},
	2077: {
		Fword:  __ccgo_ts + 14369,
		Ftype1: int8('F'),
	},
	2078: {
		Fword:  __ccgo_ts + 14376,
		Ftype1: int8('F'),
	},
	2079: {
		Fword:  __ccgo_ts + 14383,
		Ftype1: int8('F'),
	},
	2080: {
		Fword:  __ccgo_ts + 14390,
		Ftype1: int8('F'),
	},
	2081: {
		Fword:  __ccgo_ts + 14397,
		Ftype1: int8('F'),
	},
	2082: {
		Fword:  __ccgo_ts + 14404,
		Ftype1: int8('F'),
	},
	2083: {
		Fword:  __ccgo_ts + 14411,
		Ftype1: int8('F'),
	},
	2084: {
		Fword:  __ccgo_ts + 14418,
		Ftype1: int8('F'),
	},
	2085: {
		Fword:  __ccgo_ts + 14425,
		Ftype1: int8('F'),
	},
	2086: {
		Fword:  __ccgo_ts + 14431,
		Ftype1: int8('F'),
	},
	2087: {
		Fword:  __ccgo_ts + 14438,
		Ftype1: int8('F'),
	},
	2088: {
		Fword:  __ccgo_ts + 14444,
		Ftype1: int8('F'),
	},
	2089: {
		Fword:  __ccgo_ts + 14451,
		Ftype1: int8('F'),
	},
	2090: {
		Fword:  __ccgo_ts + 14458,
		Ftype1: int8('F'),
	},
	2091: {
		Fword:  __ccgo_ts + 14465,
		Ftype1: int8('F'),
	},
	2092: {
		Fword:  __ccgo_ts + 14472,
		Ftype1: int8('F'),
	},
	2093: {
		Fword:  __ccgo_ts + 14479,
		Ftype1: int8('F'),
	},
	2094: {
		Fword:  __ccgo_ts + 14486,
		Ftype1: int8('F'),
	},
	2095: {
		Fword:  __ccgo_ts + 14493,
		Ftype1: int8('F'),
	},
	2096: {
		Fword:  __ccgo_ts + 14500,
		Ftype1: int8('F'),
	},
	2097: {
		Fword:  __ccgo_ts + 14505,
		Ftype1: int8('F'),
	},
	2098: {
		Fword:  __ccgo_ts + 14511,
		Ftype1: int8('F'),
	},
	2099: {
		Fword:  __ccgo_ts + 14518,
		Ftype1: int8('F'),
	},
	2100: {
		Fword:  __ccgo_ts + 14524,
		Ftype1: int8('F'),
	},
	2101: {
		Fword:  __ccgo_ts + 14531,
		Ftype1: int8('F'),
	},
	2102: {
		Fword:  __ccgo_ts + 14538,
		Ftype1: int8('F'),
	},
	2103: {
		Fword:  __ccgo_ts + 14545,
		Ftype1: int8('F'),
	},
	2104: {
		Fword:  __ccgo_ts + 14552,
		Ftype1: int8('F'),
	},
	2105: {
		Fword:  __ccgo_ts + 14559,
		Ftype1: int8('F'),
	},
	2106: {
		Fword:  __ccgo_ts + 14566,
		Ftype1: int8('F'),
	},
	2107: {
		Fword:  __ccgo_ts + 14573,
		Ftype1: int8('F'),
	},
	2108: {
		Fword:  __ccgo_ts + 14580,
		Ftype1: int8('F'),
	},
	2109: {
		Fword:  __ccgo_ts + 14587,
		Ftype1: int8('F'),
	},
	2110: {
		Fword:  __ccgo_ts + 14594,
		Ftype1: int8('F'),
	},
	2111: {
		Fword:  __ccgo_ts + 14601,
		Ftype1: int8('F'),
	},
	2112: {
		Fword:  __ccgo_ts + 14608,
		Ftype1: int8('F'),
	},
	2113: {
		Fword:  __ccgo_ts + 14615,
		Ftype1: int8('F'),
	},
	2114: {
		Fword:  __ccgo_ts + 14622,
		Ftype1: int8('F'),
	},
	2115: {
		Fword:  __ccgo_ts + 14629,
		Ftype1: int8('F'),
	},
	2116: {
		Fword:  __ccgo_ts + 14636,
		Ftype1: int8('F'),
	},
	2117: {
		Fword:  __ccgo_ts + 14643,
		Ftype1: int8('F'),
	},
	2118: {
		Fword:  __ccgo_ts + 14650,
		Ftype1: int8('F'),
	},
	2119: {
		Fword:  __ccgo_ts + 14657,
		Ftype1: int8('F'),
	},
	2120: {
		Fword:  __ccgo_ts + 14664,
		Ftype1: int8('F'),
	},
	2121: {
		Fword:  __ccgo_ts + 14671,
		Ftype1: int8('F'),
	},
	2122: {
		Fword:  __ccgo_ts + 14678,
		Ftype1: int8('F'),
	},
	2123: {
		Fword:  __ccgo_ts + 14685,
		Ftype1: int8('F'),
	},
	2124: {
		Fword:  __ccgo_ts + 14692,
		Ftype1: int8('F'),
	},
	2125: {
		Fword:  __ccgo_ts + 14699,
		Ftype1: int8('F'),
	},
	2126: {
		Fword:  __ccgo_ts + 14706,
		Ftype1: int8('F'),
	},
	2127: {
		Fword:  __ccgo_ts + 14713,
		Ftype1: int8('F'),
	},
	2128: {
		Fword:  __ccgo_ts + 14720,
		Ftype1: int8('F'),
	},
	2129: {
		Fword:  __ccgo_ts + 14727,
		Ftype1: int8('F'),
	},
	2130: {
		Fword:  __ccgo_ts + 14734,
		Ftype1: int8('F'),
	},
	2131: {
		Fword:  __ccgo_ts + 14741,
		Ftype1: int8('F'),
	},
	2132: {
		Fword:  __ccgo_ts + 14748,
		Ftype1: int8('F'),
	},
	2133: {
		Fword:  __ccgo_ts + 14755,
		Ftype1: int8('F'),
	},
	2134: {
		Fword:  __ccgo_ts + 14762,
		Ftype1: int8('F'),
	},
	2135: {
		Fword:  __ccgo_ts + 14769,
		Ftype1: int8('F'),
	},
	2136: {
		Fword:  __ccgo_ts + 14776,
		Ftype1: int8('F'),
	},
	2137: {
		Fword:  __ccgo_ts + 14783,
		Ftype1: int8('F'),
	},
	2138: {
		Fword:  __ccgo_ts + 14790,
		Ftype1: int8('F'),
	},
	2139: {
		Fword:  __ccgo_ts + 14797,
		Ftype1: int8('F'),
	},
	2140: {
		Fword:  __ccgo_ts + 14804,
		Ftype1: int8('F'),
	},
	2141: {
		Fword:  __ccgo_ts + 14811,
		Ftype1: int8('F'),
	},
	2142: {
		Fword:  __ccgo_ts + 14818,
		Ftype1: int8('F'),
	},
	2143: {
		Fword:  __ccgo_ts + 14825,
		Ftype1: int8('F'),
	},
	2144: {
		Fword:  __ccgo_ts + 14832,
		Ftype1: int8('F'),
	},
	2145: {
		Fword:  __ccgo_ts + 14839,
		Ftype1: int8('F'),
	},
	2146: {
		Fword:  __ccgo_ts + 14846,
		Ftype1: int8('F'),
	},
	2147: {
		Fword:  __ccgo_ts + 14853,
		Ftype1: int8('F'),
	},
	2148: {
		Fword:  __ccgo_ts + 14860,
		Ftype1: int8('F'),
	},
	2149: {
		Fword:  __ccgo_ts + 14867,
		Ftype1: int8('F'),
	},
	2150: {
		Fword:  __ccgo_ts + 14874,
		Ftype1: int8('F'),
	},
	2151: {
		Fword:  __ccgo_ts + 14881,
		Ftype1: int8('F'),
	},
	2152: {
		Fword:  __ccgo_ts + 14888,
		Ftype1: int8('F'),
	},
	2153: {
		Fword:  __ccgo_ts + 14895,
		Ftype1: int8('F'),
	},
	2154: {
		Fword:  __ccgo_ts + 14902,
		Ftype1: int8('F'),
	},
	2155: {
		Fword:  __ccgo_ts + 14909,
		Ftype1: int8('F'),
	},
	2156: {
		Fword:  __ccgo_ts + 14916,
		Ftype1: int8('F'),
	},
	2157: {
		Fword:  __ccgo_ts + 14923,
		Ftype1: int8('F'),
	},
	2158: {
		Fword:  __ccgo_ts + 14930,
		Ftype1: int8('F'),
	},
	2159: {
		Fword:  __ccgo_ts + 14937,
		Ftype1: int8('F'),
	},
	2160: {
		Fword:  __ccgo_ts + 14944,
		Ftype1: int8('F'),
	},
	2161: {
		Fword:  __ccgo_ts + 14951,
		Ftype1: int8('F'),
	},
	2162: {
		Fword:  __ccgo_ts + 14958,
		Ftype1: int8('F'),
	},
	2163: {
		Fword:  __ccgo_ts + 14965,
		Ftype1: int8('F'),
	},
	2164: {
		Fword:  __ccgo_ts + 14972,
		Ftype1: int8('F'),
	},
	2165: {
		Fword:  __ccgo_ts + 14979,
		Ftype1: int8('F'),
	},
	2166: {
		Fword:  __ccgo_ts + 14986,
		Ftype1: int8('F'),
	},
	2167: {
		Fword:  __ccgo_ts + 14993,
		Ftype1: int8('F'),
	},
	2168: {
		Fword:  __ccgo_ts + 15000,
		Ftype1: int8('F'),
	},
	2169: {
		Fword:  __ccgo_ts + 15007,
		Ftype1: int8('F'),
	},
	2170: {
		Fword:  __ccgo_ts + 15014,
		Ftype1: int8('F'),
	},
	2171: {
		Fword:  __ccgo_ts + 15021,
		Ftype1: int8('F'),
	},
	2172: {
		Fword:  __ccgo_ts + 15028,
		Ftype1: int8('F'),
	},
	2173: {
		Fword:  __ccgo_ts + 15035,
		Ftype1: int8('F'),
	},
	2174: {
		Fword:  __ccgo_ts + 15042,
		Ftype1: int8('F'),
	},
	2175: {
		Fword:  __ccgo_ts + 15049,
		Ftype1: int8('F'),
	},
	2176: {
		Fword:  __ccgo_ts + 15056,
		Ftype1: int8('F'),
	},
	2177: {
		Fword:  __ccgo_ts + 15063,
		Ftype1: int8('F'),
	},
	2178: {
		Fword:  __ccgo_ts + 15070,
		Ftype1: int8('F'),
	},
	2179: {
		Fword:  __ccgo_ts + 15077,
		Ftype1: int8('F'),
	},
	2180: {
		Fword:  __ccgo_ts + 15084,
		Ftype1: int8('F'),
	},
	2181: {
		Fword:  __ccgo_ts + 15091,
		Ftype1: int8('F'),
	},
	2182: {
		Fword:  __ccgo_ts + 15098,
		Ftype1: int8('F'),
	},
	2183: {
		Fword:  __ccgo_ts + 15105,
		Ftype1: int8('F'),
	},
	2184: {
		Fword:  __ccgo_ts + 15112,
		Ftype1: int8('F'),
	},
	2185: {
		Fword:  __ccgo_ts + 15119,
		Ftype1: int8('F'),
	},
	2186: {
		Fword:  __ccgo_ts + 15126,
		Ftype1: int8('F'),
	},
	2187: {
		Fword:  __ccgo_ts + 15133,
		Ftype1: int8('F'),
	},
	2188: {
		Fword:  __ccgo_ts + 15139,
		Ftype1: int8('F'),
	},
	2189: {
		Fword:  __ccgo_ts + 15146,
		Ftype1: int8('F'),
	},
	2190: {
		Fword:  __ccgo_ts + 15153,
		Ftype1: int8('F'),
	},
	2191: {
		Fword:  __ccgo_ts + 15160,
		Ftype1: int8('F'),
	},
	2192: {
		Fword:  __ccgo_ts + 15167,
		Ftype1: int8('F'),
	},
	2193: {
		Fword:  __ccgo_ts + 15174,
		Ftype1: int8('F'),
	},
	2194: {
		Fword:  __ccgo_ts + 15181,
		Ftype1: int8('F'),
	},
	2195: {
		Fword:  __ccgo_ts + 15188,
		Ftype1: int8('F'),
	},
	2196: {
		Fword:  __ccgo_ts + 15195,
		Ftype1: int8('F'),
	},
	2197: {
		Fword:  __ccgo_ts + 15202,
		Ftype1: int8('F'),
	},
	2198: {
		Fword:  __ccgo_ts + 15209,
		Ftype1: int8('F'),
	},
	2199: {
		Fword:  __ccgo_ts + 15216,
		Ftype1: int8('F'),
	},
	2200: {
		Fword:  __ccgo_ts + 15223,
		Ftype1: int8('F'),
	},
	2201: {
		Fword:  __ccgo_ts + 15230,
		Ftype1: int8('F'),
	},
	2202: {
		Fword:  __ccgo_ts + 15237,
		Ftype1: int8('F'),
	},
	2203: {
		Fword:  __ccgo_ts + 15244,
		Ftype1: int8('F'),
	},
	2204: {
		Fword:  __ccgo_ts + 15251,
		Ftype1: int8('F'),
	},
	2205: {
		Fword:  __ccgo_ts + 15258,
		Ftype1: int8('F'),
	},
	2206: {
		Fword:  __ccgo_ts + 15265,
		Ftype1: int8('F'),
	},
	2207: {
		Fword:  __ccgo_ts + 15272,
		Ftype1: int8('F'),
	},
	2208: {
		Fword:  __ccgo_ts + 15279,
		Ftype1: int8('F'),
	},
	2209: {
		Fword:  __ccgo_ts + 15286,
		Ftype1: int8('F'),
	},
	2210: {
		Fword:  __ccgo_ts + 15293,
		Ftype1: int8('F'),
	},
	2211: {
		Fword:  __ccgo_ts + 15299,
		Ftype1: int8('F'),
	},
	2212: {
		Fword:  __ccgo_ts + 15306,
		Ftype1: int8('F'),
	},
	2213: {
		Fword:  __ccgo_ts + 15313,
		Ftype1: int8('F'),
	},
	2214: {
		Fword:  __ccgo_ts + 15320,
		Ftype1: int8('F'),
	},
	2215: {
		Fword:  __ccgo_ts + 15327,
		Ftype1: int8('F'),
	},
	2216: {
		Fword:  __ccgo_ts + 15334,
		Ftype1: int8('F'),
	},
	2217: {
		Fword:  __ccgo_ts + 15341,
		Ftype1: int8('F'),
	},
	2218: {
		Fword:  __ccgo_ts + 15348,
		Ftype1: int8('F'),
	},
	2219: {
		Fword:  __ccgo_ts + 15355,
		Ftype1: int8('F'),
	},
	2220: {
		Fword:  __ccgo_ts + 15361,
		Ftype1: int8('F'),
	},
	2221: {
		Fword:  __ccgo_ts + 15368,
		Ftype1: int8('F'),
	},
	2222: {
		Fword:  __ccgo_ts + 15375,
		Ftype1: int8('F'),
	},
	2223: {
		Fword:  __ccgo_ts + 15382,
		Ftype1: int8('F'),
	},
	2224: {
		Fword:  __ccgo_ts + 15389,
		Ftype1: int8('F'),
	},
	2225: {
		Fword:  __ccgo_ts + 15396,
		Ftype1: int8('F'),
	},
	2226: {
		Fword:  __ccgo_ts + 15403,
		Ftype1: int8('F'),
	},
	2227: {
		Fword:  __ccgo_ts + 15409,
		Ftype1: int8('F'),
	},
	2228: {
		Fword:  __ccgo_ts + 15416,
		Ftype1: int8('F'),
	},
	2229: {
		Fword:  __ccgo_ts + 15423,
		Ftype1: int8('F'),
	},
	2230: {
		Fword:  __ccgo_ts + 15430,
		Ftype1: int8('F'),
	},
	2231: {
		Fword:  __ccgo_ts + 15437,
		Ftype1: int8('F'),
	},
	2232: {
		Fword:  __ccgo_ts + 15444,
		Ftype1: int8('F'),
	},
	2233: {
		Fword:  __ccgo_ts + 15451,
		Ftype1: int8('F'),
	},
	2234: {
		Fword:  __ccgo_ts + 15458,
		Ftype1: int8('F'),
	},
	2235: {
		Fword:  __ccgo_ts + 15465,
		Ftype1: int8('F'),
	},
	2236: {
		Fword:  __ccgo_ts + 15472,
		Ftype1: int8('F'),
	},
	2237: {
		Fword:  __ccgo_ts + 15479,
		Ftype1: int8('F'),
	},
	2238: {
		Fword:  __ccgo_ts + 15486,
		Ftype1: int8('F'),
	},
	2239: {
		Fword:  __ccgo_ts + 15493,
		Ftype1: int8('F'),
	},
	2240: {
		Fword:  __ccgo_ts + 15500,
		Ftype1: int8('F'),
	},
	2241: {
		Fword:  __ccgo_ts + 15507,
		Ftype1: int8('F'),
	},
	2242: {
		Fword:  __ccgo_ts + 15514,
		Ftype1: int8('F'),
	},
	2243: {
		Fword:  __ccgo_ts + 15521,
		Ftype1: int8('F'),
	},
	2244: {
		Fword:  __ccgo_ts + 15528,
		Ftype1: int8('F'),
	},
	2245: {
		Fword:  __ccgo_ts + 15535,
		Ftype1: int8('F'),
	},
	2246: {
		Fword:  __ccgo_ts + 15542,
		Ftype1: int8('F'),
	},
	2247: {
		Fword:  __ccgo_ts + 15549,
		Ftype1: int8('F'),
	},
	2248: {
		Fword:  __ccgo_ts + 15556,
		Ftype1: int8('F'),
	},
	2249: {
		Fword:  __ccgo_ts + 15563,
		Ftype1: int8('F'),
	},
	2250: {
		Fword:  __ccgo_ts + 15570,
		Ftype1: int8('F'),
	},
	2251: {
		Fword:  __ccgo_ts + 15577,
		Ftype1: int8('F'),
	},
	2252: {
		Fword:  __ccgo_ts + 15584,
		Ftype1: int8('F'),
	},
	2253: {
		Fword:  __ccgo_ts + 15590,
		Ftype1: int8('F'),
	},
	2254: {
		Fword:  __ccgo_ts + 15597,
		Ftype1: int8('F'),
	},
	2255: {
		Fword:  __ccgo_ts + 15604,
		Ftype1: int8('F'),
	},
	2256: {
		Fword:  __ccgo_ts + 15611,
		Ftype1: int8('F'),
	},
	2257: {
		Fword:  __ccgo_ts + 15618,
		Ftype1: int8('F'),
	},
	2258: {
		Fword:  __ccgo_ts + 15625,
		Ftype1: int8('F'),
	},
	2259: {
		Fword:  __ccgo_ts + 15632,
		Ftype1: int8('F'),
	},
	2260: {
		Fword:  __ccgo_ts + 15639,
		Ftype1: int8('F'),
	},
	2261: {
		Fword:  __ccgo_ts + 15646,
		Ftype1: int8('F'),
	},
	2262: {
		Fword:  __ccgo_ts + 15652,
		Ftype1: int8('F'),
	},
	2263: {
		Fword:  __ccgo_ts + 15659,
		Ftype1: int8('F'),
	},
	2264: {
		Fword:  __ccgo_ts + 15666,
		Ftype1: int8('F'),
	},
	2265: {
		Fword:  __ccgo_ts + 15673,
		Ftype1: int8('F'),
	},
	2266: {
		Fword:  __ccgo_ts + 15680,
		Ftype1: int8('F'),
	},
	2267: {
		Fword:  __ccgo_ts + 15687,
		Ftype1: int8('F'),
	},
	2268: {
		Fword:  __ccgo_ts + 15694,
		Ftype1: int8('F'),
	},
	2269: {
		Fword:  __ccgo_ts + 15701,
		Ftype1: int8('F'),
	},
	2270: {
		Fword:  __ccgo_ts + 15708,
		Ftype1: int8('F'),
	},
	2271: {
		Fword:  __ccgo_ts + 15715,
		Ftype1: int8('F'),
	},
	2272: {
		Fword:  __ccgo_ts + 15722,
		Ftype1: int8('F'),
	},
	2273: {
		Fword:  __ccgo_ts + 15729,
		Ftype1: int8('F'),
	},
	2274: {
		Fword:  __ccgo_ts + 15736,
		Ftype1: int8('F'),
	},
	2275: {
		Fword:  __ccgo_ts + 15743,
		Ftype1: int8('F'),
	},
	2276: {
		Fword:  __ccgo_ts + 15750,
		Ftype1: int8('F'),
	},
	2277: {
		Fword:  __ccgo_ts + 15757,
		Ftype1: int8('F'),
	},
	2278: {
		Fword:  __ccgo_ts + 15764,
		Ftype1: int8('F'),
	},
	2279: {
		Fword:  __ccgo_ts + 15771,
		Ftype1: int8('F'),
	},
	2280: {
		Fword:  __ccgo_ts + 15778,
		Ftype1: int8('F'),
	},
	2281: {
		Fword:  __ccgo_ts + 15785,
		Ftype1: int8('F'),
	},
	2282: {
		Fword:  __ccgo_ts + 15792,
		Ftype1: int8('F'),
	},
	2283: {
		Fword:  __ccgo_ts + 15799,
		Ftype1: int8('F'),
	},
	2284: {
		Fword:  __ccgo_ts + 15806,
		Ftype1: int8('F'),
	},
	2285: {
		Fword:  __ccgo_ts + 15813,
		Ftype1: int8('F'),
	},
	2286: {
		Fword:  __ccgo_ts + 15820,
		Ftype1: int8('F'),
	},
	2287: {
		Fword:  __ccgo_ts + 15827,
		Ftype1: int8('F'),
	},
	2288: {
		Fword:  __ccgo_ts + 15834,
		Ftype1: int8('F'),
	},
	2289: {
		Fword:  __ccgo_ts + 15841,
		Ftype1: int8('F'),
	},
	2290: {
		Fword:  __ccgo_ts + 15848,
		Ftype1: int8('F'),
	},
	2291: {
		Fword:  __ccgo_ts + 15855,
		Ftype1: int8('F'),
	},
	2292: {
		Fword:  __ccgo_ts + 15862,
		Ftype1: int8('F'),
	},
	2293: {
		Fword:  __ccgo_ts + 15869,
		Ftype1: int8('F'),
	},
	2294: {
		Fword:  __ccgo_ts + 15876,
		Ftype1: int8('F'),
	},
	2295: {
		Fword:  __ccgo_ts + 15883,
		Ftype1: int8('F'),
	},
	2296: {
		Fword:  __ccgo_ts + 15889,
		Ftype1: int8('F'),
	},
	2297: {
		Fword:  __ccgo_ts + 15896,
		Ftype1: int8('F'),
	},
	2298: {
		Fword:  __ccgo_ts + 15903,
		Ftype1: int8('F'),
	},
	2299: {
		Fword:  __ccgo_ts + 15910,
		Ftype1: int8('F'),
	},
	2300: {
		Fword:  __ccgo_ts + 15917,
		Ftype1: int8('F'),
	},
	2301: {
		Fword:  __ccgo_ts + 15924,
		Ftype1: int8('F'),
	},
	2302: {
		Fword:  __ccgo_ts + 15931,
		Ftype1: int8('F'),
	},
	2303: {
		Fword:  __ccgo_ts + 15938,
		Ftype1: int8('F'),
	},
	2304: {
		Fword:  __ccgo_ts + 15945,
		Ftype1: int8('F'),
	},
	2305: {
		Fword:  __ccgo_ts + 15951,
		Ftype1: int8('F'),
	},
	2306: {
		Fword:  __ccgo_ts + 15958,
		Ftype1: int8('F'),
	},
	2307: {
		Fword:  __ccgo_ts + 15965,
		Ftype1: int8('F'),
	},
	2308: {
		Fword:  __ccgo_ts + 15972,
		Ftype1: int8('F'),
	},
	2309: {
		Fword:  __ccgo_ts + 15979,
		Ftype1: int8('F'),
	},
	2310: {
		Fword:  __ccgo_ts + 15986,
		Ftype1: int8('F'),
	},
	2311: {
		Fword:  __ccgo_ts + 15993,
		Ftype1: int8('F'),
	},
	2312: {
		Fword:  __ccgo_ts + 16000,
		Ftype1: int8('F'),
	},
	2313: {
		Fword:  __ccgo_ts + 16007,
		Ftype1: int8('F'),
	},
	2314: {
		Fword:  __ccgo_ts + 16014,
		Ftype1: int8('F'),
	},
	2315: {
		Fword:  __ccgo_ts + 16021,
		Ftype1: int8('F'),
	},
	2316: {
		Fword:  __ccgo_ts + 16028,
		Ftype1: int8('F'),
	},
	2317: {
		Fword:  __ccgo_ts + 16035,
		Ftype1: int8('F'),
	},
	2318: {
		Fword:  __ccgo_ts + 16042,
		Ftype1: int8('F'),
	},
	2319: {
		Fword:  __ccgo_ts + 16049,
		Ftype1: int8('F'),
	},
	2320: {
		Fword:  __ccgo_ts + 16056,
		Ftype1: int8('F'),
	},
	2321: {
		Fword:  __ccgo_ts + 16063,
		Ftype1: int8('F'),
	},
	2322: {
		Fword:  __ccgo_ts + 16070,
		Ftype1: int8('F'),
	},
	2323: {
		Fword:  __ccgo_ts + 16077,
		Ftype1: int8('F'),
	},
	2324: {
		Fword:  __ccgo_ts + 16084,
		Ftype1: int8('F'),
	},
	2325: {
		Fword:  __ccgo_ts + 16091,
		Ftype1: int8('F'),
	},
	2326: {
		Fword:  __ccgo_ts + 16098,
		Ftype1: int8('F'),
	},
	2327: {
		Fword:  __ccgo_ts + 16105,
		Ftype1: int8('F'),
	},
	2328: {
		Fword:  __ccgo_ts + 16112,
		Ftype1: int8('F'),
	},
	2329: {
		Fword:  __ccgo_ts + 16119,
		Ftype1: int8('F'),
	},
	2330: {
		Fword:  __ccgo_ts + 16126,
		Ftype1: int8('F'),
	},
	2331: {
		Fword:  __ccgo_ts + 16133,
		Ftype1: int8('F'),
	},
	2332: {
		Fword:  __ccgo_ts + 16139,
		Ftype1: int8('F'),
	},
	2333: {
		Fword:  __ccgo_ts + 16146,
		Ftype1: int8('F'),
	},
	2334: {
		Fword:  __ccgo_ts + 16153,
		Ftype1: int8('F'),
	},
	2335: {
		Fword:  __ccgo_ts + 16160,
		Ftype1: int8('F'),
	},
	2336: {
		Fword:  __ccgo_ts + 16167,
		Ftype1: int8('F'),
	},
	2337: {
		Fword:  __ccgo_ts + 16174,
		Ftype1: int8('F'),
	},
	2338: {
		Fword:  __ccgo_ts + 16181,
		Ftype1: int8('F'),
	},
	2339: {
		Fword:  __ccgo_ts + 16188,
		Ftype1: int8('F'),
	},
	2340: {
		Fword:  __ccgo_ts + 16195,
		Ftype1: int8('F'),
	},
	2341: {
		Fword:  __ccgo_ts + 16202,
		Ftype1: int8('F'),
	},
	2342: {
		Fword:  __ccgo_ts + 16209,
		Ftype1: int8('F'),
	},
	2343: {
		Fword:  __ccgo_ts + 16216,
		Ftype1: int8('F'),
	},
	2344: {
		Fword:  __ccgo_ts + 16222,
		Ftype1: int8('F'),
	},
	2345: {
		Fword:  __ccgo_ts + 16229,
		Ftype1: int8('F'),
	},
	2346: {
		Fword:  __ccgo_ts + 16236,
		Ftype1: int8('F'),
	},
	2347: {
		Fword:  __ccgo_ts + 16243,
		Ftype1: int8('F'),
	},
	2348: {
		Fword:  __ccgo_ts + 16250,
		Ftype1: int8('F'),
	},
	2349: {
		Fword:  __ccgo_ts + 16257,
		Ftype1: int8('F'),
	},
	2350: {
		Fword:  __ccgo_ts + 16264,
		Ftype1: int8('F'),
	},
	2351: {
		Fword:  __ccgo_ts + 16271,
		Ftype1: int8('F'),
	},
	2352: {
		Fword:  __ccgo_ts + 16278,
		Ftype1: int8('F'),
	},
	2353: {
		Fword:  __ccgo_ts + 16285,
		Ftype1: int8('F'),
	},
	2354: {
		Fword:  __ccgo_ts + 16292,
		Ftype1: int8('F'),
	},
	2355: {
		Fword:  __ccgo_ts + 16299,
		Ftype1: int8('F'),
	},
	2356: {
		Fword:  __ccgo_ts + 16306,
		Ftype1: int8('F'),
	},
	2357: {
		Fword:  __ccgo_ts + 16313,
		Ftype1: int8('F'),
	},
	2358: {
		Fword:  __ccgo_ts + 16320,
		Ftype1: int8('F'),
	},
	2359: {
		Fword:  __ccgo_ts + 16327,
		Ftype1: int8('F'),
	},
	2360: {
		Fword:  __ccgo_ts + 16334,
		Ftype1: int8('F'),
	},
	2361: {
		Fword:  __ccgo_ts + 16341,
		Ftype1: int8('F'),
	},
	2362: {
		Fword:  __ccgo_ts + 16348,
		Ftype1: int8('F'),
	},
	2363: {
		Fword:  __ccgo_ts + 16355,
		Ftype1: int8('F'),
	},
	2364: {
		Fword:  __ccgo_ts + 16362,
		Ftype1: int8('F'),
	},
	2365: {
		Fword:  __ccgo_ts + 16369,
		Ftype1: int8('F'),
	},
	2366: {
		Fword:  __ccgo_ts + 16376,
		Ftype1: int8('F'),
	},
	2367: {
		Fword:  __ccgo_ts + 16383,
		Ftype1: int8('F'),
	},
	2368: {
		Fword:  __ccgo_ts + 16390,
		Ftype1: int8('F'),
	},
	2369: {
		Fword:  __ccgo_ts + 16397,
		Ftype1: int8('F'),
	},
	2370: {
		Fword:  __ccgo_ts + 16403,
		Ftype1: int8('F'),
	},
	2371: {
		Fword:  __ccgo_ts + 16410,
		Ftype1: int8('F'),
	},
	2372: {
		Fword:  __ccgo_ts + 16417,
		Ftype1: int8('F'),
	},
	2373: {
		Fword:  __ccgo_ts + 16424,
		Ftype1: int8('F'),
	},
	2374: {
		Fword:  __ccgo_ts + 16431,
		Ftype1: int8('F'),
	},
	2375: {
		Fword:  __ccgo_ts + 16438,
		Ftype1: int8('F'),
	},
	2376: {
		Fword:  __ccgo_ts + 16445,
		Ftype1: int8('F'),
	},
	2377: {
		Fword:  __ccgo_ts + 16452,
		Ftype1: int8('F'),
	},
	2378: {
		Fword:  __ccgo_ts + 16459,
		Ftype1: int8('F'),
	},
	2379: {
		Fword:  __ccgo_ts + 16466,
		Ftype1: int8('F'),
	},
	2380: {
		Fword:  __ccgo_ts + 16473,
		Ftype1: int8('F'),
	},
	2381: {
		Fword:  __ccgo_ts + 16480,
		Ftype1: int8('F'),
	},
	2382: {
		Fword:  __ccgo_ts + 16487,
		Ftype1: int8('F'),
	},
	2383: {
		Fword:  __ccgo_ts + 16494,
		Ftype1: int8('F'),
	},
	2384: {
		Fword:  __ccgo_ts + 16501,
		Ftype1: int8('F'),
	},
	2385: {
		Fword:  __ccgo_ts + 16508,
		Ftype1: int8('F'),
	},
	2386: {
		Fword:  __ccgo_ts + 16515,
		Ftype1: int8('F'),
	},
	2387: {
		Fword:  __ccgo_ts + 16522,
		Ftype1: int8('F'),
	},
	2388: {
		Fword:  __ccgo_ts + 16529,
		Ftype1: int8('F'),
	},
	2389: {
		Fword:  __ccgo_ts + 16536,
		Ftype1: int8('F'),
	},
	2390: {
		Fword:  __ccgo_ts + 16543,
		Ftype1: int8('F'),
	},
	2391: {
		Fword:  __ccgo_ts + 16550,
		Ftype1: int8('F'),
	},
	2392: {
		Fword:  __ccgo_ts + 16557,
		Ftype1: int8('F'),
	},
	2393: {
		Fword:  __ccgo_ts + 16564,
		Ftype1: int8('F'),
	},
	2394: {
		Fword:  __ccgo_ts + 16571,
		Ftype1: int8('F'),
	},
	2395: {
		Fword:  __ccgo_ts + 16578,
		Ftype1: int8('F'),
	},
	2396: {
		Fword:  __ccgo_ts + 16585,
		Ftype1: int8('F'),
	},
	2397: {
		Fword:  __ccgo_ts + 16592,
		Ftype1: int8('F'),
	},
	2398: {
		Fword:  __ccgo_ts + 16599,
		Ftype1: int8('F'),
	},
	2399: {
		Fword:  __ccgo_ts + 16604,
		Ftype1: int8('F'),
	},
	2400: {
		Fword:  __ccgo_ts + 16611,
		Ftype1: int8('F'),
	},
	2401: {
		Fword:  __ccgo_ts + 16618,
		Ftype1: int8('F'),
	},
	2402: {
		Fword:  __ccgo_ts + 16625,
		Ftype1: int8('F'),
	},
	2403: {
		Fword:  __ccgo_ts + 16632,
		Ftype1: int8('F'),
	},
	2404: {
		Fword:  __ccgo_ts + 16639,
		Ftype1: int8('F'),
	},
	2405: {
		Fword:  __ccgo_ts + 16646,
		Ftype1: int8('F'),
	},
	2406: {
		Fword:  __ccgo_ts + 16652,
		Ftype1: int8('F'),
	},
	2407: {
		Fword:  __ccgo_ts + 16659,
		Ftype1: int8('F'),
	},
	2408: {
		Fword:  __ccgo_ts + 16666,
		Ftype1: int8('F'),
	},
	2409: {
		Fword:  __ccgo_ts + 16673,
		Ftype1: int8('F'),
	},
	2410: {
		Fword:  __ccgo_ts + 16680,
		Ftype1: int8('F'),
	},
	2411: {
		Fword:  __ccgo_ts + 16686,
		Ftype1: int8('F'),
	},
	2412: {
		Fword:  __ccgo_ts + 16693,
		Ftype1: int8('F'),
	},
	2413: {
		Fword:  __ccgo_ts + 16700,
		Ftype1: int8('F'),
	},
	2414: {
		Fword:  __ccgo_ts + 16707,
		Ftype1: int8('F'),
	},
	2415: {
		Fword:  __ccgo_ts + 16714,
		Ftype1: int8('F'),
	},
	2416: {
		Fword:  __ccgo_ts + 16721,
		Ftype1: int8('F'),
	},
	2417: {
		Fword:  __ccgo_ts + 16728,
		Ftype1: int8('F'),
	},
	2418: {
		Fword:  __ccgo_ts + 16735,
		Ftype1: int8('F'),
	},
	2419: {
		Fword:  __ccgo_ts + 16742,
		Ftype1: int8('F'),
	},
	2420: {
		Fword:  __ccgo_ts + 16749,
		Ftype1: int8('F'),
	},
	2421: {
		Fword:  __ccgo_ts + 16756,
		Ftype1: int8('F'),
	},
	2422: {
		Fword:  __ccgo_ts + 16763,
		Ftype1: int8('F'),
	},
	2423: {
		Fword:  __ccgo_ts + 16770,
		Ftype1: int8('F'),
	},
	2424: {
		Fword:  __ccgo_ts + 16777,
		Ftype1: int8('F'),
	},
	2425: {
		Fword:  __ccgo_ts + 16784,
		Ftype1: int8('F'),
	},
	2426: {
		Fword:  __ccgo_ts + 16790,
		Ftype1: int8('F'),
	},
	2427: {
		Fword:  __ccgo_ts + 16797,
		Ftype1: int8('F'),
	},
	2428: {
		Fword:  __ccgo_ts + 16804,
		Ftype1: int8('F'),
	},
	2429: {
		Fword:  __ccgo_ts + 16811,
		Ftype1: int8('F'),
	},
	2430: {
		Fword:  __ccgo_ts + 16818,
		Ftype1: int8('F'),
	},
	2431: {
		Fword:  __ccgo_ts + 16825,
		Ftype1: int8('F'),
	},
	2432: {
		Fword:  __ccgo_ts + 16832,
		Ftype1: int8('F'),
	},
	2433: {
		Fword:  __ccgo_ts + 16839,
		Ftype1: int8('F'),
	},
	2434: {
		Fword:  __ccgo_ts + 16846,
		Ftype1: int8('F'),
	},
	2435: {
		Fword:  __ccgo_ts + 16853,
		Ftype1: int8('F'),
	},
	2436: {
		Fword:  __ccgo_ts + 16860,
		Ftype1: int8('F'),
	},
	2437: {
		Fword:  __ccgo_ts + 16867,
		Ftype1: int8('F'),
	},
	2438: {
		Fword:  __ccgo_ts + 16874,
		Ftype1: int8('F'),
	},
	2439: {
		Fword:  __ccgo_ts + 16881,
		Ftype1: int8('F'),
	},
	2440: {
		Fword:  __ccgo_ts + 16888,
		Ftype1: int8('F'),
	},
	2441: {
		Fword:  __ccgo_ts + 16895,
		Ftype1: int8('F'),
	},
	2442: {
		Fword:  __ccgo_ts + 16902,
		Ftype1: int8('F'),
	},
	2443: {
		Fword:  __ccgo_ts + 16909,
		Ftype1: int8('F'),
	},
	2444: {
		Fword:  __ccgo_ts + 16916,
		Ftype1: int8('F'),
	},
	2445: {
		Fword:  __ccgo_ts + 16923,
		Ftype1: int8('F'),
	},
	2446: {
		Fword:  __ccgo_ts + 16930,
		Ftype1: int8('F'),
	},
	2447: {
		Fword:  __ccgo_ts + 16937,
		Ftype1: int8('F'),
	},
	2448: {
		Fword:  __ccgo_ts + 16944,
		Ftype1: int8('F'),
	},
	2449: {
		Fword:  __ccgo_ts + 16951,
		Ftype1: int8('F'),
	},
	2450: {
		Fword:  __ccgo_ts + 16958,
		Ftype1: int8('F'),
	},
	2451: {
		Fword:  __ccgo_ts + 16965,
		Ftype1: int8('F'),
	},
	2452: {
		Fword:  __ccgo_ts + 16972,
		Ftype1: int8('F'),
	},
	2453: {
		Fword:  __ccgo_ts + 16979,
		Ftype1: int8('F'),
	},
	2454: {
		Fword:  __ccgo_ts + 16986,
		Ftype1: int8('F'),
	},
	2455: {
		Fword:  __ccgo_ts + 16993,
		Ftype1: int8('F'),
	},
	2456: {
		Fword:  __ccgo_ts + 17000,
		Ftype1: int8('F'),
	},
	2457: {
		Fword:  __ccgo_ts + 17007,
		Ftype1: int8('F'),
	},
	2458: {
		Fword:  __ccgo_ts + 17014,
		Ftype1: int8('F'),
	},
	2459: {
		Fword:  __ccgo_ts + 17021,
		Ftype1: int8('F'),
	},
	2460: {
		Fword:  __ccgo_ts + 17028,
		Ftype1: int8('F'),
	},
	2461: {
		Fword:  __ccgo_ts + 17035,
		Ftype1: int8('F'),
	},
	2462: {
		Fword:  __ccgo_ts + 17042,
		Ftype1: int8('F'),
	},
	2463: {
		Fword:  __ccgo_ts + 17049,
		Ftype1: int8('F'),
	},
	2464: {
		Fword:  __ccgo_ts + 17056,
		Ftype1: int8('F'),
	},
	2465: {
		Fword:  __ccgo_ts + 17063,
		Ftype1: int8('F'),
	},
	2466: {
		Fword:  __ccgo_ts + 17070,
		Ftype1: int8('F'),
	},
	2467: {
		Fword:  __ccgo_ts + 17077,
		Ftype1: int8('F'),
	},
	2468: {
		Fword:  __ccgo_ts + 17084,
		Ftype1: int8('F'),
	},
	2469: {
		Fword:  __ccgo_ts + 17091,
		Ftype1: int8('F'),
	},
	2470: {
		Fword:  __ccgo_ts + 17098,
		Ftype1: int8('F'),
	},
	2471: {
		Fword:  __ccgo_ts + 17105,
		Ftype1: int8('F'),
	},
	2472: {
		Fword:  __ccgo_ts + 17112,
		Ftype1: int8('F'),
	},
	2473: {
		Fword:  __ccgo_ts + 17119,
		Ftype1: int8('F'),
	},
	2474: {
		Fword:  __ccgo_ts + 17126,
		Ftype1: int8('F'),
	},
	2475: {
		Fword:  __ccgo_ts + 17133,
		Ftype1: int8('F'),
	},
	2476: {
		Fword:  __ccgo_ts + 17140,
		Ftype1: int8('F'),
	},
	2477: {
		Fword:  __ccgo_ts + 17147,
		Ftype1: int8('F'),
	},
	2478: {
		Fword:  __ccgo_ts + 17154,
		Ftype1: int8('F'),
	},
	2479: {
		Fword:  __ccgo_ts + 17161,
		Ftype1: int8('F'),
	},
	2480: {
		Fword:  __ccgo_ts + 17168,
		Ftype1: int8('F'),
	},
	2481: {
		Fword:  __ccgo_ts + 17175,
		Ftype1: int8('F'),
	},
	2482: {
		Fword:  __ccgo_ts + 17182,
		Ftype1: int8('F'),
	},
	2483: {
		Fword:  __ccgo_ts + 17189,
		Ftype1: int8('F'),
	},
	2484: {
		Fword:  __ccgo_ts + 17196,
		Ftype1: int8('F'),
	},
	2485: {
		Fword:  __ccgo_ts + 17203,
		Ftype1: int8('F'),
	},
	2486: {
		Fword:  __ccgo_ts + 17210,
		Ftype1: int8('F'),
	},
	2487: {
		Fword:  __ccgo_ts + 17217,
		Ftype1: int8('F'),
	},
	2488: {
		Fword:  __ccgo_ts + 17224,
		Ftype1: int8('F'),
	},
	2489: {
		Fword:  __ccgo_ts + 17231,
		Ftype1: int8('F'),
	},
	2490: {
		Fword:  __ccgo_ts + 17238,
		Ftype1: int8('F'),
	},
	2491: {
		Fword:  __ccgo_ts + 17245,
		Ftype1: int8('F'),
	},
	2492: {
		Fword:  __ccgo_ts + 17252,
		Ftype1: int8('F'),
	},
	2493: {
		Fword:  __ccgo_ts + 17259,
		Ftype1: int8('F'),
	},
	2494: {
		Fword:  __ccgo_ts + 17266,
		Ftype1: int8('F'),
	},
	2495: {
		Fword:  __ccgo_ts + 17273,
		Ftype1: int8('F'),
	},
	2496: {
		Fword:  __ccgo_ts + 17280,
		Ftype1: int8('F'),
	},
	2497: {
		Fword:  __ccgo_ts + 17287,
		Ftype1: int8('F'),
	},
	2498: {
		Fword:  __ccgo_ts + 17294,
		Ftype1: int8('F'),
	},
	2499: {
		Fword:  __ccgo_ts + 17301,
		Ftype1: int8('F'),
	},
	2500: {
		Fword:  __ccgo_ts + 17308,
		Ftype1: int8('F'),
	},
	2501: {
		Fword:  __ccgo_ts + 17315,
		Ftype1: int8('F'),
	},
	2502: {
		Fword:  __ccgo_ts + 17322,
		Ftype1: int8('F'),
	},
	2503: {
		Fword:  __ccgo_ts + 17329,
		Ftype1: int8('F'),
	},
	2504: {
		Fword:  __ccgo_ts + 17336,
		Ftype1: int8('F'),
	},
	2505: {
		Fword:  __ccgo_ts + 17343,
		Ftype1: int8('F'),
	},
	2506: {
		Fword:  __ccgo_ts + 17350,
		Ftype1: int8('F'),
	},
	2507: {
		Fword:  __ccgo_ts + 17357,
		Ftype1: int8('F'),
	},
	2508: {
		Fword:  __ccgo_ts + 17364,
		Ftype1: int8('F'),
	},
	2509: {
		Fword:  __ccgo_ts + 17371,
		Ftype1: int8('F'),
	},
	2510: {
		Fword:  __ccgo_ts + 17378,
		Ftype1: int8('F'),
	},
	2511: {
		Fword:  __ccgo_ts + 17385,
		Ftype1: int8('F'),
	},
	2512: {
		Fword:  __ccgo_ts + 17392,
		Ftype1: int8('F'),
	},
	2513: {
		Fword:  __ccgo_ts + 17399,
		Ftype1: int8('F'),
	},
	2514: {
		Fword:  __ccgo_ts + 17406,
		Ftype1: int8('F'),
	},
	2515: {
		Fword:  __ccgo_ts + 17413,
		Ftype1: int8('F'),
	},
	2516: {
		Fword:  __ccgo_ts + 17420,
		Ftype1: int8('F'),
	},
	2517: {
		Fword:  __ccgo_ts + 17427,
		Ftype1: int8('F'),
	},
	2518: {
		Fword:  __ccgo_ts + 17434,
		Ftype1: int8('F'),
	},
	2519: {
		Fword:  __ccgo_ts + 17441,
		Ftype1: int8('F'),
	},
	2520: {
		Fword:  __ccgo_ts + 17448,
		Ftype1: int8('F'),
	},
	2521: {
		Fword:  __ccgo_ts + 17455,
		Ftype1: int8('F'),
	},
	2522: {
		Fword:  __ccgo_ts + 17462,
		Ftype1: int8('F'),
	},
	2523: {
		Fword:  __ccgo_ts + 17469,
		Ftype1: int8('F'),
	},
	2524: {
		Fword:  __ccgo_ts + 17476,
		Ftype1: int8('F'),
	},
	2525: {
		Fword:  __ccgo_ts + 17483,
		Ftype1: int8('F'),
	},
	2526: {
		Fword:  __ccgo_ts + 17490,
		Ftype1: int8('F'),
	},
	2527: {
		Fword:  __ccgo_ts + 17497,
		Ftype1: int8('F'),
	},
	2528: {
		Fword:  __ccgo_ts + 17504,
		Ftype1: int8('F'),
	},
	2529: {
		Fword:  __ccgo_ts + 17511,
		Ftype1: int8('F'),
	},
	2530: {
		Fword:  __ccgo_ts + 17516,
		Ftype1: int8('F'),
	},
	2531: {
		Fword:  __ccgo_ts + 17523,
		Ftype1: int8('F'),
	},
	2532: {
		Fword:  __ccgo_ts + 17530,
		Ftype1: int8('F'),
	},
	2533: {
		Fword:  __ccgo_ts + 17537,
		Ftype1: int8('F'),
	},
	2534: {
		Fword:  __ccgo_ts + 17544,
		Ftype1: int8('F'),
	},
	2535: {
		Fword:  __ccgo_ts + 17551,
		Ftype1: int8('F'),
	},
	2536: {
		Fword:  __ccgo_ts + 17558,
		Ftype1: int8('F'),
	},
	2537: {
		Fword:  __ccgo_ts + 17564,
		Ftype1: int8('F'),
	},
	2538: {
		Fword:  __ccgo_ts + 17571,
		Ftype1: int8('F'),
	},
	2539: {
		Fword:  __ccgo_ts + 17578,
		Ftype1: int8('F'),
	},
	2540: {
		Fword:  __ccgo_ts + 17585,
		Ftype1: int8('F'),
	},
	2541: {
		Fword:  __ccgo_ts + 17592,
		Ftype1: int8('F'),
	},
	2542: {
		Fword:  __ccgo_ts + 17598,
		Ftype1: int8('F'),
	},
	2543: {
		Fword:  __ccgo_ts + 17605,
		Ftype1: int8('F'),
	},
	2544: {
		Fword:  __ccgo_ts + 17612,
		Ftype1: int8('F'),
	},
	2545: {
		Fword:  __ccgo_ts + 17619,
		Ftype1: int8('F'),
	},
	2546: {
		Fword:  __ccgo_ts + 17626,
		Ftype1: int8('F'),
	},
	2547: {
		Fword:  __ccgo_ts + 17633,
		Ftype1: int8('F'),
	},
	2548: {
		Fword:  __ccgo_ts + 17640,
		Ftype1: int8('F'),
	},
	2549: {
		Fword:  __ccgo_ts + 17647,
		Ftype1: int8('F'),
	},
	2550: {
		Fword:  __ccgo_ts + 17654,
		Ftype1: int8('F'),
	},
	2551: {
		Fword:  __ccgo_ts + 17661,
		Ftype1: int8('F'),
	},
	2552: {
		Fword:  __ccgo_ts + 17668,
		Ftype1: int8('F'),
	},
	2553: {
		Fword:  __ccgo_ts + 17675,
		Ftype1: int8('F'),
	},
	2554: {
		Fword:  __ccgo_ts + 17682,
		Ftype1: int8('F'),
	},
	2555: {
		Fword:  __ccgo_ts + 17689,
		Ftype1: int8('F'),
	},
	2556: {
		Fword:  __ccgo_ts + 17696,
		Ftype1: int8('F'),
	},
	2557: {
		Fword:  __ccgo_ts + 17701,
		Ftype1: int8('F'),
	},
	2558: {
		Fword:  __ccgo_ts + 17707,
		Ftype1: int8('F'),
	},
	2559: {
		Fword:  __ccgo_ts + 17714,
		Ftype1: int8('F'),
	},
	2560: {
		Fword:  __ccgo_ts + 17720,
		Ftype1: int8('F'),
	},
	2561: {
		Fword:  __ccgo_ts + 17727,
		Ftype1: int8('F'),
	},
	2562: {
		Fword:  __ccgo_ts + 17734,
		Ftype1: int8('F'),
	},
	2563: {
		Fword:  __ccgo_ts + 17741,
		Ftype1: int8('F'),
	},
	2564: {
		Fword:  __ccgo_ts + 17748,
		Ftype1: int8('F'),
	},
	2565: {
		Fword:  __ccgo_ts + 17755,
		Ftype1: int8('F'),
	},
	2566: {
		Fword:  __ccgo_ts + 17762,
		Ftype1: int8('F'),
	},
	2567: {
		Fword:  __ccgo_ts + 17769,
		Ftype1: int8('F'),
	},
	2568: {
		Fword:  __ccgo_ts + 17776,
		Ftype1: int8('F'),
	},
	2569: {
		Fword:  __ccgo_ts + 17783,
		Ftype1: int8('F'),
	},
	2570: {
		Fword:  __ccgo_ts + 17790,
		Ftype1: int8('F'),
	},
	2571: {
		Fword:  __ccgo_ts + 17797,
		Ftype1: int8('F'),
	},
	2572: {
		Fword:  __ccgo_ts + 17804,
		Ftype1: int8('F'),
	},
	2573: {
		Fword:  __ccgo_ts + 17811,
		Ftype1: int8('F'),
	},
	2574: {
		Fword:  __ccgo_ts + 17818,
		Ftype1: int8('F'),
	},
	2575: {
		Fword:  __ccgo_ts + 17825,
		Ftype1: int8('F'),
	},
	2576: {
		Fword:  __ccgo_ts + 17832,
		Ftype1: int8('F'),
	},
	2577: {
		Fword:  __ccgo_ts + 17839,
		Ftype1: int8('F'),
	},
	2578: {
		Fword:  __ccgo_ts + 17846,
		Ftype1: int8('F'),
	},
	2579: {
		Fword:  __ccgo_ts + 17853,
		Ftype1: int8('F'),
	},
	2580: {
		Fword:  __ccgo_ts + 17860,
		Ftype1: int8('F'),
	},
	2581: {
		Fword:  __ccgo_ts + 17867,
		Ftype1: int8('F'),
	},
	2582: {
		Fword:  __ccgo_ts + 17874,
		Ftype1: int8('F'),
	},
	2583: {
		Fword:  __ccgo_ts + 17881,
		Ftype1: int8('F'),
	},
	2584: {
		Fword:  __ccgo_ts + 17888,
		Ftype1: int8('F'),
	},
	2585: {
		Fword:  __ccgo_ts + 17895,
		Ftype1: int8('F'),
	},
	2586: {
		Fword:  __ccgo_ts + 17902,
		Ftype1: int8('F'),
	},
	2587: {
		Fword:  __ccgo_ts + 17909,
		Ftype1: int8('F'),
	},
	2588: {
		Fword:  __ccgo_ts + 17916,
		Ftype1: int8('F'),
	},
	2589: {
		Fword:  __ccgo_ts + 17923,
		Ftype1: int8('F'),
	},
	2590: {
		Fword:  __ccgo_ts + 17930,
		Ftype1: int8('F'),
	},
	2591: {
		Fword:  __ccgo_ts + 17937,
		Ftype1: int8('F'),
	},
	2592: {
		Fword:  __ccgo_ts + 17943,
		Ftype1: int8('F'),
	},
	2593: {
		Fword:  __ccgo_ts + 17950,
		Ftype1: int8('F'),
	},
	2594: {
		Fword:  __ccgo_ts + 17957,
		Ftype1: int8('F'),
	},
	2595: {
		Fword:  __ccgo_ts + 17964,
		Ftype1: int8('F'),
	},
	2596: {
		Fword:  __ccgo_ts + 17971,
		Ftype1: int8('F'),
	},
	2597: {
		Fword:  __ccgo_ts + 17978,
		Ftype1: int8('F'),
	},
	2598: {
		Fword:  __ccgo_ts + 17985,
		Ftype1: int8('F'),
	},
	2599: {
		Fword:  __ccgo_ts + 17992,
		Ftype1: int8('F'),
	},
	2600: {
		Fword:  __ccgo_ts + 17999,
		Ftype1: int8('F'),
	},
	2601: {
		Fword:  __ccgo_ts + 18006,
		Ftype1: int8('F'),
	},
	2602: {
		Fword:  __ccgo_ts + 18013,
		Ftype1: int8('F'),
	},
	2603: {
		Fword:  __ccgo_ts + 18020,
		Ftype1: int8('F'),
	},
	2604: {
		Fword:  __ccgo_ts + 18027,
		Ftype1: int8('F'),
	},
	2605: {
		Fword:  __ccgo_ts + 18034,
		Ftype1: int8('F'),
	},
	2606: {
		Fword:  __ccgo_ts + 18041,
		Ftype1: int8('F'),
	},
	2607: {
		Fword:  __ccgo_ts + 18048,
		Ftype1: int8('F'),
	},
	2608: {
		Fword:  __ccgo_ts + 18055,
		Ftype1: int8('F'),
	},
	2609: {
		Fword:  __ccgo_ts + 18062,
		Ftype1: int8('F'),
	},
	2610: {
		Fword:  __ccgo_ts + 18069,
		Ftype1: int8('F'),
	},
	2611: {
		Fword:  __ccgo_ts + 18076,
		Ftype1: int8('F'),
	},
	2612: {
		Fword:  __ccgo_ts + 18083,
		Ftype1: int8('F'),
	},
	2613: {
		Fword:  __ccgo_ts + 18090,
		Ftype1: int8('F'),
	},
	2614: {
		Fword:  __ccgo_ts + 18097,
		Ftype1: int8('F'),
	},
	2615: {
		Fword:  __ccgo_ts + 18104,
		Ftype1: int8('F'),
	},
	2616: {
		Fword:  __ccgo_ts + 18111,
		Ftype1: int8('F'),
	},
	2617: {
		Fword:  __ccgo_ts + 18118,
		Ftype1: int8('F'),
	},
	2618: {
		Fword:  __ccgo_ts + 18125,
		Ftype1: int8('F'),
	},
	2619: {
		Fword:  __ccgo_ts + 18132,
		Ftype1: int8('F'),
	},
	2620: {
		Fword:  __ccgo_ts + 18139,
		Ftype1: int8('F'),
	},
	2621: {
		Fword:  __ccgo_ts + 18146,
		Ftype1: int8('F'),
	},
	2622: {
		Fword:  __ccgo_ts + 18153,
		Ftype1: int8('F'),
	},
	2623: {
		Fword:  __ccgo_ts + 18160,
		Ftype1: int8('F'),
	},
	2624: {
		Fword:  __ccgo_ts + 18167,
		Ftype1: int8('F'),
	},
	2625: {
		Fword:  __ccgo_ts + 18174,
		Ftype1: int8('F'),
	},
	2626: {
		Fword:  __ccgo_ts + 18181,
		Ftype1: int8('F'),
	},
	2627: {
		Fword:  __ccgo_ts + 18188,
		Ftype1: int8('F'),
	},
	2628: {
		Fword:  __ccgo_ts + 18195,
		Ftype1: int8('F'),
	},
	2629: {
		Fword:  __ccgo_ts + 18202,
		Ftype1: int8('F'),
	},
	2630: {
		Fword:  __ccgo_ts + 18209,
		Ftype1: int8('F'),
	},
	2631: {
		Fword:  __ccgo_ts + 18216,
		Ftype1: int8('F'),
	},
	2632: {
		Fword:  __ccgo_ts + 18223,
		Ftype1: int8('F'),
	},
	2633: {
		Fword:  __ccgo_ts + 18230,
		Ftype1: int8('F'),
	},
	2634: {
		Fword:  __ccgo_ts + 18237,
		Ftype1: int8('F'),
	},
	2635: {
		Fword:  __ccgo_ts + 18244,
		Ftype1: int8('F'),
	},
	2636: {
		Fword:  __ccgo_ts + 18251,
		Ftype1: int8('F'),
	},
	2637: {
		Fword:  __ccgo_ts + 18258,
		Ftype1: int8('F'),
	},
	2638: {
		Fword:  __ccgo_ts + 18265,
		Ftype1: int8('F'),
	},
	2639: {
		Fword:  __ccgo_ts + 18272,
		Ftype1: int8('F'),
	},
	2640: {
		Fword:  __ccgo_ts + 18279,
		Ftype1: int8('F'),
	},
	2641: {
		Fword:  __ccgo_ts + 18286,
		Ftype1: int8('F'),
	},
	2642: {
		Fword:  __ccgo_ts + 18293,
		Ftype1: int8('F'),
	},
	2643: {
		Fword:  __ccgo_ts + 18300,
		Ftype1: int8('F'),
	},
	2644: {
		Fword:  __ccgo_ts + 18307,
		Ftype1: int8('F'),
	},
	2645: {
		Fword:  __ccgo_ts + 18314,
		Ftype1: int8('F'),
	},
	2646: {
		Fword:  __ccgo_ts + 18321,
		Ftype1: int8('F'),
	},
	2647: {
		Fword:  __ccgo_ts + 18328,
		Ftype1: int8('F'),
	},
	2648: {
		Fword:  __ccgo_ts + 18335,
		Ftype1: int8('F'),
	},
	2649: {
		Fword:  __ccgo_ts + 18342,
		Ftype1: int8('F'),
	},
	2650: {
		Fword:  __ccgo_ts + 18349,
		Ftype1: int8('F'),
	},
	2651: {
		Fword:  __ccgo_ts + 18356,
		Ftype1: int8('F'),
	},
	2652: {
		Fword:  __ccgo_ts + 18363,
		Ftype1: int8('F'),
	},
	2653: {
		Fword:  __ccgo_ts + 18370,
		Ftype1: int8('F'),
	},
	2654: {
		Fword:  __ccgo_ts + 18377,
		Ftype1: int8('F'),
	},
	2655: {
		Fword:  __ccgo_ts + 18384,
		Ftype1: int8('F'),
	},
	2656: {
		Fword:  __ccgo_ts + 18391,
		Ftype1: int8('F'),
	},
	2657: {
		Fword:  __ccgo_ts + 18398,
		Ftype1: int8('F'),
	},
	2658: {
		Fword:  __ccgo_ts + 18405,
		Ftype1: int8('F'),
	},
	2659: {
		Fword:  __ccgo_ts + 18412,
		Ftype1: int8('F'),
	},
	2660: {
		Fword:  __ccgo_ts + 18419,
		Ftype1: int8('F'),
	},
	2661: {
		Fword:  __ccgo_ts + 18426,
		Ftype1: int8('F'),
	},
	2662: {
		Fword:  __ccgo_ts + 18433,
		Ftype1: int8('F'),
	},
	2663: {
		Fword:  __ccgo_ts + 18438,
		Ftype1: int8('F'),
	},
	2664: {
		Fword:  __ccgo_ts + 18444,
		Ftype1: int8('F'),
	},
	2665: {
		Fword:  __ccgo_ts + 18451,
		Ftype1: int8('F'),
	},
	2666: {
		Fword:  __ccgo_ts + 18457,
		Ftype1: int8('F'),
	},
	2667: {
		Fword:  __ccgo_ts + 18464,
		Ftype1: int8('F'),
	},
	2668: {
		Fword:  __ccgo_ts + 18471,
		Ftype1: int8('F'),
	},
	2669: {
		Fword:  __ccgo_ts + 18478,
		Ftype1: int8('F'),
	},
	2670: {
		Fword:  __ccgo_ts + 18485,
		Ftype1: int8('F'),
	},
	2671: {
		Fword:  __ccgo_ts + 18492,
		Ftype1: int8('F'),
	},
	2672: {
		Fword:  __ccgo_ts + 18499,
		Ftype1: int8('F'),
	},
	2673: {
		Fword:  __ccgo_ts + 18506,
		Ftype1: int8('F'),
	},
	2674: {
		Fword:  __ccgo_ts + 18513,
		Ftype1: int8('F'),
	},
	2675: {
		Fword:  __ccgo_ts + 18520,
		Ftype1: int8('F'),
	},
	2676: {
		Fword:  __ccgo_ts + 18527,
		Ftype1: int8('F'),
	},
	2677: {
		Fword:  __ccgo_ts + 18534,
		Ftype1: int8('F'),
	},
	2678: {
		Fword:  __ccgo_ts + 18541,
		Ftype1: int8('F'),
	},
	2679: {
		Fword:  __ccgo_ts + 18548,
		Ftype1: int8('F'),
	},
	2680: {
		Fword:  __ccgo_ts + 18555,
		Ftype1: int8('F'),
	},
	2681: {
		Fword:  __ccgo_ts + 18562,
		Ftype1: int8('F'),
	},
	2682: {
		Fword:  __ccgo_ts + 18569,
		Ftype1: int8('F'),
	},
	2683: {
		Fword:  __ccgo_ts + 18576,
		Ftype1: int8('F'),
	},
	2684: {
		Fword:  __ccgo_ts + 18583,
		Ftype1: int8('F'),
	},
	2685: {
		Fword:  __ccgo_ts + 18588,
		Ftype1: int8('F'),
	},
	2686: {
		Fword:  __ccgo_ts + 18595,
		Ftype1: int8('F'),
	},
	2687: {
		Fword:  __ccgo_ts + 18602,
		Ftype1: int8('F'),
	},
	2688: {
		Fword:  __ccgo_ts + 18609,
		Ftype1: int8('F'),
	},
	2689: {
		Fword:  __ccgo_ts + 18616,
		Ftype1: int8('F'),
	},
	2690: {
		Fword:  __ccgo_ts + 18623,
		Ftype1: int8('F'),
	},
	2691: {
		Fword:  __ccgo_ts + 18630,
		Ftype1: int8('F'),
	},
	2692: {
		Fword:  __ccgo_ts + 18636,
		Ftype1: int8('F'),
	},
	2693: {
		Fword:  __ccgo_ts + 18643,
		Ftype1: int8('F'),
	},
	2694: {
		Fword:  __ccgo_ts + 18650,
		Ftype1: int8('F'),
	},
	2695: {
		Fword:  __ccgo_ts + 18657,
		Ftype1: int8('F'),
	},
	2696: {
		Fword:  __ccgo_ts + 18664,
		Ftype1: int8('F'),
	},
	2697: {
		Fword:  __ccgo_ts + 18670,
		Ftype1: int8('F'),
	},
	2698: {
		Fword:  __ccgo_ts + 18677,
		Ftype1: int8('F'),
	},
	2699: {
		Fword:  __ccgo_ts + 18684,
		Ftype1: int8('F'),
	},
	2700: {
		Fword:  __ccgo_ts + 18691,
		Ftype1: int8('F'),
	},
	2701: {
		Fword:  __ccgo_ts + 18698,
		Ftype1: int8('F'),
	},
	2702: {
		Fword:  __ccgo_ts + 18705,
		Ftype1: int8('F'),
	},
	2703: {
		Fword:  __ccgo_ts + 18712,
		Ftype1: int8('F'),
	},
	2704: {
		Fword:  __ccgo_ts + 18719,
		Ftype1: int8('F'),
	},
	2705: {
		Fword:  __ccgo_ts + 18726,
		Ftype1: int8('F'),
	},
	2706: {
		Fword:  __ccgo_ts + 18733,
		Ftype1: int8('F'),
	},
	2707: {
		Fword:  __ccgo_ts + 18740,
		Ftype1: int8('F'),
	},
	2708: {
		Fword:  __ccgo_ts + 18747,
		Ftype1: int8('F'),
	},
	2709: {
		Fword:  __ccgo_ts + 18754,
		Ftype1: int8('F'),
	},
	2710: {
		Fword:  __ccgo_ts + 18761,
		Ftype1: int8('F'),
	},
	2711: {
		Fword:  __ccgo_ts + 18768,
		Ftype1: int8('F'),
	},
	2712: {
		Fword:  __ccgo_ts + 18775,
		Ftype1: int8('F'),
	},
	2713: {
		Fword:  __ccgo_ts + 18782,
		Ftype1: int8('F'),
	},
	2714: {
		Fword:  __ccgo_ts + 18789,
		Ftype1: int8('F'),
	},
	2715: {
		Fword:  __ccgo_ts + 18796,
		Ftype1: int8('F'),
	},
	2716: {
		Fword:  __ccgo_ts + 18803,
		Ftype1: int8('F'),
	},
	2717: {
		Fword:  __ccgo_ts + 18810,
		Ftype1: int8('F'),
	},
	2718: {
		Fword:  __ccgo_ts + 18817,
		Ftype1: int8('F'),
	},
	2719: {
		Fword:  __ccgo_ts + 18824,
		Ftype1: int8('F'),
	},
	2720: {
		Fword:  __ccgo_ts + 18831,
		Ftype1: int8('F'),
	},
	2721: {
		Fword:  __ccgo_ts + 18838,
		Ftype1: int8('F'),
	},
	2722: {
		Fword:  __ccgo_ts + 18845,
		Ftype1: int8('F'),
	},
	2723: {
		Fword:  __ccgo_ts + 18852,
		Ftype1: int8('F'),
	},
	2724: {
		Fword:  __ccgo_ts + 18859,
		Ftype1: int8('F'),
	},
	2725: {
		Fword:  __ccgo_ts + 18866,
		Ftype1: int8('F'),
	},
	2726: {
		Fword:  __ccgo_ts + 18873,
		Ftype1: int8('F'),
	},
	2727: {
		Fword:  __ccgo_ts + 18880,
		Ftype1: int8('F'),
	},
	2728: {
		Fword:  __ccgo_ts + 18887,
		Ftype1: int8('F'),
	},
	2729: {
		Fword:  __ccgo_ts + 18894,
		Ftype1: int8('F'),
	},
	2730: {
		Fword:  __ccgo_ts + 18901,
		Ftype1: int8('F'),
	},
	2731: {
		Fword:  __ccgo_ts + 18908,
		Ftype1: int8('F'),
	},
	2732: {
		Fword:  __ccgo_ts + 18915,
		Ftype1: int8('F'),
	},
	2733: {
		Fword:  __ccgo_ts + 18922,
		Ftype1: int8('F'),
	},
	2734: {
		Fword:  __ccgo_ts + 18929,
		Ftype1: int8('F'),
	},
	2735: {
		Fword:  __ccgo_ts + 18936,
		Ftype1: int8('F'),
	},
	2736: {
		Fword:  __ccgo_ts + 18943,
		Ftype1: int8('F'),
	},
	2737: {
		Fword:  __ccgo_ts + 18950,
		Ftype1: int8('F'),
	},
	2738: {
		Fword:  __ccgo_ts + 18957,
		Ftype1: int8('F'),
	},
	2739: {
		Fword:  __ccgo_ts + 18964,
		Ftype1: int8('F'),
	},
	2740: {
		Fword:  __ccgo_ts + 18971,
		Ftype1: int8('F'),
	},
	2741: {
		Fword:  __ccgo_ts + 18978,
		Ftype1: int8('F'),
	},
	2742: {
		Fword:  __ccgo_ts + 18985,
		Ftype1: int8('F'),
	},
	2743: {
		Fword:  __ccgo_ts + 18991,
		Ftype1: int8('F'),
	},
	2744: {
		Fword:  __ccgo_ts + 18998,
		Ftype1: int8('F'),
	},
	2745: {
		Fword:  __ccgo_ts + 19005,
		Ftype1: int8('F'),
	},
	2746: {
		Fword:  __ccgo_ts + 19012,
		Ftype1: int8('F'),
	},
	2747: {
		Fword:  __ccgo_ts + 19019,
		Ftype1: int8('F'),
	},
	2748: {
		Fword:  __ccgo_ts + 19026,
		Ftype1: int8('F'),
	},
	2749: {
		Fword:  __ccgo_ts + 19033,
		Ftype1: int8('F'),
	},
	2750: {
		Fword:  __ccgo_ts + 19040,
		Ftype1: int8('F'),
	},
	2751: {
		Fword:  __ccgo_ts + 19047,
		Ftype1: int8('F'),
	},
	2752: {
		Fword:  __ccgo_ts + 19054,
		Ftype1: int8('F'),
	},
	2753: {
		Fword:  __ccgo_ts + 19061,
		Ftype1: int8('F'),
	},
	2754: {
		Fword:  __ccgo_ts + 19068,
		Ftype1: int8('F'),
	},
	2755: {
		Fword:  __ccgo_ts + 19073,
		Ftype1: int8('F'),
	},
	2756: {
		Fword:  __ccgo_ts + 19080,
		Ftype1: int8('F'),
	},
	2757: {
		Fword:  __ccgo_ts + 19086,
		Ftype1: int8('F'),
	},
	2758: {
		Fword:  __ccgo_ts + 19093,
		Ftype1: int8('F'),
	},
	2759: {
		Fword:  __ccgo_ts + 19099,
		Ftype1: int8('F'),
	},
	2760: {
		Fword:  __ccgo_ts + 19106,
		Ftype1: int8('F'),
	},
	2761: {
		Fword:  __ccgo_ts + 19113,
		Ftype1: int8('F'),
	},
	2762: {
		Fword:  __ccgo_ts + 19120,
		Ftype1: int8('F'),
	},
	2763: {
		Fword:  __ccgo_ts + 19127,
		Ftype1: int8('F'),
	},
	2764: {
		Fword:  __ccgo_ts + 19134,
		Ftype1: int8('F'),
	},
	2765: {
		Fword:  __ccgo_ts + 19141,
		Ftype1: int8('F'),
	},
	2766: {
		Fword:  __ccgo_ts + 19148,
		Ftype1: int8('F'),
	},
	2767: {
		Fword:  __ccgo_ts + 19155,
		Ftype1: int8('F'),
	},
	2768: {
		Fword:  __ccgo_ts + 19162,
		Ftype1: int8('F'),
	},
	2769: {
		Fword:  __ccgo_ts + 19169,
		Ftype1: int8('F'),
	},
	2770: {
		Fword:  __ccgo_ts + 19176,
		Ftype1: int8('F'),
	},
	2771: {
		Fword:  __ccgo_ts + 19183,
		Ftype1: int8('F'),
	},
	2772: {
		Fword:  __ccgo_ts + 19190,
		Ftype1: int8('F'),
	},
	2773: {
		Fword:  __ccgo_ts + 19197,
		Ftype1: int8('F'),
	},
	2774: {
		Fword:  __ccgo_ts + 19204,
		Ftype1: int8('F'),
	},
	2775: {
		Fword:  __ccgo_ts + 19211,
		Ftype1: int8('F'),
	},
	2776: {
		Fword:  __ccgo_ts + 19218,
		Ftype1: int8('F'),
	},
	2777: {
		Fword:  __ccgo_ts + 19225,
		Ftype1: int8('F'),
	},
	2778: {
		Fword:  __ccgo_ts + 19232,
		Ftype1: int8('F'),
	},
	2779: {
		Fword:  __ccgo_ts + 19239,
		Ftype1: int8('F'),
	},
	2780: {
		Fword:  __ccgo_ts + 19246,
		Ftype1: int8('F'),
	},
	2781: {
		Fword:  __ccgo_ts + 19253,
		Ftype1: int8('F'),
	},
	2782: {
		Fword:  __ccgo_ts + 19260,
		Ftype1: int8('F'),
	},
	2783: {
		Fword:  __ccgo_ts + 19265,
		Ftype1: int8('F'),
	},
	2784: {
		Fword:  __ccgo_ts + 19271,
		Ftype1: int8('F'),
	},
	2785: {
		Fword:  __ccgo_ts + 19278,
		Ftype1: int8('F'),
	},
	2786: {
		Fword:  __ccgo_ts + 19284,
		Ftype1: int8('F'),
	},
	2787: {
		Fword:  __ccgo_ts + 19291,
		Ftype1: int8('F'),
	},
	2788: {
		Fword:  __ccgo_ts + 19298,
		Ftype1: int8('F'),
	},
	2789: {
		Fword:  __ccgo_ts + 19305,
		Ftype1: int8('F'),
	},
	2790: {
		Fword:  __ccgo_ts + 19312,
		Ftype1: int8('F'),
	},
	2791: {
		Fword:  __ccgo_ts + 19319,
		Ftype1: int8('F'),
	},
	2792: {
		Fword:  __ccgo_ts + 19326,
		Ftype1: int8('F'),
	},
	2793: {
		Fword:  __ccgo_ts + 19333,
		Ftype1: int8('F'),
	},
	2794: {
		Fword:  __ccgo_ts + 19340,
		Ftype1: int8('F'),
	},
	2795: {
		Fword:  __ccgo_ts + 19347,
		Ftype1: int8('F'),
	},
	2796: {
		Fword:  __ccgo_ts + 19354,
		Ftype1: int8('F'),
	},
	2797: {
		Fword:  __ccgo_ts + 19361,
		Ftype1: int8('F'),
	},
	2798: {
		Fword:  __ccgo_ts + 19368,
		Ftype1: int8('F'),
	},
	2799: {
		Fword:  __ccgo_ts + 19375,
		Ftype1: int8('F'),
	},
	2800: {
		Fword:  __ccgo_ts + 19382,
		Ftype1: int8('F'),
	},
	2801: {
		Fword:  __ccgo_ts + 19389,
		Ftype1: int8('F'),
	},
	2802: {
		Fword:  __ccgo_ts + 19396,
		Ftype1: int8('F'),
	},
	2803: {
		Fword:  __ccgo_ts + 19403,
		Ftype1: int8('F'),
	},
	2804: {
		Fword:  __ccgo_ts + 19410,
		Ftype1: int8('F'),
	},
	2805: {
		Fword:  __ccgo_ts + 19417,
		Ftype1: int8('F'),
	},
	2806: {
		Fword:  __ccgo_ts + 19424,
		Ftype1: int8('F'),
	},
	2807: {
		Fword:  __ccgo_ts + 19431,
		Ftype1: int8('F'),
	},
	2808: {
		Fword:  __ccgo_ts + 19438,
		Ftype1: int8('F'),
	},
	2809: {
		Fword:  __ccgo_ts + 19445,
		Ftype1: int8('F'),
	},
	2810: {
		Fword:  __ccgo_ts + 19452,
		Ftype1: int8('F'),
	},
	2811: {
		Fword:  __ccgo_ts + 19459,
		Ftype1: int8('F'),
	},
	2812: {
		Fword:  __ccgo_ts + 19466,
		Ftype1: int8('F'),
	},
	2813: {
		Fword:  __ccgo_ts + 19473,
		Ftype1: int8('F'),
	},
	2814: {
		Fword:  __ccgo_ts + 19480,
		Ftype1: int8('F'),
	},
	2815: {
		Fword:  __ccgo_ts + 19487,
		Ftype1: int8('F'),
	},
	2816: {
		Fword:  __ccgo_ts + 19494,
		Ftype1: int8('F'),
	},
	2817: {
		Fword:  __ccgo_ts + 19501,
		Ftype1: int8('F'),
	},
	2818: {
		Fword:  __ccgo_ts + 19508,
		Ftype1: int8('F'),
	},
	2819: {
		Fword:  __ccgo_ts + 19515,
		Ftype1: int8('F'),
	},
	2820: {
		Fword:  __ccgo_ts + 19522,
		Ftype1: int8('F'),
	},
	2821: {
		Fword:  __ccgo_ts + 19529,
		Ftype1: int8('F'),
	},
	2822: {
		Fword:  __ccgo_ts + 19536,
		Ftype1: int8('F'),
	},
	2823: {
		Fword:  __ccgo_ts + 19543,
		Ftype1: int8('F'),
	},
	2824: {
		Fword:  __ccgo_ts + 19550,
		Ftype1: int8('F'),
	},
	2825: {
		Fword:  __ccgo_ts + 19557,
		Ftype1: int8('F'),
	},
	2826: {
		Fword:  __ccgo_ts + 19564,
		Ftype1: int8('F'),
	},
	2827: {
		Fword:  __ccgo_ts + 19571,
		Ftype1: int8('F'),
	},
	2828: {
		Fword:  __ccgo_ts + 19578,
		Ftype1: int8('F'),
	},
	2829: {
		Fword:  __ccgo_ts + 19585,
		Ftype1: int8('F'),
	},
	2830: {
		Fword:  __ccgo_ts + 19592,
		Ftype1: int8('F'),
	},
	2831: {
		Fword:  __ccgo_ts + 19599,
		Ftype1: int8('F'),
	},
	2832: {
		Fword:  __ccgo_ts + 19606,
		Ftype1: int8('F'),
	},
	2833: {
		Fword:  __ccgo_ts + 19613,
		Ftype1: int8('F'),
	},
	2834: {
		Fword:  __ccgo_ts + 19620,
		Ftype1: int8('F'),
	},
	2835: {
		Fword:  __ccgo_ts + 19627,
		Ftype1: int8('F'),
	},
	2836: {
		Fword:  __ccgo_ts + 19634,
		Ftype1: int8('F'),
	},
	2837: {
		Fword:  __ccgo_ts + 19640,
		Ftype1: int8('F'),
	},
	2838: {
		Fword:  __ccgo_ts + 19647,
		Ftype1: int8('F'),
	},
	2839: {
		Fword:  __ccgo_ts + 19654,
		Ftype1: int8('F'),
	},
	2840: {
		Fword:  __ccgo_ts + 19661,
		Ftype1: int8('F'),
	},
	2841: {
		Fword:  __ccgo_ts + 19668,
		Ftype1: int8('F'),
	},
	2842: {
		Fword:  __ccgo_ts + 19675,
		Ftype1: int8('F'),
	},
	2843: {
		Fword:  __ccgo_ts + 19682,
		Ftype1: int8('F'),
	},
	2844: {
		Fword:  __ccgo_ts + 19689,
		Ftype1: int8('F'),
	},
	2845: {
		Fword:  __ccgo_ts + 19696,
		Ftype1: int8('F'),
	},
	2846: {
		Fword:  __ccgo_ts + 19703,
		Ftype1: int8('F'),
	},
	2847: {
		Fword:  __ccgo_ts + 19710,
		Ftype1: int8('F'),
	},
	2848: {
		Fword:  __ccgo_ts + 19717,
		Ftype1: int8('F'),
	},
	2849: {
		Fword:  __ccgo_ts + 19724,
		Ftype1: int8('F'),
	},
	2850: {
		Fword:  __ccgo_ts + 19731,
		Ftype1: int8('F'),
	},
	2851: {
		Fword:  __ccgo_ts + 19738,
		Ftype1: int8('F'),
	},
	2852: {
		Fword:  __ccgo_ts + 19745,
		Ftype1: int8('F'),
	},
	2853: {
		Fword:  __ccgo_ts + 19752,
		Ftype1: int8('F'),
	},
	2854: {
		Fword:  __ccgo_ts + 19759,
		Ftype1: int8('F'),
	},
	2855: {
		Fword:  __ccgo_ts + 19766,
		Ftype1: int8('F'),
	},
	2856: {
		Fword:  __ccgo_ts + 19773,
		Ftype1: int8('F'),
	},
	2857: {
		Fword:  __ccgo_ts + 19780,
		Ftype1: int8('F'),
	},
	2858: {
		Fword:  __ccgo_ts + 19787,
		Ftype1: int8('F'),
	},
	2859: {
		Fword:  __ccgo_ts + 19794,
		Ftype1: int8('F'),
	},
	2860: {
		Fword:  __ccgo_ts + 19801,
		Ftype1: int8('F'),
	},
	2861: {
		Fword:  __ccgo_ts + 19808,
		Ftype1: int8('F'),
	},
	2862: {
		Fword:  __ccgo_ts + 19815,
		Ftype1: int8('F'),
	},
	2863: {
		Fword:  __ccgo_ts + 19822,
		Ftype1: int8('F'),
	},
	2864: {
		Fword:  __ccgo_ts + 19829,
		Ftype1: int8('F'),
	},
	2865: {
		Fword:  __ccgo_ts + 19836,
		Ftype1: int8('F'),
	},
	2866: {
		Fword:  __ccgo_ts + 19843,
		Ftype1: int8('F'),
	},
	2867: {
		Fword:  __ccgo_ts + 19850,
		Ftype1: int8('F'),
	},
	2868: {
		Fword:  __ccgo_ts + 19857,
		Ftype1: int8('F'),
	},
	2869: {
		Fword:  __ccgo_ts + 19864,
		Ftype1: int8('F'),
	},
	2870: {
		Fword:  __ccgo_ts + 19871,
		Ftype1: int8('F'),
	},
	2871: {
		Fword:  __ccgo_ts + 19878,
		Ftype1: int8('F'),
	},
	2872: {
		Fword:  __ccgo_ts + 19885,
		Ftype1: int8('F'),
	},
	2873: {
		Fword:  __ccgo_ts + 19892,
		Ftype1: int8('F'),
	},
	2874: {
		Fword:  __ccgo_ts + 19899,
		Ftype1: int8('F'),
	},
	2875: {
		Fword:  __ccgo_ts + 19906,
		Ftype1: int8('F'),
	},
	2876: {
		Fword:  __ccgo_ts + 19913,
		Ftype1: int8('F'),
	},
	2877: {
		Fword:  __ccgo_ts + 19920,
		Ftype1: int8('F'),
	},
	2878: {
		Fword:  __ccgo_ts + 19927,
		Ftype1: int8('F'),
	},
	2879: {
		Fword:  __ccgo_ts + 19934,
		Ftype1: int8('F'),
	},
	2880: {
		Fword:  __ccgo_ts + 19941,
		Ftype1: int8('F'),
	},
	2881: {
		Fword:  __ccgo_ts + 19948,
		Ftype1: int8('F'),
	},
	2882: {
		Fword:  __ccgo_ts + 19955,
		Ftype1: int8('F'),
	},
	2883: {
		Fword:  __ccgo_ts + 19962,
		Ftype1: int8('F'),
	},
	2884: {
		Fword:  __ccgo_ts + 19969,
		Ftype1: int8('F'),
	},
	2885: {
		Fword:  __ccgo_ts + 19976,
		Ftype1: int8('F'),
	},
	2886: {
		Fword:  __ccgo_ts + 19983,
		Ftype1: int8('F'),
	},
	2887: {
		Fword:  __ccgo_ts + 19989,
		Ftype1: int8('F'),
	},
	2888: {
		Fword:  __ccgo_ts + 19996,
		Ftype1: int8('F'),
	},
	2889: {
		Fword:  __ccgo_ts + 20003,
		Ftype1: int8('F'),
	},
	2890: {
		Fword:  __ccgo_ts + 20010,
		Ftype1: int8('F'),
	},
	2891: {
		Fword:  __ccgo_ts + 20017,
		Ftype1: int8('F'),
	},
	2892: {
		Fword:  __ccgo_ts + 20024,
		Ftype1: int8('F'),
	},
	2893: {
		Fword:  __ccgo_ts + 20031,
		Ftype1: int8('F'),
	},
	2894: {
		Fword:  __ccgo_ts + 20038,
		Ftype1: int8('F'),
	},
	2895: {
		Fword:  __ccgo_ts + 20045,
		Ftype1: int8('F'),
	},
	2896: {
		Fword:  __ccgo_ts + 20052,
		Ftype1: int8('F'),
	},
	2897: {
		Fword:  __ccgo_ts + 20059,
		Ftype1: int8('F'),
	},
	2898: {
		Fword:  __ccgo_ts + 20066,
		Ftype1: int8('F'),
	},
	2899: {
		Fword:  __ccgo_ts + 20073,
		Ftype1: int8('F'),
	},
	2900: {
		Fword:  __ccgo_ts + 20080,
		Ftype1: int8('F'),
	},
	2901: {
		Fword:  __ccgo_ts + 20087,
		Ftype1: int8('F'),
	},
	2902: {
		Fword:  __ccgo_ts + 20094,
		Ftype1: int8('F'),
	},
	2903: {
		Fword:  __ccgo_ts + 20101,
		Ftype1: int8('F'),
	},
	2904: {
		Fword:  __ccgo_ts + 20108,
		Ftype1: int8('F'),
	},
	2905: {
		Fword:  __ccgo_ts + 20115,
		Ftype1: int8('F'),
	},
	2906: {
		Fword:  __ccgo_ts + 20122,
		Ftype1: int8('F'),
	},
	2907: {
		Fword:  __ccgo_ts + 20129,
		Ftype1: int8('F'),
	},
	2908: {
		Fword:  __ccgo_ts + 20136,
		Ftype1: int8('F'),
	},
	2909: {
		Fword:  __ccgo_ts + 20143,
		Ftype1: int8('F'),
	},
	2910: {
		Fword:  __ccgo_ts + 20149,
		Ftype1: int8('F'),
	},
	2911: {
		Fword:  __ccgo_ts + 20156,
		Ftype1: int8('F'),
	},
	2912: {
		Fword:  __ccgo_ts + 20163,
		Ftype1: int8('F'),
	},
	2913: {
		Fword:  __ccgo_ts + 20170,
		Ftype1: int8('F'),
	},
	2914: {
		Fword:  __ccgo_ts + 20177,
		Ftype1: int8('F'),
	},
	2915: {
		Fword:  __ccgo_ts + 20184,
		Ftype1: int8('F'),
	},
	2916: {
		Fword:  __ccgo_ts + 20191,
		Ftype1: int8('F'),
	},
	2917: {
		Fword:  __ccgo_ts + 20198,
		Ftype1: int8('F'),
	},
	2918: {
		Fword:  __ccgo_ts + 20205,
		Ftype1: int8('F'),
	},
	2919: {
		Fword:  __ccgo_ts + 20212,
		Ftype1: int8('F'),
	},
	2920: {
		Fword:  __ccgo_ts + 20219,
		Ftype1: int8('F'),
	},
	2921: {
		Fword:  __ccgo_ts + 20226,
		Ftype1: int8('F'),
	},
	2922: {
		Fword:  __ccgo_ts + 20233,
		Ftype1: int8('F'),
	},
	2923: {
		Fword:  __ccgo_ts + 20240,
		Ftype1: int8('F'),
	},
	2924: {
		Fword:  __ccgo_ts + 20247,
		Ftype1: int8('F'),
	},
	2925: {
		Fword:  __ccgo_ts + 20254,
		Ftype1: int8('F'),
	},
	2926: {
		Fword:  __ccgo_ts + 20261,
		Ftype1: int8('F'),
	},
	2927: {
		Fword:  __ccgo_ts + 20268,
		Ftype1: int8('F'),
	},
	2928: {
		Fword:  __ccgo_ts + 20275,
		Ftype1: int8('F'),
	},
	2929: {
		Fword:  __ccgo_ts + 20282,
		Ftype1: int8('F'),
	},
	2930: {
		Fword:  __ccgo_ts + 20289,
		Ftype1: int8('F'),
	},
	2931: {
		Fword:  __ccgo_ts + 20296,
		Ftype1: int8('F'),
	},
	2932: {
		Fword:  __ccgo_ts + 20303,
		Ftype1: int8('F'),
	},
	2933: {
		Fword:  __ccgo_ts + 20310,
		Ftype1: int8('F'),
	},
	2934: {
		Fword:  __ccgo_ts + 20317,
		Ftype1: int8('F'),
	},
	2935: {
		Fword:  __ccgo_ts + 20324,
		Ftype1: int8('F'),
	},
	2936: {
		Fword:  __ccgo_ts + 20331,
		Ftype1: int8('F'),
	},
	2937: {
		Fword:  __ccgo_ts + 20338,
		Ftype1: int8('F'),
	},
	2938: {
		Fword:  __ccgo_ts + 20345,
		Ftype1: int8('F'),
	},
	2939: {
		Fword:  __ccgo_ts + 20352,
		Ftype1: int8('F'),
	},
	2940: {
		Fword:  __ccgo_ts + 20359,
		Ftype1: int8('F'),
	},
	2941: {
		Fword:  __ccgo_ts + 20366,
		Ftype1: int8('F'),
	},
	2942: {
		Fword:  __ccgo_ts + 20373,
		Ftype1: int8('F'),
	},
	2943: {
		Fword:  __ccgo_ts + 20380,
		Ftype1: int8('F'),
	},
	2944: {
		Fword:  __ccgo_ts + 20387,
		Ftype1: int8('F'),
	},
	2945: {
		Fword:  __ccgo_ts + 20394,
		Ftype1: int8('F'),
	},
	2946: {
		Fword:  __ccgo_ts + 20401,
		Ftype1: int8('F'),
	},
	2947: {
		Fword:  __ccgo_ts + 20408,
		Ftype1: int8('F'),
	},
	2948: {
		Fword:  __ccgo_ts + 20415,
		Ftype1: int8('F'),
	},
	2949: {
		Fword:  __ccgo_ts + 20422,
		Ftype1: int8('F'),
	},
	2950: {
		Fword:  __ccgo_ts + 20429,
		Ftype1: int8('F'),
	},
	2951: {
		Fword:  __ccgo_ts + 20436,
		Ftype1: int8('F'),
	},
	2952: {
		Fword:  __ccgo_ts + 20443,
		Ftype1: int8('F'),
	},
	2953: {
		Fword:  __ccgo_ts + 20450,
		Ftype1: int8('F'),
	},
	2954: {
		Fword:  __ccgo_ts + 20457,
		Ftype1: int8('F'),
	},
	2955: {
		Fword:  __ccgo_ts + 20464,
		Ftype1: int8('F'),
	},
	2956: {
		Fword:  __ccgo_ts + 20471,
		Ftype1: int8('F'),
	},
	2957: {
		Fword:  __ccgo_ts + 20478,
		Ftype1: int8('F'),
	},
	2958: {
		Fword:  __ccgo_ts + 20485,
		Ftype1: int8('F'),
	},
	2959: {
		Fword:  __ccgo_ts + 20492,
		Ftype1: int8('F'),
	},
	2960: {
		Fword:  __ccgo_ts + 20499,
		Ftype1: int8('F'),
	},
	2961: {
		Fword:  __ccgo_ts + 20506,
		Ftype1: int8('F'),
	},
	2962: {
		Fword:  __ccgo_ts + 20513,
		Ftype1: int8('F'),
	},
	2963: {
		Fword:  __ccgo_ts + 20520,
		Ftype1: int8('F'),
	},
	2964: {
		Fword:  __ccgo_ts + 20527,
		Ftype1: int8('F'),
	},
	2965: {
		Fword:  __ccgo_ts + 20534,
		Ftype1: int8('F'),
	},
	2966: {
		Fword:  __ccgo_ts + 20541,
		Ftype1: int8('F'),
	},
	2967: {
		Fword:  __ccgo_ts + 20548,
		Ftype1: int8('F'),
	},
	2968: {
		Fword:  __ccgo_ts + 20555,
		Ftype1: int8('F'),
	},
	2969: {
		Fword:  __ccgo_ts + 20562,
		Ftype1: int8('F'),
	},
	2970: {
		Fword:  __ccgo_ts + 20569,
		Ftype1: int8('F'),
	},
	2971: {
		Fword:  __ccgo_ts + 20576,
		Ftype1: int8('F'),
	},
	2972: {
		Fword:  __ccgo_ts + 20583,
		Ftype1: int8('F'),
	},
	2973: {
		Fword:  __ccgo_ts + 20590,
		Ftype1: int8('F'),
	},
	2974: {
		Fword:  __ccgo_ts + 20597,
		Ftype1: int8('F'),
	},
	2975: {
		Fword:  __ccgo_ts + 20604,
		Ftype1: int8('F'),
	},
	2976: {
		Fword:  __ccgo_ts + 20611,
		Ftype1: int8('F'),
	},
	2977: {
		Fword:  __ccgo_ts + 20618,
		Ftype1: int8('F'),
	},
	2978: {
		Fword:  __ccgo_ts + 20625,
		Ftype1: int8('F'),
	},
	2979: {
		Fword:  __ccgo_ts + 20632,
		Ftype1: int8('F'),
	},
	2980: {
		Fword:  __ccgo_ts + 20638,
		Ftype1: int8('F'),
	},
	2981: {
		Fword:  __ccgo_ts + 20645,
		Ftype1: int8('F'),
	},
	2982: {
		Fword:  __ccgo_ts + 20652,
		Ftype1: int8('F'),
	},
	2983: {
		Fword:  __ccgo_ts + 20659,
		Ftype1: int8('F'),
	},
	2984: {
		Fword:  __ccgo_ts + 20666,
		Ftype1: int8('F'),
	},
	2985: {
		Fword:  __ccgo_ts + 20673,
		Ftype1: int8('F'),
	},
	2986: {
		Fword:  __ccgo_ts + 20680,
		Ftype1: int8('F'),
	},
	2987: {
		Fword:  __ccgo_ts + 20687,
		Ftype1: int8('F'),
	},
	2988: {
		Fword:  __ccgo_ts + 20694,
		Ftype1: int8('F'),
	},
	2989: {
		Fword:  __ccgo_ts + 20701,
		Ftype1: int8('F'),
	},
	2990: {
		Fword:  __ccgo_ts + 20708,
		Ftype1: int8('F'),
	},
	2991: {
		Fword:  __ccgo_ts + 20715,
		Ftype1: int8('F'),
	},
	2992: {
		Fword:  __ccgo_ts + 20722,
		Ftype1: int8('F'),
	},
	2993: {
		Fword:  __ccgo_ts + 20729,
		Ftype1: int8('F'),
	},
	2994: {
		Fword:  __ccgo_ts + 20736,
		Ftype1: int8('F'),
	},
	2995: {
		Fword:  __ccgo_ts + 20743,
		Ftype1: int8('F'),
	},
	2996: {
		Fword:  __ccgo_ts + 20750,
		Ftype1: int8('F'),
	},
	2997: {
		Fword:  __ccgo_ts + 20757,
		Ftype1: int8('F'),
	},
	2998: {
		Fword:  __ccgo_ts + 20764,
		Ftype1: int8('F'),
	},
	2999: {
		Fword:  __ccgo_ts + 20771,
		Ftype1: int8('F'),
	},
	3000: {
		Fword:  __ccgo_ts + 20778,
		Ftype1: int8('F'),
	},
	3001: {
		Fword:  __ccgo_ts + 20785,
		Ftype1: int8('F'),
	},
	3002: {
		Fword:  __ccgo_ts + 20792,
		Ftype1: int8('F'),
	},
	3003: {
		Fword:  __ccgo_ts + 20799,
		Ftype1: int8('F'),
	},
	3004: {
		Fword:  __ccgo_ts + 20806,
		Ftype1: int8('F'),
	},
	3005: {
		Fword:  __ccgo_ts + 20813,
		Ftype1: int8('F'),
	},
	3006: {
		Fword:  __ccgo_ts + 20820,
		Ftype1: int8('F'),
	},
	3007: {
		Fword:  __ccgo_ts + 20827,
		Ftype1: int8('F'),
	},
	3008: {
		Fword:  __ccgo_ts + 20834,
		Ftype1: int8('F'),
	},
	3009: {
		Fword:  __ccgo_ts + 20840,
		Ftype1: int8('F'),
	},
	3010: {
		Fword:  __ccgo_ts + 20847,
		Ftype1: int8('F'),
	},
	3011: {
		Fword:  __ccgo_ts + 20854,
		Ftype1: int8('F'),
	},
	3012: {
		Fword:  __ccgo_ts + 20861,
		Ftype1: int8('F'),
	},
	3013: {
		Fword:  __ccgo_ts + 20868,
		Ftype1: int8('F'),
	},
	3014: {
		Fword:  __ccgo_ts + 20875,
		Ftype1: int8('F'),
	},
	3015: {
		Fword:  __ccgo_ts + 20882,
		Ftype1: int8('F'),
	},
	3016: {
		Fword:  __ccgo_ts + 20889,
		Ftype1: int8('F'),
	},
	3017: {
		Fword:  __ccgo_ts + 20896,
		Ftype1: int8('F'),
	},
	3018: {
		Fword:  __ccgo_ts + 20903,
		Ftype1: int8('F'),
	},
	3019: {
		Fword:  __ccgo_ts + 20910,
		Ftype1: int8('F'),
	},
	3020: {
		Fword:  __ccgo_ts + 20917,
		Ftype1: int8('F'),
	},
	3021: {
		Fword:  __ccgo_ts + 20924,
		Ftype1: int8('F'),
	},
	3022: {
		Fword:  __ccgo_ts + 20931,
		Ftype1: int8('F'),
	},
	3023: {
		Fword:  __ccgo_ts + 20938,
		Ftype1: int8('F'),
	},
	3024: {
		Fword:  __ccgo_ts + 20945,
		Ftype1: int8('F'),
	},
	3025: {
		Fword:  __ccgo_ts + 20952,
		Ftype1: int8('F'),
	},
	3026: {
		Fword:  __ccgo_ts + 20959,
		Ftype1: int8('F'),
	},
	3027: {
		Fword:  __ccgo_ts + 20966,
		Ftype1: int8('F'),
	},
	3028: {
		Fword:  __ccgo_ts + 20973,
		Ftype1: int8('F'),
	},
	3029: {
		Fword:  __ccgo_ts + 20980,
		Ftype1: int8('F'),
	},
	3030: {
		Fword:  __ccgo_ts + 20987,
		Ftype1: int8('F'),
	},
	3031: {
		Fword:  __ccgo_ts + 20994,
		Ftype1: int8('F'),
	},
	3032: {
		Fword:  __ccgo_ts + 21001,
		Ftype1: int8('F'),
	},
	3033: {
		Fword:  __ccgo_ts + 21008,
		Ftype1: int8('F'),
	},
	3034: {
		Fword:  __ccgo_ts + 21015,
		Ftype1: int8('F'),
	},
	3035: {
		Fword:  __ccgo_ts + 21022,
		Ftype1: int8('F'),
	},
	3036: {
		Fword:  __ccgo_ts + 21029,
		Ftype1: int8('F'),
	},
	3037: {
		Fword:  __ccgo_ts + 21036,
		Ftype1: int8('F'),
	},
	3038: {
		Fword:  __ccgo_ts + 21043,
		Ftype1: int8('F'),
	},
	3039: {
		Fword:  __ccgo_ts + 21050,
		Ftype1: int8('F'),
	},
	3040: {
		Fword:  __ccgo_ts + 21057,
		Ftype1: int8('F'),
	},
	3041: {
		Fword:  __ccgo_ts + 21064,
		Ftype1: int8('F'),
	},
	3042: {
		Fword:  __ccgo_ts + 21070,
		Ftype1: int8('F'),
	},
	3043: {
		Fword:  __ccgo_ts + 21077,
		Ftype1: int8('F'),
	},
	3044: {
		Fword:  __ccgo_ts + 21084,
		Ftype1: int8('F'),
	},
	3045: {
		Fword:  __ccgo_ts + 21091,
		Ftype1: int8('F'),
	},
	3046: {
		Fword:  __ccgo_ts + 21098,
		Ftype1: int8('F'),
	},
	3047: {
		Fword:  __ccgo_ts + 21105,
		Ftype1: int8('F'),
	},
	3048: {
		Fword:  __ccgo_ts + 21112,
		Ftype1: int8('F'),
	},
	3049: {
		Fword:  __ccgo_ts + 21119,
		Ftype1: int8('F'),
	},
	3050: {
		Fword:  __ccgo_ts + 21126,
		Ftype1: int8('F'),
	},
	3051: {
		Fword:  __ccgo_ts + 21133,
		Ftype1: int8('F'),
	},
	3052: {
		Fword:  __ccgo_ts + 21140,
		Ftype1: int8('F'),
	},
	3053: {
		Fword:  __ccgo_ts + 21147,
		Ftype1: int8('F'),
	},
	3054: {
		Fword:  __ccgo_ts + 21154,
		Ftype1: int8('F'),
	},
	3055: {
		Fword:  __ccgo_ts + 21161,
		Ftype1: int8('F'),
	},
	3056: {
		Fword:  __ccgo_ts + 21168,
		Ftype1: int8('F'),
	},
	3057: {
		Fword:  __ccgo_ts + 21175,
		Ftype1: int8('F'),
	},
	3058: {
		Fword:  __ccgo_ts + 21182,
		Ftype1: int8('F'),
	},
	3059: {
		Fword:  __ccgo_ts + 21189,
		Ftype1: int8('F'),
	},
	3060: {
		Fword:  __ccgo_ts + 21196,
		Ftype1: int8('F'),
	},
	3061: {
		Fword:  __ccgo_ts + 21203,
		Ftype1: int8('F'),
	},
	3062: {
		Fword:  __ccgo_ts + 21210,
		Ftype1: int8('F'),
	},
	3063: {
		Fword:  __ccgo_ts + 21217,
		Ftype1: int8('F'),
	},
	3064: {
		Fword:  __ccgo_ts + 21224,
		Ftype1: int8('F'),
	},
	3065: {
		Fword:  __ccgo_ts + 21231,
		Ftype1: int8('F'),
	},
	3066: {
		Fword:  __ccgo_ts + 21238,
		Ftype1: int8('F'),
	},
	3067: {
		Fword:  __ccgo_ts + 21245,
		Ftype1: int8('F'),
	},
	3068: {
		Fword:  __ccgo_ts + 21252,
		Ftype1: int8('F'),
	},
	3069: {
		Fword:  __ccgo_ts + 21259,
		Ftype1: int8('F'),
	},
	3070: {
		Fword:  __ccgo_ts + 21266,
		Ftype1: int8('F'),
	},
	3071: {
		Fword:  __ccgo_ts + 21273,
		Ftype1: int8('F'),
	},
	3072: {
		Fword:  __ccgo_ts + 21280,
		Ftype1: int8('F'),
	},
	3073: {
		Fword:  __ccgo_ts + 21287,
		Ftype1: int8('F'),
	},
	3074: {
		Fword:  __ccgo_ts + 21294,
		Ftype1: int8('F'),
	},
	3075: {
		Fword:  __ccgo_ts + 21301,
		Ftype1: int8('F'),
	},
	3076: {
		Fword:  __ccgo_ts + 21308,
		Ftype1: int8('F'),
	},
	3077: {
		Fword:  __ccgo_ts + 21315,
		Ftype1: int8('F'),
	},
	3078: {
		Fword:  __ccgo_ts + 21322,
		Ftype1: int8('F'),
	},
	3079: {
		Fword:  __ccgo_ts + 21329,
		Ftype1: int8('F'),
	},
	3080: {
		Fword:  __ccgo_ts + 21336,
		Ftype1: int8('F'),
	},
	3081: {
		Fword:  __ccgo_ts + 21343,
		Ftype1: int8('F'),
	},
	3082: {
		Fword:  __ccgo_ts + 21350,
		Ftype1: int8('F'),
	},
	3083: {
		Fword:  __ccgo_ts + 21357,
		Ftype1: int8('F'),
	},
	3084: {
		Fword:  __ccgo_ts + 21364,
		Ftype1: int8('F'),
	},
	3085: {
		Fword:  __ccgo_ts + 21371,
		Ftype1: int8('F'),
	},
	3086: {
		Fword:  __ccgo_ts + 21378,
		Ftype1: int8('F'),
	},
	3087: {
		Fword:  __ccgo_ts + 21385,
		Ftype1: int8('F'),
	},
	3088: {
		Fword:  __ccgo_ts + 21392,
		Ftype1: int8('F'),
	},
	3089: {
		Fword:  __ccgo_ts + 21399,
		Ftype1: int8('F'),
	},
	3090: {
		Fword:  __ccgo_ts + 21406,
		Ftype1: int8('F'),
	},
	3091: {
		Fword:  __ccgo_ts + 21413,
		Ftype1: int8('F'),
	},
	3092: {
		Fword:  __ccgo_ts + 21420,
		Ftype1: int8('F'),
	},
	3093: {
		Fword:  __ccgo_ts + 21427,
		Ftype1: int8('F'),
	},
	3094: {
		Fword:  __ccgo_ts + 21434,
		Ftype1: int8('F'),
	},
	3095: {
		Fword:  __ccgo_ts + 21441,
		Ftype1: int8('F'),
	},
	3096: {
		Fword:  __ccgo_ts + 21448,
		Ftype1: int8('F'),
	},
	3097: {
		Fword:  __ccgo_ts + 21455,
		Ftype1: int8('F'),
	},
	3098: {
		Fword:  __ccgo_ts + 21462,
		Ftype1: int8('F'),
	},
	3099: {
		Fword:  __ccgo_ts + 21469,
		Ftype1: int8('F'),
	},
	3100: {
		Fword:  __ccgo_ts + 21476,
		Ftype1: int8('F'),
	},
	3101: {
		Fword:  __ccgo_ts + 21483,
		Ftype1: int8('F'),
	},
	3102: {
		Fword:  __ccgo_ts + 21490,
		Ftype1: int8('F'),
	},
	3103: {
		Fword:  __ccgo_ts + 21497,
		Ftype1: int8('F'),
	},
	3104: {
		Fword:  __ccgo_ts + 21504,
		Ftype1: int8('F'),
	},
	3105: {
		Fword:  __ccgo_ts + 21511,
		Ftype1: int8('F'),
	},
	3106: {
		Fword:  __ccgo_ts + 21518,
		Ftype1: int8('F'),
	},
	3107: {
		Fword:  __ccgo_ts + 21525,
		Ftype1: int8('F'),
	},
	3108: {
		Fword:  __ccgo_ts + 21532,
		Ftype1: int8('F'),
	},
	3109: {
		Fword:  __ccgo_ts + 21539,
		Ftype1: int8('F'),
	},
	3110: {
		Fword:  __ccgo_ts + 21546,
		Ftype1: int8('F'),
	},
	3111: {
		Fword:  __ccgo_ts + 21553,
		Ftype1: int8('F'),
	},
	3112: {
		Fword:  __ccgo_ts + 21560,
		Ftype1: int8('F'),
	},
	3113: {
		Fword:  __ccgo_ts + 21567,
		Ftype1: int8('F'),
	},
	3114: {
		Fword:  __ccgo_ts + 21574,
		Ftype1: int8('F'),
	},
	3115: {
		Fword:  __ccgo_ts + 21581,
		Ftype1: int8('F'),
	},
	3116: {
		Fword:  __ccgo_ts + 21588,
		Ftype1: int8('F'),
	},
	3117: {
		Fword:  __ccgo_ts + 21595,
		Ftype1: int8('F'),
	},
	3118: {
		Fword:  __ccgo_ts + 21602,
		Ftype1: int8('F'),
	},
	3119: {
		Fword:  __ccgo_ts + 21609,
		Ftype1: int8('F'),
	},
	3120: {
		Fword:  __ccgo_ts + 21616,
		Ftype1: int8('F'),
	},
	3121: {
		Fword:  __ccgo_ts + 21623,
		Ftype1: int8('F'),
	},
	3122: {
		Fword:  __ccgo_ts + 21630,
		Ftype1: int8('F'),
	},
	3123: {
		Fword:  __ccgo_ts + 21637,
		Ftype1: int8('F'),
	},
	3124: {
		Fword:  __ccgo_ts + 21644,
		Ftype1: int8('F'),
	},
	3125: {
		Fword:  __ccgo_ts + 21651,
		Ftype1: int8('F'),
	},
	3126: {
		Fword:  __ccgo_ts + 21658,
		Ftype1: int8('F'),
	},
	3127: {
		Fword:  __ccgo_ts + 21665,
		Ftype1: int8('F'),
	},
	3128: {
		Fword:  __ccgo_ts + 21672,
		Ftype1: int8('F'),
	},
	3129: {
		Fword:  __ccgo_ts + 21679,
		Ftype1: int8('F'),
	},
	3130: {
		Fword:  __ccgo_ts + 21686,
		Ftype1: int8('F'),
	},
	3131: {
		Fword:  __ccgo_ts + 21693,
		Ftype1: int8('F'),
	},
	3132: {
		Fword:  __ccgo_ts + 21700,
		Ftype1: int8('F'),
	},
	3133: {
		Fword:  __ccgo_ts + 21707,
		Ftype1: int8('F'),
	},
	3134: {
		Fword:  __ccgo_ts + 21714,
		Ftype1: int8('F'),
	},
	3135: {
		Fword:  __ccgo_ts + 21721,
		Ftype1: int8('F'),
	},
	3136: {
		Fword:  __ccgo_ts + 21728,
		Ftype1: int8('F'),
	},
	3137: {
		Fword:  __ccgo_ts + 21735,
		Ftype1: int8('F'),
	},
	3138: {
		Fword:  __ccgo_ts + 21742,
		Ftype1: int8('F'),
	},
	3139: {
		Fword:  __ccgo_ts + 21749,
		Ftype1: int8('F'),
	},
	3140: {
		Fword:  __ccgo_ts + 21756,
		Ftype1: int8('F'),
	},
	3141: {
		Fword:  __ccgo_ts + 21763,
		Ftype1: int8('F'),
	},
	3142: {
		Fword:  __ccgo_ts + 21770,
		Ftype1: int8('F'),
	},
	3143: {
		Fword:  __ccgo_ts + 21777,
		Ftype1: int8('F'),
	},
	3144: {
		Fword:  __ccgo_ts + 21784,
		Ftype1: int8('F'),
	},
	3145: {
		Fword:  __ccgo_ts + 21791,
		Ftype1: int8('F'),
	},
	3146: {
		Fword:  __ccgo_ts + 21798,
		Ftype1: int8('F'),
	},
	3147: {
		Fword:  __ccgo_ts + 21805,
		Ftype1: int8('F'),
	},
	3148: {
		Fword:  __ccgo_ts + 21812,
		Ftype1: int8('F'),
	},
	3149: {
		Fword:  __ccgo_ts + 21819,
		Ftype1: int8('F'),
	},
	3150: {
		Fword:  __ccgo_ts + 21826,
		Ftype1: int8('F'),
	},
	3151: {
		Fword:  __ccgo_ts + 21833,
		Ftype1: int8('F'),
	},
	3152: {
		Fword:  __ccgo_ts + 21840,
		Ftype1: int8('F'),
	},
	3153: {
		Fword:  __ccgo_ts + 21847,
		Ftype1: int8('F'),
	},
	3154: {
		Fword:  __ccgo_ts + 21854,
		Ftype1: int8('F'),
	},
	3155: {
		Fword:  __ccgo_ts + 21861,
		Ftype1: int8('F'),
	},
	3156: {
		Fword:  __ccgo_ts + 21868,
		Ftype1: int8('F'),
	},
	3157: {
		Fword:  __ccgo_ts + 21875,
		Ftype1: int8('F'),
	},
	3158: {
		Fword:  __ccgo_ts + 21882,
		Ftype1: int8('F'),
	},
	3159: {
		Fword:  __ccgo_ts + 21889,
		Ftype1: int8('F'),
	},
	3160: {
		Fword:  __ccgo_ts + 21896,
		Ftype1: int8('F'),
	},
	3161: {
		Fword:  __ccgo_ts + 21903,
		Ftype1: int8('F'),
	},
	3162: {
		Fword:  __ccgo_ts + 21910,
		Ftype1: int8('F'),
	},
	3163: {
		Fword:  __ccgo_ts + 21917,
		Ftype1: int8('F'),
	},
	3164: {
		Fword:  __ccgo_ts + 21924,
		Ftype1: int8('F'),
	},
	3165: {
		Fword:  __ccgo_ts + 21931,
		Ftype1: int8('F'),
	},
	3166: {
		Fword:  __ccgo_ts + 21938,
		Ftype1: int8('F'),
	},
	3167: {
		Fword:  __ccgo_ts + 21945,
		Ftype1: int8('F'),
	},
	3168: {
		Fword:  __ccgo_ts + 21952,
		Ftype1: int8('F'),
	},
	3169: {
		Fword:  __ccgo_ts + 21959,
		Ftype1: int8('F'),
	},
	3170: {
		Fword:  __ccgo_ts + 21966,
		Ftype1: int8('F'),
	},
	3171: {
		Fword:  __ccgo_ts + 21973,
		Ftype1: int8('F'),
	},
	3172: {
		Fword:  __ccgo_ts + 21980,
		Ftype1: int8('F'),
	},
	3173: {
		Fword:  __ccgo_ts + 21987,
		Ftype1: int8('F'),
	},
	3174: {
		Fword:  __ccgo_ts + 21994,
		Ftype1: int8('F'),
	},
	3175: {
		Fword:  __ccgo_ts + 22001,
		Ftype1: int8('F'),
	},
	3176: {
		Fword:  __ccgo_ts + 22008,
		Ftype1: int8('F'),
	},
	3177: {
		Fword:  __ccgo_ts + 22015,
		Ftype1: int8('F'),
	},
	3178: {
		Fword:  __ccgo_ts + 22022,
		Ftype1: int8('F'),
	},
	3179: {
		Fword:  __ccgo_ts + 22029,
		Ftype1: int8('F'),
	},
	3180: {
		Fword:  __ccgo_ts + 22036,
		Ftype1: int8('F'),
	},
	3181: {
		Fword:  __ccgo_ts + 22043,
		Ftype1: int8('F'),
	},
	3182: {
		Fword:  __ccgo_ts + 22050,
		Ftype1: int8('F'),
	},
	3183: {
		Fword:  __ccgo_ts + 22057,
		Ftype1: int8('F'),
	},
	3184: {
		Fword:  __ccgo_ts + 22064,
		Ftype1: int8('F'),
	},
	3185: {
		Fword:  __ccgo_ts + 22071,
		Ftype1: int8('F'),
	},
	3186: {
		Fword:  __ccgo_ts + 22078,
		Ftype1: int8('F'),
	},
	3187: {
		Fword:  __ccgo_ts + 22085,
		Ftype1: int8('F'),
	},
	3188: {
		Fword:  __ccgo_ts + 22092,
		Ftype1: int8('F'),
	},
	3189: {
		Fword:  __ccgo_ts + 22099,
		Ftype1: int8('F'),
	},
	3190: {
		Fword:  __ccgo_ts + 22106,
		Ftype1: int8('F'),
	},
	3191: {
		Fword:  __ccgo_ts + 22113,
		Ftype1: int8('F'),
	},
	3192: {
		Fword:  __ccgo_ts + 22120,
		Ftype1: int8('F'),
	},
	3193: {
		Fword:  __ccgo_ts + 22127,
		Ftype1: int8('F'),
	},
	3194: {
		Fword:  __ccgo_ts + 22134,
		Ftype1: int8('F'),
	},
	3195: {
		Fword:  __ccgo_ts + 22141,
		Ftype1: int8('F'),
	},
	3196: {
		Fword:  __ccgo_ts + 22148,
		Ftype1: int8('F'),
	},
	3197: {
		Fword:  __ccgo_ts + 22155,
		Ftype1: int8('F'),
	},
	3198: {
		Fword:  __ccgo_ts + 22162,
		Ftype1: int8('F'),
	},
	3199: {
		Fword:  __ccgo_ts + 22169,
		Ftype1: int8('F'),
	},
	3200: {
		Fword:  __ccgo_ts + 22176,
		Ftype1: int8('F'),
	},
	3201: {
		Fword:  __ccgo_ts + 22183,
		Ftype1: int8('F'),
	},
	3202: {
		Fword:  __ccgo_ts + 22190,
		Ftype1: int8('F'),
	},
	3203: {
		Fword:  __ccgo_ts + 22197,
		Ftype1: int8('F'),
	},
	3204: {
		Fword:  __ccgo_ts + 22204,
		Ftype1: int8('F'),
	},
	3205: {
		Fword:  __ccgo_ts + 22211,
		Ftype1: int8('F'),
	},
	3206: {
		Fword:  __ccgo_ts + 22218,
		Ftype1: int8('F'),
	},
	3207: {
		Fword:  __ccgo_ts + 22225,
		Ftype1: int8('F'),
	},
	3208: {
		Fword:  __ccgo_ts + 22232,
		Ftype1: int8('F'),
	},
	3209: {
		Fword:  __ccgo_ts + 22239,
		Ftype1: int8('F'),
	},
	3210: {
		Fword:  __ccgo_ts + 22246,
		Ftype1: int8('F'),
	},
	3211: {
		Fword:  __ccgo_ts + 22253,
		Ftype1: int8('F'),
	},
	3212: {
		Fword:  __ccgo_ts + 22260,
		Ftype1: int8('F'),
	},
	3213: {
		Fword:  __ccgo_ts + 22267,
		Ftype1: int8('F'),
	},
	3214: {
		Fword:  __ccgo_ts + 22274,
		Ftype1: int8('F'),
	},
	3215: {
		Fword:  __ccgo_ts + 22281,
		Ftype1: int8('F'),
	},
	3216: {
		Fword:  __ccgo_ts + 22288,
		Ftype1: int8('F'),
	},
	3217: {
		Fword:  __ccgo_ts + 22295,
		Ftype1: int8('F'),
	},
	3218: {
		Fword:  __ccgo_ts + 22302,
		Ftype1: int8('F'),
	},
	3219: {
		Fword:  __ccgo_ts + 22309,
		Ftype1: int8('F'),
	},
	3220: {
		Fword:  __ccgo_ts + 22316,
		Ftype1: int8('F'),
	},
	3221: {
		Fword:  __ccgo_ts + 22323,
		Ftype1: int8('F'),
	},
	3222: {
		Fword:  __ccgo_ts + 22330,
		Ftype1: int8('F'),
	},
	3223: {
		Fword:  __ccgo_ts + 22337,
		Ftype1: int8('F'),
	},
	3224: {
		Fword:  __ccgo_ts + 22344,
		Ftype1: int8('F'),
	},
	3225: {
		Fword:  __ccgo_ts + 22351,
		Ftype1: int8('F'),
	},
	3226: {
		Fword:  __ccgo_ts + 22358,
		Ftype1: int8('F'),
	},
	3227: {
		Fword:  __ccgo_ts + 22365,
		Ftype1: int8('F'),
	},
	3228: {
		Fword:  __ccgo_ts + 22372,
		Ftype1: int8('F'),
	},
	3229: {
		Fword:  __ccgo_ts + 22379,
		Ftype1: int8('F'),
	},
	3230: {
		Fword:  __ccgo_ts + 22386,
		Ftype1: int8('F'),
	},
	3231: {
		Fword:  __ccgo_ts + 22393,
		Ftype1: int8('F'),
	},
	3232: {
		Fword:  __ccgo_ts + 22400,
		Ftype1: int8('F'),
	},
	3233: {
		Fword:  __ccgo_ts + 22407,
		Ftype1: int8('F'),
	},
	3234: {
		Fword:  __ccgo_ts + 22414,
		Ftype1: int8('F'),
	},
	3235: {
		Fword:  __ccgo_ts + 22421,
		Ftype1: int8('F'),
	},
	3236: {
		Fword:  __ccgo_ts + 22428,
		Ftype1: int8('F'),
	},
	3237: {
		Fword:  __ccgo_ts + 22435,
		Ftype1: int8('F'),
	},
	3238: {
		Fword:  __ccgo_ts + 22442,
		Ftype1: int8('F'),
	},
	3239: {
		Fword:  __ccgo_ts + 22449,
		Ftype1: int8('F'),
	},
	3240: {
		Fword:  __ccgo_ts + 22456,
		Ftype1: int8('F'),
	},
	3241: {
		Fword:  __ccgo_ts + 22463,
		Ftype1: int8('F'),
	},
	3242: {
		Fword:  __ccgo_ts + 22470,
		Ftype1: int8('F'),
	},
	3243: {
		Fword:  __ccgo_ts + 22477,
		Ftype1: int8('F'),
	},
	3244: {
		Fword:  __ccgo_ts + 22484,
		Ftype1: int8('F'),
	},
	3245: {
		Fword:  __ccgo_ts + 22491,
		Ftype1: int8('F'),
	},
	3246: {
		Fword:  __ccgo_ts + 22498,
		Ftype1: int8('F'),
	},
	3247: {
		Fword:  __ccgo_ts + 22505,
		Ftype1: int8('F'),
	},
	3248: {
		Fword:  __ccgo_ts + 22512,
		Ftype1: int8('F'),
	},
	3249: {
		Fword:  __ccgo_ts + 22519,
		Ftype1: int8('F'),
	},
	3250: {
		Fword:  __ccgo_ts + 22526,
		Ftype1: int8('F'),
	},
	3251: {
		Fword:  __ccgo_ts + 22533,
		Ftype1: int8('F'),
	},
	3252: {
		Fword:  __ccgo_ts + 22540,
		Ftype1: int8('F'),
	},
	3253: {
		Fword:  __ccgo_ts + 22547,
		Ftype1: int8('F'),
	},
	3254: {
		Fword:  __ccgo_ts + 22554,
		Ftype1: int8('F'),
	},
	3255: {
		Fword:  __ccgo_ts + 22561,
		Ftype1: int8('F'),
	},
	3256: {
		Fword:  __ccgo_ts + 22568,
		Ftype1: int8('F'),
	},
	3257: {
		Fword:  __ccgo_ts + 22575,
		Ftype1: int8('F'),
	},
	3258: {
		Fword:  __ccgo_ts + 22582,
		Ftype1: int8('F'),
	},
	3259: {
		Fword:  __ccgo_ts + 22589,
		Ftype1: int8('F'),
	},
	3260: {
		Fword:  __ccgo_ts + 22596,
		Ftype1: int8('F'),
	},
	3261: {
		Fword:  __ccgo_ts + 22603,
		Ftype1: int8('F'),
	},
	3262: {
		Fword:  __ccgo_ts + 22610,
		Ftype1: int8('F'),
	},
	3263: {
		Fword:  __ccgo_ts + 22617,
		Ftype1: int8('F'),
	},
	3264: {
		Fword:  __ccgo_ts + 22624,
		Ftype1: int8('F'),
	},
	3265: {
		Fword:  __ccgo_ts + 22631,
		Ftype1: int8('F'),
	},
	3266: {
		Fword:  __ccgo_ts + 22638,
		Ftype1: int8('F'),
	},
	3267: {
		Fword:  __ccgo_ts + 22645,
		Ftype1: int8('F'),
	},
	3268: {
		Fword:  __ccgo_ts + 22652,
		Ftype1: int8('F'),
	},
	3269: {
		Fword:  __ccgo_ts + 22659,
		Ftype1: int8('F'),
	},
	3270: {
		Fword:  __ccgo_ts + 22666,
		Ftype1: int8('F'),
	},
	3271: {
		Fword:  __ccgo_ts + 22673,
		Ftype1: int8('F'),
	},
	3272: {
		Fword:  __ccgo_ts + 22680,
		Ftype1: int8('F'),
	},
	3273: {
		Fword:  __ccgo_ts + 22687,
		Ftype1: int8('F'),
	},
	3274: {
		Fword:  __ccgo_ts + 22694,
		Ftype1: int8('F'),
	},
	3275: {
		Fword:  __ccgo_ts + 22701,
		Ftype1: int8('F'),
	},
	3276: {
		Fword:  __ccgo_ts + 22708,
		Ftype1: int8('F'),
	},
	3277: {
		Fword:  __ccgo_ts + 22715,
		Ftype1: int8('F'),
	},
	3278: {
		Fword:  __ccgo_ts + 22722,
		Ftype1: int8('F'),
	},
	3279: {
		Fword:  __ccgo_ts + 22729,
		Ftype1: int8('F'),
	},
	3280: {
		Fword:  __ccgo_ts + 22736,
		Ftype1: int8('F'),
	},
	3281: {
		Fword:  __ccgo_ts + 22743,
		Ftype1: int8('F'),
	},
	3282: {
		Fword:  __ccgo_ts + 22750,
		Ftype1: int8('F'),
	},
	3283: {
		Fword:  __ccgo_ts + 22757,
		Ftype1: int8('F'),
	},
	3284: {
		Fword:  __ccgo_ts + 22764,
		Ftype1: int8('F'),
	},
	3285: {
		Fword:  __ccgo_ts + 22771,
		Ftype1: int8('F'),
	},
	3286: {
		Fword:  __ccgo_ts + 22778,
		Ftype1: int8('F'),
	},
	3287: {
		Fword:  __ccgo_ts + 22785,
		Ftype1: int8('F'),
	},
	3288: {
		Fword:  __ccgo_ts + 22792,
		Ftype1: int8('F'),
	},
	3289: {
		Fword:  __ccgo_ts + 22799,
		Ftype1: int8('F'),
	},
	3290: {
		Fword:  __ccgo_ts + 22806,
		Ftype1: int8('F'),
	},
	3291: {
		Fword:  __ccgo_ts + 22813,
		Ftype1: int8('F'),
	},
	3292: {
		Fword:  __ccgo_ts + 22820,
		Ftype1: int8('F'),
	},
	3293: {
		Fword:  __ccgo_ts + 22827,
		Ftype1: int8('F'),
	},
	3294: {
		Fword:  __ccgo_ts + 22834,
		Ftype1: int8('F'),
	},
	3295: {
		Fword:  __ccgo_ts + 22841,
		Ftype1: int8('F'),
	},
	3296: {
		Fword:  __ccgo_ts + 22848,
		Ftype1: int8('F'),
	},
	3297: {
		Fword:  __ccgo_ts + 22855,
		Ftype1: int8('F'),
	},
	3298: {
		Fword:  __ccgo_ts + 22862,
		Ftype1: int8('F'),
	},
	3299: {
		Fword:  __ccgo_ts + 22869,
		Ftype1: int8('F'),
	},
	3300: {
		Fword:  __ccgo_ts + 22876,
		Ftype1: int8('F'),
	},
	3301: {
		Fword:  __ccgo_ts + 22883,
		Ftype1: int8('F'),
	},
	3302: {
		Fword:  __ccgo_ts + 22890,
		Ftype1: int8('F'),
	},
	3303: {
		Fword:  __ccgo_ts + 22897,
		Ftype1: int8('F'),
	},
	3304: {
		Fword:  __ccgo_ts + 22904,
		Ftype1: int8('F'),
	},
	3305: {
		Fword:  __ccgo_ts + 22911,
		Ftype1: int8('F'),
	},
	3306: {
		Fword:  __ccgo_ts + 22918,
		Ftype1: int8('F'),
	},
	3307: {
		Fword:  __ccgo_ts + 22925,
		Ftype1: int8('F'),
	},
	3308: {
		Fword:  __ccgo_ts + 22932,
		Ftype1: int8('F'),
	},
	3309: {
		Fword:  __ccgo_ts + 22939,
		Ftype1: int8('F'),
	},
	3310: {
		Fword:  __ccgo_ts + 22946,
		Ftype1: int8('F'),
	},
	3311: {
		Fword:  __ccgo_ts + 22953,
		Ftype1: int8('F'),
	},
	3312: {
		Fword:  __ccgo_ts + 22958,
		Ftype1: int8('F'),
	},
	3313: {
		Fword:  __ccgo_ts + 22965,
		Ftype1: int8('F'),
	},
	3314: {
		Fword:  __ccgo_ts + 22972,
		Ftype1: int8('F'),
	},
	3315: {
		Fword:  __ccgo_ts + 22979,
		Ftype1: int8('F'),
	},
	3316: {
		Fword:  __ccgo_ts + 22986,
		Ftype1: int8('F'),
	},
	3317: {
		Fword:  __ccgo_ts + 22993,
		Ftype1: int8('F'),
	},
	3318: {
		Fword:  __ccgo_ts + 23000,
		Ftype1: int8('F'),
	},
	3319: {
		Fword:  __ccgo_ts + 23007,
		Ftype1: int8('F'),
	},
	3320: {
		Fword:  __ccgo_ts + 23014,
		Ftype1: int8('F'),
	},
	3321: {
		Fword:  __ccgo_ts + 23021,
		Ftype1: int8('F'),
	},
	3322: {
		Fword:  __ccgo_ts + 23028,
		Ftype1: int8('F'),
	},
	3323: {
		Fword:  __ccgo_ts + 23034,
		Ftype1: int8('F'),
	},
	3324: {
		Fword:  __ccgo_ts + 23041,
		Ftype1: int8('F'),
	},
	3325: {
		Fword:  __ccgo_ts + 23048,
		Ftype1: int8('F'),
	},
	3326: {
		Fword:  __ccgo_ts + 23055,
		Ftype1: int8('F'),
	},
	3327: {
		Fword:  __ccgo_ts + 23062,
		Ftype1: int8('F'),
	},
	3328: {
		Fword:  __ccgo_ts + 23069,
		Ftype1: int8('F'),
	},
	3329: {
		Fword:  __ccgo_ts + 23076,
		Ftype1: int8('F'),
	},
	3330: {
		Fword:  __ccgo_ts + 23083,
		Ftype1: int8('F'),
	},
	3331: {
		Fword:  __ccgo_ts + 23090,
		Ftype1: int8('F'),
	},
	3332: {
		Fword:  __ccgo_ts + 23097,
		Ftype1: int8('F'),
	},
	3333: {
		Fword:  __ccgo_ts + 23103,
		Ftype1: int8('F'),
	},
	3334: {
		Fword:  __ccgo_ts + 23110,
		Ftype1: int8('F'),
	},
	3335: {
		Fword:  __ccgo_ts + 23117,
		Ftype1: int8('F'),
	},
	3336: {
		Fword:  __ccgo_ts + 23124,
		Ftype1: int8('F'),
	},
	3337: {
		Fword:  __ccgo_ts + 23131,
		Ftype1: int8('F'),
	},
	3338: {
		Fword:  __ccgo_ts + 23138,
		Ftype1: int8('F'),
	},
	3339: {
		Fword:  __ccgo_ts + 23145,
		Ftype1: int8('F'),
	},
	3340: {
		Fword:  __ccgo_ts + 23152,
		Ftype1: int8('F'),
	},
	3341: {
		Fword:  __ccgo_ts + 23159,
		Ftype1: int8('F'),
	},
	3342: {
		Fword:  __ccgo_ts + 23166,
		Ftype1: int8('F'),
	},
	3343: {
		Fword:  __ccgo_ts + 23173,
		Ftype1: int8('F'),
	},
	3344: {
		Fword:  __ccgo_ts + 23180,
		Ftype1: int8('F'),
	},
	3345: {
		Fword:  __ccgo_ts + 23187,
		Ftype1: int8('F'),
	},
	3346: {
		Fword:  __ccgo_ts + 23194,
		Ftype1: int8('F'),
	},
	3347: {
		Fword:  __ccgo_ts + 23201,
		Ftype1: int8('F'),
	},
	3348: {
		Fword:  __ccgo_ts + 23207,
		Ftype1: int8('F'),
	},
	3349: {
		Fword:  __ccgo_ts + 23214,
		Ftype1: int8('F'),
	},
	3350: {
		Fword:  __ccgo_ts + 23221,
		Ftype1: int8('F'),
	},
	3351: {
		Fword:  __ccgo_ts + 23228,
		Ftype1: int8('F'),
	},
	3352: {
		Fword:  __ccgo_ts + 23235,
		Ftype1: int8('F'),
	},
	3353: {
		Fword:  __ccgo_ts + 23242,
		Ftype1: int8('F'),
	},
	3354: {
		Fword:  __ccgo_ts + 23249,
		Ftype1: int8('F'),
	},
	3355: {
		Fword:  __ccgo_ts + 23256,
		Ftype1: int8('F'),
	},
	3356: {
		Fword:  __ccgo_ts + 23263,
		Ftype1: int8('F'),
	},
	3357: {
		Fword:  __ccgo_ts + 23270,
		Ftype1: int8('F'),
	},
	3358: {
		Fword:  __ccgo_ts + 23277,
		Ftype1: int8('F'),
	},
	3359: {
		Fword:  __ccgo_ts + 23283,
		Ftype1: int8('F'),
	},
	3360: {
		Fword:  __ccgo_ts + 23290,
		Ftype1: int8('F'),
	},
	3361: {
		Fword:  __ccgo_ts + 23297,
		Ftype1: int8('F'),
	},
	3362: {
		Fword:  __ccgo_ts + 23304,
		Ftype1: int8('F'),
	},
	3363: {
		Fword:  __ccgo_ts + 23311,
		Ftype1: int8('F'),
	},
	3364: {
		Fword:  __ccgo_ts + 23318,
		Ftype1: int8('F'),
	},
	3365: {
		Fword:  __ccgo_ts + 23325,
		Ftype1: int8('F'),
	},
	3366: {
		Fword:  __ccgo_ts + 23332,
		Ftype1: int8('F'),
	},
	3367: {
		Fword:  __ccgo_ts + 23339,
		Ftype1: int8('F'),
	},
	3368: {
		Fword:  __ccgo_ts + 23346,
		Ftype1: int8('F'),
	},
	3369: {
		Fword:  __ccgo_ts + 23353,
		Ftype1: int8('F'),
	},
	3370: {
		Fword:  __ccgo_ts + 23360,
		Ftype1: int8('F'),
	},
	3371: {
		Fword:  __ccgo_ts + 23367,
		Ftype1: int8('F'),
	},
	3372: {
		Fword:  __ccgo_ts + 23374,
		Ftype1: int8('F'),
	},
	3373: {
		Fword:  __ccgo_ts + 23381,
		Ftype1: int8('F'),
	},
	3374: {
		Fword:  __ccgo_ts + 23388,
		Ftype1: int8('F'),
	},
	3375: {
		Fword:  __ccgo_ts + 23394,
		Ftype1: int8('F'),
	},
	3376: {
		Fword:  __ccgo_ts + 23401,
		Ftype1: int8('F'),
	},
	3377: {
		Fword:  __ccgo_ts + 23408,
		Ftype1: int8('F'),
	},
	3378: {
		Fword:  __ccgo_ts + 23415,
		Ftype1: int8('F'),
	},
	3379: {
		Fword:  __ccgo_ts + 23422,
		Ftype1: int8('F'),
	},
	3380: {
		Fword:  __ccgo_ts + 23428,
		Ftype1: int8('F'),
	},
	3381: {
		Fword:  __ccgo_ts + 23435,
		Ftype1: int8('F'),
	},
	3382: {
		Fword:  __ccgo_ts + 23442,
		Ftype1: int8('F'),
	},
	3383: {
		Fword:  __ccgo_ts + 23449,
		Ftype1: int8('F'),
	},
	3384: {
		Fword:  __ccgo_ts + 23456,
		Ftype1: int8('F'),
	},
	3385: {
		Fword:  __ccgo_ts + 23463,
		Ftype1: int8('F'),
	},
	3386: {
		Fword:  __ccgo_ts + 23470,
		Ftype1: int8('F'),
	},
	3387: {
		Fword:  __ccgo_ts + 23477,
		Ftype1: int8('F'),
	},
	3388: {
		Fword:  __ccgo_ts + 23484,
		Ftype1: int8('F'),
	},
	3389: {
		Fword:  __ccgo_ts + 23491,
		Ftype1: int8('F'),
	},
	3390: {
		Fword:  __ccgo_ts + 23498,
		Ftype1: int8('F'),
	},
	3391: {
		Fword:  __ccgo_ts + 23505,
		Ftype1: int8('F'),
	},
	3392: {
		Fword:  __ccgo_ts + 23512,
		Ftype1: int8('F'),
	},
	3393: {
		Fword:  __ccgo_ts + 23519,
		Ftype1: int8('F'),
	},
	3394: {
		Fword:  __ccgo_ts + 23526,
		Ftype1: int8('F'),
	},
	3395: {
		Fword:  __ccgo_ts + 23533,
		Ftype1: int8('F'),
	},
	3396: {
		Fword:  __ccgo_ts + 23540,
		Ftype1: int8('F'),
	},
	3397: {
		Fword:  __ccgo_ts + 23547,
		Ftype1: int8('F'),
	},
	3398: {
		Fword:  __ccgo_ts + 23554,
		Ftype1: int8('F'),
	},
	3399: {
		Fword:  __ccgo_ts + 23561,
		Ftype1: int8('F'),
	},
	3400: {
		Fword:  __ccgo_ts + 23568,
		Ftype1: int8('F'),
	},
	3401: {
		Fword:  __ccgo_ts + 23575,
		Ftype1: int8('F'),
	},
	3402: {
		Fword:  __ccgo_ts + 23582,
		Ftype1: int8('F'),
	},
	3403: {
		Fword:  __ccgo_ts + 23589,
		Ftype1: int8('F'),
	},
	3404: {
		Fword:  __ccgo_ts + 23595,
		Ftype1: int8('F'),
	},
	3405: {
		Fword:  __ccgo_ts + 23602,
		Ftype1: int8('F'),
	},
	3406: {
		Fword:  __ccgo_ts + 23609,
		Ftype1: int8('F'),
	},
	3407: {
		Fword:  __ccgo_ts + 23616,
		Ftype1: int8('F'),
	},
	3408: {
		Fword:  __ccgo_ts + 23623,
		Ftype1: int8('F'),
	},
	3409: {
		Fword:  __ccgo_ts + 23630,
		Ftype1: int8('F'),
	},
	3410: {
		Fword:  __ccgo_ts + 23637,
		Ftype1: int8('F'),
	},
	3411: {
		Fword:  __ccgo_ts + 23644,
		Ftype1: int8('F'),
	},
	3412: {
		Fword:  __ccgo_ts + 23651,
		Ftype1: int8('F'),
	},
	3413: {
		Fword:  __ccgo_ts + 23658,
		Ftype1: int8('F'),
	},
	3414: {
		Fword:  __ccgo_ts + 23665,
		Ftype1: int8('F'),
	},
	3415: {
		Fword:  __ccgo_ts + 23672,
		Ftype1: int8('F'),
	},
	3416: {
		Fword:  __ccgo_ts + 23679,
		Ftype1: int8('F'),
	},
	3417: {
		Fword:  __ccgo_ts + 23686,
		Ftype1: int8('F'),
	},
	3418: {
		Fword:  __ccgo_ts + 23693,
		Ftype1: int8('F'),
	},
	3419: {
		Fword:  __ccgo_ts + 23700,
		Ftype1: int8('F'),
	},
	3420: {
		Fword:  __ccgo_ts + 23707,
		Ftype1: int8('F'),
	},
	3421: {
		Fword:  __ccgo_ts + 23714,
		Ftype1: int8('F'),
	},
	3422: {
		Fword:  __ccgo_ts + 23721,
		Ftype1: int8('F'),
	},
	3423: {
		Fword:  __ccgo_ts + 23728,
		Ftype1: int8('F'),
	},
	3424: {
		Fword:  __ccgo_ts + 23735,
		Ftype1: int8('F'),
	},
	3425: {
		Fword:  __ccgo_ts + 23742,
		Ftype1: int8('F'),
	},
	3426: {
		Fword:  __ccgo_ts + 23749,
		Ftype1: int8('F'),
	},
	3427: {
		Fword:  __ccgo_ts + 23756,
		Ftype1: int8('F'),
	},
	3428: {
		Fword:  __ccgo_ts + 23763,
		Ftype1: int8('F'),
	},
	3429: {
		Fword:  __ccgo_ts + 23770,
		Ftype1: int8('F'),
	},
	3430: {
		Fword:  __ccgo_ts + 23777,
		Ftype1: int8('F'),
	},
	3431: {
		Fword:  __ccgo_ts + 23784,
		Ftype1: int8('F'),
	},
	3432: {
		Fword:  __ccgo_ts + 23791,
		Ftype1: int8('F'),
	},
	3433: {
		Fword:  __ccgo_ts + 23798,
		Ftype1: int8('F'),
	},
	3434: {
		Fword:  __ccgo_ts + 23805,
		Ftype1: int8('F'),
	},
	3435: {
		Fword:  __ccgo_ts + 23812,
		Ftype1: int8('F'),
	},
	3436: {
		Fword:  __ccgo_ts + 23819,
		Ftype1: int8('F'),
	},
	3437: {
		Fword:  __ccgo_ts + 23826,
		Ftype1: int8('F'),
	},
	3438: {
		Fword:  __ccgo_ts + 23833,
		Ftype1: int8('F'),
	},
	3439: {
		Fword:  __ccgo_ts + 23840,
		Ftype1: int8('F'),
	},
	3440: {
		Fword:  __ccgo_ts + 23847,
		Ftype1: int8('F'),
	},
	3441: {
		Fword:  __ccgo_ts + 23854,
		Ftype1: int8('F'),
	},
	3442: {
		Fword:  __ccgo_ts + 23861,
		Ftype1: int8('F'),
	},
	3443: {
		Fword:  __ccgo_ts + 23867,
		Ftype1: int8('F'),
	},
	3444: {
		Fword:  __ccgo_ts + 23874,
		Ftype1: int8('F'),
	},
	3445: {
		Fword:  __ccgo_ts + 23881,
		Ftype1: int8('F'),
	},
	3446: {
		Fword:  __ccgo_ts + 23888,
		Ftype1: int8('F'),
	},
	3447: {
		Fword:  __ccgo_ts + 23895,
		Ftype1: int8('F'),
	},
	3448: {
		Fword:  __ccgo_ts + 23902,
		Ftype1: int8('F'),
	},
	3449: {
		Fword:  __ccgo_ts + 23909,
		Ftype1: int8('F'),
	},
	3450: {
		Fword:  __ccgo_ts + 23916,
		Ftype1: int8('F'),
	},
	3451: {
		Fword:  __ccgo_ts + 23923,
		Ftype1: int8('F'),
	},
	3452: {
		Fword:  __ccgo_ts + 23930,
		Ftype1: int8('F'),
	},
	3453: {
		Fword:  __ccgo_ts + 23937,
		Ftype1: int8('F'),
	},
	3454: {
		Fword:  __ccgo_ts + 23944,
		Ftype1: int8('F'),
	},
	3455: {
		Fword:  __ccgo_ts + 23951,
		Ftype1: int8('F'),
	},
	3456: {
		Fword:  __ccgo_ts + 23958,
		Ftype1: int8('F'),
	},
	3457: {
		Fword:  __ccgo_ts + 23965,
		Ftype1: int8('F'),
	},
	3458: {
		Fword:  __ccgo_ts + 23971,
		Ftype1: int8('F'),
	},
	3459: {
		Fword:  __ccgo_ts + 23978,
		Ftype1: int8('F'),
	},
	3460: {
		Fword:  __ccgo_ts + 23985,
		Ftype1: int8('F'),
	},
	3461: {
		Fword:  __ccgo_ts + 23991,
		Ftype1: int8('F'),
	},
	3462: {
		Fword:  __ccgo_ts + 23998,
		Ftype1: int8('F'),
	},
	3463: {
		Fword:  __ccgo_ts + 24005,
		Ftype1: int8('F'),
	},
	3464: {
		Fword:  __ccgo_ts + 24012,
		Ftype1: int8('F'),
	},
	3465: {
		Fword:  __ccgo_ts + 24019,
		Ftype1: int8('F'),
	},
	3466: {
		Fword:  __ccgo_ts + 24026,
		Ftype1: int8('F'),
	},
	3467: {
		Fword:  __ccgo_ts + 24033,
		Ftype1: int8('F'),
	},
	3468: {
		Fword:  __ccgo_ts + 24040,
		Ftype1: int8('F'),
	},
	3469: {
		Fword:  __ccgo_ts + 24047,
		Ftype1: int8('F'),
	},
	3470: {
		Fword:  __ccgo_ts + 24054,
		Ftype1: int8('F'),
	},
	3471: {
		Fword:  __ccgo_ts + 24060,
		Ftype1: int8('F'),
	},
	3472: {
		Fword:  __ccgo_ts + 24067,
		Ftype1: int8('F'),
	},
	3473: {
		Fword:  __ccgo_ts + 24074,
		Ftype1: int8('F'),
	},
	3474: {
		Fword:  __ccgo_ts + 24081,
		Ftype1: int8('F'),
	},
	3475: {
		Fword:  __ccgo_ts + 24088,
		Ftype1: int8('F'),
	},
	3476: {
		Fword:  __ccgo_ts + 24095,
		Ftype1: int8('F'),
	},
	3477: {
		Fword:  __ccgo_ts + 24102,
		Ftype1: int8('F'),
	},
	3478: {
		Fword:  __ccgo_ts + 24109,
		Ftype1: int8('F'),
	},
	3479: {
		Fword:  __ccgo_ts + 24116,
		Ftype1: int8('F'),
	},
	3480: {
		Fword:  __ccgo_ts + 24123,
		Ftype1: int8('F'),
	},
	3481: {
		Fword:  __ccgo_ts + 24130,
		Ftype1: int8('F'),
	},
	3482: {
		Fword:  __ccgo_ts + 24137,
		Ftype1: int8('F'),
	},
	3483: {
		Fword:  __ccgo_ts + 24144,
		Ftype1: int8('F'),
	},
	3484: {
		Fword:  __ccgo_ts + 24151,
		Ftype1: int8('F'),
	},
	3485: {
		Fword:  __ccgo_ts + 24158,
		Ftype1: int8('F'),
	},
	3486: {
		Fword:  __ccgo_ts + 24165,
		Ftype1: int8('F'),
	},
	3487: {
		Fword:  __ccgo_ts + 24172,
		Ftype1: int8('F'),
	},
	3488: {
		Fword:  __ccgo_ts + 24178,
		Ftype1: int8('F'),
	},
	3489: {
		Fword:  __ccgo_ts + 24185,
		Ftype1: int8('F'),
	},
	3490: {
		Fword:  __ccgo_ts + 24192,
		Ftype1: int8('F'),
	},
	3491: {
		Fword:  __ccgo_ts + 24199,
		Ftype1: int8('F'),
	},
	3492: {
		Fword:  __ccgo_ts + 24206,
		Ftype1: int8('F'),
	},
	3493: {
		Fword:  __ccgo_ts + 24212,
		Ftype1: int8('F'),
	},
	3494: {
		Fword:  __ccgo_ts + 24219,
		Ftype1: int8('F'),
	},
	3495: {
		Fword:  __ccgo_ts + 24226,
		Ftype1: int8('F'),
	},
	3496: {
		Fword:  __ccgo_ts + 24233,
		Ftype1: int8('F'),
	},
	3497: {
		Fword:  __ccgo_ts + 24238,
		Ftype1: int8('F'),
	},
	3498: {
		Fword:  __ccgo_ts + 24245,
		Ftype1: int8('F'),
	},
	3499: {
		Fword:  __ccgo_ts + 24252,
		Ftype1: int8('F'),
	},
	3500: {
		Fword:  __ccgo_ts + 24259,
		Ftype1: int8('F'),
	},
	3501: {
		Fword:  __ccgo_ts + 24266,
		Ftype1: int8('F'),
	},
	3502: {
		Fword:  __ccgo_ts + 24273,
		Ftype1: int8('F'),
	},
	3503: {
		Fword:  __ccgo_ts + 24280,
		Ftype1: int8('F'),
	},
	3504: {
		Fword:  __ccgo_ts + 24287,
		Ftype1: int8('F'),
	},
	3505: {
		Fword:  __ccgo_ts + 24294,
		Ftype1: int8('F'),
	},
	3506: {
		Fword:  __ccgo_ts + 24301,
		Ftype1: int8('F'),
	},
	3507: {
		Fword:  __ccgo_ts + 24308,
		Ftype1: int8('F'),
	},
	3508: {
		Fword:  __ccgo_ts + 24314,
		Ftype1: int8('F'),
	},
	3509: {
		Fword:  __ccgo_ts + 24321,
		Ftype1: int8('F'),
	},
	3510: {
		Fword:  __ccgo_ts + 24328,
		Ftype1: int8('F'),
	},
	3511: {
		Fword:  __ccgo_ts + 24335,
		Ftype1: int8('F'),
	},
	3512: {
		Fword:  __ccgo_ts + 24342,
		Ftype1: int8('F'),
	},
	3513: {
		Fword:  __ccgo_ts + 24349,
		Ftype1: int8('F'),
	},
	3514: {
		Fword:  __ccgo_ts + 24356,
		Ftype1: int8('F'),
	},
	3515: {
		Fword:  __ccgo_ts + 24363,
		Ftype1: int8('F'),
	},
	3516: {
		Fword:  __ccgo_ts + 24370,
		Ftype1: int8('F'),
	},
	3517: {
		Fword:  __ccgo_ts + 24377,
		Ftype1: int8('F'),
	},
	3518: {
		Fword:  __ccgo_ts + 24383,
		Ftype1: int8('F'),
	},
	3519: {
		Fword:  __ccgo_ts + 24390,
		Ftype1: int8('F'),
	},
	3520: {
		Fword:  __ccgo_ts + 24397,
		Ftype1: int8('F'),
	},
	3521: {
		Fword:  __ccgo_ts + 24404,
		Ftype1: int8('F'),
	},
	3522: {
		Fword:  __ccgo_ts + 24411,
		Ftype1: int8('F'),
	},
	3523: {
		Fword:  __ccgo_ts + 24418,
		Ftype1: int8('F'),
	},
	3524: {
		Fword:  __ccgo_ts + 24425,
		Ftype1: int8('F'),
	},
	3525: {
		Fword:  __ccgo_ts + 24432,
		Ftype1: int8('F'),
	},
	3526: {
		Fword:  __ccgo_ts + 24439,
		Ftype1: int8('F'),
	},
	3527: {
		Fword:  __ccgo_ts + 24446,
		Ftype1: int8('F'),
	},
	3528: {
		Fword:  __ccgo_ts + 24453,
		Ftype1: int8('F'),
	},
	3529: {
		Fword:  __ccgo_ts + 24460,
		Ftype1: int8('F'),
	},
	3530: {
		Fword:  __ccgo_ts + 24467,
		Ftype1: int8('F'),
	},
	3531: {
		Fword:  __ccgo_ts + 24473,
		Ftype1: int8('F'),
	},
	3532: {
		Fword:  __ccgo_ts + 24480,
		Ftype1: int8('F'),
	},
	3533: {
		Fword:  __ccgo_ts + 24487,
		Ftype1: int8('F'),
	},
	3534: {
		Fword:  __ccgo_ts + 24494,
		Ftype1: int8('F'),
	},
	3535: {
		Fword:  __ccgo_ts + 24501,
		Ftype1: int8('F'),
	},
	3536: {
		Fword:  __ccgo_ts + 24507,
		Ftype1: int8('F'),
	},
	3537: {
		Fword:  __ccgo_ts + 24514,
		Ftype1: int8('F'),
	},
	3538: {
		Fword:  __ccgo_ts + 24521,
		Ftype1: int8('F'),
	},
	3539: {
		Fword:  __ccgo_ts + 24528,
		Ftype1: int8('F'),
	},
	3540: {
		Fword:  __ccgo_ts + 24535,
		Ftype1: int8('F'),
	},
	3541: {
		Fword:  __ccgo_ts + 24542,
		Ftype1: int8('F'),
	},
	3542: {
		Fword:  __ccgo_ts + 24549,
		Ftype1: int8('F'),
	},
	3543: {
		Fword:  __ccgo_ts + 24556,
		Ftype1: int8('F'),
	},
	3544: {
		Fword:  __ccgo_ts + 24563,
		Ftype1: int8('F'),
	},
	3545: {
		Fword:  __ccgo_ts + 24570,
		Ftype1: int8('F'),
	},
	3546: {
		Fword:  __ccgo_ts + 24577,
		Ftype1: int8('F'),
	},
	3547: {
		Fword:  __ccgo_ts + 24583,
		Ftype1: int8('F'),
	},
	3548: {
		Fword:  __ccgo_ts + 24590,
		Ftype1: int8('F'),
	},
	3549: {
		Fword:  __ccgo_ts + 24597,
		Ftype1: int8('F'),
	},
	3550: {
		Fword:  __ccgo_ts + 24604,
		Ftype1: int8('F'),
	},
	3551: {
		Fword:  __ccgo_ts + 24611,
		Ftype1: int8('F'),
	},
	3552: {
		Fword:  __ccgo_ts + 24618,
		Ftype1: int8('F'),
	},
	3553: {
		Fword:  __ccgo_ts + 24625,
		Ftype1: int8('F'),
	},
	3554: {
		Fword:  __ccgo_ts + 24632,
		Ftype1: int8('F'),
	},
	3555: {
		Fword:  __ccgo_ts + 24639,
		Ftype1: int8('F'),
	},
	3556: {
		Fword:  __ccgo_ts + 24646,
		Ftype1: int8('F'),
	},
	3557: {
		Fword:  __ccgo_ts + 24652,
		Ftype1: int8('F'),
	},
	3558: {
		Fword:  __ccgo_ts + 24659,
		Ftype1: int8('F'),
	},
	3559: {
		Fword:  __ccgo_ts + 24666,
		Ftype1: int8('F'),
	},
	3560: {
		Fword:  __ccgo_ts + 24673,
		Ftype1: int8('F'),
	},
	3561: {
		Fword:  __ccgo_ts + 24680,
		Ftype1: int8('F'),
	},
	3562: {
		Fword:  __ccgo_ts + 24687,
		Ftype1: int8('F'),
	},
	3563: {
		Fword:  __ccgo_ts + 24694,
		Ftype1: int8('F'),
	},
	3564: {
		Fword:  __ccgo_ts + 24701,
		Ftype1: int8('F'),
	},
	3565: {
		Fword:  __ccgo_ts + 24708,
		Ftype1: int8('F'),
	},
	3566: {
		Fword:  __ccgo_ts + 24714,
		Ftype1: int8('F'),
	},
	3567: {
		Fword:  __ccgo_ts + 24721,
		Ftype1: int8('F'),
	},
	3568: {
		Fword:  __ccgo_ts + 24728,
		Ftype1: int8('F'),
	},
	3569: {
		Fword:  __ccgo_ts + 24735,
		Ftype1: int8('F'),
	},
	3570: {
		Fword:  __ccgo_ts + 24742,
		Ftype1: int8('F'),
	},
	3571: {
		Fword:  __ccgo_ts + 24749,
		Ftype1: int8('F'),
	},
	3572: {
		Fword:  __ccgo_ts + 24756,
		Ftype1: int8('F'),
	},
	3573: {
		Fword:  __ccgo_ts + 24763,
		Ftype1: int8('F'),
	},
	3574: {
		Fword:  __ccgo_ts + 24770,
		Ftype1: int8('F'),
	},
	3575: {
		Fword:  __ccgo_ts + 24776,
		Ftype1: int8('F'),
	},
	3576: {
		Fword:  __ccgo_ts + 24783,
		Ftype1: int8('F'),
	},
	3577: {
		Fword:  __ccgo_ts + 24790,
		Ftype1: int8('F'),
	},
	3578: {
		Fword:  __ccgo_ts + 24797,
		Ftype1: int8('F'),
	},
	3579: {
		Fword:  __ccgo_ts + 24804,
		Ftype1: int8('F'),
	},
	3580: {
		Fword:  __ccgo_ts + 24811,
		Ftype1: int8('F'),
	},
	3581: {
		Fword:  __ccgo_ts + 24818,
		Ftype1: int8('F'),
	},
	3582: {
		Fword:  __ccgo_ts + 24825,
		Ftype1: int8('F'),
	},
	3583: {
		Fword:  __ccgo_ts + 24832,
		Ftype1: int8('F'),
	},
	3584: {
		Fword:  __ccgo_ts + 24839,
		Ftype1: int8('F'),
	},
	3585: {
		Fword:  __ccgo_ts + 24846,
		Ftype1: int8('F'),
	},
	3586: {
		Fword:  __ccgo_ts + 24853,
		Ftype1: int8('F'),
	},
	3587: {
		Fword:  __ccgo_ts + 24860,
		Ftype1: int8('F'),
	},
	3588: {
		Fword:  __ccgo_ts + 24867,
		Ftype1: int8('F'),
	},
	3589: {
		Fword:  __ccgo_ts + 24874,
		Ftype1: int8('F'),
	},
	3590: {
		Fword:  __ccgo_ts + 24881,
		Ftype1: int8('F'),
	},
	3591: {
		Fword:  __ccgo_ts + 24888,
		Ftype1: int8('F'),
	},
	3592: {
		Fword:  __ccgo_ts + 24895,
		Ftype1: int8('F'),
	},
	3593: {
		Fword:  __ccgo_ts + 24902,
		Ftype1: int8('F'),
	},
	3594: {
		Fword:  __ccgo_ts + 24909,
		Ftype1: int8('F'),
	},
	3595: {
		Fword:  __ccgo_ts + 24916,
		Ftype1: int8('F'),
	},
	3596: {
		Fword:  __ccgo_ts + 24923,
		Ftype1: int8('F'),
	},
	3597: {
		Fword:  __ccgo_ts + 24930,
		Ftype1: int8('F'),
	},
	3598: {
		Fword:  __ccgo_ts + 24937,
		Ftype1: int8('F'),
	},
	3599: {
		Fword:  __ccgo_ts + 24944,
		Ftype1: int8('F'),
	},
	3600: {
		Fword:  __ccgo_ts + 24951,
		Ftype1: int8('F'),
	},
	3601: {
		Fword:  __ccgo_ts + 24958,
		Ftype1: int8('F'),
	},
	3602: {
		Fword:  __ccgo_ts + 24965,
		Ftype1: int8('F'),
	},
	3603: {
		Fword:  __ccgo_ts + 24972,
		Ftype1: int8('F'),
	},
	3604: {
		Fword:  __ccgo_ts + 24979,
		Ftype1: int8('F'),
	},
	3605: {
		Fword:  __ccgo_ts + 24986,
		Ftype1: int8('F'),
	},
	3606: {
		Fword:  __ccgo_ts + 24993,
		Ftype1: int8('F'),
	},
	3607: {
		Fword:  __ccgo_ts + 25000,
		Ftype1: int8('F'),
	},
	3608: {
		Fword:  __ccgo_ts + 25007,
		Ftype1: int8('F'),
	},
	3609: {
		Fword:  __ccgo_ts + 25014,
		Ftype1: int8('F'),
	},
	3610: {
		Fword:  __ccgo_ts + 25020,
		Ftype1: int8('F'),
	},
	3611: {
		Fword:  __ccgo_ts + 25027,
		Ftype1: int8('F'),
	},
	3612: {
		Fword:  __ccgo_ts + 25034,
		Ftype1: int8('F'),
	},
	3613: {
		Fword:  __ccgo_ts + 25041,
		Ftype1: int8('F'),
	},
	3614: {
		Fword:  __ccgo_ts + 25048,
		Ftype1: int8('F'),
	},
	3615: {
		Fword:  __ccgo_ts + 25055,
		Ftype1: int8('F'),
	},
	3616: {
		Fword:  __ccgo_ts + 25062,
		Ftype1: int8('F'),
	},
	3617: {
		Fword:  __ccgo_ts + 25069,
		Ftype1: int8('F'),
	},
	3618: {
		Fword:  __ccgo_ts + 25076,
		Ftype1: int8('F'),
	},
	3619: {
		Fword:  __ccgo_ts + 25082,
		Ftype1: int8('F'),
	},
	3620: {
		Fword:  __ccgo_ts + 25089,
		Ftype1: int8('F'),
	},
	3621: {
		Fword:  __ccgo_ts + 25096,
		Ftype1: int8('F'),
	},
	3622: {
		Fword:  __ccgo_ts + 25103,
		Ftype1: int8('F'),
	},
	3623: {
		Fword:  __ccgo_ts + 25110,
		Ftype1: int8('F'),
	},
	3624: {
		Fword:  __ccgo_ts + 25117,
		Ftype1: int8('F'),
	},
	3625: {
		Fword:  __ccgo_ts + 25124,
		Ftype1: int8('F'),
	},
	3626: {
		Fword:  __ccgo_ts + 25130,
		Ftype1: int8('F'),
	},
	3627: {
		Fword:  __ccgo_ts + 25137,
		Ftype1: int8('F'),
	},
	3628: {
		Fword:  __ccgo_ts + 25144,
		Ftype1: int8('F'),
	},
	3629: {
		Fword:  __ccgo_ts + 25151,
		Ftype1: int8('F'),
	},
	3630: {
		Fword:  __ccgo_ts + 25158,
		Ftype1: int8('F'),
	},
	3631: {
		Fword:  __ccgo_ts + 25165,
		Ftype1: int8('F'),
	},
	3632: {
		Fword:  __ccgo_ts + 25172,
		Ftype1: int8('F'),
	},
	3633: {
		Fword:  __ccgo_ts + 25178,
		Ftype1: int8('F'),
	},
	3634: {
		Fword:  __ccgo_ts + 25185,
		Ftype1: int8('F'),
	},
	3635: {
		Fword:  __ccgo_ts + 25192,
		Ftype1: int8('F'),
	},
	3636: {
		Fword:  __ccgo_ts + 25199,
		Ftype1: int8('F'),
	},
	3637: {
		Fword:  __ccgo_ts + 25206,
		Ftype1: int8('F'),
	},
	3638: {
		Fword:  __ccgo_ts + 25213,
		Ftype1: int8('F'),
	},
	3639: {
		Fword:  __ccgo_ts + 25220,
		Ftype1: int8('F'),
	},
	3640: {
		Fword:  __ccgo_ts + 25227,
		Ftype1: int8('F'),
	},
	3641: {
		Fword:  __ccgo_ts + 25234,
		Ftype1: int8('F'),
	},
	3642: {
		Fword:  __ccgo_ts + 25241,
		Ftype1: int8('F'),
	},
	3643: {
		Fword:  __ccgo_ts + 25248,
		Ftype1: int8('F'),
	},
	3644: {
		Fword:  __ccgo_ts + 25255,
		Ftype1: int8('F'),
	},
	3645: {
		Fword:  __ccgo_ts + 25262,
		Ftype1: int8('F'),
	},
	3646: {
		Fword:  __ccgo_ts + 25269,
		Ftype1: int8('F'),
	},
	3647: {
		Fword:  __ccgo_ts + 25276,
		Ftype1: int8('F'),
	},
	3648: {
		Fword:  __ccgo_ts + 25283,
		Ftype1: int8('F'),
	},
	3649: {
		Fword:  __ccgo_ts + 25290,
		Ftype1: int8('F'),
	},
	3650: {
		Fword:  __ccgo_ts + 25297,
		Ftype1: int8('F'),
	},
	3651: {
		Fword:  __ccgo_ts + 25304,
		Ftype1: int8('F'),
	},
	3652: {
		Fword:  __ccgo_ts + 25311,
		Ftype1: int8('F'),
	},
	3653: {
		Fword:  __ccgo_ts + 25318,
		Ftype1: int8('F'),
	},
	3654: {
		Fword:  __ccgo_ts + 25325,
		Ftype1: int8('F'),
	},
	3655: {
		Fword:  __ccgo_ts + 25332,
		Ftype1: int8('F'),
	},
	3656: {
		Fword:  __ccgo_ts + 25339,
		Ftype1: int8('F'),
	},
	3657: {
		Fword:  __ccgo_ts + 25346,
		Ftype1: int8('F'),
	},
	3658: {
		Fword:  __ccgo_ts + 25353,
		Ftype1: int8('F'),
	},
	3659: {
		Fword:  __ccgo_ts + 25360,
		Ftype1: int8('F'),
	},
	3660: {
		Fword:  __ccgo_ts + 25367,
		Ftype1: int8('F'),
	},
	3661: {
		Fword:  __ccgo_ts + 25374,
		Ftype1: int8('F'),
	},
	3662: {
		Fword:  __ccgo_ts + 25381,
		Ftype1: int8('F'),
	},
	3663: {
		Fword:  __ccgo_ts + 25388,
		Ftype1: int8('F'),
	},
	3664: {
		Fword:  __ccgo_ts + 25395,
		Ftype1: int8('F'),
	},
	3665: {
		Fword:  __ccgo_ts + 25402,
		Ftype1: int8('F'),
	},
	3666: {
		Fword:  __ccgo_ts + 25409,
		Ftype1: int8('F'),
	},
	3667: {
		Fword:  __ccgo_ts + 25416,
		Ftype1: int8('F'),
	},
	3668: {
		Fword:  __ccgo_ts + 25423,
		Ftype1: int8('F'),
	},
	3669: {
		Fword:  __ccgo_ts + 25430,
		Ftype1: int8('F'),
	},
	3670: {
		Fword:  __ccgo_ts + 25437,
		Ftype1: int8('F'),
	},
	3671: {
		Fword:  __ccgo_ts + 25444,
		Ftype1: int8('F'),
	},
	3672: {
		Fword:  __ccgo_ts + 25451,
		Ftype1: int8('F'),
	},
	3673: {
		Fword:  __ccgo_ts + 25458,
		Ftype1: int8('F'),
	},
	3674: {
		Fword:  __ccgo_ts + 25465,
		Ftype1: int8('F'),
	},
	3675: {
		Fword:  __ccgo_ts + 25472,
		Ftype1: int8('F'),
	},
	3676: {
		Fword:  __ccgo_ts + 25479,
		Ftype1: int8('F'),
	},
	3677: {
		Fword:  __ccgo_ts + 25486,
		Ftype1: int8('F'),
	},
	3678: {
		Fword:  __ccgo_ts + 25493,
		Ftype1: int8('F'),
	},
	3679: {
		Fword:  __ccgo_ts + 25500,
		Ftype1: int8('F'),
	},
	3680: {
		Fword:  __ccgo_ts + 25507,
		Ftype1: int8('F'),
	},
	3681: {
		Fword:  __ccgo_ts + 25514,
		Ftype1: int8('F'),
	},
	3682: {
		Fword:  __ccgo_ts + 25521,
		Ftype1: int8('F'),
	},
	3683: {
		Fword:  __ccgo_ts + 25528,
		Ftype1: int8('F'),
	},
	3684: {
		Fword:  __ccgo_ts + 25535,
		Ftype1: int8('F'),
	},
	3685: {
		Fword:  __ccgo_ts + 25542,
		Ftype1: int8('F'),
	},
	3686: {
		Fword:  __ccgo_ts + 25549,
		Ftype1: int8('F'),
	},
	3687: {
		Fword:  __ccgo_ts + 25556,
		Ftype1: int8('F'),
	},
	3688: {
		Fword:  __ccgo_ts + 25563,
		Ftype1: int8('F'),
	},
	3689: {
		Fword:  __ccgo_ts + 25570,
		Ftype1: int8('F'),
	},
	3690: {
		Fword:  __ccgo_ts + 25577,
		Ftype1: int8('F'),
	},
	3691: {
		Fword:  __ccgo_ts + 25584,
		Ftype1: int8('F'),
	},
	3692: {
		Fword:  __ccgo_ts + 25591,
		Ftype1: int8('F'),
	},
	3693: {
		Fword:  __ccgo_ts + 25598,
		Ftype1: int8('F'),
	},
	3694: {
		Fword:  __ccgo_ts + 25605,
		Ftype1: int8('F'),
	},
	3695: {
		Fword:  __ccgo_ts + 25612,
		Ftype1: int8('F'),
	},
	3696: {
		Fword:  __ccgo_ts + 25619,
		Ftype1: int8('F'),
	},
	3697: {
		Fword:  __ccgo_ts + 25626,
		Ftype1: int8('F'),
	},
	3698: {
		Fword:  __ccgo_ts + 25633,
		Ftype1: int8('F'),
	},
	3699: {
		Fword:  __ccgo_ts + 25640,
		Ftype1: int8('F'),
	},
	3700: {
		Fword:  __ccgo_ts + 25647,
		Ftype1: int8('F'),
	},
	3701: {
		Fword:  __ccgo_ts + 25654,
		Ftype1: int8('F'),
	},
	3702: {
		Fword:  __ccgo_ts + 25661,
		Ftype1: int8('F'),
	},
	3703: {
		Fword:  __ccgo_ts + 25668,
		Ftype1: int8('F'),
	},
	3704: {
		Fword:  __ccgo_ts + 25675,
		Ftype1: int8('F'),
	},
	3705: {
		Fword:  __ccgo_ts + 25682,
		Ftype1: int8('F'),
	},
	3706: {
		Fword:  __ccgo_ts + 25689,
		Ftype1: int8('F'),
	},
	3707: {
		Fword:  __ccgo_ts + 25696,
		Ftype1: int8('F'),
	},
	3708: {
		Fword:  __ccgo_ts + 25703,
		Ftype1: int8('F'),
	},
	3709: {
		Fword:  __ccgo_ts + 25710,
		Ftype1: int8('F'),
	},
	3710: {
		Fword:  __ccgo_ts + 25717,
		Ftype1: int8('F'),
	},
	3711: {
		Fword:  __ccgo_ts + 25724,
		Ftype1: int8('F'),
	},
	3712: {
		Fword:  __ccgo_ts + 25731,
		Ftype1: int8('F'),
	},
	3713: {
		Fword:  __ccgo_ts + 25738,
		Ftype1: int8('F'),
	},
	3714: {
		Fword:  __ccgo_ts + 25745,
		Ftype1: int8('F'),
	},
	3715: {
		Fword:  __ccgo_ts + 25752,
		Ftype1: int8('F'),
	},
	3716: {
		Fword:  __ccgo_ts + 25759,
		Ftype1: int8('F'),
	},
	3717: {
		Fword:  __ccgo_ts + 25766,
		Ftype1: int8('F'),
	},
	3718: {
		Fword:  __ccgo_ts + 25773,
		Ftype1: int8('F'),
	},
	3719: {
		Fword:  __ccgo_ts + 25780,
		Ftype1: int8('F'),
	},
	3720: {
		Fword:  __ccgo_ts + 25786,
		Ftype1: int8('F'),
	},
	3721: {
		Fword:  __ccgo_ts + 25793,
		Ftype1: int8('F'),
	},
	3722: {
		Fword:  __ccgo_ts + 25800,
		Ftype1: int8('F'),
	},
	3723: {
		Fword:  __ccgo_ts + 25807,
		Ftype1: int8('F'),
	},
	3724: {
		Fword:  __ccgo_ts + 25814,
		Ftype1: int8('F'),
	},
	3725: {
		Fword:  __ccgo_ts + 25821,
		Ftype1: int8('F'),
	},
	3726: {
		Fword:  __ccgo_ts + 25828,
		Ftype1: int8('F'),
	},
	3727: {
		Fword:  __ccgo_ts + 25835,
		Ftype1: int8('F'),
	},
	3728: {
		Fword:  __ccgo_ts + 25841,
		Ftype1: int8('F'),
	},
	3729: {
		Fword:  __ccgo_ts + 25848,
		Ftype1: int8('F'),
	},
	3730: {
		Fword:  __ccgo_ts + 25855,
		Ftype1: int8('F'),
	},
	3731: {
		Fword:  __ccgo_ts + 25862,
		Ftype1: int8('F'),
	},
	3732: {
		Fword:  __ccgo_ts + 25869,
		Ftype1: int8('F'),
	},
	3733: {
		Fword:  __ccgo_ts + 25876,
		Ftype1: int8('F'),
	},
	3734: {
		Fword:  __ccgo_ts + 25883,
		Ftype1: int8('F'),
	},
	3735: {
		Fword:  __ccgo_ts + 25890,
		Ftype1: int8('F'),
	},
	3736: {
		Fword:  __ccgo_ts + 25897,
		Ftype1: int8('F'),
	},
	3737: {
		Fword:  __ccgo_ts + 25904,
		Ftype1: int8('F'),
	},
	3738: {
		Fword:  __ccgo_ts + 25911,
		Ftype1: int8('F'),
	},
	3739: {
		Fword:  __ccgo_ts + 25918,
		Ftype1: int8('F'),
	},
	3740: {
		Fword:  __ccgo_ts + 25925,
		Ftype1: int8('F'),
	},
	3741: {
		Fword:  __ccgo_ts + 25932,
		Ftype1: int8('F'),
	},
	3742: {
		Fword:  __ccgo_ts + 25939,
		Ftype1: int8('F'),
	},
	3743: {
		Fword:  __ccgo_ts + 25946,
		Ftype1: int8('F'),
	},
	3744: {
		Fword:  __ccgo_ts + 25953,
		Ftype1: int8('F'),
	},
	3745: {
		Fword:  __ccgo_ts + 25960,
		Ftype1: int8('F'),
	},
	3746: {
		Fword:  __ccgo_ts + 25967,
		Ftype1: int8('F'),
	},
	3747: {
		Fword:  __ccgo_ts + 25974,
		Ftype1: int8('F'),
	},
	3748: {
		Fword:  __ccgo_ts + 25981,
		Ftype1: int8('F'),
	},
	3749: {
		Fword:  __ccgo_ts + 25988,
		Ftype1: int8('F'),
	},
	3750: {
		Fword:  __ccgo_ts + 25995,
		Ftype1: int8('F'),
	},
	3751: {
		Fword:  __ccgo_ts + 26002,
		Ftype1: int8('F'),
	},
	3752: {
		Fword:  __ccgo_ts + 26009,
		Ftype1: int8('F'),
	},
	3753: {
		Fword:  __ccgo_ts + 26016,
		Ftype1: int8('F'),
	},
	3754: {
		Fword:  __ccgo_ts + 26023,
		Ftype1: int8('F'),
	},
	3755: {
		Fword:  __ccgo_ts + 26030,
		Ftype1: int8('F'),
	},
	3756: {
		Fword:  __ccgo_ts + 26037,
		Ftype1: int8('F'),
	},
	3757: {
		Fword:  __ccgo_ts + 26044,
		Ftype1: int8('F'),
	},
	3758: {
		Fword:  __ccgo_ts + 26051,
		Ftype1: int8('F'),
	},
	3759: {
		Fword:  __ccgo_ts + 26058,
		Ftype1: int8('F'),
	},
	3760: {
		Fword:  __ccgo_ts + 26065,
		Ftype1: int8('F'),
	},
	3761: {
		Fword:  __ccgo_ts + 26072,
		Ftype1: int8('F'),
	},
	3762: {
		Fword:  __ccgo_ts + 26079,
		Ftype1: int8('F'),
	},
	3763: {
		Fword:  __ccgo_ts + 26086,
		Ftype1: int8('F'),
	},
	3764: {
		Fword:  __ccgo_ts + 26093,
		Ftype1: int8('F'),
	},
	3765: {
		Fword:  __ccgo_ts + 26100,
		Ftype1: int8('F'),
	},
	3766: {
		Fword:  __ccgo_ts + 26107,
		Ftype1: int8('F'),
	},
	3767: {
		Fword:  __ccgo_ts + 26114,
		Ftype1: int8('F'),
	},
	3768: {
		Fword:  __ccgo_ts + 26121,
		Ftype1: int8('F'),
	},
	3769: {
		Fword:  __ccgo_ts + 26128,
		Ftype1: int8('F'),
	},
	3770: {
		Fword:  __ccgo_ts + 26135,
		Ftype1: int8('F'),
	},
	3771: {
		Fword:  __ccgo_ts + 26142,
		Ftype1: int8('F'),
	},
	3772: {
		Fword:  __ccgo_ts + 26149,
		Ftype1: int8('F'),
	},
	3773: {
		Fword:  __ccgo_ts + 26156,
		Ftype1: int8('F'),
	},
	3774: {
		Fword:  __ccgo_ts + 26163,
		Ftype1: int8('F'),
	},
	3775: {
		Fword:  __ccgo_ts + 26170,
		Ftype1: int8('F'),
	},
	3776: {
		Fword:  __ccgo_ts + 26177,
		Ftype1: int8('F'),
	},
	3777: {
		Fword:  __ccgo_ts + 26184,
		Ftype1: int8('F'),
	},
	3778: {
		Fword:  __ccgo_ts + 26191,
		Ftype1: int8('F'),
	},
	3779: {
		Fword:  __ccgo_ts + 26198,
		Ftype1: int8('F'),
	},
	3780: {
		Fword:  __ccgo_ts + 26205,
		Ftype1: int8('F'),
	},
	3781: {
		Fword:  __ccgo_ts + 26212,
		Ftype1: int8('F'),
	},
	3782: {
		Fword:  __ccgo_ts + 26218,
		Ftype1: int8('F'),
	},
	3783: {
		Fword:  __ccgo_ts + 26225,
		Ftype1: int8('F'),
	},
	3784: {
		Fword:  __ccgo_ts + 26231,
		Ftype1: int8('F'),
	},
	3785: {
		Fword:  __ccgo_ts + 26237,
		Ftype1: int8('F'),
	},
	3786: {
		Fword:  __ccgo_ts + 26244,
		Ftype1: int8('F'),
	},
	3787: {
		Fword:  __ccgo_ts + 26251,
		Ftype1: int8('F'),
	},
	3788: {
		Fword:  __ccgo_ts + 26258,
		Ftype1: int8('F'),
	},
	3789: {
		Fword:  __ccgo_ts + 26264,
		Ftype1: int8('F'),
	},
	3790: {
		Fword:  __ccgo_ts + 26271,
		Ftype1: int8('F'),
	},
	3791: {
		Fword:  __ccgo_ts + 26277,
		Ftype1: int8('F'),
	},
	3792: {
		Fword:  __ccgo_ts + 26284,
		Ftype1: int8('F'),
	},
	3793: {
		Fword:  __ccgo_ts + 26291,
		Ftype1: int8('F'),
	},
	3794: {
		Fword:  __ccgo_ts + 26298,
		Ftype1: int8('F'),
	},
	3795: {
		Fword:  __ccgo_ts + 26305,
		Ftype1: int8('F'),
	},
	3796: {
		Fword:  __ccgo_ts + 26312,
		Ftype1: int8('F'),
	},
	3797: {
		Fword:  __ccgo_ts + 26319,
		Ftype1: int8('F'),
	},
	3798: {
		Fword:  __ccgo_ts + 26326,
		Ftype1: int8('F'),
	},
	3799: {
		Fword:  __ccgo_ts + 26333,
		Ftype1: int8('F'),
	},
	3800: {
		Fword:  __ccgo_ts + 26340,
		Ftype1: int8('F'),
	},
	3801: {
		Fword:  __ccgo_ts + 26347,
		Ftype1: int8('F'),
	},
	3802: {
		Fword:  __ccgo_ts + 26354,
		Ftype1: int8('F'),
	},
	3803: {
		Fword:  __ccgo_ts + 26361,
		Ftype1: int8('F'),
	},
	3804: {
		Fword:  __ccgo_ts + 26368,
		Ftype1: int8('F'),
	},
	3805: {
		Fword:  __ccgo_ts + 26375,
		Ftype1: int8('F'),
	},
	3806: {
		Fword:  __ccgo_ts + 26382,
		Ftype1: int8('F'),
	},
	3807: {
		Fword:  __ccgo_ts + 26389,
		Ftype1: int8('F'),
	},
	3808: {
		Fword:  __ccgo_ts + 26396,
		Ftype1: int8('F'),
	},
	3809: {
		Fword:  __ccgo_ts + 26403,
		Ftype1: int8('F'),
	},
	3810: {
		Fword:  __ccgo_ts + 26410,
		Ftype1: int8('F'),
	},
	3811: {
		Fword:  __ccgo_ts + 26417,
		Ftype1: int8('F'),
	},
	3812: {
		Fword:  __ccgo_ts + 26424,
		Ftype1: int8('F'),
	},
	3813: {
		Fword:  __ccgo_ts + 26431,
		Ftype1: int8('F'),
	},
	3814: {
		Fword:  __ccgo_ts + 26438,
		Ftype1: int8('F'),
	},
	3815: {
		Fword:  __ccgo_ts + 26445,
		Ftype1: int8('F'),
	},
	3816: {
		Fword:  __ccgo_ts + 26452,
		Ftype1: int8('F'),
	},
	3817: {
		Fword:  __ccgo_ts + 26459,
		Ftype1: int8('F'),
	},
	3818: {
		Fword:  __ccgo_ts + 26466,
		Ftype1: int8('F'),
	},
	3819: {
		Fword:  __ccgo_ts + 26473,
		Ftype1: int8('F'),
	},
	3820: {
		Fword:  __ccgo_ts + 26480,
		Ftype1: int8('F'),
	},
	3821: {
		Fword:  __ccgo_ts + 26487,
		Ftype1: int8('F'),
	},
	3822: {
		Fword:  __ccgo_ts + 26494,
		Ftype1: int8('F'),
	},
	3823: {
		Fword:  __ccgo_ts + 26501,
		Ftype1: int8('F'),
	},
	3824: {
		Fword:  __ccgo_ts + 26508,
		Ftype1: int8('F'),
	},
	3825: {
		Fword:  __ccgo_ts + 26515,
		Ftype1: int8('F'),
	},
	3826: {
		Fword:  __ccgo_ts + 26522,
		Ftype1: int8('F'),
	},
	3827: {
		Fword:  __ccgo_ts + 26529,
		Ftype1: int8('F'),
	},
	3828: {
		Fword:  __ccgo_ts + 26536,
		Ftype1: int8('F'),
	},
	3829: {
		Fword:  __ccgo_ts + 26543,
		Ftype1: int8('F'),
	},
	3830: {
		Fword:  __ccgo_ts + 26550,
		Ftype1: int8('F'),
	},
	3831: {
		Fword:  __ccgo_ts + 26557,
		Ftype1: int8('F'),
	},
	3832: {
		Fword:  __ccgo_ts + 26564,
		Ftype1: int8('F'),
	},
	3833: {
		Fword:  __ccgo_ts + 26571,
		Ftype1: int8('F'),
	},
	3834: {
		Fword:  __ccgo_ts + 26578,
		Ftype1: int8('F'),
	},
	3835: {
		Fword:  __ccgo_ts + 26585,
		Ftype1: int8('F'),
	},
	3836: {
		Fword:  __ccgo_ts + 26592,
		Ftype1: int8('F'),
	},
	3837: {
		Fword:  __ccgo_ts + 26599,
		Ftype1: int8('F'),
	},
	3838: {
		Fword:  __ccgo_ts + 26606,
		Ftype1: int8('F'),
	},
	3839: {
		Fword:  __ccgo_ts + 26613,
		Ftype1: int8('F'),
	},
	3840: {
		Fword:  __ccgo_ts + 26620,
		Ftype1: int8('F'),
	},
	3841: {
		Fword:  __ccgo_ts + 26627,
		Ftype1: int8('F'),
	},
	3842: {
		Fword:  __ccgo_ts + 26634,
		Ftype1: int8('F'),
	},
	3843: {
		Fword:  __ccgo_ts + 26641,
		Ftype1: int8('F'),
	},
	3844: {
		Fword:  __ccgo_ts + 26648,
		Ftype1: int8('F'),
	},
	3845: {
		Fword:  __ccgo_ts + 26655,
		Ftype1: int8('F'),
	},
	3846: {
		Fword:  __ccgo_ts + 26662,
		Ftype1: int8('F'),
	},
	3847: {
		Fword:  __ccgo_ts + 26669,
		Ftype1: int8('F'),
	},
	3848: {
		Fword:  __ccgo_ts + 26676,
		Ftype1: int8('F'),
	},
	3849: {
		Fword:  __ccgo_ts + 26683,
		Ftype1: int8('F'),
	},
	3850: {
		Fword:  __ccgo_ts + 26689,
		Ftype1: int8('F'),
	},
	3851: {
		Fword:  __ccgo_ts + 26696,
		Ftype1: int8('F'),
	},
	3852: {
		Fword:  __ccgo_ts + 26703,
		Ftype1: int8('F'),
	},
	3853: {
		Fword:  __ccgo_ts + 26710,
		Ftype1: int8('F'),
	},
	3854: {
		Fword:  __ccgo_ts + 26717,
		Ftype1: int8('F'),
	},
	3855: {
		Fword:  __ccgo_ts + 26724,
		Ftype1: int8('F'),
	},
	3856: {
		Fword:  __ccgo_ts + 26731,
		Ftype1: int8('F'),
	},
	3857: {
		Fword:  __ccgo_ts + 26738,
		Ftype1: int8('F'),
	},
	3858: {
		Fword:  __ccgo_ts + 26745,
		Ftype1: int8('F'),
	},
	3859: {
		Fword:  __ccgo_ts + 26752,
		Ftype1: int8('F'),
	},
	3860: {
		Fword:  __ccgo_ts + 26759,
		Ftype1: int8('F'),
	},
	3861: {
		Fword:  __ccgo_ts + 26766,
		Ftype1: int8('F'),
	},
	3862: {
		Fword:  __ccgo_ts + 26773,
		Ftype1: int8('F'),
	},
	3863: {
		Fword:  __ccgo_ts + 26780,
		Ftype1: int8('F'),
	},
	3864: {
		Fword:  __ccgo_ts + 26787,
		Ftype1: int8('F'),
	},
	3865: {
		Fword:  __ccgo_ts + 26794,
		Ftype1: int8('F'),
	},
	3866: {
		Fword:  __ccgo_ts + 26801,
		Ftype1: int8('F'),
	},
	3867: {
		Fword:  __ccgo_ts + 26808,
		Ftype1: int8('F'),
	},
	3868: {
		Fword:  __ccgo_ts + 26815,
		Ftype1: int8('F'),
	},
	3869: {
		Fword:  __ccgo_ts + 26822,
		Ftype1: int8('F'),
	},
	3870: {
		Fword:  __ccgo_ts + 26829,
		Ftype1: int8('F'),
	},
	3871: {
		Fword:  __ccgo_ts + 26836,
		Ftype1: int8('F'),
	},
	3872: {
		Fword:  __ccgo_ts + 26843,
		Ftype1: int8('F'),
	},
	3873: {
		Fword:  __ccgo_ts + 26850,
		Ftype1: int8('F'),
	},
	3874: {
		Fword:  __ccgo_ts + 26857,
		Ftype1: int8('F'),
	},
	3875: {
		Fword:  __ccgo_ts + 26864,
		Ftype1: int8('F'),
	},
	3876: {
		Fword:  __ccgo_ts + 26871,
		Ftype1: int8('F'),
	},
	3877: {
		Fword:  __ccgo_ts + 26878,
		Ftype1: int8('F'),
	},
	3878: {
		Fword:  __ccgo_ts + 26885,
		Ftype1: int8('F'),
	},
	3879: {
		Fword:  __ccgo_ts + 26892,
		Ftype1: int8('F'),
	},
	3880: {
		Fword:  __ccgo_ts + 26899,
		Ftype1: int8('F'),
	},
	3881: {
		Fword:  __ccgo_ts + 26906,
		Ftype1: int8('F'),
	},
	3882: {
		Fword:  __ccgo_ts + 26913,
		Ftype1: int8('F'),
	},
	3883: {
		Fword:  __ccgo_ts + 26920,
		Ftype1: int8('F'),
	},
	3884: {
		Fword:  __ccgo_ts + 26927,
		Ftype1: int8('F'),
	},
	3885: {
		Fword:  __ccgo_ts + 26934,
		Ftype1: int8('F'),
	},
	3886: {
		Fword:  __ccgo_ts + 26941,
		Ftype1: int8('F'),
	},
	3887: {
		Fword:  __ccgo_ts + 26948,
		Ftype1: int8('F'),
	},
	3888: {
		Fword:  __ccgo_ts + 26955,
		Ftype1: int8('F'),
	},
	3889: {
		Fword:  __ccgo_ts + 26962,
		Ftype1: int8('F'),
	},
	3890: {
		Fword:  __ccgo_ts + 26969,
		Ftype1: int8('F'),
	},
	3891: {
		Fword:  __ccgo_ts + 26976,
		Ftype1: int8('F'),
	},
	3892: {
		Fword:  __ccgo_ts + 26983,
		Ftype1: int8('F'),
	},
	3893: {
		Fword:  __ccgo_ts + 26990,
		Ftype1: int8('F'),
	},
	3894: {
		Fword:  __ccgo_ts + 26997,
		Ftype1: int8('F'),
	},
	3895: {
		Fword:  __ccgo_ts + 27004,
		Ftype1: int8('F'),
	},
	3896: {
		Fword:  __ccgo_ts + 27011,
		Ftype1: int8('F'),
	},
	3897: {
		Fword:  __ccgo_ts + 27018,
		Ftype1: int8('F'),
	},
	3898: {
		Fword:  __ccgo_ts + 27025,
		Ftype1: int8('F'),
	},
	3899: {
		Fword:  __ccgo_ts + 27032,
		Ftype1: int8('F'),
	},
	3900: {
		Fword:  __ccgo_ts + 27039,
		Ftype1: int8('F'),
	},
	3901: {
		Fword:  __ccgo_ts + 27046,
		Ftype1: int8('F'),
	},
	3902: {
		Fword:  __ccgo_ts + 27053,
		Ftype1: int8('F'),
	},
	3903: {
		Fword:  __ccgo_ts + 27060,
		Ftype1: int8('F'),
	},
	3904: {
		Fword:  __ccgo_ts + 27067,
		Ftype1: int8('F'),
	},
	3905: {
		Fword:  __ccgo_ts + 27074,
		Ftype1: int8('F'),
	},
	3906: {
		Fword:  __ccgo_ts + 27081,
		Ftype1: int8('F'),
	},
	3907: {
		Fword:  __ccgo_ts + 27088,
		Ftype1: int8('F'),
	},
	3908: {
		Fword:  __ccgo_ts + 27095,
		Ftype1: int8('F'),
	},
	3909: {
		Fword:  __ccgo_ts + 27102,
		Ftype1: int8('F'),
	},
	3910: {
		Fword:  __ccgo_ts + 27109,
		Ftype1: int8('F'),
	},
	3911: {
		Fword:  __ccgo_ts + 27116,
		Ftype1: int8('F'),
	},
	3912: {
		Fword:  __ccgo_ts + 27123,
		Ftype1: int8('F'),
	},
	3913: {
		Fword:  __ccgo_ts + 27130,
		Ftype1: int8('F'),
	},
	3914: {
		Fword:  __ccgo_ts + 27137,
		Ftype1: int8('F'),
	},
	3915: {
		Fword:  __ccgo_ts + 27144,
		Ftype1: int8('F'),
	},
	3916: {
		Fword:  __ccgo_ts + 27151,
		Ftype1: int8('F'),
	},
	3917: {
		Fword:  __ccgo_ts + 27158,
		Ftype1: int8('F'),
	},
	3918: {
		Fword:  __ccgo_ts + 27165,
		Ftype1: int8('F'),
	},
	3919: {
		Fword:  __ccgo_ts + 27172,
		Ftype1: int8('F'),
	},
	3920: {
		Fword:  __ccgo_ts + 27179,
		Ftype1: int8('F'),
	},
	3921: {
		Fword:  __ccgo_ts + 27186,
		Ftype1: int8('F'),
	},
	3922: {
		Fword:  __ccgo_ts + 27193,
		Ftype1: int8('F'),
	},
	3923: {
		Fword:  __ccgo_ts + 27200,
		Ftype1: int8('F'),
	},
	3924: {
		Fword:  __ccgo_ts + 27207,
		Ftype1: int8('F'),
	},
	3925: {
		Fword:  __ccgo_ts + 27214,
		Ftype1: int8('F'),
	},
	3926: {
		Fword:  __ccgo_ts + 27221,
		Ftype1: int8('F'),
	},
	3927: {
		Fword:  __ccgo_ts + 27228,
		Ftype1: int8('F'),
	},
	3928: {
		Fword:  __ccgo_ts + 27235,
		Ftype1: int8('F'),
	},
	3929: {
		Fword:  __ccgo_ts + 27242,
		Ftype1: int8('F'),
	},
	3930: {
		Fword:  __ccgo_ts + 27249,
		Ftype1: int8('F'),
	},
	3931: {
		Fword:  __ccgo_ts + 27256,
		Ftype1: int8('F'),
	},
	3932: {
		Fword:  __ccgo_ts + 27263,
		Ftype1: int8('F'),
	},
	3933: {
		Fword:  __ccgo_ts + 27270,
		Ftype1: int8('F'),
	},
	3934: {
		Fword:  __ccgo_ts + 27275,
		Ftype1: int8('F'),
	},
	3935: {
		Fword:  __ccgo_ts + 27282,
		Ftype1: int8('F'),
	},
	3936: {
		Fword:  __ccgo_ts + 27289,
		Ftype1: int8('F'),
	},
	3937: {
		Fword:  __ccgo_ts + 27296,
		Ftype1: int8('F'),
	},
	3938: {
		Fword:  __ccgo_ts + 27303,
		Ftype1: int8('F'),
	},
	3939: {
		Fword:  __ccgo_ts + 27310,
		Ftype1: int8('F'),
	},
	3940: {
		Fword:  __ccgo_ts + 27317,
		Ftype1: int8('F'),
	},
	3941: {
		Fword:  __ccgo_ts + 27324,
		Ftype1: int8('F'),
	},
	3942: {
		Fword:  __ccgo_ts + 27331,
		Ftype1: int8('F'),
	},
	3943: {
		Fword:  __ccgo_ts + 27337,
		Ftype1: int8('F'),
	},
	3944: {
		Fword:  __ccgo_ts + 27344,
		Ftype1: int8('F'),
	},
	3945: {
		Fword:  __ccgo_ts + 27351,
		Ftype1: int8('F'),
	},
	3946: {
		Fword:  __ccgo_ts + 27358,
		Ftype1: int8('F'),
	},
	3947: {
		Fword:  __ccgo_ts + 27365,
		Ftype1: int8('F'),
	},
	3948: {
		Fword:  __ccgo_ts + 27372,
		Ftype1: int8('F'),
	},
	3949: {
		Fword:  __ccgo_ts + 27379,
		Ftype1: int8('F'),
	},
	3950: {
		Fword:  __ccgo_ts + 27386,
		Ftype1: int8('F'),
	},
	3951: {
		Fword:  __ccgo_ts + 27392,
		Ftype1: int8('F'),
	},
	3952: {
		Fword:  __ccgo_ts + 27399,
		Ftype1: int8('F'),
	},
	3953: {
		Fword:  __ccgo_ts + 27406,
		Ftype1: int8('F'),
	},
	3954: {
		Fword:  __ccgo_ts + 27413,
		Ftype1: int8('F'),
	},
	3955: {
		Fword:  __ccgo_ts + 27420,
		Ftype1: int8('F'),
	},
	3956: {
		Fword:  __ccgo_ts + 27427,
		Ftype1: int8('F'),
	},
	3957: {
		Fword:  __ccgo_ts + 27434,
		Ftype1: int8('F'),
	},
	3958: {
		Fword:  __ccgo_ts + 27441,
		Ftype1: int8('F'),
	},
	3959: {
		Fword:  __ccgo_ts + 27448,
		Ftype1: int8('F'),
	},
	3960: {
		Fword:  __ccgo_ts + 27455,
		Ftype1: int8('F'),
	},
	3961: {
		Fword:  __ccgo_ts + 27462,
		Ftype1: int8('F'),
	},
	3962: {
		Fword:  __ccgo_ts + 27469,
		Ftype1: int8('F'),
	},
	3963: {
		Fword:  __ccgo_ts + 27476,
		Ftype1: int8('F'),
	},
	3964: {
		Fword:  __ccgo_ts + 27483,
		Ftype1: int8('F'),
	},
	3965: {
		Fword:  __ccgo_ts + 27490,
		Ftype1: int8('F'),
	},
	3966: {
		Fword:  __ccgo_ts + 27497,
		Ftype1: int8('F'),
	},
	3967: {
		Fword:  __ccgo_ts + 27504,
		Ftype1: int8('F'),
	},
	3968: {
		Fword:  __ccgo_ts + 27511,
		Ftype1: int8('F'),
	},
	3969: {
		Fword:  __ccgo_ts + 27518,
		Ftype1: int8('F'),
	},
	3970: {
		Fword:  __ccgo_ts + 27525,
		Ftype1: int8('F'),
	},
	3971: {
		Fword:  __ccgo_ts + 27532,
		Ftype1: int8('F'),
	},
	3972: {
		Fword:  __ccgo_ts + 27539,
		Ftype1: int8('F'),
	},
	3973: {
		Fword:  __ccgo_ts + 27546,
		Ftype1: int8('F'),
	},
	3974: {
		Fword:  __ccgo_ts + 27553,
		Ftype1: int8('F'),
	},
	3975: {
		Fword:  __ccgo_ts + 27560,
		Ftype1: int8('F'),
	},
	3976: {
		Fword:  __ccgo_ts + 27567,
		Ftype1: int8('F'),
	},
	3977: {
		Fword:  __ccgo_ts + 27574,
		Ftype1: int8('F'),
	},
	3978: {
		Fword:  __ccgo_ts + 27581,
		Ftype1: int8('F'),
	},
	3979: {
		Fword:  __ccgo_ts + 27588,
		Ftype1: int8('F'),
	},
	3980: {
		Fword:  __ccgo_ts + 27595,
		Ftype1: int8('F'),
	},
	3981: {
		Fword:  __ccgo_ts + 27602,
		Ftype1: int8('F'),
	},
	3982: {
		Fword:  __ccgo_ts + 27609,
		Ftype1: int8('F'),
	},
	3983: {
		Fword:  __ccgo_ts + 27616,
		Ftype1: int8('F'),
	},
	3984: {
		Fword:  __ccgo_ts + 27622,
		Ftype1: int8('F'),
	},
	3985: {
		Fword:  __ccgo_ts + 27629,
		Ftype1: int8('F'),
	},
	3986: {
		Fword:  __ccgo_ts + 27636,
		Ftype1: int8('F'),
	},
	3987: {
		Fword:  __ccgo_ts + 27643,
		Ftype1: int8('F'),
	},
	3988: {
		Fword:  __ccgo_ts + 27650,
		Ftype1: int8('F'),
	},
	3989: {
		Fword:  __ccgo_ts + 27657,
		Ftype1: int8('F'),
	},
	3990: {
		Fword:  __ccgo_ts + 27664,
		Ftype1: int8('F'),
	},
	3991: {
		Fword:  __ccgo_ts + 27671,
		Ftype1: int8('F'),
	},
	3992: {
		Fword:  __ccgo_ts + 27677,
		Ftype1: int8('F'),
	},
	3993: {
		Fword:  __ccgo_ts + 27684,
		Ftype1: int8('F'),
	},
	3994: {
		Fword:  __ccgo_ts + 27691,
		Ftype1: int8('F'),
	},
	3995: {
		Fword:  __ccgo_ts + 27698,
		Ftype1: int8('F'),
	},
	3996: {
		Fword:  __ccgo_ts + 27705,
		Ftype1: int8('F'),
	},
	3997: {
		Fword:  __ccgo_ts + 27712,
		Ftype1: int8('F'),
	},
	3998: {
		Fword:  __ccgo_ts + 27719,
		Ftype1: int8('F'),
	},
	3999: {
		Fword:  __ccgo_ts + 27726,
		Ftype1: int8('F'),
	},
	4000: {
		Fword:  __ccgo_ts + 27733,
		Ftype1: int8('F'),
	},
	4001: {
		Fword:  __ccgo_ts + 27740,
		Ftype1: int8('F'),
	},
	4002: {
		Fword:  __ccgo_ts + 27747,
		Ftype1: int8('F'),
	},
	4003: {
		Fword:  __ccgo_ts + 27754,
		Ftype1: int8('F'),
	},
	4004: {
		Fword:  __ccgo_ts + 27761,
		Ftype1: int8('F'),
	},
	4005: {
		Fword:  __ccgo_ts + 27766,
		Ftype1: int8('F'),
	},
	4006: {
		Fword:  __ccgo_ts + 27773,
		Ftype1: int8('F'),
	},
	4007: {
		Fword:  __ccgo_ts + 27780,
		Ftype1: int8('F'),
	},
	4008: {
		Fword:  __ccgo_ts + 27787,
		Ftype1: int8('F'),
	},
	4009: {
		Fword:  __ccgo_ts + 27794,
		Ftype1: int8('F'),
	},
	4010: {
		Fword:  __ccgo_ts + 27801,
		Ftype1: int8('F'),
	},
	4011: {
		Fword:  __ccgo_ts + 27808,
		Ftype1: int8('F'),
	},
	4012: {
		Fword:  __ccgo_ts + 27815,
		Ftype1: int8('F'),
	},
	4013: {
		Fword:  __ccgo_ts + 27822,
		Ftype1: int8('F'),
	},
	4014: {
		Fword:  __ccgo_ts + 27828,
		Ftype1: int8('F'),
	},
	4015: {
		Fword:  __ccgo_ts + 27835,
		Ftype1: int8('F'),
	},
	4016: {
		Fword:  __ccgo_ts + 27842,
		Ftype1: int8('F'),
	},
	4017: {
		Fword:  __ccgo_ts + 27849,
		Ftype1: int8('F'),
	},
	4018: {
		Fword:  __ccgo_ts + 27856,
		Ftype1: int8('F'),
	},
	4019: {
		Fword:  __ccgo_ts + 27863,
		Ftype1: int8('F'),
	},
	4020: {
		Fword:  __ccgo_ts + 27870,
		Ftype1: int8('F'),
	},
	4021: {
		Fword:  __ccgo_ts + 27877,
		Ftype1: int8('F'),
	},
	4022: {
		Fword:  __ccgo_ts + 27883,
		Ftype1: int8('F'),
	},
	4023: {
		Fword:  __ccgo_ts + 27890,
		Ftype1: int8('F'),
	},
	4024: {
		Fword:  __ccgo_ts + 27897,
		Ftype1: int8('F'),
	},
	4025: {
		Fword:  __ccgo_ts + 27904,
		Ftype1: int8('F'),
	},
	4026: {
		Fword:  __ccgo_ts + 27911,
		Ftype1: int8('F'),
	},
	4027: {
		Fword:  __ccgo_ts + 27918,
		Ftype1: int8('F'),
	},
	4028: {
		Fword:  __ccgo_ts + 27925,
		Ftype1: int8('F'),
	},
	4029: {
		Fword:  __ccgo_ts + 27932,
		Ftype1: int8('F'),
	},
	4030: {
		Fword:  __ccgo_ts + 27939,
		Ftype1: int8('F'),
	},
	4031: {
		Fword:  __ccgo_ts + 27946,
		Ftype1: int8('F'),
	},
	4032: {
		Fword:  __ccgo_ts + 27953,
		Ftype1: int8('F'),
	},
	4033: {
		Fword:  __ccgo_ts + 27960,
		Ftype1: int8('F'),
	},
	4034: {
		Fword:  __ccgo_ts + 27967,
		Ftype1: int8('F'),
	},
	4035: {
		Fword:  __ccgo_ts + 27974,
		Ftype1: int8('F'),
	},
	4036: {
		Fword:  __ccgo_ts + 27981,
		Ftype1: int8('F'),
	},
	4037: {
		Fword:  __ccgo_ts + 27986,
		Ftype1: int8('F'),
	},
	4038: {
		Fword:  __ccgo_ts + 27993,
		Ftype1: int8('F'),
	},
	4039: {
		Fword:  __ccgo_ts + 28000,
		Ftype1: int8('F'),
	},
	4040: {
		Fword:  __ccgo_ts + 28007,
		Ftype1: int8('F'),
	},
	4041: {
		Fword:  __ccgo_ts + 28014,
		Ftype1: int8('F'),
	},
	4042: {
		Fword:  __ccgo_ts + 28021,
		Ftype1: int8('F'),
	},
	4043: {
		Fword:  __ccgo_ts + 28028,
		Ftype1: int8('F'),
	},
	4044: {
		Fword:  __ccgo_ts + 28035,
		Ftype1: int8('F'),
	},
	4045: {
		Fword:  __ccgo_ts + 28042,
		Ftype1: int8('F'),
	},
	4046: {
		Fword:  __ccgo_ts + 28048,
		Ftype1: int8('F'),
	},
	4047: {
		Fword:  __ccgo_ts + 28055,
		Ftype1: int8('F'),
	},
	4048: {
		Fword:  __ccgo_ts + 28062,
		Ftype1: int8('F'),
	},
	4049: {
		Fword:  __ccgo_ts + 28069,
		Ftype1: int8('F'),
	},
	4050: {
		Fword:  __ccgo_ts + 28076,
		Ftype1: int8('F'),
	},
	4051: {
		Fword:  __ccgo_ts + 28083,
		Ftype1: int8('F'),
	},
	4052: {
		Fword:  __ccgo_ts + 28090,
		Ftype1: int8('F'),
	},
	4053: {
		Fword:  __ccgo_ts + 28097,
		Ftype1: int8('F'),
	},
	4054: {
		Fword:  __ccgo_ts + 28103,
		Ftype1: int8('F'),
	},
	4055: {
		Fword:  __ccgo_ts + 28110,
		Ftype1: int8('F'),
	},
	4056: {
		Fword:  __ccgo_ts + 28117,
		Ftype1: int8('F'),
	},
	4057: {
		Fword:  __ccgo_ts + 28124,
		Ftype1: int8('F'),
	},
	4058: {
		Fword:  __ccgo_ts + 28131,
		Ftype1: int8('F'),
	},
	4059: {
		Fword:  __ccgo_ts + 28138,
		Ftype1: int8('F'),
	},
	4060: {
		Fword:  __ccgo_ts + 28145,
		Ftype1: int8('F'),
	},
	4061: {
		Fword:  __ccgo_ts + 28152,
		Ftype1: int8('F'),
	},
	4062: {
		Fword:  __ccgo_ts + 28159,
		Ftype1: int8('F'),
	},
	4063: {
		Fword:  __ccgo_ts + 28166,
		Ftype1: int8('F'),
	},
	4064: {
		Fword:  __ccgo_ts + 28173,
		Ftype1: int8('F'),
	},
	4065: {
		Fword:  __ccgo_ts + 28180,
		Ftype1: int8('F'),
	},
	4066: {
		Fword:  __ccgo_ts + 28184,
		Ftype1: int8('F'),
	},
	4067: {
		Fword:  __ccgo_ts + 28191,
		Ftype1: int8('F'),
	},
	4068: {
		Fword:  __ccgo_ts + 28198,
		Ftype1: int8('F'),
	},
	4069: {
		Fword:  __ccgo_ts + 28205,
		Ftype1: int8('F'),
	},
	4070: {
		Fword:  __ccgo_ts + 28212,
		Ftype1: int8('F'),
	},
	4071: {
		Fword:  __ccgo_ts + 28219,
		Ftype1: int8('F'),
	},
	4072: {
		Fword:  __ccgo_ts + 28226,
		Ftype1: int8('F'),
	},
	4073: {
		Fword:  __ccgo_ts + 28233,
		Ftype1: int8('F'),
	},
	4074: {
		Fword:  __ccgo_ts + 28240,
		Ftype1: int8('F'),
	},
	4075: {
		Fword:  __ccgo_ts + 28247,
		Ftype1: int8('F'),
	},
	4076: {
		Fword:  __ccgo_ts + 28254,
		Ftype1: int8('F'),
	},
	4077: {
		Fword:  __ccgo_ts + 28260,
		Ftype1: int8('F'),
	},
	4078: {
		Fword:  __ccgo_ts + 28267,
		Ftype1: int8('F'),
	},
	4079: {
		Fword:  __ccgo_ts + 28274,
		Ftype1: int8('F'),
	},
	4080: {
		Fword:  __ccgo_ts + 28281,
		Ftype1: int8('F'),
	},
	4081: {
		Fword:  __ccgo_ts + 28288,
		Ftype1: int8('F'),
	},
	4082: {
		Fword:  __ccgo_ts + 28295,
		Ftype1: int8('F'),
	},
	4083: {
		Fword:  __ccgo_ts + 28302,
		Ftype1: int8('F'),
	},
	4084: {
		Fword:  __ccgo_ts + 28309,
		Ftype1: int8('F'),
	},
	4085: {
		Fword:  __ccgo_ts + 28316,
		Ftype1: int8('F'),
	},
	4086: {
		Fword:  __ccgo_ts + 28323,
		Ftype1: int8('F'),
	},
	4087: {
		Fword:  __ccgo_ts + 28330,
		Ftype1: int8('F'),
	},
	4088: {
		Fword:  __ccgo_ts + 28337,
		Ftype1: int8('F'),
	},
	4089: {
		Fword:  __ccgo_ts + 28344,
		Ftype1: int8('F'),
	},
	4090: {
		Fword:  __ccgo_ts + 28351,
		Ftype1: int8('F'),
	},
	4091: {
		Fword:  __ccgo_ts + 28358,
		Ftype1: int8('F'),
	},
	4092: {
		Fword:  __ccgo_ts + 28365,
		Ftype1: int8('F'),
	},
	4093: {
		Fword:  __ccgo_ts + 28372,
		Ftype1: int8('F'),
	},
	4094: {
		Fword:  __ccgo_ts + 28379,
		Ftype1: int8('F'),
	},
	4095: {
		Fword:  __ccgo_ts + 28386,
		Ftype1: int8('F'),
	},
	4096: {
		Fword:  __ccgo_ts + 28393,
		Ftype1: int8('F'),
	},
	4097: {
		Fword:  __ccgo_ts + 28400,
		Ftype1: int8('F'),
	},
	4098: {
		Fword:  __ccgo_ts + 28407,
		Ftype1: int8('F'),
	},
	4099: {
		Fword:  __ccgo_ts + 28414,
		Ftype1: int8('F'),
	},
	4100: {
		Fword:  __ccgo_ts + 28421,
		Ftype1: int8('F'),
	},
	4101: {
		Fword:  __ccgo_ts + 28428,
		Ftype1: int8('F'),
	},
	4102: {
		Fword:  __ccgo_ts + 28435,
		Ftype1: int8('F'),
	},
	4103: {
		Fword:  __ccgo_ts + 28442,
		Ftype1: int8('F'),
	},
	4104: {
		Fword:  __ccgo_ts + 28449,
		Ftype1: int8('F'),
	},
	4105: {
		Fword:  __ccgo_ts + 28456,
		Ftype1: int8('F'),
	},
	4106: {
		Fword:  __ccgo_ts + 28463,
		Ftype1: int8('F'),
	},
	4107: {
		Fword:  __ccgo_ts + 28470,
		Ftype1: int8('F'),
	},
	4108: {
		Fword:  __ccgo_ts + 28477,
		Ftype1: int8('F'),
	},
	4109: {
		Fword:  __ccgo_ts + 28483,
		Ftype1: int8('F'),
	},
	4110: {
		Fword:  __ccgo_ts + 28490,
		Ftype1: int8('F'),
	},
	4111: {
		Fword:  __ccgo_ts + 28497,
		Ftype1: int8('F'),
	},
	4112: {
		Fword:  __ccgo_ts + 28504,
		Ftype1: int8('F'),
	},
	4113: {
		Fword:  __ccgo_ts + 28511,
		Ftype1: int8('F'),
	},
	4114: {
		Fword:  __ccgo_ts + 28518,
		Ftype1: int8('F'),
	},
	4115: {
		Fword:  __ccgo_ts + 28525,
		Ftype1: int8('F'),
	},
	4116: {
		Fword:  __ccgo_ts + 28532,
		Ftype1: int8('F'),
	},
	4117: {
		Fword:  __ccgo_ts + 28539,
		Ftype1: int8('F'),
	},
	4118: {
		Fword:  __ccgo_ts + 28546,
		Ftype1: int8('F'),
	},
	4119: {
		Fword:  __ccgo_ts + 28553,
		Ftype1: int8('F'),
	},
	4120: {
		Fword:  __ccgo_ts + 28560,
		Ftype1: int8('F'),
	},
	4121: {
		Fword:  __ccgo_ts + 28567,
		Ftype1: int8('F'),
	},
	4122: {
		Fword:  __ccgo_ts + 28574,
		Ftype1: int8('F'),
	},
	4123: {
		Fword:  __ccgo_ts + 28581,
		Ftype1: int8('F'),
	},
	4124: {
		Fword:  __ccgo_ts + 28588,
		Ftype1: int8('F'),
	},
	4125: {
		Fword:  __ccgo_ts + 28595,
		Ftype1: int8('F'),
	},
	4126: {
		Fword:  __ccgo_ts + 28602,
		Ftype1: int8('F'),
	},
	4127: {
		Fword:  __ccgo_ts + 28609,
		Ftype1: int8('F'),
	},
	4128: {
		Fword:  __ccgo_ts + 28616,
		Ftype1: int8('F'),
	},
	4129: {
		Fword:  __ccgo_ts + 28623,
		Ftype1: int8('F'),
	},
	4130: {
		Fword:  __ccgo_ts + 28630,
		Ftype1: int8('F'),
	},
	4131: {
		Fword:  __ccgo_ts + 28637,
		Ftype1: int8('F'),
	},
	4132: {
		Fword:  __ccgo_ts + 28644,
		Ftype1: int8('F'),
	},
	4133: {
		Fword:  __ccgo_ts + 28651,
		Ftype1: int8('F'),
	},
	4134: {
		Fword:  __ccgo_ts + 28658,
		Ftype1: int8('F'),
	},
	4135: {
		Fword:  __ccgo_ts + 28665,
		Ftype1: int8('F'),
	},
	4136: {
		Fword:  __ccgo_ts + 28672,
		Ftype1: int8('F'),
	},
	4137: {
		Fword:  __ccgo_ts + 28679,
		Ftype1: int8('F'),
	},
	4138: {
		Fword:  __ccgo_ts + 28686,
		Ftype1: int8('F'),
	},
	4139: {
		Fword:  __ccgo_ts + 28693,
		Ftype1: int8('F'),
	},
	4140: {
		Fword:  __ccgo_ts + 28700,
		Ftype1: int8('F'),
	},
	4141: {
		Fword:  __ccgo_ts + 28706,
		Ftype1: int8('F'),
	},
	4142: {
		Fword:  __ccgo_ts + 28713,
		Ftype1: int8('F'),
	},
	4143: {
		Fword:  __ccgo_ts + 28720,
		Ftype1: int8('F'),
	},
	4144: {
		Fword:  __ccgo_ts + 28727,
		Ftype1: int8('F'),
	},
	4145: {
		Fword:  __ccgo_ts + 28734,
		Ftype1: int8('F'),
	},
	4146: {
		Fword:  __ccgo_ts + 28741,
		Ftype1: int8('F'),
	},
	4147: {
		Fword:  __ccgo_ts + 28748,
		Ftype1: int8('F'),
	},
	4148: {
		Fword:  __ccgo_ts + 28755,
		Ftype1: int8('F'),
	},
	4149: {
		Fword:  __ccgo_ts + 28762,
		Ftype1: int8('F'),
	},
	4150: {
		Fword:  __ccgo_ts + 28769,
		Ftype1: int8('F'),
	},
	4151: {
		Fword:  __ccgo_ts + 28776,
		Ftype1: int8('F'),
	},
	4152: {
		Fword:  __ccgo_ts + 28783,
		Ftype1: int8('F'),
	},
	4153: {
		Fword:  __ccgo_ts + 28790,
		Ftype1: int8('F'),
	},
	4154: {
		Fword:  __ccgo_ts + 28797,
		Ftype1: int8('F'),
	},
	4155: {
		Fword:  __ccgo_ts + 28804,
		Ftype1: int8('F'),
	},
	4156: {
		Fword:  __ccgo_ts + 28811,
		Ftype1: int8('F'),
	},
	4157: {
		Fword:  __ccgo_ts + 28818,
		Ftype1: int8('F'),
	},
	4158: {
		Fword:  __ccgo_ts + 28825,
		Ftype1: int8('F'),
	},
	4159: {
		Fword:  __ccgo_ts + 28832,
		Ftype1: int8('F'),
	},
	4160: {
		Fword:  __ccgo_ts + 28839,
		Ftype1: int8('F'),
	},
	4161: {
		Fword:  __ccgo_ts + 28846,
		Ftype1: int8('F'),
	},
	4162: {
		Fword:  __ccgo_ts + 28853,
		Ftype1: int8('F'),
	},
	4163: {
		Fword:  __ccgo_ts + 28860,
		Ftype1: int8('F'),
	},
	4164: {
		Fword:  __ccgo_ts + 28867,
		Ftype1: int8('F'),
	},
	4165: {
		Fword:  __ccgo_ts + 28874,
		Ftype1: int8('F'),
	},
	4166: {
		Fword:  __ccgo_ts + 28881,
		Ftype1: int8('F'),
	},
	4167: {
		Fword:  __ccgo_ts + 28888,
		Ftype1: int8('F'),
	},
	4168: {
		Fword:  __ccgo_ts + 28895,
		Ftype1: int8('F'),
	},
	4169: {
		Fword:  __ccgo_ts + 28902,
		Ftype1: int8('F'),
	},
	4170: {
		Fword:  __ccgo_ts + 28909,
		Ftype1: int8('F'),
	},
	4171: {
		Fword:  __ccgo_ts + 28916,
		Ftype1: int8('F'),
	},
	4172: {
		Fword:  __ccgo_ts + 28923,
		Ftype1: int8('F'),
	},
	4173: {
		Fword:  __ccgo_ts + 28930,
		Ftype1: int8('F'),
	},
	4174: {
		Fword:  __ccgo_ts + 28937,
		Ftype1: int8('F'),
	},
	4175: {
		Fword:  __ccgo_ts + 28944,
		Ftype1: int8('F'),
	},
	4176: {
		Fword:  __ccgo_ts + 28951,
		Ftype1: int8('F'),
	},
	4177: {
		Fword:  __ccgo_ts + 28958,
		Ftype1: int8('F'),
	},
	4178: {
		Fword:  __ccgo_ts + 28965,
		Ftype1: int8('F'),
	},
	4179: {
		Fword:  __ccgo_ts + 28972,
		Ftype1: int8('F'),
	},
	4180: {
		Fword:  __ccgo_ts + 28979,
		Ftype1: int8('F'),
	},
	4181: {
		Fword:  __ccgo_ts + 28986,
		Ftype1: int8('F'),
	},
	4182: {
		Fword:  __ccgo_ts + 28993,
		Ftype1: int8('F'),
	},
	4183: {
		Fword:  __ccgo_ts + 29000,
		Ftype1: int8('F'),
	},
	4184: {
		Fword:  __ccgo_ts + 29007,
		Ftype1: int8('F'),
	},
	4185: {
		Fword:  __ccgo_ts + 29014,
		Ftype1: int8('F'),
	},
	4186: {
		Fword:  __ccgo_ts + 29021,
		Ftype1: int8('F'),
	},
	4187: {
		Fword:  __ccgo_ts + 29028,
		Ftype1: int8('F'),
	},
	4188: {
		Fword:  __ccgo_ts + 29035,
		Ftype1: int8('F'),
	},
	4189: {
		Fword:  __ccgo_ts + 29042,
		Ftype1: int8('F'),
	},
	4190: {
		Fword:  __ccgo_ts + 29049,
		Ftype1: int8('F'),
	},
	4191: {
		Fword:  __ccgo_ts + 29056,
		Ftype1: int8('F'),
	},
	4192: {
		Fword:  __ccgo_ts + 29063,
		Ftype1: int8('F'),
	},
	4193: {
		Fword:  __ccgo_ts + 29070,
		Ftype1: int8('F'),
	},
	4194: {
		Fword:  __ccgo_ts + 29077,
		Ftype1: int8('F'),
	},
	4195: {
		Fword:  __ccgo_ts + 29084,
		Ftype1: int8('F'),
	},
	4196: {
		Fword:  __ccgo_ts + 29091,
		Ftype1: int8('F'),
	},
	4197: {
		Fword:  __ccgo_ts + 29098,
		Ftype1: int8('F'),
	},
	4198: {
		Fword:  __ccgo_ts + 29105,
		Ftype1: int8('F'),
	},
	4199: {
		Fword:  __ccgo_ts + 29112,
		Ftype1: int8('F'),
	},
	4200: {
		Fword:  __ccgo_ts + 29119,
		Ftype1: int8('F'),
	},
	4201: {
		Fword:  __ccgo_ts + 29126,
		Ftype1: int8('F'),
	},
	4202: {
		Fword:  __ccgo_ts + 29133,
		Ftype1: int8('F'),
	},
	4203: {
		Fword:  __ccgo_ts + 29140,
		Ftype1: int8('F'),
	},
	4204: {
		Fword:  __ccgo_ts + 29147,
		Ftype1: int8('F'),
	},
	4205: {
		Fword:  __ccgo_ts + 29154,
		Ftype1: int8('F'),
	},
	4206: {
		Fword:  __ccgo_ts + 29161,
		Ftype1: int8('F'),
	},
	4207: {
		Fword:  __ccgo_ts + 29168,
		Ftype1: int8('F'),
	},
	4208: {
		Fword:  __ccgo_ts + 29175,
		Ftype1: int8('F'),
	},
	4209: {
		Fword:  __ccgo_ts + 29180,
		Ftype1: int8('F'),
	},
	4210: {
		Fword:  __ccgo_ts + 29187,
		Ftype1: int8('F'),
	},
	4211: {
		Fword:  __ccgo_ts + 29194,
		Ftype1: int8('F'),
	},
	4212: {
		Fword:  __ccgo_ts + 29201,
		Ftype1: int8('F'),
	},
	4213: {
		Fword:  __ccgo_ts + 29208,
		Ftype1: int8('F'),
	},
	4214: {
		Fword:  __ccgo_ts + 29215,
		Ftype1: int8('F'),
	},
	4215: {
		Fword:  __ccgo_ts + 29222,
		Ftype1: int8('F'),
	},
	4216: {
		Fword:  __ccgo_ts + 29229,
		Ftype1: int8('F'),
	},
	4217: {
		Fword:  __ccgo_ts + 29236,
		Ftype1: int8('F'),
	},
	4218: {
		Fword:  __ccgo_ts + 29243,
		Ftype1: int8('F'),
	},
	4219: {
		Fword:  __ccgo_ts + 29250,
		Ftype1: int8('F'),
	},
	4220: {
		Fword:  __ccgo_ts + 29257,
		Ftype1: int8('F'),
	},
	4221: {
		Fword:  __ccgo_ts + 29264,
		Ftype1: int8('F'),
	},
	4222: {
		Fword:  __ccgo_ts + 29271,
		Ftype1: int8('F'),
	},
	4223: {
		Fword:  __ccgo_ts + 29278,
		Ftype1: int8('F'),
	},
	4224: {
		Fword:  __ccgo_ts + 29285,
		Ftype1: int8('F'),
	},
	4225: {
		Fword:  __ccgo_ts + 29291,
		Ftype1: int8('F'),
	},
	4226: {
		Fword:  __ccgo_ts + 29298,
		Ftype1: int8('F'),
	},
	4227: {
		Fword:  __ccgo_ts + 29305,
		Ftype1: int8('F'),
	},
	4228: {
		Fword:  __ccgo_ts + 29312,
		Ftype1: int8('F'),
	},
	4229: {
		Fword:  __ccgo_ts + 29319,
		Ftype1: int8('F'),
	},
	4230: {
		Fword:  __ccgo_ts + 29326,
		Ftype1: int8('F'),
	},
	4231: {
		Fword:  __ccgo_ts + 29333,
		Ftype1: int8('F'),
	},
	4232: {
		Fword:  __ccgo_ts + 29340,
		Ftype1: int8('F'),
	},
	4233: {
		Fword:  __ccgo_ts + 29347,
		Ftype1: int8('F'),
	},
	4234: {
		Fword:  __ccgo_ts + 29354,
		Ftype1: int8('F'),
	},
	4235: {
		Fword:  __ccgo_ts + 29361,
		Ftype1: int8('F'),
	},
	4236: {
		Fword:  __ccgo_ts + 29368,
		Ftype1: int8('F'),
	},
	4237: {
		Fword:  __ccgo_ts + 29375,
		Ftype1: int8('F'),
	},
	4238: {
		Fword:  __ccgo_ts + 29382,
		Ftype1: int8('F'),
	},
	4239: {
		Fword:  __ccgo_ts + 29389,
		Ftype1: int8('F'),
	},
	4240: {
		Fword:  __ccgo_ts + 29396,
		Ftype1: int8('F'),
	},
	4241: {
		Fword:  __ccgo_ts + 29403,
		Ftype1: int8('F'),
	},
	4242: {
		Fword:  __ccgo_ts + 29410,
		Ftype1: int8('F'),
	},
	4243: {
		Fword:  __ccgo_ts + 29417,
		Ftype1: int8('F'),
	},
	4244: {
		Fword:  __ccgo_ts + 29424,
		Ftype1: int8('F'),
	},
	4245: {
		Fword:  __ccgo_ts + 29429,
		Ftype1: int8('F'),
	},
	4246: {
		Fword:  __ccgo_ts + 29436,
		Ftype1: int8('F'),
	},
	4247: {
		Fword:  __ccgo_ts + 29443,
		Ftype1: int8('F'),
	},
	4248: {
		Fword:  __ccgo_ts + 29450,
		Ftype1: int8('F'),
	},
	4249: {
		Fword:  __ccgo_ts + 29457,
		Ftype1: int8('F'),
	},
	4250: {
		Fword:  __ccgo_ts + 29464,
		Ftype1: int8('F'),
	},
	4251: {
		Fword:  __ccgo_ts + 29471,
		Ftype1: int8('F'),
	},
	4252: {
		Fword:  __ccgo_ts + 29478,
		Ftype1: int8('F'),
	},
	4253: {
		Fword:  __ccgo_ts + 29485,
		Ftype1: int8('F'),
	},
	4254: {
		Fword:  __ccgo_ts + 29492,
		Ftype1: int8('F'),
	},
	4255: {
		Fword:  __ccgo_ts + 29499,
		Ftype1: int8('F'),
	},
	4256: {
		Fword:  __ccgo_ts + 29506,
		Ftype1: int8('F'),
	},
	4257: {
		Fword:  __ccgo_ts + 29513,
		Ftype1: int8('F'),
	},
	4258: {
		Fword:  __ccgo_ts + 29520,
		Ftype1: int8('F'),
	},
	4259: {
		Fword:  __ccgo_ts + 29527,
		Ftype1: int8('F'),
	},
	4260: {
		Fword:  __ccgo_ts + 29534,
		Ftype1: int8('F'),
	},
	4261: {
		Fword:  __ccgo_ts + 29541,
		Ftype1: int8('F'),
	},
	4262: {
		Fword:  __ccgo_ts + 29548,
		Ftype1: int8('F'),
	},
	4263: {
		Fword:  __ccgo_ts + 29555,
		Ftype1: int8('F'),
	},
	4264: {
		Fword:  __ccgo_ts + 29562,
		Ftype1: int8('F'),
	},
	4265: {
		Fword:  __ccgo_ts + 29569,
		Ftype1: int8('F'),
	},
	4266: {
		Fword:  __ccgo_ts + 29576,
		Ftype1: int8('F'),
	},
	4267: {
		Fword:  __ccgo_ts + 29581,
		Ftype1: int8('F'),
	},
	4268: {
		Fword:  __ccgo_ts + 29588,
		Ftype1: int8('F'),
	},
	4269: {
		Fword:  __ccgo_ts + 29595,
		Ftype1: int8('F'),
	},
	4270: {
		Fword:  __ccgo_ts + 29602,
		Ftype1: int8('F'),
	},
	4271: {
		Fword:  __ccgo_ts + 29609,
		Ftype1: int8('F'),
	},
	4272: {
		Fword:  __ccgo_ts + 29616,
		Ftype1: int8('F'),
	},
	4273: {
		Fword:  __ccgo_ts + 29623,
		Ftype1: int8('F'),
	},
	4274: {
		Fword:  __ccgo_ts + 29629,
		Ftype1: int8('F'),
	},
	4275: {
		Fword:  __ccgo_ts + 29636,
		Ftype1: int8('F'),
	},
	4276: {
		Fword:  __ccgo_ts + 29643,
		Ftype1: int8('F'),
	},
	4277: {
		Fword:  __ccgo_ts + 29650,
		Ftype1: int8('F'),
	},
	4278: {
		Fword:  __ccgo_ts + 29657,
		Ftype1: int8('F'),
	},
	4279: {
		Fword:  __ccgo_ts + 29664,
		Ftype1: int8('F'),
	},
	4280: {
		Fword:  __ccgo_ts + 29671,
		Ftype1: int8('F'),
	},
	4281: {
		Fword:  __ccgo_ts + 29678,
		Ftype1: int8('F'),
	},
	4282: {
		Fword:  __ccgo_ts + 29685,
		Ftype1: int8('F'),
	},
	4283: {
		Fword:  __ccgo_ts + 29692,
		Ftype1: int8('F'),
	},
	4284: {
		Fword:  __ccgo_ts + 29698,
		Ftype1: int8('F'),
	},
	4285: {
		Fword:  __ccgo_ts + 29705,
		Ftype1: int8('F'),
	},
	4286: {
		Fword:  __ccgo_ts + 29712,
		Ftype1: int8('F'),
	},
	4287: {
		Fword:  __ccgo_ts + 29719,
		Ftype1: int8('F'),
	},
	4288: {
		Fword:  __ccgo_ts + 29726,
		Ftype1: int8('F'),
	},
	4289: {
		Fword:  __ccgo_ts + 29733,
		Ftype1: int8('F'),
	},
	4290: {
		Fword:  __ccgo_ts + 29740,
		Ftype1: int8('F'),
	},
	4291: {
		Fword:  __ccgo_ts + 29747,
		Ftype1: int8('F'),
	},
	4292: {
		Fword:  __ccgo_ts + 29754,
		Ftype1: int8('F'),
	},
	4293: {
		Fword:  __ccgo_ts + 29761,
		Ftype1: int8('F'),
	},
	4294: {
		Fword:  __ccgo_ts + 29768,
		Ftype1: int8('F'),
	},
	4295: {
		Fword:  __ccgo_ts + 29775,
		Ftype1: int8('F'),
	},
	4296: {
		Fword:  __ccgo_ts + 29782,
		Ftype1: int8('F'),
	},
	4297: {
		Fword:  __ccgo_ts + 29789,
		Ftype1: int8('F'),
	},
	4298: {
		Fword:  __ccgo_ts + 29796,
		Ftype1: int8('F'),
	},
	4299: {
		Fword:  __ccgo_ts + 29803,
		Ftype1: int8('F'),
	},
	4300: {
		Fword:  __ccgo_ts + 29810,
		Ftype1: int8('F'),
	},
	4301: {
		Fword:  __ccgo_ts + 29817,
		Ftype1: int8('F'),
	},
	4302: {
		Fword:  __ccgo_ts + 29824,
		Ftype1: int8('F'),
	},
	4303: {
		Fword:  __ccgo_ts + 29831,
		Ftype1: int8('F'),
	},
	4304: {
		Fword:  __ccgo_ts + 29838,
		Ftype1: int8('F'),
	},
	4305: {
		Fword:  __ccgo_ts + 29845,
		Ftype1: int8('F'),
	},
	4306: {
		Fword:  __ccgo_ts + 29852,
		Ftype1: int8('F'),
	},
	4307: {
		Fword:  __ccgo_ts + 29857,
		Ftype1: int8('F'),
	},
	4308: {
		Fword:  __ccgo_ts + 29864,
		Ftype1: int8('F'),
	},
	4309: {
		Fword:  __ccgo_ts + 29871,
		Ftype1: int8('F'),
	},
	4310: {
		Fword:  __ccgo_ts + 29878,
		Ftype1: int8('F'),
	},
	4311: {
		Fword:  __ccgo_ts + 29885,
		Ftype1: int8('F'),
	},
	4312: {
		Fword:  __ccgo_ts + 29892,
		Ftype1: int8('F'),
	},
	4313: {
		Fword:  __ccgo_ts + 29899,
		Ftype1: int8('F'),
	},
	4314: {
		Fword:  __ccgo_ts + 29905,
		Ftype1: int8('F'),
	},
	4315: {
		Fword:  __ccgo_ts + 29912,
		Ftype1: int8('F'),
	},
	4316: {
		Fword:  __ccgo_ts + 29919,
		Ftype1: int8('F'),
	},
	4317: {
		Fword:  __ccgo_ts + 29926,
		Ftype1: int8('F'),
	},
	4318: {
		Fword:  __ccgo_ts + 29933,
		Ftype1: int8('F'),
	},
	4319: {
		Fword:  __ccgo_ts + 29940,
		Ftype1: int8('F'),
	},
	4320: {
		Fword:  __ccgo_ts + 29947,
		Ftype1: int8('F'),
	},
	4321: {
		Fword:  __ccgo_ts + 29954,
		Ftype1: int8('F'),
	},
	4322: {
		Fword:  __ccgo_ts + 29961,
		Ftype1: int8('F'),
	},
	4323: {
		Fword:  __ccgo_ts + 29968,
		Ftype1: int8('F'),
	},
	4324: {
		Fword:  __ccgo_ts + 29974,
		Ftype1: int8('F'),
	},
	4325: {
		Fword:  __ccgo_ts + 29981,
		Ftype1: int8('F'),
	},
	4326: {
		Fword:  __ccgo_ts + 29988,
		Ftype1: int8('F'),
	},
	4327: {
		Fword:  __ccgo_ts + 29995,
		Ftype1: int8('F'),
	},
	4328: {
		Fword:  __ccgo_ts + 30002,
		Ftype1: int8('F'),
	},
	4329: {
		Fword:  __ccgo_ts + 30009,
		Ftype1: int8('F'),
	},
	4330: {
		Fword:  __ccgo_ts + 30016,
		Ftype1: int8('F'),
	},
	4331: {
		Fword:  __ccgo_ts + 30023,
		Ftype1: int8('F'),
	},
	4332: {
		Fword:  __ccgo_ts + 30030,
		Ftype1: int8('F'),
	},
	4333: {
		Fword:  __ccgo_ts + 30037,
		Ftype1: int8('F'),
	},
	4334: {
		Fword:  __ccgo_ts + 30044,
		Ftype1: int8('F'),
	},
	4335: {
		Fword:  __ccgo_ts + 30051,
		Ftype1: int8('F'),
	},
	4336: {
		Fword:  __ccgo_ts + 30058,
		Ftype1: int8('F'),
	},
	4337: {
		Fword:  __ccgo_ts + 30065,
		Ftype1: int8('F'),
	},
	4338: {
		Fword:  __ccgo_ts + 30072,
		Ftype1: int8('F'),
	},
	4339: {
		Fword:  __ccgo_ts + 30079,
		Ftype1: int8('F'),
	},
	4340: {
		Fword:  __ccgo_ts + 30086,
		Ftype1: int8('F'),
	},
	4341: {
		Fword:  __ccgo_ts + 30093,
		Ftype1: int8('F'),
	},
	4342: {
		Fword:  __ccgo_ts + 30100,
		Ftype1: int8('F'),
	},
	4343: {
		Fword:  __ccgo_ts + 30107,
		Ftype1: int8('F'),
	},
	4344: {
		Fword:  __ccgo_ts + 30114,
		Ftype1: int8('F'),
	},
	4345: {
		Fword:  __ccgo_ts + 30121,
		Ftype1: int8('F'),
	},
	4346: {
		Fword:  __ccgo_ts + 30128,
		Ftype1: int8('F'),
	},
	4347: {
		Fword:  __ccgo_ts + 30135,
		Ftype1: int8('F'),
	},
	4348: {
		Fword:  __ccgo_ts + 30142,
		Ftype1: int8('F'),
	},
	4349: {
		Fword:  __ccgo_ts + 30149,
		Ftype1: int8('F'),
	},
	4350: {
		Fword:  __ccgo_ts + 30156,
		Ftype1: int8('F'),
	},
	4351: {
		Fword:  __ccgo_ts + 30163,
		Ftype1: int8('F'),
	},
	4352: {
		Fword:  __ccgo_ts + 30170,
		Ftype1: int8('F'),
	},
	4353: {
		Fword:  __ccgo_ts + 30177,
		Ftype1: int8('F'),
	},
	4354: {
		Fword:  __ccgo_ts + 30184,
		Ftype1: int8('F'),
	},
	4355: {
		Fword:  __ccgo_ts + 30191,
		Ftype1: int8('F'),
	},
	4356: {
		Fword:  __ccgo_ts + 30198,
		Ftype1: int8('F'),
	},
	4357: {
		Fword:  __ccgo_ts + 30205,
		Ftype1: int8('F'),
	},
	4358: {
		Fword:  __ccgo_ts + 30212,
		Ftype1: int8('F'),
	},
	4359: {
		Fword:  __ccgo_ts + 30219,
		Ftype1: int8('F'),
	},
	4360: {
		Fword:  __ccgo_ts + 30226,
		Ftype1: int8('F'),
	},
	4361: {
		Fword:  __ccgo_ts + 30233,
		Ftype1: int8('F'),
	},
	4362: {
		Fword:  __ccgo_ts + 30240,
		Ftype1: int8('F'),
	},
	4363: {
		Fword:  __ccgo_ts + 30247,
		Ftype1: int8('F'),
	},
	4364: {
		Fword:  __ccgo_ts + 30254,
		Ftype1: int8('F'),
	},
	4365: {
		Fword:  __ccgo_ts + 30261,
		Ftype1: int8('F'),
	},
	4366: {
		Fword:  __ccgo_ts + 30268,
		Ftype1: int8('F'),
	},
	4367: {
		Fword:  __ccgo_ts + 30275,
		Ftype1: int8('F'),
	},
	4368: {
		Fword:  __ccgo_ts + 30282,
		Ftype1: int8('F'),
	},
	4369: {
		Fword:  __ccgo_ts + 30289,
		Ftype1: int8('F'),
	},
	4370: {
		Fword:  __ccgo_ts + 30296,
		Ftype1: int8('F'),
	},
	4371: {
		Fword:  __ccgo_ts + 30303,
		Ftype1: int8('F'),
	},
	4372: {
		Fword:  __ccgo_ts + 30310,
		Ftype1: int8('F'),
	},
	4373: {
		Fword:  __ccgo_ts + 30317,
		Ftype1: int8('F'),
	},
	4374: {
		Fword:  __ccgo_ts + 30324,
		Ftype1: int8('F'),
	},
	4375: {
		Fword:  __ccgo_ts + 30331,
		Ftype1: int8('F'),
	},
	4376: {
		Fword:  __ccgo_ts + 30338,
		Ftype1: int8('F'),
	},
	4377: {
		Fword:  __ccgo_ts + 30345,
		Ftype1: int8('F'),
	},
	4378: {
		Fword:  __ccgo_ts + 30352,
		Ftype1: int8('F'),
	},
	4379: {
		Fword:  __ccgo_ts + 30359,
		Ftype1: int8('F'),
	},
	4380: {
		Fword:  __ccgo_ts + 30366,
		Ftype1: int8('F'),
	},
	4381: {
		Fword:  __ccgo_ts + 30373,
		Ftype1: int8('F'),
	},
	4382: {
		Fword:  __ccgo_ts + 30380,
		Ftype1: int8('F'),
	},
	4383: {
		Fword:  __ccgo_ts + 30387,
		Ftype1: int8('F'),
	},
	4384: {
		Fword:  __ccgo_ts + 30394,
		Ftype1: int8('F'),
	},
	4385: {
		Fword:  __ccgo_ts + 30401,
		Ftype1: int8('F'),
	},
	4386: {
		Fword:  __ccgo_ts + 30408,
		Ftype1: int8('F'),
	},
	4387: {
		Fword:  __ccgo_ts + 30415,
		Ftype1: int8('F'),
	},
	4388: {
		Fword:  __ccgo_ts + 30422,
		Ftype1: int8('F'),
	},
	4389: {
		Fword:  __ccgo_ts + 30429,
		Ftype1: int8('F'),
	},
	4390: {
		Fword:  __ccgo_ts + 30436,
		Ftype1: int8('F'),
	},
	4391: {
		Fword:  __ccgo_ts + 30443,
		Ftype1: int8('F'),
	},
	4392: {
		Fword:  __ccgo_ts + 30450,
		Ftype1: int8('F'),
	},
	4393: {
		Fword:  __ccgo_ts + 30457,
		Ftype1: int8('F'),
	},
	4394: {
		Fword:  __ccgo_ts + 30464,
		Ftype1: int8('F'),
	},
	4395: {
		Fword:  __ccgo_ts + 30471,
		Ftype1: int8('F'),
	},
	4396: {
		Fword:  __ccgo_ts + 30478,
		Ftype1: int8('F'),
	},
	4397: {
		Fword:  __ccgo_ts + 30485,
		Ftype1: int8('F'),
	},
	4398: {
		Fword:  __ccgo_ts + 30492,
		Ftype1: int8('F'),
	},
	4399: {
		Fword:  __ccgo_ts + 30499,
		Ftype1: int8('F'),
	},
	4400: {
		Fword:  __ccgo_ts + 30506,
		Ftype1: int8('F'),
	},
	4401: {
		Fword:  __ccgo_ts + 30513,
		Ftype1: int8('F'),
	},
	4402: {
		Fword:  __ccgo_ts + 30520,
		Ftype1: int8('F'),
	},
	4403: {
		Fword:  __ccgo_ts + 30527,
		Ftype1: int8('F'),
	},
	4404: {
		Fword:  __ccgo_ts + 30534,
		Ftype1: int8('F'),
	},
	4405: {
		Fword:  __ccgo_ts + 30541,
		Ftype1: int8('F'),
	},
	4406: {
		Fword:  __ccgo_ts + 30548,
		Ftype1: int8('F'),
	},
	4407: {
		Fword:  __ccgo_ts + 30555,
		Ftype1: int8('F'),
	},
	4408: {
		Fword:  __ccgo_ts + 30562,
		Ftype1: int8('F'),
	},
	4409: {
		Fword:  __ccgo_ts + 30569,
		Ftype1: int8('F'),
	},
	4410: {
		Fword:  __ccgo_ts + 30576,
		Ftype1: int8('F'),
	},
	4411: {
		Fword:  __ccgo_ts + 30583,
		Ftype1: int8('F'),
	},
	4412: {
		Fword:  __ccgo_ts + 30590,
		Ftype1: int8('F'),
	},
	4413: {
		Fword:  __ccgo_ts + 30597,
		Ftype1: int8('F'),
	},
	4414: {
		Fword:  __ccgo_ts + 30604,
		Ftype1: int8('F'),
	},
	4415: {
		Fword:  __ccgo_ts + 30611,
		Ftype1: int8('F'),
	},
	4416: {
		Fword:  __ccgo_ts + 30618,
		Ftype1: int8('F'),
	},
	4417: {
		Fword:  __ccgo_ts + 30625,
		Ftype1: int8('F'),
	},
	4418: {
		Fword:  __ccgo_ts + 30632,
		Ftype1: int8('F'),
	},
	4419: {
		Fword:  __ccgo_ts + 30639,
		Ftype1: int8('F'),
	},
	4420: {
		Fword:  __ccgo_ts + 30645,
		Ftype1: int8('F'),
	},
	4421: {
		Fword:  __ccgo_ts + 30652,
		Ftype1: int8('F'),
	},
	4422: {
		Fword:  __ccgo_ts + 30659,
		Ftype1: int8('F'),
	},
	4423: {
		Fword:  __ccgo_ts + 30666,
		Ftype1: int8('F'),
	},
	4424: {
		Fword:  __ccgo_ts + 30673,
		Ftype1: int8('F'),
	},
	4425: {
		Fword:  __ccgo_ts + 30680,
		Ftype1: int8('F'),
	},
	4426: {
		Fword:  __ccgo_ts + 30687,
		Ftype1: int8('F'),
	},
	4427: {
		Fword:  __ccgo_ts + 30694,
		Ftype1: int8('F'),
	},
	4428: {
		Fword:  __ccgo_ts + 30701,
		Ftype1: int8('F'),
	},
	4429: {
		Fword:  __ccgo_ts + 30708,
		Ftype1: int8('F'),
	},
	4430: {
		Fword:  __ccgo_ts + 30715,
		Ftype1: int8('F'),
	},
	4431: {
		Fword:  __ccgo_ts + 30722,
		Ftype1: int8('F'),
	},
	4432: {
		Fword:  __ccgo_ts + 30729,
		Ftype1: int8('F'),
	},
	4433: {
		Fword:  __ccgo_ts + 30736,
		Ftype1: int8('F'),
	},
	4434: {
		Fword:  __ccgo_ts + 30743,
		Ftype1: int8('F'),
	},
	4435: {
		Fword:  __ccgo_ts + 30750,
		Ftype1: int8('F'),
	},
	4436: {
		Fword:  __ccgo_ts + 30756,
		Ftype1: int8('F'),
	},
	4437: {
		Fword:  __ccgo_ts + 30763,
		Ftype1: int8('F'),
	},
	4438: {
		Fword:  __ccgo_ts + 30770,
		Ftype1: int8('F'),
	},
	4439: {
		Fword:  __ccgo_ts + 30777,
		Ftype1: int8('F'),
	},
	4440: {
		Fword:  __ccgo_ts + 30784,
		Ftype1: int8('F'),
	},
	4441: {
		Fword:  __ccgo_ts + 30791,
		Ftype1: int8('F'),
	},
	4442: {
		Fword:  __ccgo_ts + 30798,
		Ftype1: int8('F'),
	},
	4443: {
		Fword:  __ccgo_ts + 30805,
		Ftype1: int8('F'),
	},
	4444: {
		Fword:  __ccgo_ts + 30812,
		Ftype1: int8('F'),
	},
	4445: {
		Fword:  __ccgo_ts + 30819,
		Ftype1: int8('F'),
	},
	4446: {
		Fword:  __ccgo_ts + 30826,
		Ftype1: int8('F'),
	},
	4447: {
		Fword:  __ccgo_ts + 30833,
		Ftype1: int8('F'),
	},
	4448: {
		Fword:  __ccgo_ts + 30840,
		Ftype1: int8('F'),
	},
	4449: {
		Fword:  __ccgo_ts + 30847,
		Ftype1: int8('F'),
	},
	4450: {
		Fword:  __ccgo_ts + 30854,
		Ftype1: int8('F'),
	},
	4451: {
		Fword:  __ccgo_ts + 30861,
		Ftype1: int8('F'),
	},
	4452: {
		Fword:  __ccgo_ts + 30868,
		Ftype1: int8('F'),
	},
	4453: {
		Fword:  __ccgo_ts + 30875,
		Ftype1: int8('F'),
	},
	4454: {
		Fword:  __ccgo_ts + 30882,
		Ftype1: int8('F'),
	},
	4455: {
		Fword:  __ccgo_ts + 30889,
		Ftype1: int8('F'),
	},
	4456: {
		Fword:  __ccgo_ts + 30896,
		Ftype1: int8('F'),
	},
	4457: {
		Fword:  __ccgo_ts + 30903,
		Ftype1: int8('F'),
	},
	4458: {
		Fword:  __ccgo_ts + 30910,
		Ftype1: int8('F'),
	},
	4459: {
		Fword:  __ccgo_ts + 30917,
		Ftype1: int8('F'),
	},
	4460: {
		Fword:  __ccgo_ts + 30924,
		Ftype1: int8('F'),
	},
	4461: {
		Fword:  __ccgo_ts + 30931,
		Ftype1: int8('F'),
	},
	4462: {
		Fword:  __ccgo_ts + 30938,
		Ftype1: int8('F'),
	},
	4463: {
		Fword:  __ccgo_ts + 30945,
		Ftype1: int8('F'),
	},
	4464: {
		Fword:  __ccgo_ts + 30951,
		Ftype1: int8('F'),
	},
	4465: {
		Fword:  __ccgo_ts + 30958,
		Ftype1: int8('F'),
	},
	4466: {
		Fword:  __ccgo_ts + 30965,
		Ftype1: int8('F'),
	},
	4467: {
		Fword:  __ccgo_ts + 30972,
		Ftype1: int8('F'),
	},
	4468: {
		Fword:  __ccgo_ts + 30979,
		Ftype1: int8('F'),
	},
	4469: {
		Fword:  __ccgo_ts + 30986,
		Ftype1: int8('F'),
	},
	4470: {
		Fword:  __ccgo_ts + 30993,
		Ftype1: int8('F'),
	},
	4471: {
		Fword:  __ccgo_ts + 31000,
		Ftype1: int8('F'),
	},
	4472: {
		Fword:  __ccgo_ts + 31007,
		Ftype1: int8('F'),
	},
	4473: {
		Fword:  __ccgo_ts + 31014,
		Ftype1: int8('F'),
	},
	4474: {
		Fword:  __ccgo_ts + 31021,
		Ftype1: int8('F'),
	},
	4475: {
		Fword:  __ccgo_ts + 31028,
		Ftype1: int8('F'),
	},
	4476: {
		Fword:  __ccgo_ts + 31035,
		Ftype1: int8('F'),
	},
	4477: {
		Fword:  __ccgo_ts + 31042,
		Ftype1: int8('F'),
	},
	4478: {
		Fword:  __ccgo_ts + 31049,
		Ftype1: int8('F'),
	},
	4479: {
		Fword:  __ccgo_ts + 31056,
		Ftype1: int8('F'),
	},
	4480: {
		Fword:  __ccgo_ts + 31063,
		Ftype1: int8('F'),
	},
	4481: {
		Fword:  __ccgo_ts + 31070,
		Ftype1: int8('F'),
	},
	4482: {
		Fword:  __ccgo_ts + 31077,
		Ftype1: int8('F'),
	},
	4483: {
		Fword:  __ccgo_ts + 31084,
		Ftype1: int8('F'),
	},
	4484: {
		Fword:  __ccgo_ts + 31091,
		Ftype1: int8('F'),
	},
	4485: {
		Fword:  __ccgo_ts + 31098,
		Ftype1: int8('F'),
	},
	4486: {
		Fword:  __ccgo_ts + 31105,
		Ftype1: int8('F'),
	},
	4487: {
		Fword:  __ccgo_ts + 31112,
		Ftype1: int8('F'),
	},
	4488: {
		Fword:  __ccgo_ts + 31119,
		Ftype1: int8('F'),
	},
	4489: {
		Fword:  __ccgo_ts + 31126,
		Ftype1: int8('F'),
	},
	4490: {
		Fword:  __ccgo_ts + 31133,
		Ftype1: int8('F'),
	},
	4491: {
		Fword:  __ccgo_ts + 31140,
		Ftype1: int8('F'),
	},
	4492: {
		Fword:  __ccgo_ts + 31147,
		Ftype1: int8('F'),
	},
	4493: {
		Fword:  __ccgo_ts + 31154,
		Ftype1: int8('F'),
	},
	4494: {
		Fword:  __ccgo_ts + 31161,
		Ftype1: int8('F'),
	},
	4495: {
		Fword:  __ccgo_ts + 31168,
		Ftype1: int8('F'),
	},
	4496: {
		Fword:  __ccgo_ts + 31175,
		Ftype1: int8('F'),
	},
	4497: {
		Fword:  __ccgo_ts + 31182,
		Ftype1: int8('F'),
	},
	4498: {
		Fword:  __ccgo_ts + 31189,
		Ftype1: int8('F'),
	},
	4499: {
		Fword:  __ccgo_ts + 31196,
		Ftype1: int8('F'),
	},
	4500: {
		Fword:  __ccgo_ts + 31203,
		Ftype1: int8('F'),
	},
	4501: {
		Fword:  __ccgo_ts + 31210,
		Ftype1: int8('F'),
	},
	4502: {
		Fword:  __ccgo_ts + 31217,
		Ftype1: int8('F'),
	},
	4503: {
		Fword:  __ccgo_ts + 31224,
		Ftype1: int8('F'),
	},
	4504: {
		Fword:  __ccgo_ts + 31231,
		Ftype1: int8('F'),
	},
	4505: {
		Fword:  __ccgo_ts + 31238,
		Ftype1: int8('F'),
	},
	4506: {
		Fword:  __ccgo_ts + 31245,
		Ftype1: int8('F'),
	},
	4507: {
		Fword:  __ccgo_ts + 31251,
		Ftype1: int8('F'),
	},
	4508: {
		Fword:  __ccgo_ts + 31258,
		Ftype1: int8('F'),
	},
	4509: {
		Fword:  __ccgo_ts + 31265,
		Ftype1: int8('F'),
	},
	4510: {
		Fword:  __ccgo_ts + 31272,
		Ftype1: int8('F'),
	},
	4511: {
		Fword:  __ccgo_ts + 31279,
		Ftype1: int8('F'),
	},
	4512: {
		Fword:  __ccgo_ts + 31286,
		Ftype1: int8('F'),
	},
	4513: {
		Fword:  __ccgo_ts + 31293,
		Ftype1: int8('F'),
	},
	4514: {
		Fword:  __ccgo_ts + 31300,
		Ftype1: int8('F'),
	},
	4515: {
		Fword:  __ccgo_ts + 31307,
		Ftype1: int8('F'),
	},
	4516: {
		Fword:  __ccgo_ts + 31314,
		Ftype1: int8('F'),
	},
	4517: {
		Fword:  __ccgo_ts + 31321,
		Ftype1: int8('F'),
	},
	4518: {
		Fword:  __ccgo_ts + 31328,
		Ftype1: int8('F'),
	},
	4519: {
		Fword:  __ccgo_ts + 31335,
		Ftype1: int8('F'),
	},
	4520: {
		Fword:  __ccgo_ts + 31342,
		Ftype1: int8('F'),
	},
	4521: {
		Fword:  __ccgo_ts + 31349,
		Ftype1: int8('F'),
	},
	4522: {
		Fword:  __ccgo_ts + 31356,
		Ftype1: int8('F'),
	},
	4523: {
		Fword:  __ccgo_ts + 31363,
		Ftype1: int8('F'),
	},
	4524: {
		Fword:  __ccgo_ts + 31369,
		Ftype1: int8('F'),
	},
	4525: {
		Fword:  __ccgo_ts + 31376,
		Ftype1: int8('F'),
	},
	4526: {
		Fword:  __ccgo_ts + 31383,
		Ftype1: int8('F'),
	},
	4527: {
		Fword:  __ccgo_ts + 31390,
		Ftype1: int8('F'),
	},
	4528: {
		Fword:  __ccgo_ts + 31397,
		Ftype1: int8('F'),
	},
	4529: {
		Fword:  __ccgo_ts + 31404,
		Ftype1: int8('F'),
	},
	4530: {
		Fword:  __ccgo_ts + 31411,
		Ftype1: int8('F'),
	},
	4531: {
		Fword:  __ccgo_ts + 31418,
		Ftype1: int8('F'),
	},
	4532: {
		Fword:  __ccgo_ts + 31425,
		Ftype1: int8('F'),
	},
	4533: {
		Fword:  __ccgo_ts + 31432,
		Ftype1: int8('F'),
	},
	4534: {
		Fword:  __ccgo_ts + 31439,
		Ftype1: int8('F'),
	},
	4535: {
		Fword:  __ccgo_ts + 31446,
		Ftype1: int8('F'),
	},
	4536: {
		Fword:  __ccgo_ts + 31453,
		Ftype1: int8('F'),
	},
	4537: {
		Fword:  __ccgo_ts + 31460,
		Ftype1: int8('F'),
	},
	4538: {
		Fword:  __ccgo_ts + 31467,
		Ftype1: int8('F'),
	},
	4539: {
		Fword:  __ccgo_ts + 31474,
		Ftype1: int8('F'),
	},
	4540: {
		Fword:  __ccgo_ts + 31481,
		Ftype1: int8('F'),
	},
	4541: {
		Fword:  __ccgo_ts + 31488,
		Ftype1: int8('F'),
	},
	4542: {
		Fword:  __ccgo_ts + 31495,
		Ftype1: int8('F'),
	},
	4543: {
		Fword:  __ccgo_ts + 31502,
		Ftype1: int8('F'),
	},
	4544: {
		Fword:  __ccgo_ts + 31509,
		Ftype1: int8('F'),
	},
	4545: {
		Fword:  __ccgo_ts + 31516,
		Ftype1: int8('F'),
	},
	4546: {
		Fword:  __ccgo_ts + 31523,
		Ftype1: int8('F'),
	},
	4547: {
		Fword:  __ccgo_ts + 31530,
		Ftype1: int8('F'),
	},
	4548: {
		Fword:  __ccgo_ts + 31537,
		Ftype1: int8('F'),
	},
	4549: {
		Fword:  __ccgo_ts + 31544,
		Ftype1: int8('F'),
	},
	4550: {
		Fword:  __ccgo_ts + 31551,
		Ftype1: int8('F'),
	},
	4551: {
		Fword:  __ccgo_ts + 31558,
		Ftype1: int8('F'),
	},
	4552: {
		Fword:  __ccgo_ts + 31565,
		Ftype1: int8('F'),
	},
	4553: {
		Fword:  __ccgo_ts + 31572,
		Ftype1: int8('F'),
	},
	4554: {
		Fword:  __ccgo_ts + 31579,
		Ftype1: int8('F'),
	},
	4555: {
		Fword:  __ccgo_ts + 31586,
		Ftype1: int8('F'),
	},
	4556: {
		Fword:  __ccgo_ts + 31593,
		Ftype1: int8('F'),
	},
	4557: {
		Fword:  __ccgo_ts + 31600,
		Ftype1: int8('F'),
	},
	4558: {
		Fword:  __ccgo_ts + 31607,
		Ftype1: int8('F'),
	},
	4559: {
		Fword:  __ccgo_ts + 31614,
		Ftype1: int8('F'),
	},
	4560: {
		Fword:  __ccgo_ts + 31621,
		Ftype1: int8('F'),
	},
	4561: {
		Fword:  __ccgo_ts + 31628,
		Ftype1: int8('F'),
	},
	4562: {
		Fword:  __ccgo_ts + 31635,
		Ftype1: int8('F'),
	},
	4563: {
		Fword:  __ccgo_ts + 31641,
		Ftype1: int8('F'),
	},
	4564: {
		Fword:  __ccgo_ts + 31648,
		Ftype1: int8('F'),
	},
	4565: {
		Fword:  __ccgo_ts + 31655,
		Ftype1: int8('F'),
	},
	4566: {
		Fword:  __ccgo_ts + 31662,
		Ftype1: int8('F'),
	},
	4567: {
		Fword:  __ccgo_ts + 31669,
		Ftype1: int8('F'),
	},
	4568: {
		Fword:  __ccgo_ts + 31676,
		Ftype1: int8('F'),
	},
	4569: {
		Fword:  __ccgo_ts + 31683,
		Ftype1: int8('F'),
	},
	4570: {
		Fword:  __ccgo_ts + 31690,
		Ftype1: int8('F'),
	},
	4571: {
		Fword:  __ccgo_ts + 31697,
		Ftype1: int8('F'),
	},
	4572: {
		Fword:  __ccgo_ts + 31704,
		Ftype1: int8('F'),
	},
	4573: {
		Fword:  __ccgo_ts + 31711,
		Ftype1: int8('F'),
	},
	4574: {
		Fword:  __ccgo_ts + 31718,
		Ftype1: int8('F'),
	},
	4575: {
		Fword:  __ccgo_ts + 31725,
		Ftype1: int8('F'),
	},
	4576: {
		Fword:  __ccgo_ts + 31732,
		Ftype1: int8('F'),
	},
	4577: {
		Fword:  __ccgo_ts + 31739,
		Ftype1: int8('F'),
	},
	4578: {
		Fword:  __ccgo_ts + 31746,
		Ftype1: int8('F'),
	},
	4579: {
		Fword:  __ccgo_ts + 31752,
		Ftype1: int8('F'),
	},
	4580: {
		Fword:  __ccgo_ts + 31759,
		Ftype1: int8('F'),
	},
	4581: {
		Fword:  __ccgo_ts + 31766,
		Ftype1: int8('F'),
	},
	4582: {
		Fword:  __ccgo_ts + 31773,
		Ftype1: int8('F'),
	},
	4583: {
		Fword:  __ccgo_ts + 31780,
		Ftype1: int8('F'),
	},
	4584: {
		Fword:  __ccgo_ts + 31787,
		Ftype1: int8('F'),
	},
	4585: {
		Fword:  __ccgo_ts + 31794,
		Ftype1: int8('F'),
	},
	4586: {
		Fword:  __ccgo_ts + 31801,
		Ftype1: int8('F'),
	},
	4587: {
		Fword:  __ccgo_ts + 31808,
		Ftype1: int8('F'),
	},
	4588: {
		Fword:  __ccgo_ts + 31815,
		Ftype1: int8('F'),
	},
	4589: {
		Fword:  __ccgo_ts + 31822,
		Ftype1: int8('F'),
	},
	4590: {
		Fword:  __ccgo_ts + 31829,
		Ftype1: int8('F'),
	},
	4591: {
		Fword:  __ccgo_ts + 31836,
		Ftype1: int8('F'),
	},
	4592: {
		Fword:  __ccgo_ts + 31843,
		Ftype1: int8('F'),
	},
	4593: {
		Fword:  __ccgo_ts + 31850,
		Ftype1: int8('F'),
	},
	4594: {
		Fword:  __ccgo_ts + 31857,
		Ftype1: int8('F'),
	},
	4595: {
		Fword:  __ccgo_ts + 31864,
		Ftype1: int8('F'),
	},
	4596: {
		Fword:  __ccgo_ts + 31871,
		Ftype1: int8('F'),
	},
	4597: {
		Fword:  __ccgo_ts + 31878,
		Ftype1: int8('F'),
	},
	4598: {
		Fword:  __ccgo_ts + 31885,
		Ftype1: int8('F'),
	},
	4599: {
		Fword:  __ccgo_ts + 31892,
		Ftype1: int8('F'),
	},
	4600: {
		Fword:  __ccgo_ts + 31899,
		Ftype1: int8('F'),
	},
	4601: {
		Fword:  __ccgo_ts + 31906,
		Ftype1: int8('F'),
	},
	4602: {
		Fword:  __ccgo_ts + 31913,
		Ftype1: int8('F'),
	},
	4603: {
		Fword:  __ccgo_ts + 31920,
		Ftype1: int8('F'),
	},
	4604: {
		Fword:  __ccgo_ts + 31927,
		Ftype1: int8('F'),
	},
	4605: {
		Fword:  __ccgo_ts + 31934,
		Ftype1: int8('F'),
	},
	4606: {
		Fword:  __ccgo_ts + 31941,
		Ftype1: int8('F'),
	},
	4607: {
		Fword:  __ccgo_ts + 31948,
		Ftype1: int8('F'),
	},
	4608: {
		Fword:  __ccgo_ts + 31955,
		Ftype1: int8('F'),
	},
	4609: {
		Fword:  __ccgo_ts + 31962,
		Ftype1: int8('F'),
	},
	4610: {
		Fword:  __ccgo_ts + 31969,
		Ftype1: int8('F'),
	},
	4611: {
		Fword:  __ccgo_ts + 31976,
		Ftype1: int8('F'),
	},
	4612: {
		Fword:  __ccgo_ts + 31983,
		Ftype1: int8('F'),
	},
	4613: {
		Fword:  __ccgo_ts + 31990,
		Ftype1: int8('F'),
	},
	4614: {
		Fword:  __ccgo_ts + 31997,
		Ftype1: int8('F'),
	},
	4615: {
		Fword:  __ccgo_ts + 32004,
		Ftype1: int8('F'),
	},
	4616: {
		Fword:  __ccgo_ts + 32011,
		Ftype1: int8('F'),
	},
	4617: {
		Fword:  __ccgo_ts + 32018,
		Ftype1: int8('F'),
	},
	4618: {
		Fword:  __ccgo_ts + 32025,
		Ftype1: int8('F'),
	},
	4619: {
		Fword:  __ccgo_ts + 32032,
		Ftype1: int8('F'),
	},
	4620: {
		Fword:  __ccgo_ts + 32039,
		Ftype1: int8('F'),
	},
	4621: {
		Fword:  __ccgo_ts + 32046,
		Ftype1: int8('F'),
	},
	4622: {
		Fword:  __ccgo_ts + 32053,
		Ftype1: int8('F'),
	},
	4623: {
		Fword:  __ccgo_ts + 32060,
		Ftype1: int8('F'),
	},
	4624: {
		Fword:  __ccgo_ts + 32067,
		Ftype1: int8('F'),
	},
	4625: {
		Fword:  __ccgo_ts + 32074,
		Ftype1: int8('F'),
	},
	4626: {
		Fword:  __ccgo_ts + 32081,
		Ftype1: int8('F'),
	},
	4627: {
		Fword:  __ccgo_ts + 32088,
		Ftype1: int8('F'),
	},
	4628: {
		Fword:  __ccgo_ts + 32095,
		Ftype1: int8('F'),
	},
	4629: {
		Fword:  __ccgo_ts + 32102,
		Ftype1: int8('F'),
	},
	4630: {
		Fword:  __ccgo_ts + 32109,
		Ftype1: int8('F'),
	},
	4631: {
		Fword:  __ccgo_ts + 32116,
		Ftype1: int8('F'),
	},
	4632: {
		Fword:  __ccgo_ts + 32123,
		Ftype1: int8('F'),
	},
	4633: {
		Fword:  __ccgo_ts + 32130,
		Ftype1: int8('F'),
	},
	4634: {
		Fword:  __ccgo_ts + 32137,
		Ftype1: int8('F'),
	},
	4635: {
		Fword:  __ccgo_ts + 32144,
		Ftype1: int8('F'),
	},
	4636: {
		Fword:  __ccgo_ts + 32151,
		Ftype1: int8('F'),
	},
	4637: {
		Fword:  __ccgo_ts + 32158,
		Ftype1: int8('F'),
	},
	4638: {
		Fword:  __ccgo_ts + 32165,
		Ftype1: int8('F'),
	},
	4639: {
		Fword:  __ccgo_ts + 32172,
		Ftype1: int8('F'),
	},
	4640: {
		Fword:  __ccgo_ts + 32179,
		Ftype1: int8('F'),
	},
	4641: {
		Fword:  __ccgo_ts + 32186,
		Ftype1: int8('F'),
	},
	4642: {
		Fword:  __ccgo_ts + 32193,
		Ftype1: int8('F'),
	},
	4643: {
		Fword:  __ccgo_ts + 32200,
		Ftype1: int8('F'),
	},
	4644: {
		Fword:  __ccgo_ts + 32207,
		Ftype1: int8('F'),
	},
	4645: {
		Fword:  __ccgo_ts + 32214,
		Ftype1: int8('F'),
	},
	4646: {
		Fword:  __ccgo_ts + 32221,
		Ftype1: int8('F'),
	},
	4647: {
		Fword:  __ccgo_ts + 32228,
		Ftype1: int8('F'),
	},
	4648: {
		Fword:  __ccgo_ts + 32235,
		Ftype1: int8('F'),
	},
	4649: {
		Fword:  __ccgo_ts + 32242,
		Ftype1: int8('F'),
	},
	4650: {
		Fword:  __ccgo_ts + 32249,
		Ftype1: int8('F'),
	},
	4651: {
		Fword:  __ccgo_ts + 32256,
		Ftype1: int8('F'),
	},
	4652: {
		Fword:  __ccgo_ts + 32263,
		Ftype1: int8('F'),
	},
	4653: {
		Fword:  __ccgo_ts + 32270,
		Ftype1: int8('F'),
	},
	4654: {
		Fword:  __ccgo_ts + 32277,
		Ftype1: int8('F'),
	},
	4655: {
		Fword:  __ccgo_ts + 32284,
		Ftype1: int8('F'),
	},
	4656: {
		Fword:  __ccgo_ts + 32291,
		Ftype1: int8('F'),
	},
	4657: {
		Fword:  __ccgo_ts + 32298,
		Ftype1: int8('F'),
	},
	4658: {
		Fword:  __ccgo_ts + 32305,
		Ftype1: int8('F'),
	},
	4659: {
		Fword:  __ccgo_ts + 32312,
		Ftype1: int8('F'),
	},
	4660: {
		Fword:  __ccgo_ts + 32319,
		Ftype1: int8('F'),
	},
	4661: {
		Fword:  __ccgo_ts + 32326,
		Ftype1: int8('F'),
	},
	4662: {
		Fword:  __ccgo_ts + 32333,
		Ftype1: int8('F'),
	},
	4663: {
		Fword:  __ccgo_ts + 32340,
		Ftype1: int8('F'),
	},
	4664: {
		Fword:  __ccgo_ts + 32347,
		Ftype1: int8('F'),
	},
	4665: {
		Fword:  __ccgo_ts + 32354,
		Ftype1: int8('F'),
	},
	4666: {
		Fword:  __ccgo_ts + 32361,
		Ftype1: int8('F'),
	},
	4667: {
		Fword:  __ccgo_ts + 32368,
		Ftype1: int8('F'),
	},
	4668: {
		Fword:  __ccgo_ts + 32375,
		Ftype1: int8('F'),
	},
	4669: {
		Fword:  __ccgo_ts + 32382,
		Ftype1: int8('F'),
	},
	4670: {
		Fword:  __ccgo_ts + 32389,
		Ftype1: int8('F'),
	},
	4671: {
		Fword:  __ccgo_ts + 32396,
		Ftype1: int8('F'),
	},
	4672: {
		Fword:  __ccgo_ts + 32403,
		Ftype1: int8('F'),
	},
	4673: {
		Fword:  __ccgo_ts + 32410,
		Ftype1: int8('F'),
	},
	4674: {
		Fword:  __ccgo_ts + 32417,
		Ftype1: int8('F'),
	},
	4675: {
		Fword:  __ccgo_ts + 32424,
		Ftype1: int8('F'),
	},
	4676: {
		Fword:  __ccgo_ts + 32431,
		Ftype1: int8('F'),
	},
	4677: {
		Fword:  __ccgo_ts + 32438,
		Ftype1: int8('F'),
	},
	4678: {
		Fword:  __ccgo_ts + 32445,
		Ftype1: int8('F'),
	},
	4679: {
		Fword:  __ccgo_ts + 32452,
		Ftype1: int8('F'),
	},
	4680: {
		Fword:  __ccgo_ts + 32459,
		Ftype1: int8('F'),
	},
	4681: {
		Fword:  __ccgo_ts + 32466,
		Ftype1: int8('F'),
	},
	4682: {
		Fword:  __ccgo_ts + 32473,
		Ftype1: int8('F'),
	},
	4683: {
		Fword:  __ccgo_ts + 32480,
		Ftype1: int8('F'),
	},
	4684: {
		Fword:  __ccgo_ts + 32487,
		Ftype1: int8('F'),
	},
	4685: {
		Fword:  __ccgo_ts + 32494,
		Ftype1: int8('F'),
	},
	4686: {
		Fword:  __ccgo_ts + 32501,
		Ftype1: int8('F'),
	},
	4687: {
		Fword:  __ccgo_ts + 32508,
		Ftype1: int8('F'),
	},
	4688: {
		Fword:  __ccgo_ts + 32515,
		Ftype1: int8('F'),
	},
	4689: {
		Fword:  __ccgo_ts + 32522,
		Ftype1: int8('F'),
	},
	4690: {
		Fword:  __ccgo_ts + 32529,
		Ftype1: int8('F'),
	},
	4691: {
		Fword:  __ccgo_ts + 32535,
		Ftype1: int8('F'),
	},
	4692: {
		Fword:  __ccgo_ts + 32542,
		Ftype1: int8('F'),
	},
	4693: {
		Fword:  __ccgo_ts + 32549,
		Ftype1: int8('F'),
	},
	4694: {
		Fword:  __ccgo_ts + 32556,
		Ftype1: int8('F'),
	},
	4695: {
		Fword:  __ccgo_ts + 32563,
		Ftype1: int8('F'),
	},
	4696: {
		Fword:  __ccgo_ts + 32568,
		Ftype1: int8('F'),
	},
	4697: {
		Fword:  __ccgo_ts + 32574,
		Ftype1: int8('F'),
	},
	4698: {
		Fword:  __ccgo_ts + 32579,
		Ftype1: int8('F'),
	},
	4699: {
		Fword:  __ccgo_ts + 32584,
		Ftype1: int8('F'),
	},
	4700: {
		Fword:  __ccgo_ts + 32591,
		Ftype1: int8('F'),
	},
	4701: {
		Fword:  __ccgo_ts + 32598,
		Ftype1: int8('F'),
	},
	4702: {
		Fword:  __ccgo_ts + 32605,
		Ftype1: int8('F'),
	},
	4703: {
		Fword:  __ccgo_ts + 32612,
		Ftype1: int8('F'),
	},
	4704: {
		Fword:  __ccgo_ts + 32619,
		Ftype1: int8('F'),
	},
	4705: {
		Fword:  __ccgo_ts + 32626,
		Ftype1: int8('F'),
	},
	4706: {
		Fword:  __ccgo_ts + 32633,
		Ftype1: int8('F'),
	},
	4707: {
		Fword:  __ccgo_ts + 32639,
		Ftype1: int8('F'),
	},
	4708: {
		Fword:  __ccgo_ts + 32646,
		Ftype1: int8('F'),
	},
	4709: {
		Fword:  __ccgo_ts + 32653,
		Ftype1: int8('F'),
	},
	4710: {
		Fword:  __ccgo_ts + 32660,
		Ftype1: int8('F'),
	},
	4711: {
		Fword:  __ccgo_ts + 32667,
		Ftype1: int8('F'),
	},
	4712: {
		Fword:  __ccgo_ts + 32674,
		Ftype1: int8('F'),
	},
	4713: {
		Fword:  __ccgo_ts + 32681,
		Ftype1: int8('F'),
	},
	4714: {
		Fword:  __ccgo_ts + 32688,
		Ftype1: int8('F'),
	},
	4715: {
		Fword:  __ccgo_ts + 32695,
		Ftype1: int8('F'),
	},
	4716: {
		Fword:  __ccgo_ts + 32702,
		Ftype1: int8('F'),
	},
	4717: {
		Fword:  __ccgo_ts + 32709,
		Ftype1: int8('F'),
	},
	4718: {
		Fword:  __ccgo_ts + 32716,
		Ftype1: int8('F'),
	},
	4719: {
		Fword:  __ccgo_ts + 32723,
		Ftype1: int8('F'),
	},
	4720: {
		Fword:  __ccgo_ts + 32730,
		Ftype1: int8('F'),
	},
	4721: {
		Fword:  __ccgo_ts + 32737,
		Ftype1: int8('F'),
	},
	4722: {
		Fword:  __ccgo_ts + 32743,
		Ftype1: int8('F'),
	},
	4723: {
		Fword:  __ccgo_ts + 32750,
		Ftype1: int8('F'),
	},
	4724: {
		Fword:  __ccgo_ts + 32756,
		Ftype1: int8('F'),
	},
	4725: {
		Fword:  __ccgo_ts + 32762,
		Ftype1: int8('F'),
	},
	4726: {
		Fword:  __ccgo_ts + 32769,
		Ftype1: int8('F'),
	},
	4727: {
		Fword:  __ccgo_ts + 32776,
		Ftype1: int8('F'),
	},
	4728: {
		Fword:  __ccgo_ts + 32783,
		Ftype1: int8('F'),
	},
	4729: {
		Fword:  __ccgo_ts + 32790,
		Ftype1: int8('F'),
	},
	4730: {
		Fword:  __ccgo_ts + 32796,
		Ftype1: int8('F'),
	},
	4731: {
		Fword:  __ccgo_ts + 32803,
		Ftype1: int8('F'),
	},
	4732: {
		Fword:  __ccgo_ts + 32810,
		Ftype1: int8('F'),
	},
	4733: {
		Fword:  __ccgo_ts + 32817,
		Ftype1: int8('F'),
	},
	4734: {
		Fword:  __ccgo_ts + 32824,
		Ftype1: int8('F'),
	},
	4735: {
		Fword:  __ccgo_ts + 32831,
		Ftype1: int8('F'),
	},
	4736: {
		Fword:  __ccgo_ts + 32838,
		Ftype1: int8('F'),
	},
	4737: {
		Fword:  __ccgo_ts + 32845,
		Ftype1: int8('F'),
	},
	4738: {
		Fword:  __ccgo_ts + 32852,
		Ftype1: int8('F'),
	},
	4739: {
		Fword:  __ccgo_ts + 32859,
		Ftype1: int8('F'),
	},
	4740: {
		Fword:  __ccgo_ts + 32865,
		Ftype1: int8('F'),
	},
	4741: {
		Fword:  __ccgo_ts + 32872,
		Ftype1: int8('F'),
	},
	4742: {
		Fword:  __ccgo_ts + 32879,
		Ftype1: int8('F'),
	},
	4743: {
		Fword:  __ccgo_ts + 32886,
		Ftype1: int8('F'),
	},
	4744: {
		Fword:  __ccgo_ts + 32893,
		Ftype1: int8('F'),
	},
	4745: {
		Fword:  __ccgo_ts + 32900,
		Ftype1: int8('F'),
	},
	4746: {
		Fword:  __ccgo_ts + 32907,
		Ftype1: int8('F'),
	},
	4747: {
		Fword:  __ccgo_ts + 32914,
		Ftype1: int8('F'),
	},
	4748: {
		Fword:  __ccgo_ts + 32921,
		Ftype1: int8('F'),
	},
	4749: {
		Fword:  __ccgo_ts + 32928,
		Ftype1: int8('F'),
	},
	4750: {
		Fword:  __ccgo_ts + 32935,
		Ftype1: int8('F'),
	},
	4751: {
		Fword:  __ccgo_ts + 32942,
		Ftype1: int8('F'),
	},
	4752: {
		Fword:  __ccgo_ts + 32949,
		Ftype1: int8('F'),
	},
	4753: {
		Fword:  __ccgo_ts + 32956,
		Ftype1: int8('F'),
	},
	4754: {
		Fword:  __ccgo_ts + 32963,
		Ftype1: int8('F'),
	},
	4755: {
		Fword:  __ccgo_ts + 32970,
		Ftype1: int8('F'),
	},
	4756: {
		Fword:  __ccgo_ts + 32976,
		Ftype1: int8('F'),
	},
	4757: {
		Fword:  __ccgo_ts + 32983,
		Ftype1: int8('F'),
	},
	4758: {
		Fword:  __ccgo_ts + 32990,
		Ftype1: int8('F'),
	},
	4759: {
		Fword:  __ccgo_ts + 32997,
		Ftype1: int8('F'),
	},
	4760: {
		Fword:  __ccgo_ts + 33004,
		Ftype1: int8('F'),
	},
	4761: {
		Fword:  __ccgo_ts + 33011,
		Ftype1: int8('F'),
	},
	4762: {
		Fword:  __ccgo_ts + 33018,
		Ftype1: int8('F'),
	},
	4763: {
		Fword:  __ccgo_ts + 33025,
		Ftype1: int8('F'),
	},
	4764: {
		Fword:  __ccgo_ts + 33032,
		Ftype1: int8('F'),
	},
	4765: {
		Fword:  __ccgo_ts + 33039,
		Ftype1: int8('F'),
	},
	4766: {
		Fword:  __ccgo_ts + 33046,
		Ftype1: int8('F'),
	},
	4767: {
		Fword:  __ccgo_ts + 33053,
		Ftype1: int8('F'),
	},
	4768: {
		Fword:  __ccgo_ts + 33060,
		Ftype1: int8('F'),
	},
	4769: {
		Fword:  __ccgo_ts + 33067,
		Ftype1: int8('F'),
	},
	4770: {
		Fword:  __ccgo_ts + 33073,
		Ftype1: int8('F'),
	},
	4771: {
		Fword:  __ccgo_ts + 33080,
		Ftype1: int8('F'),
	},
	4772: {
		Fword:  __ccgo_ts + 33087,
		Ftype1: int8('F'),
	},
	4773: {
		Fword:  __ccgo_ts + 33094,
		Ftype1: int8('F'),
	},
	4774: {
		Fword:  __ccgo_ts + 33101,
		Ftype1: int8('F'),
	},
	4775: {
		Fword:  __ccgo_ts + 33108,
		Ftype1: int8('F'),
	},
	4776: {
		Fword:  __ccgo_ts + 33115,
		Ftype1: int8('F'),
	},
	4777: {
		Fword:  __ccgo_ts + 33122,
		Ftype1: int8('F'),
	},
	4778: {
		Fword:  __ccgo_ts + 33129,
		Ftype1: int8('F'),
	},
	4779: {
		Fword:  __ccgo_ts + 33136,
		Ftype1: int8('F'),
	},
	4780: {
		Fword:  __ccgo_ts + 33143,
		Ftype1: int8('F'),
	},
	4781: {
		Fword:  __ccgo_ts + 33150,
		Ftype1: int8('F'),
	},
	4782: {
		Fword:  __ccgo_ts + 33157,
		Ftype1: int8('F'),
	},
	4783: {
		Fword:  __ccgo_ts + 33164,
		Ftype1: int8('F'),
	},
	4784: {
		Fword:  __ccgo_ts + 33171,
		Ftype1: int8('F'),
	},
	4785: {
		Fword:  __ccgo_ts + 33178,
		Ftype1: int8('F'),
	},
	4786: {
		Fword:  __ccgo_ts + 33185,
		Ftype1: int8('F'),
	},
	4787: {
		Fword:  __ccgo_ts + 33192,
		Ftype1: int8('F'),
	},
	4788: {
		Fword:  __ccgo_ts + 33199,
		Ftype1: int8('F'),
	},
	4789: {
		Fword:  __ccgo_ts + 33206,
		Ftype1: int8('F'),
	},
	4790: {
		Fword:  __ccgo_ts + 33213,
		Ftype1: int8('F'),
	},
	4791: {
		Fword:  __ccgo_ts + 33220,
		Ftype1: int8('F'),
	},
	4792: {
		Fword:  __ccgo_ts + 33227,
		Ftype1: int8('F'),
	},
	4793: {
		Fword:  __ccgo_ts + 33234,
		Ftype1: int8('F'),
	},
	4794: {
		Fword:  __ccgo_ts + 33241,
		Ftype1: int8('F'),
	},
	4795: {
		Fword:  __ccgo_ts + 33248,
		Ftype1: int8('F'),
	},
	4796: {
		Fword:  __ccgo_ts + 33255,
		Ftype1: int8('F'),
	},
	4797: {
		Fword:  __ccgo_ts + 33262,
		Ftype1: int8('F'),
	},
	4798: {
		Fword:  __ccgo_ts + 33268,
		Ftype1: int8('F'),
	},
	4799: {
		Fword:  __ccgo_ts + 33275,
		Ftype1: int8('F'),
	},
	4800: {
		Fword:  __ccgo_ts + 33282,
		Ftype1: int8('F'),
	},
	4801: {
		Fword:  __ccgo_ts + 33289,
		Ftype1: int8('F'),
	},
	4802: {
		Fword:  __ccgo_ts + 33296,
		Ftype1: int8('F'),
	},
	4803: {
		Fword:  __ccgo_ts + 33303,
		Ftype1: int8('F'),
	},
	4804: {
		Fword:  __ccgo_ts + 33310,
		Ftype1: int8('F'),
	},
	4805: {
		Fword:  __ccgo_ts + 33317,
		Ftype1: int8('F'),
	},
	4806: {
		Fword:  __ccgo_ts + 33324,
		Ftype1: int8('F'),
	},
	4807: {
		Fword:  __ccgo_ts + 33331,
		Ftype1: int8('F'),
	},
	4808: {
		Fword:  __ccgo_ts + 33338,
		Ftype1: int8('F'),
	},
	4809: {
		Fword:  __ccgo_ts + 33345,
		Ftype1: int8('F'),
	},
	4810: {
		Fword:  __ccgo_ts + 33352,
		Ftype1: int8('F'),
	},
	4811: {
		Fword:  __ccgo_ts + 33359,
		Ftype1: int8('F'),
	},
	4812: {
		Fword:  __ccgo_ts + 33366,
		Ftype1: int8('F'),
	},
	4813: {
		Fword:  __ccgo_ts + 33372,
		Ftype1: int8('F'),
	},
	4814: {
		Fword:  __ccgo_ts + 33379,
		Ftype1: int8('F'),
	},
	4815: {
		Fword:  __ccgo_ts + 33386,
		Ftype1: int8('F'),
	},
	4816: {
		Fword:  __ccgo_ts + 33393,
		Ftype1: int8('F'),
	},
	4817: {
		Fword:  __ccgo_ts + 33400,
		Ftype1: int8('F'),
	},
	4818: {
		Fword:  __ccgo_ts + 33407,
		Ftype1: int8('F'),
	},
	4819: {
		Fword:  __ccgo_ts + 33414,
		Ftype1: int8('F'),
	},
	4820: {
		Fword:  __ccgo_ts + 33421,
		Ftype1: int8('F'),
	},
	4821: {
		Fword:  __ccgo_ts + 33428,
		Ftype1: int8('F'),
	},
	4822: {
		Fword:  __ccgo_ts + 33435,
		Ftype1: int8('F'),
	},
	4823: {
		Fword:  __ccgo_ts + 33442,
		Ftype1: int8('F'),
	},
	4824: {
		Fword:  __ccgo_ts + 33449,
		Ftype1: int8('F'),
	},
	4825: {
		Fword:  __ccgo_ts + 33456,
		Ftype1: int8('F'),
	},
	4826: {
		Fword:  __ccgo_ts + 33463,
		Ftype1: int8('F'),
	},
	4827: {
		Fword:  __ccgo_ts + 33470,
		Ftype1: int8('F'),
	},
	4828: {
		Fword:  __ccgo_ts + 33477,
		Ftype1: int8('F'),
	},
	4829: {
		Fword:  __ccgo_ts + 33484,
		Ftype1: int8('F'),
	},
	4830: {
		Fword:  __ccgo_ts + 33491,
		Ftype1: int8('F'),
	},
	4831: {
		Fword:  __ccgo_ts + 33498,
		Ftype1: int8('F'),
	},
	4832: {
		Fword:  __ccgo_ts + 33505,
		Ftype1: int8('F'),
	},
	4833: {
		Fword:  __ccgo_ts + 33512,
		Ftype1: int8('F'),
	},
	4834: {
		Fword:  __ccgo_ts + 33519,
		Ftype1: int8('F'),
	},
	4835: {
		Fword:  __ccgo_ts + 33526,
		Ftype1: int8('F'),
	},
	4836: {
		Fword:  __ccgo_ts + 33533,
		Ftype1: int8('F'),
	},
	4837: {
		Fword:  __ccgo_ts + 33540,
		Ftype1: int8('F'),
	},
	4838: {
		Fword:  __ccgo_ts + 33547,
		Ftype1: int8('F'),
	},
	4839: {
		Fword:  __ccgo_ts + 33554,
		Ftype1: int8('F'),
	},
	4840: {
		Fword:  __ccgo_ts + 33561,
		Ftype1: int8('F'),
	},
	4841: {
		Fword:  __ccgo_ts + 33568,
		Ftype1: int8('F'),
	},
	4842: {
		Fword:  __ccgo_ts + 33573,
		Ftype1: int8('F'),
	},
	4843: {
		Fword:  __ccgo_ts + 33580,
		Ftype1: int8('F'),
	},
	4844: {
		Fword:  __ccgo_ts + 33587,
		Ftype1: int8('F'),
	},
	4845: {
		Fword:  __ccgo_ts + 33594,
		Ftype1: int8('F'),
	},
	4846: {
		Fword:  __ccgo_ts + 33601,
		Ftype1: int8('F'),
	},
	4847: {
		Fword:  __ccgo_ts + 33608,
		Ftype1: int8('F'),
	},
	4848: {
		Fword:  __ccgo_ts + 33615,
		Ftype1: int8('F'),
	},
	4849: {
		Fword:  __ccgo_ts + 33622,
		Ftype1: int8('F'),
	},
	4850: {
		Fword:  __ccgo_ts + 33629,
		Ftype1: int8('F'),
	},
	4851: {
		Fword:  __ccgo_ts + 33636,
		Ftype1: int8('F'),
	},
	4852: {
		Fword:  __ccgo_ts + 33643,
		Ftype1: int8('F'),
	},
	4853: {
		Fword:  __ccgo_ts + 33649,
		Ftype1: int8('F'),
	},
	4854: {
		Fword:  __ccgo_ts + 33656,
		Ftype1: int8('F'),
	},
	4855: {
		Fword:  __ccgo_ts + 33663,
		Ftype1: int8('F'),
	},
	4856: {
		Fword:  __ccgo_ts + 33670,
		Ftype1: int8('F'),
	},
	4857: {
		Fword:  __ccgo_ts + 33677,
		Ftype1: int8('F'),
	},
	4858: {
		Fword:  __ccgo_ts + 33684,
		Ftype1: int8('F'),
	},
	4859: {
		Fword:  __ccgo_ts + 33691,
		Ftype1: int8('F'),
	},
	4860: {
		Fword:  __ccgo_ts + 33698,
		Ftype1: int8('F'),
	},
	4861: {
		Fword:  __ccgo_ts + 33705,
		Ftype1: int8('F'),
	},
	4862: {
		Fword:  __ccgo_ts + 33712,
		Ftype1: int8('F'),
	},
	4863: {
		Fword:  __ccgo_ts + 33718,
		Ftype1: int8('F'),
	},
	4864: {
		Fword:  __ccgo_ts + 33725,
		Ftype1: int8('F'),
	},
	4865: {
		Fword:  __ccgo_ts + 33732,
		Ftype1: int8('F'),
	},
	4866: {
		Fword:  __ccgo_ts + 33739,
		Ftype1: int8('F'),
	},
	4867: {
		Fword:  __ccgo_ts + 33746,
		Ftype1: int8('F'),
	},
	4868: {
		Fword:  __ccgo_ts + 33753,
		Ftype1: int8('F'),
	},
	4869: {
		Fword:  __ccgo_ts + 33760,
		Ftype1: int8('F'),
	},
	4870: {
		Fword:  __ccgo_ts + 33767,
		Ftype1: int8('F'),
	},
	4871: {
		Fword:  __ccgo_ts + 33774,
		Ftype1: int8('F'),
	},
	4872: {
		Fword:  __ccgo_ts + 33781,
		Ftype1: int8('F'),
	},
	4873: {
		Fword:  __ccgo_ts + 33788,
		Ftype1: int8('F'),
	},
	4874: {
		Fword:  __ccgo_ts + 33795,
		Ftype1: int8('F'),
	},
	4875: {
		Fword:  __ccgo_ts + 33802,
		Ftype1: int8('F'),
	},
	4876: {
		Fword:  __ccgo_ts + 33809,
		Ftype1: int8('F'),
	},
	4877: {
		Fword:  __ccgo_ts + 33816,
		Ftype1: int8('F'),
	},
	4878: {
		Fword:  __ccgo_ts + 33822,
		Ftype1: int8('F'),
	},
	4879: {
		Fword:  __ccgo_ts + 33829,
		Ftype1: int8('F'),
	},
	4880: {
		Fword:  __ccgo_ts + 33836,
		Ftype1: int8('F'),
	},
	4881: {
		Fword:  __ccgo_ts + 33843,
		Ftype1: int8('F'),
	},
	4882: {
		Fword:  __ccgo_ts + 33850,
		Ftype1: int8('F'),
	},
	4883: {
		Fword:  __ccgo_ts + 33857,
		Ftype1: int8('F'),
	},
	4884: {
		Fword:  __ccgo_ts + 33864,
		Ftype1: int8('F'),
	},
	4885: {
		Fword:  __ccgo_ts + 33871,
		Ftype1: int8('F'),
	},
	4886: {
		Fword:  __ccgo_ts + 33878,
		Ftype1: int8('F'),
	},
	4887: {
		Fword:  __ccgo_ts + 33885,
		Ftype1: int8('F'),
	},
	4888: {
		Fword:  __ccgo_ts + 33892,
		Ftype1: int8('F'),
	},
	4889: {
		Fword:  __ccgo_ts + 33898,
		Ftype1: int8('F'),
	},
	4890: {
		Fword:  __ccgo_ts + 33905,
		Ftype1: int8('F'),
	},
	4891: {
		Fword:  __ccgo_ts + 33912,
		Ftype1: int8('F'),
	},
	4892: {
		Fword:  __ccgo_ts + 33919,
		Ftype1: int8('F'),
	},
	4893: {
		Fword:  __ccgo_ts + 33926,
		Ftype1: int8('F'),
	},
	4894: {
		Fword:  __ccgo_ts + 33933,
		Ftype1: int8('F'),
	},
	4895: {
		Fword:  __ccgo_ts + 33940,
		Ftype1: int8('F'),
	},
	4896: {
		Fword:  __ccgo_ts + 33947,
		Ftype1: int8('F'),
	},
	4897: {
		Fword:  __ccgo_ts + 33954,
		Ftype1: int8('F'),
	},
	4898: {
		Fword:  __ccgo_ts + 33961,
		Ftype1: int8('F'),
	},
	4899: {
		Fword:  __ccgo_ts + 33968,
		Ftype1: int8('F'),
	},
	4900: {
		Fword:  __ccgo_ts + 33975,
		Ftype1: int8('F'),
	},
	4901: {
		Fword:  __ccgo_ts + 33981,
		Ftype1: int8('F'),
	},
	4902: {
		Fword:  __ccgo_ts + 33988,
		Ftype1: int8('F'),
	},
	4903: {
		Fword:  __ccgo_ts + 33995,
		Ftype1: int8('F'),
	},
	4904: {
		Fword:  __ccgo_ts + 34002,
		Ftype1: int8('F'),
	},
	4905: {
		Fword:  __ccgo_ts + 34009,
		Ftype1: int8('F'),
	},
	4906: {
		Fword:  __ccgo_ts + 34015,
		Ftype1: int8('F'),
	},
	4907: {
		Fword:  __ccgo_ts + 34022,
		Ftype1: int8('F'),
	},
	4908: {
		Fword:  __ccgo_ts + 34029,
		Ftype1: int8('F'),
	},
	4909: {
		Fword:  __ccgo_ts + 34036,
		Ftype1: int8('F'),
	},
	4910: {
		Fword:  __ccgo_ts + 34043,
		Ftype1: int8('F'),
	},
	4911: {
		Fword:  __ccgo_ts + 34049,
		Ftype1: int8('F'),
	},
	4912: {
		Fword:  __ccgo_ts + 34056,
		Ftype1: int8('F'),
	},
	4913: {
		Fword:  __ccgo_ts + 34063,
		Ftype1: int8('F'),
	},
	4914: {
		Fword:  __ccgo_ts + 34070,
		Ftype1: int8('F'),
	},
	4915: {
		Fword:  __ccgo_ts + 34077,
		Ftype1: int8('F'),
	},
	4916: {
		Fword:  __ccgo_ts + 34084,
		Ftype1: int8('F'),
	},
	4917: {
		Fword:  __ccgo_ts + 34091,
		Ftype1: int8('F'),
	},
	4918: {
		Fword:  __ccgo_ts + 34098,
		Ftype1: int8('F'),
	},
	4919: {
		Fword:  __ccgo_ts + 34105,
		Ftype1: int8('F'),
	},
	4920: {
		Fword:  __ccgo_ts + 34112,
		Ftype1: int8('F'),
	},
	4921: {
		Fword:  __ccgo_ts + 34119,
		Ftype1: int8('F'),
	},
	4922: {
		Fword:  __ccgo_ts + 34126,
		Ftype1: int8('F'),
	},
	4923: {
		Fword:  __ccgo_ts + 34133,
		Ftype1: int8('F'),
	},
	4924: {
		Fword:  __ccgo_ts + 34140,
		Ftype1: int8('F'),
	},
	4925: {
		Fword:  __ccgo_ts + 34147,
		Ftype1: int8('F'),
	},
	4926: {
		Fword:  __ccgo_ts + 34154,
		Ftype1: int8('F'),
	},
	4927: {
		Fword:  __ccgo_ts + 34161,
		Ftype1: int8('F'),
	},
	4928: {
		Fword:  __ccgo_ts + 34168,
		Ftype1: int8('F'),
	},
	4929: {
		Fword:  __ccgo_ts + 34175,
		Ftype1: int8('F'),
	},
	4930: {
		Fword:  __ccgo_ts + 34182,
		Ftype1: int8('F'),
	},
	4931: {
		Fword:  __ccgo_ts + 34189,
		Ftype1: int8('F'),
	},
	4932: {
		Fword:  __ccgo_ts + 34196,
		Ftype1: int8('F'),
	},
	4933: {
		Fword:  __ccgo_ts + 34203,
		Ftype1: int8('F'),
	},
	4934: {
		Fword:  __ccgo_ts + 34210,
		Ftype1: int8('F'),
	},
	4935: {
		Fword:  __ccgo_ts + 34216,
		Ftype1: int8('F'),
	},
	4936: {
		Fword:  __ccgo_ts + 34223,
		Ftype1: int8('F'),
	},
	4937: {
		Fword:  __ccgo_ts + 34230,
		Ftype1: int8('F'),
	},
	4938: {
		Fword:  __ccgo_ts + 34237,
		Ftype1: int8('F'),
	},
	4939: {
		Fword:  __ccgo_ts + 34244,
		Ftype1: int8('F'),
	},
	4940: {
		Fword:  __ccgo_ts + 34251,
		Ftype1: int8('F'),
	},
	4941: {
		Fword:  __ccgo_ts + 34258,
		Ftype1: int8('F'),
	},
	4942: {
		Fword:  __ccgo_ts + 34265,
		Ftype1: int8('F'),
	},
	4943: {
		Fword:  __ccgo_ts + 34272,
		Ftype1: int8('F'),
	},
	4944: {
		Fword:  __ccgo_ts + 34279,
		Ftype1: int8('F'),
	},
	4945: {
		Fword:  __ccgo_ts + 34286,
		Ftype1: int8('F'),
	},
	4946: {
		Fword:  __ccgo_ts + 34293,
		Ftype1: int8('F'),
	},
	4947: {
		Fword:  __ccgo_ts + 34298,
		Ftype1: int8('F'),
	},
	4948: {
		Fword:  __ccgo_ts + 34305,
		Ftype1: int8('F'),
	},
	4949: {
		Fword:  __ccgo_ts + 34312,
		Ftype1: int8('F'),
	},
	4950: {
		Fword:  __ccgo_ts + 34319,
		Ftype1: int8('F'),
	},
	4951: {
		Fword:  __ccgo_ts + 34326,
		Ftype1: int8('F'),
	},
	4952: {
		Fword:  __ccgo_ts + 34333,
		Ftype1: int8('F'),
	},
	4953: {
		Fword:  __ccgo_ts + 34340,
		Ftype1: int8('F'),
	},
	4954: {
		Fword:  __ccgo_ts + 34347,
		Ftype1: int8('F'),
	},
	4955: {
		Fword:  __ccgo_ts + 34354,
		Ftype1: int8('F'),
	},
	4956: {
		Fword:  __ccgo_ts + 34361,
		Ftype1: int8('F'),
	},
	4957: {
		Fword:  __ccgo_ts + 34368,
		Ftype1: int8('F'),
	},
	4958: {
		Fword:  __ccgo_ts + 34374,
		Ftype1: int8('F'),
	},
	4959: {
		Fword:  __ccgo_ts + 34381,
		Ftype1: int8('F'),
	},
	4960: {
		Fword:  __ccgo_ts + 34388,
		Ftype1: int8('F'),
	},
	4961: {
		Fword:  __ccgo_ts + 34395,
		Ftype1: int8('F'),
	},
	4962: {
		Fword:  __ccgo_ts + 34402,
		Ftype1: int8('F'),
	},
	4963: {
		Fword:  __ccgo_ts + 34409,
		Ftype1: int8('F'),
	},
	4964: {
		Fword:  __ccgo_ts + 34416,
		Ftype1: int8('F'),
	},
	4965: {
		Fword:  __ccgo_ts + 34423,
		Ftype1: int8('F'),
	},
	4966: {
		Fword:  __ccgo_ts + 34430,
		Ftype1: int8('F'),
	},
	4967: {
		Fword:  __ccgo_ts + 34437,
		Ftype1: int8('F'),
	},
	4968: {
		Fword:  __ccgo_ts + 34443,
		Ftype1: int8('F'),
	},
	4969: {
		Fword:  __ccgo_ts + 34450,
		Ftype1: int8('F'),
	},
	4970: {
		Fword:  __ccgo_ts + 34457,
		Ftype1: int8('F'),
	},
	4971: {
		Fword:  __ccgo_ts + 34464,
		Ftype1: int8('F'),
	},
	4972: {
		Fword:  __ccgo_ts + 34471,
		Ftype1: int8('F'),
	},
	4973: {
		Fword:  __ccgo_ts + 34478,
		Ftype1: int8('F'),
	},
	4974: {
		Fword:  __ccgo_ts + 34485,
		Ftype1: int8('F'),
	},
	4975: {
		Fword:  __ccgo_ts + 34492,
		Ftype1: int8('F'),
	},
	4976: {
		Fword:  __ccgo_ts + 34499,
		Ftype1: int8('F'),
	},
	4977: {
		Fword:  __ccgo_ts + 34506,
		Ftype1: int8('F'),
	},
	4978: {
		Fword:  __ccgo_ts + 34513,
		Ftype1: int8('F'),
	},
	4979: {
		Fword:  __ccgo_ts + 34520,
		Ftype1: int8('F'),
	},
	4980: {
		Fword:  __ccgo_ts + 34527,
		Ftype1: int8('F'),
	},
	4981: {
		Fword:  __ccgo_ts + 34534,
		Ftype1: int8('F'),
	},
	4982: {
		Fword:  __ccgo_ts + 34540,
		Ftype1: int8('F'),
	},
	4983: {
		Fword:  __ccgo_ts + 34547,
		Ftype1: int8('F'),
	},
	4984: {
		Fword:  __ccgo_ts + 34554,
		Ftype1: int8('F'),
	},
	4985: {
		Fword:  __ccgo_ts + 34561,
		Ftype1: int8('F'),
	},
	4986: {
		Fword:  __ccgo_ts + 34568,
		Ftype1: int8('F'),
	},
	4987: {
		Fword:  __ccgo_ts + 34573,
		Ftype1: int8('F'),
	},
	4988: {
		Fword:  __ccgo_ts + 34580,
		Ftype1: int8('F'),
	},
	4989: {
		Fword:  __ccgo_ts + 34587,
		Ftype1: int8('F'),
	},
	4990: {
		Fword:  __ccgo_ts + 34594,
		Ftype1: int8('F'),
	},
	4991: {
		Fword:  __ccgo_ts + 34601,
		Ftype1: int8('F'),
	},
	4992: {
		Fword:  __ccgo_ts + 34608,
		Ftype1: int8('F'),
	},
	4993: {
		Fword:  __ccgo_ts + 34615,
		Ftype1: int8('F'),
	},
	4994: {
		Fword:  __ccgo_ts + 34622,
		Ftype1: int8('F'),
	},
	4995: {
		Fword:  __ccgo_ts + 34629,
		Ftype1: int8('F'),
	},
	4996: {
		Fword:  __ccgo_ts + 34636,
		Ftype1: int8('F'),
	},
	4997: {
		Fword:  __ccgo_ts + 34643,
		Ftype1: int8('F'),
	},
	4998: {
		Fword:  __ccgo_ts + 34649,
		Ftype1: int8('F'),
	},
	4999: {
		Fword:  __ccgo_ts + 34656,
		Ftype1: int8('F'),
	},
	5000: {
		Fword:  __ccgo_ts + 34663,
		Ftype1: int8('F'),
	},
	5001: {
		Fword:  __ccgo_ts + 34669,
		Ftype1: int8('F'),
	},
	5002: {
		Fword:  __ccgo_ts + 34676,
		Ftype1: int8('F'),
	},
	5003: {
		Fword:  __ccgo_ts + 34683,
		Ftype1: int8('F'),
	},
	5004: {
		Fword:  __ccgo_ts + 34690,
		Ftype1: int8('F'),
	},
	5005: {
		Fword:  __ccgo_ts + 34697,
		Ftype1: int8('F'),
	},
	5006: {
		Fword:  __ccgo_ts + 34704,
		Ftype1: int8('F'),
	},
	5007: {
		Fword:  __ccgo_ts + 34711,
		Ftype1: int8('F'),
	},
	5008: {
		Fword:  __ccgo_ts + 34718,
		Ftype1: int8('F'),
	},
	5009: {
		Fword:  __ccgo_ts + 34725,
		Ftype1: int8('F'),
	},
	5010: {
		Fword:  __ccgo_ts + 34732,
		Ftype1: int8('F'),
	},
	5011: {
		Fword:  __ccgo_ts + 34738,
		Ftype1: int8('F'),
	},
	5012: {
		Fword:  __ccgo_ts + 34745,
		Ftype1: int8('F'),
	},
	5013: {
		Fword:  __ccgo_ts + 34752,
		Ftype1: int8('F'),
	},
	5014: {
		Fword:  __ccgo_ts + 34759,
		Ftype1: int8('F'),
	},
	5015: {
		Fword:  __ccgo_ts + 34766,
		Ftype1: int8('F'),
	},
	5016: {
		Fword:  __ccgo_ts + 34773,
		Ftype1: int8('F'),
	},
	5017: {
		Fword:  __ccgo_ts + 34780,
		Ftype1: int8('F'),
	},
	5018: {
		Fword:  __ccgo_ts + 34787,
		Ftype1: int8('F'),
	},
	5019: {
		Fword:  __ccgo_ts + 34794,
		Ftype1: int8('F'),
	},
	5020: {
		Fword:  __ccgo_ts + 34801,
		Ftype1: int8('F'),
	},
	5021: {
		Fword:  __ccgo_ts + 34808,
		Ftype1: int8('F'),
	},
	5022: {
		Fword:  __ccgo_ts + 34815,
		Ftype1: int8('F'),
	},
	5023: {
		Fword:  __ccgo_ts + 34822,
		Ftype1: int8('F'),
	},
	5024: {
		Fword:  __ccgo_ts + 34829,
		Ftype1: int8('F'),
	},
	5025: {
		Fword:  __ccgo_ts + 34836,
		Ftype1: int8('F'),
	},
	5026: {
		Fword:  __ccgo_ts + 34843,
		Ftype1: int8('F'),
	},
	5027: {
		Fword:  __ccgo_ts + 34850,
		Ftype1: int8('F'),
	},
	5028: {
		Fword:  __ccgo_ts + 34856,
		Ftype1: int8('F'),
	},
	5029: {
		Fword:  __ccgo_ts + 34863,
		Ftype1: int8('F'),
	},
	5030: {
		Fword:  __ccgo_ts + 34870,
		Ftype1: int8('F'),
	},
	5031: {
		Fword:  __ccgo_ts + 34877,
		Ftype1: int8('F'),
	},
	5032: {
		Fword:  __ccgo_ts + 34884,
		Ftype1: int8('F'),
	},
	5033: {
		Fword:  __ccgo_ts + 34890,
		Ftype1: int8('F'),
	},
	5034: {
		Fword:  __ccgo_ts + 34897,
		Ftype1: int8('F'),
	},
	5035: {
		Fword:  __ccgo_ts + 34904,
		Ftype1: int8('F'),
	},
	5036: {
		Fword:  __ccgo_ts + 34911,
		Ftype1: int8('F'),
	},
	5037: {
		Fword:  __ccgo_ts + 34916,
		Ftype1: int8('F'),
	},
	5038: {
		Fword:  __ccgo_ts + 34923,
		Ftype1: int8('F'),
	},
	5039: {
		Fword:  __ccgo_ts + 34930,
		Ftype1: int8('F'),
	},
	5040: {
		Fword:  __ccgo_ts + 34937,
		Ftype1: int8('F'),
	},
	5041: {
		Fword:  __ccgo_ts + 34944,
		Ftype1: int8('F'),
	},
	5042: {
		Fword:  __ccgo_ts + 34951,
		Ftype1: int8('F'),
	},
	5043: {
		Fword:  __ccgo_ts + 34958,
		Ftype1: int8('F'),
	},
	5044: {
		Fword:  __ccgo_ts + 34965,
		Ftype1: int8('F'),
	},
	5045: {
		Fword:  __ccgo_ts + 34972,
		Ftype1: int8('F'),
	},
	5046: {
		Fword:  __ccgo_ts + 34979,
		Ftype1: int8('F'),
	},
	5047: {
		Fword:  __ccgo_ts + 34986,
		Ftype1: int8('F'),
	},
	5048: {
		Fword:  __ccgo_ts + 34992,
		Ftype1: int8('F'),
	},
	5049: {
		Fword:  __ccgo_ts + 34999,
		Ftype1: int8('F'),
	},
	5050: {
		Fword:  __ccgo_ts + 35006,
		Ftype1: int8('F'),
	},
	5051: {
		Fword:  __ccgo_ts + 35013,
		Ftype1: int8('F'),
	},
	5052: {
		Fword:  __ccgo_ts + 35020,
		Ftype1: int8('F'),
	},
	5053: {
		Fword:  __ccgo_ts + 35027,
		Ftype1: int8('F'),
	},
	5054: {
		Fword:  __ccgo_ts + 35034,
		Ftype1: int8('F'),
	},
	5055: {
		Fword:  __ccgo_ts + 35041,
		Ftype1: int8('F'),
	},
	5056: {
		Fword:  __ccgo_ts + 35048,
		Ftype1: int8('F'),
	},
	5057: {
		Fword:  __ccgo_ts + 35055,
		Ftype1: int8('F'),
	},
	5058: {
		Fword:  __ccgo_ts + 35061,
		Ftype1: int8('F'),
	},
	5059: {
		Fword:  __ccgo_ts + 35068,
		Ftype1: int8('F'),
	},
	5060: {
		Fword:  __ccgo_ts + 35075,
		Ftype1: int8('F'),
	},
	5061: {
		Fword:  __ccgo_ts + 35082,
		Ftype1: int8('F'),
	},
	5062: {
		Fword:  __ccgo_ts + 35089,
		Ftype1: int8('F'),
	},
	5063: {
		Fword:  __ccgo_ts + 35096,
		Ftype1: int8('F'),
	},
	5064: {
		Fword:  __ccgo_ts + 35103,
		Ftype1: int8('F'),
	},
	5065: {
		Fword:  __ccgo_ts + 35110,
		Ftype1: int8('F'),
	},
	5066: {
		Fword:  __ccgo_ts + 35117,
		Ftype1: int8('F'),
	},
	5067: {
		Fword:  __ccgo_ts + 35124,
		Ftype1: int8('F'),
	},
	5068: {
		Fword:  __ccgo_ts + 35131,
		Ftype1: int8('F'),
	},
	5069: {
		Fword:  __ccgo_ts + 35138,
		Ftype1: int8('F'),
	},
	5070: {
		Fword:  __ccgo_ts + 35145,
		Ftype1: int8('F'),
	},
	5071: {
		Fword:  __ccgo_ts + 35151,
		Ftype1: int8('F'),
	},
	5072: {
		Fword:  __ccgo_ts + 35158,
		Ftype1: int8('F'),
	},
	5073: {
		Fword:  __ccgo_ts + 35165,
		Ftype1: int8('F'),
	},
	5074: {
		Fword:  __ccgo_ts + 35172,
		Ftype1: int8('F'),
	},
	5075: {
		Fword:  __ccgo_ts + 35179,
		Ftype1: int8('F'),
	},
	5076: {
		Fword:  __ccgo_ts + 35185,
		Ftype1: int8('F'),
	},
	5077: {
		Fword:  __ccgo_ts + 35192,
		Ftype1: int8('F'),
	},
	5078: {
		Fword:  __ccgo_ts + 35199,
		Ftype1: int8('F'),
	},
	5079: {
		Fword:  __ccgo_ts + 35206,
		Ftype1: int8('F'),
	},
	5080: {
		Fword:  __ccgo_ts + 35213,
		Ftype1: int8('F'),
	},
	5081: {
		Fword:  __ccgo_ts + 35220,
		Ftype1: int8('F'),
	},
	5082: {
		Fword:  __ccgo_ts + 35227,
		Ftype1: int8('F'),
	},
	5083: {
		Fword:  __ccgo_ts + 35234,
		Ftype1: int8('F'),
	},
	5084: {
		Fword:  __ccgo_ts + 35241,
		Ftype1: int8('F'),
	},
	5085: {
		Fword:  __ccgo_ts + 35248,
		Ftype1: int8('F'),
	},
	5086: {
		Fword:  __ccgo_ts + 35255,
		Ftype1: int8('F'),
	},
	5087: {
		Fword:  __ccgo_ts + 35262,
		Ftype1: int8('F'),
	},
	5088: {
		Fword:  __ccgo_ts + 35269,
		Ftype1: int8('F'),
	},
	5089: {
		Fword:  __ccgo_ts + 35276,
		Ftype1: int8('F'),
	},
	5090: {
		Fword:  __ccgo_ts + 35283,
		Ftype1: int8('F'),
	},
	5091: {
		Fword:  __ccgo_ts + 35290,
		Ftype1: int8('F'),
	},
	5092: {
		Fword:  __ccgo_ts + 35296,
		Ftype1: int8('F'),
	},
	5093: {
		Fword:  __ccgo_ts + 35303,
		Ftype1: int8('F'),
	},
	5094: {
		Fword:  __ccgo_ts + 35310,
		Ftype1: int8('F'),
	},
	5095: {
		Fword:  __ccgo_ts + 35317,
		Ftype1: int8('F'),
	},
	5096: {
		Fword:  __ccgo_ts + 35324,
		Ftype1: int8('F'),
	},
	5097: {
		Fword:  __ccgo_ts + 35331,
		Ftype1: int8('F'),
	},
	5098: {
		Fword:  __ccgo_ts + 35338,
		Ftype1: int8('F'),
	},
	5099: {
		Fword:  __ccgo_ts + 35345,
		Ftype1: int8('F'),
	},
	5100: {
		Fword:  __ccgo_ts + 35352,
		Ftype1: int8('F'),
	},
	5101: {
		Fword:  __ccgo_ts + 35359,
		Ftype1: int8('F'),
	},
	5102: {
		Fword:  __ccgo_ts + 35365,
		Ftype1: int8('F'),
	},
	5103: {
		Fword:  __ccgo_ts + 35372,
		Ftype1: int8('F'),
	},
	5104: {
		Fword:  __ccgo_ts + 35379,
		Ftype1: int8('F'),
	},
	5105: {
		Fword:  __ccgo_ts + 35386,
		Ftype1: int8('F'),
	},
	5106: {
		Fword:  __ccgo_ts + 35393,
		Ftype1: int8('F'),
	},
	5107: {
		Fword:  __ccgo_ts + 35400,
		Ftype1: int8('F'),
	},
	5108: {
		Fword:  __ccgo_ts + 35407,
		Ftype1: int8('F'),
	},
	5109: {
		Fword:  __ccgo_ts + 35414,
		Ftype1: int8('F'),
	},
	5110: {
		Fword:  __ccgo_ts + 35421,
		Ftype1: int8('F'),
	},
	5111: {
		Fword:  __ccgo_ts + 35427,
		Ftype1: int8('F'),
	},
	5112: {
		Fword:  __ccgo_ts + 35434,
		Ftype1: int8('F'),
	},
	5113: {
		Fword:  __ccgo_ts + 35441,
		Ftype1: int8('F'),
	},
	5114: {
		Fword:  __ccgo_ts + 35448,
		Ftype1: int8('F'),
	},
	5115: {
		Fword:  __ccgo_ts + 35455,
		Ftype1: int8('F'),
	},
	5116: {
		Fword:  __ccgo_ts + 35462,
		Ftype1: int8('F'),
	},
	5117: {
		Fword:  __ccgo_ts + 35469,
		Ftype1: int8('F'),
	},
	5118: {
		Fword:  __ccgo_ts + 35476,
		Ftype1: int8('F'),
	},
	5119: {
		Fword:  __ccgo_ts + 35483,
		Ftype1: int8('F'),
	},
	5120: {
		Fword:  __ccgo_ts + 35489,
		Ftype1: int8('F'),
	},
	5121: {
		Fword:  __ccgo_ts + 35496,
		Ftype1: int8('F'),
	},
	5122: {
		Fword:  __ccgo_ts + 35503,
		Ftype1: int8('F'),
	},
	5123: {
		Fword:  __ccgo_ts + 35510,
		Ftype1: int8('F'),
	},
	5124: {
		Fword:  __ccgo_ts + 35517,
		Ftype1: int8('F'),
	},
	5125: {
		Fword:  __ccgo_ts + 35524,
		Ftype1: int8('F'),
	},
	5126: {
		Fword:  __ccgo_ts + 35531,
		Ftype1: int8('F'),
	},
	5127: {
		Fword:  __ccgo_ts + 35538,
		Ftype1: int8('F'),
	},
	5128: {
		Fword:  __ccgo_ts + 35545,
		Ftype1: int8('F'),
	},
	5129: {
		Fword:  __ccgo_ts + 35552,
		Ftype1: int8('F'),
	},
	5130: {
		Fword:  __ccgo_ts + 35559,
		Ftype1: int8('F'),
	},
	5131: {
		Fword:  __ccgo_ts + 35566,
		Ftype1: int8('F'),
	},
	5132: {
		Fword:  __ccgo_ts + 35573,
		Ftype1: int8('F'),
	},
	5133: {
		Fword:  __ccgo_ts + 35580,
		Ftype1: int8('F'),
	},
	5134: {
		Fword:  __ccgo_ts + 35587,
		Ftype1: int8('F'),
	},
	5135: {
		Fword:  __ccgo_ts + 35594,
		Ftype1: int8('F'),
	},
	5136: {
		Fword:  __ccgo_ts + 35601,
		Ftype1: int8('F'),
	},
	5137: {
		Fword:  __ccgo_ts + 35608,
		Ftype1: int8('F'),
	},
	5138: {
		Fword:  __ccgo_ts + 35615,
		Ftype1: int8('F'),
	},
	5139: {
		Fword:  __ccgo_ts + 35622,
		Ftype1: int8('F'),
	},
	5140: {
		Fword:  __ccgo_ts + 35629,
		Ftype1: int8('F'),
	},
	5141: {
		Fword:  __ccgo_ts + 35636,
		Ftype1: int8('F'),
	},
	5142: {
		Fword:  __ccgo_ts + 35643,
		Ftype1: int8('F'),
	},
	5143: {
		Fword:  __ccgo_ts + 35650,
		Ftype1: int8('F'),
	},
	5144: {
		Fword:  __ccgo_ts + 35657,
		Ftype1: int8('F'),
	},
	5145: {
		Fword:  __ccgo_ts + 35664,
		Ftype1: int8('F'),
	},
	5146: {
		Fword:  __ccgo_ts + 35671,
		Ftype1: int8('F'),
	},
	5147: {
		Fword:  __ccgo_ts + 35678,
		Ftype1: int8('F'),
	},
	5148: {
		Fword:  __ccgo_ts + 35685,
		Ftype1: int8('F'),
	},
	5149: {
		Fword:  __ccgo_ts + 35692,
		Ftype1: int8('F'),
	},
	5150: {
		Fword:  __ccgo_ts + 35699,
		Ftype1: int8('F'),
	},
	5151: {
		Fword:  __ccgo_ts + 35706,
		Ftype1: int8('F'),
	},
	5152: {
		Fword:  __ccgo_ts + 35713,
		Ftype1: int8('F'),
	},
	5153: {
		Fword:  __ccgo_ts + 35720,
		Ftype1: int8('F'),
	},
	5154: {
		Fword:  __ccgo_ts + 35727,
		Ftype1: int8('F'),
	},
	5155: {
		Fword:  __ccgo_ts + 35733,
		Ftype1: int8('F'),
	},
	5156: {
		Fword:  __ccgo_ts + 35740,
		Ftype1: int8('F'),
	},
	5157: {
		Fword:  __ccgo_ts + 35747,
		Ftype1: int8('F'),
	},
	5158: {
		Fword:  __ccgo_ts + 35754,
		Ftype1: int8('F'),
	},
	5159: {
		Fword:  __ccgo_ts + 35761,
		Ftype1: int8('F'),
	},
	5160: {
		Fword:  __ccgo_ts + 35768,
		Ftype1: int8('F'),
	},
	5161: {
		Fword:  __ccgo_ts + 35775,
		Ftype1: int8('F'),
	},
	5162: {
		Fword:  __ccgo_ts + 35782,
		Ftype1: int8('F'),
	},
	5163: {
		Fword:  __ccgo_ts + 35789,
		Ftype1: int8('F'),
	},
	5164: {
		Fword:  __ccgo_ts + 35795,
		Ftype1: int8('F'),
	},
	5165: {
		Fword:  __ccgo_ts + 35802,
		Ftype1: int8('F'),
	},
	5166: {
		Fword:  __ccgo_ts + 35809,
		Ftype1: int8('F'),
	},
	5167: {
		Fword:  __ccgo_ts + 35816,
		Ftype1: int8('F'),
	},
	5168: {
		Fword:  __ccgo_ts + 35823,
		Ftype1: int8('F'),
	},
	5169: {
		Fword:  __ccgo_ts + 35830,
		Ftype1: int8('F'),
	},
	5170: {
		Fword:  __ccgo_ts + 35837,
		Ftype1: int8('F'),
	},
	5171: {
		Fword:  __ccgo_ts + 35843,
		Ftype1: int8('F'),
	},
	5172: {
		Fword:  __ccgo_ts + 35850,
		Ftype1: int8('F'),
	},
	5173: {
		Fword:  __ccgo_ts + 35857,
		Ftype1: int8('F'),
	},
	5174: {
		Fword:  __ccgo_ts + 35864,
		Ftype1: int8('F'),
	},
	5175: {
		Fword:  __ccgo_ts + 35871,
		Ftype1: int8('F'),
	},
	5176: {
		Fword:  __ccgo_ts + 35878,
		Ftype1: int8('F'),
	},
	5177: {
		Fword:  __ccgo_ts + 35885,
		Ftype1: int8('F'),
	},
	5178: {
		Fword:  __ccgo_ts + 35891,
		Ftype1: int8('F'),
	},
	5179: {
		Fword:  __ccgo_ts + 35898,
		Ftype1: int8('F'),
	},
	5180: {
		Fword:  __ccgo_ts + 35905,
		Ftype1: int8('F'),
	},
	5181: {
		Fword:  __ccgo_ts + 35912,
		Ftype1: int8('F'),
	},
	5182: {
		Fword:  __ccgo_ts + 35919,
		Ftype1: int8('F'),
	},
	5183: {
		Fword:  __ccgo_ts + 35926,
		Ftype1: int8('F'),
	},
	5184: {
		Fword:  __ccgo_ts + 35933,
		Ftype1: int8('F'),
	},
	5185: {
		Fword:  __ccgo_ts + 35938,
		Ftype1: int8('F'),
	},
	5186: {
		Fword:  __ccgo_ts + 35945,
		Ftype1: int8('F'),
	},
	5187: {
		Fword:  __ccgo_ts + 35952,
		Ftype1: int8('F'),
	},
	5188: {
		Fword:  __ccgo_ts + 35959,
		Ftype1: int8('F'),
	},
	5189: {
		Fword:  __ccgo_ts + 35966,
		Ftype1: int8('F'),
	},
	5190: {
		Fword:  __ccgo_ts + 35973,
		Ftype1: int8('F'),
	},
	5191: {
		Fword:  __ccgo_ts + 35980,
		Ftype1: int8('F'),
	},
	5192: {
		Fword:  __ccgo_ts + 35987,
		Ftype1: int8('F'),
	},
	5193: {
		Fword:  __ccgo_ts + 35994,
		Ftype1: int8('F'),
	},
	5194: {
		Fword:  __ccgo_ts + 36001,
		Ftype1: int8('F'),
	},
	5195: {
		Fword:  __ccgo_ts + 36008,
		Ftype1: int8('F'),
	},
	5196: {
		Fword:  __ccgo_ts + 36015,
		Ftype1: int8('F'),
	},
	5197: {
		Fword:  __ccgo_ts + 36022,
		Ftype1: int8('F'),
	},
	5198: {
		Fword:  __ccgo_ts + 36029,
		Ftype1: int8('F'),
	},
	5199: {
		Fword:  __ccgo_ts + 36036,
		Ftype1: int8('F'),
	},
	5200: {
		Fword:  __ccgo_ts + 36043,
		Ftype1: int8('F'),
	},
	5201: {
		Fword:  __ccgo_ts + 36050,
		Ftype1: int8('F'),
	},
	5202: {
		Fword:  __ccgo_ts + 36057,
		Ftype1: int8('F'),
	},
	5203: {
		Fword:  __ccgo_ts + 36064,
		Ftype1: int8('F'),
	},
	5204: {
		Fword:  __ccgo_ts + 36071,
		Ftype1: int8('F'),
	},
	5205: {
		Fword:  __ccgo_ts + 36078,
		Ftype1: int8('F'),
	},
	5206: {
		Fword:  __ccgo_ts + 36085,
		Ftype1: int8('F'),
	},
	5207: {
		Fword:  __ccgo_ts + 36092,
		Ftype1: int8('F'),
	},
	5208: {
		Fword:  __ccgo_ts + 36099,
		Ftype1: int8('F'),
	},
	5209: {
		Fword:  __ccgo_ts + 36106,
		Ftype1: int8('F'),
	},
	5210: {
		Fword:  __ccgo_ts + 36113,
		Ftype1: int8('F'),
	},
	5211: {
		Fword:  __ccgo_ts + 36120,
		Ftype1: int8('F'),
	},
	5212: {
		Fword:  __ccgo_ts + 36127,
		Ftype1: int8('F'),
	},
	5213: {
		Fword:  __ccgo_ts + 36134,
		Ftype1: int8('F'),
	},
	5214: {
		Fword:  __ccgo_ts + 36141,
		Ftype1: int8('F'),
	},
	5215: {
		Fword:  __ccgo_ts + 36148,
		Ftype1: int8('F'),
	},
	5216: {
		Fword:  __ccgo_ts + 36155,
		Ftype1: int8('F'),
	},
	5217: {
		Fword:  __ccgo_ts + 36162,
		Ftype1: int8('F'),
	},
	5218: {
		Fword:  __ccgo_ts + 36169,
		Ftype1: int8('F'),
	},
	5219: {
		Fword:  __ccgo_ts + 36176,
		Ftype1: int8('F'),
	},
	5220: {
		Fword:  __ccgo_ts + 36183,
		Ftype1: int8('F'),
	},
	5221: {
		Fword:  __ccgo_ts + 36190,
		Ftype1: int8('F'),
	},
	5222: {
		Fword:  __ccgo_ts + 36197,
		Ftype1: int8('F'),
	},
	5223: {
		Fword:  __ccgo_ts + 36204,
		Ftype1: int8('F'),
	},
	5224: {
		Fword:  __ccgo_ts + 36211,
		Ftype1: int8('F'),
	},
	5225: {
		Fword:  __ccgo_ts + 36218,
		Ftype1: int8('F'),
	},
	5226: {
		Fword:  __ccgo_ts + 36225,
		Ftype1: int8('F'),
	},
	5227: {
		Fword:  __ccgo_ts + 36232,
		Ftype1: int8('F'),
	},
	5228: {
		Fword:  __ccgo_ts + 36239,
		Ftype1: int8('F'),
	},
	5229: {
		Fword:  __ccgo_ts + 36246,
		Ftype1: int8('F'),
	},
	5230: {
		Fword:  __ccgo_ts + 36253,
		Ftype1: int8('F'),
	},
	5231: {
		Fword:  __ccgo_ts + 36260,
		Ftype1: int8('F'),
	},
	5232: {
		Fword:  __ccgo_ts + 36267,
		Ftype1: int8('F'),
	},
	5233: {
		Fword:  __ccgo_ts + 36274,
		Ftype1: int8('F'),
	},
	5234: {
		Fword:  __ccgo_ts + 36281,
		Ftype1: int8('F'),
	},
	5235: {
		Fword:  __ccgo_ts + 36288,
		Ftype1: int8('F'),
	},
	5236: {
		Fword:  __ccgo_ts + 36295,
		Ftype1: int8('F'),
	},
	5237: {
		Fword:  __ccgo_ts + 36302,
		Ftype1: int8('F'),
	},
	5238: {
		Fword:  __ccgo_ts + 36309,
		Ftype1: int8('F'),
	},
	5239: {
		Fword:  __ccgo_ts + 36316,
		Ftype1: int8('F'),
	},
	5240: {
		Fword:  __ccgo_ts + 36323,
		Ftype1: int8('F'),
	},
	5241: {
		Fword:  __ccgo_ts + 36330,
		Ftype1: int8('F'),
	},
	5242: {
		Fword:  __ccgo_ts + 36337,
		Ftype1: int8('F'),
	},
	5243: {
		Fword:  __ccgo_ts + 36344,
		Ftype1: int8('F'),
	},
	5244: {
		Fword:  __ccgo_ts + 36351,
		Ftype1: int8('F'),
	},
	5245: {
		Fword:  __ccgo_ts + 36358,
		Ftype1: int8('F'),
	},
	5246: {
		Fword:  __ccgo_ts + 36365,
		Ftype1: int8('F'),
	},
	5247: {
		Fword:  __ccgo_ts + 36372,
		Ftype1: int8('F'),
	},
	5248: {
		Fword:  __ccgo_ts + 36379,
		Ftype1: int8('F'),
	},
	5249: {
		Fword:  __ccgo_ts + 36386,
		Ftype1: int8('F'),
	},
	5250: {
		Fword:  __ccgo_ts + 36393,
		Ftype1: int8('F'),
	},
	5251: {
		Fword:  __ccgo_ts + 36399,
		Ftype1: int8('F'),
	},
	5252: {
		Fword:  __ccgo_ts + 36406,
		Ftype1: int8('F'),
	},
	5253: {
		Fword:  __ccgo_ts + 36413,
		Ftype1: int8('F'),
	},
	5254: {
		Fword:  __ccgo_ts + 36420,
		Ftype1: int8('F'),
	},
	5255: {
		Fword:  __ccgo_ts + 36427,
		Ftype1: int8('F'),
	},
	5256: {
		Fword:  __ccgo_ts + 36434,
		Ftype1: int8('F'),
	},
	5257: {
		Fword:  __ccgo_ts + 36441,
		Ftype1: int8('F'),
	},
	5258: {
		Fword:  __ccgo_ts + 36448,
		Ftype1: int8('F'),
	},
	5259: {
		Fword:  __ccgo_ts + 36455,
		Ftype1: int8('F'),
	},
	5260: {
		Fword:  __ccgo_ts + 36462,
		Ftype1: int8('F'),
	},
	5261: {
		Fword:  __ccgo_ts + 36469,
		Ftype1: int8('F'),
	},
	5262: {
		Fword:  __ccgo_ts + 36476,
		Ftype1: int8('F'),
	},
	5263: {
		Fword:  __ccgo_ts + 36483,
		Ftype1: int8('F'),
	},
	5264: {
		Fword:  __ccgo_ts + 36490,
		Ftype1: int8('F'),
	},
	5265: {
		Fword:  __ccgo_ts + 36497,
		Ftype1: int8('F'),
	},
	5266: {
		Fword:  __ccgo_ts + 36504,
		Ftype1: int8('F'),
	},
	5267: {
		Fword:  __ccgo_ts + 36510,
		Ftype1: int8('F'),
	},
	5268: {
		Fword:  __ccgo_ts + 36517,
		Ftype1: int8('F'),
	},
	5269: {
		Fword:  __ccgo_ts + 36524,
		Ftype1: int8('F'),
	},
	5270: {
		Fword:  __ccgo_ts + 36531,
		Ftype1: int8('F'),
	},
	5271: {
		Fword:  __ccgo_ts + 36538,
		Ftype1: int8('F'),
	},
	5272: {
		Fword:  __ccgo_ts + 36545,
		Ftype1: int8('F'),
	},
	5273: {
		Fword:  __ccgo_ts + 36552,
		Ftype1: int8('F'),
	},
	5274: {
		Fword:  __ccgo_ts + 36559,
		Ftype1: int8('F'),
	},
	5275: {
		Fword:  __ccgo_ts + 36565,
		Ftype1: int8('F'),
	},
	5276: {
		Fword:  __ccgo_ts + 36572,
		Ftype1: int8('F'),
	},
	5277: {
		Fword:  __ccgo_ts + 36579,
		Ftype1: int8('F'),
	},
	5278: {
		Fword:  __ccgo_ts + 36586,
		Ftype1: int8('F'),
	},
	5279: {
		Fword:  __ccgo_ts + 36593,
		Ftype1: int8('F'),
	},
	5280: {
		Fword:  __ccgo_ts + 36600,
		Ftype1: int8('F'),
	},
	5281: {
		Fword:  __ccgo_ts + 36607,
		Ftype1: int8('F'),
	},
	5282: {
		Fword:  __ccgo_ts + 36614,
		Ftype1: int8('F'),
	},
	5283: {
		Fword:  __ccgo_ts + 36621,
		Ftype1: int8('F'),
	},
	5284: {
		Fword:  __ccgo_ts + 36628,
		Ftype1: int8('F'),
	},
	5285: {
		Fword:  __ccgo_ts + 36635,
		Ftype1: int8('F'),
	},
	5286: {
		Fword:  __ccgo_ts + 36642,
		Ftype1: int8('F'),
	},
	5287: {
		Fword:  __ccgo_ts + 36649,
		Ftype1: int8('F'),
	},
	5288: {
		Fword:  __ccgo_ts + 36656,
		Ftype1: int8('F'),
	},
	5289: {
		Fword:  __ccgo_ts + 36663,
		Ftype1: int8('F'),
	},
	5290: {
		Fword:  __ccgo_ts + 36670,
		Ftype1: int8('F'),
	},
	5291: {
		Fword:  __ccgo_ts + 36677,
		Ftype1: int8('F'),
	},
	5292: {
		Fword:  __ccgo_ts + 36684,
		Ftype1: int8('F'),
	},
	5293: {
		Fword:  __ccgo_ts + 36691,
		Ftype1: int8('F'),
	},
	5294: {
		Fword:  __ccgo_ts + 36698,
		Ftype1: int8('F'),
	},
	5295: {
		Fword:  __ccgo_ts + 36705,
		Ftype1: int8('F'),
	},
	5296: {
		Fword:  __ccgo_ts + 36712,
		Ftype1: int8('F'),
	},
	5297: {
		Fword:  __ccgo_ts + 36719,
		Ftype1: int8('F'),
	},
	5298: {
		Fword:  __ccgo_ts + 36726,
		Ftype1: int8('F'),
	},
	5299: {
		Fword:  __ccgo_ts + 36733,
		Ftype1: int8('F'),
	},
	5300: {
		Fword:  __ccgo_ts + 36740,
		Ftype1: int8('F'),
	},
	5301: {
		Fword:  __ccgo_ts + 36747,
		Ftype1: int8('F'),
	},
	5302: {
		Fword:  __ccgo_ts + 36754,
		Ftype1: int8('F'),
	},
	5303: {
		Fword:  __ccgo_ts + 36761,
		Ftype1: int8('F'),
	},
	5304: {
		Fword:  __ccgo_ts + 36768,
		Ftype1: int8('F'),
	},
	5305: {
		Fword:  __ccgo_ts + 36775,
		Ftype1: int8('F'),
	},
	5306: {
		Fword:  __ccgo_ts + 36782,
		Ftype1: int8('F'),
	},
	5307: {
		Fword:  __ccgo_ts + 36789,
		Ftype1: int8('F'),
	},
	5308: {
		Fword:  __ccgo_ts + 36796,
		Ftype1: int8('F'),
	},
	5309: {
		Fword:  __ccgo_ts + 36803,
		Ftype1: int8('F'),
	},
	5310: {
		Fword:  __ccgo_ts + 36810,
		Ftype1: int8('F'),
	},
	5311: {
		Fword:  __ccgo_ts + 36817,
		Ftype1: int8('F'),
	},
	5312: {
		Fword:  __ccgo_ts + 36824,
		Ftype1: int8('F'),
	},
	5313: {
		Fword:  __ccgo_ts + 36831,
		Ftype1: int8('F'),
	},
	5314: {
		Fword:  __ccgo_ts + 36838,
		Ftype1: int8('F'),
	},
	5315: {
		Fword:  __ccgo_ts + 36845,
		Ftype1: int8('F'),
	},
	5316: {
		Fword:  __ccgo_ts + 36852,
		Ftype1: int8('F'),
	},
	5317: {
		Fword:  __ccgo_ts + 36859,
		Ftype1: int8('F'),
	},
	5318: {
		Fword:  __ccgo_ts + 36866,
		Ftype1: int8('F'),
	},
	5319: {
		Fword:  __ccgo_ts + 36873,
		Ftype1: int8('F'),
	},
	5320: {
		Fword:  __ccgo_ts + 36879,
		Ftype1: int8('F'),
	},
	5321: {
		Fword:  __ccgo_ts + 36885,
		Ftype1: int8('F'),
	},
	5322: {
		Fword:  __ccgo_ts + 36892,
		Ftype1: int8('F'),
	},
	5323: {
		Fword:  __ccgo_ts + 36898,
		Ftype1: int8('F'),
	},
	5324: {
		Fword:  __ccgo_ts + 36904,
		Ftype1: int8('F'),
	},
	5325: {
		Fword:  __ccgo_ts + 36911,
		Ftype1: int8('F'),
	},
	5326: {
		Fword:  __ccgo_ts + 36918,
		Ftype1: int8('F'),
	},
	5327: {
		Fword:  __ccgo_ts + 36925,
		Ftype1: int8('F'),
	},
	5328: {
		Fword:  __ccgo_ts + 36930,
		Ftype1: int8('F'),
	},
	5329: {
		Fword:  __ccgo_ts + 36936,
		Ftype1: int8('F'),
	},
	5330: {
		Fword:  __ccgo_ts + 36943,
		Ftype1: int8('F'),
	},
	5331: {
		Fword:  __ccgo_ts + 36949,
		Ftype1: int8('F'),
	},
	5332: {
		Fword:  __ccgo_ts + 36956,
		Ftype1: int8('F'),
	},
	5333: {
		Fword:  __ccgo_ts + 36963,
		Ftype1: int8('F'),
	},
	5334: {
		Fword:  __ccgo_ts + 36970,
		Ftype1: int8('F'),
	},
	5335: {
		Fword:  __ccgo_ts + 36977,
		Ftype1: int8('F'),
	},
	5336: {
		Fword:  __ccgo_ts + 36984,
		Ftype1: int8('F'),
	},
	5337: {
		Fword:  __ccgo_ts + 36991,
		Ftype1: int8('F'),
	},
	5338: {
		Fword:  __ccgo_ts + 36998,
		Ftype1: int8('F'),
	},
	5339: {
		Fword:  __ccgo_ts + 37005,
		Ftype1: int8('F'),
	},
	5340: {
		Fword:  __ccgo_ts + 37012,
		Ftype1: int8('F'),
	},
	5341: {
		Fword:  __ccgo_ts + 37019,
		Ftype1: int8('F'),
	},
	5342: {
		Fword:  __ccgo_ts + 37026,
		Ftype1: int8('F'),
	},
	5343: {
		Fword:  __ccgo_ts + 37033,
		Ftype1: int8('F'),
	},
	5344: {
		Fword:  __ccgo_ts + 37040,
		Ftype1: int8('F'),
	},
	5345: {
		Fword:  __ccgo_ts + 37047,
		Ftype1: int8('F'),
	},
	5346: {
		Fword:  __ccgo_ts + 37054,
		Ftype1: int8('F'),
	},
	5347: {
		Fword:  __ccgo_ts + 37061,
		Ftype1: int8('F'),
	},
	5348: {
		Fword:  __ccgo_ts + 37068,
		Ftype1: int8('F'),
	},
	5349: {
		Fword:  __ccgo_ts + 37075,
		Ftype1: int8('F'),
	},
	5350: {
		Fword:  __ccgo_ts + 37082,
		Ftype1: int8('F'),
	},
	5351: {
		Fword:  __ccgo_ts + 37089,
		Ftype1: int8('F'),
	},
	5352: {
		Fword:  __ccgo_ts + 37096,
		Ftype1: int8('F'),
	},
	5353: {
		Fword:  __ccgo_ts + 37103,
		Ftype1: int8('F'),
	},
	5354: {
		Fword:  __ccgo_ts + 37110,
		Ftype1: int8('F'),
	},
	5355: {
		Fword:  __ccgo_ts + 37117,
		Ftype1: int8('F'),
	},
	5356: {
		Fword:  __ccgo_ts + 37124,
		Ftype1: int8('F'),
	},
	5357: {
		Fword:  __ccgo_ts + 37131,
		Ftype1: int8('F'),
	},
	5358: {
		Fword:  __ccgo_ts + 37138,
		Ftype1: int8('F'),
	},
	5359: {
		Fword:  __ccgo_ts + 37145,
		Ftype1: int8('F'),
	},
	5360: {
		Fword:  __ccgo_ts + 37152,
		Ftype1: int8('F'),
	},
	5361: {
		Fword:  __ccgo_ts + 37159,
		Ftype1: int8('F'),
	},
	5362: {
		Fword:  __ccgo_ts + 37166,
		Ftype1: int8('F'),
	},
	5363: {
		Fword:  __ccgo_ts + 37173,
		Ftype1: int8('F'),
	},
	5364: {
		Fword:  __ccgo_ts + 37180,
		Ftype1: int8('F'),
	},
	5365: {
		Fword:  __ccgo_ts + 37187,
		Ftype1: int8('F'),
	},
	5366: {
		Fword:  __ccgo_ts + 37194,
		Ftype1: int8('F'),
	},
	5367: {
		Fword:  __ccgo_ts + 37201,
		Ftype1: int8('F'),
	},
	5368: {
		Fword:  __ccgo_ts + 37208,
		Ftype1: int8('F'),
	},
	5369: {
		Fword:  __ccgo_ts + 37215,
		Ftype1: int8('F'),
	},
	5370: {
		Fword:  __ccgo_ts + 37222,
		Ftype1: int8('F'),
	},
	5371: {
		Fword:  __ccgo_ts + 37229,
		Ftype1: int8('F'),
	},
	5372: {
		Fword:  __ccgo_ts + 37236,
		Ftype1: int8('F'),
	},
	5373: {
		Fword:  __ccgo_ts + 37243,
		Ftype1: int8('F'),
	},
	5374: {
		Fword:  __ccgo_ts + 37250,
		Ftype1: int8('F'),
	},
	5375: {
		Fword:  __ccgo_ts + 37257,
		Ftype1: int8('F'),
	},
	5376: {
		Fword:  __ccgo_ts + 37264,
		Ftype1: int8('F'),
	},
	5377: {
		Fword:  __ccgo_ts + 37271,
		Ftype1: int8('F'),
	},
	5378: {
		Fword:  __ccgo_ts + 37278,
		Ftype1: int8('F'),
	},
	5379: {
		Fword:  __ccgo_ts + 37285,
		Ftype1: int8('F'),
	},
	5380: {
		Fword:  __ccgo_ts + 37292,
		Ftype1: int8('F'),
	},
	5381: {
		Fword:  __ccgo_ts + 37299,
		Ftype1: int8('F'),
	},
	5382: {
		Fword:  __ccgo_ts + 37306,
		Ftype1: int8('F'),
	},
	5383: {
		Fword:  __ccgo_ts + 37313,
		Ftype1: int8('F'),
	},
	5384: {
		Fword:  __ccgo_ts + 37320,
		Ftype1: int8('F'),
	},
	5385: {
		Fword:  __ccgo_ts + 37327,
		Ftype1: int8('F'),
	},
	5386: {
		Fword:  __ccgo_ts + 37334,
		Ftype1: int8('F'),
	},
	5387: {
		Fword:  __ccgo_ts + 37341,
		Ftype1: int8('F'),
	},
	5388: {
		Fword:  __ccgo_ts + 37348,
		Ftype1: int8('F'),
	},
	5389: {
		Fword:  __ccgo_ts + 37355,
		Ftype1: int8('F'),
	},
	5390: {
		Fword:  __ccgo_ts + 37361,
		Ftype1: int8('F'),
	},
	5391: {
		Fword:  __ccgo_ts + 37368,
		Ftype1: int8('F'),
	},
	5392: {
		Fword:  __ccgo_ts + 37375,
		Ftype1: int8('F'),
	},
	5393: {
		Fword:  __ccgo_ts + 37382,
		Ftype1: int8('F'),
	},
	5394: {
		Fword:  __ccgo_ts + 37389,
		Ftype1: int8('F'),
	},
	5395: {
		Fword:  __ccgo_ts + 37396,
		Ftype1: int8('F'),
	},
	5396: {
		Fword:  __ccgo_ts + 37403,
		Ftype1: int8('F'),
	},
	5397: {
		Fword:  __ccgo_ts + 37410,
		Ftype1: int8('F'),
	},
	5398: {
		Fword:  __ccgo_ts + 37417,
		Ftype1: int8('F'),
	},
	5399: {
		Fword:  __ccgo_ts + 37424,
		Ftype1: int8('F'),
	},
	5400: {
		Fword:  __ccgo_ts + 37431,
		Ftype1: int8('F'),
	},
	5401: {
		Fword:  __ccgo_ts + 37438,
		Ftype1: int8('F'),
	},
	5402: {
		Fword:  __ccgo_ts + 37445,
		Ftype1: int8('F'),
	},
	5403: {
		Fword:  __ccgo_ts + 37452,
		Ftype1: int8('F'),
	},
	5404: {
		Fword:  __ccgo_ts + 37459,
		Ftype1: int8('F'),
	},
	5405: {
		Fword:  __ccgo_ts + 37466,
		Ftype1: int8('F'),
	},
	5406: {
		Fword:  __ccgo_ts + 37473,
		Ftype1: int8('F'),
	},
	5407: {
		Fword:  __ccgo_ts + 37480,
		Ftype1: int8('F'),
	},
	5408: {
		Fword:  __ccgo_ts + 37487,
		Ftype1: int8('F'),
	},
	5409: {
		Fword:  __ccgo_ts + 37494,
		Ftype1: int8('F'),
	},
	5410: {
		Fword:  __ccgo_ts + 37501,
		Ftype1: int8('F'),
	},
	5411: {
		Fword:  __ccgo_ts + 37508,
		Ftype1: int8('F'),
	},
	5412: {
		Fword:  __ccgo_ts + 37515,
		Ftype1: int8('F'),
	},
	5413: {
		Fword:  __ccgo_ts + 37522,
		Ftype1: int8('F'),
	},
	5414: {
		Fword:  __ccgo_ts + 37529,
		Ftype1: int8('F'),
	},
	5415: {
		Fword:  __ccgo_ts + 37536,
		Ftype1: int8('F'),
	},
	5416: {
		Fword:  __ccgo_ts + 37543,
		Ftype1: int8('F'),
	},
	5417: {
		Fword:  __ccgo_ts + 37550,
		Ftype1: int8('F'),
	},
	5418: {
		Fword:  __ccgo_ts + 37557,
		Ftype1: int8('F'),
	},
	5419: {
		Fword:  __ccgo_ts + 37564,
		Ftype1: int8('F'),
	},
	5420: {
		Fword:  __ccgo_ts + 37571,
		Ftype1: int8('F'),
	},
	5421: {
		Fword:  __ccgo_ts + 37578,
		Ftype1: int8('F'),
	},
	5422: {
		Fword:  __ccgo_ts + 37585,
		Ftype1: int8('F'),
	},
	5423: {
		Fword:  __ccgo_ts + 37592,
		Ftype1: int8('F'),
	},
	5424: {
		Fword:  __ccgo_ts + 37599,
		Ftype1: int8('F'),
	},
	5425: {
		Fword:  __ccgo_ts + 37606,
		Ftype1: int8('F'),
	},
	5426: {
		Fword:  __ccgo_ts + 37613,
		Ftype1: int8('F'),
	},
	5427: {
		Fword:  __ccgo_ts + 37620,
		Ftype1: int8('F'),
	},
	5428: {
		Fword:  __ccgo_ts + 37627,
		Ftype1: int8('F'),
	},
	5429: {
		Fword:  __ccgo_ts + 37634,
		Ftype1: int8('F'),
	},
	5430: {
		Fword:  __ccgo_ts + 37641,
		Ftype1: int8('F'),
	},
	5431: {
		Fword:  __ccgo_ts + 37648,
		Ftype1: int8('F'),
	},
	5432: {
		Fword:  __ccgo_ts + 37655,
		Ftype1: int8('F'),
	},
	5433: {
		Fword:  __ccgo_ts + 37662,
		Ftype1: int8('F'),
	},
	5434: {
		Fword:  __ccgo_ts + 37669,
		Ftype1: int8('F'),
	},
	5435: {
		Fword:  __ccgo_ts + 37676,
		Ftype1: int8('F'),
	},
	5436: {
		Fword:  __ccgo_ts + 37683,
		Ftype1: int8('F'),
	},
	5437: {
		Fword:  __ccgo_ts + 37690,
		Ftype1: int8('F'),
	},
	5438: {
		Fword:  __ccgo_ts + 37697,
		Ftype1: int8('F'),
	},
	5439: {
		Fword:  __ccgo_ts + 37704,
		Ftype1: int8('F'),
	},
	5440: {
		Fword:  __ccgo_ts + 37711,
		Ftype1: int8('F'),
	},
	5441: {
		Fword:  __ccgo_ts + 37718,
		Ftype1: int8('F'),
	},
	5442: {
		Fword:  __ccgo_ts + 37725,
		Ftype1: int8('F'),
	},
	5443: {
		Fword:  __ccgo_ts + 37732,
		Ftype1: int8('F'),
	},
	5444: {
		Fword:  __ccgo_ts + 37739,
		Ftype1: int8('F'),
	},
	5445: {
		Fword:  __ccgo_ts + 37746,
		Ftype1: int8('F'),
	},
	5446: {
		Fword:  __ccgo_ts + 37753,
		Ftype1: int8('F'),
	},
	5447: {
		Fword:  __ccgo_ts + 37760,
		Ftype1: int8('F'),
	},
	5448: {
		Fword:  __ccgo_ts + 37767,
		Ftype1: int8('F'),
	},
	5449: {
		Fword:  __ccgo_ts + 37774,
		Ftype1: int8('F'),
	},
	5450: {
		Fword:  __ccgo_ts + 37781,
		Ftype1: int8('F'),
	},
	5451: {
		Fword:  __ccgo_ts + 37788,
		Ftype1: int8('F'),
	},
	5452: {
		Fword:  __ccgo_ts + 37795,
		Ftype1: int8('F'),
	},
	5453: {
		Fword:  __ccgo_ts + 37802,
		Ftype1: int8('F'),
	},
	5454: {
		Fword:  __ccgo_ts + 37809,
		Ftype1: int8('F'),
	},
	5455: {
		Fword:  __ccgo_ts + 37816,
		Ftype1: int8('F'),
	},
	5456: {
		Fword:  __ccgo_ts + 37823,
		Ftype1: int8('F'),
	},
	5457: {
		Fword:  __ccgo_ts + 37830,
		Ftype1: int8('F'),
	},
	5458: {
		Fword:  __ccgo_ts + 37837,
		Ftype1: int8('F'),
	},
	5459: {
		Fword:  __ccgo_ts + 37844,
		Ftype1: int8('F'),
	},
	5460: {
		Fword:  __ccgo_ts + 37851,
		Ftype1: int8('F'),
	},
	5461: {
		Fword:  __ccgo_ts + 37858,
		Ftype1: int8('F'),
	},
	5462: {
		Fword:  __ccgo_ts + 37865,
		Ftype1: int8('F'),
	},
	5463: {
		Fword:  __ccgo_ts + 37872,
		Ftype1: int8('F'),
	},
	5464: {
		Fword:  __ccgo_ts + 37879,
		Ftype1: int8('F'),
	},
	5465: {
		Fword:  __ccgo_ts + 37886,
		Ftype1: int8('F'),
	},
	5466: {
		Fword:  __ccgo_ts + 37893,
		Ftype1: int8('F'),
	},
	5467: {
		Fword:  __ccgo_ts + 37900,
		Ftype1: int8('F'),
	},
	5468: {
		Fword:  __ccgo_ts + 37907,
		Ftype1: int8('F'),
	},
	5469: {
		Fword:  __ccgo_ts + 37914,
		Ftype1: int8('F'),
	},
	5470: {
		Fword:  __ccgo_ts + 37921,
		Ftype1: int8('F'),
	},
	5471: {
		Fword:  __ccgo_ts + 37926,
		Ftype1: int8('F'),
	},
	5472: {
		Fword:  __ccgo_ts + 37933,
		Ftype1: int8('F'),
	},
	5473: {
		Fword:  __ccgo_ts + 37940,
		Ftype1: int8('F'),
	},
	5474: {
		Fword:  __ccgo_ts + 37947,
		Ftype1: int8('F'),
	},
	5475: {
		Fword:  __ccgo_ts + 37954,
		Ftype1: int8('F'),
	},
	5476: {
		Fword:  __ccgo_ts + 37961,
		Ftype1: int8('F'),
	},
	5477: {
		Fword:  __ccgo_ts + 37968,
		Ftype1: int8('F'),
	},
	5478: {
		Fword:  __ccgo_ts + 37975,
		Ftype1: int8('F'),
	},
	5479: {
		Fword:  __ccgo_ts + 37982,
		Ftype1: int8('F'),
	},
	5480: {
		Fword:  __ccgo_ts + 37988,
		Ftype1: int8('F'),
	},
	5481: {
		Fword:  __ccgo_ts + 37995,
		Ftype1: int8('F'),
	},
	5482: {
		Fword:  __ccgo_ts + 38002,
		Ftype1: int8('F'),
	},
	5483: {
		Fword:  __ccgo_ts + 38009,
		Ftype1: int8('F'),
	},
	5484: {
		Fword:  __ccgo_ts + 38016,
		Ftype1: int8('F'),
	},
	5485: {
		Fword:  __ccgo_ts + 38023,
		Ftype1: int8('F'),
	},
	5486: {
		Fword:  __ccgo_ts + 38030,
		Ftype1: int8('F'),
	},
	5487: {
		Fword:  __ccgo_ts + 38037,
		Ftype1: int8('F'),
	},
	5488: {
		Fword:  __ccgo_ts + 38043,
		Ftype1: int8('F'),
	},
	5489: {
		Fword:  __ccgo_ts + 38050,
		Ftype1: int8('F'),
	},
	5490: {
		Fword:  __ccgo_ts + 38057,
		Ftype1: int8('F'),
	},
	5491: {
		Fword:  __ccgo_ts + 38064,
		Ftype1: int8('F'),
	},
	5492: {
		Fword:  __ccgo_ts + 38071,
		Ftype1: int8('F'),
	},
	5493: {
		Fword:  __ccgo_ts + 38078,
		Ftype1: int8('F'),
	},
	5494: {
		Fword:  __ccgo_ts + 38085,
		Ftype1: int8('F'),
	},
	5495: {
		Fword:  __ccgo_ts + 38092,
		Ftype1: int8('F'),
	},
	5496: {
		Fword:  __ccgo_ts + 38099,
		Ftype1: int8('F'),
	},
	5497: {
		Fword:  __ccgo_ts + 38106,
		Ftype1: int8('F'),
	},
	5498: {
		Fword:  __ccgo_ts + 38113,
		Ftype1: int8('F'),
	},
	5499: {
		Fword:  __ccgo_ts + 38120,
		Ftype1: int8('F'),
	},
	5500: {
		Fword:  __ccgo_ts + 38127,
		Ftype1: int8('F'),
	},
	5501: {
		Fword:  __ccgo_ts + 38134,
		Ftype1: int8('F'),
	},
	5502: {
		Fword:  __ccgo_ts + 38141,
		Ftype1: int8('F'),
	},
	5503: {
		Fword:  __ccgo_ts + 38148,
		Ftype1: int8('F'),
	},
	5504: {
		Fword:  __ccgo_ts + 38155,
		Ftype1: int8('F'),
	},
	5505: {
		Fword:  __ccgo_ts + 38162,
		Ftype1: int8('F'),
	},
	5506: {
		Fword:  __ccgo_ts + 38169,
		Ftype1: int8('F'),
	},
	5507: {
		Fword:  __ccgo_ts + 38176,
		Ftype1: int8('F'),
	},
	5508: {
		Fword:  __ccgo_ts + 38183,
		Ftype1: int8('F'),
	},
	5509: {
		Fword:  __ccgo_ts + 38190,
		Ftype1: int8('F'),
	},
	5510: {
		Fword:  __ccgo_ts + 38197,
		Ftype1: int8('F'),
	},
	5511: {
		Fword:  __ccgo_ts + 38204,
		Ftype1: int8('F'),
	},
	5512: {
		Fword:  __ccgo_ts + 38211,
		Ftype1: int8('F'),
	},
	5513: {
		Fword:  __ccgo_ts + 38216,
		Ftype1: int8('F'),
	},
	5514: {
		Fword:  __ccgo_ts + 38223,
		Ftype1: int8('F'),
	},
	5515: {
		Fword:  __ccgo_ts + 38230,
		Ftype1: int8('F'),
	},
	5516: {
		Fword:  __ccgo_ts + 38237,
		Ftype1: int8('F'),
	},
	5517: {
		Fword:  __ccgo_ts + 38244,
		Ftype1: int8('F'),
	},
	5518: {
		Fword:  __ccgo_ts + 38251,
		Ftype1: int8('F'),
	},
	5519: {
		Fword:  __ccgo_ts + 38258,
		Ftype1: int8('F'),
	},
	5520: {
		Fword:  __ccgo_ts + 38265,
		Ftype1: int8('F'),
	},
	5521: {
		Fword:  __ccgo_ts + 38272,
		Ftype1: int8('F'),
	},
	5522: {
		Fword:  __ccgo_ts + 38278,
		Ftype1: int8('F'),
	},
	5523: {
		Fword:  __ccgo_ts + 38285,
		Ftype1: int8('F'),
	},
	5524: {
		Fword:  __ccgo_ts + 38292,
		Ftype1: int8('F'),
	},
	5525: {
		Fword:  __ccgo_ts + 38299,
		Ftype1: int8('F'),
	},
	5526: {
		Fword:  __ccgo_ts + 38306,
		Ftype1: int8('F'),
	},
	5527: {
		Fword:  __ccgo_ts + 38313,
		Ftype1: int8('F'),
	},
	5528: {
		Fword:  __ccgo_ts + 38320,
		Ftype1: int8('F'),
	},
	5529: {
		Fword:  __ccgo_ts + 38327,
		Ftype1: int8('F'),
	},
	5530: {
		Fword:  __ccgo_ts + 38333,
		Ftype1: int8('F'),
	},
	5531: {
		Fword:  __ccgo_ts + 38340,
		Ftype1: int8('F'),
	},
	5532: {
		Fword:  __ccgo_ts + 38347,
		Ftype1: int8('F'),
	},
	5533: {
		Fword:  __ccgo_ts + 38354,
		Ftype1: int8('F'),
	},
	5534: {
		Fword:  __ccgo_ts + 38361,
		Ftype1: int8('F'),
	},
	5535: {
		Fword:  __ccgo_ts + 38368,
		Ftype1: int8('F'),
	},
	5536: {
		Fword:  __ccgo_ts + 38375,
		Ftype1: int8('F'),
	},
	5537: {
		Fword:  __ccgo_ts + 38382,
		Ftype1: int8('F'),
	},
	5538: {
		Fword:  __ccgo_ts + 38389,
		Ftype1: int8('F'),
	},
	5539: {
		Fword:  __ccgo_ts + 38396,
		Ftype1: int8('F'),
	},
	5540: {
		Fword:  __ccgo_ts + 38403,
		Ftype1: int8('F'),
	},
	5541: {
		Fword:  __ccgo_ts + 38410,
		Ftype1: int8('F'),
	},
	5542: {
		Fword:  __ccgo_ts + 38417,
		Ftype1: int8('F'),
	},
	5543: {
		Fword:  __ccgo_ts + 38422,
		Ftype1: int8('F'),
	},
	5544: {
		Fword:  __ccgo_ts + 38429,
		Ftype1: int8('F'),
	},
	5545: {
		Fword:  __ccgo_ts + 38436,
		Ftype1: int8('F'),
	},
	5546: {
		Fword:  __ccgo_ts + 38443,
		Ftype1: int8('F'),
	},
	5547: {
		Fword:  __ccgo_ts + 38450,
		Ftype1: int8('F'),
	},
	5548: {
		Fword:  __ccgo_ts + 38457,
		Ftype1: int8('F'),
	},
	5549: {
		Fword:  __ccgo_ts + 38464,
		Ftype1: int8('F'),
	},
	5550: {
		Fword:  __ccgo_ts + 38471,
		Ftype1: int8('F'),
	},
	5551: {
		Fword:  __ccgo_ts + 38478,
		Ftype1: int8('F'),
	},
	5552: {
		Fword:  __ccgo_ts + 38484,
		Ftype1: int8('F'),
	},
	5553: {
		Fword:  __ccgo_ts + 38491,
		Ftype1: int8('F'),
	},
	5554: {
		Fword:  __ccgo_ts + 38498,
		Ftype1: int8('F'),
	},
	5555: {
		Fword:  __ccgo_ts + 38505,
		Ftype1: int8('F'),
	},
	5556: {
		Fword:  __ccgo_ts + 38512,
		Ftype1: int8('F'),
	},
	5557: {
		Fword:  __ccgo_ts + 38519,
		Ftype1: int8('F'),
	},
	5558: {
		Fword:  __ccgo_ts + 38526,
		Ftype1: int8('F'),
	},
	5559: {
		Fword:  __ccgo_ts + 38533,
		Ftype1: int8('F'),
	},
	5560: {
		Fword:  __ccgo_ts + 38539,
		Ftype1: int8('F'),
	},
	5561: {
		Fword:  __ccgo_ts + 38546,
		Ftype1: int8('F'),
	},
	5562: {
		Fword:  __ccgo_ts + 38553,
		Ftype1: int8('F'),
	},
	5563: {
		Fword:  __ccgo_ts + 38560,
		Ftype1: int8('F'),
	},
	5564: {
		Fword:  __ccgo_ts + 38567,
		Ftype1: int8('F'),
	},
	5565: {
		Fword:  __ccgo_ts + 38574,
		Ftype1: int8('F'),
	},
	5566: {
		Fword:  __ccgo_ts + 38581,
		Ftype1: int8('F'),
	},
	5567: {
		Fword:  __ccgo_ts + 38588,
		Ftype1: int8('F'),
	},
	5568: {
		Fword:  __ccgo_ts + 38595,
		Ftype1: int8('F'),
	},
	5569: {
		Fword:  __ccgo_ts + 38602,
		Ftype1: int8('F'),
	},
	5570: {
		Fword:  __ccgo_ts + 38609,
		Ftype1: int8('F'),
	},
	5571: {
		Fword:  __ccgo_ts + 38616,
		Ftype1: int8('F'),
	},
	5572: {
		Fword:  __ccgo_ts + 38623,
		Ftype1: int8('F'),
	},
	5573: {
		Fword:  __ccgo_ts + 38630,
		Ftype1: int8('F'),
	},
	5574: {
		Fword:  __ccgo_ts + 38637,
		Ftype1: int8('F'),
	},
	5575: {
		Fword:  __ccgo_ts + 38642,
		Ftype1: int8('F'),
	},
	5576: {
		Fword:  __ccgo_ts + 38649,
		Ftype1: int8('F'),
	},
	5577: {
		Fword:  __ccgo_ts + 38656,
		Ftype1: int8('F'),
	},
	5578: {
		Fword:  __ccgo_ts + 38663,
		Ftype1: int8('F'),
	},
	5579: {
		Fword:  __ccgo_ts + 38670,
		Ftype1: int8('F'),
	},
	5580: {
		Fword:  __ccgo_ts + 38677,
		Ftype1: int8('F'),
	},
	5581: {
		Fword:  __ccgo_ts + 38684,
		Ftype1: int8('F'),
	},
	5582: {
		Fword:  __ccgo_ts + 38691,
		Ftype1: int8('F'),
	},
	5583: {
		Fword:  __ccgo_ts + 38698,
		Ftype1: int8('F'),
	},
	5584: {
		Fword:  __ccgo_ts + 38704,
		Ftype1: int8('F'),
	},
	5585: {
		Fword:  __ccgo_ts + 38711,
		Ftype1: int8('F'),
	},
	5586: {
		Fword:  __ccgo_ts + 38718,
		Ftype1: int8('F'),
	},
	5587: {
		Fword:  __ccgo_ts + 38725,
		Ftype1: int8('F'),
	},
	5588: {
		Fword:  __ccgo_ts + 38732,
		Ftype1: int8('F'),
	},
	5589: {
		Fword:  __ccgo_ts + 38739,
		Ftype1: int8('F'),
	},
	5590: {
		Fword:  __ccgo_ts + 38746,
		Ftype1: int8('F'),
	},
	5591: {
		Fword:  __ccgo_ts + 38753,
		Ftype1: int8('F'),
	},
	5592: {
		Fword:  __ccgo_ts + 38759,
		Ftype1: int8('F'),
	},
	5593: {
		Fword:  __ccgo_ts + 38766,
		Ftype1: int8('F'),
	},
	5594: {
		Fword:  __ccgo_ts + 38773,
		Ftype1: int8('F'),
	},
	5595: {
		Fword:  __ccgo_ts + 38780,
		Ftype1: int8('F'),
	},
	5596: {
		Fword:  __ccgo_ts + 38787,
		Ftype1: int8('F'),
	},
	5597: {
		Fword:  __ccgo_ts + 38794,
		Ftype1: int8('F'),
	},
	5598: {
		Fword:  __ccgo_ts + 38801,
		Ftype1: int8('F'),
	},
	5599: {
		Fword:  __ccgo_ts + 38808,
		Ftype1: int8('F'),
	},
	5600: {
		Fword:  __ccgo_ts + 38815,
		Ftype1: int8('F'),
	},
	5601: {
		Fword:  __ccgo_ts + 38822,
		Ftype1: int8('F'),
	},
	5602: {
		Fword:  __ccgo_ts + 38829,
		Ftype1: int8('F'),
	},
	5603: {
		Fword:  __ccgo_ts + 38836,
		Ftype1: int8('F'),
	},
	5604: {
		Fword:  __ccgo_ts + 38840,
		Ftype1: int8('F'),
	},
	5605: {
		Fword:  __ccgo_ts + 38847,
		Ftype1: int8('F'),
	},
	5606: {
		Fword:  __ccgo_ts + 38854,
		Ftype1: int8('F'),
	},
	5607: {
		Fword:  __ccgo_ts + 38861,
		Ftype1: int8('F'),
	},
	5608: {
		Fword:  __ccgo_ts + 38868,
		Ftype1: int8('F'),
	},
	5609: {
		Fword:  __ccgo_ts + 38875,
		Ftype1: int8('F'),
	},
	5610: {
		Fword:  __ccgo_ts + 38882,
		Ftype1: int8('F'),
	},
	5611: {
		Fword:  __ccgo_ts + 38889,
		Ftype1: int8('F'),
	},
	5612: {
		Fword:  __ccgo_ts + 38896,
		Ftype1: int8('F'),
	},
	5613: {
		Fword:  __ccgo_ts + 38903,
		Ftype1: int8('F'),
	},
	5614: {
		Fword:  __ccgo_ts + 38910,
		Ftype1: int8('F'),
	},
	5615: {
		Fword:  __ccgo_ts + 38916,
		Ftype1: int8('F'),
	},
	5616: {
		Fword:  __ccgo_ts + 38923,
		Ftype1: int8('F'),
	},
	5617: {
		Fword:  __ccgo_ts + 38930,
		Ftype1: int8('F'),
	},
	5618: {
		Fword:  __ccgo_ts + 38937,
		Ftype1: int8('F'),
	},
	5619: {
		Fword:  __ccgo_ts + 38944,
		Ftype1: int8('F'),
	},
	5620: {
		Fword:  __ccgo_ts + 38951,
		Ftype1: int8('F'),
	},
	5621: {
		Fword:  __ccgo_ts + 38958,
		Ftype1: int8('F'),
	},
	5622: {
		Fword:  __ccgo_ts + 38965,
		Ftype1: int8('F'),
	},
	5623: {
		Fword:  __ccgo_ts + 38972,
		Ftype1: int8('F'),
	},
	5624: {
		Fword:  __ccgo_ts + 38979,
		Ftype1: int8('F'),
	},
	5625: {
		Fword:  __ccgo_ts + 38986,
		Ftype1: int8('F'),
	},
	5626: {
		Fword:  __ccgo_ts + 38993,
		Ftype1: int8('F'),
	},
	5627: {
		Fword:  __ccgo_ts + 39000,
		Ftype1: int8('F'),
	},
	5628: {
		Fword:  __ccgo_ts + 39007,
		Ftype1: int8('F'),
	},
	5629: {
		Fword:  __ccgo_ts + 39014,
		Ftype1: int8('F'),
	},
	5630: {
		Fword:  __ccgo_ts + 39021,
		Ftype1: int8('F'),
	},
	5631: {
		Fword:  __ccgo_ts + 39028,
		Ftype1: int8('F'),
	},
	5632: {
		Fword:  __ccgo_ts + 39035,
		Ftype1: int8('F'),
	},
	5633: {
		Fword:  __ccgo_ts + 39042,
		Ftype1: int8('F'),
	},
	5634: {
		Fword:  __ccgo_ts + 39049,
		Ftype1: int8('F'),
	},
	5635: {
		Fword:  __ccgo_ts + 39056,
		Ftype1: int8('F'),
	},
	5636: {
		Fword:  __ccgo_ts + 39063,
		Ftype1: int8('F'),
	},
	5637: {
		Fword:  __ccgo_ts + 39070,
		Ftype1: int8('F'),
	},
	5638: {
		Fword:  __ccgo_ts + 39077,
		Ftype1: int8('F'),
	},
	5639: {
		Fword:  __ccgo_ts + 39084,
		Ftype1: int8('F'),
	},
	5640: {
		Fword:  __ccgo_ts + 39091,
		Ftype1: int8('F'),
	},
	5641: {
		Fword:  __ccgo_ts + 39098,
		Ftype1: int8('F'),
	},
	5642: {
		Fword:  __ccgo_ts + 39105,
		Ftype1: int8('F'),
	},
	5643: {
		Fword:  __ccgo_ts + 39112,
		Ftype1: int8('F'),
	},
	5644: {
		Fword:  __ccgo_ts + 39119,
		Ftype1: int8('F'),
	},
	5645: {
		Fword:  __ccgo_ts + 39126,
		Ftype1: int8('F'),
	},
	5646: {
		Fword:  __ccgo_ts + 39133,
		Ftype1: int8('F'),
	},
	5647: {
		Fword:  __ccgo_ts + 39140,
		Ftype1: int8('F'),
	},
	5648: {
		Fword:  __ccgo_ts + 39147,
		Ftype1: int8('F'),
	},
	5649: {
		Fword:  __ccgo_ts + 39154,
		Ftype1: int8('F'),
	},
	5650: {
		Fword:  __ccgo_ts + 39161,
		Ftype1: int8('F'),
	},
	5651: {
		Fword:  __ccgo_ts + 39168,
		Ftype1: int8('F'),
	},
	5652: {
		Fword:  __ccgo_ts + 39175,
		Ftype1: int8('F'),
	},
	5653: {
		Fword:  __ccgo_ts + 39182,
		Ftype1: int8('F'),
	},
	5654: {
		Fword:  __ccgo_ts + 39189,
		Ftype1: int8('F'),
	},
	5655: {
		Fword:  __ccgo_ts + 39196,
		Ftype1: int8('F'),
	},
	5656: {
		Fword:  __ccgo_ts + 39203,
		Ftype1: int8('F'),
	},
	5657: {
		Fword:  __ccgo_ts + 39210,
		Ftype1: int8('F'),
	},
	5658: {
		Fword:  __ccgo_ts + 39217,
		Ftype1: int8('F'),
	},
	5659: {
		Fword:  __ccgo_ts + 39224,
		Ftype1: int8('F'),
	},
	5660: {
		Fword:  __ccgo_ts + 39231,
		Ftype1: int8('F'),
	},
	5661: {
		Fword:  __ccgo_ts + 39238,
		Ftype1: int8('F'),
	},
	5662: {
		Fword:  __ccgo_ts + 39245,
		Ftype1: int8('F'),
	},
	5663: {
		Fword:  __ccgo_ts + 39252,
		Ftype1: int8('F'),
	},
	5664: {
		Fword:  __ccgo_ts + 39259,
		Ftype1: int8('F'),
	},
	5665: {
		Fword:  __ccgo_ts + 39266,
		Ftype1: int8('F'),
	},
	5666: {
		Fword:  __ccgo_ts + 39273,
		Ftype1: int8('F'),
	},
	5667: {
		Fword:  __ccgo_ts + 39280,
		Ftype1: int8('F'),
	},
	5668: {
		Fword:  __ccgo_ts + 39286,
		Ftype1: int8('F'),
	},
	5669: {
		Fword:  __ccgo_ts + 39293,
		Ftype1: int8('F'),
	},
	5670: {
		Fword:  __ccgo_ts + 39300,
		Ftype1: int8('F'),
	},
	5671: {
		Fword:  __ccgo_ts + 39307,
		Ftype1: int8('F'),
	},
	5672: {
		Fword:  __ccgo_ts + 39314,
		Ftype1: int8('F'),
	},
	5673: {
		Fword:  __ccgo_ts + 39321,
		Ftype1: int8('F'),
	},
	5674: {
		Fword:  __ccgo_ts + 39328,
		Ftype1: int8('F'),
	},
	5675: {
		Fword:  __ccgo_ts + 39335,
		Ftype1: int8('F'),
	},
	5676: {
		Fword:  __ccgo_ts + 39342,
		Ftype1: int8('F'),
	},
	5677: {
		Fword:  __ccgo_ts + 39349,
		Ftype1: int8('F'),
	},
	5678: {
		Fword:  __ccgo_ts + 39356,
		Ftype1: int8('F'),
	},
	5679: {
		Fword:  __ccgo_ts + 39363,
		Ftype1: int8('F'),
	},
	5680: {
		Fword:  __ccgo_ts + 39370,
		Ftype1: int8('F'),
	},
	5681: {
		Fword:  __ccgo_ts + 39377,
		Ftype1: int8('F'),
	},
	5682: {
		Fword:  __ccgo_ts + 39384,
		Ftype1: int8('F'),
	},
	5683: {
		Fword:  __ccgo_ts + 39390,
		Ftype1: int8('F'),
	},
	5684: {
		Fword:  __ccgo_ts + 39397,
		Ftype1: int8('F'),
	},
	5685: {
		Fword:  __ccgo_ts + 39404,
		Ftype1: int8('F'),
	},
	5686: {
		Fword:  __ccgo_ts + 39411,
		Ftype1: int8('F'),
	},
	5687: {
		Fword:  __ccgo_ts + 39418,
		Ftype1: int8('F'),
	},
	5688: {
		Fword:  __ccgo_ts + 39425,
		Ftype1: int8('F'),
	},
	5689: {
		Fword:  __ccgo_ts + 39432,
		Ftype1: int8('F'),
	},
	5690: {
		Fword:  __ccgo_ts + 39439,
		Ftype1: int8('F'),
	},
	5691: {
		Fword:  __ccgo_ts + 39446,
		Ftype1: int8('F'),
	},
	5692: {
		Fword:  __ccgo_ts + 39453,
		Ftype1: int8('F'),
	},
	5693: {
		Fword:  __ccgo_ts + 39460,
		Ftype1: int8('F'),
	},
	5694: {
		Fword:  __ccgo_ts + 39467,
		Ftype1: int8('F'),
	},
	5695: {
		Fword:  __ccgo_ts + 39474,
		Ftype1: int8('F'),
	},
	5696: {
		Fword:  __ccgo_ts + 39481,
		Ftype1: int8('F'),
	},
	5697: {
		Fword:  __ccgo_ts + 39488,
		Ftype1: int8('F'),
	},
	5698: {
		Fword:  __ccgo_ts + 39495,
		Ftype1: int8('F'),
	},
	5699: {
		Fword:  __ccgo_ts + 39502,
		Ftype1: int8('F'),
	},
	5700: {
		Fword:  __ccgo_ts + 39509,
		Ftype1: int8('F'),
	},
	5701: {
		Fword:  __ccgo_ts + 39516,
		Ftype1: int8('F'),
	},
	5702: {
		Fword:  __ccgo_ts + 39523,
		Ftype1: int8('F'),
	},
	5703: {
		Fword:  __ccgo_ts + 39530,
		Ftype1: int8('F'),
	},
	5704: {
		Fword:  __ccgo_ts + 39537,
		Ftype1: int8('F'),
	},
	5705: {
		Fword:  __ccgo_ts + 39544,
		Ftype1: int8('F'),
	},
	5706: {
		Fword:  __ccgo_ts + 39551,
		Ftype1: int8('F'),
	},
	5707: {
		Fword:  __ccgo_ts + 39558,
		Ftype1: int8('F'),
	},
	5708: {
		Fword:  __ccgo_ts + 39565,
		Ftype1: int8('F'),
	},
	5709: {
		Fword:  __ccgo_ts + 39572,
		Ftype1: int8('F'),
	},
	5710: {
		Fword:  __ccgo_ts + 39579,
		Ftype1: int8('F'),
	},
	5711: {
		Fword:  __ccgo_ts + 39586,
		Ftype1: int8('F'),
	},
	5712: {
		Fword:  __ccgo_ts + 39593,
		Ftype1: int8('F'),
	},
	5713: {
		Fword:  __ccgo_ts + 39600,
		Ftype1: int8('F'),
	},
	5714: {
		Fword:  __ccgo_ts + 39607,
		Ftype1: int8('F'),
	},
	5715: {
		Fword:  __ccgo_ts + 39613,
		Ftype1: int8('F'),
	},
	5716: {
		Fword:  __ccgo_ts + 39620,
		Ftype1: int8('F'),
	},
	5717: {
		Fword:  __ccgo_ts + 39627,
		Ftype1: int8('F'),
	},
	5718: {
		Fword:  __ccgo_ts + 39634,
		Ftype1: int8('F'),
	},
	5719: {
		Fword:  __ccgo_ts + 39641,
		Ftype1: int8('F'),
	},
	5720: {
		Fword:  __ccgo_ts + 39648,
		Ftype1: int8('F'),
	},
	5721: {
		Fword:  __ccgo_ts + 39655,
		Ftype1: int8('F'),
	},
	5722: {
		Fword:  __ccgo_ts + 39662,
		Ftype1: int8('F'),
	},
	5723: {
		Fword:  __ccgo_ts + 39669,
		Ftype1: int8('F'),
	},
	5724: {
		Fword:  __ccgo_ts + 39676,
		Ftype1: int8('F'),
	},
	5725: {
		Fword:  __ccgo_ts + 39683,
		Ftype1: int8('F'),
	},
	5726: {
		Fword:  __ccgo_ts + 39690,
		Ftype1: int8('F'),
	},
	5727: {
		Fword:  __ccgo_ts + 39697,
		Ftype1: int8('F'),
	},
	5728: {
		Fword:  __ccgo_ts + 39704,
		Ftype1: int8('F'),
	},
	5729: {
		Fword:  __ccgo_ts + 39711,
		Ftype1: int8('F'),
	},
	5730: {
		Fword:  __ccgo_ts + 39718,
		Ftype1: int8('F'),
	},
	5731: {
		Fword:  __ccgo_ts + 39725,
		Ftype1: int8('F'),
	},
	5732: {
		Fword:  __ccgo_ts + 39732,
		Ftype1: int8('F'),
	},
	5733: {
		Fword:  __ccgo_ts + 39739,
		Ftype1: int8('F'),
	},
	5734: {
		Fword:  __ccgo_ts + 39746,
		Ftype1: int8('F'),
	},
	5735: {
		Fword:  __ccgo_ts + 39753,
		Ftype1: int8('F'),
	},
	5736: {
		Fword:  __ccgo_ts + 39760,
		Ftype1: int8('F'),
	},
	5737: {
		Fword:  __ccgo_ts + 39767,
		Ftype1: int8('F'),
	},
	5738: {
		Fword:  __ccgo_ts + 39774,
		Ftype1: int8('F'),
	},
	5739: {
		Fword:  __ccgo_ts + 39781,
		Ftype1: int8('F'),
	},
	5740: {
		Fword:  __ccgo_ts + 39788,
		Ftype1: int8('F'),
	},
	5741: {
		Fword:  __ccgo_ts + 39795,
		Ftype1: int8('F'),
	},
	5742: {
		Fword:  __ccgo_ts + 39802,
		Ftype1: int8('F'),
	},
	5743: {
		Fword:  __ccgo_ts + 39809,
		Ftype1: int8('F'),
	},
	5744: {
		Fword:  __ccgo_ts + 39816,
		Ftype1: int8('F'),
	},
	5745: {
		Fword:  __ccgo_ts + 39823,
		Ftype1: int8('F'),
	},
	5746: {
		Fword:  __ccgo_ts + 39830,
		Ftype1: int8('F'),
	},
	5747: {
		Fword:  __ccgo_ts + 39837,
		Ftype1: int8('F'),
	},
	5748: {
		Fword:  __ccgo_ts + 39844,
		Ftype1: int8('F'),
	},
	5749: {
		Fword:  __ccgo_ts + 39851,
		Ftype1: int8('F'),
	},
	5750: {
		Fword:  __ccgo_ts + 39858,
		Ftype1: int8('F'),
	},
	5751: {
		Fword:  __ccgo_ts + 39865,
		Ftype1: int8('F'),
	},
	5752: {
		Fword:  __ccgo_ts + 39872,
		Ftype1: int8('F'),
	},
	5753: {
		Fword:  __ccgo_ts + 39879,
		Ftype1: int8('F'),
	},
	5754: {
		Fword:  __ccgo_ts + 39886,
		Ftype1: int8('F'),
	},
	5755: {
		Fword:  __ccgo_ts + 39893,
		Ftype1: int8('F'),
	},
	5756: {
		Fword:  __ccgo_ts + 39900,
		Ftype1: int8('F'),
	},
	5757: {
		Fword:  __ccgo_ts + 39907,
		Ftype1: int8('F'),
	},
	5758: {
		Fword:  __ccgo_ts + 39914,
		Ftype1: int8('F'),
	},
	5759: {
		Fword:  __ccgo_ts + 39921,
		Ftype1: int8('F'),
	},
	5760: {
		Fword:  __ccgo_ts + 39928,
		Ftype1: int8('F'),
	},
	5761: {
		Fword:  __ccgo_ts + 39935,
		Ftype1: int8('F'),
	},
	5762: {
		Fword:  __ccgo_ts + 39942,
		Ftype1: int8('F'),
	},
	5763: {
		Fword:  __ccgo_ts + 39949,
		Ftype1: int8('F'),
	},
	5764: {
		Fword:  __ccgo_ts + 39956,
		Ftype1: int8('F'),
	},
	5765: {
		Fword:  __ccgo_ts + 39963,
		Ftype1: int8('F'),
	},
	5766: {
		Fword:  __ccgo_ts + 39970,
		Ftype1: int8('F'),
	},
	5767: {
		Fword:  __ccgo_ts + 39977,
		Ftype1: int8('F'),
	},
	5768: {
		Fword:  __ccgo_ts + 39984,
		Ftype1: int8('F'),
	},
	5769: {
		Fword:  __ccgo_ts + 39991,
		Ftype1: int8('F'),
	},
	5770: {
		Fword:  __ccgo_ts + 39998,
		Ftype1: int8('F'),
	},
	5771: {
		Fword:  __ccgo_ts + 40005,
		Ftype1: int8('F'),
	},
	5772: {
		Fword:  __ccgo_ts + 40012,
		Ftype1: int8('F'),
	},
	5773: {
		Fword:  __ccgo_ts + 40019,
		Ftype1: int8('F'),
	},
	5774: {
		Fword:  __ccgo_ts + 40026,
		Ftype1: int8('F'),
	},
	5775: {
		Fword:  __ccgo_ts + 40033,
		Ftype1: int8('F'),
	},
	5776: {
		Fword:  __ccgo_ts + 40040,
		Ftype1: int8('F'),
	},
	5777: {
		Fword:  __ccgo_ts + 40047,
		Ftype1: int8('F'),
	},
	5778: {
		Fword:  __ccgo_ts + 40054,
		Ftype1: int8('F'),
	},
	5779: {
		Fword:  __ccgo_ts + 40061,
		Ftype1: int8('F'),
	},
	5780: {
		Fword:  __ccgo_ts + 40068,
		Ftype1: int8('F'),
	},
	5781: {
		Fword:  __ccgo_ts + 40075,
		Ftype1: int8('F'),
	},
	5782: {
		Fword:  __ccgo_ts + 40082,
		Ftype1: int8('F'),
	},
	5783: {
		Fword:  __ccgo_ts + 40089,
		Ftype1: int8('F'),
	},
	5784: {
		Fword:  __ccgo_ts + 40096,
		Ftype1: int8('F'),
	},
	5785: {
		Fword:  __ccgo_ts + 40103,
		Ftype1: int8('F'),
	},
	5786: {
		Fword:  __ccgo_ts + 40108,
		Ftype1: int8('F'),
	},
	5787: {
		Fword:  __ccgo_ts + 40115,
		Ftype1: int8('F'),
	},
	5788: {
		Fword:  __ccgo_ts + 40122,
		Ftype1: int8('F'),
	},
	5789: {
		Fword:  __ccgo_ts + 40129,
		Ftype1: int8('F'),
	},
	5790: {
		Fword:  __ccgo_ts + 40136,
		Ftype1: int8('F'),
	},
	5791: {
		Fword:  __ccgo_ts + 40143,
		Ftype1: int8('F'),
	},
	5792: {
		Fword:  __ccgo_ts + 40150,
		Ftype1: int8('F'),
	},
	5793: {
		Fword:  __ccgo_ts + 40156,
		Ftype1: int8('F'),
	},
	5794: {
		Fword:  __ccgo_ts + 40163,
		Ftype1: int8('F'),
	},
	5795: {
		Fword:  __ccgo_ts + 40170,
		Ftype1: int8('F'),
	},
	5796: {
		Fword:  __ccgo_ts + 40177,
		Ftype1: int8('F'),
	},
	5797: {
		Fword:  __ccgo_ts + 40184,
		Ftype1: int8('F'),
	},
	5798: {
		Fword:  __ccgo_ts + 40191,
		Ftype1: int8('F'),
	},
	5799: {
		Fword:  __ccgo_ts + 40198,
		Ftype1: int8('F'),
	},
	5800: {
		Fword:  __ccgo_ts + 40205,
		Ftype1: int8('F'),
	},
	5801: {
		Fword:  __ccgo_ts + 40212,
		Ftype1: int8('F'),
	},
	5802: {
		Fword:  __ccgo_ts + 40219,
		Ftype1: int8('F'),
	},
	5803: {
		Fword:  __ccgo_ts + 40225,
		Ftype1: int8('F'),
	},
	5804: {
		Fword:  __ccgo_ts + 40232,
		Ftype1: int8('F'),
	},
	5805: {
		Fword:  __ccgo_ts + 40239,
		Ftype1: int8('F'),
	},
	5806: {
		Fword:  __ccgo_ts + 40246,
		Ftype1: int8('F'),
	},
	5807: {
		Fword:  __ccgo_ts + 40253,
		Ftype1: int8('F'),
	},
	5808: {
		Fword:  __ccgo_ts + 40260,
		Ftype1: int8('F'),
	},
	5809: {
		Fword:  __ccgo_ts + 40267,
		Ftype1: int8('F'),
	},
	5810: {
		Fword:  __ccgo_ts + 40274,
		Ftype1: int8('F'),
	},
	5811: {
		Fword:  __ccgo_ts + 40281,
		Ftype1: int8('F'),
	},
	5812: {
		Fword:  __ccgo_ts + 40288,
		Ftype1: int8('F'),
	},
	5813: {
		Fword:  __ccgo_ts + 40295,
		Ftype1: int8('F'),
	},
	5814: {
		Fword:  __ccgo_ts + 40302,
		Ftype1: int8('F'),
	},
	5815: {
		Fword:  __ccgo_ts + 40309,
		Ftype1: int8('F'),
	},
	5816: {
		Fword:  __ccgo_ts + 40316,
		Ftype1: int8('F'),
	},
	5817: {
		Fword:  __ccgo_ts + 40323,
		Ftype1: int8('F'),
	},
	5818: {
		Fword:  __ccgo_ts + 40330,
		Ftype1: int8('F'),
	},
	5819: {
		Fword:  __ccgo_ts + 40337,
		Ftype1: int8('F'),
	},
	5820: {
		Fword:  __ccgo_ts + 40344,
		Ftype1: int8('F'),
	},
	5821: {
		Fword:  __ccgo_ts + 40351,
		Ftype1: int8('F'),
	},
	5822: {
		Fword:  __ccgo_ts + 40358,
		Ftype1: int8('F'),
	},
	5823: {
		Fword:  __ccgo_ts + 40363,
		Ftype1: int8('F'),
	},
	5824: {
		Fword:  __ccgo_ts + 40370,
		Ftype1: int8('F'),
	},
	5825: {
		Fword:  __ccgo_ts + 40377,
		Ftype1: int8('F'),
	},
	5826: {
		Fword:  __ccgo_ts + 40384,
		Ftype1: int8('F'),
	},
	5827: {
		Fword:  __ccgo_ts + 40391,
		Ftype1: int8('F'),
	},
	5828: {
		Fword:  __ccgo_ts + 40398,
		Ftype1: int8('F'),
	},
	5829: {
		Fword:  __ccgo_ts + 40405,
		Ftype1: int8('F'),
	},
	5830: {
		Fword:  __ccgo_ts + 40411,
		Ftype1: int8('F'),
	},
	5831: {
		Fword:  __ccgo_ts + 40418,
		Ftype1: int8('F'),
	},
	5832: {
		Fword:  __ccgo_ts + 40425,
		Ftype1: int8('F'),
	},
	5833: {
		Fword:  __ccgo_ts + 40432,
		Ftype1: int8('F'),
	},
	5834: {
		Fword:  __ccgo_ts + 40439,
		Ftype1: int8('F'),
	},
	5835: {
		Fword:  __ccgo_ts + 40446,
		Ftype1: int8('F'),
	},
	5836: {
		Fword:  __ccgo_ts + 40453,
		Ftype1: int8('F'),
	},
	5837: {
		Fword:  __ccgo_ts + 40460,
		Ftype1: int8('F'),
	},
	5838: {
		Fword:  __ccgo_ts + 40467,
		Ftype1: int8('F'),
	},
	5839: {
		Fword:  __ccgo_ts + 40474,
		Ftype1: int8('F'),
	},
	5840: {
		Fword:  __ccgo_ts + 40480,
		Ftype1: int8('F'),
	},
	5841: {
		Fword:  __ccgo_ts + 40487,
		Ftype1: int8('F'),
	},
	5842: {
		Fword:  __ccgo_ts + 40494,
		Ftype1: int8('F'),
	},
	5843: {
		Fword:  __ccgo_ts + 40501,
		Ftype1: int8('F'),
	},
	5844: {
		Fword:  __ccgo_ts + 40508,
		Ftype1: int8('F'),
	},
	5845: {
		Fword:  __ccgo_ts + 40515,
		Ftype1: int8('F'),
	},
	5846: {
		Fword:  __ccgo_ts + 40522,
		Ftype1: int8('F'),
	},
	5847: {
		Fword:  __ccgo_ts + 40529,
		Ftype1: int8('F'),
	},
	5848: {
		Fword:  __ccgo_ts + 40536,
		Ftype1: int8('F'),
	},
	5849: {
		Fword:  __ccgo_ts + 40541,
		Ftype1: int8('F'),
	},
	5850: {
		Fword:  __ccgo_ts + 40548,
		Ftype1: int8('F'),
	},
	5851: {
		Fword:  __ccgo_ts + 40555,
		Ftype1: int8('F'),
	},
	5852: {
		Fword:  __ccgo_ts + 40562,
		Ftype1: int8('F'),
	},
	5853: {
		Fword:  __ccgo_ts + 40569,
		Ftype1: int8('F'),
	},
	5854: {
		Fword:  __ccgo_ts + 40576,
		Ftype1: int8('F'),
	},
	5855: {
		Fword:  __ccgo_ts + 40583,
		Ftype1: int8('F'),
	},
	5856: {
		Fword:  __ccgo_ts + 40589,
		Ftype1: int8('F'),
	},
	5857: {
		Fword:  __ccgo_ts + 40596,
		Ftype1: int8('F'),
	},
	5858: {
		Fword:  __ccgo_ts + 40603,
		Ftype1: int8('F'),
	},
	5859: {
		Fword:  __ccgo_ts + 40610,
		Ftype1: int8('F'),
	},
	5860: {
		Fword:  __ccgo_ts + 40617,
		Ftype1: int8('F'),
	},
	5861: {
		Fword:  __ccgo_ts + 40624,
		Ftype1: int8('F'),
	},
	5862: {
		Fword:  __ccgo_ts + 40631,
		Ftype1: int8('F'),
	},
	5863: {
		Fword:  __ccgo_ts + 40638,
		Ftype1: int8('F'),
	},
	5864: {
		Fword:  __ccgo_ts + 40645,
		Ftype1: int8('F'),
	},
	5865: {
		Fword:  __ccgo_ts + 40652,
		Ftype1: int8('F'),
	},
	5866: {
		Fword:  __ccgo_ts + 40658,
		Ftype1: int8('F'),
	},
	5867: {
		Fword:  __ccgo_ts + 40665,
		Ftype1: int8('F'),
	},
	5868: {
		Fword:  __ccgo_ts + 40672,
		Ftype1: int8('F'),
	},
	5869: {
		Fword:  __ccgo_ts + 40679,
		Ftype1: int8('F'),
	},
	5870: {
		Fword:  __ccgo_ts + 40686,
		Ftype1: int8('F'),
	},
	5871: {
		Fword:  __ccgo_ts + 40693,
		Ftype1: int8('F'),
	},
	5872: {
		Fword:  __ccgo_ts + 40700,
		Ftype1: int8('F'),
	},
	5873: {
		Fword:  __ccgo_ts + 40707,
		Ftype1: int8('F'),
	},
	5874: {
		Fword:  __ccgo_ts + 40714,
		Ftype1: int8('F'),
	},
	5875: {
		Fword:  __ccgo_ts + 40721,
		Ftype1: int8('F'),
	},
	5876: {
		Fword:  __ccgo_ts + 40728,
		Ftype1: int8('F'),
	},
	5877: {
		Fword:  __ccgo_ts + 40735,
		Ftype1: int8('F'),
	},
	5878: {
		Fword:  __ccgo_ts + 40742,
		Ftype1: int8('F'),
	},
	5879: {
		Fword:  __ccgo_ts + 40749,
		Ftype1: int8('F'),
	},
	5880: {
		Fword:  __ccgo_ts + 40756,
		Ftype1: int8('F'),
	},
	5881: {
		Fword:  __ccgo_ts + 40763,
		Ftype1: int8('F'),
	},
	5882: {
		Fword:  __ccgo_ts + 40770,
		Ftype1: int8('F'),
	},
	5883: {
		Fword:  __ccgo_ts + 40777,
		Ftype1: int8('F'),
	},
	5884: {
		Fword:  __ccgo_ts + 40784,
		Ftype1: int8('F'),
	},
	5885: {
		Fword:  __ccgo_ts + 40791,
		Ftype1: int8('F'),
	},
	5886: {
		Fword:  __ccgo_ts + 40798,
		Ftype1: int8('F'),
	},
	5887: {
		Fword:  __ccgo_ts + 40805,
		Ftype1: int8('F'),
	},
	5888: {
		Fword:  __ccgo_ts + 40812,
		Ftype1: int8('F'),
	},
	5889: {
		Fword:  __ccgo_ts + 40817,
		Ftype1: int8('F'),
	},
	5890: {
		Fword:  __ccgo_ts + 40824,
		Ftype1: int8('F'),
	},
	5891: {
		Fword:  __ccgo_ts + 40831,
		Ftype1: int8('F'),
	},
	5892: {
		Fword:  __ccgo_ts + 40838,
		Ftype1: int8('F'),
	},
	5893: {
		Fword:  __ccgo_ts + 40845,
		Ftype1: int8('F'),
	},
	5894: {
		Fword:  __ccgo_ts + 40852,
		Ftype1: int8('F'),
	},
	5895: {
		Fword:  __ccgo_ts + 40859,
		Ftype1: int8('F'),
	},
	5896: {
		Fword:  __ccgo_ts + 40865,
		Ftype1: int8('F'),
	},
	5897: {
		Fword:  __ccgo_ts + 40872,
		Ftype1: int8('F'),
	},
	5898: {
		Fword:  __ccgo_ts + 40879,
		Ftype1: int8('F'),
	},
	5899: {
		Fword:  __ccgo_ts + 40886,
		Ftype1: int8('F'),
	},
	5900: {
		Fword:  __ccgo_ts + 40893,
		Ftype1: int8('F'),
	},
	5901: {
		Fword:  __ccgo_ts + 40900,
		Ftype1: int8('F'),
	},
	5902: {
		Fword:  __ccgo_ts + 40907,
		Ftype1: int8('F'),
	},
	5903: {
		Fword:  __ccgo_ts + 40914,
		Ftype1: int8('F'),
	},
	5904: {
		Fword:  __ccgo_ts + 40921,
		Ftype1: int8('F'),
	},
	5905: {
		Fword:  __ccgo_ts + 40928,
		Ftype1: int8('F'),
	},
	5906: {
		Fword:  __ccgo_ts + 40934,
		Ftype1: int8('F'),
	},
	5907: {
		Fword:  __ccgo_ts + 40941,
		Ftype1: int8('F'),
	},
	5908: {
		Fword:  __ccgo_ts + 40948,
		Ftype1: int8('F'),
	},
	5909: {
		Fword:  __ccgo_ts + 40955,
		Ftype1: int8('F'),
	},
	5910: {
		Fword:  __ccgo_ts + 40962,
		Ftype1: int8('F'),
	},
	5911: {
		Fword:  __ccgo_ts + 40969,
		Ftype1: int8('F'),
	},
	5912: {
		Fword:  __ccgo_ts + 40976,
		Ftype1: int8('F'),
	},
	5913: {
		Fword:  __ccgo_ts + 40983,
		Ftype1: int8('F'),
	},
	5914: {
		Fword:  __ccgo_ts + 40990,
		Ftype1: int8('F'),
	},
	5915: {
		Fword:  __ccgo_ts + 40997,
		Ftype1: int8('F'),
	},
	5916: {
		Fword:  __ccgo_ts + 41004,
		Ftype1: int8('F'),
	},
	5917: {
		Fword:  __ccgo_ts + 41011,
		Ftype1: int8('F'),
	},
	5918: {
		Fword:  __ccgo_ts + 41018,
		Ftype1: int8('F'),
	},
	5919: {
		Fword:  __ccgo_ts + 41025,
		Ftype1: int8('F'),
	},
	5920: {
		Fword:  __ccgo_ts + 41032,
		Ftype1: int8('F'),
	},
	5921: {
		Fword:  __ccgo_ts + 41039,
		Ftype1: int8('F'),
	},
	5922: {
		Fword:  __ccgo_ts + 41046,
		Ftype1: int8('F'),
	},
	5923: {
		Fword:  __ccgo_ts + 41053,
		Ftype1: int8('F'),
	},
	5924: {
		Fword:  __ccgo_ts + 41060,
		Ftype1: int8('F'),
	},
	5925: {
		Fword:  __ccgo_ts + 41067,
		Ftype1: int8('F'),
	},
	5926: {
		Fword:  __ccgo_ts + 41074,
		Ftype1: int8('F'),
	},
	5927: {
		Fword:  __ccgo_ts + 41081,
		Ftype1: int8('F'),
	},
	5928: {
		Fword:  __ccgo_ts + 41088,
		Ftype1: int8('F'),
	},
	5929: {
		Fword:  __ccgo_ts + 41095,
		Ftype1: int8('F'),
	},
	5930: {
		Fword:  __ccgo_ts + 41102,
		Ftype1: int8('F'),
	},
	5931: {
		Fword:  __ccgo_ts + 41109,
		Ftype1: int8('F'),
	},
	5932: {
		Fword:  __ccgo_ts + 41116,
		Ftype1: int8('F'),
	},
	5933: {
		Fword:  __ccgo_ts + 41123,
		Ftype1: int8('F'),
	},
	5934: {
		Fword:  __ccgo_ts + 41130,
		Ftype1: int8('F'),
	},
	5935: {
		Fword:  __ccgo_ts + 41137,
		Ftype1: int8('F'),
	},
	5936: {
		Fword:  __ccgo_ts + 41144,
		Ftype1: int8('F'),
	},
	5937: {
		Fword:  __ccgo_ts + 41151,
		Ftype1: int8('F'),
	},
	5938: {
		Fword:  __ccgo_ts + 41158,
		Ftype1: int8('F'),
	},
	5939: {
		Fword:  __ccgo_ts + 41165,
		Ftype1: int8('F'),
	},
	5940: {
		Fword:  __ccgo_ts + 41172,
		Ftype1: int8('F'),
	},
	5941: {
		Fword:  __ccgo_ts + 41179,
		Ftype1: int8('F'),
	},
	5942: {
		Fword:  __ccgo_ts + 41186,
		Ftype1: int8('F'),
	},
	5943: {
		Fword:  __ccgo_ts + 41193,
		Ftype1: int8('F'),
	},
	5944: {
		Fword:  __ccgo_ts + 41200,
		Ftype1: int8('F'),
	},
	5945: {
		Fword:  __ccgo_ts + 41207,
		Ftype1: int8('F'),
	},
	5946: {
		Fword:  __ccgo_ts + 41214,
		Ftype1: int8('F'),
	},
	5947: {
		Fword:  __ccgo_ts + 41221,
		Ftype1: int8('F'),
	},
	5948: {
		Fword:  __ccgo_ts + 41228,
		Ftype1: int8('F'),
	},
	5949: {
		Fword:  __ccgo_ts + 41235,
		Ftype1: int8('F'),
	},
	5950: {
		Fword:  __ccgo_ts + 41242,
		Ftype1: int8('F'),
	},
	5951: {
		Fword:  __ccgo_ts + 41249,
		Ftype1: int8('F'),
	},
	5952: {
		Fword:  __ccgo_ts + 41256,
		Ftype1: int8('F'),
	},
	5953: {
		Fword:  __ccgo_ts + 41263,
		Ftype1: int8('F'),
	},
	5954: {
		Fword:  __ccgo_ts + 41270,
		Ftype1: int8('F'),
	},
	5955: {
		Fword:  __ccgo_ts + 41277,
		Ftype1: int8('F'),
	},
	5956: {
		Fword:  __ccgo_ts + 41284,
		Ftype1: int8('F'),
	},
	5957: {
		Fword:  __ccgo_ts + 41291,
		Ftype1: int8('F'),
	},
	5958: {
		Fword:  __ccgo_ts + 41298,
		Ftype1: int8('F'),
	},
	5959: {
		Fword:  __ccgo_ts + 41305,
		Ftype1: int8('F'),
	},
	5960: {
		Fword:  __ccgo_ts + 41312,
		Ftype1: int8('F'),
	},
	5961: {
		Fword:  __ccgo_ts + 41319,
		Ftype1: int8('F'),
	},
	5962: {
		Fword:  __ccgo_ts + 41326,
		Ftype1: int8('F'),
	},
	5963: {
		Fword:  __ccgo_ts + 41333,
		Ftype1: int8('F'),
	},
	5964: {
		Fword:  __ccgo_ts + 41340,
		Ftype1: int8('F'),
	},
	5965: {
		Fword:  __ccgo_ts + 41347,
		Ftype1: int8('F'),
	},
	5966: {
		Fword:  __ccgo_ts + 41354,
		Ftype1: int8('F'),
	},
	5967: {
		Fword:  __ccgo_ts + 41361,
		Ftype1: int8('F'),
	},
	5968: {
		Fword:  __ccgo_ts + 41367,
		Ftype1: int8('F'),
	},
	5969: {
		Fword:  __ccgo_ts + 41374,
		Ftype1: int8('F'),
	},
	5970: {
		Fword:  __ccgo_ts + 41381,
		Ftype1: int8('F'),
	},
	5971: {
		Fword:  __ccgo_ts + 41388,
		Ftype1: int8('F'),
	},
	5972: {
		Fword:  __ccgo_ts + 41395,
		Ftype1: int8('F'),
	},
	5973: {
		Fword:  __ccgo_ts + 41402,
		Ftype1: int8('F'),
	},
	5974: {
		Fword:  __ccgo_ts + 41409,
		Ftype1: int8('F'),
	},
	5975: {
		Fword:  __ccgo_ts + 41416,
		Ftype1: int8('F'),
	},
	5976: {
		Fword:  __ccgo_ts + 41423,
		Ftype1: int8('F'),
	},
	5977: {
		Fword:  __ccgo_ts + 41430,
		Ftype1: int8('F'),
	},
	5978: {
		Fword:  __ccgo_ts + 41437,
		Ftype1: int8('F'),
	},
	5979: {
		Fword:  __ccgo_ts + 41444,
		Ftype1: int8('F'),
	},
	5980: {
		Fword:  __ccgo_ts + 41451,
		Ftype1: int8('F'),
	},
	5981: {
		Fword:  __ccgo_ts + 41458,
		Ftype1: int8('F'),
	},
	5982: {
		Fword:  __ccgo_ts + 41465,
		Ftype1: int8('F'),
	},
	5983: {
		Fword:  __ccgo_ts + 41472,
		Ftype1: int8('F'),
	},
	5984: {
		Fword:  __ccgo_ts + 41479,
		Ftype1: int8('F'),
	},
	5985: {
		Fword:  __ccgo_ts + 41485,
		Ftype1: int8('F'),
	},
	5986: {
		Fword:  __ccgo_ts + 41492,
		Ftype1: int8('F'),
	},
	5987: {
		Fword:  __ccgo_ts + 41499,
		Ftype1: int8('F'),
	},
	5988: {
		Fword:  __ccgo_ts + 41506,
		Ftype1: int8('F'),
	},
	5989: {
		Fword:  __ccgo_ts + 41513,
		Ftype1: int8('F'),
	},
	5990: {
		Fword:  __ccgo_ts + 41520,
		Ftype1: int8('F'),
	},
	5991: {
		Fword:  __ccgo_ts + 41527,
		Ftype1: int8('F'),
	},
	5992: {
		Fword:  __ccgo_ts + 41534,
		Ftype1: int8('F'),
	},
	5993: {
		Fword:  __ccgo_ts + 41541,
		Ftype1: int8('F'),
	},
	5994: {
		Fword:  __ccgo_ts + 41548,
		Ftype1: int8('F'),
	},
	5995: {
		Fword:  __ccgo_ts + 41555,
		Ftype1: int8('F'),
	},
	5996: {
		Fword:  __ccgo_ts + 41562,
		Ftype1: int8('F'),
	},
	5997: {
		Fword:  __ccgo_ts + 41569,
		Ftype1: int8('F'),
	},
	5998: {
		Fword:  __ccgo_ts + 41576,
		Ftype1: int8('F'),
	},
	5999: {
		Fword:  __ccgo_ts + 41583,
		Ftype1: int8('F'),
	},
	6000: {
		Fword:  __ccgo_ts + 41590,
		Ftype1: int8('F'),
	},
	6001: {
		Fword:  __ccgo_ts + 41597,
		Ftype1: int8('F'),
	},
	6002: {
		Fword:  __ccgo_ts + 41604,
		Ftype1: int8('F'),
	},
	6003: {
		Fword:  __ccgo_ts + 41611,
		Ftype1: int8('F'),
	},
	6004: {
		Fword:  __ccgo_ts + 41618,
		Ftype1: int8('F'),
	},
	6005: {
		Fword:  __ccgo_ts + 41625,
		Ftype1: int8('F'),
	},
	6006: {
		Fword:  __ccgo_ts + 41632,
		Ftype1: int8('F'),
	},
	6007: {
		Fword:  __ccgo_ts + 41639,
		Ftype1: int8('F'),
	},
	6008: {
		Fword:  __ccgo_ts + 41646,
		Ftype1: int8('F'),
	},
	6009: {
		Fword:  __ccgo_ts + 41653,
		Ftype1: int8('F'),
	},
	6010: {
		Fword:  __ccgo_ts + 41660,
		Ftype1: int8('F'),
	},
	6011: {
		Fword:  __ccgo_ts + 41667,
		Ftype1: int8('F'),
	},
	6012: {
		Fword:  __ccgo_ts + 41674,
		Ftype1: int8('F'),
	},
	6013: {
		Fword:  __ccgo_ts + 41681,
		Ftype1: int8('F'),
	},
	6014: {
		Fword:  __ccgo_ts + 41688,
		Ftype1: int8('F'),
	},
	6015: {
		Fword:  __ccgo_ts + 41695,
		Ftype1: int8('F'),
	},
	6016: {
		Fword:  __ccgo_ts + 41702,
		Ftype1: int8('F'),
	},
	6017: {
		Fword:  __ccgo_ts + 41709,
		Ftype1: int8('F'),
	},
	6018: {
		Fword:  __ccgo_ts + 41716,
		Ftype1: int8('F'),
	},
	6019: {
		Fword:  __ccgo_ts + 41723,
		Ftype1: int8('F'),
	},
	6020: {
		Fword:  __ccgo_ts + 41729,
		Ftype1: int8('F'),
	},
	6021: {
		Fword:  __ccgo_ts + 41736,
		Ftype1: int8('F'),
	},
	6022: {
		Fword:  __ccgo_ts + 41743,
		Ftype1: int8('F'),
	},
	6023: {
		Fword:  __ccgo_ts + 41750,
		Ftype1: int8('F'),
	},
	6024: {
		Fword:  __ccgo_ts + 41757,
		Ftype1: int8('F'),
	},
	6025: {
		Fword:  __ccgo_ts + 41764,
		Ftype1: int8('F'),
	},
	6026: {
		Fword:  __ccgo_ts + 41771,
		Ftype1: int8('F'),
	},
	6027: {
		Fword:  __ccgo_ts + 41778,
		Ftype1: int8('F'),
	},
	6028: {
		Fword:  __ccgo_ts + 41785,
		Ftype1: int8('F'),
	},
	6029: {
		Fword:  __ccgo_ts + 41792,
		Ftype1: int8('F'),
	},
	6030: {
		Fword:  __ccgo_ts + 41799,
		Ftype1: int8('F'),
	},
	6031: {
		Fword:  __ccgo_ts + 41806,
		Ftype1: int8('F'),
	},
	6032: {
		Fword:  __ccgo_ts + 41813,
		Ftype1: int8('F'),
	},
	6033: {
		Fword:  __ccgo_ts + 41820,
		Ftype1: int8('F'),
	},
	6034: {
		Fword:  __ccgo_ts + 41827,
		Ftype1: int8('F'),
	},
	6035: {
		Fword:  __ccgo_ts + 41834,
		Ftype1: int8('F'),
	},
	6036: {
		Fword:  __ccgo_ts + 41841,
		Ftype1: int8('F'),
	},
	6037: {
		Fword:  __ccgo_ts + 41848,
		Ftype1: int8('F'),
	},
	6038: {
		Fword:  __ccgo_ts + 41855,
		Ftype1: int8('F'),
	},
	6039: {
		Fword:  __ccgo_ts + 41862,
		Ftype1: int8('F'),
	},
	6040: {
		Fword:  __ccgo_ts + 41869,
		Ftype1: int8('F'),
	},
	6041: {
		Fword:  __ccgo_ts + 41876,
		Ftype1: int8('F'),
	},
	6042: {
		Fword:  __ccgo_ts + 41883,
		Ftype1: int8('F'),
	},
	6043: {
		Fword:  __ccgo_ts + 41890,
		Ftype1: int8('F'),
	},
	6044: {
		Fword:  __ccgo_ts + 41897,
		Ftype1: int8('F'),
	},
	6045: {
		Fword:  __ccgo_ts + 41904,
		Ftype1: int8('F'),
	},
	6046: {
		Fword:  __ccgo_ts + 41911,
		Ftype1: int8('F'),
	},
	6047: {
		Fword:  __ccgo_ts + 41918,
		Ftype1: int8('F'),
	},
	6048: {
		Fword:  __ccgo_ts + 41925,
		Ftype1: int8('F'),
	},
	6049: {
		Fword:  __ccgo_ts + 41932,
		Ftype1: int8('F'),
	},
	6050: {
		Fword:  __ccgo_ts + 41939,
		Ftype1: int8('F'),
	},
	6051: {
		Fword:  __ccgo_ts + 41946,
		Ftype1: int8('F'),
	},
	6052: {
		Fword:  __ccgo_ts + 41953,
		Ftype1: int8('F'),
	},
	6053: {
		Fword:  __ccgo_ts + 41960,
		Ftype1: int8('F'),
	},
	6054: {
		Fword:  __ccgo_ts + 41967,
		Ftype1: int8('F'),
	},
	6055: {
		Fword:  __ccgo_ts + 41974,
		Ftype1: int8('F'),
	},
	6056: {
		Fword:  __ccgo_ts + 41981,
		Ftype1: int8('F'),
	},
	6057: {
		Fword:  __ccgo_ts + 41988,
		Ftype1: int8('F'),
	},
	6058: {
		Fword:  __ccgo_ts + 41995,
		Ftype1: int8('F'),
	},
	6059: {
		Fword:  __ccgo_ts + 42002,
		Ftype1: int8('F'),
	},
	6060: {
		Fword:  __ccgo_ts + 42009,
		Ftype1: int8('F'),
	},
	6061: {
		Fword:  __ccgo_ts + 42016,
		Ftype1: int8('F'),
	},
	6062: {
		Fword:  __ccgo_ts + 42023,
		Ftype1: int8('F'),
	},
	6063: {
		Fword:  __ccgo_ts + 42030,
		Ftype1: int8('F'),
	},
	6064: {
		Fword:  __ccgo_ts + 42037,
		Ftype1: int8('F'),
	},
	6065: {
		Fword:  __ccgo_ts + 42044,
		Ftype1: int8('F'),
	},
	6066: {
		Fword:  __ccgo_ts + 42051,
		Ftype1: int8('F'),
	},
	6067: {
		Fword:  __ccgo_ts + 42058,
		Ftype1: int8('F'),
	},
	6068: {
		Fword:  __ccgo_ts + 42065,
		Ftype1: int8('F'),
	},
	6069: {
		Fword:  __ccgo_ts + 42072,
		Ftype1: int8('F'),
	},
	6070: {
		Fword:  __ccgo_ts + 42079,
		Ftype1: int8('F'),
	},
	6071: {
		Fword:  __ccgo_ts + 42086,
		Ftype1: int8('F'),
	},
	6072: {
		Fword:  __ccgo_ts + 42093,
		Ftype1: int8('F'),
	},
	6073: {
		Fword:  __ccgo_ts + 42100,
		Ftype1: int8('F'),
	},
	6074: {
		Fword:  __ccgo_ts + 42107,
		Ftype1: int8('F'),
	},
	6075: {
		Fword:  __ccgo_ts + 42114,
		Ftype1: int8('F'),
	},
	6076: {
		Fword:  __ccgo_ts + 42121,
		Ftype1: int8('F'),
	},
	6077: {
		Fword:  __ccgo_ts + 42128,
		Ftype1: int8('F'),
	},
	6078: {
		Fword:  __ccgo_ts + 42135,
		Ftype1: int8('F'),
	},
	6079: {
		Fword:  __ccgo_ts + 42142,
		Ftype1: int8('F'),
	},
	6080: {
		Fword:  __ccgo_ts + 42149,
		Ftype1: int8('F'),
	},
	6081: {
		Fword:  __ccgo_ts + 42156,
		Ftype1: int8('F'),
	},
	6082: {
		Fword:  __ccgo_ts + 42163,
		Ftype1: int8('F'),
	},
	6083: {
		Fword:  __ccgo_ts + 42170,
		Ftype1: int8('F'),
	},
	6084: {
		Fword:  __ccgo_ts + 42177,
		Ftype1: int8('F'),
	},
	6085: {
		Fword:  __ccgo_ts + 42184,
		Ftype1: int8('F'),
	},
	6086: {
		Fword:  __ccgo_ts + 42191,
		Ftype1: int8('F'),
	},
	6087: {
		Fword:  __ccgo_ts + 42198,
		Ftype1: int8('F'),
	},
	6088: {
		Fword:  __ccgo_ts + 42205,
		Ftype1: int8('F'),
	},
	6089: {
		Fword:  __ccgo_ts + 42212,
		Ftype1: int8('F'),
	},
	6090: {
		Fword:  __ccgo_ts + 42219,
		Ftype1: int8('F'),
	},
	6091: {
		Fword:  __ccgo_ts + 42226,
		Ftype1: int8('F'),
	},
	6092: {
		Fword:  __ccgo_ts + 42233,
		Ftype1: int8('F'),
	},
	6093: {
		Fword:  __ccgo_ts + 42240,
		Ftype1: int8('F'),
	},
	6094: {
		Fword:  __ccgo_ts + 42247,
		Ftype1: int8('F'),
	},
	6095: {
		Fword:  __ccgo_ts + 42254,
		Ftype1: int8('F'),
	},
	6096: {
		Fword:  __ccgo_ts + 42261,
		Ftype1: int8('F'),
	},
	6097: {
		Fword:  __ccgo_ts + 42268,
		Ftype1: int8('F'),
	},
	6098: {
		Fword:  __ccgo_ts + 42275,
		Ftype1: int8('F'),
	},
	6099: {
		Fword:  __ccgo_ts + 42282,
		Ftype1: int8('F'),
	},
	6100: {
		Fword:  __ccgo_ts + 42288,
		Ftype1: int8('F'),
	},
	6101: {
		Fword:  __ccgo_ts + 42295,
		Ftype1: int8('F'),
	},
	6102: {
		Fword:  __ccgo_ts + 42302,
		Ftype1: int8('F'),
	},
	6103: {
		Fword:  __ccgo_ts + 42309,
		Ftype1: int8('F'),
	},
	6104: {
		Fword:  __ccgo_ts + 42316,
		Ftype1: int8('F'),
	},
	6105: {
		Fword:  __ccgo_ts + 42323,
		Ftype1: int8('F'),
	},
	6106: {
		Fword:  __ccgo_ts + 42330,
		Ftype1: int8('F'),
	},
	6107: {
		Fword:  __ccgo_ts + 42337,
		Ftype1: int8('F'),
	},
	6108: {
		Fword:  __ccgo_ts + 42344,
		Ftype1: int8('F'),
	},
	6109: {
		Fword:  __ccgo_ts + 42351,
		Ftype1: int8('F'),
	},
	6110: {
		Fword:  __ccgo_ts + 42358,
		Ftype1: int8('F'),
	},
	6111: {
		Fword:  __ccgo_ts + 42365,
		Ftype1: int8('F'),
	},
	6112: {
		Fword:  __ccgo_ts + 42372,
		Ftype1: int8('F'),
	},
	6113: {
		Fword:  __ccgo_ts + 42379,
		Ftype1: int8('F'),
	},
	6114: {
		Fword:  __ccgo_ts + 42386,
		Ftype1: int8('F'),
	},
	6115: {
		Fword:  __ccgo_ts + 42393,
		Ftype1: int8('F'),
	},
	6116: {
		Fword:  __ccgo_ts + 42400,
		Ftype1: int8('F'),
	},
	6117: {
		Fword:  __ccgo_ts + 42407,
		Ftype1: int8('F'),
	},
	6118: {
		Fword:  __ccgo_ts + 42414,
		Ftype1: int8('F'),
	},
	6119: {
		Fword:  __ccgo_ts + 42421,
		Ftype1: int8('F'),
	},
	6120: {
		Fword:  __ccgo_ts + 42428,
		Ftype1: int8('F'),
	},
	6121: {
		Fword:  __ccgo_ts + 42435,
		Ftype1: int8('F'),
	},
	6122: {
		Fword:  __ccgo_ts + 42442,
		Ftype1: int8('F'),
	},
	6123: {
		Fword:  __ccgo_ts + 42449,
		Ftype1: int8('F'),
	},
	6124: {
		Fword:  __ccgo_ts + 42456,
		Ftype1: int8('F'),
	},
	6125: {
		Fword:  __ccgo_ts + 42463,
		Ftype1: int8('F'),
	},
	6126: {
		Fword:  __ccgo_ts + 42470,
		Ftype1: int8('F'),
	},
	6127: {
		Fword:  __ccgo_ts + 42477,
		Ftype1: int8('F'),
	},
	6128: {
		Fword:  __ccgo_ts + 42484,
		Ftype1: int8('F'),
	},
	6129: {
		Fword:  __ccgo_ts + 42491,
		Ftype1: int8('F'),
	},
	6130: {
		Fword:  __ccgo_ts + 42498,
		Ftype1: int8('F'),
	},
	6131: {
		Fword:  __ccgo_ts + 42505,
		Ftype1: int8('F'),
	},
	6132: {
		Fword:  __ccgo_ts + 42512,
		Ftype1: int8('F'),
	},
	6133: {
		Fword:  __ccgo_ts + 42519,
		Ftype1: int8('F'),
	},
	6134: {
		Fword:  __ccgo_ts + 42526,
		Ftype1: int8('F'),
	},
	6135: {
		Fword:  __ccgo_ts + 42533,
		Ftype1: int8('F'),
	},
	6136: {
		Fword:  __ccgo_ts + 42540,
		Ftype1: int8('F'),
	},
	6137: {
		Fword:  __ccgo_ts + 42547,
		Ftype1: int8('F'),
	},
	6138: {
		Fword:  __ccgo_ts + 42554,
		Ftype1: int8('F'),
	},
	6139: {
		Fword:  __ccgo_ts + 42561,
		Ftype1: int8('F'),
	},
	6140: {
		Fword:  __ccgo_ts + 42568,
		Ftype1: int8('F'),
	},
	6141: {
		Fword:  __ccgo_ts + 42575,
		Ftype1: int8('F'),
	},
	6142: {
		Fword:  __ccgo_ts + 42582,
		Ftype1: int8('F'),
	},
	6143: {
		Fword:  __ccgo_ts + 42588,
		Ftype1: int8('F'),
	},
	6144: {
		Fword:  __ccgo_ts + 42595,
		Ftype1: int8('F'),
	},
	6145: {
		Fword:  __ccgo_ts + 42602,
		Ftype1: int8('F'),
	},
	6146: {
		Fword:  __ccgo_ts + 42609,
		Ftype1: int8('F'),
	},
	6147: {
		Fword:  __ccgo_ts + 42616,
		Ftype1: int8('F'),
	},
	6148: {
		Fword:  __ccgo_ts + 42623,
		Ftype1: int8('F'),
	},
	6149: {
		Fword:  __ccgo_ts + 42630,
		Ftype1: int8('F'),
	},
	6150: {
		Fword:  __ccgo_ts + 42637,
		Ftype1: int8('F'),
	},
	6151: {
		Fword:  __ccgo_ts + 42644,
		Ftype1: int8('F'),
	},
	6152: {
		Fword:  __ccgo_ts + 42651,
		Ftype1: int8('F'),
	},
	6153: {
		Fword:  __ccgo_ts + 42658,
		Ftype1: int8('F'),
	},
	6154: {
		Fword:  __ccgo_ts + 42665,
		Ftype1: int8('F'),
	},
	6155: {
		Fword:  __ccgo_ts + 42670,
		Ftype1: int8('F'),
	},
	6156: {
		Fword:  __ccgo_ts + 42677,
		Ftype1: int8('F'),
	},
	6157: {
		Fword:  __ccgo_ts + 42684,
		Ftype1: int8('F'),
	},
	6158: {
		Fword:  __ccgo_ts + 42691,
		Ftype1: int8('F'),
	},
	6159: {
		Fword:  __ccgo_ts + 42698,
		Ftype1: int8('F'),
	},
	6160: {
		Fword:  __ccgo_ts + 42705,
		Ftype1: int8('F'),
	},
	6161: {
		Fword:  __ccgo_ts + 42712,
		Ftype1: int8('F'),
	},
	6162: {
		Fword:  __ccgo_ts + 42719,
		Ftype1: int8('F'),
	},
	6163: {
		Fword:  __ccgo_ts + 42726,
		Ftype1: int8('F'),
	},
	6164: {
		Fword:  __ccgo_ts + 42733,
		Ftype1: int8('F'),
	},
	6165: {
		Fword:  __ccgo_ts + 42740,
		Ftype1: int8('F'),
	},
	6166: {
		Fword:  __ccgo_ts + 42747,
		Ftype1: int8('F'),
	},
	6167: {
		Fword:  __ccgo_ts + 42754,
		Ftype1: int8('F'),
	},
	6168: {
		Fword:  __ccgo_ts + 42761,
		Ftype1: int8('F'),
	},
	6169: {
		Fword:  __ccgo_ts + 42768,
		Ftype1: int8('F'),
	},
	6170: {
		Fword:  __ccgo_ts + 42775,
		Ftype1: int8('F'),
	},
	6171: {
		Fword:  __ccgo_ts + 42782,
		Ftype1: int8('F'),
	},
	6172: {
		Fword:  __ccgo_ts + 42789,
		Ftype1: int8('F'),
	},
	6173: {
		Fword:  __ccgo_ts + 42796,
		Ftype1: int8('F'),
	},
	6174: {
		Fword:  __ccgo_ts + 42803,
		Ftype1: int8('F'),
	},
	6175: {
		Fword:  __ccgo_ts + 42810,
		Ftype1: int8('F'),
	},
	6176: {
		Fword:  __ccgo_ts + 42817,
		Ftype1: int8('F'),
	},
	6177: {
		Fword:  __ccgo_ts + 42824,
		Ftype1: int8('F'),
	},
	6178: {
		Fword:  __ccgo_ts + 42831,
		Ftype1: int8('F'),
	},
	6179: {
		Fword:  __ccgo_ts + 42838,
		Ftype1: int8('F'),
	},
	6180: {
		Fword:  __ccgo_ts + 42845,
		Ftype1: int8('F'),
	},
	6181: {
		Fword:  __ccgo_ts + 42852,
		Ftype1: int8('F'),
	},
	6182: {
		Fword:  __ccgo_ts + 42859,
		Ftype1: int8('F'),
	},
	6183: {
		Fword:  __ccgo_ts + 42866,
		Ftype1: int8('F'),
	},
	6184: {
		Fword:  __ccgo_ts + 42873,
		Ftype1: int8('F'),
	},
	6185: {
		Fword:  __ccgo_ts + 42879,
		Ftype1: int8('F'),
	},
	6186: {
		Fword:  __ccgo_ts + 42886,
		Ftype1: int8('F'),
	},
	6187: {
		Fword:  __ccgo_ts + 42893,
		Ftype1: int8('F'),
	},
	6188: {
		Fword:  __ccgo_ts + 42900,
		Ftype1: int8('F'),
	},
	6189: {
		Fword:  __ccgo_ts + 42907,
		Ftype1: int8('F'),
	},
	6190: {
		Fword:  __ccgo_ts + 42914,
		Ftype1: int8('F'),
	},
	6191: {
		Fword:  __ccgo_ts + 42921,
		Ftype1: int8('F'),
	},
	6192: {
		Fword:  __ccgo_ts + 42928,
		Ftype1: int8('F'),
	},
	6193: {
		Fword:  __ccgo_ts + 42935,
		Ftype1: int8('F'),
	},
	6194: {
		Fword:  __ccgo_ts + 42942,
		Ftype1: int8('F'),
	},
	6195: {
		Fword:  __ccgo_ts + 42949,
		Ftype1: int8('F'),
	},
	6196: {
		Fword:  __ccgo_ts + 42956,
		Ftype1: int8('F'),
	},
	6197: {
		Fword:  __ccgo_ts + 42963,
		Ftype1: int8('F'),
	},
	6198: {
		Fword:  __ccgo_ts + 42970,
		Ftype1: int8('F'),
	},
	6199: {
		Fword:  __ccgo_ts + 42977,
		Ftype1: int8('F'),
	},
	6200: {
		Fword:  __ccgo_ts + 42984,
		Ftype1: int8('F'),
	},
	6201: {
		Fword:  __ccgo_ts + 42991,
		Ftype1: int8('F'),
	},
	6202: {
		Fword:  __ccgo_ts + 42997,
		Ftype1: int8('F'),
	},
	6203: {
		Fword:  __ccgo_ts + 43004,
		Ftype1: int8('F'),
	},
	6204: {
		Fword:  __ccgo_ts + 43011,
		Ftype1: int8('F'),
	},
	6205: {
		Fword:  __ccgo_ts + 43018,
		Ftype1: int8('F'),
	},
	6206: {
		Fword:  __ccgo_ts + 43025,
		Ftype1: int8('F'),
	},
	6207: {
		Fword:  __ccgo_ts + 43032,
		Ftype1: int8('F'),
	},
	6208: {
		Fword:  __ccgo_ts + 43039,
		Ftype1: int8('F'),
	},
	6209: {
		Fword:  __ccgo_ts + 43046,
		Ftype1: int8('F'),
	},
	6210: {
		Fword:  __ccgo_ts + 43053,
		Ftype1: int8('F'),
	},
	6211: {
		Fword:  __ccgo_ts + 43060,
		Ftype1: int8('F'),
	},
	6212: {
		Fword:  __ccgo_ts + 43067,
		Ftype1: int8('F'),
	},
	6213: {
		Fword:  __ccgo_ts + 43074,
		Ftype1: int8('F'),
	},
	6214: {
		Fword:  __ccgo_ts + 43081,
		Ftype1: int8('F'),
	},
	6215: {
		Fword:  __ccgo_ts + 43088,
		Ftype1: int8('F'),
	},
	6216: {
		Fword:  __ccgo_ts + 43095,
		Ftype1: int8('F'),
	},
	6217: {
		Fword:  __ccgo_ts + 43102,
		Ftype1: int8('F'),
	},
	6218: {
		Fword:  __ccgo_ts + 43109,
		Ftype1: int8('F'),
	},
	6219: {
		Fword:  __ccgo_ts + 43116,
		Ftype1: int8('F'),
	},
	6220: {
		Fword:  __ccgo_ts + 43123,
		Ftype1: int8('F'),
	},
	6221: {
		Fword:  __ccgo_ts + 43130,
		Ftype1: int8('F'),
	},
	6222: {
		Fword:  __ccgo_ts + 43137,
		Ftype1: int8('F'),
	},
	6223: {
		Fword:  __ccgo_ts + 43144,
		Ftype1: int8('F'),
	},
	6224: {
		Fword:  __ccgo_ts + 43151,
		Ftype1: int8('F'),
	},
	6225: {
		Fword:  __ccgo_ts + 43158,
		Ftype1: int8('F'),
	},
	6226: {
		Fword:  __ccgo_ts + 43165,
		Ftype1: int8('F'),
	},
	6227: {
		Fword:  __ccgo_ts + 43172,
		Ftype1: int8('F'),
	},
	6228: {
		Fword:  __ccgo_ts + 43179,
		Ftype1: int8('F'),
	},
	6229: {
		Fword:  __ccgo_ts + 43186,
		Ftype1: int8('F'),
	},
	6230: {
		Fword:  __ccgo_ts + 43192,
		Ftype1: int8('F'),
	},
	6231: {
		Fword:  __ccgo_ts + 43199,
		Ftype1: int8('F'),
	},
	6232: {
		Fword:  __ccgo_ts + 43206,
		Ftype1: int8('F'),
	},
	6233: {
		Fword:  __ccgo_ts + 43213,
		Ftype1: int8('F'),
	},
	6234: {
		Fword:  __ccgo_ts + 43220,
		Ftype1: int8('F'),
	},
	6235: {
		Fword:  __ccgo_ts + 43227,
		Ftype1: int8('F'),
	},
	6236: {
		Fword:  __ccgo_ts + 43234,
		Ftype1: int8('F'),
	},
	6237: {
		Fword:  __ccgo_ts + 43241,
		Ftype1: int8('F'),
	},
	6238: {
		Fword:  __ccgo_ts + 43248,
		Ftype1: int8('F'),
	},
	6239: {
		Fword:  __ccgo_ts + 43255,
		Ftype1: int8('F'),
	},
	6240: {
		Fword:  __ccgo_ts + 43262,
		Ftype1: int8('F'),
	},
	6241: {
		Fword:  __ccgo_ts + 43269,
		Ftype1: int8('F'),
	},
	6242: {
		Fword:  __ccgo_ts + 43276,
		Ftype1: int8('F'),
	},
	6243: {
		Fword:  __ccgo_ts + 43283,
		Ftype1: int8('F'),
	},
	6244: {
		Fword:  __ccgo_ts + 43290,
		Ftype1: int8('F'),
	},
	6245: {
		Fword:  __ccgo_ts + 43297,
		Ftype1: int8('F'),
	},
	6246: {
		Fword:  __ccgo_ts + 43304,
		Ftype1: int8('F'),
	},
	6247: {
		Fword:  __ccgo_ts + 43311,
		Ftype1: int8('F'),
	},
	6248: {
		Fword:  __ccgo_ts + 43318,
		Ftype1: int8('F'),
	},
	6249: {
		Fword:  __ccgo_ts + 43325,
		Ftype1: int8('F'),
	},
	6250: {
		Fword:  __ccgo_ts + 43330,
		Ftype1: int8('F'),
	},
	6251: {
		Fword:  __ccgo_ts + 43337,
		Ftype1: int8('F'),
	},
	6252: {
		Fword:  __ccgo_ts + 43344,
		Ftype1: int8('F'),
	},
	6253: {
		Fword:  __ccgo_ts + 43351,
		Ftype1: int8('F'),
	},
	6254: {
		Fword:  __ccgo_ts + 43358,
		Ftype1: int8('F'),
	},
	6255: {
		Fword:  __ccgo_ts + 43365,
		Ftype1: int8('F'),
	},
	6256: {
		Fword:  __ccgo_ts + 43372,
		Ftype1: int8('F'),
	},
	6257: {
		Fword:  __ccgo_ts + 43379,
		Ftype1: int8('F'),
	},
	6258: {
		Fword:  __ccgo_ts + 43386,
		Ftype1: int8('F'),
	},
	6259: {
		Fword:  __ccgo_ts + 43393,
		Ftype1: int8('F'),
	},
	6260: {
		Fword:  __ccgo_ts + 43400,
		Ftype1: int8('F'),
	},
	6261: {
		Fword:  __ccgo_ts + 43407,
		Ftype1: int8('F'),
	},
	6262: {
		Fword:  __ccgo_ts + 43414,
		Ftype1: int8('F'),
	},
	6263: {
		Fword:  __ccgo_ts + 43421,
		Ftype1: int8('F'),
	},
	6264: {
		Fword:  __ccgo_ts + 43428,
		Ftype1: int8('F'),
	},
	6265: {
		Fword:  __ccgo_ts + 43435,
		Ftype1: int8('F'),
	},
	6266: {
		Fword:  __ccgo_ts + 43442,
		Ftype1: int8('F'),
	},
	6267: {
		Fword:  __ccgo_ts + 43449,
		Ftype1: int8('F'),
	},
	6268: {
		Fword:  __ccgo_ts + 43456,
		Ftype1: int8('F'),
	},
	6269: {
		Fword:  __ccgo_ts + 43463,
		Ftype1: int8('F'),
	},
	6270: {
		Fword:  __ccgo_ts + 43470,
		Ftype1: int8('F'),
	},
	6271: {
		Fword:  __ccgo_ts + 43477,
		Ftype1: int8('F'),
	},
	6272: {
		Fword:  __ccgo_ts + 43484,
		Ftype1: int8('F'),
	},
	6273: {
		Fword:  __ccgo_ts + 43491,
		Ftype1: int8('F'),
	},
	6274: {
		Fword:  __ccgo_ts + 43497,
		Ftype1: int8('F'),
	},
	6275: {
		Fword:  __ccgo_ts + 43504,
		Ftype1: int8('F'),
	},
	6276: {
		Fword:  __ccgo_ts + 43511,
		Ftype1: int8('F'),
	},
	6277: {
		Fword:  __ccgo_ts + 43518,
		Ftype1: int8('F'),
	},
	6278: {
		Fword:  __ccgo_ts + 43525,
		Ftype1: int8('F'),
	},
	6279: {
		Fword:  __ccgo_ts + 43532,
		Ftype1: int8('F'),
	},
	6280: {
		Fword:  __ccgo_ts + 43539,
		Ftype1: int8('F'),
	},
	6281: {
		Fword:  __ccgo_ts + 43546,
		Ftype1: int8('F'),
	},
	6282: {
		Fword:  __ccgo_ts + 43553,
		Ftype1: int8('F'),
	},
	6283: {
		Fword:  __ccgo_ts + 43560,
		Ftype1: int8('F'),
	},
	6284: {
		Fword:  __ccgo_ts + 43567,
		Ftype1: int8('F'),
	},
	6285: {
		Fword:  __ccgo_ts + 43574,
		Ftype1: int8('F'),
	},
	6286: {
		Fword:  __ccgo_ts + 43581,
		Ftype1: int8('F'),
	},
	6287: {
		Fword:  __ccgo_ts + 43588,
		Ftype1: int8('F'),
	},
	6288: {
		Fword:  __ccgo_ts + 43595,
		Ftype1: int8('F'),
	},
	6289: {
		Fword:  __ccgo_ts + 43602,
		Ftype1: int8('F'),
	},
	6290: {
		Fword:  __ccgo_ts + 43609,
		Ftype1: int8('F'),
	},
	6291: {
		Fword:  __ccgo_ts + 43615,
		Ftype1: int8('F'),
	},
	6292: {
		Fword:  __ccgo_ts + 43622,
		Ftype1: int8('F'),
	},
	6293: {
		Fword:  __ccgo_ts + 43629,
		Ftype1: int8('F'),
	},
	6294: {
		Fword:  __ccgo_ts + 43636,
		Ftype1: int8('F'),
	},
	6295: {
		Fword:  __ccgo_ts + 43643,
		Ftype1: int8('F'),
	},
	6296: {
		Fword:  __ccgo_ts + 43650,
		Ftype1: int8('F'),
	},
	6297: {
		Fword:  __ccgo_ts + 43657,
		Ftype1: int8('F'),
	},
	6298: {
		Fword:  __ccgo_ts + 43664,
		Ftype1: int8('F'),
	},
	6299: {
		Fword:  __ccgo_ts + 43671,
		Ftype1: int8('F'),
	},
	6300: {
		Fword:  __ccgo_ts + 43678,
		Ftype1: int8('F'),
	},
	6301: {
		Fword:  __ccgo_ts + 43685,
		Ftype1: int8('F'),
	},
	6302: {
		Fword:  __ccgo_ts + 43692,
		Ftype1: int8('F'),
	},
	6303: {
		Fword:  __ccgo_ts + 43699,
		Ftype1: int8('F'),
	},
	6304: {
		Fword:  __ccgo_ts + 43706,
		Ftype1: int8('F'),
	},
	6305: {
		Fword:  __ccgo_ts + 43713,
		Ftype1: int8('F'),
	},
	6306: {
		Fword:  __ccgo_ts + 43720,
		Ftype1: int8('F'),
	},
	6307: {
		Fword:  __ccgo_ts + 43727,
		Ftype1: int8('F'),
	},
	6308: {
		Fword:  __ccgo_ts + 43734,
		Ftype1: int8('F'),
	},
	6309: {
		Fword:  __ccgo_ts + 43741,
		Ftype1: int8('F'),
	},
	6310: {
		Fword:  __ccgo_ts + 43748,
		Ftype1: int8('F'),
	},
	6311: {
		Fword:  __ccgo_ts + 43755,
		Ftype1: int8('F'),
	},
	6312: {
		Fword:  __ccgo_ts + 43762,
		Ftype1: int8('F'),
	},
	6313: {
		Fword:  __ccgo_ts + 43769,
		Ftype1: int8('F'),
	},
	6314: {
		Fword:  __ccgo_ts + 43776,
		Ftype1: int8('F'),
	},
	6315: {
		Fword:  __ccgo_ts + 43783,
		Ftype1: int8('F'),
	},
	6316: {
		Fword:  __ccgo_ts + 43790,
		Ftype1: int8('F'),
	},
	6317: {
		Fword:  __ccgo_ts + 43797,
		Ftype1: int8('F'),
	},
	6318: {
		Fword:  __ccgo_ts + 43804,
		Ftype1: int8('F'),
	},
	6319: {
		Fword:  __ccgo_ts + 43811,
		Ftype1: int8('F'),
	},
	6320: {
		Fword:  __ccgo_ts + 43818,
		Ftype1: int8('F'),
	},
	6321: {
		Fword:  __ccgo_ts + 43825,
		Ftype1: int8('F'),
	},
	6322: {
		Fword:  __ccgo_ts + 43832,
		Ftype1: int8('F'),
	},
	6323: {
		Fword:  __ccgo_ts + 43839,
		Ftype1: int8('F'),
	},
	6324: {
		Fword:  __ccgo_ts + 43846,
		Ftype1: int8('F'),
	},
	6325: {
		Fword:  __ccgo_ts + 43853,
		Ftype1: int8('F'),
	},
	6326: {
		Fword:  __ccgo_ts + 43860,
		Ftype1: int8('F'),
	},
	6327: {
		Fword:  __ccgo_ts + 43867,
		Ftype1: int8('F'),
	},
	6328: {
		Fword:  __ccgo_ts + 43874,
		Ftype1: int8('F'),
	},
	6329: {
		Fword:  __ccgo_ts + 43881,
		Ftype1: int8('F'),
	},
	6330: {
		Fword:  __ccgo_ts + 43887,
		Ftype1: int8('F'),
	},
	6331: {
		Fword:  __ccgo_ts + 43894,
		Ftype1: int8('F'),
	},
	6332: {
		Fword:  __ccgo_ts + 43901,
		Ftype1: int8('F'),
	},
	6333: {
		Fword:  __ccgo_ts + 43908,
		Ftype1: int8('F'),
	},
	6334: {
		Fword:  __ccgo_ts + 43915,
		Ftype1: int8('F'),
	},
	6335: {
		Fword:  __ccgo_ts + 43922,
		Ftype1: int8('F'),
	},
	6336: {
		Fword:  __ccgo_ts + 43929,
		Ftype1: int8('F'),
	},
	6337: {
		Fword:  __ccgo_ts + 43936,
		Ftype1: int8('F'),
	},
	6338: {
		Fword:  __ccgo_ts + 43943,
		Ftype1: int8('F'),
	},
	6339: {
		Fword:  __ccgo_ts + 43950,
		Ftype1: int8('F'),
	},
	6340: {
		Fword:  __ccgo_ts + 43957,
		Ftype1: int8('F'),
	},
	6341: {
		Fword:  __ccgo_ts + 43964,
		Ftype1: int8('F'),
	},
	6342: {
		Fword:  __ccgo_ts + 43971,
		Ftype1: int8('F'),
	},
	6343: {
		Fword:  __ccgo_ts + 43978,
		Ftype1: int8('F'),
	},
	6344: {
		Fword:  __ccgo_ts + 43985,
		Ftype1: int8('F'),
	},
	6345: {
		Fword:  __ccgo_ts + 43992,
		Ftype1: int8('F'),
	},
	6346: {
		Fword:  __ccgo_ts + 43999,
		Ftype1: int8('F'),
	},
	6347: {
		Fword:  __ccgo_ts + 44006,
		Ftype1: int8('F'),
	},
	6348: {
		Fword:  __ccgo_ts + 44013,
		Ftype1: int8('F'),
	},
	6349: {
		Fword:  __ccgo_ts + 44020,
		Ftype1: int8('F'),
	},
	6350: {
		Fword:  __ccgo_ts + 44027,
		Ftype1: int8('F'),
	},
	6351: {
		Fword:  __ccgo_ts + 44034,
		Ftype1: int8('F'),
	},
	6352: {
		Fword:  __ccgo_ts + 44041,
		Ftype1: int8('F'),
	},
	6353: {
		Fword:  __ccgo_ts + 44048,
		Ftype1: int8('F'),
	},
	6354: {
		Fword:  __ccgo_ts + 44055,
		Ftype1: int8('F'),
	},
	6355: {
		Fword:  __ccgo_ts + 44062,
		Ftype1: int8('F'),
	},
	6356: {
		Fword:  __ccgo_ts + 44069,
		Ftype1: int8('F'),
	},
	6357: {
		Fword:  __ccgo_ts + 44076,
		Ftype1: int8('F'),
	},
	6358: {
		Fword:  __ccgo_ts + 44083,
		Ftype1: int8('F'),
	},
	6359: {
		Fword:  __ccgo_ts + 44090,
		Ftype1: int8('F'),
	},
	6360: {
		Fword:  __ccgo_ts + 44097,
		Ftype1: int8('F'),
	},
	6361: {
		Fword:  __ccgo_ts + 44104,
		Ftype1: int8('F'),
	},
	6362: {
		Fword:  __ccgo_ts + 44111,
		Ftype1: int8('F'),
	},
	6363: {
		Fword:  __ccgo_ts + 44118,
		Ftype1: int8('F'),
	},
	6364: {
		Fword:  __ccgo_ts + 44125,
		Ftype1: int8('F'),
	},
	6365: {
		Fword:  __ccgo_ts + 44132,
		Ftype1: int8('F'),
	},
	6366: {
		Fword:  __ccgo_ts + 44139,
		Ftype1: int8('F'),
	},
	6367: {
		Fword:  __ccgo_ts + 44146,
		Ftype1: int8('F'),
	},
	6368: {
		Fword:  __ccgo_ts + 44153,
		Ftype1: int8('F'),
	},
	6369: {
		Fword:  __ccgo_ts + 44160,
		Ftype1: int8('F'),
	},
	6370: {
		Fword:  __ccgo_ts + 44167,
		Ftype1: int8('F'),
	},
	6371: {
		Fword:  __ccgo_ts + 44174,
		Ftype1: int8('F'),
	},
	6372: {
		Fword:  __ccgo_ts + 44181,
		Ftype1: int8('F'),
	},
	6373: {
		Fword:  __ccgo_ts + 44188,
		Ftype1: int8('F'),
	},
	6374: {
		Fword:  __ccgo_ts + 44195,
		Ftype1: int8('F'),
	},
	6375: {
		Fword:  __ccgo_ts + 44202,
		Ftype1: int8('F'),
	},
	6376: {
		Fword:  __ccgo_ts + 44209,
		Ftype1: int8('F'),
	},
	6377: {
		Fword:  __ccgo_ts + 44216,
		Ftype1: int8('F'),
	},
	6378: {
		Fword:  __ccgo_ts + 44223,
		Ftype1: int8('F'),
	},
	6379: {
		Fword:  __ccgo_ts + 44230,
		Ftype1: int8('F'),
	},
	6380: {
		Fword:  __ccgo_ts + 44237,
		Ftype1: int8('F'),
	},
	6381: {
		Fword:  __ccgo_ts + 44244,
		Ftype1: int8('F'),
	},
	6382: {
		Fword:  __ccgo_ts + 44251,
		Ftype1: int8('F'),
	},
	6383: {
		Fword:  __ccgo_ts + 44258,
		Ftype1: int8('F'),
	},
	6384: {
		Fword:  __ccgo_ts + 44265,
		Ftype1: int8('F'),
	},
	6385: {
		Fword:  __ccgo_ts + 44272,
		Ftype1: int8('F'),
	},
	6386: {
		Fword:  __ccgo_ts + 44279,
		Ftype1: int8('F'),
	},
	6387: {
		Fword:  __ccgo_ts + 44286,
		Ftype1: int8('F'),
	},
	6388: {
		Fword:  __ccgo_ts + 44293,
		Ftype1: int8('F'),
	},
	6389: {
		Fword:  __ccgo_ts + 44300,
		Ftype1: int8('F'),
	},
	6390: {
		Fword:  __ccgo_ts + 44307,
		Ftype1: int8('F'),
	},
	6391: {
		Fword:  __ccgo_ts + 44314,
		Ftype1: int8('F'),
	},
	6392: {
		Fword:  __ccgo_ts + 44321,
		Ftype1: int8('F'),
	},
	6393: {
		Fword:  __ccgo_ts + 44328,
		Ftype1: int8('F'),
	},
	6394: {
		Fword:  __ccgo_ts + 44335,
		Ftype1: int8('F'),
	},
	6395: {
		Fword:  __ccgo_ts + 44342,
		Ftype1: int8('F'),
	},
	6396: {
		Fword:  __ccgo_ts + 44349,
		Ftype1: int8('F'),
	},
	6397: {
		Fword:  __ccgo_ts + 44356,
		Ftype1: int8('F'),
	},
	6398: {
		Fword:  __ccgo_ts + 44363,
		Ftype1: int8('F'),
	},
	6399: {
		Fword:  __ccgo_ts + 44370,
		Ftype1: int8('F'),
	},
	6400: {
		Fword:  __ccgo_ts + 44377,
		Ftype1: int8('F'),
	},
	6401: {
		Fword:  __ccgo_ts + 44384,
		Ftype1: int8('F'),
	},
	6402: {
		Fword:  __ccgo_ts + 44391,
		Ftype1: int8('F'),
	},
	6403: {
		Fword:  __ccgo_ts + 44398,
		Ftype1: int8('F'),
	},
	6404: {
		Fword:  __ccgo_ts + 44405,
		Ftype1: int8('F'),
	},
	6405: {
		Fword:  __ccgo_ts + 44412,
		Ftype1: int8('F'),
	},
	6406: {
		Fword:  __ccgo_ts + 44419,
		Ftype1: int8('F'),
	},
	6407: {
		Fword:  __ccgo_ts + 44426,
		Ftype1: int8('F'),
	},
	6408: {
		Fword:  __ccgo_ts + 44433,
		Ftype1: int8('F'),
	},
	6409: {
		Fword:  __ccgo_ts + 44440,
		Ftype1: int8('F'),
	},
	6410: {
		Fword:  __ccgo_ts + 44447,
		Ftype1: int8('F'),
	},
	6411: {
		Fword:  __ccgo_ts + 44454,
		Ftype1: int8('F'),
	},
	6412: {
		Fword:  __ccgo_ts + 44461,
		Ftype1: int8('F'),
	},
	6413: {
		Fword:  __ccgo_ts + 44468,
		Ftype1: int8('F'),
	},
	6414: {
		Fword:  __ccgo_ts + 44475,
		Ftype1: int8('F'),
	},
	6415: {
		Fword:  __ccgo_ts + 44482,
		Ftype1: int8('F'),
	},
	6416: {
		Fword:  __ccgo_ts + 44489,
		Ftype1: int8('F'),
	},
	6417: {
		Fword:  __ccgo_ts + 44496,
		Ftype1: int8('F'),
	},
	6418: {
		Fword:  __ccgo_ts + 44503,
		Ftype1: int8('F'),
	},
	6419: {
		Fword:  __ccgo_ts + 44510,
		Ftype1: int8('F'),
	},
	6420: {
		Fword:  __ccgo_ts + 44517,
		Ftype1: int8('F'),
	},
	6421: {
		Fword:  __ccgo_ts + 44524,
		Ftype1: int8('F'),
	},
	6422: {
		Fword:  __ccgo_ts + 44531,
		Ftype1: int8('F'),
	},
	6423: {
		Fword:  __ccgo_ts + 44538,
		Ftype1: int8('F'),
	},
	6424: {
		Fword:  __ccgo_ts + 44545,
		Ftype1: int8('F'),
	},
	6425: {
		Fword:  __ccgo_ts + 44552,
		Ftype1: int8('F'),
	},
	6426: {
		Fword:  __ccgo_ts + 44559,
		Ftype1: int8('F'),
	},
	6427: {
		Fword:  __ccgo_ts + 44566,
		Ftype1: int8('F'),
	},
	6428: {
		Fword:  __ccgo_ts + 44573,
		Ftype1: int8('F'),
	},
	6429: {
		Fword:  __ccgo_ts + 44580,
		Ftype1: int8('F'),
	},
	6430: {
		Fword:  __ccgo_ts + 44587,
		Ftype1: int8('F'),
	},
	6431: {
		Fword:  __ccgo_ts + 44594,
		Ftype1: int8('F'),
	},
	6432: {
		Fword:  __ccgo_ts + 44601,
		Ftype1: int8('F'),
	},
	6433: {
		Fword:  __ccgo_ts + 44608,
		Ftype1: int8('F'),
	},
	6434: {
		Fword:  __ccgo_ts + 44615,
		Ftype1: int8('F'),
	},
	6435: {
		Fword:  __ccgo_ts + 44622,
		Ftype1: int8('F'),
	},
	6436: {
		Fword:  __ccgo_ts + 44629,
		Ftype1: int8('F'),
	},
	6437: {
		Fword:  __ccgo_ts + 44636,
		Ftype1: int8('F'),
	},
	6438: {
		Fword:  __ccgo_ts + 44643,
		Ftype1: int8('F'),
	},
	6439: {
		Fword:  __ccgo_ts + 44650,
		Ftype1: int8('F'),
	},
	6440: {
		Fword:  __ccgo_ts + 44657,
		Ftype1: int8('F'),
	},
	6441: {
		Fword:  __ccgo_ts + 44664,
		Ftype1: int8('F'),
	},
	6442: {
		Fword:  __ccgo_ts + 44671,
		Ftype1: int8('F'),
	},
	6443: {
		Fword:  __ccgo_ts + 44678,
		Ftype1: int8('F'),
	},
	6444: {
		Fword:  __ccgo_ts + 44685,
		Ftype1: int8('F'),
	},
	6445: {
		Fword:  __ccgo_ts + 44692,
		Ftype1: int8('F'),
	},
	6446: {
		Fword:  __ccgo_ts + 44699,
		Ftype1: int8('F'),
	},
	6447: {
		Fword:  __ccgo_ts + 44706,
		Ftype1: int8('F'),
	},
	6448: {
		Fword:  __ccgo_ts + 44713,
		Ftype1: int8('F'),
	},
	6449: {
		Fword:  __ccgo_ts + 44720,
		Ftype1: int8('F'),
	},
	6450: {
		Fword:  __ccgo_ts + 44726,
		Ftype1: int8('F'),
	},
	6451: {
		Fword:  __ccgo_ts + 44733,
		Ftype1: int8('F'),
	},
	6452: {
		Fword:  __ccgo_ts + 44740,
		Ftype1: int8('F'),
	},
	6453: {
		Fword:  __ccgo_ts + 44747,
		Ftype1: int8('F'),
	},
	6454: {
		Fword:  __ccgo_ts + 44754,
		Ftype1: int8('F'),
	},
	6455: {
		Fword:  __ccgo_ts + 44759,
		Ftype1: int8('F'),
	},
	6456: {
		Fword:  __ccgo_ts + 44765,
		Ftype1: int8('F'),
	},
	6457: {
		Fword:  __ccgo_ts + 44770,
		Ftype1: int8('F'),
	},
	6458: {
		Fword:  __ccgo_ts + 44775,
		Ftype1: int8('F'),
	},
	6459: {
		Fword:  __ccgo_ts + 44782,
		Ftype1: int8('F'),
	},
	6460: {
		Fword:  __ccgo_ts + 44789,
		Ftype1: int8('F'),
	},
	6461: {
		Fword:  __ccgo_ts + 44796,
		Ftype1: int8('F'),
	},
	6462: {
		Fword:  __ccgo_ts + 44803,
		Ftype1: int8('F'),
	},
	6463: {
		Fword:  __ccgo_ts + 44810,
		Ftype1: int8('F'),
	},
	6464: {
		Fword:  __ccgo_ts + 44817,
		Ftype1: int8('F'),
	},
	6465: {
		Fword:  __ccgo_ts + 44824,
		Ftype1: int8('F'),
	},
	6466: {
		Fword:  __ccgo_ts + 44830,
		Ftype1: int8('F'),
	},
	6467: {
		Fword:  __ccgo_ts + 44837,
		Ftype1: int8('F'),
	},
	6468: {
		Fword:  __ccgo_ts + 44844,
		Ftype1: int8('F'),
	},
	6469: {
		Fword:  __ccgo_ts + 44851,
		Ftype1: int8('F'),
	},
	6470: {
		Fword:  __ccgo_ts + 44858,
		Ftype1: int8('F'),
	},
	6471: {
		Fword:  __ccgo_ts + 44865,
		Ftype1: int8('F'),
	},
	6472: {
		Fword:  __ccgo_ts + 44872,
		Ftype1: int8('F'),
	},
	6473: {
		Fword:  __ccgo_ts + 44879,
		Ftype1: int8('F'),
	},
	6474: {
		Fword:  __ccgo_ts + 44886,
		Ftype1: int8('F'),
	},
	6475: {
		Fword:  __ccgo_ts + 44893,
		Ftype1: int8('F'),
	},
	6476: {
		Fword:  __ccgo_ts + 44900,
		Ftype1: int8('F'),
	},
	6477: {
		Fword:  __ccgo_ts + 44907,
		Ftype1: int8('F'),
	},
	6478: {
		Fword:  __ccgo_ts + 44914,
		Ftype1: int8('F'),
	},
	6479: {
		Fword:  __ccgo_ts + 44921,
		Ftype1: int8('F'),
	},
	6480: {
		Fword:  __ccgo_ts + 44928,
		Ftype1: int8('F'),
	},
	6481: {
		Fword:  __ccgo_ts + 44934,
		Ftype1: int8('F'),
	},
	6482: {
		Fword:  __ccgo_ts + 44941,
		Ftype1: int8('F'),
	},
	6483: {
		Fword:  __ccgo_ts + 44947,
		Ftype1: int8('F'),
	},
	6484: {
		Fword:  __ccgo_ts + 44953,
		Ftype1: int8('F'),
	},
	6485: {
		Fword:  __ccgo_ts + 44960,
		Ftype1: int8('F'),
	},
	6486: {
		Fword:  __ccgo_ts + 44967,
		Ftype1: int8('F'),
	},
	6487: {
		Fword:  __ccgo_ts + 44974,
		Ftype1: int8('F'),
	},
	6488: {
		Fword:  __ccgo_ts + 44981,
		Ftype1: int8('F'),
	},
	6489: {
		Fword:  __ccgo_ts + 44987,
		Ftype1: int8('F'),
	},
	6490: {
		Fword:  __ccgo_ts + 44994,
		Ftype1: int8('F'),
	},
	6491: {
		Fword:  __ccgo_ts + 45001,
		Ftype1: int8('F'),
	},
	6492: {
		Fword:  __ccgo_ts + 45008,
		Ftype1: int8('F'),
	},
	6493: {
		Fword:  __ccgo_ts + 45015,
		Ftype1: int8('F'),
	},
	6494: {
		Fword:  __ccgo_ts + 45022,
		Ftype1: int8('F'),
	},
	6495: {
		Fword:  __ccgo_ts + 45029,
		Ftype1: int8('F'),
	},
	6496: {
		Fword:  __ccgo_ts + 45036,
		Ftype1: int8('F'),
	},
	6497: {
		Fword:  __ccgo_ts + 45043,
		Ftype1: int8('F'),
	},
	6498: {
		Fword:  __ccgo_ts + 45050,
		Ftype1: int8('F'),
	},
	6499: {
		Fword:  __ccgo_ts + 45056,
		Ftype1: int8('F'),
	},
	6500: {
		Fword:  __ccgo_ts + 45063,
		Ftype1: int8('F'),
	},
	6501: {
		Fword:  __ccgo_ts + 45070,
		Ftype1: int8('F'),
	},
	6502: {
		Fword:  __ccgo_ts + 45077,
		Ftype1: int8('F'),
	},
	6503: {
		Fword:  __ccgo_ts + 45084,
		Ftype1: int8('F'),
	},
	6504: {
		Fword:  __ccgo_ts + 45091,
		Ftype1: int8('F'),
	},
	6505: {
		Fword:  __ccgo_ts + 45098,
		Ftype1: int8('F'),
	},
	6506: {
		Fword:  __ccgo_ts + 45105,
		Ftype1: int8('F'),
	},
	6507: {
		Fword:  __ccgo_ts + 45112,
		Ftype1: int8('F'),
	},
	6508: {
		Fword:  __ccgo_ts + 45119,
		Ftype1: int8('F'),
	},
	6509: {
		Fword:  __ccgo_ts + 45126,
		Ftype1: int8('F'),
	},
	6510: {
		Fword:  __ccgo_ts + 45133,
		Ftype1: int8('F'),
	},
	6511: {
		Fword:  __ccgo_ts + 45140,
		Ftype1: int8('F'),
	},
	6512: {
		Fword:  __ccgo_ts + 45147,
		Ftype1: int8('F'),
	},
	6513: {
		Fword:  __ccgo_ts + 45154,
		Ftype1: int8('F'),
	},
	6514: {
		Fword:  __ccgo_ts + 45161,
		Ftype1: int8('F'),
	},
	6515: {
		Fword:  __ccgo_ts + 45167,
		Ftype1: int8('F'),
	},
	6516: {
		Fword:  __ccgo_ts + 45174,
		Ftype1: int8('F'),
	},
	6517: {
		Fword:  __ccgo_ts + 45181,
		Ftype1: int8('F'),
	},
	6518: {
		Fword:  __ccgo_ts + 45188,
		Ftype1: int8('F'),
	},
	6519: {
		Fword:  __ccgo_ts + 45195,
		Ftype1: int8('F'),
	},
	6520: {
		Fword:  __ccgo_ts + 45202,
		Ftype1: int8('F'),
	},
	6521: {
		Fword:  __ccgo_ts + 45209,
		Ftype1: int8('F'),
	},
	6522: {
		Fword:  __ccgo_ts + 45216,
		Ftype1: int8('F'),
	},
	6523: {
		Fword:  __ccgo_ts + 45223,
		Ftype1: int8('F'),
	},
	6524: {
		Fword:  __ccgo_ts + 45230,
		Ftype1: int8('F'),
	},
	6525: {
		Fword:  __ccgo_ts + 45237,
		Ftype1: int8('F'),
	},
	6526: {
		Fword:  __ccgo_ts + 45244,
		Ftype1: int8('F'),
	},
	6527: {
		Fword:  __ccgo_ts + 45251,
		Ftype1: int8('F'),
	},
	6528: {
		Fword:  __ccgo_ts + 45258,
		Ftype1: int8('F'),
	},
	6529: {
		Fword:  __ccgo_ts + 45264,
		Ftype1: int8('F'),
	},
	6530: {
		Fword:  __ccgo_ts + 45271,
		Ftype1: int8('F'),
	},
	6531: {
		Fword:  __ccgo_ts + 45278,
		Ftype1: int8('F'),
	},
	6532: {
		Fword:  __ccgo_ts + 45285,
		Ftype1: int8('F'),
	},
	6533: {
		Fword:  __ccgo_ts + 45292,
		Ftype1: int8('F'),
	},
	6534: {
		Fword:  __ccgo_ts + 45299,
		Ftype1: int8('F'),
	},
	6535: {
		Fword:  __ccgo_ts + 45306,
		Ftype1: int8('F'),
	},
	6536: {
		Fword:  __ccgo_ts + 45313,
		Ftype1: int8('F'),
	},
	6537: {
		Fword:  __ccgo_ts + 45320,
		Ftype1: int8('F'),
	},
	6538: {
		Fword:  __ccgo_ts + 45327,
		Ftype1: int8('F'),
	},
	6539: {
		Fword:  __ccgo_ts + 45334,
		Ftype1: int8('F'),
	},
	6540: {
		Fword:  __ccgo_ts + 45341,
		Ftype1: int8('F'),
	},
	6541: {
		Fword:  __ccgo_ts + 45348,
		Ftype1: int8('F'),
	},
	6542: {
		Fword:  __ccgo_ts + 45355,
		Ftype1: int8('F'),
	},
	6543: {
		Fword:  __ccgo_ts + 45362,
		Ftype1: int8('F'),
	},
	6544: {
		Fword:  __ccgo_ts + 45369,
		Ftype1: int8('F'),
	},
	6545: {
		Fword:  __ccgo_ts + 45376,
		Ftype1: int8('F'),
	},
	6546: {
		Fword:  __ccgo_ts + 45383,
		Ftype1: int8('F'),
	},
	6547: {
		Fword:  __ccgo_ts + 45390,
		Ftype1: int8('F'),
	},
	6548: {
		Fword:  __ccgo_ts + 45397,
		Ftype1: int8('F'),
	},
	6549: {
		Fword:  __ccgo_ts + 45404,
		Ftype1: int8('F'),
	},
	6550: {
		Fword:  __ccgo_ts + 45411,
		Ftype1: int8('F'),
	},
	6551: {
		Fword:  __ccgo_ts + 45418,
		Ftype1: int8('F'),
	},
	6552: {
		Fword:  __ccgo_ts + 45425,
		Ftype1: int8('F'),
	},
	6553: {
		Fword:  __ccgo_ts + 45432,
		Ftype1: int8('F'),
	},
	6554: {
		Fword:  __ccgo_ts + 45439,
		Ftype1: int8('F'),
	},
	6555: {
		Fword:  __ccgo_ts + 45446,
		Ftype1: int8('F'),
	},
	6556: {
		Fword:  __ccgo_ts + 45453,
		Ftype1: int8('F'),
	},
	6557: {
		Fword:  __ccgo_ts + 45459,
		Ftype1: int8('F'),
	},
	6558: {
		Fword:  __ccgo_ts + 45466,
		Ftype1: int8('F'),
	},
	6559: {
		Fword:  __ccgo_ts + 45473,
		Ftype1: int8('F'),
	},
	6560: {
		Fword:  __ccgo_ts + 45480,
		Ftype1: int8('F'),
	},
	6561: {
		Fword:  __ccgo_ts + 45487,
		Ftype1: int8('F'),
	},
	6562: {
		Fword:  __ccgo_ts + 45494,
		Ftype1: int8('F'),
	},
	6563: {
		Fword:  __ccgo_ts + 45501,
		Ftype1: int8('F'),
	},
	6564: {
		Fword:  __ccgo_ts + 45508,
		Ftype1: int8('F'),
	},
	6565: {
		Fword:  __ccgo_ts + 45515,
		Ftype1: int8('F'),
	},
	6566: {
		Fword:  __ccgo_ts + 45522,
		Ftype1: int8('F'),
	},
	6567: {
		Fword:  __ccgo_ts + 45529,
		Ftype1: int8('F'),
	},
	6568: {
		Fword:  __ccgo_ts + 45536,
		Ftype1: int8('F'),
	},
	6569: {
		Fword:  __ccgo_ts + 45543,
		Ftype1: int8('F'),
	},
	6570: {
		Fword:  __ccgo_ts + 45550,
		Ftype1: int8('F'),
	},
	6571: {
		Fword:  __ccgo_ts + 45557,
		Ftype1: int8('F'),
	},
	6572: {
		Fword:  __ccgo_ts + 45563,
		Ftype1: int8('F'),
	},
	6573: {
		Fword:  __ccgo_ts + 45570,
		Ftype1: int8('F'),
	},
	6574: {
		Fword:  __ccgo_ts + 45577,
		Ftype1: int8('F'),
	},
	6575: {
		Fword:  __ccgo_ts + 45584,
		Ftype1: int8('F'),
	},
	6576: {
		Fword:  __ccgo_ts + 45591,
		Ftype1: int8('F'),
	},
	6577: {
		Fword:  __ccgo_ts + 45598,
		Ftype1: int8('F'),
	},
	6578: {
		Fword:  __ccgo_ts + 45605,
		Ftype1: int8('F'),
	},
	6579: {
		Fword:  __ccgo_ts + 45612,
		Ftype1: int8('F'),
	},
	6580: {
		Fword:  __ccgo_ts + 45619,
		Ftype1: int8('F'),
	},
	6581: {
		Fword:  __ccgo_ts + 45626,
		Ftype1: int8('F'),
	},
	6582: {
		Fword:  __ccgo_ts + 45633,
		Ftype1: int8('F'),
	},
	6583: {
		Fword:  __ccgo_ts + 45640,
		Ftype1: int8('F'),
	},
	6584: {
		Fword:  __ccgo_ts + 45647,
		Ftype1: int8('F'),
	},
	6585: {
		Fword:  __ccgo_ts + 45654,
		Ftype1: int8('F'),
	},
	6586: {
		Fword:  __ccgo_ts + 45661,
		Ftype1: int8('F'),
	},
	6587: {
		Fword:  __ccgo_ts + 45668,
		Ftype1: int8('F'),
	},
	6588: {
		Fword:  __ccgo_ts + 45675,
		Ftype1: int8('F'),
	},
	6589: {
		Fword:  __ccgo_ts + 45682,
		Ftype1: int8('F'),
	},
	6590: {
		Fword:  __ccgo_ts + 45689,
		Ftype1: int8('F'),
	},
	6591: {
		Fword:  __ccgo_ts + 45696,
		Ftype1: int8('F'),
	},
	6592: {
		Fword:  __ccgo_ts + 45703,
		Ftype1: int8('F'),
	},
	6593: {
		Fword:  __ccgo_ts + 45709,
		Ftype1: int8('F'),
	},
	6594: {
		Fword:  __ccgo_ts + 45716,
		Ftype1: int8('F'),
	},
	6595: {
		Fword:  __ccgo_ts + 45722,
		Ftype1: int8('F'),
	},
	6596: {
		Fword:  __ccgo_ts + 45729,
		Ftype1: int8('F'),
	},
	6597: {
		Fword:  __ccgo_ts + 45736,
		Ftype1: int8('F'),
	},
	6598: {
		Fword:  __ccgo_ts + 45743,
		Ftype1: int8('F'),
	},
	6599: {
		Fword:  __ccgo_ts + 45750,
		Ftype1: int8('F'),
	},
	6600: {
		Fword:  __ccgo_ts + 45757,
		Ftype1: int8('F'),
	},
	6601: {
		Fword:  __ccgo_ts + 45764,
		Ftype1: int8('F'),
	},
	6602: {
		Fword:  __ccgo_ts + 45770,
		Ftype1: int8('F'),
	},
	6603: {
		Fword:  __ccgo_ts + 45777,
		Ftype1: int8('F'),
	},
	6604: {
		Fword:  __ccgo_ts + 45784,
		Ftype1: int8('F'),
	},
	6605: {
		Fword:  __ccgo_ts + 45791,
		Ftype1: int8('F'),
	},
	6606: {
		Fword:  __ccgo_ts + 45798,
		Ftype1: int8('F'),
	},
	6607: {
		Fword:  __ccgo_ts + 45805,
		Ftype1: int8('F'),
	},
	6608: {
		Fword:  __ccgo_ts + 45812,
		Ftype1: int8('F'),
	},
	6609: {
		Fword:  __ccgo_ts + 45819,
		Ftype1: int8('F'),
	},
	6610: {
		Fword:  __ccgo_ts + 45826,
		Ftype1: int8('F'),
	},
	6611: {
		Fword:  __ccgo_ts + 45833,
		Ftype1: int8('F'),
	},
	6612: {
		Fword:  __ccgo_ts + 45840,
		Ftype1: int8('F'),
	},
	6613: {
		Fword:  __ccgo_ts + 45847,
		Ftype1: int8('F'),
	},
	6614: {
		Fword:  __ccgo_ts + 45854,
		Ftype1: int8('F'),
	},
	6615: {
		Fword:  __ccgo_ts + 45861,
		Ftype1: int8('F'),
	},
	6616: {
		Fword:  __ccgo_ts + 45868,
		Ftype1: int8('F'),
	},
	6617: {
		Fword:  __ccgo_ts + 45875,
		Ftype1: int8('F'),
	},
	6618: {
		Fword:  __ccgo_ts + 45882,
		Ftype1: int8('F'),
	},
	6619: {
		Fword:  __ccgo_ts + 45889,
		Ftype1: int8('F'),
	},
	6620: {
		Fword:  __ccgo_ts + 45896,
		Ftype1: int8('F'),
	},
	6621: {
		Fword:  __ccgo_ts + 45903,
		Ftype1: int8('F'),
	},
	6622: {
		Fword:  __ccgo_ts + 45910,
		Ftype1: int8('F'),
	},
	6623: {
		Fword:  __ccgo_ts + 45917,
		Ftype1: int8('F'),
	},
	6624: {
		Fword:  __ccgo_ts + 45924,
		Ftype1: int8('F'),
	},
	6625: {
		Fword:  __ccgo_ts + 45931,
		Ftype1: int8('F'),
	},
	6626: {
		Fword:  __ccgo_ts + 45938,
		Ftype1: int8('F'),
	},
	6627: {
		Fword:  __ccgo_ts + 45945,
		Ftype1: int8('F'),
	},
	6628: {
		Fword:  __ccgo_ts + 45952,
		Ftype1: int8('F'),
	},
	6629: {
		Fword:  __ccgo_ts + 45959,
		Ftype1: int8('F'),
	},
	6630: {
		Fword:  __ccgo_ts + 45966,
		Ftype1: int8('F'),
	},
	6631: {
		Fword:  __ccgo_ts + 45973,
		Ftype1: int8('F'),
	},
	6632: {
		Fword:  __ccgo_ts + 45980,
		Ftype1: int8('F'),
	},
	6633: {
		Fword:  __ccgo_ts + 45987,
		Ftype1: int8('F'),
	},
	6634: {
		Fword:  __ccgo_ts + 45994,
		Ftype1: int8('F'),
	},
	6635: {
		Fword:  __ccgo_ts + 46001,
		Ftype1: int8('F'),
	},
	6636: {
		Fword:  __ccgo_ts + 46008,
		Ftype1: int8('F'),
	},
	6637: {
		Fword:  __ccgo_ts + 46015,
		Ftype1: int8('F'),
	},
	6638: {
		Fword:  __ccgo_ts + 46022,
		Ftype1: int8('F'),
	},
	6639: {
		Fword:  __ccgo_ts + 46029,
		Ftype1: int8('F'),
	},
	6640: {
		Fword:  __ccgo_ts + 46036,
		Ftype1: int8('F'),
	},
	6641: {
		Fword:  __ccgo_ts + 46043,
		Ftype1: int8('F'),
	},
	6642: {
		Fword:  __ccgo_ts + 46050,
		Ftype1: int8('F'),
	},
	6643: {
		Fword:  __ccgo_ts + 46057,
		Ftype1: int8('F'),
	},
	6644: {
		Fword:  __ccgo_ts + 46064,
		Ftype1: int8('F'),
	},
	6645: {
		Fword:  __ccgo_ts + 46071,
		Ftype1: int8('F'),
	},
	6646: {
		Fword:  __ccgo_ts + 46078,
		Ftype1: int8('F'),
	},
	6647: {
		Fword:  __ccgo_ts + 46085,
		Ftype1: int8('F'),
	},
	6648: {
		Fword:  __ccgo_ts + 46092,
		Ftype1: int8('F'),
	},
	6649: {
		Fword:  __ccgo_ts + 46099,
		Ftype1: int8('F'),
	},
	6650: {
		Fword:  __ccgo_ts + 46106,
		Ftype1: int8('F'),
	},
	6651: {
		Fword:  __ccgo_ts + 46113,
		Ftype1: int8('F'),
	},
	6652: {
		Fword:  __ccgo_ts + 46120,
		Ftype1: int8('F'),
	},
	6653: {
		Fword:  __ccgo_ts + 46127,
		Ftype1: int8('F'),
	},
	6654: {
		Fword:  __ccgo_ts + 46134,
		Ftype1: int8('F'),
	},
	6655: {
		Fword:  __ccgo_ts + 46141,
		Ftype1: int8('F'),
	},
	6656: {
		Fword:  __ccgo_ts + 46148,
		Ftype1: int8('F'),
	},
	6657: {
		Fword:  __ccgo_ts + 46155,
		Ftype1: int8('F'),
	},
	6658: {
		Fword:  __ccgo_ts + 46162,
		Ftype1: int8('F'),
	},
	6659: {
		Fword:  __ccgo_ts + 46169,
		Ftype1: int8('F'),
	},
	6660: {
		Fword:  __ccgo_ts + 46176,
		Ftype1: int8('F'),
	},
	6661: {
		Fword:  __ccgo_ts + 46183,
		Ftype1: int8('F'),
	},
	6662: {
		Fword:  __ccgo_ts + 46190,
		Ftype1: int8('F'),
	},
	6663: {
		Fword:  __ccgo_ts + 46197,
		Ftype1: int8('F'),
	},
	6664: {
		Fword:  __ccgo_ts + 46204,
		Ftype1: int8('F'),
	},
	6665: {
		Fword:  __ccgo_ts + 46211,
		Ftype1: int8('F'),
	},
	6666: {
		Fword:  __ccgo_ts + 46218,
		Ftype1: int8('F'),
	},
	6667: {
		Fword:  __ccgo_ts + 46225,
		Ftype1: int8('F'),
	},
	6668: {
		Fword:  __ccgo_ts + 46232,
		Ftype1: int8('F'),
	},
	6669: {
		Fword:  __ccgo_ts + 46239,
		Ftype1: int8('F'),
	},
	6670: {
		Fword:  __ccgo_ts + 46246,
		Ftype1: int8('F'),
	},
	6671: {
		Fword:  __ccgo_ts + 46253,
		Ftype1: int8('F'),
	},
	6672: {
		Fword:  __ccgo_ts + 46260,
		Ftype1: int8('F'),
	},
	6673: {
		Fword:  __ccgo_ts + 46267,
		Ftype1: int8('F'),
	},
	6674: {
		Fword:  __ccgo_ts + 46274,
		Ftype1: int8('F'),
	},
	6675: {
		Fword:  __ccgo_ts + 46281,
		Ftype1: int8('F'),
	},
	6676: {
		Fword:  __ccgo_ts + 46288,
		Ftype1: int8('F'),
	},
	6677: {
		Fword:  __ccgo_ts + 46295,
		Ftype1: int8('F'),
	},
	6678: {
		Fword:  __ccgo_ts + 46302,
		Ftype1: int8('F'),
	},
	6679: {
		Fword:  __ccgo_ts + 46309,
		Ftype1: int8('F'),
	},
	6680: {
		Fword:  __ccgo_ts + 46316,
		Ftype1: int8('F'),
	},
	6681: {
		Fword:  __ccgo_ts + 46323,
		Ftype1: int8('F'),
	},
	6682: {
		Fword:  __ccgo_ts + 46330,
		Ftype1: int8('F'),
	},
	6683: {
		Fword:  __ccgo_ts + 46337,
		Ftype1: int8('F'),
	},
	6684: {
		Fword:  __ccgo_ts + 46344,
		Ftype1: int8('F'),
	},
	6685: {
		Fword:  __ccgo_ts + 46351,
		Ftype1: int8('F'),
	},
	6686: {
		Fword:  __ccgo_ts + 46358,
		Ftype1: int8('F'),
	},
	6687: {
		Fword:  __ccgo_ts + 46365,
		Ftype1: int8('F'),
	},
	6688: {
		Fword:  __ccgo_ts + 46372,
		Ftype1: int8('F'),
	},
	6689: {
		Fword:  __ccgo_ts + 46379,
		Ftype1: int8('F'),
	},
	6690: {
		Fword:  __ccgo_ts + 46386,
		Ftype1: int8('F'),
	},
	6691: {
		Fword:  __ccgo_ts + 46393,
		Ftype1: int8('F'),
	},
	6692: {
		Fword:  __ccgo_ts + 46400,
		Ftype1: int8('F'),
	},
	6693: {
		Fword:  __ccgo_ts + 46407,
		Ftype1: int8('F'),
	},
	6694: {
		Fword:  __ccgo_ts + 46414,
		Ftype1: int8('F'),
	},
	6695: {
		Fword:  __ccgo_ts + 46421,
		Ftype1: int8('F'),
	},
	6696: {
		Fword:  __ccgo_ts + 46428,
		Ftype1: int8('F'),
	},
	6697: {
		Fword:  __ccgo_ts + 46434,
		Ftype1: int8('F'),
	},
	6698: {
		Fword:  __ccgo_ts + 46441,
		Ftype1: int8('F'),
	},
	6699: {
		Fword:  __ccgo_ts + 46448,
		Ftype1: int8('F'),
	},
	6700: {
		Fword:  __ccgo_ts + 46455,
		Ftype1: int8('F'),
	},
	6701: {
		Fword:  __ccgo_ts + 46462,
		Ftype1: int8('F'),
	},
	6702: {
		Fword:  __ccgo_ts + 46469,
		Ftype1: int8('F'),
	},
	6703: {
		Fword:  __ccgo_ts + 46476,
		Ftype1: int8('F'),
	},
	6704: {
		Fword:  __ccgo_ts + 46483,
		Ftype1: int8('F'),
	},
	6705: {
		Fword:  __ccgo_ts + 46490,
		Ftype1: int8('F'),
	},
	6706: {
		Fword:  __ccgo_ts + 46497,
		Ftype1: int8('F'),
	},
	6707: {
		Fword:  __ccgo_ts + 46504,
		Ftype1: int8('F'),
	},
	6708: {
		Fword:  __ccgo_ts + 46511,
		Ftype1: int8('F'),
	},
	6709: {
		Fword:  __ccgo_ts + 46517,
		Ftype1: int8('F'),
	},
	6710: {
		Fword:  __ccgo_ts + 46524,
		Ftype1: int8('F'),
	},
	6711: {
		Fword:  __ccgo_ts + 46531,
		Ftype1: int8('F'),
	},
	6712: {
		Fword:  __ccgo_ts + 46538,
		Ftype1: int8('F'),
	},
	6713: {
		Fword:  __ccgo_ts + 46545,
		Ftype1: int8('F'),
	},
	6714: {
		Fword:  __ccgo_ts + 46552,
		Ftype1: int8('F'),
	},
	6715: {
		Fword:  __ccgo_ts + 46559,
		Ftype1: int8('F'),
	},
	6716: {
		Fword:  __ccgo_ts + 46566,
		Ftype1: int8('F'),
	},
	6717: {
		Fword:  __ccgo_ts + 46573,
		Ftype1: int8('F'),
	},
	6718: {
		Fword:  __ccgo_ts + 46580,
		Ftype1: int8('F'),
	},
	6719: {
		Fword:  __ccgo_ts + 46587,
		Ftype1: int8('F'),
	},
	6720: {
		Fword:  __ccgo_ts + 46594,
		Ftype1: int8('F'),
	},
	6721: {
		Fword:  __ccgo_ts + 46601,
		Ftype1: int8('F'),
	},
	6722: {
		Fword:  __ccgo_ts + 46608,
		Ftype1: int8('F'),
	},
	6723: {
		Fword:  __ccgo_ts + 46615,
		Ftype1: int8('F'),
	},
	6724: {
		Fword:  __ccgo_ts + 46621,
		Ftype1: int8('F'),
	},
	6725: {
		Fword:  __ccgo_ts + 46628,
		Ftype1: int8('F'),
	},
	6726: {
		Fword:  __ccgo_ts + 46635,
		Ftype1: int8('F'),
	},
	6727: {
		Fword:  __ccgo_ts + 46642,
		Ftype1: int8('F'),
	},
	6728: {
		Fword:  __ccgo_ts + 46649,
		Ftype1: int8('F'),
	},
	6729: {
		Fword:  __ccgo_ts + 46656,
		Ftype1: int8('F'),
	},
	6730: {
		Fword:  __ccgo_ts + 46663,
		Ftype1: int8('F'),
	},
	6731: {
		Fword:  __ccgo_ts + 46670,
		Ftype1: int8('F'),
	},
	6732: {
		Fword:  __ccgo_ts + 46676,
		Ftype1: int8('F'),
	},
	6733: {
		Fword:  __ccgo_ts + 46683,
		Ftype1: int8('F'),
	},
	6734: {
		Fword:  __ccgo_ts + 46690,
		Ftype1: int8('F'),
	},
	6735: {
		Fword:  __ccgo_ts + 46697,
		Ftype1: int8('F'),
	},
	6736: {
		Fword:  __ccgo_ts + 46704,
		Ftype1: int8('F'),
	},
	6737: {
		Fword:  __ccgo_ts + 46711,
		Ftype1: int8('F'),
	},
	6738: {
		Fword:  __ccgo_ts + 46718,
		Ftype1: int8('F'),
	},
	6739: {
		Fword:  __ccgo_ts + 46725,
		Ftype1: int8('F'),
	},
	6740: {
		Fword:  __ccgo_ts + 46732,
		Ftype1: int8('F'),
	},
	6741: {
		Fword:  __ccgo_ts + 46739,
		Ftype1: int8('F'),
	},
	6742: {
		Fword:  __ccgo_ts + 46746,
		Ftype1: int8('F'),
	},
	6743: {
		Fword:  __ccgo_ts + 46753,
		Ftype1: int8('F'),
	},
	6744: {
		Fword:  __ccgo_ts + 46760,
		Ftype1: int8('F'),
	},
	6745: {
		Fword:  __ccgo_ts + 46767,
		Ftype1: int8('F'),
	},
	6746: {
		Fword:  __ccgo_ts + 46774,
		Ftype1: int8('F'),
	},
	6747: {
		Fword:  __ccgo_ts + 46781,
		Ftype1: int8('F'),
	},
	6748: {
		Fword:  __ccgo_ts + 46788,
		Ftype1: int8('F'),
	},
	6749: {
		Fword:  __ccgo_ts + 46795,
		Ftype1: int8('F'),
	},
	6750: {
		Fword:  __ccgo_ts + 46802,
		Ftype1: int8('F'),
	},
	6751: {
		Fword:  __ccgo_ts + 46809,
		Ftype1: int8('F'),
	},
	6752: {
		Fword:  __ccgo_ts + 46816,
		Ftype1: int8('F'),
	},
	6753: {
		Fword:  __ccgo_ts + 46823,
		Ftype1: int8('F'),
	},
	6754: {
		Fword:  __ccgo_ts + 46830,
		Ftype1: int8('F'),
	},
	6755: {
		Fword:  __ccgo_ts + 46837,
		Ftype1: int8('F'),
	},
	6756: {
		Fword:  __ccgo_ts + 46844,
		Ftype1: int8('F'),
	},
	6757: {
		Fword:  __ccgo_ts + 46851,
		Ftype1: int8('F'),
	},
	6758: {
		Fword:  __ccgo_ts + 46858,
		Ftype1: int8('F'),
	},
	6759: {
		Fword:  __ccgo_ts + 46865,
		Ftype1: int8('F'),
	},
	6760: {
		Fword:  __ccgo_ts + 46872,
		Ftype1: int8('F'),
	},
	6761: {
		Fword:  __ccgo_ts + 46879,
		Ftype1: int8('F'),
	},
	6762: {
		Fword:  __ccgo_ts + 46886,
		Ftype1: int8('F'),
	},
	6763: {
		Fword:  __ccgo_ts + 46893,
		Ftype1: int8('F'),
	},
	6764: {
		Fword:  __ccgo_ts + 46900,
		Ftype1: int8('F'),
	},
	6765: {
		Fword:  __ccgo_ts + 46907,
		Ftype1: int8('F'),
	},
	6766: {
		Fword:  __ccgo_ts + 46914,
		Ftype1: int8('F'),
	},
	6767: {
		Fword:  __ccgo_ts + 46921,
		Ftype1: int8('F'),
	},
	6768: {
		Fword:  __ccgo_ts + 46928,
		Ftype1: int8('F'),
	},
	6769: {
		Fword:  __ccgo_ts + 46935,
		Ftype1: int8('F'),
	},
	6770: {
		Fword:  __ccgo_ts + 46942,
		Ftype1: int8('F'),
	},
	6771: {
		Fword:  __ccgo_ts + 46949,
		Ftype1: int8('F'),
	},
	6772: {
		Fword:  __ccgo_ts + 46956,
		Ftype1: int8('F'),
	},
	6773: {
		Fword:  __ccgo_ts + 46963,
		Ftype1: int8('F'),
	},
	6774: {
		Fword:  __ccgo_ts + 46970,
		Ftype1: int8('F'),
	},
	6775: {
		Fword:  __ccgo_ts + 46977,
		Ftype1: int8('F'),
	},
	6776: {
		Fword:  __ccgo_ts + 46984,
		Ftype1: int8('F'),
	},
	6777: {
		Fword:  __ccgo_ts + 46991,
		Ftype1: int8('F'),
	},
	6778: {
		Fword:  __ccgo_ts + 46998,
		Ftype1: int8('F'),
	},
	6779: {
		Fword:  __ccgo_ts + 47005,
		Ftype1: int8('F'),
	},
	6780: {
		Fword:  __ccgo_ts + 47012,
		Ftype1: int8('F'),
	},
	6781: {
		Fword:  __ccgo_ts + 47019,
		Ftype1: int8('F'),
	},
	6782: {
		Fword:  __ccgo_ts + 47026,
		Ftype1: int8('F'),
	},
	6783: {
		Fword:  __ccgo_ts + 47033,
		Ftype1: int8('F'),
	},
	6784: {
		Fword:  __ccgo_ts + 47040,
		Ftype1: int8('F'),
	},
	6785: {
		Fword:  __ccgo_ts + 47047,
		Ftype1: int8('F'),
	},
	6786: {
		Fword:  __ccgo_ts + 47054,
		Ftype1: int8('F'),
	},
	6787: {
		Fword:  __ccgo_ts + 47061,
		Ftype1: int8('F'),
	},
	6788: {
		Fword:  __ccgo_ts + 47068,
		Ftype1: int8('F'),
	},
	6789: {
		Fword:  __ccgo_ts + 47075,
		Ftype1: int8('F'),
	},
	6790: {
		Fword:  __ccgo_ts + 47082,
		Ftype1: int8('F'),
	},
	6791: {
		Fword:  __ccgo_ts + 47089,
		Ftype1: int8('F'),
	},
	6792: {
		Fword:  __ccgo_ts + 47096,
		Ftype1: int8('F'),
	},
	6793: {
		Fword:  __ccgo_ts + 47103,
		Ftype1: int8('F'),
	},
	6794: {
		Fword:  __ccgo_ts + 47110,
		Ftype1: int8('F'),
	},
	6795: {
		Fword:  __ccgo_ts + 47117,
		Ftype1: int8('F'),
	},
	6796: {
		Fword:  __ccgo_ts + 47124,
		Ftype1: int8('F'),
	},
	6797: {
		Fword:  __ccgo_ts + 47131,
		Ftype1: int8('F'),
	},
	6798: {
		Fword:  __ccgo_ts + 47138,
		Ftype1: int8('F'),
	},
	6799: {
		Fword:  __ccgo_ts + 47145,
		Ftype1: int8('F'),
	},
	6800: {
		Fword:  __ccgo_ts + 47152,
		Ftype1: int8('F'),
	},
	6801: {
		Fword:  __ccgo_ts + 47159,
		Ftype1: int8('F'),
	},
	6802: {
		Fword:  __ccgo_ts + 47164,
		Ftype1: int8('F'),
	},
	6803: {
		Fword:  __ccgo_ts + 47171,
		Ftype1: int8('F'),
	},
	6804: {
		Fword:  __ccgo_ts + 47178,
		Ftype1: int8('F'),
	},
	6805: {
		Fword:  __ccgo_ts + 47184,
		Ftype1: int8('F'),
	},
	6806: {
		Fword:  __ccgo_ts + 47191,
		Ftype1: int8('F'),
	},
	6807: {
		Fword:  __ccgo_ts + 47197,
		Ftype1: int8('F'),
	},
	6808: {
		Fword:  __ccgo_ts + 47204,
		Ftype1: int8('F'),
	},
	6809: {
		Fword:  __ccgo_ts + 47211,
		Ftype1: int8('F'),
	},
	6810: {
		Fword:  __ccgo_ts + 47218,
		Ftype1: int8('F'),
	},
	6811: {
		Fword:  __ccgo_ts + 47225,
		Ftype1: int8('F'),
	},
	6812: {
		Fword:  __ccgo_ts + 47232,
		Ftype1: int8('F'),
	},
	6813: {
		Fword:  __ccgo_ts + 47239,
		Ftype1: int8('F'),
	},
	6814: {
		Fword:  __ccgo_ts + 47246,
		Ftype1: int8('F'),
	},
	6815: {
		Fword:  __ccgo_ts + 47253,
		Ftype1: int8('F'),
	},
	6816: {
		Fword:  __ccgo_ts + 47260,
		Ftype1: int8('F'),
	},
	6817: {
		Fword:  __ccgo_ts + 47267,
		Ftype1: int8('F'),
	},
	6818: {
		Fword:  __ccgo_ts + 47274,
		Ftype1: int8('F'),
	},
	6819: {
		Fword:  __ccgo_ts + 47281,
		Ftype1: int8('F'),
	},
	6820: {
		Fword:  __ccgo_ts + 47288,
		Ftype1: int8('F'),
	},
	6821: {
		Fword:  __ccgo_ts + 47295,
		Ftype1: int8('F'),
	},
	6822: {
		Fword:  __ccgo_ts + 47302,
		Ftype1: int8('F'),
	},
	6823: {
		Fword:  __ccgo_ts + 47309,
		Ftype1: int8('F'),
	},
	6824: {
		Fword:  __ccgo_ts + 47316,
		Ftype1: int8('F'),
	},
	6825: {
		Fword:  __ccgo_ts + 47323,
		Ftype1: int8('F'),
	},
	6826: {
		Fword:  __ccgo_ts + 47330,
		Ftype1: int8('F'),
	},
	6827: {
		Fword:  __ccgo_ts + 47337,
		Ftype1: int8('F'),
	},
	6828: {
		Fword:  __ccgo_ts + 47344,
		Ftype1: int8('F'),
	},
	6829: {
		Fword:  __ccgo_ts + 47350,
		Ftype1: int8('F'),
	},
	6830: {
		Fword:  __ccgo_ts + 47357,
		Ftype1: int8('F'),
	},
	6831: {
		Fword:  __ccgo_ts + 47364,
		Ftype1: int8('F'),
	},
	6832: {
		Fword:  __ccgo_ts + 47371,
		Ftype1: int8('F'),
	},
	6833: {
		Fword:  __ccgo_ts + 47378,
		Ftype1: int8('F'),
	},
	6834: {
		Fword:  __ccgo_ts + 47385,
		Ftype1: int8('F'),
	},
	6835: {
		Fword:  __ccgo_ts + 47392,
		Ftype1: int8('F'),
	},
	6836: {
		Fword:  __ccgo_ts + 47398,
		Ftype1: int8('F'),
	},
	6837: {
		Fword:  __ccgo_ts + 47405,
		Ftype1: int8('F'),
	},
	6838: {
		Fword:  __ccgo_ts + 47412,
		Ftype1: int8('F'),
	},
	6839: {
		Fword:  __ccgo_ts + 47419,
		Ftype1: int8('F'),
	},
	6840: {
		Fword:  __ccgo_ts + 47426,
		Ftype1: int8('F'),
	},
	6841: {
		Fword:  __ccgo_ts + 47433,
		Ftype1: int8('F'),
	},
	6842: {
		Fword:  __ccgo_ts + 47439,
		Ftype1: int8('F'),
	},
	6843: {
		Fword:  __ccgo_ts + 47446,
		Ftype1: int8('F'),
	},
	6844: {
		Fword:  __ccgo_ts + 47453,
		Ftype1: int8('F'),
	},
	6845: {
		Fword:  __ccgo_ts + 47460,
		Ftype1: int8('F'),
	},
	6846: {
		Fword:  __ccgo_ts + 47467,
		Ftype1: int8('F'),
	},
	6847: {
		Fword:  __ccgo_ts + 47474,
		Ftype1: int8('F'),
	},
	6848: {
		Fword:  __ccgo_ts + 47480,
		Ftype1: int8('F'),
	},
	6849: {
		Fword:  __ccgo_ts + 47487,
		Ftype1: int8('F'),
	},
	6850: {
		Fword:  __ccgo_ts + 47494,
		Ftype1: int8('F'),
	},
	6851: {
		Fword:  __ccgo_ts + 47501,
		Ftype1: int8('F'),
	},
	6852: {
		Fword:  __ccgo_ts + 47508,
		Ftype1: int8('F'),
	},
	6853: {
		Fword:  __ccgo_ts + 47515,
		Ftype1: int8('F'),
	},
	6854: {
		Fword:  __ccgo_ts + 47522,
		Ftype1: int8('F'),
	},
	6855: {
		Fword:  __ccgo_ts + 47529,
		Ftype1: int8('F'),
	},
	6856: {
		Fword:  __ccgo_ts + 47536,
		Ftype1: int8('F'),
	},
	6857: {
		Fword:  __ccgo_ts + 47542,
		Ftype1: int8('F'),
	},
	6858: {
		Fword:  __ccgo_ts + 47549,
		Ftype1: int8('F'),
	},
	6859: {
		Fword:  __ccgo_ts + 47555,
		Ftype1: int8('F'),
	},
	6860: {
		Fword:  __ccgo_ts + 47562,
		Ftype1: int8('F'),
	},
	6861: {
		Fword:  __ccgo_ts + 47569,
		Ftype1: int8('F'),
	},
	6862: {
		Fword:  __ccgo_ts + 47576,
		Ftype1: int8('F'),
	},
	6863: {
		Fword:  __ccgo_ts + 47583,
		Ftype1: int8('F'),
	},
	6864: {
		Fword:  __ccgo_ts + 47590,
		Ftype1: int8('F'),
	},
	6865: {
		Fword:  __ccgo_ts + 47597,
		Ftype1: int8('F'),
	},
	6866: {
		Fword:  __ccgo_ts + 47604,
		Ftype1: int8('F'),
	},
	6867: {
		Fword:  __ccgo_ts + 47611,
		Ftype1: int8('F'),
	},
	6868: {
		Fword:  __ccgo_ts + 47618,
		Ftype1: int8('F'),
	},
	6869: {
		Fword:  __ccgo_ts + 47625,
		Ftype1: int8('F'),
	},
	6870: {
		Fword:  __ccgo_ts + 47630,
		Ftype1: int8('F'),
	},
	6871: {
		Fword:  __ccgo_ts + 47637,
		Ftype1: int8('F'),
	},
	6872: {
		Fword:  __ccgo_ts + 47644,
		Ftype1: int8('F'),
	},
	6873: {
		Fword:  __ccgo_ts + 47650,
		Ftype1: int8('F'),
	},
	6874: {
		Fword:  __ccgo_ts + 47657,
		Ftype1: int8('F'),
	},
	6875: {
		Fword:  __ccgo_ts + 47663,
		Ftype1: int8('F'),
	},
	6876: {
		Fword:  __ccgo_ts + 47670,
		Ftype1: int8('F'),
	},
	6877: {
		Fword:  __ccgo_ts + 47677,
		Ftype1: int8('F'),
	},
	6878: {
		Fword:  __ccgo_ts + 47684,
		Ftype1: int8('F'),
	},
	6879: {
		Fword:  __ccgo_ts + 47691,
		Ftype1: int8('F'),
	},
	6880: {
		Fword:  __ccgo_ts + 47698,
		Ftype1: int8('F'),
	},
	6881: {
		Fword:  __ccgo_ts + 47705,
		Ftype1: int8('F'),
	},
	6882: {
		Fword:  __ccgo_ts + 47712,
		Ftype1: int8('F'),
	},
	6883: {
		Fword:  __ccgo_ts + 47719,
		Ftype1: int8('F'),
	},
	6884: {
		Fword:  __ccgo_ts + 47726,
		Ftype1: int8('F'),
	},
	6885: {
		Fword:  __ccgo_ts + 47733,
		Ftype1: int8('F'),
	},
	6886: {
		Fword:  __ccgo_ts + 47740,
		Ftype1: int8('F'),
	},
	6887: {
		Fword:  __ccgo_ts + 47747,
		Ftype1: int8('F'),
	},
	6888: {
		Fword:  __ccgo_ts + 47752,
		Ftype1: int8('F'),
	},
	6889: {
		Fword:  __ccgo_ts + 47759,
		Ftype1: int8('F'),
	},
	6890: {
		Fword:  __ccgo_ts + 47766,
		Ftype1: int8('F'),
	},
	6891: {
		Fword:  __ccgo_ts + 47772,
		Ftype1: int8('F'),
	},
	6892: {
		Fword:  __ccgo_ts + 47779,
		Ftype1: int8('F'),
	},
	6893: {
		Fword:  __ccgo_ts + 47785,
		Ftype1: int8('F'),
	},
	6894: {
		Fword:  __ccgo_ts + 47792,
		Ftype1: int8('F'),
	},
	6895: {
		Fword:  __ccgo_ts + 47799,
		Ftype1: int8('F'),
	},
	6896: {
		Fword:  __ccgo_ts + 47806,
		Ftype1: int8('F'),
	},
	6897: {
		Fword:  __ccgo_ts + 47813,
		Ftype1: int8('F'),
	},
	6898: {
		Fword:  __ccgo_ts + 47820,
		Ftype1: int8('F'),
	},
	6899: {
		Fword:  __ccgo_ts + 47827,
		Ftype1: int8('F'),
	},
	6900: {
		Fword:  __ccgo_ts + 47834,
		Ftype1: int8('F'),
	},
	6901: {
		Fword:  __ccgo_ts + 47841,
		Ftype1: int8('F'),
	},
	6902: {
		Fword:  __ccgo_ts + 47848,
		Ftype1: int8('F'),
	},
	6903: {
		Fword:  __ccgo_ts + 47855,
		Ftype1: int8('F'),
	},
	6904: {
		Fword:  __ccgo_ts + 47862,
		Ftype1: int8('F'),
	},
	6905: {
		Fword:  __ccgo_ts + 47869,
		Ftype1: int8('F'),
	},
	6906: {
		Fword:  __ccgo_ts + 47876,
		Ftype1: int8('F'),
	},
	6907: {
		Fword:  __ccgo_ts + 47883,
		Ftype1: int8('F'),
	},
	6908: {
		Fword:  __ccgo_ts + 47890,
		Ftype1: int8('F'),
	},
	6909: {
		Fword:  __ccgo_ts + 47897,
		Ftype1: int8('F'),
	},
	6910: {
		Fword:  __ccgo_ts + 47904,
		Ftype1: int8('F'),
	},
	6911: {
		Fword:  __ccgo_ts + 47911,
		Ftype1: int8('F'),
	},
	6912: {
		Fword:  __ccgo_ts + 47918,
		Ftype1: int8('F'),
	},
	6913: {
		Fword:  __ccgo_ts + 47925,
		Ftype1: int8('F'),
	},
	6914: {
		Fword:  __ccgo_ts + 47932,
		Ftype1: int8('F'),
	},
	6915: {
		Fword:  __ccgo_ts + 47939,
		Ftype1: int8('F'),
	},
	6916: {
		Fword:  __ccgo_ts + 47946,
		Ftype1: int8('F'),
	},
	6917: {
		Fword:  __ccgo_ts + 47953,
		Ftype1: int8('F'),
	},
	6918: {
		Fword:  __ccgo_ts + 47960,
		Ftype1: int8('F'),
	},
	6919: {
		Fword:  __ccgo_ts + 47967,
		Ftype1: int8('F'),
	},
	6920: {
		Fword:  __ccgo_ts + 47974,
		Ftype1: int8('F'),
	},
	6921: {
		Fword:  __ccgo_ts + 47981,
		Ftype1: int8('F'),
	},
	6922: {
		Fword:  __ccgo_ts + 47988,
		Ftype1: int8('F'),
	},
	6923: {
		Fword:  __ccgo_ts + 47995,
		Ftype1: int8('F'),
	},
	6924: {
		Fword:  __ccgo_ts + 48002,
		Ftype1: int8('F'),
	},
	6925: {
		Fword:  __ccgo_ts + 48009,
		Ftype1: int8('F'),
	},
	6926: {
		Fword:  __ccgo_ts + 48016,
		Ftype1: int8('F'),
	},
	6927: {
		Fword:  __ccgo_ts + 48023,
		Ftype1: int8('F'),
	},
	6928: {
		Fword:  __ccgo_ts + 48030,
		Ftype1: int8('F'),
	},
	6929: {
		Fword:  __ccgo_ts + 48037,
		Ftype1: int8('F'),
	},
	6930: {
		Fword:  __ccgo_ts + 48044,
		Ftype1: int8('F'),
	},
	6931: {
		Fword:  __ccgo_ts + 48051,
		Ftype1: int8('F'),
	},
	6932: {
		Fword:  __ccgo_ts + 48058,
		Ftype1: int8('F'),
	},
	6933: {
		Fword:  __ccgo_ts + 48063,
		Ftype1: int8('F'),
	},
	6934: {
		Fword:  __ccgo_ts + 48070,
		Ftype1: int8('F'),
	},
	6935: {
		Fword:  __ccgo_ts + 48077,
		Ftype1: int8('F'),
	},
	6936: {
		Fword:  __ccgo_ts + 48084,
		Ftype1: int8('F'),
	},
	6937: {
		Fword:  __ccgo_ts + 48091,
		Ftype1: int8('F'),
	},
	6938: {
		Fword:  __ccgo_ts + 48098,
		Ftype1: int8('F'),
	},
	6939: {
		Fword:  __ccgo_ts + 48105,
		Ftype1: int8('F'),
	},
	6940: {
		Fword:  __ccgo_ts + 48112,
		Ftype1: int8('F'),
	},
	6941: {
		Fword:  __ccgo_ts + 48119,
		Ftype1: int8('F'),
	},
	6942: {
		Fword:  __ccgo_ts + 48126,
		Ftype1: int8('F'),
	},
	6943: {
		Fword:  __ccgo_ts + 48133,
		Ftype1: int8('F'),
	},
	6944: {
		Fword:  __ccgo_ts + 48139,
		Ftype1: int8('F'),
	},
	6945: {
		Fword:  __ccgo_ts + 48146,
		Ftype1: int8('F'),
	},
	6946: {
		Fword:  __ccgo_ts + 48153,
		Ftype1: int8('F'),
	},
	6947: {
		Fword:  __ccgo_ts + 48160,
		Ftype1: int8('F'),
	},
	6948: {
		Fword:  __ccgo_ts + 48167,
		Ftype1: int8('F'),
	},
	6949: {
		Fword:  __ccgo_ts + 48174,
		Ftype1: int8('F'),
	},
	6950: {
		Fword:  __ccgo_ts + 48181,
		Ftype1: int8('F'),
	},
	6951: {
		Fword:  __ccgo_ts + 48188,
		Ftype1: int8('F'),
	},
	6952: {
		Fword:  __ccgo_ts + 48195,
		Ftype1: int8('F'),
	},
	6953: {
		Fword:  __ccgo_ts + 48202,
		Ftype1: int8('F'),
	},
	6954: {
		Fword:  __ccgo_ts + 48208,
		Ftype1: int8('F'),
	},
	6955: {
		Fword:  __ccgo_ts + 48215,
		Ftype1: int8('F'),
	},
	6956: {
		Fword:  __ccgo_ts + 48222,
		Ftype1: int8('F'),
	},
	6957: {
		Fword:  __ccgo_ts + 48229,
		Ftype1: int8('F'),
	},
	6958: {
		Fword:  __ccgo_ts + 48236,
		Ftype1: int8('F'),
	},
	6959: {
		Fword:  __ccgo_ts + 48243,
		Ftype1: int8('F'),
	},
	6960: {
		Fword:  __ccgo_ts + 48250,
		Ftype1: int8('F'),
	},
	6961: {
		Fword:  __ccgo_ts + 48257,
		Ftype1: int8('F'),
	},
	6962: {
		Fword:  __ccgo_ts + 48264,
		Ftype1: int8('F'),
	},
	6963: {
		Fword:  __ccgo_ts + 48271,
		Ftype1: int8('F'),
	},
	6964: {
		Fword:  __ccgo_ts + 48278,
		Ftype1: int8('F'),
	},
	6965: {
		Fword:  __ccgo_ts + 48285,
		Ftype1: int8('F'),
	},
	6966: {
		Fword:  __ccgo_ts + 48292,
		Ftype1: int8('F'),
	},
	6967: {
		Fword:  __ccgo_ts + 48299,
		Ftype1: int8('F'),
	},
	6968: {
		Fword:  __ccgo_ts + 48306,
		Ftype1: int8('F'),
	},
	6969: {
		Fword:  __ccgo_ts + 48312,
		Ftype1: int8('F'),
	},
	6970: {
		Fword:  __ccgo_ts + 48319,
		Ftype1: int8('F'),
	},
	6971: {
		Fword:  __ccgo_ts + 48326,
		Ftype1: int8('F'),
	},
	6972: {
		Fword:  __ccgo_ts + 48333,
		Ftype1: int8('F'),
	},
	6973: {
		Fword:  __ccgo_ts + 48340,
		Ftype1: int8('F'),
	},
	6974: {
		Fword:  __ccgo_ts + 48347,
		Ftype1: int8('F'),
	},
	6975: {
		Fword:  __ccgo_ts + 48354,
		Ftype1: int8('F'),
	},
	6976: {
		Fword:  __ccgo_ts + 48361,
		Ftype1: int8('F'),
	},
	6977: {
		Fword:  __ccgo_ts + 48368,
		Ftype1: int8('F'),
	},
	6978: {
		Fword:  __ccgo_ts + 48375,
		Ftype1: int8('F'),
	},
	6979: {
		Fword:  __ccgo_ts + 48382,
		Ftype1: int8('F'),
	},
	6980: {
		Fword:  __ccgo_ts + 48388,
		Ftype1: int8('F'),
	},
	6981: {
		Fword:  __ccgo_ts + 48395,
		Ftype1: int8('F'),
	},
	6982: {
		Fword:  __ccgo_ts + 48402,
		Ftype1: int8('F'),
	},
	6983: {
		Fword:  __ccgo_ts + 48409,
		Ftype1: int8('F'),
	},
	6984: {
		Fword:  __ccgo_ts + 48416,
		Ftype1: int8('F'),
	},
	6985: {
		Fword:  __ccgo_ts + 48423,
		Ftype1: int8('F'),
	},
	6986: {
		Fword:  __ccgo_ts + 48430,
		Ftype1: int8('F'),
	},
	6987: {
		Fword:  __ccgo_ts + 48437,
		Ftype1: int8('F'),
	},
	6988: {
		Fword:  __ccgo_ts + 48444,
		Ftype1: int8('F'),
	},
	6989: {
		Fword:  __ccgo_ts + 48451,
		Ftype1: int8('F'),
	},
	6990: {
		Fword:  __ccgo_ts + 48458,
		Ftype1: int8('F'),
	},
	6991: {
		Fword:  __ccgo_ts + 48465,
		Ftype1: int8('F'),
	},
	6992: {
		Fword:  __ccgo_ts + 48471,
		Ftype1: int8('F'),
	},
	6993: {
		Fword:  __ccgo_ts + 48478,
		Ftype1: int8('F'),
	},
	6994: {
		Fword:  __ccgo_ts + 48485,
		Ftype1: int8('F'),
	},
	6995: {
		Fword:  __ccgo_ts + 48492,
		Ftype1: int8('F'),
	},
	6996: {
		Fword:  __ccgo_ts + 48499,
		Ftype1: int8('F'),
	},
	6997: {
		Fword:  __ccgo_ts + 48505,
		Ftype1: int8('F'),
	},
	6998: {
		Fword:  __ccgo_ts + 48512,
		Ftype1: int8('F'),
	},
	6999: {
		Fword:  __ccgo_ts + 48519,
		Ftype1: int8('F'),
	},
	7000: {
		Fword:  __ccgo_ts + 48526,
		Ftype1: int8('F'),
	},
	7001: {
		Fword:  __ccgo_ts + 48533,
		Ftype1: int8('F'),
	},
	7002: {
		Fword:  __ccgo_ts + 48539,
		Ftype1: int8('F'),
	},
	7003: {
		Fword:  __ccgo_ts + 48546,
		Ftype1: int8('F'),
	},
	7004: {
		Fword:  __ccgo_ts + 48553,
		Ftype1: int8('F'),
	},
	7005: {
		Fword:  __ccgo_ts + 48560,
		Ftype1: int8('F'),
	},
	7006: {
		Fword:  __ccgo_ts + 48567,
		Ftype1: int8('F'),
	},
	7007: {
		Fword:  __ccgo_ts + 48574,
		Ftype1: int8('F'),
	},
	7008: {
		Fword:  __ccgo_ts + 48581,
		Ftype1: int8('F'),
	},
	7009: {
		Fword:  __ccgo_ts + 48588,
		Ftype1: int8('F'),
	},
	7010: {
		Fword:  __ccgo_ts + 48595,
		Ftype1: int8('F'),
	},
	7011: {
		Fword:  __ccgo_ts + 48602,
		Ftype1: int8('F'),
	},
	7012: {
		Fword:  __ccgo_ts + 48609,
		Ftype1: int8('F'),
	},
	7013: {
		Fword:  __ccgo_ts + 48616,
		Ftype1: int8('F'),
	},
	7014: {
		Fword:  __ccgo_ts + 48623,
		Ftype1: int8('F'),
	},
	7015: {
		Fword:  __ccgo_ts + 48630,
		Ftype1: int8('F'),
	},
	7016: {
		Fword:  __ccgo_ts + 48637,
		Ftype1: int8('F'),
	},
	7017: {
		Fword:  __ccgo_ts + 48644,
		Ftype1: int8('F'),
	},
	7018: {
		Fword:  __ccgo_ts + 48651,
		Ftype1: int8('F'),
	},
	7019: {
		Fword:  __ccgo_ts + 48658,
		Ftype1: int8('F'),
	},
	7020: {
		Fword:  __ccgo_ts + 48665,
		Ftype1: int8('F'),
	},
	7021: {
		Fword:  __ccgo_ts + 48672,
		Ftype1: int8('F'),
	},
	7022: {
		Fword:  __ccgo_ts + 48679,
		Ftype1: int8('F'),
	},
	7023: {
		Fword:  __ccgo_ts + 48686,
		Ftype1: int8('F'),
	},
	7024: {
		Fword:  __ccgo_ts + 48693,
		Ftype1: int8('F'),
	},
	7025: {
		Fword:  __ccgo_ts + 48700,
		Ftype1: int8('F'),
	},
	7026: {
		Fword:  __ccgo_ts + 48706,
		Ftype1: int8('F'),
	},
	7027: {
		Fword:  __ccgo_ts + 48713,
		Ftype1: int8('F'),
	},
	7028: {
		Fword:  __ccgo_ts + 48720,
		Ftype1: int8('F'),
	},
	7029: {
		Fword:  __ccgo_ts + 48727,
		Ftype1: int8('F'),
	},
	7030: {
		Fword:  __ccgo_ts + 48734,
		Ftype1: int8('F'),
	},
	7031: {
		Fword:  __ccgo_ts + 48741,
		Ftype1: int8('F'),
	},
	7032: {
		Fword:  __ccgo_ts + 48748,
		Ftype1: int8('F'),
	},
	7033: {
		Fword:  __ccgo_ts + 48755,
		Ftype1: int8('F'),
	},
	7034: {
		Fword:  __ccgo_ts + 48762,
		Ftype1: int8('F'),
	},
	7035: {
		Fword:  __ccgo_ts + 48769,
		Ftype1: int8('F'),
	},
	7036: {
		Fword:  __ccgo_ts + 48776,
		Ftype1: int8('F'),
	},
	7037: {
		Fword:  __ccgo_ts + 48783,
		Ftype1: int8('F'),
	},
	7038: {
		Fword:  __ccgo_ts + 48788,
		Ftype1: int8('F'),
	},
	7039: {
		Fword:  __ccgo_ts + 48795,
		Ftype1: int8('F'),
	},
	7040: {
		Fword:  __ccgo_ts + 48802,
		Ftype1: int8('F'),
	},
	7041: {
		Fword:  __ccgo_ts + 48809,
		Ftype1: int8('F'),
	},
	7042: {
		Fword:  __ccgo_ts + 48816,
		Ftype1: int8('F'),
	},
	7043: {
		Fword:  __ccgo_ts + 48823,
		Ftype1: int8('F'),
	},
	7044: {
		Fword:  __ccgo_ts + 48830,
		Ftype1: int8('F'),
	},
	7045: {
		Fword:  __ccgo_ts + 48837,
		Ftype1: int8('F'),
	},
	7046: {
		Fword:  __ccgo_ts + 48844,
		Ftype1: int8('F'),
	},
	7047: {
		Fword:  __ccgo_ts + 48851,
		Ftype1: int8('F'),
	},
	7048: {
		Fword:  __ccgo_ts + 48858,
		Ftype1: int8('F'),
	},
	7049: {
		Fword:  __ccgo_ts + 48864,
		Ftype1: int8('F'),
	},
	7050: {
		Fword:  __ccgo_ts + 48871,
		Ftype1: int8('F'),
	},
	7051: {
		Fword:  __ccgo_ts + 48878,
		Ftype1: int8('F'),
	},
	7052: {
		Fword:  __ccgo_ts + 48885,
		Ftype1: int8('F'),
	},
	7053: {
		Fword:  __ccgo_ts + 48892,
		Ftype1: int8('F'),
	},
	7054: {
		Fword:  __ccgo_ts + 48899,
		Ftype1: int8('F'),
	},
	7055: {
		Fword:  __ccgo_ts + 48906,
		Ftype1: int8('F'),
	},
	7056: {
		Fword:  __ccgo_ts + 48913,
		Ftype1: int8('F'),
	},
	7057: {
		Fword:  __ccgo_ts + 48920,
		Ftype1: int8('F'),
	},
	7058: {
		Fword:  __ccgo_ts + 48927,
		Ftype1: int8('F'),
	},
	7059: {
		Fword:  __ccgo_ts + 48933,
		Ftype1: int8('F'),
	},
	7060: {
		Fword:  __ccgo_ts + 48940,
		Ftype1: int8('F'),
	},
	7061: {
		Fword:  __ccgo_ts + 48947,
		Ftype1: int8('F'),
	},
	7062: {
		Fword:  __ccgo_ts + 48954,
		Ftype1: int8('F'),
	},
	7063: {
		Fword:  __ccgo_ts + 48961,
		Ftype1: int8('F'),
	},
	7064: {
		Fword:  __ccgo_ts + 48968,
		Ftype1: int8('F'),
	},
	7065: {
		Fword:  __ccgo_ts + 48975,
		Ftype1: int8('F'),
	},
	7066: {
		Fword:  __ccgo_ts + 48982,
		Ftype1: int8('F'),
	},
	7067: {
		Fword:  __ccgo_ts + 48989,
		Ftype1: int8('F'),
	},
	7068: {
		Fword:  __ccgo_ts + 48996,
		Ftype1: int8('F'),
	},
	7069: {
		Fword:  __ccgo_ts + 49003,
		Ftype1: int8('F'),
	},
	7070: {
		Fword:  __ccgo_ts + 49010,
		Ftype1: int8('F'),
	},
	7071: {
		Fword:  __ccgo_ts + 49017,
		Ftype1: int8('F'),
	},
	7072: {
		Fword:  __ccgo_ts + 49024,
		Ftype1: int8('F'),
	},
	7073: {
		Fword:  __ccgo_ts + 49030,
		Ftype1: int8('F'),
	},
	7074: {
		Fword:  __ccgo_ts + 49037,
		Ftype1: int8('F'),
	},
	7075: {
		Fword:  __ccgo_ts + 49044,
		Ftype1: int8('F'),
	},
	7076: {
		Fword:  __ccgo_ts + 49051,
		Ftype1: int8('F'),
	},
	7077: {
		Fword:  __ccgo_ts + 49058,
		Ftype1: int8('F'),
	},
	7078: {
		Fword:  __ccgo_ts + 49063,
		Ftype1: int8('F'),
	},
	7079: {
		Fword:  __ccgo_ts + 49070,
		Ftype1: int8('F'),
	},
	7080: {
		Fword:  __ccgo_ts + 49077,
		Ftype1: int8('F'),
	},
	7081: {
		Fword:  __ccgo_ts + 49084,
		Ftype1: int8('F'),
	},
	7082: {
		Fword:  __ccgo_ts + 49091,
		Ftype1: int8('F'),
	},
	7083: {
		Fword:  __ccgo_ts + 49098,
		Ftype1: int8('F'),
	},
	7084: {
		Fword:  __ccgo_ts + 49105,
		Ftype1: int8('F'),
	},
	7085: {
		Fword:  __ccgo_ts + 49112,
		Ftype1: int8('F'),
	},
	7086: {
		Fword:  __ccgo_ts + 49119,
		Ftype1: int8('F'),
	},
	7087: {
		Fword:  __ccgo_ts + 49126,
		Ftype1: int8('F'),
	},
	7088: {
		Fword:  __ccgo_ts + 49133,
		Ftype1: int8('F'),
	},
	7089: {
		Fword:  __ccgo_ts + 49139,
		Ftype1: int8('F'),
	},
	7090: {
		Fword:  __ccgo_ts + 49146,
		Ftype1: int8('F'),
	},
	7091: {
		Fword:  __ccgo_ts + 49153,
		Ftype1: int8('F'),
	},
	7092: {
		Fword:  __ccgo_ts + 49159,
		Ftype1: int8('F'),
	},
	7093: {
		Fword:  __ccgo_ts + 49166,
		Ftype1: int8('F'),
	},
	7094: {
		Fword:  __ccgo_ts + 49173,
		Ftype1: int8('F'),
	},
	7095: {
		Fword:  __ccgo_ts + 49180,
		Ftype1: int8('F'),
	},
	7096: {
		Fword:  __ccgo_ts + 49187,
		Ftype1: int8('F'),
	},
	7097: {
		Fword:  __ccgo_ts + 49194,
		Ftype1: int8('F'),
	},
	7098: {
		Fword:  __ccgo_ts + 49201,
		Ftype1: int8('F'),
	},
	7099: {
		Fword:  __ccgo_ts + 49208,
		Ftype1: int8('F'),
	},
	7100: {
		Fword:  __ccgo_ts + 49215,
		Ftype1: int8('F'),
	},
	7101: {
		Fword:  __ccgo_ts + 49222,
		Ftype1: int8('F'),
	},
	7102: {
		Fword:  __ccgo_ts + 49228,
		Ftype1: int8('F'),
	},
	7103: {
		Fword:  __ccgo_ts + 49235,
		Ftype1: int8('F'),
	},
	7104: {
		Fword:  __ccgo_ts + 49242,
		Ftype1: int8('F'),
	},
	7105: {
		Fword:  __ccgo_ts + 49249,
		Ftype1: int8('F'),
	},
	7106: {
		Fword:  __ccgo_ts + 49256,
		Ftype1: int8('F'),
	},
	7107: {
		Fword:  __ccgo_ts + 49263,
		Ftype1: int8('F'),
	},
	7108: {
		Fword:  __ccgo_ts + 49270,
		Ftype1: int8('F'),
	},
	7109: {
		Fword:  __ccgo_ts + 49277,
		Ftype1: int8('F'),
	},
	7110: {
		Fword:  __ccgo_ts + 49284,
		Ftype1: int8('F'),
	},
	7111: {
		Fword:  __ccgo_ts + 49291,
		Ftype1: int8('F'),
	},
	7112: {
		Fword:  __ccgo_ts + 49298,
		Ftype1: int8('F'),
	},
	7113: {
		Fword:  __ccgo_ts + 49305,
		Ftype1: int8('F'),
	},
	7114: {
		Fword:  __ccgo_ts + 49312,
		Ftype1: int8('F'),
	},
	7115: {
		Fword:  __ccgo_ts + 49319,
		Ftype1: int8('F'),
	},
	7116: {
		Fword:  __ccgo_ts + 49326,
		Ftype1: int8('F'),
	},
	7117: {
		Fword:  __ccgo_ts + 49333,
		Ftype1: int8('F'),
	},
	7118: {
		Fword:  __ccgo_ts + 49340,
		Ftype1: int8('F'),
	},
	7119: {
		Fword:  __ccgo_ts + 49346,
		Ftype1: int8('F'),
	},
	7120: {
		Fword:  __ccgo_ts + 49353,
		Ftype1: int8('F'),
	},
	7121: {
		Fword:  __ccgo_ts + 49360,
		Ftype1: int8('F'),
	},
	7122: {
		Fword:  __ccgo_ts + 49367,
		Ftype1: int8('F'),
	},
	7123: {
		Fword:  __ccgo_ts + 49374,
		Ftype1: int8('F'),
	},
	7124: {
		Fword:  __ccgo_ts + 49380,
		Ftype1: int8('F'),
	},
	7125: {
		Fword:  __ccgo_ts + 49387,
		Ftype1: int8('F'),
	},
	7126: {
		Fword:  __ccgo_ts + 49394,
		Ftype1: int8('F'),
	},
	7127: {
		Fword:  __ccgo_ts + 49401,
		Ftype1: int8('F'),
	},
	7128: {
		Fword:  __ccgo_ts + 49406,
		Ftype1: int8('F'),
	},
	7129: {
		Fword:  __ccgo_ts + 49413,
		Ftype1: int8('F'),
	},
	7130: {
		Fword:  __ccgo_ts + 49420,
		Ftype1: int8('F'),
	},
	7131: {
		Fword:  __ccgo_ts + 49427,
		Ftype1: int8('F'),
	},
	7132: {
		Fword:  __ccgo_ts + 49434,
		Ftype1: int8('F'),
	},
	7133: {
		Fword:  __ccgo_ts + 49441,
		Ftype1: int8('F'),
	},
	7134: {
		Fword:  __ccgo_ts + 49448,
		Ftype1: int8('F'),
	},
	7135: {
		Fword:  __ccgo_ts + 49455,
		Ftype1: int8('F'),
	},
	7136: {
		Fword:  __ccgo_ts + 49462,
		Ftype1: int8('F'),
	},
	7137: {
		Fword:  __ccgo_ts + 49469,
		Ftype1: int8('F'),
	},
	7138: {
		Fword:  __ccgo_ts + 49476,
		Ftype1: int8('F'),
	},
	7139: {
		Fword:  __ccgo_ts + 49482,
		Ftype1: int8('F'),
	},
	7140: {
		Fword:  __ccgo_ts + 49489,
		Ftype1: int8('F'),
	},
	7141: {
		Fword:  __ccgo_ts + 49496,
		Ftype1: int8('F'),
	},
	7142: {
		Fword:  __ccgo_ts + 49503,
		Ftype1: int8('F'),
	},
	7143: {
		Fword:  __ccgo_ts + 49510,
		Ftype1: int8('F'),
	},
	7144: {
		Fword:  __ccgo_ts + 49517,
		Ftype1: int8('F'),
	},
	7145: {
		Fword:  __ccgo_ts + 49524,
		Ftype1: int8('F'),
	},
	7146: {
		Fword:  __ccgo_ts + 49531,
		Ftype1: int8('F'),
	},
	7147: {
		Fword:  __ccgo_ts + 49538,
		Ftype1: int8('F'),
	},
	7148: {
		Fword:  __ccgo_ts + 49545,
		Ftype1: int8('F'),
	},
	7149: {
		Fword:  __ccgo_ts + 49551,
		Ftype1: int8('F'),
	},
	7150: {
		Fword:  __ccgo_ts + 49558,
		Ftype1: int8('F'),
	},
	7151: {
		Fword:  __ccgo_ts + 49565,
		Ftype1: int8('F'),
	},
	7152: {
		Fword:  __ccgo_ts + 49572,
		Ftype1: int8('F'),
	},
	7153: {
		Fword:  __ccgo_ts + 49579,
		Ftype1: int8('F'),
	},
	7154: {
		Fword:  __ccgo_ts + 49586,
		Ftype1: int8('F'),
	},
	7155: {
		Fword:  __ccgo_ts + 49593,
		Ftype1: int8('F'),
	},
	7156: {
		Fword:  __ccgo_ts + 49600,
		Ftype1: int8('F'),
	},
	7157: {
		Fword:  __ccgo_ts + 49607,
		Ftype1: int8('F'),
	},
	7158: {
		Fword:  __ccgo_ts + 49614,
		Ftype1: int8('F'),
	},
	7159: {
		Fword:  __ccgo_ts + 49621,
		Ftype1: int8('F'),
	},
	7160: {
		Fword:  __ccgo_ts + 49628,
		Ftype1: int8('F'),
	},
	7161: {
		Fword:  __ccgo_ts + 49635,
		Ftype1: int8('F'),
	},
	7162: {
		Fword:  __ccgo_ts + 49641,
		Ftype1: int8('F'),
	},
	7163: {
		Fword:  __ccgo_ts + 49648,
		Ftype1: int8('F'),
	},
	7164: {
		Fword:  __ccgo_ts + 49655,
		Ftype1: int8('F'),
	},
	7165: {
		Fword:  __ccgo_ts + 49662,
		Ftype1: int8('F'),
	},
	7166: {
		Fword:  __ccgo_ts + 49669,
		Ftype1: int8('F'),
	},
	7167: {
		Fword:  __ccgo_ts + 49675,
		Ftype1: int8('F'),
	},
	7168: {
		Fword:  __ccgo_ts + 49682,
		Ftype1: int8('F'),
	},
	7169: {
		Fword:  __ccgo_ts + 49689,
		Ftype1: int8('F'),
	},
	7170: {
		Fword:  __ccgo_ts + 49696,
		Ftype1: int8('F'),
	},
	7171: {
		Fword:  __ccgo_ts + 49703,
		Ftype1: int8('F'),
	},
	7172: {
		Fword:  __ccgo_ts + 49710,
		Ftype1: int8('F'),
	},
	7173: {
		Fword:  __ccgo_ts + 49717,
		Ftype1: int8('F'),
	},
	7174: {
		Fword:  __ccgo_ts + 49724,
		Ftype1: int8('F'),
	},
	7175: {
		Fword:  __ccgo_ts + 49731,
		Ftype1: int8('F'),
	},
	7176: {
		Fword:  __ccgo_ts + 49738,
		Ftype1: int8('F'),
	},
	7177: {
		Fword:  __ccgo_ts + 49745,
		Ftype1: int8('F'),
	},
	7178: {
		Fword:  __ccgo_ts + 49752,
		Ftype1: int8('F'),
	},
	7179: {
		Fword:  __ccgo_ts + 49759,
		Ftype1: int8('F'),
	},
	7180: {
		Fword:  __ccgo_ts + 49766,
		Ftype1: int8('F'),
	},
	7181: {
		Fword:  __ccgo_ts + 49773,
		Ftype1: int8('F'),
	},
	7182: {
		Fword:  __ccgo_ts + 49780,
		Ftype1: int8('F'),
	},
	7183: {
		Fword:  __ccgo_ts + 49786,
		Ftype1: int8('F'),
	},
	7184: {
		Fword:  __ccgo_ts + 49793,
		Ftype1: int8('F'),
	},
	7185: {
		Fword:  __ccgo_ts + 49800,
		Ftype1: int8('F'),
	},
	7186: {
		Fword:  __ccgo_ts + 49807,
		Ftype1: int8('F'),
	},
	7187: {
		Fword:  __ccgo_ts + 49814,
		Ftype1: int8('F'),
	},
	7188: {
		Fword:  __ccgo_ts + 49821,
		Ftype1: int8('F'),
	},
	7189: {
		Fword:  __ccgo_ts + 49828,
		Ftype1: int8('F'),
	},
	7190: {
		Fword:  __ccgo_ts + 49835,
		Ftype1: int8('F'),
	},
	7191: {
		Fword:  __ccgo_ts + 49842,
		Ftype1: int8('F'),
	},
	7192: {
		Fword:  __ccgo_ts + 49849,
		Ftype1: int8('F'),
	},
	7193: {
		Fword:  __ccgo_ts + 49855,
		Ftype1: int8('F'),
	},
	7194: {
		Fword:  __ccgo_ts + 49862,
		Ftype1: int8('F'),
	},
	7195: {
		Fword:  __ccgo_ts + 49869,
		Ftype1: int8('F'),
	},
	7196: {
		Fword:  __ccgo_ts + 49876,
		Ftype1: int8('F'),
	},
	7197: {
		Fword:  __ccgo_ts + 49883,
		Ftype1: int8('F'),
	},
	7198: {
		Fword:  __ccgo_ts + 49890,
		Ftype1: int8('F'),
	},
	7199: {
		Fword:  __ccgo_ts + 49897,
		Ftype1: int8('F'),
	},
	7200: {
		Fword:  __ccgo_ts + 49904,
		Ftype1: int8('F'),
	},
	7201: {
		Fword:  __ccgo_ts + 49911,
		Ftype1: int8('F'),
	},
	7202: {
		Fword:  __ccgo_ts + 49917,
		Ftype1: int8('F'),
	},
	7203: {
		Fword:  __ccgo_ts + 49924,
		Ftype1: int8('F'),
	},
	7204: {
		Fword:  __ccgo_ts + 49931,
		Ftype1: int8('F'),
	},
	7205: {
		Fword:  __ccgo_ts + 49938,
		Ftype1: int8('F'),
	},
	7206: {
		Fword:  __ccgo_ts + 49945,
		Ftype1: int8('F'),
	},
	7207: {
		Fword:  __ccgo_ts + 49952,
		Ftype1: int8('F'),
	},
	7208: {
		Fword:  __ccgo_ts + 49959,
		Ftype1: int8('F'),
	},
	7209: {
		Fword:  __ccgo_ts + 49966,
		Ftype1: int8('F'),
	},
	7210: {
		Fword:  __ccgo_ts + 49973,
		Ftype1: int8('F'),
	},
	7211: {
		Fword:  __ccgo_ts + 49979,
		Ftype1: int8('F'),
	},
	7212: {
		Fword:  __ccgo_ts + 49986,
		Ftype1: int8('F'),
	},
	7213: {
		Fword:  __ccgo_ts + 49993,
		Ftype1: int8('F'),
	},
	7214: {
		Fword:  __ccgo_ts + 50000,
		Ftype1: int8('F'),
	},
	7215: {
		Fword:  __ccgo_ts + 50007,
		Ftype1: int8('F'),
	},
	7216: {
		Fword:  __ccgo_ts + 50014,
		Ftype1: int8('F'),
	},
	7217: {
		Fword:  __ccgo_ts + 50021,
		Ftype1: int8('F'),
	},
	7218: {
		Fword:  __ccgo_ts + 50028,
		Ftype1: int8('F'),
	},
	7219: {
		Fword:  __ccgo_ts + 50035,
		Ftype1: int8('F'),
	},
	7220: {
		Fword:  __ccgo_ts + 50042,
		Ftype1: int8('F'),
	},
	7221: {
		Fword:  __ccgo_ts + 50049,
		Ftype1: int8('F'),
	},
	7222: {
		Fword:  __ccgo_ts + 50056,
		Ftype1: int8('F'),
	},
	7223: {
		Fword:  __ccgo_ts + 50063,
		Ftype1: int8('F'),
	},
	7224: {
		Fword:  __ccgo_ts + 50070,
		Ftype1: int8('F'),
	},
	7225: {
		Fword:  __ccgo_ts + 50077,
		Ftype1: int8('F'),
	},
	7226: {
		Fword:  __ccgo_ts + 50084,
		Ftype1: int8('F'),
	},
	7227: {
		Fword:  __ccgo_ts + 50091,
		Ftype1: int8('F'),
	},
	7228: {
		Fword:  __ccgo_ts + 50098,
		Ftype1: int8('F'),
	},
	7229: {
		Fword:  __ccgo_ts + 50105,
		Ftype1: int8('F'),
	},
	7230: {
		Fword:  __ccgo_ts + 50112,
		Ftype1: int8('F'),
	},
	7231: {
		Fword:  __ccgo_ts + 50119,
		Ftype1: int8('F'),
	},
	7232: {
		Fword:  __ccgo_ts + 50126,
		Ftype1: int8('F'),
	},
	7233: {
		Fword:  __ccgo_ts + 50133,
		Ftype1: int8('F'),
	},
	7234: {
		Fword:  __ccgo_ts + 50140,
		Ftype1: int8('F'),
	},
	7235: {
		Fword:  __ccgo_ts + 50147,
		Ftype1: int8('F'),
	},
	7236: {
		Fword:  __ccgo_ts + 50154,
		Ftype1: int8('F'),
	},
	7237: {
		Fword:  __ccgo_ts + 50161,
		Ftype1: int8('F'),
	},
	7238: {
		Fword:  __ccgo_ts + 50168,
		Ftype1: int8('F'),
	},
	7239: {
		Fword:  __ccgo_ts + 50175,
		Ftype1: int8('F'),
	},
	7240: {
		Fword:  __ccgo_ts + 50182,
		Ftype1: int8('F'),
	},
	7241: {
		Fword:  __ccgo_ts + 50189,
		Ftype1: int8('F'),
	},
	7242: {
		Fword:  __ccgo_ts + 50196,
		Ftype1: int8('F'),
	},
	7243: {
		Fword:  __ccgo_ts + 50203,
		Ftype1: int8('F'),
	},
	7244: {
		Fword:  __ccgo_ts + 50210,
		Ftype1: int8('F'),
	},
	7245: {
		Fword:  __ccgo_ts + 50217,
		Ftype1: int8('F'),
	},
	7246: {
		Fword:  __ccgo_ts + 50223,
		Ftype1: int8('F'),
	},
	7247: {
		Fword:  __ccgo_ts + 50230,
		Ftype1: int8('F'),
	},
	7248: {
		Fword:  __ccgo_ts + 50237,
		Ftype1: int8('F'),
	},
	7249: {
		Fword:  __ccgo_ts + 50244,
		Ftype1: int8('F'),
	},
	7250: {
		Fword:  __ccgo_ts + 50251,
		Ftype1: int8('F'),
	},
	7251: {
		Fword:  __ccgo_ts + 50258,
		Ftype1: int8('F'),
	},
	7252: {
		Fword:  __ccgo_ts + 50265,
		Ftype1: int8('F'),
	},
	7253: {
		Fword:  __ccgo_ts + 50272,
		Ftype1: int8('F'),
	},
	7254: {
		Fword:  __ccgo_ts + 50279,
		Ftype1: int8('F'),
	},
	7255: {
		Fword:  __ccgo_ts + 50285,
		Ftype1: int8('F'),
	},
	7256: {
		Fword:  __ccgo_ts + 50292,
		Ftype1: int8('F'),
	},
	7257: {
		Fword:  __ccgo_ts + 50299,
		Ftype1: int8('F'),
	},
	7258: {
		Fword:  __ccgo_ts + 50306,
		Ftype1: int8('F'),
	},
	7259: {
		Fword:  __ccgo_ts + 50313,
		Ftype1: int8('F'),
	},
	7260: {
		Fword:  __ccgo_ts + 50320,
		Ftype1: int8('F'),
	},
	7261: {
		Fword:  __ccgo_ts + 50327,
		Ftype1: int8('F'),
	},
	7262: {
		Fword:  __ccgo_ts + 50333,
		Ftype1: int8('F'),
	},
	7263: {
		Fword:  __ccgo_ts + 50340,
		Ftype1: int8('F'),
	},
	7264: {
		Fword:  __ccgo_ts + 50347,
		Ftype1: int8('F'),
	},
	7265: {
		Fword:  __ccgo_ts + 50354,
		Ftype1: int8('F'),
	},
	7266: {
		Fword:  __ccgo_ts + 50361,
		Ftype1: int8('F'),
	},
	7267: {
		Fword:  __ccgo_ts + 50368,
		Ftype1: int8('F'),
	},
	7268: {
		Fword:  __ccgo_ts + 50375,
		Ftype1: int8('F'),
	},
	7269: {
		Fword:  __ccgo_ts + 50381,
		Ftype1: int8('F'),
	},
	7270: {
		Fword:  __ccgo_ts + 50388,
		Ftype1: int8('F'),
	},
	7271: {
		Fword:  __ccgo_ts + 50395,
		Ftype1: int8('F'),
	},
	7272: {
		Fword:  __ccgo_ts + 50402,
		Ftype1: int8('F'),
	},
	7273: {
		Fword:  __ccgo_ts + 50409,
		Ftype1: int8('F'),
	},
	7274: {
		Fword:  __ccgo_ts + 50416,
		Ftype1: int8('F'),
	},
	7275: {
		Fword:  __ccgo_ts + 50423,
		Ftype1: int8('F'),
	},
	7276: {
		Fword:  __ccgo_ts + 50428,
		Ftype1: int8('F'),
	},
	7277: {
		Fword:  __ccgo_ts + 50435,
		Ftype1: int8('F'),
	},
	7278: {
		Fword:  __ccgo_ts + 50442,
		Ftype1: int8('F'),
	},
	7279: {
		Fword:  __ccgo_ts + 50449,
		Ftype1: int8('F'),
	},
	7280: {
		Fword:  __ccgo_ts + 50456,
		Ftype1: int8('F'),
	},
	7281: {
		Fword:  __ccgo_ts + 50463,
		Ftype1: int8('F'),
	},
	7282: {
		Fword:  __ccgo_ts + 50470,
		Ftype1: int8('F'),
	},
	7283: {
		Fword:  __ccgo_ts + 50477,
		Ftype1: int8('F'),
	},
	7284: {
		Fword:  __ccgo_ts + 50484,
		Ftype1: int8('F'),
	},
	7285: {
		Fword:  __ccgo_ts + 50491,
		Ftype1: int8('F'),
	},
	7286: {
		Fword:  __ccgo_ts + 50498,
		Ftype1: int8('F'),
	},
	7287: {
		Fword:  __ccgo_ts + 50505,
		Ftype1: int8('F'),
	},
	7288: {
		Fword:  __ccgo_ts + 50512,
		Ftype1: int8('F'),
	},
	7289: {
		Fword:  __ccgo_ts + 50519,
		Ftype1: int8('F'),
	},
	7290: {
		Fword:  __ccgo_ts + 50526,
		Ftype1: int8('F'),
	},
	7291: {
		Fword:  __ccgo_ts + 50533,
		Ftype1: int8('F'),
	},
	7292: {
		Fword:  __ccgo_ts + 50540,
		Ftype1: int8('F'),
	},
	7293: {
		Fword:  __ccgo_ts + 50547,
		Ftype1: int8('F'),
	},
	7294: {
		Fword:  __ccgo_ts + 50554,
		Ftype1: int8('F'),
	},
	7295: {
		Fword:  __ccgo_ts + 50561,
		Ftype1: int8('F'),
	},
	7296: {
		Fword:  __ccgo_ts + 50568,
		Ftype1: int8('F'),
	},
	7297: {
		Fword:  __ccgo_ts + 50575,
		Ftype1: int8('F'),
	},
	7298: {
		Fword:  __ccgo_ts + 50582,
		Ftype1: int8('F'),
	},
	7299: {
		Fword:  __ccgo_ts + 50589,
		Ftype1: int8('F'),
	},
	7300: {
		Fword:  __ccgo_ts + 50596,
		Ftype1: int8('F'),
	},
	7301: {
		Fword:  __ccgo_ts + 50603,
		Ftype1: int8('F'),
	},
	7302: {
		Fword:  __ccgo_ts + 50610,
		Ftype1: int8('F'),
	},
	7303: {
		Fword:  __ccgo_ts + 50617,
		Ftype1: int8('F'),
	},
	7304: {
		Fword:  __ccgo_ts + 50624,
		Ftype1: int8('F'),
	},
	7305: {
		Fword:  __ccgo_ts + 50631,
		Ftype1: int8('F'),
	},
	7306: {
		Fword:  __ccgo_ts + 50638,
		Ftype1: int8('F'),
	},
	7307: {
		Fword:  __ccgo_ts + 50645,
		Ftype1: int8('F'),
	},
	7308: {
		Fword:  __ccgo_ts + 50652,
		Ftype1: int8('F'),
	},
	7309: {
		Fword:  __ccgo_ts + 50659,
		Ftype1: int8('F'),
	},
	7310: {
		Fword:  __ccgo_ts + 50666,
		Ftype1: int8('F'),
	},
	7311: {
		Fword:  __ccgo_ts + 50673,
		Ftype1: int8('F'),
	},
	7312: {
		Fword:  __ccgo_ts + 50680,
		Ftype1: int8('F'),
	},
	7313: {
		Fword:  __ccgo_ts + 50687,
		Ftype1: int8('F'),
	},
	7314: {
		Fword:  __ccgo_ts + 50694,
		Ftype1: int8('F'),
	},
	7315: {
		Fword:  __ccgo_ts + 50701,
		Ftype1: int8('F'),
	},
	7316: {
		Fword:  __ccgo_ts + 50708,
		Ftype1: int8('F'),
	},
	7317: {
		Fword:  __ccgo_ts + 50715,
		Ftype1: int8('F'),
	},
	7318: {
		Fword:  __ccgo_ts + 50722,
		Ftype1: int8('F'),
	},
	7319: {
		Fword:  __ccgo_ts + 50729,
		Ftype1: int8('F'),
	},
	7320: {
		Fword:  __ccgo_ts + 50736,
		Ftype1: int8('F'),
	},
	7321: {
		Fword:  __ccgo_ts + 50743,
		Ftype1: int8('F'),
	},
	7322: {
		Fword:  __ccgo_ts + 50750,
		Ftype1: int8('F'),
	},
	7323: {
		Fword:  __ccgo_ts + 50757,
		Ftype1: int8('F'),
	},
	7324: {
		Fword:  __ccgo_ts + 50764,
		Ftype1: int8('F'),
	},
	7325: {
		Fword:  __ccgo_ts + 50771,
		Ftype1: int8('F'),
	},
	7326: {
		Fword:  __ccgo_ts + 50778,
		Ftype1: int8('F'),
	},
	7327: {
		Fword:  __ccgo_ts + 50785,
		Ftype1: int8('F'),
	},
	7328: {
		Fword:  __ccgo_ts + 50792,
		Ftype1: int8('F'),
	},
	7329: {
		Fword:  __ccgo_ts + 50799,
		Ftype1: int8('F'),
	},
	7330: {
		Fword:  __ccgo_ts + 50806,
		Ftype1: int8('F'),
	},
	7331: {
		Fword:  __ccgo_ts + 50813,
		Ftype1: int8('F'),
	},
	7332: {
		Fword:  __ccgo_ts + 50820,
		Ftype1: int8('F'),
	},
	7333: {
		Fword:  __ccgo_ts + 50827,
		Ftype1: int8('F'),
	},
	7334: {
		Fword:  __ccgo_ts + 50834,
		Ftype1: int8('F'),
	},
	7335: {
		Fword:  __ccgo_ts + 50841,
		Ftype1: int8('F'),
	},
	7336: {
		Fword:  __ccgo_ts + 50848,
		Ftype1: int8('F'),
	},
	7337: {
		Fword:  __ccgo_ts + 50855,
		Ftype1: int8('F'),
	},
	7338: {
		Fword:  __ccgo_ts + 50862,
		Ftype1: int8('F'),
	},
	7339: {
		Fword:  __ccgo_ts + 50869,
		Ftype1: int8('F'),
	},
	7340: {
		Fword:  __ccgo_ts + 50876,
		Ftype1: int8('F'),
	},
	7341: {
		Fword:  __ccgo_ts + 50883,
		Ftype1: int8('F'),
	},
	7342: {
		Fword:  __ccgo_ts + 50889,
		Ftype1: int8('F'),
	},
	7343: {
		Fword:  __ccgo_ts + 50896,
		Ftype1: int8('F'),
	},
	7344: {
		Fword:  __ccgo_ts + 50903,
		Ftype1: int8('F'),
	},
	7345: {
		Fword:  __ccgo_ts + 50910,
		Ftype1: int8('F'),
	},
	7346: {
		Fword:  __ccgo_ts + 50917,
		Ftype1: int8('F'),
	},
	7347: {
		Fword:  __ccgo_ts + 50924,
		Ftype1: int8('F'),
	},
	7348: {
		Fword:  __ccgo_ts + 50931,
		Ftype1: int8('F'),
	},
	7349: {
		Fword:  __ccgo_ts + 50938,
		Ftype1: int8('F'),
	},
	7350: {
		Fword:  __ccgo_ts + 50945,
		Ftype1: int8('F'),
	},
	7351: {
		Fword:  __ccgo_ts + 50951,
		Ftype1: int8('F'),
	},
	7352: {
		Fword:  __ccgo_ts + 50958,
		Ftype1: int8('F'),
	},
	7353: {
		Fword:  __ccgo_ts + 50965,
		Ftype1: int8('F'),
	},
	7354: {
		Fword:  __ccgo_ts + 50972,
		Ftype1: int8('F'),
	},
	7355: {
		Fword:  __ccgo_ts + 50979,
		Ftype1: int8('F'),
	},
	7356: {
		Fword:  __ccgo_ts + 50986,
		Ftype1: int8('F'),
	},
	7357: {
		Fword:  __ccgo_ts + 50993,
		Ftype1: int8('F'),
	},
	7358: {
		Fword:  __ccgo_ts + 51000,
		Ftype1: int8('F'),
	},
	7359: {
		Fword:  __ccgo_ts + 51006,
		Ftype1: int8('F'),
	},
	7360: {
		Fword:  __ccgo_ts + 51013,
		Ftype1: int8('F'),
	},
	7361: {
		Fword:  __ccgo_ts + 51020,
		Ftype1: int8('F'),
	},
	7362: {
		Fword:  __ccgo_ts + 51027,
		Ftype1: int8('F'),
	},
	7363: {
		Fword:  __ccgo_ts + 51034,
		Ftype1: int8('F'),
	},
	7364: {
		Fword:  __ccgo_ts + 51041,
		Ftype1: int8('F'),
	},
	7365: {
		Fword:  __ccgo_ts + 51048,
		Ftype1: int8('F'),
	},
	7366: {
		Fword:  __ccgo_ts + 51055,
		Ftype1: int8('F'),
	},
	7367: {
		Fword:  __ccgo_ts + 51061,
		Ftype1: int8('F'),
	},
	7368: {
		Fword:  __ccgo_ts + 51068,
		Ftype1: int8('F'),
	},
	7369: {
		Fword:  __ccgo_ts + 51075,
		Ftype1: int8('F'),
	},
	7370: {
		Fword:  __ccgo_ts + 51082,
		Ftype1: int8('F'),
	},
	7371: {
		Fword:  __ccgo_ts + 51089,
		Ftype1: int8('F'),
	},
	7372: {
		Fword:  __ccgo_ts + 51096,
		Ftype1: int8('F'),
	},
	7373: {
		Fword:  __ccgo_ts + 51103,
		Ftype1: int8('F'),
	},
	7374: {
		Fword:  __ccgo_ts + 51110,
		Ftype1: int8('F'),
	},
	7375: {
		Fword:  __ccgo_ts + 51117,
		Ftype1: int8('F'),
	},
	7376: {
		Fword:  __ccgo_ts + 51124,
		Ftype1: int8('F'),
	},
	7377: {
		Fword:  __ccgo_ts + 51131,
		Ftype1: int8('F'),
	},
	7378: {
		Fword:  __ccgo_ts + 51138,
		Ftype1: int8('F'),
	},
	7379: {
		Fword:  __ccgo_ts + 51145,
		Ftype1: int8('F'),
	},
	7380: {
		Fword:  __ccgo_ts + 51152,
		Ftype1: int8('F'),
	},
	7381: {
		Fword:  __ccgo_ts + 51159,
		Ftype1: int8('F'),
	},
	7382: {
		Fword:  __ccgo_ts + 51166,
		Ftype1: int8('F'),
	},
	7383: {
		Fword:  __ccgo_ts + 51173,
		Ftype1: int8('F'),
	},
	7384: {
		Fword:  __ccgo_ts + 51180,
		Ftype1: int8('F'),
	},
	7385: {
		Fword:  __ccgo_ts + 51187,
		Ftype1: int8('F'),
	},
	7386: {
		Fword:  __ccgo_ts + 51194,
		Ftype1: int8('F'),
	},
	7387: {
		Fword:  __ccgo_ts + 51201,
		Ftype1: int8('F'),
	},
	7388: {
		Fword:  __ccgo_ts + 51208,
		Ftype1: int8('F'),
	},
	7389: {
		Fword:  __ccgo_ts + 51215,
		Ftype1: int8('F'),
	},
	7390: {
		Fword:  __ccgo_ts + 51222,
		Ftype1: int8('F'),
	},
	7391: {
		Fword:  __ccgo_ts + 51229,
		Ftype1: int8('F'),
	},
	7392: {
		Fword:  __ccgo_ts + 51236,
		Ftype1: int8('F'),
	},
	7393: {
		Fword:  __ccgo_ts + 51243,
		Ftype1: int8('F'),
	},
	7394: {
		Fword:  __ccgo_ts + 51250,
		Ftype1: int8('F'),
	},
	7395: {
		Fword:  __ccgo_ts + 51257,
		Ftype1: int8('F'),
	},
	7396: {
		Fword:  __ccgo_ts + 51264,
		Ftype1: int8('F'),
	},
	7397: {
		Fword:  __ccgo_ts + 51271,
		Ftype1: int8('F'),
	},
	7398: {
		Fword:  __ccgo_ts + 51278,
		Ftype1: int8('F'),
	},
	7399: {
		Fword:  __ccgo_ts + 51285,
		Ftype1: int8('F'),
	},
	7400: {
		Fword:  __ccgo_ts + 51292,
		Ftype1: int8('F'),
	},
	7401: {
		Fword:  __ccgo_ts + 51299,
		Ftype1: int8('F'),
	},
	7402: {
		Fword:  __ccgo_ts + 51306,
		Ftype1: int8('F'),
	},
	7403: {
		Fword:  __ccgo_ts + 51313,
		Ftype1: int8('F'),
	},
	7404: {
		Fword:  __ccgo_ts + 51320,
		Ftype1: int8('F'),
	},
	7405: {
		Fword:  __ccgo_ts + 51327,
		Ftype1: int8('F'),
	},
	7406: {
		Fword:  __ccgo_ts + 51334,
		Ftype1: int8('F'),
	},
	7407: {
		Fword:  __ccgo_ts + 51341,
		Ftype1: int8('F'),
	},
	7408: {
		Fword:  __ccgo_ts + 51348,
		Ftype1: int8('F'),
	},
	7409: {
		Fword:  __ccgo_ts + 51355,
		Ftype1: int8('F'),
	},
	7410: {
		Fword:  __ccgo_ts + 51362,
		Ftype1: int8('F'),
	},
	7411: {
		Fword:  __ccgo_ts + 51369,
		Ftype1: int8('F'),
	},
	7412: {
		Fword:  __ccgo_ts + 51376,
		Ftype1: int8('F'),
	},
	7413: {
		Fword:  __ccgo_ts + 51383,
		Ftype1: int8('F'),
	},
	7414: {
		Fword:  __ccgo_ts + 51390,
		Ftype1: int8('F'),
	},
	7415: {
		Fword:  __ccgo_ts + 51397,
		Ftype1: int8('F'),
	},
	7416: {
		Fword:  __ccgo_ts + 51404,
		Ftype1: int8('F'),
	},
	7417: {
		Fword:  __ccgo_ts + 51411,
		Ftype1: int8('F'),
	},
	7418: {
		Fword:  __ccgo_ts + 51418,
		Ftype1: int8('F'),
	},
	7419: {
		Fword:  __ccgo_ts + 51425,
		Ftype1: int8('F'),
	},
	7420: {
		Fword:  __ccgo_ts + 51432,
		Ftype1: int8('F'),
	},
	7421: {
		Fword:  __ccgo_ts + 51439,
		Ftype1: int8('F'),
	},
	7422: {
		Fword:  __ccgo_ts + 51446,
		Ftype1: int8('F'),
	},
	7423: {
		Fword:  __ccgo_ts + 51453,
		Ftype1: int8('F'),
	},
	7424: {
		Fword:  __ccgo_ts + 51460,
		Ftype1: int8('F'),
	},
	7425: {
		Fword:  __ccgo_ts + 51467,
		Ftype1: int8('F'),
	},
	7426: {
		Fword:  __ccgo_ts + 51474,
		Ftype1: int8('F'),
	},
	7427: {
		Fword:  __ccgo_ts + 51481,
		Ftype1: int8('F'),
	},
	7428: {
		Fword:  __ccgo_ts + 51488,
		Ftype1: int8('F'),
	},
	7429: {
		Fword:  __ccgo_ts + 51495,
		Ftype1: int8('F'),
	},
	7430: {
		Fword:  __ccgo_ts + 51502,
		Ftype1: int8('F'),
	},
	7431: {
		Fword:  __ccgo_ts + 51509,
		Ftype1: int8('F'),
	},
	7432: {
		Fword:  __ccgo_ts + 51516,
		Ftype1: int8('F'),
	},
	7433: {
		Fword:  __ccgo_ts + 51523,
		Ftype1: int8('F'),
	},
	7434: {
		Fword:  __ccgo_ts + 51530,
		Ftype1: int8('F'),
	},
	7435: {
		Fword:  __ccgo_ts + 51537,
		Ftype1: int8('F'),
	},
	7436: {
		Fword:  __ccgo_ts + 51544,
		Ftype1: int8('F'),
	},
	7437: {
		Fword:  __ccgo_ts + 51551,
		Ftype1: int8('F'),
	},
	7438: {
		Fword:  __ccgo_ts + 51558,
		Ftype1: int8('F'),
	},
	7439: {
		Fword:  __ccgo_ts + 51565,
		Ftype1: int8('F'),
	},
	7440: {
		Fword:  __ccgo_ts + 51572,
		Ftype1: int8('F'),
	},
	7441: {
		Fword:  __ccgo_ts + 51579,
		Ftype1: int8('F'),
	},
	7442: {
		Fword:  __ccgo_ts + 51586,
		Ftype1: int8('F'),
	},
	7443: {
		Fword:  __ccgo_ts + 51593,
		Ftype1: int8('F'),
	},
	7444: {
		Fword:  __ccgo_ts + 51600,
		Ftype1: int8('F'),
	},
	7445: {
		Fword:  __ccgo_ts + 51607,
		Ftype1: int8('F'),
	},
	7446: {
		Fword:  __ccgo_ts + 51614,
		Ftype1: int8('F'),
	},
	7447: {
		Fword:  __ccgo_ts + 51621,
		Ftype1: int8('F'),
	},
	7448: {
		Fword:  __ccgo_ts + 51628,
		Ftype1: int8('F'),
	},
	7449: {
		Fword:  __ccgo_ts + 51635,
		Ftype1: int8('F'),
	},
	7450: {
		Fword:  __ccgo_ts + 51642,
		Ftype1: int8('F'),
	},
	7451: {
		Fword:  __ccgo_ts + 51649,
		Ftype1: int8('F'),
	},
	7452: {
		Fword:  __ccgo_ts + 51656,
		Ftype1: int8('F'),
	},
	7453: {
		Fword:  __ccgo_ts + 51663,
		Ftype1: int8('F'),
	},
	7454: {
		Fword:  __ccgo_ts + 51670,
		Ftype1: int8('F'),
	},
	7455: {
		Fword:  __ccgo_ts + 51677,
		Ftype1: int8('F'),
	},
	7456: {
		Fword:  __ccgo_ts + 51684,
		Ftype1: int8('F'),
	},
	7457: {
		Fword:  __ccgo_ts + 51691,
		Ftype1: int8('F'),
	},
	7458: {
		Fword:  __ccgo_ts + 51698,
		Ftype1: int8('F'),
	},
	7459: {
		Fword:  __ccgo_ts + 51705,
		Ftype1: int8('F'),
	},
	7460: {
		Fword:  __ccgo_ts + 51712,
		Ftype1: int8('F'),
	},
	7461: {
		Fword:  __ccgo_ts + 51718,
		Ftype1: int8('F'),
	},
	7462: {
		Fword:  __ccgo_ts + 51725,
		Ftype1: int8('F'),
	},
	7463: {
		Fword:  __ccgo_ts + 51732,
		Ftype1: int8('F'),
	},
	7464: {
		Fword:  __ccgo_ts + 51739,
		Ftype1: int8('F'),
	},
	7465: {
		Fword:  __ccgo_ts + 51746,
		Ftype1: int8('F'),
	},
	7466: {
		Fword:  __ccgo_ts + 51753,
		Ftype1: int8('F'),
	},
	7467: {
		Fword:  __ccgo_ts + 51760,
		Ftype1: int8('F'),
	},
	7468: {
		Fword:  __ccgo_ts + 51767,
		Ftype1: int8('F'),
	},
	7469: {
		Fword:  __ccgo_ts + 51774,
		Ftype1: int8('F'),
	},
	7470: {
		Fword:  __ccgo_ts + 51781,
		Ftype1: int8('F'),
	},
	7471: {
		Fword:  __ccgo_ts + 51788,
		Ftype1: int8('F'),
	},
	7472: {
		Fword:  __ccgo_ts + 51795,
		Ftype1: int8('F'),
	},
	7473: {
		Fword:  __ccgo_ts + 51802,
		Ftype1: int8('F'),
	},
	7474: {
		Fword:  __ccgo_ts + 51809,
		Ftype1: int8('F'),
	},
	7475: {
		Fword:  __ccgo_ts + 51816,
		Ftype1: int8('F'),
	},
	7476: {
		Fword:  __ccgo_ts + 51823,
		Ftype1: int8('F'),
	},
	7477: {
		Fword:  __ccgo_ts + 51830,
		Ftype1: int8('F'),
	},
	7478: {
		Fword:  __ccgo_ts + 51837,
		Ftype1: int8('F'),
	},
	7479: {
		Fword:  __ccgo_ts + 51844,
		Ftype1: int8('F'),
	},
	7480: {
		Fword:  __ccgo_ts + 51851,
		Ftype1: int8('F'),
	},
	7481: {
		Fword:  __ccgo_ts + 51858,
		Ftype1: int8('F'),
	},
	7482: {
		Fword:  __ccgo_ts + 51865,
		Ftype1: int8('F'),
	},
	7483: {
		Fword:  __ccgo_ts + 51872,
		Ftype1: int8('F'),
	},
	7484: {
		Fword:  __ccgo_ts + 51879,
		Ftype1: int8('F'),
	},
	7485: {
		Fword:  __ccgo_ts + 51886,
		Ftype1: int8('F'),
	},
	7486: {
		Fword:  __ccgo_ts + 51893,
		Ftype1: int8('F'),
	},
	7487: {
		Fword:  __ccgo_ts + 51900,
		Ftype1: int8('F'),
	},
	7488: {
		Fword:  __ccgo_ts + 51907,
		Ftype1: int8('F'),
	},
	7489: {
		Fword:  __ccgo_ts + 51914,
		Ftype1: int8('F'),
	},
	7490: {
		Fword:  __ccgo_ts + 51921,
		Ftype1: int8('F'),
	},
	7491: {
		Fword:  __ccgo_ts + 51928,
		Ftype1: int8('F'),
	},
	7492: {
		Fword:  __ccgo_ts + 51935,
		Ftype1: int8('F'),
	},
	7493: {
		Fword:  __ccgo_ts + 51942,
		Ftype1: int8('F'),
	},
	7494: {
		Fword:  __ccgo_ts + 51949,
		Ftype1: int8('F'),
	},
	7495: {
		Fword:  __ccgo_ts + 51956,
		Ftype1: int8('F'),
	},
	7496: {
		Fword:  __ccgo_ts + 51963,
		Ftype1: int8('F'),
	},
	7497: {
		Fword:  __ccgo_ts + 51970,
		Ftype1: int8('F'),
	},
	7498: {
		Fword:  __ccgo_ts + 51977,
		Ftype1: int8('F'),
	},
	7499: {
		Fword:  __ccgo_ts + 51984,
		Ftype1: int8('F'),
	},
	7500: {
		Fword:  __ccgo_ts + 51991,
		Ftype1: int8('F'),
	},
	7501: {
		Fword:  __ccgo_ts + 51998,
		Ftype1: int8('F'),
	},
	7502: {
		Fword:  __ccgo_ts + 52005,
		Ftype1: int8('F'),
	},
	7503: {
		Fword:  __ccgo_ts + 52012,
		Ftype1: int8('F'),
	},
	7504: {
		Fword:  __ccgo_ts + 52019,
		Ftype1: int8('F'),
	},
	7505: {
		Fword:  __ccgo_ts + 52026,
		Ftype1: int8('F'),
	},
	7506: {
		Fword:  __ccgo_ts + 52033,
		Ftype1: int8('F'),
	},
	7507: {
		Fword:  __ccgo_ts + 52040,
		Ftype1: int8('F'),
	},
	7508: {
		Fword:  __ccgo_ts + 52047,
		Ftype1: int8('F'),
	},
	7509: {
		Fword:  __ccgo_ts + 52054,
		Ftype1: int8('F'),
	},
	7510: {
		Fword:  __ccgo_ts + 52061,
		Ftype1: int8('F'),
	},
	7511: {
		Fword:  __ccgo_ts + 52068,
		Ftype1: int8('F'),
	},
	7512: {
		Fword:  __ccgo_ts + 52075,
		Ftype1: int8('F'),
	},
	7513: {
		Fword:  __ccgo_ts + 52082,
		Ftype1: int8('F'),
	},
	7514: {
		Fword:  __ccgo_ts + 52089,
		Ftype1: int8('F'),
	},
	7515: {
		Fword:  __ccgo_ts + 52096,
		Ftype1: int8('F'),
	},
	7516: {
		Fword:  __ccgo_ts + 52103,
		Ftype1: int8('F'),
	},
	7517: {
		Fword:  __ccgo_ts + 52110,
		Ftype1: int8('F'),
	},
	7518: {
		Fword:  __ccgo_ts + 52117,
		Ftype1: int8('F'),
	},
	7519: {
		Fword:  __ccgo_ts + 52124,
		Ftype1: int8('F'),
	},
	7520: {
		Fword:  __ccgo_ts + 52131,
		Ftype1: int8('F'),
	},
	7521: {
		Fword:  __ccgo_ts + 52138,
		Ftype1: int8('F'),
	},
	7522: {
		Fword:  __ccgo_ts + 52145,
		Ftype1: int8('F'),
	},
	7523: {
		Fword:  __ccgo_ts + 52152,
		Ftype1: int8('F'),
	},
	7524: {
		Fword:  __ccgo_ts + 52159,
		Ftype1: int8('F'),
	},
	7525: {
		Fword:  __ccgo_ts + 52166,
		Ftype1: int8('F'),
	},
	7526: {
		Fword:  __ccgo_ts + 52173,
		Ftype1: int8('F'),
	},
	7527: {
		Fword:  __ccgo_ts + 52180,
		Ftype1: int8('F'),
	},
	7528: {
		Fword:  __ccgo_ts + 52187,
		Ftype1: int8('F'),
	},
	7529: {
		Fword:  __ccgo_ts + 52194,
		Ftype1: int8('F'),
	},
	7530: {
		Fword:  __ccgo_ts + 52201,
		Ftype1: int8('F'),
	},
	7531: {
		Fword:  __ccgo_ts + 52208,
		Ftype1: int8('F'),
	},
	7532: {
		Fword:  __ccgo_ts + 52215,
		Ftype1: int8('F'),
	},
	7533: {
		Fword:  __ccgo_ts + 52222,
		Ftype1: int8('F'),
	},
	7534: {
		Fword:  __ccgo_ts + 52229,
		Ftype1: int8('F'),
	},
	7535: {
		Fword:  __ccgo_ts + 52236,
		Ftype1: int8('F'),
	},
	7536: {
		Fword:  __ccgo_ts + 52243,
		Ftype1: int8('F'),
	},
	7537: {
		Fword:  __ccgo_ts + 52250,
		Ftype1: int8('F'),
	},
	7538: {
		Fword:  __ccgo_ts + 52257,
		Ftype1: int8('F'),
	},
	7539: {
		Fword:  __ccgo_ts + 52264,
		Ftype1: int8('F'),
	},
	7540: {
		Fword:  __ccgo_ts + 52271,
		Ftype1: int8('F'),
	},
	7541: {
		Fword:  __ccgo_ts + 52278,
		Ftype1: int8('F'),
	},
	7542: {
		Fword:  __ccgo_ts + 52283,
		Ftype1: int8('F'),
	},
	7543: {
		Fword:  __ccgo_ts + 52290,
		Ftype1: int8('F'),
	},
	7544: {
		Fword:  __ccgo_ts + 52297,
		Ftype1: int8('F'),
	},
	7545: {
		Fword:  __ccgo_ts + 52304,
		Ftype1: int8('F'),
	},
	7546: {
		Fword:  __ccgo_ts + 52311,
		Ftype1: int8('F'),
	},
	7547: {
		Fword:  __ccgo_ts + 52318,
		Ftype1: int8('F'),
	},
	7548: {
		Fword:  __ccgo_ts + 52325,
		Ftype1: int8('F'),
	},
	7549: {
		Fword:  __ccgo_ts + 52332,
		Ftype1: int8('F'),
	},
	7550: {
		Fword:  __ccgo_ts + 52339,
		Ftype1: int8('F'),
	},
	7551: {
		Fword:  __ccgo_ts + 52345,
		Ftype1: int8('F'),
	},
	7552: {
		Fword:  __ccgo_ts + 52352,
		Ftype1: int8('F'),
	},
	7553: {
		Fword:  __ccgo_ts + 52359,
		Ftype1: int8('F'),
	},
	7554: {
		Fword:  __ccgo_ts + 52366,
		Ftype1: int8('F'),
	},
	7555: {
		Fword:  __ccgo_ts + 52373,
		Ftype1: int8('F'),
	},
	7556: {
		Fword:  __ccgo_ts + 52380,
		Ftype1: int8('F'),
	},
	7557: {
		Fword:  __ccgo_ts + 52387,
		Ftype1: int8('F'),
	},
	7558: {
		Fword:  __ccgo_ts + 52394,
		Ftype1: int8('F'),
	},
	7559: {
		Fword:  __ccgo_ts + 52400,
		Ftype1: int8('F'),
	},
	7560: {
		Fword:  __ccgo_ts + 52407,
		Ftype1: int8('F'),
	},
	7561: {
		Fword:  __ccgo_ts + 52414,
		Ftype1: int8('F'),
	},
	7562: {
		Fword:  __ccgo_ts + 52421,
		Ftype1: int8('F'),
	},
	7563: {
		Fword:  __ccgo_ts + 52428,
		Ftype1: int8('F'),
	},
	7564: {
		Fword:  __ccgo_ts + 52435,
		Ftype1: int8('F'),
	},
	7565: {
		Fword:  __ccgo_ts + 52442,
		Ftype1: int8('F'),
	},
	7566: {
		Fword:  __ccgo_ts + 52449,
		Ftype1: int8('F'),
	},
	7567: {
		Fword:  __ccgo_ts + 52456,
		Ftype1: int8('F'),
	},
	7568: {
		Fword:  __ccgo_ts + 52463,
		Ftype1: int8('F'),
	},
	7569: {
		Fword:  __ccgo_ts + 52470,
		Ftype1: int8('F'),
	},
	7570: {
		Fword:  __ccgo_ts + 52477,
		Ftype1: int8('F'),
	},
	7571: {
		Fword:  __ccgo_ts + 52484,
		Ftype1: int8('F'),
	},
	7572: {
		Fword:  __ccgo_ts + 52491,
		Ftype1: int8('F'),
	},
	7573: {
		Fword:  __ccgo_ts + 52498,
		Ftype1: int8('F'),
	},
	7574: {
		Fword:  __ccgo_ts + 52505,
		Ftype1: int8('F'),
	},
	7575: {
		Fword:  __ccgo_ts + 52512,
		Ftype1: int8('F'),
	},
	7576: {
		Fword:  __ccgo_ts + 52519,
		Ftype1: int8('F'),
	},
	7577: {
		Fword:  __ccgo_ts + 52526,
		Ftype1: int8('F'),
	},
	7578: {
		Fword:  __ccgo_ts + 52533,
		Ftype1: int8('F'),
	},
	7579: {
		Fword:  __ccgo_ts + 52540,
		Ftype1: int8('F'),
	},
	7580: {
		Fword:  __ccgo_ts + 52547,
		Ftype1: int8('F'),
	},
	7581: {
		Fword:  __ccgo_ts + 52554,
		Ftype1: int8('F'),
	},
	7582: {
		Fword:  __ccgo_ts + 52561,
		Ftype1: int8('F'),
	},
	7583: {
		Fword:  __ccgo_ts + 52568,
		Ftype1: int8('F'),
	},
	7584: {
		Fword:  __ccgo_ts + 52573,
		Ftype1: int8('F'),
	},
	7585: {
		Fword:  __ccgo_ts + 52580,
		Ftype1: int8('F'),
	},
	7586: {
		Fword:  __ccgo_ts + 52587,
		Ftype1: int8('F'),
	},
	7587: {
		Fword:  __ccgo_ts + 52594,
		Ftype1: int8('F'),
	},
	7588: {
		Fword:  __ccgo_ts + 52601,
		Ftype1: int8('F'),
	},
	7589: {
		Fword:  __ccgo_ts + 52608,
		Ftype1: int8('F'),
	},
	7590: {
		Fword:  __ccgo_ts + 52615,
		Ftype1: int8('F'),
	},
	7591: {
		Fword:  __ccgo_ts + 52622,
		Ftype1: int8('F'),
	},
	7592: {
		Fword:  __ccgo_ts + 52629,
		Ftype1: int8('F'),
	},
	7593: {
		Fword:  __ccgo_ts + 52635,
		Ftype1: int8('F'),
	},
	7594: {
		Fword:  __ccgo_ts + 52642,
		Ftype1: int8('F'),
	},
	7595: {
		Fword:  __ccgo_ts + 52649,
		Ftype1: int8('F'),
	},
	7596: {
		Fword:  __ccgo_ts + 52656,
		Ftype1: int8('F'),
	},
	7597: {
		Fword:  __ccgo_ts + 52663,
		Ftype1: int8('F'),
	},
	7598: {
		Fword:  __ccgo_ts + 52670,
		Ftype1: int8('F'),
	},
	7599: {
		Fword:  __ccgo_ts + 52677,
		Ftype1: int8('F'),
	},
	7600: {
		Fword:  __ccgo_ts + 52684,
		Ftype1: int8('F'),
	},
	7601: {
		Fword:  __ccgo_ts + 52690,
		Ftype1: int8('F'),
	},
	7602: {
		Fword:  __ccgo_ts + 52697,
		Ftype1: int8('F'),
	},
	7603: {
		Fword:  __ccgo_ts + 52704,
		Ftype1: int8('F'),
	},
	7604: {
		Fword:  __ccgo_ts + 52711,
		Ftype1: int8('F'),
	},
	7605: {
		Fword:  __ccgo_ts + 52718,
		Ftype1: int8('F'),
	},
	7606: {
		Fword:  __ccgo_ts + 52725,
		Ftype1: int8('F'),
	},
	7607: {
		Fword:  __ccgo_ts + 52732,
		Ftype1: int8('F'),
	},
	7608: {
		Fword:  __ccgo_ts + 52739,
		Ftype1: int8('F'),
	},
	7609: {
		Fword:  __ccgo_ts + 52746,
		Ftype1: int8('F'),
	},
	7610: {
		Fword:  __ccgo_ts + 52753,
		Ftype1: int8('F'),
	},
	7611: {
		Fword:  __ccgo_ts + 52760,
		Ftype1: int8('F'),
	},
	7612: {
		Fword:  __ccgo_ts + 52767,
		Ftype1: int8('F'),
	},
	7613: {
		Fword:  __ccgo_ts + 52774,
		Ftype1: int8('F'),
	},
	7614: {
		Fword:  __ccgo_ts + 52779,
		Ftype1: int8('F'),
	},
	7615: {
		Fword:  __ccgo_ts + 52786,
		Ftype1: int8('F'),
	},
	7616: {
		Fword:  __ccgo_ts + 52793,
		Ftype1: int8('F'),
	},
	7617: {
		Fword:  __ccgo_ts + 52800,
		Ftype1: int8('F'),
	},
	7618: {
		Fword:  __ccgo_ts + 52807,
		Ftype1: int8('F'),
	},
	7619: {
		Fword:  __ccgo_ts + 52814,
		Ftype1: int8('F'),
	},
	7620: {
		Fword:  __ccgo_ts + 52821,
		Ftype1: int8('F'),
	},
	7621: {
		Fword:  __ccgo_ts + 52828,
		Ftype1: int8('F'),
	},
	7622: {
		Fword:  __ccgo_ts + 52835,
		Ftype1: int8('F'),
	},
	7623: {
		Fword:  __ccgo_ts + 52841,
		Ftype1: int8('F'),
	},
	7624: {
		Fword:  __ccgo_ts + 52848,
		Ftype1: int8('F'),
	},
	7625: {
		Fword:  __ccgo_ts + 52855,
		Ftype1: int8('F'),
	},
	7626: {
		Fword:  __ccgo_ts + 52862,
		Ftype1: int8('F'),
	},
	7627: {
		Fword:  __ccgo_ts + 52869,
		Ftype1: int8('F'),
	},
	7628: {
		Fword:  __ccgo_ts + 52876,
		Ftype1: int8('F'),
	},
	7629: {
		Fword:  __ccgo_ts + 52883,
		Ftype1: int8('F'),
	},
	7630: {
		Fword:  __ccgo_ts + 52890,
		Ftype1: int8('F'),
	},
	7631: {
		Fword:  __ccgo_ts + 52896,
		Ftype1: int8('F'),
	},
	7632: {
		Fword:  __ccgo_ts + 52903,
		Ftype1: int8('F'),
	},
	7633: {
		Fword:  __ccgo_ts + 52910,
		Ftype1: int8('F'),
	},
	7634: {
		Fword:  __ccgo_ts + 52917,
		Ftype1: int8('F'),
	},
	7635: {
		Fword:  __ccgo_ts + 52924,
		Ftype1: int8('F'),
	},
	7636: {
		Fword:  __ccgo_ts + 52931,
		Ftype1: int8('F'),
	},
	7637: {
		Fword:  __ccgo_ts + 52938,
		Ftype1: int8('F'),
	},
	7638: {
		Fword:  __ccgo_ts + 52945,
		Ftype1: int8('F'),
	},
	7639: {
		Fword:  __ccgo_ts + 52952,
		Ftype1: int8('F'),
	},
	7640: {
		Fword:  __ccgo_ts + 52959,
		Ftype1: int8('F'),
	},
	7641: {
		Fword:  __ccgo_ts + 52966,
		Ftype1: int8('F'),
	},
	7642: {
		Fword:  __ccgo_ts + 52973,
		Ftype1: int8('F'),
	},
	7643: {
		Fword:  __ccgo_ts + 52980,
		Ftype1: int8('F'),
	},
	7644: {
		Fword:  __ccgo_ts + 52987,
		Ftype1: int8('F'),
	},
	7645: {
		Fword:  __ccgo_ts + 52994,
		Ftype1: int8('F'),
	},
	7646: {
		Fword:  __ccgo_ts + 52999,
		Ftype1: int8('F'),
	},
	7647: {
		Fword:  __ccgo_ts + 53006,
		Ftype1: int8('F'),
	},
	7648: {
		Fword:  __ccgo_ts + 53013,
		Ftype1: int8('F'),
	},
	7649: {
		Fword:  __ccgo_ts + 53020,
		Ftype1: int8('F'),
	},
	7650: {
		Fword:  __ccgo_ts + 53027,
		Ftype1: int8('F'),
	},
	7651: {
		Fword:  __ccgo_ts + 53034,
		Ftype1: int8('F'),
	},
	7652: {
		Fword:  __ccgo_ts + 53041,
		Ftype1: int8('F'),
	},
	7653: {
		Fword:  __ccgo_ts + 53048,
		Ftype1: int8('F'),
	},
	7654: {
		Fword:  __ccgo_ts + 53055,
		Ftype1: int8('F'),
	},
	7655: {
		Fword:  __ccgo_ts + 53061,
		Ftype1: int8('F'),
	},
	7656: {
		Fword:  __ccgo_ts + 53068,
		Ftype1: int8('F'),
	},
	7657: {
		Fword:  __ccgo_ts + 53075,
		Ftype1: int8('F'),
	},
	7658: {
		Fword:  __ccgo_ts + 53082,
		Ftype1: int8('F'),
	},
	7659: {
		Fword:  __ccgo_ts + 53089,
		Ftype1: int8('F'),
	},
	7660: {
		Fword:  __ccgo_ts + 53096,
		Ftype1: int8('F'),
	},
	7661: {
		Fword:  __ccgo_ts + 53103,
		Ftype1: int8('F'),
	},
	7662: {
		Fword:  __ccgo_ts + 53110,
		Ftype1: int8('F'),
	},
	7663: {
		Fword:  __ccgo_ts + 53116,
		Ftype1: int8('F'),
	},
	7664: {
		Fword:  __ccgo_ts + 53123,
		Ftype1: int8('F'),
	},
	7665: {
		Fword:  __ccgo_ts + 53130,
		Ftype1: int8('F'),
	},
	7666: {
		Fword:  __ccgo_ts + 53137,
		Ftype1: int8('F'),
	},
	7667: {
		Fword:  __ccgo_ts + 53144,
		Ftype1: int8('F'),
	},
	7668: {
		Fword:  __ccgo_ts + 53151,
		Ftype1: int8('F'),
	},
	7669: {
		Fword:  __ccgo_ts + 53158,
		Ftype1: int8('F'),
	},
	7670: {
		Fword:  __ccgo_ts + 53165,
		Ftype1: int8('F'),
	},
	7671: {
		Fword:  __ccgo_ts + 53172,
		Ftype1: int8('F'),
	},
	7672: {
		Fword:  __ccgo_ts + 53179,
		Ftype1: int8('F'),
	},
	7673: {
		Fword:  __ccgo_ts + 53186,
		Ftype1: int8('F'),
	},
	7674: {
		Fword:  __ccgo_ts + 53193,
		Ftype1: int8('F'),
	},
	7675: {
		Fword:  __ccgo_ts + 53197,
		Ftype1: int8('F'),
	},
	7676: {
		Fword:  __ccgo_ts + 53204,
		Ftype1: int8('F'),
	},
	7677: {
		Fword:  __ccgo_ts + 53211,
		Ftype1: int8('F'),
	},
	7678: {
		Fword:  __ccgo_ts + 53218,
		Ftype1: int8('F'),
	},
	7679: {
		Fword:  __ccgo_ts + 53225,
		Ftype1: int8('F'),
	},
	7680: {
		Fword:  __ccgo_ts + 53232,
		Ftype1: int8('F'),
	},
	7681: {
		Fword:  __ccgo_ts + 53239,
		Ftype1: int8('F'),
	},
	7682: {
		Fword:  __ccgo_ts + 53246,
		Ftype1: int8('F'),
	},
	7683: {
		Fword:  __ccgo_ts + 53253,
		Ftype1: int8('F'),
	},
	7684: {
		Fword:  __ccgo_ts + 53260,
		Ftype1: int8('F'),
	},
	7685: {
		Fword:  __ccgo_ts + 53267,
		Ftype1: int8('F'),
	},
	7686: {
		Fword:  __ccgo_ts + 53273,
		Ftype1: int8('F'),
	},
	7687: {
		Fword:  __ccgo_ts + 53280,
		Ftype1: int8('F'),
	},
	7688: {
		Fword:  __ccgo_ts + 53287,
		Ftype1: int8('F'),
	},
	7689: {
		Fword:  __ccgo_ts + 53294,
		Ftype1: int8('F'),
	},
	7690: {
		Fword:  __ccgo_ts + 53301,
		Ftype1: int8('F'),
	},
	7691: {
		Fword:  __ccgo_ts + 53308,
		Ftype1: int8('F'),
	},
	7692: {
		Fword:  __ccgo_ts + 53315,
		Ftype1: int8('F'),
	},
	7693: {
		Fword:  __ccgo_ts + 53322,
		Ftype1: int8('F'),
	},
	7694: {
		Fword:  __ccgo_ts + 53329,
		Ftype1: int8('F'),
	},
	7695: {
		Fword:  __ccgo_ts + 53336,
		Ftype1: int8('F'),
	},
	7696: {
		Fword:  __ccgo_ts + 53343,
		Ftype1: int8('F'),
	},
	7697: {
		Fword:  __ccgo_ts + 53350,
		Ftype1: int8('F'),
	},
	7698: {
		Fword:  __ccgo_ts + 53357,
		Ftype1: int8('F'),
	},
	7699: {
		Fword:  __ccgo_ts + 53364,
		Ftype1: int8('F'),
	},
	7700: {
		Fword:  __ccgo_ts + 53371,
		Ftype1: int8('F'),
	},
	7701: {
		Fword:  __ccgo_ts + 53378,
		Ftype1: int8('F'),
	},
	7702: {
		Fword:  __ccgo_ts + 53385,
		Ftype1: int8('F'),
	},
	7703: {
		Fword:  __ccgo_ts + 53392,
		Ftype1: int8('F'),
	},
	7704: {
		Fword:  __ccgo_ts + 53399,
		Ftype1: int8('F'),
	},
	7705: {
		Fword:  __ccgo_ts + 53406,
		Ftype1: int8('F'),
	},
	7706: {
		Fword:  __ccgo_ts + 53413,
		Ftype1: int8('F'),
	},
	7707: {
		Fword:  __ccgo_ts + 53420,
		Ftype1: int8('F'),
	},
	7708: {
		Fword:  __ccgo_ts + 53427,
		Ftype1: int8('F'),
	},
	7709: {
		Fword:  __ccgo_ts + 53434,
		Ftype1: int8('F'),
	},
	7710: {
		Fword:  __ccgo_ts + 53441,
		Ftype1: int8('F'),
	},
	7711: {
		Fword:  __ccgo_ts + 53448,
		Ftype1: int8('F'),
	},
	7712: {
		Fword:  __ccgo_ts + 53455,
		Ftype1: int8('F'),
	},
	7713: {
		Fword:  __ccgo_ts + 53462,
		Ftype1: int8('F'),
	},
	7714: {
		Fword:  __ccgo_ts + 53469,
		Ftype1: int8('F'),
	},
	7715: {
		Fword:  __ccgo_ts + 53476,
		Ftype1: int8('F'),
	},
	7716: {
		Fword:  __ccgo_ts + 53483,
		Ftype1: int8('F'),
	},
	7717: {
		Fword:  __ccgo_ts + 53490,
		Ftype1: int8('F'),
	},
	7718: {
		Fword:  __ccgo_ts + 53497,
		Ftype1: int8('F'),
	},
	7719: {
		Fword:  __ccgo_ts + 53504,
		Ftype1: int8('F'),
	},
	7720: {
		Fword:  __ccgo_ts + 53511,
		Ftype1: int8('F'),
	},
	7721: {
		Fword:  __ccgo_ts + 53518,
		Ftype1: int8('F'),
	},
	7722: {
		Fword:  __ccgo_ts + 53525,
		Ftype1: int8('F'),
	},
	7723: {
		Fword:  __ccgo_ts + 53532,
		Ftype1: int8('F'),
	},
	7724: {
		Fword:  __ccgo_ts + 53539,
		Ftype1: int8('F'),
	},
	7725: {
		Fword:  __ccgo_ts + 53546,
		Ftype1: int8('F'),
	},
	7726: {
		Fword:  __ccgo_ts + 53553,
		Ftype1: int8('F'),
	},
	7727: {
		Fword:  __ccgo_ts + 53560,
		Ftype1: int8('F'),
	},
	7728: {
		Fword:  __ccgo_ts + 53567,
		Ftype1: int8('F'),
	},
	7729: {
		Fword:  __ccgo_ts + 53574,
		Ftype1: int8('F'),
	},
	7730: {
		Fword:  __ccgo_ts + 53581,
		Ftype1: int8('F'),
	},
	7731: {
		Fword:  __ccgo_ts + 53588,
		Ftype1: int8('F'),
	},
	7732: {
		Fword:  __ccgo_ts + 53595,
		Ftype1: int8('F'),
	},
	7733: {
		Fword:  __ccgo_ts + 53602,
		Ftype1: int8('F'),
	},
	7734: {
		Fword:  __ccgo_ts + 53609,
		Ftype1: int8('F'),
	},
	7735: {
		Fword:  __ccgo_ts + 53616,
		Ftype1: int8('F'),
	},
	7736: {
		Fword:  __ccgo_ts + 53623,
		Ftype1: int8('F'),
	},
	7737: {
		Fword:  __ccgo_ts + 53630,
		Ftype1: int8('F'),
	},
	7738: {
		Fword:  __ccgo_ts + 53637,
		Ftype1: int8('F'),
	},
	7739: {
		Fword:  __ccgo_ts + 53643,
		Ftype1: int8('F'),
	},
	7740: {
		Fword:  __ccgo_ts + 53650,
		Ftype1: int8('F'),
	},
	7741: {
		Fword:  __ccgo_ts + 53657,
		Ftype1: int8('F'),
	},
	7742: {
		Fword:  __ccgo_ts + 53664,
		Ftype1: int8('F'),
	},
	7743: {
		Fword:  __ccgo_ts + 53671,
		Ftype1: int8('F'),
	},
	7744: {
		Fword:  __ccgo_ts + 53678,
		Ftype1: int8('F'),
	},
	7745: {
		Fword:  __ccgo_ts + 53685,
		Ftype1: int8('F'),
	},
	7746: {
		Fword:  __ccgo_ts + 53692,
		Ftype1: int8('F'),
	},
	7747: {
		Fword:  __ccgo_ts + 53699,
		Ftype1: int8('F'),
	},
	7748: {
		Fword:  __ccgo_ts + 53706,
		Ftype1: int8('F'),
	},
	7749: {
		Fword:  __ccgo_ts + 53713,
		Ftype1: int8('F'),
	},
	7750: {
		Fword:  __ccgo_ts + 53720,
		Ftype1: int8('F'),
	},
	7751: {
		Fword:  __ccgo_ts + 53727,
		Ftype1: int8('F'),
	},
	7752: {
		Fword:  __ccgo_ts + 53734,
		Ftype1: int8('F'),
	},
	7753: {
		Fword:  __ccgo_ts + 53741,
		Ftype1: int8('F'),
	},
	7754: {
		Fword:  __ccgo_ts + 53747,
		Ftype1: int8('F'),
	},
	7755: {
		Fword:  __ccgo_ts + 53754,
		Ftype1: int8('F'),
	},
	7756: {
		Fword:  __ccgo_ts + 53761,
		Ftype1: int8('F'),
	},
	7757: {
		Fword:  __ccgo_ts + 53768,
		Ftype1: int8('F'),
	},
	7758: {
		Fword:  __ccgo_ts + 53775,
		Ftype1: int8('F'),
	},
	7759: {
		Fword:  __ccgo_ts + 53782,
		Ftype1: int8('F'),
	},
	7760: {
		Fword:  __ccgo_ts + 53789,
		Ftype1: int8('F'),
	},
	7761: {
		Fword:  __ccgo_ts + 53796,
		Ftype1: int8('F'),
	},
	7762: {
		Fword:  __ccgo_ts + 53803,
		Ftype1: int8('F'),
	},
	7763: {
		Fword:  __ccgo_ts + 53810,
		Ftype1: int8('F'),
	},
	7764: {
		Fword:  __ccgo_ts + 53817,
		Ftype1: int8('F'),
	},
	7765: {
		Fword:  __ccgo_ts + 53824,
		Ftype1: int8('F'),
	},
	7766: {
		Fword:  __ccgo_ts + 53831,
		Ftype1: int8('F'),
	},
	7767: {
		Fword:  __ccgo_ts + 53838,
		Ftype1: int8('F'),
	},
	7768: {
		Fword:  __ccgo_ts + 53845,
		Ftype1: int8('F'),
	},
	7769: {
		Fword:  __ccgo_ts + 53852,
		Ftype1: int8('F'),
	},
	7770: {
		Fword:  __ccgo_ts + 53859,
		Ftype1: int8('F'),
	},
	7771: {
		Fword:  __ccgo_ts + 53866,
		Ftype1: int8('F'),
	},
	7772: {
		Fword:  __ccgo_ts + 53873,
		Ftype1: int8('F'),
	},
	7773: {
		Fword:  __ccgo_ts + 53880,
		Ftype1: int8('F'),
	},
	7774: {
		Fword:  __ccgo_ts + 53887,
		Ftype1: int8('F'),
	},
	7775: {
		Fword:  __ccgo_ts + 53894,
		Ftype1: int8('F'),
	},
	7776: {
		Fword:  __ccgo_ts + 53901,
		Ftype1: int8('F'),
	},
	7777: {
		Fword:  __ccgo_ts + 53908,
		Ftype1: int8('F'),
	},
	7778: {
		Fword:  __ccgo_ts + 53915,
		Ftype1: int8('F'),
	},
	7779: {
		Fword:  __ccgo_ts + 53922,
		Ftype1: int8('F'),
	},
	7780: {
		Fword:  __ccgo_ts + 53929,
		Ftype1: int8('F'),
	},
	7781: {
		Fword:  __ccgo_ts + 53936,
		Ftype1: int8('F'),
	},
	7782: {
		Fword:  __ccgo_ts + 53943,
		Ftype1: int8('F'),
	},
	7783: {
		Fword:  __ccgo_ts + 53950,
		Ftype1: int8('F'),
	},
	7784: {
		Fword:  __ccgo_ts + 53957,
		Ftype1: int8('F'),
	},
	7785: {
		Fword:  __ccgo_ts + 53964,
		Ftype1: int8('F'),
	},
	7786: {
		Fword:  __ccgo_ts + 53970,
		Ftype1: int8('F'),
	},
	7787: {
		Fword:  __ccgo_ts + 53977,
		Ftype1: int8('F'),
	},
	7788: {
		Fword:  __ccgo_ts + 53984,
		Ftype1: int8('F'),
	},
	7789: {
		Fword:  __ccgo_ts + 53991,
		Ftype1: int8('F'),
	},
	7790: {
		Fword:  __ccgo_ts + 53998,
		Ftype1: int8('F'),
	},
	7791: {
		Fword:  __ccgo_ts + 54005,
		Ftype1: int8('F'),
	},
	7792: {
		Fword:  __ccgo_ts + 54012,
		Ftype1: int8('F'),
	},
	7793: {
		Fword:  __ccgo_ts + 54019,
		Ftype1: int8('F'),
	},
	7794: {
		Fword:  __ccgo_ts + 54026,
		Ftype1: int8('F'),
	},
	7795: {
		Fword:  __ccgo_ts + 54033,
		Ftype1: int8('F'),
	},
	7796: {
		Fword:  __ccgo_ts + 54040,
		Ftype1: int8('F'),
	},
	7797: {
		Fword:  __ccgo_ts + 54047,
		Ftype1: int8('F'),
	},
	7798: {
		Fword:  __ccgo_ts + 54054,
		Ftype1: int8('F'),
	},
	7799: {
		Fword:  __ccgo_ts + 54061,
		Ftype1: int8('F'),
	},
	7800: {
		Fword:  __ccgo_ts + 54068,
		Ftype1: int8('F'),
	},
	7801: {
		Fword:  __ccgo_ts + 54075,
		Ftype1: int8('F'),
	},
	7802: {
		Fword:  __ccgo_ts + 54082,
		Ftype1: int8('F'),
	},
	7803: {
		Fword:  __ccgo_ts + 54089,
		Ftype1: int8('F'),
	},
	7804: {
		Fword:  __ccgo_ts + 54096,
		Ftype1: int8('F'),
	},
	7805: {
		Fword:  __ccgo_ts + 54103,
		Ftype1: int8('F'),
	},
	7806: {
		Fword:  __ccgo_ts + 54110,
		Ftype1: int8('F'),
	},
	7807: {
		Fword:  __ccgo_ts + 54117,
		Ftype1: int8('F'),
	},
	7808: {
		Fword:  __ccgo_ts + 54124,
		Ftype1: int8('F'),
	},
	7809: {
		Fword:  __ccgo_ts + 54131,
		Ftype1: int8('F'),
	},
	7810: {
		Fword:  __ccgo_ts + 54138,
		Ftype1: int8('F'),
	},
	7811: {
		Fword:  __ccgo_ts + 54145,
		Ftype1: int8('F'),
	},
	7812: {
		Fword:  __ccgo_ts + 54152,
		Ftype1: int8('F'),
	},
	7813: {
		Fword:  __ccgo_ts + 54159,
		Ftype1: int8('F'),
	},
	7814: {
		Fword:  __ccgo_ts + 54166,
		Ftype1: int8('F'),
	},
	7815: {
		Fword:  __ccgo_ts + 54173,
		Ftype1: int8('F'),
	},
	7816: {
		Fword:  __ccgo_ts + 54180,
		Ftype1: int8('F'),
	},
	7817: {
		Fword:  __ccgo_ts + 54187,
		Ftype1: int8('F'),
	},
	7818: {
		Fword:  __ccgo_ts + 54194,
		Ftype1: int8('F'),
	},
	7819: {
		Fword:  __ccgo_ts + 54201,
		Ftype1: int8('F'),
	},
	7820: {
		Fword:  __ccgo_ts + 54208,
		Ftype1: int8('F'),
	},
	7821: {
		Fword:  __ccgo_ts + 54215,
		Ftype1: int8('F'),
	},
	7822: {
		Fword:  __ccgo_ts + 54222,
		Ftype1: int8('F'),
	},
	7823: {
		Fword:  __ccgo_ts + 54229,
		Ftype1: int8('F'),
	},
	7824: {
		Fword:  __ccgo_ts + 54236,
		Ftype1: int8('F'),
	},
	7825: {
		Fword:  __ccgo_ts + 54243,
		Ftype1: int8('F'),
	},
	7826: {
		Fword:  __ccgo_ts + 54250,
		Ftype1: int8('F'),
	},
	7827: {
		Fword:  __ccgo_ts + 54257,
		Ftype1: int8('F'),
	},
	7828: {
		Fword:  __ccgo_ts + 54264,
		Ftype1: int8('F'),
	},
	7829: {
		Fword:  __ccgo_ts + 54271,
		Ftype1: int8('F'),
	},
	7830: {
		Fword:  __ccgo_ts + 54278,
		Ftype1: int8('F'),
	},
	7831: {
		Fword:  __ccgo_ts + 54285,
		Ftype1: int8('F'),
	},
	7832: {
		Fword:  __ccgo_ts + 54292,
		Ftype1: int8('F'),
	},
	7833: {
		Fword:  __ccgo_ts + 54299,
		Ftype1: int8('F'),
	},
	7834: {
		Fword:  __ccgo_ts + 54306,
		Ftype1: int8('F'),
	},
	7835: {
		Fword:  __ccgo_ts + 54313,
		Ftype1: int8('F'),
	},
	7836: {
		Fword:  __ccgo_ts + 54320,
		Ftype1: int8('F'),
	},
	7837: {
		Fword:  __ccgo_ts + 54327,
		Ftype1: int8('F'),
	},
	7838: {
		Fword:  __ccgo_ts + 54334,
		Ftype1: int8('F'),
	},
	7839: {
		Fword:  __ccgo_ts + 54341,
		Ftype1: int8('F'),
	},
	7840: {
		Fword:  __ccgo_ts + 54348,
		Ftype1: int8('F'),
	},
	7841: {
		Fword:  __ccgo_ts + 54355,
		Ftype1: int8('F'),
	},
	7842: {
		Fword:  __ccgo_ts + 54362,
		Ftype1: int8('F'),
	},
	7843: {
		Fword:  __ccgo_ts + 54369,
		Ftype1: int8('F'),
	},
	7844: {
		Fword:  __ccgo_ts + 54376,
		Ftype1: int8('F'),
	},
	7845: {
		Fword:  __ccgo_ts + 54383,
		Ftype1: int8('F'),
	},
	7846: {
		Fword:  __ccgo_ts + 54390,
		Ftype1: int8('F'),
	},
	7847: {
		Fword:  __ccgo_ts + 54397,
		Ftype1: int8('F'),
	},
	7848: {
		Fword:  __ccgo_ts + 54404,
		Ftype1: int8('F'),
	},
	7849: {
		Fword:  __ccgo_ts + 54411,
		Ftype1: int8('F'),
	},
	7850: {
		Fword:  __ccgo_ts + 54418,
		Ftype1: int8('F'),
	},
	7851: {
		Fword:  __ccgo_ts + 54425,
		Ftype1: int8('F'),
	},
	7852: {
		Fword:  __ccgo_ts + 54432,
		Ftype1: int8('F'),
	},
	7853: {
		Fword:  __ccgo_ts + 54439,
		Ftype1: int8('F'),
	},
	7854: {
		Fword:  __ccgo_ts + 54446,
		Ftype1: int8('F'),
	},
	7855: {
		Fword:  __ccgo_ts + 54453,
		Ftype1: int8('F'),
	},
	7856: {
		Fword:  __ccgo_ts + 54460,
		Ftype1: int8('F'),
	},
	7857: {
		Fword:  __ccgo_ts + 54465,
		Ftype1: int8('F'),
	},
	7858: {
		Fword:  __ccgo_ts + 54472,
		Ftype1: int8('F'),
	},
	7859: {
		Fword:  __ccgo_ts + 54479,
		Ftype1: int8('F'),
	},
	7860: {
		Fword:  __ccgo_ts + 54486,
		Ftype1: int8('F'),
	},
	7861: {
		Fword:  __ccgo_ts + 54493,
		Ftype1: int8('F'),
	},
	7862: {
		Fword:  __ccgo_ts + 54500,
		Ftype1: int8('F'),
	},
	7863: {
		Fword:  __ccgo_ts + 54507,
		Ftype1: int8('F'),
	},
	7864: {
		Fword:  __ccgo_ts + 54513,
		Ftype1: int8('F'),
	},
	7865: {
		Fword:  __ccgo_ts + 54520,
		Ftype1: int8('F'),
	},
	7866: {
		Fword:  __ccgo_ts + 54527,
		Ftype1: int8('F'),
	},
	7867: {
		Fword:  __ccgo_ts + 54534,
		Ftype1: int8('F'),
	},
	7868: {
		Fword:  __ccgo_ts + 54541,
		Ftype1: int8('F'),
	},
	7869: {
		Fword:  __ccgo_ts + 54548,
		Ftype1: int8('F'),
	},
	7870: {
		Fword:  __ccgo_ts + 54555,
		Ftype1: int8('F'),
	},
	7871: {
		Fword:  __ccgo_ts + 54562,
		Ftype1: int8('F'),
	},
	7872: {
		Fword:  __ccgo_ts + 54569,
		Ftype1: int8('F'),
	},
	7873: {
		Fword:  __ccgo_ts + 54576,
		Ftype1: int8('F'),
	},
	7874: {
		Fword:  __ccgo_ts + 54582,
		Ftype1: int8('F'),
	},
	7875: {
		Fword:  __ccgo_ts + 54589,
		Ftype1: int8('F'),
	},
	7876: {
		Fword:  __ccgo_ts + 54596,
		Ftype1: int8('F'),
	},
	7877: {
		Fword:  __ccgo_ts + 54603,
		Ftype1: int8('F'),
	},
	7878: {
		Fword:  __ccgo_ts + 54610,
		Ftype1: int8('F'),
	},
	7879: {
		Fword:  __ccgo_ts + 54617,
		Ftype1: int8('F'),
	},
	7880: {
		Fword:  __ccgo_ts + 54624,
		Ftype1: int8('F'),
	},
	7881: {
		Fword:  __ccgo_ts + 54631,
		Ftype1: int8('F'),
	},
	7882: {
		Fword:  __ccgo_ts + 54638,
		Ftype1: int8('F'),
	},
	7883: {
		Fword:  __ccgo_ts + 54645,
		Ftype1: int8('F'),
	},
	7884: {
		Fword:  __ccgo_ts + 54652,
		Ftype1: int8('F'),
	},
	7885: {
		Fword:  __ccgo_ts + 54659,
		Ftype1: int8('F'),
	},
	7886: {
		Fword:  __ccgo_ts + 54666,
		Ftype1: int8('F'),
	},
	7887: {
		Fword:  __ccgo_ts + 54673,
		Ftype1: int8('F'),
	},
	7888: {
		Fword:  __ccgo_ts + 54680,
		Ftype1: int8('F'),
	},
	7889: {
		Fword:  __ccgo_ts + 54687,
		Ftype1: int8('F'),
	},
	7890: {
		Fword:  __ccgo_ts + 54694,
		Ftype1: int8('F'),
	},
	7891: {
		Fword:  __ccgo_ts + 54701,
		Ftype1: int8('F'),
	},
	7892: {
		Fword:  __ccgo_ts + 54708,
		Ftype1: int8('F'),
	},
	7893: {
		Fword:  __ccgo_ts + 54715,
		Ftype1: int8('F'),
	},
	7894: {
		Fword:  __ccgo_ts + 54720,
		Ftype1: int8('F'),
	},
	7895: {
		Fword:  __ccgo_ts + 54727,
		Ftype1: int8('F'),
	},
	7896: {
		Fword:  __ccgo_ts + 54734,
		Ftype1: int8('F'),
	},
	7897: {
		Fword:  __ccgo_ts + 54741,
		Ftype1: int8('F'),
	},
	7898: {
		Fword:  __ccgo_ts + 54748,
		Ftype1: int8('F'),
	},
	7899: {
		Fword:  __ccgo_ts + 54755,
		Ftype1: int8('F'),
	},
	7900: {
		Fword:  __ccgo_ts + 54762,
		Ftype1: int8('F'),
	},
	7901: {
		Fword:  __ccgo_ts + 54768,
		Ftype1: int8('F'),
	},
	7902: {
		Fword:  __ccgo_ts + 54775,
		Ftype1: int8('F'),
	},
	7903: {
		Fword:  __ccgo_ts + 54782,
		Ftype1: int8('F'),
	},
	7904: {
		Fword:  __ccgo_ts + 54789,
		Ftype1: int8('F'),
	},
	7905: {
		Fword:  __ccgo_ts + 54796,
		Ftype1: int8('F'),
	},
	7906: {
		Fword:  __ccgo_ts + 54803,
		Ftype1: int8('F'),
	},
	7907: {
		Fword:  __ccgo_ts + 54810,
		Ftype1: int8('F'),
	},
	7908: {
		Fword:  __ccgo_ts + 54817,
		Ftype1: int8('F'),
	},
	7909: {
		Fword:  __ccgo_ts + 54824,
		Ftype1: int8('F'),
	},
	7910: {
		Fword:  __ccgo_ts + 54831,
		Ftype1: int8('F'),
	},
	7911: {
		Fword:  __ccgo_ts + 54837,
		Ftype1: int8('F'),
	},
	7912: {
		Fword:  __ccgo_ts + 54844,
		Ftype1: int8('F'),
	},
	7913: {
		Fword:  __ccgo_ts + 54851,
		Ftype1: int8('F'),
	},
	7914: {
		Fword:  __ccgo_ts + 54858,
		Ftype1: int8('F'),
	},
	7915: {
		Fword:  __ccgo_ts + 54865,
		Ftype1: int8('F'),
	},
	7916: {
		Fword:  __ccgo_ts + 54872,
		Ftype1: int8('F'),
	},
	7917: {
		Fword:  __ccgo_ts + 54879,
		Ftype1: int8('F'),
	},
	7918: {
		Fword:  __ccgo_ts + 54886,
		Ftype1: int8('F'),
	},
	7919: {
		Fword:  __ccgo_ts + 54893,
		Ftype1: int8('F'),
	},
	7920: {
		Fword:  __ccgo_ts + 54898,
		Ftype1: int8('F'),
	},
	7921: {
		Fword:  __ccgo_ts + 54905,
		Ftype1: int8('F'),
	},
	7922: {
		Fword:  __ccgo_ts + 54912,
		Ftype1: int8('F'),
	},
	7923: {
		Fword:  __ccgo_ts + 54919,
		Ftype1: int8('F'),
	},
	7924: {
		Fword:  __ccgo_ts + 54926,
		Ftype1: int8('F'),
	},
	7925: {
		Fword:  __ccgo_ts + 54933,
		Ftype1: int8('F'),
	},
	7926: {
		Fword:  __ccgo_ts + 54940,
		Ftype1: int8('F'),
	},
	7927: {
		Fword:  __ccgo_ts + 54946,
		Ftype1: int8('F'),
	},
	7928: {
		Fword:  __ccgo_ts + 54953,
		Ftype1: int8('F'),
	},
	7929: {
		Fword:  __ccgo_ts + 54960,
		Ftype1: int8('F'),
	},
	7930: {
		Fword:  __ccgo_ts + 54967,
		Ftype1: int8('F'),
	},
	7931: {
		Fword:  __ccgo_ts + 54974,
		Ftype1: int8('F'),
	},
	7932: {
		Fword:  __ccgo_ts + 54981,
		Ftype1: int8('F'),
	},
	7933: {
		Fword:  __ccgo_ts + 54988,
		Ftype1: int8('F'),
	},
	7934: {
		Fword:  __ccgo_ts + 54995,
		Ftype1: int8('F'),
	},
	7935: {
		Fword:  __ccgo_ts + 55002,
		Ftype1: int8('F'),
	},
	7936: {
		Fword:  __ccgo_ts + 55009,
		Ftype1: int8('F'),
	},
	7937: {
		Fword:  __ccgo_ts + 55015,
		Ftype1: int8('F'),
	},
	7938: {
		Fword:  __ccgo_ts + 55022,
		Ftype1: int8('F'),
	},
	7939: {
		Fword:  __ccgo_ts + 55029,
		Ftype1: int8('F'),
	},
	7940: {
		Fword:  __ccgo_ts + 55036,
		Ftype1: int8('F'),
	},
	7941: {
		Fword:  __ccgo_ts + 55043,
		Ftype1: int8('F'),
	},
	7942: {
		Fword:  __ccgo_ts + 55050,
		Ftype1: int8('F'),
	},
	7943: {
		Fword:  __ccgo_ts + 55057,
		Ftype1: int8('F'),
	},
	7944: {
		Fword:  __ccgo_ts + 55064,
		Ftype1: int8('F'),
	},
	7945: {
		Fword:  __ccgo_ts + 55071,
		Ftype1: int8('F'),
	},
	7946: {
		Fword:  __ccgo_ts + 55078,
		Ftype1: int8('F'),
	},
	7947: {
		Fword:  __ccgo_ts + 55085,
		Ftype1: int8('F'),
	},
	7948: {
		Fword:  __ccgo_ts + 55092,
		Ftype1: int8('F'),
	},
	7949: {
		Fword:  __ccgo_ts + 55099,
		Ftype1: int8('F'),
	},
	7950: {
		Fword:  __ccgo_ts + 55106,
		Ftype1: int8('F'),
	},
	7951: {
		Fword:  __ccgo_ts + 55113,
		Ftype1: int8('F'),
	},
	7952: {
		Fword:  __ccgo_ts + 55120,
		Ftype1: int8('F'),
	},
	7953: {
		Fword:  __ccgo_ts + 55127,
		Ftype1: int8('F'),
	},
	7954: {
		Fword:  __ccgo_ts + 55134,
		Ftype1: int8('F'),
	},
	7955: {
		Fword:  __ccgo_ts + 55141,
		Ftype1: int8('F'),
	},
	7956: {
		Fword:  __ccgo_ts + 55148,
		Ftype1: int8('F'),
	},
	7957: {
		Fword:  __ccgo_ts + 55155,
		Ftype1: int8('F'),
	},
	7958: {
		Fword:  __ccgo_ts + 55162,
		Ftype1: int8('F'),
	},
	7959: {
		Fword:  __ccgo_ts + 55169,
		Ftype1: int8('F'),
	},
	7960: {
		Fword:  __ccgo_ts + 55174,
		Ftype1: int8('F'),
	},
	7961: {
		Fword:  __ccgo_ts + 55181,
		Ftype1: int8('F'),
	},
	7962: {
		Fword:  __ccgo_ts + 55188,
		Ftype1: int8('F'),
	},
	7963: {
		Fword:  __ccgo_ts + 55195,
		Ftype1: int8('F'),
	},
	7964: {
		Fword:  __ccgo_ts + 55202,
		Ftype1: int8('F'),
	},
	7965: {
		Fword:  __ccgo_ts + 55209,
		Ftype1: int8('F'),
	},
	7966: {
		Fword:  __ccgo_ts + 55216,
		Ftype1: int8('F'),
	},
	7967: {
		Fword:  __ccgo_ts + 55222,
		Ftype1: int8('F'),
	},
	7968: {
		Fword:  __ccgo_ts + 55229,
		Ftype1: int8('F'),
	},
	7969: {
		Fword:  __ccgo_ts + 55236,
		Ftype1: int8('F'),
	},
	7970: {
		Fword:  __ccgo_ts + 55243,
		Ftype1: int8('F'),
	},
	7971: {
		Fword:  __ccgo_ts + 55250,
		Ftype1: int8('F'),
	},
	7972: {
		Fword:  __ccgo_ts + 55257,
		Ftype1: int8('F'),
	},
	7973: {
		Fword:  __ccgo_ts + 55264,
		Ftype1: int8('F'),
	},
	7974: {
		Fword:  __ccgo_ts + 55271,
		Ftype1: int8('F'),
	},
	7975: {
		Fword:  __ccgo_ts + 55278,
		Ftype1: int8('F'),
	},
	7976: {
		Fword:  __ccgo_ts + 55285,
		Ftype1: int8('F'),
	},
	7977: {
		Fword:  __ccgo_ts + 55291,
		Ftype1: int8('F'),
	},
	7978: {
		Fword:  __ccgo_ts + 55298,
		Ftype1: int8('F'),
	},
	7979: {
		Fword:  __ccgo_ts + 55305,
		Ftype1: int8('F'),
	},
	7980: {
		Fword:  __ccgo_ts + 55312,
		Ftype1: int8('F'),
	},
	7981: {
		Fword:  __ccgo_ts + 55319,
		Ftype1: int8('F'),
	},
	7982: {
		Fword:  __ccgo_ts + 55326,
		Ftype1: int8('F'),
	},
	7983: {
		Fword:  __ccgo_ts + 55333,
		Ftype1: int8('F'),
	},
	7984: {
		Fword:  __ccgo_ts + 55340,
		Ftype1: int8('F'),
	},
	7985: {
		Fword:  __ccgo_ts + 55347,
		Ftype1: int8('F'),
	},
	7986: {
		Fword:  __ccgo_ts + 55354,
		Ftype1: int8('F'),
	},
	7987: {
		Fword:  __ccgo_ts + 55361,
		Ftype1: int8('F'),
	},
	7988: {
		Fword:  __ccgo_ts + 55368,
		Ftype1: int8('F'),
	},
	7989: {
		Fword:  __ccgo_ts + 55375,
		Ftype1: int8('F'),
	},
	7990: {
		Fword:  __ccgo_ts + 55382,
		Ftype1: int8('F'),
	},
	7991: {
		Fword:  __ccgo_ts + 55389,
		Ftype1: int8('F'),
	},
	7992: {
		Fword:  __ccgo_ts + 55396,
		Ftype1: int8('F'),
	},
	7993: {
		Fword:  __ccgo_ts + 55403,
		Ftype1: int8('F'),
	},
	7994: {
		Fword:  __ccgo_ts + 55410,
		Ftype1: int8('F'),
	},
	7995: {
		Fword:  __ccgo_ts + 55417,
		Ftype1: int8('F'),
	},
	7996: {
		Fword:  __ccgo_ts + 55424,
		Ftype1: int8('F'),
	},
	7997: {
		Fword:  __ccgo_ts + 55431,
		Ftype1: int8('F'),
	},
	7998: {
		Fword:  __ccgo_ts + 55438,
		Ftype1: int8('F'),
	},
	7999: {
		Fword:  __ccgo_ts + 55445,
		Ftype1: int8('F'),
	},
	8000: {
		Fword:  __ccgo_ts + 55452,
		Ftype1: int8('F'),
	},
	8001: {
		Fword:  __ccgo_ts + 55459,
		Ftype1: int8('F'),
	},
	8002: {
		Fword:  __ccgo_ts + 55466,
		Ftype1: int8('F'),
	},
	8003: {
		Fword:  __ccgo_ts + 55473,
		Ftype1: int8('F'),
	},
	8004: {
		Fword:  __ccgo_ts + 55480,
		Ftype1: int8('F'),
	},
	8005: {
		Fword:  __ccgo_ts + 55487,
		Ftype1: int8('F'),
	},
	8006: {
		Fword:  __ccgo_ts + 55494,
		Ftype1: int8('F'),
	},
	8007: {
		Fword:  __ccgo_ts + 55501,
		Ftype1: int8('F'),
	},
	8008: {
		Fword:  __ccgo_ts + 55508,
		Ftype1: int8('F'),
	},
	8009: {
		Fword:  __ccgo_ts + 55515,
		Ftype1: int8('F'),
	},
	8010: {
		Fword:  __ccgo_ts + 55522,
		Ftype1: int8('F'),
	},
	8011: {
		Fword:  __ccgo_ts + 55529,
		Ftype1: int8('F'),
	},
	8012: {
		Fword:  __ccgo_ts + 55536,
		Ftype1: int8('F'),
	},
	8013: {
		Fword:  __ccgo_ts + 55543,
		Ftype1: int8('F'),
	},
	8014: {
		Fword:  __ccgo_ts + 55550,
		Ftype1: int8('F'),
	},
	8015: {
		Fword:  __ccgo_ts + 55557,
		Ftype1: int8('F'),
	},
	8016: {
		Fword:  __ccgo_ts + 55564,
		Ftype1: int8('F'),
	},
	8017: {
		Fword:  __ccgo_ts + 55571,
		Ftype1: int8('F'),
	},
	8018: {
		Fword:  __ccgo_ts + 55578,
		Ftype1: int8('F'),
	},
	8019: {
		Fword:  __ccgo_ts + 55585,
		Ftype1: int8('F'),
	},
	8020: {
		Fword:  __ccgo_ts + 55592,
		Ftype1: int8('F'),
	},
	8021: {
		Fword:  __ccgo_ts + 55599,
		Ftype1: int8('F'),
	},
	8022: {
		Fword:  __ccgo_ts + 55606,
		Ftype1: int8('F'),
	},
	8023: {
		Fword:  __ccgo_ts + 55613,
		Ftype1: int8('F'),
	},
	8024: {
		Fword:  __ccgo_ts + 55620,
		Ftype1: int8('F'),
	},
	8025: {
		Fword:  __ccgo_ts + 55627,
		Ftype1: int8('F'),
	},
	8026: {
		Fword:  __ccgo_ts + 55634,
		Ftype1: int8('F'),
	},
	8027: {
		Fword:  __ccgo_ts + 55641,
		Ftype1: int8('F'),
	},
	8028: {
		Fword:  __ccgo_ts + 55648,
		Ftype1: int8('F'),
	},
	8029: {
		Fword:  __ccgo_ts + 55655,
		Ftype1: int8('F'),
	},
	8030: {
		Fword:  __ccgo_ts + 55662,
		Ftype1: int8('F'),
	},
	8031: {
		Fword:  __ccgo_ts + 55669,
		Ftype1: int8('F'),
	},
	8032: {
		Fword:  __ccgo_ts + 55676,
		Ftype1: int8('F'),
	},
	8033: {
		Fword:  __ccgo_ts + 55683,
		Ftype1: int8('F'),
	},
	8034: {
		Fword:  __ccgo_ts + 55690,
		Ftype1: int8('F'),
	},
	8035: {
		Fword:  __ccgo_ts + 55697,
		Ftype1: int8('F'),
	},
	8036: {
		Fword:  __ccgo_ts + 55704,
		Ftype1: int8('F'),
	},
	8037: {
		Fword:  __ccgo_ts + 55711,
		Ftype1: int8('F'),
	},
	8038: {
		Fword:  __ccgo_ts + 55718,
		Ftype1: int8('F'),
	},
	8039: {
		Fword:  __ccgo_ts + 55725,
		Ftype1: int8('F'),
	},
	8040: {
		Fword:  __ccgo_ts + 55732,
		Ftype1: int8('F'),
	},
	8041: {
		Fword:  __ccgo_ts + 55739,
		Ftype1: int8('F'),
	},
	8042: {
		Fword:  __ccgo_ts + 55746,
		Ftype1: int8('F'),
	},
	8043: {
		Fword:  __ccgo_ts + 55753,
		Ftype1: int8('F'),
	},
	8044: {
		Fword:  __ccgo_ts + 55760,
		Ftype1: int8('F'),
	},
	8045: {
		Fword:  __ccgo_ts + 55767,
		Ftype1: int8('F'),
	},
	8046: {
		Fword:  __ccgo_ts + 55774,
		Ftype1: int8('F'),
	},
	8047: {
		Fword:  __ccgo_ts + 55781,
		Ftype1: int8('F'),
	},
	8048: {
		Fword:  __ccgo_ts + 55786,
		Ftype1: int8('F'),
	},
	8049: {
		Fword:  __ccgo_ts + 55793,
		Ftype1: int8('F'),
	},
	8050: {
		Fword:  __ccgo_ts + 55800,
		Ftype1: int8('F'),
	},
	8051: {
		Fword:  __ccgo_ts + 55807,
		Ftype1: int8('F'),
	},
	8052: {
		Fword:  __ccgo_ts + 55814,
		Ftype1: int8('F'),
	},
	8053: {
		Fword:  __ccgo_ts + 55821,
		Ftype1: int8('F'),
	},
	8054: {
		Fword:  __ccgo_ts + 55828,
		Ftype1: int8('F'),
	},
	8055: {
		Fword:  __ccgo_ts + 55835,
		Ftype1: int8('F'),
	},
	8056: {
		Fword:  __ccgo_ts + 55842,
		Ftype1: int8('F'),
	},
	8057: {
		Fword:  __ccgo_ts + 55849,
		Ftype1: int8('F'),
	},
	8058: {
		Fword:  __ccgo_ts + 55856,
		Ftype1: int8('F'),
	},
	8059: {
		Fword:  __ccgo_ts + 55863,
		Ftype1: int8('F'),
	},
	8060: {
		Fword:  __ccgo_ts + 55870,
		Ftype1: int8('F'),
	},
	8061: {
		Fword:  __ccgo_ts + 55877,
		Ftype1: int8('F'),
	},
	8062: {
		Fword:  __ccgo_ts + 55884,
		Ftype1: int8('F'),
	},
	8063: {
		Fword:  __ccgo_ts + 55891,
		Ftype1: int8('F'),
	},
	8064: {
		Fword:  __ccgo_ts + 55898,
		Ftype1: int8('F'),
	},
	8065: {
		Fword:  __ccgo_ts + 55905,
		Ftype1: int8('F'),
	},
	8066: {
		Fword:  __ccgo_ts + 55912,
		Ftype1: int8('F'),
	},
	8067: {
		Fword:  __ccgo_ts + 55919,
		Ftype1: int8('F'),
	},
	8068: {
		Fword:  __ccgo_ts + 55926,
		Ftype1: int8('F'),
	},
	8069: {
		Fword:  __ccgo_ts + 55933,
		Ftype1: int8('F'),
	},
	8070: {
		Fword:  __ccgo_ts + 55940,
		Ftype1: int8('F'),
	},
	8071: {
		Fword:  __ccgo_ts + 55947,
		Ftype1: int8('F'),
	},
	8072: {
		Fword:  __ccgo_ts + 55954,
		Ftype1: int8('F'),
	},
	8073: {
		Fword:  __ccgo_ts + 55961,
		Ftype1: int8('F'),
	},
	8074: {
		Fword:  __ccgo_ts + 55968,
		Ftype1: int8('F'),
	},
	8075: {
		Fword:  __ccgo_ts + 55975,
		Ftype1: int8('F'),
	},
	8076: {
		Fword:  __ccgo_ts + 55982,
		Ftype1: int8('F'),
	},
	8077: {
		Fword:  __ccgo_ts + 55989,
		Ftype1: int8('F'),
	},
	8078: {
		Fword:  __ccgo_ts + 55995,
		Ftype1: int8('F'),
	},
	8079: {
		Fword:  __ccgo_ts + 56002,
		Ftype1: int8('F'),
	},
	8080: {
		Fword:  __ccgo_ts + 56009,
		Ftype1: int8('F'),
	},
	8081: {
		Fword:  __ccgo_ts + 56016,
		Ftype1: int8('F'),
	},
	8082: {
		Fword:  __ccgo_ts + 56023,
		Ftype1: int8('F'),
	},
	8083: {
		Fword:  __ccgo_ts + 56030,
		Ftype1: int8('F'),
	},
	8084: {
		Fword:  __ccgo_ts + 56037,
		Ftype1: int8('F'),
	},
	8085: {
		Fword:  __ccgo_ts + 56044,
		Ftype1: int8('F'),
	},
	8086: {
		Fword:  __ccgo_ts + 56051,
		Ftype1: int8('F'),
	},
	8087: {
		Fword:  __ccgo_ts + 56058,
		Ftype1: int8('F'),
	},
	8088: {
		Fword:  __ccgo_ts + 56065,
		Ftype1: int8('F'),
	},
	8089: {
		Fword:  __ccgo_ts + 56072,
		Ftype1: int8('F'),
	},
	8090: {
		Fword:  __ccgo_ts + 56079,
		Ftype1: int8('F'),
	},
	8091: {
		Fword:  __ccgo_ts + 56086,
		Ftype1: int8('F'),
	},
	8092: {
		Fword:  __ccgo_ts + 56093,
		Ftype1: int8('F'),
	},
	8093: {
		Fword:  __ccgo_ts + 56100,
		Ftype1: int8('F'),
	},
	8094: {
		Fword:  __ccgo_ts + 56107,
		Ftype1: int8('F'),
	},
	8095: {
		Fword:  __ccgo_ts + 56113,
		Ftype1: int8('F'),
	},
	8096: {
		Fword:  __ccgo_ts + 56120,
		Ftype1: int8('F'),
	},
	8097: {
		Fword:  __ccgo_ts + 56127,
		Ftype1: int8('F'),
	},
	8098: {
		Fword:  __ccgo_ts + 56134,
		Ftype1: int8('F'),
	},
	8099: {
		Fword:  __ccgo_ts + 56141,
		Ftype1: int8('F'),
	},
	8100: {
		Fword:  __ccgo_ts + 56148,
		Ftype1: int8('F'),
	},
	8101: {
		Fword:  __ccgo_ts + 56155,
		Ftype1: int8('F'),
	},
	8102: {
		Fword:  __ccgo_ts + 56162,
		Ftype1: int8('F'),
	},
	8103: {
		Fword:  __ccgo_ts + 56169,
		Ftype1: int8('F'),
	},
	8104: {
		Fword:  __ccgo_ts + 56176,
		Ftype1: int8('F'),
	},
	8105: {
		Fword:  __ccgo_ts + 56183,
		Ftype1: int8('F'),
	},
	8106: {
		Fword:  __ccgo_ts + 56190,
		Ftype1: int8('F'),
	},
	8107: {
		Fword:  __ccgo_ts + 56197,
		Ftype1: int8('F'),
	},
	8108: {
		Fword:  __ccgo_ts + 56204,
		Ftype1: int8('F'),
	},
	8109: {
		Fword:  __ccgo_ts + 56211,
		Ftype1: int8('F'),
	},
	8110: {
		Fword:  __ccgo_ts + 56218,
		Ftype1: int8('F'),
	},
	8111: {
		Fword:  __ccgo_ts + 56225,
		Ftype1: int8('F'),
	},
	8112: {
		Fword:  __ccgo_ts + 56232,
		Ftype1: int8('F'),
	},
	8113: {
		Fword:  __ccgo_ts + 56239,
		Ftype1: int8('F'),
	},
	8114: {
		Fword:  __ccgo_ts + 56246,
		Ftype1: int8('F'),
	},
	8115: {
		Fword:  __ccgo_ts + 56253,
		Ftype1: int8('F'),
	},
	8116: {
		Fword:  __ccgo_ts + 56260,
		Ftype1: int8('F'),
	},
	8117: {
		Fword:  __ccgo_ts + 56267,
		Ftype1: int8('F'),
	},
	8118: {
		Fword:  __ccgo_ts + 56274,
		Ftype1: int8('F'),
	},
	8119: {
		Fword:  __ccgo_ts + 56281,
		Ftype1: int8('F'),
	},
	8120: {
		Fword:  __ccgo_ts + 56288,
		Ftype1: int8('F'),
	},
	8121: {
		Fword:  __ccgo_ts + 56295,
		Ftype1: int8('F'),
	},
	8122: {
		Fword:  __ccgo_ts + 56302,
		Ftype1: int8('F'),
	},
	8123: {
		Fword:  __ccgo_ts + 56308,
		Ftype1: int8('F'),
	},
	8124: {
		Fword:  __ccgo_ts + 56315,
		Ftype1: int8('F'),
	},
	8125: {
		Fword:  __ccgo_ts + 56322,
		Ftype1: int8('F'),
	},
	8126: {
		Fword:  __ccgo_ts + 56329,
		Ftype1: int8('F'),
	},
	8127: {
		Fword:  __ccgo_ts + 56336,
		Ftype1: int8('F'),
	},
	8128: {
		Fword:  __ccgo_ts + 56343,
		Ftype1: int8('F'),
	},
	8129: {
		Fword:  __ccgo_ts + 56350,
		Ftype1: int8('F'),
	},
	8130: {
		Fword:  __ccgo_ts + 56357,
		Ftype1: int8('F'),
	},
	8131: {
		Fword:  __ccgo_ts + 56364,
		Ftype1: int8('F'),
	},
	8132: {
		Fword:  __ccgo_ts + 56371,
		Ftype1: int8('F'),
	},
	8133: {
		Fword:  __ccgo_ts + 56378,
		Ftype1: int8('F'),
	},
	8134: {
		Fword:  __ccgo_ts + 56385,
		Ftype1: int8('F'),
	},
	8135: {
		Fword:  __ccgo_ts + 56392,
		Ftype1: int8('F'),
	},
	8136: {
		Fword:  __ccgo_ts + 56399,
		Ftype1: int8('F'),
	},
	8137: {
		Fword:  __ccgo_ts + 56406,
		Ftype1: int8('F'),
	},
	8138: {
		Fword:  __ccgo_ts + 56413,
		Ftype1: int8('F'),
	},
	8139: {
		Fword:  __ccgo_ts + 56420,
		Ftype1: int8('F'),
	},
	8140: {
		Fword:  __ccgo_ts + 56427,
		Ftype1: int8('F'),
	},
	8141: {
		Fword:  __ccgo_ts + 56434,
		Ftype1: int8('F'),
	},
	8142: {
		Fword:  __ccgo_ts + 56441,
		Ftype1: int8('F'),
	},
	8143: {
		Fword:  __ccgo_ts + 56448,
		Ftype1: int8('F'),
	},
	8144: {
		Fword:  __ccgo_ts + 56455,
		Ftype1: int8('F'),
	},
	8145: {
		Fword:  __ccgo_ts + 56462,
		Ftype1: int8('F'),
	},
	8146: {
		Fword:  __ccgo_ts + 56469,
		Ftype1: int8('F'),
	},
	8147: {
		Fword:  __ccgo_ts + 56476,
		Ftype1: int8('F'),
	},
	8148: {
		Fword:  __ccgo_ts + 56483,
		Ftype1: int8('F'),
	},
	8149: {
		Fword:  __ccgo_ts + 56490,
		Ftype1: int8('F'),
	},
	8150: {
		Fword:  __ccgo_ts + 56497,
		Ftype1: int8('F'),
	},
	8151: {
		Fword:  __ccgo_ts + 56504,
		Ftype1: int8('F'),
	},
	8152: {
		Fword:  __ccgo_ts + 56511,
		Ftype1: int8('F'),
	},
	8153: {
		Fword:  __ccgo_ts + 56518,
		Ftype1: int8('F'),
	},
	8154: {
		Fword:  __ccgo_ts + 56525,
		Ftype1: int8('F'),
	},
	8155: {
		Fword:  __ccgo_ts + 56532,
		Ftype1: int8('F'),
	},
	8156: {
		Fword:  __ccgo_ts + 56539,
		Ftype1: int8('F'),
	},
	8157: {
		Fword:  __ccgo_ts + 56546,
		Ftype1: int8('F'),
	},
	8158: {
		Fword:  __ccgo_ts + 56553,
		Ftype1: int8('F'),
	},
	8159: {
		Fword:  __ccgo_ts + 56560,
		Ftype1: int8('F'),
	},
	8160: {
		Fword:  __ccgo_ts + 56567,
		Ftype1: int8('F'),
	},
	8161: {
		Fword:  __ccgo_ts + 56574,
		Ftype1: int8('F'),
	},
	8162: {
		Fword:  __ccgo_ts + 56581,
		Ftype1: int8('F'),
	},
	8163: {
		Fword:  __ccgo_ts + 56588,
		Ftype1: int8('F'),
	},
	8164: {
		Fword:  __ccgo_ts + 56595,
		Ftype1: int8('F'),
	},
	8165: {
		Fword:  __ccgo_ts + 56602,
		Ftype1: int8('F'),
	},
	8166: {
		Fword:  __ccgo_ts + 56609,
		Ftype1: int8('F'),
	},
	8167: {
		Fword:  __ccgo_ts + 56616,
		Ftype1: int8('F'),
	},
	8168: {
		Fword:  __ccgo_ts + 56623,
		Ftype1: int8('F'),
	},
	8169: {
		Fword:  __ccgo_ts + 56630,
		Ftype1: int8('F'),
	},
	8170: {
		Fword:  __ccgo_ts + 56637,
		Ftype1: int8('F'),
	},
	8171: {
		Fword:  __ccgo_ts + 56644,
		Ftype1: int8('F'),
	},
	8172: {
		Fword:  __ccgo_ts + 56651,
		Ftype1: int8('F'),
	},
	8173: {
		Fword:  __ccgo_ts + 56658,
		Ftype1: int8('F'),
	},
	8174: {
		Fword:  __ccgo_ts + 56665,
		Ftype1: int8('F'),
	},
	8175: {
		Fword:  __ccgo_ts + 56672,
		Ftype1: int8('F'),
	},
	8176: {
		Fword:  __ccgo_ts + 56679,
		Ftype1: int8('F'),
	},
	8177: {
		Fword:  __ccgo_ts + 56686,
		Ftype1: int8('F'),
	},
	8178: {
		Fword:  __ccgo_ts + 56693,
		Ftype1: int8('F'),
	},
	8179: {
		Fword:  __ccgo_ts + 56700,
		Ftype1: int8('F'),
	},
	8180: {
		Fword:  __ccgo_ts + 56707,
		Ftype1: int8('F'),
	},
	8181: {
		Fword:  __ccgo_ts + 56714,
		Ftype1: int8('F'),
	},
	8182: {
		Fword:  __ccgo_ts + 56721,
		Ftype1: int8('F'),
	},
	8183: {
		Fword:  __ccgo_ts + 56728,
		Ftype1: int8('F'),
	},
	8184: {
		Fword:  __ccgo_ts + 56735,
		Ftype1: int8('F'),
	},
	8185: {
		Fword:  __ccgo_ts + 56742,
		Ftype1: int8('F'),
	},
	8186: {
		Fword:  __ccgo_ts + 56749,
		Ftype1: int8('F'),
	},
	8187: {
		Fword:  __ccgo_ts + 56756,
		Ftype1: int8('F'),
	},
	8188: {
		Fword:  __ccgo_ts + 56763,
		Ftype1: int8('F'),
	},
	8189: {
		Fword:  __ccgo_ts + 56770,
		Ftype1: int8('F'),
	},
	8190: {
		Fword:  __ccgo_ts + 56777,
		Ftype1: int8('F'),
	},
	8191: {
		Fword:  __ccgo_ts + 56784,
		Ftype1: int8('F'),
	},
	8192: {
		Fword:  __ccgo_ts + 56791,
		Ftype1: int8('F'),
	},
	8193: {
		Fword:  __ccgo_ts + 56798,
		Ftype1: int8('F'),
	},
	8194: {
		Fword:  __ccgo_ts + 56805,
		Ftype1: int8('F'),
	},
	8195: {
		Fword:  __ccgo_ts + 56812,
		Ftype1: int8('F'),
	},
	8196: {
		Fword:  __ccgo_ts + 56819,
		Ftype1: int8('F'),
	},
	8197: {
		Fword:  __ccgo_ts + 56826,
		Ftype1: int8('F'),
	},
	8198: {
		Fword:  __ccgo_ts + 56833,
		Ftype1: int8('F'),
	},
	8199: {
		Fword:  __ccgo_ts + 56840,
		Ftype1: int8('F'),
	},
	8200: {
		Fword:  __ccgo_ts + 56847,
		Ftype1: int8('F'),
	},
	8201: {
		Fword:  __ccgo_ts + 56854,
		Ftype1: int8('F'),
	},
	8202: {
		Fword:  __ccgo_ts + 56861,
		Ftype1: int8('F'),
	},
	8203: {
		Fword:  __ccgo_ts + 56868,
		Ftype1: int8('F'),
	},
	8204: {
		Fword:  __ccgo_ts + 56875,
		Ftype1: int8('F'),
	},
	8205: {
		Fword:  __ccgo_ts + 56882,
		Ftype1: int8('F'),
	},
	8206: {
		Fword:  __ccgo_ts + 56889,
		Ftype1: int8('F'),
	},
	8207: {
		Fword:  __ccgo_ts + 56896,
		Ftype1: int8('F'),
	},
	8208: {
		Fword:  __ccgo_ts + 56903,
		Ftype1: int8('F'),
	},
	8209: {
		Fword:  __ccgo_ts + 56910,
		Ftype1: int8('F'),
	},
	8210: {
		Fword:  __ccgo_ts + 56917,
		Ftype1: int8('F'),
	},
	8211: {
		Fword:  __ccgo_ts + 56924,
		Ftype1: int8('F'),
	},
	8212: {
		Fword:  __ccgo_ts + 56931,
		Ftype1: int8('F'),
	},
	8213: {
		Fword:  __ccgo_ts + 56938,
		Ftype1: int8('F'),
	},
	8214: {
		Fword:  __ccgo_ts + 56945,
		Ftype1: int8('F'),
	},
	8215: {
		Fword:  __ccgo_ts + 56952,
		Ftype1: int8('F'),
	},
	8216: {
		Fword:  __ccgo_ts + 56959,
		Ftype1: int8('F'),
	},
	8217: {
		Fword:  __ccgo_ts + 56966,
		Ftype1: int8('F'),
	},
	8218: {
		Fword:  __ccgo_ts + 56973,
		Ftype1: int8('F'),
	},
	8219: {
		Fword:  __ccgo_ts + 56980,
		Ftype1: int8('F'),
	},
	8220: {
		Fword:  __ccgo_ts + 56987,
		Ftype1: int8('F'),
	},
	8221: {
		Fword:  __ccgo_ts + 56994,
		Ftype1: int8('F'),
	},
	8222: {
		Fword:  __ccgo_ts + 57001,
		Ftype1: int8('F'),
	},
	8223: {
		Fword:  __ccgo_ts + 57008,
		Ftype1: int8('F'),
	},
	8224: {
		Fword:  __ccgo_ts + 57015,
		Ftype1: int8('F'),
	},
	8225: {
		Fword:  __ccgo_ts + 57022,
		Ftype1: int8('F'),
	},
	8226: {
		Fword:  __ccgo_ts + 57029,
		Ftype1: int8('F'),
	},
	8227: {
		Fword:  __ccgo_ts + 57036,
		Ftype1: int8('F'),
	},
	8228: {
		Fword:  __ccgo_ts + 57043,
		Ftype1: int8('F'),
	},
	8229: {
		Fword:  __ccgo_ts + 57050,
		Ftype1: int8('F'),
	},
	8230: {
		Fword:  __ccgo_ts + 57057,
		Ftype1: int8('F'),
	},
	8231: {
		Fword:  __ccgo_ts + 57064,
		Ftype1: int8('F'),
	},
	8232: {
		Fword:  __ccgo_ts + 57071,
		Ftype1: int8('F'),
	},
	8233: {
		Fword:  __ccgo_ts + 57075,
		Ftype1: int8('F'),
	},
	8234: {
		Fword:  __ccgo_ts + 57082,
		Ftype1: int8('F'),
	},
	8235: {
		Fword:  __ccgo_ts + 57089,
		Ftype1: int8('F'),
	},
	8236: {
		Fword:  __ccgo_ts + 57096,
		Ftype1: int8('F'),
	},
	8237: {
		Fword:  __ccgo_ts + 57103,
		Ftype1: int8('F'),
	},
	8238: {
		Fword:  __ccgo_ts + 57110,
		Ftype1: int8('F'),
	},
	8239: {
		Fword:  __ccgo_ts + 57117,
		Ftype1: int8('F'),
	},
	8240: {
		Fword:  __ccgo_ts + 57124,
		Ftype1: int8('F'),
	},
	8241: {
		Fword:  __ccgo_ts + 57131,
		Ftype1: int8('F'),
	},
	8242: {
		Fword:  __ccgo_ts + 57138,
		Ftype1: int8('F'),
	},
	8243: {
		Fword:  __ccgo_ts + 57145,
		Ftype1: int8('F'),
	},
	8244: {
		Fword:  __ccgo_ts + 57152,
		Ftype1: int8('F'),
	},
	8245: {
		Fword:  __ccgo_ts + 57159,
		Ftype1: int8('F'),
	},
	8246: {
		Fword:  __ccgo_ts + 57166,
		Ftype1: int8('F'),
	},
	8247: {
		Fword:  __ccgo_ts + 57173,
		Ftype1: int8('F'),
	},
	8248: {
		Fword:  __ccgo_ts + 57180,
		Ftype1: int8('F'),
	},
	8249: {
		Fword:  __ccgo_ts + 57187,
		Ftype1: int8('F'),
	},
	8250: {
		Fword:  __ccgo_ts + 57194,
		Ftype1: int8('F'),
	},
	8251: {
		Fword:  __ccgo_ts + 57201,
		Ftype1: int8('F'),
	},
	8252: {
		Fword:  __ccgo_ts + 57207,
		Ftype1: int8('F'),
	},
	8253: {
		Fword:  __ccgo_ts + 57214,
		Ftype1: int8('F'),
	},
	8254: {
		Fword:  __ccgo_ts + 57221,
		Ftype1: int8('F'),
	},
	8255: {
		Fword:  __ccgo_ts + 57228,
		Ftype1: int8('F'),
	},
	8256: {
		Fword:  __ccgo_ts + 57235,
		Ftype1: int8('F'),
	},
	8257: {
		Fword:  __ccgo_ts + 57240,
		Ftype1: int8('F'),
	},
	8258: {
		Fword:  __ccgo_ts + 57246,
		Ftype1: int8('F'),
	},
	8259: {
		Fword:  __ccgo_ts + 57251,
		Ftype1: int8('F'),
	},
	8260: {
		Fword:  __ccgo_ts + 57256,
		Ftype1: int8('F'),
	},
	8261: {
		Fword:  __ccgo_ts + 57263,
		Ftype1: int8('F'),
	},
	8262: {
		Fword:  __ccgo_ts + 57270,
		Ftype1: int8('F'),
	},
	8263: {
		Fword:  __ccgo_ts + 57277,
		Ftype1: int8('F'),
	},
	8264: {
		Fword:  __ccgo_ts + 57284,
		Ftype1: int8('F'),
	},
	8265: {
		Fword:  __ccgo_ts + 57291,
		Ftype1: int8('F'),
	},
	8266: {
		Fword:  __ccgo_ts + 57298,
		Ftype1: int8('F'),
	},
	8267: {
		Fword:  __ccgo_ts + 57305,
		Ftype1: int8('F'),
	},
	8268: {
		Fword:  __ccgo_ts + 57311,
		Ftype1: int8('F'),
	},
	8269: {
		Fword:  __ccgo_ts + 57318,
		Ftype1: int8('F'),
	},
	8270: {
		Fword:  __ccgo_ts + 57325,
		Ftype1: int8('F'),
	},
	8271: {
		Fword:  __ccgo_ts + 57332,
		Ftype1: int8('F'),
	},
	8272: {
		Fword:  __ccgo_ts + 57339,
		Ftype1: int8('F'),
	},
	8273: {
		Fword:  __ccgo_ts + 57346,
		Ftype1: int8('F'),
	},
	8274: {
		Fword:  __ccgo_ts + 57353,
		Ftype1: int8('F'),
	},
	8275: {
		Fword:  __ccgo_ts + 57360,
		Ftype1: int8('F'),
	},
	8276: {
		Fword:  __ccgo_ts + 57367,
		Ftype1: int8('F'),
	},
	8277: {
		Fword:  __ccgo_ts + 57374,
		Ftype1: int8('F'),
	},
	8278: {
		Fword:  __ccgo_ts + 57381,
		Ftype1: int8('F'),
	},
	8279: {
		Fword:  __ccgo_ts + 57388,
		Ftype1: int8('F'),
	},
	8280: {
		Fword:  __ccgo_ts + 57395,
		Ftype1: int8('F'),
	},
	8281: {
		Fword:  __ccgo_ts + 57402,
		Ftype1: int8('F'),
	},
	8282: {
		Fword:  __ccgo_ts + 57409,
		Ftype1: int8('F'),
	},
	8283: {
		Fword:  __ccgo_ts + 57415,
		Ftype1: int8('F'),
	},
	8284: {
		Fword:  __ccgo_ts + 57422,
		Ftype1: int8('F'),
	},
	8285: {
		Fword:  __ccgo_ts + 57428,
		Ftype1: int8('F'),
	},
	8286: {
		Fword:  __ccgo_ts + 57434,
		Ftype1: int8('F'),
	},
	8287: {
		Fword:  __ccgo_ts + 57441,
		Ftype1: int8('F'),
	},
	8288: {
		Fword:  __ccgo_ts + 57448,
		Ftype1: int8('F'),
	},
	8289: {
		Fword:  __ccgo_ts + 57455,
		Ftype1: int8('F'),
	},
	8290: {
		Fword:  __ccgo_ts + 57462,
		Ftype1: int8('F'),
	},
	8291: {
		Fword:  __ccgo_ts + 57468,
		Ftype1: int8('F'),
	},
	8292: {
		Fword:  __ccgo_ts + 57475,
		Ftype1: int8('F'),
	},
	8293: {
		Fword:  __ccgo_ts + 57482,
		Ftype1: int8('F'),
	},
	8294: {
		Fword:  __ccgo_ts + 57489,
		Ftype1: int8('F'),
	},
	8295: {
		Fword:  __ccgo_ts + 57496,
		Ftype1: int8('F'),
	},
	8296: {
		Fword:  __ccgo_ts + 57503,
		Ftype1: int8('F'),
	},
	8297: {
		Fword:  __ccgo_ts + 57510,
		Ftype1: int8('F'),
	},
	8298: {
		Fword:  __ccgo_ts + 57517,
		Ftype1: int8('F'),
	},
	8299: {
		Fword:  __ccgo_ts + 57524,
		Ftype1: int8('F'),
	},
	8300: {
		Fword:  __ccgo_ts + 57531,
		Ftype1: int8('F'),
	},
	8301: {
		Fword:  __ccgo_ts + 57537,
		Ftype1: int8('F'),
	},
	8302: {
		Fword:  __ccgo_ts + 57544,
		Ftype1: int8('F'),
	},
	8303: {
		Fword:  __ccgo_ts + 57551,
		Ftype1: int8('F'),
	},
	8304: {
		Fword:  __ccgo_ts + 57558,
		Ftype1: int8('F'),
	},
	8305: {
		Fword:  __ccgo_ts + 57565,
		Ftype1: int8('F'),
	},
	8306: {
		Fword:  __ccgo_ts + 57572,
		Ftype1: int8('F'),
	},
	8307: {
		Fword:  __ccgo_ts + 57579,
		Ftype1: int8('F'),
	},
	8308: {
		Fword:  __ccgo_ts + 57586,
		Ftype1: int8('F'),
	},
	8309: {
		Fword:  __ccgo_ts + 57593,
		Ftype1: int8('F'),
	},
	8310: {
		Fword:  __ccgo_ts + 57600,
		Ftype1: int8('F'),
	},
	8311: {
		Fword:  __ccgo_ts + 57607,
		Ftype1: int8('F'),
	},
	8312: {
		Fword:  __ccgo_ts + 57614,
		Ftype1: int8('F'),
	},
	8313: {
		Fword:  __ccgo_ts + 57621,
		Ftype1: int8('F'),
	},
	8314: {
		Fword:  __ccgo_ts + 57628,
		Ftype1: int8('F'),
	},
	8315: {
		Fword:  __ccgo_ts + 57635,
		Ftype1: int8('F'),
	},
	8316: {
		Fword:  __ccgo_ts + 57642,
		Ftype1: int8('F'),
	},
	8317: {
		Fword:  __ccgo_ts + 57648,
		Ftype1: int8('F'),
	},
	8318: {
		Fword:  __ccgo_ts + 57655,
		Ftype1: int8('F'),
	},
	8319: {
		Fword:  __ccgo_ts + 57662,
		Ftype1: int8('F'),
	},
	8320: {
		Fword:  __ccgo_ts + 57669,
		Ftype1: int8('F'),
	},
	8321: {
		Fword:  __ccgo_ts + 57676,
		Ftype1: int8('F'),
	},
	8322: {
		Fword:  __ccgo_ts + 57683,
		Ftype1: int8('F'),
	},
	8323: {
		Fword:  __ccgo_ts + 57690,
		Ftype1: int8('F'),
	},
	8324: {
		Fword:  __ccgo_ts + 57697,
		Ftype1: int8('F'),
	},
	8325: {
		Fword:  __ccgo_ts + 57704,
		Ftype1: int8('F'),
	},
	8326: {
		Fword:  __ccgo_ts + 57711,
		Ftype1: int8('F'),
	},
	8327: {
		Fword:  __ccgo_ts + 57718,
		Ftype1: int8('F'),
	},
	8328: {
		Fword:  __ccgo_ts + 57725,
		Ftype1: int8('F'),
	},
	8329: {
		Fword:  __ccgo_ts + 57732,
		Ftype1: int8('F'),
	},
	8330: {
		Fword:  __ccgo_ts + 57739,
		Ftype1: int8('F'),
	},
	8331: {
		Fword:  __ccgo_ts + 57745,
		Ftype1: int8('F'),
	},
	8332: {
		Fword:  __ccgo_ts + 57752,
		Ftype1: int8('F'),
	},
	8333: {
		Fword:  __ccgo_ts + 57759,
		Ftype1: int8('F'),
	},
	8334: {
		Fword:  __ccgo_ts + 57766,
		Ftype1: int8('F'),
	},
	8335: {
		Fword:  __ccgo_ts + 57773,
		Ftype1: int8('F'),
	},
	8336: {
		Fword:  __ccgo_ts + 57780,
		Ftype1: int8('F'),
	},
	8337: {
		Fword:  __ccgo_ts + 57787,
		Ftype1: int8('F'),
	},
	8338: {
		Fword:  __ccgo_ts + 57794,
		Ftype1: int8('F'),
	},
	8339: {
		Fword:  __ccgo_ts + 57801,
		Ftype1: int8('F'),
	},
	8340: {
		Fword:  __ccgo_ts + 57808,
		Ftype1: int8('F'),
	},
	8341: {
		Fword:  __ccgo_ts + 57815,
		Ftype1: int8('F'),
	},
	8342: {
		Fword:  __ccgo_ts + 57822,
		Ftype1: int8('F'),
	},
	8343: {
		Fword:  __ccgo_ts + 57829,
		Ftype1: int8('F'),
	},
	8344: {
		Fword:  __ccgo_ts + 57836,
		Ftype1: int8('F'),
	},
	8345: {
		Fword:  __ccgo_ts + 57843,
		Ftype1: int8('F'),
	},
	8346: {
		Fword:  __ccgo_ts + 57850,
		Ftype1: int8('F'),
	},
	8347: {
		Fword:  __ccgo_ts + 57857,
		Ftype1: int8('F'),
	},
	8348: {
		Fword:  __ccgo_ts + 57864,
		Ftype1: int8('F'),
	},
	8349: {
		Fword:  __ccgo_ts + 57871,
		Ftype1: int8('F'),
	},
	8350: {
		Fword:  __ccgo_ts + 57878,
		Ftype1: int8('F'),
	},
	8351: {
		Fword:  __ccgo_ts + 57885,
		Ftype1: int8('F'),
	},
	8352: {
		Fword:  __ccgo_ts + 57892,
		Ftype1: int8('F'),
	},
	8353: {
		Fword:  __ccgo_ts + 57899,
		Ftype1: int8('F'),
	},
	8354: {
		Fword:  __ccgo_ts + 57906,
		Ftype1: int8('F'),
	},
	8355: {
		Fword:  __ccgo_ts + 57913,
		Ftype1: int8('F'),
	},
	8356: {
		Fword:  __ccgo_ts + 57920,
		Ftype1: int8('F'),
	},
	8357: {
		Fword:  __ccgo_ts + 57927,
		Ftype1: int8('F'),
	},
	8358: {
		Fword:  __ccgo_ts + 57934,
		Ftype1: int8('F'),
	},
	8359: {
		Fword:  __ccgo_ts + 57940,
		Ftype1: int8('F'),
	},
	8360: {
		Fword:  __ccgo_ts + 57947,
		Ftype1: int8('F'),
	},
	8361: {
		Fword:  __ccgo_ts + 57954,
		Ftype1: int8('F'),
	},
	8362: {
		Fword:  __ccgo_ts + 57961,
		Ftype1: int8('F'),
	},
	8363: {
		Fword:  __ccgo_ts + 57968,
		Ftype1: int8('F'),
	},
	8364: {
		Fword:  __ccgo_ts + 57975,
		Ftype1: int8('F'),
	},
	8365: {
		Fword:  __ccgo_ts + 57982,
		Ftype1: int8('F'),
	},
	8366: {
		Fword:  __ccgo_ts + 57989,
		Ftype1: int8('F'),
	},
	8367: {
		Fword:  __ccgo_ts + 57996,
		Ftype1: int8('F'),
	},
	8368: {
		Fword:  __ccgo_ts + 58003,
		Ftype1: int8('F'),
	},
	8369: {
		Fword:  __ccgo_ts + 58010,
		Ftype1: int8('F'),
	},
	8370: {
		Fword:  __ccgo_ts + 58017,
		Ftype1: int8('F'),
	},
	8371: {
		Fword:  __ccgo_ts + 58024,
		Ftype1: int8('F'),
	},
	8372: {
		Fword:  __ccgo_ts + 58031,
		Ftype1: int8('F'),
	},
	8373: {
		Fword:  __ccgo_ts + 58038,
		Ftype1: int8('F'),
	},
	8374: {
		Fword:  __ccgo_ts + 58044,
		Ftype1: int8('F'),
	},
	8375: {
		Fword:  __ccgo_ts + 58051,
		Ftype1: int8('F'),
	},
	8376: {
		Fword:  __ccgo_ts + 58058,
		Ftype1: int8('F'),
	},
	8377: {
		Fword:  __ccgo_ts + 58065,
		Ftype1: int8('F'),
	},
	8378: {
		Fword:  __ccgo_ts + 58068,
		Ftype1: int8('o'),
	},
	8379: {
		Fword:  __ccgo_ts + 58071,
		Ftype1: int8('o'),
	},
	8380: {
		Fword:  __ccgo_ts + 58074,
		Ftype1: int8('o'),
	},
	8381: {
		Fword:  __ccgo_ts + 58077,
		Ftype1: int8('o'),
	},
	8382: {
		Fword:  __ccgo_ts + 58080,
		Ftype1: int8('o'),
	},
	8383: {
		Fword:  __ccgo_ts + 58083,
		Ftype1: int8('o'),
	},
	8384: {
		Fword:  __ccgo_ts + 58086,
		Ftype1: int8('o'),
	},
	8385: {
		Fword:  __ccgo_ts + 58089,
		Ftype1: int8('o'),
	},
	8386: {
		Fword:  __ccgo_ts + 58092,
		Ftype1: int8('o'),
	},
	8387: {
		Fword:  __ccgo_ts + 58095,
		Ftype1: int8('k'),
	},
	8388: {
		Fword:  __ccgo_ts + 58101,
		Ftype1: int8('f'),
	},
	8389: {
		Fword:  __ccgo_ts + 58105,
		Ftype1: int8('k'),
	},
	8390: {
		Fword:  __ccgo_ts + 58116,
		Ftype1: int8('f'),
	},
	8391: {
		Fword:  __ccgo_ts + 58121,
		Ftype1: int8('f'),
	},
	8392: {
		Fword:  __ccgo_ts + 58129,
		Ftype1: int8('f'),
	},
	8393: {
		Fword:  __ccgo_ts + 58137,
		Ftype1: int8('f'),
	},
	8394: {
		Fword:  __ccgo_ts + 58149,
		Ftype1: int8('f'),
	},
	8395: {
		Fword:  __ccgo_ts + 58161,
		Ftype1: int8('k'),
	},
	8396: {
		Fword:  __ccgo_ts + 58169,
		Ftype1: int8('f'),
	},
	8397: {
		Fword:  __ccgo_ts + 58173,
		Ftype1: int8('k'),
	},
	8398: {
		Fword:  __ccgo_ts + 58183,
		Ftype1: int8('k'),
	},
	8399: {
		Fword:  __ccgo_ts + 58189,
		Ftype1: int8('k'),
	},
	8400: {
		Fword:  __ccgo_ts + 58202,
		Ftype1: int8('k'),
	},
	8401: {
		Fword:  __ccgo_ts + 58214,
		Ftype1: int8('k'),
	},
	8402: {
		Fword:  __ccgo_ts + 58222,
		Ftype1: int8('&'),
	},
	8403: {
		Fword:  __ccgo_ts + 58226,
		Ftype1: int8('f'),
	},
	8404: {
		Fword:  __ccgo_ts + 58230,
		Ftype1: int8('t'),
	},
	8405: {
		Fword:  __ccgo_ts + 58239,
		Ftype1: int8('t'),
	},
	8406: {
		Fword:  __ccgo_ts + 58250,
		Ftype1: int8('t'),
	},
	8407: {
		Fword:  __ccgo_ts + 58261,
		Ftype1: int8('f'),
	},
	8408: {
		Fword:  __ccgo_ts + 58274,
		Ftype1: int8('f'),
	},
	8409: {
		Fword:  __ccgo_ts + 58287,
		Ftype1: int8('f'),
	},
	8410: {
		Fword:  __ccgo_ts + 58296,
		Ftype1: int8('f'),
	},
	8411: {
		Fword:  __ccgo_ts + 58306,
		Ftype1: int8('f'),
	},
	8412: {
		Fword:  __ccgo_ts + 58316,
		Ftype1: int8('f'),
	},
	8413: {
		Fword:  __ccgo_ts + 58326,
		Ftype1: int8('f'),
	},
	8414: {
		Fword:  __ccgo_ts + 58337,
		Ftype1: int8('f'),
	},
	8415: {
		Fword:  __ccgo_ts + 58350,
		Ftype1: int8('f'),
	},
	8416: {
		Fword:  __ccgo_ts + 58362,
		Ftype1: int8('f'),
	},
	8417: {
		Fword:  __ccgo_ts + 58374,
		Ftype1: int8('f'),
	},
	8418: {
		Fword:  __ccgo_ts + 58388,
		Ftype1: int8('f'),
	},
	8419: {
		Fword:  __ccgo_ts + 58402,
		Ftype1: int8('f'),
	},
	8420: {
		Fword:  __ccgo_ts + 58418,
		Ftype1: int8('f'),
	},
	8421: {
		Fword:  __ccgo_ts + 58430,
		Ftype1: int8('k'),
	},
	8422: {
		Fword:  __ccgo_ts + 58433,
		Ftype1: int8('k'),
	},
	8423: {
		Fword:  __ccgo_ts + 58437,
		Ftype1: int8('f'),
	},
	8424: {
		Fword:  __ccgo_ts + 58443,
		Ftype1: int8('k'),
	},
	8425: {
		Fword:  __ccgo_ts + 58454,
		Ftype1: int8('f'),
	},
	8426: {
		Fword:  __ccgo_ts + 58459,
		Ftype1: int8('f'),
	},
	8427: {
		Fword:  __ccgo_ts + 58476,
		Ftype1: int8('f'),
	},
	8428: {
		Fword:  __ccgo_ts + 58487,
		Ftype1: int8('n'),
	},
	8429: {
		Fword:  __ccgo_ts + 58495,
		Ftype1: int8('k'),
	},
	8430: {
		Fword:  __ccgo_ts + 58508,
		Ftype1: int8('f'),
	},
	8431: {
		Fword:  __ccgo_ts + 58513,
		Ftype1: int8('f'),
	},
	8432: {
		Fword:  __ccgo_ts + 58519,
		Ftype1: int8('k'),
	},
	8433: {
		Fword:  __ccgo_ts + 58533,
		Ftype1: int8('f'),
	},
	8434: {
		Fword:  __ccgo_ts + 58537,
		Ftype1: int8('k'),
	},
	8435: {
		Fword:  __ccgo_ts + 58544,
		Ftype1: int8('T'),
	},
	8436: {
		Fword:  __ccgo_ts + 58550,
		Ftype1: int8('T'),
	},
	8437: {
		Fword:  __ccgo_ts + 58564,
		Ftype1: int8('T'),
	},
	8438: {
		Fword:  __ccgo_ts + 58575,
		Ftype1: int8('T'),
	},
	8439: {
		Fword:  __ccgo_ts + 58585,
		Ftype1: int8('T'),
	},
	8440: {
		Fword:  __ccgo_ts + 58603,
		Ftype1: int8('f'),
	},
	8441: {
		Fword:  __ccgo_ts + 58613,
		Ftype1: int8('o'),
	},
	8442: {
		Fword:  __ccgo_ts + 58621,
		Ftype1: int8('t'),
	},
	8443: {
		Fword:  __ccgo_ts + 58628,
		Ftype1: int8('t'),
	},
	8444: {
		Fword:  __ccgo_ts + 58638,
		Ftype1: int8('f'),
	},
	8445: {
		Fword:  __ccgo_ts + 58642,
		Ftype1: int8('t'),
	},
	8446: {
		Fword:  __ccgo_ts + 58649,
		Ftype1: int8('1'),
	},
	8447: {
		Fword:  __ccgo_ts + 58672,
		Ftype1: int8('1'),
	},
	8448: {
		Fword:  __ccgo_ts + 58690,
		Ftype1: int8('1'),
	},
	8449: {
		Fword:  __ccgo_ts + 58712,
		Ftype1: int8('1'),
	},
	8450: {
		Fword:  __ccgo_ts + 58729,
		Ftype1: int8('f'),
	},
	8451: {
		Fword:  __ccgo_ts + 58739,
		Ftype1: int8('f'),
	},
	8452: {
		Fword:  __ccgo_ts + 58747,
		Ftype1: int8('f'),
	},
	8453: {
		Fword:  __ccgo_ts + 58757,
		Ftype1: int8('f'),
	},
	8454: {
		Fword:  __ccgo_ts + 58768,
		Ftype1: int8('f'),
	},
	8455: {
		Fword:  __ccgo_ts + 58775,
		Ftype1: int8('f'),
	},
	8456: {
		Fword:  __ccgo_ts + 58783,
		Ftype1: int8('k'),
	},
	8457: {
		Fword:  __ccgo_ts + 58788,
		Ftype1: int8('t'),
	},
	8458: {
		Fword:  __ccgo_ts + 58796,
		Ftype1: int8('f'),
	},
	8459: {
		Fword:  __ccgo_ts + 58805,
		Ftype1: int8('f'),
	},
	8460: {
		Fword:  __ccgo_ts + 58813,
		Ftype1: int8('k'),
	},
	8461: {
		Fword:  __ccgo_ts + 58818,
		Ftype1: int8('f'),
	},
	8462: {
		Fword:  __ccgo_ts + 58824,
		Ftype1: int8('n'),
	},
	8463: {
		Fword:  __ccgo_ts + 58827,
		Ftype1: int8('t'),
	},
	8464: {
		Fword:  __ccgo_ts + 58833,
		Ftype1: int8('T'),
	},
	8465: {
		Fword:  __ccgo_ts + 58838,
		Ftype1: int8('k'),
	},
	8466: {
		Fword:  __ccgo_ts + 58846,
		Ftype1: int8('E'),
	},
	8467: {
		Fword:  __ccgo_ts + 58851,
		Ftype1: int8('f'),
	},
	8468: {
		Fword:  __ccgo_ts + 58856,
		Ftype1: int8('f'),
	},
	8469: {
		Fword:  __ccgo_ts + 58862,
		Ftype1: int8('f'),
	},
	8470: {
		Fword:  __ccgo_ts + 58867,
		Ftype1: int8('f'),
	},
	8471: {
		Fword:  __ccgo_ts + 58873,
		Ftype1: int8('f'),
	},
	8472: {
		Fword:  __ccgo_ts + 58878,
		Ftype1: int8('f'),
	},
	8473: {
		Fword:  __ccgo_ts + 58884,
		Ftype1: int8('f'),
	},
	8474: {
		Fword:  __ccgo_ts + 58889,
		Ftype1: int8('f'),
	},
	8475: {
		Fword:  __ccgo_ts + 58894,
		Ftype1: int8('f'),
	},
	8476: {
		Fword:  __ccgo_ts + 58902,
		Ftype1: int8('f'),
	},
	8477: {
		Fword:  __ccgo_ts + 58914,
		Ftype1: int8('f'),
	},
	8478: {
		Fword:  __ccgo_ts + 58929,
		Ftype1: int8('f'),
	},
	8479: {
		Fword:  __ccgo_ts + 58937,
		Ftype1: int8('f'),
	},
	8480: {
		Fword:  __ccgo_ts + 58951,
		Ftype1: int8('k'),
	},
	8481: {
		Fword:  __ccgo_ts + 58958,
		Ftype1: int8('f'),
	},
	8482: {
		Fword:  __ccgo_ts + 58966,
		Ftype1: int8('f'),
	},
	8483: {
		Fword:  __ccgo_ts + 58971,
		Ftype1: int8('t'),
	},
	8484: {
		Fword:  __ccgo_ts + 58981,
		Ftype1: int8('t'),
	},
	8485: {
		Fword:  __ccgo_ts + 58999,
		Ftype1: int8('f'),
	},
	8486: {
		Fword:  __ccgo_ts + 59016,
		Ftype1: int8('f'),
	},
	8487: {
		Fword:  __ccgo_ts + 59026,
		Ftype1: int8('f'),
	},
	8488: {
		Fword:  __ccgo_ts + 59034,
		Ftype1: int8('f'),
	},
	8489: {
		Fword:  __ccgo_ts + 59046,
		Ftype1: int8('f'),
	},
	8490: {
		Fword:  __ccgo_ts + 59052,
		Ftype1: int8('f'),
	},
	8491: {
		Fword:  __ccgo_ts + 59060,
		Ftype1: int8('n'),
	},
	8492: {
		Fword:  __ccgo_ts + 59066,
		Ftype1: int8('f'),
	},
	8493: {
		Fword:  __ccgo_ts + 59079,
		Ftype1: int8('f'),
	},
	8494: {
		Fword:  __ccgo_ts + 59086,
		Ftype1: int8('f'),
	},
	8495: {
		Fword:  __ccgo_ts + 59090,
		Ftype1: int8('f'),
	},
	8496: {
		Fword:  __ccgo_ts + 59095,
		Ftype1: int8('f'),
	},
	8497: {
		Fword:  __ccgo_ts + 59100,
		Ftype1: int8('f'),
	},
	8498: {
		Fword:  __ccgo_ts + 59116,
		Ftype1: int8('f'),
	},
	8499: {
		Fword:  __ccgo_ts + 59125,
		Ftype1: int8('f'),
	},
	8500: {
		Fword:  __ccgo_ts + 59138,
		Ftype1: int8('A'),
	},
	8501: {
		Fword:  __ccgo_ts + 59146,
		Ftype1: int8('f'),
	},
	8502: {
		Fword:  __ccgo_ts + 59156,
		Ftype1: int8('f'),
	},
	8503: {
		Fword:  __ccgo_ts + 59174,
		Ftype1: int8('k'),
	},
	8504: {
		Fword:  __ccgo_ts + 59181,
		Ftype1: int8('f'),
	},
	8505: {
		Fword:  __ccgo_ts + 59196,
		Ftype1: int8('f'),
	},
	8506: {
		Fword:  __ccgo_ts + 59212,
		Ftype1: int8('f'),
	},
	8507: {
		Fword:  __ccgo_ts + 59223,
		Ftype1: int8('f'),
	},
	8508: {
		Fword:  __ccgo_ts + 59232,
		Ftype1: int8('f'),
	},
	8509: {
		Fword:  __ccgo_ts + 59241,
		Ftype1: int8('f'),
	},
	8510: {
		Fword:  __ccgo_ts + 59248,
		Ftype1: int8('f'),
	},
	8511: {
		Fword:  __ccgo_ts + 59258,
		Ftype1: int8('k'),
	},
	8512: {
		Fword:  __ccgo_ts + 59268,
		Ftype1: int8('f'),
	},
	8513: {
		Fword:  __ccgo_ts + 59282,
		Ftype1: int8('k'),
	},
	8514: {
		Fword:  __ccgo_ts + 59293,
		Ftype1: int8('k'),
	},
	8515: {
		Fword:  __ccgo_ts + 59302,
		Ftype1: int8('f'),
	},
	8516: {
		Fword:  __ccgo_ts + 59307,
		Ftype1: int8('f'),
	},
	8517: {
		Fword:  __ccgo_ts + 59315,
		Ftype1: int8('f'),
	},
	8518: {
		Fword:  __ccgo_ts + 59328,
		Ftype1: int8('f'),
	},
	8519: {
		Fword:  __ccgo_ts + 59339,
		Ftype1: int8('f'),
	},
	8520: {
		Fword:  __ccgo_ts + 59350,
		Ftype1: int8('f'),
	},
	8521: {
		Fword:  __ccgo_ts + 59354,
		Ftype1: int8('f'),
	},
	8522: {
		Fword:  __ccgo_ts + 59358,
		Ftype1: int8('f'),
	},
	8523: {
		Fword:  __ccgo_ts + 59364,
		Ftype1: int8('k'),
	},
	8524: {
		Fword:  __ccgo_ts + 59374,
		Ftype1: int8('f'),
	},
	8525: {
		Fword:  __ccgo_ts + 59380,
		Ftype1: int8('E'),
	},
	8526: {
		Fword:  __ccgo_ts + 59387,
		Ftype1: int8('n'),
	},
	8527: {
		Fword:  __ccgo_ts + 59397,
		Ftype1: int8('T'),
	},
	8528: {
		Fword:  __ccgo_ts + 59415,
		Ftype1: int8('n'),
	},
	8529: {
		Fword:  __ccgo_ts + 59421,
		Ftype1: int8('k'),
	},
	8530: {
		Fword:  __ccgo_ts + 59432,
		Ftype1: int8('f'),
	},
	8531: {
		Fword:  __ccgo_ts + 59437,
		Ftype1: int8('t'),
	},
	8532: {
		Fword:  __ccgo_ts + 59445,
		Ftype1: int8('f'),
	},
	8533: {
		Fword:  __ccgo_ts + 59463,
		Ftype1: int8('f'),
	},
	8534: {
		Fword:  __ccgo_ts + 59473,
		Ftype1: int8('f'),
	},
	8535: {
		Fword:  __ccgo_ts + 59481,
		Ftype1: int8('f'),
	},
	8536: {
		Fword:  __ccgo_ts + 59488,
		Ftype1: int8('v'),
	},
	8537: {
		Fword:  __ccgo_ts + 59501,
		Ftype1: int8('v'),
	},
	8538: {
		Fword:  __ccgo_ts + 59516,
		Ftype1: int8('v'),
	},
	8539: {
		Fword:  __ccgo_ts + 59533,
		Ftype1: int8('v'),
	},
	8540: {
		Fword:  __ccgo_ts + 59555,
		Ftype1: int8('v'),
	},
	8541: {
		Fword:  __ccgo_ts + 59568,
		Ftype1: int8('v'),
	},
	8542: {
		Fword:  __ccgo_ts + 59583,
		Ftype1: int8('v'),
	},
	8543: {
		Fword:  __ccgo_ts + 59598,
		Ftype1: int8('v'),
	},
	8544: {
		Fword:  __ccgo_ts + 59611,
		Ftype1: int8('v'),
	},
	8545: {
		Fword:  __ccgo_ts + 59628,
		Ftype1: int8('f'),
	},
	8546: {
		Fword:  __ccgo_ts + 59640,
		Ftype1: int8('f'),
	},
	8547: {
		Fword:  __ccgo_ts + 59657,
		Ftype1: int8('v'),
	},
	8548: {
		Fword:  __ccgo_ts + 59670,
		Ftype1: int8('v'),
	},
	8549: {
		Fword:  __ccgo_ts + 59683,
		Ftype1: int8('f'),
	},
	8550: {
		Fword:  __ccgo_ts + 59697,
		Ftype1: int8('f'),
	},
	8551: {
		Fword:  __ccgo_ts + 59712,
		Ftype1: int8('f'),
	},
	8552: {
		Fword:  __ccgo_ts + 59728,
		Ftype1: int8('v'),
	},
	8553: {
		Fword:  __ccgo_ts + 59743,
		Ftype1: int8('f'),
	},
	8554: {
		Fword:  __ccgo_ts + 59759,
		Ftype1: int8('v'),
	},
	8555: {
		Fword:  __ccgo_ts + 59772,
		Ftype1: int8('v'),
	},
	8556: {
		Fword:  __ccgo_ts + 59790,
		Ftype1: int8('v'),
	},
	8557: {
		Fword:  __ccgo_ts + 59807,
		Ftype1: int8('v'),
	},
	8558: {
		Fword:  __ccgo_ts + 59820,
		Ftype1: int8('f'),
	},
	8559: {
		Fword:  __ccgo_ts + 59828,
		Ftype1: int8('k'),
	},
	8560: {
		Fword:  __ccgo_ts + 59835,
		Ftype1: int8('f'),
	},
	8561: {
		Fword:  __ccgo_ts + 59849,
		Ftype1: int8('f'),
	},
	8562: {
		Fword:  __ccgo_ts + 59857,
		Ftype1: int8('f'),
	},
	8563: {
		Fword:  __ccgo_ts + 59862,
		Ftype1: int8('n'),
	},
	8564: {
		Fword:  __ccgo_ts + 59871,
		Ftype1: int8('f'),
	},
	8565: {
		Fword:  __ccgo_ts + 59890,
		Ftype1: int8('k'),
	},
	8566: {
		Fword:  __ccgo_ts + 59900,
		Ftype1: int8('f'),
	},
	8567: {
		Fword:  __ccgo_ts + 59922,
		Ftype1: int8('f'),
	},
	8568: {
		Fword:  __ccgo_ts + 59933,
		Ftype1: int8('f'),
	},
	8569: {
		Fword:  __ccgo_ts + 59938,
		Ftype1: int8('f'),
	},
	8570: {
		Fword:  __ccgo_ts + 59946,
		Ftype1: int8('f'),
	},
	8571: {
		Fword:  __ccgo_ts + 59955,
		Ftype1: int8('f'),
	},
	8572: {
		Fword:  __ccgo_ts + 59969,
		Ftype1: int8('f'),
	},
	8573: {
		Fword:  __ccgo_ts + 59978,
		Ftype1: int8('f'),
	},
	8574: {
		Fword:  __ccgo_ts + 59987,
		Ftype1: int8('f'),
	},
	8575: {
		Fword:  __ccgo_ts + 59998,
		Ftype1: int8('f'),
	},
	8576: {
		Fword:  __ccgo_ts + 60017,
		Ftype1: int8('f'),
	},
	8577: {
		Fword:  __ccgo_ts + 60035,
		Ftype1: int8('f'),
	},
	8578: {
		Fword:  __ccgo_ts + 60059,
		Ftype1: int8('f'),
	},
	8579: {
		Fword:  __ccgo_ts + 60069,
		Ftype1: int8('f'),
	},
	8580: {
		Fword:  __ccgo_ts + 60078,
		Ftype1: int8('f'),
	},
	8581: {
		Fword:  __ccgo_ts + 60090,
		Ftype1: int8('f'),
	},
	8582: {
		Fword:  __ccgo_ts + 60100,
		Ftype1: int8('f'),
	},
	8583: {
		Fword:  __ccgo_ts + 60109,
		Ftype1: int8('f'),
	},
	8584: {
		Fword:  __ccgo_ts + 60120,
		Ftype1: int8('f'),
	},
	8585: {
		Fword:  __ccgo_ts + 60125,
		Ftype1: int8('f'),
	},
	8586: {
		Fword:  __ccgo_ts + 60129,
		Ftype1: int8('f'),
	},
	8587: {
		Fword:  __ccgo_ts + 60137,
		Ftype1: int8('f'),
	},
	8588: {
		Fword:  __ccgo_ts + 60148,
		Ftype1: int8('f'),
	},
	8589: {
		Fword:  __ccgo_ts + 60158,
		Ftype1: int8('f'),
	},
	8590: {
		Fword:  __ccgo_ts + 60168,
		Ftype1: int8('k'),
	},
	8591: {
		Fword:  __ccgo_ts + 60177,
		Ftype1: int8('k'),
	},
	8592: {
		Fword:  __ccgo_ts + 60193,
		Ftype1: int8('k'),
	},
	8593: {
		Fword:  __ccgo_ts + 60204,
		Ftype1: int8('k'),
	},
	8594: {
		Fword:  __ccgo_ts + 60215,
		Ftype1: int8('f'),
	},
	8595: {
		Fword:  __ccgo_ts + 60231,
		Ftype1: int8('f'),
	},
	8596: {
		Fword:  __ccgo_ts + 60257,
		Ftype1: int8('f'),
	},
	8597: {
		Fword:  __ccgo_ts + 60287,
		Ftype1: int8('f'),
	},
	8598: {
		Fword:  __ccgo_ts + 60293,
		Ftype1: int8('f'),
	},
	8599: {
		Fword:  __ccgo_ts + 60301,
		Ftype1: int8('f'),
	},
	8600: {
		Fword:  __ccgo_ts + 60308,
		Ftype1: int8('k'),
	},
	8601: {
		Fword:  __ccgo_ts + 60312,
		Ftype1: int8('t'),
	},
	8602: {
		Fword:  __ccgo_ts + 60320,
		Ftype1: int8('T'),
	},
	8603: {
		Fword:  __ccgo_ts + 60328,
		Ftype1: int8('f'),
	},
	8604: {
		Fword:  __ccgo_ts + 60335,
		Ftype1: int8('f'),
	},
	8605: {
		Fword:  __ccgo_ts + 60351,
		Ftype1: int8('f'),
	},
	8606: {
		Fword:  __ccgo_ts + 60365,
		Ftype1: int8('f'),
	},
	8607: {
		Fword:  __ccgo_ts + 60378,
		Ftype1: int8('f'),
	},
	8608: {
		Fword:  __ccgo_ts + 60399,
		Ftype1: int8('f'),
	},
	8609: {
		Fword:  __ccgo_ts + 60419,
		Ftype1: int8('k'),
	},
	8610: {
		Fword:  __ccgo_ts + 60427,
		Ftype1: int8('f'),
	},
	8611: {
		Fword:  __ccgo_ts + 60435,
		Ftype1: int8('k'),
	},
	8612: {
		Fword:  __ccgo_ts + 60441,
		Ftype1: int8('k'),
	},
	8613: {
		Fword:  __ccgo_ts + 60449,
		Ftype1: int8('T'),
	},
	8614: {
		Fword:  __ccgo_ts + 60456,
		Ftype1: int8('f'),
	},
	8615: {
		Fword:  __ccgo_ts + 60467,
		Ftype1: int8('k'),
	},
	8616: {
		Fword:  __ccgo_ts + 60472,
		Ftype1: int8('k'),
	},
	8617: {
		Fword:  __ccgo_ts + 60481,
		Ftype1: int8('f'),
	},
	8618: {
		Fword:  __ccgo_ts + 60493,
		Ftype1: int8('f'),
	},
	8619: {
		Fword:  __ccgo_ts + 60505,
		Ftype1: int8('k'),
	},
	8620: {
		Fword:  __ccgo_ts + 60519,
		Ftype1: int8('f'),
	},
	8621: {
		Fword:  __ccgo_ts + 60526,
		Ftype1: int8('f'),
	},
	8622: {
		Fword:  __ccgo_ts + 60537,
		Ftype1: int8('k'),
	},
	8623: {
		Fword:  __ccgo_ts + 60546,
		Ftype1: int8('k'),
	},
	8624: {
		Fword:  __ccgo_ts + 60558,
		Ftype1: int8('o'),
	},
	8625: {
		Fword:  __ccgo_ts + 60562,
		Ftype1: int8('f'),
	},
	8626: {
		Fword:  __ccgo_ts + 60568,
		Ftype1: int8('f'),
	},
	8627: {
		Fword:  __ccgo_ts + 60576,
		Ftype1: int8('f'),
	},
	8628: {
		Fword:  __ccgo_ts + 60581,
		Ftype1: int8('f'),
	},
	8629: {
		Fword:  __ccgo_ts + 60586,
		Ftype1: int8('n'),
	},
	8630: {
		Fword:  __ccgo_ts + 60589,
		Ftype1: int8('t'),
	},
	8631: {
		Fword:  __ccgo_ts + 60596,
		Ftype1: int8('t'),
	},
	8632: {
		Fword:  __ccgo_ts + 60613,
		Ftype1: int8('T'),
	},
	8633: {
		Fword:  __ccgo_ts + 60618,
		Ftype1: int8('f'),
	},
	8634: {
		Fword:  __ccgo_ts + 60623,
		Ftype1: int8('n'),
	},
	8635: {
		Fword:  __ccgo_ts + 60628,
		Ftype1: int8('k'),
	},
	8636: {
		Fword:  __ccgo_ts + 60633,
		Ftype1: int8('k'),
	},
	8637: {
		Fword:  __ccgo_ts + 60638,
		Ftype1: int8('k'),
	},
	8638: {
		Fword:  __ccgo_ts + 60645,
		Ftype1: int8('f'),
	},
	8639: {
		Fword:  __ccgo_ts + 60649,
		Ftype1: int8('k'),
	},
	8640: {
		Fword:  __ccgo_ts + 60658,
		Ftype1: int8('f'),
	},
	8641: {
		Fword:  __ccgo_ts + 60665,
		Ftype1: int8('f'),
	},
	8642: {
		Fword:  __ccgo_ts + 60673,
		Ftype1: int8('f'),
	},
	8643: {
		Fword:  __ccgo_ts + 60689,
		Ftype1: int8('f'),
	},
	8644: {
		Fword:  __ccgo_ts + 60703,
		Ftype1: int8('f'),
	},
	8645: {
		Fword:  __ccgo_ts + 60716,
		Ftype1: int8('f'),
	},
	8646: {
		Fword:  __ccgo_ts + 60736,
		Ftype1: int8('f'),
	},
	8647: {
		Fword:  __ccgo_ts + 60747,
		Ftype1: int8('f'),
	},
	8648: {
		Fword:  __ccgo_ts + 60757,
		Ftype1: int8('f'),
	},
	8649: {
		Fword:  __ccgo_ts + 60768,
		Ftype1: int8('f'),
	},
	8650: {
		Fword:  __ccgo_ts + 60776,
		Ftype1: int8('o'),
	},
	8651: {
		Fword:  __ccgo_ts + 60780,
		Ftype1: int8('k'),
	},
	8652: {
		Fword:  __ccgo_ts + 60788,
		Ftype1: int8('f'),
	},
	8653: {
		Fword:  __ccgo_ts + 60798,
		Ftype1: int8('U'),
	},
	8654: {
		Fword:  __ccgo_ts + 60805,
		Ftype1: int8('T'),
	},
	8655: {
		Fword:  __ccgo_ts + 60810,
		Ftype1: int8('T'),
	},
	8656: {
		Fword:  __ccgo_ts + 60818,
		Ftype1: int8('E'),
	},
	8657: {
		Fword:  __ccgo_ts + 60829,
		Ftype1: int8('E'),
	},
	8658: {
		Fword:  __ccgo_ts + 60846,
		Ftype1: int8('f'),
	},
	8659: {
		Fword:  __ccgo_ts + 60853,
		Ftype1: int8('k'),
	},
	8660: {
		Fword:  __ccgo_ts + 60858,
		Ftype1: int8('f'),
	},
	8661: {
		Fword:  __ccgo_ts + 60862,
		Ftype1: int8('k'),
	},
	8662: {
		Fword:  __ccgo_ts + 60870,
		Ftype1: int8('f'),
	},
	8663: {
		Fword:  __ccgo_ts + 60881,
		Ftype1: int8('f'),
	},
	8664: {
		Fword:  __ccgo_ts + 60889,
		Ftype1: int8('f'),
	},
	8665: {
		Fword:  __ccgo_ts + 60902,
		Ftype1: int8('f'),
	},
	8666: {
		Fword:  __ccgo_ts + 60916,
		Ftype1: int8('1'),
	},
	8667: {
		Fword:  __ccgo_ts + 60922,
		Ftype1: int8('k'),
	},
	8668: {
		Fword:  __ccgo_ts + 60928,
		Ftype1: int8('f'),
	},
	8669: {
		Fword:  __ccgo_ts + 60934,
		Ftype1: int8('f'),
	},
	8670: {
		Fword:  __ccgo_ts + 60947,
		Ftype1: int8('f'),
	},
	8671: {
		Fword:  __ccgo_ts + 60965,
		Ftype1: int8('f'),
	},
	8672: {
		Fword:  __ccgo_ts + 60978,
		Ftype1: int8('f'),
	},
	8673: {
		Fword:  __ccgo_ts + 60993,
		Ftype1: int8('f'),
	},
	8674: {
		Fword:  __ccgo_ts + 61001,
		Ftype1: int8('f'),
	},
	8675: {
		Fword:  __ccgo_ts + 61014,
		Ftype1: int8('f'),
	},
	8676: {
		Fword:  __ccgo_ts + 61025,
		Ftype1: int8('f'),
	},
	8677: {
		Fword:  __ccgo_ts + 61036,
		Ftype1: int8('f'),
	},
	8678: {
		Fword:  __ccgo_ts + 61044,
		Ftype1: int8('f'),
	},
	8679: {
		Fword:  __ccgo_ts + 61054,
		Ftype1: int8('f'),
	},
	8680: {
		Fword:  __ccgo_ts + 61064,
		Ftype1: int8('f'),
	},
	8681: {
		Fword:  __ccgo_ts + 61076,
		Ftype1: int8('f'),
	},
	8682: {
		Fword:  __ccgo_ts + 61088,
		Ftype1: int8('t'),
	},
	8683: {
		Fword:  __ccgo_ts + 61094,
		Ftype1: int8('t'),
	},
	8684: {
		Fword:  __ccgo_ts + 61101,
		Ftype1: int8('t'),
	},
	8685: {
		Fword:  __ccgo_ts + 61108,
		Ftype1: int8('f'),
	},
	8686: {
		Fword:  __ccgo_ts + 61114,
		Ftype1: int8('f'),
	},
	8687: {
		Fword:  __ccgo_ts + 61134,
		Ftype1: int8('n'),
	},
	8688: {
		Fword:  __ccgo_ts + 61138,
		Ftype1: int8('k'),
	},
	8689: {
		Fword:  __ccgo_ts + 61149,
		Ftype1: int8('k'),
	},
	8690: {
		Fword:  __ccgo_ts + 61167,
		Ftype1: int8('k'),
	},
	8691: {
		Fword:  __ccgo_ts + 61181,
		Ftype1: int8('k'),
	},
	8692: {
		Fword:  __ccgo_ts + 61197,
		Ftype1: int8('k'),
	},
	8693: {
		Fword:  __ccgo_ts + 61220,
		Ftype1: int8('k'),
	},
	8694: {
		Fword:  __ccgo_ts + 61236,
		Ftype1: int8('k'),
	},
	8695: {
		Fword:  __ccgo_ts + 61242,
		Ftype1: int8('k'),
	},
	8696: {
		Fword:  __ccgo_ts + 61250,
		Ftype1: int8('f'),
	},
	8697: {
		Fword:  __ccgo_ts + 61257,
		Ftype1: int8('f'),
	},
	8698: {
		Fword:  __ccgo_ts + 61268,
		Ftype1: int8('k'),
	},
	8699: {
		Fword:  __ccgo_ts + 61273,
		Ftype1: int8('f'),
	},
	8700: {
		Fword:  __ccgo_ts + 61285,
		Ftype1: int8('f'),
	},
	8701: {
		Fword:  __ccgo_ts + 61295,
		Ftype1: int8('f'),
	},
	8702: {
		Fword:  __ccgo_ts + 61309,
		Ftype1: int8('k'),
	},
	8703: {
		Fword:  __ccgo_ts + 61319,
		Ftype1: int8('k'),
	},
	8704: {
		Fword:  __ccgo_ts + 61330,
		Ftype1: int8('k'),
	},
	8705: {
		Fword:  __ccgo_ts + 61346,
		Ftype1: int8('k'),
	},
	8706: {
		Fword:  __ccgo_ts + 61355,
		Ftype1: int8('f'),
	},
	8707: {
		Fword:  __ccgo_ts + 61379,
		Ftype1: int8('f'),
	},
	8708: {
		Fword:  __ccgo_ts + 61403,
		Ftype1: int8('k'),
	},
	8709: {
		Fword:  __ccgo_ts + 61412,
		Ftype1: int8('f'),
	},
	8710: {
		Fword:  __ccgo_ts + 61428,
		Ftype1: int8('f'),
	},
	8711: {
		Fword:  __ccgo_ts + 61448,
		Ftype1: int8('f'),
	},
	8712: {
		Fword:  __ccgo_ts + 61456,
		Ftype1: int8('f'),
	},
	8713: {
		Fword:  __ccgo_ts + 61464,
		Ftype1: int8('f'),
	},
	8714: {
		Fword:  __ccgo_ts + 61475,
		Ftype1: int8('f'),
	},
	8715: {
		Fword:  __ccgo_ts + 61483,
		Ftype1: int8('f'),
	},
	8716: {
		Fword:  __ccgo_ts + 61492,
		Ftype1: int8('f'),
	},
	8717: {
		Fword:  __ccgo_ts + 61503,
		Ftype1: int8('f'),
	},
	8718: {
		Fword:  __ccgo_ts + 61512,
		Ftype1: int8('T'),
	},
	8719: {
		Fword:  __ccgo_ts + 61515,
		Ftype1: int8('T'),
	},
	8720: {
		Fword:  __ccgo_ts + 61520,
		Ftype1: int8('k'),
	},
	8721: {
		Fword:  __ccgo_ts + 61526,
		Ftype1: int8('f'),
	},
	8722: {
		Fword:  __ccgo_ts + 61535,
		Ftype1: int8('n'),
	},
	8723: {
		Fword:  __ccgo_ts + 61541,
		Ftype1: int8('B'),
	},
	8724: {
		Fword:  __ccgo_ts + 61550,
		Ftype1: int8('f'),
	},
	8725: {
		Fword:  __ccgo_ts + 61559,
		Ftype1: int8('f'),
	},
	8726: {
		Fword:  __ccgo_ts + 61571,
		Ftype1: int8('f'),
	},
	8727: {
		Fword:  __ccgo_ts + 61584,
		Ftype1: int8('T'),
	},
	8728: {
		Fword:  __ccgo_ts + 61592,
		Ftype1: int8('f'),
	},
	8729: {
		Fword:  __ccgo_ts + 61602,
		Ftype1: int8('f'),
	},
	8730: {
		Fword:  __ccgo_ts + 61620,
		Ftype1: int8('B'),
	},
	8731: {
		Fword:  __ccgo_ts + 61627,
		Ftype1: int8('f'),
	},
	8732: {
		Fword:  __ccgo_ts + 61631,
		Ftype1: int8('k'),
	},
	8733: {
		Fword:  __ccgo_ts + 61645,
		Ftype1: int8('f'),
	},
	8734: {
		Fword:  __ccgo_ts + 61655,
		Ftype1: int8('f'),
	},
	8735: {
		Fword:  __ccgo_ts + 61660,
		Ftype1: int8('k'),
	},
	8736: {
		Fword:  __ccgo_ts + 61677,
		Ftype1: int8('k'),
	},
	8737: {
		Fword:  __ccgo_ts + 61689,
		Ftype1: int8('k'),
	},
	8738: {
		Fword:  __ccgo_ts + 61701,
		Ftype1: int8('f'),
	},
	8739: {
		Fword:  __ccgo_ts + 61710,
		Ftype1: int8('f'),
	},
	8740: {
		Fword:  __ccgo_ts + 61724,
		Ftype1: int8('f'),
	},
	8741: {
		Fword:  __ccgo_ts + 61735,
		Ftype1: int8('f'),
	},
	8742: {
		Fword:  __ccgo_ts + 61746,
		Ftype1: int8('f'),
	},
	8743: {
		Fword:  __ccgo_ts + 61749,
		Ftype1: int8('f'),
	},
	8744: {
		Fword:  __ccgo_ts + 61759,
		Ftype1: int8('f'),
	},
	8745: {
		Fword:  __ccgo_ts + 61766,
		Ftype1: int8('f'),
	},
	8746: {
		Fword:  __ccgo_ts + 61780,
		Ftype1: int8('f'),
	},
	8747: {
		Fword:  __ccgo_ts + 61784,
		Ftype1: int8('f'),
	},
	8748: {
		Fword:  __ccgo_ts + 61791,
		Ftype1: int8('k'),
	},
	8749: {
		Fword:  __ccgo_ts + 61798,
		Ftype1: int8('f'),
	},
	8750: {
		Fword:  __ccgo_ts + 61802,
		Ftype1: int8('k'),
	},
	8751: {
		Fword:  __ccgo_ts + 61805,
		Ftype1: int8('n'),
	},
	8752: {
		Fword:  __ccgo_ts + 61816,
		Ftype1: int8('k'),
	},
	8753: {
		Fword:  __ccgo_ts + 61832,
		Ftype1: int8('k'),
	},
	8754: {
		Fword:  __ccgo_ts + 61838,
		Ftype1: int8('f'),
	},
	8755: {
		Fword:  __ccgo_ts + 61856,
		Ftype1: int8('f'),
	},
	8756: {
		Fword:  __ccgo_ts + 61870,
		Ftype1: int8('f'),
	},
	8757: {
		Fword:  __ccgo_ts + 61880,
		Ftype1: int8('f'),
	},
	8758: {
		Fword:  __ccgo_ts + 61890,
		Ftype1: int8('f'),
	},
	8759: {
		Fword:  __ccgo_ts + 61900,
		Ftype1: int8('k'),
	},
	8760: {
		Fword:  __ccgo_ts + 61907,
		Ftype1: int8('f'),
	},
	8761: {
		Fword:  __ccgo_ts + 61915,
		Ftype1: int8('k'),
	},
	8762: {
		Fword:  __ccgo_ts + 61921,
		Ftype1: int8('k'),
	},
	8763: {
		Fword:  __ccgo_ts + 61932,
		Ftype1: int8('k'),
	},
	8764: {
		Fword:  __ccgo_ts + 61938,
		Ftype1: int8('k'),
	},
	8765: {
		Fword:  __ccgo_ts + 61950,
		Ftype1: int8('E'),
	},
	8766: {
		Fword:  __ccgo_ts + 61957,
		Ftype1: int8('E'),
	},
	8767: {
		Fword:  __ccgo_ts + 61972,
		Ftype1: int8('T'),
	},
	8768: {
		Fword:  __ccgo_ts + 61992,
		Ftype1: int8('E'),
	},
	8769: {
		Fword:  __ccgo_ts + 62013,
		Ftype1: int8('T'),
	},
	8770: {
		Fword:  __ccgo_ts + 62039,
		Ftype1: int8('E'),
	},
	8771: {
		Fword:  __ccgo_ts + 62053,
		Ftype1: int8('T'),
	},
	8772: {
		Fword:  __ccgo_ts + 62072,
		Ftype1: int8('T'),
	},
	8773: {
		Fword:  __ccgo_ts + 62084,
		Ftype1: int8('E'),
	},
	8774: {
		Fword:  __ccgo_ts + 62104,
		Ftype1: int8('T'),
	},
	8775: {
		Fword:  __ccgo_ts + 62129,
		Ftype1: int8('f'),
	},
	8776: {
		Fword:  __ccgo_ts + 62135,
		Ftype1: int8('f'),
	},
	8777: {
		Fword:  __ccgo_ts + 62144,
		Ftype1: int8('t'),
	},
	8778: {
		Fword:  __ccgo_ts + 62148,
		Ftype1: int8('t'),
	},
	8779: {
		Fword:  __ccgo_ts + 62153,
		Ftype1: int8('t'),
	},
	8780: {
		Fword:  __ccgo_ts + 62158,
		Ftype1: int8('t'),
	},
	8781: {
		Fword:  __ccgo_ts + 62163,
		Ftype1: int8('t'),
	},
	8782: {
		Fword:  __ccgo_ts + 62168,
		Ftype1: int8('t'),
	},
	8783: {
		Fword:  __ccgo_ts + 62173,
		Ftype1: int8('t'),
	},
	8784: {
		Fword:  __ccgo_ts + 62181,
		Ftype1: int8('U'),
	},
	8785: {
		Fword:  __ccgo_ts + 62191,
		Ftype1: int8('U'),
	},
	8786: {
		Fword:  __ccgo_ts + 62205,
		Ftype1: int8('k'),
	},
	8787: {
		Fword:  __ccgo_ts + 62214,
		Ftype1: int8('k'),
	},
	8788: {
		Fword:  __ccgo_ts + 62219,
		Ftype1: int8('k'),
	},
	8789: {
		Fword:  __ccgo_ts + 62233,
		Ftype1: int8('k'),
	},
	8790: {
		Fword:  __ccgo_ts + 62246,
		Ftype1: int8('o'),
	},
	8791: {
		Fword:  __ccgo_ts + 62249,
		Ftype1: int8('n'),
	},
	8792: {
		Fword:  __ccgo_ts + 62261,
		Ftype1: int8('o'),
	},
	8793: {
		Fword:  __ccgo_ts + 62278,
		Ftype1: int8('o'),
	},
	8794: {
		Fword:  __ccgo_ts + 62285,
		Ftype1: int8('n'),
	},
	8795: {
		Fword:  __ccgo_ts + 62301,
		Ftype1: int8('o'),
	},
	8796: {
		Fword:  __ccgo_ts + 62322,
		Ftype1: int8('f'),
	},
	8797: {
		Fword:  __ccgo_ts + 62329,
		Ftype1: int8('f'),
	},
	8798: {
		Fword:  __ccgo_ts + 62337,
		Ftype1: int8('f'),
	},
	8799: {
		Fword:  __ccgo_ts + 62346,
		Ftype1: int8('f'),
	},
	8800: {
		Fword:  __ccgo_ts + 62353,
		Ftype1: int8('f'),
	},
	8801: {
		Fword:  __ccgo_ts + 62363,
		Ftype1: int8('f'),
	},
	8802: {
		Fword:  __ccgo_ts + 62376,
		Ftype1: int8('f'),
	},
	8803: {
		Fword:  __ccgo_ts + 62386,
		Ftype1: int8('f'),
	},
	8804: {
		Fword:  __ccgo_ts + 62402,
		Ftype1: int8('f'),
	},
	8805: {
		Fword:  __ccgo_ts + 62416,
		Ftype1: int8('f'),
	},
	8806: {
		Fword:  __ccgo_ts + 62433,
		Ftype1: int8('f'),
	},
	8807: {
		Fword:  __ccgo_ts + 62446,
		Ftype1: int8('k'),
	},
	8808: {
		Fword:  __ccgo_ts + 62454,
		Ftype1: int8('k'),
	},
	8809: {
		Fword:  __ccgo_ts + 62459,
		Ftype1: int8('f'),
	},
	8810: {
		Fword:  __ccgo_ts + 62469,
		Ftype1: int8('f'),
	},
	8811: {
		Fword:  __ccgo_ts + 62479,
		Ftype1: int8('f'),
	},
	8812: {
		Fword:  __ccgo_ts + 62492,
		Ftype1: int8('f'),
	},
	8813: {
		Fword:  __ccgo_ts + 62506,
		Ftype1: int8('f'),
	},
	8814: {
		Fword:  __ccgo_ts + 62523,
		Ftype1: int8('k'),
	},
	8815: {
		Fword:  __ccgo_ts + 62528,
		Ftype1: int8('f'),
	},
	8816: {
		Fword:  __ccgo_ts + 62537,
		Ftype1: int8('f'),
	},
	8817: {
		Fword:  __ccgo_ts + 62544,
		Ftype1: int8('k'),
	},
	8818: {
		Fword:  __ccgo_ts + 62549,
		Ftype1: int8('f'),
	},
	8819: {
		Fword:  __ccgo_ts + 62553,
		Ftype1: int8('f'),
	},
	8820: {
		Fword:  __ccgo_ts + 62561,
		Ftype1: int8('f'),
	},
	8821: {
		Fword:  __ccgo_ts + 62576,
		Ftype1: int8('f'),
	},
	8822: {
		Fword:  __ccgo_ts + 62594,
		Ftype1: int8('f'),
	},
	8823: {
		Fword:  __ccgo_ts + 62605,
		Ftype1: int8('f'),
	},
	8824: {
		Fword:  __ccgo_ts + 62611,
		Ftype1: int8('f'),
	},
	8825: {
		Fword:  __ccgo_ts + 62616,
		Ftype1: int8('k'),
	},
	8826: {
		Fword:  __ccgo_ts + 62624,
		Ftype1: int8('f'),
	},
	8827: {
		Fword:  __ccgo_ts + 62630,
		Ftype1: int8('k'),
	},
	8828: {
		Fword:  __ccgo_ts + 62636,
		Ftype1: int8('f'),
	},
	8829: {
		Fword:  __ccgo_ts + 62641,
		Ftype1: int8('k'),
	},
	8830: {
		Fword:  __ccgo_ts + 62651,
		Ftype1: int8('k'),
	},
	8831: {
		Fword:  __ccgo_ts + 62662,
		Ftype1: int8('k'),
	},
	8832: {
		Fword:  __ccgo_ts + 62678,
		Ftype1: int8('f'),
	},
	8833: {
		Fword:  __ccgo_ts + 62685,
		Ftype1: int8('o'),
	},
	8834: {
		Fword:  __ccgo_ts + 62690,
		Ftype1: int8('B'),
	},
	8835: {
		Fword:  __ccgo_ts + 62696,
		Ftype1: int8('k'),
	},
	8836: {
		Fword:  __ccgo_ts + 62703,
		Ftype1: int8('k'),
	},
	8837: {
		Fword:  __ccgo_ts + 62709,
		Ftype1: int8('f'),
	},
	8838: {
		Fword:  __ccgo_ts + 62712,
		Ftype1: int8('k'),
	},
	8839: {
		Fword:  __ccgo_ts + 62717,
		Ftype1: int8('T'),
	},
	8840: {
		Fword:  __ccgo_ts + 62727,
		Ftype1: int8('T'),
	},
	8841: {
		Fword:  __ccgo_ts + 62736,
		Ftype1: int8('f'),
	},
	8842: {
		Fword:  __ccgo_ts + 62751,
		Ftype1: int8('f'),
	},
	8843: {
		Fword:  __ccgo_ts + 62761,
		Ftype1: int8('v'),
	},
	8844: {
		Fword:  __ccgo_ts + 62771,
		Ftype1: int8('v'),
	},
	8845: {
		Fword:  __ccgo_ts + 62786,
		Ftype1: int8('f'),
	},
	8846: {
		Fword:  __ccgo_ts + 62793,
		Ftype1: int8('n'),
	},
	8847: {
		Fword:  __ccgo_ts + 62798,
		Ftype1: int8('n'),
	},
	8848: {
		Fword:  __ccgo_ts + 62806,
		Ftype1: int8('n'),
	},
	8849: {
		Fword:  __ccgo_ts + 62820,
		Ftype1: int8('k'),
	},
	8850: {
		Fword:  __ccgo_ts + 62839,
		Ftype1: int8('k'),
	},
	8851: {
		Fword:  __ccgo_ts + 62850,
		Ftype1: int8('k'),
	},
	8852: {
		Fword:  __ccgo_ts + 62862,
		Ftype1: int8('f'),
	},
	8853: {
		Fword:  __ccgo_ts + 62866,
		Ftype1: int8('f'),
	},
	8854: {
		Fword:  __ccgo_ts + 62872,
		Ftype1: int8('f'),
	},
	8855: {
		Fword:  __ccgo_ts + 62877,
		Ftype1: int8('k'),
	},
	8856: {
		Fword:  __ccgo_ts + 62886,
		Ftype1: int8('k'),
	},
	8857: {
		Fword:  __ccgo_ts + 62895,
		Ftype1: int8('k'),
	},
	8858: {
		Fword:  __ccgo_ts + 62900,
		Ftype1: int8('f'),
	},
	8859: {
		Fword:  __ccgo_ts + 62906,
		Ftype1: int8('f'),
	},
	8860: {
		Fword:  __ccgo_ts + 62916,
		Ftype1: int8('f'),
	},
	8861: {
		Fword:  __ccgo_ts + 62926,
		Ftype1: int8('k'),
	},
	8862: {
		Fword:  __ccgo_ts + 62939,
		Ftype1: int8('f'),
	},
	8863: {
		Fword:  __ccgo_ts + 62944,
		Ftype1: int8('f'),
	},
	8864: {
		Fword:  __ccgo_ts + 62950,
		Ftype1: int8('f'),
	},
	8865: {
		Fword:  __ccgo_ts + 62959,
		Ftype1: int8('f'),
	},
	8866: {
		Fword:  __ccgo_ts + 62968,
		Ftype1: int8('f'),
	},
	8867: {
		Fword:  __ccgo_ts + 62976,
		Ftype1: int8('k'),
	},
	8868: {
		Fword:  __ccgo_ts + 62988,
		Ftype1: int8('f'),
	},
	8869: {
		Fword:  __ccgo_ts + 63004,
		Ftype1: int8('k'),
	},
	8870: {
		Fword:  __ccgo_ts + 63034,
		Ftype1: int8('k'),
	},
	8871: {
		Fword:  __ccgo_ts + 63040,
		Ftype1: int8('f'),
	},
	8872: {
		Fword:  __ccgo_ts + 63044,
		Ftype1: int8('k'),
	},
	8873: {
		Fword:  __ccgo_ts + 63053,
		Ftype1: int8('f'),
	},
	8874: {
		Fword:  __ccgo_ts + 63057,
		Ftype1: int8('k'),
	},
	8875: {
		Fword:  __ccgo_ts + 63068,
		Ftype1: int8('k'),
	},
	8876: {
		Fword:  __ccgo_ts + 63078,
		Ftype1: int8('k'),
	},
	8877: {
		Fword:  __ccgo_ts + 63089,
		Ftype1: int8('k'),
	},
	8878: {
		Fword:  __ccgo_ts + 63095,
		Ftype1: int8('f'),
	},
	8879: {
		Fword:  __ccgo_ts + 63107,
		Ftype1: int8('f'),
	},
	8880: {
		Fword:  __ccgo_ts + 63111,
		Ftype1: int8('k'),
	},
	8881: {
		Fword:  __ccgo_ts + 63121,
		Ftype1: int8('f'),
	},
	8882: {
		Fword:  __ccgo_ts + 63125,
		Ftype1: int8('f'),
	},
	8883: {
		Fword:  __ccgo_ts + 63132,
		Ftype1: int8('k'),
	},
	8884: {
		Fword:  __ccgo_ts + 63151,
		Ftype1: int8('k'),
	},
	8885: {
		Fword:  __ccgo_ts + 63165,
		Ftype1: int8('f'),
	},
	8886: {
		Fword:  __ccgo_ts + 63171,
		Ftype1: int8('o'),
	},
	8887: {
		Fword:  __ccgo_ts + 63175,
		Ftype1: int8('n'),
	},
	8888: {
		Fword:  __ccgo_ts + 63180,
		Ftype1: int8('k'),
	},
	8889: {
		Fword:  __ccgo_ts + 63189,
		Ftype1: int8('t'),
	},
	8890: {
		Fword:  __ccgo_ts + 63195,
		Ftype1: int8('f'),
	},
	8891: {
		Fword:  __ccgo_ts + 63201,
		Ftype1: int8('f'),
	},
	8892: {
		Fword:  __ccgo_ts + 63211,
		Ftype1: int8('f'),
	},
	8893: {
		Fword:  __ccgo_ts + 63222,
		Ftype1: int8('n'),
	},
	8894: {
		Fword:  __ccgo_ts + 63230,
		Ftype1: int8('k'),
	},
	8895: {
		Fword:  __ccgo_ts + 63243,
		Ftype1: int8('k'),
	},
	8896: {
		Fword:  __ccgo_ts + 63267,
		Ftype1: int8('k'),
	},
	8897: {
		Fword:  __ccgo_ts + 63281,
		Ftype1: int8('k'),
	},
	8898: {
		Fword:  __ccgo_ts + 63294,
		Ftype1: int8('k'),
	},
	8899: {
		Fword:  __ccgo_ts + 63307,
		Ftype1: int8('k'),
	},
	8900: {
		Fword:  __ccgo_ts + 63326,
		Ftype1: int8('k'),
	},
	8901: {
		Fword:  __ccgo_ts + 63350,
		Ftype1: int8('k'),
	},
	8902: {
		Fword:  __ccgo_ts + 63364,
		Ftype1: int8('k'),
	},
	8903: {
		Fword:  __ccgo_ts + 63378,
		Ftype1: int8('k'),
	},
	8904: {
		Fword:  __ccgo_ts + 63403,
		Ftype1: int8('f'),
	},
	8905: {
		Fword:  __ccgo_ts + 63411,
		Ftype1: int8('n'),
	},
	8906: {
		Fword:  __ccgo_ts + 63422,
		Ftype1: int8('k'),
	},
	8907: {
		Fword:  __ccgo_ts + 63437,
		Ftype1: int8('f'),
	},
	8908: {
		Fword:  __ccgo_ts + 63445,
		Ftype1: int8('o'),
	},
	8909: {
		Fword:  __ccgo_ts + 63449,
		Ftype1: int8('o'),
	},
	8910: {
		Fword:  __ccgo_ts + 63461,
		Ftype1: int8('k'),
	},
	8911: {
		Fword:  __ccgo_ts + 63468,
		Ftype1: int8('o'),
	},
	8912: {
		Fword:  __ccgo_ts + 63477,
		Ftype1: int8('o'),
	},
	8913: {
		Fword:  __ccgo_ts + 63488,
		Ftype1: int8('o'),
	},
	8914: {
		Fword:  __ccgo_ts + 63498,
		Ftype1: int8('o'),
	},
	8915: {
		Fword:  __ccgo_ts + 63510,
		Ftype1: int8('o'),
	},
	8916: {
		Fword:  __ccgo_ts + 63525,
		Ftype1: int8('k'),
	},
	8917: {
		Fword:  __ccgo_ts + 63533,
		Ftype1: int8('f'),
	},
	8918: {
		Fword:  __ccgo_ts + 63537,
		Ftype1: int8('k'),
	},
	8919: {
		Fword:  __ccgo_ts + 63544,
		Ftype1: int8('k'),
	},
	8920: {
		Fword:  __ccgo_ts + 63563,
		Ftype1: int8('f'),
	},
	8921: {
		Fword:  __ccgo_ts + 63573,
		Ftype1: int8('f'),
	},
	8922: {
		Fword:  __ccgo_ts + 63579,
		Ftype1: int8('v'),
	},
	8923: {
		Fword:  __ccgo_ts + 63584,
		Ftype1: int8('f'),
	},
	8924: {
		Fword:  __ccgo_ts + 63591,
		Ftype1: int8('t'),
	},
	8925: {
		Fword:  __ccgo_ts + 63599,
		Ftype1: int8('f'),
	},
	8926: {
		Fword:  __ccgo_ts + 63602,
		Ftype1: int8('f'),
	},
	8927: {
		Fword:  __ccgo_ts + 63617,
		Ftype1: int8('f'),
	},
	8928: {
		Fword:  __ccgo_ts + 63634,
		Ftype1: int8('f'),
	},
	8929: {
		Fword:  __ccgo_ts + 63652,
		Ftype1: int8('f'),
	},
	8930: {
		Fword:  __ccgo_ts + 63662,
		Ftype1: int8('f'),
	},
	8931: {
		Fword:  __ccgo_ts + 63674,
		Ftype1: int8('f'),
	},
	8932: {
		Fword:  __ccgo_ts + 63693,
		Ftype1: int8('f'),
	},
	8933: {
		Fword:  __ccgo_ts + 63697,
		Ftype1: int8('f'),
	},
	8934: {
		Fword:  __ccgo_ts + 63710,
		Ftype1: int8('k'),
	},
	8935: {
		Fword:  __ccgo_ts + 63717,
		Ftype1: int8('t'),
	},
	8936: {
		Fword:  __ccgo_ts + 63721,
		Ftype1: int8('f'),
	},
	8937: {
		Fword:  __ccgo_ts + 63734,
		Ftype1: int8('k'),
	},
	8938: {
		Fword:  __ccgo_ts + 63743,
		Ftype1: int8('k'),
	},
	8939: {
		Fword:  __ccgo_ts + 63748,
		Ftype1: int8('f'),
	},
	8940: {
		Fword:  __ccgo_ts + 63763,
		Ftype1: int8('f'),
	},
	8941: {
		Fword:  __ccgo_ts + 63773,
		Ftype1: int8('f'),
	},
	8942: {
		Fword:  __ccgo_ts + 63784,
		Ftype1: int8('f'),
	},
	8943: {
		Fword:  __ccgo_ts + 63792,
		Ftype1: int8('k'),
	},
	8944: {
		Fword:  __ccgo_ts + 63801,
		Ftype1: int8('k'),
	},
	8945: {
		Fword:  __ccgo_ts + 63808,
		Ftype1: int8('k'),
	},
	8946: {
		Fword:  __ccgo_ts + 63819,
		Ftype1: int8('&'),
	},
	8947: {
		Fword:  __ccgo_ts + 63822,
		Ftype1: int8('f'),
	},
	8948: {
		Fword:  __ccgo_ts + 63826,
		Ftype1: int8('n'),
	},
	8949: {
		Fword:  __ccgo_ts + 63832,
		Ftype1: int8('B'),
	},
	8950: {
		Fword:  __ccgo_ts + 63841,
		Ftype1: int8('f'),
	},
	8951: {
		Fword:  __ccgo_ts + 63858,
		Ftype1: int8('f'),
	},
	8952: {
		Fword:  __ccgo_ts + 63873,
		Ftype1: int8('n'),
	},
	8953: {
		Fword:  __ccgo_ts + 63877,
		Ftype1: int8('n'),
	},
	8954: {
		Fword:  __ccgo_ts + 63883,
		Ftype1: int8('k'),
	},
	8955: {
		Fword:  __ccgo_ts + 63891,
		Ftype1: int8('f'),
	},
	8956: {
		Fword:  __ccgo_ts + 63900,
		Ftype1: int8('f'),
	},
	8957: {
		Fword:  __ccgo_ts + 63908,
		Ftype1: int8('k'),
	},
	8958: {
		Fword:  __ccgo_ts + 63914,
		Ftype1: int8('B'),
	},
	8959: {
		Fword:  __ccgo_ts + 63923,
		Ftype1: int8('f'),
	},
	8960: {
		Fword:  __ccgo_ts + 63933,
		Ftype1: int8('k'),
	},
	8961: {
		Fword:  __ccgo_ts + 63943,
		Ftype1: int8('B'),
	},
	8962: {
		Fword:  __ccgo_ts + 63956,
		Ftype1: int8('n'),
	},
	8963: {
		Fword:  __ccgo_ts + 63965,
		Ftype1: int8('f'),
	},
	8964: {
		Fword:  __ccgo_ts + 63975,
		Ftype1: int8('f'),
	},
	8965: {
		Fword:  __ccgo_ts + 63984,
		Ftype1: int8('f'),
	},
	8966: {
		Fword:  __ccgo_ts + 64001,
		Ftype1: int8('f'),
	},
	8967: {
		Fword:  __ccgo_ts + 64017,
		Ftype1: int8('f'),
	},
	8968: {
		Fword:  __ccgo_ts + 64033,
		Ftype1: int8('f'),
	},
	8969: {
		Fword:  __ccgo_ts + 64046,
		Ftype1: int8('f'),
	},
	8970: {
		Fword:  __ccgo_ts + 64057,
		Ftype1: int8('f'),
	},
	8971: {
		Fword:  __ccgo_ts + 64069,
		Ftype1: int8('f'),
	},
	8972: {
		Fword:  __ccgo_ts + 64081,
		Ftype1: int8('f'),
	},
	8973: {
		Fword:  __ccgo_ts + 64098,
		Ftype1: int8('f'),
	},
	8974: {
		Fword:  __ccgo_ts + 64113,
		Ftype1: int8('f'),
	},
	8975: {
		Fword:  __ccgo_ts + 64131,
		Ftype1: int8('f'),
	},
	8976: {
		Fword:  __ccgo_ts + 64150,
		Ftype1: int8('f'),
	},
	8977: {
		Fword:  __ccgo_ts + 64168,
		Ftype1: int8('f'),
	},
	8978: {
		Fword:  __ccgo_ts + 64192,
		Ftype1: int8('f'),
	},
	8979: {
		Fword:  __ccgo_ts + 64204,
		Ftype1: int8('f'),
	},
	8980: {
		Fword:  __ccgo_ts + 64222,
		Ftype1: int8('f'),
	},
	8981: {
		Fword:  __ccgo_ts + 64246,
		Ftype1: int8('f'),
	},
	8982: {
		Fword:  __ccgo_ts + 64268,
		Ftype1: int8('f'),
	},
	8983: {
		Fword:  __ccgo_ts + 64278,
		Ftype1: int8('f'),
	},
	8984: {
		Fword:  __ccgo_ts + 64296,
		Ftype1: int8('f'),
	},
	8985: {
		Fword:  __ccgo_ts + 64321,
		Ftype1: int8('f'),
	},
	8986: {
		Fword:  __ccgo_ts + 64341,
		Ftype1: int8('f'),
	},
	8987: {
		Fword:  __ccgo_ts + 64354,
		Ftype1: int8('f'),
	},
	8988: {
		Fword:  __ccgo_ts + 64369,
		Ftype1: int8('f'),
	},
	8989: {
		Fword:  __ccgo_ts + 64387,
		Ftype1: int8('f'),
	},
	8990: {
		Fword:  __ccgo_ts + 64396,
		Ftype1: int8('f'),
	},
	8991: {
		Fword:  __ccgo_ts + 64412,
		Ftype1: int8('f'),
	},
	8992: {
		Fword:  __ccgo_ts + 64425,
		Ftype1: int8('f'),
	},
	8993: {
		Fword:  __ccgo_ts + 64440,
		Ftype1: int8('f'),
	},
	8994: {
		Fword:  __ccgo_ts + 64455,
		Ftype1: int8('f'),
	},
	8995: {
		Fword:  __ccgo_ts + 64476,
		Ftype1: int8('f'),
	},
	8996: {
		Fword:  __ccgo_ts + 64493,
		Ftype1: int8('f'),
	},
	8997: {
		Fword:  __ccgo_ts + 64496,
		Ftype1: int8('f'),
	},
	8998: {
		Fword:  __ccgo_ts + 64505,
		Ftype1: int8('f'),
	},
	8999: {
		Fword:  __ccgo_ts + 64509,
		Ftype1: int8('f'),
	},
	9000: {
		Fword:  __ccgo_ts + 64515,
		Ftype1: int8('k'),
	},
	9001: {
		Fword:  __ccgo_ts + 64525,
		Ftype1: int8('n'),
	},
	9002: {
		Fword:  __ccgo_ts + 64540,
		Ftype1: int8('k'),
	},
	9003: {
		Fword:  __ccgo_ts + 64559,
		Ftype1: int8('k'),
	},
	9004: {
		Fword:  __ccgo_ts + 64567,
		Ftype1: int8('T'),
	},
	9005: {
		Fword:  __ccgo_ts + 64573,
		Ftype1: int8('k'),
	},
	9006: {
		Fword:  __ccgo_ts + 64583,
		Ftype1: int8('f'),
	},
	9007: {
		Fword:  __ccgo_ts + 64601,
		Ftype1: int8('f'),
	},
	9008: {
		Fword:  __ccgo_ts + 64622,
		Ftype1: int8('k'),
	},
	9009: {
		Fword:  __ccgo_ts + 64628,
		Ftype1: int8('f'),
	},
	9010: {
		Fword:  __ccgo_ts + 64639,
		Ftype1: int8('f'),
	},
	9011: {
		Fword:  __ccgo_ts + 64650,
		Ftype1: int8('f'),
	},
	9012: {
		Fword:  __ccgo_ts + 64658,
		Ftype1: int8('f'),
	},
	9013: {
		Fword:  __ccgo_ts + 64664,
		Ftype1: int8('f'),
	},
	9014: {
		Fword:  __ccgo_ts + 64674,
		Ftype1: int8('f'),
	},
	9015: {
		Fword:  __ccgo_ts + 64686,
		Ftype1: int8('f'),
	},
	9016: {
		Fword:  __ccgo_ts + 64700,
		Ftype1: int8('f'),
	},
	9017: {
		Fword:  __ccgo_ts + 64715,
		Ftype1: int8('f'),
	},
	9018: {
		Fword:  __ccgo_ts + 64723,
		Ftype1: int8('E'),
	},
	9019: {
		Fword:  __ccgo_ts + 64734,
		Ftype1: int8('f'),
	},
	9020: {
		Fword:  __ccgo_ts + 64739,
		Ftype1: int8('f'),
	},
	9021: {
		Fword:  __ccgo_ts + 64746,
		Ftype1: int8('f'),
	},
	9022: {
		Fword:  __ccgo_ts + 64757,
		Ftype1: int8('k'),
	},
	9023: {
		Fword:  __ccgo_ts + 64763,
		Ftype1: int8('f'),
	},
	9024: {
		Fword:  __ccgo_ts + 64768,
		Ftype1: int8('k'),
	},
	9025: {
		Fword:  __ccgo_ts + 64773,
		Ftype1: int8('k'),
	},
	9026: {
		Fword:  __ccgo_ts + 64784,
		Ftype1: int8('k'),
	},
	9027: {
		Fword:  __ccgo_ts + 64790,
		Ftype1: int8('k'),
	},
	9028: {
		Fword:  __ccgo_ts + 64801,
		Ftype1: int8('t'),
	},
	9029: {
		Fword:  __ccgo_ts + 64806,
		Ftype1: int8('k'),
	},
	9030: {
		Fword:  __ccgo_ts + 64817,
		Ftype1: int8('t'),
	},
	9031: {
		Fword:  __ccgo_ts + 64826,
		Ftype1: int8('t'),
	},
	9032: {
		Fword:  __ccgo_ts + 64836,
		Ftype1: int8('t'),
	},
	9033: {
		Fword:  __ccgo_ts + 64850,
		Ftype1: int8('o'),
	},
	9034: {
		Fword:  __ccgo_ts + 64857,
		Ftype1: int8('f'),
	},
	9035: {
		Fword:  __ccgo_ts + 64870,
		Ftype1: int8('f'),
	},
	9036: {
		Fword:  __ccgo_ts + 64885,
		Ftype1: int8('f'),
	},
	9037: {
		Fword:  __ccgo_ts + 64900,
		Ftype1: int8('f'),
	},
	9038: {
		Fword:  __ccgo_ts + 64922,
		Ftype1: int8('f'),
	},
	9039: {
		Fword:  __ccgo_ts + 64944,
		Ftype1: int8('f'),
	},
	9040: {
		Fword:  __ccgo_ts + 64958,
		Ftype1: int8('t'),
	},
	9041: {
		Fword:  __ccgo_ts + 64966,
		Ftype1: int8('t'),
	},
	9042: {
		Fword:  __ccgo_ts + 64978,
		Ftype1: int8('t'),
	},
	9043: {
		Fword:  __ccgo_ts + 64986,
		Ftype1: int8('t'),
	},
	9044: {
		Fword:  __ccgo_ts + 64999,
		Ftype1: int8('t'),
	},
	9045: {
		Fword:  __ccgo_ts + 65007,
		Ftype1: int8('k'),
	},
	9046: {
		Fword:  __ccgo_ts + 65015,
		Ftype1: int8('f'),
	},
	9047: {
		Fword:  __ccgo_ts + 65028,
		Ftype1: int8('k'),
	},
	9048: {
		Fword:  __ccgo_ts + 65035,
		Ftype1: int8('k'),
	},
	9049: {
		Fword:  __ccgo_ts + 65042,
		Ftype1: int8('k'),
	},
	9050: {
		Fword:  __ccgo_ts + 65050,
		Ftype1: int8('f'),
	},
	9051: {
		Fword:  __ccgo_ts + 65060,
		Ftype1: int8('k'),
	},
	9052: {
		Fword:  __ccgo_ts + 65068,
		Ftype1: int8('k'),
	},
	9053: {
		Fword:  __ccgo_ts + 65077,
		Ftype1: int8('k'),
	},
	9054: {
		Fword:  __ccgo_ts + 65086,
		Ftype1: int8('k'),
	},
	9055: {
		Fword:  __ccgo_ts + 65093,
		Ftype1: int8('f'),
	},
	9056: {
		Fword:  __ccgo_ts + 65101,
		Ftype1: int8('k'),
	},
	9057: {
		Fword:  __ccgo_ts + 65108,
		Ftype1: int8('n'),
	},
	9058: {
		Fword:  __ccgo_ts + 65114,
		Ftype1: int8('k'),
	},
	9059: {
		Fword:  __ccgo_ts + 65125,
		Ftype1: int8('k'),
	},
	9060: {
		Fword:  __ccgo_ts + 65137,
		Ftype1: int8('k'),
	},
	9061: {
		Fword:  __ccgo_ts + 65154,
		Ftype1: int8('o'),
	},
	9062: {
		Fword:  __ccgo_ts + 65160,
		Ftype1: int8('f'),
	},
	9063: {
		Fword:  __ccgo_ts + 65166,
		Ftype1: int8('f'),
	},
	9064: {
		Fword:  __ccgo_ts + 65170,
		Ftype1: int8('f'),
	},
	9065: {
		Fword:  __ccgo_ts + 65180,
		Ftype1: int8('f'),
	},
	9066: {
		Fword:  __ccgo_ts + 65191,
		Ftype1: int8('f'),
	},
	9067: {
		Fword:  __ccgo_ts + 65203,
		Ftype1: int8('f'),
	},
	9068: {
		Fword:  __ccgo_ts + 65208,
		Ftype1: int8('f'),
	},
	9069: {
		Fword:  __ccgo_ts + 65214,
		Ftype1: int8('f'),
	},
	9070: {
		Fword:  __ccgo_ts + 65226,
		Ftype1: int8('k'),
	},
	9071: {
		Fword:  __ccgo_ts + 65233,
		Ftype1: int8('k'),
	},
	9072: {
		Fword:  __ccgo_ts + 65241,
		Ftype1: int8('f'),
	},
	9073: {
		Fword:  __ccgo_ts + 65251,
		Ftype1: int8('f'),
	},
	9074: {
		Fword:  __ccgo_ts + 65266,
		Ftype1: int8('k'),
	},
	9075: {
		Fword:  __ccgo_ts + 65285,
		Ftype1: int8('f'),
	},
	9076: {
		Fword:  __ccgo_ts + 65297,
		Ftype1: int8('E'),
	},
	9077: {
		Fword:  __ccgo_ts + 65304,
		Ftype1: int8('E'),
	},
	9078: {
		Fword:  __ccgo_ts + 65315,
		Ftype1: int8('E'),
	},
	9079: {
		Fword:  __ccgo_ts + 65331,
		Ftype1: int8('k'),
	},
	9080: {
		Fword:  __ccgo_ts + 65341,
		Ftype1: int8('k'),
	},
	9081: {
		Fword:  __ccgo_ts + 65351,
		Ftype1: int8('t'),
	},
	9082: {
		Fword:  __ccgo_ts + 65358,
		Ftype1: int8('t'),
	},
	9083: {
		Fword:  __ccgo_ts + 65366,
		Ftype1: int8('t'),
	},
	9084: {
		Fword:  __ccgo_ts + 65374,
		Ftype1: int8('t'),
	},
	9085: {
		Fword:  __ccgo_ts + 65382,
		Ftype1: int8('f'),
	},
	9086: {
		Fword:  __ccgo_ts + 65397,
		Ftype1: int8('f'),
	},
	9087: {
		Fword:  __ccgo_ts + 65410,
		Ftype1: int8('E'),
	},
	9088: {
		Fword:  __ccgo_ts + 65414,
		Ftype1: int8('f'),
	},
	9089: {
		Fword:  __ccgo_ts + 65422,
		Ftype1: int8('f'),
	},
	9090: {
		Fword:  __ccgo_ts + 65430,
		Ftype1: int8('f'),
	},
	9091: {
		Fword:  __ccgo_ts + 65437,
		Ftype1: int8('f'),
	},
	9092: {
		Fword:  __ccgo_ts + 65445,
		Ftype1: int8('f'),
	},
	9093: {
		Fword:  __ccgo_ts + 65454,
		Ftype1: int8('f'),
	},
	9094: {
		Fword:  __ccgo_ts + 65465,
		Ftype1: int8('f'),
	},
	9095: {
		Fword:  __ccgo_ts + 65477,
		Ftype1: int8('f'),
	},
	9096: {
		Fword:  __ccgo_ts + 65481,
		Ftype1: int8('f'),
	},
	9097: {
		Fword:  __ccgo_ts + 65486,
		Ftype1: int8('f'),
	},
	9098: {
		Fword:  __ccgo_ts + 65491,
		Ftype1: int8('n'),
	},
	9099: {
		Fword:  __ccgo_ts + 65496,
		Ftype1: int8('T'),
	},
	9100: {
		Fword:  __ccgo_ts + 65505,
		Ftype1: int8('f'),
	},
	9101: {
		Fword:  __ccgo_ts + 65510,
		Ftype1: int8('k'),
	},
	9102: {
		Fword:  __ccgo_ts + 65517,
		Ftype1: int8('f'),
	},
	9103: {
		Fword:  __ccgo_ts + 65530,
		Ftype1: int8('f'),
	},
	9104: {
		Fword:  __ccgo_ts + 65541,
		Ftype1: int8('k'),
	},
	9105: {
		Fword:  __ccgo_ts + 65549,
		Ftype1: int8('o'),
	},
	9106: {
		Fword:  __ccgo_ts + 65560,
		Ftype1: int8('f'),
	},
	9107: {
		Fword:  __ccgo_ts + 65564,
		Ftype1: int8('f'),
	},
	9108: {
		Fword:  __ccgo_ts + 65570,
		Ftype1: int8('f'),
	},
	9109: {
		Fword:  __ccgo_ts + 65593,
		Ftype1: int8('t'),
	},
	9110: {
		Fword:  __ccgo_ts + 65602,
		Ftype1: int8('t'),
	},
	9111: {
		Fword:  __ccgo_ts + 65614,
		Ftype1: int8('f'),
	},
	9112: {
		Fword:  __ccgo_ts + 65619,
		Ftype1: int8('f'),
	},
	9113: {
		Fword:  __ccgo_ts + 65627,
		Ftype1: int8('o'),
	},
	9114: {
		Fword:  __ccgo_ts + 65634,
		Ftype1: int8('o'),
	},
	9115: {
		Fword:  __ccgo_ts + 65646,
		Ftype1: int8('f'),
	},
	9116: {
		Fword:  __ccgo_ts + 65652,
		Ftype1: int8('k'),
	},
	9117: {
		Fword:  __ccgo_ts + 65660,
		Ftype1: int8('k'),
	},
	9118: {
		Fword:  __ccgo_ts + 65669,
		Ftype1: int8('f'),
	},
	9119: {
		Fword:  __ccgo_ts + 65680,
		Ftype1: int8('k'),
	},
	9120: {
		Fword:  __ccgo_ts + 65684,
		Ftype1: int8('k'),
	},
	9121: {
		Fword:  __ccgo_ts + 65697,
		Ftype1: int8('f'),
	},
	9122: {
		Fword:  __ccgo_ts + 65712,
		Ftype1: int8('k'),
	},
	9123: {
		Fword:  __ccgo_ts + 65721,
		Ftype1: int8('k'),
	},
	9124: {
		Fword:  __ccgo_ts + 65732,
		Ftype1: int8('k'),
	},
	9125: {
		Fword:  __ccgo_ts + 65747,
		Ftype1: int8('k'),
	},
	9126: {
		Fword:  __ccgo_ts + 65765,
		Ftype1: int8('k'),
	},
	9127: {
		Fword:  __ccgo_ts + 65775,
		Ftype1: int8('k'),
	},
	9128: {
		Fword:  __ccgo_ts + 65795,
		Ftype1: int8('k'),
	},
	9129: {
		Fword:  __ccgo_ts + 65808,
		Ftype1: int8('k'),
	},
	9130: {
		Fword:  __ccgo_ts + 65825,
		Ftype1: int8('f'),
	},
	9131: {
		Fword:  __ccgo_ts + 65846,
		Ftype1: int8('f'),
	},
	9132: {
		Fword:  __ccgo_ts + 65851,
		Ftype1: int8('k'),
	},
	9133: {
		Fword:  __ccgo_ts + 65855,
		Ftype1: int8('k'),
	},
	9134: {
		Fword:  __ccgo_ts + 65864,
		Ftype1: int8('f'),
	},
	9135: {
		Fword:  __ccgo_ts + 65884,
		Ftype1: int8('f'),
	},
	9136: {
		Fword:  __ccgo_ts + 65895,
		Ftype1: int8('f'),
	},
	9137: {
		Fword:  __ccgo_ts + 65902,
		Ftype1: int8('f'),
	},
	9138: {
		Fword:  __ccgo_ts + 65913,
		Ftype1: int8('f'),
	},
	9139: {
		Fword:  __ccgo_ts + 65925,
		Ftype1: int8('k'),
	},
	9140: {
		Fword:  __ccgo_ts + 65939,
		Ftype1: int8('f'),
	},
	9141: {
		Fword:  __ccgo_ts + 65946,
		Ftype1: int8('f'),
	},
	9142: {
		Fword:  __ccgo_ts + 65954,
		Ftype1: int8('f'),
	},
	9143: {
		Fword:  __ccgo_ts + 65962,
		Ftype1: int8('f'),
	},
	9144: {
		Fword:  __ccgo_ts + 65973,
		Ftype1: int8('f'),
	},
	9145: {
		Fword:  __ccgo_ts + 65989,
		Ftype1: int8('f'),
	},
	9146: {
		Fword:  __ccgo_ts + 65996,
		Ftype1: int8('f'),
	},
	9147: {
		Fword:  __ccgo_ts + 66008,
		Ftype1: int8('f'),
	},
	9148: {
		Fword:  __ccgo_ts + 66014,
		Ftype1: int8('f'),
	},
	9149: {
		Fword:  __ccgo_ts + 66022,
		Ftype1: int8('f'),
	},
	9150: {
		Fword:  __ccgo_ts + 66029,
		Ftype1: int8('f'),
	},
	9151: {
		Fword:  __ccgo_ts + 66039,
		Ftype1: int8('f'),
	},
	9152: {
		Fword:  __ccgo_ts + 66055,
		Ftype1: int8('f'),
	},
	9153: {
		Fword:  __ccgo_ts + 66063,
		Ftype1: int8('f'),
	},
	9154: {
		Fword:  __ccgo_ts + 66067,
		Ftype1: int8('f'),
	},
	9155: {
		Fword:  __ccgo_ts + 66076,
		Ftype1: int8('f'),
	},
	9156: {
		Fword:  __ccgo_ts + 66087,
		Ftype1: int8('f'),
	},
	9157: {
		Fword:  __ccgo_ts + 66097,
		Ftype1: int8('f'),
	},
	9158: {
		Fword:  __ccgo_ts + 66109,
		Ftype1: int8('f'),
	},
	9159: {
		Fword:  __ccgo_ts + 66121,
		Ftype1: int8('n'),
	},
	9160: {
		Fword:  __ccgo_ts + 66139,
		Ftype1: int8('f'),
	},
	9161: {
		Fword:  __ccgo_ts + 66166,
		Ftype1: int8('f'),
	},
	9162: {
		Fword:  __ccgo_ts + 66188,
		Ftype1: int8('f'),
	},
	9163: {
		Fword:  __ccgo_ts + 66210,
		Ftype1: int8('f'),
	},
	9164: {
		Fword:  __ccgo_ts + 66221,
		Ftype1: int8('k'),
	},
	9165: {
		Fword:  __ccgo_ts + 66232,
		Ftype1: int8('f'),
	},
	9166: {
		Fword:  __ccgo_ts + 66240,
		Ftype1: int8('f'),
	},
	9167: {
		Fword:  __ccgo_ts + 66252,
		Ftype1: int8('f'),
	},
	9168: {
		Fword:  __ccgo_ts + 66270,
		Ftype1: int8('k'),
	},
	9169: {
		Fword:  __ccgo_ts + 66281,
		Ftype1: int8('f'),
	},
	9170: {
		Fword:  __ccgo_ts + 66293,
		Ftype1: int8('k'),
	},
	9171: {
		Fword:  __ccgo_ts + 66302,
		Ftype1: int8('f'),
	},
	9172: {
		Fword:  __ccgo_ts + 66316,
		Ftype1: int8('n'),
	},
	9173: {
		Fword:  __ccgo_ts + 66322,
		Ftype1: int8('f'),
	},
	9174: {
		Fword:  __ccgo_ts + 66326,
		Ftype1: int8('k'),
	},
	9175: {
		Fword:  __ccgo_ts + 66337,
		Ftype1: int8('f'),
	},
	9176: {
		Fword:  __ccgo_ts + 66354,
		Ftype1: int8('t'),
	},
	9177: {
		Fword:  __ccgo_ts + 66359,
		Ftype1: int8('f'),
	},
	9178: {
		Fword:  __ccgo_ts + 66367,
		Ftype1: int8('f'),
	},
	9179: {
		Fword:  __ccgo_ts + 66375,
		Ftype1: int8('f'),
	},
	9180: {
		Fword:  __ccgo_ts + 66385,
		Ftype1: int8('k'),
	},
	9181: {
		Fword:  __ccgo_ts + 66390,
		Ftype1: int8('k'),
	},
	9182: {
		Fword:  __ccgo_ts + 66395,
		Ftype1: int8('f'),
	},
	9183: {
		Fword:  __ccgo_ts + 66404,
		Ftype1: int8('f'),
	},
	9184: {
		Fword:  __ccgo_ts + 66418,
		Ftype1: int8('f'),
	},
	9185: {
		Fword:  __ccgo_ts + 66428,
		Ftype1: int8('f'),
	},
	9186: {
		Fword:  __ccgo_ts + 66439,
		Ftype1: int8('t'),
	},
	9187: {
		Fword:  __ccgo_ts + 66449,
		Ftype1: int8('f'),
	},
	9188: {
		Fword:  __ccgo_ts + 66462,
		Ftype1: int8('f'),
	},
	9189: {
		Fword:  __ccgo_ts + 66472,
		Ftype1: int8('f'),
	},
	9190: {
		Fword:  __ccgo_ts + 66484,
		Ftype1: int8('f'),
	},
	9191: {
		Fword:  __ccgo_ts + 66496,
		Ftype1: int8('k'),
	},
	9192: {
		Fword:  __ccgo_ts + 66505,
		Ftype1: int8('k'),
	},
	9193: {
		Fword:  __ccgo_ts + 66513,
		Ftype1: int8('k'),
	},
	9194: {
		Fword:  __ccgo_ts + 66522,
		Ftype1: int8('f'),
	},
	9195: {
		Fword:  __ccgo_ts + 66539,
		Ftype1: int8('k'),
	},
	9196: {
		Fword:  __ccgo_ts + 66543,
		Ftype1: int8('f'),
	},
	9197: {
		Fword:  __ccgo_ts + 66549,
		Ftype1: int8('f'),
	},
	9198: {
		Fword:  __ccgo_ts + 66563,
		Ftype1: int8('f'),
	},
	9199: {
		Fword:  __ccgo_ts + 66572,
		Ftype1: int8('f'),
	},
	9200: {
		Fword:  __ccgo_ts + 66582,
		Ftype1: int8('f'),
	},
	9201: {
		Fword:  __ccgo_ts + 66590,
		Ftype1: int8('f'),
	},
	9202: {
		Fword:  __ccgo_ts + 66598,
		Ftype1: int8('f'),
	},
	9203: {
		Fword:  __ccgo_ts + 66606,
		Ftype1: int8('f'),
	},
	9204: {
		Fword:  __ccgo_ts + 66613,
		Ftype1: int8('f'),
	},
	9205: {
		Fword:  __ccgo_ts + 66623,
		Ftype1: int8('f'),
	},
	9206: {
		Fword:  __ccgo_ts + 66634,
		Ftype1: int8('f'),
	},
	9207: {
		Fword:  __ccgo_ts + 66647,
		Ftype1: int8('n'),
	},
	9208: {
		Fword:  __ccgo_ts + 66656,
		Ftype1: int8('f'),
	},
	9209: {
		Fword:  __ccgo_ts + 66678,
		Ftype1: int8('f'),
	},
	9210: {
		Fword:  __ccgo_ts + 66688,
		Ftype1: int8('k'),
	},
	9211: {
		Fword:  __ccgo_ts + 66696,
		Ftype1: int8('f'),
	},
	9212: {
		Fword:  __ccgo_ts + 66714,
		Ftype1: int8('f'),
	},
	9213: {
		Fword:  __ccgo_ts + 66719,
		Ftype1: int8('1'),
	},
	9214: {
		Fword:  __ccgo_ts + 66724,
		Ftype1: int8('f'),
	},
	9215: {
		Fword:  __ccgo_ts + 66730,
		Ftype1: int8('f'),
	},
	9216: {
		Fword:  __ccgo_ts + 66739,
		Ftype1: int8('T'),
	},
	9217: {
		Fword:  __ccgo_ts + 66743,
		Ftype1: int8('f'),
	},
	9218: {
		Fword:  __ccgo_ts + 66752,
		Ftype1: int8('f'),
	},
	9219: {
		Fword:  __ccgo_ts + 66764,
		Ftype1: int8('f'),
	},
	9220: {
		Fword:  __ccgo_ts + 66774,
		Ftype1: int8('f'),
	},
	9221: {
		Fword:  __ccgo_ts + 66781,
		Ftype1: int8('f'),
	},
	9222: {
		Fword:  __ccgo_ts + 66794,
		Ftype1: int8('f'),
	},
	9223: {
		Fword:  __ccgo_ts + 66802,
		Ftype1: int8('f'),
	},
	9224: {
		Fword:  __ccgo_ts + 66812,
		Ftype1: int8('f'),
	},
	9225: {
		Fword:  __ccgo_ts + 66818,
		Ftype1: int8('o'),
	},
	9226: {
		Fword:  __ccgo_ts + 66826,
		Ftype1: int8('f'),
	},
	9227: {
		Fword:  __ccgo_ts + 66837,
		Ftype1: int8('f'),
	},
	9228: {
		Fword:  __ccgo_ts + 66855,
		Ftype1: int8('k'),
	},
	9229: {
		Fword:  __ccgo_ts + 66860,
		Ftype1: int8('f'),
	},
	9230: {
		Fword:  __ccgo_ts + 66866,
		Ftype1: int8('f'),
	},
	9231: {
		Fword:  __ccgo_ts + 66874,
		Ftype1: int8('U'),
	},
	9232: {
		Fword:  __ccgo_ts + 66880,
		Ftype1: int8('U'),
	},
	9233: {
		Fword:  __ccgo_ts + 66890,
		Ftype1: int8('U'),
	},
	9234: {
		Fword:  __ccgo_ts + 66909,
		Ftype1: int8('U'),
	},
	9235: {
		Fword:  __ccgo_ts + 66924,
		Ftype1: int8('U'),
	},
	9236: {
		Fword:  __ccgo_ts + 66943,
		Ftype1: int8('n'),
	},
	9237: {
		Fword:  __ccgo_ts + 66950,
		Ftype1: int8('f'),
	},
	9238: {
		Fword:  __ccgo_ts + 66965,
		Ftype1: int8('U'),
	},
	9239: {
		Fword:  __ccgo_ts + 66972,
		Ftype1: int8('v'),
	},
	9240: {
		Fword:  __ccgo_ts + 66980,
		Ftype1: int8('k'),
	},
	9241: {
		Fword:  __ccgo_ts + 66987,
		Ftype1: int8('f'),
	},
	9242: {
		Fword:  __ccgo_ts + 66994,
		Ftype1: int8('k'),
	},
	9243: {
		Fword:  __ccgo_ts + 67003,
		Ftype1: int8('E'),
	},
	9244: {
		Fword:  __ccgo_ts + 67010,
		Ftype1: int8('f'),
	},
	9245: {
		Fword:  __ccgo_ts + 67020,
		Ftype1: int8('f'),
	},
	9246: {
		Fword:  __ccgo_ts + 67026,
		Ftype1: int8('f'),
	},
	9247: {
		Fword:  __ccgo_ts + 67036,
		Ftype1: int8('f'),
	},
	9248: {
		Fword:  __ccgo_ts + 67046,
		Ftype1: int8('k'),
	},
	9249: {
		Fword:  __ccgo_ts + 67052,
		Ftype1: int8('T'),
	},
	9250: {
		Fword:  __ccgo_ts + 67056,
		Ftype1: int8('n'),
	},
	9251: {
		Fword:  __ccgo_ts + 67061,
		Ftype1: int8('n'),
	},
	9252: {
		Fword:  __ccgo_ts + 67069,
		Ftype1: int8('f'),
	},
	9253: {
		Fword:  __ccgo_ts + 67085,
		Ftype1: int8('n'),
	},
	9254: {
		Fword:  __ccgo_ts + 67095,
		Ftype1: int8('f'),
	},
	9255: {
		Fword:  __ccgo_ts + 67101,
		Ftype1: int8('k'),
	},
	9256: {
		Fword:  __ccgo_ts + 67110,
		Ftype1: int8('k'),
	},
	9257: {
		Fword:  __ccgo_ts + 67119,
		Ftype1: int8('k'),
	},
	9258: {
		Fword:  __ccgo_ts + 67133,
		Ftype1: int8('f'),
	},
	9259: {
		Fword:  __ccgo_ts + 67150,
		Ftype1: int8('f'),
	},
	9260: {
		Fword:  __ccgo_ts + 67178,
		Ftype1: int8('f'),
	},
	9261: {
		Fword:  __ccgo_ts + 67203,
		Ftype1: int8('f'),
	},
	9262: {
		Fword:  __ccgo_ts + 67208,
		Ftype1: int8('f'),
	},
	9263: {
		Fword:  __ccgo_ts + 67219,
		Ftype1: int8('k'),
	},
	9264: {
		Fword:  __ccgo_ts + 67226,
		Ftype1: int8('f'),
	},
	9265: {
		Fword:  __ccgo_ts + 67230,
		Ftype1: int8('k'),
	},
	9266: {
		Fword:  __ccgo_ts + 67240,
		Ftype1: int8('t'),
	},
	9267: {
		Fword:  __ccgo_ts + 67248,
		Ftype1: int8('k'),
	},
	9268: {
		Fword:  __ccgo_ts + 67261,
		Ftype1: int8('f'),
	},
	9269: {
		Fword:  __ccgo_ts + 67270,
		Ftype1: int8('f'),
	},
	9270: {
		Fword:  __ccgo_ts + 67275,
		Ftype1: int8('k'),
	},
	9271: {
		Fword:  __ccgo_ts + 67283,
		Ftype1: int8('f'),
	},
	9272: {
		Fword:  __ccgo_ts + 67291,
		Ftype1: int8('f'),
	},
	9273: {
		Fword:  __ccgo_ts + 67300,
		Ftype1: int8('f'),
	},
	9274: {
		Fword:  __ccgo_ts + 67321,
		Ftype1: int8('f'),
	},
	9275: {
		Fword:  __ccgo_ts + 67340,
		Ftype1: int8('f'),
	},
	9276: {
		Fword:  __ccgo_ts + 67348,
		Ftype1: int8('t'),
	},
	9277: {
		Fword:  __ccgo_ts + 67353,
		Ftype1: int8('k'),
	},
	9278: {
		Fword:  __ccgo_ts + 67358,
		Ftype1: int8('n'),
	},
	9279: {
		Fword:  __ccgo_ts + 67366,
		Ftype1: int8('E'),
	},
	9280: {
		Fword:  __ccgo_ts + 67380,
		Ftype1: int8('E'),
	},
	9281: {
		Fword:  __ccgo_ts + 67396,
		Ftype1: int8('E'),
	},
	9282: {
		Fword:  __ccgo_ts + 67409,
		Ftype1: int8('f'),
	},
	9283: {
		Fword:  __ccgo_ts + 67414,
		Ftype1: int8('f'),
	},
	9284: {
		Fword:  __ccgo_ts + 67422,
		Ftype1: int8('f'),
	},
	9285: {
		Fword:  __ccgo_ts + 67434,
		Ftype1: int8('f'),
	},
	9286: {
		Fword:  __ccgo_ts + 67445,
		Ftype1: int8('k'),
	},
	9287: {
		Fword:  __ccgo_ts + 67450,
		Ftype1: int8('k'),
	},
	9288: {
		Fword:  __ccgo_ts + 67456,
		Ftype1: int8('T'),
	},
	9289: {
		Fword:  __ccgo_ts + 67462,
		Ftype1: int8('f'),
	},
	9290: {
		Fword:  __ccgo_ts + 67475,
		Ftype1: int8('n'),
	},
	9291: {
		Fword:  __ccgo_ts + 67480,
		Ftype1: int8('k'),
	},
	9292: {
		Fword:  __ccgo_ts + 67492,
		Ftype1: int8('f'),
	},
	9293: {
		Fword:  __ccgo_ts + 67499,
		Ftype1: int8('f'),
	},
	9294: {
		Fword:  __ccgo_ts + 67510,
		Ftype1: int8('f'),
	},
	9295: {
		Fword:  __ccgo_ts + 67520,
		Ftype1: int8('f'),
	},
	9296: {
		Fword:  __ccgo_ts + 67531,
		Ftype1: int8('f'),
	},
	9297: {
		Fword:  __ccgo_ts + 67541,
		Ftype1: int8('f'),
	},
	9298: {
		Fword:  __ccgo_ts + 67551,
		Ftype1: int8('f'),
	},
	9299: {
		Fword:  __ccgo_ts + 67561,
		Ftype1: int8('f'),
	},
	9300: {
		Fword:  __ccgo_ts + 67567,
		Ftype1: int8('f'),
	},
	9301: {
		Fword:  __ccgo_ts + 67575,
		Ftype1: int8('f'),
	},
	9302: {
		Fword:  __ccgo_ts + 67583,
		Ftype1: int8('f'),
	},
	9303: {
		Fword:  __ccgo_ts + 67602,
		Ftype1: int8('&'),
	},
	9304: {
		Fword:  __ccgo_ts + 67606,
		Ftype1: int8('f'),
	},
	9305: {
		Fword:  __ccgo_ts + 67612,
		Ftype1: int8('f'),
	},
	9306: {
		Fword:  __ccgo_ts + 67625,
		Ftype1: int8('k'),
	},
	9307: {
		Fword:  __ccgo_ts + 67642,
		Ftype1: int8('f'),
	},
	9308: {
		Fword:  __ccgo_ts + 67647,
		Ftype1: int8('f'),
	},
	9309: {
		Fword:  __ccgo_ts + 67656,
		Ftype1: int8('k'),
	},
	9310: {
		Fword:  __ccgo_ts + 67667,
		Ftype1: int8('f'),
	},
	9311: {
		Fword:  __ccgo_ts + 67676,
		Ftype1: int8('k'),
	},
	9312: {
		Fword:  __ccgo_ts + 67685,
		Ftype1: int8('o'),
	},
	9313: {
		Fword:  __ccgo_ts + 67688,
		Ftype1: int8('t'),
	},
	9314: {
		Fword:  __ccgo_ts + 67698,
		Ftype1: int8('t'),
	},
	9315: {
		Fword:  __ccgo_ts + 67705,
		Ftype1: int8('t'),
	},
	9316: {
		Fword:  __ccgo_ts + 67711,
		Ftype1: int8('t'),
	},
	9317: {
		Fword:  __ccgo_ts + 67719,
		Ftype1: int8('t'),
	},
	9318: {
		Fword:  __ccgo_ts + 67727,
		Ftype1: int8('t'),
	},
	9319: {
		Fword:  __ccgo_ts + 67735,
		Ftype1: int8('t'),
	},
	9320: {
		Fword:  __ccgo_ts + 67743,
		Ftype1: int8('t'),
	},
	9321: {
		Fword:  __ccgo_ts + 67750,
		Ftype1: int8('t'),
	},
	9322: {
		Fword:  __ccgo_ts + 67757,
		Ftype1: int8('t'),
	},
	9323: {
		Fword:  __ccgo_ts + 67764,
		Ftype1: int8('t'),
	},
	9324: {
		Fword:  __ccgo_ts + 67771,
		Ftype1: int8('t'),
	},
	9325: {
		Fword:  __ccgo_ts + 67777,
		Ftype1: int8('t'),
	},
	9326: {
		Fword:  __ccgo_ts + 67786,
		Ftype1: int8('t'),
	},
	9327: {
		Fword:  __ccgo_ts + 67793,
		Ftype1: int8('t'),
	},
	9328: {
		Fword:  __ccgo_ts + 67801,
		Ftype1: int8('t'),
	},
	9329: {
		Fword:  __ccgo_ts + 67806,
		Ftype1: int8('t'),
	},
	9330: {
		Fword:  __ccgo_ts + 67815,
		Ftype1: int8('t'),
	},
	9331: {
		Fword:  __ccgo_ts + 67822,
		Ftype1: int8('t'),
	},
	9332: {
		Fword:  __ccgo_ts + 67830,
		Ftype1: int8('t'),
	},
	9333: {
		Fword:  __ccgo_ts + 67835,
		Ftype1: int8('t'),
	},
	9334: {
		Fword:  __ccgo_ts + 67844,
		Ftype1: int8('t'),
	},
	9335: {
		Fword:  __ccgo_ts + 67851,
		Ftype1: int8('t'),
	},
	9336: {
		Fword:  __ccgo_ts + 67858,
		Ftype1: int8('t'),
	},
	9337: {
		Fword:  __ccgo_ts + 67866,
		Ftype1: int8('t'),
	},
	9338: {
		Fword:  __ccgo_ts + 67874,
		Ftype1: int8('t'),
	},
	9339: {
		Fword:  __ccgo_ts + 67882,
		Ftype1: int8('t'),
	},
	9340: {
		Fword:  __ccgo_ts + 67890,
		Ftype1: int8('t'),
	},
	9341: {
		Fword:  __ccgo_ts + 67897,
		Ftype1: int8('t'),
	},
	9342: {
		Fword:  __ccgo_ts + 67907,
		Ftype1: int8('t'),
	},
	9343: {
		Fword:  __ccgo_ts + 67913,
		Ftype1: int8('t'),
	},
	9344: {
		Fword:  __ccgo_ts + 67919,
		Ftype1: int8('t'),
	},
	9345: {
		Fword:  __ccgo_ts + 67927,
		Ftype1: int8('t'),
	},
	9346: {
		Fword:  __ccgo_ts + 67933,
		Ftype1: int8('t'),
	},
	9347: {
		Fword:  __ccgo_ts + 67939,
		Ftype1: int8('t'),
	},
	9348: {
		Fword:  __ccgo_ts + 67945,
		Ftype1: int8('o'),
	},
	9349: {
		Fword:  __ccgo_ts + 67948,
		Ftype1: int8('o'),
	},
	9350: {
		Fword:  __ccgo_ts + 67951,
		Ftype1: int8('&'),
	},
	9351: {
		Fword:  __ccgo_ts + 67954,
		Ftype1: int8('o'),
	},
}
var sql_keywords_sz = uint64(9352)

/* faster than calling out to libc isdigit */

// C documentation
//
//	/*
//	 * not making public just yet
//	 */
type sqli_token_types = int32

const TYPE_NONE = 0
const TYPE_KEYWORD = 107
const TYPE_UNION = 85
const TYPE_GROUP = 66
const TYPE_EXPRESSION = 69
const TYPE_SQLTYPE = 116
const TYPE_FUNCTION = 102
const TYPE_BAREWORD = 110
const TYPE_NUMBER = 49
const TYPE_VARIABLE = 118
const TYPE_STRING = 115
const TYPE_OPERATOR = 111
const TYPE_LOGIC_OPERATOR = 38
const TYPE_COMMENT = 99
const TYPE_COLLATE = 65
const TYPE_LEFTPARENS = 40
const TYPE_RIGHTPARENS = 41
const TYPE_LEFTBRACE = 123
const TYPE_RIGHTBRACE = 125
const TYPE_DOT = 46
const TYPE_COMMA = 44
const TYPE_COLON = 58
const TYPE_SEMICOLON = 59
const TYPE_TSQL = 84
const TYPE_UNKNOWN = 63
const TYPE_EVIL = 88
const TYPE_FINGERPRINT = 70
const TYPE_BACKSLASH = 92

// C documentation
//
//	/**
//	 * Initializes parsing state
//	 *
//	 */
func flag2delim(tls *libc.TLS, flag int32) (r int8) {
	if flag&int32(FLAG_QUOTE_SINGLE) != 0 {
		return int8('\'')
	} else {
		if flag&int32(FLAG_QUOTE_DOUBLE) != 0 {
			return int8('"')
		} else {
			return int8('\000')
		}
	}
	return r
}

// C documentation
//
//	/* memchr2 finds a string of 2 characters inside another string
//	 * This a specialized version of "memmem" or "memchr".
//	 * 'memmem' doesn't exist on all platforms
//	 *
//	 * Porting notes: this is just a special version of
//	 *    astring.find("AB")
//	 *
//	 */
func memchr2(tls *libc.TLS, haystack uintptr, haystack_len size_t, c0 int8, c1 int8) (r uintptr) {
	var cur, last uintptr
	_, _ = cur, last
	cur = haystack
	last = haystack + uintptr(haystack_len) - uintptr(1)
	if haystack_len < uint64(2) {
		return libc.UintptrFromInt32(0)
	}
	for cur < last {
		/* safe since cur < len - 1 always */
		if int32(**(**int8)(__ccgo_up(cur))) == int32(c0) && int32(**(**int8)(__ccgo_up(cur + 1))) == int32(c1) {
			return cur
		}
		cur = cur + uintptr(1)
	}
	return libc.UintptrFromInt32(0)
}

// C documentation
//
//	/**
//	 * memmem might not exist on some systems
//	 */
func my_memmem(tls *libc.TLS, haystack uintptr, hlen size_t, needle uintptr, nlen size_t) (r uintptr) {
	var cur, last uintptr
	var v1 bool
	_, _, _ = cur, last, v1
	if v1 = haystack != 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+67957, __ccgo_ts+67966, int32(133), uintptr(unsafe.Pointer(&__func__)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if v1 = needle != 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+68020, __ccgo_ts+67966, int32(134), uintptr(unsafe.Pointer(&__func__)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if v1 = nlen > uint64(1); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+68027, __ccgo_ts+67966, int32(135), uintptr(unsafe.Pointer(&__func__)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	last = haystack + uintptr(hlen) - uintptr(nlen)
	cur = haystack
	for {
		if !(cur <= last) {
			break
		}
		if int32(**(**int8)(__ccgo_up(cur))) == int32(**(**int8)(__ccgo_up(needle))) && libc.Xmemcmp(tls, cur, needle, nlen) == 0 {
			return cur
		}
		goto _4
	_4:
		;
		cur = cur + 1
	}
	return libc.UintptrFromInt32(0)
}

var __func__ = [10]int8{'m', 'y', '_', 'm', 'e', 'm', 'm', 'e', 'm'}

// C documentation
//
//	/** Find largest string containing certain characters.
//	 *
//	 * C Standard library 'strspn' only works for 'c-strings' (null terminated)
//	 * This works on arbitrary length.
//	 *
//	 * Performance notes:
//	 *   not critical
//	 *
//	 * Porting notes:
//	 *   if accept is 'ABC', then this function would be similar to
//	 *   a_regexp.match(a_str, '[ABC]*'),
//	 */
func strlenspn(tls *libc.TLS, s uintptr, len1 size_t, accept uintptr) (r size_t) {
	var i size_t
	_ = i
	i = uint64(0)
	for {
		if !(i < len1) {
			break
		}
		/* likely we can do better by inlining this function
		 * but this works for now
		 */
		if libc.Xstrchr(tls, accept, int32(**(**int8)(__ccgo_up(s + uintptr(i))))) == libc.UintptrFromInt32(0) {
			return i
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return len1
}

func strlencspn(tls *libc.TLS, s uintptr, len1 size_t, accept uintptr) (r size_t) {
	var i size_t
	_ = i
	i = uint64(0)
	for {
		if !(i < len1) {
			break
		}
		/* likely we can do better by inlining this function
		 * but this works for now
		 */
		if libc.Xstrchr(tls, accept, int32(**(**int8)(__ccgo_up(s + uintptr(i))))) != libc.UintptrFromInt32(0) {
			return i
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return len1
}

func char_is_white(tls *libc.TLS, ch int8) (r int32) {
	/* ' '  space is 0x32
	   '\t  0x09 \011 horizontal tab
	   '\n' 0x0a \012 new line
	   '\v' 0x0b \013 vertical tab
	   '\f' 0x0c \014 new page
	   '\r' 0x0d \015 carriage return
	        0x00 \000 null (oracle)
	        0xa0 \240 is Latin-1
	*/
	return libc.BoolInt32(libc.Xstrchr(tls, __ccgo_ts+68036, int32(ch)) != libc.UintptrFromInt32(0))
}

// C documentation
//
//	/* DANGER DANGER
//	 * This is -very specialized function-
//	 *
//	 * this compares a ALL_UPPER CASE C STRING
//	 * with a *arbitrary memory* + length
//	 *
//	 * Sane people would just make a copy, up-case
//	 * and use a hash table.
//	 *
//	 * Required since libc version uses the current locale
//	 * and is much slower.
//	 */
func cstrcasecmp(tls *libc.TLS, a uintptr, b uintptr, n size_t) (r int32) {
	var cb int8
	var v2 int32
	_, _ = cb, v2
	for {
		if !(n > uint64(0)) {
			break
		}
		cb = **(**int8)(__ccgo_up(b))
		if int32(cb) >= int32('a') && int32(cb) <= int32('z') {
			cb = int8(int32(cb) - libc.Int32FromInt32(0x20))
		}
		if int32(**(**int8)(__ccgo_up(a))) != int32(cb) {
			return int32(**(**int8)(__ccgo_up(a))) - int32(cb)
		} else {
			if int32(**(**int8)(__ccgo_up(a))) == int32('\000') {
				return -int32(1)
			}
		}
		goto _1
	_1:
		;
		a = a + 1
		b = b + 1
		n = n - 1
	}
	if int32(**(**int8)(__ccgo_up(a))) == 0 {
		v2 = 0
	} else {
		v2 = int32(1)
	}
	return v2
}

// C documentation
//
//	/**
//	 * Case sensitive string compare.
//	 *  Here only to make code more readable
//	 */
func streq(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	return libc.BoolInt32(libc.Xstrcmp(tls, a, b) == 0)
}

/**
 *
 *
 *
 * Porting Notes:
 *  given a mapping/hash of string to char
 *  this is just
 *    typecode = mapping[key.upper()]
 */

func bsearch_keyword_type(tls *libc.TLS, key uintptr, len1 size_t, keywords uintptr, numb size_t) (r int8) {
	var left, pos, right size_t
	_, _, _ = left, pos, right
	left = uint64(0)
	right = numb - uint64(1)
	for left < right {
		pos = (left + right) >> int32(1)
		/* arg0 = upper case only, arg1 = mixed case */
		if cstrcasecmp(tls, (**(**keyword_t)(__ccgo_up(keywords + uintptr(pos)*16))).Fword, key, len1) < 0 {
			left = pos + uint64(1)
		} else {
			right = pos
		}
	}
	if left == right && cstrcasecmp(tls, (**(**keyword_t)(__ccgo_up(keywords + uintptr(left)*16))).Fword, key, len1) == 0 {
		return (**(**keyword_t)(__ccgo_up(keywords + uintptr(left)*16))).Ftype1
	} else {
		return int8('\000')
	}
	return r
}

func is_keyword(tls *libc.TLS, key uintptr, len1 size_t) (r int8) {
	return bsearch_keyword_type(tls, key, len1, uintptr(unsafe.Pointer(&sql_keywords)), sql_keywords_sz)
}

/* st_token methods
 *
 * The following functions manipulates the stoken_t type
 *
 *
 */

func st_clear(tls *libc.TLS, st uintptr) {
	libc.Xmemset(tls, st, 0, uint64(56))
}

func st_assign_char(tls *libc.TLS, st uintptr, stype int8, pos size_t, len1 size_t, value int8) {
	/* done to eliminate unused warning */
	_ = len1
	(*stoken_t)(unsafe.Pointer(st)).Ftype1 = stype
	(*stoken_t)(unsafe.Pointer(st)).Fpos = pos
	(*stoken_t)(unsafe.Pointer(st)).Flen1 = uint64(1)
	**(**int8)(__ccgo_up(st + 23)) = value
	**(**int8)(__ccgo_up(st + 23 + 1)) = int8('\000')
}

func st_assign(tls *libc.TLS, st uintptr, stype int8, pos size_t, len1 size_t, value uintptr) {
	var MSIZE, last size_t
	var v1 uint64
	_, _, _ = MSIZE, last, v1
	MSIZE = uint64(32)
	if len1 < MSIZE {
		v1 = len1
	} else {
		v1 = MSIZE - uint64(1)
	}
	last = v1
	(*stoken_t)(unsafe.Pointer(st)).Ftype1 = stype
	(*stoken_t)(unsafe.Pointer(st)).Fpos = pos
	(*stoken_t)(unsafe.Pointer(st)).Flen1 = last
	libc.Xmemcpy(tls, st+23, value, last)
	**(**int8)(__ccgo_up(st + 23 + uintptr(last))) = int8('\000')
}

func st_copy(tls *libc.TLS, dest uintptr, src uintptr) {
	libc.Xmemcpy(tls, dest, src, uint64(56))
}

func st_is_arithmetic_op(tls *libc.TLS, st uintptr) (r int32) {
	var ch int8
	_ = ch
	ch = **(**int8)(__ccgo_up(st + 23))
	return libc.BoolInt32(int32((*stoken_t)(unsafe.Pointer(st)).Ftype1) == int32(TYPE_OPERATOR) && (*stoken_t)(unsafe.Pointer(st)).Flen1 == uint64(1) && (int32(ch) == int32('*') || int32(ch) == int32('/') || int32(ch) == int32('-') || int32(ch) == int32('+') || int32(ch) == int32('%')))
}

func st_is_unary_op(tls *libc.TLS, st uintptr) (r int32) {
	var len1 size_t
	var str uintptr
	_, _ = len1, str
	str = st + 23
	len1 = (*stoken_t)(unsafe.Pointer(st)).Flen1
	if int32((*stoken_t)(unsafe.Pointer(st)).Ftype1) != int32(TYPE_OPERATOR) {
		return FALSE
	}
	switch len1 {
	case uint64(1):
		return libc.BoolInt32(int32(**(**int8)(__ccgo_up(str))) == int32('+') || int32(**(**int8)(__ccgo_up(str))) == int32('-') || int32(**(**int8)(__ccgo_up(str))) == int32('!') || int32(**(**int8)(__ccgo_up(str))) == int32('~'))
	case uint64(2):
		return libc.BoolInt32(int32(**(**int8)(__ccgo_up(str))) == int32('!') && int32(**(**int8)(__ccgo_up(str + 1))) == int32('!'))
	case uint64(3):
		return libc.BoolInt32(cstrcasecmp(tls, __ccgo_ts+63445, str, uint64(3)) == 0)
	default:
		return FALSE
	}
	return r
}

/* Parsers
 *
 *
 */

func parse_white(tls *libc.TLS, sf uintptr) (r size_t) {
	return (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos + uint64(1)
}

func parse_operator1(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos size_t
	_, _ = cs, pos
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_OPERATOR), pos, uint64(1), **(**int8)(__ccgo_up(cs + uintptr(pos))))
	return pos + uint64(1)
}

func parse_other(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos size_t
	_, _ = cs, pos
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_UNKNOWN), pos, uint64(1), **(**int8)(__ccgo_up(cs + uintptr(pos))))
	return pos + uint64(1)
}

func parse_char(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos size_t
	_, _ = cs, pos
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, **(**int8)(__ccgo_up(cs + uintptr(pos))), pos, uint64(1), **(**int8)(__ccgo_up(cs + uintptr(pos))))
	return pos + uint64(1)
}

func parse_eol_comment(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs, endpos uintptr
	var pos, slen size_t
	_, _, _, _ = cs, endpos, pos, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	endpos = libc.Xmemchr(tls, cs+uintptr(pos), int32('\n'), slen-pos)
	if endpos == libc.UintptrFromInt32(0) {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_COMMENT), pos, slen-pos, cs+uintptr(pos))
		return slen
	} else {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_COMMENT), pos, libc.Uint64FromInt64(int64(endpos)-int64(cs))-pos, cs+uintptr(pos))
		return libc.Uint64FromInt64(int64(endpos) - int64(cs) + libc.Int64FromInt32(1))
	}
	return r
}

// C documentation
//
//	/** In ANSI mode, hash is an operator
//	 *  In MYSQL mode, it's a EOL comment like '--'
//	 */
func parse_hash(tls *libc.TLS, sf uintptr) (r size_t) {
	**(**int32)(__ccgo_up(sf + 528)) += int32(1)
	if (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fflags&int32(FLAG_SQL_MYSQL) != 0 {
		**(**int32)(__ccgo_up(sf + 528)) += int32(1)
		return parse_eol_comment(tls, sf)
	} else {
		st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_OPERATOR), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos, uint64(1), int8('#'))
		return (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos + uint64(1)
	}
	return r
}

func parse_dash(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen size_t
	_, _, _ = cs, pos, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	/*
	 * five cases
	 * 1) --[white]  this is always a SQL comment
	 * 2) --[EOF]    this is a comment
	 * 3) --[notwhite] in MySQL this is NOT a comment but two unary operators
	 * 4) --[notwhite] everyone else thinks this is a comment
	 * 5) -[not dash]  '-' is a unary operator
	 */
	if pos+uint64(2) < slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('-') && char_is_white(tls, **(**int8)(__ccgo_up(cs + uintptr(pos+uint64(2))))) != 0 {
		return parse_eol_comment(tls, sf)
	} else {
		if pos+uint64(2) == slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('-') {
			return parse_eol_comment(tls, sf)
		} else {
			if pos+uint64(1) < slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('-') && (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fflags&int32(FLAG_SQL_ANSI) != 0 {
				/* --[not-white] not-white case:
				 *
				 */
				**(**int32)(__ccgo_up(sf + 520)) += int32(1)
				return parse_eol_comment(tls, sf)
			} else {
				st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_OPERATOR), pos, uint64(1), int8('-'))
				return pos + uint64(1)
			}
		}
	}
	return r
}

// C documentation
//
//	/** This detects MySQL comments, comments that
//	 * start with /x!   We just ban these now but
//	 * previously we attempted to parse the inside
//	 *
//	 * For reference:
//	 * the form of /x![anything]x/ or /x!12345[anything] x/
//	 *
//	 * Mysql 3 (maybe 4), allowed this:
//	 *    /x!0selectx/ 1;
//	 * where 0 could be any number.
//	 *
//	 * The last version of MySQL 3 was in 2003.
//
//	 * It is unclear if the MySQL 3 syntax was allowed
//	 * in MySQL 4.  The last version of MySQL 4 was in 2008
//	 *
//	 */
func is_mysql_comment(tls *libc.TLS, cs uintptr, len1 size_t, pos size_t) (r size_t) {
	/* so far...
	 * cs[pos] == '/' && cs[pos+1] == '*'
	 */
	if pos+uint64(2) >= len1 {
		/* not a mysql comment */
		return uint64(0)
	}
	if int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(2))))) != int32('!') {
		/* not a mysql comment */
		return uint64(0)
	}
	/*
	 * this is a mysql comment
	 *  got "/x!"
	 */
	return uint64(1)
}

func parse_slash(tls *libc.TLS, sf uintptr) (r size_t) {
	var clen, pos, pos1, slen size_t
	var cs, cur, ptr uintptr
	var ctype int8
	_, _, _, _, _, _, _, _ = clen, cs, ctype, cur, pos, pos1, ptr, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	cur = cs + uintptr(pos)
	ctype = int8(TYPE_COMMENT)
	pos1 = pos + uint64(1)
	if pos1 == slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos1)))) != int32('*') {
		return parse_operator1(tls, sf)
	}
	/*
	 * skip over initial '/x'
	 */
	ptr = memchr2(tls, cur+uintptr(2), slen-(pos+uint64(2)), int8('*'), int8('/'))
	/*
	 * (ptr == NULL) causes false positive in cppcheck 1.61
	 * casting to type seems to fix it
	 */
	if ptr == libc.UintptrFromInt32(0) {
		/* till end of line */
		clen = slen - pos
	} else {
		clen = libc.Uint64FromInt64(int64(ptr+libc.UintptrFromInt32(2)) - int64(cur))
	}
	/*
	 * postgresql allows nested comments which makes
	 * this is incompatible with parsing so
	 * if we find a '/x' inside the coment, then
	 * make a new token.
	 *
	 * Also, Mysql's "conditional" comments for version
	 *  are an automatic black ban!
	 */
	if memchr2(tls, cur+uintptr(2), libc.Uint64FromInt64(int64(ptr)-int64(cur+libc.UintptrFromInt32(1))), int8('/'), int8('*')) != libc.UintptrFromInt32(0) {
		ctype = int8(TYPE_EVIL)
	} else {
		if is_mysql_comment(tls, cs, slen, pos) != 0 {
			ctype = int8(TYPE_EVIL)
		}
	}
	st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, ctype, pos, clen, cs+uintptr(pos))
	return pos + clen
}

func parse_backslash(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen size_t
	_, _, _ = cs, pos, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	/*
	 * Weird MySQL alias for NULL, "\N" (capital N only)
	 */
	if pos+uint64(1) < slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('N') {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_NUMBER), pos, uint64(2), cs+uintptr(pos))
		return pos + uint64(2)
	} else {
		st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BACKSLASH), pos, uint64(1), **(**int8)(__ccgo_up(cs + uintptr(pos))))
		return pos + uint64(1)
	}
	return r
}

func parse_operator2(tls *libc.TLS, sf uintptr) (r size_t) {
	var ch int8
	var cs uintptr
	var pos, slen size_t
	_, _, _, _ = ch, cs, pos, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	if pos+uint64(1) >= slen {
		return parse_operator1(tls, sf)
	}
	if pos+uint64(2) < slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('<') && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('=') && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(2))))) == int32('>') {
		/*
		 * special 3-char operator
		 */
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_OPERATOR), pos, uint64(3), cs+uintptr(pos))
		return pos + uint64(3)
	}
	ch = (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup})))(tls, sf, int32(LOOKUP_OPERATOR), cs+uintptr(pos), uint64(2))
	if int32(ch) != int32('\000') {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, ch, pos, uint64(2), cs+uintptr(pos))
		return pos + uint64(2)
	}
	/*
	 * not an operator.. what to do with the two
	 * characters we got?
	 */
	if int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32(':') {
		/* ':' is not an operator */
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_COLON), pos, uint64(1), cs+uintptr(pos))
		return pos + uint64(1)
	} else {
		/*
		 * must be a single char operator
		 */
		return parse_operator1(tls, sf)
	}
	return r
}

// C documentation
//
//	/*
//	 * Ok!   "  \"   "  one backslash = escaped!
//	 *       " \\"   "  two backslash = not escaped!
//	 *       "\\\"   "  three backslash = escaped!
//	 */
func is_backslash_escaped(tls *libc.TLS, end uintptr, start uintptr) (r int32) {
	var ptr uintptr
	_ = ptr
	ptr = end
	for {
		if !(ptr >= start) {
			break
		}
		if int32(**(**int8)(__ccgo_up(ptr))) != int32('\\') {
			break
		}
		goto _1
	_1:
		;
		ptr = ptr - 1
	}
	/* if number of backslashes is odd, it is escaped */
	return int32((int64(end) - int64(ptr)) & int64(1))
}

func is_double_delim_escaped(tls *libc.TLS, cur uintptr, end uintptr) (r size_t) {
	return libc.BoolUint64(cur+uintptr(1) < end && int32(**(**int8)(__ccgo_up(cur + libc.UintptrFromInt32(1)))) == int32(**(**int8)(__ccgo_up(cur))))
}

// C documentation
//
//	/* Look forward for doubling of delimiter
//	 *
//	 * case 'foo''bar' --> foo''bar
//	 *
//	 * ending quote isn't duplicated (i.e. escaped)
//	 * since it's the wrong char or EOL
//	 *
//	 */
func parse_string_core(tls *libc.TLS, cs uintptr, len1 size_t, pos size_t, st uintptr, delim int8, offset size_t) (r size_t) {
	var qpos uintptr
	_ = qpos
	/*
	 * offset is to skip the perhaps first quote char
	 */
	qpos = libc.Xmemchr(tls, cs+uintptr(pos)+uintptr(offset), int32(delim), len1-pos-offset)
	/*
	 * then keep string open/close info
	 */
	if offset > uint64(0) {
		/*
		 * this is real quote
		 */
		(*stoken_t)(unsafe.Pointer(st)).Fstr_open = delim
	} else {
		/*
		 * this was a simulated quote
		 */
		(*stoken_t)(unsafe.Pointer(st)).Fstr_open = int8('\000')
	}
	for int32(TRUE) != 0 {
		if qpos == libc.UintptrFromInt32(0) {
			/*
			 * string ended with no trailing quote
			 * assign what we have
			 */
			st_assign(tls, st, int8(TYPE_STRING), pos+offset, len1-pos-offset, cs+uintptr(pos)+uintptr(offset))
			(*stoken_t)(unsafe.Pointer(st)).Fstr_close = int8('\000')
			return len1
		} else {
			if is_backslash_escaped(tls, qpos-uintptr(1), cs+uintptr(pos)+uintptr(offset)) != 0 {
				/* keep going, move ahead one character */
				qpos = libc.Xmemchr(tls, qpos+libc.UintptrFromInt32(1), int32(delim), libc.Uint64FromInt64(int64(cs+uintptr(len1))-int64(qpos+libc.UintptrFromInt32(1))))
				continue
			} else {
				if is_double_delim_escaped(tls, qpos, cs+uintptr(len1)) != 0 {
					/* keep going, move ahead two characters */
					qpos = libc.Xmemchr(tls, qpos+libc.UintptrFromInt32(2), int32(delim), libc.Uint64FromInt64(int64(cs+uintptr(len1))-int64(qpos+libc.UintptrFromInt32(2))))
					continue
				} else {
					/* hey it's a normal string */
					st_assign(tls, st, int8(TYPE_STRING), pos+offset, libc.Uint64FromInt64(int64(qpos)-int64(cs+uintptr(pos)+uintptr(offset))), cs+uintptr(pos)+uintptr(offset))
					(*stoken_t)(unsafe.Pointer(st)).Fstr_close = delim
					return libc.Uint64FromInt64(int64(qpos) - int64(cs) + libc.Int64FromInt32(1))
				}
			}
		}
	}
	return r
}

// C documentation
//
//	/**
//	 * Used when first char is a ' or "
//	 */
func parse_string(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen size_t
	_, _, _ = cs, pos, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	/*
	 * assert cs[pos] == single or double quote
	 */
	return parse_string_core(tls, cs, slen, pos, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, **(**int8)(__ccgo_up(cs + uintptr(pos))), uint64(1))
}

// C documentation
//
//	/**
//	 * Used when first char is:
//	 *    N or n:  mysql "National Character set"
//	 *    E     :  psql  "Escaped String"
//	 */
func parse_estring(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen size_t
	_, _, _ = cs, pos, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	if pos+uint64(2) >= slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) != int32('\'') {
		return parse_word(tls, sf)
	}
	return parse_string_core(tls, cs, slen, pos, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8('\''), uint64(2))
}

func parse_ustring(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen size_t
	_, _, _ = cs, pos, slen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	if pos+uint64(2) < slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('&') && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(2))))) == int32('\'') {
		**(**size_t)(__ccgo_up(sf + 40)) += uint64(2)
		pos = parse_string(tls, sf)
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_open = int8('u')
		if int32((*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close) == int32('\'') {
			(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close = int8('u')
		}
		return pos
	} else {
		return parse_word(tls, sf)
	}
	return r
}

func parse_qstring_core(tls *libc.TLS, sf uintptr, offset size_t) (r size_t) {
	var ch int8
	var cs, strend uintptr
	var pos, slen size_t
	_, _, _, _, _ = ch, cs, pos, slen, strend
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos + offset
	/* if we are already at end of string..
	   if current char is not q or Q
	   if we don't have 2 more chars
	   if char2 != a single quote
	   then, just treat as word
	*/
	if pos >= slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) != int32('q') && int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) != int32('Q') || pos+uint64(2) >= slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) != int32('\'') {
		return parse_word(tls, sf)
	}
	ch = **(**int8)(__ccgo_up(cs + uintptr(pos+uint64(2))))
	/* the ch > 127 is un-needed since
	 * we assume char is signed
	 */
	if int32(ch) < int32(33) {
		return parse_word(tls, sf)
	}
	switch int32(ch) {
	case int32('('):
		ch = int8(')')
	case int32('['):
		ch = int8(']')
	case int32('{'):
		ch = int8('}')
	case int32('<'):
		ch = int8('>')
		break
	}
	strend = memchr2(tls, cs+uintptr(pos)+uintptr(3), slen-pos-uint64(3), ch, int8('\''))
	if strend == libc.UintptrFromInt32(0) {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_STRING), pos+uint64(3), slen-pos-uint64(3), cs+uintptr(pos)+uintptr(3))
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_open = int8('q')
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close = int8('\000')
		return slen
	} else {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_STRING), pos+uint64(3), libc.Uint64FromInt64(int64(strend)-int64(cs))-pos-uint64(3), cs+uintptr(pos)+uintptr(3))
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_open = int8('q')
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close = int8('q')
		return libc.Uint64FromInt64(int64(strend) - int64(cs) + libc.Int64FromInt32(2))
	}
	return r
}

// C documentation
//
//	/*
//	 * Oracle's q string
//	 */
func parse_qstring(tls *libc.TLS, sf uintptr) (r size_t) {
	return parse_qstring_core(tls, sf, uint64(0))
}

// C documentation
//
//	/*
//	 * mysql's N'STRING' or
//	 * ...  Oracle's nq string
//	 */
func parse_nqstring(tls *libc.TLS, sf uintptr) (r size_t) {
	var pos, slen size_t
	_, _ = pos, slen
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	if pos+uint64(2) < slen && int32(**(**int8)(__ccgo_up((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs + uintptr(pos+uint64(1))))) == int32('\'') {
		return parse_estring(tls, sf)
	}
	return parse_qstring_core(tls, sf, uint64(1))
}

// C documentation
//
//	/*
//	 * binary literal string
//	 * re: [bB]'[01]*'
//	 */
func parse_bstring(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen, wlen size_t
	_, _, _, _ = cs, pos, slen, wlen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	/* need at least 2 more characters
	 * if next char isn't a single quote, then
	 * continue as normal word
	 */
	if pos+uint64(2) >= slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) != int32('\'') {
		return parse_word(tls, sf)
	}
	wlen = strlenspn(tls, cs+uintptr(pos)+uintptr(2), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen-pos-uint64(2), __ccgo_ts+68045)
	if pos+uint64(2)+wlen >= slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(2)+wlen)))) != int32('\'') {
		return parse_word(tls, sf)
	}
	st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_NUMBER), pos, wlen+uint64(3), cs+uintptr(pos))
	return pos + uint64(2) + wlen + uint64(1)
}

// C documentation
//
//	/*
//	 * hex literal string
//	 * re: [xX]'[0123456789abcdefABCDEF]*'
//	 * mysql has requirement of having EVEN number of chars,
//	 *  but pgsql does not
//	 */
func parse_xstring(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen, wlen size_t
	_, _, _, _ = cs, pos, slen, wlen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	/* need at least 2 more characters
	 * if next char isn't a single quote, then
	 * continue as normal word
	 */
	if pos+uint64(2) >= slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) != int32('\'') {
		return parse_word(tls, sf)
	}
	wlen = strlenspn(tls, cs+uintptr(pos)+uintptr(2), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen-pos-uint64(2), __ccgo_ts+68048)
	if pos+uint64(2)+wlen >= slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(2)+wlen)))) != int32('\'') {
		return parse_word(tls, sf)
	}
	st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_NUMBER), pos, wlen+uint64(3), cs+uintptr(pos))
	return pos + uint64(2) + wlen + uint64(1)
}

// C documentation
//
//	/**
//	 * This handles MS SQLSERVER bracket words
//	 * http://stackoverflow.com/questions/3551284/sql-serverwhat-do-brackets-mean-around-column-name
//	 *
//	 */
func parse_bword(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs, endptr uintptr
	var pos size_t
	_, _, _ = cs, endptr, pos
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	endptr = libc.Xmemchr(tls, cs+uintptr(pos), int32(']'), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen-pos)
	if endptr == libc.UintptrFromInt32(0) {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), pos, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen-pos, cs+uintptr(pos))
		return (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	} else {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), pos, libc.Uint64FromInt64(int64(endptr)-int64(cs))-pos+uint64(1), cs+uintptr(pos))
		return libc.Uint64FromInt64(int64(endptr) - int64(cs) + libc.Int64FromInt32(1))
	}
	return r
}

func parse_word(tls *libc.TLS, sf uintptr) (r size_t) {
	var ch, delim int8
	var cs uintptr
	var i, pos, wlen size_t
	_, _, _, _, _, _ = ch, cs, delim, i, pos, wlen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	wlen = strlencspn(tls, cs+uintptr(pos), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen-pos, __ccgo_ts+68071)
	st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), pos, wlen, cs+uintptr(pos))
	/* now we need to look inside what we good for "." and "`"
	 * and see if what is before is a keyword or not
	 */
	i = uint64(0)
	for {
		if !(i < (*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Flen1) {
			break
		}
		delim = **(**int8)(__ccgo_up((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent + 23 + uintptr(i)))
		if int32(delim) == int32('.') || int32(delim) == int32('`') {
			ch = (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup})))(tls, sf, int32(LOOKUP_WORD), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent+23, i)
			if int32(ch) != int32(TYPE_NONE) && int32(ch) != int32(TYPE_BAREWORD) {
				/* needed for swig */
				st_clear(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)
				/*
				 * we got something like "SELECT.1"
				 * or SELECT`column`
				 */
				st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, ch, pos, i, cs+uintptr(pos))
				return pos + i
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	/*
	 * do normal lookup with word including '.'
	 */
	if wlen < uint64(32) {
		ch = (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup})))(tls, sf, int32(LOOKUP_WORD), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent+23, wlen)
		if int32(ch) == int32('\000') {
			ch = int8(TYPE_BAREWORD)
		}
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1 = ch
	}
	return pos + wlen
}

// C documentation
//
//	/* MySQL backticks are a cross between string and
//	 * and a bare word.
//	 *
//	 */
func parse_tick(tls *libc.TLS, sf uintptr) (r size_t) {
	var ch int8
	var pos size_t
	_, _ = ch, pos
	pos = parse_string_core(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8('`'), uint64(1))
	/* we could check to see if start and end of
	 * of string are both "`", i.e. make sure we have
	 * matching set.  `foo` vs. `foo
	 * but I don't think it matters much
	 */
	/* check value of string to see if it's a keyword,
	 * function, operator, etc
	 */
	ch = (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup})))(tls, sf, int32(LOOKUP_WORD), (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent+23, (*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Flen1)
	if int32(ch) == int32(TYPE_FUNCTION) {
		/* if it's a function, then convert token */
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1 = int8(TYPE_FUNCTION)
	} else {
		/* otherwise it's a 'n' type -- mysql treats
		 * everything as a bare word
		 */
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1 = int8(TYPE_BAREWORD)
	}
	return pos
}

func parse_var(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs uintptr
	var pos, slen, xlen size_t
	_, _, _, _ = cs, pos, slen, xlen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos + uint64(1)
	/*
	 * var_count is only used to reconstruct
	 * the input.  It counts the number of '@'
	 * seen 0 in the case of NULL, 1 or 2
	 */
	/*
	 * move past optional other '@'
	 */
	if pos < slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('@') {
		pos = pos + uint64(1)
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fcount = int32(2)
	} else {
		(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fcount = int32(1)
	}
	/*
	 * MySQL allows @@`version`
	 */
	if pos < slen {
		if int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('`') {
			(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos = pos
			pos = parse_tick(tls, sf)
			(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1 = int8(TYPE_VARIABLE)
			return pos
		} else {
			if int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('\'') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('"') {
				(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos = pos
				pos = parse_string(tls, sf)
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1 = int8(TYPE_VARIABLE)
				return pos
			}
		}
	}
	xlen = strlencspn(tls, cs+uintptr(pos), slen-pos, __ccgo_ts+68108)
	if xlen == uint64(0) {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_VARIABLE), pos, uint64(0), cs+uintptr(pos))
		return pos
	} else {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_VARIABLE), pos, xlen, cs+uintptr(pos))
		return pos + xlen
	}
	return r
}

func parse_money(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs, strend uintptr
	var pos, slen, xlen size_t
	_, _, _, _, _ = cs, pos, slen, strend, xlen
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	if pos+uint64(1) == slen {
		/* end of line */
		st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), pos, uint64(1), int8('$'))
		return slen
	}
	/*
	 * $1,000.00 or $1.000,00 ok!
	 * This also parses $....,,,111 but that's ok
	 */
	xlen = strlenspn(tls, cs+uintptr(pos)+uintptr(1), slen-pos-uint64(1), __ccgo_ts+68141)
	if xlen == uint64(0) {
		if int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('$') {
			/* we have $$ .. find ending $$ and make string */
			strend = memchr2(tls, cs+uintptr(pos)+uintptr(2), slen-pos-uint64(2), int8('$'), int8('$'))
			if strend == libc.UintptrFromInt32(0) {
				/* fell off edge */
				st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_STRING), pos+uint64(2), slen-(pos+uint64(2)), cs+uintptr(pos)+uintptr(2))
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_open = int8('$')
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close = int8('\000')
				return slen
			} else {
				st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_STRING), pos+uint64(2), libc.Uint64FromInt64(int64(strend)-int64(cs+uintptr(pos)+libc.UintptrFromInt32(2))), cs+uintptr(pos)+uintptr(2))
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_open = int8('$')
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close = int8('$')
				return libc.Uint64FromInt64(int64(strend) - int64(cs) + libc.Int64FromInt32(2))
			}
		} else {
			/* ok it's not a number or '$$', but maybe it's pgsql "$ quoted strings" */
			xlen = strlenspn(tls, cs+uintptr(pos)+uintptr(1), slen-pos-uint64(1), __ccgo_ts+68154)
			if xlen == uint64(0) {
				/* hmm it's "$" _something_ .. just add $ and keep going*/
				st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), pos, uint64(1), int8('$'))
				return pos + uint64(1)
			}
			/* we have $foobar????? */
			/* is it $foobar$ */
			if pos+xlen+uint64(1) == slen || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+xlen+uint64(1))))) != int32('$') {
				/* not $foobar$, or fell off edge */
				st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), pos, uint64(1), int8('$'))
				return pos + uint64(1)
			}
			/* we have $foobar$ ... find it again */
			strend = my_memmem(tls, cs+uintptr(xlen)+uintptr(2), slen-(pos+xlen+uint64(2)), cs+uintptr(pos), xlen+uint64(2))
			if strend == libc.UintptrFromInt32(0) || libc.Uint64FromInt64(int64(strend)-int64(cs)) < pos+xlen+uint64(2) {
				/* fell off edge */
				st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_STRING), pos+xlen+uint64(2), slen-pos-xlen-uint64(2), cs+uintptr(pos)+uintptr(xlen)+uintptr(2))
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_open = int8('$')
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close = int8('\000')
				return slen
			} else {
				/* got one */
				st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_STRING), pos+xlen+uint64(2), libc.Uint64FromInt64(int64(strend)-int64(cs+uintptr(pos)+uintptr(xlen)+libc.UintptrFromInt32(2))), cs+uintptr(pos)+uintptr(xlen)+uintptr(2))
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_open = int8('$')
				(*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Fstr_close = int8('$')
				return libc.Uint64FromInt64(int64(strend+uintptr(xlen)+libc.UintptrFromInt32(2)) - int64(cs))
			}
		}
	} else {
		if xlen == uint64(1) && int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('.') {
			/* $. should parsed as a word */
			return parse_word(tls, sf)
		} else {
			st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_NUMBER), pos, uint64(1)+xlen, cs+uintptr(pos))
			return pos + uint64(1) + xlen
		}
	}
	return r
}

func parse_number(tls *libc.TLS, sf uintptr) (r size_t) {
	var cs, digits uintptr
	var have_e, have_exp int32
	var pos, slen, start, xlen size_t
	_, _, _, _, _, _, _, _ = cs, digits, have_e, have_exp, pos, slen, start, xlen
	digits = libc.UintptrFromInt32(0)
	cs = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	pos = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fpos
	have_e = 0
	have_exp = 0
	/* cs[pos] == '0' has 1/10 chance of being true,
	 * while pos+1< slen is almost always true
	 */
	if int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('0') && pos+uint64(1) < slen {
		if int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('X') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('x') {
			digits = __ccgo_ts + 68048
		} else {
			if int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('B') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('b') {
				digits = __ccgo_ts + 68045
			}
		}
		if digits != 0 {
			xlen = strlenspn(tls, cs+uintptr(pos)+uintptr(2), slen-pos-uint64(2), digits)
			if xlen == uint64(0) {
				st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), pos, uint64(2), cs+uintptr(pos))
				return pos + uint64(2)
			} else {
				st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_NUMBER), pos, uint64(2)+xlen, cs+uintptr(pos))
				return pos + uint64(2) + xlen
			}
		}
	}
	start = pos
	for pos < slen && libc.Uint32FromInt32(int32(**(**int8)(__ccgo_up(cs + uintptr(pos))))-libc.Int32FromUint8('0')) <= uint32(9) {
		pos = pos + uint64(1)
	}
	if pos < slen && int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('.') {
		pos = pos + uint64(1)
		for pos < slen && libc.Uint32FromInt32(int32(**(**int8)(__ccgo_up(cs + uintptr(pos))))-libc.Int32FromUint8('0')) <= uint32(9) {
			pos = pos + uint64(1)
		}
		if pos-start == uint64(1) {
			/* only one character read so far */
			st_assign_char(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_DOT), start, uint64(1), int8('.'))
			return pos
		}
	}
	if pos < slen {
		if int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('E') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('e') {
			have_e = int32(1)
			pos = pos + uint64(1)
			if pos < slen && (int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('+') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('-')) {
				pos = pos + uint64(1)
			}
			for pos < slen && libc.Uint32FromInt32(int32(**(**int8)(__ccgo_up(cs + uintptr(pos))))-libc.Int32FromUint8('0')) <= uint32(9) {
				have_exp = int32(1)
				pos = pos + uint64(1)
			}
		}
	}
	/* oracle's ending float or double suffix
	 * http://docs.oracle.com/cd/B19306_01/server.102/b14200/sql_elements003.htm#i139891
	 */
	if pos < slen && (int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('d') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('D') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('f') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos)))) == int32('F')) {
		if pos+uint64(1) == slen {
			/* line ends evaluate "... 1.2f$" as '1.2f' */
			pos = pos + uint64(1)
		} else {
			if char_is_white(tls, **(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) != 0 || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32(';') {
				/*
				 * easy case, evaluate "... 1.2f ... as '1.2f'
				 */
				pos = pos + uint64(1)
			} else {
				if int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('u') || int32(**(**int8)(__ccgo_up(cs + uintptr(pos+uint64(1))))) == int32('U') {
					/*
					 * a bit of a hack but makes '1fUNION' parse as '1f UNION'
					 */
					pos = pos + uint64(1)
				} else {
					/* it's like "123FROM" */
					/* parse as "123" only */
				}
			}
		}
	}
	if have_e == int32(1) && have_exp == 0 {
		/* very special form of
		 * "1234.e"
		 * "10.10E"
		 * ".E"
		 * this is a WORD not a number!! */
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_BAREWORD), start, pos-start, cs+uintptr(start))
	} else {
		st_assign(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent, int8(TYPE_NUMBER), start, pos-start, cs+uintptr(start))
	}
	return pos
}

// C documentation
//
//	/*
//	 * API to return version.  This allows us to increment the version
//	 * without having to regenerated the SWIG (or other binding) in minor
//	 * releases.
//	 */
func libinjection_version(tls *libc.TLS) (r uintptr) {
	return __ccgo_ts + 68207
}

func libinjection_sqli_tokenize(tls *libc.TLS, sf uintptr) (r int32) {
	var ch uint8
	var current, pos, s uintptr
	var fnptr pt2Function
	var slen size_t
	_, _, _, _, _, _ = ch, current, fnptr, pos, s, slen
	pos = sf + 40
	current = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent
	s = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen
	if slen == uint64(0) {
		return FALSE
	}
	st_clear(tls, current)
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent = current
	/*
	 * if we are at beginning of string
	 *  and in single-quote or double quote mode
	 *  then pretend the input starts with a quote
	 */
	if **(**size_t)(__ccgo_up(pos)) == uint64(0) && (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fflags&(int32(FLAG_QUOTE_SINGLE)|int32(FLAG_QUOTE_DOUBLE)) != 0 {
		**(**size_t)(__ccgo_up(pos)) = parse_string_core(tls, s, slen, uint64(0), current, flag2delim(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fflags), uint64(0))
		**(**int32)(__ccgo_up(sf + 536)) += int32(1)
		return int32(TRUE)
	}
	for **(**size_t)(__ccgo_up(pos)) < slen {
		/*
		 * get current character
		 */
		ch = libc.Uint8FromInt8(**(**int8)(__ccgo_up(s + uintptr(**(**size_t)(__ccgo_up(pos))))))
		/*
		 * look up the parser, and call it
		 *
		 * Porting Note: this is mapping of char to function
		 *   charparsers[ch]()
		 */
		fnptr = char_parse_map[ch]
		**(**size_t)(__ccgo_up(pos)) = (*(*func(*libc.TLS, uintptr) size_t)(unsafe.Pointer(&struct{ uintptr }{fnptr})))(tls, sf)
		/*
		 *
		 */
		if int32((*stoken_t)(unsafe.Pointer(current)).Ftype1) != int32('\000') {
			**(**int32)(__ccgo_up(sf + 536)) += int32(1)
			return int32(TRUE)
		}
	}
	return FALSE
}

func libinjection_sqli_init(tls *libc.TLS, sf uintptr, s uintptr, len1 size_t, flags int32) {
	if flags == 0 {
		flags = int32(FLAG_QUOTE_NONE) | int32(FLAG_SQL_ANSI)
	}
	libc.Xmemset(tls, sf, 0, uint64(544))
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs = s
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen = len1
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup = __ccgo_fp(libinjection_sqli_lookup_word)
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fuserdata = uintptr(0)
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fflags = flags
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent = sf + 48
}

func libinjection_sqli_reset(tls *libc.TLS, sf uintptr, flags int32) {
	var lookup ptr_lookup_fn
	var userdata uintptr
	_, _ = lookup, userdata
	userdata = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fuserdata
	lookup = (*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup
	if flags == 0 {
		flags = int32(FLAG_QUOTE_NONE) | int32(FLAG_SQL_ANSI)
	}
	libinjection_sqli_init(tls, sf, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fs, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fslen, flags)
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup = lookup
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fuserdata = userdata
}

type __ccgo_fp__Xlibinjection_sqli_callback_1 = func(*libc.TLS, uintptr, int32, uintptr, uint64) int8

func libinjection_sqli_callback(tls *libc.TLS, sf uintptr, __ccgo_fp_fn ptr_lookup_fn, userdata uintptr) {
	if __ccgo_fp_fn == libc.UintptrFromInt32(0) {
		(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup = __ccgo_fp(libinjection_sqli_lookup_word)
		(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fuserdata = libc.UintptrFromInt32(0)
	} else {
		(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup = __ccgo_fp_fn
		(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fuserdata = userdata
	}
}

// C documentation
//
//	/** See if two tokens can be merged since they are compound SQL phrases.
//	 *
//	 * This takes two tokens, and, if they are the right type,
//	 * merges their values together.  Then checks to see if the
//	 * new value is special using the PHRASES mapping.
//	 *
//	 * Example: "UNION" + "ALL" ==> "UNION ALL"
//	 *
//	 * C Security Notes: this is safe to use C-strings (null-terminated)
//	 *  since the types involved by definition do not have embedded nulls
//	 *  (e.g. there is no keyword with embedded null)
//	 *
//	 * Porting Notes: since this is C, it's oddly complicated.
//	 *  This is just:  multikeywords[token.value + ' ' + token2.value]
//	 *
//	 */
func syntax_merge_words(tls *libc.TLS, sf uintptr, a uintptr, b uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ch int8
	var sz1, sz2, sz3 size_t
	var _ /* tmp at bp+0 */ [32]int8
	_, _, _, _ = ch, sz1, sz2, sz3
	/* first token is of right type? */
	if !(int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_KEYWORD) || int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_BAREWORD) || int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_OPERATOR) || int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_UNION) || int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_FUNCTION) || int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_EXPRESSION) || int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_TSQL) || int32((*stoken_t)(unsafe.Pointer(a)).Ftype1) == int32(TYPE_SQLTYPE)) {
		return FALSE
	}
	if !(int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_KEYWORD) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_BAREWORD) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_OPERATOR) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_UNION) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_FUNCTION) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_EXPRESSION) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_TSQL) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_SQLTYPE) || int32((*stoken_t)(unsafe.Pointer(b)).Ftype1) == int32(TYPE_LOGIC_OPERATOR)) {
		return FALSE
	}
	sz1 = (*stoken_t)(unsafe.Pointer(a)).Flen1
	sz2 = (*stoken_t)(unsafe.Pointer(b)).Flen1
	sz3 = sz1 + sz2 + uint64(1) /* +1 for space in the middle */
	if sz3 >= uint64(32) {      /* make sure there is room for ending null */
		return FALSE
	}
	/*
	 * oddly annoying  last.val + ' ' + current.val
	 */
	libc.Xmemcpy(tls, bp, a+23, sz1)
	(**(**[32]int8)(__ccgo_up(bp)))[sz1] = int8(' ')
	libc.Xmemcpy(tls, bp+uintptr(sz1)+uintptr(1), b+23, sz2)
	(**(**[32]int8)(__ccgo_up(bp)))[sz3] = int8('\000')
	ch = (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sf)).Flookup})))(tls, sf, int32(LOOKUP_WORD), bp, sz3)
	if int32(ch) != int32('\000') {
		st_assign(tls, a, ch, (*stoken_t)(unsafe.Pointer(a)).Fpos, sz3, bp)
		return int32(TRUE)
	} else {
		return FALSE
	}
	return r
}

func libinjection_sqli_fold(tls *libc.TLS, sf uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var left, pos size_t
	var more int32
	var v1 bool
	var _ /* last_comment at bp+0 */ stoken_t
	_, _, _, _ = left, more, pos, v1
	/* POS is the position of where the NEXT token goes */
	pos = uint64(0)
	/* LEFT is a count of how many tokens that are already
	   folded or processed (i.e. part of the fingerprint) */
	left = uint64(0)
	more = int32(1)
	st_clear(tls, bp)
	/* Skip all initial comments, right-parens ( and unary operators
	 *
	 */
	(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent = sf + 48
	for more != 0 {
		more = libinjection_sqli_tokenize(tls, sf)
		if !(int32((*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1) == int32(TYPE_COMMENT) || int32((*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1) == int32(TYPE_LEFTPARENS) || int32((*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1) == int32(TYPE_SQLTYPE) || st_is_unary_op(tls, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent) != 0) {
			break
		}
	}
	if !(more != 0) {
		/* If input was only comments, unary or (, then exit */
		return 0
	} else {
		/* it's some other token */
		pos = pos + uint64(1)
	}
	for int32(1) != 0 {
		/* do we have all the max number of tokens?  if so do
		 * some special cases for 5 tokens
		 */
		if pos >= uint64(LIBINJECTION_SQLI_MAX_TOKENS) {
			if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48))).Ftype1) == int32(TYPE_NUMBER) && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 1*56))).Ftype1) == int32(TYPE_OPERATOR) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 1*56))).Ftype1) == int32(TYPE_COMMA)) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 2*56))).Ftype1) == int32(TYPE_LEFTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 3*56))).Ftype1) == int32(TYPE_NUMBER) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 4*56))).Ftype1) == int32(TYPE_RIGHTPARENS) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48))).Ftype1) == int32(TYPE_BAREWORD) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 1*56))).Ftype1) == int32(TYPE_OPERATOR) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 2*56))).Ftype1) == int32(TYPE_LEFTPARENS) && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 3*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 3*56))).Ftype1) == int32(TYPE_NUMBER)) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 4*56))).Ftype1) == int32(TYPE_RIGHTPARENS) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48))).Ftype1) == int32(TYPE_NUMBER) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 1*56))).Ftype1) == int32(TYPE_RIGHTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 2*56))).Ftype1) == int32(TYPE_COMMA) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 3*56))).Ftype1) == int32(TYPE_LEFTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 4*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48))).Ftype1) == int32(TYPE_BAREWORD) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 1*56))).Ftype1) == int32(TYPE_RIGHTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 2*56))).Ftype1) == int32(TYPE_OPERATOR) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 3*56))).Ftype1) == int32(TYPE_LEFTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + 4*56))).Ftype1) == int32(TYPE_BAREWORD) {
				if pos > uint64(LIBINJECTION_SQLI_MAX_TOKENS) {
					st_copy(tls, sf+48+1*56, sf+48+5*56)
					pos = uint64(2)
					left = uint64(0)
				} else {
					pos = uint64(1)
					left = uint64(0)
				}
			}
		}
		if !(more != 0) || left >= uint64(LIBINJECTION_SQLI_MAX_TOKENS) {
			left = pos
			break
		}
		/* get up to two tokens */
		for more != 0 && pos <= uint64(LIBINJECTION_SQLI_MAX_TOKENS) && pos-left < uint64(2) {
			(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent = sf + 48 + uintptr(pos)*56
			more = libinjection_sqli_tokenize(tls, sf)
			if more != 0 {
				if int32((*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1) == int32(TYPE_COMMENT) {
					st_copy(tls, bp, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)
				} else {
					(**(**stoken_t)(__ccgo_up(bp))).Ftype1 = int8('\000')
					pos = pos + uint64(1)
				}
			}
		}
		/* did we get 2 tokens? if not then we are done */
		if pos-left < uint64(2) {
			left = pos
			continue
		}
		/* FOLD: "ss" -> "s"
		 * "foo" "bar" is valid SQL
		 * just ignore second string
		 */
		if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_STRING) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_STRING) {
			pos = pos - uint64(1)
			**(**int32)(__ccgo_up(sf + 532)) += int32(1)
			continue
		} else {
			if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_SEMICOLON) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_SEMICOLON) {
				/* not sure how various engines handle
				 * 'select 1;;drop table foo' or
				 * 'select 1; /x foo x/; drop table foo'
				 * to prevent surprises, just fold away repeated semicolons
				 */
				pos = pos - uint64(1)
				**(**int32)(__ccgo_up(sf + 532)) += int32(1)
				continue
			} else {
				if (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_OPERATOR) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_LOGIC_OPERATOR)) && (st_is_unary_op(tls, sf+48+uintptr(left+uint64(1))*56) != 0 || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_SQLTYPE)) {
					pos = pos - uint64(1)
					**(**int32)(__ccgo_up(sf + 532)) += int32(1)
					left = uint64(0)
					continue
				} else {
					if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_LEFTPARENS) && st_is_unary_op(tls, sf+48+uintptr(left+uint64(1))*56) != 0 {
						pos = pos - uint64(1)
						**(**int32)(__ccgo_up(sf + 532)) += int32(1)
						if left > uint64(0) {
							left = left - uint64(1)
						}
						continue
					} else {
						if syntax_merge_words(tls, sf, sf+48+uintptr(left)*56, sf+48+uintptr(left+uint64(1))*56) != 0 {
							pos = pos - uint64(1)
							**(**int32)(__ccgo_up(sf + 532)) += int32(1)
							if left > uint64(0) {
								left = left - uint64(1)
							}
							continue
						} else {
							if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_SEMICOLON) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_FUNCTION) && (int32(**(**int8)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56 + 23))) == int32('I') || int32(**(**int8)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56 + 23))) == int32('i')) && (int32(**(**int8)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56 + 23 + 1))) == int32('F') || int32(**(**int8)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56 + 23 + 1))) == int32('f')) {
								/* IF is normally a function, except in Transact-SQL where it can be used as a
								 * standalone control flow operator, e.g. ; IF 1=1 ...
								 * if found after a semicolon, convert from 'f' type to 'T' type
								 */
								(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1 = int8(TYPE_TSQL)
								/* left += 2; */
								continue /* reparse everything, but we probably can advance left, and pos */
							} else {
								if (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_VARIABLE)) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_LEFTPARENS) && (cstrcasecmp(tls, __ccgo_ts+67061, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+67085, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+59862, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+63956, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+67056, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+59807, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+59657, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+59759, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+59772, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+62761, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+62771, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0) {
									/* pos is the same
									 * other conversions need to go here... for instance
									 * password CAN be a function, coalesce CAN be a function
									 */
									(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1 = int8(TYPE_FUNCTION)
									continue
								} else {
									if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_KEYWORD) && (cstrcasecmp(tls, __ccgo_ts+61802, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+63461, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0) {
										if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_LEFTPARENS) {
											/* got .... IN ( ...  (or 'NOT IN')
											 * it's an operator
											 */
											(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1 = int8(TYPE_OPERATOR)
										} else {
											/*
											 * it's a nothing
											 */
											(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1 = int8(TYPE_BAREWORD)
										}
										/* "IN" can be used as "IN BOOLEAN MODE" for mysql
										 *  in which case merging of words can be done later
										 * other wise it acts as an equality operator __ IN (values..)
										 *
										 * here we got "IN" "(" so it's an operator.
										 * also back track to handle "NOT IN"
										 * might need to do the same with like
										 * two use cases   "foo" LIKE "BAR" (normal operator)
										 *  "foo" = LIKE(1,2)
										 */
										continue
									} else {
										if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_OPERATOR) && (cstrcasecmp(tls, __ccgo_ts+62685, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 || cstrcasecmp(tls, __ccgo_ts+63468, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0) {
											if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_LEFTPARENS) {
												/* SELECT LIKE(...
												 * it's a function
												 */
												(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1 = int8(TYPE_FUNCTION)
											}
										} else {
											if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_SQLTYPE) && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_SQLTYPE) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_LEFTPARENS) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_FUNCTION) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_VARIABLE) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_STRING)) {
												st_copy(tls, sf+48+uintptr(left)*56, sf+48+uintptr(left+uint64(1))*56)
												pos = pos - uint64(1)
												**(**int32)(__ccgo_up(sf + 532)) += int32(1)
												left = uint64(0)
												continue
											} else {
												if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_COLLATE) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_BAREWORD) {
													/*
													 * there are too many collation types.. so if the bareword has a "_"
													 * then it's TYPE_SQLTYPE
													 */
													if libc.Xstrchr(tls, sf+48+uintptr(left+uint64(1))*56+23, int32('_')) != libc.UintptrFromInt32(0) {
														(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1 = int8(TYPE_SQLTYPE)
														left = uint64(0)
													}
												} else {
													if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_BACKSLASH) {
														if st_is_arithmetic_op(tls, sf+48+uintptr(left+uint64(1))*56) != 0 {
															/* very weird case in TSQL where '\%1' is parsed as '0 % 1', etc */
															(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1 = int8(TYPE_NUMBER)
														} else {
															/* just ignore it.. Again T-SQL seems to parse \1 as "1" */
															st_copy(tls, sf+48+uintptr(left)*56, sf+48+uintptr(left+uint64(1))*56)
															pos = pos - uint64(1)
															**(**int32)(__ccgo_up(sf + 532)) += int32(1)
														}
														left = uint64(0)
														continue
													} else {
														if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_LEFTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_LEFTPARENS) {
															pos = pos - uint64(1)
															left = uint64(0)
															**(**int32)(__ccgo_up(sf + 532)) += int32(1)
															continue
														} else {
															if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_RIGHTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_RIGHTPARENS) {
																pos = pos - uint64(1)
																left = uint64(0)
																**(**int32)(__ccgo_up(sf + 532)) += int32(1)
																continue
															} else {
																if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_LEFTBRACE) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_BAREWORD) {
																	/*
																	 * MySQL Degenerate case --
																	 *
																	 *   select { ``.``.id };  -- valid !!!
																	 *   select { ``.``.``.id };  -- invalid
																	 *   select ``.``.id; -- invalid
																	 *   select { ``.id }; -- invalid
																	 *
																	 * so it appears {``.``.id} is a magic case
																	 * I suspect this is "current database, current table, field id"
																	 *
																	 * The folding code can't look at more than 3 tokens, and
																	 * I don't want to make two passes.
																	 *
																	 * Since "{ ``" so rare, we are just going to blacklist it.
																	 *
																	 * Highly likely this will need revisiting!
																	 *
																	 * CREDIT @rsalgado 2013-11-25
																	 */
																	if (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Flen1 == uint64(0) {
																		(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1 = int8(TYPE_EVIL)
																		return libc.Int32FromUint64(left + libc.Uint64FromInt32(2))
																	}
																	/* weird ODBC / MYSQL  {foo expr} --> expr
																	 * but for this rule we just strip away the "{ foo" part
																	 */
																	left = uint64(0)
																	pos = pos - uint64(2)
																	**(**int32)(__ccgo_up(sf + 532)) += int32(2)
																	continue
																} else {
																	if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_RIGHTBRACE) {
																		pos = pos - uint64(1)
																		left = uint64(0)
																		**(**int32)(__ccgo_up(sf + 532)) += int32(1)
																		continue
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		/* all cases of handing 2 tokens is done
		   and nothing matched.  Get one more token
		*/
		for more != 0 && pos <= uint64(LIBINJECTION_SQLI_MAX_TOKENS) && pos-left < uint64(3) {
			(*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent = sf + 48 + uintptr(pos)*56
			more = libinjection_sqli_tokenize(tls, sf)
			if more != 0 {
				if int32((*libinjection_sqli_token)(unsafe.Pointer((*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)).Ftype1) == int32(TYPE_COMMENT) {
					st_copy(tls, bp, (*libinjection_sqli_state)(unsafe.Pointer(sf)).Fcurrent)
				} else {
					(**(**stoken_t)(__ccgo_up(bp))).Ftype1 = int8('\000')
					pos = pos + uint64(1)
				}
			}
		}
		/* do we have three tokens? If not then we are done */
		if pos-left < uint64(3) {
			left = pos
			continue
		}
		/*
		 * now look for three token folding
		 */
		if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_NUMBER) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_OPERATOR) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_NUMBER) {
			pos = pos - uint64(2)
			left = uint64(0)
			continue
		} else {
			if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_OPERATOR) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) != int32(TYPE_LEFTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_OPERATOR) {
				left = uint64(0)
				pos = pos - uint64(2)
				continue
			} else {
				if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_LOGIC_OPERATOR) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_LOGIC_OPERATOR) {
					pos = pos - uint64(2)
					left = uint64(0)
					continue
				} else {
					if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_VARIABLE) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_OPERATOR) && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_VARIABLE) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_BAREWORD)) {
						pos = pos - uint64(2)
						left = uint64(0)
						continue
					} else {
						if (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_NUMBER)) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_OPERATOR) && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_BAREWORD)) {
							pos = pos - uint64(2)
							left = uint64(0)
							continue
						} else {
							if (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_VARIABLE) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_STRING)) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_OPERATOR) && streq(tls, sf+48+uintptr(left+uint64(1))*56+23, __ccgo_ts+58068) != 0 && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_SQLTYPE) {
								pos = pos - uint64(2)
								left = uint64(0)
								**(**int32)(__ccgo_up(sf + 532)) += int32(2)
								continue
							} else {
								if (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_STRING) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_VARIABLE)) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_COMMA) && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_STRING) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_VARIABLE)) {
									pos = pos - uint64(2)
									left = uint64(0)
									continue
								} else {
									if (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_EXPRESSION) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_GROUP) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_COMMA)) && st_is_unary_op(tls, sf+48+uintptr(left+uint64(1))*56) != 0 && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_LEFTPARENS) {
										/* got something like SELECT + (, LIMIT + (
										 * remove unary operator
										 */
										st_copy(tls, sf+48+uintptr(left+uint64(1))*56, sf+48+uintptr(left+uint64(2))*56)
										pos = pos - uint64(1)
										left = uint64(0)
										continue
									} else {
										if (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_KEYWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_EXPRESSION) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_GROUP)) && st_is_unary_op(tls, sf+48+uintptr(left+uint64(1))*56) != 0 && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_VARIABLE) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_STRING) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_FUNCTION)) {
											/* remove unary operators
											 * select - 1
											 */
											st_copy(tls, sf+48+uintptr(left+uint64(1))*56, sf+48+uintptr(left+uint64(2))*56)
											pos = pos - uint64(1)
											left = uint64(0)
											continue
										} else {
											if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_COMMA) && st_is_unary_op(tls, sf+48+uintptr(left+uint64(1))*56) != 0 && (int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_NUMBER) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_BAREWORD) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_VARIABLE) || int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_STRING)) {
												/*
												 * interesting case    turn ", -1"  ->> ",1" PLUS we need to back up
												 * one token if possible to see if more folding can be done
												 * "1,-1" --> "1"
												 */
												st_copy(tls, sf+48+uintptr(left+uint64(1))*56, sf+48+uintptr(left+uint64(2))*56)
												left = uint64(0)
												/* pos is >= 3 so this is safe */
												if v1 = pos >= uint64(3); !v1 {
													libc.X__assert_fail(tls, __ccgo_ts+68213, __ccgo_ts+67966, int32(1813), uintptr(unsafe.Pointer(&__func__1)))
												}
												_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
												pos = pos - uint64(3)
												continue
											} else {
												if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_COMMA) && st_is_unary_op(tls, sf+48+uintptr(left+uint64(1))*56) != 0 && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_FUNCTION) {
													/* Separate case from above since you end up with
													 * 1,-sin(1) --> 1 (1)
													 * Here, just do
													 * 1,-sin(1) --> 1,sin(1)
													 * just remove unary operator
													 */
													st_copy(tls, sf+48+uintptr(left+uint64(1))*56, sf+48+uintptr(left+uint64(2))*56)
													pos = pos - uint64(1)
													left = uint64(0)
													continue
												} else {
													if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_BAREWORD) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_DOT) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_BAREWORD) {
														/* ignore the '.n'
														 * typically is this databasename.table
														 */
														if v1 = pos >= uint64(3); !v1 {
															libc.X__assert_fail(tls, __ccgo_ts+68213, __ccgo_ts+67966, int32(1836), uintptr(unsafe.Pointer(&__func__1)))
														}
														_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
														pos = pos - uint64(2)
														left = uint64(0)
														continue
													} else {
														if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_EXPRESSION) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_DOT) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) == int32(TYPE_BAREWORD) {
															/* select . `foo` --> select `foo` */
															st_copy(tls, sf+48+uintptr(left+uint64(1))*56, sf+48+uintptr(left+uint64(2))*56)
															pos = pos - uint64(1)
															left = uint64(0)
															continue
														} else {
															if int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1) == int32(TYPE_FUNCTION) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(1))*56))).Ftype1) == int32(TYPE_LEFTPARENS) && int32((**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left+uint64(2))*56))).Ftype1) != int32(TYPE_RIGHTPARENS) {
																/*
																 * whats going on here
																 * Some SQL functions like USER() have 0 args
																 * if we get User(foo), then User is not a function
																 * This should be expanded since it eliminated a lot of false
																 * positives.
																 */
																if cstrcasecmp(tls, __ccgo_ts+67056, sf+48+uintptr(left)*56+23, (**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Flen1) == 0 {
																	(**(**libinjection_sqli_token)(__ccgo_up(sf + 48 + uintptr(left)*56))).Ftype1 = int8(TYPE_BAREWORD)
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		/* no folding -- assume left-most token is
		   is good, now use the existing 2 tokens --
		   do not get another
		*/
		left = left + uint64(1)
	} /* while(1) */
	/* if we have 4 or less tokens, and we had a comment token
	 * at the end, add it back
	 */
	if left < uint64(LIBINJECTION_SQLI_MAX_TOKENS) && int32((**(**stoken_t)(__ccgo_up(bp))).Ftype1) == int32(TYPE_COMMENT) {
		st_copy(tls, sf+48+uintptr(left)*56, bp)
		left = left + uint64(1)
	}
	/* sometimes we grab a 6th token to help
	   determine the type of token 5.
	*/
	if left > uint64(LIBINJECTION_SQLI_MAX_TOKENS) {
		left = uint64(LIBINJECTION_SQLI_MAX_TOKENS)
	}
	return libc.Int32FromUint64(left)
}

var __func__1 = [23]int8{'l', 'i', 'b', 'i', 'n', 'j', 'e', 'c', 't', 'i', 'o', 'n', '_', 's', 'q', 'l', 'i', '_', 'f', 'o', 'l', 'd'}

// C documentation
//
//	/* secondary api: detects SQLi in a string, GIVEN a context.
//	 *
//	 * A context can be:
//	 *   *  CHAR_NULL (\0), process as is
//	 *   *  CHAR_SINGLE ('), process pretending input started with a
//	 *          single quote.
//	 *   *  CHAR_DOUBLE ("), process pretending input started with a
//	 *          double quote.
//	 *
//	 */
func libinjection_sqli_fingerprint(tls *libc.TLS, sql_state uintptr, flags int32) (r uintptr) {
	var i, tlen int32
	_, _ = i, tlen
	tlen = 0
	libinjection_sqli_reset(tls, sql_state, flags)
	tlen = libinjection_sqli_fold(tls, sql_state)
	/* Check for magic PHP backquote comment
	 * If:
	 * * last token is of type "bareword"
	 * * And is quoted in a backtick
	 * * And isn't closed
	 * * And it's empty?
	 * Then convert it to comment
	 */
	if tlen > int32(2) && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + uintptr(tlen-int32(1))*56))).Ftype1) == int32(TYPE_BAREWORD) && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + uintptr(tlen-int32(1))*56))).Fstr_open) == int32('`') && (**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + uintptr(tlen-int32(1))*56))).Flen1 == uint64(0) && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + uintptr(tlen-int32(1))*56))).Fstr_close) == int32('\000') {
		(**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + uintptr(tlen-int32(1))*56))).Ftype1 = int8(TYPE_COMMENT)
	}
	i = 0
	for {
		if !(i < tlen) {
			break
		}
		**(**int8)(__ccgo_up(sql_state + 504 + uintptr(i))) = (**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + uintptr(i)*56))).Ftype1
		goto _1
	_1:
		;
		i = i + 1
	}
	/*
	 * make the fingerprint pattern a c-string (null delimited)
	 */
	**(**int8)(__ccgo_up(sql_state + 504 + uintptr(tlen))) = int8('\000')
	/*
	 * check for 'X' in pattern, and then
	 * clear out all tokens
	 *
	 * this means parsing could not be done
	 * accurately due to pgsql's double comments
	 * or other syntax that isn't consistent.
	 * Should be very rare false positive
	 */
	if libc.Xstrchr(tls, sql_state+504, int32(TYPE_EVIL)) != 0 {
		/*  needed for SWIG */
		libc.Xmemset(tls, sql_state+504, 0, libc.Uint64FromInt32(libc.Int32FromInt32(LIBINJECTION_SQLI_MAX_TOKENS)+libc.Int32FromInt32(1)))
		libc.Xmemset(tls, sql_state+48+23, 0, uint64(32))
		**(**int8)(__ccgo_up(sql_state + 504)) = int8(TYPE_EVIL)
		(**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Ftype1 = int8(TYPE_EVIL)
		**(**int8)(__ccgo_up(sql_state + 48 + 23)) = int8(TYPE_EVIL)
		(**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 1*56))).Ftype1 = int8('\000')
	}
	return sql_state + 504
}

func libinjection_sqli_check_fingerprint(tls *libc.TLS, sql_state uintptr) (r int32) {
	return libc.BoolInt32(libinjection_sqli_blacklist(tls, sql_state) != 0 && libinjection_sqli_not_whitelist(tls, sql_state) != 0)
}

func libinjection_sqli_lookup_word(tls *libc.TLS, sql_state uintptr, lookup_type int32, str uintptr, len1 size_t) (r int8) {
	var v1 int32
	_ = v1
	if lookup_type == int32(LOOKUP_FINGERPRINT) {
		if libinjection_sqli_check_fingerprint(tls, sql_state) != 0 {
			v1 = int32('X')
		} else {
			v1 = int32('\000')
		}
		return int8(v1)
	} else {
		return bsearch_keyword_type(tls, str, len1, uintptr(unsafe.Pointer(&sql_keywords)), sql_keywords_sz)
	}
	return r
}

func libinjection_sqli_blacklist(tls *libc.TLS, sql_state uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ch int8
	var i, len1 size_t
	var patmatch int32
	var _ /* fp2 at bp+0 */ [8]int8
	_, _, _, _ = ch, i, len1, patmatch
	len1 = libc.Xstrlen(tls, sql_state+504)
	if len1 < uint64(1) {
		(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(1989)
		return FALSE
	}
	/*
	   to keep everything compatible, convert the
	   v0 fingerprint pattern to v1
	   v0: up to 5 chars, mixed case
	   v1: 1 char is '0', up to 5 more chars, upper case
	*/
	(**(**[8]int8)(__ccgo_up(bp)))[0] = int8('0')
	i = uint64(0)
	for {
		if !(i < len1) {
			break
		}
		ch = **(**int8)(__ccgo_up(sql_state + 504 + uintptr(i)))
		if int32(ch) >= int32('a') && int32(ch) <= int32('z') {
			ch = int8(int32(ch) - libc.Int32FromInt32(0x20))
		}
		(**(**[8]int8)(__ccgo_up(bp)))[i+uint64(1)] = ch
		goto _1
	_1:
		;
		i = i + 1
	}
	(**(**[8]int8)(__ccgo_up(bp)))[i+uint64(1)] = int8('\000')
	patmatch = libc.BoolInt32(int32(is_keyword(tls, bp, len1+uint64(1))) == int32(TYPE_FINGERPRINT))
	/*
	 * No match.
	 *
	 * Set sql_state->reason to current line number
	 * only for debugging purposes.
	 */
	if !(patmatch != 0) {
		(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2019)
		return FALSE
	}
	return int32(TRUE)
}

// C documentation
//
//	/*
//	 * return TRUE if SQLi, false is benign
//	 */
func libinjection_sqli_not_whitelist(tls *libc.TLS, sql_state uintptr) (r int32) {
	var ch int8
	var tlen size_t
	_, _ = ch, tlen
	tlen = libc.Xstrlen(tls, sql_state+504)
	if tlen > uint64(1) && int32(**(**int8)(__ccgo_up(sql_state + 504 + uintptr(tlen-uint64(1))))) == int32(TYPE_COMMENT) {
		/*
		 * if ending comment is contains 'sp_password' then it's SQLi!
		 * MS Audit log apparently ignores anything with
		 * 'sp_password' in it. Unable to find primary reference to
		 * this "feature" of SQL Server but seems to be known SQLi
		 * technique
		 */
		if my_memmem(tls, (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fs, (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fslen, __ccgo_ts+68222, libc.Xstrlen(tls, __ccgo_ts+68222)) != 0 {
			(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2049)
			return int32(TRUE)
		}
	}
	switch tlen {
	case uint64(2):
		/*
		 * case 2 are "very small SQLi" which make them
		 * hard to tell from normal input...
		 */
		if int32(**(**int8)(__ccgo_up(sql_state + 504 + 1))) == int32(TYPE_UNION) {
			if (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fstats_tokens == int32(2) {
				/* not sure why but 1U comes up in SQLi attack
				 * likely part of parameter splitting/etc.
				 * lots of reasons why "1 union" might be normal
				 * input, so beep only if other SQLi things are present
				 */
				/* it really is a number and 'union'
				 * other wise it has folding or comments
				 */
				(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2071)
				return FALSE
			} else {
				(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2074)
				return int32(TRUE)
			}
		}
		/*
		 * if 'comment' is '#' ignore.. too many FP
		 */
		if int32(**(**int8)(__ccgo_up(sql_state + 48 + 1*56 + 23))) == int32('#') {
			(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2082)
			return FALSE
		}
		/*
		 * for fingerprint like 'nc', only comments of /x are treated
		 * as SQL... ending comments of "--" and "#" are not SQLi
		 */
		if int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Ftype1) == int32(TYPE_BAREWORD) && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 1*56))).Ftype1) == int32(TYPE_COMMENT) && int32(**(**int8)(__ccgo_up(sql_state + 48 + 1*56 + 23))) != int32('/') {
			(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2093)
			return FALSE
		}
		/*
		 * if '1c' ends with '/x' then it's SQLi
		 */
		if int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Ftype1) == int32(TYPE_NUMBER) && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 1*56))).Ftype1) == int32(TYPE_COMMENT) && int32(**(**int8)(__ccgo_up(sql_state + 48 + 1*56 + 23))) == int32('/') {
			return int32(TRUE)
		}
		/**
		 * there are some odd base64-looking query string values
		 * 1234-ABCDEFEhfhihwuefi--
		 * which evaluate to "1c"... these are not SQLi
		 * but 1234-- probably is.
		 * Make sure the "1" in "1c" is actually a true decimal number
		 *
		 * Need to check -original- string since the folding step
		 * may have merged tokens, e.g. "1+FOO" is folded into "1"
		 *
		 * Note: evasion: 1*1--
		 */
		if int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Ftype1) == int32(TYPE_NUMBER) && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 1*56))).Ftype1) == int32(TYPE_COMMENT) {
			if (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fstats_tokens > int32(2) {
				/* we have some folding going on, highly likely SQLi */
				(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2122)
				return int32(TRUE)
			}
			/*
			 * we check that next character after the number is either whitespace,
			 * or '/' or a '-' ==> SQLi.
			 */
			ch = **(**int8)(__ccgo_up((*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fs + uintptr((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Flen1)))
			if int32(ch) <= int32(32) {
				/* next char was whitespace,e.g. "1234 --"
				 * this isn't exactly correct.. ideally we should skip over all whitespace
				 * but this seems to be ok for now
				 */
				return int32(TRUE)
			}
			if int32(ch) == int32('/') && int32(**(**int8)(__ccgo_up((*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fs + uintptr((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Flen1+uint64(1))))) == int32('*') {
				return int32(TRUE)
			}
			if int32(ch) == int32('-') && int32(**(**int8)(__ccgo_up((*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fs + uintptr((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Flen1+uint64(1))))) == int32('-') {
				return int32(TRUE)
			}
			(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2144)
			return FALSE
		}
		/*
		 * detect obvious SQLi scans.. many people put '--' in plain text
		 * so only detect if input ends with '--', e.g. 1-- but not 1-- foo
		 */
		if (**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 1*56))).Flen1 > uint64(2) && int32(**(**int8)(__ccgo_up(sql_state + 48 + 1*56 + 23))) == int32('-') {
			(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2154)
			return FALSE
		}
		break
		/* case 2 */
		fallthrough
	case uint64(3):
		/*
		 * ...foo' + 'bar...
		 * no opening quote, no closing quote
		 * and each string has data
		 */
		if streq(tls, sql_state+504, __ccgo_ts+68234) != 0 || streq(tls, sql_state+504, __ccgo_ts+68238) != 0 {
			if int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Fstr_open) == int32('\000') && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 2*56))).Fstr_close) == int32('\000') && int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48))).Fstr_close) == int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 2*56))).Fstr_open) {
				/*
				 * if ....foo" + "bar....
				 */
				(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2176)
				return int32(TRUE)
			}
			if (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fstats_tokens == int32(3) {
				(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2180)
				return FALSE
			}
			/*
			 * not SQLi
			 */
			(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2187)
			return FALSE
		} else {
			if streq(tls, sql_state+504, __ccgo_ts+68242) != 0 || streq(tls, sql_state+504, __ccgo_ts+68246) != 0 || streq(tls, sql_state+504, __ccgo_ts+68250) != 0 || streq(tls, sql_state+504, __ccgo_ts+68254) != 0 || streq(tls, sql_state+504, __ccgo_ts+68258) != 0 {
				/* 'sexy and 17' not SQLi
				 * 'sexy and 17<18'  SQLi
				 */
				if (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fstats_tokens == int32(3) {
					(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2198)
					return FALSE
				}
			} else {
				if int32((**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 1*56))).Ftype1) == int32(TYPE_KEYWORD) {
					if (**(**libinjection_sqli_token)(__ccgo_up(sql_state + 48 + 1*56))).Flen1 < uint64(5) || cstrcasecmp(tls, __ccgo_ts+62214, sql_state+48+1*56+23, uint64(4)) != 0 {
						/* if it's not "INTO OUTFILE", or "INTO DUMPFILE" (MySQL)
						 * then treat as safe
						 */
						(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Freason = int32(2207)
						return FALSE
					}
				}
			}
		}
		break
		/* case 3 */
		fallthrough
	case uint64(4):
		fallthrough
	case uint64(5):
		/* nothing right now */
		break
		/* case 5 */
	} /* end switch */
	return int32(TRUE)
}

// C documentation
//
//	/**  Main API, detects SQLi in an input.
//	 *
//	 *
//	 */
func reparse_as_mysql(tls *libc.TLS, sql_state uintptr) (r int32) {
	return libc.BoolInt32((*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fstats_comment_ddx != 0 || (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fstats_comment_hash != 0)
}

// C documentation
//
//	/*
//	 * This function is mostly use with SWIG
//	 */
func libinjection_sqli_get_token(tls *libc.TLS, sql_state uintptr, i int32) (r uintptr) {
	if i < 0 || i > libc.Int32FromInt32(LIBINJECTION_SQLI_MAX_TOKENS) {
		return libc.UintptrFromInt32(0)
	}
	return sql_state + 48 + uintptr(i)*56
}

func libinjection_is_sqli(tls *libc.TLS, sql_state uintptr) (r int32) {
	var s uintptr
	var slen size_t
	_, _ = s, slen
	s = (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fs
	slen = (*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Fslen
	/*
	 * no input? not SQLi
	 */
	if slen == uint64(0) {
		return FALSE
	}
	/*
	 * test input "as-is"
	 */
	libinjection_sqli_fingerprint(tls, sql_state, int32(FLAG_QUOTE_NONE)|int32(FLAG_SQL_ANSI))
	if (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Flookup})))(tls, sql_state, int32(LOOKUP_FINGERPRINT), sql_state+504, libc.Xstrlen(tls, sql_state+504)) != 0 {
		return int32(TRUE)
	} else {
		if reparse_as_mysql(tls, sql_state) != 0 {
			libinjection_sqli_fingerprint(tls, sql_state, int32(FLAG_QUOTE_NONE)|int32(FLAG_SQL_MYSQL))
			if (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Flookup})))(tls, sql_state, int32(LOOKUP_FINGERPRINT), sql_state+504, libc.Xstrlen(tls, sql_state+504)) != 0 {
				return int32(TRUE)
			}
		}
	}
	/*
	 * if input has a single_quote, then
	 * test as if input was actually '
	 * example: if input if "1' = 1", then pretend it's
	 *   "'1' = 1"
	 * Porting Notes: example the same as doing
	 *   is_string_sqli(sql_state, "'" + s, slen+1, NULL, fn, arg)
	 *
	 */
	if libc.Xmemchr(tls, s, int32('\''), slen) != 0 {
		libinjection_sqli_fingerprint(tls, sql_state, int32(FLAG_QUOTE_SINGLE)|int32(FLAG_SQL_ANSI))
		if (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Flookup})))(tls, sql_state, int32(LOOKUP_FINGERPRINT), sql_state+504, libc.Xstrlen(tls, sql_state+504)) != 0 {
			return int32(TRUE)
		} else {
			if reparse_as_mysql(tls, sql_state) != 0 {
				libinjection_sqli_fingerprint(tls, sql_state, int32(FLAG_QUOTE_SINGLE)|int32(FLAG_SQL_MYSQL))
				if (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Flookup})))(tls, sql_state, int32(LOOKUP_FINGERPRINT), sql_state+504, libc.Xstrlen(tls, sql_state+504)) != 0 {
					return int32(TRUE)
				}
			}
		}
	}
	/*
	 * same as above but with a double-quote "
	 */
	if libc.Xmemchr(tls, s, int32('"'), slen) != 0 {
		libinjection_sqli_fingerprint(tls, sql_state, int32(FLAG_QUOTE_DOUBLE)|int32(FLAG_SQL_MYSQL))
		if (*(*func(*libc.TLS, uintptr, int32, uintptr, size_t) int8)(unsafe.Pointer(&struct{ uintptr }{(*libinjection_sqli_state)(unsafe.Pointer(sql_state)).Flookup})))(tls, sql_state, int32(LOOKUP_FINGERPRINT), sql_state+504, libc.Xstrlen(tls, sql_state+504)) != 0 {
			return int32(TRUE)
		}
	}
	/*
	 * Hurray, input is not SQLi
	 */
	return FALSE
}

func libinjection_sqli(tls *libc.TLS, input uintptr, slen size_t, fingerprint uintptr) (r int32) {
	bp := tls.Alloc(544)
	defer tls.Free(544)
	var issqli int32
	var _ /* state at bp+0 */ libinjection_sqli_state
	_ = issqli
	libinjection_sqli_init(tls, bp, input, slen, 0)
	issqli = libinjection_is_sqli(tls, bp)
	if issqli != 0 {
		libc.Xstrcpy(tls, fingerprint, bp+504)
	} else {
		**(**int8)(__ccgo_up(fingerprint)) = int8('\000')
	}
	return issqli
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "!!\x00!<\x00!=\x00!>\x00%=\x00&&\x00&=\x00*=\x00+=\x00-=\x00/=\x000&(1)O\x000&(1)U\x000&(1O(\x000&(1OF\x000&(1OS\x000&(1OV\x000&(F()\x000&(F(1\x000&(F(F\x000&(F(N\x000&(F(S\x000&(F(V\x000&(N)O\x000&(N)U\x000&(NO(\x000&(NOF\x000&(NOS\x000&(NOV\x000&(S)O\x000&(S)U\x000&(SO(\x000&(SO1\x000&(SOF\x000&(SON\x000&(SOS\x000&(SOV\x000&(V)O\x000&(V)U\x000&(VO(\x000&(VOF\x000&(VOS\x000&1O(1\x000&1O(F\x000&1O(N\x000&1O(S\x000&1O(V\x000&1OF(\x000&1OS(\x000&1OS1\x000&1OSF\x000&1OSU\x000&1OSV\x000&1OV(\x000&1OVF\x000&1OVO\x000&1OVS\x000&1OVU\x000&1UE(\x000&1UE1\x000&1UEF\x000&1UEK\x000&1UEN\x000&1UES\x000&1UEV\x000&F()O\x000&F()U\x000&F(1)\x000&F(1O\x000&F(F(\x000&F(N)\x000&F(NO\x000&F(S)\x000&F(SO\x000&F(V)\x000&F(VO\x000&NO(1\x000&NO(F\x000&NO(N\x000&NO(S\x000&NO(V\x000&NOF(\x000&NOS(\x000&NOS1\x000&NOSF\x000&NOSU\x000&NOSV\x000&NOV(\x000&NOVF\x000&NOVO\x000&NOVS\x000&NOVU\x000&NUE(\x000&NUE1\x000&NUEF\x000&NUEK\x000&NUEN\x000&NUES\x000&NUEV\x000&SO(1\x000&SO(F\x000&SO(N\x000&SO(S\x000&SO(V\x000&SO1(\x000&SO1F\x000&SO1N\x000&SO1S\x000&SO1U\x000&SO1V\x000&SOF(\x000&SON(\x000&SON1\x000&SONF\x000&SONU\x000&SOS(\x000&SOS1\x000&SOSF\x000&SOSU\x000&SOSV\x000&SOV(\x000&SOVF\x000&SOVO\x000&SOVS\x000&SOVU\x000&SUE(\x000&SUE1\x000&SUEF\x000&SUEK\x000&SUEN\x000&SUES\x000&SUEV\x000&VO(1\x000&VO(F\x000&VO(N\x000&VO(S\x000&VO(V\x000&VOF(\x000&VOS(\x000&VOS1\x000&VOSF\x000&VOSU\x000&VOSV\x000&VUE(\x000&VUE1\x000&VUEF\x000&VUEK\x000&VUEN\x000&VUES\x000&VUEV\x000)&(EK\x000)&(EN\x000)UE(1\x000)UE(F\x000)UE(N\x000)UE(S\x000)UE(V\x000)UE1K\x000)UE1O\x000)UEF(\x000)UEK(\x000)UEK1\x000)UEKF\x000)UEKN\x000)UEKS\x000)UEKV\x000)UENK\x000)UENO\x000)UESK\x000)UESO\x000)UEVK\x000)UEVO\x0001&(1&\x0001&(1)\x0001&(1,\x0001&(1O\x0001&(E(\x0001&(E1\x0001&(EF\x0001&(EK\x0001&(EN\x0001&(EO\x0001&(ES\x0001&(EV\x0001&(F(\x0001&(N&\x0001&(N)\x0001&(N,\x0001&(NO\x0001&(S&\x0001&(S)\x0001&(S,\x0001&(SO\x0001&(V&\x0001&(V)\x0001&(V,\x0001&(VO\x0001&1\x0001&1&(\x0001&1&1\x0001&1&F\x0001&1&N\x0001&1&S\x0001&1&V\x0001&1)&\x0001&1)C\x0001&1)O\x0001&1)U\x0001&1;\x0001&1;C\x0001&1;E\x0001&1;T\x0001&1B(\x0001&1B1\x0001&1BF\x0001&1BN\x0001&1BS\x0001&1BV\x0001&1C\x0001&1EK\x0001&1EN\x0001&1F(\x0001&1K(\x0001&1K1\x0001&1KF\x0001&1KN\x0001&1KS\x0001&1KV\x0001&1O(\x0001&1OF\x0001&1OS\x0001&1OV\x0001&1TN\x0001&1U\x0001&1U(\x0001&1U;\x0001&1UC\x0001&1UE\x0001&E(1\x0001&E(F\x0001&E(N\x0001&E(O\x0001&E(S\x0001&E(V\x0001&E1\x0001&E1;\x0001&E1C\x0001&E1K\x0001&E1O\x0001&EF(\x0001&EK(\x0001&EK1\x0001&EKF\x0001&EKN\x0001&EKS\x0001&EKU\x0001&EKV\x0001&EN\x0001&EN;\x0001&ENC\x0001&ENK\x0001&ENO\x0001&ES\x0001&ES;\x0001&ESC\x0001&ESK\x0001&ESO\x0001&EUE\x0001&EV\x0001&EV;\x0001&EVC\x0001&EVK\x0001&EVO\x0001&F()\x0001&F(1\x0001&F(E\x0001&F(F\x0001&F(N\x0001&F(S\x0001&F(V\x0001&K&(\x0001&K&1\x0001&K&F\x0001&K&N\x0001&K&S\x0001&K&V\x0001&K(1\x0001&K(F\x0001&K(N\x0001&K(S\x0001&K(V\x0001&K1O\x0001&KC\x0001&KF(\x0001&KNK\x0001&KO(\x0001&KO1\x0001&KOF\x0001&KOK\x0001&KON\x0001&KOS\x0001&KOV\x0001&KSO\x0001&KVO\x0001&N&(\x0001&N&1\x0001&N&F\x0001&N&N\x0001&N&S\x0001&N&V\x0001&N)&\x0001&N)C\x0001&N)O\x0001&N)U\x0001&N;\x0001&N;C\x0001&N;E\x0001&N;T\x0001&NB(\x0001&NB1\x0001&NBF\x0001&NBN\x0001&NBS\x0001&NBV\x0001&NC\x0001&NEN\x0001&NF(\x0001&NK(\x0001&NK1\x0001&NKF\x0001&NKN\x0001&NKS\x0001&NKV\x0001&NO(\x0001&NOF\x0001&NOS\x0001&NOV\x0001&NTN\x0001&NU\x0001&NU(\x0001&NU;\x0001&NUC\x0001&NUE\x0001&S\x0001&S&(\x0001&S&1\x0001&S&F\x0001&S&N\x0001&S&S\x0001&S&V\x0001&S)&\x0001&S)C\x0001&S)O\x0001&S)U\x0001&S1\x0001&S1;\x0001&S1C\x0001&S;\x0001&S;C\x0001&S;E\x0001&S;T\x0001&SB(\x0001&SB1\x0001&SBF\x0001&SBN\x0001&SBS\x0001&SBV\x0001&SC\x0001&SEK\x0001&SEN\x0001&SF(\x0001&SK(\x0001&SK1\x0001&SKF\x0001&SKN\x0001&SKS\x0001&SKV\x0001&SO(\x0001&SO1\x0001&SOF\x0001&SON\x0001&SOS\x0001&SOV\x0001&STN\x0001&SU\x0001&SU(\x0001&SU;\x0001&SUC\x0001&SUE\x0001&SV\x0001&SV;\x0001&SVC\x0001&SVO\x0001&V\x0001&V&(\x0001&V&1\x0001&V&F\x0001&V&N\x0001&V&S\x0001&V&V\x0001&V)&\x0001&V)C\x0001&V)O\x0001&V)U\x0001&V;\x0001&V;C\x0001&V;E\x0001&V;T\x0001&VB(\x0001&VB1\x0001&VBF\x0001&VBN\x0001&VBS\x0001&VBV\x0001&VC\x0001&VEK\x0001&VEN\x0001&VF(\x0001&VK(\x0001&VK1\x0001&VKF\x0001&VKN\x0001&VKS\x0001&VKV\x0001&VO(\x0001&VOF\x0001&VOS\x0001&VS\x0001&VS;\x0001&VSC\x0001&VSO\x0001&VTN\x0001&VU\x0001&VU(\x0001&VU;\x0001&VUC\x0001&VUE\x0001(EF(\x0001(EKF\x0001(EKN\x0001(ENK\x0001(U(E\x0001)&(1\x0001)&(E\x0001)&(F\x0001)&(N\x0001)&(S\x0001)&(V\x0001)&1\x0001)&1&\x0001)&1)\x0001)&1;\x0001)&1B\x0001)&1C\x0001)&1F\x0001)&1O\x0001)&1U\x0001)&F(\x0001)&N\x0001)&N&\x0001)&N)\x0001)&N;\x0001)&NB\x0001)&NC\x0001)&NF\x0001)&NO\x0001)&NU\x0001)&S\x0001)&S&\x0001)&S)\x0001)&S;\x0001)&SB\x0001)&SC\x0001)&SF\x0001)&SO\x0001)&SU\x0001)&V\x0001)&V&\x0001)&V)\x0001)&V;\x0001)&VB\x0001)&VC\x0001)&VF\x0001)&VO\x0001)&VU\x0001),(1\x0001),(F\x0001),(N\x0001),(S\x0001),(V\x0001);E(\x0001);E1\x0001);EF\x0001);EK\x0001);EN\x0001);EO\x0001);ES\x0001);EV\x0001);T(\x0001);T1\x0001);TF\x0001);TK\x0001);TN\x0001);TO\x0001);TS\x0001);TV\x0001)B(1\x0001)B(F\x0001)B(N\x0001)B(S\x0001)B(V\x0001)B1\x0001)B1&\x0001)B1;\x0001)B1C\x0001)B1K\x0001)B1N\x0001)B1O\x0001)B1U\x0001)BF(\x0001)BN\x0001)BN&\x0001)BN;\x0001)BNC\x0001)BNK\x0001)BNO\x0001)BNU\x0001)BS\x0001)BS&\x0001)BS;\x0001)BSC\x0001)BSK\x0001)BSO\x0001)BSU\x0001)BV\x0001)BV&\x0001)BV;\x0001)BVC\x0001)BVK\x0001)BVO\x0001)BVU\x0001)C\x0001)E(1\x0001)E(F\x0001)E(N\x0001)E(S\x0001)E(V\x0001)E1C\x0001)E1O\x0001)EF(\x0001)EK(\x0001)EK1\x0001)EKF\x0001)EKN\x0001)EKS\x0001)EKV\x0001)ENC\x0001)ENO\x0001)ESC\x0001)ESO\x0001)EVC\x0001)EVO\x0001)F(F\x0001)K(1\x0001)K(F\x0001)K(N\x0001)K(S\x0001)K(V\x0001)K1&\x0001)K1;\x0001)K1B\x0001)K1E\x0001)K1O\x0001)K1U\x0001)KB(\x0001)KB1\x0001)KBF\x0001)KBN\x0001)KBS\x0001)KBV\x0001)KF(\x0001)KN&\x0001)KN;\x0001)KNB\x0001)KNC\x0001)KNE\x0001)KNK\x0001)KNU\x0001)KS&\x0001)KS;\x0001)KSB\x0001)KSE\x0001)KSO\x0001)KSU\x0001)KUE\x0001)KV&\x0001)KV;\x0001)KVB\x0001)KVE\x0001)KVO\x0001)KVU\x0001)O(1\x0001)O(E\x0001)O(F\x0001)O(N\x0001)O(S\x0001)O(V\x0001)O1\x0001)O1&\x0001)O1)\x0001)O1;\x0001)O1B\x0001)O1C\x0001)O1K\x0001)O1U\x0001)OF(\x0001)ON&\x0001)ON)\x0001)ON;\x0001)ONB\x0001)ONC\x0001)ONK\x0001)ONU\x0001)OS\x0001)OS&\x0001)OS)\x0001)OS;\x0001)OSB\x0001)OSC\x0001)OSK\x0001)OSU\x0001)OV\x0001)OV&\x0001)OV)\x0001)OV;\x0001)OVB\x0001)OVC\x0001)OVK\x0001)OVO\x0001)OVU\x0001)U(E\x0001)UE(\x0001)UE1\x0001)UEF\x0001)UEK\x0001)UEN\x0001)UES\x0001)UEV\x0001,(1)\x0001,(1O\x0001,(E(\x0001,(E1\x0001,(EF\x0001,(EK\x0001,(EN\x0001,(ES\x0001,(EV\x0001,(F(\x0001,(N)\x0001,(NO\x0001,(S)\x0001,(SO\x0001,(V)\x0001,(VO\x0001,F()\x0001,F(1\x0001,F(F\x0001,F(N\x0001,F(S\x0001,F(V\x0001;E(1\x0001;E(E\x0001;E(F\x0001;E(N\x0001;E(S\x0001;E(V\x0001;E1,\x0001;E1;\x0001;E1C\x0001;E1K\x0001;E1O\x0001;E1T\x0001;EF(\x0001;EK(\x0001;EK1\x0001;EKF\x0001;EKN\x0001;EKO\x0001;EKS\x0001;EKV\x0001;EN,\x0001;EN;\x0001;ENC\x0001;ENE\x0001;ENK\x0001;ENO\x0001;ENT\x0001;ES,\x0001;ES;\x0001;ESC\x0001;ESK\x0001;ESO\x0001;EST\x0001;EV,\x0001;EV;\x0001;EVC\x0001;EVK\x0001;EVO\x0001;EVT\x0001;N:T\x0001;T(1\x0001;T(C\x0001;T(E\x0001;T(F\x0001;T(N\x0001;T(S\x0001;T(V\x0001;T1(\x0001;T1,\x0001;T1;\x0001;T1C\x0001;T1F\x0001;T1K\x0001;T1O\x0001;T1T\x0001;T;\x0001;T;C\x0001;TF(\x0001;TK(\x0001;TK1\x0001;TKF\x0001;TKK\x0001;TKN\x0001;TKO\x0001;TKS\x0001;TKV\x0001;TN(\x0001;TN,\x0001;TN1\x0001;TN;\x0001;TNC\x0001;TNF\x0001;TNK\x0001;TNN\x0001;TNO\x0001;TNS\x0001;TNT\x0001;TNV\x0001;TO(\x0001;TS(\x0001;TS,\x0001;TS;\x0001;TSC\x0001;TSF\x0001;TSK\x0001;TSO\x0001;TST\x0001;TTN\x0001;TV(\x0001;TV,\x0001;TV;\x0001;TVC\x0001;TVF\x0001;TVK\x0001;TVO\x0001;TVT\x0001A(F(\x0001A(N)\x0001A(NO\x0001A(S)\x0001A(SO\x0001A(V)\x0001A(VO\x0001AF()\x0001AF(1\x0001AF(F\x0001AF(N\x0001AF(S\x0001AF(V\x0001ASO(\x0001ASO1\x0001ASOF\x0001ASON\x0001ASOS\x0001ASOV\x0001ASUE\x0001ATO(\x0001ATO1\x0001ATOF\x0001ATON\x0001ATOS\x0001ATOV\x0001ATUE\x0001AVO(\x0001AVOF\x0001AVOS\x0001AVUE\x0001B(1)\x0001B(1O\x0001B(F(\x0001B(NO\x0001B(S)\x0001B(SO\x0001B(V)\x0001B(VO\x0001B1\x0001B1&(\x0001B1&1\x0001B1&F\x0001B1&N\x0001B1&S\x0001B1&V\x0001B1,(\x0001B1,F\x0001B1;\x0001B1;C\x0001B1B(\x0001B1B1\x0001B1BF\x0001B1BN\x0001B1BS\x0001B1BV\x0001B1C\x0001B1K(\x0001B1K1\x0001B1KF\x0001B1KN\x0001B1KS\x0001B1KV\x0001B1O(\x0001B1OF\x0001B1OS\x0001B1OV\x0001B1U(\x0001B1UE\x0001BE(1\x0001BE(F\x0001BE(N\x0001BE(S\x0001BE(V\x0001BEK(\x0001BF()\x0001BF(1\x0001BF(F\x0001BF(N\x0001BF(S\x0001BF(V\x0001BN\x0001BN&(\x0001BN&1\x0001BN&F\x0001BN&N\x0001BN&S\x0001BN&V\x0001BN,(\x0001BN,F\x0001BN;\x0001BN;C\x0001BNB(\x0001BNB1\x0001BNBF\x0001BNBN\x0001BNBS\x0001BNBV\x0001BNC\x0001BNK(\x0001BNK1\x0001BNKF\x0001BNKN\x0001BNKS\x0001BNKV\x0001BNO(\x0001BNOF\x0001BNOS\x0001BNOV\x0001BNU(\x0001BNUE\x0001BS\x0001BS&(\x0001BS&1\x0001BS&F\x0001BS&N\x0001BS&S\x0001BS&V\x0001BS,(\x0001BS,F\x0001BS;\x0001BS;C\x0001BSB(\x0001BSB1\x0001BSBF\x0001BSBN\x0001BSBS\x0001BSBV\x0001BSC\x0001BSK(\x0001BSK1\x0001BSKF\x0001BSKN\x0001BSKS\x0001BSKV\x0001BSO(\x0001BSO1\x0001BSOF\x0001BSON\x0001BSOS\x0001BSOV\x0001BSU(\x0001BSUE\x0001BV\x0001BV&(\x0001BV&1\x0001BV&F\x0001BV&N\x0001BV&S\x0001BV&V\x0001BV,(\x0001BV,F\x0001BV;\x0001BV;C\x0001BVB(\x0001BVB1\x0001BVBF\x0001BVBN\x0001BVBS\x0001BVBV\x0001BVC\x0001BVK(\x0001BVK1\x0001BVKF\x0001BVKN\x0001BVKS\x0001BVKV\x0001BVO(\x0001BVOF\x0001BVOS\x0001BVU(\x0001BVUE\x0001C\x0001E(1)\x0001E(1O\x0001E(F(\x0001E(N)\x0001E(NO\x0001E(S)\x0001E(SO\x0001E(V)\x0001E(VO\x0001E1;T\x0001E1C\x0001E1O(\x0001E1OF\x0001E1OS\x0001E1OV\x0001E1T(\x0001E1T1\x0001E1TF\x0001E1TN\x0001E1TS\x0001E1TV\x0001E1UE\x0001EF()\x0001EF(1\x0001EF(F\x0001EF(N\x0001EF(S\x0001EF(V\x0001EK(1\x0001EK(E\x0001EK(F\x0001EK(N\x0001EK(S\x0001EK(V\x0001EK1;\x0001EK1C\x0001EK1O\x0001EK1T\x0001EK1U\x0001EKF(\x0001EKN;\x0001EKNC\x0001EKNE\x0001EKNT\x0001EKNU\x0001EKOK\x0001EKS;\x0001EKSC\x0001EKSO\x0001EKST\x0001EKSU\x0001EKU(\x0001EKU1\x0001EKUE\x0001EKUF\x0001EKUS\x0001EKUV\x0001EKV;\x0001EKVC\x0001EKVO\x0001EKVT\x0001EKVU\x0001EN;T\x0001ENC\x0001ENEN\x0001ENO(\x0001ENOF\x0001ENOS\x0001ENOV\x0001ENT(\x0001ENT1\x0001ENTF\x0001ENTN\x0001ENTS\x0001ENTV\x0001ENUE\x0001EOKN\x0001ES;T\x0001ESC\x0001ESO(\x0001ESO1\x0001ESOF\x0001ESON\x0001ESOS\x0001ESOV\x0001EST(\x0001EST1\x0001ESTF\x0001ESTN\x0001ESTS\x0001ESTV\x0001ESUE\x0001EU(1\x0001EU(F\x0001EU(N\x0001EU(S\x0001EU(V\x0001EU1,\x0001EU1C\x0001EU1O\x0001EUEF\x0001EUEK\x0001EUF(\x0001EUS,\x0001EUSC\x0001EUSO\x0001EUV,\x0001EUVC\x0001EUVO\x0001EV;T\x0001EVC\x0001EVO(\x0001EVOF\x0001EVOS\x0001EVT(\x0001EVT1\x0001EVTF\x0001EVTN\x0001EVTS\x0001EVTV\x0001EVUE\x0001F()1\x0001F()F\x0001F()K\x0001F()N\x0001F()O\x0001F()S\x0001F()U\x0001F()V\x0001F(1)\x0001F(1N\x0001F(1O\x0001F(E(\x0001F(E1\x0001F(EF\x0001F(EK\x0001F(EN\x0001F(ES\x0001F(EV\x0001F(F(\x0001F(N)\x0001F(N,\x0001F(NO\x0001F(S)\x0001F(SO\x0001F(V)\x0001F(VO\x0001K(1O\x0001K(F(\x0001K(N)\x0001K(NO\x0001K(S)\x0001K(SO\x0001K(V)\x0001K(VO\x0001K)&(\x0001K)&1\x0001K)&F\x0001K)&N\x0001K)&S\x0001K)&V\x0001K);E\x0001K);T\x0001K)B(\x0001K)B1\x0001K)BF\x0001K)BN\x0001K)BS\x0001K)BV\x0001K)E(\x0001K)E1\x0001K)EF\x0001K)EK\x0001K)EN\x0001K)ES\x0001K)EV\x0001K)F(\x0001K)O(\x0001K)OF\x0001K)UE\x0001K1\x0001K1&(\x0001K1&1\x0001K1&F\x0001K1&N\x0001K1&S\x0001K1&V\x0001K1;\x0001K1;C\x0001K1;E\x0001K1;T\x0001K1B(\x0001K1B1\x0001K1BF\x0001K1BN\x0001K1BS\x0001K1BV\x0001K1C\x0001K1E(\x0001K1E1\x0001K1EF\x0001K1EK\x0001K1EN\x0001K1ES\x0001K1EV\x0001K1O(\x0001K1OF\x0001K1OS\x0001K1OV\x0001K1U(\x0001K1UE\x0001KF()\x0001KF(1\x0001KF(F\x0001KF(N\x0001KF(S\x0001KF(V\x0001KN\x0001KN&(\x0001KN&1\x0001KN&F\x0001KN&N\x0001KN&S\x0001KN&V\x0001KN;\x0001KN;C\x0001KN;E\x0001KN;T\x0001KNB(\x0001KNB1\x0001KNBF\x0001KNBN\x0001KNBS\x0001KNBV\x0001KNC\x0001KNE(\x0001KNE1\x0001KNEF\x0001KNEN\x0001KNES\x0001KNEV\x0001KNU(\x0001KNUE\x0001KS\x0001KS&(\x0001KS&1\x0001KS&F\x0001KS&N\x0001KS&S\x0001KS&V\x0001KS;\x0001KS;C\x0001KS;E\x0001KS;T\x0001KSB(\x0001KSB1\x0001KSBF\x0001KSBN\x0001KSBS\x0001KSBV\x0001KSC\x0001KSE(\x0001KSE1\x0001KSEF\x0001KSEK\x0001KSEN\x0001KSES\x0001KSEV\x0001KSO(\x0001KSO1\x0001KSOF\x0001KSON\x0001KSOS\x0001KSOV\x0001KSU(\x0001KSUE\x0001KUE(\x0001KUE1\x0001KUEF\x0001KUEK\x0001KUEN\x0001KUES\x0001KUEV\x0001KV\x0001KV&(\x0001KV&1\x0001KV&F\x0001KV&N\x0001KV&S\x0001KV&V\x0001KV;\x0001KV;C\x0001KV;E\x0001KV;T\x0001KVB(\x0001KVB1\x0001KVBF\x0001KVBN\x0001KVBS\x0001KVBV\x0001KVC\x0001KVE(\x0001KVE1\x0001KVEF\x0001KVEK\x0001KVEN\x0001KVES\x0001KVEV\x0001KVO(\x0001KVOF\x0001KVOS\x0001KVU(\x0001KVUE\x0001N&F(\x0001N(1O\x0001N(F(\x0001N(S)\x0001N(SO\x0001N(V)\x0001N(VO\x0001N)UE\x0001N,F(\x0001NE(1\x0001NE(F\x0001NE(N\x0001NE(S\x0001NE(V\x0001NE1C\x0001NE1O\x0001NEF(\x0001NENC\x0001NENO\x0001NESC\x0001NESO\x0001NEVC\x0001NEVO\x0001NU(E\x0001NUE\x0001NUE(\x0001NUE1\x0001NUE;\x0001NUEC\x0001NUEF\x0001NUEK\x0001NUEN\x0001NUES\x0001NUEV\x0001O(1&\x0001O(1)\x0001O(1,\x0001O(1O\x0001O(E(\x0001O(E1\x0001O(EE\x0001O(EF\x0001O(EK\x0001O(EN\x0001O(EO\x0001O(ES\x0001O(EV\x0001O(F(\x0001O(N&\x0001O(N)\x0001O(N,\x0001O(NO\x0001O(S&\x0001O(S)\x0001O(S,\x0001O(SO\x0001O(V&\x0001O(V)\x0001O(V,\x0001O(VO\x0001OF()\x0001OF(1\x0001OF(E\x0001OF(F\x0001OF(N\x0001OF(S\x0001OF(V\x0001OK&(\x0001OK&1\x0001OK&F\x0001OK&N\x0001OK&S\x0001OK&V\x0001OK(1\x0001OK(F\x0001OK(N\x0001OK(S\x0001OK(V\x0001OK1C\x0001OK1O\x0001OKF(\x0001OKNC\x0001OKO(\x0001OKO1\x0001OKOF\x0001OKON\x0001OKOS\x0001OKOV\x0001OKSC\x0001OKSO\x0001OKVC\x0001OKVO\x0001ONSU\x0001OS&(\x0001OS&1\x0001OS&E\x0001OS&F\x0001OS&K\x0001OS&N\x0001OS&S\x0001OS&U\x0001OS&V\x0001OS(E\x0001OS(U\x0001OS)&\x0001OS),\x0001OS);\x0001OS)B\x0001OS)C\x0001OS)E\x0001OS)F\x0001OS)K\x0001OS)O\x0001OS)U\x0001OS,(\x0001OS,F\x0001OS1(\x0001OS1F\x0001OS1N\x0001OS1S\x0001OS1U\x0001OS1V\x0001OS;\x0001OS;C\x0001OS;E\x0001OS;N\x0001OS;T\x0001OSA(\x0001OSAF\x0001OSAS\x0001OSAT\x0001OSAV\x0001OSB(\x0001OSB1\x0001OSBE\x0001OSBF\x0001OSBN\x0001OSBS\x0001OSBV\x0001OSC\x0001OSE(\x0001OSE1\x0001OSEF\x0001OSEK\x0001OSEN\x0001OSEO\x0001OSES\x0001OSEU\x0001OSEV\x0001OSF(\x0001OSK(\x0001OSK)\x0001OSK1\x0001OSKB\x0001OSKF\x0001OSKN\x0001OSKS\x0001OSKU\x0001OSKV\x0001OST(\x0001OST1\x0001OSTE\x0001OSTF\x0001OSTN\x0001OSTS\x0001OSTT\x0001OSTV\x0001OSU\x0001OSU(\x0001OSU1\x0001OSU;\x0001OSUC\x0001OSUE\x0001OSUF\x0001OSUK\x0001OSUO\x0001OSUS\x0001OSUT\x0001OSUV\x0001OSV(\x0001OSVF\x0001OSVO\x0001OSVS\x0001OSVU\x0001OU(E\x0001OUEK\x0001OUEN\x0001OV\x0001OV&(\x0001OV&1\x0001OV&E\x0001OV&F\x0001OV&K\x0001OV&N\x0001OV&S\x0001OV&U\x0001OV&V\x0001OV(E\x0001OV(U\x0001OV)&\x0001OV),\x0001OV);\x0001OV)B\x0001OV)C\x0001OV)E\x0001OV)F\x0001OV)K\x0001OV)O\x0001OV)U\x0001OV,(\x0001OV,F\x0001OV;\x0001OV;C\x0001OV;E\x0001OV;N\x0001OV;T\x0001OVA(\x0001OVAF\x0001OVAS\x0001OVAT\x0001OVAV\x0001OVB(\x0001OVB1\x0001OVBE\x0001OVBF\x0001OVBN\x0001OVBS\x0001OVBV\x0001OVC\x0001OVE(\x0001OVE1\x0001OVEF\x0001OVEK\x0001OVEN\x0001OVEO\x0001OVES\x0001OVEU\x0001OVEV\x0001OVF(\x0001OVK(\x0001OVK)\x0001OVK1\x0001OVKB\x0001OVKF\x0001OVKN\x0001OVKS\x0001OVKU\x0001OVKV\x0001OVO(\x0001OVOF\x0001OVOK\x0001OVOS\x0001OVOU\x0001OVS(\x0001OVS1\x0001OVSF\x0001OVSO\x0001OVSU\x0001OVSV\x0001OVT(\x0001OVT1\x0001OVTE\x0001OVTF\x0001OVTN\x0001OVTS\x0001OVTT\x0001OVTV\x0001OVU\x0001OVU(\x0001OVU1\x0001OVU;\x0001OVUC\x0001OVUE\x0001OVUF\x0001OVUK\x0001OVUO\x0001OVUS\x0001OVUT\x0001OVUV\x0001SF()\x0001SF(1\x0001SF(F\x0001SF(N\x0001SF(S\x0001SF(V\x0001SUE\x0001SUE;\x0001SUEC\x0001SUEK\x0001SV\x0001SV;\x0001SV;C\x0001SVC\x0001SVO(\x0001SVOF\x0001SVOS\x0001T(1)\x0001T(1O\x0001T(F(\x0001T(N)\x0001T(NO\x0001T(S)\x0001T(SO\x0001T(V)\x0001T(VO\x0001T1(F\x0001T1O(\x0001T1OF\x0001T1OS\x0001T1OV\x0001TE(1\x0001TE(F\x0001TE(N\x0001TE(S\x0001TE(V\x0001TE1N\x0001TE1O\x0001TEF(\x0001TEK(\x0001TEK1\x0001TEKF\x0001TEKN\x0001TEKS\x0001TEKV\x0001TENN\x0001TENO\x0001TESN\x0001TESO\x0001TEVN\x0001TEVO\x0001TF()\x0001TF(1\x0001TF(F\x0001TF(N\x0001TF(S\x0001TF(V\x0001TN(1\x0001TN(F\x0001TN(S\x0001TN(V\x0001TN1C\x0001TN1O\x0001TN;E\x0001TN;N\x0001TN;T\x0001TNE(\x0001TNE1\x0001TNEF\x0001TNEN\x0001TNES\x0001TNEV\x0001TNF(\x0001TNKN\x0001TNN:\x0001TNNC\x0001TNNO\x0001TNO(\x0001TNOF\x0001TNOS\x0001TNOV\x0001TNSC\x0001TNSO\x0001TNT(\x0001TNT1\x0001TNTF\x0001TNTN\x0001TNTS\x0001TNTV\x0001TNVC\x0001TNVO\x0001TS(F\x0001TSO(\x0001TSO1\x0001TSOF\x0001TSON\x0001TSOS\x0001TSOV\x0001TTNE\x0001TTNK\x0001TTNN\x0001TTNT\x0001TV(1\x0001TV(F\x0001TVO(\x0001TVOF\x0001TVOS\x0001U\x0001U(1)\x0001U(1O\x0001U(E(\x0001U(E1\x0001U(EF\x0001U(EK\x0001U(EN\x0001U(ES\x0001U(EV\x0001U(F(\x0001U(N)\x0001U(NO\x0001U(S)\x0001U(SO\x0001U(V)\x0001U(VO\x0001U1,(\x0001U1,F\x0001U1C\x0001U1O(\x0001U1OF\x0001U1OS\x0001U1OV\x0001U;\x0001U;C\x0001UC\x0001UE\x0001UE(1\x0001UE(E\x0001UE(F\x0001UE(N\x0001UE(O\x0001UE(S\x0001UE(V\x0001UE1\x0001UE1&\x0001UE1(\x0001UE1)\x0001UE1,\x0001UE1;\x0001UE1B\x0001UE1C\x0001UE1F\x0001UE1K\x0001UE1N\x0001UE1O\x0001UE1S\x0001UE1U\x0001UE1V\x0001UE;\x0001UE;C\x0001UEC\x0001UEF\x0001UEF(\x0001UEF,\x0001UEF;\x0001UEFC\x0001UEK\x0001UEK(\x0001UEK1\x0001UEK;\x0001UEKC\x0001UEKF\x0001UEKN\x0001UEKO\x0001UEKS\x0001UEKV\x0001UEN\x0001UEN&\x0001UEN(\x0001UEN)\x0001UEN,\x0001UEN1\x0001UEN;\x0001UENB\x0001UENC\x0001UENF\x0001UENK\x0001UENN\x0001UENO\x0001UENS\x0001UENU\x0001UEOK\x0001UEON\x0001UES\x0001UES&\x0001UES(\x0001UES)\x0001UES,\x0001UES1\x0001UES;\x0001UESB\x0001UESC\x0001UESF\x0001UESK\x0001UESO\x0001UESU\x0001UESV\x0001UEV\x0001UEV&\x0001UEV(\x0001UEV)\x0001UEV,\x0001UEV;\x0001UEVB\x0001UEVC\x0001UEVF\x0001UEVK\x0001UEVN\x0001UEVO\x0001UEVS\x0001UEVU\x0001UF()\x0001UF(1\x0001UF(F\x0001UF(N\x0001UF(S\x0001UF(V\x0001UK(E\x0001UO(E\x0001UON(\x0001UON1\x0001UONF\x0001UONS\x0001US,(\x0001US,F\x0001USC\x0001USO(\x0001USO1\x0001USOF\x0001USON\x0001USOS\x0001USOV\x0001UTN(\x0001UTN1\x0001UTNF\x0001UTNN\x0001UTNS\x0001UTNV\x0001UV,(\x0001UV,F\x0001UVC\x0001UVO(\x0001UVOF\x0001UVOS\x0001VF()\x0001VF(1\x0001VF(F\x0001VF(N\x0001VF(S\x0001VF(V\x0001VO(1\x0001VO(F\x0001VO(N\x0001VO(S\x0001VO(V\x0001VOF(\x0001VOS(\x0001VOS1\x0001VOSF\x0001VOSU\x0001VOSV\x0001VS\x0001VS;\x0001VS;C\x0001VSC\x0001VSO(\x0001VSO1\x0001VSOF\x0001VSON\x0001VSOS\x0001VSOV\x0001VUE\x0001VUE;\x0001VUEC\x0001VUEK\x000;T(EF\x000;T(EK\x000;TKNC\x000E(1&(\x000E(1&1\x000E(1&F\x000E(1&N\x000E(1&S\x000E(1&V\x000E(1)&\x000E(1),\x000E(1)1\x000E(1);\x000E(1)B\x000E(1)C\x000E(1)F\x000E(1)K\x000E(1)N\x000E(1)O\x000E(1)S\x000E(1)U\x000E(1)V\x000E(1,F\x000E(1F(\x000E(1N)\x000E(1O(\x000E(1OF\x000E(1OS\x000E(1OV\x000E(1S)\x000E(1V)\x000E(1VO\x000E(E(1\x000E(E(E\x000E(E(F\x000E(E(N\x000E(E(S\x000E(E(V\x000E(E1&\x000E(E1)\x000E(E1O\x000E(EF(\x000E(EK(\x000E(EK1\x000E(EKF\x000E(EKN\x000E(EKS\x000E(EKV\x000E(EN&\x000E(EN)\x000E(ENO\x000E(ES&\x000E(ES)\x000E(ESO\x000E(EV&\x000E(EV)\x000E(EVO\x000E(F()\x000E(F(1\x000E(F(E\x000E(F(F\x000E(F(N\x000E(F(S\x000E(F(V\x000E(N&(\x000E(N&1\x000E(N&F\x000E(N&N\x000E(N&S\x000E(N&V\x000E(N(1\x000E(N(F\x000E(N(S\x000E(N(V\x000E(N)&\x000E(N),\x000E(N)1\x000E(N);\x000E(N)B\x000E(N)C\x000E(N)F\x000E(N)K\x000E(N)N\x000E(N)O\x000E(N)S\x000E(N)U\x000E(N)V\x000E(N,F\x000E(N1)\x000E(N1O\x000E(NF(\x000E(NO(\x000E(NOF\x000E(NOS\x000E(NOV\x000E(S&(\x000E(S&1\x000E(S&F\x000E(S&N\x000E(S&S\x000E(S&V\x000E(S)&\x000E(S),\x000E(S)1\x000E(S);\x000E(S)B\x000E(S)C\x000E(S)F\x000E(S)K\x000E(S)N\x000E(S)O\x000E(S)S\x000E(S)U\x000E(S)V\x000E(S,F\x000E(S1)\x000E(SF(\x000E(SO(\x000E(SO1\x000E(SOF\x000E(SON\x000E(SOS\x000E(SOV\x000E(SV)\x000E(SVO\x000E(V&(\x000E(V&1\x000E(V&F\x000E(V&N\x000E(V&S\x000E(V&V\x000E(V)&\x000E(V),\x000E(V)1\x000E(V);\x000E(V)B\x000E(V)C\x000E(V)F\x000E(V)K\x000E(V)N\x000E(V)O\x000E(V)S\x000E(V)U\x000E(V)V\x000E(V,F\x000E(VF(\x000E(VO(\x000E(VOF\x000E(VOS\x000E(VS)\x000E(VSO\x000E1&(1\x000E1&(E\x000E1&(F\x000E1&(N\x000E1&(S\x000E1&(V\x000E1&1)\x000E1&1O\x000E1&F(\x000E1&N)\x000E1&NO\x000E1&S)\x000E1&SO\x000E1&V)\x000E1&VO\x000E1)\x000E1)&(\x000E1)&1\x000E1)&F\x000E1)&N\x000E1)&S\x000E1)&V\x000E1);\x000E1);(\x000E1);C\x000E1);E\x000E1);T\x000E1)C\x000E1)KN\x000E1)O(\x000E1)O1\x000E1)OF\x000E1)ON\x000E1)OS\x000E1)OV\x000E1)UE\x000E1,(1\x000E1,(F\x000E1,(N\x000E1,(S\x000E1,(V\x000E1,F(\x000E1;(E\x000E1B(1\x000E1B(F\x000E1B(N\x000E1B(S\x000E1B(V\x000E1B1)\x000E1B1O\x000E1BF(\x000E1BN)\x000E1BNO\x000E1BS)\x000E1BSO\x000E1BV)\x000E1BVO\x000E1F()\x000E1F(1\x000E1F(F\x000E1F(N\x000E1F(S\x000E1F(V\x000E1K(1\x000E1K(E\x000E1K(F\x000E1K(N\x000E1K(S\x000E1K(V\x000E1K1)\x000E1K1K\x000E1K1O\x000E1KF(\x000E1KN\x000E1KN)\x000E1KN;\x000E1KNC\x000E1KNK\x000E1KNU\x000E1KS)\x000E1KSK\x000E1KSO\x000E1KV)\x000E1KVK\x000E1KVO\x000E1N)U\x000E1N;\x000E1N;C\x000E1NC\x000E1NKN\x000E1O(1\x000E1O(E\x000E1O(F\x000E1O(N\x000E1O(S\x000E1O(V\x000E1OF(\x000E1OS&\x000E1OS(\x000E1OS)\x000E1OS,\x000E1OS1\x000E1OS;\x000E1OSB\x000E1OSF\x000E1OSK\x000E1OSU\x000E1OSV\x000E1OV&\x000E1OV(\x000E1OV)\x000E1OV,\x000E1OV;\x000E1OVB\x000E1OVF\x000E1OVK\x000E1OVO\x000E1OVS\x000E1OVU\x000E1S;\x000E1S;C\x000E1SC\x000E1U(E\x000E1UE(\x000E1UE1\x000E1UEF\x000E1UEK\x000E1UEN\x000E1UES\x000E1UEV\x000E1V\x000E1V;\x000E1V;C\x000E1VC\x000E1VO(\x000E1VOF\x000E1VOS\x000EE(F(\x000EEK(F\x000EF()&\x000EF(),\x000EF()1\x000EF();\x000EF()B\x000EF()F\x000EF()K\x000EF()N\x000EF()O\x000EF()S\x000EF()U\x000EF()V\x000EF(1&\x000EF(1)\x000EF(1,\x000EF(1O\x000EF(E(\x000EF(E1\x000EF(EF\x000EF(EK\x000EF(EN\x000EF(ES\x000EF(EV\x000EF(F(\x000EF(N&\x000EF(N)\x000EF(N,\x000EF(NO\x000EF(O)\x000EF(S&\x000EF(S)\x000EF(S,\x000EF(SO\x000EF(V&\x000EF(V)\x000EF(V,\x000EF(VO\x000EK(1&\x000EK(1(\x000EK(1)\x000EK(1,\x000EK(1F\x000EK(1N\x000EK(1O\x000EK(1S\x000EK(1V\x000EK(E(\x000EK(E1\x000EK(EF\x000EK(EK\x000EK(EN\x000EK(ES\x000EK(EV\x000EK(F(\x000EK(N&\x000EK(N(\x000EK(N)\x000EK(N,\x000EK(N1\x000EK(NF\x000EK(NO\x000EK(S&\x000EK(S(\x000EK(S)\x000EK(S,\x000EK(S1\x000EK(SF\x000EK(SO\x000EK(SV\x000EK(V&\x000EK(V(\x000EK(V)\x000EK(V,\x000EK(VF\x000EK(VO\x000EK(VS\x000EK1&(\x000EK1&1\x000EK1&F\x000EK1&N\x000EK1&S\x000EK1&V\x000EK1)\x000EK1)&\x000EK1);\x000EK1)C\x000EK1)K\x000EK1)O\x000EK1)U\x000EK1,(\x000EK1,F\x000EK1;(\x000EK1B(\x000EK1B1\x000EK1BF\x000EK1BN\x000EK1BS\x000EK1BV\x000EK1F(\x000EK1K(\x000EK1K1\x000EK1KF\x000EK1KN\x000EK1KS\x000EK1KV\x000EK1N\x000EK1N)\x000EK1N;\x000EK1NC\x000EK1NK\x000EK1O(\x000EK1OF\x000EK1OS\x000EK1OV\x000EK1S\x000EK1S;\x000EK1SC\x000EK1SF\x000EK1SK\x000EK1U(\x000EK1UE\x000EK1V\x000EK1V;\x000EK1VC\x000EK1VF\x000EK1VK\x000EK1VO\x000EKE(F\x000EKEK(\x000EKF()\x000EKF(1\x000EKF(E\x000EKF(F\x000EKF(N\x000EKF(O\x000EKF(S\x000EKF(V\x000EKN&(\x000EKN&1\x000EKN&F\x000EKN&N\x000EKN&S\x000EKN&V\x000EKN(1\x000EKN(F\x000EKN(S\x000EKN(V\x000EKN)\x000EKN)&\x000EKN);\x000EKN)C\x000EKN)K\x000EKN)O\x000EKN)U\x000EKN,(\x000EKN,F\x000EKN1\x000EKN1;\x000EKN1C\x000EKN1K\x000EKN1O\x000EKN;(\x000EKNB(\x000EKNB1\x000EKNBF\x000EKNBN\x000EKNBS\x000EKNBV\x000EKNF(\x000EKNK(\x000EKNK1\x000EKNKF\x000EKNKN\x000EKNKS\x000EKNKV\x000EKNU(\x000EKNUE\x000EKO(1\x000EKO(F\x000EKO(N\x000EKO(S\x000EKO(V\x000EKOK(\x000EKOKN\x000EKS&(\x000EKS&1\x000EKS&F\x000EKS&N\x000EKS&S\x000EKS&V\x000EKS)\x000EKS)&\x000EKS);\x000EKS)C\x000EKS)K\x000EKS)O\x000EKS)U\x000EKS,(\x000EKS,F\x000EKS1\x000EKS1;\x000EKS1C\x000EKS1F\x000EKS1K\x000EKS;(\x000EKSB(\x000EKSB1\x000EKSBF\x000EKSBN\x000EKSBS\x000EKSBV\x000EKSF(\x000EKSK(\x000EKSK1\x000EKSKF\x000EKSKN\x000EKSKS\x000EKSKV\x000EKSO(\x000EKSO1\x000EKSOF\x000EKSON\x000EKSOS\x000EKSOV\x000EKSU(\x000EKSUE\x000EKSV\x000EKSV;\x000EKSVC\x000EKSVF\x000EKSVK\x000EKSVO\x000EKV&(\x000EKV&1\x000EKV&F\x000EKV&N\x000EKV&S\x000EKV&V\x000EKV)\x000EKV)&\x000EKV);\x000EKV)C\x000EKV)K\x000EKV)O\x000EKV)U\x000EKV,(\x000EKV,F\x000EKV;(\x000EKVB(\x000EKVB1\x000EKVBF\x000EKVBN\x000EKVBS\x000EKVBV\x000EKVF(\x000EKVK(\x000EKVK1\x000EKVKF\x000EKVKN\x000EKVKS\x000EKVKV\x000EKVO(\x000EKVOF\x000EKVOS\x000EKVS\x000EKVS;\x000EKVSC\x000EKVSF\x000EKVSK\x000EKVSO\x000EKVU(\x000EKVUE\x000EN&(1\x000EN&(E\x000EN&(F\x000EN&(N\x000EN&(S\x000EN&(V\x000EN&1)\x000EN&1O\x000EN&F(\x000EN&N)\x000EN&NO\x000EN&S)\x000EN&SO\x000EN&V)\x000EN&VO\x000EN(1O\x000EN(F(\x000EN(S)\x000EN(SO\x000EN(V)\x000EN(VO\x000EN)\x000EN)&(\x000EN)&1\x000EN)&F\x000EN)&N\x000EN)&S\x000EN)&V\x000EN);\x000EN);(\x000EN);C\x000EN);E\x000EN);T\x000EN)C\x000EN)KN\x000EN)O(\x000EN)O1\x000EN)OF\x000EN)ON\x000EN)OS\x000EN)OV\x000EN)UE\x000EN,(1\x000EN,(F\x000EN,(N\x000EN,(S\x000EN,(V\x000EN,F(\x000EN1;\x000EN1;C\x000EN1O(\x000EN1OF\x000EN1OS\x000EN1OV\x000EN;(E\x000ENB(1\x000ENB(F\x000ENB(N\x000ENB(S\x000ENB(V\x000ENB1)\x000ENB1O\x000ENBF(\x000ENBN)\x000ENBNO\x000ENBS)\x000ENBSO\x000ENBV)\x000ENBVO\x000ENF()\x000ENF(1\x000ENF(F\x000ENF(N\x000ENF(S\x000ENF(V\x000ENK(1\x000ENK(E\x000ENK(F\x000ENK(N\x000ENK(S\x000ENK(V\x000ENK1)\x000ENK1K\x000ENK1O\x000ENKF(\x000ENKN)\x000ENKN,\x000ENKN;\x000ENKNB\x000ENKNC\x000ENKNK\x000ENKNU\x000ENKS)\x000ENKSK\x000ENKSO\x000ENKV)\x000ENKVK\x000ENKVO\x000ENO(1\x000ENO(E\x000ENO(F\x000ENO(N\x000ENO(S\x000ENO(V\x000ENOF(\x000ENOS&\x000ENOS(\x000ENOS)\x000ENOS,\x000ENOS1\x000ENOS;\x000ENOSB\x000ENOSF\x000ENOSK\x000ENOSU\x000ENOSV\x000ENOV&\x000ENOV(\x000ENOV)\x000ENOV,\x000ENOV;\x000ENOVB\x000ENOVF\x000ENOVK\x000ENOVO\x000ENOVS\x000ENOVU\x000ENU(E\x000ENUE(\x000ENUE1\x000ENUEF\x000ENUEK\x000ENUEN\x000ENUES\x000ENUEV\x000EOK(E\x000EOKNK\x000ES&(1\x000ES&(E\x000ES&(F\x000ES&(N\x000ES&(S\x000ES&(V\x000ES&1)\x000ES&1O\x000ES&F(\x000ES&N)\x000ES&NO\x000ES&S)\x000ES&SO\x000ES&V)\x000ES&VO\x000ES)\x000ES)&(\x000ES)&1\x000ES)&F\x000ES)&N\x000ES)&S\x000ES)&V\x000ES);\x000ES);(\x000ES);C\x000ES);E\x000ES);T\x000ES)C\x000ES)KN\x000ES)O(\x000ES)O1\x000ES)OF\x000ES)ON\x000ES)OS\x000ES)OV\x000ES)UE\x000ES,(1\x000ES,(F\x000ES,(N\x000ES,(S\x000ES,(V\x000ES,F(\x000ES1\x000ES1;\x000ES1;C\x000ES1C\x000ES;(E\x000ESB(1\x000ESB(F\x000ESB(N\x000ESB(S\x000ESB(V\x000ESB1)\x000ESB1O\x000ESBF(\x000ESBN)\x000ESBNO\x000ESBS)\x000ESBSO\x000ESBV)\x000ESBVO\x000ESF()\x000ESF(1\x000ESF(F\x000ESF(N\x000ESF(S\x000ESF(V\x000ESK(1\x000ESK(E\x000ESK(F\x000ESK(N\x000ESK(S\x000ESK(V\x000ESK1)\x000ESK1K\x000ESK1O\x000ESKF(\x000ESKN\x000ESKN)\x000ESKN;\x000ESKNC\x000ESKNK\x000ESKNU\x000ESKS)\x000ESKSK\x000ESKSO\x000ESKV)\x000ESKVK\x000ESKVO\x000ESO(1\x000ESO(E\x000ESO(F\x000ESO(N\x000ESO(S\x000ESO(V\x000ESO1&\x000ESO1(\x000ESO1)\x000ESO1,\x000ESO1;\x000ESO1B\x000ESO1F\x000ESO1K\x000ESO1N\x000ESO1S\x000ESO1U\x000ESO1V\x000ESOF(\x000ESON&\x000ESON(\x000ESON)\x000ESON,\x000ESON1\x000ESON;\x000ESONB\x000ESONF\x000ESONK\x000ESONU\x000ESOS&\x000ESOS(\x000ESOS)\x000ESOS,\x000ESOS1\x000ESOS;\x000ESOSB\x000ESOSF\x000ESOSK\x000ESOSU\x000ESOSV\x000ESOV&\x000ESOV(\x000ESOV)\x000ESOV,\x000ESOV;\x000ESOVB\x000ESOVF\x000ESOVK\x000ESOVO\x000ESOVS\x000ESOVU\x000ESU(E\x000ESUE(\x000ESUE1\x000ESUEF\x000ESUEK\x000ESUEN\x000ESUES\x000ESUEV\x000ESV\x000ESV;\x000ESV;C\x000ESVC\x000ESVO(\x000ESVOF\x000ESVOS\x000EV&(1\x000EV&(E\x000EV&(F\x000EV&(N\x000EV&(S\x000EV&(V\x000EV&1)\x000EV&1O\x000EV&F(\x000EV&N)\x000EV&NO\x000EV&S)\x000EV&SO\x000EV&V)\x000EV&VO\x000EV)\x000EV)&(\x000EV)&1\x000EV)&F\x000EV)&N\x000EV)&S\x000EV)&V\x000EV);\x000EV);(\x000EV);C\x000EV);E\x000EV);T\x000EV)C\x000EV)KN\x000EV)O(\x000EV)O1\x000EV)OF\x000EV)ON\x000EV)OS\x000EV)OV\x000EV)UE\x000EV,(1\x000EV,(F\x000EV,(N\x000EV,(S\x000EV,(V\x000EV,F(\x000EV;(E\x000EVB(1\x000EVB(F\x000EVB(N\x000EVB(S\x000EVB(V\x000EVB1)\x000EVB1O\x000EVBF(\x000EVBN)\x000EVBNO\x000EVBS)\x000EVBSO\x000EVBV)\x000EVBVO\x000EVF()\x000EVF(1\x000EVF(F\x000EVF(N\x000EVF(S\x000EVF(V\x000EVK(1\x000EVK(E\x000EVK(F\x000EVK(N\x000EVK(S\x000EVK(V\x000EVK1)\x000EVK1K\x000EVK1O\x000EVKF(\x000EVKN\x000EVKN)\x000EVKN;\x000EVKNC\x000EVKNK\x000EVKNU\x000EVKS)\x000EVKSK\x000EVKSO\x000EVKV)\x000EVKVK\x000EVKVO\x000EVN\x000EVN)U\x000EVN;\x000EVN;C\x000EVNC\x000EVNKN\x000EVNO(\x000EVNOF\x000EVNOS\x000EVNOV\x000EVO(1\x000EVO(E\x000EVO(F\x000EVO(N\x000EVO(S\x000EVO(V\x000EVOF(\x000EVOS&\x000EVOS(\x000EVOS)\x000EVOS,\x000EVOS1\x000EVOS;\x000EVOSB\x000EVOSF\x000EVOSK\x000EVOSU\x000EVOSV\x000EVS\x000EVS;\x000EVS;C\x000EVSC\x000EVSO(\x000EVSO1\x000EVSOF\x000EVSON\x000EVSOS\x000EVSOV\x000EVU(E\x000EVUE(\x000EVUE1\x000EVUEF\x000EVUEK\x000EVUEN\x000EVUES\x000EVUEV\x000F()&(\x000F()&1\x000F()&E\x000F()&F\x000F()&K\x000F()&N\x000F()&S\x000F()&V\x000F(),(\x000F(),1\x000F(),F\x000F(),N\x000F(),S\x000F(),V\x000F()1(\x000F()1F\x000F()1N\x000F()1O\x000F()1S\x000F()1U\x000F()1V\x000F();E\x000F();N\x000F();T\x000F()A(\x000F()AF\x000F()AS\x000F()AT\x000F()AV\x000F()B(\x000F()B1\x000F()BE\x000F()BF\x000F()BN\x000F()BS\x000F()BV\x000F()C\x000F()E(\x000F()E1\x000F()EF\x000F()EK\x000F()EN\x000F()EO\x000F()ES\x000F()EU\x000F()EV\x000F()F(\x000F()K(\x000F()K)\x000F()K1\x000F()KF\x000F()KN\x000F()KS\x000F()KU\x000F()KV\x000F()N&\x000F()N(\x000F()N)\x000F()N,\x000F()N1\x000F()NE\x000F()NF\x000F()NO\x000F()NU\x000F()O(\x000F()O1\x000F()OF\x000F()OK\x000F()ON\x000F()OS\x000F()OU\x000F()OV\x000F()S(\x000F()S1\x000F()SF\x000F()SO\x000F()SU\x000F()SV\x000F()T(\x000F()T1\x000F()TE\x000F()TF\x000F()TN\x000F()TS\x000F()TT\x000F()TV\x000F()U\x000F()U(\x000F()U1\x000F()U;\x000F()UC\x000F()UE\x000F()UF\x000F()UK\x000F()UO\x000F()US\x000F()UT\x000F()UV\x000F()V(\x000F()VF\x000F()VO\x000F()VS\x000F()VU\x000F(1&(\x000F(1&1\x000F(1&F\x000F(1&N\x000F(1&S\x000F(1&V\x000F(1)\x000F(1)&\x000F(1),\x000F(1)1\x000F(1);\x000F(1)A\x000F(1)B\x000F(1)C\x000F(1)E\x000F(1)F\x000F(1)K\x000F(1)N\x000F(1)O\x000F(1)S\x000F(1)T\x000F(1)U\x000F(1)V\x000F(1,(\x000F(1,F\x000F(1O(\x000F(1OF\x000F(1OS\x000F(1OV\x000F(E(1\x000F(E(E\x000F(E(F\x000F(E(N\x000F(E(S\x000F(E(V\x000F(E1&\x000F(E1)\x000F(E1K\x000F(E1O\x000F(EF(\x000F(EK(\x000F(EK1\x000F(EKF\x000F(EKN\x000F(EKO\x000F(EKS\x000F(EKV\x000F(EN&\x000F(EN)\x000F(ENK\x000F(ENO\x000F(EOK\x000F(ES&\x000F(ES)\x000F(ESK\x000F(ESO\x000F(EV&\x000F(EV)\x000F(EVK\x000F(EVO\x000F(F()\x000F(F(1\x000F(F(E\x000F(F(F\x000F(F(N\x000F(F(S\x000F(F(V\x000F(K()\x000F(K,(\x000F(K,F\x000F(N&(\x000F(N&1\x000F(N&F\x000F(N&N\x000F(N&S\x000F(N&V\x000F(N)\x000F(N)&\x000F(N),\x000F(N)1\x000F(N);\x000F(N)A\x000F(N)B\x000F(N)C\x000F(N)E\x000F(N)F\x000F(N)K\x000F(N)N\x000F(N)O\x000F(N)S\x000F(N)T\x000F(N)U\x000F(N)V\x000F(N,(\x000F(N,F\x000F(NO(\x000F(NOF\x000F(NOS\x000F(NOV\x000F(S&(\x000F(S&1\x000F(S&F\x000F(S&N\x000F(S&S\x000F(S&V\x000F(S)\x000F(S)&\x000F(S),\x000F(S)1\x000F(S);\x000F(S)A\x000F(S)B\x000F(S)C\x000F(S)E\x000F(S)F\x000F(S)K\x000F(S)N\x000F(S)O\x000F(S)S\x000F(S)T\x000F(S)U\x000F(S)V\x000F(S,(\x000F(S,F\x000F(SO(\x000F(SO1\x000F(SOF\x000F(SON\x000F(SOS\x000F(SOV\x000F(T,(\x000F(T,F\x000F(V&(\x000F(V&1\x000F(V&F\x000F(V&N\x000F(V&S\x000F(V&V\x000F(V)\x000F(V)&\x000F(V),\x000F(V)1\x000F(V);\x000F(V)A\x000F(V)B\x000F(V)C\x000F(V)E\x000F(V)F\x000F(V)K\x000F(V)N\x000F(V)O\x000F(V)S\x000F(V)T\x000F(V)U\x000F(V)V\x000F(V,(\x000F(V,F\x000F(VO(\x000F(VOF\x000F(VOS\x000K(1),\x000K(1)A\x000K(1)K\x000K(1)O\x000K(1O(\x000K(1OF\x000K(1OS\x000K(1OV\x000K(F()\x000K(F(1\x000K(F(F\x000K(F(N\x000K(F(S\x000K(F(V\x000K(N),\x000K(N)A\x000K(N)K\x000K(N)O\x000K(NO(\x000K(NOF\x000K(NOS\x000K(NOV\x000K(S),\x000K(S)A\x000K(S)K\x000K(S)O\x000K(SO(\x000K(SO1\x000K(SOF\x000K(SON\x000K(SOS\x000K(SOV\x000K(V),\x000K(V)A\x000K(V)K\x000K(V)O\x000K(VO(\x000K(VOF\x000K(VOS\x000K1,(1\x000K1,(F\x000K1,(N\x000K1,(S\x000K1,(V\x000K1,F(\x000K1A(F\x000K1A(N\x000K1A(S\x000K1A(V\x000K1AF(\x000K1ASO\x000K1AVO\x000K1K(1\x000K1K(F\x000K1K(N\x000K1K(S\x000K1K(V\x000K1K1O\x000K1K1U\x000K1KF(\x000K1KNU\x000K1KSO\x000K1KSU\x000K1KVO\x000K1KVU\x000K1O(1\x000K1O(F\x000K1O(N\x000K1O(S\x000K1O(V\x000K1OF(\x000K1OS(\x000K1OS,\x000K1OS1\x000K1OSA\x000K1OSF\x000K1OSK\x000K1OSV\x000K1OV(\x000K1OV,\x000K1OVA\x000K1OVF\x000K1OVK\x000K1OVO\x000K1OVS\x000KF(),\x000KF()A\x000KF()K\x000KF()O\x000KF(1)\x000KF(1O\x000KF(F(\x000KF(N)\x000KF(NO\x000KF(S)\x000KF(SO\x000KF(V)\x000KF(VO\x000KN,(1\x000KN,(F\x000KN,(N\x000KN,(S\x000KN,(V\x000KN,F(\x000KNA(F\x000KNA(N\x000KNA(S\x000KNA(V\x000KNAF(\x000KNASO\x000KNAVO\x000KNK(1\x000KNK(F\x000KNK(N\x000KNK(S\x000KNK(V\x000KNK1O\x000KNK1U\x000KNKF(\x000KNKNU\x000KNKSO\x000KNKSU\x000KNKVO\x000KNKVU\x000KS,(1\x000KS,(F\x000KS,(N\x000KS,(S\x000KS,(V\x000KS,F(\x000KSA(F\x000KSA(N\x000KSA(S\x000KSA(V\x000KSAF(\x000KSASO\x000KSAVO\x000KSK(1\x000KSK(F\x000KSK(N\x000KSK(S\x000KSK(V\x000KSK1O\x000KSK1U\x000KSKF(\x000KSKNU\x000KSKSO\x000KSKSU\x000KSKVO\x000KSKVU\x000KSO(1\x000KSO(F\x000KSO(N\x000KSO(S\x000KSO(V\x000KSO1(\x000KSO1,\x000KSO1A\x000KSO1F\x000KSO1K\x000KSO1N\x000KSO1S\x000KSO1V\x000KSOF(\x000KSON(\x000KSON,\x000KSON1\x000KSONA\x000KSONF\x000KSONK\x000KSOS(\x000KSOS,\x000KSOS1\x000KSOSA\x000KSOSF\x000KSOSK\x000KSOSV\x000KSOV(\x000KSOV,\x000KSOVA\x000KSOVF\x000KSOVK\x000KSOVO\x000KSOVS\x000KV,(1\x000KV,(F\x000KV,(N\x000KV,(S\x000KV,(V\x000KV,F(\x000KVA(F\x000KVA(N\x000KVA(S\x000KVA(V\x000KVAF(\x000KVASO\x000KVAVO\x000KVK(1\x000KVK(F\x000KVK(N\x000KVK(S\x000KVK(V\x000KVK1O\x000KVK1U\x000KVKF(\x000KVKNU\x000KVKSO\x000KVKSU\x000KVKVO\x000KVKVU\x000KVO(1\x000KVO(F\x000KVO(N\x000KVO(S\x000KVO(V\x000KVOF(\x000KVOS(\x000KVOS,\x000KVOS1\x000KVOSA\x000KVOSF\x000KVOSK\x000KVOSV\x000N&(1&\x000N&(1)\x000N&(1,\x000N&(1O\x000N&(E(\x000N&(E1\x000N&(EF\x000N&(EK\x000N&(EN\x000N&(EO\x000N&(ES\x000N&(EV\x000N&(F(\x000N&(N&\x000N&(N)\x000N&(N,\x000N&(NO\x000N&(S&\x000N&(S)\x000N&(S,\x000N&(SO\x000N&(V&\x000N&(V)\x000N&(V,\x000N&(VO\x000N&1\x000N&1&(\x000N&1&1\x000N&1&F\x000N&1&N\x000N&1&S\x000N&1&V\x000N&1)&\x000N&1)C\x000N&1)O\x000N&1)U\x000N&1;\x000N&1;C\x000N&1;E\x000N&1;T\x000N&1B(\x000N&1B1\x000N&1BF\x000N&1BN\x000N&1BS\x000N&1BV\x000N&1C\x000N&1EK\x000N&1EN\x000N&1F(\x000N&1K(\x000N&1K1\x000N&1KF\x000N&1KN\x000N&1KS\x000N&1KV\x000N&1O(\x000N&1OF\x000N&1OS\x000N&1OV\x000N&1TN\x000N&1U\x000N&1U(\x000N&1U;\x000N&1UC\x000N&1UE\x000N&E(1\x000N&E(F\x000N&E(N\x000N&E(O\x000N&E(S\x000N&E(V\x000N&E1\x000N&E1;\x000N&E1C\x000N&E1K\x000N&E1O\x000N&EF(\x000N&EK(\x000N&EK1\x000N&EKF\x000N&EKN\x000N&EKS\x000N&EKV\x000N&EN;\x000N&ENC\x000N&ENK\x000N&ENO\x000N&ES\x000N&ES;\x000N&ESC\x000N&ESK\x000N&ESO\x000N&EV\x000N&EV;\x000N&EVC\x000N&EVK\x000N&EVO\x000N&F()\x000N&F(1\x000N&F(E\x000N&F(F\x000N&F(N\x000N&F(S\x000N&F(V\x000N&K&(\x000N&K&1\x000N&K&F\x000N&K&N\x000N&K&S\x000N&K&V\x000N&K(1\x000N&K(F\x000N&K(N\x000N&K(S\x000N&K(V\x000N&K1O\x000N&KC\x000N&KF(\x000N&KNK\x000N&KO(\x000N&KO1\x000N&KOF\x000N&KOK\x000N&KON\x000N&KOS\x000N&KOV\x000N&KSO\x000N&KVO\x000N&N&(\x000N&N&1\x000N&N&F\x000N&N&S\x000N&N&V\x000N&N)&\x000N&N)C\x000N&N)O\x000N&N)U\x000N&N;C\x000N&N;E\x000N&N;T\x000N&NB(\x000N&NB1\x000N&NBF\x000N&NBS\x000N&NBV\x000N&NF(\x000N&NK(\x000N&NK1\x000N&NKF\x000N&NKS\x000N&NKV\x000N&NO(\x000N&NOF\x000N&NOS\x000N&NOV\x000N&NU\x000N&NU(\x000N&NU;\x000N&NUC\x000N&NUE\x000N&S&(\x000N&S&1\x000N&S&F\x000N&S&N\x000N&S&S\x000N&S&V\x000N&S)&\x000N&S)C\x000N&S)O\x000N&S)U\x000N&S1\x000N&S1;\x000N&S1C\x000N&S;\x000N&S;C\x000N&S;E\x000N&S;T\x000N&SB(\x000N&SB1\x000N&SBF\x000N&SBN\x000N&SBS\x000N&SBV\x000N&SC\x000N&SEK\x000N&SEN\x000N&SF(\x000N&SK(\x000N&SK1\x000N&SKF\x000N&SKN\x000N&SKS\x000N&SKV\x000N&SO(\x000N&SO1\x000N&SOF\x000N&SON\x000N&SOS\x000N&SOV\x000N&STN\x000N&SU\x000N&SU(\x000N&SU;\x000N&SUC\x000N&SUE\x000N&SV\x000N&SV;\x000N&SVC\x000N&SVO\x000N&V\x000N&V&(\x000N&V&1\x000N&V&F\x000N&V&N\x000N&V&S\x000N&V&V\x000N&V)&\x000N&V)C\x000N&V)O\x000N&V)U\x000N&V;\x000N&V;C\x000N&V;E\x000N&V;T\x000N&VB(\x000N&VB1\x000N&VBF\x000N&VBN\x000N&VBS\x000N&VBV\x000N&VC\x000N&VEK\x000N&VEN\x000N&VF(\x000N&VK(\x000N&VK1\x000N&VKF\x000N&VKN\x000N&VKS\x000N&VKV\x000N&VO(\x000N&VOF\x000N&VOS\x000N&VS\x000N&VS;\x000N&VSC\x000N&VSO\x000N&VTN\x000N&VU\x000N&VU(\x000N&VU;\x000N&VUC\x000N&VUE\x000N)&(1\x000N)&(E\x000N)&(F\x000N)&(N\x000N)&(S\x000N)&(V\x000N)&1\x000N)&1&\x000N)&1)\x000N)&1;\x000N)&1B\x000N)&1C\x000N)&1F\x000N)&1O\x000N)&1U\x000N)&F(\x000N)&N\x000N)&N&\x000N)&N)\x000N)&N;\x000N)&NB\x000N)&NC\x000N)&NF\x000N)&NO\x000N)&NU\x000N)&S\x000N)&S&\x000N)&S)\x000N)&S;\x000N)&SB\x000N)&SC\x000N)&SF\x000N)&SO\x000N)&SU\x000N)&V\x000N)&V&\x000N)&V)\x000N)&V;\x000N)&VB\x000N)&VC\x000N)&VF\x000N)&VO\x000N)&VU\x000N),(1\x000N),(F\x000N),(N\x000N),(S\x000N),(V\x000N);E(\x000N);E1\x000N);EF\x000N);EK\x000N);EN\x000N);EO\x000N);ES\x000N);EV\x000N);T(\x000N);T1\x000N);TF\x000N);TK\x000N);TN\x000N);TO\x000N);TS\x000N);TV\x000N)B(1\x000N)B(F\x000N)B(N\x000N)B(S\x000N)B(V\x000N)B1\x000N)B1&\x000N)B1;\x000N)B1C\x000N)B1K\x000N)B1N\x000N)B1O\x000N)B1U\x000N)BF(\x000N)BN\x000N)BN&\x000N)BN;\x000N)BNC\x000N)BNK\x000N)BNO\x000N)BNU\x000N)BS\x000N)BS&\x000N)BS;\x000N)BSC\x000N)BSK\x000N)BSO\x000N)BSU\x000N)BV\x000N)BV&\x000N)BV;\x000N)BVC\x000N)BVK\x000N)BVO\x000N)BVU\x000N)E(1\x000N)E(F\x000N)E(N\x000N)E(S\x000N)E(V\x000N)E1C\x000N)E1O\x000N)EF(\x000N)EK(\x000N)EK1\x000N)EKF\x000N)EKN\x000N)EKS\x000N)EKV\x000N)ENC\x000N)ENO\x000N)ESC\x000N)ESO\x000N)EVC\x000N)EVO\x000N)F(F\x000N)K(1\x000N)K(F\x000N)K(N\x000N)K(S\x000N)K(V\x000N)K1&\x000N)K1;\x000N)K1B\x000N)K1E\x000N)K1O\x000N)K1U\x000N)KB(\x000N)KB1\x000N)KBF\x000N)KBN\x000N)KBS\x000N)KBV\x000N)KF(\x000N)KN&\x000N)KN;\x000N)KNB\x000N)KNC\x000N)KNE\x000N)KNK\x000N)KNU\x000N)KS&\x000N)KS;\x000N)KSB\x000N)KSE\x000N)KSO\x000N)KSU\x000N)KUE\x000N)KV&\x000N)KV;\x000N)KVB\x000N)KVE\x000N)KVO\x000N)KVU\x000N)O(1\x000N)O(E\x000N)O(F\x000N)O(N\x000N)O(S\x000N)O(V\x000N)O1&\x000N)O1)\x000N)O1;\x000N)O1B\x000N)O1C\x000N)O1K\x000N)O1U\x000N)OF(\x000N)ON&\x000N)ON)\x000N)ON;\x000N)ONB\x000N)ONC\x000N)ONK\x000N)ONU\x000N)OS\x000N)OS&\x000N)OS)\x000N)OS;\x000N)OSB\x000N)OSC\x000N)OSK\x000N)OSU\x000N)OV\x000N)OV&\x000N)OV)\x000N)OV;\x000N)OVB\x000N)OVC\x000N)OVK\x000N)OVO\x000N)OVU\x000N)U(E\x000N)UE(\x000N)UE1\x000N)UEF\x000N)UEK\x000N)UEN\x000N)UES\x000N)UEV\x000N,(1)\x000N,(1O\x000N,(E(\x000N,(E1\x000N,(EF\x000N,(EK\x000N,(EN\x000N,(ES\x000N,(EV\x000N,(F(\x000N,(NO\x000N,(S)\x000N,(SO\x000N,(V)\x000N,(VO\x000N,F()\x000N,F(1\x000N,F(F\x000N,F(N\x000N,F(S\x000N,F(V\x000N1O(1\x000N1O(F\x000N1O(N\x000N1O(S\x000N1O(V\x000N1OF(\x000N1OS(\x000N1OS1\x000N1OSF\x000N1OSU\x000N1OSV\x000N1OV(\x000N1OVF\x000N1OVO\x000N1OVS\x000N1OVU\x000N1S;\x000N1S;C\x000N1SC\x000N1UE\x000N1UE;\x000N1UEC\x000N1UEK\x000N1V;\x000N1V;C\x000N1VC\x000N1VO(\x000N1VOF\x000N1VOS\x000N;E(1\x000N;E(E\x000N;E(F\x000N;E(N\x000N;E(S\x000N;E(V\x000N;E1,\x000N;E1;\x000N;E1C\x000N;E1K\x000N;E1O\x000N;E1T\x000N;EF(\x000N;EK(\x000N;EK1\x000N;EKF\x000N;EKN\x000N;EKO\x000N;EKS\x000N;EKV\x000N;EN,\x000N;EN;\x000N;ENC\x000N;ENE\x000N;ENK\x000N;ENO\x000N;ENT\x000N;ES,\x000N;ES;\x000N;ESC\x000N;ESK\x000N;ESO\x000N;EST\x000N;EV,\x000N;EV;\x000N;EVC\x000N;EVK\x000N;EVO\x000N;EVT\x000N;N:T\x000N;T(1\x000N;T(C\x000N;T(E\x000N;T(F\x000N;T(N\x000N;T(S\x000N;T(V\x000N;T1(\x000N;T1,\x000N;T1;\x000N;T1C\x000N;T1F\x000N;T1K\x000N;T1O\x000N;T1T\x000N;T;\x000N;T;C\x000N;TF(\x000N;TK(\x000N;TK1\x000N;TKF\x000N;TKK\x000N;TKO\x000N;TKS\x000N;TKV\x000N;TN(\x000N;TN,\x000N;TN1\x000N;TN;\x000N;TNC\x000N;TNE\x000N;TNF\x000N;TNK\x000N;TNN\x000N;TNO\x000N;TNS\x000N;TNT\x000N;TNV\x000N;TO(\x000N;TS(\x000N;TS,\x000N;TS;\x000N;TSC\x000N;TSF\x000N;TSK\x000N;TSO\x000N;TST\x000N;TTN\x000N;TV(\x000N;TV,\x000N;TV;\x000N;TVC\x000N;TVF\x000N;TVK\x000N;TVO\x000N;TVT\x000NA(F(\x000NA(N)\x000NA(NO\x000NA(S)\x000NA(SO\x000NA(V)\x000NA(VO\x000NAF()\x000NAF(1\x000NAF(F\x000NAF(N\x000NAF(S\x000NAF(V\x000NASO(\x000NASO1\x000NASOF\x000NASON\x000NASOS\x000NASOV\x000NASUE\x000NATO(\x000NATO1\x000NATOF\x000NATON\x000NATOS\x000NATOV\x000NATUE\x000NAVO(\x000NAVOF\x000NAVOS\x000NAVUE\x000NB(1&\x000NB(1)\x000NB(1O\x000NB(F(\x000NB(N&\x000NB(NO\x000NB(S&\x000NB(S)\x000NB(SO\x000NB(V&\x000NB(V)\x000NB(VO\x000NB1\x000NB1&(\x000NB1&1\x000NB1&F\x000NB1&N\x000NB1&S\x000NB1&V\x000NB1,(\x000NB1,F\x000NB1;\x000NB1;C\x000NB1B(\x000NB1B1\x000NB1BF\x000NB1BN\x000NB1BS\x000NB1BV\x000NB1C\x000NB1K(\x000NB1K1\x000NB1KF\x000NB1KN\x000NB1KS\x000NB1KV\x000NB1O(\x000NB1OF\x000NB1OS\x000NB1OV\x000NB1U(\x000NB1UE\x000NBE(1\x000NBE(F\x000NBE(N\x000NBE(S\x000NBE(V\x000NBEK(\x000NBF()\x000NBF(1\x000NBF(F\x000NBF(N\x000NBF(S\x000NBF(V\x000NBN&(\x000NBN&1\x000NBN&F\x000NBN&N\x000NBN&S\x000NBN&V\x000NBN,(\x000NBN,F\x000NBN;\x000NBN;C\x000NBNB(\x000NBNB1\x000NBNBF\x000NBNBN\x000NBNBS\x000NBNBV\x000NBNC\x000NBNK(\x000NBNK1\x000NBNKF\x000NBNKN\x000NBNKS\x000NBNKV\x000NBNO(\x000NBNOF\x000NBNOS\x000NBNOV\x000NBNU(\x000NBNUE\x000NBS\x000NBS&(\x000NBS&1\x000NBS&F\x000NBS&N\x000NBS&S\x000NBS&V\x000NBS,(\x000NBS,F\x000NBS;\x000NBS;C\x000NBSB(\x000NBSB1\x000NBSBF\x000NBSBN\x000NBSBS\x000NBSBV\x000NBSC\x000NBSK(\x000NBSK1\x000NBSKF\x000NBSKN\x000NBSKS\x000NBSKV\x000NBSO(\x000NBSO1\x000NBSOF\x000NBSON\x000NBSOS\x000NBSOV\x000NBSU(\x000NBSUE\x000NBV\x000NBV&(\x000NBV&1\x000NBV&F\x000NBV&N\x000NBV&S\x000NBV&V\x000NBV,(\x000NBV,F\x000NBV;\x000NBV;C\x000NBVB(\x000NBVB1\x000NBVBF\x000NBVBN\x000NBVBS\x000NBVBV\x000NBVC\x000NBVK(\x000NBVK1\x000NBVKF\x000NBVKN\x000NBVKS\x000NBVKV\x000NBVO(\x000NBVOF\x000NBVOS\x000NBVU(\x000NBVUE\x000NC\x000NE(1)\x000NE(1O\x000NE(F(\x000NE(N)\x000NE(NO\x000NE(S)\x000NE(SO\x000NE(V)\x000NE(VO\x000NE1;T\x000NE1C\x000NE1O(\x000NE1OF\x000NE1OS\x000NE1OV\x000NE1T(\x000NE1T1\x000NE1TF\x000NE1TN\x000NE1TS\x000NE1TV\x000NE1UE\x000NEF()\x000NEF(1\x000NEF(F\x000NEF(N\x000NEF(S\x000NEF(V\x000NEN;T\x000NENO(\x000NENOF\x000NENOS\x000NENOV\x000NENT(\x000NENT1\x000NENTF\x000NENTN\x000NENTS\x000NENTV\x000NENUE\x000NEOKN\x000NES;T\x000NESC\x000NESO(\x000NESO1\x000NESOF\x000NESON\x000NESOS\x000NESOV\x000NEST(\x000NEST1\x000NESTF\x000NESTN\x000NESTS\x000NESTV\x000NESUE\x000NEU(1\x000NEU(F\x000NEU(N\x000NEU(S\x000NEU(V\x000NEU1,\x000NEU1C\x000NEU1O\x000NEUEF\x000NEUEK\x000NEUF(\x000NEUS,\x000NEUSC\x000NEUSO\x000NEUV,\x000NEUVC\x000NEUVO\x000NEV;T\x000NEVC\x000NEVO(\x000NEVOF\x000NEVOS\x000NEVT(\x000NEVT1\x000NEVTF\x000NEVTN\x000NEVTS\x000NEVTV\x000NEVUE\x000NF()1\x000NF()F\x000NF()K\x000NF()N\x000NF()O\x000NF()S\x000NF()U\x000NF()V\x000NF(1)\x000NF(1O\x000NF(E(\x000NF(E1\x000NF(EF\x000NF(EK\x000NF(EN\x000NF(ES\x000NF(EV\x000NF(F(\x000NF(N,\x000NF(NO\x000NF(S)\x000NF(SO\x000NF(V)\x000NF(VO\x000NK(1)\x000NK(1O\x000NK(F(\x000NK(NO\x000NK(S)\x000NK(SO\x000NK(V)\x000NK(VO\x000NK)&(\x000NK)&1\x000NK)&F\x000NK)&N\x000NK)&S\x000NK)&V\x000NK);E\x000NK);T\x000NK)B(\x000NK)B1\x000NK)BF\x000NK)BN\x000NK)BS\x000NK)BV\x000NK)E(\x000NK)E1\x000NK)EF\x000NK)EK\x000NK)EN\x000NK)ES\x000NK)EV\x000NK)F(\x000NK)O(\x000NK)OF\x000NK)UE\x000NK1\x000NK1&(\x000NK1&1\x000NK1&F\x000NK1&N\x000NK1&S\x000NK1&V\x000NK1;C\x000NK1;E\x000NK1;T\x000NK1B(\x000NK1B1\x000NK1BF\x000NK1BN\x000NK1BS\x000NK1BV\x000NK1C\x000NK1E(\x000NK1E1\x000NK1EF\x000NK1EK\x000NK1EN\x000NK1ES\x000NK1EV\x000NK1O(\x000NK1OF\x000NK1OS\x000NK1OV\x000NK1U(\x000NK1UE\x000NKF()\x000NKF(1\x000NKF(F\x000NKF(N\x000NKF(S\x000NKF(V\x000NKN\x000NKN&(\x000NKN&1\x000NKN&F\x000NKN&S\x000NKN&V\x000NKN;C\x000NKN;E\x000NKN;T\x000NKNB(\x000NKNB1\x000NKNBF\x000NKNBN\x000NKNBS\x000NKNBV\x000NKNE(\x000NKNE1\x000NKNEF\x000NKNES\x000NKNEV\x000NKNU(\x000NKNUE\x000NKS\x000NKS&(\x000NKS&1\x000NKS&F\x000NKS&N\x000NKS&S\x000NKS&V\x000NKS;\x000NKS;C\x000NKS;E\x000NKS;T\x000NKSB(\x000NKSB1\x000NKSBF\x000NKSBN\x000NKSBS\x000NKSBV\x000NKSC\x000NKSE(\x000NKSE1\x000NKSEF\x000NKSEK\x000NKSEN\x000NKSES\x000NKSEV\x000NKSO(\x000NKSO1\x000NKSOF\x000NKSON\x000NKSOS\x000NKSOV\x000NKSU(\x000NKSUE\x000NKUE(\x000NKUE1\x000NKUEF\x000NKUEK\x000NKUEN\x000NKUES\x000NKUEV\x000NKV\x000NKV&(\x000NKV&1\x000NKV&F\x000NKV&N\x000NKV&S\x000NKV&V\x000NKV;\x000NKV;C\x000NKV;E\x000NKV;T\x000NKVB(\x000NKVB1\x000NKVBF\x000NKVBN\x000NKVBS\x000NKVBV\x000NKVC\x000NKVE(\x000NKVE1\x000NKVEF\x000NKVEK\x000NKVEN\x000NKVES\x000NKVEV\x000NKVO(\x000NKVOF\x000NKVOS\x000NKVU(\x000NKVUE\x000NO(1&\x000NO(1)\x000NO(1,\x000NO(1O\x000NO(E(\x000NO(E1\x000NO(EE\x000NO(EF\x000NO(EK\x000NO(EN\x000NO(EO\x000NO(ES\x000NO(EV\x000NO(F(\x000NO(N&\x000NO(N)\x000NO(N,\x000NO(NO\x000NO(S&\x000NO(S)\x000NO(S,\x000NO(SO\x000NO(V&\x000NO(V)\x000NO(V,\x000NO(VO\x000NOF()\x000NOF(1\x000NOF(E\x000NOF(F\x000NOF(N\x000NOF(S\x000NOF(V\x000NOK&(\x000NOK(1\x000NOK(F\x000NOK(N\x000NOK(S\x000NOK(V\x000NOK1C\x000NOK1O\x000NOKF(\x000NOKNC\x000NOKO(\x000NOKO1\x000NOKOF\x000NOKON\x000NOKOS\x000NOKOV\x000NOKSC\x000NOKSO\x000NOKVC\x000NOKVO\x000NONSU\x000NOS&(\x000NOS&1\x000NOS&E\x000NOS&F\x000NOS&K\x000NOS&N\x000NOS&S\x000NOS&U\x000NOS&V\x000NOS(E\x000NOS(U\x000NOS)&\x000NOS),\x000NOS);\x000NOS)B\x000NOS)C\x000NOS)E\x000NOS)F\x000NOS)K\x000NOS)O\x000NOS)U\x000NOS,(\x000NOS,F\x000NOS1(\x000NOS1F\x000NOS1N\x000NOS1S\x000NOS1U\x000NOS1V\x000NOS;\x000NOS;C\x000NOS;E\x000NOS;T\x000NOSA(\x000NOSAF\x000NOSAS\x000NOSAT\x000NOSAV\x000NOSB(\x000NOSB1\x000NOSBE\x000NOSBF\x000NOSBN\x000NOSBS\x000NOSBV\x000NOSC\x000NOSE(\x000NOSE1\x000NOSEF\x000NOSEK\x000NOSEN\x000NOSEO\x000NOSES\x000NOSEU\x000NOSEV\x000NOSF(\x000NOSK(\x000NOSK)\x000NOSK1\x000NOSKB\x000NOSKF\x000NOSKN\x000NOSKS\x000NOSKU\x000NOSKV\x000NOST(\x000NOST1\x000NOSTE\x000NOSTF\x000NOSTN\x000NOSTS\x000NOSTT\x000NOSTV\x000NOSU\x000NOSU(\x000NOSU1\x000NOSU;\x000NOSUC\x000NOSUE\x000NOSUF\x000NOSUK\x000NOSUO\x000NOSUS\x000NOSUT\x000NOSUV\x000NOSV(\x000NOSVF\x000NOSVO\x000NOSVS\x000NOSVU\x000NOU(E\x000NOUEK\x000NOUEN\x000NOV&(\x000NOV&1\x000NOV&E\x000NOV&F\x000NOV&K\x000NOV&N\x000NOV&S\x000NOV&U\x000NOV&V\x000NOV(E\x000NOV(U\x000NOV)&\x000NOV),\x000NOV);\x000NOV)B\x000NOV)C\x000NOV)E\x000NOV)F\x000NOV)K\x000NOV)O\x000NOV)U\x000NOV,(\x000NOV,F\x000NOV;\x000NOV;C\x000NOV;E\x000NOV;N\x000NOV;T\x000NOVA(\x000NOVAF\x000NOVAS\x000NOVAT\x000NOVAV\x000NOVB(\x000NOVB1\x000NOVBE\x000NOVBF\x000NOVBN\x000NOVBS\x000NOVBV\x000NOVC\x000NOVE(\x000NOVE1\x000NOVEF\x000NOVEK\x000NOVEN\x000NOVEO\x000NOVES\x000NOVEU\x000NOVEV\x000NOVF(\x000NOVK(\x000NOVK)\x000NOVK1\x000NOVKB\x000NOVKF\x000NOVKN\x000NOVKS\x000NOVKU\x000NOVKV\x000NOVO(\x000NOVOF\x000NOVOK\x000NOVOS\x000NOVOU\x000NOVS(\x000NOVS1\x000NOVSF\x000NOVSO\x000NOVSU\x000NOVSV\x000NOVT(\x000NOVT1\x000NOVTE\x000NOVTF\x000NOVTN\x000NOVTS\x000NOVTT\x000NOVTV\x000NOVU\x000NOVU(\x000NOVU1\x000NOVU;\x000NOVUC\x000NOVUE\x000NOVUF\x000NOVUK\x000NOVUO\x000NOVUS\x000NOVUT\x000NOVUV\x000NSO1U\x000NSONU\x000NSOSU\x000NSOVU\x000NSUE\x000NSUE;\x000NSUEC\x000NSUEK\x000NT(1)\x000NT(1O\x000NT(F(\x000NT(N)\x000NT(NO\x000NT(S)\x000NT(SO\x000NT(V)\x000NT(VO\x000NT1(F\x000NT1O(\x000NT1OF\x000NT1OS\x000NT1OV\x000NTE(1\x000NTE(F\x000NTE(N\x000NTE(S\x000NTE(V\x000NTE1N\x000NTE1O\x000NTEF(\x000NTEK(\x000NTEK1\x000NTEKF\x000NTEKN\x000NTEKS\x000NTEKV\x000NTENN\x000NTENO\x000NTESN\x000NTESO\x000NTEVN\x000NTEVO\x000NTF()\x000NTF(1\x000NTF(F\x000NTF(N\x000NTF(S\x000NTF(V\x000NTN(1\x000NTN(F\x000NTN(S\x000NTN(V\x000NTN1C\x000NTN1O\x000NTN;E\x000NTN;N\x000NTN;T\x000NTNE(\x000NTNE1\x000NTNEF\x000NTNEN\x000NTNES\x000NTNEV\x000NTNF(\x000NTNKN\x000NTNN:\x000NTNNC\x000NTNNO\x000NTNO(\x000NTNOF\x000NTNOS\x000NTNOV\x000NTNSC\x000NTNSO\x000NTNT(\x000NTNT1\x000NTNTF\x000NTNTN\x000NTNTS\x000NTNTV\x000NTNVC\x000NTNVO\x000NTS(F\x000NTSO(\x000NTSO1\x000NTSOF\x000NTSON\x000NTSOS\x000NTSOV\x000NTTNE\x000NTTNK\x000NTTNN\x000NTTNT\x000NTV(1\x000NTV(F\x000NTVO(\x000NTVOF\x000NTVOS\x000NU(1)\x000NU(1O\x000NU(E(\x000NU(E1\x000NU(EF\x000NU(EK\x000NU(EN\x000NU(ES\x000NU(EV\x000NU(F(\x000NU(N)\x000NU(NO\x000NU(S)\x000NU(SO\x000NU(V)\x000NU(VO\x000NU1,(\x000NU1,F\x000NU1C\x000NU1O(\x000NU1OF\x000NU1OS\x000NU1OV\x000NU;\x000NU;C\x000NUC\x000NUE\x000NUE(1\x000NUE(E\x000NUE(F\x000NUE(N\x000NUE(O\x000NUE(S\x000NUE(V\x000NUE1\x000NUE1&\x000NUE1(\x000NUE1)\x000NUE1,\x000NUE1;\x000NUE1B\x000NUE1C\x000NUE1F\x000NUE1K\x000NUE1N\x000NUE1O\x000NUE1S\x000NUE1U\x000NUE1V\x000NUE;\x000NUE;C\x000NUEC\x000NUEF\x000NUEF(\x000NUEF,\x000NUEF;\x000NUEFC\x000NUEK\x000NUEK(\x000NUEK1\x000NUEK;\x000NUEKC\x000NUEKF\x000NUEKN\x000NUEKO\x000NUEKS\x000NUEKV\x000NUEN\x000NUEN&\x000NUEN(\x000NUEN)\x000NUEN,\x000NUEN1\x000NUEN;\x000NUENB\x000NUENC\x000NUENF\x000NUENK\x000NUENO\x000NUENS\x000NUENU\x000NUEOK\x000NUEON\x000NUES\x000NUES&\x000NUES(\x000NUES)\x000NUES,\x000NUES1\x000NUES;\x000NUESB\x000NUESC\x000NUESF\x000NUESK\x000NUESO\x000NUESU\x000NUESV\x000NUEV\x000NUEV&\x000NUEV(\x000NUEV)\x000NUEV,\x000NUEV;\x000NUEVB\x000NUEVC\x000NUEVF\x000NUEVK\x000NUEVN\x000NUEVO\x000NUEVS\x000NUEVU\x000NUF()\x000NUF(1\x000NUF(F\x000NUF(N\x000NUF(S\x000NUF(V\x000NUK(E\x000NUO(E\x000NUON(\x000NUON1\x000NUONF\x000NUONS\x000NUS,(\x000NUS,F\x000NUSC\x000NUSO(\x000NUSO1\x000NUSOF\x000NUSON\x000NUSOS\x000NUSOV\x000NUTN(\x000NUTN1\x000NUTNF\x000NUTNN\x000NUTNS\x000NUTNV\x000NUV,(\x000NUV,F\x000NUVC\x000NUVO(\x000NUVOF\x000NUVOS\x000S&(1&\x000S&(1)\x000S&(1,\x000S&(1O\x000S&(E(\x000S&(E1\x000S&(EF\x000S&(EK\x000S&(EN\x000S&(EO\x000S&(ES\x000S&(EV\x000S&(F(\x000S&(N&\x000S&(N)\x000S&(N,\x000S&(NO\x000S&(S&\x000S&(S)\x000S&(S,\x000S&(SO\x000S&(V&\x000S&(V)\x000S&(V,\x000S&(VO\x000S&1\x000S&1&(\x000S&1&1\x000S&1&F\x000S&1&N\x000S&1&S\x000S&1&V\x000S&1)&\x000S&1)C\x000S&1)O\x000S&1)U\x000S&1;\x000S&1;C\x000S&1;E\x000S&1;T\x000S&1B(\x000S&1B1\x000S&1BF\x000S&1BN\x000S&1BS\x000S&1BV\x000S&1C\x000S&1EK\x000S&1EN\x000S&1F(\x000S&1K(\x000S&1K1\x000S&1KF\x000S&1KN\x000S&1KS\x000S&1KV\x000S&1O(\x000S&1OF\x000S&1OS\x000S&1OV\x000S&1TN\x000S&1U\x000S&1U(\x000S&1U;\x000S&1UC\x000S&1UE\x000S&E(1\x000S&E(F\x000S&E(N\x000S&E(O\x000S&E(S\x000S&E(V\x000S&E1\x000S&E1;\x000S&E1C\x000S&E1K\x000S&E1O\x000S&EF(\x000S&EK(\x000S&EK1\x000S&EKF\x000S&EKN\x000S&EKS\x000S&EKV\x000S&EN\x000S&EN;\x000S&ENC\x000S&ENK\x000S&ENO\x000S&ES\x000S&ES;\x000S&ESC\x000S&ESK\x000S&ESO\x000S&EV\x000S&EV;\x000S&EVC\x000S&EVK\x000S&EVO\x000S&F()\x000S&F(1\x000S&F(E\x000S&F(F\x000S&F(N\x000S&F(S\x000S&F(V\x000S&K&(\x000S&K&1\x000S&K&F\x000S&K&N\x000S&K&S\x000S&K&V\x000S&K(1\x000S&K(F\x000S&K(N\x000S&K(S\x000S&K(V\x000S&K1O\x000S&KC\x000S&KF(\x000S&KNK\x000S&KO(\x000S&KO1\x000S&KOF\x000S&KOK\x000S&KON\x000S&KOS\x000S&KOV\x000S&KSO\x000S&KVO\x000S&N\x000S&N&(\x000S&N&1\x000S&N&F\x000S&N&N\x000S&N&S\x000S&N&V\x000S&N)&\x000S&N)C\x000S&N)O\x000S&N)U\x000S&N;\x000S&N;C\x000S&N;E\x000S&N;T\x000S&NB(\x000S&NB1\x000S&NBF\x000S&NBN\x000S&NBS\x000S&NBV\x000S&NC\x000S&NEN\x000S&NF(\x000S&NK(\x000S&NK1\x000S&NKF\x000S&NKN\x000S&NKS\x000S&NKV\x000S&NO(\x000S&NOF\x000S&NOS\x000S&NOV\x000S&NTN\x000S&NU\x000S&NU(\x000S&NU;\x000S&NUC\x000S&NUE\x000S&S\x000S&S&(\x000S&S&1\x000S&S&F\x000S&S&N\x000S&S&S\x000S&S&V\x000S&S)&\x000S&S)C\x000S&S)O\x000S&S)U\x000S&S1\x000S&S1;\x000S&S1C\x000S&S;\x000S&S;C\x000S&S;E\x000S&S;T\x000S&SB(\x000S&SB1\x000S&SBF\x000S&SBN\x000S&SBS\x000S&SBV\x000S&SC\x000S&SEK\x000S&SEN\x000S&SF(\x000S&SK(\x000S&SK1\x000S&SKF\x000S&SKN\x000S&SKS\x000S&SKV\x000S&SO(\x000S&SO1\x000S&SOF\x000S&SON\x000S&SOS\x000S&SOV\x000S&STN\x000S&SU\x000S&SU(\x000S&SU;\x000S&SUC\x000S&SUE\x000S&SV\x000S&SV;\x000S&SVC\x000S&SVO\x000S&V\x000S&V&(\x000S&V&1\x000S&V&F\x000S&V&N\x000S&V&S\x000S&V&V\x000S&V)&\x000S&V)C\x000S&V)O\x000S&V)U\x000S&V;\x000S&V;C\x000S&V;E\x000S&V;T\x000S&VB(\x000S&VB1\x000S&VBF\x000S&VBN\x000S&VBS\x000S&VBV\x000S&VC\x000S&VEK\x000S&VEN\x000S&VF(\x000S&VK(\x000S&VK1\x000S&VKF\x000S&VKN\x000S&VKS\x000S&VKV\x000S&VO(\x000S&VOF\x000S&VOS\x000S&VS\x000S&VS;\x000S&VSC\x000S&VSO\x000S&VTN\x000S&VU\x000S&VU(\x000S&VU;\x000S&VUC\x000S&VUE\x000S(EF(\x000S(EKF\x000S(EKN\x000S(ENK\x000S(U(E\x000S)&(1\x000S)&(E\x000S)&(F\x000S)&(N\x000S)&(S\x000S)&(V\x000S)&1\x000S)&1&\x000S)&1)\x000S)&1;\x000S)&1B\x000S)&1C\x000S)&1F\x000S)&1O\x000S)&1U\x000S)&F(\x000S)&N\x000S)&N&\x000S)&N)\x000S)&N;\x000S)&NB\x000S)&NC\x000S)&NF\x000S)&NO\x000S)&NU\x000S)&S\x000S)&S&\x000S)&S)\x000S)&S;\x000S)&SB\x000S)&SC\x000S)&SF\x000S)&SO\x000S)&SU\x000S)&V\x000S)&V&\x000S)&V)\x000S)&V;\x000S)&VB\x000S)&VC\x000S)&VF\x000S)&VO\x000S)&VU\x000S),(1\x000S),(F\x000S),(N\x000S),(S\x000S),(V\x000S);E(\x000S);E1\x000S);EF\x000S);EK\x000S);EN\x000S);EO\x000S);ES\x000S);EV\x000S);T(\x000S);T1\x000S);TF\x000S);TK\x000S);TN\x000S);TO\x000S);TS\x000S);TV\x000S)B(1\x000S)B(F\x000S)B(N\x000S)B(S\x000S)B(V\x000S)B1\x000S)B1&\x000S)B1;\x000S)B1C\x000S)B1K\x000S)B1N\x000S)B1O\x000S)B1U\x000S)BF(\x000S)BN\x000S)BN&\x000S)BN;\x000S)BNC\x000S)BNK\x000S)BNO\x000S)BNU\x000S)BS\x000S)BS&\x000S)BS;\x000S)BSC\x000S)BSK\x000S)BSO\x000S)BSU\x000S)BV\x000S)BV&\x000S)BV;\x000S)BVC\x000S)BVK\x000S)BVO\x000S)BVU\x000S)C\x000S)E(1\x000S)E(F\x000S)E(N\x000S)E(S\x000S)E(V\x000S)E1C\x000S)E1O\x000S)EF(\x000S)EK(\x000S)EK1\x000S)EKF\x000S)EKN\x000S)EKS\x000S)EKV\x000S)ENC\x000S)ENO\x000S)ESC\x000S)ESO\x000S)EVC\x000S)EVO\x000S)F(F\x000S)K(1\x000S)K(F\x000S)K(N\x000S)K(S\x000S)K(V\x000S)K1&\x000S)K1;\x000S)K1B\x000S)K1E\x000S)K1O\x000S)K1U\x000S)KB(\x000S)KB1\x000S)KBF\x000S)KBN\x000S)KBS\x000S)KBV\x000S)KF(\x000S)KN&\x000S)KN;\x000S)KNB\x000S)KNC\x000S)KNE\x000S)KNK\x000S)KNU\x000S)KS&\x000S)KS;\x000S)KSB\x000S)KSE\x000S)KSO\x000S)KSU\x000S)KUE\x000S)KV&\x000S)KV;\x000S)KVB\x000S)KVE\x000S)KVO\x000S)KVU\x000S)O(1\x000S)O(E\x000S)O(F\x000S)O(N\x000S)O(S\x000S)O(V\x000S)O1\x000S)O1&\x000S)O1)\x000S)O1;\x000S)O1B\x000S)O1C\x000S)O1K\x000S)O1U\x000S)OF(\x000S)ON&\x000S)ON)\x000S)ON;\x000S)ONB\x000S)ONC\x000S)ONK\x000S)ONU\x000S)OS\x000S)OS&\x000S)OS)\x000S)OS;\x000S)OSB\x000S)OSC\x000S)OSK\x000S)OSU\x000S)OV\x000S)OV&\x000S)OV)\x000S)OV;\x000S)OVB\x000S)OVC\x000S)OVK\x000S)OVO\x000S)OVU\x000S)U(E\x000S)UE(\x000S)UE1\x000S)UEF\x000S)UEK\x000S)UEN\x000S)UES\x000S)UEV\x000S,(1)\x000S,(1O\x000S,(E(\x000S,(E1\x000S,(EF\x000S,(EK\x000S,(EN\x000S,(ES\x000S,(EV\x000S,(F(\x000S,(N)\x000S,(NO\x000S,(S)\x000S,(SO\x000S,(V)\x000S,(VO\x000S,F()\x000S,F(1\x000S,F(F\x000S,F(N\x000S,F(S\x000S,F(V\x000S1F()\x000S1F(1\x000S1F(F\x000S1F(N\x000S1F(S\x000S1F(V\x000S1NC\x000S1S;\x000S1S;C\x000S1SC\x000S1UE\x000S1UE;\x000S1UEC\x000S1UEK\x000S1V\x000S1V;\x000S1V;C\x000S1VC\x000S1VO(\x000S1VOF\x000S1VOS\x000S;E(1\x000S;E(E\x000S;E(F\x000S;E(N\x000S;E(S\x000S;E(V\x000S;E1,\x000S;E1;\x000S;E1C\x000S;E1K\x000S;E1O\x000S;E1T\x000S;EF(\x000S;EK(\x000S;EK1\x000S;EKF\x000S;EKN\x000S;EKO\x000S;EKS\x000S;EKV\x000S;EN,\x000S;EN;\x000S;ENC\x000S;ENE\x000S;ENK\x000S;ENO\x000S;ENT\x000S;ES,\x000S;ES;\x000S;ESC\x000S;ESK\x000S;ESO\x000S;EST\x000S;EV,\x000S;EV;\x000S;EVC\x000S;EVK\x000S;EVO\x000S;EVT\x000S;N:T\x000S;T(1\x000S;T(C\x000S;T(E\x000S;T(F\x000S;T(N\x000S;T(S\x000S;T(V\x000S;T1(\x000S;T1,\x000S;T1;\x000S;T1C\x000S;T1F\x000S;T1K\x000S;T1O\x000S;T1T\x000S;T;\x000S;T;C\x000S;TF(\x000S;TK(\x000S;TK1\x000S;TKF\x000S;TKK\x000S;TKN\x000S;TKO\x000S;TKS\x000S;TKV\x000S;TN(\x000S;TN,\x000S;TN1\x000S;TN;\x000S;TNC\x000S;TNE\x000S;TNF\x000S;TNK\x000S;TNN\x000S;TNO\x000S;TNS\x000S;TNT\x000S;TNV\x000S;TO(\x000S;TS(\x000S;TS,\x000S;TS;\x000S;TSC\x000S;TSF\x000S;TSK\x000S;TSO\x000S;TST\x000S;TTN\x000S;TV(\x000S;TV,\x000S;TV;\x000S;TVC\x000S;TVF\x000S;TVK\x000S;TVO\x000S;TVT\x000SA(F(\x000SA(N)\x000SA(NO\x000SA(S)\x000SA(SO\x000SA(V)\x000SA(VO\x000SAF()\x000SAF(1\x000SAF(F\x000SAF(N\x000SAF(S\x000SAF(V\x000SASO(\x000SASO1\x000SASOF\x000SASON\x000SASOS\x000SASOV\x000SASUE\x000SATO(\x000SATO1\x000SATOF\x000SATON\x000SATOS\x000SATOV\x000SATUE\x000SAVO(\x000SAVOF\x000SAVOS\x000SAVUE\x000SB(1)\x000SB(1O\x000SB(F(\x000SB(NO\x000SB(S)\x000SB(SO\x000SB(V)\x000SB(VO\x000SB1\x000SB1&(\x000SB1&1\x000SB1&F\x000SB1&N\x000SB1&S\x000SB1&V\x000SB1,(\x000SB1,F\x000SB1;\x000SB1;C\x000SB1B(\x000SB1B1\x000SB1BF\x000SB1BN\x000SB1BS\x000SB1BV\x000SB1C\x000SB1K(\x000SB1K1\x000SB1KF\x000SB1KN\x000SB1KS\x000SB1KV\x000SB1O(\x000SB1OF\x000SB1OS\x000SB1OV\x000SB1U(\x000SB1UE\x000SBE(1\x000SBE(F\x000SBE(N\x000SBE(S\x000SBE(V\x000SBEK(\x000SBF()\x000SBF(1\x000SBF(F\x000SBF(N\x000SBF(S\x000SBF(V\x000SBN\x000SBN&(\x000SBN&1\x000SBN&F\x000SBN&N\x000SBN&S\x000SBN&V\x000SBN,(\x000SBN,F\x000SBN;\x000SBN;C\x000SBNB(\x000SBNB1\x000SBNBF\x000SBNBN\x000SBNBS\x000SBNBV\x000SBNC\x000SBNK(\x000SBNK1\x000SBNKF\x000SBNKN\x000SBNKS\x000SBNKV\x000SBNO(\x000SBNOF\x000SBNOS\x000SBNOV\x000SBNU(\x000SBNUE\x000SBS\x000SBS&(\x000SBS&1\x000SBS&F\x000SBS&N\x000SBS&S\x000SBS&V\x000SBS,(\x000SBS,F\x000SBS;\x000SBS;C\x000SBSB(\x000SBSB1\x000SBSBF\x000SBSBN\x000SBSBS\x000SBSBV\x000SBSC\x000SBSK(\x000SBSK1\x000SBSKF\x000SBSKN\x000SBSKS\x000SBSKV\x000SBSO(\x000SBSO1\x000SBSOF\x000SBSON\x000SBSOS\x000SBSOV\x000SBSU(\x000SBSUE\x000SBV\x000SBV&(\x000SBV&1\x000SBV&F\x000SBV&N\x000SBV&S\x000SBV&V\x000SBV,(\x000SBV,F\x000SBV;\x000SBV;C\x000SBVB(\x000SBVB1\x000SBVBF\x000SBVBN\x000SBVBS\x000SBVBV\x000SBVC\x000SBVK(\x000SBVK1\x000SBVKF\x000SBVKN\x000SBVKS\x000SBVKV\x000SBVO(\x000SBVOF\x000SBVOS\x000SBVU(\x000SBVUE\x000SC\x000SE(1)\x000SE(1O\x000SE(F(\x000SE(N)\x000SE(NO\x000SE(S)\x000SE(SO\x000SE(V)\x000SE(VO\x000SE1;T\x000SE1C\x000SE1O(\x000SE1OF\x000SE1OS\x000SE1OV\x000SE1T(\x000SE1T1\x000SE1TF\x000SE1TN\x000SE1TS\x000SE1TV\x000SE1UE\x000SEF()\x000SEF(1\x000SEF(F\x000SEF(N\x000SEF(S\x000SEF(V\x000SEK(1\x000SEK(E\x000SEK(F\x000SEK(N\x000SEK(S\x000SEK(V\x000SEK1;\x000SEK1C\x000SEK1O\x000SEK1T\x000SEK1U\x000SEKF(\x000SEKN;\x000SEKNC\x000SEKNE\x000SEKNT\x000SEKNU\x000SEKOK\x000SEKS;\x000SEKSC\x000SEKSO\x000SEKST\x000SEKSU\x000SEKU(\x000SEKU1\x000SEKUE\x000SEKUF\x000SEKUS\x000SEKUV\x000SEKV;\x000SEKVC\x000SEKVO\x000SEKVT\x000SEKVU\x000SEN;T\x000SENC\x000SENEN\x000SENO(\x000SENOF\x000SENOS\x000SENOV\x000SENT(\x000SENT1\x000SENTF\x000SENTN\x000SENTS\x000SENTV\x000SENUE\x000SEOKN\x000SES;T\x000SESC\x000SESO(\x000SESO1\x000SESOF\x000SESON\x000SESOS\x000SESOV\x000SEST(\x000SEST1\x000SESTF\x000SESTN\x000SESTS\x000SESTV\x000SESUE\x000SEU(1\x000SEU(F\x000SEU(N\x000SEU(S\x000SEU(V\x000SEU1,\x000SEU1C\x000SEU1O\x000SEUEF\x000SEUEK\x000SEUF(\x000SEUS,\x000SEUSC\x000SEUSO\x000SEUV,\x000SEUVC\x000SEUVO\x000SEV;T\x000SEVC\x000SEVO(\x000SEVOF\x000SEVOS\x000SEVT(\x000SEVT1\x000SEVTF\x000SEVTN\x000SEVTS\x000SEVTV\x000SEVUE\x000SF()1\x000SF()F\x000SF()K\x000SF()N\x000SF()O\x000SF()S\x000SF()U\x000SF()V\x000SF(1)\x000SF(1N\x000SF(1O\x000SF(E(\x000SF(E1\x000SF(EF\x000SF(EK\x000SF(EN\x000SF(ES\x000SF(EV\x000SF(F(\x000SF(N)\x000SF(N,\x000SF(NO\x000SF(S)\x000SF(SO\x000SF(V)\x000SF(VO\x000SK(1)\x000SK(1O\x000SK(F(\x000SK(N)\x000SK(NO\x000SK(S)\x000SK(SO\x000SK(V)\x000SK(VO\x000SK)&(\x000SK)&1\x000SK)&F\x000SK)&N\x000SK)&S\x000SK)&V\x000SK);E\x000SK);T\x000SK)B(\x000SK)B1\x000SK)BF\x000SK)BN\x000SK)BS\x000SK)BV\x000SK)E(\x000SK)E1\x000SK)EF\x000SK)EK\x000SK)EN\x000SK)ES\x000SK)EV\x000SK)F(\x000SK)O(\x000SK)OF\x000SK)UE\x000SK1\x000SK1&(\x000SK1&1\x000SK1&F\x000SK1&N\x000SK1&S\x000SK1&V\x000SK1;\x000SK1;C\x000SK1;E\x000SK1;T\x000SK1B(\x000SK1B1\x000SK1BF\x000SK1BN\x000SK1BS\x000SK1BV\x000SK1C\x000SK1E(\x000SK1E1\x000SK1EF\x000SK1EK\x000SK1EN\x000SK1ES\x000SK1EV\x000SK1O(\x000SK1OF\x000SK1OS\x000SK1OV\x000SK1U(\x000SK1UE\x000SKF()\x000SKF(1\x000SKF(F\x000SKF(N\x000SKF(S\x000SKF(V\x000SKN\x000SKN&(\x000SKN&1\x000SKN&F\x000SKN&N\x000SKN&S\x000SKN&V\x000SKN;\x000SKN;C\x000SKN;E\x000SKN;T\x000SKNB(\x000SKNB1\x000SKNBF\x000SKNBN\x000SKNBS\x000SKNBV\x000SKNC\x000SKNE(\x000SKNE1\x000SKNEF\x000SKNEN\x000SKNES\x000SKNEV\x000SKNU(\x000SKNUE\x000SKS\x000SKS&(\x000SKS&1\x000SKS&F\x000SKS&N\x000SKS&S\x000SKS&V\x000SKS;\x000SKS;C\x000SKS;E\x000SKS;T\x000SKSB(\x000SKSB1\x000SKSBF\x000SKSBN\x000SKSBS\x000SKSBV\x000SKSC\x000SKSE(\x000SKSE1\x000SKSEF\x000SKSEK\x000SKSEN\x000SKSES\x000SKSEV\x000SKSO(\x000SKSO1\x000SKSOF\x000SKSON\x000SKSOS\x000SKSOV\x000SKSU(\x000SKSUE\x000SKUE(\x000SKUE1\x000SKUEF\x000SKUEK\x000SKUEN\x000SKUES\x000SKUEV\x000SKV\x000SKV&(\x000SKV&1\x000SKV&F\x000SKV&N\x000SKV&S\x000SKV&V\x000SKV;\x000SKV;C\x000SKV;E\x000SKV;T\x000SKVB(\x000SKVB1\x000SKVBF\x000SKVBN\x000SKVBS\x000SKVBV\x000SKVC\x000SKVE(\x000SKVE1\x000SKVEF\x000SKVEK\x000SKVEN\x000SKVES\x000SKVEV\x000SKVO(\x000SKVOF\x000SKVOS\x000SKVU(\x000SKVUE\x000SO(1&\x000SO(1)\x000SO(1,\x000SO(1O\x000SO(E(\x000SO(E1\x000SO(EE\x000SO(EF\x000SO(EK\x000SO(EN\x000SO(EO\x000SO(ES\x000SO(EV\x000SO(F(\x000SO(N&\x000SO(N)\x000SO(N,\x000SO(NO\x000SO(S&\x000SO(S)\x000SO(S,\x000SO(SO\x000SO(V&\x000SO(V)\x000SO(V,\x000SO(VO\x000SO1&(\x000SO1&1\x000SO1&E\x000SO1&F\x000SO1&K\x000SO1&N\x000SO1&S\x000SO1&U\x000SO1&V\x000SO1(E\x000SO1(U\x000SO1)&\x000SO1),\x000SO1);\x000SO1)B\x000SO1)C\x000SO1)E\x000SO1)F\x000SO1)K\x000SO1)O\x000SO1)U\x000SO1,(\x000SO1,F\x000SO1;\x000SO1;C\x000SO1;E\x000SO1;N\x000SO1;T\x000SO1A(\x000SO1AF\x000SO1AS\x000SO1AT\x000SO1AV\x000SO1B(\x000SO1B1\x000SO1BE\x000SO1BF\x000SO1BN\x000SO1BS\x000SO1BV\x000SO1C\x000SO1E(\x000SO1E1\x000SO1EF\x000SO1EK\x000SO1EN\x000SO1EO\x000SO1ES\x000SO1EU\x000SO1EV\x000SO1F(\x000SO1K(\x000SO1K)\x000SO1K1\x000SO1KB\x000SO1KF\x000SO1KN\x000SO1KS\x000SO1KU\x000SO1KV\x000SO1N&\x000SO1N(\x000SO1N,\x000SO1NE\x000SO1NU\x000SO1SU\x000SO1SV\x000SO1T(\x000SO1T1\x000SO1TE\x000SO1TF\x000SO1TN\x000SO1TS\x000SO1TT\x000SO1TV\x000SO1U\x000SO1U(\x000SO1U1\x000SO1U;\x000SO1UC\x000SO1UE\x000SO1UF\x000SO1UK\x000SO1UO\x000SO1US\x000SO1UT\x000SO1UV\x000SO1V(\x000SO1VF\x000SO1VO\x000SO1VS\x000SO1VU\x000SOF()\x000SOF(1\x000SOF(E\x000SOF(F\x000SOF(N\x000SOF(S\x000SOF(V\x000SOK&(\x000SOK&1\x000SOK&F\x000SOK&N\x000SOK&S\x000SOK&V\x000SOK(1\x000SOK(F\x000SOK(N\x000SOK(S\x000SOK(V\x000SOK1C\x000SOK1O\x000SOKF(\x000SOKNC\x000SOKO(\x000SOKO1\x000SOKOF\x000SOKON\x000SOKOS\x000SOKOV\x000SOKSC\x000SOKSO\x000SOKVC\x000SOKVO\x000SON&(\x000SON&1\x000SON&E\x000SON&F\x000SON&K\x000SON&N\x000SON&S\x000SON&U\x000SON&V\x000SON(1\x000SON(E\x000SON(F\x000SON(S\x000SON(U\x000SON(V\x000SON)&\x000SON),\x000SON);\x000SON)B\x000SON)C\x000SON)E\x000SON)F\x000SON)K\x000SON)O\x000SON)U\x000SON,(\x000SON,F\x000SON1(\x000SON1O\x000SON1U\x000SON1V\x000SON;\x000SON;C\x000SON;E\x000SON;N\x000SON;T\x000SONA(\x000SONAF\x000SONAS\x000SONAT\x000SONAV\x000SONB(\x000SONB1\x000SONBE\x000SONBF\x000SONBN\x000SONBS\x000SONBV\x000SONE(\x000SONE1\x000SONEF\x000SONEN\x000SONEO\x000SONES\x000SONEU\x000SONEV\x000SONF(\x000SONK(\x000SONK)\x000SONK1\x000SONKB\x000SONKF\x000SONKS\x000SONKU\x000SONKV\x000SONSU\x000SONT(\x000SONT1\x000SONTE\x000SONTF\x000SONTN\x000SONTS\x000SONTT\x000SONTV\x000SONU\x000SONU(\x000SONU1\x000SONU;\x000SONUC\x000SONUE\x000SONUF\x000SONUK\x000SONUO\x000SONUS\x000SONUT\x000SONUV\x000SOS\x000SOS&(\x000SOS&1\x000SOS&E\x000SOS&F\x000SOS&K\x000SOS&N\x000SOS&S\x000SOS&U\x000SOS&V\x000SOS(E\x000SOS(U\x000SOS)&\x000SOS),\x000SOS);\x000SOS)B\x000SOS)C\x000SOS)E\x000SOS)F\x000SOS)K\x000SOS)O\x000SOS)U\x000SOS,(\x000SOS,F\x000SOS1(\x000SOS1F\x000SOS1N\x000SOS1S\x000SOS1U\x000SOS1V\x000SOS;\x000SOS;C\x000SOS;E\x000SOS;N\x000SOS;T\x000SOSA(\x000SOSAF\x000SOSAS\x000SOSAT\x000SOSAV\x000SOSB(\x000SOSB1\x000SOSBE\x000SOSBF\x000SOSBN\x000SOSBS\x000SOSBV\x000SOSC\x000SOSE(\x000SOSE1\x000SOSEF\x000SOSEK\x000SOSEN\x000SOSEO\x000SOSES\x000SOSEU\x000SOSEV\x000SOSF(\x000SOSK(\x000SOSK)\x000SOSK1\x000SOSKB\x000SOSKF\x000SOSKN\x000SOSKS\x000SOSKU\x000SOSKV\x000SOST(\x000SOST1\x000SOSTE\x000SOSTF\x000SOSTN\x000SOSTS\x000SOSTT\x000SOSTV\x000SOSU\x000SOSU(\x000SOSU1\x000SOSU;\x000SOSUC\x000SOSUE\x000SOSUF\x000SOSUK\x000SOSUO\x000SOSUS\x000SOSUT\x000SOSUV\x000SOSV(\x000SOSVF\x000SOSVO\x000SOSVS\x000SOSVU\x000SOU(E\x000SOUEK\x000SOUEN\x000SOV\x000SOV&(\x000SOV&1\x000SOV&E\x000SOV&F\x000SOV&K\x000SOV&N\x000SOV&S\x000SOV&U\x000SOV&V\x000SOV(E\x000SOV(U\x000SOV)&\x000SOV),\x000SOV);\x000SOV)B\x000SOV)C\x000SOV)E\x000SOV)F\x000SOV)K\x000SOV)O\x000SOV)U\x000SOV,(\x000SOV,F\x000SOV;\x000SOV;C\x000SOV;E\x000SOV;N\x000SOV;T\x000SOVA(\x000SOVAF\x000SOVAS\x000SOVAT\x000SOVAV\x000SOVB(\x000SOVB1\x000SOVBE\x000SOVBF\x000SOVBN\x000SOVBS\x000SOVBV\x000SOVC\x000SOVE(\x000SOVE1\x000SOVEF\x000SOVEK\x000SOVEN\x000SOVEO\x000SOVES\x000SOVEU\x000SOVEV\x000SOVF(\x000SOVK(\x000SOVK)\x000SOVK1\x000SOVKB\x000SOVKF\x000SOVKN\x000SOVKS\x000SOVKU\x000SOVKV\x000SOVO(\x000SOVOF\x000SOVOK\x000SOVOS\x000SOVOU\x000SOVS(\x000SOVS1\x000SOVSF\x000SOVSO\x000SOVSU\x000SOVSV\x000SOVT(\x000SOVT1\x000SOVTE\x000SOVTF\x000SOVTN\x000SOVTS\x000SOVTT\x000SOVTV\x000SOVU\x000SOVU(\x000SOVU1\x000SOVU;\x000SOVUC\x000SOVUE\x000SOVUF\x000SOVUK\x000SOVUO\x000SOVUS\x000SOVUT\x000SOVUV\x000ST(1)\x000ST(1O\x000ST(F(\x000ST(N)\x000ST(NO\x000ST(S)\x000ST(SO\x000ST(V)\x000ST(VO\x000ST1(F\x000ST1O(\x000ST1OF\x000ST1OS\x000ST1OV\x000STE(1\x000STE(F\x000STE(N\x000STE(S\x000STE(V\x000STE1N\x000STE1O\x000STEF(\x000STEK(\x000STEK1\x000STEKF\x000STEKN\x000STEKS\x000STEKV\x000STENN\x000STENO\x000STESN\x000STESO\x000STEVN\x000STEVO\x000STF()\x000STF(1\x000STF(F\x000STF(N\x000STF(S\x000STF(V\x000STN(1\x000STN(F\x000STN(S\x000STN(V\x000STN1C\x000STN1O\x000STN;E\x000STN;N\x000STN;T\x000STNE(\x000STNE1\x000STNEF\x000STNEN\x000STNES\x000STNEV\x000STNF(\x000STNKN\x000STNN:\x000STNNC\x000STNNO\x000STNO(\x000STNOF\x000STNOS\x000STNOV\x000STNSC\x000STNSO\x000STNT(\x000STNT1\x000STNTF\x000STNTN\x000STNTS\x000STNTV\x000STNVC\x000STNVO\x000STS(F\x000STSO(\x000STSO1\x000STSOF\x000STSON\x000STSOS\x000STSOV\x000STTNE\x000STTNK\x000STTNN\x000STTNT\x000STV(1\x000STV(F\x000STVO(\x000STVOF\x000STVOS\x000SU(1)\x000SU(1O\x000SU(E(\x000SU(E1\x000SU(EF\x000SU(EK\x000SU(EN\x000SU(ES\x000SU(EV\x000SU(F(\x000SU(N)\x000SU(NO\x000SU(S)\x000SU(SO\x000SU(V)\x000SU(VO\x000SU1,(\x000SU1,F\x000SU1C\x000SU1O(\x000SU1OF\x000SU1OS\x000SU1OV\x000SU;\x000SU;C\x000SUC\x000SUE\x000SUE(1\x000SUE(E\x000SUE(F\x000SUE(N\x000SUE(O\x000SUE(S\x000SUE(V\x000SUE1\x000SUE1&\x000SUE1(\x000SUE1)\x000SUE1,\x000SUE1;\x000SUE1B\x000SUE1C\x000SUE1F\x000SUE1K\x000SUE1N\x000SUE1O\x000SUE1S\x000SUE1U\x000SUE1V\x000SUE;\x000SUE;C\x000SUEC\x000SUEF\x000SUEF(\x000SUEF,\x000SUEF;\x000SUEFC\x000SUEK\x000SUEK(\x000SUEK1\x000SUEK;\x000SUEKC\x000SUEKF\x000SUEKN\x000SUEKO\x000SUEKS\x000SUEKV\x000SUEN\x000SUEN&\x000SUEN(\x000SUEN)\x000SUEN,\x000SUEN1\x000SUEN;\x000SUENB\x000SUENC\x000SUENF\x000SUENK\x000SUENO\x000SUENS\x000SUENU\x000SUEOK\x000SUEON\x000SUES\x000SUES&\x000SUES(\x000SUES)\x000SUES,\x000SUES1\x000SUES;\x000SUESB\x000SUESC\x000SUESF\x000SUESK\x000SUESO\x000SUESU\x000SUESV\x000SUEV\x000SUEV&\x000SUEV(\x000SUEV)\x000SUEV,\x000SUEV;\x000SUEVB\x000SUEVC\x000SUEVF\x000SUEVK\x000SUEVN\x000SUEVO\x000SUEVS\x000SUEVU\x000SUF()\x000SUF(1\x000SUF(F\x000SUF(N\x000SUF(S\x000SUF(V\x000SUK(E\x000SUO(E\x000SUON(\x000SUON1\x000SUONF\x000SUONS\x000SUS,(\x000SUS,F\x000SUSC\x000SUSO(\x000SUSO1\x000SUSOF\x000SUSON\x000SUSOS\x000SUSOV\x000SUTN(\x000SUTN1\x000SUTNF\x000SUTNN\x000SUTNS\x000SUTNV\x000SUV,(\x000SUV,F\x000SUVC\x000SUVO(\x000SUVOF\x000SUVOS\x000SVF()\x000SVF(1\x000SVF(F\x000SVF(N\x000SVF(S\x000SVF(V\x000SVO(1\x000SVO(F\x000SVO(N\x000SVO(S\x000SVO(V\x000SVOF(\x000SVOS(\x000SVOS1\x000SVOSF\x000SVOSU\x000SVOSV\x000SVS;\x000SVS;C\x000SVSC\x000SVSO(\x000SVSO1\x000SVSOF\x000SVSON\x000SVSOS\x000SVSOV\x000SVUE\x000SVUE;\x000SVUEC\x000SVUEK\x000T(1)F\x000T(1)O\x000T(1F(\x000T(1N)\x000T(1O(\x000T(1OF\x000T(1OS\x000T(1OV\x000T(1S)\x000T(1V)\x000T(1VO\x000T(F()\x000T(F(1\x000T(F(F\x000T(F(N\x000T(F(S\x000T(F(V\x000T(N(1\x000T(N(F\x000T(N(S\x000T(N(V\x000T(N)F\x000T(N)O\x000T(N1)\x000T(N1O\x000T(NF(\x000T(NN)\x000T(NNO\x000T(NO(\x000T(NOF\x000T(NOS\x000T(NOV\x000T(NS)\x000T(NSO\x000T(NV)\x000T(NVO\x000T(S)F\x000T(S)O\x000T(S1)\x000T(SF(\x000T(SN)\x000T(SNO\x000T(SO(\x000T(SO1\x000T(SOF\x000T(SON\x000T(SOS\x000T(SOV\x000T(SV)\x000T(SVO\x000T(V)F\x000T(V)O\x000T(VF(\x000T(VO(\x000T(VOF\x000T(VOS\x000T(VS)\x000T(VSO\x000T(VV)\x000T1F(1\x000T1F(F\x000T1F(N\x000T1F(S\x000T1F(V\x000T1O(1\x000T1O(F\x000T1O(N\x000T1O(S\x000T1O(V\x000T1OF(\x000T1OSF\x000T1OVF\x000T1OVO\x000TF()F\x000TF()O\x000TF(1)\x000TF(1O\x000TF(F(\x000TF(N)\x000TF(NO\x000TF(S)\x000TF(SO\x000TF(V)\x000TF(VO\x000TN(1)\x000TN(1O\x000TN(F(\x000TN(S)\x000TN(SO\x000TN(V)\x000TN(VO\x000TN1;\x000TN1;C\x000TN1O(\x000TN1OF\x000TN1OS\x000TN1OV\x000TNF()\x000TNF(1\x000TNF(F\x000TNF(N\x000TNF(S\x000TNF(V\x000TNN;\x000TNN;C\x000TNNO(\x000TNNOF\x000TNNOS\x000TNNOV\x000TNO(1\x000TNO(F\x000TNO(N\x000TNO(S\x000TNO(V\x000TNOF(\x000TNOSF\x000TNOVF\x000TNOVO\x000TNS;\x000TNS;C\x000TNSO(\x000TNSO1\x000TNSOF\x000TNSON\x000TNSOS\x000TNSOV\x000TNV;\x000TNV;C\x000TNVO(\x000TNVOF\x000TNVOS\x000TSF(1\x000TSF(F\x000TSF(N\x000TSF(S\x000TSF(V\x000TSO(1\x000TSO(F\x000TSO(N\x000TSO(S\x000TSO(V\x000TSO1F\x000TSOF(\x000TSONF\x000TSOSF\x000TSOVF\x000TSOVO\x000TVF(1\x000TVF(F\x000TVF(N\x000TVF(S\x000TVF(V\x000TVO(1\x000TVO(F\x000TVO(N\x000TVO(S\x000TVO(V\x000TVOF(\x000TVOSF\x000U(E(1\x000U(E(F\x000U(E(K\x000U(E(N\x000U(E(S\x000U(E(V\x000U(E1)\x000U(E1O\x000U(EF(\x000U(EK(\x000U(EK1\x000U(EKF\x000U(EKN\x000U(EKO\x000U(EKS\x000U(EKV\x000U(EN)\x000U(ENK\x000U(ENO\x000U(EOK\x000U(ES)\x000U(ESO\x000U(EV)\x000U(EVO\x000UE(1)\x000UE(1,\x000UE(1O\x000UE(F(\x000UE(N)\x000UE(N,\x000UE(NO\x000UE(S)\x000UE(S,\x000UE(SO\x000UE(V)\x000UE(V,\x000UE(VO\x000UE1\x000UE1,(\x000UE1,F\x000UE1;\x000UE1;C\x000UE1C\x000UE1K(\x000UE1K1\x000UE1KF\x000UE1KN\x000UE1KS\x000UE1KV\x000UE1O(\x000UE1OF\x000UE1OS\x000UE1OV\x000UEF()\x000UEF(1\x000UEF(F\x000UEF(N\x000UEF(S\x000UEF(V\x000UEK(1\x000UEK(F\x000UEK(N\x000UEK(S\x000UEK(V\x000UEK1\x000UEK1,\x000UEK1;\x000UEK1C\x000UEK1K\x000UEK1O\x000UEKF(\x000UEKN\x000UEKN(\x000UEKN,\x000UEKN;\x000UEKNC\x000UEKNK\x000UEKS\x000UEKS,\x000UEKS;\x000UEKSC\x000UEKSK\x000UEKSO\x000UEKV\x000UEKV,\x000UEKV;\x000UEKVC\x000UEKVK\x000UEKVO\x000UEN()\x000UEN,(\x000UEN,F\x000UEN;\x000UEN;C\x000UENC\x000UENK(\x000UENK1\x000UENKF\x000UENKN\x000UENKS\x000UENKV\x000UENO(\x000UENOF\x000UENOS\x000UENOV\x000UES\x000UES,(\x000UES,F\x000UES;\x000UES;C\x000UESC\x000UESK(\x000UESK1\x000UESKF\x000UESKN\x000UESKS\x000UESKV\x000UESO(\x000UESO1\x000UESOF\x000UESON\x000UESOS\x000UESOV\x000UEV\x000UEV,(\x000UEV,F\x000UEV;\x000UEV;C\x000UEVC\x000UEVK(\x000UEVK1\x000UEVKF\x000UEVKN\x000UEVKS\x000UEVKV\x000UEVO(\x000UEVOF\x000UEVOS\x000UF(1O\x000UF(F(\x000UF(NO\x000UF(SO\x000UF(VO\x000V&(1&\x000V&(1)\x000V&(1,\x000V&(1O\x000V&(E(\x000V&(E1\x000V&(EF\x000V&(EK\x000V&(EN\x000V&(EO\x000V&(ES\x000V&(EV\x000V&(F(\x000V&(N&\x000V&(N)\x000V&(N,\x000V&(NO\x000V&(S&\x000V&(S)\x000V&(S,\x000V&(SO\x000V&(V&\x000V&(V)\x000V&(V,\x000V&(VO\x000V&1\x000V&1&(\x000V&1&1\x000V&1&F\x000V&1&N\x000V&1&S\x000V&1&V\x000V&1)&\x000V&1)C\x000V&1)O\x000V&1)U\x000V&1;\x000V&1;C\x000V&1;E\x000V&1;T\x000V&1B(\x000V&1B1\x000V&1BF\x000V&1BN\x000V&1BS\x000V&1BV\x000V&1C\x000V&1EK\x000V&1EN\x000V&1F(\x000V&1K(\x000V&1K1\x000V&1KF\x000V&1KN\x000V&1KS\x000V&1KV\x000V&1O(\x000V&1OF\x000V&1OS\x000V&1OV\x000V&1TN\x000V&1U\x000V&1U(\x000V&1U;\x000V&1UC\x000V&1UE\x000V&E(1\x000V&E(F\x000V&E(N\x000V&E(O\x000V&E(S\x000V&E(V\x000V&E1\x000V&E1;\x000V&E1C\x000V&E1K\x000V&E1O\x000V&EF(\x000V&EK(\x000V&EK1\x000V&EKF\x000V&EKN\x000V&EKS\x000V&EKV\x000V&EN\x000V&EN;\x000V&ENC\x000V&ENK\x000V&ENO\x000V&ES\x000V&ES;\x000V&ESC\x000V&ESK\x000V&ESO\x000V&EV\x000V&EV;\x000V&EVC\x000V&EVK\x000V&EVO\x000V&F()\x000V&F(1\x000V&F(E\x000V&F(F\x000V&F(N\x000V&F(S\x000V&F(V\x000V&K&(\x000V&K&1\x000V&K&F\x000V&K&N\x000V&K&S\x000V&K&V\x000V&K(1\x000V&K(F\x000V&K(N\x000V&K(S\x000V&K(V\x000V&K1O\x000V&KC\x000V&KF(\x000V&KNK\x000V&KO(\x000V&KO1\x000V&KOF\x000V&KOK\x000V&KON\x000V&KOS\x000V&KOV\x000V&KSO\x000V&KVO\x000V&N\x000V&N&(\x000V&N&1\x000V&N&F\x000V&N&N\x000V&N&S\x000V&N&V\x000V&N)&\x000V&N)C\x000V&N)O\x000V&N)U\x000V&N;\x000V&N;C\x000V&N;E\x000V&N;T\x000V&NB(\x000V&NB1\x000V&NBF\x000V&NBN\x000V&NBS\x000V&NBV\x000V&NC\x000V&NEN\x000V&NF(\x000V&NK(\x000V&NK1\x000V&NKF\x000V&NKN\x000V&NKS\x000V&NKV\x000V&NO(\x000V&NOF\x000V&NOS\x000V&NOV\x000V&NTN\x000V&NU\x000V&NU(\x000V&NU;\x000V&NUC\x000V&NUE\x000V&S\x000V&S&(\x000V&S&1\x000V&S&F\x000V&S&N\x000V&S&S\x000V&S&V\x000V&S)&\x000V&S)C\x000V&S)O\x000V&S)U\x000V&S1\x000V&S1;\x000V&S1C\x000V&S;\x000V&S;C\x000V&S;E\x000V&S;T\x000V&SB(\x000V&SB1\x000V&SBF\x000V&SBN\x000V&SBS\x000V&SBV\x000V&SC\x000V&SEK\x000V&SEN\x000V&SF(\x000V&SK(\x000V&SK1\x000V&SKF\x000V&SKN\x000V&SKS\x000V&SKV\x000V&SO(\x000V&SO1\x000V&SOF\x000V&SON\x000V&SOS\x000V&SOV\x000V&STN\x000V&SU\x000V&SU(\x000V&SU;\x000V&SUC\x000V&SUE\x000V&SV\x000V&SV;\x000V&SVC\x000V&SVO\x000V&V\x000V&V&(\x000V&V&1\x000V&V&F\x000V&V&N\x000V&V&S\x000V&V&V\x000V&V)&\x000V&V)C\x000V&V)O\x000V&V)U\x000V&V;\x000V&V;C\x000V&V;E\x000V&V;T\x000V&VB(\x000V&VB1\x000V&VBF\x000V&VBN\x000V&VBS\x000V&VBV\x000V&VC\x000V&VEK\x000V&VEN\x000V&VF(\x000V&VK(\x000V&VK1\x000V&VKF\x000V&VKN\x000V&VKS\x000V&VKV\x000V&VO(\x000V&VOF\x000V&VOS\x000V&VS\x000V&VS;\x000V&VSC\x000V&VSO\x000V&VTN\x000V&VU\x000V&VU(\x000V&VU;\x000V&VUC\x000V&VUE\x000V(EF(\x000V(EKF\x000V(EKN\x000V(ENK\x000V(U(E\x000V)&(1\x000V)&(E\x000V)&(F\x000V)&(N\x000V)&(S\x000V)&(V\x000V)&1\x000V)&1&\x000V)&1)\x000V)&1;\x000V)&1B\x000V)&1C\x000V)&1F\x000V)&1O\x000V)&1U\x000V)&F(\x000V)&N\x000V)&N&\x000V)&N)\x000V)&N;\x000V)&NB\x000V)&NC\x000V)&NF\x000V)&NO\x000V)&NU\x000V)&S\x000V)&S&\x000V)&S)\x000V)&S;\x000V)&SB\x000V)&SC\x000V)&SF\x000V)&SO\x000V)&SU\x000V)&V\x000V)&V&\x000V)&V)\x000V)&V;\x000V)&VB\x000V)&VC\x000V)&VF\x000V)&VO\x000V)&VU\x000V),(1\x000V),(F\x000V),(N\x000V),(S\x000V),(V\x000V);E(\x000V);E1\x000V);EF\x000V);EK\x000V);EN\x000V);EO\x000V);ES\x000V);EV\x000V);T(\x000V);T1\x000V);TF\x000V);TK\x000V);TN\x000V);TO\x000V);TS\x000V);TV\x000V)B(1\x000V)B(F\x000V)B(N\x000V)B(S\x000V)B(V\x000V)B1\x000V)B1&\x000V)B1;\x000V)B1C\x000V)B1K\x000V)B1N\x000V)B1O\x000V)B1U\x000V)BF(\x000V)BN\x000V)BN&\x000V)BN;\x000V)BNC\x000V)BNK\x000V)BNO\x000V)BNU\x000V)BS\x000V)BS&\x000V)BS;\x000V)BSC\x000V)BSK\x000V)BSO\x000V)BSU\x000V)BV\x000V)BV&\x000V)BV;\x000V)BVC\x000V)BVK\x000V)BVO\x000V)BVU\x000V)C\x000V)E(1\x000V)E(F\x000V)E(N\x000V)E(S\x000V)E(V\x000V)E1C\x000V)E1O\x000V)EF(\x000V)EK(\x000V)EK1\x000V)EKF\x000V)EKN\x000V)EKS\x000V)EKV\x000V)ENC\x000V)ENO\x000V)ESC\x000V)ESO\x000V)EVC\x000V)EVO\x000V)F(F\x000V)K(1\x000V)K(F\x000V)K(N\x000V)K(S\x000V)K(V\x000V)K1&\x000V)K1;\x000V)K1B\x000V)K1E\x000V)K1O\x000V)K1U\x000V)KB(\x000V)KB1\x000V)KBF\x000V)KBN\x000V)KBS\x000V)KBV\x000V)KF(\x000V)KN&\x000V)KN;\x000V)KNB\x000V)KNC\x000V)KNE\x000V)KNK\x000V)KNU\x000V)KS&\x000V)KS;\x000V)KSB\x000V)KSE\x000V)KSO\x000V)KSU\x000V)KUE\x000V)KV&\x000V)KV;\x000V)KVB\x000V)KVE\x000V)KVO\x000V)KVU\x000V)O(1\x000V)O(E\x000V)O(F\x000V)O(N\x000V)O(S\x000V)O(V\x000V)O1\x000V)O1&\x000V)O1)\x000V)O1;\x000V)O1B\x000V)O1C\x000V)O1K\x000V)O1U\x000V)OF(\x000V)ON\x000V)ON&\x000V)ON)\x000V)ON;\x000V)ONB\x000V)ONC\x000V)ONK\x000V)ONU\x000V)OS\x000V)OS&\x000V)OS)\x000V)OS;\x000V)OSB\x000V)OSC\x000V)OSK\x000V)OSU\x000V)OV\x000V)OV&\x000V)OV)\x000V)OV;\x000V)OVB\x000V)OVC\x000V)OVK\x000V)OVO\x000V)OVU\x000V)U(E\x000V)UE(\x000V)UE1\x000V)UEF\x000V)UEK\x000V)UEN\x000V)UES\x000V)UEV\x000V,(1)\x000V,(1O\x000V,(E(\x000V,(E1\x000V,(EF\x000V,(EK\x000V,(EN\x000V,(ES\x000V,(EV\x000V,(F(\x000V,(N)\x000V,(NO\x000V,(S)\x000V,(SO\x000V,(V)\x000V,(VO\x000V,F()\x000V,F(1\x000V,F(F\x000V,F(N\x000V,F(S\x000V,F(V\x000V;E(1\x000V;E(E\x000V;E(F\x000V;E(N\x000V;E(S\x000V;E(V\x000V;E1,\x000V;E1;\x000V;E1C\x000V;E1K\x000V;E1O\x000V;E1T\x000V;EF(\x000V;EK(\x000V;EK1\x000V;EKF\x000V;EKN\x000V;EKO\x000V;EKS\x000V;EKV\x000V;EN,\x000V;EN;\x000V;ENC\x000V;ENE\x000V;ENK\x000V;ENO\x000V;ENT\x000V;ES,\x000V;ES;\x000V;ESC\x000V;ESK\x000V;ESO\x000V;EST\x000V;EV,\x000V;EV;\x000V;EVC\x000V;EVK\x000V;EVO\x000V;EVT\x000V;N:T\x000V;T(1\x000V;T(C\x000V;T(E\x000V;T(F\x000V;T(N\x000V;T(S\x000V;T(V\x000V;T1(\x000V;T1,\x000V;T1;\x000V;T1C\x000V;T1F\x000V;T1K\x000V;T1O\x000V;T1T\x000V;T;\x000V;T;C\x000V;TF(\x000V;TK(\x000V;TK1\x000V;TKF\x000V;TKK\x000V;TKN\x000V;TKO\x000V;TKS\x000V;TKV\x000V;TN(\x000V;TN,\x000V;TN1\x000V;TN;\x000V;TNC\x000V;TNE\x000V;TNF\x000V;TNK\x000V;TNN\x000V;TNO\x000V;TNS\x000V;TNT\x000V;TNV\x000V;TO(\x000V;TS(\x000V;TS,\x000V;TS;\x000V;TSC\x000V;TSF\x000V;TSK\x000V;TSO\x000V;TST\x000V;TTN\x000V;TV(\x000V;TV,\x000V;TV;\x000V;TVC\x000V;TVF\x000V;TVK\x000V;TVO\x000V;TVT\x000VA(F(\x000VA(N)\x000VA(NO\x000VA(S)\x000VA(SO\x000VA(V)\x000VA(VO\x000VAF()\x000VAF(1\x000VAF(F\x000VAF(N\x000VAF(S\x000VAF(V\x000VASO(\x000VASO1\x000VASOF\x000VASON\x000VASOS\x000VASOV\x000VASUE\x000VATO(\x000VATO1\x000VATOF\x000VATON\x000VATOS\x000VATOV\x000VATUE\x000VAVO(\x000VAVOF\x000VAVOS\x000VAVUE\x000VB(1)\x000VB(1O\x000VB(F(\x000VB(NO\x000VB(S)\x000VB(SO\x000VB(V)\x000VB(VO\x000VB1\x000VB1&(\x000VB1&1\x000VB1&F\x000VB1&N\x000VB1&S\x000VB1&V\x000VB1,(\x000VB1,F\x000VB1;\x000VB1;C\x000VB1B(\x000VB1B1\x000VB1BF\x000VB1BN\x000VB1BS\x000VB1BV\x000VB1C\x000VB1K(\x000VB1K1\x000VB1KF\x000VB1KN\x000VB1KS\x000VB1KV\x000VB1O(\x000VB1OF\x000VB1OS\x000VB1OV\x000VB1U(\x000VB1UE\x000VBE(1\x000VBE(F\x000VBE(N\x000VBE(S\x000VBE(V\x000VBEK(\x000VBF()\x000VBF(1\x000VBF(F\x000VBF(N\x000VBF(S\x000VBF(V\x000VBN\x000VBN&(\x000VBN&1\x000VBN&F\x000VBN&N\x000VBN&S\x000VBN&V\x000VBN,(\x000VBN,F\x000VBN;\x000VBN;C\x000VBNB(\x000VBNB1\x000VBNBF\x000VBNBN\x000VBNBS\x000VBNBV\x000VBNC\x000VBNK(\x000VBNK1\x000VBNKF\x000VBNKN\x000VBNKS\x000VBNKV\x000VBNO(\x000VBNOF\x000VBNOS\x000VBNOV\x000VBNU(\x000VBNUE\x000VBS\x000VBS&(\x000VBS&1\x000VBS&F\x000VBS&N\x000VBS&S\x000VBS&V\x000VBS,(\x000VBS,F\x000VBS;\x000VBS;C\x000VBSB(\x000VBSB1\x000VBSBF\x000VBSBN\x000VBSBS\x000VBSBV\x000VBSC\x000VBSK(\x000VBSK1\x000VBSKF\x000VBSKN\x000VBSKS\x000VBSKV\x000VBSO(\x000VBSO1\x000VBSOF\x000VBSON\x000VBSOS\x000VBSOV\x000VBSU(\x000VBSUE\x000VBV\x000VBV&(\x000VBV&1\x000VBV&F\x000VBV&N\x000VBV&S\x000VBV&V\x000VBV,(\x000VBV,F\x000VBV;\x000VBV;C\x000VBVB(\x000VBVB1\x000VBVBF\x000VBVBN\x000VBVBS\x000VBVBV\x000VBVC\x000VBVK(\x000VBVK1\x000VBVKF\x000VBVKN\x000VBVKS\x000VBVKV\x000VBVO(\x000VBVOF\x000VBVOS\x000VBVU(\x000VBVUE\x000VC\x000VE(1)\x000VE(1O\x000VE(F(\x000VE(N)\x000VE(NO\x000VE(S)\x000VE(SO\x000VE(V)\x000VE(VO\x000VE1;T\x000VE1C\x000VE1O(\x000VE1OF\x000VE1OS\x000VE1OV\x000VE1T(\x000VE1T1\x000VE1TF\x000VE1TN\x000VE1TS\x000VE1TV\x000VE1UE\x000VEF()\x000VEF(1\x000VEF(F\x000VEF(N\x000VEF(S\x000VEF(V\x000VEK(1\x000VEK(E\x000VEK(F\x000VEK(N\x000VEK(S\x000VEK(V\x000VEK1;\x000VEK1C\x000VEK1O\x000VEK1T\x000VEK1U\x000VEKF(\x000VEKN;\x000VEKNC\x000VEKNE\x000VEKNT\x000VEKNU\x000VEKOK\x000VEKS;\x000VEKSC\x000VEKSO\x000VEKST\x000VEKSU\x000VEKU(\x000VEKU1\x000VEKUE\x000VEKUF\x000VEKUS\x000VEKUV\x000VEKV;\x000VEKVC\x000VEKVO\x000VEKVT\x000VEKVU\x000VEN;T\x000VENC\x000VENEN\x000VENO(\x000VENOF\x000VENOS\x000VENOV\x000VENT(\x000VENT1\x000VENTF\x000VENTN\x000VENTS\x000VENTV\x000VENUE\x000VEOKN\x000VES;T\x000VESC\x000VESO(\x000VESO1\x000VESOF\x000VESON\x000VESOS\x000VESOV\x000VEST(\x000VEST1\x000VESTF\x000VESTN\x000VESTS\x000VESTV\x000VESUE\x000VEU(1\x000VEU(F\x000VEU(N\x000VEU(S\x000VEU(V\x000VEU1,\x000VEU1C\x000VEU1O\x000VEUEF\x000VEUEK\x000VEUF(\x000VEUS,\x000VEUSC\x000VEUSO\x000VEUV,\x000VEUVC\x000VEUVO\x000VEV;T\x000VEVC\x000VEVO(\x000VEVOF\x000VEVOS\x000VEVT(\x000VEVT1\x000VEVTF\x000VEVTN\x000VEVTS\x000VEVTV\x000VEVUE\x000VF()1\x000VF()F\x000VF()K\x000VF()N\x000VF()O\x000VF()S\x000VF()U\x000VF()V\x000VF(1)\x000VF(1N\x000VF(1O\x000VF(E(\x000VF(E1\x000VF(EF\x000VF(EK\x000VF(EN\x000VF(ES\x000VF(EV\x000VF(F(\x000VF(N)\x000VF(N,\x000VF(NO\x000VF(S)\x000VF(SO\x000VF(V)\x000VF(VO\x000VK(1)\x000VK(1O\x000VK(F(\x000VK(N)\x000VK(NO\x000VK(S)\x000VK(SO\x000VK(V)\x000VK(VO\x000VK)&(\x000VK)&1\x000VK)&F\x000VK)&N\x000VK)&S\x000VK)&V\x000VK);E\x000VK);T\x000VK)B(\x000VK)B1\x000VK)BF\x000VK)BN\x000VK)BS\x000VK)BV\x000VK)E(\x000VK)E1\x000VK)EF\x000VK)EK\x000VK)EN\x000VK)ES\x000VK)EV\x000VK)F(\x000VK)O(\x000VK)OF\x000VK)UE\x000VK1\x000VK1&(\x000VK1&1\x000VK1&F\x000VK1&N\x000VK1&S\x000VK1&V\x000VK1;\x000VK1;C\x000VK1;E\x000VK1;T\x000VK1B(\x000VK1B1\x000VK1BF\x000VK1BN\x000VK1BS\x000VK1BV\x000VK1C\x000VK1E(\x000VK1E1\x000VK1EF\x000VK1EK\x000VK1EN\x000VK1ES\x000VK1EV\x000VK1O(\x000VK1OF\x000VK1OS\x000VK1OV\x000VK1U(\x000VK1UE\x000VKF()\x000VKF(1\x000VKF(F\x000VKF(N\x000VKF(S\x000VKF(V\x000VKN\x000VKN&(\x000VKN&1\x000VKN&F\x000VKN&N\x000VKN&S\x000VKN&V\x000VKN;\x000VKN;C\x000VKN;E\x000VKN;T\x000VKNB(\x000VKNB1\x000VKNBF\x000VKNBN\x000VKNBS\x000VKNBV\x000VKNC\x000VKNE(\x000VKNE1\x000VKNEF\x000VKNEN\x000VKNES\x000VKNEV\x000VKNU(\x000VKNUE\x000VKS\x000VKS&(\x000VKS&1\x000VKS&F\x000VKS&N\x000VKS&S\x000VKS&V\x000VKS;\x000VKS;C\x000VKS;E\x000VKS;T\x000VKSB(\x000VKSB1\x000VKSBF\x000VKSBN\x000VKSBS\x000VKSBV\x000VKSC\x000VKSE(\x000VKSE1\x000VKSEF\x000VKSEK\x000VKSEN\x000VKSES\x000VKSEV\x000VKSO(\x000VKSO1\x000VKSOF\x000VKSON\x000VKSOS\x000VKSOV\x000VKSU(\x000VKSUE\x000VKUE(\x000VKUE1\x000VKUEF\x000VKUEK\x000VKUEN\x000VKUES\x000VKUEV\x000VKV\x000VKV&(\x000VKV&1\x000VKV&F\x000VKV&N\x000VKV&S\x000VKV&V\x000VKV;\x000VKV;C\x000VKV;E\x000VKV;T\x000VKVB(\x000VKVB1\x000VKVBF\x000VKVBN\x000VKVBS\x000VKVBV\x000VKVC\x000VKVE(\x000VKVE1\x000VKVEF\x000VKVEK\x000VKVEN\x000VKVES\x000VKVEV\x000VKVO(\x000VKVOF\x000VKVOS\x000VKVU(\x000VKVUE\x000VO(1&\x000VO(1)\x000VO(1,\x000VO(1O\x000VO(E(\x000VO(E1\x000VO(EE\x000VO(EF\x000VO(EK\x000VO(EN\x000VO(EO\x000VO(ES\x000VO(EV\x000VO(F(\x000VO(N&\x000VO(N)\x000VO(N,\x000VO(NO\x000VO(S&\x000VO(S)\x000VO(S,\x000VO(SO\x000VO(V&\x000VO(V)\x000VO(V,\x000VO(VO\x000VOF()\x000VOF(1\x000VOF(E\x000VOF(F\x000VOF(N\x000VOF(S\x000VOF(V\x000VOK&(\x000VOK&1\x000VOK&F\x000VOK&N\x000VOK&S\x000VOK&V\x000VOK(1\x000VOK(F\x000VOK(N\x000VOK(S\x000VOK(V\x000VOK1C\x000VOK1O\x000VOKF(\x000VOKNC\x000VOKO(\x000VOKO1\x000VOKOF\x000VOKON\x000VOKOS\x000VOKOV\x000VOKSC\x000VOKSO\x000VOKVC\x000VOKVO\x000VOS\x000VOS&(\x000VOS&1\x000VOS&E\x000VOS&F\x000VOS&K\x000VOS&N\x000VOS&S\x000VOS&U\x000VOS&V\x000VOS(E\x000VOS(U\x000VOS)&\x000VOS),\x000VOS);\x000VOS)B\x000VOS)C\x000VOS)E\x000VOS)F\x000VOS)K\x000VOS)O\x000VOS)U\x000VOS,(\x000VOS,F\x000VOS1(\x000VOS1F\x000VOS1N\x000VOS1S\x000VOS1U\x000VOS1V\x000VOS;\x000VOS;C\x000VOS;E\x000VOS;N\x000VOS;T\x000VOSA(\x000VOSAF\x000VOSAS\x000VOSAT\x000VOSAV\x000VOSB(\x000VOSB1\x000VOSBE\x000VOSBF\x000VOSBN\x000VOSBS\x000VOSBV\x000VOSC\x000VOSE(\x000VOSE1\x000VOSEF\x000VOSEK\x000VOSEN\x000VOSEO\x000VOSES\x000VOSEU\x000VOSEV\x000VOSF(\x000VOSK(\x000VOSK)\x000VOSK1\x000VOSKB\x000VOSKF\x000VOSKN\x000VOSKS\x000VOSKU\x000VOSKV\x000VOST(\x000VOST1\x000VOSTE\x000VOSTF\x000VOSTN\x000VOSTS\x000VOSTT\x000VOSTV\x000VOSU\x000VOSU(\x000VOSU1\x000VOSU;\x000VOSUC\x000VOSUE\x000VOSUF\x000VOSUK\x000VOSUO\x000VOSUS\x000VOSUT\x000VOSUV\x000VOSV(\x000VOSVF\x000VOSVO\x000VOSVS\x000VOSVU\x000VOU(E\x000VOUEK\x000VOUEN\x000VT(1)\x000VT(1O\x000VT(F(\x000VT(N)\x000VT(NO\x000VT(S)\x000VT(SO\x000VT(V)\x000VT(VO\x000VT1(F\x000VT1O(\x000VT1OF\x000VT1OS\x000VT1OV\x000VTE(1\x000VTE(F\x000VTE(N\x000VTE(S\x000VTE(V\x000VTE1N\x000VTE1O\x000VTEF(\x000VTEK(\x000VTEK1\x000VTEKF\x000VTEKN\x000VTEKS\x000VTEKV\x000VTENN\x000VTENO\x000VTESN\x000VTESO\x000VTEVN\x000VTEVO\x000VTF()\x000VTF(1\x000VTF(F\x000VTF(N\x000VTF(S\x000VTF(V\x000VTN(1\x000VTN(F\x000VTN(S\x000VTN(V\x000VTN1C\x000VTN1O\x000VTN;E\x000VTN;N\x000VTN;T\x000VTNE(\x000VTNE1\x000VTNEF\x000VTNEN\x000VTNES\x000VTNEV\x000VTNF(\x000VTNKN\x000VTNN:\x000VTNNC\x000VTNNO\x000VTNO(\x000VTNOF\x000VTNOS\x000VTNOV\x000VTNSC\x000VTNSO\x000VTNT(\x000VTNT1\x000VTNTF\x000VTNTN\x000VTNTS\x000VTNTV\x000VTNVC\x000VTNVO\x000VTS(F\x000VTSO(\x000VTSO1\x000VTSOF\x000VTSON\x000VTSOS\x000VTSOV\x000VTTNE\x000VTTNK\x000VTTNN\x000VTTNT\x000VTV(1\x000VTV(F\x000VTVO(\x000VTVOF\x000VTVOS\x000VU\x000VU(1)\x000VU(1O\x000VU(E(\x000VU(E1\x000VU(EF\x000VU(EK\x000VU(EN\x000VU(ES\x000VU(EV\x000VU(F(\x000VU(N)\x000VU(NO\x000VU(S)\x000VU(SO\x000VU(V)\x000VU(VO\x000VU1,(\x000VU1,F\x000VU1C\x000VU1O(\x000VU1OF\x000VU1OS\x000VU1OV\x000VU;\x000VU;C\x000VUC\x000VUE\x000VUE(1\x000VUE(E\x000VUE(F\x000VUE(N\x000VUE(O\x000VUE(S\x000VUE(V\x000VUE1\x000VUE1&\x000VUE1(\x000VUE1)\x000VUE1,\x000VUE1;\x000VUE1B\x000VUE1C\x000VUE1F\x000VUE1K\x000VUE1N\x000VUE1O\x000VUE1S\x000VUE1U\x000VUE1V\x000VUE;\x000VUE;C\x000VUEC\x000VUEF\x000VUEF(\x000VUEF,\x000VUEF;\x000VUEFC\x000VUEK\x000VUEK(\x000VUEK1\x000VUEK;\x000VUEKC\x000VUEKF\x000VUEKN\x000VUEKO\x000VUEKS\x000VUEKV\x000VUEN\x000VUEN&\x000VUEN(\x000VUEN)\x000VUEN,\x000VUEN1\x000VUEN;\x000VUENB\x000VUENC\x000VUENF\x000VUENK\x000VUENO\x000VUENS\x000VUENU\x000VUEOK\x000VUEON\x000VUES\x000VUES&\x000VUES(\x000VUES)\x000VUES,\x000VUES1\x000VUES;\x000VUESB\x000VUESC\x000VUESF\x000VUESK\x000VUESO\x000VUESU\x000VUESV\x000VUEV\x000VUEV&\x000VUEV(\x000VUEV)\x000VUEV,\x000VUEV;\x000VUEVB\x000VUEVC\x000VUEVF\x000VUEVK\x000VUEVN\x000VUEVO\x000VUEVS\x000VUEVU\x000VUF()\x000VUF(1\x000VUF(F\x000VUF(N\x000VUF(S\x000VUF(V\x000VUK(E\x000VUO(E\x000VUON(\x000VUON1\x000VUONF\x000VUONS\x000VUS,(\x000VUS,F\x000VUSC\x000VUSO(\x000VUSO1\x000VUSOF\x000VUSON\x000VUSOS\x000VUSOV\x000VUTN(\x000VUTN1\x000VUTNF\x000VUTNN\x000VUTNS\x000VUTNV\x000VUV,(\x000VUV,F\x000VUVC\x000VUVO(\x000VUVOF\x000VUVOS\x000X\x00::\x00:=\x00<<\x00<=\x00<>\x00<@\x00>=\x00>>\x00@>\x00ABORT\x00ABS\x00ACCESSIBLE\x00ACOS\x00ADDDATE\x00ADDTIME\x00AES_DECRYPT\x00AES_ENCRYPT\x00AGAINST\x00AGE\x00ALL_USERS\x00ALTER\x00ALTER DOMAIN\x00ALTER TABLE\x00ANALYZE\x00AND\x00ANY\x00ANYARRAY\x00ANYELEMENT\x00ANYNONARRY\x00APPLOCK_MODE\x00APPLOCK_TEST\x00APP_NAME\x00ARRAY_AGG\x00ARRAY_CAT\x00ARRAY_DIM\x00ARRAY_FILL\x00ARRAY_LENGTH\x00ARRAY_LOWER\x00ARRAY_NDIMS\x00ARRAY_PREPEND\x00ARRAY_TO_JSON\x00ARRAY_TO_STRING\x00ARRAY_UPPER\x00AS\x00ASC\x00ASCII\x00ASENSITIVE\x00ASIN\x00ASSEMBLYPROPERTY\x00ASYMKEY_ID\x00AT TIME\x00AT TIME ZONE\x00ATAN\x00ATAN2\x00AUTOINCREMENT\x00AVG\x00BEFORE\x00BEGIN\x00BEGIN DECLARE\x00BEGIN GOTO\x00BEGIN TRY\x00BEGIN TRY DECLARE\x00BENCHMARK\x00BETWEEN\x00BIGINT\x00BIGSERIAL\x00BIN\x00BINARY\x00BINARY_DOUBLE_INFINITY\x00BINARY_DOUBLE_NAN\x00BINARY_FLOAT_INFINITY\x00BINARY_FLOAT_NAN\x00BINBINARY\x00BIT_AND\x00BIT_COUNT\x00BIT_LENGTH\x00BIT_OR\x00BIT_XOR\x00BLOB\x00BOOLEAN\x00BOOL_AND\x00BOOL_OR\x00BOTH\x00BTRIM\x00BY\x00BYTEA\x00CALL\x00CASCADE\x00CASE\x00CAST\x00CBOOL\x00CBRT\x00CBYTE\x00CCUR\x00CDATE\x00CDBL\x00CEIL\x00CEILING\x00CERTENCODED\x00CERTPRIVATEKEY\x00CERT_ID\x00CERT_PROPERTY\x00CHANGE\x00CHANGES\x00CHAR\x00CHARACTER\x00CHARACTER VARYING\x00CHARACTER_LENGTH\x00CHARINDEX\x00CHARSET\x00CHAR_LENGTH\x00CHDIR\x00CHDRIVE\x00CHECK\x00CHECKSUM_AGG\x00CHOOSE\x00CHR\x00CINT\x00CLNG\x00CLOCK_TIMESTAMP\x00COALESCE\x00COERCIBILITY\x00COLLATE\x00COLLATION\x00COLLATIONPROPERTY\x00COLUMN\x00COLUMNPROPERTY\x00COLUMNS_UPDATED\x00COL_LENGTH\x00COL_NAME\x00COMPRESS\x00CONCAT\x00CONCAT_WS\x00CONDITION\x00CONNECTION_ID\x00CONSTRAINT\x00CONTINUE\x00CONV\x00CONVERT\x00CONVERT_FROM\x00CONVERT_TO\x00CONVERT_TZ\x00COS\x00COT\x00COUNT\x00COUNT_BIG\x00CRC32\x00CREATE\x00CREATE OR\x00CREATE OR REPLACE\x00CROSS\x00CROSS JOIN\x00CSNG\x00CSTRING\x00CTXSYS.DRITHSX.SN\x00CUME_DIST\x00CURDATE\x00CURDIR\x00CURRENT DATE\x00CURRENT DEGREE\x00CURRENT FUNCTION\x00CURRENT FUNCTION PATH\x00CURRENT PATH\x00CURRENT SCHEMA\x00CURRENT SERVER\x00CURRENT TIME\x00CURRENT TIMEZONE\x00CURRENTUSER\x00CURRENT_DATABASE\x00CURRENT_DATE\x00CURRENT_PATH\x00CURRENT_QUERY\x00CURRENT_SCHEMA\x00CURRENT_SCHEMAS\x00CURRENT_SERVER\x00CURRENT_SETTING\x00CURRENT_TIME\x00CURRENT_TIMESTAMP\x00CURRENT_TIMEZONE\x00CURRENT_USER\x00CURRVAL\x00CURSOR\x00CURSOR_STATUS\x00CURTIME\x00CVAR\x00DATABASE\x00DATABASEPROPERTYEX\x00DATABASES\x00DATABASE_PRINCIPAL_ID\x00DATALENGTH\x00DATE\x00DATEADD\x00DATEDIFF\x00DATEFROMPARTS\x00DATENAME\x00DATEPART\x00DATESERIAL\x00DATETIME2FROMPARTS\x00DATETIMEFROMPARTS\x00DATETIMEOFFSETFROMPARTS\x00DATEVALUE\x00DATE_ADD\x00DATE_FORMAT\x00DATE_PART\x00DATE_SUB\x00DATE_TRUNC\x00DAVG\x00DAY\x00DAYNAME\x00DAYOFMONTH\x00DAYOFWEEK\x00DAYOFYEAR\x00DAY_HOUR\x00DAY_MICROSECOND\x00DAY_MINUTE\x00DAY_SECOND\x00DBMS_LOCK.SLEEP\x00DBMS_PIPE.RECEIVE_MESSAGE\x00DBMS_UTILITY.SQLID_TO_SQLHASH\x00DB_ID\x00DB_NAME\x00DCOUNT\x00DEC\x00DECIMAL\x00DECLARE\x00DECODE\x00DECRYPTBYASMKEY\x00DECRYPTBYCERT\x00DECRYPTBYKEY\x00DECRYPTBYKEYAUTOCERT\x00DECRYPTBYPASSPHRASE\x00DEFAULT\x00DEGREES\x00DELAY\x00DELAYED\x00DELETE\x00DENSE_RANK\x00DESC\x00DESCRIBE\x00DES_DECRYPT\x00DES_ENCRYPT\x00DETERMINISTIC\x00DFIRST\x00DIFFERENCE\x00DISTINCT\x00DISTINCTROW\x00DIV\x00DLAST\x00DLOOKUP\x00DMAX\x00DMIN\x00DO\x00DOUBLE\x00DOUBLE PRECISION\x00DROP\x00DSUM\x00DUAL\x00EACH\x00ELSE\x00ELSEIF\x00ELT\x00ENCLOSED\x00ENCODE\x00ENCRYPT\x00ENCRYPTBYASMKEY\x00ENCRYPTBYCERT\x00ENCRYPTBYKEY\x00ENCRYPTBYPASSPHRASE\x00ENUM_FIRST\x00ENUM_LAST\x00ENUM_RANGE\x00EOMONTH\x00EQV\x00ESCAPED\x00EVENTDATA\x00EXCEPT\x00EXEC\x00EXECUTE\x00EXECUTE AS\x00EXECUTE AS LOGIN\x00EXISTS\x00EXIT\x00EXP\x00EXPLAIN\x00EXPORT_SET\x00EXTRACT\x00EXTRACTVALUE\x00EXTRACT_VALUE\x00FALSE\x00FETCH\x00FIELD\x00FILEDATETIME\x00FILEGROUPPROPERTY\x00FILEGROUP_ID\x00FILEGROUP_NAME\x00FILELEN\x00FILEPROPERTY\x00FILETOBLOB\x00FILETOCLOB\x00FILE_ID\x00FILE_IDEX\x00FILE_NAME\x00FIND_IN_SET\x00FIRST_VALUE\x00FLOAT\x00FLOAT4\x00FLOAT8\x00FLOOR\x00FN_VIRTUALFILESTATS\x00FOR\x00FOR UPDATE\x00FOR UPDATE NOWAIT\x00FOR UPDATE OF\x00FOR UPDATE SKIP\x00FOR UPDATE SKIP LOCKED\x00FOR UPDATE WAIT\x00FORCE\x00FOREIGN\x00FORMAT\x00FOUND_ROWS\x00FROM\x00FROM_BASE64\x00FROM_DAYS\x00FROM_UNIXTIME\x00FULL JOIN\x00FULL OUTER\x00FULL OUTER JOIN\x00FULLTEXT\x00FULLTEXTCATALOGPROPERTY\x00FULLTEXTSERVICEPROPERTY\x00FUNCTION\x00GENERATE_SERIES\x00GENERATE_SUBSCRIPTS\x00GETATTR\x00GETDATE\x00GETUTCDATE\x00GET_BIT\x00GET_BYTE\x00GET_FORMAT\x00GET_LOCK\x00GO\x00GOTO\x00GRANT\x00GREATEST\x00GROUP\x00GROUP BY\x00GROUPING\x00GROUPING_ID\x00GROUP_CONCAT\x00HANDLER\x00HASHBYTES\x00HAS_PERMS_BY_NAME\x00HAVING\x00HEX\x00HIGH_PRIORITY\x00HOST_NAME\x00HOUR\x00HOUR_MICROSECOND\x00HOUR_MINUTE\x00HOUR_SECOND\x00IDENTIFY\x00IDENT_CURRENT\x00IDENT_INCR\x00IDENT_SEED\x00IF\x00IF EXISTS\x00IF NOT\x00IF NOT EXISTS\x00IFF\x00IFNULL\x00IGNORE\x00IIF\x00IN\x00IN BOOLEAN\x00IN BOOLEAN MODE\x00INDEX\x00INDEXKEY_PROPERTY\x00INDEXPROPERTY\x00INDEX_COL\x00INET_ATON\x00INET_NTOA\x00INFILE\x00INITCAP\x00INNER\x00INNER JOIN\x00INOUT\x00INSENSITIVE\x00INSERT\x00INSERT DELAYED\x00INSERT DELAYED INTO\x00INSERT HIGH_PRIORITY\x00INSERT HIGH_PRIORITY INTO\x00INSERT IGNORE\x00INSERT IGNORE INTO\x00INSERT INTO\x00INSERT LOW_PRIORITY\x00INSERT LOW_PRIORITY INTO\x00INSTR\x00INSTRREV\x00INT\x00INT1\x00INT2\x00INT3\x00INT4\x00INT8\x00INTEGER\x00INTERSECT\x00INTERSECT ALL\x00INTERVAL\x00INTO\x00INTO DUMPFILE\x00INTO OUTFILE\x00IS\x00IS DISTINCT\x00IS DISTINCT FROM\x00IS NOT\x00IS NOT DISTINCT\x00IS NOT DISTINCT FROM\x00ISDATE\x00ISEMPTY\x00ISFINITE\x00ISNULL\x00ISNUMERIC\x00IS_FREE_LOCK\x00IS_MEMBER\x00IS_OBJECTSIGNED\x00IS_ROLEMEMBER\x00IS_SRVROLEMEMBER\x00IS_USED_LOCK\x00ITERATE\x00JOIN\x00JSON_KEYS\x00JULIANDAY\x00JUSTIFY_DAYS\x00JUSTIFY_HOURS\x00JUSTIFY_INTERVAL\x00KEYS\x00KEY_GUID\x00KEY_ID\x00KILL\x00LAG\x00LASTVAL\x00LAST_INSERT_ID\x00LAST_INSERT_ROWID\x00LAST_VALUE\x00LCASE\x00LEAD\x00LEADING\x00LEAST\x00LEAVE\x00LEFT\x00LEFT JOIN\x00LEFT OUTER\x00LEFT OUTER JOIN\x00LENGTH\x00LIKE\x00LIMIT\x00LINEAR\x00LINES\x00LN\x00LOAD\x00LOAD DATA\x00LOAD XML\x00LOAD_EXTENSION\x00LOAD_FILE\x00LOCALTIME\x00LOCALTIMESTAMP\x00LOCATE\x00LOCK\x00LOCK IN\x00LOCK IN SHARE\x00LOCK IN SHARE MODE\x00LOCK TABLE\x00LOCK TABLES\x00LOG\x00LOG10\x00LOG2\x00LONGBLOB\x00LONGTEXT\x00LOOP\x00LOWER\x00LOWER_INC\x00LOWER_INF\x00LOW_PRIORITY\x00LPAD\x00LTRIM\x00MAKEDATE\x00MAKE_SET\x00MASKLEN\x00MASTER_BIND\x00MASTER_POS_WAIT\x00MASTER_SSL_VERIFY_SERVER_CERT\x00MATCH\x00MAX\x00MAXVALUE\x00MD5\x00MEDIUMBLOB\x00MEDIUMINT\x00MEDIUMTEXT\x00MERGE\x00MICROSECOND\x00MID\x00MIDDLEINT\x00MIN\x00MINUTE\x00MINUTE_MICROSECOND\x00MINUTE_SECOND\x00MKDIR\x00MOD\x00MODE\x00MODIFIES\x00MONEY\x00MONTH\x00MONTHNAME\x00NAME_CONST\x00NATURAL\x00NATURAL FULL\x00NATURAL FULL OUTER JOIN\x00NATURAL INNER\x00NATURAL JOIN\x00NATURAL LEFT\x00NATURAL LEFT OUTER\x00NATURAL LEFT OUTER JOIN\x00NATURAL OUTER\x00NATURAL RIGHT\x00NATURAL RIGHT OUTER JOIN\x00NETMASK\x00NEXT VALUE\x00NEXT VALUE FOR\x00NEXTVAL\x00NOT\x00NOT BETWEEN\x00NOT IN\x00NOT LIKE\x00NOT REGEXP\x00NOT RLIKE\x00NOT SIMILAR\x00NOT SIMILAR TO\x00NOTNULL\x00NOW\x00NOWAIT\x00NO_WRITE_TO_BINLOG\x00NTH_VALUE\x00NTILE\x00NULL\x00NULLIF\x00NUMERIC\x00NZ\x00OBJECTPROPERTY\x00OBJECTPROPERTYEX\x00OBJECT_DEFINITION\x00OBJECT_ID\x00OBJECT_NAME\x00OBJECT_SCHEMA_NAME\x00OCT\x00OCTET_LENGTH\x00OFFSET\x00OID\x00OLD_PASSWORD\x00ONE_SHOT\x00OPEN\x00OPENDATASOURCE\x00OPENQUERY\x00OPENROWSET\x00OPENXML\x00OPTIMIZE\x00OPTION\x00OPTIONALLY\x00OR\x00ORD\x00ORDER\x00ORDER BY\x00ORIGINAL_DB_NAME\x00ORIGINAL_LOGIN\x00OUT\x00OUTER\x00OUTFILE\x00OVERLAPS\x00OVERLAY\x00OWN3D\x00OWN3D BY\x00PARSENAME\x00PARTITION\x00PARTITION BY\x00PASSWORD\x00PATHINDEX\x00PATINDEX\x00PERCENTILE_COUNT\x00PERCENTILE_DISC\x00PERCENTILE_RANK\x00PERCENT_RANK\x00PERIOD_ADD\x00PERIOD_DIFF\x00PERMISSIONS\x00PG_ADVISORY_LOCK\x00PG_BACKEND_PID\x00PG_CANCEL_BACKEND\x00PG_CLIENT_ENCODING\x00PG_CONF_LOAD_TIME\x00PG_CREATE_RESTORE_POINT\x00PG_HAS_ROLE\x00PG_IS_IN_RECOVERY\x00PG_IS_OTHER_TEMP_SCHEMA\x00PG_LISTENING_CHANNELS\x00PG_LS_DIR\x00PG_MY_TEMP_SCHEMA\x00PG_POSTMASTER_START_TIME\x00PG_READ_BINARY_FILE\x00PG_READ_FILE\x00PG_RELOAD_CONF\x00PG_ROTATE_LOGFILE\x00PG_SLEEP\x00PG_START_BACKUP\x00PG_STAT_FILE\x00PG_STOP_BACKUP\x00PG_SWITCH_XLOG\x00PG_TERMINATE_BACKEND\x00PG_TRIGGER_DEPTH\x00PI\x00POSITION\x00POW\x00POWER\x00PRECISION\x00PREVIOUS VALUE\x00PREVIOUS VALUE FOR\x00PRIMARY\x00PRINT\x00PROCEDURE\x00PROCEDURE ANALYSE\x00PUBLISHINGSERVERNAME\x00PURGE\x00PWDCOMPARE\x00PWDENCRYPT\x00QUARTER\x00QUOTE\x00QUOTENAME\x00QUOTE_IDENT\x00QUOTE_LITERAL\x00QUOTE_NULLABLE\x00RADIANS\x00RAISEERROR\x00RAND\x00RANDOM\x00RANDOMBLOB\x00RANGE\x00RANK\x00READ\x00READ WRITE\x00READS\x00READ_WRITE\x00REAL\x00REFERENCES\x00REGCLASS\x00REGCONFIG\x00REGDICTIONARY\x00REGEXP\x00REGEXP_INSTR\x00REGEXP_MATCHES\x00REGEXP_REPLACE\x00REGEXP_SPLIT_TO_ARRAY\x00REGEXP_SPLIT_TO_TABLE\x00REGEXP_SUBSTR\x00REGOPER\x00REGOPERATOR\x00REGPROC\x00REGPROCEDURE\x00REGTYPE\x00RELEASE\x00RELEASE_LOCK\x00RENAME\x00REPEAT\x00REPLACE\x00REPLICATE\x00REQUIRE\x00RESIGNAL\x00RESTRICT\x00RETURN\x00REVERSE\x00REVOKE\x00RIGHT\x00RIGHT JOIN\x00RIGHT OUTER\x00RIGHT OUTER JOIN\x00RLIKE\x00ROUND\x00ROW\x00ROW_COUNT\x00ROW_NUMBER\x00ROW_TO_JSON\x00RPAD\x00RTRIM\x00SCHAMA_NAME\x00SCHEMA\x00SCHEMAS\x00SCHEMA_ID\x00SCOPE_IDENTITY\x00SECOND_MICROSECOND\x00SEC_TO_TIME\x00SELECT\x00SELECT ALL\x00SELECT DISTINCT\x00SENSITIVE\x00SEPARATOR\x00SERIAL\x00SERIAL2\x00SERIAL4\x00SERIAL8\x00SERVERPROPERTY\x00SESSION_USER\x00SET\x00SETATTR\x00SETSEED\x00SETVAL\x00SET_BIT\x00SET_BYTE\x00SET_CONFIG\x00SET_MASKLEN\x00SHA\x00SHA1\x00SHA2\x00SHOW\x00SHUTDOWN\x00SIGN\x00SIGNAL\x00SIGNBYASMKEY\x00SIGNBYCERT\x00SIMILAR\x00SIMILAR TO\x00SIN\x00SLEEP\x00SMALLDATETIMEFROMPARTS\x00SMALLINT\x00SMALLSERIAL\x00SOME\x00SOUNDEX\x00SOUNDS\x00SOUNDS LIKE\x00SPACE\x00SPATIAL\x00SPECIFIC\x00SPLIT_PART\x00SQL\x00SQLEXCEPTION\x00SQLITE_VERSION\x00SQLSTATE\x00SQLWARNING\x00SQL_BIG_RESULT\x00SQL_BUFFER_RESULT\x00SQL_CACHE\x00SQL_CALC_FOUND_ROWS\x00SQL_NO_CACHE\x00SQL_SMALL_RESULT\x00SQL_VARIANT_PROPERTY\x00SQRT\x00SSL\x00STARTING\x00STATEMENT_TIMESTAMP\x00STATS_DATE\x00STDDEV\x00STDDEV_POP\x00STDDEV_SAMP\x00STRAIGHT_JOIN\x00STRCMP\x00STRCOMP\x00STRCONV\x00STRING_AGG\x00STRING_TO_ARRAY\x00STRPOS\x00STR_TO_DATE\x00STUFF\x00SUBDATE\x00SUBSTR\x00SUBSTRING\x00SUBSTRING_INDEX\x00SUBTIME\x00SUM\x00SUSER_ID\x00SUSER_NAME\x00SUSER_SID\x00SUSER_SNAME\x00SWITCHOFFET\x00SYS.DATABASE_NAME\x00SYS.FN_BUILTIN_PERMISSIONS\x00SYS.FN_GET_AUDIT_FILE\x00SYS.FN_MY_PERMISSIONS\x00SYS.STRAGG\x00SYSCOLUMNS\x00SYSDATE\x00SYSDATETIME\x00SYSDATETIMEOFFSET\x00SYSOBJECTS\x00SYSTEM_USER\x00SYSUSERS\x00SYSUTCDATETME\x00TABLE\x00TAN\x00TERMINATED\x00TERTIARY_WEIGHTS\x00TEXT\x00TEXTPOS\x00TEXTPTR\x00TEXTVALID\x00THEN\x00TIME\x00TIMEDIFF\x00TIMEFROMPARTS\x00TIMEOFDAY\x00TIMESERIAL\x00TIMESTAMP\x00TIMESTAMPADD\x00TIMEVALUE\x00TIME_FORMAT\x00TIME_TO_SEC\x00TINYBLOB\x00TINYINT\x00TINYTEXT\x00TODATETIMEOFFSET\x00TOP\x00TOTAL\x00TOTAL_CHANGES\x00TO_ASCII\x00TO_BASE64\x00TO_CHAR\x00TO_DATE\x00TO_DAYS\x00TO_HEX\x00TO_NUMBER\x00TO_SECONDS\x00TO_TIMESTAMP\x00TRAILING\x00TRANSACTION_TIMESTAMP\x00TRANSLATE\x00TRIGGER\x00TRIGGER_NESTLEVEL\x00TRIM\x00TRUE\x00TRUNC\x00TRUNCATE\x00TRY\x00TRY_CAST\x00TRY_CONVERT\x00TRY_PARSE\x00TYPEOF\x00TYPEPROPERTY\x00TYPE_ID\x00TYPE_NAME\x00UCASE\x00UESCAPE\x00UNCOMPRESS\x00UNCOMPRESS_LENGTH\x00UNDO\x00UNHEX\x00UNICODE\x00UNION\x00UNION ALL\x00UNION ALL DISTINCT\x00UNION DISTINCT\x00UNION DISTINCT ALL\x00UNIQUE\x00UNIX_TIMESTAMP\x00UNI_ON\x00UNKNOWN\x00UNLOCK\x00UNNEST\x00UNSIGNED\x00UPDATE\x00UPDATEXML\x00UPPER\x00UPPER_INC\x00UPPER_INF\x00USAGE\x00USE\x00USER\x00USER_ID\x00USER_LOCK.SLEEP\x00USER_NAME\x00USING\x00UTC_DATE\x00UTC_TIME\x00UTC_TIMESTAMP\x00UTL_HTTP.REQUEST\x00UTL_INADDR.GET_HOST_ADDRESS\x00UTL_INADDR.GET_HOST_NAME\x00UUID\x00UUID_SHORT\x00VALUES\x00VAR\x00VARBINARY\x00VARCHAR\x00VARCHARACTER\x00VARIANCE\x00VARP\x00VARYING\x00VAR_POP\x00VAR_SAMP\x00VERIFYSIGNEDBYASMKEY\x00VERIFYSIGNEDBYCERT\x00VERSION\x00VOID\x00WAIT\x00WAITFOR\x00WAITFOR DELAY\x00WAITFOR RECEIVE\x00WAITFOR TIME\x00WEEK\x00WEEKDAY\x00WEEKDAYNAME\x00WEEKOFYEAR\x00WHEN\x00WHERE\x00WHILE\x00WIDTH_BUCKET\x00WITH\x00WITH ROLLUP\x00XMLAGG\x00XMLCOMMENT\x00XMLCONCAT\x00XMLELEMENT\x00XMLEXISTS\x00XMLFOREST\x00XMLFORMAT\x00XMLPI\x00XMLROOT\x00XMLTYPE\x00XML_IS_WELL_FORMED\x00XOR\x00XPATH\x00XPATH_EXISTS\x00XP_EXECRESULTSET\x00YEAR\x00YEARWEEK\x00YEAR_MONTH\x00ZEROBLOB\x00ZEROFILL\x00^=\x00_ARMSCII8\x00_ASCII\x00_BIG5\x00_BINARY\x00_CP1250\x00_CP1251\x00_CP1257\x00_CP850\x00_CP852\x00_CP866\x00_CP932\x00_DEC8\x00_EUCJPMS\x00_EUCKR\x00_GB2312\x00_GBK\x00_GEOSTD8\x00_GREEK\x00_HEBREW\x00_HP8\x00_KEYBCS2\x00_KOI8R\x00_KOI8U\x00_LATIN1\x00_LATIN2\x00_LATIN5\x00_LATIN7\x00_MACCE\x00_MACROMAN\x00_SJIS\x00_SWE7\x00_TIS620\x00_UJIS\x00_USC2\x00_UTF8\x00|/\x00|=\x00||\x00~*\x00haystack\x00spec/dogfood/cycles/20260810f/libinjection_sqli/src.c\x00needle\x00nlen > 1\x00 \t\n\v\f\r\xa0\x00\x0001\x000123456789ABCDEFabcdef\x00 []{}<>:\\?=@!#~+-*/&|^%(),';\t\n\v\f\r\"\xa0\x00\x00 <>:\\?=@!#~+-*/&|^%(),';\t\n\v\f\r'`\"\x000123456789.,\x00abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\x003.9.2\x00pos >= 3\x00sp_password\x00sos\x00s&s\x00s&n\x00n&1\x001&1\x001&v\x001&s\x00"
