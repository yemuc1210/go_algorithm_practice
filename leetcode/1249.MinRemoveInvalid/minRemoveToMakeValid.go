package lt1249
/*
最容易想到的思路是利用栈判断括号匹配是否有效。这个思路可行，时间复杂度也只是 O(n)。
不用栈，可以 2 次循环遍历，正向遍历一次，标记出多余的 '(' ，逆向遍历一次，再标记出多余的 ')'，最后将所有这些标记多余的字符删掉即可。这种解法写出来的代码也很简洁，时间复杂度也是 O(n)。
针对上面的解法再改进一点。正向遍历的时候不仅标记出多余的 '('，还可以顺手把多余的 ')' 删除。这样只用循环一次。最后再删除掉多余的 '(' 即可。时间复杂度还是 O(n)。
*/
func minRemoveToMakeValid(s string) string {
	res, opens := []byte{}, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			opens++
		} else if s[i] == ')' {
			if opens == 0 {
				continue
			}
			opens--
		}
		res = append(res, s[i])
	}
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == '(' && opens > 0 {
			opens--
			res = append(res[:i], res[i+1:]...)
		}
	}
	return string(res)
}