package lt225
/*lt232 栈实现队列  简单
中 使用队列实现栈
*/
type MyStack struct {
	//定义两个队列
	enque []int
	deque []int
}

func Constructor() MyStack {
	return MyStack{[]int{},[]int{}}
}


func (this *MyStack) Push(x int)  {
	this.enque = append(this.enque, x)
}

func (this *MyStack) Pop() int {
	length := len(this.enque)
	for i:=0;i<length-1;i++{
		//转移
		this.deque = append(this.deque, this.enque[0]) //倾倒
		this.enque = this.enque[1:]  //出对
	}
	//最后一个元素
	res := this.enque[0]
	//交换指针
	this.enque = this.deque
	this.deque = nil

	return res
}


func (this *MyStack) Top() int {
	top := this.Pop()
	this.enque = append(this.enque, top)
	return top
}


func (this *MyStack) Empty() bool {
	if len(this.enque) == 0 {
		return true
	}

	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */