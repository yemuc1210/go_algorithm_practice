package lt1137

/*
求第N个泰波那契数
T0=0
T1=1
T2=1

Tn+3 = Tn + Tn+1 + Tn+2

dp
*/
func tribonacci(n int) int {
	if n == 0 || n==1{
		return n
	}
	if n == 2 {
		return 1
	}

	dp0,dp1,dp2,dp3 := 0,1,1,2

	for i:=3;i<=n;i++{
		dp3 = dp0 + dp1 + dp2
		dp2,dp1,dp0 = dp3,dp2,dp1
	}
	return dp3
}