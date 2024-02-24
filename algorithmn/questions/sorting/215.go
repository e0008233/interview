package sorting

func FindKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k+1)
}

func quickSort(nums []int, left, right int, k int) int {

	if left == right {
		return nums[left]
	}
	pivot := partition(nums, left, right, k)

	if k-1 == pivot {
		return nums[pivot]
	} else if k-1 < pivot {
		return quickSort(nums, left, pivot-1, k)
	} else {
		return quickSort(nums, pivot+1, right, k)
	}
}

func partition(nums []int, left int, right int, k int) int {
	current := nums[left : right+1]
	println(current)
	pivotNum := nums[left]
	pivotIndex := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivotNum {
			pivotIndex++
			swap(nums, pivotIndex, i)

		}
	}
	swap(nums, pivotIndex, left)
	return pivotIndex
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
