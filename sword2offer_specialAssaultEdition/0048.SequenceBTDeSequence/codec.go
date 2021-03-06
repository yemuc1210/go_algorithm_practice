package offerS48
import (
	"strconv"
	"strings"
	"go_practice/structs"
)
type TreeNode = structs.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
层次遍历   BFS   注意nil值
*/
type Codec struct {
    
}

func Constructor() (_ Codec) {
    return 
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    //BFS
	q := []*TreeNode{root}
	res := []string{}
	for len(q)!=0{
		cur := q[0]
		q=q[1:]
		if cur != nil{
			res = append(res, strconv.Itoa(cur.Val))
			q = append(q, cur.Left)
			q = append(q, cur.Right)   //即便是nil 也加入
		}else{
			res = append(res, "X")
		}
	}
	return strings.Join(res,",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    //BFS
	/*
	先转换数组，遇到空值跳过
	*/
	if data == "X" {
		return nil
	}
	list := strings.Split(data,",")
	Val,_ := strconv.Atoi(list[0])
	root := &TreeNode{Val: Val}
	q := []*TreeNode{root}
	cursor := 1
	for cursor < len(list) {
		cur := q[0]
		q = q[1:]
		leftVal := list[cursor]
		rightVal := list[cursor+1]
		if leftVal != "X" {
			v,_ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: v}
			cur.Left = leftNode
			q = append(q, leftNode)
		}
		if rightVal != "X"{
			v,_ := strconv.Atoi(rightVal)
			rightNode := &TreeNode{Val: v}
			cur.Right = rightNode
			q = append(q, rightNode)
		}
		cursor += 2
	}
	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
