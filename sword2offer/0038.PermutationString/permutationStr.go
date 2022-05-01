package offer38
/*
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

*/
// func permutation(s string) []string {
// 	if len(s) == 0 {
// 		return []string{}
// 	}
// 	used, res := make([]bool, len(s)), []string{}
// 	p :=[]byte{}
// 	generatePermutation(s, 0, p, &res, &used)
// 	return res
// }
// func generatePermutation(nums string, index int, p []byte, res *[]string, used *[]bool) {
// 	if index == len(nums) {
// 		var temp []byte
// 		copy(temp,p)
// 		*res = append(*res, string(temp))
// 		return
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if !(*used)[i] {
// 			(*used)[i] = true
// 			p = append(p, nums[i])
// 			generatePermutation(nums, index+1, p, res, used)
// 			p = p[:len(p)-1]
// 			(*used)[i] = false
// 		}
// 	}
// 	return
// }
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
