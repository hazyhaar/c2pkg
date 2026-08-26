// Code generated for linux/amd64 by 'ccgo --package-name=df_tiny_regex -o spec/dogfood/cycles/20260810g/tiny_regex/raw.go -I spec/dogfood/cycles/20260810g/tiny_regex spec/dogfood/cycles/20260810g/tiny_regex/src.c', DO NOT EDIT.

//go:build linux && amd64

package df_tiny_regex

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
const FILENAME_MAX = 4096
const FOPEN_MAX = 1000
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MAX_CHAR_CLASS_LEN = 40
const MAX_REGEXP_OBJECTS = 30
const P_tmpdir = "/tmp"
const RE_DOT_MATCHES_NEWLINE = 1
const TMP_MAX = 10000
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
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type re_t = uintptr

type regex_t = struct {
	Ftype1 uint8
	Fu     struct {
		Fccl         [0]uintptr
		Fch          uint8
		F__ccgo_pad2 [7]byte
	}
}

type size_t = uint64

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

type locale_t = uintptr

const UNUSED = 0
const DOT = 1
const BEGIN = 2
const END = 3
const QUESTIONMARK = 4
const STAR = 5
const PLUS = 6
const CHAR = 7
const CHAR_CLASS = 8
const INV_CHAR_CLASS = 9
const DIGIT = 10
const NOT_DIGIT = 11
const ALPHA = 12
const NOT_ALPHA = 13
const WHITESPACE = 14
const NOT_WHITESPACE = 15

// C documentation
//
//	/* Public functions: */
func re_match(tls *libc.TLS, pattern uintptr, text uintptr, matchlength uintptr) (r int32) {
	return re_matchp(tls, re_compile(tls, pattern), text, matchlength)
}

func re_matchp(tls *libc.TLS, pattern re_t, text uintptr, matchlength uintptr) (r int32) {
	var idx, v1 int32
	var v2 uintptr
	_, _, _ = idx, v1, v2
	**(**int32)(__ccgo_up(matchlength)) = 0
	if pattern != uintptr(0) {
		if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern))).Ftype1) == int32(BEGIN) {
			if matchpattern(tls, pattern+1*16, text, matchlength) != 0 {
				v1 = 0
			} else {
				v1 = -int32(1)
			}
			return v1
		} else {
			idx = -int32(1)
			for {
				idx = idx + int32(1)
				if matchpattern(tls, pattern, text, matchlength) != 0 {
					if int32(**(**int8)(__ccgo_up(text))) == int32('\000') {
						return -int32(1)
					}
					return idx
				}
				goto _3
			_3:
				;
				v2 = text
				text = text + 1
				if !(int32(**(**int8)(__ccgo_up(v2))) != int32('\000')) {
					break
				}
			}
		}
	}
	return -int32(1)
}

