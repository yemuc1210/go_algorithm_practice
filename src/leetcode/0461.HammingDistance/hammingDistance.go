package lt461
/*
461. 汉明距离
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
*/
/*
利用的位操作的是 X &= (X - 1) 不断的清除最低位的 1 。
先将这两个数异或，异或以后清除低位的 1 就是最终答案。
*/
func hammingDistance(x int, y int) int {
	distance := 0
	for xor := x^y;xor!=0;xor &= (xor-1) {
		distance++
	}  //数异或结果里的1的个数
	return distance
}