package main

// import "fmt"

//1 2 3 4 5 6 7      8*7/2 = 28
//2 3 4 5 6 7 8      28 + 7 = 35
//3 4 5 6 7 8 9      28 + 2*7 = 42
//4 5 6 7 8 9 10     28 + 3*7 = 49
func totalMoney(n int) int {
	if n <=7 {
		return (1+n)*n/2
	}
	//否则 超过7天 根据等差序列的公式计算
	fullWeeks := n/7
	days := n%7
	// 28+(fullWeeks-1)*7  上一周的  
	//等差序列求和公式   Sn=n*a1 +n*(n-1)*d/2  Sn=n*(a1+an)/2
	//现在 求an即可   an= 1+fullWeeks + (days-1)     
	//a1-an-1求和先
	an := 28 + (fullWeeks-1) * 7  
	sum := fullWeeks*(28 + an)/2  
	lastWeekLastDay := fullWeeks+days   
	// println(curWeekFirstDay)
	for i:=0;i<days;i++{  //days 记录最后一周有几天
		sum += lastWeekLastDay   //从最后一天的数开始往前加
		lastWeekLastDay --
	}
	return sum
}

func main() {
	res := totalMoney(10)
	println(res)
}