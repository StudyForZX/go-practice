package ltmath

func LT121_MaxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	max := 0

	for i := 1; i < len(prices); i++ {

		if prices[i] < buy {
			buy = prices[i]
		} else {
			tmp := prices[i] - buy

			if tmp > max {
				max = tmp
			}
		}

	}

	return max
}
