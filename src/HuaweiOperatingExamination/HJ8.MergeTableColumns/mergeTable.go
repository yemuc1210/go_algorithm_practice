package main

import (
	"fmt"
	"sort"
)

/*
HJ8 合并表记录
描述
数据表记录包含表索引和数值（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照key值升序进行输出。
提示:
0 <= index <= 11111111
1 <= value <= 100000

输入描述：
先输入键值对的个数n（1 <= n <= 500）
然后输入成对的index和value值，以空格隔开

输出描述：
输出合并后的键值对（多行）
*/
func main(){
	var N int
	fmt.Scanln(&N)
	m := make(map[int]int)
	var key,val int
	var keys []int
	for i:=0;i<N;i++{
		//读取N对数
		fmt.Scanln(&key,&val)
		if _,ok := m[key]; ok {
			m[key] += val
		}else if _,ok :=m[key]; !ok{
			m[key] = val
			keys = append(keys, key)
		}
	}
	//按照Key值升序
	sort.Ints(keys)
	for i:=0;i<len(keys);i++{
		fmt.Println(keys[i], m[keys[i]])
	}
}