func re_compile(tls *libc.TLS, pattern uintptr) (r re_t) {
	var buf_begin, ccl_bufidx, i, j, v1, v2 int32
	var c int8
	_, _, _, _, _, _, _ = buf_begin, c, ccl_bufidx, i, j, v1, v2
	ccl_bufidx = int32(1) /* current char in pattern   */
	i = 0                 /* index into pattern        */
	j = 0                 /* index into re_compiled    */
	for int32(**(**int8)(__ccgo_up(pattern + uintptr(i)))) != int32('\000') && j+int32(1) < int32(MAX_REGEXP_OBJECTS) {
		c = **(**int8)(__ccgo_up(pattern + uintptr(i)))
		switch int32(c) {
		/* Meta-characters: */
		case int32('^'):
			re_compiled[j].Ftype1 = uint8(BEGIN)
		case int32('$'):
			re_compiled[j].Ftype1 = uint8(END)
		case int32('.'):
			re_compiled[j].Ftype1 = uint8(DOT)
		case int32('*'):
			re_compiled[j].Ftype1 = uint8(STAR)
		case int32('+'):
			re_compiled[j].Ftype1 = uint8(PLUS)
		case int32('?'):
			re_compiled[j].Ftype1 = uint8(QUESTIONMARK)
			break
			/*    case '|': {    re_compiled[j].type = BRANCH;          } break; <-- not working properly */
			/* Escaped character-classes (\s \w ...): */
			fallthrough
		case int32('\\'):
			if int32(**(**int8)(__ccgo_up(pattern + uintptr(i+int32(1))))) != int32('\000') {
				/* Skip the escape-char '\\' */
				i = i + int32(1)
				/* ... and check the next */
				switch int32(**(**int8)(__ccgo_up(pattern + uintptr(i)))) {
				/* Meta-character: */
				case int32('d'):
					re_compiled[j].Ftype1 = uint8(DIGIT)
				case int32('D'):
					re_compiled[j].Ftype1 = uint8(NOT_DIGIT)
				case int32('w'):
					re_compiled[j].Ftype1 = uint8(ALPHA)
				case int32('W'):
					re_compiled[j].Ftype1 = uint8(NOT_ALPHA)
				case int32('s'):
					re_compiled[j].Ftype1 = uint8(WHITESPACE)
				case int32('S'):
					re_compiled[j].Ftype1 = uint8(NOT_WHITESPACE)
					break
					/* Escaped character, e.g. '.' or '$' */
					fallthrough
				default:
					re_compiled[j].Ftype1 = uint8(CHAR)
					*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&re_compiled)) + uintptr(j)*16 + 8)) = libc.Uint8FromInt8(**(**int8)(__ccgo_up(pattern + uintptr(i))))
					break
				}
			}
			/* '\\' as last char in pattern -> invalid regular expression. */
			/*
			   else
			   {
			     re_compiled[j].type = CHAR;
			     re_compiled[j].ch = pattern[i];
			   }
			*/
			break
			/* Character class: */
			fallthrough
		case int32('['):
			/* Remember where the char-buffer starts. */
			buf_begin = ccl_bufidx
			/* Look-ahead to determine if negated */
			if int32(**(**int8)(__ccgo_up(pattern + uintptr(i+int32(1))))) == int32('^') {
				re_compiled[j].Ftype1 = uint8(INV_CHAR_CLASS)
				i = i + int32(1)                                                      /* Increment i to avoid including '^' in the char-buffer */
				if int32(**(**int8)(__ccgo_up(pattern + uintptr(i+int32(1))))) == 0 { /* incomplete pattern, missing non-zero char after '^' */
					return uintptr(0)
				}
			} else {
				re_compiled[j].Ftype1 = uint8(CHAR_CLASS)
			}
			/* Copy characters inside [..] to buffer */
			for {
				i = i + 1
				v1 = i
				if !(int32(**(**int8)(__ccgo_up(pattern + uintptr(v1)))) != int32(']') && int32(**(**int8)(__ccgo_up(pattern + uintptr(i)))) != int32('\000')) {
					break
				} /* Missing ] */
				if int32(**(**int8)(__ccgo_up(pattern + uintptr(i)))) == int32('\\') {
					if ccl_bufidx >= libc.Int32FromInt32(MAX_CHAR_CLASS_LEN)-libc.Int32FromInt32(1) {
						//fputs("exceeded internal buffer!\n", stderr);
						return uintptr(0)
					}
					if int32(**(**int8)(__ccgo_up(pattern + uintptr(i+int32(1))))) == 0 { /* incomplete pattern, missing non-zero char after '\\' */
						return uintptr(0)
					}
					v1 = ccl_bufidx
					ccl_bufidx = ccl_bufidx + 1
					v2 = i
					i = i + 1
					ccl_buf[v1] = libc.Uint8FromInt8(**(**int8)(__ccgo_up(pattern + uintptr(v2))))
				} else {
					if ccl_bufidx >= int32(MAX_CHAR_CLASS_LEN) {
						//fputs("exceeded internal buffer!\n", stderr);
						return uintptr(0)
					}
				}
				v1 = ccl_bufidx
				ccl_bufidx = ccl_bufidx + 1
				ccl_buf[v1] = libc.Uint8FromInt8(**(**int8)(__ccgo_up(pattern + uintptr(i))))
			}
			if ccl_bufidx >= int32(MAX_CHAR_CLASS_LEN) {
				/* Catches cases such as [00000000000000000000000000000000000000][ */
				//fputs("exceeded internal buffer!\n", stderr);
				return uintptr(0)
			}
			/* Null-terminate string end */
			v1 = ccl_bufidx
			ccl_bufidx = ccl_bufidx + 1
			ccl_buf[v1] = uint8(0)
			*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&re_compiled)) + uintptr(j)*16 + 8)) = uintptr(unsafe.Pointer(&ccl_buf)) + uintptr(buf_begin)
			break
			/* Other characters: */
			fallthrough
		default:
			re_compiled[j].Ftype1 = uint8(CHAR)
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&re_compiled)) + uintptr(j)*16 + 8)) = libc.Uint8FromInt8(c)
			break
		}
		/* no buffer-out-of-bounds access on invalid patterns - see https://github.com/kokke/tiny-regex-c/commit/1a279e04014b70b0695fba559a7c05d55e6ee90b */
		if int32(**(**int8)(__ccgo_up(pattern + uintptr(i)))) == 0 {
			return uintptr(0)
		}
		i = i + int32(1)
		j = j + int32(1)
	}
	/* 'UNUSED' is a sentinel used to indicate end-of-pattern */
	re_compiled[j].Ftype1 = uint8(UNUSED)
	return uintptr(unsafe.Pointer(&re_compiled))
}

/*
The sizes of the two static arrays below substantiates the static RAM usage of this module.

	MAX_REGEXP_OBJECTS is the max number of symbols in the expression.
	MAX_CHAR_CLASS_LEN determines the size of buffer for chars in all char-classes in the expression.
*/
var re_compiled [30]regex_t

var ccl_buf [40]uint8

func re_print(tls *libc.TLS, pattern uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var c int8
	var i, j int32
	var types [17]uintptr
	_, _, _, _ = c, i, j, types
	types = [17]uintptr{
		0:  __ccgo_ts,
		1:  __ccgo_ts + 7,
		2:  __ccgo_ts + 11,
		3:  __ccgo_ts + 17,
		4:  __ccgo_ts + 21,
		5:  __ccgo_ts + 34,
		6:  __ccgo_ts + 39,
		7:  __ccgo_ts + 44,
		8:  __ccgo_ts + 49,
		9:  __ccgo_ts + 60,
		10: __ccgo_ts + 75,
		11: __ccgo_ts + 81,
		12: __ccgo_ts + 91,
		13: __ccgo_ts + 97,
		14: __ccgo_ts + 107,
		15: __ccgo_ts + 118,
		16: __ccgo_ts + 133,
	}
	i = 0
	for {
		if !(i < int32(MAX_REGEXP_OBJECTS)) {
			break
		}
		if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + uintptr(i)*16))).Ftype1) == int32(UNUSED) {
			break
		}
		libc.Xprintf(tls, __ccgo_ts+140, libc.VaList(bp+8, types[(**(**regex_t)(__ccgo_up(pattern + uintptr(i)*16))).Ftype1]))
		if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + uintptr(i)*16))).Ftype1) == int32(CHAR_CLASS) || libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + uintptr(i)*16))).Ftype1) == int32(INV_CHAR_CLASS) {
			libc.Xprintf(tls, __ccgo_ts+149, 0)
			j = 0
			for {
				if !(j < int32(MAX_CHAR_CLASS_LEN)) {
					break
				}
				c = libc.Int8FromUint8(**(**uint8)(__ccgo_up(*(*uintptr)(unsafe.Pointer(pattern + uintptr(i)*16 + 8)) + uintptr(j))))
				if int32(c) == int32('\000') || int32(c) == int32(']') {
					break
				}
				libc.Xprintf(tls, __ccgo_ts+152, libc.VaList(bp+8, int32(c)))
				goto _2
			_2:
				;
				j = j + 1
			}
			libc.Xprintf(tls, __ccgo_ts+155, 0)
		} else {
			if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + uintptr(i)*16))).Ftype1) == int32(CHAR) {
				libc.Xprintf(tls, __ccgo_ts+157, libc.VaList(bp+8, libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(pattern + uintptr(i)*16 + 8)))))
			}
		}
		libc.Xprintf(tls, __ccgo_ts+163, 0)
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	/* Private functions: */
func matchdigit(tls *libc.TLS, c int8) (r int32) {
	return libc.BoolInt32(libc.Uint32FromInt8(c)-uint32('0') < uint32(10))
}

func matchalpha(tls *libc.TLS, c int8) (r int32) {
	return libc.BoolInt32(libc.Uint32FromInt8(c)|uint32(32)-uint32('a') < uint32(26))
}

func matchwhitespace(tls *libc.TLS, c int8) (r int32) {
	var v1, v2 int32
	_, _ = v1, v2
	v1 = int32(c)
	v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
	goto _3
_3:
	return v2
}

func matchalphanum(tls *libc.TLS, c int8) (r int32) {
	return libc.BoolInt32(int32(c) == int32('_') || matchalpha(tls, c) != 0 || matchdigit(tls, c) != 0)
}

func matchrange(tls *libc.TLS, c int8, str uintptr) (r int32) {
	return libc.BoolInt32(int32(c) != int32('-') && int32(**(**int8)(__ccgo_up(str))) != int32('\000') && int32(**(**int8)(__ccgo_up(str))) != int32('-') && int32(**(**int8)(__ccgo_up(str + 1))) == int32('-') && int32(**(**int8)(__ccgo_up(str + 2))) != int32('\000') && (int32(c) >= int32(**(**int8)(__ccgo_up(str))) && int32(c) <= int32(**(**int8)(__ccgo_up(str + 2)))))
}

