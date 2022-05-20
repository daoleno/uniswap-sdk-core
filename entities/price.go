package entities

import (
	"errors"
	"math/big"
)

var (
	ErrDifferentCurrencies = errors.New("different currencies")
)

type Price struct {
	*Fraction
	BaseCurrency  Currency  // input i.e. denominator
	QuoteCurrency Currency  // output i.e. numerator
	Scalar        *Fraction // used to adjust the raw fraction w/r/t the decimals of the {base,quote}Token
}

// Construct a price, either with the base and quote currency amount, or the args
func NewPrice(baseCurrency, quoteCurrency Currency, denominator, numerator *big.Int) *Price {
	return &Price{
		Fraction:      NewFraction(numerator, denominator),
		BaseCurrency:  baseCurrency,
		QuoteCurrency: quoteCurrency,
		Scalar: NewFraction(
			new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(baseCurrency.Decimals())), nil),
			new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(quoteCurrency.Decimals())), nil)),
	}
}

// Invert flips the price, switching the base and quote currency
func (p *Price) Invert() *Price {
	return NewPrice(p.QuoteCurrency, p.BaseCurrency, p.Numerator, p.Denominator)
}

// Multiply Multiplies the price by another price, returning a new price. The other price must have the same base currency as this price's quote currency
func (p *Price) Multiply(other *Price) (*Price, error) {
	if !other.BaseCurrency.Equal(p.QuoteCurrency) {
		return nil, ErrDifferentCurrencies
	}

	fraction := p.Fraction.Multiply(other.Fraction)
	return NewPrice(p.BaseCurrency, other.QuoteCurrency, fraction.Denominator, fraction.Numerator), nil
}

// Quote returns the amount of quote currency corresponding to a given amount of the base currency
func (p *Price) Quote(currencyAmount *CurrencyAmount) (*CurrencyAmount, error) {
	if !currencyAmount.Currency.Equal(p.BaseCurrency) {
		return nil, ErrDifferentCurrencies
	}

	result := p.Fraction.Multiply(currencyAmount.Fraction)
	return newCurrencyAmount(p.QuoteCurrency, result.Numerator, result.Denominator), nil
}

// adjustedForDecimals Get the value scaled by decimals for formatting
func (p *Price) adjustedForDecimals() *Fraction {
	return p.Fraction.Multiply(p.Scalar)
}

func (p *Price) ToSignificant(significantDigits int32) string {
	return p.adjustedForDecimals().ToSignificant(significantDigits)
}

func (p *Price) ToFixed(decimalPlaces int32) string {
	return p.adjustedForDecimals().ToFixed(decimalPlaces)
}
