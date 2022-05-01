package lt367

import "fmt"

/*
367. 有效的完全平方数
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，
则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如  sqrt 。

完全平方数：一个正整数是另一个整数的平方
（1）个位只能是0，1，4，5，6，9
（2）任何偶数的平方一定能被4整除
*/
func isPerfectSquare(num int) bool {
	if num == 0 {
		return true
	}
	for i:=1;i<=num;i++{
		fmt.Println(i,num,i*i)
		if i*i == num{
			return true
		}
	}
	return false
}