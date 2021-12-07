package offer31
/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
*/

//8ms 84.70%            3.8MB 75.56%
func validateStackSequences(pushed []int, popped []int) bool {
	//使用辅助栈
	//初始化辅助
	stack := []int{}
	//遍历压栈序列
	for _,v:=range pushed{
		stack = append(stack, v)
		//若元素=出栈栈顶元素， 就出栈
		for popped[0]==stack[len(stack)-1]{
			popped=popped[1:]
			stack=stack[:len(stack)-1]
			if len(stack)==0{
				break
			}
		}
	}
	return len(popped)==0
}