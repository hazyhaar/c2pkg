// Code generated for linux/amd64 by 'ccgo --package-name=df_lz4 -o spec/dogfood/cycles/20260810g/lz4/raw.go -I spec/dogfood/cycles/20260810g/lz4 spec/dogfood/cycles/20260810g/lz4/src.c', DO NOT EDIT.

//go:build linux && amd64

package df_lz4

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
const CHARCLASS_NAME_MAX = 14
const CHAR_BIT = 8
const CHAR_MAX = 255
const CHAR_MIN = 0
const COLL_WEIGHTS_MAX = 2
const DELAYTIMER_MAX = 0x7fffffff
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXPR_NEST_MAX = 32
const FASTLOOP_SAFE_DISTANCE = 64
const FILESIZEBITS = 64
const HOST_NAME_MAX = 255
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
const INT_MAX = 0x7fffffff
const IOV_MAX = 1024
const LASTLITERALS = 5
const LINE_MAX = 4096
const LLONG_MAX = 0x7fffffffffffffff
const LOGIN_NAME_MAX = 256
const LONG_BIT = 64
const LONG_MAX = "__LONG_MAX"
const LZ4LIB_API = "LZ4LIB_VISIBILITY"
const LZ4_ACCELERATION_DEFAULT = 1
const LZ4_ACCELERATION_MAX = 65537
const LZ4_ALIGN_TEST = 1
const LZ4_DISTANCE_ABSOLUTE_MAX = 65535
const LZ4_DISTANCE_MAX = 65535
const LZ4_FAST_DEC_LOOP = 1
const LZ4_FORCE_MEMORY_ACCESS = 1
const LZ4_FREESTANDING = 0
const LZ4_HEAPMODE = 0
const LZ4_MAX_INPUT_SIZE = 2113929216
const LZ4_MEMORY_USAGE = "LZ4_MEMORY_USAGE_DEFAULT"
const LZ4_MEMORY_USAGE_DEFAULT = 14
const LZ4_MEMORY_USAGE_MAX = 20
const LZ4_MEMORY_USAGE_MIN = 10
const LZ4_SRC_INCLUDED = 1
const LZ4_STREAMDECODE_MINSIZE = 32
const LZ4_VERSION_MAJOR = 1
const LZ4_VERSION_MINOR = 9
const LZ4_VERSION_RELEASE = 4
const LZ4_memmove = "__builtin_memmove"
const MB_LEN_MAX = 4
const MFLIMIT = 12
const MINMATCH = 4
const ML_BITS = 4
const MQ_PRIO_MAX = 32768
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
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const RE_DUP_MAX = 255
const SCHAR_MAX = 127
const SEM_NSEMS_MAX = 256
const SEM_VALUE_MAX = 0x7fffffff
const SHRT_MAX = 0x7fff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "LONG_MAX"
const SYMLOOP_MAX = 40
const TTY_NAME_MAX = 32
const TZNAME_MAX = 6
const UCHAR_MAX = 255
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
const UINT_MAX = 0xffffffff
const USHRT_MAX = 0xffff
const WILDCOPYLENGTH = 8
const WINT_MAX = "UINT32_MAX"
const WINT_MIN = 0
const WNOHANG = 1
const WORD_BIT = 32
const WUNTRACED = 2
const _GNU_SOURCE = 1
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
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type wchar_t = int32

type max_align_t = struct {
	F__ll int64
	F__ld float64
}

type size_t = uint64

type ptrdiff_t = int64

type LZ4_stream_t = struct {
	Finternal_donotuse [0]LZ4_stream_t_internal
	FminStateSize      [16416]int8
}

type LZ4_stream_u = LZ4_stream_t

type LZ4_streamDecode_t = struct {
	Finternal_donotuse [0]LZ4_streamDecode_t_internal
	FminStateSize      [32]int8
}

type LZ4_streamDecode_u = LZ4_streamDecode_t

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

type LZ4_i8 = int8

type LZ4_byte = uint8

type LZ4_u16 = uint16

type LZ4_u32 = uint32

type LZ4_stream_t_internal = struct {
	FhashTable     [4096]LZ4_u32
	Fdictionary    uintptr
	FdictCtx       uintptr
	FcurrentOffset LZ4_u32
	FtableType     LZ4_u32
	FdictSize      LZ4_u32
}

type LZ4_streamDecode_t_internal = struct {
	FexternalDict uintptr
	FprefixEnd    uintptr
	FextDictSize  size_t
	FprefixSize   size_t
}

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

type locale_t = uintptr

/*-************************************
*  Common Constants
**************************************/

var LZ4_minLength = libc.Int32FromInt32(MFLIMIT) + libc.Int32FromInt32(1)

/*-************************************
*  Error detection
**************************************/

func LZ4_isAligned(tls *libc.TLS, ptr uintptr, alignment size_t) (r int32) {
	return libc.BoolInt32(uint64(ptr)&(alignment-uint64(1)) == uint64(0))
}

/*-************************************
*  Types
**************************************/

/* Support signed or unsigned plain-char */

/* Implementation choices... */

/* Arbitrary numbers... */

/* POSIX/SUS requirements follow. These numbers come directly
 * from SUS and have nothing to do with the host system. */
type BYTE = uint8

type U16 = uint16

type U32 = uint32

type S32 = int32

type U64 = uint64

type uptrval = uint64

type reg_t = uint64

/* 64-bits in x32 mode */

type limitedOutput_directive = int32

const notLimited = 0
const limitedOutput = 1
const fillOutput = 2

/*-************************************
*  Reading and writing into memory
**************************************/

/**
 * LZ4 relies on memcpy with a constant size being inlined. In freestanding
 * environments, the compiler can't assume the implementation of memcpy() is
 * standard compliant, so it can't apply its specialized memcpy() inlining
 * logic. When possible, use __builtin_memcpy() to tell the compiler to analyze
 * memcpy() as if it were standard compliant, so it can inline it in freestanding
 * environments. This is needed when decompressing the Linux Kernel, for example.
 */

func LZ4_isLittleEndian(tls *libc.TLS) (r uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* one at bp+0 */ struct {
		Fc [0][4]BYTE
		Fu U32
	}
	*(*struct {
		Fc [0][4]BYTE
		Fu U32
	})(unsafe.Pointer(bp)) = struct {
		Fc [0][4]BYTE
		Fu U32
	}{}
	*(*uint32)(unsafe.Pointer(bp)) = uint32(1) /* don't use static : performance detrimental */
	return uint32(**(**BYTE)(__ccgo_up(bp)))
}

// C documentation
//
//	/* __pack instructions are safer, but compiler specific, hence potentially problematic for some compilers */
//	/* currently only defined for gcc and icc */
type LZ4_unalign = struct {
	Fu32         [0]U32
	FuArch       [0]reg_t
	Fu16         U16
	F__ccgo_pad3 [6]byte
}

func LZ4_read16(tls *libc.TLS, ptr uintptr) (r U16) {
	return (*LZ4_unalign)(unsafe.Pointer(ptr)).Fu16
}

func LZ4_read32(tls *libc.TLS, ptr uintptr) (r U32) {
	return *(*U32)(unsafe.Pointer(ptr))
}

func LZ4_read_ARCH(tls *libc.TLS, ptr uintptr) (r reg_t) {
	return *(*reg_t)(unsafe.Pointer(ptr))
}

func LZ4_write16(tls *libc.TLS, memPtr uintptr, value U16) {
	(*LZ4_unalign)(unsafe.Pointer(memPtr)).Fu16 = value
}

func LZ4_write32(tls *libc.TLS, memPtr uintptr, value U32) {
	*(*U32)(unsafe.Pointer(memPtr)) = value
}

func LZ4_readLE16(tls *libc.TLS, memPtr uintptr) (r U16) {
	var p uintptr
	_ = p
	if LZ4_isLittleEndian(tls) != 0 {
		return LZ4_read16(tls, memPtr)
	} else {
		p = memPtr
		return libc.Uint16FromInt32(libc.Int32FromUint16(uint16(**(**BYTE)(__ccgo_up(p)))) + libc.Int32FromUint8(**(**BYTE)(__ccgo_up(p + 1)))<<libc.Int32FromInt32(8))
	}
	return r
}

func LZ4_writeLE16(tls *libc.TLS, memPtr uintptr, value U16) {
	var p uintptr
	_ = p
	if LZ4_isLittleEndian(tls) != 0 {
		LZ4_write16(tls, memPtr, value)
	} else {
		p = memPtr
		**(**BYTE)(__ccgo_up(p)) = uint8(value)
		**(**BYTE)(__ccgo_up(p + 1)) = libc.Uint8FromInt32(libc.Int32FromUint16(value) >> libc.Int32FromInt32(8))
	}
}

// C documentation
//
//	/* customized variant of memcpy, which can overwrite up to 8 bytes beyond dstEnd */
func LZ4_wildCopy8(tls *libc.TLS, dstPtr uintptr, srcPtr uintptr, dstEnd uintptr) {
	var d, e, s uintptr
	_, _, _ = d, e, s
	d = dstPtr
	s = srcPtr
	e = dstEnd
	for cond := true; cond; cond = d < e {
		libc.X__builtin_memcpy(tls, d, s, uint64(8))
		d = d + uintptr(8)
		s = s + uintptr(8)
	}
}

var inc32table = [8]uint32{
	1: uint32(1),
	2: uint32(2),
	3: uint32(1),
	5: uint32(4),
	6: uint32(4),
	7: uint32(4),
}
var dec64table = [8]int32{
	3: -int32(1),
	4: -int32(4),
	5: int32(1),
	6: int32(2),
	7: int32(3),
}

func LZ4_memcpy_using_offset_base(tls *libc.TLS, dstPtr uintptr, srcPtr uintptr, dstEnd uintptr, offset size_t) {
	if offset < uint64(8) {
		LZ4_write32(tls, dstPtr, uint32(0)) /* silence an msan warning when offset==0 */
		**(**BYTE)(__ccgo_up(dstPtr)) = **(**BYTE)(__ccgo_up(srcPtr))
		**(**BYTE)(__ccgo_up(dstPtr + 1)) = **(**BYTE)(__ccgo_up(srcPtr + 1))
		**(**BYTE)(__ccgo_up(dstPtr + 2)) = **(**BYTE)(__ccgo_up(srcPtr + 2))
		**(**BYTE)(__ccgo_up(dstPtr + 3)) = **(**BYTE)(__ccgo_up(srcPtr + 3))
		srcPtr = srcPtr + uintptr(inc32table[offset])
		libc.X__builtin_memcpy(tls, dstPtr+uintptr(4), srcPtr, uint64(4))
		srcPtr = srcPtr - uintptr(dec64table[offset])
		dstPtr = dstPtr + uintptr(8)
	} else {
		libc.X__builtin_memcpy(tls, dstPtr, srcPtr, uint64(8))
		dstPtr = dstPtr + uintptr(8)
		srcPtr = srcPtr + uintptr(8)
	}
	LZ4_wildCopy8(tls, dstPtr, srcPtr, dstEnd)
}

// C documentation
//
//	/* customized variant of memcpy, which can overwrite up to 32 bytes beyond dstEnd
//	 * this version copies two times 16 bytes (instead of one time 32 bytes)
//	 * because it must be compatible with offsets >= 16. */
func LZ4_wildCopy32(tls *libc.TLS, dstPtr uintptr, srcPtr uintptr, dstEnd uintptr) {
	var d, e, s uintptr
	_, _, _ = d, e, s
	d = dstPtr
	s = srcPtr
	e = dstEnd
	for cond := true; cond; cond = d < e {
		libc.X__builtin_memcpy(tls, d, s, uint64(16))
		libc.X__builtin_memcpy(tls, d+uintptr(16), s+uintptr(16), uint64(16))
		d = d + uintptr(32)
		s = s + uintptr(32)
	}
}

// C documentation
//
//	/* LZ4_memcpy_using_offset()  presumes :
//	 * - dstEnd >= dstPtr + MINMATCH
//	 * - there is at least 8 bytes available to write after dstEnd */
func LZ4_memcpy_using_offset(tls *libc.TLS, dstPtr uintptr, srcPtr uintptr, dstEnd uintptr, offset size_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* v at bp+0 */ [8]BYTE
	switch offset {
	case uint64(1):
		libc.Xmemset(tls, bp, libc.Int32FromUint8(**(**BYTE)(__ccgo_up(srcPtr))), libc.Uint64FromInt32(libc.Int32FromInt32(8)))
	case uint64(2):
		libc.X__builtin_memcpy(tls, bp, srcPtr, uint64(2))
		libc.X__builtin_memcpy(tls, bp+2, srcPtr, uint64(2))
		libc.X__builtin_memcpy(tls, bp+4, bp, uint64(4))
	case uint64(4):
		libc.X__builtin_memcpy(tls, bp, srcPtr, uint64(4))
		libc.X__builtin_memcpy(tls, bp+4, srcPtr, uint64(4))
	default:
		LZ4_memcpy_using_offset_base(tls, dstPtr, srcPtr, dstEnd, offset)
		return
	}
	libc.X__builtin_memcpy(tls, dstPtr, bp, uint64(8))
	dstPtr = dstPtr + uintptr(8)
	for dstPtr < dstEnd {
		libc.X__builtin_memcpy(tls, dstPtr, bp, uint64(8))
		dstPtr = dstPtr + uintptr(8)
	}
}

// C documentation
//
//	/*-************************************
//	*  Common functions
//	**************************************/
func LZ4_NbCommonBytes(tls *libc.TLS, val reg_t) (r uint32) {
	if LZ4_isLittleEndian(tls) != 0 {
		if uint64(8) == uint64(8) {
			return libc.Uint32FromInt32(libc.X__builtin_ctzll(tls, val)) >> int32(3)
		} else { /* 32 bits */
			return libc.Uint32FromInt32(libc.X__builtin_ctz(tls, uint32(val))) >> int32(3)
		}
	} else { /* Big Endian CPU */
		if uint64(8) == uint64(8) {
			return libc.Uint32FromInt32(libc.X__builtin_clzll(tls, val)) >> int32(3)
		} else { /* 32 bits */
			return libc.Uint32FromInt32(libc.X__builtin_clz(tls, uint32(val))) >> int32(3)
		}
	}
	return r
}

func LZ4_count(tls *libc.TLS, pIn uintptr, pMatch uintptr, pInLimit uintptr) (r uint32) {
	var diff, diff1 reg_t
	var pStart uintptr
	_, _, _ = diff, diff1, pStart
	pStart = pIn
	if libc.BoolInt64(libc.BoolInt32(pIn < pInLimit-uintptr(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) != libc.Int32FromInt32(0)) != 0 {
		diff = LZ4_read_ARCH(tls, pMatch) ^ LZ4_read_ARCH(tls, pIn)
		if !(diff != 0) {
			pIn = pIn + uintptr(8)
			pMatch = pMatch + uintptr(8)
		} else {
			return LZ4_NbCommonBytes(tls, diff)
		}
	}
	for libc.BoolInt64(libc.BoolInt32(pIn < pInLimit-uintptr(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) != libc.Int32FromInt32(0)) != 0 {
		diff1 = LZ4_read_ARCH(tls, pMatch) ^ LZ4_read_ARCH(tls, pIn)
		if !(diff1 != 0) {
			pIn = pIn + uintptr(8)
			pMatch = pMatch + uintptr(8)
			continue
		}
		pIn = pIn + uintptr(LZ4_NbCommonBytes(tls, diff1))
		return libc.Uint32FromInt64(int64(pIn) - int64(pStart))
	}
	if libc.Bool(uint64(8) == uint64(8)) && pIn < pInLimit-libc.UintptrFromInt32(3) && LZ4_read32(tls, pMatch) == LZ4_read32(tls, pIn) {
		pIn = pIn + uintptr(4)
		pMatch = pMatch + uintptr(4)
	}
	if pIn < pInLimit-libc.UintptrFromInt32(1) && libc.Int32FromUint16(LZ4_read16(tls, pMatch)) == libc.Int32FromUint16(LZ4_read16(tls, pIn)) {
		pIn = pIn + uintptr(2)
		pMatch = pMatch + uintptr(2)
	}
	if pIn < pInLimit && libc.Int32FromUint8(**(**BYTE)(__ccgo_up(pMatch))) == libc.Int32FromUint8(**(**BYTE)(__ccgo_up(pIn))) {
		pIn = pIn + 1
	}
	return libc.Uint32FromInt64(int64(pIn) - int64(pStart))
}

// C documentation
//
//	/*-************************************
//	*  Local Constants
//	**************************************/
var LZ4_64Klimit = libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)) + (libc.Int32FromInt32(MFLIMIT) - libc.Int32FromInt32(1))
var LZ4_skipTrigger = uint32(6)

/* Increase this value ==> compression run slower on incompressible data */

// C documentation
//
//	/*-************************************
//	*  Local Structures and types
//	**************************************/
type tableType_t = int32

const clearedTable = 0
const byPtr = 1
const byU32 = 2
const byU16 = 3

// C documentation
//
//	/**
//	 * This enum distinguishes several different modes of accessing previous
//	 * content in the stream.
//	 *
//	 * - noDict        : There is no preceding content.
//	 * - withPrefix64k : Table entries up to ctx->dictSize before the current blob
//	 *                   blob being compressed are valid and refer to the preceding
//	 *                   content (of length ctx->dictSize), which is available
//	 *                   contiguously preceding in memory the content currently
//	 *                   being compressed.
//	 * - usingExtDict  : Like withPrefix64k, but the preceding content is somewhere
//	 *                   else in memory, starting at ctx->dictionary with length
//	 *                   ctx->dictSize.
//	 * - usingDictCtx  : Everything concerning the preceding content is
//	 *                   in a separate context, pointed to by ctx->dictCtx.
//	 *                   ctx->dictionary, ctx->dictSize, and table entries
//	 *                   in the current context that refer to positions
//	 *                   preceding the beginning of the current compression are
//	 *                   ignored. Instead, ctx->dictCtx->dictionary and ctx->dictCtx
//	 *                   ->dictSize describe the location and size of the preceding
//	 *                   content, and matches are found by looking in the ctx
//	 *                   ->dictCtx->hashTable.
//	 */
type dict_directive = int32

const noDict = 0
const withPrefix64k = 1
const usingExtDict = 2
const usingDictCtx = 3

type dictIssue_directive = int32

const noDictIssue = 0
const dictSmall = 1

// C documentation
//
//	/*-************************************
//	*  Local Utils
//	**************************************/
func LZ4_versionNumber(tls *libc.TLS) (r int32) {
	return libc.Int32FromInt32(LZ4_VERSION_MAJOR)*libc.Int32FromInt32(100)*libc.Int32FromInt32(100) + libc.Int32FromInt32(LZ4_VERSION_MINOR)*libc.Int32FromInt32(100) + libc.Int32FromInt32(LZ4_VERSION_RELEASE)
}

func LZ4_versionString(tls *libc.TLS) (r uintptr) {
	return __ccgo_ts
}

func LZ4_compressBound(tls *libc.TLS, isize int32) (r int32) {
	var v1 int32
	_ = v1
	if libc.Uint32FromInt32(isize) > libc.Uint32FromInt32(LZ4_MAX_INPUT_SIZE) {
		v1 = 0
	} else {
		v1 = isize + isize/int32(255) + int32(16)
	}
	return v1
}

func LZ4_sizeofState(tls *libc.TLS) (r int32) {
	return int32(16416)
}

// C documentation
//
//	/*-******************************
//	*  Compression functions
//	********************************/
func LZ4_hash4(tls *libc.TLS, sequence U32, tableType tableType_t) (r U32) {
	if tableType == int32(byU16) {
		return sequence * libc.Uint32FromUint32(2654435761) >> (libc.Int32FromInt32(MINMATCH)*libc.Int32FromInt32(8) - (libc.Int32FromInt32(LZ4_MEMORY_USAGE_DEFAULT) - libc.Int32FromInt32(2) + libc.Int32FromInt32(1)))
	} else {
		return sequence * libc.Uint32FromUint32(2654435761) >> (libc.Int32FromInt32(MINMATCH)*libc.Int32FromInt32(8) - (libc.Int32FromInt32(LZ4_MEMORY_USAGE_DEFAULT) - libc.Int32FromInt32(2)))
	}
	return r
}

func LZ4_hash5(tls *libc.TLS, sequence U64, tableType tableType_t) (r U32) {
	var hashLog U32
	var prime5bytes, prime8bytes U64
	var v1 int32
	_, _, _, _ = hashLog, prime5bytes, prime8bytes, v1
	if tableType == int32(byU16) {
		v1 = libc.Int32FromInt32(LZ4_MEMORY_USAGE_DEFAULT) - libc.Int32FromInt32(2) + libc.Int32FromInt32(1)
	} else {
		v1 = libc.Int32FromInt32(LZ4_MEMORY_USAGE_DEFAULT) - libc.Int32FromInt32(2)
	}
	hashLog = libc.Uint32FromInt32(v1)
	if LZ4_isLittleEndian(tls) != 0 {
		prime5bytes = uint64(889523592379)
		return uint32(sequence << libc.Int32FromInt32(24) * prime5bytes >> (libc.Uint32FromInt32(64) - hashLog))
	} else {
		prime8bytes = uint64(11400714785074694791)
		return uint32(sequence >> libc.Int32FromInt32(24) * prime8bytes >> (libc.Uint32FromInt32(64) - hashLog))
	}
	return r
}

func LZ4_hashPosition(tls *libc.TLS, p uintptr, tableType tableType_t) (r U32) {
	if libc.Bool(uint64(8) == uint64(8)) && tableType != int32(byU16) {
		return LZ4_hash5(tls, LZ4_read_ARCH(tls, p), tableType)
	}
	return LZ4_hash4(tls, LZ4_read32(tls, p), tableType)
}

func LZ4_clearHash(tls *libc.TLS, h U32, tableBase uintptr, tableType tableType_t) {
	var hashTable, hashTable1, hashTable2 uintptr
	_, _, _ = hashTable, hashTable1, hashTable2
	switch tableType {
	default: /* fallthrough */
		fallthrough
	case int32(clearedTable):
		return
	case int32(byPtr):
		hashTable = tableBase
		**(**uintptr)(__ccgo_up(hashTable + uintptr(h)*8)) = libc.UintptrFromInt32(0)
		return
	case int32(byU32):
		hashTable1 = tableBase
		**(**U32)(__ccgo_up(hashTable1 + uintptr(h)*4)) = uint32(0)
		return
	case int32(byU16):
		hashTable2 = tableBase
		**(**U16)(__ccgo_up(hashTable2 + uintptr(h)*2)) = uint16(0)
		return
	}
}

func LZ4_putIndexOnHash(tls *libc.TLS, idx U32, h U32, tableBase uintptr, tableType tableType_t) {
	var hashTable, hashTable1 uintptr
	_, _ = hashTable, hashTable1
	switch tableType {
	default: /* fallthrough */
		fallthrough
	case int32(clearedTable): /* fallthrough */
		fallthrough
	case int32(byPtr):
		return
	case int32(byU32):
		hashTable = tableBase
		**(**U32)(__ccgo_up(hashTable + uintptr(h)*4)) = idx
		return
	case int32(byU16):
		hashTable1 = tableBase
		**(**U16)(__ccgo_up(hashTable1 + uintptr(h)*2)) = uint16(idx)
		return
	}
}

func LZ4_putPositionOnHash(tls *libc.TLS, p uintptr, h U32, tableBase uintptr, tableType tableType_t, srcBase uintptr) {
	var hashTable, hashTable1, hashTable2 uintptr
	_, _, _ = hashTable, hashTable1, hashTable2
	switch tableType {
	case int32(clearedTable):
		return
	case int32(byPtr):
		hashTable = tableBase
		**(**uintptr)(__ccgo_up(hashTable + uintptr(h)*8)) = p
		return
	case int32(byU32):
		hashTable1 = tableBase
		**(**U32)(__ccgo_up(hashTable1 + uintptr(h)*4)) = libc.Uint32FromInt64(int64(p) - int64(srcBase))
		return
	case int32(byU16):
		hashTable2 = tableBase
		**(**U16)(__ccgo_up(hashTable2 + uintptr(h)*2)) = libc.Uint16FromInt64(int64(p) - int64(srcBase))
		return
	}
}

func LZ4_putPosition(tls *libc.TLS, p uintptr, tableBase uintptr, tableType tableType_t, srcBase uintptr) {
	var h U32
	_ = h
	h = LZ4_hashPosition(tls, p, tableType)
	LZ4_putPositionOnHash(tls, p, h, tableBase, tableType, srcBase)
}

// C documentation
//
//	/* LZ4_getIndexOnHash() :
//	 * Index of match position registered in hash table.
//	 * hash position must be calculated by using base+index, or dictBase+index.
//	 * Assumption 1 : only valid if tableType == byU32 or byU16.
//	 * Assumption 2 : h is presumed valid (within limits of hash table)
//	 */
func LZ4_getIndexOnHash(tls *libc.TLS, h U32, tableBase uintptr, tableType tableType_t) (r U32) {
	var hashTable, hashTable1 uintptr
	_, _ = hashTable, hashTable1
	if tableType == int32(byU32) {
		hashTable = tableBase
		return **(**U32)(__ccgo_up(hashTable + uintptr(h)*4))
	}
	if tableType == int32(byU16) {
		hashTable1 = tableBase
		return uint32(**(**U16)(__ccgo_up(hashTable1 + uintptr(h)*2)))
	}
	return uint32(0) /* forbidden case */
}

func LZ4_getPositionOnHash(tls *libc.TLS, h U32, tableBase uintptr, tableType tableType_t, srcBase uintptr) (r uintptr) {
	var hashTable, hashTable1, hashTable2 uintptr
	_, _, _ = hashTable, hashTable1, hashTable2
	if tableType == int32(byPtr) {
		hashTable = tableBase
		return **(**uintptr)(__ccgo_up(hashTable + uintptr(h)*8))
	}
	if tableType == int32(byU32) {
		hashTable1 = tableBase
		return uintptr(**(**U32)(__ccgo_up(hashTable1 + uintptr(h)*4))) + srcBase
	}
	hashTable2 = tableBase
	return uintptr(**(**U16)(__ccgo_up(hashTable2 + uintptr(h)*2))) + srcBase /* default, to ensure a return */
	return r
}

func LZ4_getPosition(tls *libc.TLS, p uintptr, tableBase uintptr, tableType tableType_t, srcBase uintptr) (r uintptr) {
	var h U32
	_ = h
	h = LZ4_hashPosition(tls, p, tableType)
	return LZ4_getPositionOnHash(tls, h, tableBase, tableType, srcBase)
}

