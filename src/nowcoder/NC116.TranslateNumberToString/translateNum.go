package NC116



//剑指offer 46 有区别

/**
 * 解码
 *字母编码成数字：a-1 b-2   z-26
 * 进行解码，将数字解码成字符串，返回可能的结果
 * @param nums string字符串 数字串
 * @return int整型
 */

/*
dp
翻译规则：
	字符串第i位置 可以单独翻译
	如果第i-1和第i位组成数字在10-25之间，可以连起来翻译
f(i)表示以第i位结尾的前缀串翻译的方案数
f(i)=f(i-1)+f(i-2)  i-1>=0     10<=x<=25   i-2和i位一起翻译
f(-1)=0  f(1)=0

滚动数组压缩空间，因为只有三个变量
*/
func solve( nums string ) int {
	dp := make([]int,len(nums)+1)
	if len(nums)==0 || nums=="0" {
		return 0
	}
	dp[0]=1
	for i:=1;i<=len(nums);i++{
		if nums[i-1]=='0' {
			dp[i]=0
		}else{
			dp[i]=dp[i-1]
		}
		if (i>1) && ((nums[i-2]=='1') || (nums[i-2]=='2' && nums[i-1]<='6')) {
			dp[i] += dp[i-2]
		}
	
	}
	return dp[len(nums)]
	// p, q, r := 0, 0, 1 //前三项
	// for i := 0; i < len(nums); i++ {
	// 	p, q, r = q, r, 0 //f(i-2)  f(i-1)  f(i)
	// 	r += q            //f(i) 一定包含f(i-1)
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	pre := nums[i-1 : i+1] //【i-1,i]
	// 	if pre <= "26" && pre >= "10" {
	// 		r += p
	// 	}
	// }
	// return r
}
