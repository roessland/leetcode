package leetcode53

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSubArray(nums []int) int {
	m := nums[0]
	maxM := m
	for i := 1; i < len(nums); i++ {
		m = Max(m+nums[i], nums[i])
		maxM = Max(maxM, m)
	}
	return maxM
}

var maxSubArray = MaxSubArray
