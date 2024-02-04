package main

import "fmt"

func calculateCore(s string) (count01, count10 int) {
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			count01 += countOnes(s[i+1:])
		} else if s[i] == '1' {
			count10 += countZeros(s[i+1:])
		}
	}
	return count01, count10
}

func countOnes(s string) int {
	count := 0
	for _, char := range s {
		if char == '1' {
			count++
		}
	}
	return count
}

func countZeros(s string) int {
	count := 0
	for _, char := range s {
		if char == '0' {
			count++
		}
	}
	return count
}

func main() {
	input := "10101" // Replace with your input string

	count01, count10 := calculateCore(input)
	fmt.Println("Count of '01' patterns:", count01)
	fmt.Println("Count of '10' patterns:", count10)
}
