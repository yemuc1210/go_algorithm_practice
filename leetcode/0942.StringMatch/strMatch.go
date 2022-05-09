package main

import "fmt"
// 贪心法
func diStringMatch(s string) []int {
    n := len(s)
    perm := make([]int, n+1)
    lo, hi := 0, n
    for i, ch := range s {
        if ch == 'I' {
            perm[i] = lo
            lo++
        } else {
            perm[i] = hi
            hi--
        }
    }
    perm[n] = lo // 最后剩下一个数，此时 lo == hi
    return perm
}

func main() {
	s := "IDID"
	res := diStringMatch(s)
	fmt.Println(res)
}