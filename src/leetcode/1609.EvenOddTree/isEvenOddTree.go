package lt1609

import (
	// "fmt"
	"math"
	. "structs"
)

/*
奇偶树：
如果一棵二叉树满足下述几个条件，则可以称为 奇偶树 ：
	二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推。
	偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
	奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减

解法：
	bfs
*/
func isEvenOddTree(root *TreeNode) bool {
    q := []*TreeNode{root}
    for level := 0; len(q) > 0; level++ {
        prev := 0
        if level%2 == 1 {
            prev = math.MaxInt32
        }
        size := len(q)
        for _, node := range q {
            val := node.Val
            if val%2 == level%2 || level%2 == 0 && val <= prev || level%2 == 1 && val >= prev {
                return false
            }
            prev = val
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        q = q[size:]
    }
    return true
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/even-odd-tree/solution/qi-ou-shu-by-leetcode-solution-l7bw/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。