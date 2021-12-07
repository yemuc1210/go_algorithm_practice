package interview01

import "strings"

/*
面试题 01.09. 字符串轮转

字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

示例1:

 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
示例2:

 输入：s1 = "aa", s2 = "aba"
 输出：False


 s2*2   erbottlewaterbottlewat   s1包含在其中
*/
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) == 0 && len(s2) == 0{
		return true
	}
	s := s1+s1
	//strings.Index  返回s中字串s2的起始位置
	return len(s2)*2 == len(s) && strings.Index(s,s2) > -1

}

// func myStringsIndex(s,s1 string) int{
// 	if len(s)<len(s1){
// 		return -1
// 	}
// 	length := len(s1)
// 	//从s中找到s1的首字符
// 	for i:=0;i<len(s);i++{
// 		if 
// 	}
// }
