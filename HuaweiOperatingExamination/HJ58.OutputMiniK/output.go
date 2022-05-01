package main

import (
    "fmt"
    "sort"
)

func main() {
    var n,k int
    for {
        m,_ := fmt.Scanln(&n,&k)
        if m==0 {
            break
        }
        nums := make([]int,n)
        for i:=0;i<n;i++{
            fmt.Scanf("%d%c",&nums[i])
        }
		// fmt.Println(nums)
        sort.Ints(nums)
        for i:=0;i<k;i++{
            fmt.Printf("%d ",nums[i])
        }
		fmt.Println()
    }
}