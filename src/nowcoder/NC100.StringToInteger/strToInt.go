package NC100

import (
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 字符串转换成整数
 * myAtoI    lt8   32位
 * @param s string字符串
 * @return int整型
 */
func StrToInt( s string ) int {
    //预处理，丢弃无用的前导空格inde
	res,index,n,factor :=0,0,len(s),1 //index标记有效的起始位置下标
	for index < n && s[index] == ' '{
		index ++
	}
	//下面第一个字符是正负号，
	// fmt.Println(s[index])
	if index < n{
		if s[index] == '-'{
			factor = -1   //乘子，控制正负号
			index ++
		}else if s[index] =='+'{
			factor = 1
			index ++
		}
	}
	//读入下一个字符，直到到达下一个非数字字符 或 达到输入的末尾
	for index < n && s[index]>='0' && s[index]<='9' {
		res = 10 * res + int(s[index]-'0')
		if factor*res < math.MinInt32 {
			return math.MinInt32
		}else if factor*res >math.MaxInt32{
			return math.MaxInt32
		}
		index++

	}
	return res*factor
}