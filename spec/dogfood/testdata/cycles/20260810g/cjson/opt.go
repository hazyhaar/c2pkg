// Code generated for linux/amd64 by 'ccgo --package-name=df_cjson -o spec/dogfood/cycles/20260810g/cjson/raw.go -I spec/dogfood/cycles/20260810g/cjson spec/dogfood/cycles/20260810g/cjson/src.c', DO NOT EDIT.

//go:build linux && amd64

package df_cjson

import (
	"math"
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ = math.Pi
var _ reflect.Type
var _ unsafe.Pointer

const ARG_MAX = 131072
const BC_BASE_MAX = 99
const BC_DIM_MAX = 2048
const BC_SCALE_MAX = 99
const BC_STRING_MAX = 1000
const BUFSIZ = 1024
const CHARCLASS_NAME_MAX = 14
const CHAR_BIT = 8
const CHAR_MAX = 255
const CHAR_MIN = 0
const CJSON_NESTING_LIMIT = 1000
const CJSON_VERSION_MAJOR = 1
const CJSON_VERSION_MINOR = 7
const CJSON_VERSION_PATCH = 18
const COLL_WEIGHTS_MAX = 2
const DBL_DECIMAL_DIG = 17
const DBL_DIG = 15
const DBL_EPSILON = 2.220446049250313e-16
const DBL_HAS_SUBNORM = 1
const DBL_MANT_DIG = 53
const DBL_MAX = 1.79769313486231570815e+308
const DBL_MAX_10_EXP = 308
const DBL_MAX_EXP = 1024
const DBL_MIN = 2.22507385850720138309e-308
const DBL_TRUE_MIN = 4.94065645841246544177e-324
const DECIMAL_DIG = 17
const DELAYTIMER_MAX = 0x7fffffff
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXPR_NEST_MAX = 32
const FILENAME_MAX = 4096
const FILESIZEBITS = 64
const FLT_DECIMAL_DIG = 9
const FLT_DIG = 6
const FLT_EPSILON = 1.1920928955078125e-07
const FLT_EVAL_METHOD = 0
const FLT_HAS_SUBNORM = 1
const FLT_MANT_DIG = 24
const FLT_MAX = 3.40282346638528859812e+38
const FLT_MAX_10_EXP = 38
const FLT_MAX_EXP = 128
const FLT_MIN = 1.17549435082228750797e-38
const FLT_RADIX = 2
const FLT_TRUE_MIN = 1.40129846432481707092e-45
const FOPEN_MAX = 1000
const FP_ILOGB0 = "FP_ILOGBNAN"
const FP_INFINITE = 1
const FP_NAN = 0
const FP_NORMAL = 4
const FP_SUBNORMAL = 3
const FP_ZERO = 2
const HOST_NAME_MAX = 255
const HUGE = 3.40282346638528859812e+38
const HUGE_VALF = "INFINITY"
const INT_MAX = 2147483647
const IOV_MAX = 1024
const LDBL_DECIMAL_DIG = "DECIMAL_DIG"
const LDBL_DIG = 15
const LDBL_EPSILON = 2.22044604925031308085e-16
const LDBL_HAS_SUBNORM = 1
const LDBL_MANT_DIG = 53
const LDBL_MAX = 1.79769313486231570815e+308
const LDBL_MAX_10_EXP = 308
const LDBL_MAX_EXP = 1024
const LDBL_MIN = 2.22507385850720138309e-308
const LDBL_TRUE_MIN = 4.94065645841246544177e-324
const LINE_MAX = 4096
const LLONG_MAX = 0x7fffffffffffffff
const LOGIN_NAME_MAX = 256
const LONG_BIT = 64
const LONG_MAX = "__LONG_MAX"
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MATH_ERREXCEPT = 2
const MATH_ERRNO = 1
const MB_LEN_MAX = 4
const MQ_PRIO_MAX = 32768
const M_1_PI = 0.31830988618379067154
const M_2_PI = 0.63661977236758134308
const M_2_SQRTPI = 1.12837916709551257390
const M_E = 2.7182818284590452354
const M_LN10 = 2.30258509299404568402
const M_LN2 = 0.69314718055994530942
const M_LOG10E = 0.43429448190325182765
const M_LOG2E = 1.4426950408889634074
const M_PI = 3.14159265358979323846
const M_PI_2 = 1.57079632679489661923
const M_PI_4 = 0.78539816339744830962
const M_SQRT1_2 = 0.70710678118654752440
const M_SQRT2 = 1.41421356237309504880
const NAME_MAX = 255
const NGROUPS_MAX = 32
const NL_ARGMAX = 9
const NL_LANGMAX = 32
const NL_MSGMAX = 32767
const NL_NMAX = 16
const NL_SETMAX = 255
const NL_TEXTMAX = 2048
const NZERO = 20
const PAGESIZE = 4096
const PAGE_SIZE = "PAGESIZE"
const PATH_MAX = 4096
const PIPE_BUF = 4096
const PTHREAD_DESTRUCTOR_ITERATIONS = 4
const PTHREAD_KEYS_MAX = 128
const PTHREAD_STACK_MIN = 2048
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const RE_DUP_MAX = 255
const SCHAR_MAX = 127
const SEM_NSEMS_MAX = 256
const SEM_VALUE_MAX = 0x7fffffff
const SHRT_MAX = 0x7fff
const SSIZE_MAX = "LONG_MAX"
const SYMLOOP_MAX = 40
const TMP_MAX = 10000
const TTY_NAME_MAX = 32
const TZNAME_MAX = 6
const UCHAR_MAX = 255
const UINT_MAX = 0xffffffff
const USHRT_MAX = 0xffff
const WNOHANG = 1
const WORD_BIT = 32
const WUNTRACED = 2
const _GNU_SOURCE = 1
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const _LP64 = 1
const _POSIX2_BC_BASE_MAX = 99
const _POSIX2_BC_DIM_MAX = 2048
const _POSIX2_BC_SCALE_MAX = 99
const _POSIX2_BC_STRING_MAX = 1000
const _POSIX2_CHARCLASS_NAME_MAX = 14
const _POSIX2_COLL_WEIGHTS_MAX = 2
const _POSIX2_EXPR_NEST_MAX = 32
const _POSIX2_LINE_MAX = 2048
const _POSIX2_RE_DUP_MAX = 255
const _POSIX_AIO_LISTIO_MAX = 2
const _POSIX_AIO_MAX = 1
const _POSIX_ARG_MAX = 4096
const _POSIX_CHILD_MAX = 25
const _POSIX_CLOCKRES_MIN = 20000000
const _POSIX_DELAYTIMER_MAX = 32
const _POSIX_HOST_NAME_MAX = 255
const _POSIX_LINK_MAX = 8
const _POSIX_LOGIN_NAME_MAX = 9
const _POSIX_MAX_CANON = 255
const _POSIX_MAX_INPUT = 255
const _POSIX_MQ_OPEN_MAX = 8
const _POSIX_MQ_PRIO_MAX = 32
const _POSIX_NAME_MAX = 14
const _POSIX_NGROUPS_MAX = 8
const _POSIX_OPEN_MAX = 20
const _POSIX_PATH_MAX = 256
const _POSIX_PIPE_BUF = 512
const _POSIX_RE_DUP_MAX = 255
const _POSIX_RTSIG_MAX = 8
const _POSIX_SEM_NSEMS_MAX = 256
const _POSIX_SEM_VALUE_MAX = 32767
const _POSIX_SIGQUEUE_MAX = 32
const _POSIX_SSIZE_MAX = 32767
const _POSIX_SS_REPL_MAX = 4
const _POSIX_STREAM_MAX = 8
const _POSIX_SYMLINK_MAX = 255
const _POSIX_SYMLOOP_MAX = 8
const _POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const _POSIX_THREAD_KEYS_MAX = 128
const _POSIX_THREAD_THREADS_MAX = 64
const _POSIX_TIMER_MAX = 32
const _POSIX_TRACE_EVENT_NAME_MAX = 30
const _POSIX_TRACE_NAME_MAX = 8
const _POSIX_TRACE_SYS_MAX = 8
const _POSIX_TRACE_USER_EVENT_MAX = 32
const _POSIX_TTY_NAME_MAX = 9
const _POSIX_TZNAME_MAX = 6
const _STDC_PREDEF_H = 1
const _XOPEN_IOV_MAX = 16
const _XOPEN_NAME_MAX = 255
const _XOPEN_PATH_MAX = 1024
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
const cJSON_Invalid = 0
const cJSON_IsReference = 256
const cJSON_StringIsConst = 512
const internal_free = "free"
const internal_malloc = "malloc"
const internal_realloc = "realloc"
const linux = 1
const math_errhandling = 2
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type size_t = uint64

type locale_t = uintptr

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

type float_t = float32

type double_t = float64

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

type max_align_t = struct {
	F__ll int64
	F__ld float64
}

type ptrdiff_t = int64

type cJSON = struct {
	Fnext        uintptr
	Fprev        uintptr
	Fchild       uintptr
	Ftype1       int32
	Fvaluestring uintptr
	Fvalueint    int32
	Fvaluedouble float64
	Fstring1     uintptr
}

type cJSON_Hooks = struct {
	Fmalloc_fn uintptr
	Ffree_fn   uintptr
}

type cJSON_bool = int32

/* define our own boolean type */

/* define isnan and isinf for ANSI C, if in C99 or above, isnan and isinf has been defined in math.h */

type error1 = struct {
	Fjson     uintptr
	Fposition size_t
}

var global_error = error1{}

func cJSON_GetErrorPtr() (r uintptr) {
	return global_error.Fjson + uintptr(global_error.Fposition)
}

func cJSON_GetStringValue(item uintptr) (r uintptr) {
	if !(cJSON_IsString(item) != 0) {
		return libc.UintptrFromInt32(0)
	}
	return (*cJSON)(unsafe.Pointer(item)).Fvaluestring
}

func cJSON_GetNumberValue(tls *libc.TLS, item uintptr) (r float64) {
	if !(cJSON_IsNumber(item) != 0) {
		return float64(libc.X__builtin_nanf(tls, __ccgo_ts))
	}
	return (*cJSON)(unsafe.Pointer(item)).Fvaluedouble
}

/* This is a safeguard to prevent copy-pasters from using incompatible C and header files */

func cJSON_Version(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	libc.Xsprintf(tls, uintptr(unsafe.Pointer(&version)), __ccgo_ts+1, libc.VaList(bp+8, int32(CJSON_VERSION_MAJOR), int32(CJSON_VERSION_MINOR), int32(CJSON_VERSION_PATCH)))
	return uintptr(unsafe.Pointer(&version))
}

var version [15]int8

// C documentation
//
//	/* Case insensitive string comparison, doesn't consider two NULL pointers equal though */
func case_insensitive_strcmp(tls *libc.TLS, string1 uintptr, string2 uintptr) (r int32) {
	if string1 == libc.UintptrFromInt32(0) || string2 == libc.UintptrFromInt32(0) {
		return int32(1)
	}
	if string1 == string2 {
		return 0
	}
	for {
		if !(libc.Xtolower(tls, libc.Int32FromUint8(**(**uint8)(__ccgo_up(string1)))) == libc.Xtolower(tls, libc.Int32FromUint8(**(**uint8)(__ccgo_up(string2))))) {
			break
		}
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(string1))) == int32('\000') {
			return 0
		}
		goto _1
	_1:
		;
		string1 = string1 + 1
		string2 = string2 + 1
	}
	return libc.Xtolower(tls, libc.Int32FromUint8(**(**uint8)(__ccgo_up(string1)))) - libc.Xtolower(tls, libc.Int32FromUint8(**(**uint8)(__ccgo_up(string2))))
}

type internal_hooks = struct {
	Fallocate   uintptr
	Fdeallocate uintptr
	Freallocate uintptr
}

/* strlen of character literals resolved at compile time */

var global_hooks = internal_hooks{}

func init() {
	p := unsafe.Pointer(&global_hooks)
	*(*uintptr)(unsafe.Add(p, 0)) = __ccgo_fp(libc.Xmalloc)
	*(*uintptr)(unsafe.Add(p, 8)) = __ccgo_fp(libc.Xfree)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(libc.Xrealloc)
}

func cJSON_strdup(tls *libc.TLS, string1 uintptr, hooks uintptr) (r uintptr) {
	var copy1 uintptr
	var length size_t
	_, _ = copy1, length
	length = uint64(0)
	copy1 = libc.UintptrFromInt32(0)
	if string1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	length = libc.Xstrlen(tls, string1) + uint64(1)
	copy1 = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fallocate})))(tls, length)
	if copy1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	libc.Xmemcpy(tls, copy1, string1, length)
	return copy1
}

func cJSON_InitHooks(hooks uintptr) {
	if hooks == libc.UintptrFromInt32(0) {
		/* Reset hooks */
		global_hooks.Fallocate = __ccgo_fp(libc.Xmalloc)
		global_hooks.Fdeallocate = __ccgo_fp(libc.Xfree)
		global_hooks.Freallocate = __ccgo_fp(libc.Xrealloc)
		return
	}
	global_hooks.Fallocate = __ccgo_fp(libc.Xmalloc)
	if (*cJSON_Hooks)(unsafe.Pointer(hooks)).Fmalloc_fn != libc.UintptrFromInt32(0) {
		global_hooks.Fallocate = (*cJSON_Hooks)(unsafe.Pointer(hooks)).Fmalloc_fn
	}
	global_hooks.Fdeallocate = __ccgo_fp(libc.Xfree)
	if (*cJSON_Hooks)(unsafe.Pointer(hooks)).Ffree_fn != libc.UintptrFromInt32(0) {
		global_hooks.Fdeallocate = (*cJSON_Hooks)(unsafe.Pointer(hooks)).Ffree_fn
	}
	/* use realloc only if both free and malloc are used */
	global_hooks.Freallocate = libc.UintptrFromInt32(0)
	if global_hooks.Fallocate == __ccgo_fp(libc.Xmalloc) && global_hooks.Fdeallocate == __ccgo_fp(libc.Xfree) {
		global_hooks.Freallocate = __ccgo_fp(libc.Xrealloc)
	}
}

// C documentation
//
//	/* Internal constructor. */
func cJSON_New_Item(tls *libc.TLS, hooks uintptr) (r uintptr) {
	var node uintptr
	_ = node
	node = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fallocate})))(tls, uint64(64))
	if node != 0 {
		libc.Xmemset(tls, node, int32('\000'), uint64(64))
	}
	return node
}

