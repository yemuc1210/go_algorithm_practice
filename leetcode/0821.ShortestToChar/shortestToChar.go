package main

import (
	"fmt"
	"reflect"
)

/*
给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。

返回一个整数数组 answer ，其中 answer.length == s.length 且
	answer[i] 是 s 中从下标 i 到离它 最近 的字符 c 的 距离 。

两个下标 i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数。

*/

func shortestToChar(s string, c byte) []int {
	// 找出c出现的所有位置
	mLoc := make(map[int]bool)
	for i:=0;i<len(s);i++{
		if s[i]==c{
			mLoc[i] = true
			fmt.Println(i)
		}
 	}
	res := make([]int, len(s))
	for i:=0;i<len(s);i++{
		if s[i]==c {
			res[i] = 0
		}else{
			// 遍历mLoc
			min := len(s)+1
			for k := range mLoc {
				fmt.Println("abs(k,i)",k,i,abs(k,i))
				if abs(k,i) < min {
					min = abs(k,i)
				}
			}
			res[i] =min
		}
	}
	return res
}

func abs(a,b int) int {
	if a>b {return a-b}
	return b-a
}
// 两次遍历
func shortestToChar1(s string, c byte) []int {
    n := len(s)
    ans := make([]int, n)

    idx := -n
    for i, ch := range s {
        if byte(ch) == c {
            idx = i
        }
        ans[i] = i - idx
    }

    idx = n * 2
    for i := n - 1; i >= 0; i-- {
        if s[i] == c {
            idx = i
        }
        ans[i] = min(ans[i], idx-i)
    }
    return ans
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func main() {
	s := "loveleetcode"
	var c byte = 'e'
	ans := []int{3,2,1,0,1,0,0,1,2,2,1,0}
	res := shortestToChar(s,c)
	fmt.Println(res)
	if reflect.DeepEqual(ans,res) {
		fmt.Println("pass")
	}	
}