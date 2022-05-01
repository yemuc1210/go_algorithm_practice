package lt1446
/*
1446. 连续字符
给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。
请你返回字符串的能量。
*/
/*
子序列只有一种字符的最长长度
窗口
*/
func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1{
		return 1
	}

	left,right,n := 0, 0,len(s)
	res := 0
	for left <= right && right < n{
		if s[left] == s[right] {
			right ++
		} else{
			//更新
			res = max(res, right-left)
			left = right
		}
	}
	//有一种情况 字符串只有一个字符 跳出时 res=0
	res = max(res, right - left)  //加一步更新操作
	return res
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}