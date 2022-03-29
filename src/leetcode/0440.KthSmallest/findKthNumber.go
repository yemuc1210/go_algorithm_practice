package main

import (
	"fmt"
)
func getSteps(cur, n int) (steps int) {
	// first 左侧节点   last 右侧节点
    first, last := cur, cur
    for first <= n {
		// 当前层个数
        steps += min(last, n) - first + 1
        first *= 10
        last = last*10 + 9
    }
    return
}

func findKthNumber(n, k int) int {
    cur := 1
    k--
    for k > 0 {
        steps := getSteps(cur, n)
        if steps <= k { // 不在cur节点的子树中，以cur节点为根的子树有steps个节点
            k -= steps
            cur++
        } else {  // 在cur节点子树中，向下一层继续遍历
            cur *= 10
            k--
        }
    }
    return cur
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
 
func main() {
	res := findKthNumber(13, 2)
	fmt.Println("vim-go", res)
}
