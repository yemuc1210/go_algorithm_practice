package tencent50
/*简单
7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。
*/
func reverse(x int) int {
	//确定正负号
	sign := -1
	if x >= 0{
		sign = 1
	}
	x = abs(x)

	res := x % 10
	x = x/10
	for x > 0 {
		res = res * 10 + x%10
		x = x/10
	}
	if res > 1<<31-1 || res < -(1<<31) {
		return 0
	}
	return res*sign
}
func abs(a int) int{
	if a>=0 {
		return a
	}
	return -a
}