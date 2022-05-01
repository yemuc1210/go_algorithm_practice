package offerS40
func maximalRectangle(matrix []string) int {
	var res int
	m := len(matrix)
	if m == 0 {
		return res
	}
	n := len(matrix[0])
	if n == 0 {
		return res
	}
	left := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				left[j] = 0
			} else if i == 0 {
				left[j] = 1
			} else {
				left[j] ++
			}
		}
		stack := []int{-1}
		for j, l := range left {
			for len(stack) > 1 && l < left[stack[len(stack)-1]] {
				w := j - stack[len(stack)-2] - 1
				res = max(res, left[stack[len(stack)-1]]*w)
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, j)
		}
		for len(stack) > 1 {
			w := n - stack[len(stack)-2] - 1
			res = max(res, left[stack[len(stack)-1]]*w)
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

