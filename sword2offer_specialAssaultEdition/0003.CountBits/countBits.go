package offerS3
/*
剑指 Offer II 003. 前 n 个数字二进制中 1 的个数
给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。

移位
*/
func countBits(n int) []int {
	res := make([]int,n+1)
	/*
	如果最后一位为0，则下一个数字的1个数会+1
	如果最后一位为1，下一个数字会导致最后一位变成0，进一位，相当于右移一位之后+1的数字中1的个数
	*/
	for i:=0;i<=n;i++ {
		res[i] = res[i>>1] + (i&0x1)
	}
	return res
}