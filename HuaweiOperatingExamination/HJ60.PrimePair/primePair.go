package main

import (
    "fmt"
)

func main() {
    var num int
    for {
        n,_ := fmt.Scanln(&num)
        if n==0 {
            break
        }
        // 输出两个素数
        // 两个素质差值最小的素数对
        // 以输出为中心，向两边发散  双指针
        l,r := num/2,num/2
        for !isPrime(l) || !isPrime(r){
            l--
            r++
        }
        fmt.Println(l)
        fmt.Println(r)
    }
}

// 判断是否是素数
func isPrime(n int) bool {
    // 判断方法，从2-sqrt(n)
    for i:=2;i*i<=n;i++{
        if n%i == 0 {
            //能被整除 则
            return false
        }
    }
    return true
}