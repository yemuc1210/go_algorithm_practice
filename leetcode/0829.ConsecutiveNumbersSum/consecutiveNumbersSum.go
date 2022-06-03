package main

import "fmt"

// 连续整数求和
//   返回连续正整数满足和为n的组数
// n=5 输出2    ([5][2,3])
func consecutiveNumbersSum(n int) int {
	if n<=0 {
		return 0
	}
	res :=0
	// 判断当有i个连续数字时是否满足
	// 当有i+1个数字时，第i+1和第1个数字差为i
	// 所以每次循环减i.
	// 若能乘除i,则这些连续整数开始的元素是n/i
	for i:=1;n>0;i++ {
		if n%i==0 {
			res++
		}
		n -=i
	}
	return res
}

func main() {
	res := consecutiveNumbersSum(5)
	fmt.Println(res)
}