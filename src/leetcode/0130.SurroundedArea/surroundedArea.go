package lt130

/*
算法20天    dfs/bfs

给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
 如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。


算法：
	对边界上的0，从其开始搜索，标记与其直接相连或间接相邻的0
	最后遍历矩阵，对每个字母：
		如果被标记过，则该字母为没有被X包围的0，还原
		如果没有被标记，则属于被包围的0，修改为X
*/
var rows,cols int

func solve(board [][]byte)  {
	if board==nil || len(board)==0{
		return 
	}
	rows,cols = len(board),len(board[0])

	for i:=0;i<rows;i++{
		dfs(board,i,0) // i行0列   最左边一列
		dfs(board,i,cols-1) //最右边一列
	}

	//行
	for i:=1;i<cols-1;i++{ //因为第一个和最后一个在上面遍历过了
		dfs(board,0,i)  //最上面一行
		dfs(board,rows-1, i)
	}

	//处理
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			if board[i][j] == 'Y' {
				//说明被标记过，不是被包围的
				board[i][j] = 'O' //恢复
			}else if board[i][j] == 'O' {
				//未被访问过的，是被包围的
				board[i][j] = 'X'
			}
		}
	}
	
}

func dfs(board [][]byte, x,y int){
	//边界检查
	if x<0 || x>=rows ||y<0 || y>=cols ||board[x][y] != 'O'{
		return 
	}
	//只对'0'  错了是'O'
	board[x][y] = 'Y' //做个标记
	//对上下左右遍历
	dfs(board, x+1,y)
	dfs(board,x-1,y)
	dfs(board,x,y+1)
	dfs(board,x,y-1)
}