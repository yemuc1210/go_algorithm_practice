package NC31

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 第一个只出现一次的字符
 * @param str string字符串 
 * @return int整型
*/
func FirstNotRepeatingChar( str string ) int {
    // 遍历完才知道是否只出现一次   第一个出现的 不能排序，这会改变顺序
	m := make(map[byte]int)
	for i:=0;i<len(str);i++{
		m[str[i]] ++
	}
	//再次遍历
	for i:=0;i<len(str);i++{
		if m[str[i]] == 1{
			return i
		}
	}
	return -1
}