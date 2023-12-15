package pivotIndex

import "fmt"


func PivotIndex(nums []int) int {

    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Printf("total of nums is : %d\n", total)

    left := 0
    for i, n := range nums {
        fmt.Println("till == : ", total - left - n)
       if left  == total - left - n {
           return i
       }
        fmt.Println("current total of left: ", left)
        left += n
    }

    return -1 
}