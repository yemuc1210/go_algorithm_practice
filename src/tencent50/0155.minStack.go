package tencent50
/*
155. 最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
	push(x) —— 将元素 x 推入栈中。
	pop() —— 删除栈顶的元素。
	top() —— 获取栈顶元素。
	getMin() —— 检索栈中的最小元素。
 
*/
/*
最小元素 存在另一个数据结构
其他操作就是一般的栈
*/
type MinStack struct {
	elements,min []int
	length int
}


func Constructor155() MinStack {
	return MinStack{
		elements: []int{},
		min: []int{},
		length: 0,
	}
}

//最复杂的逻辑
func (this *MinStack) Push(val int)  {
	this.elements = append(this.elements, val)  //入栈
	if this.length == 0 {
		//初始时
		this.min = append(this.min, val)
	}else{  //找出最小值
		preMin := this.GetMin()  //   this.min[this.length-1]
		//更新最小值
		if val < preMin {
			this.min = append(this.min, val)
		}else{
			this.min = append(this.min, preMin)  //那么preMin这个值存了不止一次
		}
	}
	
	this.length++
}

//easy
func (this *MinStack) Pop()  {
	this.length -- 
	this.min = this.min[:this.length] //弹出最后一个
	this.elements = this.elements[:this.length]
}


func (this *MinStack) Top() int {
	return this.elements[this.length-1]
}


func (this *MinStack) GetMin() int {
	return this.min[this.length-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */