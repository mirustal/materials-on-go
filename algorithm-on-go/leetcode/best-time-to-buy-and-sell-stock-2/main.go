package best_time_to_buy_and_sell_stock_2

func maxProfit(prices []int) int {
	total := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			total += prices[i] - prices[i-1]
		}
	}

	return total
}
