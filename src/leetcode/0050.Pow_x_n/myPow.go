package leetcode

import "math"

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
输入：x = 2.00000, n = 10
输出：1024.00000

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

提示
-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104
*/
// 存在超时问能题
func myPow(x float64, n int) float64 {
	if (n == 0 && x != 0.0) || (x == 1 && n != 0) {
		return 1
	}
	if x == -1 && n%2 == 0 {
		return 1
	}
	if x == -1 && n%2 != 0 {
		return -1
	}
	var res float64 = 1.0
	var lastRes float64 = res
	if n > 0 { //×
		for n > 0 {
			lastRes = res
			res *= x
			if math.Abs(res-lastRes) < 1e-10 {
				break
			}
			n -= 1
		}
	} else if n < 0 { //➗
		n = -n
		for n > 0 {
			lastRes = res
			res = res * (1 / x)
			if math.Abs(res-lastRes) < 1e-10 {
				break
			}
			n -= 1
		}
	}
	return res
}
