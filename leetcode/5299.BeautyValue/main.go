package main

import (
	"fmt"
	"strconv"
)

// num 的 k美丽值 = 子串长度且能整除num
func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)
	if k > len(s) {
		return 0
	}
	bs := []byte(s) 
	n := len(s)
	var res int
	for i:=0;i<=n-k;i++{
		// 子串
		subStr := bs[i:i+k]
		subNum,_ := strconv.Atoi(string(subStr))
		// fmt.Println(i,n,res,num,string(subStr),subNum)
		if subNum == 0 {
			continue
		}
		if num%subNum == 0 {
			res ++
		} 
	}
	return res
}

func main() {
	res := divisorSubstrings(430043,2)
	fmt.Println(res)
}