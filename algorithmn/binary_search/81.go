package binary_search

func Search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return true
		}

		if nums[middle] < nums[right] {
			if nums[middle] < target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else if nums[left] < nums[middle] {
			if nums[middle] > target && target >= nums[left] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if nums[middle] == nums[right] {
			right--
		} else {
			left++
		}

	}

	return false
}
