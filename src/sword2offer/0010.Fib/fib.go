package fib

//递归超时
func fib(n int)int{
	if n==0||n==1{
        return n
    }
    return fib(n-1)+fib(n-2)
}
//动态规划法，转移方程就是定义式
func fib_1(n int)int{
	//非递归写
	f0,f1 := 0,1
	for i:=0;i<n;i++{
		f0,f1 = f1,(f0+f1)%1000000007
	}
	return f0
}