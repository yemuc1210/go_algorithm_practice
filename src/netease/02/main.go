package main

import "fmt"

var m map[int]int = make(map[int]int)
var arr []byte
var score []int = make([]int, 26)

func main() {
	var s string
	fmt.Scanln(&s)
	// 拿到字符
	arr = []byte(s)
	// 贪心法  每次选其中最大的一组.预处理时,计算出每个字符可能获得的最大分数(相邻字符的情况).
	// 不过每次选择过后,会对原来的结果造成影响
	// 三个三个一组来,计算1-2 2-3组合谁更大
	for i := 0; i < 26; i++ {
		score[i] = i + 1
	}
	fmt.Println(getMax(0))
}
func getMax(index int) int {
	if index >= len(arr)-1 {
		return 0
	}
	if _, ok := m[index]; ok {
		return m[index]
	}
	var cur int
	if abs(arr[index],arr[index+1]) <= 1 {
		pre := int(arr[index] - 'a')
		aft := int(arr[index+1] - 'a')
		cur = score[pre] + score[aft]
	}
	max1 := max(cur+getMax(index+2), getMax(index+1))
	m[index] = max1
	return max1

}
func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
func abs(a, b byte) int{
	if a > b {
		return int(a-'a') - int(b-'a')
	}
	return int(b-'a') - int(a-'a')
}
