package main

import (
    "fmt"
)

func main() {
    var num int32
    for {
        n,_ := fmt.Scanln(&num)
        if n==0 {
            break
        }
        // 计算num二进制中1的个数
        var res int32
        for num!=0{
            res += num&1
            num >>= 1
        }
        fmt.Println(res)       
    }
}