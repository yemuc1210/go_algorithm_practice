package offerS105
/*lt695
剑指 Offer II 105. 岛屿的最大面积
给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。

一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，
这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

 
*/

//图的题  深度有限搜索
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] == 1{
				maxArea = max(maxArea, dfs(grid,i,j))
			}
		}
	}
	return maxArea
}

//从（i，j)开始搜索
func dfs(grid [][]int,  i,j int) int{
	//边界判断
	if i<0 || j <0 || i>=len(grid)||j>=len(grid[0]) || grid[i][j]==0{
		return 0  //要从土地块开始遍历
	}

	//遍历
	area := 1  //现在，起码 当前块（i，j)是土地
	grid[i][j] = 0  //避免重复遍历回来  反正本地只关注最大面积 对原图修改无妨
	area += dfs(grid, i+1,j)
	area += dfs(grid, i-1, j)
	area += dfs(grid, i, j+1)
	area += dfs(grid, i, j-1)
	return area
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}