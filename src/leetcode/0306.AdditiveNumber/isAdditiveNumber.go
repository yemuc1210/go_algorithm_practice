package lt306

import (
	// "crypto/ecdsa"
	"strconv"
	"strings"
)

//累加数是一个字符串 组成它的数字形成累加序列
// 除了最开始两个 其余数是之前两个数之和
//fib序列   所以 字符串长度至少是2
//问题是 第一个数和第二个数的长度  dfs
//非法情况 以0开头

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	//两个数的边界 firstEnd secondEnd
	for firstEnd:=0; firstEnd<len(num)/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 0 {
			break   //非法情况
		}
		//第一个数
		first,_ := strconv.Atoi(num[:firstEnd+1])
		//遍历第二个数的各种情况
		for secondEnd := firstEnd+1; max(firstEnd, secondEnd-firstEnd) <= len(num)-secondEnd; secondEnd++{
			if num[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break        //第二个数以0开头
			}
			second,_ := strconv.Atoi(num[firstEnd+1:secondEnd+1])  //闭区间
			//判断 后面的字符串中是否以这两个的和为开头。
			if recursiveCheck(num, first,second,  secondEnd+1) {
				return true
			}
		}
	}
	return false
}
func recursiveCheck(num string, x1,x2 int, left int) bool {
	if left==len(num) {  //到头了都是合法的
		return true
	}
	// 后面的字符串中是否以和为开头。
	if strings.HasPrefix(num[left:],   strconv.Itoa(x1+x2)) {
		return recursiveCheck(num, x2, x1+x2, left+len(strconv.Itoa(x1+x2)))
	}
	//否则
	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