func LZ4_prepareTable(tls *libc.TLS, cctx uintptr, inputSize int32, tableType tableType_t) {
	/* If the table hasn't been used, it's guaranteed to be zeroed out, and is
	 * therefore safe to use no matter what mode we're in. Otherwise, we figure
	 * out if it's safe to leave as is or whether it needs to be reset.
	 */
	if libc.Int32FromUint32((*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FtableType) != int32(clearedTable) {
		if libc.Int32FromUint32((*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FtableType) != tableType || tableType == int32(byU16) && (*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FcurrentOffset+libc.Uint32FromInt32(inputSize) >= uint32(0xFFFF) || tableType == int32(byU32) && (*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FcurrentOffset > libc.Uint32FromInt32(1)*(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(30)) || tableType == int32(byPtr) || inputSize >= libc.Int32FromInt32(4)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)) {
			libc.Xmemset(tls, cctx, 0, libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(LZ4_MEMORY_USAGE_DEFAULT)))
			(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FcurrentOffset = uint32(0)
			(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FtableType = uint32(clearedTable)
		} else {
		}
	}
	/* Adding a gap, so all previous entries are > LZ4_DISTANCE_MAX back,
	 * is faster than compressing without a gap.
	 * However, compressing with currentOffset == 0 is faster still,
	 * so we preserve that case.
	 */
	if (*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FcurrentOffset != uint32(0) && tableType == int32(byU32) {
		**(**LZ4_u32)(__ccgo_up(cctx + 16400)) += libc.Uint32FromInt32(libc.Int32FromInt32(64) * (libc.Int32FromInt32(1) << libc.Int32FromInt32(10)))
	}
	/* Finally, clear history */
	(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FdictCtx = libc.UintptrFromInt32(0)
	(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).Fdictionary = libc.UintptrFromInt32(0)
	(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FdictSize = uint32(0)
}

// C documentation
//
//	/** LZ4_compress_generic() :
//	 *  inlined, to ensure branches are decided at compilation time.
//	 *  Presumed already validated at this stage:
//	 *  - source != NULL
//	 *  - inputSize > 0
//	 */
func LZ4_compress_generic_validated(tls *libc.TLS, cctx uintptr, source uintptr, dest uintptr, inputSize int32, inputConsumed uintptr, maxOutputSize int32, outputDirective limitedOutput_directive, tableType tableType_t, dictDirective dict_directive, dictIssue dictIssue_directive, acceleration int32) (r int32) {
	var accumulator, lastRun size_t
	var anchor, base, dictBase, dictCtx, dictEnd, dictionary, filledIp, forwardIp, forwardIp1, iend, ip, limit, lowLimit, match, matchlimit, mflimitPlusOne, olimit, op, ptr, token, v1, v4, v5, v6 uintptr
	var current, current1, dictDelta, dictSize, forwardH, h, h1, h2, h3, matchIndex, matchIndex1, newMatchCode, offset, prefixIdxLimit, startIndex U32
	var len1, maybe_extMem, result, searchMatchNb, searchMatchNb1, step, step1, v10, v9 int32
	var litLength, matchCode, more, v2, v3 uint32
	var v22 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = accumulator, anchor, base, current, current1, dictBase, dictCtx, dictDelta, dictEnd, dictSize, dictionary, filledIp, forwardH, forwardIp, forwardIp1, h, h1, h2, h3, iend, ip, lastRun, len1, limit, litLength, lowLimit, match, matchCode, matchIndex, matchIndex1, matchlimit, maybe_extMem, mflimitPlusOne, more, newMatchCode, offset, olimit, op, prefixIdxLimit, ptr, result, searchMatchNb, searchMatchNb1, startIndex, step, step1, token, v1, v10, v2, v22, v3, v4, v5, v6, v9
	ip = source
	startIndex = (*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FcurrentOffset
	base = source - uintptr(startIndex)
	dictCtx = (*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FdictCtx
	if dictDirective == int32(usingDictCtx) {
		v1 = (*LZ4_stream_t_internal)(unsafe.Pointer(dictCtx)).Fdictionary
	} else {
		v1 = (*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).Fdictionary
	}
	dictionary = v1
	if dictDirective == int32(usingDictCtx) {
		v2 = (*LZ4_stream_t_internal)(unsafe.Pointer(dictCtx)).FdictSize
	} else {
		v2 = (*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FdictSize
	}
	dictSize = v2
	if dictDirective == int32(usingDictCtx) {
		v3 = startIndex - (*LZ4_stream_t_internal)(unsafe.Pointer(dictCtx)).FcurrentOffset
	} else {
		v3 = uint32(0)
	}
	dictDelta = v3 /* make indexes in dictCtx comparable with index in current context */
	maybe_extMem = libc.BoolInt32(dictDirective == int32(usingExtDict) || dictDirective == int32(usingDictCtx))
	prefixIdxLimit = startIndex - dictSize
	if dictionary != 0 {
		v4 = dictionary + uintptr(dictSize)
	} else {
		v4 = dictionary
	} /* used when dictDirective == dictSmall */
	dictEnd = v4
	anchor = source
	iend = ip + uintptr(inputSize)
	mflimitPlusOne = iend - uintptr(MFLIMIT) + uintptr(1)
	matchlimit = iend - uintptr(LASTLITERALS)
	if dictionary == libc.UintptrFromInt32(0) {
		v5 = libc.UintptrFromInt32(0)
	} else {
		if dictDirective == int32(usingDictCtx) {
			v6 = dictionary + uintptr(dictSize) - uintptr((*LZ4_stream_t_internal)(unsafe.Pointer(dictCtx)).FcurrentOffset)
		} else {
			v6 = dictionary + uintptr(dictSize) - uintptr(startIndex)
		}
		v5 = v6
	}
	/* the dictCtx currentOffset is indexed on the start of the dictionary,
	 * while a dictionary in the current context precedes the currentOffset */
	dictBase = v5
	op = dest
	olimit = op + uintptr(maxOutputSize)
	offset = uint32(0)
	/* If init conditions are not met, we don't have to mark stream
	 * as having dirty context, since no action was taken yet */
	if outputDirective == int32(fillOutput) && maxOutputSize < int32(1) {
		return 0
	} /* Impossible to store anything */
	if tableType == int32(byU16) && inputSize >= LZ4_64Klimit {
		return 0
	} /* Size too large (not within 64K limit) */
	if tableType == int32(byPtr) {
	} /* only supported use case with byPtr */
	if dictDirective == int32(withPrefix64k) {
		v2 = dictSize
	} else {
		v2 = uint32(0)
	}
	lowLimit = source - uintptr(v2)
	/* Update context state */
	if dictDirective == int32(usingDictCtx) {
		/* Subsequent linked blocks can't use the dictionary. */
		/* Instead, they use the block we just compressed. */
		(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FdictCtx = libc.UintptrFromInt32(0)
		(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FdictSize = libc.Uint32FromInt32(inputSize)
	} else {
		**(**LZ4_u32)(__ccgo_up(cctx + 16408)) += libc.Uint32FromInt32(inputSize)
	}
	**(**LZ4_u32)(__ccgo_up(cctx + 16400)) += libc.Uint32FromInt32(inputSize)
	(*LZ4_stream_t_internal)(unsafe.Pointer(cctx)).FtableType = libc.Uint32FromInt32(tableType)
	if inputSize < LZ4_minLength {
		goto _last_literals
	} /* Input too small, no compression (all literals) */
	/* First Byte */
	LZ4_putPosition(tls, ip, cctx, tableType, base)
	ip = ip + 1
	forwardH = LZ4_hashPosition(tls, ip, tableType)
	/* Main Loop */
	for {
		/* Find a match */
		if tableType == int32(byPtr) {
			forwardIp = ip
			step = int32(1)
			searchMatchNb = acceleration << LZ4_skipTrigger
			for cond := true; cond; cond = match+uintptr(LZ4_DISTANCE_MAX) < ip || LZ4_read32(tls, match) != LZ4_read32(tls, ip) {
				h = forwardH
				ip = forwardIp
				forwardIp = forwardIp + uintptr(step)
				v9 = searchMatchNb
				searchMatchNb = searchMatchNb + 1
				step = v9 >> LZ4_skipTrigger
				if libc.BoolInt64(libc.BoolInt32(forwardIp > mflimitPlusOne) != libc.Int32FromInt32(0)) != 0 {
					goto _last_literals
				}
				match = LZ4_getPositionOnHash(tls, h, cctx, tableType, base)
				forwardH = LZ4_hashPosition(tls, forwardIp, tableType)
				LZ4_putPositionOnHash(tls, ip, h, cctx, tableType, base)
			}
		} else { /* byU32, byU16 */
			forwardIp1 = ip
			step1 = int32(1)
			searchMatchNb1 = acceleration << LZ4_skipTrigger
			for cond := true; cond; cond = int32(1) != 0 {
				h1 = forwardH
				current = libc.Uint32FromInt64(int64(forwardIp1) - int64(base))
				matchIndex = LZ4_getIndexOnHash(tls, h1, cctx, tableType)
				ip = forwardIp1
				forwardIp1 = forwardIp1 + uintptr(step1)
				v9 = searchMatchNb1
				searchMatchNb1 = searchMatchNb1 + 1
				step1 = v9 >> LZ4_skipTrigger
				if libc.BoolInt64(libc.BoolInt32(forwardIp1 > mflimitPlusOne) != libc.Int32FromInt32(0)) != 0 {
					goto _last_literals
				}
				if dictDirective == int32(usingDictCtx) {
					if matchIndex < startIndex {
						/* there was no match, try the dictionary */
						matchIndex = LZ4_getIndexOnHash(tls, h1, dictCtx, int32(byU32))
						match = dictBase + uintptr(matchIndex)
						matchIndex = matchIndex + dictDelta /* make dictCtx index comparable with current context */
						lowLimit = dictionary
					} else {
						match = base + uintptr(matchIndex)
						lowLimit = source
					}
				} else {
					if dictDirective == int32(usingExtDict) {
						if matchIndex < startIndex {
							match = dictBase + uintptr(matchIndex)
							lowLimit = dictionary
						} else {
							match = base + uintptr(matchIndex)
							lowLimit = source
						}
					} else { /* single continuous memory segment */
						match = base + uintptr(matchIndex)
					}
				}
				forwardH = LZ4_hashPosition(tls, forwardIp1, tableType)
				LZ4_putIndexOnHash(tls, current, h1, cctx, tableType)
				if dictIssue == int32(dictSmall) && matchIndex < prefixIdxLimit {
					continue
				} /* match outside of valid area */
				if (tableType != int32(byU16) || libc.Bool(int32(LZ4_DISTANCE_MAX) < int32(LZ4_DISTANCE_ABSOLUTE_MAX))) && matchIndex+uint32(LZ4_DISTANCE_MAX) < current {
					continue
				} /* too far */
				/* match now expected within distance */
				if LZ4_read32(tls, match) == LZ4_read32(tls, ip) {
					if maybe_extMem != 0 {
						offset = current - matchIndex
					}
					break /* match found */
				}
			}
		}
		/* Catch up */
		filledIp = ip
		for libc.BoolInt32(ip > anchor)&libc.BoolInt32(match > lowLimit) != 0 && libc.BoolInt64(libc.BoolInt32(libc.Int32FromUint8(**(**BYTE)(__ccgo_up(ip + uintptr(-libc.Int32FromInt32(1))))) == libc.Int32FromUint8(**(**BYTE)(__ccgo_up(match + uintptr(-libc.Int32FromInt32(1)))))) != libc.Int32FromInt32(0)) != 0 {
			ip = ip - 1
			match = match - 1
		}
		/* Encode Literals */
		litLength = libc.Uint32FromInt64(int64(ip) - int64(anchor))
		v1 = op
		op = op + 1
		token = v1
		if outputDirective == int32(limitedOutput) && libc.BoolInt64(libc.BoolInt32(op+uintptr(litLength)+uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(1)+libc.Int32FromInt32(LASTLITERALS))+uintptr(litLength/libc.Uint32FromInt32(255)) > olimit) != libc.Int32FromInt32(0)) != 0 {
			return 0 /* cannot compress within `dst` budget. Stored indexes in hash table are nonetheless fine */
		}
		if outputDirective == int32(fillOutput) && libc.BoolInt64(libc.BoolInt32(op+uintptr((litLength+libc.Uint32FromInt32(240))/libc.Uint32FromInt32(255))+uintptr(litLength)+libc.UintptrFromInt32(2)+libc.UintptrFromInt32(1)+libc.UintptrFromInt32(MFLIMIT)-libc.UintptrFromInt32(MINMATCH) > olimit) != libc.Int32FromInt32(0)) != 0 {
			op = op - 1
			goto _last_literals
		}
		if litLength >= libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1) {
			len1 = libc.Int32FromUint32(litLength - (libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS)) - libc.Uint32FromInt32(1)))
			**(**BYTE)(__ccgo_up(token)) = uint8((libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS)) - libc.Uint32FromInt32(1)) << libc.Int32FromInt32(ML_BITS))
			for {
				if !(len1 >= int32(255)) {
					break
				}
				v1 = op
				op = op + 1
				**(**BYTE)(__ccgo_up(v1)) = uint8(255)
				goto _12
			_12:
				;
				len1 = len1 - int32(255)
			}
			v1 = op
			op = op + 1
			**(**BYTE)(__ccgo_up(v1)) = libc.Uint8FromInt32(len1)
		} else {
			**(**BYTE)(__ccgo_up(token)) = uint8(litLength << libc.Int32FromInt32(ML_BITS))
		}
		/* Copy Literals */
		LZ4_wildCopy8(tls, op, anchor, op+uintptr(litLength))
		op = op + uintptr(litLength)
		goto _next_match
	_next_match:
		;
		/* at this stage, the following variables must be correctly set :
		 * - ip : at start of LZ operation
		 * - match : at start of previous pattern occurrence; can be within current prefix, or within extDict
		 * - offset : if maybe_ext_memSegment==1 (constant)
		 * - lowLimit : must be == dictionary to mean "match is within extDict"; must be == source otherwise
		 * - token and *token : position to write 4-bits for match length; higher 4-bits for literal length supposed already written
		 */
		if outputDirective == int32(fillOutput) && op+uintptr(2)+uintptr(1)+uintptr(MFLIMIT)-uintptr(MINMATCH) > olimit {
			/* the match was too close to the end, rewind and go to last literals */
			op = token
			goto _last_literals
		}
		/* Encode Offset */
		if maybe_extMem != 0 { /* static test */
			LZ4_writeLE16(tls, op, uint16(offset))
			op = op + uintptr(2)
		} else {
			LZ4_writeLE16(tls, op, libc.Uint16FromInt64(int64(ip)-int64(match)))
			op = op + uintptr(2)
		}
		/* Encode MatchLength */
		if (dictDirective == int32(usingExtDict) || dictDirective == int32(usingDictCtx)) && lowLimit == dictionary {
			limit = ip + uintptr(int64(dictEnd)-int64(match))
			if limit > matchlimit {
				limit = matchlimit
			}
			matchCode = LZ4_count(tls, ip+uintptr(MINMATCH), match+uintptr(MINMATCH), limit)
			ip = ip + uintptr(uint64(matchCode)+uint64(MINMATCH))
			if ip == limit {
				more = LZ4_count(tls, limit, source, matchlimit)
				matchCode = matchCode + more
				ip = ip + uintptr(more)
			}
		} else {
			matchCode = LZ4_count(tls, ip+uintptr(MINMATCH), match+uintptr(MINMATCH), matchlimit)
			ip = ip + uintptr(uint64(matchCode)+uint64(MINMATCH))
		}
		if outputDirective != 0 && libc.BoolInt64(libc.BoolInt32(op+uintptr(libc.Int32FromInt32(1)+libc.Int32FromInt32(LASTLITERALS))+uintptr((matchCode+libc.Uint32FromInt32(240))/libc.Uint32FromInt32(255)) > olimit) != libc.Int32FromInt32(0)) != 0 {
			if outputDirective == int32(fillOutput) {
				/* Match description too long : reduce it */
				newMatchCode = libc.Uint32FromInt32(libc.Int32FromInt32(15)-libc.Int32FromInt32(1)) + (libc.Uint32FromInt64(int64(olimit)-int64(op))-uint32(1)-uint32(LASTLITERALS))*uint32(255)
				ip = ip - uintptr(matchCode-newMatchCode)
				matchCode = newMatchCode
				if libc.BoolInt64(libc.BoolInt32(ip <= filledIp) != libc.Int32FromInt32(0)) != 0 {
					ptr = ip
					for {
						if !(ptr <= filledIp) {
							break
						}
						h2 = LZ4_hashPosition(tls, ptr, tableType)
						LZ4_clearHash(tls, h2, cctx, tableType)
						goto _15
					_15:
						;
						ptr = ptr + 1
					}
				}
			} else {
				return 0 /* cannot compress within `dst` budget. Stored indexes in hash table are nonetheless fine */
			}
		}
		if matchCode >= libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS)-libc.Uint32FromInt32(1) {
			v1 = token
			*(*BYTE)(unsafe.Pointer(v1)) = BYTE(uint32(*(*BYTE)(unsafe.Pointer(v1))) + (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS) - libc.Uint32FromInt32(1)))
			matchCode = matchCode - (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS) - libc.Uint32FromInt32(1))
			LZ4_write32(tls, op, uint32(0xFFFFFFFF))
			for matchCode >= libc.Uint32FromInt32(libc.Int32FromInt32(4)*libc.Int32FromInt32(255)) {
				op = op + uintptr(4)
				LZ4_write32(tls, op, uint32(0xFFFFFFFF))
				matchCode = matchCode - libc.Uint32FromInt32(libc.Int32FromInt32(4)*libc.Int32FromInt32(255))
			}
			op = op + uintptr(matchCode/uint32(255))
			v1 = op
			op = op + 1
			**(**BYTE)(__ccgo_up(v1)) = uint8(matchCode % libc.Uint32FromInt32(255))
		} else {
			v1 = token
			*(*BYTE)(unsafe.Pointer(v1)) = BYTE(int32(*(*BYTE)(unsafe.Pointer(v1))) + libc.Int32FromUint8(uint8(matchCode)))
		}
		/* Ensure we have enough space for the last literals. */
		anchor = ip
		/* Test end of chunk */
		if ip >= mflimitPlusOne {
			break
		}
		/* Fill table */
		LZ4_putPosition(tls, ip-uintptr(2), cctx, tableType, base)
		/* Test next position */
		if tableType == int32(byPtr) {
			match = LZ4_getPosition(tls, ip, cctx, tableType, base)
			LZ4_putPosition(tls, ip, cctx, tableType, base)
			if match+uintptr(LZ4_DISTANCE_MAX) >= ip && LZ4_read32(tls, match) == LZ4_read32(tls, ip) {
				v1 = op
				op = op + 1
				token = v1
				**(**BYTE)(__ccgo_up(token)) = uint8(0)
				goto _next_match
			}
		} else { /* byU32, byU16 */
			h3 = LZ4_hashPosition(tls, ip, tableType)
			current1 = libc.Uint32FromInt64(int64(ip) - int64(base))
			matchIndex1 = LZ4_getIndexOnHash(tls, h3, cctx, tableType)
			if dictDirective == int32(usingDictCtx) {
				if matchIndex1 < startIndex {
					/* there was no match, try the dictionary */
					matchIndex1 = LZ4_getIndexOnHash(tls, h3, dictCtx, int32(byU32))
					match = dictBase + uintptr(matchIndex1)
					lowLimit = dictionary /* required for match length counter */
					matchIndex1 = matchIndex1 + dictDelta
				} else {
					match = base + uintptr(matchIndex1)
					lowLimit = source /* required for match length counter */
				}
			} else {
				if dictDirective == int32(usingExtDict) {
					if matchIndex1 < startIndex {
						match = dictBase + uintptr(matchIndex1)
						lowLimit = dictionary /* required for match length counter */
					} else {
						match = base + uintptr(matchIndex1)
						lowLimit = source /* required for match length counter */
					}
				} else { /* single memory segment */
					match = base + uintptr(matchIndex1)
				}
			}
			LZ4_putIndexOnHash(tls, current1, h3, cctx, tableType)
			if dictIssue == int32(dictSmall) {
				v9 = libc.BoolInt32(matchIndex1 >= prefixIdxLimit)
			} else {
				v9 = int32(1)
			}
			if v22 = v9 != 0; v22 {
				if tableType == int32(byU16) && libc.Bool(true) {
					v10 = int32(1)
				} else {
					v10 = libc.BoolInt32(matchIndex1+uint32(LZ4_DISTANCE_MAX) >= current1)
				}
			}
			if v22 && v10 != 0 && LZ4_read32(tls, match) == LZ4_read32(tls, ip) {
				v1 = op
				op = op + 1
				token = v1
				**(**BYTE)(__ccgo_up(token)) = uint8(0)
				if maybe_extMem != 0 {
					offset = current1 - matchIndex1
				}
				goto _next_match
			}
		}
		/* Prepare next loop */
		ip = ip + 1
		v1 = ip
		forwardH = LZ4_hashPosition(tls, v1, tableType)
		goto _8
	_8:
	}
	goto _last_literals
_last_literals:
	;
	/* Encode Last Literals */
	lastRun = libc.Uint64FromInt64(int64(iend) - int64(anchor))
	if outputDirective != 0 && op+uintptr(lastRun)+uintptr(1)+uintptr((lastRun+libc.Uint64FromInt32(255)-uint64(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)))/libc.Uint64FromInt32(255)) > olimit {
		if outputDirective == int32(fillOutput) {
			/* adapt lastRun to fill 'dst' */
			lastRun = libc.Uint64FromInt64(int64(olimit)-int64(op)) - uint64(1)
			lastRun = lastRun - (lastRun+uint64(256)-uint64(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)))/uint64(256) /*additional length tokens*/
		} else {
			return 0 /* cannot compress within `dst` budget. Stored indexes in hash table are nonetheless fine */
		}
	}
	if lastRun >= uint64(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)) {
		accumulator = lastRun - uint64(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1))
		v1 = op
		op = op + 1
		**(**BYTE)(__ccgo_up(v1)) = uint8((libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS)) - libc.Uint32FromInt32(1)) << libc.Int32FromInt32(ML_BITS))
		for {
			if !(accumulator >= uint64(255)) {
				break
			}
			v1 = op
			op = op + 1
			**(**BYTE)(__ccgo_up(v1)) = uint8(255)
			goto _26
		_26:
			;
			accumulator = accumulator - uint64(255)
		}
		v1 = op
		op = op + 1
		**(**BYTE)(__ccgo_up(v1)) = uint8(accumulator)
	} else {
		v1 = op
		op = op + 1
		**(**BYTE)(__ccgo_up(v1)) = uint8(lastRun << libc.Int32FromInt32(ML_BITS))
	}
	libc.X__builtin_memcpy(tls, op, anchor, lastRun)
	ip = anchor + uintptr(lastRun)
	op = op + uintptr(lastRun)
	if outputDirective == int32(fillOutput) {
		**(**int32)(__ccgo_up(inputConsumed)) = int32(int64(ip) - int64(source))
	}
	result = int32(int64(op) - int64(dest))
	return result
}

