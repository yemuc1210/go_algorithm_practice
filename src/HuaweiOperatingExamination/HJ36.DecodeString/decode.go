package main

import (
	"fmt"
	// "strings"
)

func main() {
	//首先建立密文表
	var key, s string
	for {
		// n, err := fmt.Scanf("%s\n%s",&key,&s) 
		var n int
		var err error
		n,err = fmt.Scanln(&key)
		if n==0 || err!=nil{
			break
		}
		fmt.Scanln(&s)
		//加密
		fmt.Println(encrypt(key,s))
	}
}

func encrypt(key,src string) string{
	var flag [26]bool 
	var table []rune
	//对key去重
	for _,v := range key {
		//转为同一个大小写
		if v>='a' && v<='z' {  //A-65  a-97
			v -= 32
		}
		if !flag[int(v-'A')] {
			//未出现过的字符
			table = append(table, v)
			flag[int(v-'A')] = true
		}
	}

	//得到总表  剩下字符按顺序
	var tmp []rune = make([]rune, len(table))
	copy(tmp,table)
	// fmt.Println("tmp:",string(tmp))
	for i:='A';i<='Z';i++{
		//如果i不在tmp中  tmp不用map 因为map是无序的
		if !find(tmp,i) {
			table = append(table, i)  //保证顺序
		}
	}
	// fmt.Println("table:",string(table))
	//下面可以进行加密
	res := []rune{}
	for _,v := range src {
		if v>='A' && v<='Z' {
			res = append(res,table[int(v-'Z')])
		}else {
			//小写字符
			res = append(res, toLower(table[int(v-'a')]))
		}
	}
	return string(res)
}

func toLower(ch rune) rune {
	if ch>='a' && ch<='z'{
		return ch
	}else if ch>='A' && ch<='Z' {
		return ch+32
	}
	return ch
}
func find(tmp []rune, ch rune) bool {
	for i:=0;i<len(tmp);i++{
		if tmp[i] == ch {
			return true
		}
	}
	return false
}