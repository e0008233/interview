package binary_search

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			begin := middle
			end := middle
			for begin >= 0 {
				if nums[begin] != target {
					break
				}
				begin--
			}

			for end <= len(nums)-1 {
				if nums[end] != target {
					break
				}
				end++
			}
			return []int{begin + 1, end - 1}
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}

	}
	return []int{-1, -1}
}
