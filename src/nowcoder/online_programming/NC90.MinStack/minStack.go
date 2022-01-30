package NC90

type MinStack struct {
	elements, min []int
	l             int
}
var minStack = MinStack{make([]int, 0), make([]int, 0), 0}


func Push(node int) {
    // write code here
	minStack.elements = append(minStack.elements, node)
	if minStack.l == 0 {
		minStack.min = append(minStack.min, node)
	}else{
		min := Min()
		if node<min{
			minStack.min = append(minStack.min, node)
		}else{
			minStack.min = append(minStack.min, min)
		}
	}
	minStack.l ++
}
func Pop() {
    // write code here
	minStack.l--
	minStack.min = minStack.min[:minStack.l]
	minStack.elements = minStack.elements[:minStack.l]
}
func Top() int {
    // write code here
	return minStack.elements[minStack.l-1]
}
func Min() int {
    // write code here
	return minStack.min[minStack.l-1]
}