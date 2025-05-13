package best_time_to_buy_and_sell_stock_with_transaction_fee

import "math"

func maxProfit(prices []int, fee int) int {
	profit, hold := 0, -math.MaxInt32

	for _, v := range prices {
		profit, hold = max(profit, hold+v), max(hold, profit-v-fee)
	}

	return profit
}
