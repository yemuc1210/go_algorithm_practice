package main

import "fmt"

func maximumGain(s string, x int, y int) int {
	// 贪心，根据x y大小觉得先消除“ab"还是”ba"
	// 栈  括号消除匹配
	res := 0 
	var goals []string
	var scores []int
	if x>y {
		// 先消除“ab"
		goals = []string{"ab","ba"}
		scores = []int{x,y}
	}else{
		goals = []string{"ba","ab"}
		scores = []int{y,x}
	}
	// 消除
	for i,goal:=range goals{
		score := scores[i]
		// 本轮消除goal
		var tmp = []byte{}
		for j:=0;j<len(s);j++{
			tmp = append(tmp, s[j])
			if len(tmp)<2 {
				continue
			}
			// 判断是否出现模板
			if string(tmp[len(tmp)-2:])!=goal {
				continue
			}
			// 出现目标
			res += score
			// 删除
			tmp = tmp[:len(tmp)-2]
		}
		// s 变化
		s = string(tmp)
	}
	return res
}

func main() {
	s := "cdbcbbaaabab"
	x := 4
	y := 5
	res := maximumGain(s,x,y)
	fmt.Println(res)
}