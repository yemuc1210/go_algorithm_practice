package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scanln(&n, &m, &k)
	//n 个页面投放m个广告，间隔k
	// 后续m行
	for i := 0; i < m; i++ {
		var nums []int = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d%c", &nums[i])
		}
		// 判断nums是否合理
		if judge(nums, k) {
			fmt.Print("1 ")
		} else {
			fmt.Print("0 ")
		}
		fmt.Println()
	}
}

func judge(nums []int, k int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for page := range m {
		_, e1 := m[page+k+1]
		_, e2 := m[page-k-1]
		if e1 || e2 {
			return false
		}
		delete(m, page)
	}
	return len(m) == 0
}
