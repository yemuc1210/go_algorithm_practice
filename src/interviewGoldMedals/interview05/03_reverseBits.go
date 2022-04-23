package interview05

/*
面试题 05.03. 翻转数位
给定一个32位整数 num，你可以将一个数位从0变为1。
请编写一个程序，找出你能够获得的最长的一串1的长度。
*/
func reverseBits(num int) int {
	cur := 0 // 当前位置连续1的个数  遇0归零
	insert := 0  // 当前位置变为1，往前数连续1的最大个数，遇到0 cur+1,遇到1+1
	res := 1
	for i:=0;i<32;i++{
		if num & 1 == 1{ // 倒数第i位是否为1    每次num都会移位
			cur += 1
			insert += 1
		}else{
			insert = cur+1
			cur = 0
		}
        num >>= 1   // 右移
		res = max(res,insert)
	}
	return res
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}