package interview16
/*
面试题 16.01. 交换数字
编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。

异或 
a ^ a = 0
0 ^ b = b
*/
func swapNumbers(numbers []int) []int {
	a1,b1 := numbers[0],numbers[1]

	numbers[0] = a1^a1^b1
	numbers[1] = b1^b1^a1
	return numbers
}