package interview01
/*

面试题 01.02. 判定是否互为字符重排

给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

判断每个字符个数是否相等

排序、map

*/
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}
	m := make(map[byte]int)
	for i:=0;i<len(s1);i++ {
		m[s1[i]] ++
	}
	//比较是否每个元素的个数都相同
	for i:=0;i<len(s2);i++ {
		m[s2[i]] -- 
		// if m[s2[i]] >0 {
		// 	return false
		// }  
	}
	for _,v := range m{
		if v>0{
			return false
		}
	}
	return true

}