package offer46

import "strconv"

/*
给定一个数字，我们按照如下规则把它翻译为字符串：
0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法

组合题
dp
翻译规则：
	字符串第i位置 可以单独翻译
	如果第i-1和第i位组成数字在10-25之间，可以连起来翻译
f(i)表示以第i位结尾的前缀串翻译的方案数
f(i)=f(i-1)+f(i-2)  i-1>=0     10<=x<=25   i-2和i位一起翻译
f(-1)=0  f(1)=0

滚动数组压缩空间，因为只有三个变量
*/
func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1 //前三项
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0 //f(i-2)  f(i-1)  f(i)
		r += q            //f(i) 一定包含f(i-1)
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1] //【i-1,i]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}
