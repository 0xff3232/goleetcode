package mergedZeros

func MergeZeros(nums []int) []int {
    i, j := 0, 0

    for j < len(nums) {
        if nums[j] != 0 {
            nums[i] = nums[j] // if not 0 nums[i] = nums[j] 
            i++ // this will be our len of nums that aren't 0
        }
        j++ // move up every iteration
    }

    // while i < nums we add a 0 to the array
    // filling out the end of our array with 0s
    for i < len(nums) {
        nums[i] = 0
        i++
    }
    return nums
}