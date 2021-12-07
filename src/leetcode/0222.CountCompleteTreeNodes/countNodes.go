package lt222

import (
	"sort"
	"structs"
)
type TreeNode = structs.TreeNode
/*中
222. 完全二叉树的节点个数
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
*/
/*直接bfs
bfs完全可，不过或许可以利用一下 完全二叉树的特性
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	resCnt := 0
	queue := []*TreeNode{root}
	for len(queue) > 0{
		curNode := queue[0]
		queue = queue[1:]  
		resCnt ++
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}else{
			//该节点是最后叶子节点的父亲节点
			//可以根据父亲节点的序号 推测最后一层的数量 节省时间
			//尤其是最后一次节点一般很多  
			//还有一种情况 最后一次是满的 或者没有单个孩子的节点  麻烦，直接bfs
		}
	}
	return resCnt
}

/*
利用完全二叉树的特性
寻找最后叶子节点的父亲节点
根据节点个数范围的上下界得到当前需要判断的节点个数 k，如果第 k 个节点存在，
则节点个数一定大于或等于 k，如果第 k 个节点不存在，则节点个数一定小于 k，
由此可以将查找的范围缩小一半，直到得到节点个数。

如何判断第 k 个节点是否存在呢？
如果第 k 个节点位于第 h 层，则 k 的二进制表示包含 h+1 位，
其中最高位是 1，其余各位从高到低表示从根节点到第 k 个节点的路径，
0 表示移动到左子节点，1 表示移动到右子节点。
通过位运算得到第 k 个节点对应的路径，判断该路径对应的节点是否存在，
即可判断第 k 个节点是否存在。

*/
func countNodes1(root *TreeNode) int {
	if root == nil {
        return 0
    }
    level := 0
    for node := root; node.Left != nil; node = node.Left {
        level++
    }
    return sort.Search(1<<(level+1), func(k int) bool {
        if k <= 1<<level {
            return false
        }
        bits := 1 << (level - 1)
        node := root
        for node != nil && bits > 0 {
            if bits&k == 0 {
                node = node.Left
            } else {
                node = node.Right
            }
            bits >>= 1
        }
        return node == nil
    }) - 1
}