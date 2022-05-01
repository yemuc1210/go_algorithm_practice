package NC159

import (
	"container/list"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 最小生成树
 * 返回最小的花费代价使得这n户人家连接起来
 * @param n int n户人家的村庄
 * @param m int m条路
 * @param cost int二维数组 一维3个参数，表示连接1个村庄到另外1个村庄的花费的代价
 * @return int
 */
func miniSpanningTree( n int ,  m int ,  cost [][]int ) int {
    // 科鲁斯卡尔算法  prim
	getPoints := make(map[int]int)
	edgeList := list.New()
	sort.Slice(cost, func(i, j int) bool {
		return cost[i][2] < cost[j][2]
	})
	for _,data := range cost{
		edgeList.PushBack(data)
	}
	res := 0 
	res += cost[0][2]
	getPoints[cost[0][0]] ++
	getPoints[cost[0][1]] ++

	for {
		for i:=1;i<edgeList.Len();i++{
			// nextVal := edgeList.Front().Value[0]
			
			// if getPoints[edgeList.Front().Value]{}
		}
	}
	return res
}