package main

import (
    "fmt"
    "github.com/0xff3232/goleetcode/db/transposeMatrix"  // Assume this is the correct path
)

func main() {
    fmt.Println("Solution I'm currently working on..")

    var x [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

    t := transposeMatrix.TransposeMatrix(x)
    fmt.Println("Original Matrix:")
    fmt.Println(x)
    fmt.Println("Transposed Matrix:")
    fmt.Println(t)
}
