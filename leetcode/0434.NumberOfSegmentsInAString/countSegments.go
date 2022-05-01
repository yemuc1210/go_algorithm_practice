package lt434

import "fmt"

/*
434. 字符串中的单词数
统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。
示例:

输入: "Hello, my name is John"
输出: 5
解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。

*/
func countSegments(s string) int {
	if len(s) == 0{
		return 0
	}

	res := 0

	//首先需要跳过前置空格
	index := 0 
	for index < len(s) && s[index] ==' '{
		index ++
	}

	//判断
	for index < len(s) {
		tmpIdx := index
		for tmpIdx<len(s) && s[tmpIdx] != ' '{
			tmpIdx++
		}
		res ++ 
		fmt.Println(res,tmpIdx)
		//找到下一个非空
		index = tmpIdx
		for index < len(s) && s[index] ==' '{
			index ++
		}
	
	}
	return res
}