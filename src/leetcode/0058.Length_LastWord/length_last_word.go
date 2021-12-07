/*
给你一个字符串 s，由若干单词组成，单词之间用 空格 隔开。
返回字符串中最后一个单词的长度。
如果不存在最后一个单词，请返回 0 。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
*/
package leetcode

//直接从后往前遍历？？？？
func lengthOfLastWord(s string) int {
	if len(s)==0{
		return 0
	}
	//注意：“ ”长度为1
	index :=0
	n:= len(s)
	for index<n && s[index]==' '{
		index ++
	}
	if index>=n{
		return 0
	}
	index = n-1
	for index<n&&index>=0&&s[index]==' '{
		index --
	}
	//此时inedx 执行第一个不为空的位置
	n = index
	for index>=0 && s[index]!=' '{
		index --
	}
	return n-index
}