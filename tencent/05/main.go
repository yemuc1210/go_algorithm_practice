package main

import (
	"fmt"
	"math"
)

func main() {
	var n,m int
	fmt.Scanf("%d %d",&n,&m)
	// 第二行 n个整数
	a := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d%c",&a[i])
	}
	// n,m := 6,2
	// a := []int{2,3,1,1,1,2}
	tmpM := m  // 备份
	before := math.MaxInt32   
	count := 0  // 只有股票
	for i:=0;i<n;i++{
		if before < a[i] && count> 0 {
			m += a[i]   // 卖出 
			count--
		}else if before>a[i] && m >= a[i] {
			m -= a[i]   // 买入
			count++
		}else if before==a[i] && m>=a[i] {
			m -= a[i]   // 价格持平 能买则买
			count++
		}
		before = a[i]
	}
	res := m+count *a[n-1]
	if res < tmpM {
		fmt.Println(tmpM)
	}else{
		fmt.Println(res)
	}
}
