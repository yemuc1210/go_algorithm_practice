package main

import (
    "fmt"
    "strconv"
)

func main() {
    var m int
    for {
        n,_ := fmt.Scanln(&m)
        if n==0 {
            break
        }
        mid := m*m
        res := []int{}
        
        if mid%2 == 1 {
            // 奇数
            m -= 1
            half := m/2
            res = append(res, mid)
            for i:=1;i<=half;i++{
                res = append(res, mid+2*i)
            }
            for i:=1;i<=half;i++{
                res = append([]int{mid-2*i},res...)
            }
        }else{ //偶数
            half := m/2
            for i:=0;i<half;i++{
                res = append(res, mid+2*i+1)
            }
            for i:=0;i<half;i++{
                res = append([]int{mid-2*i-1},res...)
            }
        }
        s := ""
        for i:=0;i<len(res);i++{
            if i==len(res)-1{
                s += strconv.Itoa(res[i])
            }else{
                s += strconv.Itoa(res[i]) +"+"
            }
        }
        fmt.Println(s)
    }

}