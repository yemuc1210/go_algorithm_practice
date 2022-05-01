package main

import "fmt"

func main(){
	// m := make(map[byte]int)
	var helper = [26]int{2,2,2,3,3,3,4,4,4,5,5,5,6,6,6,7,7,7,7,8,8,8,9,9,9,9}
	//把密码中小写字母变成数字   数字和其他符合不变
	var s string
	//输入包含多个测试
	for {
		n,_ := fmt.Scan(&s)
		// fmt.Println(s,n)
		if n==0 {
			break
		}else{
			//长度不超过100
			if n > 100 {
				break
			}
			res := make([]byte,len(s))  //密文
			for i:=0;i<len(s);i++{
				if s[i]>='a' && s[i]<='z' {
					//小写字母变成数字
					// fmt.Println(s[i]-'a',helper[s[i]-'a'])
					res[i] = byte(helper[s[i]-'a'] + 48)
					// fmt.Println(res[i])
				}else if s[i]>='A' && s[i]<='Z'{
					//大写字母变小写 再后移动 大写65-90  小写97-122
					ch := ( (s[i]-'A'+1)%26 + 65 ) + 32
					res[i] = ch
				}else{
					//数字和其他数字不变
					res[i] = s[i]
				}
			}
			for i:=0;i<len(res);i++{
				fmt.Printf("%c",res[i])
			}
			fmt.Println()
		}
	}
}