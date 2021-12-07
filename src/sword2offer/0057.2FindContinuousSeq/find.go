package offer57_2
/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

连续正整数
滑动窗口法
*/
func findContinuousSequence(target int) [][]int {
	i,j,sum := 1,2,3
	res := [][]int{}
	for i<=target/2 {
		if target>sum{
			//右边界右移
			j++
			sum+=j
		}else{
			if target==sum{  //记录
				tmp := make([]int,j-i+1)
				for k:=i;k<=j;k++{
					tmp[k-i] = k
				}
				res =append(res, tmp)

			}
			//左边界左移
			sum -= i
			i++
		}
	}
	return res
}