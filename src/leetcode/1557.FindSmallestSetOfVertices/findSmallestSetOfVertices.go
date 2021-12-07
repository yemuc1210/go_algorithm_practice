package lt1557
/*
给定有向无环图  节点编号0~n-1
边数组 edge[i] = [from , to]   有向边
找到最小的点集，使得从这些点可以到达图中所有的点。
题目保证解存在且唯一

找到入度为0的集合
由于给定的图是有向无环图，基于上述分析可知，对于任意入度不为零的节点 x，一定存在另一个节点 z，使得从节点 z 出发可以到达节点 x。
为了获得最小的点集，只有入度为零的节点才应该加入最小的点集。


*/
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	endPointsMap := make(map[int]bool)   //入度是否为0
	res := []int{}

	for i:=0;i<len(edges);i++{
		if _,ok := endPointsMap[edges[i][1]]; !ok {
			endPointsMap[edges[i][1]] = true
		}
	}
	for i:=0;i<n;i++{
		if _,ok := endPointsMap[i]; !ok {
			//入度为0
			res = append(res, i)
		}
	}
	return res

}
// class Solution:
	// def findSmallestSetOfVertices(self, n: int, edges: List[List[int]]) -> List[int]:
	// 	a = set()
	// 	b = set()
	// 	for x, y in edges:
	// 		a.add(x)
	// 		b.add(y)
	// 	return list(a - b)