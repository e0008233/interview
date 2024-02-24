package google

// https://leetcode.cn/problems/next-permutation/solutions/80560/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 {
		if nums[j] > nums[i] {
			break
		}
		i--
		j--
	}
	k := 0
	if i >= 0 {
		for k = len(nums) - 1; k >= j; k-- {
			if nums[k] > nums[i] {
				break
			}
		}
		temp := nums[k]
		nums[k] = nums[i]
		nums[i] = temp
	}

	k = len(nums) - 1
	for j < k {
		temp := nums[k]
		nums[k] = nums[j]
		nums[j] = temp

		j++
		k--
	}
}
