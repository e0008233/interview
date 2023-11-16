package sorting

import "sort"

func TopKFrequent(nums []int, k int) []int {
	myMap := make(map[int]int, 0)
	valueArray := make([]int, 0)

	for _, v := range nums {

		_, ok := myMap[v]
		if ok {
			myMap[v] = myMap[v] + 1
		} else {
			myMap[v] = 1
			valueArray = append(valueArray, v)
		}
	}

	sort.Slice(valueArray, func(i, j int) bool {
		return myMap[valueArray[i]] > myMap[valueArray[j]]
	})

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, valueArray[i])
	}
	return result
}
