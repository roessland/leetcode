package leetcode654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func split(nums []int) ([]int, int, []int) {
	if len(nums) == 0 {
		panic("must supply at least one number")
	}
	imax := 0
	max := nums[0]
	for i, num := range nums {
		if num > max {
			imax = i
			max = num
		}
	}

	return nums[0:imax], max, nums[imax+1 : len(nums)]
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	left, val, right := split(nums)
	return &TreeNode{val, constructMaximumBinaryTree(left), constructMaximumBinaryTree(right)}
}
