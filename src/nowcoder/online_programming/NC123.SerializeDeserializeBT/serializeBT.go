package NC123
import . "structs"
import (
	// . "structs"
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
 */
func Serialize( root *TreeNode ) string {
    // write code here
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
func Deserialize( data string ) *TreeNode {
    // write code here
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