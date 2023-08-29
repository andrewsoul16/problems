package arrays

func productExceptSelf(nums []int) []int {
	var (
		res  = make([]int, len(nums))
		mult int
	)

	for i := 0; i < len(nums); i++ {
		mult *= nums[i]
	}

	return res
}
