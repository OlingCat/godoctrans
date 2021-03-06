// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package strconv implements conversions to and from string representations of
// basic data types.
package strconv

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target
// type.
var ErrSyntax = errors.New("invalid syntax")

// AppendBool appends "true" or "false", according to the value of b, to dst and
// returns the extended buffer.
func AppendBool(dst []byte, b bool) []byte

// AppendFloat appends the string form of the floating-point number f, as generated
// by FormatFloat, to dst and returns the extended buffer.
func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte

// AppendInt appends the string form of the integer i, as generated by FormatInt,
// to dst and returns the extended buffer.
func AppendInt(dst []byte, i int64, base int) []byte

// AppendQuote appends a double-quoted Go string literal representing s, as
// generated by Quote, to dst and returns the extended buffer.
func AppendQuote(dst []byte, s string) []byte

// AppendQuoteRune appends a single-quoted Go character literal representing the
// rune, as generated by QuoteRune, to dst and returns the extended buffer.
func AppendQuoteRune(dst []byte, r rune) []byte

// AppendQuoteRuneToASCII appends a single-quoted Go character literal representing
// the rune, as generated by QuoteRuneToASCII, to dst and returns the extended
// buffer.
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte

// AppendQuoteToASCII appends a double-quoted Go string literal representing s, as
// generated by QuoteToASCII, to dst and returns the extended buffer.
func AppendQuoteToASCII(dst []byte, s string) []byte

// AppendUint appends the string form of the unsigned integer i, as generated by
// FormatUint, to dst and returns the extended buffer.
func AppendUint(dst []byte, i uint64, base int) []byte

// Atoi is shorthand for ParseInt(s, 10, 0).
func Atoi(s string) (i int, err error)

// CanBackquote reports whether the string s can be represented unchanged as a
// single-line backquoted string without control characters other than tab.
func CanBackquote(s string) bool

// FormatBool returns "true" or "false" according to the value of b
func FormatBool(b bool) string

// FormatFloat converts the floating-point number f to a string, according to the
// format fmt and precision prec. It rounds the result assuming that the original
// was obtained from a floating-point value of bitSize bits (32 for float32, 64 for
// float64).
//
// The format fmt is one of 'b' (-ddddp±ddd, a binary exponent), 'e' (-d.dddde±dd,
// a decimal exponent), 'E' (-d.ddddE±dd, a decimal exponent), 'f' (-ddd.dddd, no
// exponent), 'g' ('e' for large exponents, 'f' otherwise), or 'G' ('E' for large
// exponents, 'f' otherwise).
//
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', and 'G' formats. For 'e', 'E', and 'f' it is
// the number of digits after the decimal point. For 'g' and 'G' it is the total
// number of digits. The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

// FormatInt returns the string representation of i in the given base, for 2 <=
// base <= 36. The result uses the lower-case letters 'a' to 'z' for digit values
// >= 10.
func FormatInt(i int64, base int) string

// FormatUint returns the string representation of i in the given base, for 2 <=
// base <= 36. The result uses the lower-case letters 'a' to 'z' for digit values
// >= 10.
func FormatUint(i uint64, base int) string

// IsPrint reports whether the rune is defined as printable by Go, with the same
// definition as unicode.IsPrint: letters, numbers, punctuation, symbols and ASCII
// space.
func IsPrint(r rune) bool

// Itoa is shorthand for FormatInt(i, 10).
func Itoa(i int) string

// ParseBool returns the boolean value represented by the string. It accepts 1, t,
// T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an
// error.
func ParseBool(str string) (value bool, err error)

// ParseFloat converts the string s to a floating-point number with the precision
// specified by bitSize: 32 for float32, or 64 for float64. When bitSize=32, the
// result still has type float64, but it will be convertible to float32 without
// changing its value.
//
// If s is well-formed and near a valid floating point number, ParseFloat returns
// the nearest floating point number rounded using IEEE754 unbiased rounding.
//
// The errors that ParseFloat returns have concrete type *NumError and include
// err.Num = s.
//
// If s is not syntactically well-formed, ParseFloat returns err.Err = ErrSyntax.
//
// If s is syntactically well-formed but is more than 1/2 ULP away from the largest
// floating point number of the given size, ParseFloat returns f = ±Inf, err.Err =
// ErrRange.
func ParseFloat(s string, bitSize int) (f float64, err error)

// ParseInt interprets a string s in the given base (2 to 36) and returns the
// corresponding value i. If base == 0, the base is implied by the string's prefix:
// base 16 for "0x", base 8 for "0", and base 10 otherwise.
//
// The bitSize argument specifies the integer type that the result must fit into.
// Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64.
//
// The errors that ParseInt returns have concrete type *NumError and include
// err.Num = s. If s is empty or contains invalid digits, err.Err = ErrSyntax and
// the returned value is 0; if the value corresponding to s cannot be represented
// by a signed integer of the given size, err.Err = ErrRange and the returned value
// is the maximum magnitude integer of the appropriate bitSize and sign.
func ParseInt(s string, base int, bitSize int) (i int64, err error)

// ParseUint is like ParseInt but for unsigned numbers.
func ParseUint(s string, base int, bitSize int) (n uint64, err error)

// Quote returns a double-quoted Go string literal representing s. The returned
// string uses Go escape sequences (\t, \n, \xFF, \u0100) for control characters
// and non-printable characters as defined by IsPrint.
func Quote(s string) string

// QuoteRune returns a single-quoted Go character literal representing the rune.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for control
// characters and non-printable characters as defined by IsPrint.
func QuoteRune(r rune) string

// QuoteRuneToASCII returns a single-quoted Go character literal representing the
// rune. The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// non-ASCII characters and non-printable characters as defined by IsPrint.
func QuoteRuneToASCII(r rune) string

// QuoteToASCII returns a double-quoted Go string literal representing s. The
// returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for non-ASCII
// characters and non-printable characters as defined by IsPrint.
func QuoteToASCII(s string) string

// Unquote interprets s as a single-quoted, double-quoted, or backquoted Go string
// literal, returning the string value that s quotes. (If s is single-quoted, it
// would be a Go character literal; Unquote returns the corresponding one-character
// string.)
func Unquote(s string) (t string, err error)

// UnquoteChar decodes the first character or byte in the escaped string or
// character literal represented by the string s. It returns four values:
//
//	1) value, the decoded Unicode code point or byte value;
//	2) multibyte, a boolean indicating whether the decoded character requires a multibyte UTF-8 representation;
//	3) tail, the remainder of the string after the character; and
//	4) an error that will be nil if the character is syntactically valid.
//
// The second argument, quote, specifies the type of literal being parsed and
// therefore which escaped quote character is permitted. If set to a single quote,
// it permits the sequence \' and disallows unescaped '. If set to a double quote,
// it permits \" and disallows unescaped ". If set to zero, it does not permit
// either escape and allows both quote characters to appear unescaped.
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)

// A NumError records a failed conversion.
type NumError struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
	Num  string // the input
	Err  error  // the reason the conversion failed (ErrRange, ErrSyntax)
}

func (e *NumError) Error() string
