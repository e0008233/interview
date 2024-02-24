package dp

import "fmt"

func MaxProfit(k int, prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	if k >= len(prices)/2 {
		ans := 0

		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				ans = ans + prices[i] - prices[i-1]
			}
		}
		return ans

	}
	// buy[i] means max profits when the last transaction is buy at ith day
	// sell[i] means max profits when  the last transaction is sell at ith day
	// introduce another dimension k for when buying and selling with k transaction
	buy := make([][]int, len(prices)) // dp
	sell := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}

	for i := 0; i <= k; i++ {
		buy[0][i] = 0 - prices[0] // although no k transactions on the k day
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			// buy it, or do not buy
			buy[i][j] = max3(buy[i-1][j], sell[i-1][j-1]-prices[i])

			// buy it, or do not buy
			sell[i][j] = max3(sell[i-1][j], buy[i-1][j]+prices[i])
		}
	}

	fmt.Println(buy)

	fmt.Println(sell)
	maxPrice := sell[len(prices)-1][0]
	for _, value := range sell[len(prices)-1] {
		if value > maxPrice {
			maxPrice = value
		}
	}

	return maxPrice
}

func max3(i int, i2 int) int {
	if i > i2 {
		return i
	} else {
		return i2
	}
}
