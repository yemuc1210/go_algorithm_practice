package NC46

import "sort"

/*
算法计划  递归回溯

给定一个数组 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中  只能使用一次。

注意：解集不能包含重复的组合。
*/
func combinationSum2(candidates []int, target int) (ans [][]int) {
    sort.Ints(candidates)
    var freq [][2]int
    for _, num := range candidates {
        if freq == nil || num != freq[len(freq)-1][0] {
            freq = append(freq, [2]int{num, 1})
        } else {
            freq[len(freq)-1][1]++
        }
    }

    var sequence []int
    var dfs func(pos, rest int)
    dfs = func(pos, rest int) {
        if rest == 0 {
            ans = append(ans, append([]int(nil), sequence...))
            return
        }
        if pos == len(freq) || rest < freq[pos][0] {
            return
        }

        dfs(pos+1, rest)

        most := min(rest/freq[pos][0], freq[pos][1])
        for i := 1; i <= most; i++ {
            sequence = append(sequence, freq[pos][0])
            dfs(pos+1, rest-i*freq[pos][0])
        }
        sequence = sequence[:len(sequence)-most]
    }
    dfs(0, target)
	sort.Slice(ans, func(i, j int) bool {
		n1,n2 := len(ans[i]),len(ans[j])
		minLen := min(n1,n2)
		k := 0
		for ;k<minLen;k++{
			if ans[i][k] != ans[j][k] {
				return ans[i][k] < ans[j][k]
			}
		}
		if minLen == n1 {
			return i<j
		}
		return j<i
	})
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// package main

// import (
// 	"fmt"
// 	"sort"
// )
// func main() {
// 	ans := [][]int{
// 		{10,70},
// 		{20,60},
// 		{10,20,50},
// 		{10,10,60},
// 	}
// 	sort.Slice(ans, func(i, j int) bool {
// 		n1,n2 := len(ans[i]),len(ans[j])
// 		minLen := min(n1,n2)
// 		k := 0
// 		for ;k<minLen;k++{
// 			if ans[i][k] != ans[j][k] {
// 				return ans[i][k] < ans[j][k]
// 			}
// 		}
// 		if minLen == n1 {
// 			return i<j
// 		}
// 		return j<i
// 	})
// 	fmt.Println(ans)
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }