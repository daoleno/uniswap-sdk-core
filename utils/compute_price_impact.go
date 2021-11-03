package utils

import "github.com/daoleno/uniswap-sdk-core/entities"

/**
 * Returns the percent difference between the mid price and the execution price, i.e. price impact.
 * @param midPrice mid price before the trade
 * @param inputAmount the input amount of the trade
 * @param outputAmount the output amount of the trade
 */
func ComputePriceImpact(midPrice *entities.Price, inputAmount, outputAmount *entities.CurrencyAmount) (*entities.Percent, error) {
	quotedOutputAmount, err := midPrice.Quote(inputAmount)
	if err != nil {
		return nil, err
	}
	priceImpact := quotedOutputAmount.Subtract(outputAmount).Divide(quotedOutputAmount.Fraction)
	return entities.NewPercent(priceImpact.Numerator, priceImpact.Denominator), nil
}
