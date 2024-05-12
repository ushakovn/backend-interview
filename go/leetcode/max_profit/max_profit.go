package max_profit

func maxProfitV2(prices []int) int {
	var (
		minPrice  int
		maxProfit int
		curProfit int
	)
	if len(prices) > 0 {
		minPrice = prices[0]
	} else {
		return 0
	}
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if curProfit = price - minPrice; curProfit > maxProfit {
			maxProfit = curProfit
		}
	}
	return maxProfit
}

func maxProfit(prices []int) int {
	var (
		profit int
		max    int
	)
	for idx := 0; idx < len(prices)-1; idx++ {
		for jdx := idx + 1; jdx < len(prices); jdx++ {
			if profit = prices[jdx] - prices[idx]; profit > max {
				max = profit
			}
		}
	}
	return max
}
