package main

import (
	"fmt"
	"math"
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

    test1 := findPoints(coords1)
    test2 := findPoints(coords2)

    fmt.Println(test1)
    fmt.Println(test2)
}
func findPoints(lx [][]int) int {
    sec := 0
    N := len(lx) - 1
    fmt.Printf("length of array: %d\n", N)
    for i := 0; i < N; i++{
        x1, y1 := lx[i][0], lx[i][1]
        x2, y2 := lx[i+1][0], lx[i+1][1]
        dx := int(math.Abs(float64(x2 - x1)))
        dy := int(math.Abs(float64(y2 - y1)))
        dis := int(math.Max(float64(dx), float64(dy)))
        sec += dis
    } 
    return sec 
}
