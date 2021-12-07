package offerS115
/*lt444   图 graph
剑指 Offer II 115. 重建序列
请判断原始的序列 org 是否可以从序列集 seqs 中唯一地 重建 。

序列 org 是 1 到 n 整数的排列，其中 1 ≤ n ≤ 104。
重建 是指在序列集 seqs 中构建最短的公共超序列，
即  seqs 中的任意序列都是该最短序列的子序列。

关键是 唯一
示例 1：
	输入: org = [1,2,3], seqs = [[1,2],[1,3]]
	输出: false
	解释：[1,2,3] 不是可以被重建的唯一的序列，因为 [1,3,2] 也是一个合法的序列。
*/
/*
	怎么用图 每个整数作为一个节点，序列给出的是有向边
	序列给的是排列，所以整数之间满足大小关系的
*/
func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	graph := make([][]bool, n+1)
	for i := range graph {
		graph[i] = make([]bool, n+1)
	}
	inDegrees := make([]int, n+1)   //每个节点的入度
	for i := range inDegrees {
		inDegrees[i] = -1
	}
	for _, seq := range seqs {   //遍历序列，构建图和入度
		for _, num := range seq {
			if num < 1 || num > n{
				return false
			}
			if inDegrees[num] == -1 {  //记录这个点是否在序列中出现
				inDegrees[num] = 0
			}
		}
		m := len(seq)  //序列的长度不定
		for i := 0; i < m-1; i++ {
			if graph[seq[i]][seq[i+1]] == false {
				graph[seq[i]][seq[i+1]] = true  //构建图
				inDegrees[seq[i+1]]++  //入度
			}
		}
	}
	var queue []int
	var index int
	for i, d := range inDegrees[1:] {
		if d == 0 {
			if len(queue) == 0 && i+1 == org[index] {
				queue = append(queue, i+1)
				index++
			} else {
				return false
			}
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for v, b := range graph[u] {
			if b == true {
				inDegrees[v]--
				if inDegrees[v] == 0 {
					if len(queue) == 0 && v == org[index] {
						queue = append(queue, v)
						index++
					} else {
						return false
					}
				}
			}
		}
	}
	return index == n
}
