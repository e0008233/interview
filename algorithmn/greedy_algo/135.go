package greedy_algo

import "fmt"

func Candy(ratings []int) int {

	if len(ratings) == 1 {
		return 1
	}

	if len(ratings) == 2 {
		if ratings[0] == ratings[1] {
			return 2
		} else {
			return 3
		}
	}

	cookies := make([]int, len(ratings))
	for k, _ := range cookies {
		cookies[k] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			cookies[i] = cookies[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			cookies[i] = max(cookies[i], cookies[i+1]+1)
		}
	}
	fmt.Println(cookies)
	result := 0
	for _, v := range cookies {
		result = result + v
	}

	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
