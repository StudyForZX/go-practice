package ltmath

import "sort"

func LT123_MaxProfit(prices []int) int {

	// 现在最多交易两次，所以定义两组buy以及sell
	firstBuy, firstSell := -prices[0], 0
	secondBuy, secondSell := -prices[0], 0

	for _, price := range prices {
		// 第一次买 -price
		firstBuy = max(firstBuy, -price)
		// 第一次卖 firstBuy + price
		firstSell = max(firstSell, firstBuy+price)
		// 第一次卖了后现在买 firstSell - price
		secondBuy = max(secondBuy, firstSell-price)
		// 第二次买了后现在卖 secondBuy + price
		secondSell = max(secondSell, secondBuy+price)
	}

	return secondSell
}

// 目前这个题解有问题  需要进行修改
func LT123_MaxProfitByFor(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	sums := []int{}
	sum := 0
	mark := false

	for i := 0; i < len(prices)-1; i++ {

		if prices[i] < prices[i+1] {

			tmp := prices[i+1] - prices[i]

			sum = tmp + sum

			mark = true

		} else {

			if mark {
				sums = append(sums, sum)
				sum = 0
				mark = false
			}

		}

	}

	sums = append(sums, sum)

	sort.Ints(sums)

	sum1, sum2 := 0, 0

	if len(sums) == 0 {
		return 0
	}

	if len(sums) == 1 {
		return sums[0]
	}

	sum1 = sums[len(sums)-1]
	sum2 = sums[len(sums)-2]
	return sum1 + sum2
}
