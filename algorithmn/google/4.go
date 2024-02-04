package google

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	isOdd := false
	if (len(nums1)+len(nums2))%2 == 1 {
		isOdd = true
	}

	target := (len(nums1) + len(nums2)) / 2
	if isOdd {
		target++
	}

	count := 0
	num1Index := -1
	num2Index := -1
	for {

		if num1Index >= len(nums1)-1 {
			count++
			num2Index++
			if count == target {
				if isOdd {
					return float64(nums2[num2Index])
				} else {
					return float64(nums2[num2Index]+nums2[num2Index+1]) / 2.0
				}
			}

		} else if num2Index >= len(nums2)-1 {
			count++
			num1Index++
			if count == target {
				if isOdd {
					return float64(nums1[num1Index])
				} else {
					return float64(nums1[num1Index]+nums1[num1Index+1]) / 2.0
				}
			}
		} else {
			count++
			if nums2[num2Index+1] < nums1[num1Index+1] {
				num2Index++
				if count == target {
					if isOdd {
						return float64(nums2[num2Index])
					} else {
						next := getNext(nums1, num1Index, nums2, num2Index)
						return float64(nums2[num2Index]+next) / 2.0
					}
				}
			} else {
				num1Index++
				if count == target {
					if isOdd {
						return float64(nums1[num1Index])
					} else {
						next := getNext(nums1, num1Index, nums2, num2Index)
						return float64(nums1[num1Index]+next) / 2.0
					}
				}
			}
		}

	}
}

func getNext(nums1 []int, num1Index int, nums2 []int, num2Index int) int {
	if num1Index >= len(nums1)-1 {
		return nums2[num2Index+1]
	} else if num2Index >= len(nums2)-1 {
		return nums1[num1Index+1]
	} else {
		if nums2[num2Index+1] < nums1[num1Index+1] {
			return nums2[num2Index+1]
		} else {
			return nums1[num1Index+1]
		}
	}
}
