package lt839

/* 困难 lt839 图
剑指 Offer II 117. 相似的字符串
如果交换字符串 X 中的两个不同位置的字母，使得它和字符串 Y 相等，那么称 X 和 Y 两个字符串相似。
如果这两个字符串本身是相等的，那它们也是相似的。

例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)； "rats" 和 "arts" 也是相似的，
但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。

总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。
注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。
形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

给定一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串的一个 字母异位词 。
请问 strs 中有多少个相似字符串组？

字母异位词（anagram），一种把某个字符串的字母的位置（顺序）加以改换所形成的新词。
*/
/*
- 给出一个字符串数组，要求找出这个数组中，"不相似"的字符串有多少种。
相似的字符串的定义是：如果 A 和 B 字符串只需要交换一次字母的位置就能变成两个相等的字符串，那么 A 和 B 是相似的。
- 这一题的解题思路是用并查集。先将题目中的“相似”的定义转换一下，
	A 和 B 相似的意思是，A 和 B 中只有 2 个字符不相等，其他字符都相等，
	这样交换一次才能完全相等。
	有没有可能这两个字符交换了也不相等呢？
	这种情况不用考虑，因为题目中提到了给的字符串都是 `anagram` 的(`anagram` 的意思是，字符串的任意排列组合)。
	那么这题就比较简单了，只需要判断每两个字符串是否“相似”，如果相似就 `union()`，最后看并查集中有几个集合即可。
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
func numSimilarGroups(A []string) int {
	uf := UnionFind{}
	uf.Init(len(A))
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if isSimilar(A[i], A[j]) {
				uf.Union(i, j)
			}
		}
	}
	return uf.TotalCount()
}

func isSimilar(a, b string) bool {
	var n int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			n++
			if n > 2 {
				return false
			}
		}
	}
	return true
}

