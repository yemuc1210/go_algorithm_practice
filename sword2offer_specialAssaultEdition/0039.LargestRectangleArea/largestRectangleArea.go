package offerS39
func largestRectangleArea(heights []int) int {
	heights = append([]int{-1}, append(heights, 0)...)
	var res int
	var stack []int
	for i, h := range heights {
		for len(stack) > 1 && h < heights[stack[len(stack)-1]] {
			w := i - stack[len(stack)-2] - 1
			res = max(res, heights[stack[len(stack)-1]]*w)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

