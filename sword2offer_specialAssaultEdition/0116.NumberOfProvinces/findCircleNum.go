package offerS116
/*lt547  图
剑指 Offer II 116. 省份数量
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，
且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，
其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，
而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。
*/
/*
寻找图的连通子图  社区划分
BFS/DFS
*/
func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	res := 0
	//一次dfs就是一个省份
	for i:=0;i<len(isConnected);i++{
		if !visited[i] {
			dfs(isConnected,visited,i)
			//dfs结束，得到一个省份
			res++
		}
	}

	return res
}

func dfs(isConnected [][]int, visited []bool, i int) {
	// panic("unimplemented")
	//从i开始搜索 搜索与i有边的
	for j:=0;j<len(isConnected);j++{
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(isConnected,visited,j)  //顺着j继续往下
		}
	}
}
