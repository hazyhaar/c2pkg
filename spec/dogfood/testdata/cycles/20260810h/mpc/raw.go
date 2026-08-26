// Code generated for linux/amd64 by 'ccgo --package-name=df_mpc -o spec/dogfood/cycles/20260810g/mpc/raw.go -I spec/dogfood/cycles/20260810g/mpc spec/dogfood/cycles/20260810g/mpc/src.c', DO NOT EDIT.

//go:build linux && amd64

package df_mpc

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
const E2BIG = 7
const EACCES = 13
const EADDRINUSE = 98
const EADDRNOTAVAIL = 99
const EADV = 68
const EAFNOSUPPORT = 97
const EAGAIN = 11
const EALREADY = 114
const EBADE = 52
const EBADF = 9
const EBADFD = 77
const EBADMSG = 74
const EBADR = 53
const EBADRQC = 56
const EBADSLT = 57
const EBFONT = 59
const EBUSY = 16
const ECANCELED = 125
const ECHILD = 10
const ECHRNG = 44
const ECOMM = 70
const ECONNABORTED = 103
const ECONNREFUSED = 111
const ECONNRESET = 104
const EDEADLK = 35
const EDEADLOCK = "EDEADLK"
const EDESTADDRREQ = 89
const EDOM = 33
const EDOTDOT = 73
const EDQUOT = 122
const EEXIST = 17
const EFAULT = 14
const EFBIG = 27
const EHOSTDOWN = 112
const EHOSTUNREACH = 113
const EHWPOISON = 133
const EIDRM = 43
const EILSEQ = 84
const EINPROGRESS = 115
const EINTR = 4
const EINVAL = 22
const EIO = 5
const EISCONN = 106
const EISDIR = 21
const EISNAM = 120
const EKEYEXPIRED = 127
const EKEYREJECTED = 129
const EKEYREVOKED = 128
const EL2HLT = 51
const EL2NSYNC = 45
const EL3HLT = 46
const EL3RST = 47
const ELIBACC = 79
const ELIBBAD = 80
const ELIBEXEC = 83
const ELIBMAX = 82
const ELIBSCN = 81
const ELNRNG = 48
const ELOOP = 40
const EMEDIUMTYPE = 124
const EMFILE = 24
const EMLINK = 31
const EMSGSIZE = 90
const EMULTIHOP = 72
const ENAMETOOLONG = 36
const ENAVAIL = 119
const ENETDOWN = 100
const ENETRESET = 102
const ENETUNREACH = 101
const ENFILE = 23
const ENOANO = 55
const ENOBUFS = 105
const ENOCSI = 50
const ENODATA = 61
const ENODEV = 19
const ENOENT = 2
const ENOEXEC = 8
const ENOKEY = 126
const ENOLCK = 37
const ENOLINK = 67
const ENOMEDIUM = 123
const ENOMEM = 12
const ENOMSG = 42
const ENONET = 64
const ENOPKG = 65
const ENOPROTOOPT = 92
const ENOSPC = 28
const ENOSR = 63
const ENOSTR = 60
const ENOSYS = 38
const ENOTBLK = 15
const ENOTCONN = 107
const ENOTDIR = 20
const ENOTEMPTY = 39
const ENOTNAM = 118
const ENOTRECOVERABLE = 131
const ENOTSOCK = 88
const ENOTSUP = "EOPNOTSUPP"
const ENOTTY = 25
const ENOTUNIQ = 76
const ENXIO = 6
const EOPNOTSUPP = 95
const EOVERFLOW = 75
const EOWNERDEAD = 130
const EPERM = 1
const EPFNOSUPPORT = 96
const EPIPE = 32
const EPROTO = 71
const EPROTONOSUPPORT = 93
const EPROTOTYPE = 91
const ERANGE = 34
const EREMCHG = 78
const EREMOTE = 66
const EREMOTEIO = 121
const ERESTART = 85
const ERFKILL = 132
const EROFS = 30
const ESHUTDOWN = 108
const ESOCKTNOSUPPORT = 94
const ESPIPE = 29
const ESRCH = 3
const ESRMNT = 69
const ESTALE = 116
const ESTRPIPE = 86
const ETIME = 62
const ETIMEDOUT = 110
const ETOOMANYREFS = 109
const ETXTBSY = 26
const EUCLEAN = 117
const EUNATCH = 49
const EUSERS = 87
const EWOULDBLOCK = "EAGAIN"
const EXDEV = 18
const EXFULL = 54
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const FILENAME_MAX = 4096
const FOPEN_MAX = 1000
const FP_ILOGB0 = "FP_ILOGBNAN"
const FP_INFINITE = 1
const FP_NAN = 0
const FP_NORMAL = 4
const FP_SUBNORMAL = 3
const FP_ZERO = 2
const HUGE = 3.40282346638528859812e+38
const HUGE_VALF = "INFINITY"
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MATH_ERREXCEPT = 2
const MATH_ERRNO = 1
const MPC_MAX_RECURSION_DEPTH = 1000
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
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const TMP_MAX = 10000
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
const math_errhandling = 2
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type wchar_t = int32

type size_t = uint64

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

type locale_t = uintptr

type float_t = float32

type double_t = float64

type mpc_state_t = struct {
	Fpos  int64
	Frow  int64
	Fcol  int64
	Fterm int32
}

type mpc_err_t = struct {
	Fstate        mpc_state_t
	Fexpected_num int32
	Ffilename     uintptr
	Ffailure      uintptr
	Fexpected     uintptr
	Freceived     int8
}

type mpc_val_t = struct{}

type mpc_result_t = struct {
	Foutput [0]uintptr
	Ferror1 uintptr
}

type mpc_parser_t = struct {
	Fname     uintptr
	Fdata     mpc_pdata_t
	Ftype1    int8
	Fretained int8
}

type mpc_dtor_t = uintptr

type mpc_ctor_t = uintptr

type mpc_apply_t = uintptr

type mpc_apply_to_t = uintptr

type mpc_fold_t = uintptr

type mpc_check_t = uintptr

type mpc_check_with_t = uintptr

const MPC_RE_DEFAULT = 0
const MPC_RE_M = 1
const MPC_RE_S = 2
const MPC_RE_MULTILINE = 1
const MPC_RE_DOTALL = 2

type mpc_ast_t = struct {
	Ftag          uintptr
	Fcontents     uintptr
	Fstate        mpc_state_t
	Fchildren_num int32
	Fchildren     uintptr
}

type mpc_ast_trav_order_t = int32

const mpc_ast_trav_order_pre = 0
const mpc_ast_trav_order_post = 1

type mpc_ast_trav_t = struct {
	Fcurr_node  uintptr
	Fparent     uintptr
	Fcurr_child int32
	Forder      mpc_ast_trav_order_t
}

const MPCA_LANG_DEFAULT = 0
const MPCA_LANG_PREDICTIVE = 1
const MPCA_LANG_WHITESPACE_SENSITIVE = 2

/*
** State Type
 */
func mpc_state_invalid(tls *libc.TLS) (r mpc_state_t) {
	var s mpc_state_t
	_ = s
	s.Fpos = int64(-int32(1))
	s.Frow = int64(-int32(1))
	s.Fcol = int64(-int32(1))
	s.Fterm = 0
	return s
}

func mpc_state_new(tls *libc.TLS) (r mpc_state_t) {
	var s mpc_state_t
	_ = s
	s.Fpos = 0
	s.Frow = 0
	s.Fcol = 0
	s.Fterm = 0
	return s
}

const MPC_INPUT_STRING = 0
const MPC_INPUT_FILE = 1
const MPC_INPUT_PIPE = 2
const MPC_INPUT_MARKS_MIN = 32
const MPC_INPUT_MEM_NUM = 512

type mpc_mem_t = struct {
	Fmem [64]int8
}

type mpc_input_t = struct {
	Ftype1       int32
	Ffilename    uintptr
	Fstate       mpc_state_t
	Fstring1     uintptr
	Fbuffer      uintptr
	Ffile        uintptr
	Fsuppress    int32
	Fbacktrack   int32
	Fmarks_slots int32
	Fmarks_num   int32
	Fmarks       uintptr
	Flasts       uintptr
	Flast        int8
	Fmem_index   size_t
	Fmem_full    [512]int8
	Fmem         [512]mpc_mem_t
}

func mpc_input_new_string(tls *libc.TLS, filename uintptr, string1 uintptr) (r uintptr) {
	var i uintptr
	_ = i
	i = libc.Xmalloc(tls, uint64(33400))
	(*mpc_input_t)(unsafe.Pointer(i)).Ffilename = libc.Xmalloc(tls, libc.Xstrlen(tls, filename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename, filename)
	(*mpc_input_t)(unsafe.Pointer(i)).Ftype1 = int32(MPC_INPUT_STRING)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstate = mpc_state_new(tls)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstring1 = libc.Xmalloc(tls, libc.Xstrlen(tls, string1)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fstring1, string1)
	(*mpc_input_t)(unsafe.Pointer(i)).Fbuffer = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Ffile = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Fsuppress = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack = int32(1)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots = int32(MPC_INPUT_MARKS_MIN)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks = libc.Xmalloc(tls, uint64(32)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flasts = libc.Xmalloc(tls, uint64(1)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flast = int8('\000')
	(*mpc_input_t)(unsafe.Pointer(i)).Fmem_index = uint64(0)
	libc.Xmemset(tls, i+120, 0, libc.Uint64FromInt64(1)*uint64(MPC_INPUT_MEM_NUM))
	return i
}

func mpc_input_new_nstring(tls *libc.TLS, filename uintptr, string1 uintptr, length size_t) (r uintptr) {
	var i uintptr
	_ = i
	i = libc.Xmalloc(tls, uint64(33400))
	(*mpc_input_t)(unsafe.Pointer(i)).Ffilename = libc.Xmalloc(tls, libc.Xstrlen(tls, filename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename, filename)
	(*mpc_input_t)(unsafe.Pointer(i)).Ftype1 = int32(MPC_INPUT_STRING)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstate = mpc_state_new(tls)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstring1 = libc.Xmalloc(tls, length+uint64(1))
	libc.Xstrncpy(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fstring1, string1, length)
	**(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fstring1 + uintptr(length))) = int8('\000')
	(*mpc_input_t)(unsafe.Pointer(i)).Fbuffer = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Ffile = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Fsuppress = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack = int32(1)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots = int32(MPC_INPUT_MARKS_MIN)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks = libc.Xmalloc(tls, uint64(32)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flasts = libc.Xmalloc(tls, uint64(1)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flast = int8('\000')
	(*mpc_input_t)(unsafe.Pointer(i)).Fmem_index = uint64(0)
	libc.Xmemset(tls, i+120, 0, libc.Uint64FromInt64(1)*uint64(MPC_INPUT_MEM_NUM))
	return i
}

func mpc_input_new_pipe(tls *libc.TLS, filename uintptr, pipe uintptr) (r uintptr) {
	var i uintptr
	_ = i
	i = libc.Xmalloc(tls, uint64(33400))
	(*mpc_input_t)(unsafe.Pointer(i)).Ffilename = libc.Xmalloc(tls, libc.Xstrlen(tls, filename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename, filename)
	(*mpc_input_t)(unsafe.Pointer(i)).Ftype1 = int32(MPC_INPUT_PIPE)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstate = mpc_state_new(tls)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstring1 = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Fbuffer = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Ffile = pipe
	(*mpc_input_t)(unsafe.Pointer(i)).Fsuppress = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack = int32(1)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots = int32(MPC_INPUT_MARKS_MIN)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks = libc.Xmalloc(tls, uint64(32)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flasts = libc.Xmalloc(tls, uint64(1)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flast = int8('\000')
	(*mpc_input_t)(unsafe.Pointer(i)).Fmem_index = uint64(0)
	libc.Xmemset(tls, i+120, 0, libc.Uint64FromInt64(1)*uint64(MPC_INPUT_MEM_NUM))
	return i
}

func mpc_input_new_file(tls *libc.TLS, filename uintptr, file uintptr) (r uintptr) {
	var i uintptr
	_ = i
	i = libc.Xmalloc(tls, uint64(33400))
	(*mpc_input_t)(unsafe.Pointer(i)).Ffilename = libc.Xmalloc(tls, libc.Xstrlen(tls, filename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename, filename)
	(*mpc_input_t)(unsafe.Pointer(i)).Ftype1 = int32(MPC_INPUT_FILE)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstate = mpc_state_new(tls)
	(*mpc_input_t)(unsafe.Pointer(i)).Fstring1 = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Fbuffer = libc.UintptrFromInt32(0)
	(*mpc_input_t)(unsafe.Pointer(i)).Ffile = file
	(*mpc_input_t)(unsafe.Pointer(i)).Fsuppress = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack = int32(1)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num = 0
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots = int32(MPC_INPUT_MARKS_MIN)
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks = libc.Xmalloc(tls, uint64(32)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flasts = libc.Xmalloc(tls, uint64(1)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	(*mpc_input_t)(unsafe.Pointer(i)).Flast = int8('\000')
	(*mpc_input_t)(unsafe.Pointer(i)).Fmem_index = uint64(0)
	libc.Xmemset(tls, i+120, 0, libc.Uint64FromInt64(1)*uint64(MPC_INPUT_MEM_NUM))
	return i
}

func mpc_input_delete(tls *libc.TLS, i uintptr) {
	libc.Xfree(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename)
	if (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 == int32(MPC_INPUT_STRING) {
		libc.Xfree(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fstring1)
	}
	if (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 == int32(MPC_INPUT_PIPE) {
		libc.Xfree(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer)
	}
	libc.Xfree(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fmarks)
	libc.Xfree(tls, (*mpc_input_t)(unsafe.Pointer(i)).Flasts)
	libc.Xfree(tls, i)
}

func mpc_mem_ptr(tls *libc.TLS, i uintptr, p uintptr) (r int32) {
	return libc.BoolInt32(p >= i+632 && p < i+632+uintptr(uint64(MPC_INPUT_MEM_NUM)*libc.Uint64FromInt64(64)))
}

func mpc_malloc(tls *libc.TLS, i uintptr, n size_t) (r uintptr) {
	var j size_t
	var p uintptr
	_, _ = j, p
	if n > uint64(64) {
		return libc.Xmalloc(tls, n)
	}
	j = (*mpc_input_t)(unsafe.Pointer(i)).Fmem_index
	for cond := true; cond; cond = j != (*mpc_input_t)(unsafe.Pointer(i)).Fmem_index {
		if !(**(**int8)(__ccgo_up(i + 120 + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fmem_index))) != 0) {
			p = i + 632 + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fmem_index)*64
			**(**int8)(__ccgo_up(i + 120 + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fmem_index))) = int8(1)
			(*mpc_input_t)(unsafe.Pointer(i)).Fmem_index = ((*mpc_input_t)(unsafe.Pointer(i)).Fmem_index + uint64(1)) % uint64(MPC_INPUT_MEM_NUM)
			return p
		}
		(*mpc_input_t)(unsafe.Pointer(i)).Fmem_index = ((*mpc_input_t)(unsafe.Pointer(i)).Fmem_index + uint64(1)) % uint64(MPC_INPUT_MEM_NUM)
	}
	return libc.Xmalloc(tls, n)
}

func mpc_calloc(tls *libc.TLS, i uintptr, n size_t, m size_t) (r uintptr) {
	var x uintptr
	_ = x
	x = mpc_malloc(tls, i, n*m)
	libc.Xmemset(tls, x, 0, n*m)
	return x
}

func mpc_free(tls *libc.TLS, i uintptr, p uintptr) {
	var j size_t
	_ = j
	if !(mpc_mem_ptr(tls, i, p) != 0) {
		libc.Xfree(tls, p)
		return
	}
	j = libc.Uint64FromInt64(int64(p)-int64(i+632)) / uint64(64)
	**(**int8)(__ccgo_up(i + 120 + uintptr(j))) = 0
}

func mpc_realloc(tls *libc.TLS, i uintptr, p uintptr, n size_t) (r uintptr) {
	var q uintptr
	_ = q
	q = libc.UintptrFromInt32(0)
	if !(mpc_mem_ptr(tls, i, p) != 0) {
		return libc.Xrealloc(tls, p, n)
	}
	if n > uint64(64) {
		q = libc.Xmalloc(tls, n)
		libc.Xmemcpy(tls, q, p, uint64(64))
		mpc_free(tls, i, p)
		return q
	}
	return p
}

func mpc_export(tls *libc.TLS, i uintptr, p uintptr) (r uintptr) {
	var q uintptr
	_ = q
	q = libc.UintptrFromInt32(0)
	if !(mpc_mem_ptr(tls, i, p) != 0) {
		return p
	}
	q = libc.Xmalloc(tls, uint64(64))
	libc.Xmemcpy(tls, q, p, uint64(64))
	mpc_free(tls, i, p)
	return q
}

func mpc_input_backtrack_disable(tls *libc.TLS, i uintptr) {
	(*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack = (*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack - 1
}

func mpc_input_backtrack_enable(tls *libc.TLS, i uintptr) {
	(*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack = (*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack + 1
}

func mpc_input_suppress_disable(tls *libc.TLS, i uintptr) {
	(*mpc_input_t)(unsafe.Pointer(i)).Fsuppress = (*mpc_input_t)(unsafe.Pointer(i)).Fsuppress - 1
}

func mpc_input_suppress_enable(tls *libc.TLS, i uintptr) {
	(*mpc_input_t)(unsafe.Pointer(i)).Fsuppress = (*mpc_input_t)(unsafe.Pointer(i)).Fsuppress + 1
}

func mpc_input_mark(tls *libc.TLS, i uintptr) {
	if (*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack < int32(1) {
		return
	}
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num = (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num + 1
	if (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num > (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots {
		(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots = (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num + (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num/int32(2)
		(*mpc_input_t)(unsafe.Pointer(i)).Fmarks = libc.Xrealloc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fmarks, uint64(32)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
		(*mpc_input_t)(unsafe.Pointer(i)).Flasts = libc.Xrealloc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Flasts, uint64(1)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	}
	**(**mpc_state_t)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fmarks + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num-int32(1))*32)) = (*mpc_input_t)(unsafe.Pointer(i)).Fstate
	**(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Flasts + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num-int32(1)))) = (*mpc_input_t)(unsafe.Pointer(i)).Flast
	if (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 == int32(MPC_INPUT_PIPE) && (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num == int32(1) {
		(*mpc_input_t)(unsafe.Pointer(i)).Fbuffer = libc.Xcalloc(tls, uint64(1), uint64(1))
	}
}

func mpc_input_unmark(tls *libc.TLS, i uintptr) {
	var j, v1 int32
	_, _ = j, v1
	if (*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack < int32(1) {
		return
	}
	(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num = (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num - 1
	if (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots > (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num+(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num/int32(2) && (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots > int32(MPC_INPUT_MARKS_MIN) {
		if (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num > int32(MPC_INPUT_MARKS_MIN) {
			v1 = (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num
		} else {
			v1 = int32(MPC_INPUT_MARKS_MIN)
		}
		(*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots = v1
		(*mpc_input_t)(unsafe.Pointer(i)).Fmarks = libc.Xrealloc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fmarks, uint64(32)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
		(*mpc_input_t)(unsafe.Pointer(i)).Flasts = libc.Xrealloc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Flasts, uint64(1)*libc.Uint64FromInt32((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_slots))
	}
	if (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 == int32(MPC_INPUT_PIPE) && (*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num == 0 {
		j = libc.Int32FromUint64(libc.Xstrlen(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer) - uint64(1))
		for {
			if !(j >= 0) {
				break
			}
			libc.Xungetc(tls, int32(**(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fbuffer + uintptr(j)))), (*mpc_input_t)(unsafe.Pointer(i)).Ffile)
			goto _2
		_2:
			;
			j = j - 1
		}
		libc.Xfree(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer)
		(*mpc_input_t)(unsafe.Pointer(i)).Fbuffer = libc.UintptrFromInt32(0)
	}
}

func mpc_input_rewind(tls *libc.TLS, i uintptr) {
	if (*mpc_input_t)(unsafe.Pointer(i)).Fbacktrack < int32(1) {
		return
	}
	(*mpc_input_t)(unsafe.Pointer(i)).Fstate = **(**mpc_state_t)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fmarks + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num-int32(1))*32))
	(*mpc_input_t)(unsafe.Pointer(i)).Flast = **(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Flasts + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fmarks_num-int32(1))))
	if (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 == int32(MPC_INPUT_FILE) {
		libc.Xfseek(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile, (*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fpos, 0)
	}
	mpc_input_unmark(tls, i)
}

func mpc_input_buffer_in_range(tls *libc.TLS, i uintptr) (r int32) {
	return libc.BoolInt32((*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fpos < libc.Int64FromUint64(libc.Xstrlen(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer)+libc.Uint64FromInt64((**(**mpc_state_t)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fmarks))).Fpos)))
}

func mpc_input_buffer_get(tls *libc.TLS, i uintptr) (r int8) {
	return **(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fbuffer + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fpos-(**(**mpc_state_t)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fmarks))).Fpos)))
}

func mpc_input_getc(tls *libc.TLS, i uintptr) (r int8) {
	var c int8
	_ = c
	c = int8('\000')
	switch (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 {
	case int32(MPC_INPUT_STRING):
		return **(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fstring1 + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fpos)))
	case int32(MPC_INPUT_FILE):
		c = int8(libc.Xfgetc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile))
		return c
	case int32(MPC_INPUT_PIPE):
		if !((*mpc_input_t)(unsafe.Pointer(i)).Fbuffer != 0) {
			c = int8(libc.Xgetc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile))
			return c
		}
		if (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer != 0 && mpc_input_buffer_in_range(tls, i) != 0 {
			c = mpc_input_buffer_get(tls, i)
			return c
		} else {
			c = int8(libc.Xgetc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile))
			return c
		}
		fallthrough
	default:
		return c
	}
	return r
}

func mpc_input_peekc(tls *libc.TLS, i uintptr) (r int8) {
	var c int8
	_ = c
	c = int8('\000')
	switch (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 {
	case int32(MPC_INPUT_STRING):
		return **(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fstring1 + uintptr((*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fpos)))
	case int32(MPC_INPUT_FILE):
		c = int8(libc.Xfgetc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile))
		if libc.Xfeof(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile) != 0 {
			return int8('\000')
		}
		libc.Xfseek(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile, int64(-int32(1)), int32(1))
		return c
	case int32(MPC_INPUT_PIPE):
		if !((*mpc_input_t)(unsafe.Pointer(i)).Fbuffer != 0) {
			c = int8(libc.Xgetc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile))
			if libc.Xfeof(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile) != 0 {
				return int8('\000')
			}
			libc.Xungetc(tls, int32(c), (*mpc_input_t)(unsafe.Pointer(i)).Ffile)
			return c
		}
		if (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer != 0 && mpc_input_buffer_in_range(tls, i) != 0 {
			return mpc_input_buffer_get(tls, i)
		} else {
			c = int8(libc.Xgetc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile))
			if libc.Xfeof(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile) != 0 {
				return int8('\000')
			}
			libc.Xungetc(tls, int32(c), (*mpc_input_t)(unsafe.Pointer(i)).Ffile)
			return c
		}
		fallthrough
	default:
		return c
	}
	return r
}

func mpc_input_terminated(tls *libc.TLS, i uintptr) (r int32) {
	return libc.BoolInt32(int32(mpc_input_peekc(tls, i)) == int32('\000'))
}

func mpc_input_failure(tls *libc.TLS, i uintptr, c int8) (r int32) {
	switch (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 {
	case int32(MPC_INPUT_STRING):
	case int32(MPC_INPUT_FILE):
		libc.Xfseek(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffile, int64(-int32(1)), int32(1))
	case int32(MPC_INPUT_PIPE):
		if !((*mpc_input_t)(unsafe.Pointer(i)).Fbuffer != 0) {
			libc.Xungetc(tls, int32(c), (*mpc_input_t)(unsafe.Pointer(i)).Ffile)
			break
		}
		if (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer != 0 && mpc_input_buffer_in_range(tls, i) != 0 {
			break
		} else {
			libc.Xungetc(tls, int32(c), (*mpc_input_t)(unsafe.Pointer(i)).Ffile)
		}
		fallthrough
	default:
		break
	}
	return 0
}

func mpc_input_success(tls *libc.TLS, i uintptr, c int8, o uintptr) (r int32) {
	if (*mpc_input_t)(unsafe.Pointer(i)).Ftype1 == int32(MPC_INPUT_PIPE) && (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer != 0 && !(mpc_input_buffer_in_range(tls, i) != 0) {
		(*mpc_input_t)(unsafe.Pointer(i)).Fbuffer = libc.Xrealloc(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer, libc.Xstrlen(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer)+uint64(2))
		**(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fbuffer + uintptr(libc.Xstrlen(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer)+uint64(1)))) = int8('\000')
		**(**int8)(__ccgo_up((*mpc_input_t)(unsafe.Pointer(i)).Fbuffer + uintptr(libc.Xstrlen(tls, (*mpc_input_t)(unsafe.Pointer(i)).Fbuffer)+uint64(0)))) = c
	}
	(*mpc_input_t)(unsafe.Pointer(i)).Flast = c
	(*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fpos = (*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fpos + 1
	(*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fcol = (*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fcol + 1
	if int32(c) == int32('\n') {
		(*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fcol = 0
		(*mpc_input_t)(unsafe.Pointer(i)).Fstate.Frow = (*mpc_input_t)(unsafe.Pointer(i)).Fstate.Frow + 1
	}
	if o != 0 {
		**(**uintptr)(__ccgo_up(o)) = mpc_malloc(tls, i, uint64(2))
		**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(o)))) = c
		**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(o)) + 1)) = int8('\000')
	}
	return int32(1)
}

func mpc_input_any(tls *libc.TLS, i uintptr, o uintptr) (r int32) {
	var x int8
	_ = x
	if mpc_input_terminated(tls, i) != 0 {
		return 0
	}
	x = mpc_input_getc(tls, i)
	return mpc_input_success(tls, i, x, o)
}

func mpc_input_char(tls *libc.TLS, i uintptr, c int8, o uintptr) (r int32) {
	var x int8
	var v1 int32
	_, _ = x, v1
	if mpc_input_terminated(tls, i) != 0 {
		return 0
	}
	x = mpc_input_getc(tls, i)
	if int32(x) == int32(c) {
		v1 = mpc_input_success(tls, i, x, o)
	} else {
		v1 = mpc_input_failure(tls, i, x)
	}
	return v1
}

func mpc_input_range(tls *libc.TLS, i uintptr, c int8, d int8, o uintptr) (r int32) {
	var x int8
	var v1 int32
	_, _ = x, v1
	if mpc_input_terminated(tls, i) != 0 {
		return 0
	}
	x = mpc_input_getc(tls, i)
	if int32(x) >= int32(c) && int32(x) <= int32(d) {
		v1 = mpc_input_success(tls, i, x, o)
	} else {
		v1 = mpc_input_failure(tls, i, x)
	}
	return v1
}

func mpc_input_oneof(tls *libc.TLS, i uintptr, c uintptr, o uintptr) (r int32) {
	var x int8
	var v1 int32
	_, _ = x, v1
	if mpc_input_terminated(tls, i) != 0 {
		return 0
	}
	x = mpc_input_getc(tls, i)
	if libc.Xstrchr(tls, c, int32(x)) != uintptr(0) {
		v1 = mpc_input_success(tls, i, x, o)
	} else {
		v1 = mpc_input_failure(tls, i, x)
	}
	return v1
}

func mpc_input_noneof(tls *libc.TLS, i uintptr, c uintptr, o uintptr) (r int32) {
	var x int8
	var v1 int32
	_, _ = x, v1
	if mpc_input_terminated(tls, i) != 0 {
		return 0
	}
	x = mpc_input_getc(tls, i)
	if libc.Xstrchr(tls, c, int32(x)) == uintptr(0) {
		v1 = mpc_input_success(tls, i, x, o)
	} else {
		v1 = mpc_input_failure(tls, i, x)
	}
	return v1
}

func mpc_input_satisfy(tls *libc.TLS, i uintptr, __ccgo_fp_cond uintptr, o uintptr) (r int32) {
	var x int8
	var v1 int32
	_, _ = x, v1
	if mpc_input_terminated(tls, i) != 0 {
		return 0
	}
	x = mpc_input_getc(tls, i)
	if (*(*func(*libc.TLS, int8) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_cond})))(tls, x) != 0 {
		v1 = mpc_input_success(tls, i, x, o)
	} else {
		v1 = mpc_input_failure(tls, i, x)
	}
	return v1
}

func mpc_input_string(tls *libc.TLS, i uintptr, c uintptr, o uintptr) (r int32) {
	var x uintptr
	_ = x
	x = c
	mpc_input_mark(tls, i)
	for **(**int8)(__ccgo_up(x)) != 0 {
		if !(mpc_input_char(tls, i, **(**int8)(__ccgo_up(x)), libc.UintptrFromInt32(0)) != 0) {
			mpc_input_rewind(tls, i)
			return 0
		}
		x = x + 1
	}
	mpc_input_unmark(tls, i)
	**(**uintptr)(__ccgo_up(o)) = mpc_malloc(tls, i, libc.Xstrlen(tls, c)+uint64(1))
	libc.Xstrcpy(tls, **(**uintptr)(__ccgo_up(o)), c)
	return int32(1)
}

func mpc_input_anchor(tls *libc.TLS, i uintptr, __ccgo_fp_f uintptr, o uintptr) (r int32) {
	**(**uintptr)(__ccgo_up(o)) = libc.UintptrFromInt32(0)
	return (*(*func(*libc.TLS, int8, int8) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_f})))(tls, (*mpc_input_t)(unsafe.Pointer(i)).Flast, mpc_input_peekc(tls, i))
}

func mpc_input_soi(tls *libc.TLS, i uintptr, o uintptr) (r int32) {
	**(**uintptr)(__ccgo_up(o)) = libc.UintptrFromInt32(0)
	return libc.BoolInt32(int32((*mpc_input_t)(unsafe.Pointer(i)).Flast) == int32('\000'))
}

func mpc_input_eoi(tls *libc.TLS, i uintptr, o uintptr) (r int32) {
	**(**uintptr)(__ccgo_up(o)) = libc.UintptrFromInt32(0)
	if (*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fterm != 0 {
		return 0
	} else {
		if mpc_input_terminated(tls, i) != 0 {
			(*mpc_input_t)(unsafe.Pointer(i)).Fstate.Fterm = int32(1)
			return int32(1)
		} else {
			return 0
		}
	}
	return r
}

func mpc_input_state_copy(tls *libc.TLS, i uintptr) (r1 uintptr) {
	var r uintptr
	_ = r
	r = mpc_malloc(tls, i, uint64(32))
	libc.Xmemcpy(tls, r, i+16, uint64(32))
	return r
}

/*
** Error Type
 */
func mpc_err_delete(tls *libc.TLS, x uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num) {
			break
		}
		libc.Xfree(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfree(tls, (*mpc_err_t)(unsafe.Pointer(x)).Fexpected)
	libc.Xfree(tls, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename)
	libc.Xfree(tls, (*mpc_err_t)(unsafe.Pointer(x)).Ffailure)
	libc.Xfree(tls, x)
}

func mpc_err_print(tls *libc.TLS, x uintptr) {
	mpc_err_print_to(tls, x, libc.Xstdout)
}

func mpc_err_print_to(tls *libc.TLS, x uintptr, f uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var str uintptr
	_ = str
	str = mpc_err_string(tls, x)
	libc.Xfprintf(tls, f, __ccgo_ts, libc.VaList(bp+8, str))
	libc.Xfree(tls, str)
}

func mpc_err_string_cat(tls *libc.TLS, buffer uintptr, pos uintptr, max uintptr, fmt uintptr, va1 uintptr) {
	var left int32
	var va va_list
	_, _ = left, va
	/* TODO: Error Checking on Length */
	left = **(**int32)(__ccgo_up(max)) - **(**int32)(__ccgo_up(pos))
	va = va1
	if left < 0 {
		left = 0
	}
	**(**int32)(__ccgo_up(pos)) += libc.Xvsprintf(tls, buffer+uintptr(**(**int32)(__ccgo_up(pos))), fmt, va)
	_ = va
}

func mpc_err_char_unescape(tls *libc.TLS, c int8, char_unescape_buffer uintptr) (r uintptr) {
	**(**int8)(__ccgo_up(char_unescape_buffer)) = int8('\'')
	**(**int8)(__ccgo_up(char_unescape_buffer + 1)) = int8(' ')
	**(**int8)(__ccgo_up(char_unescape_buffer + 2)) = int8('\'')
	**(**int8)(__ccgo_up(char_unescape_buffer + 3)) = int8('\000')
	switch int32(c) {
	case int32('\a'):
		return __ccgo_ts + 3
	case int32('\b'):
		return __ccgo_ts + 8
	case int32('\f'):
		return __ccgo_ts + 18
	case int32('\r'):
		return __ccgo_ts + 27
	case int32('\v'):
		return __ccgo_ts + 43
	case int32('\000'):
		return __ccgo_ts + 56
	case int32('\n'):
		return __ccgo_ts + 69
	case int32('\t'):
		return __ccgo_ts + 77
	case int32(' '):
		return __ccgo_ts + 81
	default:
		**(**int8)(__ccgo_up(char_unescape_buffer + 1)) = c
		return char_unescape_buffer
	}
	return r
}

func mpc_err_string(tls *libc.TLS, x uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var buffer uintptr
	var i int32
	var _ /* char_unescape_buffer at bp+8 */ [4]int8
	var _ /* max at bp+4 */ int32
	var _ /* pos at bp+0 */ int32
	_, _ = buffer, i
	**(**int32)(__ccgo_up(bp)) = 0
	**(**int32)(__ccgo_up(bp + 4)) = int32(1023)
	buffer = libc.Xcalloc(tls, uint64(1), uint64(1024))
	if (*mpc_err_t)(unsafe.Pointer(x)).Ffailure != 0 {
		mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts+87, libc.VaList(bp+24, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename, (*mpc_err_t)(unsafe.Pointer(x)).Ffailure))
		return buffer
	}
	mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts+102, libc.VaList(bp+24, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename, (*mpc_err_t)(unsafe.Pointer(x)).Fstate.Frow+int64(1), (*mpc_err_t)(unsafe.Pointer(x)).Fstate.Fcol+int64(1)))
	if (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num == 0 {
		mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts+131, 0)
	}
	if (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num == int32(1) {
		mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts, libc.VaList(bp+24, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected))))
	}
	if (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num >= int32(2) {
		i = 0
		for {
			if !(i < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(2)) {
				break
			}
			mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts+155, libc.VaList(bp+24, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(i)*8))))
			goto _1
		_1:
			;
			i = i + 1
		}
		mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts+160, libc.VaList(bp+24, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(2))*8)), **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(1))*8))))
	}
	mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts+169, 0)
	mpc_err_string_cat(tls, buffer, bp, bp+4, mpc_err_char_unescape(tls, (*mpc_err_t)(unsafe.Pointer(x)).Freceived, bp+8), 0)
	mpc_err_string_cat(tls, buffer, bp, bp+4, __ccgo_ts+174, 0)
	return libc.Xrealloc(tls, buffer, libc.Xstrlen(tls, buffer)+uint64(1))
}

func mpc_err_new(tls *libc.TLS, i uintptr, expected uintptr) (r uintptr) {
	var x uintptr
	_ = x
	if (*mpc_input_t)(unsafe.Pointer(i)).Fsuppress != 0 {
		return libc.UintptrFromInt32(0)
	}
	x = mpc_malloc(tls, i, uint64(72))
	(*mpc_err_t)(unsafe.Pointer(x)).Ffilename = mpc_malloc(tls, i, libc.Xstrlen(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename)
	(*mpc_err_t)(unsafe.Pointer(x)).Fstate = (*mpc_input_t)(unsafe.Pointer(i)).Fstate
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num = int32(1)
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected = mpc_malloc(tls, i, uint64(8))
	**(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)) = mpc_malloc(tls, i, libc.Xstrlen(tls, expected)+uint64(1))
	libc.Xstrcpy(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)), expected)
	(*mpc_err_t)(unsafe.Pointer(x)).Ffailure = libc.UintptrFromInt32(0)
	(*mpc_err_t)(unsafe.Pointer(x)).Freceived = mpc_input_peekc(tls, i)
	return x
}

