package greedy

import "sort"

// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
// 对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
// 并且每块饼干 j，都有一个尺寸 s[j] 。
// 如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。
// 你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
func findContentChildren(g []int, s []int) int {
	// 分配：满足其胃口的最小尺寸
	sort.Ints(s)   // 饼干从小到大分配   尽可能多分配嘛，先将尺寸小的尽可能分配出去
	sort.Ints(g)   // 排序复杂度O(log(n))
	var numContent int
	var startIdxOfS int 
	for i:=0;i<len(g);i++ {
		for j:=startIdxOfS;j<len(s);j ++{
			if s[j] >= g[i] {
				// 分配s[j]
				// s[j] = 0
				startIdxOfS = j+1   
				numContent ++ 
				break
			}
		}
		// 判断是否有剩余饼干
		if startIdxOfS == len(s) {
			break  
		}
	}
	return numContent
}