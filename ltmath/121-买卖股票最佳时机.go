package ltmath

func LT121MaxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	maxRes := 0
	buy := prices[0]

	for _, price := range prices {

		if price < buy {
			buy = price
		} else {
			maxRes = max(price-buy, maxRes)
		}
	}

	return maxRes
}