// C documentation
//
//	/* Delete a cJSON structure. */
func cJSON_Delete(tls *libc.TLS, item uintptr) {
	var next uintptr
	_ = next
	next = libc.UintptrFromInt32(0)
	for item != libc.UintptrFromInt32(0) {
		next = (*cJSON)(unsafe.Pointer(item)).Fnext
		if !((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(cJSON_IsReference) != 0) && (*cJSON)(unsafe.Pointer(item)).Fchild != libc.UintptrFromInt32(0) {
			cJSON_Delete(tls, (*cJSON)(unsafe.Pointer(item)).Fchild)
		}
		if !((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(cJSON_IsReference) != 0) && (*cJSON)(unsafe.Pointer(item)).Fvaluestring != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{global_hooks.Fdeallocate})))(tls, (*cJSON)(unsafe.Pointer(item)).Fvaluestring)
			(*cJSON)(unsafe.Pointer(item)).Fvaluestring = libc.UintptrFromInt32(0)
		}
		if !((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(cJSON_StringIsConst) != 0) && (*cJSON)(unsafe.Pointer(item)).Fstring1 != libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{global_hooks.Fdeallocate})))(tls, (*cJSON)(unsafe.Pointer(item)).Fstring1)
			(*cJSON)(unsafe.Pointer(item)).Fstring1 = libc.UintptrFromInt32(0)
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{global_hooks.Fdeallocate})))(tls, item)
		item = next
	}
}

// C documentation
//
//	/* get the decimal point character of the current locale */
func get_decimal_point() (r uint8) {
	return uint8('.')
}

type parse_buffer = struct {
	Fcontent uintptr
	Flength  size_t
	Foffset  size_t
	Fdepth   size_t
	Fhooks   internal_hooks
}

/* check if the given size is left to read in a given parse buffer (starting with 1) */
/* check if the buffer can be accessed at the given index (starting with 0) */
/* get a pointer to the buffer at the position */

// C documentation
//
//	/* Parse the input text to generate a number, and populate the result into item. */
func parse_number(tls *libc.TLS, item uintptr, input_buffer uintptr) (r cJSON_bool) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var decimal_point uint8
	var i size_t
	var number float64
	var _ /* after_end at bp+0 */ uintptr
	var _ /* number_c_string at bp+8 */ [64]uint8
	_, _, _ = decimal_point, i, number
	number = libc.Float64FromInt32(0)
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	decimal_point = get_decimal_point()
	i = uint64(0)
	if input_buffer == libc.UintptrFromInt32(0) || (*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* copy the number into a temporary buffer and replace '.' with the decimal point
	 * of the current locale (for strtod)
	 * This also takes care of '\0' not necessarily being available for marking the end of the input */
	i = uint64(0)
	for {
		if !(i < libc.Uint64FromInt64(64)-uint64(1) && (input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+i < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength)) {
			break
		}
		switch libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset) + uintptr(i)))) {
		case int32('0'):
			fallthrough
		case int32('1'):
			fallthrough
		case int32('2'):
			fallthrough
		case int32('3'):
			fallthrough
		case int32('4'):
			fallthrough
		case int32('5'):
			fallthrough
		case int32('6'):
			fallthrough
		case int32('7'):
			fallthrough
		case int32('8'):
			fallthrough
		case int32('9'):
			fallthrough
		case int32('+'):
			fallthrough
		case int32('-'):
			fallthrough
		case int32('e'):
			fallthrough
		case int32('E'):
			(**(**[64]uint8)(__ccgo_up(bp + 8)))[i] = **(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset) + uintptr(i)))
		case int32('.'):
			(**(**[64]uint8)(__ccgo_up(bp + 8)))[i] = decimal_point
		default:
			goto loop_end
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	goto loop_end
loop_end:
	;
	(**(**[64]uint8)(__ccgo_up(bp + 8)))[i] = uint8('\000')
	number = libc.Xstrtod(tls, bp+8, bp)
	if bp+8 == **(**uintptr)(__ccgo_up(bp)) {
		return int32(0) /* parse_error */
	}
	(*cJSON)(unsafe.Pointer(item)).Fvaluedouble = number
	/* use saturation in case of overflow */
	if number >= libc.Float64FromInt32(INT_MAX) {
		(*cJSON)(unsafe.Pointer(item)).Fvalueint = int32(INT_MAX)
	} else {
		if number <= float64(-int32(1)-int32(0x7fffffff)) {
			(*cJSON)(unsafe.Pointer(item)).Fvalueint = -int32(1) - int32(0x7fffffff)
		} else {
			(*cJSON)(unsafe.Pointer(item)).Fvalueint = int32(number)
		}
	}
	(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(3)
	**(**size_t)(__ccgo_up(input_buffer + 16)) += libc.Uint64FromInt64(int64(**(**uintptr)(__ccgo_up(bp))) - __predefined_ptrdiff_t(bp+8))
	return int32(1)
}

// C documentation
//
//	/* don't ask me, but the original cJSON_SetNumberValue returns an integer or double */
func cJSON_SetNumberHelper(object uintptr, number float64) (r float64) {
	var v1 float64
	_ = v1
	if number >= libc.Float64FromInt32(INT_MAX) {
		(*cJSON)(unsafe.Pointer(object)).Fvalueint = int32(INT_MAX)
	} else {
		if number <= float64(-int32(1)-int32(0x7fffffff)) {
			(*cJSON)(unsafe.Pointer(object)).Fvalueint = -int32(1) - int32(0x7fffffff)
		} else {
			(*cJSON)(unsafe.Pointer(object)).Fvalueint = int32(number)
		}
	}
	v1 = number
	(*cJSON)(unsafe.Pointer(object)).Fvaluedouble = v1
	return v1
}

// C documentation
//
//	/* Note: when passing a NULL valuestring, cJSON_SetValuestring treats this as an error and return NULL */
func cJSON_SetValuestring(tls *libc.TLS, object uintptr, valuestring uintptr) (r uintptr) {
	var copy1 uintptr
	_ = copy1
	copy1 = libc.UintptrFromInt32(0)
	/* if object's type is not cJSON_String or is cJSON_IsReference, it should not set valuestring */
	if object == libc.UintptrFromInt32(0) || !((*cJSON)(unsafe.Pointer(object)).Ftype1&(int32(1)<<int32(4)) != 0) || (*cJSON)(unsafe.Pointer(object)).Ftype1&int32(cJSON_IsReference) != 0 {
		return libc.UintptrFromInt32(0)
	}
	/* return NULL if the object is corrupted or valuestring is NULL */
	if (*cJSON)(unsafe.Pointer(object)).Fvaluestring == libc.UintptrFromInt32(0) || valuestring == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if libc.Xstrlen(tls, valuestring) <= libc.Xstrlen(tls, (*cJSON)(unsafe.Pointer(object)).Fvaluestring) {
		libc.Xstrcpy(tls, (*cJSON)(unsafe.Pointer(object)).Fvaluestring, valuestring)
		return (*cJSON)(unsafe.Pointer(object)).Fvaluestring
	}
	copy1 = cJSON_strdup(tls, valuestring, uintptr(unsafe.Pointer(&global_hooks)))
	if copy1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if (*cJSON)(unsafe.Pointer(object)).Fvaluestring != libc.UintptrFromInt32(0) {
		cJSON_free(tls, (*cJSON)(unsafe.Pointer(object)).Fvaluestring)
	}
	(*cJSON)(unsafe.Pointer(object)).Fvaluestring = copy1
	return copy1
}

type printbuffer = struct {
	Fbuffer  uintptr
	Flength  size_t
	Foffset  size_t
	Fdepth   size_t
	Fnoalloc cJSON_bool
	Fformat  cJSON_bool
	Fhooks   internal_hooks
}

// C documentation
//
//	/* realloc printbuffer if necessary to have at least "needed" bytes more */
func ensure(tls *libc.TLS, p uintptr, needed size_t) (r uintptr) {
	var newbuffer uintptr
	var newsize size_t
	_, _ = newbuffer, newsize
	newbuffer = libc.UintptrFromInt32(0)
	newsize = uint64(0)
	if p == libc.UintptrFromInt32(0) || (*printbuffer)(unsafe.Pointer(p)).Fbuffer == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if (*printbuffer)(unsafe.Pointer(p)).Flength > uint64(0) && (*printbuffer)(unsafe.Pointer(p)).Foffset >= (*printbuffer)(unsafe.Pointer(p)).Flength {
		/* make sure that offset is valid */
		return libc.UintptrFromInt32(0)
	}
	if needed > uint64(INT_MAX) {
		/* sizes bigger than INT_MAX are currently not supported */
		return libc.UintptrFromInt32(0)
	}
	needed = needed + ((*printbuffer)(unsafe.Pointer(p)).Foffset + uint64(1))
	if needed <= (*printbuffer)(unsafe.Pointer(p)).Flength {
		return (*printbuffer)(unsafe.Pointer(p)).Fbuffer + uintptr((*printbuffer)(unsafe.Pointer(p)).Foffset)
	}
	if (*printbuffer)(unsafe.Pointer(p)).Fnoalloc != 0 {
		return libc.UintptrFromInt32(0)
	}
	/* calculate new buffer size */
	if needed > uint64(int32(INT_MAX)/int32(2)) {
		/* overflow of int, use INT_MAX if possible */
		if needed <= uint64(INT_MAX) {
			newsize = uint64(INT_MAX)
		} else {
			return libc.UintptrFromInt32(0)
		}
	} else {
		newsize = needed * uint64(2)
	}
	if (*printbuffer)(unsafe.Pointer(p)).Fhooks.Freallocate != libc.UintptrFromInt32(0) {
		/* reallocate with realloc if available */
		newbuffer = (*(*func(*libc.TLS, uintptr, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*printbuffer)(unsafe.Pointer(p)).Fhooks.Freallocate})))(tls, (*printbuffer)(unsafe.Pointer(p)).Fbuffer, newsize)
		if newbuffer == libc.UintptrFromInt32(0) {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*printbuffer)(unsafe.Pointer(p)).Fhooks.Fdeallocate})))(tls, (*printbuffer)(unsafe.Pointer(p)).Fbuffer)
			(*printbuffer)(unsafe.Pointer(p)).Flength = uint64(0)
			(*printbuffer)(unsafe.Pointer(p)).Fbuffer = libc.UintptrFromInt32(0)
			return libc.UintptrFromInt32(0)
		}
	} else {
		/* otherwise reallocate manually */
		newbuffer = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*printbuffer)(unsafe.Pointer(p)).Fhooks.Fallocate})))(tls, newsize)
		if !(newbuffer != 0) {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*printbuffer)(unsafe.Pointer(p)).Fhooks.Fdeallocate})))(tls, (*printbuffer)(unsafe.Pointer(p)).Fbuffer)
			(*printbuffer)(unsafe.Pointer(p)).Flength = uint64(0)
			(*printbuffer)(unsafe.Pointer(p)).Fbuffer = libc.UintptrFromInt32(0)
			return libc.UintptrFromInt32(0)
		}
		libc.Xmemcpy(tls, newbuffer, (*printbuffer)(unsafe.Pointer(p)).Fbuffer, (*printbuffer)(unsafe.Pointer(p)).Foffset+uint64(1))
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*printbuffer)(unsafe.Pointer(p)).Fhooks.Fdeallocate})))(tls, (*printbuffer)(unsafe.Pointer(p)).Fbuffer)
	}
	(*printbuffer)(unsafe.Pointer(p)).Flength = newsize
	(*printbuffer)(unsafe.Pointer(p)).Fbuffer = newbuffer
	return newbuffer + uintptr((*printbuffer)(unsafe.Pointer(p)).Foffset)
}

// C documentation
//
//	/* calculate the new length of the string in a printbuffer and update the offset */
func update_offset(tls *libc.TLS, buffer uintptr) {
	var buffer_pointer uintptr
	_ = buffer_pointer
	buffer_pointer = libc.UintptrFromInt32(0)
	if buffer == libc.UintptrFromInt32(0) || (*printbuffer)(unsafe.Pointer(buffer)).Fbuffer == libc.UintptrFromInt32(0) {
		return
	}
	buffer_pointer = (*printbuffer)(unsafe.Pointer(buffer)).Fbuffer + uintptr((*printbuffer)(unsafe.Pointer(buffer)).Foffset)
	**(**size_t)(__ccgo_up(buffer + 16)) += libc.Xstrlen(tls, buffer_pointer)
}

// C documentation
//
//	/* securely comparison of floating-point variables */
func compare_double(tls *libc.TLS, a float64, b float64) (r cJSON_bool) {
	var maxVal, v1 float64
	_, _ = maxVal, v1
	if libc.Xfabs(tls, a) > libc.Xfabs(tls, b) {
		v1 = libc.Xfabs(tls, a)
	} else {
		v1 = libc.Xfabs(tls, b)
	}
	maxVal = v1
	return libc.BoolInt32(libc.Xfabs(tls, a-b) <= float64(maxVal*libc.Float64FromFloat64(2.220446049250313e-16)))
}

// C documentation
//
//	/* Render the number nicely from the given item into a string. */
func print_number(tls *libc.TLS, item uintptr, output_buffer uintptr) (r cJSON_bool) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var d float64
	var decimal_point uint8
	var i size_t
	var length int32
	var output_pointer uintptr
	var v1, v3 uint64
	var v5 bool
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint64
		F__f float64
	}
	var _ /* number_buffer at bp+8 */ [26]uint8
	var _ /* test at bp+40 */ float64
	_, _, _, _, _, _, _, _ = d, decimal_point, i, length, output_pointer, v1, v3, v5
	output_pointer = libc.UintptrFromInt32(0)
	d = (*cJSON)(unsafe.Pointer(item)).Fvaluedouble
	length = 0
	i = uint64(0)
	**(**[26]uint8)(__ccgo_up(bp + 8)) = [26]uint8{} /* temporary buffer to print the number into */
	decimal_point = get_decimal_point()
	**(**float64)(__ccgo_up(bp + 40)) = float64(0)
	if output_buffer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* This checks for NaN and Infinity */
	*(*float64)(unsafe.Pointer(bp)) = d
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	;
	if v5 = libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>int32(1)) > libc.Uint64FromUint64(0x7ff)<<int32(52)) != 0; !v5 {
		*(*float64)(unsafe.Pointer(bp)) = d
		v3 = *(*uint64)(unsafe.Pointer(bp))
		goto _4
	_4:
	}
	if v5 || libc.BoolInt32(v3&(-libc.Uint64FromUint64(1)>>int32(1)) == libc.Uint64FromUint64(0x7ff)<<int32(52)) != 0 {
		length = libc.Xsprintf(tls, bp+8, __ccgo_ts+10, 0)
	} else {
		if d == float64((*cJSON)(unsafe.Pointer(item)).Fvalueint) {
			length = libc.Xsprintf(tls, bp+8, __ccgo_ts+15, libc.VaList(bp+56, (*cJSON)(unsafe.Pointer(item)).Fvalueint))
		} else {
			/* Try 15 decimal places of precision to avoid nonsignificant nonzero digits */
			length = libc.Xsprintf(tls, bp+8, __ccgo_ts+18, libc.VaList(bp+56, d))
			/* Check whether the original double can be recovered */
			if libc.Xsscanf(tls, bp+8, __ccgo_ts+25, libc.VaList(bp+56, bp+40)) != int32(1) || !(compare_double(tls, **(**float64)(__ccgo_up(bp + 40)), d) != 0) {
				/* If not, print with 17 decimal places of precision */
				length = libc.Xsprintf(tls, bp+8, __ccgo_ts+29, libc.VaList(bp+56, d))
			}
		}
	}
	/* sprintf failed or buffer overrun occurred */
	if length < 0 || length > libc.Int32FromUint64(libc.Uint64FromInt64(26)-uint64(1)) {
		return int32(0)
	}
	/* reserve appropriate space in the output */
	output_pointer = ensure(tls, output_buffer, uint64(length)+uint64(1))
	if output_pointer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* copy the printed number to the output and replace locale
	 * dependent decimal point with '.' */
	i = uint64(0)
	for {
		if !(i < uint64(length)) {
			break
		}
		if libc.Int32FromUint8((**(**[26]uint8)(__ccgo_up(bp + 8)))[i]) == libc.Int32FromUint8(decimal_point) {
			**(**uint8)(__ccgo_up(output_pointer + uintptr(i))) = uint8('.')
			goto _6
		}
		**(**uint8)(__ccgo_up(output_pointer + uintptr(i))) = (**(**[26]uint8)(__ccgo_up(bp + 8)))[i]
		goto _6
	_6:
		;
		i = i + 1
	}
	**(**uint8)(__ccgo_up(output_pointer + uintptr(i))) = uint8('\000')
	**(**size_t)(__ccgo_up(output_buffer + 16)) += uint64(length)
	return int32(1)
}

