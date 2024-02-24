package google

func intToRoman(num int) string {
	numberRepresentation := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanRepresentation := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	index := 0
	result := ""
	for num > 0 {
		if num >= numberRepresentation[index] {
			num = num - numberRepresentation[index]
			result = result + romanRepresentation[index]
		} else {
			index++
		}
	}

	return result
}