// C documentation
//
//	/** LZ4_compress_generic() :
//	 *  inlined, to ensure branches are decided at compilation time;
//	 *  takes care of src == (NULL, 0)
//	 *  and forward the rest to LZ4_compress_generic_validated */
func LZ4_compress_generic(tls *libc.TLS, cctx uintptr, src uintptr, dst uintptr, srcSize int32, inputConsumed uintptr, dstCapacity int32, outputDirective limitedOutput_directive, tableType tableType_t, dictDirective dict_directive, dictIssue dictIssue_directive, acceleration int32) (r int32) {
	if libc.Uint32FromInt32(srcSize) > libc.Uint32FromInt32(LZ4_MAX_INPUT_SIZE) {
		return 0
	} /* Unsupported srcSize, too large (or negative) */
	if srcSize == 0 { /* src == NULL supported if srcSize == 0 */
		if outputDirective != int32(notLimited) && dstCapacity <= 0 {
			return 0
		} /* no output, can't write anything */
		**(**int8)(__ccgo_up(dst)) = 0
		if outputDirective == int32(fillOutput) {
			**(**int32)(__ccgo_up(inputConsumed)) = 0
		}
		return int32(1)
	}
	return LZ4_compress_generic_validated(tls, cctx, src, dst, srcSize, inputConsumed, dstCapacity, outputDirective, tableType, dictDirective, dictIssue, acceleration)
}

func LZ4_compress_fast_extState(tls *libc.TLS, state uintptr, source uintptr, dest uintptr, inputSize int32, maxOutputSize int32, acceleration int32) (r int32) {
	var ctx uintptr
	var tableType, tableType1 tableType_t
	_, _, _ = ctx, tableType, tableType1
	ctx = LZ4_initStream(tls, state, uint64(16416))
	if acceleration < int32(1) {
		acceleration = int32(LZ4_ACCELERATION_DEFAULT)
	}
	if acceleration > int32(LZ4_ACCELERATION_MAX) {
		acceleration = int32(LZ4_ACCELERATION_MAX)
	}
	if maxOutputSize >= LZ4_compressBound(tls, inputSize) {
		if inputSize < LZ4_64Klimit {
			return LZ4_compress_generic(tls, ctx, source, dest, inputSize, libc.UintptrFromInt32(0), 0, int32(notLimited), int32(byU16), int32(noDict), int32(noDictIssue), acceleration)
		} else {
			tableType = int32(byU32)
			return LZ4_compress_generic(tls, ctx, source, dest, inputSize, libc.UintptrFromInt32(0), 0, int32(notLimited), tableType, int32(noDict), int32(noDictIssue), acceleration)
		}
	} else {
		if inputSize < LZ4_64Klimit {
			return LZ4_compress_generic(tls, ctx, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), int32(byU16), int32(noDict), int32(noDictIssue), acceleration)
		} else {
			tableType1 = int32(byU32)
			return LZ4_compress_generic(tls, ctx, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), tableType1, int32(noDict), int32(noDictIssue), acceleration)
		}
	}
	return r
}

// C documentation
//
//	/**
//	 * LZ4_compress_fast_extState_fastReset() :
//	 * A variant of LZ4_compress_fast_extState().
//	 *
//	 * Using this variant avoids an expensive initialization step. It is only safe
//	 * to call if the state buffer is known to be correctly initialized already
//	 * (see comment in lz4.h on LZ4_resetStream_fast() for a definition of
//	 * "correctly initialized").
//	 */
func LZ4_compress_fast_extState_fastReset(tls *libc.TLS, state uintptr, src uintptr, dst uintptr, srcSize int32, dstCapacity int32, acceleration int32) (r int32) {
	var ctx uintptr
	var tableType, tableType1, tableType2, tableType3 tableType_t
	_, _, _, _, _ = ctx, tableType, tableType1, tableType2, tableType3
	ctx = state
	if acceleration < int32(1) {
		acceleration = int32(LZ4_ACCELERATION_DEFAULT)
	}
	if acceleration > int32(LZ4_ACCELERATION_MAX) {
		acceleration = int32(LZ4_ACCELERATION_MAX)
	}
	if dstCapacity >= LZ4_compressBound(tls, srcSize) {
		if srcSize < LZ4_64Klimit {
			tableType = int32(byU16)
			LZ4_prepareTable(tls, ctx, srcSize, tableType)
			if (*LZ4_stream_t_internal)(unsafe.Pointer(ctx)).FcurrentOffset != 0 {
				return LZ4_compress_generic(tls, ctx, src, dst, srcSize, libc.UintptrFromInt32(0), 0, int32(notLimited), tableType, int32(noDict), int32(dictSmall), acceleration)
			} else {
				return LZ4_compress_generic(tls, ctx, src, dst, srcSize, libc.UintptrFromInt32(0), 0, int32(notLimited), tableType, int32(noDict), int32(noDictIssue), acceleration)
			}
		} else {
			tableType1 = int32(byU32)
			LZ4_prepareTable(tls, ctx, srcSize, tableType1)
			return LZ4_compress_generic(tls, ctx, src, dst, srcSize, libc.UintptrFromInt32(0), 0, int32(notLimited), tableType1, int32(noDict), int32(noDictIssue), acceleration)
		}
	} else {
		if srcSize < LZ4_64Klimit {
			tableType2 = int32(byU16)
			LZ4_prepareTable(tls, ctx, srcSize, tableType2)
			if (*LZ4_stream_t_internal)(unsafe.Pointer(ctx)).FcurrentOffset != 0 {
				return LZ4_compress_generic(tls, ctx, src, dst, srcSize, libc.UintptrFromInt32(0), dstCapacity, int32(limitedOutput), tableType2, int32(noDict), int32(dictSmall), acceleration)
			} else {
				return LZ4_compress_generic(tls, ctx, src, dst, srcSize, libc.UintptrFromInt32(0), dstCapacity, int32(limitedOutput), tableType2, int32(noDict), int32(noDictIssue), acceleration)
			}
		} else {
			tableType3 = int32(byU32)
			LZ4_prepareTable(tls, ctx, srcSize, tableType3)
			return LZ4_compress_generic(tls, ctx, src, dst, srcSize, libc.UintptrFromInt32(0), dstCapacity, int32(limitedOutput), tableType3, int32(noDict), int32(noDictIssue), acceleration)
		}
	}
	return r
}

func LZ4_compress_fast(tls *libc.TLS, source uintptr, dest uintptr, inputSize int32, maxOutputSize int32, acceleration int32) (r int32) {
	bp := tls.Alloc(16416)
	defer tls.Free(16416)
	var ctxPtr uintptr
	var result int32
	var _ /* ctx at bp+0 */ LZ4_stream_t
	_, _ = ctxPtr, result
	ctxPtr = bp
	result = LZ4_compress_fast_extState(tls, ctxPtr, source, dest, inputSize, maxOutputSize, acceleration)
	return result
}

func LZ4_compress_default(tls *libc.TLS, src uintptr, dst uintptr, srcSize int32, maxOutputSize int32) (r int32) {
	return LZ4_compress_fast(tls, src, dst, srcSize, maxOutputSize, int32(1))
}

// C documentation
//
//	/* Note!: This function leaves the stream in an unclean/broken state!
//	 * It is not safe to subsequently use the same state with a _fastReset() or
//	 * _continue() call without resetting it. */
func LZ4_compress_destSize_extState(tls *libc.TLS, state uintptr, src uintptr, dst uintptr, srcSizePtr uintptr, targetDstSize int32) (r int32) {
	var addrMode tableType_t
	var s uintptr
	_, _ = addrMode, s
	s = LZ4_initStream(tls, state, uint64(16416))
	_ = s
	if targetDstSize >= LZ4_compressBound(tls, **(**int32)(__ccgo_up(srcSizePtr))) { /* compression success is guaranteed */
		return LZ4_compress_fast_extState(tls, state, src, dst, **(**int32)(__ccgo_up(srcSizePtr)), targetDstSize, int32(1))
	} else {
		if **(**int32)(__ccgo_up(srcSizePtr)) < LZ4_64Klimit {
			return LZ4_compress_generic(tls, state, src, dst, **(**int32)(__ccgo_up(srcSizePtr)), srcSizePtr, targetDstSize, int32(fillOutput), int32(byU16), int32(noDict), int32(noDictIssue), int32(1))
		} else {
			addrMode = int32(byU32)
			return LZ4_compress_generic(tls, state, src, dst, **(**int32)(__ccgo_up(srcSizePtr)), srcSizePtr, targetDstSize, int32(fillOutput), addrMode, int32(noDict), int32(noDictIssue), int32(1))
		}
	}
	return r
}

func LZ4_compress_destSize(tls *libc.TLS, src uintptr, dst uintptr, srcSizePtr uintptr, targetDstSize int32) (r int32) {
	bp := tls.Alloc(16416)
	defer tls.Free(16416)
	var ctx uintptr
	var result int32
	var _ /* ctxBody at bp+0 */ LZ4_stream_t
	_, _ = ctx, result
	ctx = bp
	result = LZ4_compress_destSize_extState(tls, ctx, src, dst, srcSizePtr, targetDstSize)
	return result
}

/*-******************************
*  Streaming functions
********************************/

func LZ4_createStream(tls *libc.TLS) (r uintptr) {
	var lz4s uintptr
	_ = lz4s
	lz4s = libc.Xmalloc(tls, uint64(16416))
	if lz4s == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	LZ4_initStream(tls, lz4s, uint64(16416))
	return lz4s
}

func LZ4_stream_t_alignment(tls *libc.TLS) (r size_t) {
	return libc.Uint64FromInt64(16424) - libc.Uint64FromInt64(16416)
}

