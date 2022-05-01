package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
HJ26 字符串排序
描述
编写一个程序，将输入字符串中的字符按如下规则排序。

规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。

如，输入： Type 输出： epTy

规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。

如，输入： BabA 输出： aABb

规则 3 ：非英文字母的其它字符保持原来的位置。


如，输入： By?e 输出： Be?y
*/
func main() {
	var str []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str = append(str, input.Text())
	}

	for loop := 0; loop < len(str); loop++ {
		NewPrint(str[loop])
        fmt.Print("\n")
	}

}
 
func NewPrint(str string) {
	var tmpStr []string
	tmpMap := make(map[int]byte)
	for loop := 0; loop < len(str); loop++ {
		if (str[loop] >= 'a' && str[loop] <= 'z') || (str[loop] >= 'A' && str[loop] <= 'Z') {
			tmpStr = append(tmpStr, string(str[loop]))
		} else {
			tmpMap[loop]++
		}
	}
	sort.SliceStable(tmpStr, func(i, j int) bool {
		return strings.ToLower(tmpStr[i]) < strings.ToLower(tmpStr[j])
	})

	var loop2 int
	for loop := 0; loop < len(str); loop++ {
		if _, ok := tmpMap[loop]; ok {
			fmt.Print(string(str[loop]))
		} else {
			fmt.Print(tmpStr[loop2])
			loop2++
		}
	}
}

//公共