// C documentation
//
//	/* parse 4 digit hexadecimal number */
func parse_hex4(input uintptr) (r uint32) {
	var h uint32
	var i size_t
	_, _ = h, i
	h = uint32(0)
	i = uint64(0)
	i = uint64(0)
	for {
		if !(i < uint64(4)) {
			break
		}
		/* parse digit */
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(input + uintptr(i)))) >= int32('0') && libc.Int32FromUint8(**(**uint8)(__ccgo_up(input + uintptr(i)))) <= int32('9') {
			h = h + (uint32(**(**uint8)(__ccgo_up(input + uintptr(i)))) - uint32('0'))
		} else {
			if libc.Int32FromUint8(**(**uint8)(__ccgo_up(input + uintptr(i)))) >= int32('A') && libc.Int32FromUint8(**(**uint8)(__ccgo_up(input + uintptr(i)))) <= int32('F') {
				h = h + (uint32(10) + uint32(**(**uint8)(__ccgo_up(input + uintptr(i)))) - uint32('A'))
			} else {
				if libc.Int32FromUint8(**(**uint8)(__ccgo_up(input + uintptr(i)))) >= int32('a') && libc.Int32FromUint8(**(**uint8)(__ccgo_up(input + uintptr(i)))) <= int32('f') {
					h = h + (uint32(10) + uint32(**(**uint8)(__ccgo_up(input + uintptr(i)))) - uint32('a'))
				} else { /* invalid */
					return uint32(0)
				}
			}
		}
		if i < uint64(3) {
			/* shift left to make place for the next nibble */
			h = h << int32(4)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return h
}

// C documentation
//
//	/* converts a UTF-16 literal to UTF-8
//	 * A literal can be one or two sequences of the form \uXXXX */
func utf16_literal_to_utf8(input_pointer uintptr, input_end uintptr, output_pointer uintptr) (r uint8) {
	var codepoint uint64
	var first_byte_mark, sequence_length, utf8_length, utf8_position uint8
	var first_code, second_code uint32
	var first_sequence, second_sequence uintptr
	_, _, _, _, _, _, _, _, _ = codepoint, first_byte_mark, first_code, first_sequence, second_code, second_sequence, sequence_length, utf8_length, utf8_position
	codepoint = uint64(0)
	first_code = uint32(0)
	first_sequence = input_pointer
	utf8_length = uint8(0)
	utf8_position = uint8(0)
	sequence_length = uint8(0)
	first_byte_mark = uint8(0)
	if int64(input_end)-int64(first_sequence) < int64(6) {
		/* input ends unexpectedly */
		goto fail
	}
	/* get the first utf16 sequence */
	first_code = parse_hex4(first_sequence + uintptr(2))
	/* check that the code is valid */
	if first_code >= uint32(0xDC00) && first_code <= uint32(0xDFFF) {
		goto fail
	}
	/* UTF16 surrogate pair */
	if first_code >= uint32(0xD800) && first_code <= uint32(0xDBFF) {
		second_sequence = first_sequence + uintptr(6)
		second_code = uint32(0)
		sequence_length = uint8(12) /* \uXXXX\uXXXX */
		if int64(input_end)-int64(second_sequence) < int64(6) {
			/* input ends unexpectedly */
			goto fail
		}
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(second_sequence))) != int32('\\') || libc.Int32FromUint8(**(**uint8)(__ccgo_up(second_sequence + 1))) != int32('u') {
			/* missing second half of the surrogate pair */
			goto fail
		}
		/* get the second utf16 sequence */
		second_code = parse_hex4(second_sequence + uintptr(2))
		/* check that the code is valid */
		if second_code < uint32(0xDC00) || second_code > uint32(0xDFFF) {
			/* invalid second half of the surrogate pair */
			goto fail
		}
		/* calculate the unicode codepoint from the surrogate pair */
		codepoint = uint64(uint32(0x10000) + (first_code&uint32(0x3FF)<<int32(10) | second_code&uint32(0x3FF)))
	} else {
		sequence_length = uint8(6) /* \uXXXX */
		codepoint = uint64(first_code)
	}
	/* encode as UTF-8
	 * takes at maximum 4 bytes to encode:
	 * 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx */
	if codepoint < uint64(0x80) {
		/* normal ascii, encoding 0xxxxxxx */
		utf8_length = uint8(1)
	} else {
		if codepoint < uint64(0x800) {
			/* two bytes, encoding 110xxxxx 10xxxxxx */
			utf8_length = uint8(2)
			first_byte_mark = uint8(0xC0) /* 11000000 */
		} else {
			if codepoint < uint64(0x10000) {
				/* three bytes, encoding 1110xxxx 10xxxxxx 10xxxxxx */
				utf8_length = uint8(3)
				first_byte_mark = uint8(0xE0) /* 11100000 */
			} else {
				if codepoint <= uint64(0x10FFFF) {
					/* four bytes, encoding 1110xxxx 10xxxxxx 10xxxxxx 10xxxxxx */
					utf8_length = uint8(4)
					first_byte_mark = uint8(0xF0) /* 11110000 */
				} else {
					/* invalid unicode codepoint */
					goto fail
				}
			}
		}
	}
	/* encode as utf8 */
	utf8_position = libc.Uint8FromInt32(libc.Int32FromUint8(utf8_length) - int32(1))
	for {
		if !(libc.Int32FromUint8(utf8_position) > 0) {
			break
		}
		/* 10xxxxxx */
		**(**uint8)(__ccgo_up(**(**uintptr)(__ccgo_up(output_pointer)) + uintptr(utf8_position))) = uint8((codepoint | uint64(0x80)) & uint64(0xBF))
		codepoint = codepoint >> uint64(6)
		goto _1
	_1:
		;
		utf8_position = utf8_position - 1
	}
	/* encode first byte */
	if libc.Int32FromUint8(utf8_length) > int32(1) {
		**(**uint8)(__ccgo_up(**(**uintptr)(__ccgo_up(output_pointer)))) = uint8((codepoint | uint64(first_byte_mark)) & uint64(0xFF))
	} else {
		**(**uint8)(__ccgo_up(**(**uintptr)(__ccgo_up(output_pointer)))) = uint8(codepoint & uint64(0x7F))
	}
	**(**uintptr)(__ccgo_up(output_pointer)) += uintptr(utf8_length)
	return sequence_length
	goto fail
fail:
	;
	return uint8(0)
	return r
}

// C documentation
//
//	/* Parse the input text into an unescaped cinput, and populate item. */
func parse_string(tls *libc.TLS, item uintptr, input_buffer uintptr) (r cJSON_bool) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var allocation_length, skipped_bytes size_t
	var input_end, input_pointer, output, v1, v2 uintptr
	var sequence_length uint8
	var _ /* output_pointer at bp+0 */ uintptr
	_, _, _, _, _, _, _, _ = allocation_length, input_end, input_pointer, output, sequence_length, skipped_bytes, v1, v2
	input_pointer = (*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset) + uintptr(1)
	input_end = (*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset) + uintptr(1)
	**(**uintptr)(__ccgo_up(bp)) = libc.UintptrFromInt32(0)
	output = libc.UintptrFromInt32(0)
	/* not a string */
	if libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) != int32('"') {
		goto fail
	}
	/* calculate approximate size of the output (overestimate) */
	allocation_length = uint64(0)
	skipped_bytes = uint64(0)
	for libc.Uint64FromInt64(int64(input_end)-int64((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent)) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_end))) != int32('"') {
		/* is escape sequence */
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_end))) == int32('\\') {
			if libc.Uint64FromInt64(int64(input_end+libc.UintptrFromInt32(1))-int64((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent)) >= (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength {
				/* prevent buffer overflow when last input character is a backslash */
				goto fail
			}
			skipped_bytes = skipped_bytes + 1
			input_end = input_end + 1
		}
		input_end = input_end + 1
	}
	if libc.Uint64FromInt64(int64(input_end)-int64((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent)) >= (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength || libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_end))) != int32('"') {
		goto fail /* string ended unexpectedly */
	}
	/* This is at most how much we need for the output */
	allocation_length = libc.Uint64FromInt64(int64(input_end)-int64((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent+uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset))) - skipped_bytes
	output = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*parse_buffer)(unsafe.Pointer(input_buffer)).Fhooks.Fallocate})))(tls, allocation_length+uint64(1))
	if output == libc.UintptrFromInt32(0) {
		goto fail /* allocation failure */
	}
	**(**uintptr)(__ccgo_up(bp)) = output
	/* loop through the string literal */
	for input_pointer < input_end {
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) != int32('\\') {
			v1 = **(**uintptr)(__ccgo_up(bp))
			**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
			v2 = input_pointer
			input_pointer = input_pointer + 1
			**(**uint8)(__ccgo_up(v1)) = **(**uint8)(__ccgo_up(v2))
		} else {
			sequence_length = uint8(2)
			if int64(input_end)-int64(input_pointer) < int64(1) {
				goto fail
			}
			switch libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer + 1))) {
			case int32('b'):
				v1 = **(**uintptr)(__ccgo_up(bp))
				**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
				**(**uint8)(__ccgo_up(v1)) = uint8('\b')
			case int32('f'):
				v1 = **(**uintptr)(__ccgo_up(bp))
				**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
				**(**uint8)(__ccgo_up(v1)) = uint8('\f')
			case int32('n'):
				v1 = **(**uintptr)(__ccgo_up(bp))
				**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
				**(**uint8)(__ccgo_up(v1)) = uint8('\n')
			case int32('r'):
				v1 = **(**uintptr)(__ccgo_up(bp))
				**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
				**(**uint8)(__ccgo_up(v1)) = uint8('\r')
			case int32('t'):
				v1 = **(**uintptr)(__ccgo_up(bp))
				**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
				**(**uint8)(__ccgo_up(v1)) = uint8('\t')
			case int32('"'):
				fallthrough
			case int32('\\'):
				fallthrough
			case int32('/'):
				v1 = **(**uintptr)(__ccgo_up(bp))
				**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
				**(**uint8)(__ccgo_up(v1)) = **(**uint8)(__ccgo_up(input_pointer + 1))
				break
				/* UTF-16 literal */
				fallthrough
			case int32('u'):
				sequence_length = utf16_literal_to_utf8(input_pointer, input_end, bp)
				if libc.Int32FromUint8(sequence_length) == 0 {
					/* failed to convert UTF16-literal to UTF-8 */
					goto fail
				}
			default:
				goto fail
			}
			input_pointer = input_pointer + uintptr(sequence_length)
		}
	}
	/* zero terminate the output */
	**(**uint8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)))) = uint8('\000')
	(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(4)
	(*cJSON)(unsafe.Pointer(item)).Fvaluestring = output
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = libc.Uint64FromInt64(int64(input_end) - int64((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent))
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
	return int32(1)
	goto fail
fail:
	;
	if output != libc.UintptrFromInt32(0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*parse_buffer)(unsafe.Pointer(input_buffer)).Fhooks.Fdeallocate})))(tls, output)
		output = libc.UintptrFromInt32(0)
	}
	if input_pointer != libc.UintptrFromInt32(0) {
		(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = libc.Uint64FromInt64(int64(input_pointer) - int64((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent))
	}
	return int32(0)
}

// C documentation
//
//	/* Render the cstring provided to an escaped version that can be printed. */
func print_string_ptr(tls *libc.TLS, input uintptr, output_buffer uintptr) (r cJSON_bool) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var escape_characters, output_length size_t
	var input_pointer, output, output_pointer, v3 uintptr
	_, _, _, _, _, _ = escape_characters, input_pointer, output, output_length, output_pointer, v3
	input_pointer = libc.UintptrFromInt32(0)
	output = libc.UintptrFromInt32(0)
	output_pointer = libc.UintptrFromInt32(0)
	output_length = uint64(0)
	/* numbers of additional characters needed for escaping */
	escape_characters = uint64(0)
	if output_buffer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* empty string */
	if input == libc.UintptrFromInt32(0) {
		output = ensure(tls, output_buffer, uint64(3))
		if output == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		libc.Xstrcpy(tls, output, __ccgo_ts+36)
		return int32(1)
	}
	/* set "flag" to 1 if something needs to be escaped */
	input_pointer = input
	for {
		if !(**(**uint8)(__ccgo_up(input_pointer)) != 0) {
			break
		}
		switch libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) {
		case int32('"'):
			fallthrough
		case int32('\\'):
			fallthrough
		case int32('\b'):
			fallthrough
		case int32('\f'):
			fallthrough
		case int32('\n'):
			fallthrough
		case int32('\r'):
			fallthrough
		case int32('\t'):
			/* one character escape sequence */
			escape_characters = escape_characters + 1
		default:
			if libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) < int32(32) {
				/* UTF-16 escape sequence uXXXX */
				escape_characters = escape_characters + uint64(5)
			}
			break
		}
		goto _1
	_1:
		;
		input_pointer = input_pointer + 1
	}
	output_length = libc.Uint64FromInt64(int64(input_pointer)-int64(input)) + escape_characters
	output = ensure(tls, output_buffer, output_length+uint64(3))
	if output == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* no characters have to be escaped */
	if escape_characters == uint64(0) {
		**(**uint8)(__ccgo_up(output)) = uint8('"')
		libc.Xmemcpy(tls, output+uintptr(1), input, output_length)
		**(**uint8)(__ccgo_up(output + uintptr(output_length+uint64(1)))) = uint8('"')
		**(**uint8)(__ccgo_up(output + uintptr(output_length+uint64(2)))) = uint8('\000')
		return int32(1)
	}
	**(**uint8)(__ccgo_up(output)) = uint8('"')
	output_pointer = output + uintptr(1)
	/* copy the string */
	input_pointer = input
	for {
		if !(libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) != int32('\000')) {
			break
		}
		if libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) > int32(31) && libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) != int32('"') && libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) != int32('\\') {
			/* normal character, copy */
			**(**uint8)(__ccgo_up(output_pointer)) = **(**uint8)(__ccgo_up(input_pointer))
		} else {
			/* character needs to be escaped */
			v3 = output_pointer
			output_pointer = output_pointer + 1
			**(**uint8)(__ccgo_up(v3)) = uint8('\\')
			switch libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer))) {
			case int32('\\'):
				**(**uint8)(__ccgo_up(output_pointer)) = uint8('\\')
			case int32('"'):
				**(**uint8)(__ccgo_up(output_pointer)) = uint8('"')
			case int32('\b'):
				**(**uint8)(__ccgo_up(output_pointer)) = uint8('b')
			case int32('\f'):
				**(**uint8)(__ccgo_up(output_pointer)) = uint8('f')
			case int32('\n'):
				**(**uint8)(__ccgo_up(output_pointer)) = uint8('n')
			case int32('\r'):
				**(**uint8)(__ccgo_up(output_pointer)) = uint8('r')
			case int32('\t'):
				**(**uint8)(__ccgo_up(output_pointer)) = uint8('t')
			default:
				/* escape and print as unicode codepoint */
				libc.Xsprintf(tls, output_pointer, __ccgo_ts+39, libc.VaList(bp+8, libc.Int32FromUint8(**(**uint8)(__ccgo_up(input_pointer)))))
				output_pointer = output_pointer + uintptr(4)
				break
			}
		}
		goto _2
	_2:
		;
		input_pointer = input_pointer + 1
		output_pointer = output_pointer + 1
	}
	**(**uint8)(__ccgo_up(output + uintptr(output_length+uint64(1)))) = uint8('"')
	**(**uint8)(__ccgo_up(output + uintptr(output_length+uint64(2)))) = uint8('\000')
	return int32(1)
}

