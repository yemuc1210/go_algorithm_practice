package main

import (
	"fmt"
	"strconv"
)
// 6056 字符串中最大的 3 位相同数字
// 给你一个字符串 num ，表示一个大整数。如果一个整数满足下述所有条件，则认为该整数是一个 优质整数 ：

// 该整数是 num 的一个长度为 3 的 子字符串 。
// 该整数由唯一一个数字重复 3 次组成。
// 以字符串形式返回 最大的优质整数 。如果不存在满足要求的整数，则返回一个空字符串 "" 。

// 注意：

// 子字符串 是字符串中的一个连续字符序列。
// num 或优质整数中可能存在 前导零 。
func largestGoodInteger(num string) string {
	n := len(num)
	var maxNum int = -1
	var res string 
	for i:=0;i<n-2;{
		curNum,_ := strconv.Atoi(num[i:i+1])
		curStrNum,_ := strconv.Atoi(num[i:i+3])
		if curNum*100+curNum*10+curNum == curStrNum {
			if curStrNum > maxNum {
				maxNum = curStrNum
				res = num[i:i+3]
			}
			i += 3
			continue
		}
		i++
	}
	if maxNum == 0 {
		return "000"
	}
	return res
}

func main(){
	num := "222"
	res := largestGoodInteger(num)
	fmt.Println(res)
}