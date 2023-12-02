package main

import (
	"fmt"
)

func main() {
    var nums = []int{1,0,1,1,1,0,1} 
    N := len(nums)
    
    var count int = 0
    var max_int int = 0

    for i := 0; i < N; i++ {
        if nums[i] == 1 {
            count += 1
        } else if nums[i] == 0 {
            max_int = count
            count = 0
        }
    }
    fmt.Println(max_int)
}

