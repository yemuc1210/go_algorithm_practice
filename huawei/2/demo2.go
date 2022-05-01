
package main

import (
	"fmt"
	"math"
)

func main() {
	// please define the input here.
	// For example: r := bufio.NewReader(os.Stdin) input, err := r.ReadString("\n")
	// please finish the function body here.
	// please define the output here. For example: fmt.Println(input)
	// 假设是5*5的地图
	var rows,cols int
	fmt.Scanln(&rows,&cols)
	// 起点坐标
	var startX,startY int
	fmt.Scanln(&startX,&startY)
	// 终点坐标
	var endX,endY int
	fmt.Scanln(&endX,&endY)
	// 构建图
	graph := make([][]int, rows)
	for i:=range graph {
		graph[i] = make([]int,cols)
	}
	// 湖泊个数
	var poolNum int
	fmt.Scanln(&poolNum)
	// 湖泊坐标
	var pools [][]int =make([][]int,poolNum)
	for i:=range pools{
		pools[i] = make([]int,2)
	}
	var poolX,poolY int
	for i:=0;i<poolNum;i++{
		fmt.Scanln(&poolX,&poolY)
		pools[i][0] = poolX
		pools[i][1] = poolY
		graph[pools[i][0]][pools[i][1]] = 2
	}


	// 0-可走，1-起点 2-湖泊 3-终端
	// 求起点到终点得最短路径个数
	// dp[i][j]: 起点到（i，j）的最短路径长度
	// dfs 求路径
	if startX == endX && startY==endY {
		fmt.Printf("1 0")
		return
	}
	//var res [][]pair
	var dirs = [][]int{
		{-1,0},{1,0},{0,-1},{0,1},
	}

	visited := make([][]bool, rows)
	for i:=range visited {
		visited[i] = make([]bool ,cols)
	}
	var count int = 0
	minPathLen := math.MaxInt32

	var solve func(curX,curY int,sumLen int)
	solve = func(curX,curY int,sumLen int){

		// 越界
		if curX<0 || curX>=rows || curY<0 || curY>=cols || visited[curX][curY] || graph[curX][curY]==2 {
			return
		}
		if curX==endX && curY==endY {
			// 得到一条路径
			if sumLen < minPathLen {
				minPathLen = sumLen
				count = 1
			}else if sumLen == minPathLen {
				count++
			}
			return
		}
		// 判断是否可达
		// 四个方向
		visited[curX][curY] = true
		for _,dir:= range dirs{
			//nextX,nextY := curX+dir[0],curY+dir[1]
			solve(curX+dir[0], curY+dir[1], sumLen+1)
		}
		visited[curX][curY] = false
	}
	solve(startX,startY,0)
	fmt.Printf("%d %d",count,minPathLen)
}

