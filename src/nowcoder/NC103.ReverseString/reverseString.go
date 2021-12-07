package NC103

/**
 * 反转字符串
 * @param str string字符串 
 * @return string字符串
*/
func solve( str string ) string {
    // write code here
	//存在一个问题 string类型是不可变量
	s := []byte(str)
	n := len(s)
	for i:=0;i<n/2;i++{
		s[i],s[n-i-1] = s[n-i-1],s[i]
	}
	return string(s)
}