package main

import (
	"fmt"
)

func lenOk(s string) bool{ // 长度超过8位
	return len(s) > 8
}
func kindOK(s string) bool{ // 包括大小写字母.数字.其它符号,以上四种至少三种
	n,A,a,c := 0,0,0,0 // 数字 大写字母 小写字母 特殊字符标识
	for i := 0 ; i < len(s) ; i ++{
		chr := s[i]
		if chr >= 48 && chr <= 57 {
			n = 1
			continue
		}else if chr >= 65 && chr <= 90 {
			A = 1
			continue
		}else if chr >= 97 && chr <= 122 {
			a = 1
			continue
		}else {
			c = 1
		}
		if n+A+a+c >= 3 { // 满足条件提前终止循环
			break
		}
	}
	//fmt.Println(n,a,A,c)
	return n+A+a+c >= 3
}
func repeatOK(s string) bool { //不能有相同长度大于2的子串重复
	tag := true
	cnt := make(map[string]int) // 计数字串
	for i := 0; i < len(s)-2; i++ {
		j := i + 3
		slice := s[i:j]
		//fmt.Println(slice)
		if _,ok := cnt[slice] ; ok {
			return false
		}else {
			cnt[slice] = 1
		}
	}
	return tag
}

func main() {
	pwd := ""
	for {
		n, _ := fmt.Scan(&pwd)
		if n == 0 {
			break
		} else {
			//fmt.Println(lenOk(pwd),kindOK(pwd),repeatOK(pwd))
			//fmt.Printf("%s\n",pwd)
			if lenOk(pwd)&&kindOK(pwd)&&repeatOK(pwd){
				fmt.Println("OK")
			}else {
				fmt.Println("NG")
			}
		}
	}
}