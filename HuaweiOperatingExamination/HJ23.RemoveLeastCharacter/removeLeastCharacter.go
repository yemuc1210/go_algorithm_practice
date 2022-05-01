package main

import "fmt"

//删除字符串中出现次数最少的字符 若次数一样 全删除
//保持原顺序
func main(){
	var s string
	//有多个输入
	for {
		n,_ := fmt.Scan(&s)
		if n==0 {
			break
		}
		// fmt.Println(s)
		//只包含小写字母 不考虑非法输入  输入长度小于等于20
		var isChValid = make([]bool, len(s))  //标记当前字符是否该删除
		m := make(map[byte]int)  //字符计数
		for i:=0;i<len(s);i++{
			m[s[i]] ++
		}
		//找出出现次数最少的
		min := len(s)+1
		for i:=0;i<len(s);i++{
			if m[s[i]] < min {
				min = m[s[i]]
			}
		}
		// fmt.Println(min)
		//删除
		for i:=0;i<len(s);i++{
			if m[s[i]] == min {
				// fmt.Printf("%c",s[i])
				isChValid[i] = true
			}
		}
		//输出
		for i:=0;i<len(s);i++{
			if !isChValid[i] {
				fmt.Printf("%c",s[i])
			}
		}
		fmt.Println()
	}
}