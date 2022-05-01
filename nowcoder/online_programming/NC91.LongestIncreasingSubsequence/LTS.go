package NC91

/**
 * 最长递增子序列 lt300     
 * @param arr int整型一维数组 the array
 * @return int整型一维数组  最长上升子序列
*/
/*
动态规划+二分
贪心：让上升子序列尽可能长，则上升得需要尽可能慢  末尾添加数字需要尽可能小
d[i] 长度为i的最长上升子序列的末尾数字的最小值
length:当前上升子序列长度
d[i] 关于i递增 
遍历arr 更新d length
	arr[i] > d[length]  更新
	否则二分找到d[i-1] < arr[j] <d[i]  的下标i进行更新
p数组辅助 p[i]当前
*/
func LIS( arr []int ) []int {
    // dp[i]  以arr[i]结尾的递增序列长度
	//状态转移方程  dp[i] = max{ dp[j] } + 1   
	n := len(arr)
	d := make([]int,n+1)
	p := make([]int, n)
	for i:=0;i<n+1;i++{
		d[i] = -1
	}
	//二分+dp
	length := 1
	d[length] = arr[0]
	p[0] = 1
	for i:=1;i<n;i++{
		if arr[i] > d[length] {
			//此时将数字增加到末尾
			length++
			d[length] = arr[i]
			p[i] = length
		}else {
			//二分查合适的位置
			left,right,pos := 1,length,0
			for left <= right {
				mid := (left+right)/2
				if d[mid] < arr[i] {
					pos = mid
					left = mid+1
				}else{
					right = mid-1
				}
			}
			//对该位置更新
			d[pos+1] = arr[i]
			p[i] = pos+1
		}
	}
	res := make([]int,length)
	for i:=n-1;i>=0;i--{
		if p[i] == length {
			length--
			res[length] = arr[i]
		}
	}
	return res
}