func mpc_err_fail(tls *libc.TLS, i uintptr, failure uintptr) (r uintptr) {
	var x uintptr
	_ = x
	if (*mpc_input_t)(unsafe.Pointer(i)).Fsuppress != 0 {
		return libc.UintptrFromInt32(0)
	}
	x = mpc_malloc(tls, i, uint64(72))
	(*mpc_err_t)(unsafe.Pointer(x)).Ffilename = mpc_malloc(tls, i, libc.Xstrlen(tls, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename, (*mpc_input_t)(unsafe.Pointer(i)).Ffilename)
	(*mpc_err_t)(unsafe.Pointer(x)).Fstate = (*mpc_input_t)(unsafe.Pointer(i)).Fstate
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num = 0
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected = libc.UintptrFromInt32(0)
	(*mpc_err_t)(unsafe.Pointer(x)).Ffailure = mpc_malloc(tls, i, libc.Xstrlen(tls, failure)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_err_t)(unsafe.Pointer(x)).Ffailure, failure)
	(*mpc_err_t)(unsafe.Pointer(x)).Freceived = int8(' ')
	return x
}

func mpc_err_file(tls *libc.TLS, filename uintptr, failure uintptr) (r uintptr) {
	var x uintptr
	_ = x
	x = libc.Xmalloc(tls, uint64(72))
	(*mpc_err_t)(unsafe.Pointer(x)).Ffilename = libc.Xmalloc(tls, libc.Xstrlen(tls, filename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename, filename)
	(*mpc_err_t)(unsafe.Pointer(x)).Fstate = mpc_state_new(tls)
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num = 0
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected = libc.UintptrFromInt32(0)
	(*mpc_err_t)(unsafe.Pointer(x)).Ffailure = libc.Xmalloc(tls, libc.Xstrlen(tls, failure)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_err_t)(unsafe.Pointer(x)).Ffailure, failure)
	(*mpc_err_t)(unsafe.Pointer(x)).Freceived = int8(' ')
	return x
}

func mpc_err_delete_internal(tls *libc.TLS, i uintptr, x uintptr) {
	var j int32
	_ = j
	if x == libc.UintptrFromInt32(0) {
		return
	}
	j = 0
	for {
		if !(j < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num) {
			break
		}
		mpc_free(tls, i, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(j)*8)))
		goto _1
	_1:
		;
		j = j + 1
	}
	mpc_free(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Fexpected)
	mpc_free(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename)
	mpc_free(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Ffailure)
	mpc_free(tls, i, x)
}

func mpc_err_export(tls *libc.TLS, i uintptr, x uintptr) (r uintptr) {
	var j int32
	_ = j
	j = 0
	for {
		if !(j < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num) {
			break
		}
		**(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(j)*8)) = mpc_export(tls, i, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(j)*8)))
		goto _1
	_1:
		;
		j = j + 1
	}
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected = mpc_export(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Fexpected)
	(*mpc_err_t)(unsafe.Pointer(x)).Ffilename = mpc_export(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Ffilename)
	(*mpc_err_t)(unsafe.Pointer(x)).Ffailure = mpc_export(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Ffailure)
	return mpc_export(tls, i, x)
}

func mpc_err_contains_expected(tls *libc.TLS, i uintptr, x uintptr, expected uintptr) (r int32) {
	var j int32
	_ = j
	_ = i
	j = 0
	for {
		if !(j < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num) {
			break
		}
		if libc.Xstrcmp(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(j)*8)), expected) == 0 {
			return int32(1)
		}
		goto _1
	_1:
		;
		j = j + 1
	}
	return 0
}

func mpc_err_add_expected(tls *libc.TLS, i uintptr, x uintptr, expected uintptr) {
	_ = i
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num = (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num + 1
	(*mpc_err_t)(unsafe.Pointer(x)).Fexpected = mpc_realloc(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Fexpected, uint64(8)*libc.Uint64FromInt32((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num))
	**(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(1))*8)) = mpc_malloc(tls, i, libc.Xstrlen(tls, expected)+uint64(1))
	libc.Xstrcpy(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(1))*8)), expected)
}

func mpc_err_or(tls *libc.TLS, i uintptr, x uintptr, n int32) (r uintptr) {
	var e uintptr
	var fst, j, k int32
	_, _, _, _ = e, fst, j, k
	fst = -int32(1)
	j = 0
	for {
		if !(j < n) {
			break
		}
		if **(**uintptr)(__ccgo_up(x + uintptr(j)*8)) != libc.UintptrFromInt32(0) {
			fst = j
		}
		goto _1
	_1:
		;
		j = j + 1
	}
	if fst == -int32(1) {
		return libc.UintptrFromInt32(0)
	}
	e = mpc_malloc(tls, i, uint64(72))
	(*mpc_err_t)(unsafe.Pointer(e)).Fstate = mpc_state_invalid(tls)
	(*mpc_err_t)(unsafe.Pointer(e)).Fexpected_num = 0
	(*mpc_err_t)(unsafe.Pointer(e)).Fexpected = libc.UintptrFromInt32(0)
	(*mpc_err_t)(unsafe.Pointer(e)).Ffailure = libc.UintptrFromInt32(0)
	(*mpc_err_t)(unsafe.Pointer(e)).Ffilename = mpc_malloc(tls, i, libc.Xstrlen(tls, (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(fst)*8)))).Ffilename)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_err_t)(unsafe.Pointer(e)).Ffilename, (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(fst)*8)))).Ffilename)
	j = 0
	for {
		if !(j < n) {
			break
		}
		if **(**uintptr)(__ccgo_up(x + uintptr(j)*8)) == libc.UintptrFromInt32(0) {
			goto _2
		}
		if (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Fstate.Fpos > (*mpc_err_t)(unsafe.Pointer(e)).Fstate.Fpos {
			(*mpc_err_t)(unsafe.Pointer(e)).Fstate = (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Fstate
		}
		goto _2
	_2:
		;
		j = j + 1
	}
	j = 0
	for {
		if !(j < n) {
			break
		}
		if **(**uintptr)(__ccgo_up(x + uintptr(j)*8)) == libc.UintptrFromInt32(0) {
			goto _3
		}
		if (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Fstate.Fpos < (*mpc_err_t)(unsafe.Pointer(e)).Fstate.Fpos {
			goto _3
		}
		if (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Ffailure != 0 {
			(*mpc_err_t)(unsafe.Pointer(e)).Ffailure = mpc_malloc(tls, i, libc.Xstrlen(tls, (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Ffailure)+uint64(1))
			libc.Xstrcpy(tls, (*mpc_err_t)(unsafe.Pointer(e)).Ffailure, (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Ffailure)
			break
		}
		(*mpc_err_t)(unsafe.Pointer(e)).Freceived = (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Freceived
		k = 0
		for {
			if !(k < (*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Fexpected_num) {
				break
			}
			if !(mpc_err_contains_expected(tls, i, e, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Fexpected + uintptr(k)*8))) != 0) {
				mpc_err_add_expected(tls, i, e, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(x + uintptr(j)*8)))).Fexpected + uintptr(k)*8)))
			}
			goto _4
		_4:
			;
			k = k + 1
		}
		goto _3
	_3:
		;
		j = j + 1
	}
	j = 0
	for {
		if !(j < n) {
			break
		}
		if **(**uintptr)(__ccgo_up(x + uintptr(j)*8)) == libc.UintptrFromInt32(0) {
			goto _5
		}
		mpc_err_delete_internal(tls, i, **(**uintptr)(__ccgo_up(x + uintptr(j)*8)))
		goto _5
	_5:
		;
		j = j + 1
	}
	return e
}

func mpc_err_repeat(tls *libc.TLS, i uintptr, x uintptr, prefix uintptr) (r uintptr) {
	var expect uintptr
	var j int32
	var l size_t
	_, _, _ = expect, j, l
	j = 0
	l = uint64(0)
	expect = libc.UintptrFromInt32(0)
	if x == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num == 0 {
		expect = mpc_calloc(tls, i, uint64(1), uint64(1))
		(*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num = int32(1)
		(*mpc_err_t)(unsafe.Pointer(x)).Fexpected = mpc_realloc(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Fexpected, uint64(8)*libc.Uint64FromInt32((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num))
		**(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)) = expect
		return x
	} else {
		if (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num == int32(1) {
			expect = mpc_malloc(tls, i, libc.Xstrlen(tls, prefix)+libc.Xstrlen(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)))+uint64(1))
			libc.Xstrcpy(tls, expect, prefix)
			libc.Xstrcat(tls, expect, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)))
			mpc_free(tls, i, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)))
			**(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)) = expect
			return x
		} else {
			if (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num > int32(1) {
				l = l + libc.Xstrlen(tls, prefix)
				j = 0
				for {
					if !(j < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(2)) {
						break
					}
					l = l + (libc.Xstrlen(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(j)*8))) + libc.Xstrlen(tls, __ccgo_ts+176))
					goto _1
				_1:
					;
					j = j + 1
				}
				l = l + libc.Xstrlen(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(2))*8)))
				l = l + libc.Xstrlen(tls, __ccgo_ts+179)
				l = l + libc.Xstrlen(tls, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(1))*8)))
				expect = mpc_malloc(tls, i, l+uint64(1))
				libc.Xstrcpy(tls, expect, prefix)
				j = 0
				for {
					if !(j < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(2)) {
						break
					}
					libc.Xstrcat(tls, expect, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(j)*8)))
					libc.Xstrcat(tls, expect, __ccgo_ts+176)
					goto _2
				_2:
					;
					j = j + 1
				}
				libc.Xstrcat(tls, expect, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(2))*8)))
				libc.Xstrcat(tls, expect, __ccgo_ts+179)
				libc.Xstrcat(tls, expect, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num-int32(1))*8)))
				j = 0
				for {
					if !(j < (*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num) {
						break
					}
					mpc_free(tls, i, **(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected + uintptr(j)*8)))
					goto _3
				_3:
					;
					j = j + 1
				}
				(*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num = int32(1)
				(*mpc_err_t)(unsafe.Pointer(x)).Fexpected = mpc_realloc(tls, i, (*mpc_err_t)(unsafe.Pointer(x)).Fexpected, uint64(8)*libc.Uint64FromInt32((*mpc_err_t)(unsafe.Pointer(x)).Fexpected_num))
				**(**uintptr)(__ccgo_up((*mpc_err_t)(unsafe.Pointer(x)).Fexpected)) = expect
				return x
			}
		}
	}
	return libc.UintptrFromInt32(0)
}

func mpc_err_many1(tls *libc.TLS, i uintptr, x uintptr) (r uintptr) {
	return mpc_err_repeat(tls, i, x, __ccgo_ts+184)
}

func mpc_err_count(tls *libc.TLS, i uintptr, x uintptr, n int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var digits int32
	var prefix, y uintptr
	_, _, _ = digits, prefix, y
	digits = n/int32(10) + int32(1)
	prefix = mpc_malloc(tls, i, libc.Uint64FromInt32(digits)+libc.Xstrlen(tls, __ccgo_ts+200)+uint64(1))
	if !(prefix != 0) {
		return libc.UintptrFromInt32(0)
	}
	libc.Xsprintf(tls, prefix, __ccgo_ts+205, libc.VaList(bp+8, n))
	y = mpc_err_repeat(tls, i, x, prefix)
	mpc_free(tls, i, prefix)
	return y
}

func mpc_err_merge(tls *libc.TLS, i uintptr, x uintptr, y uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* errs at bp+0 */ [2]uintptr
	(**(**[2]uintptr)(__ccgo_up(bp)))[0] = x
	(**(**[2]uintptr)(__ccgo_up(bp)))[int32(1)] = y
	return mpc_err_or(tls, i, bp, int32(2))
}

const MPC_TYPE_UNDEFINED = 0
const MPC_TYPE_PASS = 1
const MPC_TYPE_FAIL = 2
const MPC_TYPE_LIFT = 3
const MPC_TYPE_LIFT_VAL = 4
const MPC_TYPE_EXPECT = 5
const MPC_TYPE_ANCHOR = 6
const MPC_TYPE_STATE = 7
const MPC_TYPE_ANY = 8
const MPC_TYPE_SINGLE = 9
const MPC_TYPE_ONEOF = 10
const MPC_TYPE_NONEOF = 11
const MPC_TYPE_RANGE = 12
const MPC_TYPE_SATISFY = 13
const MPC_TYPE_STRING = 14
const MPC_TYPE_APPLY = 15
const MPC_TYPE_APPLY_TO = 16
const MPC_TYPE_PREDICT = 17
const MPC_TYPE_NOT = 18
const MPC_TYPE_MAYBE = 19
const MPC_TYPE_MANY = 20
const MPC_TYPE_MANY1 = 21
const MPC_TYPE_COUNT = 22
const MPC_TYPE_OR = 23
const MPC_TYPE_AND = 24
const MPC_TYPE_CHECK = 25
const MPC_TYPE_CHECK_WITH = 26
const MPC_TYPE_SOI = 27
const MPC_TYPE_EOI = 28
const MPC_TYPE_SEPBY1 = 29

type mpc_pdata_fail_t = struct {
	Fm uintptr
}

type mpc_pdata_lift_t = struct {
	Flf mpc_ctor_t
	Fx  uintptr
}

type mpc_pdata_expect_t = struct {
	Fx uintptr
	Fm uintptr
}

type mpc_pdata_anchor_t = struct {
	Ff uintptr
}

type mpc_pdata_single_t = struct {
	Fx int8
}

type mpc_pdata_range_t = struct {
	Fx int8
	Fy int8
}

type mpc_pdata_satisfy_t = struct {
	Ff uintptr
}

type mpc_pdata_string_t = struct {
	Fx uintptr
}

type mpc_pdata_apply_t = struct {
	Fx uintptr
	Ff mpc_apply_t
}

type mpc_pdata_apply_to_t = struct {
	Fx uintptr
	Ff mpc_apply_to_t
	Fd uintptr
}

type mpc_pdata_check_t = struct {
	Fx  uintptr
	Fdx mpc_dtor_t
	Ff  mpc_check_t
	Fe  uintptr
}

type mpc_pdata_check_with_t = struct {
	Fx  uintptr
	Fdx mpc_dtor_t
	Ff  mpc_check_with_t
	Fd  uintptr
	Fe  uintptr
}

