package lt326

/*
每日一题  简单

326. 3的幂
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3^x
*/
func isPowerOfThree(n int) bool {
	//能被3整除
	if n<1{   //0  负数都不行
		return false
	}
	res := 0 
	for n > 1 {
		res = n%3  //记录余数
		if res != 0 {
			return false
		}
		n = n / 3    //循环除
	}
	return res == 0
}