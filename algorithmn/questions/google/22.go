package google

func GenerateParenthesis(n int) []string {
	result := make([]string, 0)

	backtrack(0, 0, n, "", &result)

	return result
}

func backtrack(left, right, n int, s string, result *[]string) {
	if right > left {
		return
	}
	if left > n {
		return
	}

	if right > n {
		return
	}
	if left+right == 2*n {
		*result = append(*result, s)
		return
	}

	s = s + "("
	left = left + 1
	backtrack(left, right, n, s, result)

	s = s[0 : len(s)-1]
	left = left - 1

	s = s + ")"
	right = right + 1
	backtrack(left, right, n, s, result)

	s = s[0 : len(s)-1]
	right = right - 1
}
