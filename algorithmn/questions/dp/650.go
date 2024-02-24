package dp

func minSteps(n int) int {
	if n <= 1 {
		return 0
	}

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i // copy and paste 'A' i times
		for j := 2; j < i-1; j = j + 1 {
			if i%j == 0 {
				dp[i] = min2(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}

func min2(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
