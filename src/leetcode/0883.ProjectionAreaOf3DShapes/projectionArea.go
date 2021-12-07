package lt883

/*
883. 三维形体投影面积
在 N * N 的网格中，我们放置了一些与 x，y，z 三轴对齐的 1 * 1 * 1 立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。

现在，我们查看这些立方体在 xy、yz 和 zx 平面上的投影。

投影就像影子，将三维形体映射到一个二维平面上。

在这里，从顶部、前面和侧面看立方体时，我们会看到“影子”。

返回所有三个投影的总面积。
*/
/*
正视图：每一行最大值之和；
侧视图：每一列最大值之和；
俯视图：柱子个数；
*/
func projectionArea(grid [][]int) int {
	//定义三个变量 记录三个面的面积
	xy := 0 //  xy面 grid值>1则+1
	xz := 0   //正视图  +grid每一行最大值
	yz := 0 // 侧视图   gird每一列最大值
	col := make([]int, len(grid[0]))  //每一列最大值
	for i:=0;i<len(grid);i++{
		row := 0   //记录每一行最大值 一行一行遍历 所以不需要  row数组
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] > 0 {
				xy ++
			}
			row = max(row, grid[i][j])
			col[j] = max(col[j], grid[i][j])
		}
		//可以得到每一行的最大值
		xz += row
	}
	//处理侧视图 每一列的最大值
	for i:=0;i<len(col);i++{
		yz += col[i]
	}
	return xy+xz+yz
}
func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}