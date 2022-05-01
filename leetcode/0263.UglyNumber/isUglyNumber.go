package  lt263
/*简单
判断是否是丑数
丑数：只包含质因数2 3和/或5的正整数
因式分解 
*/
func isUgly(n int) bool {
	if n<1 {
		return false
	}
	for n%5==0{
		n /=5
	}
	for n%3==0{
		n /=3
	}
	for n%2==0{
		n/=2
	}
	return n==1
}