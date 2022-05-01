package tencent50
/*
59. 螺旋矩阵 II
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，
且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/
/* 螺旋 烦
1-n^2    所以正方形 
*/
func generateMatrix1(n int) [][]int {
	if n==0 {
		return [][]int{}
	}
	if n==1{
		return [][]int{{1}}
	}
	j:=0
	curNum := 1
	res := make([][]int,n)
	for i := range res {
		res[i] = make([]int, n)
	} //先初始化二维数组   然后根据下标生成对应位置数字
	for curNum <= n*n{  
		//j  轮次
		for i := j; i < n - j; i++{
			res[j][i] = curNum
			curNum ++
		}
        for i:=j+1;i<n-j;i++{
			res[i][n-j-1] = curNum
			curNum++
		}

        for i:=n-j-2;i>=j;i--{
			res[n-j-1][i] = curNum
			curNum++
		}
		for i:=n-j-2;i>j;i--{
			res[i][j] = curNum
			curNum++
		}
    	j++
	}
	return res
}

func generateMatrix(n int) [][]int{
	if n==0 {
		return [][]int{}
	}
	if n==1 {
		return [][]int{{1}}
	}
	res,visit,round,x,y,spDir := make([][]int,n),make([][]int,n),0,0,0,[][]int{
		{0,1}, //右
		{1,0}, //下
		{0,-1}, //left
		{-1,0},//up
	}
	//初始化 visit
	for i:=0;i<n;i++{
		visit[i] = make([]int, n)
		res[i] = make([]int, n)
	}
	visit[x][y] = 1
	res[x][y] = 1
	//迭代
	for i:=0;i<n*n;i++{
		x += spDir[round%4][0]
		y += spDir[round%4][1]  //x y表示方向
		if (x==0 && y==n-1) || (x==n-1 && y==n-1) || (y==0 && x==n-1) {
			round ++ //换方向
		}
		if x>n-1 || y>n-1 || x<0 || y<0 {
			return res  //break
		}
		if visit[x][y]==0{
			visit[x][y] = 1
			res[x][y] = i+2  //
		}
		switch round%4 {
		case 0:
			if y+1 <= n-1 && visit[x][y+1]==1 {
				round ++
				continue
			}
		case 1:
			if x+1 <= n-1 && visit[x+1][y] == 1 {
					round++
					continue
			}
		case 2:
			if y-1 >= 0 && visit[x][y-1] == 1 {
					round++
					continue
			}
		case 3:
			if x-1 >= 0 && visit[x-1][y] == 1 {
					round++
					continue
			}
		}
	}
	return res
}