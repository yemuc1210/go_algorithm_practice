package lt133
import "go_practice/structs"
type Node = structs.Node
/*中
133. 克隆图
给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。

图中的每个节点都包含它的值 val（int） 和其邻居的列表
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
/*
遍历+构造
使用哈希表记录遍历过的和克隆的节点 key-val:原始-克隆
从给定节点遍历，若节点被访问过，则返回val；若不在，可创建克隆节点 放入hash表

*/
func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}
	var traversal func(node *Node) *Node
	traversal = func(node *Node) *Node {
		if node == nil {
			return node
		}
		//若访问过
		if _,ok := visited[node]; ok {
			return visited[node]
		}
		//否则 克隆
		cloneNode := &Node{node.Val, []*Node{}}
		visited[node] = cloneNode

		//遍历该节点邻居列表 克隆
		for _,n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, traversal(n))
		}
		return cloneNode
	}
	return traversal(node)

}