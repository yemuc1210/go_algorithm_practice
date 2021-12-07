/*
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/

package leetcode

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	//已知非空字符串，不过还是判断下吧
	if len(a) == 0 && len(b)==0{
		return ""
	}
	if len(a)==0{
		return b
	}
	if len(b)==0{
		return b
	}

	index_a := len(a) - 1
	index_b := len(b) - 1
	jinwei := 0
	result := ""
	for index_a >=0 && index_b >= 0 {
		ia := a[index_a] - '0' // -'0' 可以得到对应数字
		ib := b[index_b] - '0'
		sum := int(ia) + int(ib) + jinwei

		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum % 2 + '0'
		result = fmt.Sprintf("%c%s",res,result)
		index_a--
		index_b--
	}

	for index_a >= 0 {
		ia := a[index_a] - '0'
		sum := int(ia) + jinwei
		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum % 2 + '0'
		result = fmt.Sprintf("%c%s",res,result)
		index_a--
	}
	for index_b >= 0 {
		ib := b[index_b] - '0'
		sum := int(ib) + jinwei
		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum % 2 + '0'
		result = fmt.Sprintf("%c%s",res,result)
		index_b--
	}
	if jinwei == 1 {
		result = fmt.Sprintf("1%s",result)
	}
	return result

}

func addBinary_1(a string, b string) string {
	ai, _ := new(big.Int).SetString(a,2)
	bo, _ := new(big.Int).SetString(b,2)
	res := ai.Add(ai,bo)
	return res.Text(2)

}


