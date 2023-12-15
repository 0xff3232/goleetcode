package largestNumAtLeastTwiceOfOthers
import "fmt"

func LargestNumTwice(nums []int) int {

    m, i := Max(nums) 
    
    fmt.Println("largest of nums: ", m)

    for _, n := range nums {   
        // ensure m doesn't compare itself.
        if m < n*2 && m != n {
            return -1
        }
    } 
    
    return i  // i is index from Max
}

func Max(nums []int) (int, int) {
    largest := nums[0]
    largestIndex := 0

    for i, num := range nums{
        if num > largest {
            largest = num
            largestIndex = i
        }
    } 
    return largest, largestIndex
}