package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	// 有效的电子邮件地址：本地名@域名
	// 除了小写字母，还可以含有一个或多个'.' '+'
	// 本地名中介有句点'.'，则忽略
	// 本地名有'+'，则忽略第一个加号后面所有内容
	
	res := 0
	// m := make(map[string]int)
	// 1. 判断有么@
	for _,email := range emails {
		exist := strings.ContainsRune(email,'@')
		if !exist {
			continue
		}
		// 解析email
		names := strings.Split(email,"@")
		fmt.Println(names)
	}
	return res
}

func main() {
	emails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com"}
	res := numUniqueEmails(emails)
	fmt.Println(res)
}

// func main() {

// }