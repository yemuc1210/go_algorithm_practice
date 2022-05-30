package main

import "go_practice/structs"

type TreeNode = structs.TreeNode
// 每个节点值都是0或1
// 根到叶的路径表示一个从最高位开始的二进制数
// 对于每一个叶子，找到根到叶子的路径所表示的数字，返回数字之和
// 问题:寻找路径 先序遍历
// 左右中  后序
func sumRootToLeaf(root *TreeNode) (ans int) {
    val, st := 0, []*TreeNode{}
    var pre *TreeNode
    for root != nil || len(st) > 0 {
        for root != nil {
			// val = val*2 + root.Val
            val = val<<1 | root.Val
            st = append(st, root)
            root = root.Left
        }
		// 此时root==nil
		// 访问了一条路径

		// 指向叶子节点
        root = st[len(st)-1]
        if root.Right == nil || root.Right == pre {
            if root.Left == nil && root.Right == nil {  // 叶子节点
                ans += val   // 更新：加上 一条路径的值
            }
            val >>= 1  // val  = val / 2    
            st = st[:len(st)-1]   // 出栈
            pre = root   // 更新pre   pre: 当前遍历树的根（左右中的中                          ）
            root = nil
        } else {
            root = root.Right
        }
    }
    return
}
