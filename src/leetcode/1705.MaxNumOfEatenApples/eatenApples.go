package lt1705

import "container/heap"

/*
1705. 吃苹果的最大数目

有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。
在第 i 天，树上会长出 apples[i] 个苹果，
这些苹果将会在 days[i] 天后（也就是说，第 i + days[i] 天时）腐烂，变得无法食用。
也可能有那么几天，树上不会长出新的苹果，此时用 apples[i] == 0 且 days[i] == 0 表示。

你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。

给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。
*/

//贪心 吃距离保质期最近的
func eatenApples(apples []int, days []int) int {
	h := hp{} 
	idx := 0
	res := 0
	for ;idx<len(apples);idx++{
		for len(h)>0 && h[0].end <= idx {  //优先权队列h[0]总是end最小的也就是快过期的
			heap.Pop(&h)   //抛弃过期的
		}
		if apples[idx] > 0 {
			//产生苹果
			heap.Push(&h, pair{idx+days[idx], apples[idx]})
		}
		if len(h) > 0 {
			h[0].left -- 
			if h[0].left == 0 {
				heap.Pop(&h)  //调整
			}
			res ++
		}
	}
	//n天后可以继续吃
	for len(h)>0 {
		for len(h)>0 && h[0].end <= idx {
			heap.Pop(&h)
		}
		if len(h) == 0 {
			break  
		}
		p := heap.Pop(&h).(pair)

		num := min(p.end - idx,   p.left)
		res += num   //一次性处理num天
		idx += num
	}
	return res
}

//优先权队列
type pair struct{
	end int
	left int //剩余的
}
type hp []pair

//Len Less  Swap  实现接口用
func (h hp) Len() int {return len(h)}
func (h hp) Less(i,j int) bool {return h[i].end<h[j].end}
func (h hp) Swap(i,j int) {h[i],h[j] = h[j],h[i]}

func(h *hp) Push(v interface{}){
	*h = append(*h, v.(pair))   //v.(pair) 断言 好像是这个 强制类型转换
}
func(h *hp) Pop() interface{}{
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}