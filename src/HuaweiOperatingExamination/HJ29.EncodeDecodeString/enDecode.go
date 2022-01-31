package main

import "fmt"

/*
字符串加密
描述
1、对输入的字符串进行加解密，并输出。
2、加密方法为：
当内容是英文字母时则用该英文字母的后一个字母替换，同时字母变换大小写,如字母a时则替换为B；字母Z时则替换为a；
当内容是数字时则把该数字加1，如0替换1，1替换2，9替换0；
其他字符不做变化。
3、解密方法为加密的逆过程。
*/
func main() {
	var toEncode string
	var encoded string
	for {
		//输入一串要加密的密码
		// 输入一串加过密的密码
		n,_ := fmt.Scanln(&toEncode)
		if n==0 {
			break
		}
		fmt.Scanln(&encoded)
		//输出加密后的
		fmt.Println(solve(toEncode,1))
		//输出解密后的
		fmt.Println(solve(encoded,-1))
	}
}

func solve(src string,isEncode int) string {
	/*
	当内容是英文字母时则用该英文字母的后一个字母替换，同时字母变换大小写,如字母a时则替换为B；字母Z时则替换为a；
	当内容是数字时则把该数字加1，如0替换1，1替换2，9替换0；
	*/
	bytes := []byte(src)
	for i:=0;i<len(bytes);i++ {
		if isDigit(bytes[i]) {
			d := int(bytes[i]-'0')
			bytes[i] = byte((d+isEncode+10)%10+'0')
		}else if f1,f2:=isChar(bytes[i]);f1 {
			if f2 {
				//大写
				d := int(bytes[i]-'A')
				// 转小写
				bytes[i] = byte((d+isEncode+26)%26+'A') + 32
				
			}else {
				//小写
				d := int(bytes[i]-'a')
				bytes[i] = byte((d+isEncode+26)%26+'a') - 32
			}
		}
	}
	return string(bytes)
}

func isDigit(bt byte) bool {
	if bt>='0' && bt<='9' {
		return true
	}
	return false
}
func isChar(bt byte)(bool,bool) {
	if bt>='A' && bt<='Z' {
		return true,true
	}else if bt>='a' && bt<='z' {
		return true,false
	}
	return false,false
}