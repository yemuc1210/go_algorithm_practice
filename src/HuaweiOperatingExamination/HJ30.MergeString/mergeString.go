package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1,s2 string
	for {
		n,_ := fmt.Scan(&s1,&s2)
		if n==0 {
			break
		}
		//输出一行
		// 第一步 合并
		s := s1+s2
		// 第二步 排序
		var b1,b2 []byte
		for i:=0;i<len(s);i++ {
			if i%2==0 {  // 偶数
				b2 = append(b2, s[i])
			}else{ // 奇数
				b1 = append(b1, s[i])
			}
		}
		sort.Slice(b1,func(i, j int) bool {
			return b1[i]<b1[j]
		})
		sort.Slice(b2,func(i, j int) bool {
			return b2[i]<b2[j]
		})
		// 合并归位
		var b []byte
		var turn int=2
		for i:=0;i<len(s);i++{
			if turn%2==0 {
				// 偶数位
				b = append(b, b2[0])
				b2 = b2[1:] //弹出开头的
				turn = 1
			}else {
				// 奇数
				b = append(b, b1[0])
				b1 = b1[1:]
				turn = 2
			}
		}
		// 第三步 
		for i:=0;i<len(b);i++{
			//如果是0-9 ok
			if b[i]>='0' && b[i]<='9' {
				a := toInteger(toBytes(int(b[i]-'0')))
				b[i] = int2Ch(a)
			}else if b[i]>='a' && b[i]<='f' {
				a := int(b[i]-'a') + 10 
				a = toInteger(toBytes(a))
				// 变为字符
				b[i] = int2Ch(a)
			}else if b[i]>='A' && b[i]<='F' {
				a := int(b[i]-'A') + 10
				a = toInteger(toBytes(a))
				b[i] = int2Ch(a)
			}
		}
		// 输出
		fmt.Println(string(b))
	}
}

// 将数字变为二进制表示
func toBytes(n int) []int {
	var res []int
	for n!=0 {
		res = append(res,n&1)  //最后一位是否是1
		// 移位
		n = n >> 1  //向右移位
	}
	return res   // 正好是倒叙的
}
func toInteger(b []int) int {
	//将01序列变为整数
	factor := 1
	var res int
	for i:=len(b)-1;i>=0;i-- {
		res = res + b[i]*factor
		factor *= 2
	}
	return res
}
func int2Ch(n int) byte {
	// 十进制数字变十六进制
	if n>=0 && n<=9 {
		return byte('0'+n)
	}else{
		return byte(n-10+'A')
	}
	
}