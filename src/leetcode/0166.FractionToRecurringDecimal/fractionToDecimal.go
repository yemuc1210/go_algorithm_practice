package lt166

import (
	"errors"
	"fmt"

	"strconv"
)

/*
166. 分数到小数   medium
给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
如果小数部分为循环小数，则将循环的部分括在括号内。
如果存在多个答案，只需返回 任意一个 。
对于所有给定的输入，保证 答案字符串的长度小于 104 。

示例 1：
	输入：numerator = 1, denominator = 2
	输出："0.5"
示例 3：
	输入：numerator = 2, denominator = 3
	输出："0.(6)"
*/

//关键在于两数相除，只能是有限或无限循环，不存在无限不循环
//所以用set记录余数出现位置，出现了相同余数，则从出现位置到当前结尾为循环部分
func fractionToDecimal(numerator int, denominator int) string {
	//规范性约束  健壮性
	if denominator == 0{
		return ""// 分母不能为0
		defer errors.New("error, denominator can not be 0!")   //先return结束再运行该句，应该没毛病
	}

	if numerator ==0 {
		return "0"
	}

	if numerator % denominator == 0 {
		//整除
		return strconv.Itoa(numerator/denominator)
	}

	res := ""
	//首先追加正负号
	if numerator*denominator < 0 {
		res += "-"
	}
	//计算整数部分
	numerator,denominator = abs(numerator),abs(denominator)
	// fmt.Println(numerator,denominator)
	// fmt.Println(numerator/denominator)
	res += strconv.Itoa(numerator/denominator) + "."

	// fmt.Println(res)
	//计算小数部分
	numerator %= denominator   //余数
	m := make(map[int]int)

	for numerator != 0 { //辗转相除的过程
		//记录当前余数所在答案的位置
		m[numerator] = len(res)

		//模拟除法过程  余数补0  也就是×10
		numerator *= 10
		res += strconv.Itoa(numerator/denominator)
		numerator %= denominator
		// fmt.Println(numerator)
	}
	//判断当前的余数是否出现过
	if m[numerator]!=0 {
		u := m[numerator]  //得到开始位置
		return fmt.Sprintf("%s(%s)",  res[0:u],res[u:] )
	}
	return res

}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
// func isCirculated(a,b int) {
// 	//判断分母是否是2、5、10的N次幂

// }