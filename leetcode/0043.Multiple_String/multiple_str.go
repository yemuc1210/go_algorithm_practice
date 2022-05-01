package leetcode

import (
	"fmt"
	"math"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，
返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"
*/

/*
问题关键是整数的长度小于110，这会大于int64的取值范围
*/
func my_atoi(s string)int{
	//预处理，丢弃无用的前导空格inde
	res,index,n,factor :=0,0,len(s),1 //index标记有效的起始位置下标
	//读入下一个字符，直到到达下一个非数字字符 或 达到输入的末尾
	for index < n && s[index]>='0' && s[index]<='9' {
		res = 10 * res + int(s[index]-'0')
		if factor*res < math.MinInt64 {
			return math.MinInt32
		}else if factor*res >math.MaxInt64{
			return math.MaxInt32
		}
		index++

	}
	return res*factor

}
func multiply(num1 string, num2 string) string {
	//已知nums1 nums2非负
	//将字符串转换成整数，不需要处理前置空格
	n1:=my_atoi(num1)
	n2:=my_atoi(num2)

	return fmt.Sprintf("%v",n1*n2)
	
}
func multiply1(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    // num1[i] * num[j] 值必定在resArr[i+j] resArr[i+j+1]上，i+j+1存储地位
    resArr := make([]int, len(num1) + len(num2))
    for i := len(num2)-1; i >= 0; i -- {
        n2 := int(num2[i] - '0')
        for j := len(num1)-1; j >= 0; j -- {
            n1 := int(num1[j] - '0')
            sum := n2 * n1 + resArr[i+j+1]
            resArr[i+j+1] = sum % 10
            resArr[i+j] += sum / 10
        }
    }
    
    res := ""
    for k, v := range resArr {
        if k == 0 && v == 0 {
            continue
        }
        res += string(v + '0')
    }
    
    return res
}

