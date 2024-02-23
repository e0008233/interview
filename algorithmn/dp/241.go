package dp

import "strconv"

func DiffWaysToCompute(expression string) []int {

	cache := make(map[string][]int)

	return diffWaysToComputeHelper(expression, &cache)
}

func diffWaysToComputeHelper(expression string, cache *map[string][]int) []int {
	v, ok := (*cache)[expression]
	if ok {
		return v
	}
	result := make([]int, 0)

	for i := 0; i < len(expression); i++ {

		if expression[i] == '+' {
			left := diffWaysToComputeHelper(expression[0:i], cache)
			right := diffWaysToComputeHelper(expression[i+1:len(expression)], cache)
			for _, num1 := range left {
				for _, num2 := range right {
					result = append(result, num1+num2)
				}
			}
		} else if expression[i] == '-' {
			left := diffWaysToComputeHelper(expression[0:i], cache)
			right := diffWaysToComputeHelper(expression[i+1:len(expression)], cache)
			for _, num1 := range left {
				for _, num2 := range right {
					result = append(result, num1-num2)
				}
			}
		} else if expression[i] == '*' {
			left := diffWaysToComputeHelper(expression[0:i], cache)
			right := diffWaysToComputeHelper(expression[i+1:len(expression)], cache)
			for _, num1 := range left {
				for _, num2 := range right {
					result = append(result, num1*num2)
				}
			}

		}
	}

	if len(result) == 0 { // single number case, e.g. (2)
		i, _ := strconv.Atoi(expression)
		result = append(result, i)
	}

	(*cache)[expression] = result
	return result
}
