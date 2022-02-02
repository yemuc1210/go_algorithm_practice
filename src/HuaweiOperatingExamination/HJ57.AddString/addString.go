package main

import (
	"fmt"
)

func main() {
	var s1,s2 string
	for {
		n1,_ := fmt.Scanln(&s1)
		if n1==0 {
			break
		}
		n2,_ :=fmt.Scanln(&s2)
		
		//做s1+s2
		var res []int
		// 令s1指向长度较大的  s2指向长度较小的 模拟竖式加法
		n1,n2 = len(s1),len(s2)
		// fmt.Println(n1,n2)
		if n1<n2 {
			// 交换
			s1,s2 = s2,s1
			n1,n2 = n2,n1
		}
		// fmt.Println(n1,n2, s1,s2)
		i,j := n1-1,n2-1
		var carry int  // 进位
 		for i>=0 || j>=0 {
			// ans := s1[i]+s2[j]
			var a1,b1 int
			if i>=0 {
				a1 = int(s1[i]-'0') 
				i--
			}
			if j>=0 {
				b1 = int(s2[j]-'0')
				j--
			}
			sum := a1+b1+carry
			res = append(res, sum%10)  //先倒置没关系 待会计算最后结果就从下标0开始计算
			carry = sum/10
			// fmt.Println(a1,b1,carry,sum)
			// fmt.Println(res)
		}
		// carry
		if carry != 0 {
			res =append(res, carry)
		}

		// 根据数组计算最终结果
		// var sum string
		for i:=len(res)-1;i>=0;i-- {
			fmt.Print(res[i])
		}
		fmt.Println()
	}
}