package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	var res int

	// 用栈stack := []int{ops[0]}
	stack := []int{}
	var i int = 0
	for i < len(ops) {
		if num, err := strconv.Atoi(ops[i]); err == nil {
			stack = append(stack, num)
		} else if ops[i] == "+" {
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			//		stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		} else if ops[i] == "D" {
			a := stack[len(stack)-1]

			stack = append(stack, 2*a)
		} else if ops[i] == "C" {
			stack = stack[:len(stack)-1]
		}
		i++
	}
	for _, v := range stack {
		res += v
	}
	return res
}

func isDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	ops := []string{"5", "2", "C", "D", "+"}
	res := calPoints(ops)
	fmt.Println(res)
}