type mpc_pdata_predict_t = struct {
	Fx uintptr
}

type mpc_pdata_not_t = struct {
	Fx  uintptr
	Fdx mpc_dtor_t
	Flf mpc_ctor_t
}

type mpc_pdata_repeat_t = struct {
	Fn  int32
	Ff  mpc_fold_t
	Fx  uintptr
	Fdx mpc_dtor_t
}

type mpc_pdata_or_t = struct {
	Fn  int32
	Fxs uintptr
}

type mpc_pdata_and_t = struct {
	Fn   int32
	Ff   mpc_fold_t
	Fxs  uintptr
	Fdxs uintptr
}

type mpc_pdata_sepby1 = struct {
	Fn   int32
	Ff   mpc_fold_t
	Fx   uintptr
	Fsep uintptr
}

type mpc_pdata_t = struct {
	Flift         [0]mpc_pdata_lift_t
	Fexpect       [0]mpc_pdata_expect_t
	Fanchor       [0]mpc_pdata_anchor_t
	Fsingle       [0]mpc_pdata_single_t
	Frange1       [0]mpc_pdata_range_t
	Fsatisfy      [0]mpc_pdata_satisfy_t
	Fstring1      [0]mpc_pdata_string_t
	Fapply        [0]mpc_pdata_apply_t
	Fapply_to     [0]mpc_pdata_apply_to_t
	Fcheck        [0]mpc_pdata_check_t
	Fcheck_with   [0]mpc_pdata_check_with_t
	Fpredict      [0]mpc_pdata_predict_t
	Fnot          [0]mpc_pdata_not_t
	Frepeat       [0]mpc_pdata_repeat_t
	Fand          [0]mpc_pdata_and_t
	For           [0]mpc_pdata_or_t
	Fsepby1       [0]mpc_pdata_sepby1
	Ffail         mpc_pdata_fail_t
	F__ccgo_pad18 [32]byte
}

func mpcf_input_nth_free(tls *libc.TLS, i uintptr, n int32, xs uintptr, x int32) (r uintptr) {
	var j int32
	_ = j
	j = 0
	for {
		if !(j < n) {
			break
		}
		if j != x {
			mpc_free(tls, i, **(**uintptr)(__ccgo_up(xs + uintptr(j)*8)))
		}
		goto _1
	_1:
		;
		j = j + 1
	}
	return **(**uintptr)(__ccgo_up(xs + uintptr(x)*8))
}

func mpcf_input_fst_free(tls *libc.TLS, i uintptr, n int32, xs uintptr) (r uintptr) {
	return mpcf_input_nth_free(tls, i, n, xs, 0)
}

func mpcf_input_snd_free(tls *libc.TLS, i uintptr, n int32, xs uintptr) (r uintptr) {
	return mpcf_input_nth_free(tls, i, n, xs, int32(1))
}

func mpcf_input_trd_free(tls *libc.TLS, i uintptr, n int32, xs uintptr) (r uintptr) {
	return mpcf_input_nth_free(tls, i, n, xs, int32(2))
}

func mpcf_input_strfold(tls *libc.TLS, i uintptr, n int32, xs uintptr) (r uintptr) {
	var j int32
	var l size_t
	_, _ = j, l
	l = uint64(0)
	if n == 0 {
		return mpc_calloc(tls, i, uint64(1), uint64(1))
	}
	j = 0
	for {
		if !(j < n) {
			break
		}
		l = l + libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(xs + uintptr(j)*8)))
		goto _1
	_1:
		;
		j = j + 1
	}
	**(**uintptr)(__ccgo_up(xs)) = mpc_realloc(tls, i, **(**uintptr)(__ccgo_up(xs)), l+uint64(1))
	j = int32(1)
	for {
		if !(j < n) {
			break
		}
		libc.Xstrcat(tls, **(**uintptr)(__ccgo_up(xs)), **(**uintptr)(__ccgo_up(xs + uintptr(j)*8)))
		mpc_free(tls, i, **(**uintptr)(__ccgo_up(xs + uintptr(j)*8)))
		goto _2
	_2:
		;
		j = j + 1
	}
	return **(**uintptr)(__ccgo_up(xs))
}

func mpcf_input_state_ast(tls *libc.TLS, i uintptr, n int32, xs uintptr) (r uintptr) {
	var a, s uintptr
	_, _ = a, s
	s = **(**uintptr)(__ccgo_up(xs))
	a = **(**uintptr)(__ccgo_up(xs + 1*8))
	a = mpc_ast_state(tls, a, **(**mpc_state_t)(__ccgo_up(s)))
	mpc_free(tls, i, s)
	_ = n
	return a
}

func mpc_parse_fold(tls *libc.TLS, i uintptr, __ccgo_fp_f mpc_fold_t, n int32, xs uintptr) (r uintptr) {
	var j int32
	_ = j
	if __ccgo_fp_f == __ccgo_fp(mpcf_null) {
		return mpcf_null(tls, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_fst) {
		return mpcf_fst(tls, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_snd) {
		return mpcf_snd(tls, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_trd) {
		return mpcf_trd(tls, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_fst_free) {
		return mpcf_input_fst_free(tls, i, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_snd_free) {
		return mpcf_input_snd_free(tls, i, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_trd_free) {
		return mpcf_input_trd_free(tls, i, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_strfold) {
		return mpcf_input_strfold(tls, i, n, xs)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_state_ast) {
		return mpcf_input_state_ast(tls, i, n, xs)
	}
	j = 0
	for {
		if !(j < n) {
			break
		}
		**(**uintptr)(__ccgo_up(xs + uintptr(j)*8)) = mpc_export(tls, i, **(**uintptr)(__ccgo_up(xs + uintptr(j)*8)))
		goto _1
	_1:
		;
		j = j + 1
	}
	return (*(*func(*libc.TLS, int32, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_f})))(tls, j, xs)
}

func mpcf_input_free(tls *libc.TLS, i uintptr, x uintptr) (r uintptr) {
	mpc_free(tls, i, x)
	return libc.UintptrFromInt32(0)
}

func mpcf_input_str_ast(tls *libc.TLS, i uintptr, c uintptr) (r uintptr) {
	var a uintptr
	_ = a
	a = mpc_ast_new(tls, __ccgo_ts+212, c)
	mpc_free(tls, i, c)
	return a
}

func mpc_parse_apply(tls *libc.TLS, i uintptr, __ccgo_fp_f mpc_apply_t, x uintptr) (r uintptr) {
	if __ccgo_fp_f == __ccgo_fp(mpcf_free) {
		return mpcf_input_free(tls, i, x)
	}
	if __ccgo_fp_f == __ccgo_fp(mpcf_str_ast) {
		return mpcf_input_str_ast(tls, i, x)
	}
	return (*(*func(*libc.TLS, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_f})))(tls, mpc_export(tls, i, x))
}

func mpc_parse_apply_to(tls *libc.TLS, i uintptr, __ccgo_fp_f mpc_apply_to_t, x uintptr, d uintptr) (r uintptr) {
	return (*(*func(*libc.TLS, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_f})))(tls, mpc_export(tls, i, x), d)
}

func mpc_parse_dtor(tls *libc.TLS, i uintptr, __ccgo_fp_d mpc_dtor_t, x uintptr) {
	if __ccgo_fp_d == __ccgo_fp(libc.Xfree) {
		mpc_free(tls, i, x)
		return
	}
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_d})))(tls, mpc_export(tls, i, x))
}

const MPC_PARSE_STACK_MIN = 4

func mpc_grow_results(tls *libc.TLS, i uintptr, j int32, results_stk uintptr, results uintptr) (r uintptr) {
	var results_slots, results_slots1 int32
	var tmp_results uintptr
	_, _, _ = results_slots, results_slots1, tmp_results
	tmp_results = results
	if j == int32(MPC_PARSE_STACK_MIN) {
		results_slots = j + j/int32(2)
		tmp_results = mpc_malloc(tls, i, uint64(8)*libc.Uint64FromInt32(results_slots))
		libc.Xmemcpy(tls, tmp_results, results_stk, libc.Uint64FromInt64(8)*uint64(MPC_PARSE_STACK_MIN))
	} else {
		if j >= int32(MPC_PARSE_STACK_MIN) {
			results_slots1 = j + j/int32(2)
			tmp_results = mpc_realloc(tls, i, tmp_results, uint64(8)*libc.Uint64FromInt32(results_slots1))
		}
	}
	return tmp_results
}

func mpc_parse_run(tls *libc.TLS, i uintptr, p uintptr, r uintptr, e uintptr, depth int32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var j, k int32
	var results, v1 uintptr
	var _ /* results_stk at bp+0 */ [4]mpc_result_t
	_, _, _, _ = j, k, results, v1
	j = 0
	k = 0
	if depth == int32(MPC_MAX_RECURSION_DEPTH) {
		(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_fail(tls, i, __ccgo_ts+213)
		return 0
	}
	switch int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) {
	/* Basic Parsers */
	case int32(MPC_TYPE_ANY):
		if mpc_input_any(tls, i, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_SINGLE):
		if mpc_input_char(tls, i, (*(*mpc_pdata_single_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_RANGE):
		if mpc_input_range(tls, i, (*(*mpc_pdata_range_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, (*(*mpc_pdata_range_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fy, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_ONEOF):
		if mpc_input_oneof(tls, i, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_NONEOF):
		if mpc_input_noneof(tls, i, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_SATISFY):
		if mpc_input_satisfy(tls, i, (*(*mpc_pdata_satisfy_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_STRING):
		if mpc_input_string(tls, i, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_ANCHOR):
		if mpc_input_anchor(tls, i, (*(*mpc_pdata_anchor_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_SOI):
		if mpc_input_soi(tls, i, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_EOI):
		if mpc_input_eoi(tls, i, r) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
			return 0
		}
		/* Other parsers */
		fallthrough
	case int32(MPC_TYPE_UNDEFINED):
		(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_fail(tls, i, __ccgo_ts+247)
		return 0
	case int32(MPC_TYPE_PASS):
		*(*uintptr)(unsafe.Pointer(r)) = libc.UintptrFromInt32(0)
		return int32(1)
	case int32(MPC_TYPE_FAIL):
		(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_fail(tls, i, (*mpc_parser_t)(unsafe.Pointer(p)).Fdata.Ffail.Fm)
		return 0
	case int32(MPC_TYPE_LIFT):
		*(*uintptr)(unsafe.Pointer(r)) = (*(*func(*libc.TLS) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*(*mpc_pdata_lift_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Flf})))(tls)
		return int32(1)
	case int32(MPC_TYPE_LIFT_VAL):
		*(*uintptr)(unsafe.Pointer(r)) = (*(*mpc_pdata_lift_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx
		return int32(1)
	case int32(MPC_TYPE_STATE):
		*(*uintptr)(unsafe.Pointer(r)) = mpc_input_state_copy(tls, i)
		return int32(1)
		/* Application Parsers */
		fallthrough
	case int32(MPC_TYPE_APPLY):
		if mpc_parse_run(tls, i, (*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r, e, depth+int32(1)) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = mpc_parse_apply(tls, i, (*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, *(*uintptr)(unsafe.Pointer(r)))
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = *(*uintptr)(unsafe.Pointer(r))
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_APPLY_TO):
		if mpc_parse_run(tls, i, (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r, e, depth+int32(1)) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = mpc_parse_apply_to(tls, i, (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, *(*uintptr)(unsafe.Pointer(r)), (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fd)
			return int32(1)
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = (*mpc_result_t)(unsafe.Pointer(r)).Ferror1
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_CHECK):
		if mpc_parse_run(tls, i, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r, e, depth+int32(1)) != 0 {
			if (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff})))(tls, r) != 0 {
				*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
				return int32(1)
			} else {
				mpc_parse_dtor(tls, i, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdx, *(*uintptr)(unsafe.Pointer(r)))
				(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_fail(tls, i, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fe)
				return 0
			}
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = (*mpc_result_t)(unsafe.Pointer(r)).Ferror1
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_CHECK_WITH):
		if mpc_parse_run(tls, i, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fx, r, e, depth+int32(1)) != 0 {
			if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Ff})))(tls, r, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fd) != 0 {
				*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
				return int32(1)
			} else {
				mpc_parse_dtor(tls, i, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdx, *(*uintptr)(unsafe.Pointer(r)))
				(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_fail(tls, i, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fe)
				return 0
			}
		} else {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = (*mpc_result_t)(unsafe.Pointer(r)).Ferror1
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_EXPECT):
		mpc_input_suppress_enable(tls, i)
		if mpc_parse_run(tls, i, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r, e, depth+int32(1)) != 0 {
			mpc_input_suppress_disable(tls, i)
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			mpc_input_suppress_disable(tls, i)
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_new(tls, i, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm)
			return 0
		}
		fallthrough
	case int32(MPC_TYPE_PREDICT):
		mpc_input_backtrack_disable(tls, i)
		if mpc_parse_run(tls, i, (*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r, e, depth+int32(1)) != 0 {
			mpc_input_backtrack_enable(tls, i)
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			mpc_input_backtrack_enable(tls, i)
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = (*mpc_result_t)(unsafe.Pointer(r)).Ferror1
			return 0
		}
		/* Optional Parsers */
		/* TODO: Update Not Error Message */
		fallthrough
	case int32(MPC_TYPE_NOT):
		mpc_input_mark(tls, i)
		mpc_input_suppress_enable(tls, i)
		if mpc_parse_run(tls, i, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r, e, depth+int32(1)) != 0 {
			mpc_input_rewind(tls, i)
			mpc_input_suppress_disable(tls, i)
			mpc_parse_dtor(tls, i, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdx, *(*uintptr)(unsafe.Pointer(r)))
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_new(tls, i, __ccgo_ts+265)
			return 0
		} else {
			mpc_input_unmark(tls, i)
			mpc_input_suppress_disable(tls, i)
			*(*uintptr)(unsafe.Pointer(r)) = (*(*func(*libc.TLS) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Flf})))(tls)
			return int32(1)
		}
		fallthrough
	case int32(MPC_TYPE_MAYBE):
		if mpc_parse_run(tls, i, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, r, e, depth+int32(1)) != 0 {
			*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(r))
			return int32(1)
		} else {
			**(**uintptr)(__ccgo_up(e)) = mpc_err_merge(tls, i, **(**uintptr)(__ccgo_up(e)), (*mpc_result_t)(unsafe.Pointer(r)).Ferror1)
			*(*uintptr)(unsafe.Pointer(r)) = (*(*func(*libc.TLS) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Flf})))(tls)
			return int32(1)
		}
		/* Repeat Parsers */
		fallthrough
	case int32(MPC_TYPE_MANY):
		results = bp
		for mpc_parse_run(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, results+uintptr(j)*8, e, depth+int32(1)) != 0 {
			j = j + 1
			results = mpc_grow_results(tls, i, j, bp, results)
		}
		**(**uintptr)(__ccgo_up(e)) = mpc_err_merge(tls, i, **(**uintptr)(__ccgo_up(e)), *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8)))
		*(*uintptr)(unsafe.Pointer(r)) = mpc_parse_fold(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, j, results)
		if j >= int32(MPC_PARSE_STACK_MIN) {
			mpc_free(tls, i, results)
		}
		return int32(1)
	case int32(MPC_TYPE_MANY1):
		results = bp
		for mpc_parse_run(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, results+uintptr(j)*8, e, depth+int32(1)) != 0 {
			j = j + 1
			results = mpc_grow_results(tls, i, j, bp, results)
		}
		if j == 0 {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_many1(tls, i, *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8)))
			if j >= int32(MPC_PARSE_STACK_MIN) {
				mpc_free(tls, i, results)
			}
			return 0
		} else {
			**(**uintptr)(__ccgo_up(e)) = mpc_err_merge(tls, i, **(**uintptr)(__ccgo_up(e)), *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8)))
			*(*uintptr)(unsafe.Pointer(r)) = mpc_parse_fold(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, j, results)
			if j >= int32(MPC_PARSE_STACK_MIN) {
				mpc_free(tls, i, results)
			}
			return int32(1)
		}
		fallthrough
	case int32(MPC_TYPE_SEPBY1):
		results = bp
		if mpc_parse_run(tls, i, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, results+uintptr(j)*8, e, depth+int32(1)) != 0 {
			j = j + 1
			results = mpc_grow_results(tls, i, j, bp, results)
			for mpc_parse_run(tls, i, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fsep, results+uintptr(j)*8, e, depth+int32(1)) != 0 && mpc_parse_run(tls, i, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, results+uintptr(j)*8, e, depth+int32(1)) != 0 {
				j = j + 1
				results = mpc_grow_results(tls, i, j, bp, results)
			}
		}
		if j == 0 {
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_many1(tls, i, *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8)))
			if j >= int32(MPC_PARSE_STACK_MIN) {
				mpc_free(tls, i, results)
			}
			return 0
		} else {
			**(**uintptr)(__ccgo_up(e)) = mpc_err_merge(tls, i, **(**uintptr)(__ccgo_up(e)), *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8)))
			*(*uintptr)(unsafe.Pointer(r)) = mpc_parse_fold(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, j, results)
			if j >= int32(MPC_PARSE_STACK_MIN) {
				mpc_free(tls, i, results)
			}
			return int32(1)
		}
		fallthrough
	case int32(MPC_TYPE_COUNT):
		if (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
			v1 = mpc_malloc(tls, i, uint64(8)*libc.Uint64FromInt32((*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn))
		} else {
			v1 = bp
		}
		results = v1
		for mpc_parse_run(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, results+uintptr(j)*8, e, depth+int32(1)) != 0 {
			j = j + 1
			if j == (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn {
				break
			}
		}
		if j == (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn {
			*(*uintptr)(unsafe.Pointer(r)) = mpc_parse_fold(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, j, results)
			if (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
				mpc_free(tls, i, results)
			}
			return int32(1)
		} else {
			k = 0
			for {
				if !(k < j) {
					break
				}
				mpc_parse_dtor(tls, i, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdx, *(*uintptr)(unsafe.Pointer(results + uintptr(k)*8)))
				goto _2
			_2:
				;
				k = k + 1
			}
			(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_count(tls, i, *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8)), (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn)
			if (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
				mpc_free(tls, i, results)
			}
			return 0
		}
		/* Combinatory Parsers */
		fallthrough
	case int32(MPC_TYPE_OR):
		if (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn == 0 {
			*(*uintptr)(unsafe.Pointer(r)) = libc.UintptrFromInt32(0)
			return int32(1)
		}
		if (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
			v1 = mpc_malloc(tls, i, uint64(8)*libc.Uint64FromInt32((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn))
		} else {
			v1 = bp
		}
		results = v1
		j = 0
		for {
			if !(j < (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
				break
			}
			if mpc_parse_run(tls, i, **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(j)*8)), results+uintptr(j)*8, e, depth+int32(1)) != 0 {
				*(*uintptr)(unsafe.Pointer(r)) = *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8))
				if (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
					mpc_free(tls, i, results)
				}
				return int32(1)
			} else {
				**(**uintptr)(__ccgo_up(e)) = mpc_err_merge(tls, i, **(**uintptr)(__ccgo_up(e)), *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8)))
			}
			goto _4
		_4:
			;
			j = j + 1
		}
		(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = libc.UintptrFromInt32(0)
		if (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
			mpc_free(tls, i, results)
		}
		return 0
	case int32(MPC_TYPE_AND):
		if (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn == 0 {
			*(*uintptr)(unsafe.Pointer(r)) = libc.UintptrFromInt32(0)
			return int32(1)
		}
		if (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
			v1 = mpc_malloc(tls, i, uint64(8)*libc.Uint64FromInt32((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn))
		} else {
			v1 = bp
		}
		results = v1
		mpc_input_mark(tls, i)
		j = 0
		for {
			if !(j < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
				break
			}
			if !(mpc_parse_run(tls, i, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(j)*8)), results+uintptr(j)*8, e, depth+int32(1)) != 0) {
				mpc_input_rewind(tls, i)
				k = 0
				for {
					if !(k < j) {
						break
					}
					mpc_parse_dtor(tls, i, **(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(k)*8)), *(*uintptr)(unsafe.Pointer(results + uintptr(k)*8)))
					goto _7
				_7:
					;
					k = k + 1
				}
				(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = *(*uintptr)(unsafe.Pointer(results + uintptr(j)*8))
				if (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
					mpc_free(tls, i, results)
				}
				return 0
			}
			goto _6
		_6:
			;
			j = j + 1
		}
		mpc_input_unmark(tls, i)
		*(*uintptr)(unsafe.Pointer(r)) = mpc_parse_fold(tls, i, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff, j, results)
		if (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn > int32(MPC_PARSE_STACK_MIN) {
			mpc_free(tls, i, results)
		}
		return int32(1)
		/* End */
		fallthrough
	default:
		(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_fail(tls, i, __ccgo_ts+274)
		return 0
	}
	return 0
}

func mpc_parse_input(tls *libc.TLS, i uintptr, p uintptr, r uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var x int32
	var _ /* e at bp+0 */ uintptr
	_ = x
	**(**uintptr)(__ccgo_up(bp)) = mpc_err_fail(tls, i, __ccgo_ts+298)
	(*mpc_err_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(bp)))).Fstate = mpc_state_invalid(tls)
	x = mpc_parse_run(tls, i, p, r, bp, 0)
	if x != 0 {
		mpc_err_delete_internal(tls, i, **(**uintptr)(__ccgo_up(bp)))
		*(*uintptr)(unsafe.Pointer(r)) = mpc_export(tls, i, *(*uintptr)(unsafe.Pointer(r)))
	} else {
		(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_export(tls, i, mpc_err_merge(tls, i, **(**uintptr)(__ccgo_up(bp)), (*mpc_result_t)(unsafe.Pointer(r)).Ferror1))
	}
	return x
}

func mpc_parse(tls *libc.TLS, filename uintptr, string1 uintptr, p uintptr, r uintptr) (r1 int32) {
	var i uintptr
	var x int32
	_, _ = i, x
	i = mpc_input_new_string(tls, filename, string1)
	x = mpc_parse_input(tls, i, p, r)
	mpc_input_delete(tls, i)
	return x
}

func mpc_nparse(tls *libc.TLS, filename uintptr, string1 uintptr, length size_t, p uintptr, r uintptr) (r1 int32) {
	var i uintptr
	var x int32
	_, _ = i, x
	i = mpc_input_new_nstring(tls, filename, string1, length)
	x = mpc_parse_input(tls, i, p, r)
	mpc_input_delete(tls, i)
	return x
}

func mpc_parse_file(tls *libc.TLS, filename uintptr, file uintptr, p uintptr, r uintptr) (r1 int32) {
	var i uintptr
	var x int32
	_, _ = i, x
	i = mpc_input_new_file(tls, filename, file)
	x = mpc_parse_input(tls, i, p, r)
	mpc_input_delete(tls, i)
	return x
}

func mpc_parse_pipe(tls *libc.TLS, filename uintptr, pipe uintptr, p uintptr, r uintptr) (r1 int32) {
	var i uintptr
	var x int32
	_, _ = i, x
	i = mpc_input_new_pipe(tls, filename, pipe)
	x = mpc_parse_input(tls, i, p, r)
	mpc_input_delete(tls, i)
	return x
}

func mpc_parse_contents(tls *libc.TLS, filename uintptr, p uintptr, r uintptr) (r1 int32) {
	var f uintptr
	var res int32
	_, _ = f, res
	f = libc.Xfopen(tls, filename, __ccgo_ts+312)
	if f == libc.UintptrFromInt32(0) {
		*(*uintptr)(unsafe.Pointer(r)) = libc.UintptrFromInt32(0)
		(*mpc_result_t)(unsafe.Pointer(r)).Ferror1 = mpc_err_file(tls, filename, __ccgo_ts+315)
		return 0
	}
	res = mpc_parse_file(tls, filename, f, p, r)
	libc.Xfclose(tls, f)
	return res
}

func mpc_undefine_or(tls *libc.TLS, p uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
			break
		}
		mpc_undefine_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfree(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)
}

func mpc_undefine_and(tls *libc.TLS, p uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
			break
		}
		mpc_undefine_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)
	libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs)
}

func mpc_undefine_unretained(tls *libc.TLS, p uintptr, force int32) {
	if (*mpc_parser_t)(unsafe.Pointer(p)).Fretained != 0 && !(force != 0) {
		return
	}
	switch int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) {
	case int32(MPC_TYPE_FAIL):
		libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fdata.Ffail.Fm)
	case int32(MPC_TYPE_ONEOF):
		fallthrough
	case int32(MPC_TYPE_NONEOF):
		fallthrough
	case int32(MPC_TYPE_STRING):
		libc.Xfree(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx)
	case int32(MPC_TYPE_APPLY):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	case int32(MPC_TYPE_APPLY_TO):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	case int32(MPC_TYPE_PREDICT):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	case int32(MPC_TYPE_MAYBE):
		fallthrough
	case int32(MPC_TYPE_NOT):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	case int32(MPC_TYPE_EXPECT):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xfree(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm)
	case int32(MPC_TYPE_MANY):
		fallthrough
	case int32(MPC_TYPE_MANY1):
		fallthrough
	case int32(MPC_TYPE_COUNT):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	case int32(MPC_TYPE_SEPBY1):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		mpc_undefine_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fsep, 0)
	case int32(MPC_TYPE_OR):
		mpc_undefine_or(tls, p)
	case int32(MPC_TYPE_AND):
		mpc_undefine_and(tls, p)
	case int32(MPC_TYPE_CHECK):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xfree(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fe)
	case int32(MPC_TYPE_CHECK_WITH):
		mpc_undefine_unretained(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fx, 0)
		libc.Xfree(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fe)
	default:
		break
	}
	if !(force != 0) {
		libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname)
		libc.Xfree(tls, p)
	}
}

