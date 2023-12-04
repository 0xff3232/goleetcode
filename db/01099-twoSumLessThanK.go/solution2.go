// this solution was overkill, just wanted to see how to implement a quicksort and how it would effect it.
package main

import (
	"fmt"
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

func quickSort(arr []int) []int{
    if len(arr)<2{
        return arr
    }
    left, right := 0, len(arr)-1

    pivotIndex := len(arr)/2
    arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

    for i := range arr{
        if arr[i] < arr[right]{
            arr[i], arr[left] = arr[left], arr[i]
            left++
        }
    }

    arr[left], arr[right] = arr[right], arr[left]

    quickSort(arr[:left])
    quickSort(arr[left+1:])
    
    return arr
}

func twoSumLessThanK(nums []int, k int) int {

    quickSort(nums)
    less := -1
    lp, rp := 0, len(nums) - 1
    

    for lp < rp {
        sum := nums[lp] + nums[rp]

        if sum < k {
            if sum > less {
                less = sum
            }
            lp++
        } else {
            rp--
        }
    }

    return less 
}