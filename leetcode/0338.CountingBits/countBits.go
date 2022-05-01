package lt338
/*
338. 比特位计数
给你一个整数 n ，对于 0 <= i <= n 中的  每个 i ，
计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
*/
func countBits(n int) []int {
	//x&1==1 判断奇偶 >0奇数
	//x = x &(x-1)  清零最低位的1
	//x & -x   得到最低位的1
	// x&-x  0
	bits := make([]int,n+1)
	for i:=1;i<=n;i++{
		bits[i] = bits[i] + bits[i&(i-1)] + 1
	}
	return bits
}