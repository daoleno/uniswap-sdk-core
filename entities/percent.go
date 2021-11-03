package entities

import "math/big"

var OneHundred = NewFraction(big.NewInt(100), big.NewInt(1))

type Percent struct {
	*Fraction
}

/**
 * Converts a fraction to a percent
 * @param fraction the fraction to convert
 */
func toPercent(fraction *Fraction) *Percent {
	return NewPercent(fraction.Numerator, fraction.Denominator)
}

// NewPercent creates a new Percent
func NewPercent(numerator, denominator *big.Int) *Percent {
	return &Percent{NewFraction(numerator, denominator)}
}

// Add adds two Percent
func (p *Percent) Add(other *Percent) *Percent {
	return toPercent(p.Fraction.Add(other.Fraction))
}

// Subtract subtracts two Percent
func (p *Percent) Subtract(other *Percent) *Percent {
	return toPercent(p.Fraction.Subtract(other.Fraction))
}

// Multiply multiplies two Percent
func (p *Percent) Multiply(other *Percent) *Percent {
	return toPercent(p.Fraction.Multiply(other.Fraction))
}

// Divide divides two Percent
func (p *Percent) Divide(other *Percent) *Percent {
	return toPercent(p.Fraction.Divide(other.Fraction))
}

// ToSignificant converts a Percent to a string with a given number of significant figures
func (p *Percent) ToSignificant(significantDigits int32) string {
	return p.Fraction.Multiply(OneHundred).ToSignificant(significantDigits)
}

// ToFixedFigures converts a Percent to a string with a given number of fixed figures
func (p *Percent) ToFixed(decimalPlaces int32) string {
	return p.Fraction.Multiply(OneHundred).ToFixed(decimalPlaces)
}
