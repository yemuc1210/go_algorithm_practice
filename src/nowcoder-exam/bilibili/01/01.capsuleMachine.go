package bilibili

// import "fmt"

//第一题 扭蛋机
/*
2号扭蛋机：2x+1
3号扭蛋机：2x+2
题目要求：最后恰好扭到N个  不是总共N
dp[n] 表示扭到N需要的最小次数
dp[1] = 1   //2#
dp[2] = 1   //3#
dp[3] = 3   //no 因为是轮流
dp[4] = 2   //2#+
dp[i] = min { dp[i-1] + 1   i=2(i-1)+1 || ,  2*dp[i-1]+2}
这道题可以简化成走方格
初始位置是x=0，然后每步有两种选择
1、第一种选择是走到编号为 2*x+1 的方格，代号为 ‘2’  2号鸡
2、第一种选择是走到编号为 2*x+2 的方格，代号为 ‘3’   3号鸡
求一种可以正好走到 编号为 n 的方式

# 逆推法：根据最后一步的奇偶，判断前一步的选择
n = int(input())
res = ''
while n > 0:
    if (n-2) % 2:
        n = (n-1) // 2
        res = '2' + res
    else:
        n = (n-2) // 2
        res = '3' + res
print(res)
*/
func capsuleMachine(n int) string {
	var res string
	for n>0 { 
		//根据奇偶判断前一步
		if (n-2) % 2 != 0 {  //前一步是走的2号
			res = "2" + res
			n = (n-1)/2
		}else{
			res = "3" + res
			n = (n-2)/2
		}
	}
	return res
}

// func main(){
//     var n int
//     fmt.Scanln(&n)
//     fmt.Println(capsuleMachine(n))
// }