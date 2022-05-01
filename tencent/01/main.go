package main

import (
    "fmt"
    "strconv"
    "sort"
)
func main() {
    var n int
    fmt.Scanln(&n)
    // 
    var str string
    fmt.Scanln(&str)
    var m = len(str)
    nums := make([]int,m)
    for i:=0;i<m;i++{
        num,_ := strconv.Atoi(str[i:i+1])
        nums[i] = num
    }
    for i:=1;i<n;i++{
        var str1 string
        fmt.Scanln(&str1)
        for j:=0;j<m;j++{
            num,_ := strconv.Atoi(str1[j:j+1])
            nums[j] = nums[j]*10 + num
        }
    }
    // 排序
    sort.Slice(nums, func(i,j int)bool {
        return nums[i]<nums[j]
    })
    for i:=0;i<m;i++{
        fmt.Printf("%d ",nums[i])
    }
}