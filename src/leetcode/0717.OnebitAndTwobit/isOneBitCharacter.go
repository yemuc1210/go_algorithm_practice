package main

import "fmt"

/*
有两种特殊字符：

第一种字符可以用一比特 0 表示
第二种字符可以用两比特（10 或 11）表示
给你一个以 0 结尾的二进制数组 bits ，如果最后一个字符必须是一个一比特字符，则返回 true 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/1-bit-and-2-bit-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isOneBitCharacter(bits []int) bool {
	// 1. 最后一个需要是0
	// 2. 考虑最后一个0前面连续1的个数  偶数为true 奇数为false
	if bits[len(bits)-1] != 0 {
		return false
	}
	n := len(bits)
	cntOne := 0
	for i:=n-2;i>=0;i-- {
		//找到最后一个0
		if bits[i] == 1 {
			cntOne ++
		}else{
			break
		}
	}
	return cntOne&1 == 0
}

func main() {
	res := isOneBitCharacter([]int{1,0,0})
	fmt.Println(res)
}