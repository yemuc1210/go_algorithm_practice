package offer05

import "strings"

/*
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/
func replaceSpace(s string) string {
	return strings.ReplaceAll(s," ","%20")
}


//思路2
//模拟
func replaceSpace1(s string) string {
	var res string
	for i:=0;i<len(s);i++{
		if s[i] != ' '{
			res += string(s[i])
		}else {
			res += "%20"
		}
	}
	return res
}