package plusOne

func PlusOne(nums []int) []int {
	n := len(nums)

	for i := n - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}

	// If all the digits are 9, then the result will be 1 followed by n zeros.
	// Create a new slice with an additional length for the leading 1.
	result := make([]int, n+1)
	result[0] = 1

	return result
}
