package two_pointers

import (
	"fmt"
	"slices"
	"sort"
)

type triplets [][]int

func newTriplets() triplets {
	return make([][]int, 0)
}

func newTriplet(a, b, c int) []int {
	return []int{a, b, c}
}

func (t *triplets) push(triplet []int) {
	*t = append(*t, triplet)
}

func (t triplets) haveTripletAlready(triplet []int) bool {
	sort.Ints(triplet)
	for _, v := range t {
		if slices.Equal(v, triplet) {
			return true
		}
	}
	return false
}

func check(trips triplets, nums []int, i, j, k int) {
	if i == 1 && j == 4 && k == 5 {
		fmt.Println("dfgldfg")
	}
	if k == i || k == j || i == j {
		return
	}
	if nums[i]+nums[j]+nums[k] != 0 {
		return
	}
	t := newTriplet(nums[i], nums[j], nums[k])
	if trips.haveTripletAlready(t) {
		return
	}
	trips.push(t)
}

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var trips = newTriplets()
	for i := 0; i < len(nums); i++ {
		// пропускаем элемент для первой позиции, если он уже был, так как мы его в этом случае уже проверили
		if i > 0 && nums[i] == nums[i]-1 {
			continue
		}
		target := nums[i]
		m := map[int]int{}
		for j := i + 1; j < len(nums)-1; j++ {
			k, ok := m[nums[j]]
			if !ok {
				m[target-nums[j]] = j
				continue
			}
			trips.push([]int{nums[i], nums[j], nums[k]})
		}
	}
	return trips
}