// C documentation
//
//	/* Invoke print_string_ptr (which is useful) on an item. */
func print_string(tls *libc.TLS, item uintptr, p uintptr) (r cJSON_bool) {
	return print_string_ptr(tls, (*cJSON)(unsafe.Pointer(item)).Fvaluestring, p)
}

// C documentation
//
//	/* Utility to jump whitespace and cr/lf */
func buffer_skip_whitespace(buffer uintptr) (r uintptr) {
	if buffer == libc.UintptrFromInt32(0) || (*parse_buffer)(unsafe.Pointer(buffer)).Fcontent == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if !(buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(buffer)).Flength) {
		return buffer
	}
	for buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(buffer)).Foffset)))) <= int32(32) {
		(*parse_buffer)(unsafe.Pointer(buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(buffer)).Foffset + 1
	}
	if (*parse_buffer)(unsafe.Pointer(buffer)).Foffset == (*parse_buffer)(unsafe.Pointer(buffer)).Flength {
		(*parse_buffer)(unsafe.Pointer(buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(buffer)).Foffset - 1
	}
	return buffer
}

// C documentation
//
//	/* skip the UTF-8 BOM (byte order mark) if it is at the beginning of a buffer */
func skip_utf8_bom(tls *libc.TLS, buffer uintptr) (r uintptr) {
	if buffer == libc.UintptrFromInt32(0) || (*parse_buffer)(unsafe.Pointer(buffer)).Fcontent == libc.UintptrFromInt32(0) || (*parse_buffer)(unsafe.Pointer(buffer)).Foffset != uint64(0) {
		return libc.UintptrFromInt32(0)
	}
	if buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(buffer)).Foffset+uint64(4) < (*parse_buffer)(unsafe.Pointer(buffer)).Flength && libc.Xstrncmp(tls, (*parse_buffer)(unsafe.Pointer(buffer)).Fcontent+uintptr((*parse_buffer)(unsafe.Pointer(buffer)).Foffset), __ccgo_ts+45, uint64(3)) == 0 {
		**(**size_t)(__ccgo_up(buffer + 16)) += uint64(3)
	}
	return buffer
}

func cJSON_ParseWithOpts(tls *libc.TLS, value uintptr, return_parse_end uintptr, require_null_terminated cJSON_bool) (r uintptr) {
	var buffer_length size_t
	_ = buffer_length
	if libc.UintptrFromInt32(0) == value {
		return libc.UintptrFromInt32(0)
	}
	/* Adding null character size due to require_null_terminated. */
	buffer_length = libc.Xstrlen(tls, value) + uint64(1)
	return cJSON_ParseWithLengthOpts(tls, value, buffer_length, return_parse_end, require_null_terminated)
}

// C documentation
//
//	/* Parse an object - create a new root, and populate. */
func cJSON_ParseWithLengthOpts(tls *libc.TLS, value uintptr, buffer_length size_t, return_parse_end uintptr, require_null_terminated cJSON_bool) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var item uintptr
	var local_error error1
	var _ /* buffer at bp+0 */ parse_buffer
	_, _ = item, local_error
	**(**parse_buffer)(__ccgo_up(bp)) = parse_buffer{}
	item = libc.UintptrFromInt32(0)
	/* reset error position */
	global_error.Fjson = libc.UintptrFromInt32(0)
	global_error.Fposition = uint64(0)
	if value == libc.UintptrFromInt32(0) || uint64(0) == buffer_length {
		goto fail
	}
	(**(**parse_buffer)(__ccgo_up(bp))).Fcontent = value
	(**(**parse_buffer)(__ccgo_up(bp))).Flength = buffer_length
	(**(**parse_buffer)(__ccgo_up(bp))).Foffset = uint64(0)
	(**(**parse_buffer)(__ccgo_up(bp))).Fhooks = global_hooks
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item == libc.UintptrFromInt32(0) { /* memory fail */
		goto fail
	}
	if !(parse_value(tls, item, buffer_skip_whitespace(skip_utf8_bom(tls, bp))) != 0) {
		/* parse failure. ep is set. */
		goto fail
	}
	/* if we require null-terminated JSON without appended garbage, skip and then check for a null terminator */
	if require_null_terminated != 0 {
		buffer_skip_whitespace(bp)
		if (**(**parse_buffer)(__ccgo_up(bp))).Foffset >= (**(**parse_buffer)(__ccgo_up(bp))).Flength || libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(bp)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(bp)).Foffset)))) != int32('\000') {
			goto fail
		}
	}
	if return_parse_end != 0 {
		**(**uintptr)(__ccgo_up(return_parse_end)) = (*parse_buffer)(unsafe.Pointer(bp)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(bp)).Foffset)
	}
	return item
	goto fail
fail:
	;
	if item != libc.UintptrFromInt32(0) {
		cJSON_Delete(tls, item)
	}
	if value != libc.UintptrFromInt32(0) {
		local_error.Fjson = value
		local_error.Fposition = uint64(0)
		if (**(**parse_buffer)(__ccgo_up(bp))).Foffset < (**(**parse_buffer)(__ccgo_up(bp))).Flength {
			local_error.Fposition = (**(**parse_buffer)(__ccgo_up(bp))).Foffset
		} else {
			if (**(**parse_buffer)(__ccgo_up(bp))).Flength > uint64(0) {
				local_error.Fposition = (**(**parse_buffer)(__ccgo_up(bp))).Flength - uint64(1)
			}
		}
		if return_parse_end != libc.UintptrFromInt32(0) {
			**(**uintptr)(__ccgo_up(return_parse_end)) = local_error.Fjson + uintptr(local_error.Fposition)
		}
		global_error = local_error
	}
	return libc.UintptrFromInt32(0)
}

// C documentation
//
//	/* Default options for cJSON_Parse */
func cJSON_Parse(tls *libc.TLS, value uintptr) (r uintptr) {
	return cJSON_ParseWithOpts(tls, value, uintptr(0), 0)
}

func cJSON_ParseWithLength(tls *libc.TLS, value uintptr, buffer_length size_t) (r uintptr) {
	return cJSON_ParseWithLengthOpts(tls, value, buffer_length, uintptr(0), 0)
}

func print1(tls *libc.TLS, item uintptr, format cJSON_bool, hooks uintptr) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var printed uintptr
	var v1 uint64
	var _ /* buffer at bp+0 */ [1]printbuffer
	_, _ = printed, v1
	printed = libc.UintptrFromInt32(0)
	libc.Xmemset(tls, bp, 0, uint64(64))
	/* create buffer */
	(*printbuffer)(unsafe.Pointer(bp)).Fbuffer = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fallocate})))(tls, default_buffer_size)
	(*printbuffer)(unsafe.Pointer(bp)).Flength = default_buffer_size
	(*printbuffer)(unsafe.Pointer(bp)).Fformat = format
	(*printbuffer)(unsafe.Pointer(bp)).Fhooks = **(**internal_hooks)(__ccgo_up(hooks))
	if (*printbuffer)(unsafe.Pointer(bp)).Fbuffer == libc.UintptrFromInt32(0) {
		goto fail
	}
	/* print the value */
	if !(print_value(tls, item, bp) != 0) {
		goto fail
	}
	update_offset(tls, bp)
	/* check if reallocate is available */
	if (*internal_hooks)(unsafe.Pointer(hooks)).Freallocate != libc.UintptrFromInt32(0) {
		printed = (*(*func(*libc.TLS, uintptr, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Freallocate})))(tls, (*printbuffer)(unsafe.Pointer(bp)).Fbuffer, (*printbuffer)(unsafe.Pointer(bp)).Foffset+uint64(1))
		if printed == libc.UintptrFromInt32(0) {
			goto fail
		}
		(*printbuffer)(unsafe.Pointer(bp)).Fbuffer = libc.UintptrFromInt32(0)
	} else { /* otherwise copy the JSON over to a new buffer */
		printed = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fallocate})))(tls, (*printbuffer)(unsafe.Pointer(bp)).Foffset+uint64(1))
		if printed == libc.UintptrFromInt32(0) {
			goto fail
		}
		if (*printbuffer)(unsafe.Pointer(bp)).Flength < (*printbuffer)(unsafe.Pointer(bp)).Foffset+uint64(1) {
			v1 = (*printbuffer)(unsafe.Pointer(bp)).Flength
		} else {
			v1 = (*printbuffer)(unsafe.Pointer(bp)).Foffset + uint64(1)
		}
		libc.Xmemcpy(tls, printed, (*printbuffer)(unsafe.Pointer(bp)).Fbuffer, v1)
		**(**uint8)(__ccgo_up(printed + uintptr((*printbuffer)(unsafe.Pointer(bp)).Foffset))) = uint8('\000') /* just to be sure */
		/* free the buffer */
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fdeallocate})))(tls, (*printbuffer)(unsafe.Pointer(bp)).Fbuffer)
		(*printbuffer)(unsafe.Pointer(bp)).Fbuffer = libc.UintptrFromInt32(0)
	}
	return printed
	goto fail
fail:
	;
	if (*printbuffer)(unsafe.Pointer(bp)).Fbuffer != libc.UintptrFromInt32(0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fdeallocate})))(tls, (*printbuffer)(unsafe.Pointer(bp)).Fbuffer)
		(*printbuffer)(unsafe.Pointer(bp)).Fbuffer = libc.UintptrFromInt32(0)
	}
	if printed != libc.UintptrFromInt32(0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fdeallocate})))(tls, printed)
		printed = libc.UintptrFromInt32(0)
	}
	return libc.UintptrFromInt32(0)
}

var default_buffer_size = uint64(256)

// C documentation
//
//	/* Render a cJSON item/entity/structure to text. */
func cJSON_Print(tls *libc.TLS, item uintptr) (r uintptr) {
	return print1(tls, item, int32(1), uintptr(unsafe.Pointer(&global_hooks)))
}

func cJSON_PrintUnformatted(tls *libc.TLS, item uintptr) (r uintptr) {
	return print1(tls, item, int32(0), uintptr(unsafe.Pointer(&global_hooks)))
}

func cJSON_PrintBuffered(tls *libc.TLS, item uintptr, prebuffer int32, fmt cJSON_bool) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* p at bp+0 */ printbuffer
	**(**printbuffer)(__ccgo_up(bp)) = printbuffer{}
	if prebuffer < 0 {
		return libc.UintptrFromInt32(0)
	}
	(**(**printbuffer)(__ccgo_up(bp))).Fbuffer = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{global_hooks.Fallocate})))(tls, uint64(prebuffer))
	if !((**(**printbuffer)(__ccgo_up(bp))).Fbuffer != 0) {
		return libc.UintptrFromInt32(0)
	}
	(**(**printbuffer)(__ccgo_up(bp))).Flength = uint64(prebuffer)
	(**(**printbuffer)(__ccgo_up(bp))).Foffset = uint64(0)
	(**(**printbuffer)(__ccgo_up(bp))).Fnoalloc = int32(0)
	(**(**printbuffer)(__ccgo_up(bp))).Fformat = fmt
	(**(**printbuffer)(__ccgo_up(bp))).Fhooks = global_hooks
	if !(print_value(tls, item, bp) != 0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{global_hooks.Fdeallocate})))(tls, (**(**printbuffer)(__ccgo_up(bp))).Fbuffer)
		(**(**printbuffer)(__ccgo_up(bp))).Fbuffer = libc.UintptrFromInt32(0)
		return libc.UintptrFromInt32(0)
	}
	return (**(**printbuffer)(__ccgo_up(bp))).Fbuffer
}

func cJSON_PrintPreallocated(tls *libc.TLS, item uintptr, buffer uintptr, length int32, format cJSON_bool) (r cJSON_bool) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* p at bp+0 */ printbuffer
	**(**printbuffer)(__ccgo_up(bp)) = printbuffer{}
	if length < 0 || buffer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	(**(**printbuffer)(__ccgo_up(bp))).Fbuffer = buffer
	(**(**printbuffer)(__ccgo_up(bp))).Flength = uint64(length)
	(**(**printbuffer)(__ccgo_up(bp))).Foffset = uint64(0)
	(**(**printbuffer)(__ccgo_up(bp))).Fnoalloc = int32(1)
	(**(**printbuffer)(__ccgo_up(bp))).Fformat = format
	(**(**printbuffer)(__ccgo_up(bp))).Fhooks = global_hooks
	return print_value(tls, item, bp)
}

