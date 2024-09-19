package arrays

import "math"

// реализация без приведения к строке
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	nums := make([]int, 0)
	var a, xx int
	xx = x
	for {
		a = xx % 10
		nums = append(nums, a)
		xx /= 10
		if xx == 0 {
			break
		}
	}
	if len(nums) == 0 {
		return false
	}
	newNum := nums[0] * int(math.Pow10(len(nums)-1)) // идем с последней цифры
	for i := 1; i < len(nums); i++ {
		newNum += nums[i] * int(math.Pow10(len(nums)-1-i))
	}

	return newNum == x
}
