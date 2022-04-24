package main

import "fmt"

func binaryGap(n int) int {
	// 中间必须隔着0
	var last int = -1
	var res int
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				res = max(res, i-last)
			}
			//else
			last = i
		}
		// 移位
		n >>= 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int = 22
	res := binaryGap(n)
	fmt.Println(res)
}
