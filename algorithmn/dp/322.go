package dp

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}

	for j := 1; j <= amount; j++ {
		dp[0][j] = amount + 1 // impossible with 0 coins
	}
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {

			if coins[i-1] <= j {
				dp[i][j] = min(dp[i][j-coins[i-1]]+1, dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j] // the same as not taking the current coins
			}
		}
	}

	if dp[len(coins)][amount] == amount+1 {
		return -1
	}
	return dp[len(coins)][amount]
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
