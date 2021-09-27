package entities

import (
	"math/big"
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
	return NewFraction(
		new(big.Int).Add(new(big.Int).Mul(f.Numerator, other.Denominator), other.Numerator),
		new(big.Int).Mul(f.Denominator, other.Denominator))
}

// Subtract subtracts two fractions
func (f *Fraction) Subtract(other *Fraction) *Fraction {
	return NewFraction(
		new(big.Int).Sub(new(big.Int).Mul(f.Numerator, other.Denominator), other.Numerator),
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

// ToSignificant returns a string representation of the fraction
func (f *Fraction) ToSignificant(significantDigits uint) string {
	// TODO
	return ""
}

// ToFixed returns a string representation of the fraction
func (f *Fraction) ToFixed(decimalPlaces uint) string {
	// TODO
	return ""
}
