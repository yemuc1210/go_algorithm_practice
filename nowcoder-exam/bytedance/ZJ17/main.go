package main

import (
	"fmt"
	// "io"
)

//水仙花数
//各位数字的立方和等于其本身   穷举
func main() {
	var m, n int
	num, _ := fmt.Scanf("%d %d", &m, &n)

	for num != 0 {

		//输出给定范围内的水仙花数  100<=m<=n<=999
		//穷举法
		curRes := []int{}
		for i := m; i <= n; i++ {
			// fmt.Println(i/100,i/10%10,i%100%10)
			if solve(i/100, i/10%10, i%100%10, i) {
				curRes = append(curRes, i)
			}
		}
		if len(curRes) == 0 {
			fmt.Println("no")
		} else {
			for i := range curRes {
				if i == len(curRes)-1 {
					fmt.Printf("%d\n", curRes[i])
				} else {
					fmt.Printf("%d ", curRes[i])
				}
			}
			// fmt.Println(curRes)
		}
		num, _ = fmt.Scanf("%d %d", &m, &n)

	}
}

func solve(a, b, c int, num int) bool {
	//a b c分别是百 十 个位
	return a*a*a+b*b*b+c*c*c == num
}