func mpc_delete(tls *libc.TLS, p uintptr) {
	if (*mpc_parser_t)(unsafe.Pointer(p)).Fretained != 0 {
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) != int32(MPC_TYPE_UNDEFINED) {
			mpc_undefine_unretained(tls, p, 0)
		}
		libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname)
		libc.Xfree(tls, p)
	} else {
		mpc_undefine_unretained(tls, p, 0)
	}
}

func mpc_soft_delete(tls *libc.TLS, x uintptr) {
	mpc_undefine_unretained(tls, x, 0)
}

func mpc_undefined(tls *libc.TLS) (r uintptr) {
	var p uintptr
	_ = p
	p = libc.Xcalloc(tls, uint64(1), uint64(56))
	(*mpc_parser_t)(unsafe.Pointer(p)).Fretained = 0
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_UNDEFINED)
	(*mpc_parser_t)(unsafe.Pointer(p)).Fname = libc.UintptrFromInt32(0)
	return p
}

func mpc_new(tls *libc.TLS, name uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Fretained = int8(1)
	(*mpc_parser_t)(unsafe.Pointer(p)).Fname = libc.Xrealloc(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname, libc.Xstrlen(tls, name)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname, name)
	return p
}

func mpc_copy(tls *libc.TLS, a uintptr) (r uintptr) {
	var i int32
	var p uintptr
	_, _ = i, p
	i = 0
	if (*mpc_parser_t)(unsafe.Pointer(a)).Fretained != 0 {
		return a
	}
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Fretained = (*mpc_parser_t)(unsafe.Pointer(a)).Fretained
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = (*mpc_parser_t)(unsafe.Pointer(a)).Ftype1
	(*mpc_parser_t)(unsafe.Pointer(p)).Fdata = (*mpc_parser_t)(unsafe.Pointer(a)).Fdata
	if (*mpc_parser_t)(unsafe.Pointer(a)).Fname != 0 {
		(*mpc_parser_t)(unsafe.Pointer(p)).Fname = libc.Xmalloc(tls, libc.Xstrlen(tls, (*mpc_parser_t)(unsafe.Pointer(a)).Fname)+uint64(1))
		libc.Xstrcpy(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname, (*mpc_parser_t)(unsafe.Pointer(a)).Fname)
	}
	switch int32((*mpc_parser_t)(unsafe.Pointer(a)).Ftype1) {
	case int32(MPC_TYPE_FAIL):
		(*mpc_parser_t)(unsafe.Pointer(p)).Fdata.Ffail.Fm = libc.Xmalloc(tls, libc.Xstrlen(tls, (*mpc_parser_t)(unsafe.Pointer(a)).Fdata.Ffail.Fm)+uint64(1))
		libc.Xstrcpy(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fdata.Ffail.Fm, (*mpc_parser_t)(unsafe.Pointer(a)).Fdata.Ffail.Fm)
	case int32(MPC_TYPE_ONEOF):
		fallthrough
	case int32(MPC_TYPE_NONEOF):
		fallthrough
	case int32(MPC_TYPE_STRING):
		(*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = libc.Xmalloc(tls, libc.Xstrlen(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)+uint64(1))
		libc.Xstrcpy(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
	case int32(MPC_TYPE_APPLY):
		(*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
	case int32(MPC_TYPE_APPLY_TO):
		(*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
	case int32(MPC_TYPE_PREDICT):
		(*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
	case int32(MPC_TYPE_MAYBE):
		fallthrough
	case int32(MPC_TYPE_NOT):
		(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
	case int32(MPC_TYPE_EXPECT):
		(*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
		(*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm = libc.Xmalloc(tls, libc.Xstrlen(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fm)+uint64(1))
		libc.Xstrcpy(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fm)
	case int32(MPC_TYPE_MANY):
		fallthrough
	case int32(MPC_TYPE_MANY1):
		fallthrough
	case int32(MPC_TYPE_COUNT):
		(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
	case int32(MPC_TYPE_SEPBY1):
		(*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
		(*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fsep = mpc_copy(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fsep)
	case int32(MPC_TYPE_OR):
		(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xmalloc(tls, libc.Uint64FromInt32((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fn)*uint64(8))
		i = 0
		for {
			if !(i < (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fn) {
				break
			}
			**(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)) = mpc_copy(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fxs + uintptr(i)*8)))
			goto _1
		_1:
			;
			i = i + 1
		}
	case int32(MPC_TYPE_AND):
		(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xmalloc(tls, libc.Uint64FromInt32((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fn)*uint64(8))
		i = 0
		for {
			if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fn) {
				break
			}
			**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)) = mpc_copy(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fxs + uintptr(i)*8)))
			goto _2
		_2:
			;
			i = i + 1
		}
		if (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fn > 0 {
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs = libc.Xmalloc(tls, libc.Uint64FromInt32((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fn-libc.Int32FromInt32(1))*uint64(8))
			i = 0
			for {
				if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fn-int32(1)) {
					break
				}
				**(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(i)*8)) = **(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fdxs + uintptr(i)*8))
				goto _3
			_3:
				;
				i = i + 1
			}
		}
	case int32(MPC_TYPE_CHECK):
		(*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = mpc_copy(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fx)
		(*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fe = libc.Xmalloc(tls, libc.Xstrlen(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fe)+uint64(1))
		libc.Xstrcpy(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fe, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(a)).Fdata))).Fe)
	case int32(MPC_TYPE_CHECK_WITH):
		(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fx = mpc_copy(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(a + 8))).Fx)
		(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fe = libc.Xmalloc(tls, libc.Xstrlen(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(a + 8))).Fe)+uint64(1))
		libc.Xstrcpy(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fe, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(a + 8))).Fe)
	default:
		break
	}
	return p
}

func mpc_undefine(tls *libc.TLS, p uintptr) (r uintptr) {
	mpc_undefine_unretained(tls, p, int32(1))
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_UNDEFINED)
	return p
}

func mpc_define(tls *libc.TLS, p uintptr, a uintptr) (r uintptr) {
	var a2 uintptr
	_ = a2
	if (*mpc_parser_t)(unsafe.Pointer(p)).Fretained != 0 {
		(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = (*mpc_parser_t)(unsafe.Pointer(a)).Ftype1
		(*mpc_parser_t)(unsafe.Pointer(p)).Fdata = (*mpc_parser_t)(unsafe.Pointer(a)).Fdata
	} else {
		a2 = mpc_failf(tls, __ccgo_ts+336, 0)
		(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = (*mpc_parser_t)(unsafe.Pointer(a2)).Ftype1
		(*mpc_parser_t)(unsafe.Pointer(p)).Fdata = (*mpc_parser_t)(unsafe.Pointer(a2)).Fdata
		libc.Xfree(tls, a2)
	}
	libc.Xfree(tls, a)
	return p
}

func mpc_cleanup(tls *libc.TLS, n int32, va1 uintptr) {
	var i int32
	var list uintptr
	var va va_list
	_, _, _ = i, list, va
	list = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n))
	va = va1
	i = 0
	for {
		if !(i < n) {
			break
		}
		**(**uintptr)(__ccgo_up(list + uintptr(i)*8)) = libc.VaUintptr(&va)
		goto _1
	_1:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < n) {
			break
		}
		mpc_undefine(tls, **(**uintptr)(__ccgo_up(list + uintptr(i)*8)))
		goto _2
	_2:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < n) {
			break
		}
		mpc_delete(tls, **(**uintptr)(__ccgo_up(list + uintptr(i)*8)))
		goto _3
	_3:
		;
		i = i + 1
	}
	_ = va
	libc.Xfree(tls, list)
}

func mpc_pass(tls *libc.TLS) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_PASS)
	return p
}

func mpc_fail(tls *libc.TLS, m uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_FAIL)
	(*mpc_parser_t)(unsafe.Pointer(p)).Fdata.Ffail.Fm = libc.Xmalloc(tls, libc.Xstrlen(tls, m)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fdata.Ffail.Fm, m)
	return p
}

/*
** As `snprintf` is not ANSI standard this
** function `mpc_failf` should be considered
** unsafe.
**
** You have a few options if this is going to be
** trouble.
**
** - Ensure the format string does not exceed
**   the buffer length using precision specifiers
**   such as `%.512s`.
**
** - Patch this function in your code base to
**   use `snprintf` or whatever variant your
**   system supports.
**
** - Avoid it altogether.
**
 */
func mpc_failf(tls *libc.TLS, fmt uintptr, va1 uintptr) (r uintptr) {
	var buffer, p uintptr
	var va va_list
	_, _, _ = buffer, p, va
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_FAIL)
	va = va1
	buffer = libc.Xmalloc(tls, uint64(2048))
	if !(buffer != 0) {
		return libc.UintptrFromInt32(0)
	}
	libc.Xvsprintf(tls, buffer, fmt, va)
	_ = va
	buffer = libc.Xrealloc(tls, buffer, libc.Xstrlen(tls, buffer)+uint64(1))
	(*mpc_parser_t)(unsafe.Pointer(p)).Fdata.Ffail.Fm = buffer
	return p
}

func mpc_lift_val(tls *libc.TLS, x uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_LIFT_VAL)
	(*(*mpc_pdata_lift_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = x
	return p
}

type __ccgo_fp__Xmpc_lift_0 = func(*libc.TLS) uintptr

func mpc_lift(tls *libc.TLS, __ccgo_fp_lf mpc_ctor_t) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_LIFT)
	(*(*mpc_pdata_lift_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Flf = __ccgo_fp_lf
	return p
}

type __ccgo_fp__Xmpc_anchor_0 = func(*libc.TLS, int8, int8) int32

func mpc_anchor(tls *libc.TLS, __ccgo_fp_f uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_ANCHOR)
	(*(*mpc_pdata_anchor_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	return mpc_expect(tls, p, __ccgo_ts+376)
}

func mpc_state(tls *libc.TLS) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_STATE)
	return p
}

func mpc_expect(tls *libc.TLS, a uintptr, expected uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_EXPECT)
	(*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm = libc.Xmalloc(tls, libc.Xstrlen(tls, expected)+uint64(1))
	libc.Xstrcpy(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm, expected)
	return p
}

/*
** As `snprintf` is not ANSI standard this
** function `mpc_expectf` should be considered
** unsafe.
**
** You have a few options if this is going to be
** trouble.
**
** - Ensure the format string does not exceed
**   the buffer length using precision specifiers
**   such as `%.512s`.
**
** - Patch this function in your code base to
**   use `snprintf` or whatever variant your
**   system supports.
**
** - Avoid it altogether.
**
 */
func mpc_expectf(tls *libc.TLS, a uintptr, fmt uintptr, va1 uintptr) (r uintptr) {
	var buffer, p uintptr
	var va va_list
	_, _, _ = buffer, p, va
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_EXPECT)
	va = va1
	buffer = libc.Xmalloc(tls, uint64(2048))
	if !(buffer != 0) {
		return libc.UintptrFromInt32(0)
	}
	libc.Xvsprintf(tls, buffer, fmt, va)
	_ = va
	buffer = libc.Xrealloc(tls, buffer, libc.Xstrlen(tls, buffer)+uint64(1))
	(*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm = buffer
	return p
}

/*
** Basic Parsers
 */
func mpc_any(tls *libc.TLS) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_ANY)
	return mpc_expect(tls, p, __ccgo_ts+383)
}

func mpc_char(tls *libc.TLS, c int8) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_SINGLE)
	(*(*mpc_pdata_single_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = c
	return mpc_expectf(tls, p, __ccgo_ts+397, libc.VaList(bp+8, int32(c)))
}

func mpc_range(tls *libc.TLS, s int8, e int8) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_RANGE)
	(*(*mpc_pdata_range_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = s
	(*(*mpc_pdata_range_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fy = e
	return mpc_expectf(tls, p, __ccgo_ts+402, libc.VaList(bp+8, int32(s), int32(e)))
}

func mpc_oneof(tls *libc.TLS, s uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_ONEOF)
	(*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = libc.Xmalloc(tls, libc.Xstrlen(tls, s)+uint64(1))
	libc.Xstrcpy(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, s)
	return mpc_expectf(tls, p, __ccgo_ts+434, libc.VaList(bp+8, s))
}

func mpc_noneof(tls *libc.TLS, s uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_NONEOF)
	(*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = libc.Xmalloc(tls, libc.Xstrlen(tls, s)+uint64(1))
	libc.Xstrcpy(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, s)
	return mpc_expectf(tls, p, __ccgo_ts+446, libc.VaList(bp+8, s))
}

type __ccgo_fp__Xmpc_satisfy_0 = func(*libc.TLS, int8) int32

func mpc_satisfy(tls *libc.TLS, __ccgo_fp_f uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_SATISFY)
	(*(*mpc_pdata_satisfy_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	return mpc_expectf(tls, p, __ccgo_ts+459, libc.VaList(bp+8, __ccgo_fp_f))
}

func mpc_string(tls *libc.TLS, s uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_STRING)
	(*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = libc.Xmalloc(tls, libc.Xstrlen(tls, s)+uint64(1))
	libc.Xstrcpy(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, s)
	return mpc_expectf(tls, p, __ccgo_ts+492, libc.VaList(bp+8, s))
}

type __ccgo_fp__Xmpc_apply_1 = func(*libc.TLS, uintptr) uintptr

/*
** Core Parsers
 */
func mpc_apply(tls *libc.TLS, a uintptr, __ccgo_fp_f mpc_apply_t) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_APPLY)
	(*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	return p
}

type __ccgo_fp__Xmpc_apply_to_1 = func(*libc.TLS, uintptr, uintptr) uintptr

func mpc_apply_to(tls *libc.TLS, a uintptr, __ccgo_fp_f mpc_apply_to_t, x uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_APPLY_TO)
	(*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	(*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fd = x
	return p
}

type __ccgo_fp__Xmpc_check_1 = func(*libc.TLS, uintptr)

type __ccgo_fp__Xmpc_check_2 = func(*libc.TLS, uintptr) int32

func mpc_check(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t, __ccgo_fp_f mpc_check_t, e uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_CHECK)
	(*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdx = __ccgo_fp_da
	(*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	(*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fe = libc.Xmalloc(tls, libc.Xstrlen(tls, e)+uint64(1))
	libc.Xstrcpy(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fe, e)
	return p
}

type __ccgo_fp__Xmpc_check_with_1 = func(*libc.TLS, uintptr)

type __ccgo_fp__Xmpc_check_with_2 = func(*libc.TLS, uintptr, uintptr) int32

func mpc_check_with(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t, __ccgo_fp_f mpc_check_with_t, x uintptr, e uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_CHECK_WITH)
	(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fx = a
	(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fdx = __ccgo_fp_da
	(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Ff = __ccgo_fp_f
	(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fd = x
	(*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fe = libc.Xmalloc(tls, libc.Xstrlen(tls, e)+uint64(1))
	libc.Xstrcpy(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fe, e)
	return p
}

type __ccgo_fp__Xmpc_checkf_1 = func(*libc.TLS, uintptr)

type __ccgo_fp__Xmpc_checkf_2 = func(*libc.TLS, uintptr) int32

func mpc_checkf(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t, __ccgo_fp_f mpc_check_t, fmt uintptr, va1 uintptr) (r uintptr) {
	var buffer, p uintptr
	var va va_list
	_, _, _ = buffer, p, va
	va = va1
	buffer = libc.Xmalloc(tls, uint64(2048))
	libc.Xvsprintf(tls, buffer, fmt, va)
	_ = va
	p = mpc_check(tls, a, __ccgo_fp_da, __ccgo_fp_f, buffer)
	libc.Xfree(tls, buffer)
	return p
}

type __ccgo_fp__Xmpc_check_withf_1 = func(*libc.TLS, uintptr)

type __ccgo_fp__Xmpc_check_withf_2 = func(*libc.TLS, uintptr, uintptr) int32

func mpc_check_withf(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t, __ccgo_fp_f mpc_check_with_t, x uintptr, fmt uintptr, va1 uintptr) (r uintptr) {
	var buffer, p uintptr
	var va va_list
	_, _, _ = buffer, p, va
	va = va1
	buffer = libc.Xmalloc(tls, uint64(2048))
	libc.Xvsprintf(tls, buffer, fmt, va)
	_ = va
	p = mpc_check_with(tls, a, __ccgo_fp_da, __ccgo_fp_f, x, buffer)
	libc.Xfree(tls, buffer)
	return p
}

func mpc_predictive(tls *libc.TLS, a uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_PREDICT)
	(*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	return p
}

type __ccgo_fp__Xmpc_not_lift_1 = func(*libc.TLS, uintptr)

type __ccgo_fp__Xmpc_not_lift_2 = func(*libc.TLS) uintptr

func mpc_not_lift(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t, __ccgo_fp_lf mpc_ctor_t) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_NOT)
	(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdx = __ccgo_fp_da
	(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Flf = __ccgo_fp_lf
	return p
}

type __ccgo_fp__Xmpc_not_1 = func(*libc.TLS, uintptr)

func mpc_not(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t) (r uintptr) {
	return mpc_not_lift(tls, a, __ccgo_fp_da, __ccgo_fp(mpcf_ctor_null))
}

type __ccgo_fp__Xmpc_maybe_lift_1 = func(*libc.TLS) uintptr

func mpc_maybe_lift(tls *libc.TLS, a uintptr, __ccgo_fp_lf mpc_ctor_t) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_MAYBE)
	(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Flf = __ccgo_fp_lf
	return p
}

func mpc_maybe(tls *libc.TLS, a uintptr) (r uintptr) {
	return mpc_maybe_lift(tls, a, __ccgo_fp(mpcf_ctor_null))
}

type __ccgo_fp__Xmpc_many_0 = func(*libc.TLS, int32, uintptr) uintptr

func mpc_many(tls *libc.TLS, __ccgo_fp_f mpc_fold_t, a uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_MANY)
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	return p
}

type __ccgo_fp__Xmpc_many1_0 = func(*libc.TLS, int32, uintptr) uintptr

func mpc_many1(tls *libc.TLS, __ccgo_fp_f mpc_fold_t, a uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_MANY1)
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	return p
}

type __ccgo_fp__Xmpc_count_1 = func(*libc.TLS, int32, uintptr) uintptr

type __ccgo_fp__Xmpc_count_3 = func(*libc.TLS, uintptr)

func mpc_count(tls *libc.TLS, n int32, __ccgo_fp_f mpc_fold_t, a uintptr, __ccgo_fp_da mpc_dtor_t) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_COUNT)
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdx = __ccgo_fp_da
	return p
}

type __ccgo_fp__Xmpc_sepby1_0 = func(*libc.TLS, int32, uintptr) uintptr

func mpc_sepby1(tls *libc.TLS, __ccgo_fp_f mpc_fold_t, sep uintptr, a uintptr) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_SEPBY1)
	(*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx = a
	(*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	(*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fsep = sep
	return p
}

func mpc_or(tls *libc.TLS, n int32, va1 uintptr) (r uintptr) {
	var i int32
	var p uintptr
	var va va_list
	_, _, _ = i, p, va
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_OR)
	(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n
	(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n))
	va = va1
	i = 0
	for {
		if !(i < n) {
			break
		}
		**(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)) = libc.VaUintptr(&va)
		goto _1
	_1:
		;
		i = i + 1
	}
	_ = va
	return p
}

type __ccgo_fp__Xmpc_and_1 = func(*libc.TLS, int32, uintptr) uintptr

func mpc_and(tls *libc.TLS, n int32, __ccgo_fp_f mpc_fold_t, va1 uintptr) (r uintptr) {
	var i int32
	var p uintptr
	var va va_list
	_, _, _ = i, p, va
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_AND)
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp_f
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n))
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n-libc.Int32FromInt32(1)))
	va = va1
	i = 0
	for {
		if !(i < n) {
			break
		}
		**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)) = libc.VaUintptr(&va)
		goto _1
	_1:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < n-int32(1)) {
			break
		}
		**(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(i)*8)) = libc.VaUintptr(&va)
		goto _2
	_2:
		;
		i = i + 1
	}
	_ = va
	return p
}

/*
** Common Parsers
 */
func mpc_soi(tls *libc.TLS) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_SOI)
	return mpc_expect(tls, p, __ccgo_ts+497)
}

func mpc_eoi(tls *libc.TLS) (r uintptr) {
	var p uintptr
	_ = p
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_EOI)
	return mpc_expect(tls, p, __ccgo_ts+56)
}

func mpc_boundary_anchor(tls *libc.TLS, prev int8, next int8) (r int32) {
	var word uintptr
	_ = word
	word = __ccgo_ts + 512
	if libc.Xstrchr(tls, word, int32(next)) != 0 && int32(prev) == int32('\000') {
		return int32(1)
	}
	if libc.Xstrchr(tls, word, int32(prev)) != 0 && int32(next) == int32('\000') {
		return int32(1)
	}
	if libc.Xstrchr(tls, word, int32(next)) != 0 && !(libc.Xstrchr(tls, word, int32(prev)) != 0) {
		return int32(1)
	}
	if !(libc.Xstrchr(tls, word, int32(next)) != 0) && libc.Xstrchr(tls, word, int32(prev)) != 0 {
		return int32(1)
	}
	return 0
}

func mpc_boundary_newline_anchor(tls *libc.TLS, prev int8, next int8) (r int32) {
	_ = next
	return libc.BoolInt32(int32(prev) == int32('\n'))
}

func mpc_boundary(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_anchor(tls, __ccgo_fp(mpc_boundary_anchor)), __ccgo_ts+576)
}

func mpc_boundary_newline(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_anchor(tls, __ccgo_fp(mpc_boundary_newline_anchor)), __ccgo_ts+590)
}

func mpc_whitespace(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_oneof(tls, __ccgo_ts+607), __ccgo_ts+614)
}

func mpc_whitespaces(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_many(tls, __ccgo_fp(mpcf_strfold), mpc_whitespace(tls)), __ccgo_ts+625)
}

func mpc_blank(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_apply(tls, mpc_whitespaces(tls), __ccgo_fp(mpcf_free)), __ccgo_ts+614)
}

func mpc_newline(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_char(tls, int8('\n')), __ccgo_ts+69)
}

func mpc_tab(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_char(tls, int8('\t')), __ccgo_ts+77)
}

func mpc_escape(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_strfold), libc.VaList(bp+8, mpc_char(tls, int8('\\')), mpc_any(tls), __ccgo_fp(libc.Xfree)))
}

func mpc_digit(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_oneof(tls, __ccgo_ts+632), __ccgo_ts+643)
}

func mpc_hexdigit(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_oneof(tls, __ccgo_ts+649), __ccgo_ts+672)
}

func mpc_octdigit(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_oneof(tls, __ccgo_ts+682), __ccgo_ts+691)
}

func mpc_digits(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_many1(tls, __ccgo_fp(mpcf_strfold), mpc_digit(tls)), __ccgo_ts+701)
}

func mpc_hexdigits(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_many1(tls, __ccgo_fp(mpcf_strfold), mpc_hexdigit(tls)), __ccgo_ts+708)
}

func mpc_octdigits(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_many1(tls, __ccgo_fp(mpcf_strfold), mpc_octdigit(tls)), __ccgo_ts+719)
}

func mpc_lower(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_oneof(tls, __ccgo_ts+730), __ccgo_ts+757)
}

func mpc_upper(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_oneof(tls, __ccgo_ts+774), __ccgo_ts+801)
}

func mpc_alpha(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_oneof(tls, __ccgo_ts+818), __ccgo_ts+871)
}

func mpc_underscore(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_char(tls, int8('_')), __ccgo_ts+878)
}

func mpc_alphanum(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_expect(tls, mpc_or(tls, int32(3), libc.VaList(bp+8, mpc_alpha(tls), mpc_digit(tls), mpc_underscore(tls))), __ccgo_ts+889)
}

func mpc_int(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_apply(tls, mpc_digits(tls), __ccgo_fp(mpcf_int)), __ccgo_ts+902)
}

func mpc_hex(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_apply(tls, mpc_hexdigits(tls), __ccgo_fp(mpcf_hex)), __ccgo_ts+910)
}

func mpc_oct(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_apply(tls, mpc_octdigits(tls), __ccgo_fp(mpcf_oct)), __ccgo_ts+922)
}

func mpc_number(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_expect(tls, mpc_or(tls, int32(3), libc.VaList(bp+8, mpc_int(tls), mpc_hex(tls), mpc_oct(tls))), __ccgo_ts+934)
}

