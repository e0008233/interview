package google

// explanation: https://www.zhihu.com/tardis/bd/art/403653338?source_id=1001
func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	if start == end {
		return 0
	}

	max := 0

	for end > start {
		if height[start] >= height[end] {
			current := (end - start) * height[end]
			if current > max {
				max = current
			}
			end--
		} else {
			current := (end - start) * height[start]
			if current > max {
				max = current
			}
			start++
		}
	}

	return max
}
