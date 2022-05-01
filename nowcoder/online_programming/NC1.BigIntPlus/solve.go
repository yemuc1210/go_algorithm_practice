package nc1

import "strconv"

/** lt415
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve( s string ,  t string ) string {
    // write code here
	add := 0   //进位
    ans := ""
	//字符串仅由字符‘0’-‘9’构成 故数字是非负的
    for i, j := len(s) - 1, len(t) - 1; i >= 0 || j >= 0 || add != 0; i, j = i - 1, j - 1 {
        var x, y int
        if i >= 0 {
            x = int(s[i] - '0')
        }
        if j >= 0 {
            y = int(t[j] - '0')
        }
        result := x + y + add
        ans = strconv.Itoa(result%10) + ans
        add = result / 10
    }
    return ans
}