func mpc_real(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var p0, p1, p2, p3, p30, p31, p32 uintptr
	_, _, _, _, _, _, _ = p0, p1, p2, p3, p30, p31, p32
	p0 = mpc_maybe_lift(tls, mpc_oneof(tls, __ccgo_ts+941), __ccgo_fp(mpcf_ctor_str))
	p1 = mpc_digits(tls)
	p2 = mpc_maybe_lift(tls, mpc_and(tls, int32(2), __ccgo_fp(mpcf_strfold), libc.VaList(bp+8, mpc_char(tls, int8('.')), mpc_digits(tls), __ccgo_fp(libc.Xfree))), __ccgo_fp(mpcf_ctor_str))
	p30 = mpc_oneof(tls, __ccgo_ts+944)
	p31 = mpc_maybe_lift(tls, mpc_oneof(tls, __ccgo_ts+941), __ccgo_fp(mpcf_ctor_str))
	p32 = mpc_digits(tls)
	p3 = mpc_maybe_lift(tls, mpc_and(tls, int32(3), __ccgo_fp(mpcf_strfold), libc.VaList(bp+8, p30, p31, p32, __ccgo_fp(libc.Xfree), __ccgo_fp(libc.Xfree))), __ccgo_fp(mpcf_ctor_str))
	return mpc_expect(tls, mpc_and(tls, int32(4), __ccgo_fp(mpcf_strfold), libc.VaList(bp+8, p0, p1, p2, p3, __ccgo_fp(libc.Xfree), __ccgo_fp(libc.Xfree), __ccgo_fp(libc.Xfree))), __ccgo_ts+947)
}

func mpc_float(tls *libc.TLS) (r uintptr) {
	return mpc_expect(tls, mpc_apply(tls, mpc_real(tls), __ccgo_fp(mpcf_float)), __ccgo_ts+952)
}

func mpc_char_lit(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_expect(tls, mpc_between(tls, mpc_or(tls, int32(2), libc.VaList(bp+8, mpc_escape(tls), mpc_any(tls))), __ccgo_fp(libc.Xfree), __ccgo_ts+958, __ccgo_ts+958), __ccgo_ts+960)
}

func mpc_string_lit(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var strchar uintptr
	_ = strchar
	strchar = mpc_or(tls, int32(2), libc.VaList(bp+8, mpc_escape(tls), mpc_noneof(tls, __ccgo_ts+965)))
	return mpc_expect(tls, mpc_between(tls, mpc_many(tls, __ccgo_fp(mpcf_strfold), strchar), __ccgo_fp(libc.Xfree), __ccgo_ts+965, __ccgo_ts+965), __ccgo_ts+967)
}

func mpc_regex_lit(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var regexchar uintptr
	_ = regexchar
	regexchar = mpc_or(tls, int32(2), libc.VaList(bp+8, mpc_escape(tls), mpc_noneof(tls, __ccgo_ts+974)))
	return mpc_expect(tls, mpc_between(tls, mpc_many(tls, __ccgo_fp(mpcf_strfold), regexchar), __ccgo_fp(libc.Xfree), __ccgo_ts+974, __ccgo_ts+974), __ccgo_ts+976)
}

func mpc_ident(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var p0, p1 uintptr
	_, _ = p0, p1
	p0 = mpc_or(tls, int32(2), libc.VaList(bp+8, mpc_alpha(tls), mpc_underscore(tls)))
	p1 = mpc_many(tls, __ccgo_fp(mpcf_strfold), mpc_alphanum(tls))
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_strfold), libc.VaList(bp+8, p0, p1, __ccgo_fp(libc.Xfree)))
}

/*
** Useful Parsers
 */
func mpc_startwith(tls *libc.TLS, a uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_soi(tls), a, __ccgo_fp(mpcf_dtor_null)))
}

type __ccgo_fp__Xmpc_endwith_1 = func(*libc.TLS, uintptr)

func mpc_endwith(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_fst), libc.VaList(bp+8, a, mpc_eoi(tls), __ccgo_fp_da))
}

type __ccgo_fp__Xmpc_whole_1 = func(*libc.TLS, uintptr)

func mpc_whole(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	return mpc_and(tls, int32(3), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_soi(tls), a, mpc_eoi(tls), __ccgo_fp(mpcf_dtor_null), __ccgo_fp_da))
}

func mpc_stripl(tls *libc.TLS, a uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_blank(tls), a, __ccgo_fp(mpcf_dtor_null)))
}

func mpc_stripr(tls *libc.TLS, a uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_fst), libc.VaList(bp+8, a, mpc_blank(tls), __ccgo_fp(mpcf_dtor_null)))
}

func mpc_strip(tls *libc.TLS, a uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	return mpc_and(tls, int32(3), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_blank(tls), a, mpc_blank(tls), __ccgo_fp(mpcf_dtor_null), __ccgo_fp(mpcf_dtor_null)))
}

func mpc_tok(tls *libc.TLS, a uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_fst), libc.VaList(bp+8, a, mpc_blank(tls), __ccgo_fp(mpcf_dtor_null)))
}

func mpc_sym(tls *libc.TLS, s uintptr) (r uintptr) {
	return mpc_tok(tls, mpc_string(tls, s))
}

type __ccgo_fp__Xmpc_total_1 = func(*libc.TLS, uintptr)

func mpc_total(tls *libc.TLS, a uintptr, __ccgo_fp_da mpc_dtor_t) (r uintptr) {
	return mpc_whole(tls, mpc_strip(tls, a), __ccgo_fp_da)
}

type __ccgo_fp__Xmpc_between_1 = func(*libc.TLS, uintptr)

func mpc_between(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t, o uintptr, c uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	return mpc_and(tls, int32(3), __ccgo_fp(mpcf_snd_free), libc.VaList(bp+8, mpc_string(tls, o), a, mpc_string(tls, c), __ccgo_fp(libc.Xfree), __ccgo_fp_ad))
}

type __ccgo_fp__Xmpc_parens_1 = func(*libc.TLS, uintptr)

func mpc_parens(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_between(tls, a, __ccgo_fp_ad, __ccgo_ts+982, __ccgo_ts+984)
}

type __ccgo_fp__Xmpc_braces_1 = func(*libc.TLS, uintptr)

func mpc_braces(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_between(tls, a, __ccgo_fp_ad, __ccgo_ts+986, __ccgo_ts+988)
}

type __ccgo_fp__Xmpc_brackets_1 = func(*libc.TLS, uintptr)

func mpc_brackets(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_between(tls, a, __ccgo_fp_ad, __ccgo_ts+990, __ccgo_ts+992)
}

type __ccgo_fp__Xmpc_squares_1 = func(*libc.TLS, uintptr)

func mpc_squares(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_between(tls, a, __ccgo_fp_ad, __ccgo_ts+994, __ccgo_ts+996)
}

type __ccgo_fp__Xmpc_tok_between_1 = func(*libc.TLS, uintptr)

func mpc_tok_between(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t, o uintptr, c uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	return mpc_and(tls, int32(3), __ccgo_fp(mpcf_snd_free), libc.VaList(bp+8, mpc_sym(tls, o), mpc_tok(tls, a), mpc_sym(tls, c), __ccgo_fp(libc.Xfree), __ccgo_fp_ad))
}

type __ccgo_fp__Xmpc_tok_parens_1 = func(*libc.TLS, uintptr)

func mpc_tok_parens(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_tok_between(tls, a, __ccgo_fp_ad, __ccgo_ts+982, __ccgo_ts+984)
}

type __ccgo_fp__Xmpc_tok_braces_1 = func(*libc.TLS, uintptr)

func mpc_tok_braces(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_tok_between(tls, a, __ccgo_fp_ad, __ccgo_ts+986, __ccgo_ts+988)
}

type __ccgo_fp__Xmpc_tok_brackets_1 = func(*libc.TLS, uintptr)

func mpc_tok_brackets(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_tok_between(tls, a, __ccgo_fp_ad, __ccgo_ts+990, __ccgo_ts+992)
}

type __ccgo_fp__Xmpc_tok_squares_1 = func(*libc.TLS, uintptr)

func mpc_tok_squares(tls *libc.TLS, a uintptr, __ccgo_fp_ad mpc_dtor_t) (r uintptr) {
	return mpc_tok_between(tls, a, __ccgo_fp_ad, __ccgo_ts+994, __ccgo_ts+996)
}

/*
** Regular Expression Parsers
 */

/*
** So here is a cute bootstrapping.
**
** I'm using the previously defined
** mpc constructs and functions to
** parse the user regex string and
** construct a parser from it.
**
** As it turns out lots of the standard
** mpc functions look a lot like `fold`
** functions and so can be used indirectly
** by many of the parsing functions to build
** a parser directly - as we are parsing.
**
** This is certainly something that
** would be less elegant/interesting
** in a two-phase parser which first
** builds an AST and then traverses it
** to generate the object.
**
** This whole thing acts as a great
** case study for how trivial it can be
** to write a great parser in a few
** lines of code using mpc.
 */

/*
**
**  ### Regular Expression Grammar
**
**      <regex> : <term> | (<term> "|" <regex>)
**
**      <term> : <factor>*
**
**      <factor> : <base>
**               | <base> "*"
**               | <base> "+"
**               | <base> "?"
**               | <base> "{" <digits> "}"
**
**      <base> : <char>
**             | "\" <char>
**             | "(" <regex> ")"
**             | "[" <range> "]"
 */
func mpcf_re_or(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	_ = n
	if **(**uintptr)(__ccgo_up(xs + 1*8)) == libc.UintptrFromInt32(0) {
		return **(**uintptr)(__ccgo_up(xs))
	} else {
		return mpc_or(tls, int32(2), libc.VaList(bp+8, **(**uintptr)(__ccgo_up(xs)), **(**uintptr)(__ccgo_up(xs + 1*8))))
	}
	return r
}

func mpcf_re_and(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var p uintptr
	_, _ = i, p
	p = mpc_lift(tls, __ccgo_fp(mpcf_ctor_str))
	i = 0
	for {
		if !(i < n) {
			break
		}
		p = mpc_and(tls, int32(2), __ccgo_fp(mpcf_strfold), libc.VaList(bp+8, p, **(**uintptr)(__ccgo_up(xs + uintptr(i)*8)), __ccgo_fp(libc.Xfree)))
		goto _1
	_1:
		;
		i = i + 1
	}
	return p
}

func mpcf_re_repeat(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var num int32
	_ = num
	_ = n
	if **(**uintptr)(__ccgo_up(xs + 1*8)) == libc.UintptrFromInt32(0) {
		return **(**uintptr)(__ccgo_up(xs))
	}
	switch int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(xs + 1*8))))) {
	case int32('*'):
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
		return mpc_many(tls, __ccgo_fp(mpcf_strfold), **(**uintptr)(__ccgo_up(xs)))
	case int32('+'):
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
		return mpc_many1(tls, __ccgo_fp(mpcf_strfold), **(**uintptr)(__ccgo_up(xs)))
	case int32('?'):
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
		return mpc_maybe_lift(tls, **(**uintptr)(__ccgo_up(xs)), __ccgo_fp(mpcf_ctor_str))
	default:
		num = **(**int32)(__ccgo_up(**(**uintptr)(__ccgo_up(xs + 1*8))))
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
	}
	return mpc_count(tls, num, __ccgo_fp(mpcf_strfold), **(**uintptr)(__ccgo_up(xs)), __ccgo_fp(libc.Xfree))
}

func mpc_re_escape_char(tls *libc.TLS, c int8) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	switch int32(c) {
	case int32('a'):
		return mpc_char(tls, int8('\a'))
	case int32('f'):
		return mpc_char(tls, int8('\f'))
	case int32('n'):
		return mpc_char(tls, int8('\n'))
	case int32('r'):
		return mpc_char(tls, int8('\r'))
	case int32('t'):
		return mpc_char(tls, int8('\t'))
	case int32('v'):
		return mpc_char(tls, int8('\v'))
	case int32('b'):
		return mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_boundary(tls), mpc_lift(tls, __ccgo_fp(mpcf_ctor_str)), __ccgo_fp(libc.Xfree)))
	case int32('B'):
		return mpc_not_lift(tls, mpc_boundary(tls), __ccgo_fp(libc.Xfree), __ccgo_fp(mpcf_ctor_str))
	case int32('A'):
		return mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_soi(tls), mpc_lift(tls, __ccgo_fp(mpcf_ctor_str)), __ccgo_fp(libc.Xfree)))
	case int32('Z'):
		return mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_eoi(tls), mpc_lift(tls, __ccgo_fp(mpcf_ctor_str)), __ccgo_fp(libc.Xfree)))
	case int32('d'):
		return mpc_digit(tls)
	case int32('D'):
		return mpc_not_lift(tls, mpc_digit(tls), __ccgo_fp(libc.Xfree), __ccgo_fp(mpcf_ctor_str))
	case int32('s'):
		return mpc_whitespace(tls)
	case int32('S'):
		return mpc_not_lift(tls, mpc_whitespace(tls), __ccgo_fp(libc.Xfree), __ccgo_fp(mpcf_ctor_str))
	case int32('w'):
		return mpc_alphanum(tls)
	case int32('W'):
		return mpc_not_lift(tls, mpc_alphanum(tls), __ccgo_fp(libc.Xfree), __ccgo_fp(mpcf_ctor_str))
	default:
		return libc.UintptrFromInt32(0)
	}
	return r
}

func mpcf_re_escape(tls *libc.TLS, x uintptr, data uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var mode int32
	var p, s, v1 uintptr
	_, _, _, _ = mode, p, s, v1
	mode = **(**int32)(__ccgo_up(data))
	s = x
	/* Any Character */
	if int32(**(**int8)(__ccgo_up(s))) == int32('.') {
		libc.Xfree(tls, s)
		if mode&int32(MPC_RE_DOTALL) != 0 {
			return mpc_any(tls)
		} else {
			return mpc_expect(tls, mpc_noneof(tls, __ccgo_ts+174), __ccgo_ts+998)
		}
	}
	/* Start of Input */
	if int32(**(**int8)(__ccgo_up(s))) == int32('^') {
		libc.Xfree(tls, s)
		if mode&int32(MPC_RE_MULTILINE) != 0 {
			return mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_or(tls, int32(2), libc.VaList(bp+8, mpc_soi(tls), mpc_boundary_newline(tls))), mpc_lift(tls, __ccgo_fp(mpcf_ctor_str)), __ccgo_fp(libc.Xfree)))
		} else {
			return mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_soi(tls), mpc_lift(tls, __ccgo_fp(mpcf_ctor_str)), __ccgo_fp(libc.Xfree)))
		}
	}
	/* End of Input */
	if int32(**(**int8)(__ccgo_up(s))) == int32('$') {
		libc.Xfree(tls, s)
		if mode&int32(MPC_RE_MULTILINE) != 0 {
			return mpc_or(tls, int32(2), libc.VaList(bp+8, mpc_newline(tls), mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_eoi(tls), mpc_lift(tls, __ccgo_fp(mpcf_ctor_str)), __ccgo_fp(libc.Xfree)))))
		} else {
			return mpc_or(tls, int32(2), libc.VaList(bp+8, mpc_and(tls, int32(2), __ccgo_fp(mpcf_fst), libc.VaList(bp+8, mpc_newline(tls), mpc_eoi(tls), __ccgo_fp(libc.Xfree))), mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd), libc.VaList(bp+8, mpc_eoi(tls), mpc_lift(tls, __ccgo_fp(mpcf_ctor_str)), __ccgo_fp(libc.Xfree)))))
		}
	}
	/* Regex Escape */
	if int32(**(**int8)(__ccgo_up(s))) == int32('\\') {
		p = mpc_re_escape_char(tls, **(**int8)(__ccgo_up(s + 1)))
		if p == libc.UintptrFromInt32(0) {
			v1 = mpc_char(tls, **(**int8)(__ccgo_up(s + 1)))
		} else {
			v1 = p
		}
		p = v1
		libc.Xfree(tls, s)
		return p
	}
	/* Regex Standard */
	p = mpc_char(tls, **(**int8)(__ccgo_up(s)))
	libc.Xfree(tls, s)
	return p
}

func mpc_re_range_escape_char(tls *libc.TLS, c int8) (r uintptr) {
	switch int32(c) {
	case int32('-'):
		return __ccgo_ts + 1029
	case int32('a'):
		return __ccgo_ts + 1031
	case int32('f'):
		return __ccgo_ts + 1033
	case int32('n'):
		return __ccgo_ts + 174
	case int32('r'):
		return __ccgo_ts + 1035
	case int32('t'):
		return __ccgo_ts + 1037
	case int32('v'):
		return __ccgo_ts + 1039
	case int32('b'):
		return __ccgo_ts + 1041
	case int32('d'):
		return __ccgo_ts + 632
	case int32('s'):
		return __ccgo_ts + 607
	case int32('w'):
		return __ccgo_ts + 512
	default:
		return libc.UintptrFromInt32(0)
	}
	return r
}

