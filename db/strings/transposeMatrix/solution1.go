package transposeMatrix

func TransposeMatrix(m [][]int) [][]int {
    if len(m) == 0 || len(m[0]) == 0 {
        return m
    }

    rows := len(m)
    cols := len(m[0])
    transposed := make([][]int, cols)

    for i := range transposed{
        transposed[i] = make([]int, rows)
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            transposed[j][i] = m[i][j]
        }
    }

    return transposed
}
