package tencent50
/*简单
9. 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
*/

func isPalindrome(x int) bool {
	//负数不是
	if x < 0 {
		return false
	}
	//只有一位数的是
	if x>=0 && x<=9 {
		return true
	}
	//现在处理正数  多位数
	//得到逆序数，然后对比一下
	reverseNum := 0
	tmp := x
	for tmp > 0 {
		reverseNum = reverseNum*10 + tmp%10
		tmp /= 10
	}
	if reverseNum == x {
		return true
	}
	return false
}