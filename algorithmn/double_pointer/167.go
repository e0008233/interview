package double_pointer

func twoSum(numbers []int, target int) []int {

	//sort.Slice(numbers,func(i, j int)bool {
	//	return numbers[i]<numbers[j]
	//})
	start := 0
	end := len(numbers) - 1

	for start <= end {
		if numbers[start]+numbers[end] == target {
			result := []int{start + 1, end + 1}
			return result
		} else if numbers[start]+numbers[end] < target {
			start++
		} else {
			end--
		}
	}

	return []int{start + 1, end + 1}
}