// C documentation
//
//	/* Parser core - when encountering text, process appropriately. */
func parse_value(tls *libc.TLS, item uintptr, input_buffer uintptr) (r cJSON_bool) {
	if input_buffer == libc.UintptrFromInt32(0) || (*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent == libc.UintptrFromInt32(0) {
		return int32(0) /* no input */
	}
	/* parse the different types of values */
	/* null */
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(4) <= (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Xstrncmp(tls, (*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent+uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset), __ccgo_ts+10, uint64(4)) == 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(2)
		**(**size_t)(__ccgo_up(input_buffer + 16)) += uint64(4)
		return int32(1)
	}
	/* false */
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(5) <= (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Xstrncmp(tls, (*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent+uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset), __ccgo_ts+49, uint64(5)) == 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(0)
		**(**size_t)(__ccgo_up(input_buffer + 16)) += uint64(5)
		return int32(1)
	}
	/* true */
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(4) <= (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Xstrncmp(tls, (*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent+uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset), __ccgo_ts+55, uint64(4)) == 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(1)
		(*cJSON)(unsafe.Pointer(item)).Fvalueint = int32(1)
		**(**size_t)(__ccgo_up(input_buffer + 16)) += uint64(4)
		return int32(1)
	}
	/* string */
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32('"') {
		return parse_string(tls, item, input_buffer)
	}
	/* number */
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && (libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32('-') || libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) >= int32('0') && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) <= int32('9')) {
		return parse_number(tls, item, input_buffer)
	}
	/* array */
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32('[') {
		return parse_array(tls, item, input_buffer)
	}
	/* object */
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32('{') {
		return parse_object(tls, item, input_buffer)
	}
	return int32(0)
}

// C documentation
//
//	/* Render a value to text. */
func print_value(tls *libc.TLS, item uintptr, output_buffer uintptr) (r cJSON_bool) {
	var output uintptr
	var raw_length size_t
	_, _ = output, raw_length
	output = libc.UintptrFromInt32(0)
	if item == libc.UintptrFromInt32(0) || output_buffer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	switch (*cJSON)(unsafe.Pointer(item)).Ftype1 & int32(0xFF) {
	case int32(1) << int32(2):
		output = ensure(tls, output_buffer, uint64(5))
		if output == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		libc.Xstrcpy(tls, output, __ccgo_ts+10)
		return int32(1)
	case int32(1) << int32(0):
		output = ensure(tls, output_buffer, uint64(6))
		if output == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		libc.Xstrcpy(tls, output, __ccgo_ts+49)
		return int32(1)
	case int32(1) << int32(1):
		output = ensure(tls, output_buffer, uint64(5))
		if output == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		libc.Xstrcpy(tls, output, __ccgo_ts+55)
		return int32(1)
	case int32(1) << int32(3):
		return print_number(tls, item, output_buffer)
	case int32(1) << int32(7):
		raw_length = uint64(0)
		if (*cJSON)(unsafe.Pointer(item)).Fvaluestring == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		raw_length = libc.Xstrlen(tls, (*cJSON)(unsafe.Pointer(item)).Fvaluestring) + uint64(1)
		output = ensure(tls, output_buffer, raw_length)
		if output == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		libc.Xmemcpy(tls, output, (*cJSON)(unsafe.Pointer(item)).Fvaluestring, raw_length)
		return int32(1)
	case int32(1) << int32(4):
		return print_string(tls, item, output_buffer)
	case int32(1) << int32(5):
		return print_array(tls, item, output_buffer)
	case int32(1) << int32(6):
		return print_object(tls, item, output_buffer)
	default:
		return int32(0)
	}
	return r
}

