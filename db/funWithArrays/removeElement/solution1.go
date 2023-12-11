package removeElement

import "fmt"

func RemoveElement(nums []int, val int) int { 

    left, right := 0, len(nums) - 1

    for left <= right {
        if nums[left] == val {
            nums[left], nums[right] = nums[right], nums[left]
            right--
        } else {  
            left++
        }
    }

    fmt.Println(nums)
    return left
}