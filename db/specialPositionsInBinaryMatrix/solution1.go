package specialPositionsInBinaryMatrix

import "fmt"

func SpecialPositions(mat [][]int) int {
    rows := len(mat)
    cols := len(mat[0])

    rowCount := make([]int, rows)
    colCount := make([]int, cols)

    for row := 0; row < rows; row++{  
        for col := 0; col < cols; col++{
            if mat[row][col] == 1{
                rowCount[row]++
                colCount[col]++
            }
        }
    }

    fmt.Println("Row counts:", rowCount)
    fmt.Println("Column counts:", colCount)

    ans := 0

    // Finding special positions
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if mat[row][col] == 1 && rowCount[row] == 1 && colCount[col] == 1 {
                ans++
            }
        }
    }

    return ans    
}