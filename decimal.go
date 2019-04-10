// Package d128 represents a decimal with a 128 bit sized coefficent.
//
// Right now this is a prototype for API design only. Do not use.
//
// This is not designed to be a IEEE complient decimal type. Rather,
// it should do the logical thing for Go. The main decimal value
// should not be a pointer. The context containing the rounding mode
// and precision for operations should be separate from the decimal value.
// Lastly, while there should be an option to not panic on bad values,
// or too large of values, it should be default act like other Go numeric
// operations.
//
// The operations it nativly supports should be basic arithmatic.
// The API should be simple to use correctly and easy to understand. The
// API should not have many options.
//
// Infinity is not supported. An invalid operation (out of range or divide
// by zero) will be default panic, NaN and Zero do not have negative forms.
//
// In C#, Java, PostgreSQL, and SQL Server decimals also do not support Infinity.
// PostgreSQL also supports NaN, and NaNs are useful to get non-panic'ing operations.
// The largest use case for decimals is handling money. Generally money can fit
// fully in a fixed decimal128 representation.
package d128

import (
	"fmt"
	"math/big"
)

// Decimal represents a number in the form of (neg) coef * 10 ^ exponent form.
// It also can be set to NaN and Infinite states.
// Exponent must be within the range of [-1000, 1000].
type Decimal struct { // 24 byte size.
	hi  uint64
	lo  uint64
	exp int32
	neg bool
	nan bool
}

// Negative returns true if the decimal is negative.
func (x Decimal) Negative() bool {
	if x.nan {
		return false
	}
	if x.hi == 0 && x.lo == 0 {
		return false
	}
	return x.neg
}

// Negate x and return the result.
func (x Decimal) Negate() Decimal {
	x.neg = !x.neg
	return x
}

// Abs returns the absolute value of x.
func (x Decimal) Abs() Decimal {
	x.neg = false
	return x
}

// IsNaN returns true if the decimal is a NaN.
func (x Decimal) IsNaN() bool {
	return x.nan
}

// IsInt reports whether x is an integer. ±Inf values are not integers.
func (x Decimal) IsInt() bool {
	panic("TODO")
}

// IsInt reports whether x is an integer. ±Inf values are not integers.
func (x Decimal) IsZero() bool {
	if x.nan {
		return false
	}
	return x.hi == 0 && x.lo == 0
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y (incl. NaN == NaN)
//   +1 if x >  y
//
// A NaN will always be greater then any non-NaN Decimal.
func (x Decimal) Cmp(y Decimal) int {
	panic("TODO")
}

func (x Decimal) Format(f fmt.State, c rune) {
	panic("TODO")
}
func (x Decimal) String() string {
	panic("TODO")
}
func (x *Decimal) Scan(state fmt.ScanState, verb rune) error {
	panic("TODO")
}

func (x Decimal) UnmarshalText(text []byte) error {
	panic("TODO")
}
func (x Decimal) MarshalText() (text []byte, err error) {
	panic("TODO")
}

func (x Decimal) MarshalBinary() (data []byte, err error) {
	// byte version + negative + form
	//    First 4 bits is version, allows for 16 versions.
	//    bit-5 is negative
	//    bit-6 is nan
	// Varint exp
	// Uvarint lo+hi
	panic("TODO")
}

func (x *Decimal) UnmarshalBinary(data []byte) error {
	panic("TODO")
}

// Compose sets the internal decimal value from parts. If the value cannot be
// represented then an error should be returned.
func (x *Decimal) Compose(form byte, negative bool, coefficient []byte, exponent int32) error {
	panic("TODO")
}

// Decompose returns the internal decimal state into parts.
// If the provided buf has sufficient capacity, buf may be returned as the coefficient with
// the value set and length set as appropriate.
func (x Decimal) Decompose(buf []byte) (form byte, negative bool, coefficient []byte, exponent int32) {
	panic("TODO")
}

func (x Decimal) CoefExp() (coef int64, exp int, err error) {
	panic("TODO")
}

func (x Decimal) IntNano() (integer int64, nano int, err error) {
	panic("TODO")
}

type Operation struct {
	// The scale of the decimal to return from an
	// operation. Scale == -Exponent.
	Scale int

	// Rounding Mode for operations.
	RoundingMode big.RoundingMode

	// Return a NaN Decimal and do not panic if true on an invalid operation.
	ReturnNaN bool
}

func NewOperation(scale int) Operation {
	return Operation{Scale: scale}
}

// Round decimal x and according to the operation Scale and
// RoundingMode and return the result.
func (o Operation) Round(x Decimal) Decimal {
	panic("TODO")
}

// Add returns (x+y).
func (o Operation) Add(x, y Decimal) Decimal {
	panic("TODO")
}

// Sub returns (x-y).
func (o Operation) Sub(x, y Decimal) Decimal {
	panic("TODO")
}

// Mul returns (x*y).
func (o Operation) Mul(x, y Decimal) Decimal {
	panic("TODO")
}

// Div returns (x/y).
func (o Operation) Div(x, y Decimal) Decimal {
	panic("TODO")
}

// Parse the string s.
// Parsing "NaN" will return an error unless ReturnNaN is true.
func (o Operation) Parse(s string) (Decimal, error) {
	panic("TODO")
}

// CoefExp returns a decimal from a coefficent and exponent
// in the form of "decimal = coef * 10 ^ exp".
func (o Operation) CoefExp(coef int64, exp int) Decimal {
	panic("TODO")
}

// IntNano returns a decimal in the form of "decimal = integer + (nano * 10 ^-9)".
func (o Operation) IntNano(integer int64, nano int) Decimal {
	panic("TODO")
}
