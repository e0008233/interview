package array

func subarraySum(nums []int, k int) int {
	result := 0

	sum := 0
	sumCount := make(map[int]int)
	sumCount[0] = 1
	for _, v := range nums {
		sum = sum + v
		value, ok := sumCount[sum-k]
		if ok {
			result = result + value
		}

		value, ok = sumCount[sum]
		if ok {
			sumCount[sum] = value + 1
		} else {
			sumCount[sum] = 1
		}
	}

	return result
}