func LZ4_initStream(tls *libc.TLS, buffer uintptr, size size_t) (r uintptr) {
	if buffer == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if size < uint64(16416) {
		return libc.UintptrFromInt32(0)
	}
	if !(LZ4_isAligned(tls, buffer, LZ4_stream_t_alignment(tls)) != 0) {
		return libc.UintptrFromInt32(0)
	}
	libc.Xmemset(tls, buffer, 0, libc.Uint64FromInt64(16416))
	return buffer
}

// C documentation
//
//	/* resetStream is now deprecated,
//	 * prefer initStream() which is more general */
func LZ4_resetStream(tls *libc.TLS, LZ4_stream uintptr) {
	libc.Xmemset(tls, LZ4_stream, 0, libc.Uint64FromInt64(16416))
}

func LZ4_resetStream_fast(tls *libc.TLS, ctx uintptr) {
	LZ4_prepareTable(tls, ctx, 0, int32(byU32))
}

func LZ4_freeStream(tls *libc.TLS, LZ4_stream uintptr) (r int32) {
	if !(LZ4_stream != 0) {
		return 0
	} /* support free on NULL */
	libc.Xfree(tls, LZ4_stream)
	return 0
}

func LZ4_loadDict(tls *libc.TLS, LZ4_dict uintptr, dictionary uintptr, dictSize int32) (r int32) {
	var base, dict, dictEnd, p uintptr
	var tableType tableType_t
	_, _, _, _, _ = base, dict, dictEnd, p, tableType
	dict = LZ4_dict
	tableType = int32(byU32)
	p = dictionary
	dictEnd = p + uintptr(dictSize)
	/* It's necessary to reset the context,
	 * and not just continue it with prepareTable()
	 * to avoid any risk of generating overflowing matchIndex
	 * when compressing using this dictionary */
	LZ4_resetStream(tls, LZ4_dict)
	/* We always increment the offset by 64 KB, since, if the dict is longer,
	 * we truncate it to the last 64k, and if it's shorter, we still want to
	 * advance by a whole window length so we can provide the guarantee that
	 * there are only valid offsets in the window, which allows an optimization
	 * in LZ4_compress_fast_continue() where it uses noDictIssue even when the
	 * dictionary isn't a full 64k. */
	**(**LZ4_u32)(__ccgo_up(dict + 16400)) += libc.Uint32FromInt32(libc.Int32FromInt32(64) * (libc.Int32FromInt32(1) << libc.Int32FromInt32(10)))
	if dictSize < libc.Int32FromInt64(8) {
		return 0
	}
	if int64(dictEnd)-int64(p) > int64(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))) {
		p = dictEnd - uintptr(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)))
	}
	base = dictEnd - uintptr((*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FcurrentOffset)
	(*LZ4_stream_t_internal)(unsafe.Pointer(dict)).Fdictionary = p
	(*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FdictSize = libc.Uint32FromInt64(int64(dictEnd) - int64(p))
	(*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FtableType = libc.Uint32FromInt32(tableType)
	for p <= dictEnd-uintptr(8) {
		LZ4_putPosition(tls, p, dict, tableType, base)
		p = p + uintptr(3)
	}
	return libc.Int32FromUint32((*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FdictSize)
}

func LZ4_attach_dictionary(tls *libc.TLS, workingStream uintptr, dictionaryStream uintptr) {
	var dictCtx, v1 uintptr
	_, _ = dictCtx, v1
	if dictionaryStream == libc.UintptrFromInt32(0) {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = dictionaryStream
	}
	dictCtx = v1
	if dictCtx != libc.UintptrFromInt32(0) {
		/* If the current offset is zero, we will never look in the
		 * external dictionary context, since there is no value a table
		 * entry can take that indicate a miss. In that case, we need
		 * to bump the offset to something non-zero.
		 */
		if (*(*LZ4_stream_t_internal)(unsafe.Pointer(workingStream))).FcurrentOffset == uint32(0) {
			(*(*LZ4_stream_t_internal)(unsafe.Pointer(workingStream))).FcurrentOffset = libc.Uint32FromInt32(libc.Int32FromInt32(64) * (libc.Int32FromInt32(1) << libc.Int32FromInt32(10)))
		}
		/* Don't actually attach an empty dictionary.
		 */
		if (*LZ4_stream_t_internal)(unsafe.Pointer(dictCtx)).FdictSize == uint32(0) {
			dictCtx = libc.UintptrFromInt32(0)
		}
	}
	(*(*LZ4_stream_t_internal)(unsafe.Pointer(workingStream))).FdictCtx = dictCtx
}

func LZ4_renormDictT(tls *libc.TLS, LZ4_dict uintptr, nextSize int32) {
	var delta U32
	var dictEnd uintptr
	var i int32
	_, _, _ = delta, dictEnd, i
	if (*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).FcurrentOffset+libc.Uint32FromInt32(nextSize) > uint32(0x80000000) { /* potential ptrdiff_t overflow (32-bits mode) */
		/* rescale hash table */
		delta = (*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).FcurrentOffset - libc.Uint32FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)))
		dictEnd = (*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).Fdictionary + uintptr((*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).FdictSize)
		i = 0
		for {
			if !(i < libc.Int32FromInt32(1)<<(libc.Int32FromInt32(LZ4_MEMORY_USAGE_DEFAULT)-libc.Int32FromInt32(2))) {
				break
			}
			if **(**LZ4_u32)(__ccgo_up(LZ4_dict + uintptr(i)*4)) < delta {
				**(**LZ4_u32)(__ccgo_up(LZ4_dict + uintptr(i)*4)) = uint32(0)
			} else {
				**(**LZ4_u32)(__ccgo_up(LZ4_dict + uintptr(i)*4)) -= delta
			}
			goto _1
		_1:
			;
			i = i + 1
		}
		(*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).FcurrentOffset = libc.Uint32FromInt32(libc.Int32FromInt32(64) * (libc.Int32FromInt32(1) << libc.Int32FromInt32(10)))
		if (*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).FdictSize > libc.Uint32FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))) {
			(*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).FdictSize = libc.Uint32FromInt32(libc.Int32FromInt32(64) * (libc.Int32FromInt32(1) << libc.Int32FromInt32(10)))
		}
		(*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).Fdictionary = dictEnd - uintptr((*LZ4_stream_t_internal)(unsafe.Pointer(LZ4_dict)).FdictSize)
	}
}

func LZ4_compress_fast_continue(tls *libc.TLS, LZ4_stream uintptr, source uintptr, dest uintptr, inputSize int32, maxOutputSize int32, acceleration int32) (r int32) {
	var dictEnd, sourceEnd, streamPtr, v1 uintptr
	var result int32
	var tableType tableType_t
	_, _, _, _, _, _ = dictEnd, result, sourceEnd, streamPtr, tableType, v1
	tableType = int32(byU32)
	streamPtr = LZ4_stream
	if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize != 0 {
		v1 = (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).Fdictionary + uintptr((*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize)
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	dictEnd = v1
	LZ4_renormDictT(tls, streamPtr, inputSize) /* fix index overflow */
	if acceleration < int32(1) {
		acceleration = int32(LZ4_ACCELERATION_DEFAULT)
	}
	if acceleration > int32(LZ4_ACCELERATION_MAX) {
		acceleration = int32(LZ4_ACCELERATION_MAX)
	}
	/* invalidate tiny dictionaries */
	if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < uint32(4) && dictEnd != source && inputSize > 0 && (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictCtx == libc.UintptrFromInt32(0) {
		/* remove dictionary existence from history, to employ faster prefix mode */
		(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize = uint32(0)
		(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).Fdictionary = source
		dictEnd = source
	}
	/* Check overlapping input/dictionary space */
	sourceEnd = source + uintptr(inputSize)
	if sourceEnd > (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).Fdictionary && sourceEnd < dictEnd {
		(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize = libc.Uint32FromInt64(int64(dictEnd) - int64(sourceEnd))
		if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize > libc.Uint32FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))) {
			(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize = libc.Uint32FromInt32(libc.Int32FromInt32(64) * (libc.Int32FromInt32(1) << libc.Int32FromInt32(10)))
		}
		if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < uint32(4) {
			(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize = uint32(0)
		}
		(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).Fdictionary = dictEnd - uintptr((*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize)
	}
	/* prefix mode : source data follows dictionary */
	if dictEnd == source {
		if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < libc.Uint32FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))) && (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FcurrentOffset {
			return LZ4_compress_generic(tls, streamPtr, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), tableType, int32(withPrefix64k), int32(dictSmall), acceleration)
		} else {
			return LZ4_compress_generic(tls, streamPtr, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), tableType, int32(withPrefix64k), int32(noDictIssue), acceleration)
		}
	}
	/* external dictionary mode */
	if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictCtx != 0 {
		/* We depend here on the fact that dictCtx'es (produced by
		 * LZ4_loadDict) guarantee that their tables contain no references
		 * to offsets between dictCtx->currentOffset - 64 KB and
		 * dictCtx->currentOffset - dictCtx->dictSize. This makes it safe
		 * to use noDictIssue even when the dict isn't a full 64 KB.
		 */
		if inputSize > libc.Int32FromInt32(4)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)) {
			/* For compressing large blobs, it is faster to pay the setup
			 * cost to copy the dictionary's tables into the active context,
			 * so that the compression loop is only looking into one table.
			 */
			libc.X__builtin_memcpy(tls, streamPtr, (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictCtx, uint64(16416))
			result = LZ4_compress_generic(tls, streamPtr, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), tableType, int32(usingExtDict), int32(noDictIssue), acceleration)
		} else {
			result = LZ4_compress_generic(tls, streamPtr, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), tableType, int32(usingDictCtx), int32(noDictIssue), acceleration)
		}
	} else { /* small data <= 4 KB */
		if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < libc.Uint32FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))) && (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FcurrentOffset {
			result = LZ4_compress_generic(tls, streamPtr, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), tableType, int32(usingExtDict), int32(dictSmall), acceleration)
		} else {
			result = LZ4_compress_generic(tls, streamPtr, source, dest, inputSize, libc.UintptrFromInt32(0), maxOutputSize, int32(limitedOutput), tableType, int32(usingExtDict), int32(noDictIssue), acceleration)
		}
	}
	(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).Fdictionary = source
	(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize = libc.Uint32FromInt32(inputSize)
	return result
	return r
}

// C documentation
//
//	/* Hidden debug function, to force-test external dictionary mode */
func LZ4_compress_forceExtDict(tls *libc.TLS, LZ4_dict uintptr, source uintptr, dest uintptr, srcSize int32) (r int32) {
	var result int32
	var streamPtr uintptr
	_, _ = result, streamPtr
	streamPtr = LZ4_dict
	LZ4_renormDictT(tls, streamPtr, srcSize)
	if (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < libc.Uint32FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))) && (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize < (*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FcurrentOffset {
		result = LZ4_compress_generic(tls, streamPtr, source, dest, srcSize, libc.UintptrFromInt32(0), 0, int32(notLimited), int32(byU32), int32(usingExtDict), int32(dictSmall), int32(1))
	} else {
		result = LZ4_compress_generic(tls, streamPtr, source, dest, srcSize, libc.UintptrFromInt32(0), 0, int32(notLimited), int32(byU32), int32(usingExtDict), int32(noDictIssue), int32(1))
	}
	(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).Fdictionary = source
	(*LZ4_stream_t_internal)(unsafe.Pointer(streamPtr)).FdictSize = libc.Uint32FromInt32(srcSize)
	return result
}

// C documentation
//
//	/*! LZ4_saveDict() :
//	 *  If previously compressed data block is not guaranteed to remain available at its memory location,
//	 *  save it into a safer place (char* safeBuffer).
//	 *  Note : no need to call LZ4_loadDict() afterwards, dictionary is immediately usable,
//	 *         one can therefore call LZ4_compress_fast_continue() right after.
//	 * @return : saved dictionary size in bytes (necessarily <= dictSize), or 0 if error.
//	 */
func LZ4_saveDict(tls *libc.TLS, LZ4_dict uintptr, safeBuffer uintptr, dictSize int32) (r int32) {
	var dict, previousDictEnd uintptr
	_, _ = dict, previousDictEnd
	dict = LZ4_dict
	if libc.Uint32FromInt32(dictSize) > libc.Uint32FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))) {
		dictSize = libc.Int32FromInt32(64) * (libc.Int32FromInt32(1) << libc.Int32FromInt32(10))
	} /* useless to define a dictionary > 64 KB */
	if libc.Uint32FromInt32(dictSize) > (*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FdictSize {
		dictSize = libc.Int32FromUint32((*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FdictSize)
	}
	if safeBuffer == libc.UintptrFromInt32(0) {
	}
	if dictSize > 0 {
		previousDictEnd = (*LZ4_stream_t_internal)(unsafe.Pointer(dict)).Fdictionary + uintptr((*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FdictSize)
		iqlibc.__builtin_memmove(tls, safeBuffer, previousDictEnd-uintptr(dictSize), libc.Uint64FromInt32(dictSize))
	}
	(*LZ4_stream_t_internal)(unsafe.Pointer(dict)).Fdictionary = safeBuffer
	(*LZ4_stream_t_internal)(unsafe.Pointer(dict)).FdictSize = libc.Uint32FromInt32(dictSize)
	return dictSize
}

/*-*******************************
 *  Decompression functions
 ********************************/

type earlyEnd_directive = int32

const decode_full_block = 0
const partial_decode = 1

// C documentation
//
//	/* variant for decompress_unsafe()
//	 * does not know end of input
//	 * presumes input is well formed
//	 * note : will consume at least one byte */
func read_long_length_no_check(tls *libc.TLS, pp uintptr) (r size_t) {
	var b, l size_t
	_, _ = b, l
	l = uint64(0)
	for cond := true; cond; cond = b == uint64(255) {
		b = uint64(**(**BYTE)(__ccgo_up(**(**uintptr)(__ccgo_up(pp)))))
		**(**uintptr)(__ccgo_up(pp)) = **(**uintptr)(__ccgo_up(pp)) + 1
		l = l + b
	}
	return l
}

// C documentation
//
//	/* core decoder variant for LZ4_decompress_fast*()
//	 * for legacy support only : these entry points are deprecated.
//	 * - Presumes input is correctly formed (no defense vs malformed inputs)
//	 * - Does not know input size (presume input buffer is "large enough")
//	 * - Decompress a full block (only)
//	 * @return : nb of bytes read from input.
//	 * Note : this variant is not optimized for speed, just for maintenance.
//	 *        the goal is to remove support of decompress_fast*() variants by v2.0
//	**/
func LZ4_decompress_unsafe_generic(tls *libc.TLS, istart uintptr, ostart uintptr, decompressedSize int32, prefixSize size_t, dictStart uintptr, dictSize size_t) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var dictEnd, extMatch, match, oend, op, prefixStart, v1 uintptr
	var extml, ll, ml, offset, u size_t
	var token uint32
	var _ /* ip at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = dictEnd, extMatch, extml, ll, match, ml, oend, offset, op, prefixStart, token, u, v1
	**(**uintptr)(__ccgo_up(bp)) = istart
	op = ostart
	oend = ostart + uintptr(decompressedSize)
	prefixStart = ostart - uintptr(prefixSize)
	if dictStart == libc.UintptrFromInt32(0) {
	}
	for int32(1) != 0 {
		v1 = **(**uintptr)(__ccgo_up(bp))
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
		/* start new sequence */
		token = uint32(**(**BYTE)(__ccgo_up(v1)))
		/* literals */
		ll = uint64(token >> int32(ML_BITS))
		if ll == uint64(15) {
			/* long literal length */
			ll = ll + read_long_length_no_check(tls, bp)
		}
		if libc.Uint64FromInt64(int64(oend)-int64(op)) < ll {
			return -int32(1)
		} /* output buffer overflow */
		iqlibc.__builtin_memmove(tls, op, **(**uintptr)(__ccgo_up(bp)), ll) /* support in-place decompression */
		op = op + uintptr(ll)
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(ll)
		if libc.Uint64FromInt64(int64(oend)-int64(op)) < uint64(MFLIMIT) {
			if op == oend {
				break
			} /* end of block */
			/* incorrect end of block :
			 * last match must start at least MFLIMIT==12 bytes before end of output block */
			return -int32(1)
		}
		/* match */
		ml = uint64(token & uint32(15))
		offset = uint64(LZ4_readLE16(tls, **(**uintptr)(__ccgo_up(bp))))
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(2)
		if ml == uint64(15) {
			/* long literal length */
			ml = ml + read_long_length_no_check(tls, bp)
		}
		ml = ml + uint64(MINMATCH)
		if libc.Uint64FromInt64(int64(oend)-int64(op)) < ml {
			return -int32(1)
		} /* output buffer overflow */
		match = op - uintptr(offset)
		/* out of range */
		if offset > libc.Uint64FromInt64(int64(op)-int64(prefixStart))+dictSize {
			return -int32(1)
		}
		/* check special case : extDict */
		if offset > libc.Uint64FromInt64(int64(op)-int64(prefixStart)) {
			/* extDict scenario */
			dictEnd = dictStart + uintptr(dictSize)
			extMatch = dictEnd - uintptr(offset-libc.Uint64FromInt64(int64(op)-int64(prefixStart)))
			extml = libc.Uint64FromInt64(int64(dictEnd) - int64(extMatch))
			if extml > ml {
				/* match entirely within extDict */
				iqlibc.__builtin_memmove(tls, op, extMatch, ml)
				op = op + uintptr(ml)
				ml = uint64(0)
			} else {
				/* match split between extDict & prefix */
				iqlibc.__builtin_memmove(tls, op, extMatch, extml)
				op = op + uintptr(extml)
				ml = ml - extml
			}
			match = prefixStart
		}
		/* match copy - slow variant, supporting overlap copy */
		u = uint64(0)
		for {
			if !(u < ml) {
				break
			}
			**(**BYTE)(__ccgo_up(op + uintptr(u))) = **(**BYTE)(__ccgo_up(match + uintptr(u)))
			goto _2
		_2:
			;
			u = u + 1
		}
		op = op + uintptr(ml)
		if libc.Uint64FromInt64(int64(oend)-int64(op)) < uint64(LASTLITERALS) {
			/* incorrect end of block :
			 * last match must stop at least LASTLITERALS==5 bytes before end of output block */
			return -int32(1)
		}
		/* match */
	} /* main loop */
	return int32(int64(**(**uintptr)(__ccgo_up(bp))) - int64(istart))
}

