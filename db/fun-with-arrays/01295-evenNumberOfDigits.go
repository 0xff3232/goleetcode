package main

import (
	"fmt"
    "strconv"
)

func main() {
    var nums = []int {555, 901, 482, 1771}
    var nums2 = []int {5555, 901, 482, 1771, 1212, 292, 9999}

    even := Solution(nums)
    even2 := Solution(nums2)

          
    fmt.Println("------TEST------")
    fmt.Printf("even value is: %d\n", even)
    fmt.Println("Output should be: 1") 
    fmt.Println("------END------")   
  
    fmt.Println("------TEST------")
    fmt.Printf("even value is: %d\n", even2)
    fmt.Println("Output should be: 4") 
    fmt.Println("------END------")   

}

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
