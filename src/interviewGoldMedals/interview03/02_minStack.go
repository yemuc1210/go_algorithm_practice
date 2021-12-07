package interview03

/*
包含最小函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
调用 min、push 及 pop 的时间复杂度都是 O(1)。

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

*/
type MinStack struct {
	elements, min []int
	l             int
}


/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}


func (this *MinStack) Push(x int)  {
	this.elements = append(this.elements, x)
	if this.l == 0 {
		this.min = append(this.min, x)
	} else {
		min := this.GetMin()  //获得当前最小元素
		if x < min {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.l++
}


func (this *MinStack) Pop()  {
	this.l--
	this.min = this.min[:this.l]
	this.elements = this.elements[:this.l]
}


func (this *MinStack) Top() int {
	return this.elements[this.l-1]
}


func (this *MinStack) GetMin() int {
	return this.min[this.l-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */