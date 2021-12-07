package offerS43

import "structs"
type TreeNode = structs.TreeNode
/*
剑指 Offer II 043. 往完全二叉树添加节点
完全二叉树是每一层（除最后一层外）都是完全填充
（即，节点数达到最大，第 n 层有 2n-1 个节点）的，并且所有的节点都尽可能地集中在左侧。

设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：

CBTInserter(TreeNode root) 使用根节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个新节点，节点类型为 TreeNode，值为 v 。使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
CBTInserter.get_root() 将返回树的根节点。
 

双端队列
*/

type CBTInserter struct {
	queue []*TreeNode

}


func Constructor(root *TreeNode) CBTInserter {
	rootNow := &CBTInserter{
		queue: []*TreeNode{root},
	}
	//插入左右节点 
	for i:=0;i<len(rootNow.queue);i++{
		if rootNow.queue[i].Left != nil{
			rootNow.queue = append(rootNow.queue, rootNow.queue[i].Left)
		}
		if rootNow.queue[i].Right != nil{
			rootNow.queue = append(rootNow.queue, rootNow.queue[i].Right)
		}
	}
	return *rootNow
}


func (this *CBTInserter) Insert(v int) int {
	n := len(this.queue)
	fatherNode := this.queue[(n-1) >> 1]   //÷2
	this.queue = append(this.queue, &TreeNode{Val: v})

	if n & 1 == 1{  //奇数
		fatherNode.Left = this.queue[n]  
	}else{
		fatherNode.Right = this.queue[n]
	}
	return fatherNode.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.queue[0]
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */