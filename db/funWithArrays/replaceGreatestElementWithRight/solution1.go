package replaceGreatestElementWithRight


func ReplaceGreatestElementWithRight(nums []int) []int {
	index := 0
	for index < len(nums)-1 {
		max := Max(nums[index+1:])
		nums[index] = max
		index++
	}
	nums[len(nums)-1] = -1
	return nums
}

func Max(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
