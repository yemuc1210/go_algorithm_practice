package NC10

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 大数乘法 思路是模拟竖式乘法
 * @param s string字符串 第一个整数
 * @param t string字符串 第二个整数
 * @return string字符串
*/
func solve( num1 string ,  num2 string ) string {
    // write code here
	if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    // num1[i] * num[j] 值必定在resArr[i+j] resArr[i+j+1]上，i+j+1存储
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