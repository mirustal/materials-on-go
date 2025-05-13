package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {

	bestPrice := 0
	var low int
	for i := range prices {
		if prices[low] < prices[i] {
			bestPrice = max(bestPrice, prices[i]-prices[low])
		} else {
			low = i
		}
	}

	return bestPrice
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
