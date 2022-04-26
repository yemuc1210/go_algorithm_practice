package main

import "sort"

// 账户合并
// accounts[i]: accounts[i][0]是名称,其余是emails
// 如果两个账户有一些相同得邮箱地址,则两个账户属于同一个人.
// 返回: 每个账户得第一个元素是名称,其余元素是按字母ASCII排序得邮箱
// 并查集: 为每个邮箱创建并查集,则包含相同邮箱的,属于同一个人
func accountsMerge(accounts [][]string) [][]string {
	// 首先,确定并查集的大小,需要知道一个有多少个邮箱地址
	emailIdx := make(map[string]int)
	emailName := make(map[string]string)   // email: name1


	// 遍历
	for _,account := range accounts {
		name := account[0]
		for _,email := range account[1:] {
			if _,has := emailIdx[email];!has {
				// 存入
				emailIdx[email] = len(emailIdx)
				emailName[email] = name
			}
		}
	}
	// 通过并查集进行合并, 遍历每个账户,对账户中得邮箱进行合并
	// 并查集存储得邮箱地址对应得编号,针对编号进行合并
	// 完成并查集得合并操作之后,即可知道有多少不同得账户.
	// 遍历所有邮箱地址,对每个邮箱地址,通过并查集得到该邮箱属于哪个合并后得账户,即可整理出每个账户包含哪些邮箱地址
	uf := NewUnionFind(len(emailIdx))
	for _,account := range accounts {
		firstIdx := emailIdx[account[1]]
		for _,email := range account[2:] {
			uf.Union(emailIdx[email],  firstIdx)
		}
	}
	idx2Email := map[int][]string{}
	for email,idx := range emailIdx {
		idx := uf.Find(idx) 
		idx2Email[idx]  = append(idx2Email[idx], email)
	}
	var res [][]string
	for _,emails := range idx2Email{
		sort.Strings(emails)
		account := append([]string{emailName[emails[0]]}, emails...)
		res = append(res,account)
	}
	return res
}

type UnionFind struct {
	parent []int
	// 用秩优化
	rank []int 
	count int
}
func NewUnionFind(n int ) UnionFind {
	u := UnionFind{
		count: n,
		parent: make([]int,n),
		rank: make([]int,n),
	}
	// 初始化parent
	for i:=range u.parent {
		u.parent[i] = i   //  i节点的父亲为其自身
	}
	return u 
}

// 实现基本的union  find 方法

// find: find时完成路径压缩
func (u *UnionFind) Find(p int) int {
	root := p
	for root!=u.parent[root] {  // 只有根的父亲是自己
		root = u.parent[root]
	}  // 从叶子向上迭代,找到树根
	// 此时找到根
	// 压缩路径
	for p != u.parent[p] {  // 效果是把路径上的都指向root
		// p不是根
		tmp := u.parent[p]
		u.parent[p] = root   // p直接指向根root
		p = tmp    // p向上移动,指向原先的parent
	}   // 完成路径压缩
	return root  
}
// Union: 包含p q的两个并查集合并;依靠秩觉得合并的主体
func (u *UnionFind) Union(p,q int) { 
	// 首先,find
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		// 同一个并查集
		return 
	}
	// 准备合并
	// 判断秩
	if u.rank[qRoot] > u.rank[pRoot] {
		u.parent[pRoot] = qRoot
	}else{ // <=   pRoot为新的根
		u.parent[qRoot] = pRoot    
		if u.rank[pRoot] == u.rank[qRoot] {
			u.rank[pRoot] ++ // ? u.rank[pRoot] += u.rank[qRoot]
		}
	}
	u.count --   // 并查集的个数
}
func (u *UnionFind) TotalCount()int{return u.count}

func main() {

}