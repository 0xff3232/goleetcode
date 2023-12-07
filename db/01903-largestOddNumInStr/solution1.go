package main

import (
	"fmt"
)

func main() {
    t := "35426"

    fmt.Println(largestOddNumber(t))
}

func largestOddNumber(num string) string {
    odd := map[byte]bool{
        '1':true,
        '3':true,
        '5':true,
        '7':true,
        '9':true,
    }

    // loop from end until we find odd
    for i := len(num) - 1; i >= 0; i-- {
        if _, exists := odd[num[i]]; exists { 
            return num[:i+1] // if odd exists, return split + 1
        }
    }
    return ""
}
