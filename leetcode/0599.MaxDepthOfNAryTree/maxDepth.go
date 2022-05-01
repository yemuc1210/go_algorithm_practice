package lt599

/**
 * Definition for a Node*/
type Node struct {
	Val      int
	Children []*Node
}

/*
层序遍历罢了
*/
func maxDepth(root *Node) (ans int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, node := range q {
			queue = append(queue, node.Children...)
		}
		ans++
	}
	return
}
