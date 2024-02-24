package greedy_algo

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	result := 0
	childIndex := 0
	for _, value := range s {
		if g[childIndex] <= value {
			result++
			childIndex++
		}
		if childIndex >= len(g) {
			return result
		}
	}
	return result
}
