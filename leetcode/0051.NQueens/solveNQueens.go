package leetcode

import "strings"

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，
该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

//回溯法
func solveNQueens(n int) [][]string {
	res:=[][]string{}
	board := make([][]string,n)
	for i:=range board{
		board[i]=make([]string, n)
		for j:=range board[i]{
			board[i][j]="."
		}
	}  //构建棋盘
	solve(0,board,&res,n)
	//下面求解
	return res
}
func solve(r int,board [][]string,res *[][]string,n int){
	if r==n {
		tmp:=make([]string,len(board))
		for i:=0;i<n;i++{
			tmp[i]=strings.Join(board[i],"")
		}
		*res = append(*res, tmp)
		return
	}
	for c:=0;c<n;c++{
		if isValid(r,c,n,board){
			board[r][c]="Q"
			solve(r+1,board,res,n)
			board[r][c]="."  //否则回复状态
		}
	}
}

func isValid(r,c int, n int, board [][]string)bool{
	for i := 0; i < r; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}