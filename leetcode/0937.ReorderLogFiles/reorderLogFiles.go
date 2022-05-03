package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	sort.Slice(logs, func(i, j int) bool {
		s,t := logs[i],logs[j]
		// 分割
		s1 := strings.SplitN(s," ",2)[1]
		s2 := strings.SplitN(t," ",2)[1]  // 内容
		// 判断日志类型
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		// 排序规则
		if isDigit1 && isDigit2 {
			// 数字日志顺序不变
			return false
		}
		if !isDigit1 && !isDigit2 {
			// 字母日志  首先根据内容 然后根据类型排序
			return s1<s2 || (s1==s2 && s<t)
		}
		// 否则 字母在前
		return isDigit1
	})
	return logs
}

func main(){
	logs :=[]string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	res := reorderLogFiles(logs)
	fmt.Println(res)
}