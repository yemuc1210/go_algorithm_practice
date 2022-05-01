package lt1221

/*
1221. 分割平衡字符串
在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。

给你一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。

注意：分割得到的每个字符串都必须是平衡字符串。

返回可以通过分割得到的平衡字符串的 最大数量 。



括号匹配问题：R进栈  L出栈  栈空匹配 res++
*/
func balancedStringSplit(s string) int {
	res := 0

	num := 0
	for i:=0;i<len(s);i++ {
		if s[i] == 'R' {
			num ++
		}
		if s[i] == 'L' {
			num--
		}
		if num == 0{
			res ++
		}
	}
	return res
}