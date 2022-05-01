package main

import "fmt"

// 强整数：x^i+y^j 则认为该整数的一个强整数
// 给定x y bound 返回小于等于bound的所有强整数
func powerfulIntegers(x int, y int, bound int) []int {
	// x 的阶乘   y的阶乘   两个阶乘列表的和的组合
	// x y 大于0
	m := make(map[int]struct{})
	for a:=1;a<bound;a*=x {  // a 是x的阶乘
		for b:=1;a+b<=bound;b*=y { // b 是y的阶乘
			m[a+b] = struct{}{}
			// 若 y == 1,因为它的任何次幂都是 1，
			// 那么内层循环 a 只需要与 它 计算一次即可
			if y==1 {
				break
			}
		}
		//  x == 1 同理，只需要外层只需要计算一次即可，所有可能都在内层进行计算了
		if x==1 {
			break
		}
	}
	res := []int{}
	for k:=range m {
		res = append(res, k)
	}
	return res
}

func main() {
	x := 2
	y := 3
	bound := 10
	res := powerfulIntegers(x,y,bound)
	fmt.Println(res)
}