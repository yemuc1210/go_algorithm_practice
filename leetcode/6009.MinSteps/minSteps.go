package main

import (
	"fmt"
	// "sort"
)

// 给你两个字符串 s 和 t 。在一步操作中，你可以给 s 或者 t 追加 任一字符 。
// 返回使 s 和 t 互为 字母异位词 所需的最少步骤数。
// 字母异位词 指字母相同但是顺序不同（或者相同）的字符串。
func minSteps(s string, t string) int {
	// 比较bs,bt ， 删去相同的字母
	// deeel   as
	ms,mt := make(map[rune]int),make(map[rune]int)
	for _,v := range s {
		ms[v]++
	}
	for _,v := range t {
		mt[v]++
	}
	for k,v := range ms {
		if v1,ok:=mt[k];ok {
			if v>v1 {
				delete(mt,k)
				ms[k] -= v1
			}else {
				delete(ms,k)
				mt[k] -= v
			}
		}
	}
	var res int
	for _,v := range ms {
		res += v
	}
	for _,v := range mt {
		res += v
	}
	return res
}

func main() {
	s:="leetcode"   // 排序：cdeeelot  8
	t:="coats"   // 排序：acots  5
	res := minSteps(s,t)
	fmt.Println(res)
}