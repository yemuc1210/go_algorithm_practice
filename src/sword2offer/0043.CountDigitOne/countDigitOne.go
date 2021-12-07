package offer43

import (
	"math"
	"strconv"
)

/*
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

*/
func countDigitOne(n int) int {
	if n == 0{
		return 0
	}
	if n<10 {
		return 1
	}

	s := strconv.Itoa(n)
	length := len(s)

	base := (length - 1) *(int(math.Pow(10,float64(length)-2)))

	high := int(s[0] - '0')

	cur := int(math.Pow(10,float64(length)-1))  //当前数量级

	if high == 1 {
		//最高位为1，1+n-cur就是1000~1234中由千位数提供的1的个数，剩下的f函数就是求1000~1234中由234产生的1的个数
		return base+1+n-cur + countDigitOne(n-high*cur)
	}else{
		return base *high +cur + countDigitOne(n - high*cur)
	}

}