func mpcf_re_range(tls *libc.TLS, x uintptr) (r uintptr) {
	var comp, v1 int32
	var end, i, j, start size_t
	var out, range1, s, tmp, v4 uintptr
	_, _, _, _, _, _, _, _, _, _, _ = comp, end, i, j, out, range1, s, start, tmp, v1, v4
	tmp = libc.UintptrFromInt32(0)
	s = x
	if int32(**(**int8)(__ccgo_up(s))) == int32('^') {
		v1 = int32(1)
	} else {
		v1 = 0
	}
	comp = v1
	range1 = libc.Xcalloc(tls, uint64(1), uint64(1))
	if int32(**(**int8)(__ccgo_up(s))) == int32('\000') {
		libc.Xfree(tls, range1)
		libc.Xfree(tls, x)
		return mpc_fail(tls, __ccgo_ts+1043)
	}
	if int32(**(**int8)(__ccgo_up(s))) == int32('^') && int32(**(**int8)(__ccgo_up(s + 1))) == int32('\000') {
		libc.Xfree(tls, range1)
		libc.Xfree(tls, x)
		return mpc_fail(tls, __ccgo_ts+1043)
	}
	i = libc.Uint64FromInt32(comp)
	for {
		if !(i < libc.Xstrlen(tls, s)) {
			break
		}
		/* Regex Range Escape */
		if int32(**(**int8)(__ccgo_up(s + uintptr(i)))) == int32('\\') {
			tmp = mpc_re_range_escape_char(tls, **(**int8)(__ccgo_up(s + uintptr(i+uint64(1)))))
			if tmp != libc.UintptrFromInt32(0) {
				range1 = libc.Xrealloc(tls, range1, libc.Xstrlen(tls, range1)+libc.Xstrlen(tls, tmp)+uint64(1))
				libc.Xstrcat(tls, range1, tmp)
			} else {
				range1 = libc.Xrealloc(tls, range1, libc.Xstrlen(tls, range1)+uint64(1)+uint64(1))
				**(**int8)(__ccgo_up(range1 + uintptr(libc.Xstrlen(tls, range1)+uint64(1)))) = int8('\000')
				**(**int8)(__ccgo_up(range1 + uintptr(libc.Xstrlen(tls, range1)+uint64(0)))) = **(**int8)(__ccgo_up(s + uintptr(i+uint64(1))))
			}
			i = i + 1
		} else {
			if int32(**(**int8)(__ccgo_up(s + uintptr(i)))) == int32('-') {
				if int32(**(**int8)(__ccgo_up(s + uintptr(i+uint64(1))))) == int32('\000') || i == uint64(0) {
					range1 = libc.Xrealloc(tls, range1, libc.Xstrlen(tls, range1)+libc.Xstrlen(tls, __ccgo_ts+1029)+uint64(1))
					libc.Xstrcat(tls, range1, __ccgo_ts+1029)
				} else {
					start = libc.Uint64FromInt32(int32(**(**int8)(__ccgo_up(s + uintptr(i-uint64(1))))) + int32(1))
					end = libc.Uint64FromInt32(int32(**(**int8)(__ccgo_up(s + uintptr(i+uint64(1))))) - int32(1))
					j = start
					for {
						if !(j <= end) {
							break
						}
						range1 = libc.Xrealloc(tls, range1, libc.Xstrlen(tls, range1)+uint64(1)+uint64(1)+uint64(1))
						**(**int8)(__ccgo_up(range1 + uintptr(libc.Xstrlen(tls, range1)+uint64(1)))) = int8('\000')
						**(**int8)(__ccgo_up(range1 + uintptr(libc.Xstrlen(tls, range1)+uint64(0)))) = libc.Int8FromUint64(j)
						goto _3
					_3:
						;
						j = j + 1
					}
				}
			} else {
				range1 = libc.Xrealloc(tls, range1, libc.Xstrlen(tls, range1)+uint64(1)+uint64(1))
				**(**int8)(__ccgo_up(range1 + uintptr(libc.Xstrlen(tls, range1)+uint64(1)))) = int8('\000')
				**(**int8)(__ccgo_up(range1 + uintptr(libc.Xstrlen(tls, range1)+uint64(0)))) = **(**int8)(__ccgo_up(s + uintptr(i)))
			}
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	if comp == int32(1) {
		v4 = mpc_noneof(tls, range1)
	} else {
		v4 = mpc_oneof(tls, range1)
	}
	out = v4
	libc.Xfree(tls, x)
	libc.Xfree(tls, range1)
	return out
}

func mpc_re(tls *libc.TLS, re uintptr) (r uintptr) {
	return mpc_re_mode(tls, re, int32(MPC_RE_DEFAULT))
}

func mpc_re_mode(tls *libc.TLS, re uintptr, _mode int32) (r uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*int32)(unsafe.Pointer(bp)) = _mode
	var Base, Factor, Range, Regex, RegexEnclose, Term, err_msg, err_out uintptr
	var _ /* r at bp+8 */ mpc_result_t
	_, _, _, _, _, _, _, _ = Base, Factor, Range, Regex, RegexEnclose, Term, err_msg, err_out
	Regex = mpc_new(tls, __ccgo_ts+976)
	Term = mpc_new(tls, __ccgo_ts+1074)
	Factor = mpc_new(tls, __ccgo_ts+1079)
	Base = mpc_new(tls, __ccgo_ts+1086)
	Range = mpc_new(tls, __ccgo_ts+1091)
	mpc_define(tls, Regex, mpc_and(tls, int32(2), __ccgo_fp(mpcf_re_or), libc.VaList(bp+24, Term, mpc_maybe(tls, mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd_free), libc.VaList(bp+24, mpc_char(tls, int8('|')), Regex, __ccgo_fp(libc.Xfree)))), __ccgo_fp(mpc_delete))))
	mpc_define(tls, Term, mpc_many(tls, __ccgo_fp(mpcf_re_and), Factor))
	mpc_define(tls, Factor, mpc_and(tls, int32(2), __ccgo_fp(mpcf_re_repeat), libc.VaList(bp+24, Base, mpc_or(tls, int32(5), libc.VaList(bp+24, mpc_char(tls, int8('*')), mpc_char(tls, int8('+')), mpc_char(tls, int8('?')), mpc_brackets(tls, mpc_int(tls), __ccgo_fp(libc.Xfree)), mpc_pass(tls))), __ccgo_fp(mpc_delete))))
	mpc_define(tls, Base, mpc_or(tls, int32(4), libc.VaList(bp+24, mpc_parens(tls, Regex, __ccgo_fp(mpc_delete)), mpc_squares(tls, Range, __ccgo_fp(mpc_delete)), mpc_apply_to(tls, mpc_escape(tls), __ccgo_fp(mpcf_re_escape), bp), mpc_apply_to(tls, mpc_noneof(tls, __ccgo_ts+1097), __ccgo_fp(mpcf_re_escape), bp))))
	mpc_define(tls, Range, mpc_apply(tls, mpc_many(tls, __ccgo_fp(mpcf_strfold), mpc_or(tls, int32(2), libc.VaList(bp+24, mpc_escape(tls), mpc_noneof(tls, __ccgo_ts+996)))), __ccgo_fp(mpcf_re_range)))
	RegexEnclose = mpc_whole(tls, mpc_predictive(tls, Regex), __ccgo_fp(mpc_delete))
	mpc_optimise(tls, RegexEnclose)
	mpc_optimise(tls, Regex)
	mpc_optimise(tls, Term)
	mpc_optimise(tls, Factor)
	mpc_optimise(tls, Base)
	mpc_optimise(tls, Range)
	if !(mpc_parse(tls, __ccgo_ts+1100, re, RegexEnclose, bp+8) != 0) {
		err_msg = mpc_err_string(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
		err_out = mpc_failf(tls, __ccgo_ts+1118, libc.VaList(bp+24, err_msg))
		mpc_err_delete(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
		libc.Xfree(tls, err_msg)
		*(*uintptr)(unsafe.Pointer(bp + 8)) = err_out
	}
	mpc_cleanup(tls, int32(6), libc.VaList(bp+24, RegexEnclose, Regex, Term, Factor, Base, Range))
	mpc_optimise(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
	return *(*uintptr)(unsafe.Pointer(bp + 8))
}

/*
** Common Fold Functions
 */
func mpcf_dtor_null(tls *libc.TLS, x uintptr) {
	_ = x
	return
}

func mpcf_ctor_null(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func mpcf_ctor_str(tls *libc.TLS) (r uintptr) {
	return libc.Xcalloc(tls, uint64(1), uint64(1))
}

func mpcf_free(tls *libc.TLS, x uintptr) (r uintptr) {
	libc.Xfree(tls, x)
	return libc.UintptrFromInt32(0)
}

func mpcf_int(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = libc.Xmalloc(tls, uint64(4))
	**(**int32)(__ccgo_up(y)) = int32(libc.Xstrtol(tls, x, libc.UintptrFromInt32(0), int32(10)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_hex(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = libc.Xmalloc(tls, uint64(4))
	**(**int32)(__ccgo_up(y)) = int32(libc.Xstrtol(tls, x, libc.UintptrFromInt32(0), int32(16)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_oct(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = libc.Xmalloc(tls, uint64(4))
	**(**int32)(__ccgo_up(y)) = int32(libc.Xstrtol(tls, x, libc.UintptrFromInt32(0), int32(8)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_float(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = libc.Xmalloc(tls, uint64(4))
	**(**float32)(__ccgo_up(y)) = float32(libc.Xstrtod(tls, x, libc.UintptrFromInt32(0)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_strtriml(tls *libc.TLS, x uintptr) (r uintptr) {
	var s uintptr
	var v1, v2 int32
	_, _, _ = s, v1, v2
	s = x
	for {
		v1 = libc.Int32FromUint8(libc.Uint8FromInt8(**(**int8)(__ccgo_up(s))))
		v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
		goto _3
	_3:
		if !(v2 != 0) {
			break
		}
		libc.Xmemmove(tls, s, s+uintptr(1), libc.Xstrlen(tls, s))
	}
	return s
}

func mpcf_strtrimr(tls *libc.TLS, x uintptr) (r uintptr) {
	var l size_t
	var s uintptr
	var v1, v2 int32
	var v4 bool
	_, _, _, _, _ = l, s, v1, v2, v4
	s = x
	l = libc.Xstrlen(tls, s)
	for {
		if v4 = l > uint64(0); v4 {
			v1 = libc.Int32FromUint8(libc.Uint8FromInt8(**(**int8)(__ccgo_up(s + uintptr(l-uint64(1))))))
			v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
			goto _3
		_3:
		}
		if !(v4 && v2 != 0) {
			break
		}
		**(**int8)(__ccgo_up(s + uintptr(l-uint64(1)))) = int8('\000')
		l = l - 1
	}
	return s
}

func mpcf_strtrim(tls *libc.TLS, x uintptr) (r uintptr) {
	return mpcf_strtriml(tls, mpcf_strtrimr(tls, x))
}

var mpc_escape_input_c = [11]int8{
	0: int8('\a'),
	1: int8('\b'),
	2: int8('\f'),
	3: int8('\n'),
	4: int8('\r'),
	5: int8('\t'),
	6: int8('\v'),
	7: int8('\\'),
	8: int8('\''),
	9: int8('"'),
}

var mpc_escape_output_c = [12]uintptr{
	0:  __ccgo_ts + 1136,
	1:  __ccgo_ts + 1139,
	2:  __ccgo_ts + 1142,
	3:  __ccgo_ts + 1145,
	4:  __ccgo_ts + 1148,
	5:  __ccgo_ts + 1151,
	6:  __ccgo_ts + 1154,
	7:  __ccgo_ts + 1157,
	8:  __ccgo_ts + 1160,
	9:  __ccgo_ts + 1163,
	10: __ccgo_ts + 1166,
	11: libc.UintptrFromInt32(0),
}

var mpc_escape_input_raw_re = [1]int8{
	0: int8('/'),
}
var mpc_escape_output_raw_re = [2]uintptr{
	0: __ccgo_ts + 1169,
	1: libc.UintptrFromInt32(0),
}

var mpc_escape_input_raw_cstr = [1]int8{
	0: int8('"'),
}
var mpc_escape_output_raw_cstr = [2]uintptr{
	0: __ccgo_ts + 1163,
	1: libc.UintptrFromInt32(0),
}

var mpc_escape_input_raw_cchar = [1]int8{
	0: int8('\''),
}
var mpc_escape_output_raw_cchar = [2]uintptr{
	0: __ccgo_ts + 1160,
	1: libc.UintptrFromInt32(0),
}

func mpcf_escape_new(tls *libc.TLS, x uintptr, input uintptr, output uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var found, i int32
	var s, y uintptr
	var _ /* buff at bp+0 */ [2]int8
	_, _, _, _ = found, i, s, y
	s = x
	y = libc.Xcalloc(tls, uint64(1), uint64(1))
	for **(**int8)(__ccgo_up(s)) != 0 {
		i = 0
		found = 0
		for **(**uintptr)(__ccgo_up(output + uintptr(i)*8)) != 0 {
			if int32(**(**int8)(__ccgo_up(s))) == int32(**(**int8)(__ccgo_up(input + uintptr(i)))) {
				y = libc.Xrealloc(tls, y, libc.Xstrlen(tls, y)+libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(output + uintptr(i)*8)))+uint64(1))
				libc.Xstrcat(tls, y, **(**uintptr)(__ccgo_up(output + uintptr(i)*8)))
				found = int32(1)
				break
			}
			i = i + 1
		}
		if !(found != 0) {
			y = libc.Xrealloc(tls, y, libc.Xstrlen(tls, y)+uint64(2))
			(**(**[2]int8)(__ccgo_up(bp)))[0] = **(**int8)(__ccgo_up(s))
			(**(**[2]int8)(__ccgo_up(bp)))[int32(1)] = int8('\000')
			libc.Xstrcat(tls, y, bp)
		}
		s = s + 1
	}
	return y
}

func mpcf_unescape_new(tls *libc.TLS, x uintptr, input uintptr, output uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var found, i int32
	var s, y uintptr
	var _ /* buff at bp+0 */ [2]int8
	_, _, _, _ = found, i, s, y
	found = 0
	s = x
	y = libc.Xcalloc(tls, uint64(1), uint64(1))
	for **(**int8)(__ccgo_up(s)) != 0 {
		i = 0
		found = 0
		for **(**uintptr)(__ccgo_up(output + uintptr(i)*8)) != 0 {
			if int32(**(**int8)(__ccgo_up(s + libc.UintptrFromInt32(0)))) == int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(output + uintptr(i)*8))))) && int32(**(**int8)(__ccgo_up(s + libc.UintptrFromInt32(1)))) == int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(output + uintptr(i)*8)) + 1))) {
				y = libc.Xrealloc(tls, y, libc.Xstrlen(tls, y)+uint64(1)+uint64(1))
				(**(**[2]int8)(__ccgo_up(bp)))[0] = **(**int8)(__ccgo_up(input + uintptr(i)))
				(**(**[2]int8)(__ccgo_up(bp)))[int32(1)] = int8('\000')
				libc.Xstrcat(tls, y, bp)
				found = int32(1)
				s = s + 1
				break
			}
			i = i + 1
		}
		if !(found != 0) {
			y = libc.Xrealloc(tls, y, libc.Xstrlen(tls, y)+uint64(1)+uint64(1))
			(**(**[2]int8)(__ccgo_up(bp)))[0] = **(**int8)(__ccgo_up(s))
			(**(**[2]int8)(__ccgo_up(bp)))[int32(1)] = int8('\000')
			libc.Xstrcat(tls, y, bp)
		}
		if int32(**(**int8)(__ccgo_up(s))) == int32('\000') {
			break
		} else {
			s = s + 1
		}
	}
	return y
}

func mpcf_escape(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_escape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_unescape(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_unescape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_escape_regex(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_escape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_raw_re)), uintptr(unsafe.Pointer(&mpc_escape_output_raw_re)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_unescape_regex(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_unescape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_raw_re)), uintptr(unsafe.Pointer(&mpc_escape_output_raw_re)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_escape_string_raw(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_escape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_raw_cstr)), uintptr(unsafe.Pointer(&mpc_escape_output_raw_cstr)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_unescape_string_raw(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_unescape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_raw_cstr)), uintptr(unsafe.Pointer(&mpc_escape_output_raw_cstr)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_escape_char_raw(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_escape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_raw_cchar)), uintptr(unsafe.Pointer(&mpc_escape_output_raw_cchar)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_unescape_char_raw(tls *libc.TLS, x uintptr) (r uintptr) {
	var y uintptr
	_ = y
	y = mpcf_unescape_new(tls, x, uintptr(unsafe.Pointer(&mpc_escape_input_raw_cchar)), uintptr(unsafe.Pointer(&mpc_escape_output_raw_cchar)))
	libc.Xfree(tls, x)
	return y
}

func mpcf_null(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	_ = n
	_ = xs
	return libc.UintptrFromInt32(0)
}

func mpcf_fst(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	_ = n
	return **(**uintptr)(__ccgo_up(xs))
}

func mpcf_snd(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	_ = n
	return **(**uintptr)(__ccgo_up(xs + 1*8))
}

func mpcf_trd(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	_ = n
	return **(**uintptr)(__ccgo_up(xs + 2*8))
}

func mpcf_nth_free(tls *libc.TLS, n int32, xs uintptr, x int32) (r uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < n) {
			break
		}
		if i != x {
			libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + uintptr(i)*8)))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return **(**uintptr)(__ccgo_up(xs + uintptr(x)*8))
}

func mpcf_fst_free(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	return mpcf_nth_free(tls, n, xs, 0)
}

func mpcf_snd_free(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	return mpcf_nth_free(tls, n, xs, int32(1))
}

func mpcf_trd_free(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	return mpcf_nth_free(tls, n, xs, int32(2))
}

func mpcf_all_free(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < n) {
			break
		}
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.UintptrFromInt32(0)
}

func mpcf_strfold(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var i int32
	var l size_t
	_, _ = i, l
	l = uint64(0)
	if n == 0 {
		return libc.Xcalloc(tls, uint64(1), uint64(1))
	}
	i = 0
	for {
		if !(i < n) {
			break
		}
		l = l + libc.Xstrlen(tls, **(**uintptr)(__ccgo_up(xs + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	**(**uintptr)(__ccgo_up(xs)) = libc.Xrealloc(tls, **(**uintptr)(__ccgo_up(xs)), l+uint64(1))
	i = int32(1)
	for {
		if !(i < n) {
			break
		}
		libc.Xstrcat(tls, **(**uintptr)(__ccgo_up(xs)), **(**uintptr)(__ccgo_up(xs + uintptr(i)*8)))
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + uintptr(i)*8)))
		goto _2
	_2:
		;
		i = i + 1
	}
	return **(**uintptr)(__ccgo_up(xs))
}

/*
** Printing
 */
func mpc_print_unretained(tls *libc.TLS, p uintptr, force int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var e, s uintptr
	var i int32
	var _ /* buff at bp+0 */ [2]int8
	_, _, _ = e, i, s
	if (*mpc_parser_t)(unsafe.Pointer(p)).Fretained != 0 && !(force != 0) {
		if (*mpc_parser_t)(unsafe.Pointer(p)).Fname != 0 {
			libc.Xprintf(tls, __ccgo_ts+1172, libc.VaList(bp+16, (*mpc_parser_t)(unsafe.Pointer(p)).Fname))
		} else {
			libc.Xprintf(tls, __ccgo_ts+1177, 0)
		}
		return
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_UNDEFINED) {
		libc.Xprintf(tls, __ccgo_ts+1184, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_PASS) {
		libc.Xprintf(tls, __ccgo_ts+1188, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_FAIL) {
		libc.Xprintf(tls, __ccgo_ts+1192, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_LIFT) {
		libc.Xprintf(tls, __ccgo_ts+1196, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_STATE) {
		libc.Xprintf(tls, __ccgo_ts+1200, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_ANCHOR) {
		libc.Xprintf(tls, __ccgo_ts+1204, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_EXPECT) {
		libc.Xprintf(tls, __ccgo_ts, libc.VaList(bp+16, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fm))
		/*mpc_print_unretained(p->data.expect.x, 0);*/
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_ANY) {
		libc.Xprintf(tls, __ccgo_ts+1208, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_SATISFY) {
		libc.Xprintf(tls, __ccgo_ts+1212, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_SINGLE) {
		(**(**[2]int8)(__ccgo_up(bp)))[0] = (*(*mpc_pdata_single_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx
		(**(**[2]int8)(__ccgo_up(bp)))[int32(1)] = int8('\000')
		s = mpcf_escape_new(tls, bp, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
		libc.Xprintf(tls, __ccgo_ts+1216, libc.VaList(bp+16, s))
		libc.Xfree(tls, s)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_RANGE) {
		(**(**[2]int8)(__ccgo_up(bp)))[0] = (*(*mpc_pdata_range_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx
		(**(**[2]int8)(__ccgo_up(bp)))[int32(1)] = int8('\000')
		s = mpcf_escape_new(tls, bp, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
		(**(**[2]int8)(__ccgo_up(bp)))[0] = (*(*mpc_pdata_range_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fy
		(**(**[2]int8)(__ccgo_up(bp)))[int32(1)] = int8('\000')
		e = mpcf_escape_new(tls, bp, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
		libc.Xprintf(tls, __ccgo_ts+1221, libc.VaList(bp+16, s, e))
		libc.Xfree(tls, s)
		libc.Xfree(tls, e)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_ONEOF) {
		s = mpcf_escape_new(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
		libc.Xprintf(tls, __ccgo_ts+1229, libc.VaList(bp+16, s))
		libc.Xfree(tls, s)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_NONEOF) {
		s = mpcf_escape_new(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
		libc.Xprintf(tls, __ccgo_ts+1234, libc.VaList(bp+16, s))
		libc.Xfree(tls, s)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_STRING) {
		s = mpcf_escape_new(tls, (*(*mpc_pdata_string_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, uintptr(unsafe.Pointer(&mpc_escape_input_c)), uintptr(unsafe.Pointer(&mpc_escape_output_c)))
		libc.Xprintf(tls, __ccgo_ts+492, libc.VaList(bp+16, s))
		libc.Xfree(tls, s)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_APPLY) {
		mpc_print_unretained(tls, (*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_APPLY_TO) {
		mpc_print_unretained(tls, (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_PREDICT) {
		mpc_print_unretained(tls, (*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_NOT) {
		mpc_print_unretained(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1240, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MAYBE) {
		mpc_print_unretained(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1242, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MANY) {
		mpc_print_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1244, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MANY1) {
		mpc_print_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1246, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_COUNT) {
		mpc_print_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1248, libc.VaList(bp+16, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn))
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_SEPBY1) {
		mpc_print_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1253, 0)
		mpc_print_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fsep, 0)
		libc.Xprintf(tls, __ccgo_ts+1256, 0)
		mpc_print_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+984, 0)
		libc.Xprintf(tls, __ccgo_ts+1244, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_OR) {
		libc.Xprintf(tls, __ccgo_ts+982, 0)
		i = 0
		for {
			if !(i < (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1)) {
				break
			}
			mpc_print_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
			libc.Xprintf(tls, __ccgo_ts+1258, 0)
			goto _1
		_1:
			;
			i = i + 1
		}
		mpc_print_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)), 0)
		libc.Xprintf(tls, __ccgo_ts+984, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) {
		libc.Xprintf(tls, __ccgo_ts+982, 0)
		i = 0
		for {
			if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1)) {
				break
			}
			mpc_print_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
			libc.Xprintf(tls, __ccgo_ts+1256, 0)
			goto _2
		_2:
			;
			i = i + 1
		}
		mpc_print_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)), 0)
		libc.Xprintf(tls, __ccgo_ts+984, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_CHECK) {
		mpc_print_unretained(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1262, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_CHECK_WITH) {
		mpc_print_unretained(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fx, 0)
		libc.Xprintf(tls, __ccgo_ts+1262, 0)
	}
}

func mpc_print(tls *libc.TLS, p uintptr) {
	mpc_print_unretained(tls, p, int32(1))
	libc.Xprintf(tls, __ccgo_ts+174, 0)
}

type __ccgo_fp__Xmpc_test_fail_3 = func(*libc.TLS, uintptr, uintptr) int32

type __ccgo_fp__Xmpc_test_fail_4 = func(*libc.TLS, uintptr)

type __ccgo_fp__Xmpc_test_fail_5 = func(*libc.TLS, uintptr)

/*
** Testing
 */

/*
** These functions are slightly unwieldy and
** also the whole of the testing suite for mpc
** mpc is pretty shaky.
**
** It could do with a lot more tests and more
** precision. Currently I am only really testing
** changes off of the examples.
**
 */
func mpc_test_fail(tls *libc.TLS, p uintptr, s uintptr, d uintptr, __ccgo_fp_tester uintptr, __ccgo_fp_destructor mpc_dtor_t, __ccgo_fp_printer uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* r at bp+0 */ mpc_result_t
	_ = __ccgo_fp_printer
	if mpc_parse(tls, __ccgo_ts+1266, s, p, bp) != 0 {
		if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_tester})))(tls, *(*uintptr)(unsafe.Pointer(bp)), d) != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_destructor})))(tls, *(*uintptr)(unsafe.Pointer(bp)))
			return 0
		} else {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_destructor})))(tls, *(*uintptr)(unsafe.Pointer(bp)))
			return int32(1)
		}
	} else {
		mpc_err_delete(tls, *(*uintptr)(unsafe.Pointer(bp)))
		return int32(1)
	}
	return r
}

type __ccgo_fp__Xmpc_test_pass_3 = func(*libc.TLS, uintptr, uintptr) int32

type __ccgo_fp__Xmpc_test_pass_4 = func(*libc.TLS, uintptr)

type __ccgo_fp__Xmpc_test_pass_5 = func(*libc.TLS, uintptr)

func mpc_test_pass(tls *libc.TLS, p uintptr, s uintptr, d uintptr, __ccgo_fp_tester uintptr, __ccgo_fp_destructor mpc_dtor_t, __ccgo_fp_printer uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* r at bp+0 */ mpc_result_t
	if mpc_parse(tls, __ccgo_ts+1266, s, p, bp) != 0 {
		if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_tester})))(tls, *(*uintptr)(unsafe.Pointer(bp)), d) != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_destructor})))(tls, *(*uintptr)(unsafe.Pointer(bp)))
			return int32(1)
		} else {
			libc.Xprintf(tls, __ccgo_ts+1273, 0)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_printer})))(tls, *(*uintptr)(unsafe.Pointer(bp)))
			libc.Xprintf(tls, __ccgo_ts+174, 0)
			libc.Xprintf(tls, __ccgo_ts+1278, 0)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_printer})))(tls, d)
			libc.Xprintf(tls, __ccgo_ts+174, 0)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_destructor})))(tls, *(*uintptr)(unsafe.Pointer(bp)))
			return 0
		}
	} else {
		mpc_err_print(tls, *(*uintptr)(unsafe.Pointer(bp)))
		mpc_err_delete(tls, *(*uintptr)(unsafe.Pointer(bp)))
		return 0
	}
	return r
}

/*
** AST
 */
func mpc_ast_delete(tls *libc.TLS, a uintptr) {
	var i int32
	_ = i
	if a == libc.UintptrFromInt32(0) {
		return
	}
	i = 0
	for {
		if !(i < (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren_num) {
			break
		}
		mpc_ast_delete(tls, **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(a)).Fchildren + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfree(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren)
	libc.Xfree(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag)
	libc.Xfree(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Fcontents)
	libc.Xfree(tls, a)
}

func mpc_ast_delete_no_children(tls *libc.TLS, a uintptr) {
	libc.Xfree(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren)
	libc.Xfree(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag)
	libc.Xfree(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Fcontents)
	libc.Xfree(tls, a)
}

func mpc_ast_new(tls *libc.TLS, tag uintptr, contents uintptr) (r uintptr) {
	var a uintptr
	_ = a
	a = libc.Xmalloc(tls, uint64(64))
	(*mpc_ast_t)(unsafe.Pointer(a)).Ftag = libc.Xmalloc(tls, libc.Xstrlen(tls, tag)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, tag)
	(*mpc_ast_t)(unsafe.Pointer(a)).Fcontents = libc.Xmalloc(tls, libc.Xstrlen(tls, contents)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Fcontents, contents)
	(*mpc_ast_t)(unsafe.Pointer(a)).Fstate = mpc_state_new(tls)
	(*mpc_ast_t)(unsafe.Pointer(a)).Fchildren_num = 0
	(*mpc_ast_t)(unsafe.Pointer(a)).Fchildren = libc.UintptrFromInt32(0)
	return a
}

func mpc_ast_build(tls *libc.TLS, n int32, tag uintptr, va1 uintptr) (r uintptr) {
	var a uintptr
	var i int32
	var va va_list
	_, _, _ = a, i, va
	a = mpc_ast_new(tls, tag, __ccgo_ts+212)
	va = va1
	i = 0
	for {
		if !(i < n) {
			break
		}
		mpc_ast_add_child(tls, a, libc.VaUintptr(&va))
		goto _1
	_1:
		;
		i = i + 1
	}
	_ = va
	return a
}

func mpc_ast_add_root(tls *libc.TLS, a uintptr) (r1 uintptr) {
	var r uintptr
	_ = r
	if a == libc.UintptrFromInt32(0) {
		return a
	}
	if (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren_num == 0 {
		return a
	}
	if (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren_num == int32(1) {
		return a
	}
	r = mpc_ast_new(tls, __ccgo_ts+988, __ccgo_ts+212)
	mpc_ast_add_child(tls, r, a)
	return r
}

func mpc_ast_eq(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	var i int32
	_ = i
	if libc.Xstrcmp(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, (*mpc_ast_t)(unsafe.Pointer(b)).Ftag) != 0 {
		return 0
	}
	if libc.Xstrcmp(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Fcontents, (*mpc_ast_t)(unsafe.Pointer(b)).Fcontents) != 0 {
		return 0
	}
	if (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren_num != (*mpc_ast_t)(unsafe.Pointer(b)).Fchildren_num {
		return 0
	}
	i = 0
	for {
		if !(i < (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren_num) {
			break
		}
		if !(mpc_ast_eq(tls, **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(a)).Fchildren + uintptr(i)*8)), **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(b)).Fchildren + uintptr(i)*8))) != 0) {
			return 0
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return int32(1)
}

func mpc_ast_add_child(tls *libc.TLS, r uintptr, a uintptr) (r1 uintptr) {
	(*mpc_ast_t)(unsafe.Pointer(r)).Fchildren_num = (*mpc_ast_t)(unsafe.Pointer(r)).Fchildren_num + 1
	(*mpc_ast_t)(unsafe.Pointer(r)).Fchildren = libc.Xrealloc(tls, (*mpc_ast_t)(unsafe.Pointer(r)).Fchildren, uint64(8)*libc.Uint64FromInt32((*mpc_ast_t)(unsafe.Pointer(r)).Fchildren_num))
	**(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(r)).Fchildren + uintptr((*mpc_ast_t)(unsafe.Pointer(r)).Fchildren_num-int32(1))*8)) = a
	return r
}

func mpc_ast_add_tag(tls *libc.TLS, a uintptr, t uintptr) (r uintptr) {
	if a == libc.UintptrFromInt32(0) {
		return a
	}
	(*mpc_ast_t)(unsafe.Pointer(a)).Ftag = libc.Xrealloc(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, libc.Xstrlen(tls, t)+uint64(1)+libc.Xstrlen(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag)+uint64(1))
	libc.Xmemmove(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag+uintptr(libc.Xstrlen(tls, t))+uintptr(1), (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, libc.Xstrlen(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag)+uint64(1))
	libc.Xmemmove(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, t, libc.Xstrlen(tls, t))
	libc.Xmemmove(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag+uintptr(libc.Xstrlen(tls, t)), __ccgo_ts+1288, uint64(1))
	return a
}

func mpc_ast_add_root_tag(tls *libc.TLS, a uintptr, t uintptr) (r uintptr) {
	if a == libc.UintptrFromInt32(0) {
		return a
	}
	(*mpc_ast_t)(unsafe.Pointer(a)).Ftag = libc.Xrealloc(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, libc.Xstrlen(tls, t)-uint64(1)+libc.Xstrlen(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag)+uint64(1))
	libc.Xmemmove(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag+uintptr(libc.Xstrlen(tls, t)-libc.Uint64FromInt32(1)), (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, libc.Xstrlen(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag)+uint64(1))
	libc.Xmemmove(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, t, libc.Xstrlen(tls, t)-libc.Uint64FromInt32(1))
	return a
}

func mpc_ast_tag(tls *libc.TLS, a uintptr, t uintptr) (r uintptr) {
	(*mpc_ast_t)(unsafe.Pointer(a)).Ftag = libc.Xrealloc(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, libc.Xstrlen(tls, t)+uint64(1))
	libc.Xstrcpy(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, t)
	return a
}

func mpc_ast_state(tls *libc.TLS, a uintptr, s mpc_state_t) (r uintptr) {
	if a == libc.UintptrFromInt32(0) {
		return a
	}
	(*mpc_ast_t)(unsafe.Pointer(a)).Fstate = s
	return a
}

func mpc_ast_print_depth(tls *libc.TLS, a uintptr, d int32, fp uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i int32
	_ = i
	if a == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, fp, __ccgo_ts+1290, 0)
		return
	}
	i = 0
	for {
		if !(i < d) {
			break
		}
		libc.Xfprintf(tls, fp, __ccgo_ts+1296, 0)
		goto _1
	_1:
		;
		i = i + 1
	}
	if libc.Xstrlen(tls, (*mpc_ast_t)(unsafe.Pointer(a)).Fcontents) != 0 {
		libc.Xfprintf(tls, fp, __ccgo_ts+1299, libc.VaList(bp+8, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag, libc.Uint64FromInt64((*mpc_ast_t)(unsafe.Pointer(a)).Fstate.Frow+libc.Int64FromInt32(1)), libc.Uint64FromInt64((*mpc_ast_t)(unsafe.Pointer(a)).Fstate.Fcol+libc.Int64FromInt32(1)), (*mpc_ast_t)(unsafe.Pointer(a)).Fcontents))
	} else {
		libc.Xfprintf(tls, fp, __ccgo_ts+1316, libc.VaList(bp+8, (*mpc_ast_t)(unsafe.Pointer(a)).Ftag))
	}
	i = 0
	for {
		if !(i < (*mpc_ast_t)(unsafe.Pointer(a)).Fchildren_num) {
			break
		}
		mpc_ast_print_depth(tls, **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(a)).Fchildren + uintptr(i)*8)), d+int32(1), fp)
		goto _2
	_2:
		;
		i = i + 1
	}
}

