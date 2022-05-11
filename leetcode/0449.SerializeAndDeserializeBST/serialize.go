package main

import (
	"go_practice/structs"
	"math"
	"sort"
	"strconv"
	"strings"
)
type TreeNode = structs.TreeNode
// 序列化和反序列化二叉搜索树
// BST中序有序，因此给定先序或后序遍历即可得到两个序列
// 先序+中序，后序+中序 都可以恢复一棵树

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{} 
}

// Serializes a tree to a single string.
// 序列化，后序遍历 因为后序遍历的最后一个是root
func (this *Codec) serialize(root *TreeNode) string {
    var res string
	var nums []string
	// var preOrder func(node *TreeNode)
	// preOrder = func(node *TreeNode) {
	// 	// 使用栈模拟递归
	// 	if node == nil {
	// 		return 
	// 	}
	// 	sk := []*TreeNode{node}
	// 	for len(sk)>0 {
	// 		// 左节点入队
	// 		for node.Left != nil {
	// 			sk = append(sk, node.Left)
	// 			node = node.Left
	// 		}
	// 		// 此时node指向最左节点，访问之
	// 		// nums = append(nums, node.Val)
	// 		num := strconv.Itoa(node.Val)
	// 		// 逗号分割
	// 		// res += num+","
	// 		nums = append(nums, num)
	// 		// 开始访问右节点
	// 		node = node.Right
	// 	}
	// }	
	// preOrder(root)
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return 
		}
		postOrder(node.Left)
		postOrder(node.Right)
		nums = append(nums, strconv.Itoa(node.Val))
	}	
	postOrder(root)
	res = strings.Join(nums,",")
	return res
}

// Deserializes your encoded data to tree.
// 反序列化
func (this *Codec) deserialize(data string) *TreeNode {    
	if data == "" {
		return nil
	}
	// 根据逗号分割
	arr := strings.Split(data,",")
	var construct func(lower,upper int) *TreeNode
	construct = func(lower, upper int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val,_ := strconv.Atoi(arr[len(arr)-1])  //root值  后序最后一个
		if val<lower || val>upper {
			// 越界  数组分割
			return nil
		}
		// 递归构建
		arr = arr[:len(arr)-1]
		return &TreeNode{Val:val, Right: construct(val,upper), Left: construct(lower,val)}
	}
	return construct(math.MinInt32,math.MaxInt32)
}


func main() {

}