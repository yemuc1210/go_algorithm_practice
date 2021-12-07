package lt204
/*中
204. 计数质数
统计所有小于非负整数 n 的质数的数量。
提示：
	0 <= n <= 5 * 106
*/
/*
计算质数  借鉴计算丑数的思路 利用倍数关系
0 1既不是质数也不是合数
2 质数  2的倍数不是 划去
3 质数  3的倍数不是 划去
4 
*/
func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	//构建长度为n的数组  空间换时间？  根据提示 n最大为5M
	nonPrimes := make([]bool, n)
	for i:=2;i*i<n;i++{
		if nonPrimes[i] { //非质数
			continue
		}
		//否则 是质数
		for j:=i*i; j<n;j = j+i { //j遍历i的倍数
			nonPrimes[j] = true  // 是质数倍数关系的 不是质数
		}
	}
	cnt := 0 
	for i:=2;i<n;i++{
		if !nonPrimes[i] {
			cnt++
		}
	}
	return cnt
}

