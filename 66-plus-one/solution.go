package leetcode66

func PlusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; 0 <= i; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		if carry == 0 {
			break
		}
		digits[i] = digits[i] % 10
	}
	if carry == 1 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}

var plusOne = PlusOne
