package main

import (
    "fmt"
)

func main() {
    var N int
    //包含多组样例
    for {
        n,err := fmt.Scanln(&N)
        if n==0 {
            break
        }
        if err!=nil {
            panic(err)
        }
        var snake = make([][]int,N)
        for i:=range snake {
            snake[i] = make([]int, N-i)
        }
        getOutputSnakeMatrix(N, &snake)
        PrintSnakeMatrix(snake)
    }
}

func getOutputSnakeMatrix(N int, snake *[][]int){
    curNum := 1
    curCnt := 0
    
    for curCnt < N {
        //当前处理第curCnt 斜行
        //下标位置[curCnt,0] - [0,curCnt]   i--;j++
        i,j := curCnt,0
        for i>=0 || j<=curCnt {
            (*snake)[i][j] = curNum
            curNum ++
            i --
            j ++
        }
        curCnt ++
    }
    
}

func PrintSnakeMatrix(snake [][]int) {
     //输出
    for i:=0;i<len(snake);i++{
        for j:=0;j<len(snake[i]);j++{
            if j==len(snake[i])-1 {
                fmt.Println(snake[i][j])
                continue
            }
            fmt.Printf("%d ", snake[i][j])
        }
    }
}