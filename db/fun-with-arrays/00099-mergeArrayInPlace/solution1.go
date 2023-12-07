package main

import (
	"fmt"
)

// this solution wasn't correct, I need to modify the array in place.
func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	m := 3
	var nums2 = []int{2, 5, 6}
	n := 3

	t := Merge(nums1, m, nums2, n)
	fmt.Println(t)
}

func Merge(nums1 []int, m int, nums2 []int, n int) (merged []int) {
	merged = nums1[:m]
	var toMerge []int = nums2[:n]
	merged = append(merged, toMerge...)
    BubbleSort(merged)

	return merged
}

func BubbleSort(arr []int) []int {

    N := len(arr)
    for i := 0; i < N-1; i++ {
        for j := 0; j < N-i-1; j++ {

            if arr[j] > arr[j+1]{
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
   return arr 
}