// C documentation
//
//	/* Read the variable-length literal or match length.
//	 *
//	 * @ip : input pointer
//	 * @ilimit : position after which if length is not decoded, the input is necessarily corrupted.
//	 * @initial_check - check ip >= ipmax before start of loop.  Returns initial_error if so.
//	 * @error (output) - error code.  Must be set to 0 before call.
//	**/
type Rvl_t = uint64

var rvl_error = libc.Uint64FromInt32(-libc.Int32FromInt32(1))

func read_variable_length(tls *libc.TLS, ip uintptr, ilimit uintptr, initial_check int32) (r Rvl_t) {
	var length, s Rvl_t
	_, _ = length, s
	length = uint64(0)
	if initial_check != 0 && libc.BoolInt64(libc.BoolInt32(**(**uintptr)(__ccgo_up(ip)) >= ilimit) != libc.Int32FromInt32(0)) != 0 { /* read limit reached */
		return rvl_error
	}
	for cond := true; cond; cond = s == uint64(255) {
		s = uint64(**(**BYTE)(__ccgo_up(**(**uintptr)(__ccgo_up(ip)))))
		**(**uintptr)(__ccgo_up(ip)) = **(**uintptr)(__ccgo_up(ip)) + 1
		length = length + s
		if libc.BoolInt64(libc.BoolInt32(**(**uintptr)(__ccgo_up(ip)) > ilimit) != libc.Int32FromInt32(0)) != 0 { /* read limit reached */
			return rvl_error
		}
		/* accumulator overflow detection (32-bit mode only) */
		if libc.Bool(uint64(8) < uint64(8)) && libc.BoolInt64(libc.BoolInt32(length > libc.Uint64FromInt32(-libc.Int32FromInt32(1))/libc.Uint64FromInt32(2)) != libc.Int32FromInt32(0)) != 0 {
			return rvl_error
		}
	}
	return length
}

// C documentation
//
//	/*! LZ4_decompress_generic() :
//	 *  This generic decompression function covers all use cases.
//	 *  It shall be instantiated several times, using different sets of directives.
//	 *  Note that it is important for performance that this function really get inlined,
//	 *  in order to remove useless branches during compilation optimization.
//	 */
func LZ4_decompress_generic(tls *libc.TLS, src uintptr, dst uintptr, srcSize int32, outputSize int32, partialDecoding earlyEnd_directive, dict dict_directive, lowPrefix uintptr, dictStart uintptr, dictSize size_t) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var addl, addl1, addl2, addl3, copySize, copySize1, length, mlen, offset, restSize, restSize1 size_t
	var checkOffset, v2 int32
	var copyEnd, copyFrom, copyFrom1, cpy, dictEnd, endOfMatch, endOfMatch1, iend, match, matchEnd, oCopyLimit, oend, op, shortiend, shortoend, v1, v3 uintptr
	var token uint32
	var v4 uint64
	var _ /* ip at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = addl, addl1, addl2, addl3, checkOffset, copyEnd, copyFrom, copyFrom1, copySize, copySize1, cpy, dictEnd, endOfMatch, endOfMatch1, iend, length, match, matchEnd, mlen, oCopyLimit, oend, offset, op, restSize, restSize1, shortiend, shortoend, token, v1, v2, v3, v4
	if src == libc.UintptrFromInt32(0) || outputSize < 0 {
		return -int32(1)
	}
	**(**uintptr)(__ccgo_up(bp)) = src
	iend = **(**uintptr)(__ccgo_up(bp)) + uintptr(srcSize)
	op = dst
	oend = op + uintptr(outputSize)
	if dictStart == libc.UintptrFromInt32(0) {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = dictStart + uintptr(dictSize)
	}
	dictEnd = v1
	checkOffset = libc.BoolInt32(dictSize < libc.Uint64FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))))
	/* Set up the "end" pointers for the shortcut. */
	shortiend = iend - uintptr(14) - uintptr(2)
	shortoend = oend - uintptr(14) - uintptr(18)
	/* Special cases */
	if libc.BoolInt64(libc.BoolInt32(outputSize == libc.Int32FromInt32(0)) != libc.Int32FromInt32(0)) != 0 {
		/* Empty output buffer */
		if partialDecoding != 0 {
			return 0
		}
		if srcSize == int32(1) && libc.Int32FromUint8(**(**BYTE)(__ccgo_up(**(**uintptr)(__ccgo_up(bp))))) == 0 {
			v2 = 0
		} else {
			v2 = -int32(1)
		}
		return v2
	}
	if libc.BoolInt64(libc.BoolInt32(srcSize == libc.Int32FromInt32(0)) != libc.Int32FromInt32(0)) != 0 {
		return -int32(1)
	}
	/* LZ4_FAST_DEC_LOOP:
	 * designed for modern OoO performance cpus,
	 * where copying reliably 32-bytes is preferable to an unpredictable branch.
	 * note : fast loop may show a regression for some client arm chips. */
	if int64(oend)-int64(op) < int64(FASTLOOP_SAFE_DISTANCE) {
		goto safe_decode
	}
	/* Fast loop : decode sequences as long as output < oend-FASTLOOP_SAFE_DISTANCE */
	for int32(1) != 0 {
		/* Main fastloop assertion: We can always wildcopy FASTLOOP_SAFE_DISTANCE */
		v1 = **(**uintptr)(__ccgo_up(bp))
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
		token = uint32(**(**BYTE)(__ccgo_up(v1)))
		length = uint64(token >> int32(ML_BITS)) /* literal length */
		/* decode literal length */
		if length == uint64(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)) {
			addl = read_variable_length(tls, bp, iend-uintptr(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)), int32(1))
			if addl == rvl_error {
				goto _output_error
			}
			length = length + addl
			if libc.BoolInt64(libc.BoolInt32(uint64(op)+length < uint64(op)) != libc.Int32FromInt32(0)) != 0 {
				goto _output_error
			} /* overflow detection */
			if libc.BoolInt64(libc.BoolInt32(uint64(**(**uintptr)(__ccgo_up(bp)))+length < uint64(**(**uintptr)(__ccgo_up(bp)))) != libc.Int32FromInt32(0)) != 0 {
				goto _output_error
			} /* overflow detection */
			/* copy literals */
			cpy = op + uintptr(length)
			if cpy > oend-uintptr(32) || **(**uintptr)(__ccgo_up(bp))+uintptr(length) > iend-uintptr(32) {
				goto safe_literal_copy
			}
			LZ4_wildCopy32(tls, op, **(**uintptr)(__ccgo_up(bp)), cpy)
			**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(length)
			op = cpy
		} else {
			cpy = op + uintptr(length)
			/* We don't need to check oend, since we check it once for each loop below */
			if **(**uintptr)(__ccgo_up(bp)) > iend-uintptr(libc.Int32FromInt32(16)+libc.Int32FromInt32(1)) {
				goto safe_literal_copy
			}
			/* Literals can only be <= 14, but hope compilers optimize better when copy by a register size */
			libc.X__builtin_memcpy(tls, op, **(**uintptr)(__ccgo_up(bp)), uint64(16))
			**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(length)
			op = cpy
		}
		/* get offset */
		offset = uint64(LZ4_readLE16(tls, **(**uintptr)(__ccgo_up(bp))))
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(2)
		match = op - uintptr(offset)
		/* overflow check */
		/* get matchlength */
		length = uint64(token & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS) - libc.Uint32FromInt32(1)))
		if length == uint64(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS)-libc.Uint32FromInt32(1)) {
			addl1 = read_variable_length(tls, bp, iend-uintptr(LASTLITERALS)+uintptr(1), 0)
			if addl1 == rvl_error {
				goto _output_error
			}
			length = length + addl1
			length = length + uint64(MINMATCH)
			if libc.BoolInt64(libc.BoolInt32(uint64(op)+length < uint64(op)) != libc.Int32FromInt32(0)) != 0 {
				goto _output_error
			} /* overflow detection */
			if checkOffset != 0 && libc.BoolInt64(libc.BoolInt32(match+uintptr(dictSize) < lowPrefix) != libc.Int32FromInt32(0)) != 0 {
				goto _output_error
			} /* Error : offset outside buffers */
			if op+uintptr(length) >= oend-uintptr(FASTLOOP_SAFE_DISTANCE) {
				goto safe_match_copy
			}
		} else {
			length = length + uint64(MINMATCH)
			if op+uintptr(length) >= oend-uintptr(FASTLOOP_SAFE_DISTANCE) {
				goto safe_match_copy
			}
			/* Fastpath check: skip LZ4_wildCopy32 when true */
			if dict == int32(withPrefix64k) || match >= lowPrefix {
				if offset >= uint64(8) {
					libc.X__builtin_memcpy(tls, op, match, uint64(8))
					libc.X__builtin_memcpy(tls, op+uintptr(8), match+uintptr(8), uint64(8))
					libc.X__builtin_memcpy(tls, op+uintptr(16), match+uintptr(16), uint64(2))
					op = op + uintptr(length)
					continue
				}
			}
		}
		if checkOffset != 0 && libc.BoolInt64(libc.BoolInt32(match+uintptr(dictSize) < lowPrefix) != libc.Int32FromInt32(0)) != 0 {
			goto _output_error
		} /* Error : offset outside buffers */
		/* match starting within external dictionary */
		if dict == int32(usingExtDict) && match < lowPrefix {
			if libc.BoolInt64(libc.BoolInt32(op+uintptr(length) > oend-libc.UintptrFromInt32(LASTLITERALS)) != libc.Int32FromInt32(0)) != 0 {
				if partialDecoding != 0 {
					if length < libc.Uint64FromInt64(int64(oend)-int64(op)) {
						v4 = length
					} else {
						v4 = libc.Uint64FromInt64(int64(oend) - int64(op))
					}
					length = v4
				} else {
					goto _output_error /* end-of-block condition violated */
				}
			}
			if length <= libc.Uint64FromInt64(int64(lowPrefix)-int64(match)) {
				/* match fits entirely within external dictionary : just copy */
				iqlibc.__builtin_memmove(tls, op, dictEnd-uintptr(int64(lowPrefix)-int64(match)), length)
				op = op + uintptr(length)
			} else {
				/* match stretches into both external dictionary and current block */
				copySize = libc.Uint64FromInt64(int64(lowPrefix) - int64(match))
				restSize = length - copySize
				libc.X__builtin_memcpy(tls, op, dictEnd-uintptr(copySize), copySize)
				op = op + uintptr(copySize)
				if restSize > libc.Uint64FromInt64(int64(op)-int64(lowPrefix)) { /* overlap copy */
					endOfMatch = op + uintptr(restSize)
					copyFrom = lowPrefix
					for op < endOfMatch {
						v1 = op
						op = op + 1
						v3 = copyFrom
						copyFrom = copyFrom + 1
						**(**BYTE)(__ccgo_up(v1)) = **(**BYTE)(__ccgo_up(v3))
					}
				} else {
					libc.X__builtin_memcpy(tls, op, lowPrefix, restSize)
					op = op + uintptr(restSize)
				}
			}
			continue
		}
		/* copy match within block */
		cpy = op + uintptr(length)
		if libc.BoolInt64(libc.BoolInt32(offset < libc.Uint64FromInt32(16)) != libc.Int32FromInt32(0)) != 0 {
			LZ4_memcpy_using_offset(tls, op, match, cpy, offset)
		} else {
			LZ4_wildCopy32(tls, op, match, cpy)
		}
		op = cpy /* wildcopy correction */
	}
	goto safe_decode
safe_decode:
	;
	/* Main Loop : decode remaining sequences where output < FASTLOOP_SAFE_DISTANCE */
