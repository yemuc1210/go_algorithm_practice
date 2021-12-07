package interview04

import "structs"

type ListNode = structs.ListNode


/*
面试题 04.03. 特定深度节点链表
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（
比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

示例：

输入：[1,2,3,4,5,null,7,8]

        1
       /  \ 
      2    3
     / \    \ 
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]

层次遍历
*/
func listOfDepth(tree *TreeNode) []*ListNode {

	if tree == nil {
		return []*ListNode{}
	}

	res := []*ListNode{}

	queue := []*TreeNode{}
	queue = append(queue, tree)

	for len(queue)>0{
		curSize := len(queue)   //当前层的尺寸
		//存储下一层节点
		nextLevel := []*TreeNode{}

		head := new(ListNode)
		curNode := head

		//遍历当前层
		for i:=0;i<curSize;i++{
			curNode.Next = &ListNode{Val: queue[i].Val}
			curNode = curNode.Next  //后移
			//左右节点入队
			if queue[i].Left != nil{
				nextLevel = append(nextLevel, queue[i].Left)
			}

			if queue[i].Right != nil{
				nextLevel = append(nextLevel, queue[i].Right)
			}
		}

		queue = nextLevel   
		res = append(res, head.Next)
	}
	return res
}