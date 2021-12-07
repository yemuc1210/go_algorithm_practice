package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//字符串排序

func main(){
	//输入 一个正整数
	reader := bufio.NewScanner(os.Stdin)
	var N int64
	fmt.Scanln(&N)
	// fmt.Println(N)
	//下面N行字符串
	inputs := make([]string, N)
	var i int64
	for i=0;i<N;i++{
		// fmt.Scanln(&s)   //每一行的字符串
		reader.Scan()
		s := reader.Text()
		// fmt.Sscanln(s)
		// fmt.Println(s)
		inputs[i] = s
	}
	// fmt.Println(inputs)
	//输出
	fmt.Println(len(inputs))
	sort.Strings(inputs)
	fmt.Println(len(inputs))
	for j:=0;j<len(inputs);j++{
		fmt.Println(inputs[j])
	}
}