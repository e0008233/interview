package google

import "math"

func divide(dividend int, divisor int) int {

	if dividend == math.MaxInt32 && divisor == -1 {
		return math.MinInt32 + 1
	}

	if divisor == 0 {
		return math.MaxInt32
	}
	flag := 1
	if dividend < 0 {
		flag *= -1
		dividend *= -1
	}
	if divisor < 0 {
		flag *= -1
		divisor *= -1
	}
	var multiple int
	for divisor <= dividend {
		t := divisor // t = divisor
		count := 1
		for t<<1 <= dividend {
			t <<= 1     // divisor * 2
			count <<= 1 // 倍数 * 2
		}
		multiple += count
		dividend -= t
	}
	if multiple >= math.MaxInt32 {
		if flag == 1 {
			return math.MaxInt32
		}
		return math.MinInt32
	}
	return multiple * flag
}
