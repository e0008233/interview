package google

func romanToInt(s string) int {
	romanIntMap := make(map[string]int)
	romanIntMap["M"] = 1000
	romanIntMap["CM"] = 900
	romanIntMap["D"] = 500
	romanIntMap["CD"] = 400
	romanIntMap["C"] = 100
	romanIntMap["XC"] = 90
	romanIntMap["L"] = 50
	romanIntMap["XL"] = 40
	romanIntMap["X"] = 10
	romanIntMap["IX"] = 9
	romanIntMap["V"] = 5
	romanIntMap["IV"] = 4
	romanIntMap["I"] = 1

	result := 0
	index := 0

	for index <= len(s)-1 {
		if index+1 <= len(s)-1 {
			v, ok := romanIntMap[s[index:index+2]]
			if ok {
				index = index + 2
				result = result + v
				continue
			}
		}
		v, ok := romanIntMap[s[index:index+1]]
		if ok {
			index = index + 1
			result = result + v
			continue
		}
	}

	return result
}