// C documentation
//
//	/* Build an array from input text. */
func parse_array(tls *libc.TLS, item uintptr, input_buffer uintptr) (r cJSON_bool) {
	var current_item, head, new_item, v1 uintptr
	_, _, _, _ = current_item, head, new_item, v1
	head = libc.UintptrFromInt32(0) /* head of the linked list */
	current_item = libc.UintptrFromInt32(0)
	if (*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth >= uint64(CJSON_NESTING_LIMIT) {
		return int32(0) /* to deeply nested */
	}
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth = (*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth + 1
	if libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) != int32('[') {
		/* not an array */
		goto fail
	}
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
	buffer_skip_whitespace(input_buffer)
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32(']') {
		/* empty array */
		goto success
	}
	/* check if we skipped to the end of the buffer */
	if !(input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength) {
		(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset - 1
		goto fail
	}
	/* step back to character in front of the first element */
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset - 1
	/* loop through the comma separated array elements */
	for cond := true; cond; cond = input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32(',') {
		/* allocate next item */
		new_item = cJSON_New_Item(tls, input_buffer+32)
		if new_item == libc.UintptrFromInt32(0) {
			goto fail /* allocation failure */
		}
		/* attach next item to list */
		if head == libc.UintptrFromInt32(0) {
			/* start the linked list */
			v1 = new_item
			head = v1
			current_item = v1
		} else {
			/* add to the end and advance */
			(*cJSON)(unsafe.Pointer(current_item)).Fnext = new_item
			(*cJSON)(unsafe.Pointer(new_item)).Fprev = current_item
			current_item = new_item
		}
		/* parse next value */
		(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
		buffer_skip_whitespace(input_buffer)
		if !(parse_value(tls, current_item, input_buffer) != 0) {
			goto fail /* failed to parse value */
		}
		buffer_skip_whitespace(input_buffer)
	}
	if !(input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength) || libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) != int32(']') {
		goto fail /* expected end of array */
	}
	goto success
success:
	;
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth = (*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth - 1
	if head != libc.UintptrFromInt32(0) {
		(*cJSON)(unsafe.Pointer(head)).Fprev = current_item
	}
	(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(5)
	(*cJSON)(unsafe.Pointer(item)).Fchild = head
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
	return int32(1)
	goto fail
fail:
	;
	if head != libc.UintptrFromInt32(0) {
		cJSON_Delete(tls, head)
	}
	return int32(0)
}

// C documentation
//
//	/* Render an array to text */
func print_array(tls *libc.TLS, item uintptr, output_buffer uintptr) (r cJSON_bool) {
	var current_element, output_pointer, v2 uintptr
	var length size_t
	var v1 int32
	_, _, _, _, _ = current_element, length, output_pointer, v1, v2
	output_pointer = libc.UintptrFromInt32(0)
	length = uint64(0)
	current_element = (*cJSON)(unsafe.Pointer(item)).Fchild
	if output_buffer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* Compose the output array. */
	/* opening square bracket */
	output_pointer = ensure(tls, output_buffer, uint64(1))
	if output_pointer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	**(**uint8)(__ccgo_up(output_pointer)) = uint8('[')
	(*printbuffer)(unsafe.Pointer(output_buffer)).Foffset = (*printbuffer)(unsafe.Pointer(output_buffer)).Foffset + 1
	(*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth = (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth + 1
	for current_element != libc.UintptrFromInt32(0) {
		if !(print_value(tls, current_element, output_buffer) != 0) {
			return int32(0)
		}
		update_offset(tls, output_buffer)
		if (*cJSON)(unsafe.Pointer(current_element)).Fnext != 0 {
			if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
				v1 = int32(2)
			} else {
				v1 = int32(1)
			}
			length = uint64(v1)
			output_pointer = ensure(tls, output_buffer, length+uint64(1))
			if output_pointer == libc.UintptrFromInt32(0) {
				return int32(0)
			}
			v2 = output_pointer
			output_pointer = output_pointer + 1
			**(**uint8)(__ccgo_up(v2)) = uint8(',')
			if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
				v2 = output_pointer
				output_pointer = output_pointer + 1
				**(**uint8)(__ccgo_up(v2)) = uint8(' ')
			}
			**(**uint8)(__ccgo_up(output_pointer)) = uint8('\000')
			**(**size_t)(__ccgo_up(output_buffer + 16)) += length
		}
		current_element = (*cJSON)(unsafe.Pointer(current_element)).Fnext
	}
	output_pointer = ensure(tls, output_buffer, uint64(2))
	if output_pointer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	v2 = output_pointer
	output_pointer = output_pointer + 1
	**(**uint8)(__ccgo_up(v2)) = uint8(']')
	**(**uint8)(__ccgo_up(output_pointer)) = uint8('\000')
	(*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth = (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth - 1
	return int32(1)
}

// C documentation
//
//	/* Build an object from the text. */
func parse_object(tls *libc.TLS, item uintptr, input_buffer uintptr) (r cJSON_bool) {
	var current_item, head, new_item, v1 uintptr
	_, _, _, _ = current_item, head, new_item, v1
	head = libc.UintptrFromInt32(0) /* linked list head */
	current_item = libc.UintptrFromInt32(0)
	if (*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth >= uint64(CJSON_NESTING_LIMIT) {
		return int32(0) /* to deeply nested */
	}
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth = (*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth + 1
	if !(input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength) || libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) != int32('{') {
		goto fail /* not an object */
	}
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
	buffer_skip_whitespace(input_buffer)
	if input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32('}') {
		goto success /* empty object */
	}
	/* check if we skipped to the end of the buffer */
	if !(input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength) {
		(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset - 1
		goto fail
	}
	/* step back to character in front of the first element */
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset - 1
	/* loop through the comma separated array elements */
	for cond := true; cond; cond = input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength && libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) == int32(',') {
		/* allocate next item */
		new_item = cJSON_New_Item(tls, input_buffer+32)
		if new_item == libc.UintptrFromInt32(0) {
			goto fail /* allocation failure */
		}
		/* attach next item to list */
		if head == libc.UintptrFromInt32(0) {
			/* start the linked list */
			v1 = new_item
			head = v1
			current_item = v1
		} else {
			/* add to the end and advance */
			(*cJSON)(unsafe.Pointer(current_item)).Fnext = new_item
			(*cJSON)(unsafe.Pointer(new_item)).Fprev = current_item
			current_item = new_item
		}
		if !(input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(1) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength) {
			goto fail /* nothing comes after the comma */
		}
		/* parse the name of the child */
		(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
		buffer_skip_whitespace(input_buffer)
		if !(parse_string(tls, current_item, input_buffer) != 0) {
			goto fail /* failed to parse name */
		}
		buffer_skip_whitespace(input_buffer)
		/* swap valuestring and string, because we parsed the name */
		(*cJSON)(unsafe.Pointer(current_item)).Fstring1 = (*cJSON)(unsafe.Pointer(current_item)).Fvaluestring
		(*cJSON)(unsafe.Pointer(current_item)).Fvaluestring = libc.UintptrFromInt32(0)
		if !(input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength) || libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) != int32(':') {
			goto fail /* invalid object */
		}
		/* parse the value */
		(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
		buffer_skip_whitespace(input_buffer)
		if !(parse_value(tls, current_item, input_buffer) != 0) {
			goto fail /* failed to parse value */
		}
		buffer_skip_whitespace(input_buffer)
	}
	if !(input_buffer != libc.UintptrFromInt32(0) && (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset+uint64(0) < (*parse_buffer)(unsafe.Pointer(input_buffer)).Flength) || libc.Int32FromUint8(**(**uint8)(__ccgo_up((*parse_buffer)(unsafe.Pointer(input_buffer)).Fcontent + uintptr((*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset)))) != int32('}') {
		goto fail /* expected end of object */
	}
	goto success
success:
	;
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth = (*parse_buffer)(unsafe.Pointer(input_buffer)).Fdepth - 1
	if head != libc.UintptrFromInt32(0) {
		(*cJSON)(unsafe.Pointer(head)).Fprev = current_item
	}
	(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(6)
	(*cJSON)(unsafe.Pointer(item)).Fchild = head
	(*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset = (*parse_buffer)(unsafe.Pointer(input_buffer)).Foffset + 1
	return int32(1)
	goto fail
fail:
	;
	if head != libc.UintptrFromInt32(0) {
		cJSON_Delete(tls, head)
	}
	return int32(0)
}

// C documentation
//
//	/* Render an object to text. */
func print_object(tls *libc.TLS, item uintptr, output_buffer uintptr) (r cJSON_bool) {
	var current_item, output_pointer, v2 uintptr
	var i, i1, length size_t
	var v1, v6 int32
	var v13 uint64
	_, _, _, _, _, _, _, _, _ = current_item, i, i1, length, output_pointer, v1, v13, v2, v6
	output_pointer = libc.UintptrFromInt32(0)
	length = uint64(0)
	current_item = (*cJSON)(unsafe.Pointer(item)).Fchild
	if output_buffer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* Compose the output: */
	if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
		v1 = int32(2)
	} else {
		v1 = int32(1)
	}
	length = uint64(v1) /* fmt: {\n */
	output_pointer = ensure(tls, output_buffer, length+uint64(1))
	if output_pointer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	v2 = output_pointer
	output_pointer = output_pointer + 1
	**(**uint8)(__ccgo_up(v2)) = uint8('{')
	(*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth = (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth + 1
	if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
		v2 = output_pointer
		output_pointer = output_pointer + 1
		**(**uint8)(__ccgo_up(v2)) = uint8('\n')
	}
	**(**size_t)(__ccgo_up(output_buffer + 16)) += length
	for current_item != 0 {
		if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
			output_pointer = ensure(tls, output_buffer, (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth)
			if output_pointer == libc.UintptrFromInt32(0) {
				return int32(0)
			}
			i = uint64(0)
			for {
				if !(i < (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth) {
					break
				}
				v2 = output_pointer
				output_pointer = output_pointer + 1
				**(**uint8)(__ccgo_up(v2)) = uint8('\t')
				goto _4
			_4:
				;
				i = i + 1
			}
			**(**size_t)(__ccgo_up(output_buffer + 16)) += (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth
		}
		/* print key */
		if !(print_string_ptr(tls, (*cJSON)(unsafe.Pointer(current_item)).Fstring1, output_buffer) != 0) {
			return int32(0)
		}
		update_offset(tls, output_buffer)
		if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
			v1 = int32(2)
		} else {
			v1 = int32(1)
		}
		length = uint64(v1)
		output_pointer = ensure(tls, output_buffer, length)
		if output_pointer == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		v2 = output_pointer
		output_pointer = output_pointer + 1
		**(**uint8)(__ccgo_up(v2)) = uint8(':')
		if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
			v2 = output_pointer
			output_pointer = output_pointer + 1
			**(**uint8)(__ccgo_up(v2)) = uint8('\t')
		}
		**(**size_t)(__ccgo_up(output_buffer + 16)) += length
		/* print value */
		if !(print_value(tls, current_item, output_buffer) != 0) {
			return int32(0)
		}
		update_offset(tls, output_buffer)
		/* print comma if not last */
		if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
			v1 = int32(1)
		} else {
			v1 = 0
		}
		if (*cJSON)(unsafe.Pointer(current_item)).Fnext != 0 {
			v6 = int32(1)
		} else {
			v6 = 0
		}
		length = uint64(v1) + uint64(v6)
		output_pointer = ensure(tls, output_buffer, length+uint64(1))
		if output_pointer == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		if (*cJSON)(unsafe.Pointer(current_item)).Fnext != 0 {
			v2 = output_pointer
			output_pointer = output_pointer + 1
			**(**uint8)(__ccgo_up(v2)) = uint8(',')
		}
		if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
			v2 = output_pointer
			output_pointer = output_pointer + 1
			**(**uint8)(__ccgo_up(v2)) = uint8('\n')
		}
		**(**uint8)(__ccgo_up(output_pointer)) = uint8('\000')
		**(**size_t)(__ccgo_up(output_buffer + 16)) += length
		current_item = (*cJSON)(unsafe.Pointer(current_item)).Fnext
	}
	if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
		v13 = (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth + uint64(1)
	} else {
		v13 = uint64(2)
	}
	output_pointer = ensure(tls, output_buffer, v13)
	if output_pointer == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	if (*printbuffer)(unsafe.Pointer(output_buffer)).Fformat != 0 {
		i1 = uint64(0)
		for {
			if !(i1 < (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth-uint64(1)) {
				break
			}
			v2 = output_pointer
			output_pointer = output_pointer + 1
			**(**uint8)(__ccgo_up(v2)) = uint8('\t')
			goto _14
		_14:
			;
			i1 = i1 + 1
		}
	}
	v2 = output_pointer
	output_pointer = output_pointer + 1
	**(**uint8)(__ccgo_up(v2)) = uint8('}')
	**(**uint8)(__ccgo_up(output_pointer)) = uint8('\000')
	(*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth = (*printbuffer)(unsafe.Pointer(output_buffer)).Fdepth - 1
	return int32(1)
}

// C documentation
//
//	/* Get Array size/item / object item. */
func cJSON_GetArraySize(array uintptr) (r int32) {
	var child uintptr
	var size size_t
	_, _ = child, size
	child = libc.UintptrFromInt32(0)
	size = uint64(0)
	if array == libc.UintptrFromInt32(0) {
		return 0
	}
	child = (*cJSON)(unsafe.Pointer(array)).Fchild
	for child != libc.UintptrFromInt32(0) {
		size = size + 1
		child = (*cJSON)(unsafe.Pointer(child)).Fnext
	}
	/* FIXME: Can overflow here. Cannot be fixed without breaking the API */
	return libc.Int32FromUint64(size)
}

func get_array_item(array uintptr, index size_t) (r uintptr) {
	var current_child uintptr
	_ = current_child
	current_child = libc.UintptrFromInt32(0)
	if array == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	current_child = (*cJSON)(unsafe.Pointer(array)).Fchild
	for current_child != libc.UintptrFromInt32(0) && index > uint64(0) {
		index = index - 1
		current_child = (*cJSON)(unsafe.Pointer(current_child)).Fnext
	}
	return current_child
}

func cJSON_GetArrayItem(array uintptr, index int32) (r uintptr) {
	if index < 0 {
		return libc.UintptrFromInt32(0)
	}
	return get_array_item(array, uint64(index))
}

func get_object_item(tls *libc.TLS, object uintptr, name uintptr, case_sensitive cJSON_bool) (r uintptr) {
	var current_element uintptr
	_ = current_element
	current_element = libc.UintptrFromInt32(0)
	if object == libc.UintptrFromInt32(0) || name == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	current_element = (*cJSON)(unsafe.Pointer(object)).Fchild
	if case_sensitive != 0 {
		for current_element != libc.UintptrFromInt32(0) && (*cJSON)(unsafe.Pointer(current_element)).Fstring1 != libc.UintptrFromInt32(0) && libc.Xstrcmp(tls, name, (*cJSON)(unsafe.Pointer(current_element)).Fstring1) != 0 {
			current_element = (*cJSON)(unsafe.Pointer(current_element)).Fnext
		}
	} else {
		for current_element != libc.UintptrFromInt32(0) && case_insensitive_strcmp(tls, name, (*cJSON)(unsafe.Pointer(current_element)).Fstring1) != 0 {
			current_element = (*cJSON)(unsafe.Pointer(current_element)).Fnext
		}
	}
	if current_element == libc.UintptrFromInt32(0) || (*cJSON)(unsafe.Pointer(current_element)).Fstring1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	return current_element
}

func cJSON_GetObjectItem(tls *libc.TLS, object uintptr, string1 uintptr) (r uintptr) {
	return get_object_item(tls, object, string1, int32(0))
}

func cJSON_GetObjectItemCaseSensitive(tls *libc.TLS, object uintptr, string1 uintptr) (r uintptr) {
	return get_object_item(tls, object, string1, int32(1))
}

func cJSON_HasObjectItem(tls *libc.TLS, object uintptr, string1 uintptr) (r cJSON_bool) {
	var v1 int32
	_ = v1
	if cJSON_GetObjectItem(tls, object, string1) != 0 {
		v1 = int32(1)
	} else {
		v1 = 0
	}
	return v1
}

// C documentation
//
//	/* Utility for array list handling. */
func suffix_object(prev uintptr, item uintptr) {
	(*cJSON)(unsafe.Pointer(prev)).Fnext = item
	(*cJSON)(unsafe.Pointer(item)).Fprev = prev
}

// C documentation
//
//	/* Utility for handling references. */
func create_reference(tls *libc.TLS, item uintptr, hooks uintptr) (r uintptr) {
	var reference, v1 uintptr
	_, _ = reference, v1
	reference = libc.UintptrFromInt32(0)
	if item == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	reference = cJSON_New_Item(tls, hooks)
	if reference == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	libc.Xmemcpy(tls, reference, item, uint64(64))
	(*cJSON)(unsafe.Pointer(reference)).Fstring1 = libc.UintptrFromInt32(0)
	**(**int32)(__ccgo_up(reference + 24)) |= int32(cJSON_IsReference)
	v1 = libc.UintptrFromInt32(0)
	(*cJSON)(unsafe.Pointer(reference)).Fprev = v1
	(*cJSON)(unsafe.Pointer(reference)).Fnext = v1
	return reference
}

func add_item_to_array(array uintptr, item uintptr) (r cJSON_bool) {
	var child uintptr
	_ = child
	child = libc.UintptrFromInt32(0)
	if item == libc.UintptrFromInt32(0) || array == libc.UintptrFromInt32(0) || array == item {
		return int32(0)
	}
	child = (*cJSON)(unsafe.Pointer(array)).Fchild
	/*
	 * To find the last item in array quickly, we use prev in array
	 */
	if child == libc.UintptrFromInt32(0) {
		/* list is empty, start new one */
		(*cJSON)(unsafe.Pointer(array)).Fchild = item
		(*cJSON)(unsafe.Pointer(item)).Fprev = item
		(*cJSON)(unsafe.Pointer(item)).Fnext = libc.UintptrFromInt32(0)
	} else {
		/* append to the end */
		if (*cJSON)(unsafe.Pointer(child)).Fprev != 0 {
			suffix_object((*cJSON)(unsafe.Pointer(child)).Fprev, item)
			(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(array)).Fchild)).Fprev = item
		}
	}
	return int32(1)
}

// C documentation
//
//	/* Add item to array/object. */
func cJSON_AddItemToArray(array uintptr, item uintptr) (r cJSON_bool) {
	return add_item_to_array(array, item)
}

// C documentation
//
//	/* helper function to cast away const */
func cast_away_const(string1 uintptr) (r uintptr) {
	return string1
}

func add_item_to_object(tls *libc.TLS, object uintptr, string1 uintptr, item uintptr, hooks uintptr, constant_key cJSON_bool) (r cJSON_bool) {
	var new_key uintptr
	var new_type int32
	_, _ = new_key, new_type
	new_key = libc.UintptrFromInt32(0)
	new_type = cJSON_Invalid
	if object == libc.UintptrFromInt32(0) || string1 == libc.UintptrFromInt32(0) || item == libc.UintptrFromInt32(0) || object == item {
		return int32(0)
	}
	if constant_key != 0 {
		new_key = cast_away_const(string1)
		new_type = (*cJSON)(unsafe.Pointer(item)).Ftype1 | int32(cJSON_StringIsConst)
	} else {
		new_key = cJSON_strdup(tls, string1, hooks)
		if new_key == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		new_type = (*cJSON)(unsafe.Pointer(item)).Ftype1 & ^int32(cJSON_StringIsConst)
	}
	if !((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(cJSON_StringIsConst) != 0) && (*cJSON)(unsafe.Pointer(item)).Fstring1 != libc.UintptrFromInt32(0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*internal_hooks)(unsafe.Pointer(hooks)).Fdeallocate})))(tls, (*cJSON)(unsafe.Pointer(item)).Fstring1)
	}
	(*cJSON)(unsafe.Pointer(item)).Fstring1 = new_key
	(*cJSON)(unsafe.Pointer(item)).Ftype1 = new_type
	return add_item_to_array(object, item)
}

func cJSON_AddItemToObject(tls *libc.TLS, object uintptr, string1 uintptr, item uintptr) (r cJSON_bool) {
	return add_item_to_object(tls, object, string1, item, uintptr(unsafe.Pointer(&global_hooks)), int32(0))
}

// C documentation
//
//	/* Add an item to an object with constant string as key */
func cJSON_AddItemToObjectCS(tls *libc.TLS, object uintptr, string1 uintptr, item uintptr) (r cJSON_bool) {
	return add_item_to_object(tls, object, string1, item, uintptr(unsafe.Pointer(&global_hooks)), int32(1))
}

func cJSON_AddItemReferenceToArray(tls *libc.TLS, array uintptr, item uintptr) (r cJSON_bool) {
	if array == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return add_item_to_array(array, create_reference(tls, item, uintptr(unsafe.Pointer(&global_hooks))))
}

func cJSON_AddItemReferenceToObject(tls *libc.TLS, object uintptr, string1 uintptr, item uintptr) (r cJSON_bool) {
	if object == libc.UintptrFromInt32(0) || string1 == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return add_item_to_object(tls, object, string1, create_reference(tls, item, uintptr(unsafe.Pointer(&global_hooks))), uintptr(unsafe.Pointer(&global_hooks)), int32(0))
}

func cJSON_AddNullToObject(tls *libc.TLS, object uintptr, name uintptr) (r uintptr) {
	var null uintptr
	_ = null
	null = cJSON_CreateNull(tls)
	if add_item_to_object(tls, object, name, null, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return null
	}
	cJSON_Delete(tls, null)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddTrueToObject(tls *libc.TLS, object uintptr, name uintptr) (r uintptr) {
	var true_item uintptr
	_ = true_item
	true_item = cJSON_CreateTrue(tls)
	if add_item_to_object(tls, object, name, true_item, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return true_item
	}
	cJSON_Delete(tls, true_item)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddFalseToObject(tls *libc.TLS, object uintptr, name uintptr) (r uintptr) {
	var false_item uintptr
	_ = false_item
	false_item = cJSON_CreateFalse(tls)
	if add_item_to_object(tls, object, name, false_item, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return false_item
	}
	cJSON_Delete(tls, false_item)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddBoolToObject(tls *libc.TLS, object uintptr, name uintptr, boolean cJSON_bool) (r uintptr) {
	var bool_item uintptr
	_ = bool_item
	bool_item = cJSON_CreateBool(tls, boolean)
	if add_item_to_object(tls, object, name, bool_item, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return bool_item
	}
	cJSON_Delete(tls, bool_item)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddNumberToObject(tls *libc.TLS, object uintptr, name uintptr, number float64) (r uintptr) {
	var number_item uintptr
	_ = number_item
	number_item = cJSON_CreateNumber(tls, number)
	if add_item_to_object(tls, object, name, number_item, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return number_item
	}
	cJSON_Delete(tls, number_item)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddStringToObject(tls *libc.TLS, object uintptr, name uintptr, string1 uintptr) (r uintptr) {
	var string_item uintptr
	_ = string_item
	string_item = cJSON_CreateString(tls, string1)
	if add_item_to_object(tls, object, name, string_item, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return string_item
	}
	cJSON_Delete(tls, string_item)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddRawToObject(tls *libc.TLS, object uintptr, name uintptr, raw uintptr) (r uintptr) {
	var raw_item uintptr
	_ = raw_item
	raw_item = cJSON_CreateRaw(tls, raw)
	if add_item_to_object(tls, object, name, raw_item, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return raw_item
	}
	cJSON_Delete(tls, raw_item)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddObjectToObject(tls *libc.TLS, object uintptr, name uintptr) (r uintptr) {
	var object_item uintptr
	_ = object_item
	object_item = cJSON_CreateObject(tls)
	if add_item_to_object(tls, object, name, object_item, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return object_item
	}
	cJSON_Delete(tls, object_item)
	return libc.UintptrFromInt32(0)
}

func cJSON_AddArrayToObject(tls *libc.TLS, object uintptr, name uintptr) (r uintptr) {
	var array uintptr
	_ = array
	array = cJSON_CreateArray(tls)
	if add_item_to_object(tls, object, name, array, uintptr(unsafe.Pointer(&global_hooks)), int32(0)) != 0 {
		return array
	}
	cJSON_Delete(tls, array)
	return libc.UintptrFromInt32(0)
}

func cJSON_DetachItemViaPointer(parent uintptr, item uintptr) (r uintptr) {
	if parent == libc.UintptrFromInt32(0) || item == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if item != (*cJSON)(unsafe.Pointer(parent)).Fchild {
		/* not the first element */
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(item)).Fprev)).Fnext = (*cJSON)(unsafe.Pointer(item)).Fnext
	}
	if (*cJSON)(unsafe.Pointer(item)).Fnext != libc.UintptrFromInt32(0) {
		/* not the last element */
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(item)).Fnext)).Fprev = (*cJSON)(unsafe.Pointer(item)).Fprev
	}
	if item == (*cJSON)(unsafe.Pointer(parent)).Fchild {
		/* first element */
		(*cJSON)(unsafe.Pointer(parent)).Fchild = (*cJSON)(unsafe.Pointer(item)).Fnext
	} else {
		if (*cJSON)(unsafe.Pointer(item)).Fnext == libc.UintptrFromInt32(0) {
			/* last element */
			(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(parent)).Fchild)).Fprev = (*cJSON)(unsafe.Pointer(item)).Fprev
		}
	}
	/* make sure the detached item doesn't point anywhere anymore */
	(*cJSON)(unsafe.Pointer(item)).Fprev = libc.UintptrFromInt32(0)
	(*cJSON)(unsafe.Pointer(item)).Fnext = libc.UintptrFromInt32(0)
	return item
}

func cJSON_DetachItemFromArray(array uintptr, which int32) (r uintptr) {
	if which < 0 {
		return libc.UintptrFromInt32(0)
	}
	return cJSON_DetachItemViaPointer(array, get_array_item(array, uint64(which)))
}

func cJSON_DeleteItemFromArray(tls *libc.TLS, array uintptr, which int32) {
	cJSON_Delete(tls, cJSON_DetachItemFromArray(array, which))
}

func cJSON_DetachItemFromObject(tls *libc.TLS, object uintptr, string1 uintptr) (r uintptr) {
	var to_detach uintptr
	_ = to_detach
	to_detach = cJSON_GetObjectItem(tls, object, string1)
	return cJSON_DetachItemViaPointer(object, to_detach)
}

func cJSON_DetachItemFromObjectCaseSensitive(tls *libc.TLS, object uintptr, string1 uintptr) (r uintptr) {
	var to_detach uintptr
	_ = to_detach
	to_detach = cJSON_GetObjectItemCaseSensitive(tls, object, string1)
	return cJSON_DetachItemViaPointer(object, to_detach)
}

func cJSON_DeleteItemFromObject(tls *libc.TLS, object uintptr, string1 uintptr) {
	cJSON_Delete(tls, cJSON_DetachItemFromObject(tls, object, string1))
}

func cJSON_DeleteItemFromObjectCaseSensitive(tls *libc.TLS, object uintptr, string1 uintptr) {
	cJSON_Delete(tls, cJSON_DetachItemFromObjectCaseSensitive(tls, object, string1))
}

// C documentation
//
//	/* Replace array/object items with new ones. */
func cJSON_InsertItemInArray(array uintptr, which int32, newitem uintptr) (r cJSON_bool) {
	var after_inserted uintptr
	_ = after_inserted
	after_inserted = libc.UintptrFromInt32(0)
	if which < 0 || newitem == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	after_inserted = get_array_item(array, uint64(which))
	if after_inserted == libc.UintptrFromInt32(0) {
		return add_item_to_array(array, newitem)
	}
	if after_inserted != (*cJSON)(unsafe.Pointer(array)).Fchild && (*cJSON)(unsafe.Pointer(after_inserted)).Fprev == libc.UintptrFromInt32(0) {
		/* return false if after_inserted is a corrupted array item */
		return int32(0)
	}
	(*cJSON)(unsafe.Pointer(newitem)).Fnext = after_inserted
	(*cJSON)(unsafe.Pointer(newitem)).Fprev = (*cJSON)(unsafe.Pointer(after_inserted)).Fprev
	(*cJSON)(unsafe.Pointer(after_inserted)).Fprev = newitem
	if after_inserted == (*cJSON)(unsafe.Pointer(array)).Fchild {
		(*cJSON)(unsafe.Pointer(array)).Fchild = newitem
	} else {
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(newitem)).Fprev)).Fnext = newitem
	}
	return int32(1)
}

func cJSON_ReplaceItemViaPointer(tls *libc.TLS, parent uintptr, item uintptr, replacement uintptr) (r cJSON_bool) {
	if parent == libc.UintptrFromInt32(0) || (*cJSON)(unsafe.Pointer(parent)).Fchild == libc.UintptrFromInt32(0) || replacement == libc.UintptrFromInt32(0) || item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	if replacement == item {
		return int32(1)
	}
	(*cJSON)(unsafe.Pointer(replacement)).Fnext = (*cJSON)(unsafe.Pointer(item)).Fnext
	(*cJSON)(unsafe.Pointer(replacement)).Fprev = (*cJSON)(unsafe.Pointer(item)).Fprev
	if (*cJSON)(unsafe.Pointer(replacement)).Fnext != libc.UintptrFromInt32(0) {
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(replacement)).Fnext)).Fprev = replacement
	}
	if (*cJSON)(unsafe.Pointer(parent)).Fchild == item {
		if (*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(parent)).Fchild)).Fprev == (*cJSON)(unsafe.Pointer(parent)).Fchild {
			(*cJSON)(unsafe.Pointer(replacement)).Fprev = replacement
		}
		(*cJSON)(unsafe.Pointer(parent)).Fchild = replacement
	} else {
		/*
		 * To find the last item in array quickly, we use prev in array.
		 * We can't modify the last item's next pointer where this item was the parent's child
		 */
		if (*cJSON)(unsafe.Pointer(replacement)).Fprev != libc.UintptrFromInt32(0) {
			(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(replacement)).Fprev)).Fnext = replacement
		}
		if (*cJSON)(unsafe.Pointer(replacement)).Fnext == libc.UintptrFromInt32(0) {
			(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(parent)).Fchild)).Fprev = replacement
		}
	}
	(*cJSON)(unsafe.Pointer(item)).Fnext = libc.UintptrFromInt32(0)
	(*cJSON)(unsafe.Pointer(item)).Fprev = libc.UintptrFromInt32(0)
	cJSON_Delete(tls, item)
	return int32(1)
}

