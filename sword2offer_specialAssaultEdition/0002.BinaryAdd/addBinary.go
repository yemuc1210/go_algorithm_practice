package offerS2

import "math/big"

/*
剑指 Offer II 002. 二进制加法
给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。

输入为 非空 字符串且只包含数字 1 和 0。
*/
func addBinary(a string, b string) string {
	ai, _ := new(big.Int).SetString(a,2)
	bo, _ := new(big.Int).SetString(b,2)
	res := ai.Add(ai,bo)
	return res.Text(2)

}