func mpc_ast_print(tls *libc.TLS, a uintptr) {
	mpc_ast_print_depth(tls, a, 0, libc.Xstdout)
}

func mpc_ast_print_to(tls *libc.TLS, a uintptr, fp uintptr) {
	mpc_ast_print_depth(tls, a, 0, fp)
}

func mpc_ast_get_index(tls *libc.TLS, ast uintptr, tag uintptr) (r int32) {
	return mpc_ast_get_index_lb(tls, ast, tag, 0)
}

func mpc_ast_get_index_lb(tls *libc.TLS, ast uintptr, tag uintptr, lb int32) (r int32) {
	var i int32
	_ = i
	i = lb
	for {
		if !(i < (*mpc_ast_t)(unsafe.Pointer(ast)).Fchildren_num) {
			break
		}
		if libc.Xstrcmp(tls, (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(ast)).Fchildren + uintptr(i)*8)))).Ftag, tag) == 0 {
			return i
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return -int32(1)
}

func mpc_ast_get_child(tls *libc.TLS, ast uintptr, tag uintptr) (r uintptr) {
	return mpc_ast_get_child_lb(tls, ast, tag, 0)
}

func mpc_ast_get_child_lb(tls *libc.TLS, ast uintptr, tag uintptr, lb int32) (r uintptr) {
	var i int32
	_ = i
	i = lb
	for {
		if !(i < (*mpc_ast_t)(unsafe.Pointer(ast)).Fchildren_num) {
			break
		}
		if libc.Xstrcmp(tls, (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(ast)).Fchildren + uintptr(i)*8)))).Ftag, tag) == 0 {
			return **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(ast)).Fchildren + uintptr(i)*8))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.UintptrFromInt32(0)
}

func mpc_ast_traverse_start(tls *libc.TLS, ast uintptr, order mpc_ast_trav_order_t) (r uintptr) {
	var cnode, n_trav, trav uintptr
	_, _, _ = cnode, n_trav, trav
	cnode = ast
	/* Create the traversal structure */
	trav = libc.Xmalloc(tls, uint64(24))
	(*mpc_ast_trav_t)(unsafe.Pointer(trav)).Fcurr_node = cnode
	(*mpc_ast_trav_t)(unsafe.Pointer(trav)).Fparent = libc.UintptrFromInt32(0)
	(*mpc_ast_trav_t)(unsafe.Pointer(trav)).Fcurr_child = 0
	(*mpc_ast_trav_t)(unsafe.Pointer(trav)).Forder = order
	/* Get start node */
	switch order {
	case int32(mpc_ast_trav_order_pre):
		goto _1
	case int32(mpc_ast_trav_order_post):
		goto _2
	default:
		goto _3
	}
	goto _4
_1:
	;
	/* Nothing else is needed for pre order start */
	goto _4
_2:
	;
_6:
	;
	if !((*mpc_ast_t)(unsafe.Pointer(cnode)).Fchildren_num > 0) {
		goto _5
	}
	cnode = **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(cnode)).Fchildren))
	n_trav = libc.Xmalloc(tls, uint64(24))
	(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fcurr_node = cnode
	(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fparent = trav
	(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fcurr_child = 0
	(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Forder = order
	trav = n_trav
	goto _6
_5:
	;
	goto _4
_3:
	;
	/* Unreachable, but compiler complaints */
	goto _4
_4:
	;
	return trav
}

func mpc_ast_traverse_next(tls *libc.TLS, trav uintptr) (r uintptr) {
	var cchild int32
	var n_trav, ret, to_free uintptr
	_, _, _, _ = cchild, n_trav, ret, to_free
	ret = libc.UintptrFromInt32(0)
	/* The end of traversal was reached */
	if **(**uintptr)(__ccgo_up(trav)) == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	switch (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Forder {
	case int32(mpc_ast_trav_order_pre):
		ret = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_node
		/* If there aren't any more children, go up */
		for **(**uintptr)(__ccgo_up(trav)) != libc.UintptrFromInt32(0) && (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child >= (*mpc_ast_t)(unsafe.Pointer((*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_node)).Fchildren_num {
			to_free = **(**uintptr)(__ccgo_up(trav))
			**(**uintptr)(__ccgo_up(trav)) = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fparent
			libc.Xfree(tls, to_free)
		}
		/* If trav is NULL, the end was reached */
		if **(**uintptr)(__ccgo_up(trav)) == libc.UintptrFromInt32(0) {
			break
		}
		/* Go to next child */
		n_trav = libc.Xmalloc(tls, uint64(24))
		cchild = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child
		(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fcurr_node = **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer((*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_node)).Fchildren + uintptr(cchild)*8))
		(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fparent = **(**uintptr)(__ccgo_up(trav))
		(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fcurr_child = 0
		(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Forder = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Forder
		(*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child + 1
		**(**uintptr)(__ccgo_up(trav)) = n_trav
	case int32(mpc_ast_trav_order_post):
		ret = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_node
		/* Move up tree to the parent If the parent doesn't have any more nodes,
		 * then this is the current node. If it does, move down to its left most
		 * child. Also, free the previous traversal node */
		to_free = **(**uintptr)(__ccgo_up(trav))
		**(**uintptr)(__ccgo_up(trav)) = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fparent
		libc.Xfree(tls, to_free)
		if **(**uintptr)(__ccgo_up(trav)) == libc.UintptrFromInt32(0) {
			break
		}
		/* Next child */
		(*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child + 1
		/* If there aren't any more children, this is the next node */
		if (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child >= (*mpc_ast_t)(unsafe.Pointer((*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_node)).Fchildren_num {
			break
		}
		/* If there are still more children, find the leftmost child from this
		 * node */
		for (*mpc_ast_t)(unsafe.Pointer((*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_node)).Fchildren_num > 0 {
			n_trav = libc.Xmalloc(tls, uint64(24))
			cchild = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_child
			(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fcurr_node = **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer((*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fcurr_node)).Fchildren + uintptr(cchild)*8))
			(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fparent = **(**uintptr)(__ccgo_up(trav))
			(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Fcurr_child = 0
			(*mpc_ast_trav_t)(unsafe.Pointer(n_trav)).Forder = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Forder
			**(**uintptr)(__ccgo_up(trav)) = n_trav
		}
		fallthrough
	default:
		/* Unreachable, but compiler complaints */
		break
	}
	return ret
}

func mpc_ast_traverse_free(tls *libc.TLS, trav uintptr) {
	var n_trav uintptr
	_ = n_trav
	/* Go through parents until all are free */
	for **(**uintptr)(__ccgo_up(trav)) != libc.UintptrFromInt32(0) {
		n_trav = (*mpc_ast_trav_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(trav)))).Fparent
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(trav)))
		**(**uintptr)(__ccgo_up(trav)) = n_trav
	}
}

func mpcf_fold_ast(tls *libc.TLS, n int32, xs uintptr) (r1 uintptr) {
	var as, r uintptr
	var i, j int32
	_, _, _, _ = as, i, j, r
	as = xs
	if n == 0 {
		return libc.UintptrFromInt32(0)
	}
	if n == int32(1) {
		return **(**uintptr)(__ccgo_up(xs))
	}
	if n == int32(2) && **(**uintptr)(__ccgo_up(xs + 1*8)) == libc.UintptrFromInt32(0) {
		return **(**uintptr)(__ccgo_up(xs))
	}
	if n == int32(2) && **(**uintptr)(__ccgo_up(xs)) == libc.UintptrFromInt32(0) {
		return **(**uintptr)(__ccgo_up(xs + 1*8))
	}
	r = mpc_ast_new(tls, __ccgo_ts+988, __ccgo_ts+212)
	i = 0
	for {
		if !(i < n) {
			break
		}
		if **(**uintptr)(__ccgo_up(as + uintptr(i)*8)) == libc.UintptrFromInt32(0) {
			goto _1
		}
		if **(**uintptr)(__ccgo_up(as + uintptr(i)*8)) != 0 && (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(as + uintptr(i)*8)))).Fchildren_num == 0 {
			mpc_ast_add_child(tls, r, **(**uintptr)(__ccgo_up(as + uintptr(i)*8)))
		} else {
			if **(**uintptr)(__ccgo_up(as + uintptr(i)*8)) != 0 && (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(as + uintptr(i)*8)))).Fchildren_num == int32(1) {
				mpc_ast_add_child(tls, r, mpc_ast_add_root_tag(tls, **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(as + uintptr(i)*8)))).Fchildren)), (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(as + uintptr(i)*8)))).Ftag))
				mpc_ast_delete_no_children(tls, **(**uintptr)(__ccgo_up(as + uintptr(i)*8)))
			} else {
				if **(**uintptr)(__ccgo_up(as + uintptr(i)*8)) != 0 && (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(as + uintptr(i)*8)))).Fchildren_num >= int32(2) {
					j = 0
					for {
						if !(j < (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(as + uintptr(i)*8)))).Fchildren_num) {
							break
						}
						mpc_ast_add_child(tls, r, **(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up(as + uintptr(i)*8)))).Fchildren + uintptr(j)*8)))
						goto _2
					_2:
						;
						j = j + 1
					}
					mpc_ast_delete_no_children(tls, **(**uintptr)(__ccgo_up(as + uintptr(i)*8)))
				}
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if (*mpc_ast_t)(unsafe.Pointer(r)).Fchildren_num != 0 {
		(*mpc_ast_t)(unsafe.Pointer(r)).Fstate = (*mpc_ast_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*mpc_ast_t)(unsafe.Pointer(r)).Fchildren)))).Fstate
	}
	return r
}

func mpcf_str_ast(tls *libc.TLS, c uintptr) (r uintptr) {
	var a uintptr
	_ = a
	a = mpc_ast_new(tls, __ccgo_ts+212, c)
	libc.Xfree(tls, c)
	return a
}

func mpcf_state_ast(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var a, s uintptr
	_, _ = a, s
	s = **(**uintptr)(__ccgo_up(xs))
	a = **(**uintptr)(__ccgo_up(xs + 1*8))
	_ = n
	a = mpc_ast_state(tls, a, **(**mpc_state_t)(__ccgo_up(s)))
	libc.Xfree(tls, s)
	return a
}

func mpca_state(tls *libc.TLS, a uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return mpc_and(tls, int32(2), __ccgo_fp(mpcf_state_ast), libc.VaList(bp+8, mpc_state(tls), a, __ccgo_fp(libc.Xfree)))
}

func mpca_tag(tls *libc.TLS, a uintptr, t uintptr) (r uintptr) {
	return mpc_apply_to(tls, a, __ccgo_fp(mpc_ast_tag), t)
}

func mpca_add_tag(tls *libc.TLS, a uintptr, t uintptr) (r uintptr) {
	return mpc_apply_to(tls, a, __ccgo_fp(mpc_ast_add_tag), t)
}

func mpca_root(tls *libc.TLS, a uintptr) (r uintptr) {
	return mpc_apply(tls, a, __ccgo_fp(mpc_ast_add_root))
}

func mpca_not(tls *libc.TLS, a uintptr) (r uintptr) {
	return mpc_not(tls, a, __ccgo_fp(mpc_ast_delete))
}

func mpca_maybe(tls *libc.TLS, a uintptr) (r uintptr) {
	return mpc_maybe(tls, a)
}

func mpca_many(tls *libc.TLS, a uintptr) (r uintptr) {
	return mpc_many(tls, __ccgo_fp(mpcf_fold_ast), a)
}

func mpca_many1(tls *libc.TLS, a uintptr) (r uintptr) {
	return mpc_many1(tls, __ccgo_fp(mpcf_fold_ast), a)
}

func mpca_count(tls *libc.TLS, n int32, a uintptr) (r uintptr) {
	return mpc_count(tls, n, __ccgo_fp(mpcf_fold_ast), a, __ccgo_fp(mpc_ast_delete))
}

func mpca_or(tls *libc.TLS, n int32, va1 uintptr) (r uintptr) {
	var i int32
	var p uintptr
	var va va_list
	_, _, _ = i, p, va
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_OR)
	(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n
	(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n))
	va = va1
	i = 0
	for {
		if !(i < n) {
			break
		}
		**(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)) = libc.VaUintptr(&va)
		goto _1
	_1:
		;
		i = i + 1
	}
	_ = va
	return p
}

func mpca_and(tls *libc.TLS, n int32, va1 uintptr) (r uintptr) {
	var i int32
	var p uintptr
	var va va_list
	_, _, _ = i, p, va
	p = mpc_undefined(tls)
	(*mpc_parser_t)(unsafe.Pointer(p)).Ftype1 = int8(MPC_TYPE_AND)
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff = __ccgo_fp(mpcf_fold_ast)
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n))
	(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n-libc.Int32FromInt32(1)))
	va = va1
	i = 0
	for {
		if !(i < n) {
			break
		}
		**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)) = libc.VaUintptr(&va)
		goto _1
	_1:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < n-int32(1)) {
			break
		}
		**(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(i)*8)) = __ccgo_fp(mpc_ast_delete)
		goto _2
	_2:
		;
		i = i + 1
	}
	_ = va
	return p
}

func mpca_total(tls *libc.TLS, a uintptr) (r uintptr) {
	return mpc_total(tls, a, __ccgo_fp(mpc_ast_delete))
}

/*
** Grammar Parser
 */

/*
** This is another interesting bootstrapping.
**
** Having a general purpose AST type allows
** users to specify the grammar alone and
** let all fold rules be automatically taken
** care of by existing functions.
**
** You don't get to control the type spat
** out but this means you can make a nice
** parser to take in some grammar in nice
** syntax and spit out a parser that works.
**
** The grammar for this looks surprisingly
** like regex but the main difference is that
** it is now whitespace insensitive and the
** base type takes literals of some form.
 */

/*
**
**  ### Grammar Grammar
**
**      <grammar> : (<term> "|" <grammar>) | <term>
**
**      <term> : <factor>*
**
**      <factor> : <base>
**               | <base> "*"
**               | <base> "+"
**               | <base> "?"
**               | <base> "{" <digits> "}"
**
**      <base> : "<" (<digits> | <ident>) ">"
**             | <string_lit>
**             | <char_lit>
**             | <regex_lit> <regex_mode>
**             | "(" <grammar> ")"
 */
type mpca_grammar_st_t = struct {
	Fva          uintptr
	Fparsers_num int32
	Fparsers     uintptr
	Fflags       int32
}

func mpcaf_grammar_or(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	_ = n
	if **(**uintptr)(__ccgo_up(xs + 1*8)) == libc.UintptrFromInt32(0) {
		return **(**uintptr)(__ccgo_up(xs))
	} else {
		return mpca_or(tls, int32(2), libc.VaList(bp+8, **(**uintptr)(__ccgo_up(xs)), **(**uintptr)(__ccgo_up(xs + 1*8))))
	}
	return r
}

func mpcaf_grammar_and(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var p uintptr
	_, _ = i, p
	p = mpc_pass(tls)
	i = 0
	for {
		if !(i < n) {
			break
		}
		if **(**uintptr)(__ccgo_up(xs + uintptr(i)*8)) != libc.UintptrFromInt32(0) {
			p = mpca_and(tls, int32(2), libc.VaList(bp+8, p, **(**uintptr)(__ccgo_up(xs + uintptr(i)*8))))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return p
}

func mpcaf_grammar_repeat(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var num int32
	_ = num
	_ = n
	if **(**uintptr)(__ccgo_up(xs + 1*8)) == libc.UintptrFromInt32(0) {
		return **(**uintptr)(__ccgo_up(xs))
	}
	switch int32(**(**int8)(__ccgo_up(**(**uintptr)(__ccgo_up(xs + 1*8))))) {
	case int32('*'):
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
		return mpca_many(tls, **(**uintptr)(__ccgo_up(xs)))
	case int32('+'):
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
		return mpca_many1(tls, **(**uintptr)(__ccgo_up(xs)))
	case int32('?'):
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
		return mpca_maybe(tls, **(**uintptr)(__ccgo_up(xs)))
	case int32('!'):
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
		return mpca_not(tls, **(**uintptr)(__ccgo_up(xs)))
	default:
		num = **(**int32)(__ccgo_up(**(**uintptr)(__ccgo_up(xs + 1*8))))
		libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 1*8)))
	}
	return mpca_count(tls, num, **(**uintptr)(__ccgo_up(xs)))
}