_8:
	;
	if !(int32(1) != 0) {
		goto _7
	}
	v1 = **(**uintptr)(__ccgo_up(bp))
	**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + 1
	token = uint32(**(**BYTE)(__ccgo_up(v1)))
	length = uint64(token >> int32(ML_BITS)) /* literal length */
	/* A two-stage shortcut for the most common case:
	 * 1) If the literal length is 0..14, and there is enough space,
	 * enter the shortcut and copy 16 bytes on behalf of the literals
	 * (in the fast mode, only 8 bytes can be safely copied this way).
	 * 2) Further if the match length is 4..18, copy 18 bytes in a similar
	 * manner; but we ensure that there's enough space in the output for
	 * those 18 bytes earlier, upon entering the shortcut (in other words,
	 * there is a combined check for both stages).
	 */
	if length != uint64(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)) && libc.BoolInt64(libc.BoolInt32(**(**uintptr)(__ccgo_up(bp)) < shortiend)&libc.BoolInt32(op <= shortoend) != libc.Int32FromInt32(0)) != 0 {
		/* Copy the literals */
		libc.X__builtin_memcpy(tls, op, **(**uintptr)(__ccgo_up(bp)), uint64(16))
		op = op + uintptr(length)
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(length)
		/* The second stage: prepare for match copying, decode full info.
		 * If it doesn't work out, the info won't be wasted. */
		length = uint64(token & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS) - libc.Uint32FromInt32(1))) /* match length */
		offset = uint64(LZ4_readLE16(tls, **(**uintptr)(__ccgo_up(bp))))
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(2)
		match = op - uintptr(offset)
		/* check overflow */
		/* Do not deal with overlapping matches. */
		if length != uint64(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS)-libc.Uint32FromInt32(1)) && offset >= uint64(8) && (dict == int32(withPrefix64k) || match >= lowPrefix) {
			/* Copy the match. */
			libc.X__builtin_memcpy(tls, op+uintptr(0), match+uintptr(0), uint64(8))
			libc.X__builtin_memcpy(tls, op+uintptr(8), match+uintptr(8), uint64(8))
			libc.X__builtin_memcpy(tls, op+uintptr(16), match+uintptr(16), uint64(2))
			op = op + uintptr(length+uint64(MINMATCH))
			/* Both stages worked, load the next token. */
			goto _8
		}
		/* The second stage didn't work out, but the info is ready.
		 * Propel it right to the point of match copying. */
		goto _copy_match
	}
	/* decode literal length */
	if length == uint64(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)) {
		addl2 = read_variable_length(tls, bp, iend-uintptr(libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(ML_BITS))-libc.Uint32FromInt32(1)), int32(1))
		if addl2 == rvl_error {
			goto _output_error
		}
		length = length + addl2
		if libc.BoolInt64(libc.BoolInt32(uint64(op)+length < uint64(op)) != libc.Int32FromInt32(0)) != 0 {
			goto _output_error
		} /* overflow detection */
		if libc.BoolInt64(libc.BoolInt32(uint64(**(**uintptr)(__ccgo_up(bp)))+length < uint64(**(**uintptr)(__ccgo_up(bp)))) != libc.Int32FromInt32(0)) != 0 {
			goto _output_error
		} /* overflow detection */
	}
	/* copy literals */
	cpy = op + uintptr(length)
	goto safe_literal_copy
safe_literal_copy:
	;
	if cpy > oend-uintptr(MFLIMIT) || **(**uintptr)(__ccgo_up(bp))+uintptr(length) > iend-uintptr(libc.Int32FromInt32(2)+libc.Int32FromInt32(1)+libc.Int32FromInt32(LASTLITERALS)) {
		/* We've either hit the input parsing restriction or the output parsing restriction.
		 * In the normal scenario, decoding a full block, it must be the last sequence,
		 * otherwise it's an error (invalid input or dimensions).
		 * In partialDecoding scenario, it's necessary to ensure there is no buffer overflow.
		 */
		if partialDecoding != 0 {
			/* Since we are partial decoding we may be in this block because of the output parsing
			 * restriction, which is not valid since the output buffer is allowed to be undersized.
			 */
			/* Finishing in the middle of a literals segment,
			 * due to lack of input.
			 */
			if **(**uintptr)(__ccgo_up(bp))+uintptr(length) > iend {
				length = libc.Uint64FromInt64(int64(iend) - int64(**(**uintptr)(__ccgo_up(bp))))
				cpy = op + uintptr(length)
			}
			/* Finishing in the middle of a literals segment,
			 * due to lack of output space.
			 */
			if cpy > oend {
				cpy = oend
				length = libc.Uint64FromInt64(int64(oend) - int64(op))
			}
		} else {
			/* We must be on the last sequence (or invalid) because of the parsing limitations
			 * so check that we exactly consume the input and don't overrun the output buffer.
			 */
			if **(**uintptr)(__ccgo_up(bp))+uintptr(length) != iend || cpy > oend {
				goto _output_error
			}
		}
		iqlibc.__builtin_memmove(tls, op, **(**uintptr)(__ccgo_up(bp)), length) /* supports overlapping memory regions, for in-place decompression scenarios */
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(length)
		op = op + uintptr(length)
		/* Necessarily EOF when !partialDecoding.
		 * When partialDecoding, it is EOF if we've either
		 * filled the output buffer or
		 * can't proceed with reading an offset for following match.
		 */
		if !(partialDecoding != 0) || cpy == oend || **(**uintptr)(__ccgo_up(bp)) >= iend-libc.UintptrFromInt32(2) {
			goto _7
		}
	} else {
		LZ4_wildCopy8(tls, op, **(**uintptr)(__ccgo_up(bp)), cpy) /* can overwrite up to 8 bytes beyond cpy */
		**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(length)
		op = cpy
	}
	/* get offset */
	offset = uint64(LZ4_readLE16(tls, **(**uintptr)(__ccgo_up(bp))))
	**(**uintptr)(__ccgo_up(bp)) = **(**uintptr)(__ccgo_up(bp)) + uintptr(2)
	match = op - uintptr(offset)
	/* get matchlength */
	length = uint64(token & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS) - libc.Uint32FromInt32(1)))
	goto _copy_match
_copy_match:
	;
	if length == uint64(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(ML_BITS)-libc.Uint32FromInt32(1)) {
		addl3 = read_variable_length(tls, bp, iend-uintptr(LASTLITERALS)+uintptr(1), 0)
		if addl3 == rvl_error {
			goto _output_error
		}
		length = length + addl3
		if libc.BoolInt64(libc.BoolInt32(uint64(op)+length < uint64(op)) != libc.Int32FromInt32(0)) != 0 {
			goto _output_error
		} /* overflow detection */
	}
	length = length + uint64(MINMATCH)
	goto safe_match_copy
safe_match_copy:
	;
	if checkOffset != 0 && libc.BoolInt64(libc.BoolInt32(match+uintptr(dictSize) < lowPrefix) != libc.Int32FromInt32(0)) != 0 {
		goto _output_error
	} /* Error : offset outside buffers */
	/* match starting within external dictionary */
	if dict == int32(usingExtDict) && match < lowPrefix {
		if libc.BoolInt64(libc.BoolInt32(op+uintptr(length) > oend-libc.UintptrFromInt32(LASTLITERALS)) != libc.Int32FromInt32(0)) != 0 {
			if partialDecoding != 0 {
				if length < libc.Uint64FromInt64(int64(oend)-int64(op)) {
					v4 = length
				} else {
					v4 = libc.Uint64FromInt64(int64(oend) - int64(op))
				}
				length = v4
			} else {
				goto _output_error
			} /* doesn't respect parsing restriction */
		}
		if length <= libc.Uint64FromInt64(int64(lowPrefix)-int64(match)) {
			/* match fits entirely within external dictionary : just copy */
			iqlibc.__builtin_memmove(tls, op, dictEnd-uintptr(int64(lowPrefix)-int64(match)), length)
			op = op + uintptr(length)
		} else {
			/* match stretches into both external dictionary and current block */
			copySize1 = libc.Uint64FromInt64(int64(lowPrefix) - int64(match))
			restSize1 = length - copySize1
			libc.X__builtin_memcpy(tls, op, dictEnd-uintptr(copySize1), copySize1)
			op = op + uintptr(copySize1)
			if restSize1 > libc.Uint64FromInt64(int64(op)-int64(lowPrefix)) { /* overlap copy */
				endOfMatch1 = op + uintptr(restSize1)
				copyFrom1 = lowPrefix
				for op < endOfMatch1 {
					v1 = op
					op = op + 1
					v3 = copyFrom1
					copyFrom1 = copyFrom1 + 1
					**(**BYTE)(__ccgo_up(v1)) = **(**BYTE)(__ccgo_up(v3))
				}
			} else {
				libc.X__builtin_memcpy(tls, op, lowPrefix, restSize1)
				op = op + uintptr(restSize1)
			}
		}
		goto _8
	}
	/* copy match within block */
	cpy = op + uintptr(length)
	/* partialDecoding : may end anywhere within the block */
	if partialDecoding != 0 && cpy > oend-uintptr(libc.Int32FromInt32(2)*libc.Int32FromInt32(WILDCOPYLENGTH)-libc.Int32FromInt32(MINMATCH)) {
		if length < libc.Uint64FromInt64(int64(oend)-int64(op)) {
			v4 = length
		} else {
			v4 = libc.Uint64FromInt64(int64(oend) - int64(op))
		}
		mlen = v4
		matchEnd = match + uintptr(mlen)
		copyEnd = op + uintptr(mlen)
		if matchEnd > op { /* overlap copy */
			for op < copyEnd {
				v1 = op
				op = op + 1
				v3 = match
				match = match + 1
				**(**BYTE)(__ccgo_up(v1)) = **(**BYTE)(__ccgo_up(v3))
			}
		} else {
			libc.X__builtin_memcpy(tls, op, match, mlen)
		}
		op = copyEnd
		if op == oend {
			goto _7
		}
		goto _8
	}
	if libc.BoolInt64(libc.BoolInt32(offset < libc.Uint64FromInt32(8)) != libc.Int32FromInt32(0)) != 0 {
		LZ4_write32(tls, op, uint32(0)) /* silence msan warning when offset==0 */
		**(**BYTE)(__ccgo_up(op)) = **(**BYTE)(__ccgo_up(match))
		**(**BYTE)(__ccgo_up(op + 1)) = **(**BYTE)(__ccgo_up(match + 1))
		**(**BYTE)(__ccgo_up(op + 2)) = **(**BYTE)(__ccgo_up(match + 2))
		**(**BYTE)(__ccgo_up(op + 3)) = **(**BYTE)(__ccgo_up(match + 3))
		match = match + uintptr(inc32table[offset])
		libc.X__builtin_memcpy(tls, op+uintptr(4), match, uint64(4))
		match = match - uintptr(dec64table[offset])
	} else {
		libc.X__builtin_memcpy(tls, op, match, uint64(8))
		match = match + uintptr(8)
	}
	op = op + uintptr(8)
	if libc.BoolInt64(libc.BoolInt32(cpy > oend-uintptr(libc.Int32FromInt32(2)*libc.Int32FromInt32(WILDCOPYLENGTH)-libc.Int32FromInt32(MINMATCH))) != libc.Int32FromInt32(0)) != 0 {
		oCopyLimit = oend - uintptr(libc.Int32FromInt32(WILDCOPYLENGTH)-libc.Int32FromInt32(1))
		if cpy > oend-uintptr(LASTLITERALS) {
			goto _output_error
		} /* Error : last LASTLITERALS bytes must be literals (uncompressed) */
		if op < oCopyLimit {
			LZ4_wildCopy8(tls, op, match, oCopyLimit)
			match = match + uintptr(int64(oCopyLimit)-int64(op))
			op = oCopyLimit
		}
		for op < cpy {
			v1 = op
			op = op + 1
			v3 = match
			match = match + 1
			**(**BYTE)(__ccgo_up(v1)) = **(**BYTE)(__ccgo_up(v3))
		}
	} else {
		libc.X__builtin_memcpy(tls, op, match, uint64(8))
		if length > uint64(16) {
			LZ4_wildCopy8(tls, op+uintptr(8), match+uintptr(8), cpy)
		}
	}
	op = cpy /* wildcopy correction */
	goto _8
_7:
	;
	/* end of decoding */
	return int32(int64(op) - int64(dst)) /* Nb of output bytes decoded */
	/* Overflow error detected */
	goto _output_error
_output_error:
	;
	return int32(-(int64(**(**uintptr)(__ccgo_up(bp))) - int64(src))) - int32(1)
	return r
}

/*===== Instantiate the API decoding functions. =====*/

func LZ4_decompress_safe(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, maxDecompressedSize int32) (r int32) {
	return LZ4_decompress_generic(tls, source, dest, compressedSize, maxDecompressedSize, int32(decode_full_block), int32(noDict), dest, libc.UintptrFromInt32(0), uint64(0))
}

func LZ4_decompress_safe_partial(tls *libc.TLS, src uintptr, dst uintptr, compressedSize int32, targetOutputSize int32, dstCapacity int32) (r int32) {
	var v1 int32
	_ = v1
	if targetOutputSize < dstCapacity {
		v1 = targetOutputSize
	} else {
		v1 = dstCapacity
	}
	dstCapacity = v1
	return LZ4_decompress_generic(tls, src, dst, compressedSize, dstCapacity, int32(partial_decode), int32(noDict), dst, libc.UintptrFromInt32(0), uint64(0))
}

func LZ4_decompress_fast(tls *libc.TLS, source uintptr, dest uintptr, originalSize int32) (r int32) {
	return LZ4_decompress_unsafe_generic(tls, source, dest, originalSize, uint64(0), libc.UintptrFromInt32(0), uint64(0))
}

/*===== Instantiate a few more decoding cases, used more than once. =====*/

// C documentation
//
//	/* Exported, an obsolete API function. */
func LZ4_decompress_safe_withPrefix64k(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, maxOutputSize int32) (r int32) {
	return LZ4_decompress_generic(tls, source, dest, compressedSize, maxOutputSize, int32(decode_full_block), int32(withPrefix64k), dest-uintptr(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))), libc.UintptrFromInt32(0), uint64(0))
}

func LZ4_decompress_safe_partial_withPrefix64k(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, targetOutputSize int32, dstCapacity int32) (r int32) {
	var v1 int32
	_ = v1
	if targetOutputSize < dstCapacity {
		v1 = targetOutputSize
	} else {
		v1 = dstCapacity
	}
	dstCapacity = v1
	return LZ4_decompress_generic(tls, source, dest, compressedSize, dstCapacity, int32(partial_decode), int32(withPrefix64k), dest-uintptr(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))), libc.UintptrFromInt32(0), uint64(0))
}

// C documentation
//
//	/* Another obsolete API function, paired with the previous one. */
func LZ4_decompress_fast_withPrefix64k(tls *libc.TLS, source uintptr, dest uintptr, originalSize int32) (r int32) {
	return LZ4_decompress_unsafe_generic(tls, source, dest, originalSize, libc.Uint64FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))), libc.UintptrFromInt32(0), uint64(0))
}

