package lt807

// import "fmt"

/*
lt218
807. 保持城市天际线
给你一座由 n x n 个街区组成的城市，每个街区都包含一座立方体建筑。
给你一个下标从 0 开始的 n x n 整数矩阵 grid ，其中 grid[r][c] 表示坐落于 r 行 c 列的建筑物的 高度 。

城市的 天际线 是从远处观察城市时，所有建筑物形成的外部轮廓。
从东、南、西、北四个主要方向观测到的 天际线 可能不同。

我们被允许为 任意数量的建筑物 的高度增加 任意增量（不同建筑物的增量可能不同） 。
高度为 0 的建筑物的高度也可以增加。
然而，增加的建筑物高度 不能影响 从任何主要方向观察城市得到的 天际线 。

在 不改变 从任何主要方向观测到的城市 天际线 的前提下，返回建筑物可以增加的 最大高度增量总和 。
*/
/*
南北视角相反   东西也是
增加高度后，不能影响各个角度的天际线   其实也就两个方向，选择西和南

则可以找出每行每列的最大值，某个数的最大允许的高度，同时受行列的最大值约束
*/
func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)  //n*n
	//row最大值对应每行的最大值 西方视角     col最大值每列 南方
	rowMaxHeight,colMaxHeight := make([]int,n),make([]int,n)
	//初始化
	for i:=0;i<n;i++{
		rowMaxHeight[i] = grid[i][0]
		colMaxHeight[i] = grid[0][i]
	}
	for i:=0;i<n;i++{
		//找到每行的最大值
		for j:=0;j<n;j++{
			if grid[i][j] > rowMaxHeight[i] {
				rowMaxHeight[i] = grid[i][j]
			}
			if grid[i][j] > colMaxHeight[j] {
				colMaxHeight[j] = grid[i][j]
			}
		}
	}
	// fmt.Println(rowMaxHeight)
	// fmt.Println(colMaxHeight)
	//测试 上面求正确
	res := 0
	//下面计算最大增量 某个格子的最大增量由行/列最大值同时约束
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			//求当前元素的最大增量
			res +=  myAbs(min(rowMaxHeight[i],  colMaxHeight[j]), grid[i][j])
		}
	}
	return res
}
func myAbs(a,b int) int {
	if a>b {
		return a-b
	}
	return b-a
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}
