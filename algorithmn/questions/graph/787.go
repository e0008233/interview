package graph

import "math"

// https://www.youtube.com/watch?v=5eIK3zUdYmE
// Bellman-Ford (k layers of iterations of BFS)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = math.MaxInt32
	}
	prices[src] = 0

	for i := 0; i < k+1; i++ {
		copy(temp, prices)
		for _, v := range flights {
			source := v[0]
			dest := v[1]
			price := v[2]

			// not reachable yet
			if prices[source] == math.MaxInt32 {
				continue
			}

			// use prices[source] is important, as temp[source] may be updated
			if prices[source]+price < temp[dest] {
				temp[dest] = prices[source] + price

			}
		}

		copy(prices, temp)
	}

	if prices[dst] == math.MaxInt32 {
		return -1
	}

	return prices[dst]
}
