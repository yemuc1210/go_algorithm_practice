package interview03
/*

面试题 03.04. 化栈为队
使用两个栈实现一个队列

leetcode 232
剑指offer 9
*/



/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

 type MyQueue struct {
	Stack *[]int
	Queue *[]int
}

/** Initialize your data structure here. */
func Constructor4() MyQueue {
	tmp1, tmp2 := []int{}, []int{}
	return MyQueue{Stack: &tmp1, Queue: &tmp2}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	*this.Stack = append(*this.Stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	popped := (*this.Queue)[len(*this.Queue)-1]
	*this.Queue = (*this.Queue)[:len(*this.Queue)-1]
	return popped
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	return (*this.Queue)[len(*this.Queue)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*this.Stack)+len(*this.Queue) == 0
}

func (this *MyQueue) fromStackToQueue(s, q *[]int) {
	for len(*s) > 0 {
		popped := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]

		*q = append(*q, popped)
	}
}
