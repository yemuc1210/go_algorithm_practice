package NC112

/**
 * 进制转换
 * @param M int整型 给定整数
 * @param N int整型 转换到的进制
 * @return string字符串
 */

/*
保留符号
2<=N<=16  用大写字母表示整数 A-10 B-11
M转换成N进制数
除以 N 取余，然后倒序排列，高位补零
*/
func solve( M int ,  N int ) string {
    // write code here
	m := abs(M)
	res := ""
	s := "0123456789ABCDEF";
	for m>0 {
		res += string(s[m%N])
		m /= N
	}

	if M<0 {
		res += "-"
	}
	//逆置res
	bytes := []byte(res)
	n := len(bytes)
	for i:=0;i<n/2;i++{
		bytes[i],bytes[n-i-1] = bytes[n-i-1],bytes[i]
	}
	return string(bytes)
}

func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}