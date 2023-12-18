package maxDifBetweenTwoPairs

import (
	"sort"
)

func MaxProductDifference(nums []int) int {
	sort.Ints(nums)
	t := (nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1])
	return t
}