func cJSON_ReplaceItemInArray(tls *libc.TLS, array uintptr, which int32, newitem uintptr) (r cJSON_bool) {
	if which < 0 {
		return int32(0)
	}
	return cJSON_ReplaceItemViaPointer(tls, array, get_array_item(array, uint64(which)), newitem)
}

func replace_item_in_object(tls *libc.TLS, object uintptr, string1 uintptr, replacement uintptr, case_sensitive cJSON_bool) (r cJSON_bool) {
	if replacement == libc.UintptrFromInt32(0) || string1 == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	/* replace the name in the replacement */
	if !((*cJSON)(unsafe.Pointer(replacement)).Ftype1&int32(cJSON_StringIsConst) != 0) && (*cJSON)(unsafe.Pointer(replacement)).Fstring1 != libc.UintptrFromInt32(0) {
		cJSON_free(tls, (*cJSON)(unsafe.Pointer(replacement)).Fstring1)
	}
	(*cJSON)(unsafe.Pointer(replacement)).Fstring1 = cJSON_strdup(tls, string1, uintptr(unsafe.Pointer(&global_hooks)))
	if (*cJSON)(unsafe.Pointer(replacement)).Fstring1 == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	**(**int32)(__ccgo_up(replacement + 24)) &= ^int32(cJSON_StringIsConst)
	return cJSON_ReplaceItemViaPointer(tls, object, get_object_item(tls, object, string1, case_sensitive), replacement)
}

func cJSON_ReplaceItemInObject(tls *libc.TLS, object uintptr, string1 uintptr, newitem uintptr) (r cJSON_bool) {
	return replace_item_in_object(tls, object, string1, newitem, int32(0))
}

func cJSON_ReplaceItemInObjectCaseSensitive(tls *libc.TLS, object uintptr, string1 uintptr, newitem uintptr) (r cJSON_bool) {
	return replace_item_in_object(tls, object, string1, newitem, int32(1))
}

// C documentation
//
//	/* Create basic types: */
func cJSON_CreateNull(tls *libc.TLS) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(2)
	}
	return item
}

func cJSON_CreateTrue(tls *libc.TLS) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(1)
	}
	return item
}

func cJSON_CreateFalse(tls *libc.TLS) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(0)
	}
	return item
}

func cJSON_CreateBool(tls *libc.TLS, boolean cJSON_bool) (r uintptr) {
	var item uintptr
	var v1 int32
	_, _ = item, v1
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		if boolean != 0 {
			v1 = int32(1) << int32(1)
		} else {
			v1 = int32(1) << int32(0)
		}
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = v1
	}
	return item
}

func cJSON_CreateNumber(tls *libc.TLS, num float64) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(3)
		(*cJSON)(unsafe.Pointer(item)).Fvaluedouble = num
		/* use saturation in case of overflow */
		if num >= libc.Float64FromInt32(INT_MAX) {
			(*cJSON)(unsafe.Pointer(item)).Fvalueint = int32(INT_MAX)
		} else {
			if num <= float64(-int32(1)-int32(0x7fffffff)) {
				(*cJSON)(unsafe.Pointer(item)).Fvalueint = -int32(1) - int32(0x7fffffff)
			} else {
				(*cJSON)(unsafe.Pointer(item)).Fvalueint = int32(num)
			}
		}
	}
	return item
}

func cJSON_CreateString(tls *libc.TLS, string1 uintptr) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(4)
		(*cJSON)(unsafe.Pointer(item)).Fvaluestring = cJSON_strdup(tls, string1, uintptr(unsafe.Pointer(&global_hooks)))
		if !((*cJSON)(unsafe.Pointer(item)).Fvaluestring != 0) {
			cJSON_Delete(tls, item)
			return libc.UintptrFromInt32(0)
		}
	}
	return item
}

func cJSON_CreateStringReference(tls *libc.TLS, string1 uintptr) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != libc.UintptrFromInt32(0) {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1)<<int32(4) | int32(cJSON_IsReference)
		(*cJSON)(unsafe.Pointer(item)).Fvaluestring = cast_away_const(string1)
	}
	return item
}

func cJSON_CreateObjectReference(tls *libc.TLS, child uintptr) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != libc.UintptrFromInt32(0) {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1)<<int32(6) | int32(cJSON_IsReference)
		(*cJSON)(unsafe.Pointer(item)).Fchild = cast_away_const(child)
	}
	return item
}

func cJSON_CreateArrayReference(tls *libc.TLS, child uintptr) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != libc.UintptrFromInt32(0) {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1)<<int32(5) | int32(cJSON_IsReference)
		(*cJSON)(unsafe.Pointer(item)).Fchild = cast_away_const(child)
	}
	return item
}

func cJSON_CreateRaw(tls *libc.TLS, raw uintptr) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(7)
		(*cJSON)(unsafe.Pointer(item)).Fvaluestring = cJSON_strdup(tls, raw, uintptr(unsafe.Pointer(&global_hooks)))
		if !((*cJSON)(unsafe.Pointer(item)).Fvaluestring != 0) {
			cJSON_Delete(tls, item)
			return libc.UintptrFromInt32(0)
		}
	}
	return item
}

func cJSON_CreateArray(tls *libc.TLS) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(5)
	}
	return item
}

func cJSON_CreateObject(tls *libc.TLS) (r uintptr) {
	var item uintptr
	_ = item
	item = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if item != 0 {
		(*cJSON)(unsafe.Pointer(item)).Ftype1 = int32(1) << int32(6)
	}
	return item
}

// C documentation
//
//	/* Create Arrays: */
func cJSON_CreateIntArray(tls *libc.TLS, numbers uintptr, count int32) (r uintptr) {
	var a, n, p uintptr
	var i size_t
	_, _, _, _ = a, i, n, p
	i = uint64(0)
	n = libc.UintptrFromInt32(0)
	p = libc.UintptrFromInt32(0)
	a = libc.UintptrFromInt32(0)
	if count < 0 || numbers == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	a = cJSON_CreateArray(tls)
	i = uint64(0)
	for {
		if !(a != 0 && i < uint64(count)) {
			break
		}
		n = cJSON_CreateNumber(tls, float64(**(**int32)(__ccgo_up(numbers + uintptr(i)*4))))
		if !(n != 0) {
			cJSON_Delete(tls, a)
			return libc.UintptrFromInt32(0)
		}
		if !(i != 0) {
			(*cJSON)(unsafe.Pointer(a)).Fchild = n
		} else {
			suffix_object(p, n)
		}
		p = n
		goto _1
	_1:
		;
		i = i + 1
	}
	if a != 0 && (*cJSON)(unsafe.Pointer(a)).Fchild != 0 {
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(a)).Fchild)).Fprev = n
	}
	return a
}

func cJSON_CreateFloatArray(tls *libc.TLS, numbers uintptr, count int32) (r uintptr) {
	var a, n, p uintptr
	var i size_t
	_, _, _, _ = a, i, n, p
	i = uint64(0)
	n = libc.UintptrFromInt32(0)
	p = libc.UintptrFromInt32(0)
	a = libc.UintptrFromInt32(0)
	if count < 0 || numbers == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	a = cJSON_CreateArray(tls)
	i = uint64(0)
	for {
		if !(a != 0 && i < uint64(count)) {
			break
		}
		n = cJSON_CreateNumber(tls, float64(**(**float32)(__ccgo_up(numbers + uintptr(i)*4))))
		if !(n != 0) {
			cJSON_Delete(tls, a)
			return libc.UintptrFromInt32(0)
		}
		if !(i != 0) {
			(*cJSON)(unsafe.Pointer(a)).Fchild = n
		} else {
			suffix_object(p, n)
		}
		p = n
		goto _1
	_1:
		;
		i = i + 1
	}
	if a != 0 && (*cJSON)(unsafe.Pointer(a)).Fchild != 0 {
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(a)).Fchild)).Fprev = n
	}
	return a
}

func cJSON_CreateDoubleArray(tls *libc.TLS, numbers uintptr, count int32) (r uintptr) {
	var a, n, p uintptr
	var i size_t
	_, _, _, _ = a, i, n, p
	i = uint64(0)
	n = libc.UintptrFromInt32(0)
	p = libc.UintptrFromInt32(0)
	a = libc.UintptrFromInt32(0)
	if count < 0 || numbers == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	a = cJSON_CreateArray(tls)
	i = uint64(0)
	for {
		if !(a != 0 && i < uint64(count)) {
			break
		}
		n = cJSON_CreateNumber(tls, **(**float64)(__ccgo_up(numbers + uintptr(i)*8)))
		if !(n != 0) {
			cJSON_Delete(tls, a)
			return libc.UintptrFromInt32(0)
		}
		if !(i != 0) {
			(*cJSON)(unsafe.Pointer(a)).Fchild = n
		} else {
			suffix_object(p, n)
		}
		p = n
		goto _1
	_1:
		;
		i = i + 1
	}
	if a != 0 && (*cJSON)(unsafe.Pointer(a)).Fchild != 0 {
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(a)).Fchild)).Fprev = n
	}
	return a
}

func cJSON_CreateStringArray(tls *libc.TLS, strings uintptr, count int32) (r uintptr) {
	var a, n, p uintptr
	var i size_t
	_, _, _, _ = a, i, n, p
	i = uint64(0)
	n = libc.UintptrFromInt32(0)
	p = libc.UintptrFromInt32(0)
	a = libc.UintptrFromInt32(0)
	if count < 0 || strings == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	a = cJSON_CreateArray(tls)
	i = uint64(0)
	for {
		if !(a != 0 && i < uint64(count)) {
			break
		}
		n = cJSON_CreateString(tls, **(**uintptr)(__ccgo_up(strings + uintptr(i)*8)))
		if !(n != 0) {
			cJSON_Delete(tls, a)
			return libc.UintptrFromInt32(0)
		}
		if !(i != 0) {
			(*cJSON)(unsafe.Pointer(a)).Fchild = n
		} else {
			suffix_object(p, n)
		}
		p = n
		goto _1
	_1:
		;
		i = i + 1
	}
	if a != 0 && (*cJSON)(unsafe.Pointer(a)).Fchild != 0 {
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(a)).Fchild)).Fprev = n
	}
	return a
}