func matchdot(tls *libc.TLS, c int8) (r int32) {
	_ = c
	return int32(1)
}

func ismetachar(tls *libc.TLS, c int8) (r int32) {
	return libc.BoolInt32(int32(c) == int32('s') || int32(c) == int32('S') || int32(c) == int32('w') || int32(c) == int32('W') || int32(c) == int32('d') || int32(c) == int32('D'))
}

func matchmetachar(tls *libc.TLS, c int8, str uintptr) (r int32) {
	switch int32(**(**int8)(__ccgo_up(str))) {
	case int32('d'):
		return matchdigit(tls, c)
	case int32('D'):
		return libc.BoolInt32(!(matchdigit(tls, c) != 0))
	case int32('w'):
		return matchalphanum(tls, c)
	case int32('W'):
		return libc.BoolInt32(!(matchalphanum(tls, c) != 0))
	case int32('s'):
		return matchwhitespace(tls, c)
	case int32('S'):
		return libc.BoolInt32(!(matchwhitespace(tls, c) != 0))
	default:
		return libc.BoolInt32(int32(c) == int32(**(**int8)(__ccgo_up(str))))
	}
	return r
}

func matchcharclass(tls *libc.TLS, c int8, str uintptr) (r int32) {
	var v1 uintptr
	_ = v1
	for {
		if matchrange(tls, c, str) != 0 {
			return int32(1)
		} else {
			if int32(**(**int8)(__ccgo_up(str))) == int32('\\') {
				/* Escape-char: increment str-ptr and match on next char */
				str = str + uintptr(1)
				if matchmetachar(tls, c, str) != 0 {
					return int32(1)
				} else {
					if int32(c) == int32(**(**int8)(__ccgo_up(str))) && !(ismetachar(tls, c) != 0) {
						return int32(1)
					}
				}
			} else {
				if int32(c) == int32(**(**int8)(__ccgo_up(str))) {
					if int32(c) == int32('-') {
						return libc.BoolInt32(int32(**(**int8)(__ccgo_up(str + uintptr(-libc.Int32FromInt32(1))))) == int32('\000') || int32(**(**int8)(__ccgo_up(str + 1))) == int32('\000'))
					} else {
						return int32(1)
					}
				}
			}
		}
		goto _2
	_2:
		;
		v1 = str
		str = str + 1
		if !(int32(**(**int8)(__ccgo_up(v1))) != int32('\000')) {
			break
		}
	}
	return 0
}

func matchone(tls *libc.TLS, _p regex_t, c int8) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*regex_t)(unsafe.Pointer(bp)) = _p
	switch libc.Int32FromUint8((**(**regex_t)(__ccgo_up(bp))).Ftype1) {
	case int32(DOT):
		return matchdot(tls, c)
	case int32(CHAR_CLASS):
		return matchcharclass(tls, c, *(*uintptr)(unsafe.Pointer(bp + 8)))
	case int32(INV_CHAR_CLASS):
		return libc.BoolInt32(!(matchcharclass(tls, c, *(*uintptr)(unsafe.Pointer(bp + 8))) != 0))
	case int32(DIGIT):
		return matchdigit(tls, c)
	case int32(NOT_DIGIT):
		return libc.BoolInt32(!(matchdigit(tls, c) != 0))
	case int32(ALPHA):
		return matchalphanum(tls, c)
	case int32(NOT_ALPHA):
		return libc.BoolInt32(!(matchalphanum(tls, c) != 0))
	case int32(WHITESPACE):
		return matchwhitespace(tls, c)
	case int32(NOT_WHITESPACE):
		return libc.BoolInt32(!(matchwhitespace(tls, c) != 0))
	default:
		return libc.BoolInt32(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(bp + 8))) == int32(c))
	}
	return r
}

func matchstar(tls *libc.TLS, p regex_t, pattern uintptr, text uintptr, matchlength uintptr) (r int32) {
	var prelen int32
	var prepoint, v1 uintptr
	_, _, _ = prelen, prepoint, v1
	prelen = **(**int32)(__ccgo_up(matchlength))
	prepoint = text
	for int32(**(**int8)(__ccgo_up(text))) != int32('\000') && matchone(tls, p, **(**int8)(__ccgo_up(text))) != 0 {
		text = text + 1
		**(**int32)(__ccgo_up(matchlength)) = **(**int32)(__ccgo_up(matchlength)) + 1
	}
	for text >= prepoint {
		v1 = text
		text = text - 1
		if matchpattern(tls, pattern, v1, matchlength) != 0 {
			return int32(1)
		}
		**(**int32)(__ccgo_up(matchlength)) = **(**int32)(__ccgo_up(matchlength)) - 1
	}
	**(**int32)(__ccgo_up(matchlength)) = prelen
	return 0
}

