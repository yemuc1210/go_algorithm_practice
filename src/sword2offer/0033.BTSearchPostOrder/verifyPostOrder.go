package offer33

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

注意   数组长度 <= 1000

此外，判断的二叉搜索树
搜索树特点，左子树上的值都小于根，右子树都大于根

*/
//0ms  100%    2MB  63.02%
func verifyPostorder(postorder []int) bool {
	if len(postorder)<=2{
		return true
	}
	//递归查找 最后一个元素是根
	root :=postorder[len(postorder)-1]
	left := -1
	
	for i:=len(postorder)-2; i>=0; i--{
		if postorder[i] < root {
			left = i
			break
		}
	}
	//找到第一个小于根的了

	//判断，如果左子树有大于根的，U报错
	for _,v := range postorder[:left+1] {
		if v > root {
			return false
		}
	}
	return verifyPostorder(postorder[0:left+1]) && verifyPostorder(postorder[left+1:len(postorder)-1])

}