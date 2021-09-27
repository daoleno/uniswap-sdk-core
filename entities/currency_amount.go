package entities

import "math/big"

type CurrencyAmount struct {
	*Fraction
	Currency     *Currency
	DecimalScale *big.Int
}

/**
 * Returns a new currency amount instance from the unitless amount of token, i.e. the raw amount
 * @param currency the currency in the amount
 * @param rawAmount the raw token or ether amount
 */
func FromRawAmount(currency *Currency, rawAmount *big.Int) *CurrencyAmount {
	return NewCurrencyAmount(currency, rawAmount, big.NewInt(1))
}

/**
 * Construct a currency amount with a denominator that is not equal to 1
 * @param currency the currency
 * @param numerator the numerator of the fractional token amount
 * @param denominator the denominator of the fractional token amount
 */
func FromFractionalAmount(currency *Currency, numerator *big.Int, denominator *big.Int) *CurrencyAmount {
	return NewCurrencyAmount(currency, numerator, denominator)
}

// NewCurrencyAmount creates a new CurrencyAmount instance
func NewCurrencyAmount(currency *Currency, numerator, denominator *big.Int) *CurrencyAmount {
	return &CurrencyAmount{
		Currency:     currency,
		Fraction:     NewFraction(numerator, denominator),
		DecimalScale: new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(currency.Decimals)), nil),
	}
}

// Add adds two currency amounts together
func (ca *CurrencyAmount) Add(other *CurrencyAmount) *CurrencyAmount {
	added := ca.Fraction.Add(other.Fraction)
	return FromFractionalAmount(ca.Currency, added.Numerator, added.Denominator)
}

// Subtract subtracts one currency amount from another
func (ca *CurrencyAmount) Subtract(other *CurrencyAmount) *CurrencyAmount {
	subtracted := ca.Fraction.Subtract(other.Fraction)
	return FromFractionalAmount(ca.Currency, subtracted.Numerator, subtracted.Denominator)
}

// Multiply multiplies two currency amounts
func (ca *CurrencyAmount) Multiply(other *CurrencyAmount) *CurrencyAmount {
	multiplied := ca.Fraction.Multiply(other.Fraction)
	return FromFractionalAmount(ca.Currency, multiplied.Numerator, multiplied.Denominator)
}

// Divide divides one currency amount by another
func (ca *CurrencyAmount) Divide(other *CurrencyAmount) *CurrencyAmount {
	divided := ca.Fraction.Divide(other.Fraction)
	return FromFractionalAmount(ca.Currency, divided.Numerator, divided.Denominator)
}

// ToSignificant returns the currency amount as a string with the most significant digits
func (ca *CurrencyAmount) ToSignificant(significantDigits int) string {
	// TODO
	return ""
}

// ToFixed returns the currency amount as a string with the specified number of digits after the decimal
func (ca *CurrencyAmount) ToFixed(decimalPlaces int) string {
	// TODO
	return ""
}

// ToExact returns the currency amount as a string with the specified number of digits after the decimal
func (ca *CurrencyAmount) ToExact(decimalPlaces int) string {
	// TODO
	return ""
}
