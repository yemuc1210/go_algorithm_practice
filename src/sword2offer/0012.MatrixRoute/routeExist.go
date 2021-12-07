package offer12
/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

*/
//dfs
func exist(board [][]byte, word string) bool {
	// idx := 0   //迭代的深度   len(word)==idx 最后一次
	m,n := len(board),len(board[0])   //m行n列
	var search func(board [][]byte,i,j,dep int, word string)bool
	search = func(board [][]byte,i,j,dep int, word string)bool{
		//在board中dfs搜索
		if dep == len(word){       //递归到最后一个树
			return true
		}
		if i<0 || j<0 || i==len(board) || j==len(board[0]){
			return false   //对i j 的约束
		}
		//dfs遍历流程
		if board[i][j]==word[dep]{   //找到一个字符，递归进入下一个
			//需要还原，所以记录一下当前值
			tmp:=board[i][j]
			board[i][j]=' '
			if search(board,i,j+1,dep+1,word) || search(board,i+1,j,dep+1,word) ||
				search(board,i,j-1,dep+1,word) || search(board,i-1,j,dep+1,word){
					return true   //在其他位置找到
			}else{
				board[i][j]=tmp   //还原  回溯
			}
		}
		return  false
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if search(board,i,j,0,word){
				return true
			}
		}
	}
	return false
}

