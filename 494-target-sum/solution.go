package leetcode494

// FindTargetSumWays solves LeetCode 494: Target Sum
// Runs out of time -- O(2^n)
func SlowExponentialFindTargetSumWays(nums []int, S int) int {
	ways := 0

	var i, j, n uint32
	n = uint32(len(nums))
	for i = 0; i < 1<<n; i++ {
		sum := 0
		for j = 0; j < n; j++ {
			if i&(1<<j) > 0 {
				sum += nums[j]
			} else {
				sum -= nums[j]
			}
		}
		if sum == S {
			ways++
		}
	}
	return ways
}

// FindTargetSumWays solves LeetCode 494: Target Sum
func FindTargetSumWays(nums []int, S int) int {
	if len(nums) == 1 {
		ways := 0
		if S == nums[0] {
			ways++
		}
		if S == -nums[0] {
			ways++
		}
		return ways
	}
	return FindTargetSumWays(nums[1:], S-nums[0]) + FindTargetSumWays(nums[1:], S+nums[0])
}

var findTargetSumWays = FindTargetSumWays
