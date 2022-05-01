package lt371
/*
371. 两整数之和
给你两个整数 a 和 b ，不使用 运算符 + 和 - ​​​​​​​，计算并返回两整数之和。


位运算
	0+0=0
	0+1=1
	1+0=1
	1+1=0（进位）

	异或运算符合不进位的情况
	进位情况：   (a&b)<<1
*/
func getSum(a int, b int) int {
	for b!=0{
		carry:= uint(a&b) << 1   //进位
		a ^= b    //异或
		b = int(carry)
	}
	return a
}