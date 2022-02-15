package lt1380

func luckyNumbers(matrix [][]int) (ans []int) {
    minRow := make([]int, len(matrix))
    maxCol := make([]int, len(matrix[0]))
    for i, row := range matrix {
        minRow[i] = row[0]
        for j, x := range row {
            minRow[i] = min(minRow[i], x)
            maxCol[j] = max(maxCol[j], x)
        }
    }
    for i, row := range matrix {
        for j, x := range row {
            if x == minRow[i] && x == maxCol[j] {
                ans = append(ans, x)
            }
        }
    }
    return
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
