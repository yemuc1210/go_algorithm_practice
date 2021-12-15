package lt851
/*
851. 喧闹和富有
有一组 n 个人作为实验对象，从 0 到 n - 1 编号，其中每个人都有不同数目的钱，
以及不同程度的安静值（quietness）。
为了方便起见，我们将编号为 x 的人简称为 "person x "。

给你一个数组 richer ，其中 richer[i] = [ai, bi] 表示 person ai 比 person bi 更有钱。
另给你一个整数数组 quiet ，其中 quiet[i] 是 person i 的安静值。
richer 中所给出的数据 逻辑自恰
（也就是说，在 person x 比 person y 更有钱的同时，不会出现 person y 比 person x 更有钱的情况 ）。

现在，返回一个整数数组 answer 作为答案，
其中 answer[x] = y 的前提是，在所有拥有的钱肯定不少于 person x 的人中，person y 是最安静的人（也就是安静值 quiet[y] 最小的人）。
*/

/*
根据richer可以得到财富间的大小关系  有向无环图的结构 
首先使用字典建立图的关系，然后找到最富的人，遍历BFS，不断使用富有的人，但是安静值更小的人更新子节点

2 拓扑排序   richer描述边关系  quiet是权值
用数组记录答案，初始时ans[i]=i，然后对原图进行拓扑排序
对于每一条边，如果发现quiet[ans[v]] > quiet[ans[u]] 则ans[v] = ans[u]
*/
func loudAndRich(richer [][]int, quiet []int) []int {
	edges := make([][]int, len(quiet))
	for i := range edges {
		edges[i] = []int{}
	}
	indegrees := make([]int, len(quiet))
	for _, edge := range richer {
		n1, n2 := edge[0], edge[1]
		edges[n1] = append(edges[n1], n2)
		indegrees[n2]++
	}
	res := make([]int, len(quiet))
	for i := range res {
		res[i] = i
	}
	queue := []int{}
	for i, v := range indegrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		nexts := []int{}
		for _, n1 := range queue {
			for _, n2 := range edges[n1] {
				indegrees[n2]--
				if quiet[res[n2]] > quiet[res[n1]] {
					res[n2] = res[n1]
				}
				if indegrees[n2] == 0 {
					nexts = append(nexts, n2)
				}
			}
		}
		queue = nexts
	}
	return res
}