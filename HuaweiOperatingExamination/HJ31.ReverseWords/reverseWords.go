package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var words string
	//输入 一行
	// fmt.Sscanln(words)
	// fmt.Scan(&words)
	// fmt.Scanln(&words)
	reader := bufio.NewReader(os.Stdin)
	b, _, _ := reader.ReadLine()
	words = string(b)
	fmt.Println(solve(words))
}

func solve(s string) string {
	//逆置单词
	//去除前导空格
	idx := 0
	for idx < len(s) && !isChar(s[idx])  {
		idx++
	}

	words := []string{}
	for idx < len(s) {
		//获取单词 假设前导空格都去除了
		word := []byte{}
		for idx < len(s) && isChar(s[idx])  {
			word = append(word, s[idx])
			idx++
		}
		if len(word) > 0 {
			words = append(words, string(word))
				//去除中间空格
			
		}
		for idx < len(s) && !isChar(s[idx]) {
			idx++
		}
	
	}
	var res string
	n := len(words)
	// fmt.Println(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}
	// fmt.Println(words)
	for i := 0; i < n; i++ {
		// words[i],words[n-i-1] = words[n-i-1],words[i]
		if i == n-1 {
			res += words[i]
		} else {
			res += words[i] + " "
		}
	}
	return res
}

//单词字符只有大小写英文字符
func isChar(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
