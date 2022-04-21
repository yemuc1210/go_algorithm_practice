package main

import "fmt"

// 空格替换不是数字的字符
// 剩下的整数以空格隔开
// 返回不同整数的个数
func numDifferentIntegers(word string) int {
	if len(word)==0 {
		return 0
	}
	// map 去重
	res := make(map[string]struct{})
	var curNum []byte
	// 处理字符串
	for i:=0;i<len(word);{
		if word[i]>='0' && word[i]<='9' {
			for i<len(word) && word[i]>='0' && word[i]<='9' {
				curNum =append(curNum,  word[i])
				i++
			}
			// 数据处理，去除前置0
			var j int
			for j=0;j<len(curNum)&&curNum[j]=='0';j++{}
			curNum = curNum[j:]
			res[string(curNum)] = struct{}{}
			curNum = []byte{}
		}else{
			i++
		}
	}
	return len(res)
}

func main() {
	word := "192383183928778851682383a2089984061937879119"
	res := numDifferentIntegers(word)
	fmt.Println(res)
}