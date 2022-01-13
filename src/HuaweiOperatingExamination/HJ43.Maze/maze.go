package main

import (
    "fmt"
)

func main() {
    var cols,rows int
    var n int
    var err error
    for {
        n,err = fmt.Scan(&cols,&rows)
		// fmt.Println(cols,rows)
        if n==0 {
            break
        }
        if err!=nil {
            panic(err)
        }
        var mazeMtrix = make([][]int,cols)
        for i:=range mazeMtrix {
            mazeMtrix[i] = make([]int,rows)
            for j:=0;j<rows;j++{
                fmt.Scan(&mazeMtrix[i][j])
            }
        }
        //求解
		var ans [][]int
        // dfs1(0,0,&ans, mazeMtrix, cols,rows)   //cols*rows
		dfs(mazeMtrix, 0, 0, cols, rows, &ans)
		length := len(ans)
		for i:=0;i<length;i++{
			// fmt.Println(ans[i])
			fmt.Printf("(%d,%d)\n", ans[i][0], ans[i][1])
		}
    }
}

func dfs1(i,j int, path *[]string, maze [][]int, m,n int) bool {
	if i==m-1 && j==n-1 {
		*path = append(*path, fmt.Sprintf("(%d,%d)",i,j))
		maze[i][j] = 3
		return true
	}
	//边界  m*n
	if i<0 || i>m-1 || j<0 || j> n-1 {
		return false
	}else{
		if maze[i][j] == 1 || maze[i][j] == 3{
			//不能访问
			return false
		}else{
			*path = append(*path, fmt.Sprintf("(%d,%d)",i,j))
			//打上访问标记
			maze[i][j] = 3
		}
	}

	if dfs1(i-1,j, path,maze, m,n) || dfs1(i+1,j,path,maze,m,m) || dfs1(i,j-1,path,maze,m,n) || dfs1(i,j+1,path,maze,m,n) {
		return true
	} else{
		//回溯上一个状态
		*path = (*path)[:len(*path)-1]
		maze[i][j] = 0
	}
	return false
}

func dfs(keep [][]int, i, j, n, m int, ans *[][]int) bool {
    // 终止条件
    if i == n-1 && j == m-1 {
        *ans = append(*ans, []int{i, j})
        keep[i][j] = 3
        return true
    }
    //边界条件
    if i < 0 || i > n-1 || j < 0 || j > m-1 {
        return false
    } else {
        if keep[i][j] == 1 || keep[i][j] == 3 {
            return false
        } else {
            *ans = append(*ans, []int{i, j})
            keep[i][j] = 3
        }
    }
 
    if dfs(keep, i-1, j, n, m, ans) || dfs(keep, i+1, j, n, m, ans) || dfs(keep, i, j-1, n, m, ans) || dfs(keep, i, j+1, n, m, ans) {
        return true
    } else {
        //回溯 回到上一个状态
        *ans = (*ans)[:len(*ans)-1]
            keep[i][j] = 0
    }
    return false
}