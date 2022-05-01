package main

import (
	"fmt"
	"sort"
)

func canReorderDoubled(arr []int) bool {
	// 哈希表+计数
	// 将数字分为一对一对的，每一对（x,2*x)
	// 0 只能和 0配对
	// m[x] == m[2x]
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	// 模拟
	if v, ok := m[0]; ok && v%2 != 0 {
		return false
	} else {
		delete(m, 0)
	}
	// 最好是按键值排序
	for k, v := range m {
		if v1, exist := m[2*k]; exist && v1 == v {
			// 配对成功
			delete(m, k)
			delete(m, 2*k)
		} else if v2, ok := m[k/2]; ok && v2 == v {
			delete(m, k)
			delete(m, k/2)
		} else {
			return false
		}
	}
	return len(m) == 0

}
func canReorderDoubled11(arr []int) bool {
	cnt := make(map[int]int, len(arr))
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 == 1 {
		return false
	}

	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })

	for _, x := range vals {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("vim-go")
	arr := []int{3, 1, 3, 6}
	fmt.Println(canReorderDoubled(arr))
	arr1 := []int{4, -2, 2, -4}
	fmt.Println(canReorderDoubled(arr1))
}
