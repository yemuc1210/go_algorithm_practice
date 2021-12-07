package structs

import "errors"

// // Stack 是用于存放 int 的 栈
// type Stack struct {
// 	nums []int
// }
type Stack []interface{}
// Push 把 n 放入 栈
func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

// Pop 从 s 中取出最后放入 栈 的值
func (stack *Stack) Pop() (interface{}, error)  {
    theStack := *stack
    if len(theStack) == 0 {
        return nil, errors.New("Out of index, len is 0")
    }
    value := theStack[len(theStack) - 1]
    *stack = theStack[:len(theStack) - 1]
    return value, nil
}
func (stack Stack) Top() (interface{}, error)  {
    if len(stack) == 0 {
        return nil, errors.New("Out of index, len is 0")
    }
    return stack[len(stack) - 1], nil
}
// Len 返回 s 的长度
func (stack Stack) Len() int {
    return len(stack)
}

// IsEmpty 反馈 s 是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
func (stack Stack) Cap() int {
    return cap(stack)
}
