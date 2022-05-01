package main

import (
	"fmt"
	"math"
)

/*
给定一个整数 n ，返回 可表示为两个 n 位整数乘积的 最大回文整数 。
因为答案可能非常大，所以返回它对 1337 取余 。
*/

/*
枚举，从大到小枚举回文数，枚举左半部分即可
两个n位数乘积 至多 2n位   从10^n-1开始枚举左半部分
得到回文数p，然后判断是否可以分解成两个n位数
	从10^n-1从大到小（sqrt(p)）枚举x
*/


func largestPalindrome(n int) int {
	if n==1 {
		return 9
	}
	upper := int(math.Pow10(n)-1)
	//枚举左半部分
	for left:=upper;;left-- {
		p := left
		// 构建回文数
		for x:=left;x>0;x/=10 {
			// left   left 最后一位x%10 
			p = p*10 + x%10
		}
		// 判断能否分解
		for x:=upper;x*x>=p;x-- {
			if p%x==0 {
				//可以分解
				return p%1337
			}

		}
	}
}

func main() {
	n:=2
	ans:=987
	res := largestPalindrome(n)
	if res==ans{
		fmt.Println("pass")
	}
	fmt.Println(largestPalindrome(n))
}