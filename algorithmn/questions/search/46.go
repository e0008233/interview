package search

func Permute(nums []int) [][]int {

	result := make([][]int, 0)
	used := make([]bool, len(nums))
	currentSet := make([]int, 0)
	backTrack(nums, &result, currentSet, used)
	return result
}

func backTrack(nums []int, result *[][]int, currentSet []int, used []bool) {
	if len(currentSet) == len(nums) {
		// Create a copy of currentSet before appending to result
		temp := make([]int, len(currentSet))
		copy(temp, currentSet)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			currentSet = append(currentSet, nums[i])
			backTrack(nums, result, currentSet, used)
			currentSet = currentSet[:len(currentSet)-1]
			used[i] = false
		}
	}

}
