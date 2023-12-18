package spiralMatrix

func SpiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }

    var result []int
    /*
    [1,2,3]
    [4,5,6]
    [7,8,9]
    */

    // Defining boundaries for our spiral circle.
    // top, bottom = rows
    // left, right = cols
    top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

    for left <= right && top <= bottom { // case to stop and return ans

        // Move right
        // Since left is 0 and right is len(row) we iterate left <= right to get the whole first row
        for i := left; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
        // we += to top so that it moves the top down from row 0 to row 1 (happens after we go the whole row)
        top += 1 // [1,2,3]
                //..[5]

        // Move down
        // i will start at 1, because of code above, we then iterate from 1[right] which would be the last col.
        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        // Once we iterate col, we then minus right so it brings in that col to len(mat)-2
        right -= 1 // [6, 9]

        // These if statements (if top <= bottom and if left <= right),
        // are used to determine whether there are still rows or columns left to traverse in the respective directions.
        if top <= bottom {
            // Move left
            // go opposite way, use opposite logic.
            for i := right; i >= left; i-- {
                result = append(result, matrix[bottom][i])
            }
            bottom -= 1 // [8, 7]
        }

        if left <= right {
            // Move up
            // go opposite way, use opposite logic.
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left += 1 // [4]
        }
    } 
    return result
}