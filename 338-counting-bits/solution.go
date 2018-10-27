package leetcode338

import "math/bits"

func countBitsEasy(num int) []int {
	counts := make([]int, num+1)
	for n := uint32(0); n <= uint32(num); n++ {
		counts[n] = bits.OnesCount32(n)
	}
	return counts
}

func countBits(num int) []int {
	counts := make([]int, num+1)
	for n := 0; n <= num; n++ {
		counts[n] = counts[n/2] + n%2
	}
	return counts
}
