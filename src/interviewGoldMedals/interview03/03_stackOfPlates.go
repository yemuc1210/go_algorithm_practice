package interview03
/*
面试题 03.03. 堆盘子
堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。
请实现数据结构SetOfStacks，模拟这种行为。
SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。
此外，SetOfStacks.push()和SetOfStacks.pop()应该与普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 

进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。

当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.

*/
type StackOfPlates struct {
	cap int //容量
	plates [][]int  //栈数组
}


func Constructor(cap int) StackOfPlates {
	return  StackOfPlates{cap: cap,  plates: make([][]int, 0)}
}


func (this *StackOfPlates) Push(val int)  {
	if this.cap == 0{
		return
	}
	if len(this.plates) == 0 {
		//当前无盘子
		newPlates := make([]int,0)    //新堆
		newPlates = append(newPlates, val)
		this.plates = append(this.plates, newPlates)
		return
	}

	//否则，已经有若干堆盘子
	lastPlates := this.plates[len(this.plates)-1] 
	if len(lastPlates) == this.cap {
		//最后一摞已经满了
		newPlates := make([]int,0)    //新堆
		newPlates = append(newPlates, val)
		this.plates = append(this.plates, newPlates)
		return
	}
	//否则 可以插入
	lastPlates = append(lastPlates, val)
	this.plates[len(this.plates)-1] = lastPlates   //必须加，为啥呢，不是数组引用吗
}


func (this *StackOfPlates) Pop() int {
	return this.PopAt(len(this.plates)-1)
}


func (this *StackOfPlates) PopAt(index int) int {
	if len(this.plates)==0 || index >= len(this.plates) {
		return -1
	}
	plates := this.plates[index]   //取某一摞
	val := plates[len(plates)-1]
	plates = plates[:len(plates)-1]  //弹出
	this.plates[index] = plates
	if len(plates) == 0{
		//删除
		this.plates = append(this.plates[:index], this.plates[index+1:]...)
	}
	return val
}


/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */