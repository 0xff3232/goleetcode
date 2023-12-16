package diagonalTraversal

import (
	"fmt"
	"sort"
)


func DiagonalTraversal(mat [][]int) []int {
    rows := len(mat)
    cols := len(mat[0])

    ans := make([]int, 0, rows*cols)
    diagonals := MakeDiagonal(mat)

    for k := 0; k <= rows+cols-2; k++ {
        if k%2==0 {
            for l:= len(diagonals[k]) - 1; l >= 0; l-- {
                ans = append(ans, diagonals[k][l])
            }
        } else {
            ans = append(ans, diagonals[k]...)     
        }
    }
    return ans
}

// Create a map of diagonals from a matrix.
func MakeDiagonal(mat [][]int) map[int][]int {
    rows := len(mat)
    cols := len(mat[0])
    diagonals := make(map[int][]int)
    
 
    for i := 0; i < rows; i++ { 
        for j := 0; j < cols; j++ {
            // i+j creates our diagonal to add.
            newDiagonal := i + j
            diagonals[newDiagonal] = append(diagonals[newDiagonal], mat[i][j])
        }  
    }

    PrintDiagonals(diagonals) 
    return diagonals
}

func PrintDiagonals(d map[int][]int) {
    fmt.Println()
    fmt.Println("### TEST-START ###")

    // print each diagonal and its values for testing, in order:
    keys := []int {} // to hold keys we sort.
    for k := range d {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for _, k := range keys {
        fmt.Printf("Diagonal: %d, %v\n", k, d[k])
    }
    // test end.
    fmt.Println("### END-TEST ###")
    fmt.Println()
}