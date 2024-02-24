package greedy_algo

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	removed := 0
	prev := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prev {
			removed++
		} else {
			prev = intervals[i][1]
		}

	}

	return removed
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