func matchplus(tls *libc.TLS, p regex_t, pattern uintptr, text uintptr, matchlength uintptr) (r int32) {
	var prepoint, v1 uintptr
	_, _ = prepoint, v1
	prepoint = text
	for int32(**(**int8)(__ccgo_up(text))) != int32('\000') && matchone(tls, p, **(**int8)(__ccgo_up(text))) != 0 {
		text = text + 1
		**(**int32)(__ccgo_up(matchlength)) = **(**int32)(__ccgo_up(matchlength)) + 1
	}
	for text > prepoint {
		v1 = text
		text = text - 1
		if matchpattern(tls, pattern, v1, matchlength) != 0 {
			return int32(1)
		}
		**(**int32)(__ccgo_up(matchlength)) = **(**int32)(__ccgo_up(matchlength)) - 1
	}
	return 0
}

func matchquestion(tls *libc.TLS, p regex_t, pattern uintptr, text uintptr, matchlength uintptr) (r int32) {
	var v1 uintptr
	var v2 bool
	_, _ = v1, v2
	if libc.Int32FromUint8(p.Ftype1) == int32(UNUSED) {
		return int32(1)
	}
	if matchpattern(tls, pattern, text, matchlength) != 0 {
		return int32(1)
	}
	if v2 = **(**int8)(__ccgo_up(text)) != 0; v2 {
		v1 = text
		text = text + 1
	}
	if v2 && matchone(tls, p, **(**int8)(__ccgo_up(v1))) != 0 {
		if matchpattern(tls, pattern, text, matchlength) != 0 {
			**(**int32)(__ccgo_up(matchlength)) = **(**int32)(__ccgo_up(matchlength)) + 1
			return int32(1)
		}
	}
	return 0
}

// C documentation
//
//	/* Iterative matching */
func matchpattern(tls *libc.TLS, pattern uintptr, text uintptr, matchlength uintptr) (r int32) {
	var pre int32
	var v1, v2 uintptr
	var v3 bool
	_, _, _, _ = pre, v1, v2, v3
	pre = **(**int32)(__ccgo_up(matchlength))
	for {
		if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern))).Ftype1) == int32(UNUSED) || libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + 1*16))).Ftype1) == int32(QUESTIONMARK) {
			return matchquestion(tls, **(**regex_t)(__ccgo_up(pattern)), pattern+2*16, text, matchlength)
		} else {
			if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + 1*16))).Ftype1) == int32(STAR) {
				return matchstar(tls, **(**regex_t)(__ccgo_up(pattern)), pattern+2*16, text, matchlength)
			} else {
				if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + 1*16))).Ftype1) == int32(PLUS) {
					return matchplus(tls, **(**regex_t)(__ccgo_up(pattern)), pattern+2*16, text, matchlength)
				} else {
					if libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern))).Ftype1) == int32(END) && libc.Int32FromUint8((**(**regex_t)(__ccgo_up(pattern + 1*16))).Ftype1) == int32(UNUSED) {
						return libc.BoolInt32(int32(**(**int8)(__ccgo_up(text))) == int32('\000'))
					}
				}
			}
		}
		/*  Branching is not working properly
		    else if (pattern[1].type == BRANCH)
		    {
		      return (matchpattern(pattern, text) || matchpattern(&pattern[2], text));
		    }
		*/
		**(**int32)(__ccgo_up(matchlength)) = **(**int32)(__ccgo_up(matchlength)) + 1
		goto _4
	_4:
		;
		if v3 = int32(**(**int8)(__ccgo_up(text))) != int32('\000'); v3 {
			v1 = pattern
			pattern += 16
			v2 = text
			text = text + 1
		}
		if !(v3 && matchone(tls, **(**regex_t)(__ccgo_up(v1)), **(**int8)(__ccgo_up(v2))) != 0) {
			break
		}
	}
	**(**int32)(__ccgo_up(matchlength)) = pre
	return 0
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "UNUSED\x00DOT\x00BEGIN\x00END\x00QUESTIONMARK\x00STAR\x00PLUS\x00CHAR\x00CHAR_CLASS\x00INV_CHAR_CLASS\x00DIGIT\x00NOT_DIGIT\x00ALPHA\x00NOT_ALPHA\x00WHITESPACE\x00NOT_WHITESPACE\x00BRANCH\x00type: %s\x00 [\x00%c\x00]\x00 '%c'\x00\n\x00"
