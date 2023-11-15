package double_pointer

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1End := m - 1
	nums2End := n - 1
	for i := len(nums1) - 1; i >= 0; i-- {
		if nums1End < 0 {
			nums1[i] = nums2[nums2End]
			nums2End--
		} else if nums2End < 0 {
			nums1[i] = nums1[nums1End]
			nums1End--
		} else if nums1[nums1End] < nums2[nums2End] {
			nums1[i] = nums2[nums2End]
			nums2End--
		} else {
			nums1[i] = nums1[nums1End]
			nums1End--
		}
	}
}
