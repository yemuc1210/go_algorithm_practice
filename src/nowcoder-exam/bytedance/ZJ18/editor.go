package main

import (
	"fmt"
	// "runtime/internal/atomic"
)



func main() {
	var N int
	fmt.Scanln(&N)
	var s string
	for i:=0;i<N;i++{
		
		fmt.Scanln(&s)
		fmt.Println(output(s))
		//输出结果
		
	}
}
/*
三个同样的字母一定错误，去掉一个
两对一样的字母 去掉第二对一个字母
优先从左到右

自动机
输入A->状态0(A) 输入A->状态1(AA)   输入A-> 状态1(AA) 
			输入!A->状态0(AB)    输入B->状态2（AAB) 输入B->状态2（AAB)
											   输入!B ->状态0 (AABC) 从C开始
*/
func output(s string) string {
	last := s[0]  //第一个字符
	var res []byte
	res = append(res, s[0])

	//喂入状态机
	var state int
	for i:=1;i<len(s);i++ {
		cur := s[i]
		switch state {
		case 0:
			if cur == last {
				state = 1  //AA
			}else {
				state = 0  //AB 
			}
		case 1:
			if cur == last {
				// state = 1 //AAA
				continue
			}else{
				state = 2  //AAB
			}
		case 2:
			if cur == last {
				// state = 2  //AABB->AAB
				continue     //忽略
			}else{
				state = 0  //AABC  合法从C开始
			}
		default:
			
		}
		res = append(res, cur)
		last = cur
	}
	return string(res)
}