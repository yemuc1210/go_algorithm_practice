package main

import (
    "fmt"
)

func main(){
    var y, m, d int
    for {
        _, err := fmt.Scan(&y, &m, &d)
        if err != nil {
            break
        }
        res := OutPrint(y, m, d)
        fmt.Println(res)
    }
}

func OutPrint(y, m, d int) int {
    var res int
    dicD := []int{31,28,31,30,31,30,31,31,30,31,30,31}
    dicR := []int{31,29,31,30,31,30,31,31,30,31,30,31}
    if isRui(y) {
        for loop :=0; loop < m-1; loop++{
            res += dicR[loop]
        }
        return res + d
    }else{
        for loop :=0; loop < m-1; loop++{
            res += dicD[loop]
        }
        return res + d
    }
}

func isRui(y int) bool {
    if ( y%4 == 0 && y %100 !=0 )|| y% 400 ==0 {
        return true
    }
    return false
}
