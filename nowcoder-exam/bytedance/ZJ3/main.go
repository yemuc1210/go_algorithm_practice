package main

import (
    "fmt"
)

func main(){
    var n,m int
    var s string
    fmt.Scanln(&n,&m)    //长度n  操作次数上限m
    fmt.Scanln(&s)
    //滑动窗口
    length := m  //初始长度为m，
    var num_a,num_b int
    for i:=0;i<m;i++{
        if s[i] == 'a' {
            num_a ++
        }else if s[i] =='b' {
            num_b ++
        }
    }
    startIdx := 0  
    for i:=m;i<n;i++{
        if s[i] == 'a' {
            num_a ++
        }else{
            num_b ++
        }
        if Min(num_a,num_b) <= m {
            //m个替换操作 可以将这段s[startIdx,i]替换成一种
            length = Max(length, i-startIdx+1)  //更新
        }else if s[startIdx] == 'a' {
            //a b 的个数都超过m
            //m个替换操作不足以替换成一种
            num_a -- 
            startIdx ++
        }else{
            num_b --
            startIdx ++
        }
    }
    
    fmt.Println(length)
}

func Min(a,b int) int {
    if a<b {
        return a
        
    }
    return b
}

func Max(a,b int) int {
    if a>b {
        return a
    }
    return b
}