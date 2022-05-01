package lt405

import "strings"

/*

405. 数字转换为十六进制数
给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。

注意:

1 十六进制中所有字母(a-f)都必须是小写。
2 十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；
对于其他情况，十六进制字符串中的第一个字符将不会是0字符。

3 给定的数确保在32位有符号整数范围内。
4 不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
*/
func toHex(num int) string {
	if num == 0 {
        return "0"
    }
    sb := &strings.Builder{}
    for i := 7; i >= 0; i-- {
        val := num >> (4 * i) & 0xf
        if val > 0 || sb.Len() > 0 {
            var digit byte
            if val < 10 {
                digit = '0' + byte(val)
            } else {
                digit = 'a' + byte(val-10)
            }
            sb.WriteByte(digit)
        }
    }
    return sb.String()
}

func toHex1(num int) string {
	if num == 0 {
		return "0"
	}
	helper := "0123456789abcdef"
	ans := ""
	for num != 0 && len(ans) < 8 {
		temp := num & 15
		ans = string(helper[temp]) + ans
		num >>= 4
	}
	return ans
}

