package leetcode

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
    // in case if all digits were 999 we should get 1000
	a := make([]int, n+1)
	a[0] = 1
	return a
}
