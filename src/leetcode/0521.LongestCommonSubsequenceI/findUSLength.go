package main

import "reflect"

// 最长特殊序列
// a有b没有的子序列是特殊子序列
func findUSlength(a,b string) int {
	// a==b return -1
	if reflect.DeepEqual(a,b) {
		return -1
	}
	// 长度不同
	// 那么长度长的本身，就不可能是短的子序列
	if len(a) > len(b) {
		return len(a)
	}else{
		return len(b)
	}
}


func main() {

}