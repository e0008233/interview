package double_pointer

func MinWindow(s string, t string) string {

	counter := make(map[string]int)
	for _, v := range t {

		_, ok := counter[string(v)]
		if ok {
			counter[string(v)] = counter[string(v)] + 1
		} else {
			counter[string(v)] = 1
		}
	}

	result := ""

	l := 0
	min_size := 9999999
	for i := 0; i < len(s); i++ {
		_, ok := counter[string(s[i])]
		if ok {
			counter[string(s[i])] = counter[string(s[i])] - 1

			for isMatched(counter) {
				size := i - l + 1
				if size < min_size {
					min_size = size
					result = string(s[l : i+1])
				}

				_, ok := counter[string(s[l])]
				if ok {
					counter[string(s[l])] = counter[string(s[l])] + 1
				}
				l++

			}

		}

	}

	return result
}

func isMatched(counter map[string]int) bool {
	for _, v := range counter {
		if v > 0 {
			return false
		}
	}
	return true
}
