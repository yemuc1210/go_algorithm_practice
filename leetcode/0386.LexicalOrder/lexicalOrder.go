/*
 * @Author: ye fei
 * @Date: 2022-04-18 09:54:10
 * @LastEditTime: 2022-04-18 09:59:10
 * @FilePath: /go_practice1/0386.LexicalOrder/lexicalOrder.go
 * @Description:leetcode 386 字典序排数；要求时间复杂度O(N)，空间复杂度O(1)
 */
package main

import (
	"fmt"
)

func lexicalOrder(n int) []int {
	// dfs的感觉
	// 尝试给数字增加0-9 然后回溯
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		// num后面加数字
		if num*10 <= n {
			num *= 10
		} else {
			// 添加1-9
			for num%10 == 9 || num+1 > n {
				// 回溯
				num /= 10
			}
			// 否则++
			num++ //添加1-9
		}
	}
	return ans
}

func main() {
	n := 13
	ans := lexicalOrder(n)
	fmt.Println(ans)
}
