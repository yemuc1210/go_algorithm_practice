package NC121

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 字符串的排列  offer38
 * @param str string字符串 
 * @return string字符串一维数组
*/
func Permutation( str string ) []string {
    // write code here
	res := []string{}
	bytes := []byte(str)
	var dfs func(x int)
	dfs = func(x int) {
		if x==len(bytes)-1{
			res = append(res, string(bytes))
		}
		dict := map[byte]bool{}
		for i:=x;i<len(bytes);i++{
			if !dict[bytes[i]] {
				bytes[x],bytes[i] = bytes[i],bytes[x]
				dict[bytes[x]] = true
				dfs(x+1)
				bytes[x], bytes[i] = bytes[i], bytes[x]
			}
		}
	}
	dfs(0)
	return res
}

func permutation(s string) []string {
	res := []string{}
	bytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(bytes)-1 {
			res = append(res, string(bytes))
		}
		dict := map[byte]bool{}
		for i := x; i < len(bytes); i++ {
			if !dict[bytes[i]] {
				bytes[x], bytes[i] = bytes[i], bytes[x]
				dict[bytes[x]] = true
				dfs(x + 1)
				bytes[x], bytes[i] = bytes[i], bytes[x]
			}
		}
	}
	dfs(0)
	return res
}