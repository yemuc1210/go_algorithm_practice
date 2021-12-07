package main

import "fmt"

/*
HJ10 字符个数统计
描述
编写一个函数，计算字符串中含有的不同字符的个数。字符在 ASCII 码范围内( 0~127 ，包括 0 和 127 )，换行表示结束符，不算在字符里。不在范围内的不作统计。多个相同的字符只计算一次
例如，对于字符串 abaca 而言，有 a、b、c 三种不同的字符，因此输出 3 。

数据范围：
输入描述：
输入一行没有空格的字符串。

输出描述：
输出 输入字符串 中范围在(0~127，包括0和127)字符的种数。
*/
func main(){
	//输入是没有空格的字符串 
	var input string
	fmt.Scanln(&input)
	//统计出现字符的种类
	var cnt int
	m := make(map[byte]int)
	for i:=0;i<len(input);i++{
		if _,ok := m[input[i]]; !ok {
			m[input[i]] ++
			cnt ++
		}else  {
			m[input[i]] ++
			continue
		}
	}
	fmt.Println(cnt)
}