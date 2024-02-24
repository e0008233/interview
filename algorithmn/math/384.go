package math

import "math/rand"

type Solution struct {
	Origin []int
}

func Constructor(nums []int) Solution {
	return Solution{Origin: nums}
}

func (this *Solution) Reset() []int {
	return this.Origin
}

func (this *Solution) Shuffle() []int {
	i, j := len(this.Origin), 0
	n := make([]int, len(this.Origin))
	copy(n, this.Origin)

	for 1 < i {
		j = rand.Intn(i)
		n[i-1], n[j] = n[j], n[i-1]
		i--
	}

	return n
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
