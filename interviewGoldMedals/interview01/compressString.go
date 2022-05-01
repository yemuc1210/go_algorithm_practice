package interview01

import (
	"fmt"
	"strconv"
)

/*
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
比如，字符串aabcccccaaa会变为a2b1c5a3。
若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"

 示例2:

 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
*/
func compressString(S string) string {
	if len(S) == 0 || len(S)==1{
		return S
	}
	//字符计数
	curByte,curCnt := S[0],1
	var res string
	fmt.Println("len(S)=",len(S))
	for i:=1;i<len(S);{
		fmt.Printf("curByte=%c,S[%v]=%c\n",curByte,i,S[i])
		if S[i] == curByte {
			curCnt++
			i++
		}else{
			res += string(curByte) + strconv.Itoa(curCnt)
			curByte = S[i]
			curCnt = 1
			i++
		}
	}
	//最后一组需要另外处理
	res += string(curByte) + strconv.Itoa(curCnt)
	if len(res) >= len(S) {
		return S
	}
	return res
}