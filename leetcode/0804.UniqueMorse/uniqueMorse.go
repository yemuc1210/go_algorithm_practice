package main

import (
	"fmt"
)

func uniqueMorseeRepresentations(words []string) int {
	// 定义摩斯密码
	var m = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	// 解析单词
	var parseRes = make(map[string]int)
	for _, word := range words {
		// 解析单词
		var parse []byte
		for i := 0; i < len(word); i++ {
			ch := int(word[i]-'a')
			parse = append(parse, []byte(m[ch])...)
		}
		parseRes[string(parse)]++
	}
	return len(parseRes)
}

func main() {
	fmt.Println("vim-go")
	res := uniqueMorseeRepresentations([]string{"gin", "zen", "gig", "msg"})
	fmt.Println(res)
}
