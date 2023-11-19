package search

func combine(n int, k int) [][]int {

	used := make([]bool, n+1)
	result := make([][]int, 0)
	currentSet := make([]int, 0)
	backTrack2(n, k, &result, currentSet, used, 0)

	return result
}

func backTrack2(n int, k int, result *[][]int, set []int, used []bool, next int) {
	if k == len(set) {
		temp := make([]int, k)
		copy(temp, set)
		*result = append(*result, temp)
		return
	}

	for i := next + 1; i <= n; i++ {
		if !used[i] {
			used[i] = true
			set = append(set, i)
			backTrack2(n, k, result, set, used, i)

			used[i] = false
			set = set[:len(set)-1]

		}
	}

}
