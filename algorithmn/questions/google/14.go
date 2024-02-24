package google

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := ""

	for index := 0; index <= len(strs[0])-1; index++ {
		isBroken := false
		current := strs[0][index]
		for i := 1; i <= len(strs)-1; i++ {
			if index > len(strs[i])-1 {
				isBroken = true
				break
			}
			if strs[i][index] != current {
				isBroken = true
				break
			}
		}
		if isBroken {
			return result
		} else {
			//result = strs[0][0:index+1]
			result = result + string(current)
		}
	}
	return result
}
