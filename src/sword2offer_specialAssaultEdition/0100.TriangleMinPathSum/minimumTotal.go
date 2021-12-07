package offerS100
/*lt120  dp
剑指 Offer II 100. 三角形中最小路径之和
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 
在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

还是向下 向右下  
贪心、dp
*/

/*
从下往上推到
*/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0 //nil
	}

	//从下网上推到
	for row:=len(triangle)-2; row>=0;row--{//行--
		for col :=0;col<len(triangle[row]);col++{//列++
			triangle[row][col] += min(triangle[row+1][col],   triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

func min(a,b int) int {
	if a<b{
		return a
	}
	return b
}