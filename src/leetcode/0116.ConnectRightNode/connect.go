package lt116
/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

题目本身希望我们将二叉树的每一层节点都连接起来形成一个链表。因此直观的做法我们可以对二叉树进行层次遍历，
在层次遍历的过程中将我们将二叉树每一层的节点拿出来遍历并连接。

层次遍历基于广度优先搜索，它与广度优先搜索的不同之处在于，广度优先搜索每次只会取出一个节点来拓展，
而层次遍历会每次将队列中的所有元素都拿出来拓展，这样能保证每次从队列中拿出来遍历的元素都是属于同一层的，
因此我们可以在遍历的过程中修改每个节点的 next 指针，同时拓展下一层的新队列。


*/
type Node struct {
	Val int
	Left *Node
	Right *Node
		Next *Node
}
func connect(root *Node) *Node {
    if root == nil {
        return root
    }

    // 初始化队列同时将第一层节点加入队列中，即根节点
    queue := []*Node{root}

    // 循环迭代的是层数
    for len(queue) > 0 {
        tmp := queue
        queue = nil

        // 遍历这一层的所有节点
        for i, node := range tmp {
            // 连接
            if i+1 < len(tmp) {
                node.Next = tmp[i+1]
            }

            // 拓展下一层节点
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    // 返回根节点
    return root
}

