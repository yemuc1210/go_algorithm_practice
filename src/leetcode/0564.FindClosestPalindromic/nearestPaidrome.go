package lt564

import (
	"fmt"
	// "math"
	"strconv"
	// "github.com/sbinet/go-python"
)

// 思路分析：

// 如果数组的字符串长度 == 1，数字n - 1
// 开头为1，9**9为一个候选答案        例：100000，答案为99999
// 开头为9, 10**01为一个候选答案       例：99999，答案为100001
// 如果本身对称，则把最中间的一个（或两个）位数减（如果0则加）一
// 	例：123321，答案为122221
//     例：120021，答案为121121
// 如果不对称：
//     把前半部分逆序替换掉后半部分       例：1223，答案为1221
//     把最中间的一个（或两个）位数加一   例：1283，答案为1331，而非1221
//     把最中间的一个（或两个）位数减一   例：1800，答案为1771，而非1881
// ---------------------
//寻找最近的回文数
//给定整数字符串，返回与它 最近 的回文整数，若不止一个，返回最小的
func nearestPalidrome(n string) string {
	// n 中间划分  分为a,b 两个数  a是奇数时
	// a a+1 a-1在左边构建回文，n长度为奇数时左边最后一个字符不能复制过去，取距离n最近的返回
	// 小于10的直接返回
	num,_ := strconv.Atoi(n)
	if num<10 {
		return strconv.Itoa(num-1)
	}
	// 10的幂，但会n-1
	if isPow(num) {
		return strconv.Itoa(num-1)
	}
	if num == 11 {
		return strconv.Itoa(9)
	}
	// 若干个9，返回n+2
	if isAllNine(n) {
		return strconv.Itoa(num+2)
	}
	// 一般情况

	s1,s2 := n[:(len(n)+1)/2],n[(len(n)+1)/2:]//1283 -> 12 83  
	// 构建回文 注意n长度为奇数的时候左边的最后一个字符不能复制过去。 所以需要用到s2长度   i+i[len(s2)-1::-1]
	t1 := s1 + reverseStr(s1)  // 12 21
	n1,_ := strconv.Atoi(s1)
	l1 := n1 - 1
	r1 := n1 + 1
	sl1 := strconv.Itoa(l1)
	sr1 := strconv.Itoa(r1)
	t2 := sl1 + reverseStr(sl1)
	t3 := sr1 + reverseStr(sr1)
	fmt.Println(t1,t2,t3,s2)
	return ""
}
func reverseStr(s string) string {
	n := len(s)
	bytes := []byte(s)
	for i:=0;i<n/2;i++{
		bytes[i],bytes[n-i-1] = bytes[n-i-1],bytes[i]
	}
	return string(bytes)
}
func isAllNine(n string) bool{
	for i:=0;i<len(n);i++{
		if n[i] != '9' {
			return false
		}
	}
	return true
}

func isPow(n int) bool {
	for n!=1 {
		if n%10!=0{
			return false
		}
		n = n/10
	}
	return n==1
}
func isPalidrome(num int) bool {
	a := 0
	for num!=0 {
		// t := num%10  // 最后一位
		a = a*10 + num%10
		num /= 10
	}
	return a==num
}

func min(a,b int) int {
	if a<b {return a}
	return b
}
func diffs(a,b int) int{
	if a<b {
		return b-a
	}
	return a-b
}