package main

import "fmt"

func hasAlternatingBits(n int) bool {
	// 给正整数，检查对应二进制是否0，1交替
	var lastIsOne int = n & 1
	n = n >> 1
	for n != 0 {
		last := n & 1
		n = n >> 1
		if (last == 1 && lastIsOne == 0) || (last == 0 && lastIsOne == 1) {
			lastIsOne = last
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
