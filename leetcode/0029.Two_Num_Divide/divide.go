/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，
要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8
以及 truncate(-2.7335) = -2

*/
package leetcode

import "math"

func divide(dividend int, divisor int) int {
	//位运算？？
	if dividend == math.MinInt32 && divisor==-1{
		return math.MaxInt32
	}

	res := 0
	sign := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}
	
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor <0 {
		divisor = -divisor
	}

	for dividend>=divisor{
		tmp:=divisor
		m:=1
		for tmp<<1 <=dividend{
			tmp<<=1
			m<<=1
		}
		dividend -= tmp
		res += m
	}
	return sign*res
	
}
