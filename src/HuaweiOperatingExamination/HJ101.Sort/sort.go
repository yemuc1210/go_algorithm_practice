package main
import (
    "fmt"
    "sort"
)
func main() {
    var num,isUpper int
    fmt.Scanln(&num)
    nums := make([]int,num)
    for i:=0;i<num;i++{
        fmt.Scanf("%d%c",&nums[i])
    }
    fmt.Scanln(&isUpper)
    sort.Ints(nums)
    if isUpper==0 {
        for i:=0;i<num;i++{
            if i==num-1 {
                fmt.Printf("%d\n",nums[i])
                continue
            }
            fmt.Printf("%d ",nums[i])
        }
    }else{
        for i:=num-1;i>=0;i--{
            if i==0 {
                fmt.Printf("%d\n",nums[i])
                continue
            }
            fmt.Printf("%d ",nums[i])
        }
    }
}