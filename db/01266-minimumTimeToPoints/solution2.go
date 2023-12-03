package main

import (
	"fmt"
)

func main() {
    var coords1 = [][]int{
        {1, 1},
        {3, 4},
        {-1, 0},
    }
    
    var coords2 = [][]int{
        {1, 1},
        {3, 4},
    }

    test1 := findpoints(coords1)
    test2 := findpoints(coords2)

    fmt.Println(test1)
    fmt.Println(test2)  
}
func findpoints(p [][]int) (sum int) {
   for i := 0; i < len(p)- 1; i++{
      sum += max(abs(p[i+1][0] - p[i][0]), abs(p[i+1][1] - p[i][1])) 
   }
    return sum
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func max(a, b int) int {
    if a > b {
        return a 
    }
    return b 
}