// C documentation
//
//	/* Duplication */
func cJSON_Duplicate(tls *libc.TLS, item uintptr, recurse cJSON_bool) (r uintptr) {
	var child, newchild, newitem, next, v1 uintptr
	_, _, _, _, _ = child, newchild, newitem, next, v1
	newitem = libc.UintptrFromInt32(0)
	child = libc.UintptrFromInt32(0)
	next = libc.UintptrFromInt32(0)
	newchild = libc.UintptrFromInt32(0)
	/* Bail on bad ptr */
	if !(item != 0) {
		goto fail
	}
	/* Create new item */
	newitem = cJSON_New_Item(tls, uintptr(unsafe.Pointer(&global_hooks)))
	if !(newitem != 0) {
		goto fail
	}
	/* Copy over all vars */
	(*cJSON)(unsafe.Pointer(newitem)).Ftype1 = (*cJSON)(unsafe.Pointer(item)).Ftype1 & ^int32(cJSON_IsReference)
	(*cJSON)(unsafe.Pointer(newitem)).Fvalueint = (*cJSON)(unsafe.Pointer(item)).Fvalueint
	(*cJSON)(unsafe.Pointer(newitem)).Fvaluedouble = (*cJSON)(unsafe.Pointer(item)).Fvaluedouble
	if (*cJSON)(unsafe.Pointer(item)).Fvaluestring != 0 {
		(*cJSON)(unsafe.Pointer(newitem)).Fvaluestring = cJSON_strdup(tls, (*cJSON)(unsafe.Pointer(item)).Fvaluestring, uintptr(unsafe.Pointer(&global_hooks)))
		if !((*cJSON)(unsafe.Pointer(newitem)).Fvaluestring != 0) {
			goto fail
		}
	}
	if (*cJSON)(unsafe.Pointer(item)).Fstring1 != 0 {
		if (*cJSON)(unsafe.Pointer(item)).Ftype1&int32(cJSON_StringIsConst) != 0 {
			v1 = (*cJSON)(unsafe.Pointer(item)).Fstring1
		} else {
			v1 = cJSON_strdup(tls, (*cJSON)(unsafe.Pointer(item)).Fstring1, uintptr(unsafe.Pointer(&global_hooks)))
		}
		(*cJSON)(unsafe.Pointer(newitem)).Fstring1 = v1
		if !((*cJSON)(unsafe.Pointer(newitem)).Fstring1 != 0) {
			goto fail
		}
	}
	/* If non-recursive, then we're done! */
	if !(recurse != 0) {
		return newitem
	}
	/* Walk the ->next chain for the child. */
	child = (*cJSON)(unsafe.Pointer(item)).Fchild
	for child != libc.UintptrFromInt32(0) {
		newchild = cJSON_Duplicate(tls, child, int32(1)) /* Duplicate (with recurse) each item in the ->next chain */
		if !(newchild != 0) {
			goto fail
		}
		if next != libc.UintptrFromInt32(0) {
			/* If newitem->child already set, then crosswire ->prev and ->next and move on */
			(*cJSON)(unsafe.Pointer(next)).Fnext = newchild
			(*cJSON)(unsafe.Pointer(newchild)).Fprev = next
			next = newchild
		} else {
			/* Set newitem->child and move to it */
			(*cJSON)(unsafe.Pointer(newitem)).Fchild = newchild
			next = newchild
		}
		child = (*cJSON)(unsafe.Pointer(child)).Fnext
	}
	if newitem != 0 && (*cJSON)(unsafe.Pointer(newitem)).Fchild != 0 {
		(*cJSON)(unsafe.Pointer((*cJSON)(unsafe.Pointer(newitem)).Fchild)).Fprev = newchild
	}
	return newitem
	goto fail
fail:
	;
	if newitem != libc.UintptrFromInt32(0) {
		cJSON_Delete(tls, newitem)
	}
	return libc.UintptrFromInt32(0)
}

func skip_oneline_comment(input uintptr) {
	**(**uintptr)(__ccgo_up(input)) += uintptr(libc.Uint64FromInt64(3) - libc.Uint64FromInt64(1))
	for {
		if !(int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))) != int32('\000')) {
			break
		}
		if int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))) == int32('\n') {
			**(**uintptr)(__ccgo_up(input)) += uintptr(libc.Uint64FromInt64(2) - libc.Uint64FromInt64(1))
			return
		}
		goto _1
	_1:
		;
		**(**uintptr)(__ccgo_up(input)) = **(**uintptr)(__ccgo_up(input)) + 1
	}
}

func skip_multiline_comment(input uintptr) {
	**(**uintptr)(__ccgo_up(input)) += uintptr(libc.Uint64FromInt64(3) - libc.Uint64FromInt64(1))
	for {
		if !(int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))) != int32('\000')) {
			break
		}
		if int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))) == int32('*') && int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input)) + 1))) == int32('/') {
			**(**uintptr)(__ccgo_up(input)) += uintptr(libc.Uint64FromInt64(3) - libc.Uint64FromInt64(1))
			return
		}
		goto _1
	_1:
		;
		**(**uintptr)(__ccgo_up(input)) = **(**uintptr)(__ccgo_up(input)) + 1
	}
}

func minify_string(input uintptr, output uintptr) {
	**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(output)))) = **(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))
	**(**uintptr)(__ccgo_up(input)) += uintptr(libc.Uint64FromInt64(2) - libc.Uint64FromInt64(1))
	**(**uintptr)(__ccgo_up(output)) += uintptr(libc.Uint64FromInt64(2) - libc.Uint64FromInt64(1))
	for {
		if !(int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))) != int32('\000')) {
			break
		}
		**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(output)))) = **(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))
		if int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))) == int32('"') {
			**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(output)))) = int8('"')
			**(**uintptr)(__ccgo_up(input)) += uintptr(libc.Uint64FromInt64(2) - libc.Uint64FromInt64(1))
			**(**uintptr)(__ccgo_up(output)) += uintptr(libc.Uint64FromInt64(2) - libc.Uint64FromInt64(1))
			return
		} else {
			if int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input))))) == int32('\\') && int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input)) + 1))) == int32('"') {
				**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(output)) + 1)) = **(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(input)) + 1))
				**(**uintptr)(__ccgo_up(input)) += uintptr(libc.Uint64FromInt64(2) - libc.Uint64FromInt64(1))
				**(**uintptr)(__ccgo_up(output)) += uintptr(libc.Uint64FromInt64(2) - libc.Uint64FromInt64(1))
			}
		}
		goto _1
	_1:
		;
		**(**uintptr)(__ccgo_up(input)) = **(**uintptr)(__ccgo_up(input)) + 1
		**(**uintptr)(__ccgo_up(output)) = **(**uintptr)(__ccgo_up(output)) + 1
	}
}

func cJSON_Minify(tls *libc.TLS, _json uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uintptr)(unsafe.Pointer(bp)) = _json
	var _ /* into at bp+8 */ uintptr
	**(**uintptr)(__ccgo_up(bp + 8)) = **(**uintptr)(__ccgo_up(bp))
	if **(**uintptr)(__ccgo_up(bp)) == libc.UintptrFromInt32(0) {
		return
	}
	for int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp))))) != int32('\000') {
		switch int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp))))) {
		case int32(' '):
			fallthrough
		case int32('\t'):
			fallthrough
		case int32('\r'):
			fallthrough
		case int32('\n'):
			**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
		case int32('/'):
			if int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + 1))) == int32('/') {
				skip_oneline_comment(bp)
			} else {
				if int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp)) + 1))) == int32('*') {
					skip_multiline_comment(bp)
				} else {
					**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
				}
			}
		case int32('"'):
			minify_string(bp, bp+8)
		default:
			**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)))) = **(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp))))
			**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
			**(**uintptr)(__ccgo_up(bp + 8)) = **(**uintptr)(__ccgo_up(bp + 8)) + 1
		}
	}
	/* and null-terminate. */
	**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(bp + 8)))) = int8('\000')
}

func cJSON_IsInvalid(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == cJSON_Invalid)
}

func cJSON_IsFalse(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == int32(1)<<int32(0))
}

func cJSON_IsTrue(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xff) == int32(1)<<int32(1))
}

func cJSON_IsBool(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&(int32(1)<<int32(1)|int32(1)<<int32(0)) != 0)
}

func cJSON_IsNull(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == int32(1)<<int32(2))
}

func cJSON_IsNumber(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == int32(1)<<int32(3))
}

func cJSON_IsString(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == int32(1)<<int32(4))
}

func cJSON_IsArray(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == int32(1)<<int32(5))
}

func cJSON_IsObject(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == int32(1)<<int32(6))
}

func cJSON_IsRaw(item uintptr) (r cJSON_bool) {
	if item == libc.UintptrFromInt32(0) {
		return int32(0)
	}
	return libc.BoolInt32((*cJSON)(unsafe.Pointer(item)).Ftype1&int32(0xFF) == int32(1)<<int32(7))
}

func cJSON_Compare(tls *libc.TLS, a uintptr, b uintptr, case_sensitive cJSON_bool) (r cJSON_bool) {
	var a_element, a_element1, b_element, b_element1, v3 uintptr
	_, _, _, _, _ = a_element, a_element1, b_element, b_element1, v3
	if a == libc.UintptrFromInt32(0) || b == libc.UintptrFromInt32(0) || (*cJSON)(unsafe.Pointer(a)).Ftype1&int32(0xFF) != (*cJSON)(unsafe.Pointer(b)).Ftype1&int32(0xFF) {
		return int32(0)
	}
	/* check if type is valid */
	switch (*cJSON)(unsafe.Pointer(a)).Ftype1 & int32(0xFF) {
	case int32(1) << int32(0):
		fallthrough
	case int32(1) << int32(1):
		fallthrough
	case int32(1) << int32(2):
		fallthrough
	case int32(1) << int32(3):
		fallthrough
	case int32(1) << int32(4):
		fallthrough
	case int32(1) << int32(7):
		fallthrough
	case int32(1) << int32(5):
		fallthrough
	case int32(1) << int32(6):
	default:
		return int32(0)
	}
	/* identical objects are equal */
	if a == b {
		return int32(1)
	}
	switch (*cJSON)(unsafe.Pointer(a)).Ftype1 & int32(0xFF) {
	/* in these cases and equal type is enough */
	case int32(1) << int32(0):
		fallthrough
	case int32(1) << int32(1):
		fallthrough
	case int32(1) << int32(2):
		return int32(1)
	case int32(1) << int32(3):
		if compare_double(tls, (*cJSON)(unsafe.Pointer(a)).Fvaluedouble, (*cJSON)(unsafe.Pointer(b)).Fvaluedouble) != 0 {
			return int32(1)
		}
		return int32(0)
	case int32(1) << int32(4):
		fallthrough
	case int32(1) << int32(7):
		if (*cJSON)(unsafe.Pointer(a)).Fvaluestring == libc.UintptrFromInt32(0) || (*cJSON)(unsafe.Pointer(b)).Fvaluestring == libc.UintptrFromInt32(0) {
			return int32(0)
		}
		if libc.Xstrcmp(tls, (*cJSON)(unsafe.Pointer(a)).Fvaluestring, (*cJSON)(unsafe.Pointer(b)).Fvaluestring) == 0 {
			return int32(1)
		}
		return int32(0)
	case int32(1) << int32(5):
		a_element = (*cJSON)(unsafe.Pointer(a)).Fchild
		b_element = (*cJSON)(unsafe.Pointer(b)).Fchild
		for {
			if !(a_element != libc.UintptrFromInt32(0) && b_element != libc.UintptrFromInt32(0)) {
				break
			}
			if !(cJSON_Compare(tls, a_element, b_element, case_sensitive) != 0) {
				return int32(0)
			}
			a_element = (*cJSON)(unsafe.Pointer(a_element)).Fnext
			b_element = (*cJSON)(unsafe.Pointer(b_element)).Fnext
			goto _1
		_1:
		}
		/* one of the arrays is longer than the other */
		if a_element != b_element {
			return int32(0)
		}
		return int32(1)
	case int32(1) << int32(6):
		a_element1 = libc.UintptrFromInt32(0)
		b_element1 = libc.UintptrFromInt32(0)
		if a != libc.UintptrFromInt32(0) {
			v3 = (*cJSON)(unsafe.Pointer(a)).Fchild
		} else {
			v3 = libc.UintptrFromInt32(0)
		}
		a_element1 = v3
		for {
			if !(a_element1 != libc.UintptrFromInt32(0)) {
				break
			}
			/* TODO This has O(n^2) runtime, which is horrible! */
			b_element1 = get_object_item(tls, b, (*cJSON)(unsafe.Pointer(a_element1)).Fstring1, case_sensitive)
			if b_element1 == libc.UintptrFromInt32(0) {
				return int32(0)
			}
			if !(cJSON_Compare(tls, a_element1, b_element1, case_sensitive) != 0) {
				return int32(0)
			}
			goto _2
		_2:
			;
			a_element1 = (*cJSON)(unsafe.Pointer(a_element1)).Fnext
		}
		/* doing this twice, once on a and b to prevent true comparison if a subset of b
		 * TODO: Do this the proper way, this is just a fix for now */
		if b != libc.UintptrFromInt32(0) {
			v3 = (*cJSON)(unsafe.Pointer(b)).Fchild
		} else {
			v3 = libc.UintptrFromInt32(0)
		}
		b_element1 = v3
		for {
			if !(b_element1 != libc.UintptrFromInt32(0)) {
				break
			}
			a_element1 = get_object_item(tls, a, (*cJSON)(unsafe.Pointer(b_element1)).Fstring1, case_sensitive)
			if a_element1 == libc.UintptrFromInt32(0) {
				return int32(0)
			}
			if !(cJSON_Compare(tls, b_element1, a_element1, case_sensitive) != 0) {
				return int32(0)
			}
			goto _4
		_4:
			;
			b_element1 = (*cJSON)(unsafe.Pointer(b_element1)).Fnext
		}
		return int32(1)
	default:
		return int32(0)
	}
	return r
}

func cJSON_malloc(tls *libc.TLS, size size_t) (r uintptr) {
	return (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{global_hooks.Fallocate})))(tls, size)
}

func cJSON_free(tls *libc.TLS, object uintptr) {
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{global_hooks.Fdeallocate})))(tls, object)
	object = libc.UintptrFromInt32(0)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "\x00%i.%i.%i\x00null\x00%d\x00%1.15g\x00%lg\x00%1.17g\x00\"\"\x00u%04x\x00\ufeff\x00false\x00true\x00"
