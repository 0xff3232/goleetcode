package evenNumberOfDigits

import (
	"fmt"
    "strconv"
)

func Solution(nums []int) int {
    even := 0
    fmt.Println(nums)

    for i := 0; i < len(nums); i++ {
        s := strconv.Itoa(nums[i])
        if len(s) % 2 == 0 {
            even++
        }
    }
    return even
}


    //var nums = []int {555, 901, 482, 1771}
    //var nums2 = []int {5555, 901, 482, 1771, 1212, 292, 9999}

    
