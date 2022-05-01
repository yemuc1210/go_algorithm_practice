package lt84
/*困难
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10   5、6合并
*/

/*
dp的感觉
dp[0] heights[0]
dp[1] = max ( min{h[0],h[1]}*2, dp[0] )
然而 随着宽度的增加，需要min max比较的规模就越大了
用二维dp?
	dp[i][j]  i-j直接合并的最大值 

单调栈
用单调栈依次保存直方图的高度下标，一旦出现高度比栈顶元素小的情况就取出栈顶元素，
单独计算一下这个栈顶元素的矩形的高度。
然后停在这里(外层循环中的 i--，再 ++，就相当于停在这里了)，
继续取出当前最大栈顶的前一个元素，即连续弹出 2 个最大的，
以稍小的一个作为矩形的边，宽就是 2 计算面积…………
如果停在这里的下标代表的高度一直比栈里面的元素小，就一直弹出，
取出最后一个比当前下标大的高度作为矩形的边。
宽就是最后一个比当前下标大的高度和当前下标 i 的差值。
计算出面积以后不断的更新 maxArea 即可。
*/
func largestRectangleArea(heights []int) int {
	maxArea, stack, height := 0, []int{}, 0
	for i := 0; i <= len(heights); i++ {
		if i == len(heights) {
			height = 0
		} else {
			height = heights[i]
		}
		if len(stack) == 0 || height >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			tmp := stack[len(stack)-1]
			// fmt.Printf("1. tmp = %v stack = %v\n", tmp, stack)
			stack = stack[:len(stack)-1]
			length := 0
			if len(stack) == 0 {
				length = i
			} else {
				length = i - 1 - stack[len(stack)-1]
				// fmt.Printf("2. length = %v stack = %v i = %v\n", length, stack, i)
			}
			maxArea = max(maxArea, heights[tmp]*length)
			// fmt.Printf("3. maxArea = %v heights[tmp]*length = %v\n", maxArea, heights[tmp]*length)
			i--
		}
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
