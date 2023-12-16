package spiralMatrix

import "fmt"

func SpiralMatrix(mat [][]int)[]int{
    rows := len(mat)
    cols := len(mat[0]) 
    ans := []int{}
    
    // midPoint := FindMid(mat) ?

    for r := 0; r < rows; r++ {
        fmt.Println(mat[r])
        for c := 0; c < cols; c++ {     
            fmt.Println(mat[r][c])
            // iterate till end of row
            // then iterate till end of col until mid or /2 > right of

            // mat[1][2] -> mid
            
            // mat[0][0-2] [1,2,3]   +3    
            // mat[1-2][2] [6,9]     +2
            // mat[2][2-1] [8, 7]    -2
            // mat[1][1]   [4]       -1
            // mat[1][2]   [5]       +1 (if mid == mat[i][j]) end
        }
    } 
    return ans
}

