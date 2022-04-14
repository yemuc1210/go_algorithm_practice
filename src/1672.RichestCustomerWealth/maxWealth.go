package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	// accounts[i][j]: 用户i 在 银行j的资产
	// 求最富有客户的资产    对行求和
	// 加法操作 x^y
	// 进位操作  (x&y)<<1
	max := accounts[0][0]
	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			accounts[i][0] = add(accounts[i][0], accounts[i][j])
		}
		if accounts[i][0] > max {
			max = accounts[i][0]
		}
	}
	return max
}

func add(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	var cur int
	var carrier int
	for {
		cur = a ^ b            // 不进位加法
		carrier = (a & b) << 1 // 进位
		if carrier == 0 {
			break
		}
		// 继续
		a = cur
		b = carrier
	}
	return cur
}

func main() {
	fmt.Println("vim-go")
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(maximumWealth(accounts))
}
