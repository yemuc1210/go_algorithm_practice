package offerS111
/*lt399   图
剑指 Offer II 111. 计算除法
给定一个变量对数组 equations 和一个实数值数组 values 作为已知条件，
其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。
每个 Ai 或 Bi 是一个表示单个变量的字符串。

另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，
请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。
如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

注意：输入总是有效的。可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
*/
/*
Dj = 0 边界
构建图 有向图 
节点表示equations中的某个变量   a->b的权值可以表示a/b  值从values数组来

最后query数组中的问题（边）可能是不存在的，就需要构建边  搜索dfs/bfs
*/


type stringUnionFind struct {
	parents map[string]string
	vals    map[string]float64
}

func (suf stringUnionFind) add(x string) {
	if _, ok := suf.parents[x]; ok {
		return
	}
	suf.parents[x] = x
	suf.vals[x] = 1.0
}

func (suf stringUnionFind) find(x string) string {
	p := ""
	if v, ok := suf.parents[x]; ok {
		p = v
	} else {
		p = x
	}
	if x != p {
		pp := suf.find(p)
		suf.vals[x] *= suf.vals[p]
		suf.parents[x] = pp
	}
	if v, ok := suf.parents[x]; ok {
		return v
	}
	return x
}

func (suf stringUnionFind) union(x, y string, v float64) {
	suf.add(x)
	suf.add(y)
	px, py := suf.find(x), suf.find(y)
	suf.parents[px] = py
	// x / px = vals[x]
	// x / y = v
	// 由上面 2 个式子就可以得出 px = v * vals[y] / vals[x]
	suf.vals[px] = v * suf.vals[y] / suf.vals[x]
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res, suf := make([]float64, len(queries)), stringUnionFind{parents: map[string]string{}, vals: map[string]float64{}}
	for i := 0; i < len(values); i++ {
		suf.union(equations[i][0], equations[i][1], values[i])
	}
	for i := 0; i < len(queries); i++ {
		x, y := queries[i][0], queries[i][1]
		if _, ok := suf.parents[x]; ok {
			if _, ok := suf.parents[y]; ok {
				if suf.find(x) == suf.find(y) {
					res[i] = suf.vals[x] / suf.vals[y]
				} else {
					res[i] = -1
				}
			} else {
				res[i] = -1
			}
		} else {
			res[i] = -1
		}
	}
	return res
}
