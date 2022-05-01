package interview17

import (
	"fmt"
	"strconv"
	"strings"
)

/*
面试题 17.07. 婴儿名字
每年，政府都会公布一万个最常见的婴儿名字和它们出现的频率，也就是同名婴儿的数量。
有些名字有多种拼法，例如，John 和 Jon 本质上是相同的名字，但被当成了两个名字公布出来。
给定两个列表，一个是名字及对应的频率，另一个是本质相同的名字对。
设计一个算法打印出每个真实名字的实际频率。
注意，如果 John 和 Jon 是相同的，并且 Jon 和 Johnny 相同，则 John 与 Johnny 也相同，
	即它们有传递和对称性。

在结果列表中，选择 字典序最小 的名字作为真实名字。
*/
/*
根据synonyms 得到传递关系  构建边 节点有权值
图 查找连通子图的数量
	使用邻接数组：len(names)*len(names)
	使用邻接表； 有点复杂
选择字典序最小的名字作为真实名字

并查集
*/
// type Node struct{
// 	name string //名称
// 	count int  //频次
// }
// func trulyMostPopular(names []string, synonyms []string) []string {
// 	if len(names) == 0 {
// 		return nil
// 	}
// }
func trulyMostPopular(names []string, synonyms []string) []string {
	if len(names) <= 1 {
		return names
	}
	idMap := make(map[string]int, 0)
	var uf UnionFind = NewMapUnionFindWithRank(0, RankSize)
	for _, synonym := range synonyms {
		a, b := getSynonym(synonym)
		if _, ok := idMap[a]; !ok {
			idMap[a] = len(idMap)
		}
		if _, ok := idMap[b]; !ok {
			idMap[b] = len(idMap)
		}
		uf.Union(idMap[a], idMap[b])
	}

	type tmp struct {
		index int
		name  string
	}
	rootMap := make(map[int]*tmp, 0)
	result := []string{}
	for _, v := range names {
		name, count := getNameAndCount(v)
		if _, ok := idMap[name]; !ok {
			result = append(result, v)
		} else {
			rootID := uf.Find(idMap[name])
			if _, ok := rootMap[rootID]; !ok {
				rootMap[rootID] = &tmp{
					index: len(result),
					name:  name,
				}
				result = append(result, v)
			} else {
				if name < rootMap[rootID].name {
					rootMap[rootID].name = name
				}
				rootIndex := rootMap[rootID].index
				vName, vCount := getNameAndCount(result[rootIndex])
				vName, vCount = rootMap[rootID].name, vCount+count
				result[rootIndex] = formatName(vName, vCount)
			}
		}
	}
	return result
}

func getSynonym(s string) (string, string) {
	arr := strings.Split(s[1:len(s)-1], ",")
	return arr[0], arr[1]
}

func getNameAndCount(s string) (string, int) {
	arr := strings.Split(s[:len(s)-1], "(")
	count, _ := strconv.Atoi(arr[1])
	return arr[0], count
}

func formatName(name string, count int) string {
	return fmt.Sprintf("%s(%d)", name, count)
}


//RankType 按秩合并类型
type RankType int

//按秩合并类型
const (
	RankNone   RankType = iota //不按秩合并
	RankHeight                 //按树的高度合并
	RankSize                   //按树的节点数合并
)

//UnionFind 并查集
type UnionFind interface {
	Union(x, y int) bool
	Find(x int) int
	Rank(x int) int
	Count() int
	IsConnected(x, y int) bool
}


//MapUnionFind 并查集（map存储）
type MapUnionFind struct {
	parent   map[int]int
	rank     map[int]int
	count    int
	rankType RankType
}

//NewMapUnionFind 初始化并查集
func NewMapUnionFind() *MapUnionFind {
	uf := &MapUnionFind{
		parent: make(map[int]int, 0),
	}
	return uf
}

//NewMapUnionFindWithRank 初始化带秩的并查集
func NewMapUnionFindWithRank(n int, rankType RankType) *MapUnionFind {
	uf := &MapUnionFind{
		parent:   make(map[int]int, n),
		rank:     make(map[int]int, n),
		rankType: rankType,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = -1
		uf.rank[i] = 1
	}
	return uf
}

//Union 合并两个节点,按秩合并
func (s *MapUnionFind) Union(x, y int) bool {
	rootX, rootY := s.Find(x), s.Find(y)
	if rootX != rootY {
		switch s.rankType {
		case RankNone:
			s.parent[rootX] = rootY
		case RankHeight:
			if s.rank[rootX] == s.rank[rootY] {
				s.rank[rootY]++
			} else if s.rank[rootX] > s.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			s.parent[rootX] = rootY
		case RankSize:
			if s.rank[rootX] >= s.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			s.parent[rootX] = rootY
			s.rank[rootY] += s.rank[rootX]
		}
		s.count--
		return true
	}
	return false
}

//Find 路径压缩
func (s *MapUnionFind) Find(x int) int {
	if _, ok := s.parent[x]; !ok {
		s.parent[x] = x
		s.count++
	}
	if x != s.parent[x] {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

//IsConnected 节点是否连通
func (s *MapUnionFind) IsConnected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}

//Count ..
func (s *MapUnionFind) Count() int {
	return s.count
}

//Rank ..
func (s *MapUnionFind) Rank(x int) int {
	return s.rank[x]
}

 