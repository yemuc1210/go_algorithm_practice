package lt684
/*lt684  图 中
剑指 Offer II 118. 多余的边
树可以看成是一个连通且 无环 的 无向 图。

给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。
*/
/*

- 给出一个连通无环无向图和一些连通的边，要求在这些边中删除一条边以后，
	图中的 N 个节点依旧是连通的。如果有多条边，输出最后一条。
- 这一题可以用并查集直接秒杀。依次扫描所有的边，把边的两端点都合并 `union()` 到一起。
	如果遇到一条边的两端点已经在一个集合里面了，就说明是多余边，删除。最后输出这些边即可。
*/

// UnionFind defind
// 路径压缩 + 秩优化
type UnionFind struct {
	parent, rank []int
	count        int
}

// Init define
func (uf *UnionFind) Init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

// Find define
func (uf *UnionFind) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	// compress path
	for p != uf.parent[p] {
		tmp := uf.parent[p]
		uf.parent[p] = root
		p = tmp
	}
	return root
}

// Union define
func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot
	} else {
		uf.parent[qroot] = proot
		if uf.rank[proot] == uf.rank[qroot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}

// TotalCount define
func (uf *UnionFind) TotalCount() int {
	return uf.count
}

// UnionFindCount define
// 计算每个集合中元素的个数 + 最大集合元素个数
type UnionFindCount struct {
	parent, count []int
	maxUnionCount int
}

// Init define
func (uf *UnionFindCount) Init(n int) {
	uf.parent = make([]int, n)
	uf.count = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
		uf.count[i] = 1
	}
}

// Find define
func (uf *UnionFindCount) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	return root
}

// 不进行秩压缩，时间复杂度爆炸，太高了
// func (uf *UnionFindCount) union(p, q int) {
// 	proot := uf.find(p)
// 	qroot := uf.find(q)
// 	if proot == qroot {
// 		return
// 	}
// 	if proot != qroot {
// 		uf.parent[proot] = qroot
// 		uf.count[qroot] += uf.count[proot]
// 	}
// }

// Union define
func (uf *UnionFindCount) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if proot == len(uf.parent)-1 {
		//proot is root
	} else if qroot == len(uf.parent)-1 {
		// qroot is root, always attach to root
		proot, qroot = qroot, proot
	} else if uf.count[qroot] > uf.count[proot] {
		proot, qroot = qroot, proot
	}

	//set relation[0] as parent
	uf.maxUnionCount = max(uf.maxUnionCount, (uf.count[proot] + uf.count[qroot]))
	uf.parent[qroot] = proot
	uf.count[proot] += uf.count[qroot]
}

// Count define
func (uf *UnionFindCount) Count() []int {
	return uf.count
}

// MaxUnionCount define
func (uf *UnionFindCount) MaxUnionCount() int {
	return uf.maxUnionCount
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findRedundantConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	uf, res := UnionFind{}, []int{}
	uf.Init(len(edges) + 1)
	for i := 0; i < len(edges); i++ {
		if uf.Find(edges[i][0]) != uf.Find(edges[i][1]) {
			uf.Union(edges[i][0], edges[i][1])
		} else {
			res = append(res, edges[i][0])
			res = append(res, edges[i][1])
		}
	}
	return res
}