func LZ4_decompress_safe_withSmallPrefix(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, maxOutputSize int32, prefixSize size_t) (r int32) {
	return LZ4_decompress_generic(tls, source, dest, compressedSize, maxOutputSize, int32(decode_full_block), int32(noDict), dest-uintptr(prefixSize), libc.UintptrFromInt32(0), uint64(0))
}

func LZ4_decompress_safe_partial_withSmallPrefix(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, targetOutputSize int32, dstCapacity int32, prefixSize size_t) (r int32) {
	var v1 int32
	_ = v1
	if targetOutputSize < dstCapacity {
		v1 = targetOutputSize
	} else {
		v1 = dstCapacity
	}
	dstCapacity = v1
	return LZ4_decompress_generic(tls, source, dest, compressedSize, dstCapacity, int32(partial_decode), int32(noDict), dest-uintptr(prefixSize), libc.UintptrFromInt32(0), uint64(0))
}

func LZ4_decompress_safe_forceExtDict(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, maxOutputSize int32, dictStart uintptr, dictSize size_t) (r int32) {
	return LZ4_decompress_generic(tls, source, dest, compressedSize, maxOutputSize, int32(decode_full_block), int32(usingExtDict), dest, dictStart, dictSize)
}

func LZ4_decompress_safe_partial_forceExtDict(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, targetOutputSize int32, dstCapacity int32, dictStart uintptr, dictSize size_t) (r int32) {
	var v1 int32
	_ = v1
	if targetOutputSize < dstCapacity {
		v1 = targetOutputSize
	} else {
		v1 = dstCapacity
	}
	dstCapacity = v1
	return LZ4_decompress_generic(tls, source, dest, compressedSize, dstCapacity, int32(partial_decode), int32(usingExtDict), dest, dictStart, dictSize)
}

func LZ4_decompress_fast_extDict(tls *libc.TLS, source uintptr, dest uintptr, originalSize int32, dictStart uintptr, dictSize size_t) (r int32) {
	return LZ4_decompress_unsafe_generic(tls, source, dest, originalSize, uint64(0), dictStart, dictSize)
}

// C documentation
//
//	/* The "double dictionary" mode, for use with e.g. ring buffers: the first part
//	 * of the dictionary is passed as prefix, and the second via dictStart + dictSize.
//	 * These routines are used only once, in LZ4_decompress_*_continue().
//	 */
func LZ4_decompress_safe_doubleDict(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, maxOutputSize int32, prefixSize size_t, dictStart uintptr, dictSize size_t) (r int32) {
	return LZ4_decompress_generic(tls, source, dest, compressedSize, maxOutputSize, int32(decode_full_block), int32(usingExtDict), dest-uintptr(prefixSize), dictStart, dictSize)
}

/*===== streaming decompression functions =====*/

func LZ4_createStreamDecode(tls *libc.TLS) (r uintptr) {
	return libc.Xcalloc(tls, uint64(1), uint64(32))
}

func LZ4_freeStreamDecode(tls *libc.TLS, LZ4_stream uintptr) (r int32) {
	if LZ4_stream == libc.UintptrFromInt32(0) {
		return 0
	} /* support free on NULL */
	libc.Xfree(tls, LZ4_stream)
	return 0
}

// C documentation
//
//	/*! LZ4_setStreamDecode() :
//	 *  Use this function to instruct where to find the dictionary.
//	 *  This function is not necessary if previous data is still available where it was decoded.
//	 *  Loading a size of 0 is allowed (same effect as no dictionary).
//	 * @return : 1 if OK, 0 if error
//	 */
func LZ4_setStreamDecode(tls *libc.TLS, LZ4_streamDecode uintptr, dictionary uintptr, dictSize int32) (r int32) {
	var lz4sd uintptr
	_ = lz4sd
	lz4sd = LZ4_streamDecode
	(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize = libc.Uint64FromInt32(dictSize)
	if dictSize != 0 {
		(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd = dictionary + uintptr(dictSize)
	} else {
		(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd = dictionary
	}
	(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FexternalDict = libc.UintptrFromInt32(0)
	(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize = uint64(0)
	return int32(1)
}

// C documentation
//
//	/*! LZ4_decoderRingBufferSize() :
//	 *  when setting a ring buffer for streaming decompression (optional scenario),
//	 *  provides the minimum size of this ring buffer
//	 *  to be compatible with any source respecting maxBlockSize condition.
//	 *  Note : in a ring buffer scenario,
//	 *  blocks are presumed decompressed next to each other.
//	 *  When not enough space remains for next block (remainingSize < maxBlockSize),
//	 *  decoding resumes from beginning of ring buffer.
//	 * @return : minimum ring buffer size,
//	 *           or 0 if there is an error (invalid maxBlockSize).
//	 */
func LZ4_decoderRingBufferSize(tls *libc.TLS, maxBlockSize int32) (r int32) {
	if maxBlockSize < 0 {
		return 0
	}
	if maxBlockSize > int32(LZ4_MAX_INPUT_SIZE) {
		return 0
	}
	if maxBlockSize < int32(16) {
		maxBlockSize = int32(16)
	}
	return libc.Int32FromInt32(65536) + libc.Int32FromInt32(14) + maxBlockSize
}

/*
*_continue() :
    These decoding functions allow decompression of multiple blocks in "streaming" mode.
    Previously decoded blocks must still be available at the memory position where they were decoded.
    If it's not possible, save the relevant part of decoded data into a safe buffer,
    and indicate where it stands using LZ4_setStreamDecode()
*/

func LZ4_decompress_safe_continue(tls *libc.TLS, LZ4_streamDecode uintptr, source uintptr, dest uintptr, compressedSize int32, maxOutputSize int32) (r int32) {
	var lz4sd uintptr
	var result int32
	_, _ = lz4sd, result
	lz4sd = LZ4_streamDecode
	if (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize == uint64(0) {
		/* The first call, no dictionary yet. */
		result = LZ4_decompress_safe(tls, source, dest, compressedSize, maxOutputSize)
		if result <= 0 {
			return result
		}
		(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize = libc.Uint64FromInt32(result)
		(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd = dest + uintptr(result)
	} else {
		if (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd == dest {
			/* They're rolling the current segment. */
			if (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize >= libc.Uint64FromInt32(libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))-libc.Int32FromInt32(1)) {
				result = LZ4_decompress_safe_withPrefix64k(tls, source, dest, compressedSize, maxOutputSize)
			} else {
				if (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize == uint64(0) {
					result = LZ4_decompress_safe_withSmallPrefix(tls, source, dest, compressedSize, maxOutputSize, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize)
				} else {
					result = LZ4_decompress_safe_doubleDict(tls, source, dest, compressedSize, maxOutputSize, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FexternalDict, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize)
				}
			}
			if result <= 0 {
				return result
			}
			**(**size_t)(__ccgo_up(lz4sd + 24)) += libc.Uint64FromInt32(result)
			**(**uintptr)(__ccgo_up(lz4sd + 8)) += uintptr(result)
		} else {
			/* The buffer wraps around, or they're switching to another buffer. */
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize = (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FexternalDict = (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd - uintptr((*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize)
			result = LZ4_decompress_safe_forceExtDict(tls, source, dest, compressedSize, maxOutputSize, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FexternalDict, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize)
			if result <= 0 {
				return result
			}
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize = libc.Uint64FromInt32(result)
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd = dest + uintptr(result)
		}
	}
	return result
}

func LZ4_decompress_fast_continue(tls *libc.TLS, LZ4_streamDecode uintptr, source uintptr, dest uintptr, originalSize int32) (r int32) {
	var lz4sd uintptr
	var result int32
	_, _ = lz4sd, result
	lz4sd = LZ4_streamDecode
	if (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize == uint64(0) {
		result = LZ4_decompress_fast(tls, source, dest, originalSize)
		if result <= 0 {
			return result
		}
		(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize = libc.Uint64FromInt32(originalSize)
		(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd = dest + uintptr(originalSize)
	} else {
		if (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd == dest {
			result = LZ4_decompress_unsafe_generic(tls, source, dest, originalSize, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FexternalDict, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize)
			if result <= 0 {
				return result
			}
			**(**size_t)(__ccgo_up(lz4sd + 24)) += libc.Uint64FromInt32(originalSize)
			**(**uintptr)(__ccgo_up(lz4sd + 8)) += uintptr(originalSize)
		} else {
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize = (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FexternalDict = (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd - uintptr((*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize)
			result = LZ4_decompress_fast_extDict(tls, source, dest, originalSize, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FexternalDict, (*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FextDictSize)
			if result <= 0 {
				return result
			}
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixSize = libc.Uint64FromInt32(originalSize)
			(*LZ4_streamDecode_t_internal)(unsafe.Pointer(lz4sd)).FprefixEnd = dest + uintptr(originalSize)
		}
	}
	return result
}

/*
Advanced decoding functions :
*_usingDict() :
    These decoding functions work the same as "_continue" ones,
    the dictionary must be explicitly provided within parameters
*/

func LZ4_decompress_safe_usingDict(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, maxOutputSize int32, dictStart uintptr, dictSize int32) (r int32) {
	if dictSize == 0 {
		return LZ4_decompress_safe(tls, source, dest, compressedSize, maxOutputSize)
	}
	if dictStart+uintptr(dictSize) == dest {
		if dictSize >= libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))-libc.Int32FromInt32(1) {
			return LZ4_decompress_safe_withPrefix64k(tls, source, dest, compressedSize, maxOutputSize)
		}
		return LZ4_decompress_safe_withSmallPrefix(tls, source, dest, compressedSize, maxOutputSize, libc.Uint64FromInt32(dictSize))
	}
	return LZ4_decompress_safe_forceExtDict(tls, source, dest, compressedSize, maxOutputSize, dictStart, libc.Uint64FromInt32(dictSize))
}

func LZ4_decompress_safe_partial_usingDict(tls *libc.TLS, source uintptr, dest uintptr, compressedSize int32, targetOutputSize int32, dstCapacity int32, dictStart uintptr, dictSize int32) (r int32) {
	if dictSize == 0 {
		return LZ4_decompress_safe_partial(tls, source, dest, compressedSize, targetOutputSize, dstCapacity)
	}
	if dictStart+uintptr(dictSize) == dest {
		if dictSize >= libc.Int32FromInt32(64)*(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10))-libc.Int32FromInt32(1) {
			return LZ4_decompress_safe_partial_withPrefix64k(tls, source, dest, compressedSize, targetOutputSize, dstCapacity)
		}
		return LZ4_decompress_safe_partial_withSmallPrefix(tls, source, dest, compressedSize, targetOutputSize, dstCapacity, libc.Uint64FromInt32(dictSize))
	}
	return LZ4_decompress_safe_partial_forceExtDict(tls, source, dest, compressedSize, targetOutputSize, dstCapacity, dictStart, libc.Uint64FromInt32(dictSize))
}

func LZ4_decompress_fast_usingDict(tls *libc.TLS, source uintptr, dest uintptr, originalSize int32, dictStart uintptr, dictSize int32) (r int32) {
	if dictSize == 0 || dictStart+uintptr(dictSize) == dest {
		return LZ4_decompress_unsafe_generic(tls, source, dest, originalSize, libc.Uint64FromInt32(dictSize), libc.UintptrFromInt32(0), uint64(0))
	}
	return LZ4_decompress_fast_extDict(tls, source, dest, originalSize, dictStart, libc.Uint64FromInt32(dictSize))
}

// C documentation
//
//	/*=*************************************************
//	*  Obsolete Functions
//	***************************************************/
//	/* obsolete compression functions */
func LZ4_compress_limitedOutput(tls *libc.TLS, source uintptr, dest uintptr, inputSize int32, maxOutputSize int32) (r int32) {
	return LZ4_compress_default(tls, source, dest, inputSize, maxOutputSize)
}

func LZ4_compress(tls *libc.TLS, src uintptr, dest uintptr, srcSize int32) (r int32) {
	return LZ4_compress_default(tls, src, dest, srcSize, LZ4_compressBound(tls, srcSize))
}

func LZ4_compress_limitedOutput_withState(tls *libc.TLS, state uintptr, src uintptr, dst uintptr, srcSize int32, dstSize int32) (r int32) {
	return LZ4_compress_fast_extState(tls, state, src, dst, srcSize, dstSize, int32(1))
}

func LZ4_compress_withState(tls *libc.TLS, state uintptr, src uintptr, dst uintptr, srcSize int32) (r int32) {
	return LZ4_compress_fast_extState(tls, state, src, dst, srcSize, LZ4_compressBound(tls, srcSize), int32(1))
}

func LZ4_compress_limitedOutput_continue(tls *libc.TLS, LZ4_stream uintptr, src uintptr, dst uintptr, srcSize int32, dstCapacity int32) (r int32) {
	return LZ4_compress_fast_continue(tls, LZ4_stream, src, dst, srcSize, dstCapacity, int32(1))
}

func LZ4_compress_continue(tls *libc.TLS, LZ4_stream uintptr, source uintptr, dest uintptr, inputSize int32) (r int32) {
	return LZ4_compress_fast_continue(tls, LZ4_stream, source, dest, inputSize, LZ4_compressBound(tls, inputSize), int32(1))
}

// C documentation
//
//	/*
//	These decompression functions are deprecated and should no longer be used.
//	They are only provided here for compatibility with older user programs.
//	- LZ4_uncompress is totally equivalent to LZ4_decompress_fast
//	- LZ4_uncompress_unknownOutputSize is totally equivalent to LZ4_decompress_safe
//	*/
func LZ4_uncompress(tls *libc.TLS, source uintptr, dest uintptr, outputSize int32) (r int32) {
	return LZ4_decompress_fast(tls, source, dest, outputSize)
}

func LZ4_uncompress_unknownOutputSize(tls *libc.TLS, source uintptr, dest uintptr, isize int32, maxOutputSize int32) (r int32) {
	return LZ4_decompress_safe(tls, source, dest, isize, maxOutputSize)
}

/* Obsolete Streaming functions */

func LZ4_sizeofStreamState(tls *libc.TLS) (r int32) {
	return int32(16416)
}

func LZ4_resetStreamState(tls *libc.TLS, state uintptr, inputBuffer uintptr) (r int32) {
	_ = inputBuffer
	LZ4_resetStream(tls, state)
	return 0
}

func LZ4_create(tls *libc.TLS, inputBuffer uintptr) (r uintptr) {
	_ = inputBuffer
	return LZ4_createStream(tls)
}

func LZ4_slideInputBuffer(tls *libc.TLS, state uintptr) (r uintptr) {
	/* avoid const char * -> char * conversion warning */
	return uintptr(uint64((*(*LZ4_stream_t_internal)(unsafe.Pointer(state))).Fdictionary))
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "1.9.4\x00"
