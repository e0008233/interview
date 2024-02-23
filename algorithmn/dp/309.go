package dp

type Key2 struct {
	day             int
	emptyStockState bool
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	cache := make(map[Key2]int)
	return helper(prices, 0, true, &cache)
}

func helper(prices []int, i int, emptyStockState bool, cache *map[Key2]int) int {
	if i >= len(prices) {
		return 0
	}

	v, ok := (*cache)[Key2{i, emptyStockState}]
	if ok {
		return v
	}
	result := 0
	if emptyStockState {
		// can buy
		buyResult := helper(prices, i+1, false, cache) - prices[i]
		doNothingResult := helper(prices, i+1, emptyStockState, cache)
		result = max4(buyResult, doNothingResult)
	} else {
		sellResult := helper(prices, i+2, true, cache) + prices[i]
		doNothingResult := helper(prices, i+1, emptyStockState, cache)
		result = max4(sellResult, doNothingResult)
	}
	(*cache)[Key2{i, emptyStockState}] = result
	return result

}

func max4(i int, i2 int) int {
	if i > i2 {
		return i
	} else {
		return i2
	}
}