func mpcaf_grammar_string(tls *libc.TLS, x uintptr, s uintptr) (r uintptr) {
	var p, st, y, v1 uintptr
	_, _, _, _ = p, st, y, v1
	st = s
	y = mpcf_unescape(tls, x)
	if (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fflags&int32(MPCA_LANG_WHITESPACE_SENSITIVE) != 0 {
		v1 = mpc_string(tls, y)
	} else {
		v1 = mpc_tok(tls, mpc_string(tls, y))
	}
	p = v1
	libc.Xfree(tls, y)
	return mpca_state(tls, mpca_tag(tls, mpc_apply(tls, p, __ccgo_fp(mpcf_str_ast)), __ccgo_ts+967))
}

func mpcaf_grammar_char(tls *libc.TLS, x uintptr, s uintptr) (r uintptr) {
	var p, st, y, v1 uintptr
	_, _, _, _ = p, st, y, v1
	st = s
	y = mpcf_unescape(tls, x)
	if (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fflags&int32(MPCA_LANG_WHITESPACE_SENSITIVE) != 0 {
		v1 = mpc_char(tls, **(**int8)(__ccgo_up(y)))
	} else {
		v1 = mpc_tok(tls, mpc_char(tls, **(**int8)(__ccgo_up(y))))
	}
	p = v1
	libc.Xfree(tls, y)
	return mpca_state(tls, mpca_tag(tls, mpc_apply(tls, p, __ccgo_fp(mpcf_str_ast)), __ccgo_ts+960))
}

func mpcaf_fold_regex(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var m, p, st, y, v1 uintptr
	var mode int32
	_, _, _, _, _, _ = m, mode, p, st, y, v1
	y = **(**uintptr)(__ccgo_up(xs))
	m = **(**uintptr)(__ccgo_up(xs + 1*8))
	st = **(**uintptr)(__ccgo_up(xs + 2*8))
	mode = int32(MPC_RE_DEFAULT)
	_ = n
	if libc.Xstrchr(tls, m, int32('m')) != 0 {
		mode = mode | int32(MPC_RE_MULTILINE)
	}
	if libc.Xstrchr(tls, m, int32('s')) != 0 {
		mode = mode | int32(MPC_RE_DOTALL)
	}
	y = mpcf_unescape_regex(tls, y)
	if (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fflags&int32(MPCA_LANG_WHITESPACE_SENSITIVE) != 0 {
		v1 = mpc_re_mode(tls, y, mode)
	} else {
		v1 = mpc_tok(tls, mpc_re_mode(tls, y, mode))
	}
	p = v1
	libc.Xfree(tls, y)
	libc.Xfree(tls, m)
	return mpca_state(tls, mpca_tag(tls, mpc_apply(tls, p, __ccgo_fp(mpcf_str_ast)), __ccgo_ts+976))
}

// C documentation
//
//	/* Should this just use `isdigit` instead? */
func is_number(tls *libc.TLS, s uintptr) (r int32) {
	var i size_t
	_ = i
	i = uint64(0)
	for {
		if !(i < libc.Xstrlen(tls, s)) {
			break
		}
		if !(libc.Xstrchr(tls, __ccgo_ts+632, int32(**(**int8)(__ccgo_up(s + uintptr(i))))) != 0) {
			return 0
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return int32(1)
}

func mpca_grammar_find_parser(tls *libc.TLS, x uintptr, st uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var p, q uintptr
	_, _, _ = i, p, q
	/* Case of Number */
	if is_number(tls, x) != 0 {
		i = int32(libc.Xstrtol(tls, x, libc.UintptrFromInt32(0), int32(10)))
		for (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num <= i {
			(*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num = (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num + 1
			(*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers = libc.Xrealloc(tls, (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers, uint64(8)*libc.Uint64FromInt32((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num))
			**(**uintptr)(__ccgo_up((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers + uintptr((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num-int32(1))*8)) = libc.VaUintptr(&**(**va_list)(__ccgo_up((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fva)))
			if **(**uintptr)(__ccgo_up((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers + uintptr((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num-int32(1))*8)) == libc.UintptrFromInt32(0) {
				return mpc_failf(tls, __ccgo_ts+1321, libc.VaList(bp+8, i, (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num))
			}
		}
		return **(**uintptr)(__ccgo_up((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers + uintptr((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num-int32(1))*8))
		/* Case of Identifier */
	} else {
		/* Search Existing Parsers */
		i = 0
		for {
			if !(i < (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num) {
				break
			}
			q = **(**uintptr)(__ccgo_up((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers + uintptr(i)*8))
			if q == libc.UintptrFromInt32(0) {
				return mpc_failf(tls, __ccgo_ts+1373, libc.VaList(bp+8, x))
			}
			if (*mpc_parser_t)(unsafe.Pointer(q)).Fname != 0 && libc.Xstrcmp(tls, (*mpc_parser_t)(unsafe.Pointer(q)).Fname, x) == 0 {
				return q
			}
			goto _1
		_1:
			;
			i = i + 1
		}
		/* Search New Parsers */
		for int32(1) != 0 {
			p = libc.VaUintptr(&**(**va_list)(__ccgo_up((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fva)))
			(*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num = (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num + 1
			(*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers = libc.Xrealloc(tls, (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers, uint64(8)*libc.Uint64FromInt32((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num))
			**(**uintptr)(__ccgo_up((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers + uintptr((*mpca_grammar_st_t)(unsafe.Pointer(st)).Fparsers_num-int32(1))*8)) = p
			if p == libc.UintptrFromInt32(0) || (*mpc_parser_t)(unsafe.Pointer(p)).Fname == libc.UintptrFromInt32(0) {
				return mpc_failf(tls, __ccgo_ts+1373, libc.VaList(bp+8, x))
			}
			if (*mpc_parser_t)(unsafe.Pointer(p)).Fname != 0 && libc.Xstrcmp(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname, x) == 0 {
				return p
			}
		}
	}
	return r
}

func mpcaf_grammar_id(tls *libc.TLS, x uintptr, s uintptr) (r uintptr) {
	var p, st uintptr
	_, _ = p, st
	st = s
	p = mpca_grammar_find_parser(tls, x, st)
	libc.Xfree(tls, x)
	if (*mpc_parser_t)(unsafe.Pointer(p)).Fname != 0 {
		return mpca_state(tls, mpca_root(tls, mpca_add_tag(tls, p, (*mpc_parser_t)(unsafe.Pointer(p)).Fname)))
	} else {
		return mpca_state(tls, mpca_root(tls, p))
	}
	return r
}

func mpca_grammar_st(tls *libc.TLS, grammar uintptr, st uintptr) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var Base, Factor, Grammar, GrammarTotal, Term, err_msg, err_out, v1 uintptr
	var _ /* r at bp+0 */ mpc_result_t
	_, _, _, _, _, _, _, _ = Base, Factor, Grammar, GrammarTotal, Term, err_msg, err_out, v1
	GrammarTotal = mpc_new(tls, __ccgo_ts+1394)
	Grammar = mpc_new(tls, __ccgo_ts+1408)
	Term = mpc_new(tls, __ccgo_ts+1074)
	Factor = mpc_new(tls, __ccgo_ts+1079)
	Base = mpc_new(tls, __ccgo_ts+1086)
	mpc_define(tls, GrammarTotal, mpc_predictive(tls, mpc_total(tls, Grammar, __ccgo_fp(mpc_soft_delete))))
	mpc_define(tls, Grammar, mpc_and(tls, int32(2), __ccgo_fp(mpcaf_grammar_or), libc.VaList(bp+16, Term, mpc_maybe(tls, mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd_free), libc.VaList(bp+16, mpc_sym(tls, __ccgo_ts+1288), Grammar, __ccgo_fp(libc.Xfree)))), __ccgo_fp(mpc_soft_delete))))
	mpc_define(tls, Term, mpc_many1(tls, __ccgo_fp(mpcaf_grammar_and), Factor))
	mpc_define(tls, Factor, mpc_and(tls, int32(2), __ccgo_fp(mpcaf_grammar_repeat), libc.VaList(bp+16, Base, mpc_or(tls, int32(6), libc.VaList(bp+16, mpc_sym(tls, __ccgo_ts+1244), mpc_sym(tls, __ccgo_ts+1246), mpc_sym(tls, __ccgo_ts+1242), mpc_sym(tls, __ccgo_ts+1240), mpc_tok_brackets(tls, mpc_int(tls), __ccgo_fp(libc.Xfree)), mpc_pass(tls))), __ccgo_fp(mpc_soft_delete))))
	mpc_define(tls, Base, mpc_or(tls, int32(5), libc.VaList(bp+16, mpc_apply_to(tls, mpc_tok(tls, mpc_string_lit(tls)), __ccgo_fp(mpcaf_grammar_string), st), mpc_apply_to(tls, mpc_tok(tls, mpc_char_lit(tls)), __ccgo_fp(mpcaf_grammar_char), st), mpc_tok(tls, mpc_and(tls, int32(3), __ccgo_fp(mpcaf_fold_regex), libc.VaList(bp+16, mpc_regex_lit(tls), mpc_many(tls, __ccgo_fp(mpcf_strfold), mpc_oneof(tls, __ccgo_ts+1416)), mpc_lift_val(tls, st), __ccgo_fp(libc.Xfree), __ccgo_fp(libc.Xfree)))), mpc_apply_to(tls, mpc_tok_braces(tls, mpc_or(tls, int32(2), libc.VaList(bp+16, mpc_digits(tls), mpc_ident(tls))), __ccgo_fp(libc.Xfree)), __ccgo_fp(mpcaf_grammar_id), st), mpc_tok_parens(tls, Grammar, __ccgo_fp(mpc_soft_delete)))))
	mpc_optimise(tls, GrammarTotal)
	mpc_optimise(tls, Grammar)
	mpc_optimise(tls, Factor)
	mpc_optimise(tls, Term)
	mpc_optimise(tls, Base)
	if !(mpc_parse(tls, __ccgo_ts+1419, grammar, GrammarTotal, bp) != 0) {
		err_msg = mpc_err_string(tls, *(*uintptr)(unsafe.Pointer(bp)))
		err_out = mpc_failf(tls, __ccgo_ts+1442, libc.VaList(bp+16, err_msg))
		mpc_err_delete(tls, *(*uintptr)(unsafe.Pointer(bp)))
		libc.Xfree(tls, err_msg)
		*(*uintptr)(unsafe.Pointer(bp)) = err_out
	}
	mpc_cleanup(tls, int32(5), libc.VaList(bp+16, GrammarTotal, Grammar, Term, Factor, Base))
	mpc_optimise(tls, *(*uintptr)(unsafe.Pointer(bp)))
	if (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fflags&int32(MPCA_LANG_PREDICTIVE) != 0 {
		v1 = mpc_predictive(tls, *(*uintptr)(unsafe.Pointer(bp)))
	} else {
		v1 = *(*uintptr)(unsafe.Pointer(bp))
	}
	return v1
}

func mpca_grammar(tls *libc.TLS, flags int32, grammar uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var res uintptr
	var _ /* st at bp+0 */ mpca_grammar_st_t
	var _ /* va at bp+32 */ va_list
	_ = res
	**(**va_list)(__ccgo_up(bp + 32)) = va
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fva = bp + 32
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers_num = 0
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers = libc.UintptrFromInt32(0)
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fflags = flags
	res = mpca_grammar_st(tls, grammar, bp)
	libc.Xfree(tls, (**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers)
	_ = **(**va_list)(__ccgo_up(bp + 32))
	return res
}

type mpca_stmt_t = struct {
	Fident   uintptr
	Fname    uintptr
	Fgrammar uintptr
}

func mpca_stmt_afold(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var stmt uintptr
	_ = stmt
	stmt = libc.Xmalloc(tls, uint64(24))
	(*mpca_stmt_t)(unsafe.Pointer(stmt)).Fident = **(**uintptr)(__ccgo_up(xs))
	(*mpca_stmt_t)(unsafe.Pointer(stmt)).Fname = **(**uintptr)(__ccgo_up(xs + 1*8))
	(*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar = **(**uintptr)(__ccgo_up(xs + 3*8))
	_ = n
	libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 2*8)))
	libc.Xfree(tls, **(**uintptr)(__ccgo_up(xs + 4*8)))
	return stmt
}

func mpca_stmt_fold(tls *libc.TLS, n int32, xs uintptr) (r uintptr) {
	var i int32
	var stmts uintptr
	_, _ = i, stmts
	stmts = libc.Xmalloc(tls, uint64(8)*libc.Uint64FromInt32(n+libc.Int32FromInt32(1)))
	i = 0
	for {
		if !(i < n) {
			break
		}
		**(**uintptr)(__ccgo_up(stmts + uintptr(i)*8)) = **(**uintptr)(__ccgo_up(xs + uintptr(i)*8))
		goto _1
	_1:
		;
		i = i + 1
	}
	**(**uintptr)(__ccgo_up(stmts + uintptr(n)*8)) = libc.UintptrFromInt32(0)
	return stmts
}

func mpca_stmt_list_delete(tls *libc.TLS, x uintptr) {
	var stmt, stmts uintptr
	_, _ = stmt, stmts
	stmts = x
	for **(**uintptr)(__ccgo_up(stmts)) != 0 {
		stmt = **(**uintptr)(__ccgo_up(stmts))
		libc.Xfree(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fident)
		libc.Xfree(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fname)
		mpc_soft_delete(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar)
		libc.Xfree(tls, stmt)
		stmts += 8
	}
	libc.Xfree(tls, x)
}

func mpca_stmt_list_apply_to(tls *libc.TLS, x uintptr, s uintptr) (r uintptr) {
	var left, st, stmt, stmts uintptr
	_, _, _, _ = left, st, stmt, stmts
	st = s
	stmts = x
	for **(**uintptr)(__ccgo_up(stmts)) != 0 {
		stmt = **(**uintptr)(__ccgo_up(stmts))
		left = mpca_grammar_find_parser(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fident, st)
		if (*mpca_grammar_st_t)(unsafe.Pointer(st)).Fflags&int32(MPCA_LANG_PREDICTIVE) != 0 {
			(*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar = mpc_predictive(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar)
		}
		if (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fname != 0 {
			(*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar = mpc_expect(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fname)
		}
		mpc_optimise(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar)
		mpc_define(tls, left, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fgrammar)
		libc.Xfree(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fident)
		libc.Xfree(tls, (*mpca_stmt_t)(unsafe.Pointer(stmt)).Fname)
		libc.Xfree(tls, stmt)
		stmts += 8
	}
	libc.Xfree(tls, x)
	return libc.UintptrFromInt32(0)
}

func mpca_lang_st(tls *libc.TLS, i uintptr, st uintptr) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var Base, Factor, Grammar, Lang, Stmt, Term, e uintptr
	var _ /* r at bp+0 */ mpc_result_t
	_, _, _, _, _, _, _ = Base, Factor, Grammar, Lang, Stmt, Term, e
	Lang = mpc_new(tls, __ccgo_ts+1462)
	Stmt = mpc_new(tls, __ccgo_ts+1467)
	Grammar = mpc_new(tls, __ccgo_ts+1408)
	Term = mpc_new(tls, __ccgo_ts+1074)
	Factor = mpc_new(tls, __ccgo_ts+1079)
	Base = mpc_new(tls, __ccgo_ts+1086)
	mpc_define(tls, Lang, mpc_apply_to(tls, mpc_total(tls, mpc_predictive(tls, mpc_many(tls, __ccgo_fp(mpca_stmt_fold), Stmt)), __ccgo_fp(mpca_stmt_list_delete)), __ccgo_fp(mpca_stmt_list_apply_to), st))
	mpc_define(tls, Stmt, mpc_and(tls, int32(5), __ccgo_fp(mpca_stmt_afold), libc.VaList(bp+16, mpc_tok(tls, mpc_ident(tls)), mpc_maybe(tls, mpc_tok(tls, mpc_string_lit(tls))), mpc_sym(tls, __ccgo_ts+1472), Grammar, mpc_sym(tls, __ccgo_ts+1474), __ccgo_fp(libc.Xfree), __ccgo_fp(libc.Xfree), __ccgo_fp(libc.Xfree), __ccgo_fp(mpc_soft_delete))))
	mpc_define(tls, Grammar, mpc_and(tls, int32(2), __ccgo_fp(mpcaf_grammar_or), libc.VaList(bp+16, Term, mpc_maybe(tls, mpc_and(tls, int32(2), __ccgo_fp(mpcf_snd_free), libc.VaList(bp+16, mpc_sym(tls, __ccgo_ts+1288), Grammar, __ccgo_fp(libc.Xfree)))), __ccgo_fp(mpc_soft_delete))))
	mpc_define(tls, Term, mpc_many1(tls, __ccgo_fp(mpcaf_grammar_and), Factor))
	mpc_define(tls, Factor, mpc_and(tls, int32(2), __ccgo_fp(mpcaf_grammar_repeat), libc.VaList(bp+16, Base, mpc_or(tls, int32(6), libc.VaList(bp+16, mpc_sym(tls, __ccgo_ts+1244), mpc_sym(tls, __ccgo_ts+1246), mpc_sym(tls, __ccgo_ts+1242), mpc_sym(tls, __ccgo_ts+1240), mpc_tok_brackets(tls, mpc_int(tls), __ccgo_fp(libc.Xfree)), mpc_pass(tls))), __ccgo_fp(mpc_soft_delete))))
	mpc_define(tls, Base, mpc_or(tls, int32(5), libc.VaList(bp+16, mpc_apply_to(tls, mpc_tok(tls, mpc_string_lit(tls)), __ccgo_fp(mpcaf_grammar_string), st), mpc_apply_to(tls, mpc_tok(tls, mpc_char_lit(tls)), __ccgo_fp(mpcaf_grammar_char), st), mpc_tok(tls, mpc_and(tls, int32(3), __ccgo_fp(mpcaf_fold_regex), libc.VaList(bp+16, mpc_regex_lit(tls), mpc_many(tls, __ccgo_fp(mpcf_strfold), mpc_oneof(tls, __ccgo_ts+1416)), mpc_lift_val(tls, st), __ccgo_fp(libc.Xfree), __ccgo_fp(libc.Xfree)))), mpc_apply_to(tls, mpc_tok_braces(tls, mpc_or(tls, int32(2), libc.VaList(bp+16, mpc_digits(tls), mpc_ident(tls))), __ccgo_fp(libc.Xfree)), __ccgo_fp(mpcaf_grammar_id), st), mpc_tok_parens(tls, Grammar, __ccgo_fp(mpc_soft_delete)))))
	mpc_optimise(tls, Lang)
	mpc_optimise(tls, Stmt)
	mpc_optimise(tls, Grammar)
	mpc_optimise(tls, Term)
	mpc_optimise(tls, Factor)
	mpc_optimise(tls, Base)
	if !(mpc_parse_input(tls, i, Lang, bp) != 0) {
		e = *(*uintptr)(unsafe.Pointer(bp))
	} else {
		e = libc.UintptrFromInt32(0)
	}
	mpc_cleanup(tls, int32(6), libc.VaList(bp+16, Lang, Stmt, Grammar, Term, Factor, Base))
	return e
}

func mpca_lang_file(tls *libc.TLS, flags int32, f uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var err, i uintptr
	var _ /* st at bp+0 */ mpca_grammar_st_t
	var _ /* va at bp+32 */ va_list
	_, _ = err, i
	**(**va_list)(__ccgo_up(bp + 32)) = va
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fva = bp + 32
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers_num = 0
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers = libc.UintptrFromInt32(0)
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fflags = flags
	i = mpc_input_new_file(tls, __ccgo_ts+1476, f)
	err = mpca_lang_st(tls, i, bp)
	mpc_input_delete(tls, i)
	libc.Xfree(tls, (**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers)
	_ = **(**va_list)(__ccgo_up(bp + 32))
	return err
}

func mpca_lang_pipe(tls *libc.TLS, flags int32, p uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var err, i uintptr
	var _ /* st at bp+0 */ mpca_grammar_st_t
	var _ /* va at bp+32 */ va_list
	_, _ = err, i
	**(**va_list)(__ccgo_up(bp + 32)) = va
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fva = bp + 32
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers_num = 0
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers = libc.UintptrFromInt32(0)
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fflags = flags
	i = mpc_input_new_pipe(tls, __ccgo_ts+1493, p)
	err = mpca_lang_st(tls, i, bp)
	mpc_input_delete(tls, i)
	libc.Xfree(tls, (**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers)
	_ = **(**va_list)(__ccgo_up(bp + 32))
	return err
}

func mpca_lang(tls *libc.TLS, flags int32, language uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var err, i uintptr
	var _ /* st at bp+0 */ mpca_grammar_st_t
	var _ /* va at bp+32 */ va_list
	_, _ = err, i
	**(**va_list)(__ccgo_up(bp + 32)) = va
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fva = bp + 32
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers_num = 0
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers = libc.UintptrFromInt32(0)
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fflags = flags
	i = mpc_input_new_string(tls, __ccgo_ts+1510, language)
	err = mpca_lang_st(tls, i, bp)
	mpc_input_delete(tls, i)
	libc.Xfree(tls, (**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers)
	_ = **(**va_list)(__ccgo_up(bp + 32))
	return err
}

func mpca_lang_contents(tls *libc.TLS, flags int32, filename uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var err, f, i uintptr
	var _ /* st at bp+0 */ mpca_grammar_st_t
	var _ /* va at bp+32 */ va_list
	_, _, _ = err, f, i
	f = libc.Xfopen(tls, filename, __ccgo_ts+312)
	if f == libc.UintptrFromInt32(0) {
		err = mpc_err_file(tls, filename, __ccgo_ts+315)
		return err
	}
	**(**va_list)(__ccgo_up(bp + 32)) = va
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fva = bp + 32
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers_num = 0
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers = libc.UintptrFromInt32(0)
	(**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fflags = flags
	i = mpc_input_new_file(tls, filename, f)
	err = mpca_lang_st(tls, i, bp)
	mpc_input_delete(tls, i)
	libc.Xfree(tls, (**(**mpca_grammar_st_t)(__ccgo_up(bp))).Fparsers)
	_ = **(**va_list)(__ccgo_up(bp + 32))
	libc.Xfclose(tls, f)
	return err
}

func mpc_nodecount_unretained(tls *libc.TLS, p uintptr, force int32) (r int32) {
	var i, total int32
	_, _ = i, total
	if (*mpc_parser_t)(unsafe.Pointer(p)).Fretained != 0 && !(force != 0) {
		return 0
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_EXPECT) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_APPLY) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_APPLY_TO) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_PREDICT) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_CHECK) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_CHECK_WITH) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_NOT) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MAYBE) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MANY) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MANY1) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_COUNT) {
		return int32(1) + mpc_nodecount_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_SEPBY1) {
		total = int32(1)
		total = total + mpc_nodecount_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		total = total + mpc_nodecount_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fsep, 0)
		total = total + mpc_nodecount_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		return total
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_OR) {
		total = int32(1)
		i = 0
		for {
			if !(i < (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
				break
			}
			total = total + mpc_nodecount_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
			goto _1
		_1:
			;
			i = i + 1
		}
		return total
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) {
		total = int32(1)
		i = 0
		for {
			if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
				break
			}
			total = total + mpc_nodecount_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
			goto _2
		_2:
			;
			i = i + 1
		}
		return total
	}
	return int32(1)
}

func mpc_stats(tls *libc.TLS, p uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libc.Xprintf(tls, __ccgo_ts+1522, 0)
	libc.Xprintf(tls, __ccgo_ts+1529, 0)
	libc.Xprintf(tls, __ccgo_ts+1536, libc.VaList(bp+8, mpc_nodecount_unretained(tls, p, int32(1))))
}

func mpc_optimise_unretained(tls *libc.TLS, p uintptr, force int32) {
	var i, m, n int32
	var t uintptr
	_, _, _, _ = i, m, n, t
	if (*mpc_parser_t)(unsafe.Pointer(p)).Fretained != 0 && !(force != 0) {
		return
	}
	/* Optimise Subexpressions */
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_EXPECT) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_expect_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_APPLY) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_apply_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_APPLY_TO) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_apply_to_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_CHECK) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_check_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_CHECK_WITH) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_check_with_t)(unsafe.Pointer(p + 8))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_PREDICT) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_predict_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_NOT) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MAYBE) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_not_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MANY) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_MANY1) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_COUNT) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_repeat_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_SEPBY1) {
		mpc_optimise_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fx, 0)
		mpc_optimise_unretained(tls, (*(*mpc_pdata_sepby1)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fsep, 0)
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_OR) {
		i = 0
		for {
			if !(i < (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
				break
			}
			mpc_optimise_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) {
		i = 0
		for {
			if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn) {
				break
			}
			mpc_optimise_unretained(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr(i)*8)), 0)
			goto _2
		_2:
			;
			i = i + 1
		}
	}
	/* Perform optimisations */
	for int32(1) != 0 {
		/* Merge rhs `or` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_OR) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Ftype1) == int32(MPC_TYPE_OR) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Fretained != 0) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8))
			n = (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn
			m = (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fn
			(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n + m - int32(1)
			(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xrealloc(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)))
			libc.Xmemmove(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(n)*8-uintptr(1)*8, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs, libc.Uint64FromInt32(m)*uint64(8))
			libc.Xfree(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(t)).Fname)
			libc.Xfree(tls, t)
			continue
		}
		/* Merge lhs `or` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_OR) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Ftype1) == int32(MPC_TYPE_OR) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fretained != 0) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs))
			n = (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn
			m = (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fn
			(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n + m - int32(1)
			(*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xrealloc(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)))
			libc.Xmemmove(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(m)*8, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(1)*8, libc.Uint64FromInt32(n-libc.Int32FromInt32(1))*uint64(8))
			libc.Xmemmove(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs, libc.Uint64FromInt32(m)*uint64(8))
			libc.Xfree(tls, (*(*mpc_pdata_or_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(t)).Fname)
			libc.Xfree(tls, t)
			continue
		}
		/* Remove ast `pass` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn == int32(2) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Ftype1) == int32(MPC_TYPE_PASS) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fretained != 0) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff == __ccgo_fp(mpcf_fold_ast) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + 1*8))
			mpc_delete(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname)
			libc.Xmemcpy(tls, p, t, uint64(56))
			libc.Xfree(tls, t)
			continue
		}
		/* Merge ast lhs `and` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff == __ccgo_fp(mpcf_fold_ast) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Ftype1) == int32(MPC_TYPE_AND) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fretained != 0) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fdata))).Ff == __ccgo_fp(mpcf_fold_ast) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs))
			n = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn
			m = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fn
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n + m - int32(1)
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)))
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)-libc.Int32FromInt32(1)))
			libc.Xmemmove(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(m)*8, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(1)*8, libc.Uint64FromInt32(n-libc.Int32FromInt32(1))*uint64(8))
			libc.Xmemmove(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs, libc.Uint64FromInt32(m)*uint64(8))
			i = 0
			for {
				if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1)) {
					break
				}
				**(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(i)*8)) = __ccgo_fp(mpc_ast_delete)
				goto _3
			_3:
				;
				i = i + 1
			}
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs)
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fdxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(t)).Fname)
			libc.Xfree(tls, t)
			continue
		}
		/* Merge ast rhs `and` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff == __ccgo_fp(mpcf_fold_ast) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Ftype1) == int32(MPC_TYPE_AND) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Fretained != 0) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Fdata))).Ff == __ccgo_fp(mpcf_fold_ast) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8))
			n = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn
			m = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fn
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n + m - int32(1)
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)))
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)-libc.Int32FromInt32(1)))
			libc.Xmemmove(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(n)*8-uintptr(1)*8, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs, libc.Uint64FromInt32(m)*uint64(8))
			i = 0
			for {
				if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1)) {
					break
				}
				**(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(i)*8)) = __ccgo_fp(mpc_ast_delete)
				goto _4
			_4:
				;
				i = i + 1
			}
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs)
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fdxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(t)).Fname)
			libc.Xfree(tls, t)
			continue
		}
		/* Remove re `lift` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn == int32(2) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Ftype1) == int32(MPC_TYPE_LIFT) && (*(*mpc_pdata_lift_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fdata))).Flf == __ccgo_fp(mpcf_ctor_str) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fretained != 0) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff == __ccgo_fp(mpcf_strfold) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + 1*8))
			mpc_delete(tls, **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(p)).Fname)
			libc.Xmemcpy(tls, p, t, uint64(56))
			libc.Xfree(tls, t)
			continue
		}
		/* Merge re lhs `and` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff == __ccgo_fp(mpcf_strfold) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Ftype1) == int32(MPC_TYPE_AND) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fretained != 0) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs)))).Fdata))).Ff == __ccgo_fp(mpcf_strfold) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs))
			n = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn
			m = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fn
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n + m - int32(1)
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)))
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)-libc.Int32FromInt32(1)))
			libc.Xmemmove(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(m)*8, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(1)*8, libc.Uint64FromInt32(n-libc.Int32FromInt32(1))*uint64(8))
			libc.Xmemmove(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs, libc.Uint64FromInt32(m)*uint64(8))
			i = 0
			for {
				if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1)) {
					break
				}
				**(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(i)*8)) = __ccgo_fp(libc.Xfree)
				goto _5
			_5:
				;
				i = i + 1
			}
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs)
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fdxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(t)).Fname)
			libc.Xfree(tls, t)
			continue
		}
		/* Merge re rhs `and` */
		if int32((*mpc_parser_t)(unsafe.Pointer(p)).Ftype1) == int32(MPC_TYPE_AND) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Ff == __ccgo_fp(mpcf_strfold) && int32((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Ftype1) == int32(MPC_TYPE_AND) && !((*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Fretained != 0) && (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(**(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8)))).Fdata))).Ff == __ccgo_fp(mpcf_strfold) {
			t = **(**uintptr)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs + uintptr((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1))*8))
			n = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn
			m = (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fn
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn = n + m - int32(1)
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)))
			(*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs = libc.Xrealloc(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs, uint64(8)*libc.Uint64FromInt32(n+m-libc.Int32FromInt32(1)-libc.Int32FromInt32(1)))
			libc.Xmemmove(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fxs+uintptr(n)*8-uintptr(1)*8, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs, libc.Uint64FromInt32(m)*uint64(8))
			i = 0
			for {
				if !(i < (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fn-int32(1)) {
					break
				}
				**(**mpc_dtor_t)(__ccgo_up((*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(p)).Fdata))).Fdxs + uintptr(i)*8)) = __ccgo_fp(libc.Xfree)
				goto _6
			_6:
				;
				i = i + 1
			}
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fxs)
			libc.Xfree(tls, (*(*mpc_pdata_and_t)(unsafe.Pointer(&(*mpc_parser_t)(unsafe.Pointer(t)).Fdata))).Fdxs)
			libc.Xfree(tls, (*mpc_parser_t)(unsafe.Pointer(t)).Fname)
			libc.Xfree(tls, t)
			continue
		}
		return
	}
}

func mpc_optimise(tls *libc.TLS, p uintptr) {
	mpc_optimise_unretained(tls, p, int32(1))
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

func __ccgo_up(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(&n)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "%s\x00bell\x00backspace\x00formfeed\x00carriage return\x00vertical tab\x00end of input\x00newline\x00tab\x00space\x00%s: error: %s\n\x00%s:%li:%li: error: expected \x00ERROR: NOTHING EXPECTED\x00%s, \x00%s or %s\x00 at \x00\n\x00, \x00 or \x00one or more of \x00 of \x00%i of \x00\x00Maximum recursion depth exceeded!\x00Parser Undefined!\x00opposite\x00Unknown Parser Type Id!\x00Unknown Error\x00rb\x00Unable to open file!\x00Attempt to assign to Unretained Parser!\x00anchor\x00any character\x00'%c'\x00character between '%c' and '%c'\x00one of '%s'\x00none of '%s'\x00character satisfying function %p\x00\"%s\"\x00start of input\x00abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_\x00word boundary\x00start of newline\x00 \f\n\r\t\v\x00whitespace\x00spaces\x000123456789\x00digit\x000123456789ABCDEFabcdef\x00hex digit\x0001234567\x00oct digit\x00digits\x00hex digits\x00oct digits\x00abcdefghijklmnopqrstuvwxyz\x00lowercase letter\x00ABCDEFGHIJKLMNOPQRSTUVWXYZ\x00uppercase letter\x00abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\x00letter\x00underscore\x00alphanumeric\x00integer\x00hexadecimal\x00octadecimal\x00number\x00+-\x00eE\x00real\x00float\x00'\x00char\x00\"\x00string\x00/\x00regex\x00(\x00)\x00<\x00>\x00{\x00}\x00[\x00]\x00any character except a newline\x00-\x00\a\x00\f\x00\r\x00\t\x00\v\x00\b\x00Invalid Regex Range Expression\x00term\x00factor\x00base\x00range\x00)|\x00<mpc_re_compiler>\x00Invalid Regex: %s\x00\\a\x00\\b\x00\\f\x00\\n\x00\\r\x00\\t\x00\\v\x00\\\\\x00\\'\x00\\\"\x00\\0\x00\\/\x00<%s>\x00<anon>\x00<?>\x00<:>\x00<!>\x00<#>\x00<S>\x00<@>\x00<.>\x00<f>\x00'%s'\x00[%s-%s]\x00[%s]\x00[^%s]\x00!\x00?\x00*\x00+\x00{%i}\x00 (\x00 \x00 | \x00->?\x00<test>\x00Got \x00Expected \x00|\x00NULL\n\x00  \x00%s:%lu:%lu '%s'\n\x00%s \n\x00No Parser in position %i! Only supplied %i Parsers!\x00Unknown Parser '%s'!\x00grammar_total\x00grammar\x00ms\x00<mpc_grammar_compiler>\x00Invalid Grammar: %s\x00lang\x00stmt\x00:\x00;\x00<mpca_lang_file>\x00<mpca_lang_pipe>\x00<mpca_lang>\x00Stats\n\x00=====\n\x00Node Count: %i\n\x00"
