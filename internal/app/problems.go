package app

import (
	"fmt"
	"problems/internal/arrays"
)

func Run() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Source array: ", nums)
	fmt.Println("Longest sequence: ", arrays.LongestConsecutive(nums))
}
