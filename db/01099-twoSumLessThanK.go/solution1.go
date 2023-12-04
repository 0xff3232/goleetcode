package main

import (
	"fmt"
    "sort"
)

func main() {
    var nums1 = []int{ 34,23,1,24,75,33,54,8 }
    k1 := 60

    var nums2 = []int{10, 20, 30}
    k2 := 15
    
    t1 := twoSumLessThanK(nums1, k1)
    t2 := twoSumLessThanK(nums2, k2)
    fmt.Println(t1)
    fmt.Println(t2)
 
}

func twoSumLessThanK(nums []int , k int) int {

    sort.Ints(nums)
    less := -1
    lp, rp := 0, len(nums) - 1
    
    
    for lp < rp {
        sum := nums[lp] + nums[rp]
   
         if sum < k {
            if sum > less{
                less = sum
            }    
            lp++
        } else {
            rp--
        }
    }
   
    return less
}
