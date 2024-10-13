package ltmath

func LT122_MaxProfit(prices []int) int {

	sum := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			tmp := prices[i+1] - prices[i]
			sum = sum + tmp
		}
	}

	return sum
}
