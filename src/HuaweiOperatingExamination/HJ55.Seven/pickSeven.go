package main

import "fmt"

func main(){
	var N int
	for {
		n,_ := fmt.Scanln(&N)
		if n== 0 {
			break
		}
		//输出不大于n的与7相关的数字的个数
		fmt.Println(solve(N))
	}

}

func solve(n int) int {
	cnt := 0
	vis := make([]int,n+1) 
	for i:=1;i<=n ;i++{
		if vis[i] == 0 {
			vis[i] = 1
			if hasDigit(i,7) || isMultiple(i,7) {
				cnt ++
				fmt.Println(i, isMultiple(i,7))
				if isMultiple(i,7) {
					//将其倍数也访问
					for j:=i;j<=n;j+=7{
						if vis[j]==0 {
							vis[j] = 1
							cnt ++
							fmt.Println(j)
						}
						
					}
				}
			}
		}
	}
	return cnt
}
// 是否有7
func hasDigit(num int, target int) bool {
	for num > 0 {
		a := num%10   //获得最后一位
		if a== target {
			return true
		}
		num = num/10  // 去掉最后一位
	}
	return false
}
func isMultiple(num int,target int) bool {
	return num%target == 0
}