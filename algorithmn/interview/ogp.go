package interview

func Solution(A []int) int {
	n := len(A)
	if n < 2 {
		return 0
	}

	// Map to store sums and their corresponding segment start indices
	sumSegments := make(map[int][]int)

	// Compute sums of all possible segments of length 2
	for i := 0; i < n-1; i++ {
		sum := A[i] + A[i+1]
		sumSegments[sum] = append(sumSegments[sum], i)
	}

	maxSegments := 0

	// Iterate through each sum's segments to find non-intersecting segments
	for _, indices := range sumSegments {
		count := 0
		lastUsed := -2 // Initialize to -2 to avoid conflict with the first segment

		for _, idx := range indices {
			// Check if the current segment does not intersect with the previous one
			if idx > lastUsed+1 {
				count++
				lastUsed = idx
			}
		}

		if count > maxSegments {
			maxSegments = count
		}
	}

	return maxSegments
}
