package lcp35

import (
	"math"
	"sort"
)

/*
LCP 35. 电动车游城市
小明的电动车电量充满时可行驶距离为 cnt，每行驶 1 单位距离消耗 1 单位电量，且花费 1 单位时间。
小明想选择电动车作为代步工具。地图上共有 N 个景点，景点编号为 0 ~ N-1。
他将地图信息以 [城市 A 编号,城市 B 编号,两城市间距离] 格式整理在在二维数组 paths，
表示城市 A、B 间存在双向通路。初始状态，电动车电量为 0。
每个城市都设有充电桩，charge[i] 表示第 i 个城市每充 1 单位电量需要花费的单位时间。
请返回小明最少需要花费多少单位时间从起点城市 start 抵达终点城市 end。


(city,charge)视为节点建图   （start,0)作为起点  Dijkstra
*/
type path struct {
	next int
	dist int
}

type node struct {
	cost  int
	index int
	power int
}

type NodePriorityQueue struct {
	nodes []node
}

func (n *NodePriorityQueue) Len() int {
	return len(n.nodes)
}

func (n *NodePriorityQueue) Less(i, j int) bool {
	if n.nodes[i].cost < n.nodes[j].cost {
		return true
	}
	return false
}

func (n *NodePriorityQueue) Swap(i, j int) {
	tmp := n.nodes[i]
	n.nodes[i] = n.nodes[j]
	n.nodes[j] = tmp
}

func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	size := len(charge)
	dp := make([][]int, size)
	for i:=0; i<size; i++ {
		dp[i] = make([]int, cnt+1)
		for j:=0; j<cnt+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[start][0] = 0

	p := make([][]path, size)
	for _,v := range paths {
		p[v[0]] = append(p[v[0]], path{next: v[1], dist: v[2]})
		p[v[1]] = append(p[v[1]], path{next: v[0], dist: v[2]})
	}
	queue := &NodePriorityQueue{}
	queue.nodes = append(queue.nodes, node{cost: 0, index: start, power: 0})

	for queue.Len() > 0 {
		n := queue.nodes[0]
		queue.nodes = queue.nodes[1:]
		if n.cost > dp[n.index][n.power] {
			continue
		}
		if n.index == end {
			return n.cost
		}

		if n.power < cnt {
			nt := n.cost + charge[n.index]
			if nt < dp[n.index][n.power+1] {
				dp[n.index][n.power+1] = nt
				queue.nodes = append(queue.nodes, node{cost: nt, index: n.index, power: n.power+1})
				sort.Sort(queue)
			}
		}

		for _,v := range p[n.index] {
			if n.power >= v.dist && n.cost+v.dist < dp[v.next][n.power-v.dist] {
				dp[v.next][n.power-v.dist] = n.cost+v.dist
				queue.nodes = append(queue.nodes, node{index: v.next, power: n.power-v.dist, cost: n.cost+v.dist})
				sort.Sort(queue)
			}
		}
	}


	return -100000
}