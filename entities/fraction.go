package entities

import (
	"fmt"
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int
}

// NewFraction creates a new fraction
func NewFraction(numerator, denominator *big.Int) *Fraction {
	return &Fraction{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

// Quotient performs floor division
func (f *Fraction) Quotient() *big.Int {
	return new(big.Int).Div(f.Numerator, f.Denominator)
}

// Remainder remainders after floor division
func (f *Fraction) Remainder() *Fraction {
	return NewFraction(new(big.Int).Rem(f.Numerator, f.Denominator), f.Denominator)
}

// Invert inverts the fraction
func (f *Fraction) Invert() *Fraction {
	return NewFraction(f.Denominator, f.Numerator)
}

// Add adds two fractions
func (f *Fraction) Add(other *Fraction) *Fraction {
	if f.Denominator.Cmp(other.Denominator) == 0 {
		return NewFraction(new(big.Int).Add(f.Numerator, other.Numerator), f.Denominator)
	}
	return NewFraction(
		new(big.Int).Add(new(big.Int).Mul(f.Numerator, other.Denominator), new(big.Int).Mul(other.Numerator, f.Denominator)),
		new(big.Int).Mul(f.Denominator, other.Denominator))
}

// Subtract subtracts two fractions
func (f *Fraction) Subtract(other *Fraction) *Fraction {
	if f.Denominator.Cmp(other.Denominator) == 0 {
		return NewFraction(new(big.Int).Sub(f.Numerator, other.Numerator), f.Denominator)
	}
	return NewFraction(
		new(big.Int).Sub(new(big.Int).Mul(f.Numerator, other.Denominator), new(big.Int).Mul(other.Numerator, f.Denominator)),
		new(big.Int).Mul(f.Denominator, other.Denominator))
}

// LessThan returns true if the fraction is less than the other fraction
func (f *Fraction) LessThan(other *Fraction) bool {
	return new(big.Int).Mul(f.Numerator, other.Denominator).Cmp(new(big.Int).Mul(other.Numerator, f.Denominator)) < 0
}

// EqualTo returns true if the fraction is equal to the other fraction
func (f *Fraction) EqualTo(other *Fraction) bool {
	return new(big.Int).Mul(f.Numerator, other.Denominator).Cmp(new(big.Int).Mul(other.Numerator, f.Denominator)) == 0
}

// GreaterThan returns true if the fraction is greater than the other fraction
func (f *Fraction) GreaterThan(other *Fraction) bool {
	return new(big.Int).Mul(f.Numerator, other.Denominator).Cmp(new(big.Int).Mul(other.Numerator, f.Denominator)) > 0
}

// Multiply multiplies two fractions
func (f *Fraction) Multiply(other *Fraction) *Fraction {
	return NewFraction(new(big.Int).Mul(f.Numerator, other.Numerator), new(big.Int).Mul(f.Denominator, other.Denominator))
}

// Divide divides two fractions
func (f *Fraction) Divide(other *Fraction) *Fraction {
	return NewFraction(new(big.Int).Mul(f.Numerator, other.Denominator), new(big.Int).Mul(f.Denominator, other.Numerator))
}

// ToSignificant returns a significant string representation of the fraction
// Example: NewFraction(big.NewInt(125), big.NewInt(1)).ToSignificant(2) // output: "130"
func (f *Fraction) ToSignificant(significantDigits int32) string {
	return roundToSignificantFigures(f, significantDigits).String()

}

// ToFixed returns a fixed string representation of the fraction
func (f *Fraction) ToFixed(decimalPlaces int32) string {
	return decimal.NewFromBigInt(f.Numerator, 0).Div(decimal.NewFromBigInt(f.Denominator, 0)).StringFixed(decimalPlaces)
}

var (
	oneInt = big.NewInt(1)
	twoInt = big.NewInt(2)
	tenInt = big.NewInt(10)
)

// roundToSignificantFigures returns a copy of a decimal rounded to specified number of significant figures.
// For negative or zero figures the function returns zero.
// Result is normalized without trailing zeros.
//
// Example:
// 	5.45 figures:2 --> 5.51
// 	545 figures:2 --> 550
// Ref: https://github.com/shopspring/decimal/pull/117/files#diff-84512ce9971597d4817ea830422f80d41b8f8f050b5998bd49499d8d1eebb16dR922
func roundToSignificantFigures(f *Fraction, figures int32) decimal.Decimal {
	if figures <= 0 {
		return decimal.Zero
	}
	d := decimal.NewFromBigInt(f.Numerator, 0).Div(decimal.NewFromBigInt(f.Denominator, 0))
	twoMant := d.Mul(decimal.NewFromFloat(math.Pow10(decimal.DivisionPrecision))).BigInt()
	twoMant.Abs(twoMant)
	twoMant.Mul(twoMant, twoInt)
	upper := big.NewInt(int64(figures))
	upper.Exp(tenInt, upper, nil)
	upper.Mul(upper, twoInt)
	upper.Sub(upper, oneInt)
	m := int64(0)
	for twoMant.Cmp(upper) >= 0 {
		upper.Mul(upper, tenInt)
		m++
	}
	if int64(d.Exponent())+m > int64(math.MaxInt32) {
		panic(fmt.Sprintf("exponent %d overflows an int32", int64(d.Exponent())+m))
	}
	return d.Round(-d.Exponent() - int32(m))
}
