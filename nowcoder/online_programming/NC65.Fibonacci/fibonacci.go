package NC65
/**
 * 斐波那契数列
 * @param n int整型 
 * @return int整型
*/
func Fibonacci( n int ) int {
    // 迭代计算
	if n == 1 || n == 2{
		return 1
	}
	fib := make([]int, n)
	fib[0],fib[1] = 1,1
	for i:=2;i<n;i++{
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n-1]
}