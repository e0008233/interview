package array

func Search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	// [start, end]
	for start <= end {
		middle := (start + end) / 2
		if target == nums[middle] {
			return middle
		} else if target > nums[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}

		if start == end {
			if nums[start] == target {
				return start
			} else {
				return -1
			}
		}
	}
	return -1
}
