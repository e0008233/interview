package google

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	currentSum := nums[0] + nums[1] + nums[2]
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	currentDiff := math.Abs(float64(currentSum - target))
	for index1 := 0; index1 < len(nums)-2; index1++ {
		index2 := index1 + 1
		index3 := len(nums) - 1
		for index2 < index3 {
			current := nums[index1] + nums[index2] + nums[index3]
			if current == target {
				return current
			}
			diff := math.Abs(float64(current - target))
			if diff < currentDiff {
				currentDiff = diff
				currentSum = current
			}

			if current < target {
				index2++
			} else {
				index3--
			}
		}
	}

	return currentSum
}
