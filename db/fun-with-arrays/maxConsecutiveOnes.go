package main

import (
	"fmt"
)

func maxConsecutiveOnes() {
    var nums = []int{1,0,1,1,1,0,1} 
    N := len(nums)
    
    var count int = 0
    var max_int int = 0

    for i := 0; i < N; i++ {
        if nums[i] == 1 {
            count += 1
            if count > max_int {
                max_int = count
            }
        } else if nums[i] == 0 {
            count = 0
        }
    }
    
    fmt.Println("------TEST------")
    fmt.Printf("max_int value: %d\n", max_int)
    fmt.Println("value should be: 3")
    fmt.Println("------END------")
}

