package main

import "fmt"

// 整数的数组形式的加法
// 返回num+k 的数组形式
func addToArrayForm(num []int, k int) []int {
	// 全部转为数组，然后数组逆序，从下标0出开始竖式运算
	// 若转为整数，可能会发生越界错误
	numK := []int{}
	for k != 0 {
		numK = append(numK, k%10)
		k /= 10
	} // 此时numK是逆序数组
	// 把 num逆序
	m := len(numK)
	n := len(num)
	for i := 0; i < n/2; i++ {
		num[i], num[n-i-1] = num[n-i-1], num[i]
	}
	// 准备相加
	if n < m {
		num, numK = numK, num
		n, m = m, n
	} // 保证num始终是较长的那个
	carry := 0
	var b int
	for i := 0; i < n; i++ {
		if i >= m { // 不存在的位置，用0代替
			b = 0
		} else {
			b = numK[i]
		}
		sum := (num[i] + b + carry)
		// 计算进位
		num[i] = sum % 10
		carry = sum / 10
	}
	// 仍有进位
	if carry != 0 {
		num = append(num, carry)
		n++
	}
	// 和逆置
	for i := 0; i < n/2; i++ {
		num[i], num[n-i-1] = num[n-i-1], num[i]
	}
	return num
}

func main() {
	num := []int{2, 1, 5}
	k := 806
	res := addToArrayForm(num, k)
	fmt.Println(res)
}
