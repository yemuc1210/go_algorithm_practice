package interview03

/*
01 三合一。描述如何只用一个数组来实现三个栈。

使用最前或最后4个元素存储容量、栈顶坐标
*/
type TripleInOne struct {
	arr []int
}


func Constructor1(stackSize int) TripleInOne {
	arr := make([]int, stackSize*3 + 4)
	copy(arr,  []int{4,5,6,stackSize*3 + 3})
	return TripleInOne{arr: arr}
}


func (this *TripleInOne) Push(stackNum int, value int)  {
	if this.arr[stackNum] <= this.arr[3] {
		this.arr[this.arr[stackNum]] = value
		this.arr[stackNum] += 3
	}
}


func (this *TripleInOne) Pop(stackNum int) int {
	if this.arr[stackNum] > 6 {
		this.arr[stackNum] -= 3
		return this.arr[this.arr[stackNum]]
	}
	return -1
}


func (this *TripleInOne) Peek(stackNum int) int {
	if this.arr[stackNum] <= 6 {
		return -1
	}
	return this.arr[this.arr[stackNum] - 3]
}


func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.arr[stackNum] <= 6
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */

//  type TripleInOne struct {
// 	data     []int
// 	stackPtr [3]int
// }

// func Constructor(stackSize int) TripleInOne {
// 	return TripleInOne{
// 		make([]int, stackSize*3),
// 		[3]int{0, stackSize, stackSize * 2},
// 	}
// }

// func (this *TripleInOne) Push(stackNum int, value int) {
// 	if this.IsFull(stackNum) {
// 		return
// 	}
// 	this.data[this.stackPtr[stackNum]] = value
// 	this.stackPtr[stackNum]++
// }

// func (this *TripleInOne) Pop(stackNum int) int {
// 	if this.IsEmpty(stackNum) {
// 		return -1
// 	}
// 	ele := this.data[this.stackPtr[stackNum]-1]
// 	this.stackPtr[stackNum]--
// 	return ele
// }

// func (this *TripleInOne) Peek(stackNum int) int {
// 	if this.IsEmpty(stackNum) {
// 		return -1
// 	}
// 	return this.data[this.stackPtr[stackNum]-1]
// }

// func (this *TripleInOne) IsEmpty(stackNum int) bool {
// 	switch stackNum {
// 	case 0:
// 		return this.stackPtr[0] == 0
// 	case 1:
// 		return this.stackPtr[1] == len(this.data)/3
// 	case 2:
// 		return this.stackPtr[2] == (len(this.data)/3)*2
// 	}
// 	return true
// }

// func (this *TripleInOne) IsFull(stackNum int) bool {
// 	switch stackNum {
// 	case 0:
// 		return this.stackPtr[0] == len(this.data)/3
// 	case 1:
// 		return this.stackPtr[1] == (len(this.data)/3)*2
// 	case 2:
// 		return this.stackPtr[2] == len(this.data)
// 	}
// 	return true
// }

