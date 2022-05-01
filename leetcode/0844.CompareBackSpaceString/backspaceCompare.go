package lt844

/*
给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。

双指针题目   退格表示删除之前的元素

双指针思路：
一个字符是否会被删掉，只取决于该字符后面的退格符，而与该字符前面的退格符无关。因此当我们  逆序   地遍历字符串，就可以立即确定当前字符是否会被删掉。

具体地，我们定义 skip 表示当前待删除的字符的数量。每次我们遍历到一个字符：

	若该字符为退格符，则我们需要多删除一个普通字符，我们让 skip 加 1；

	若该字符为普通字符：

		若 skip 为 0，则说明当前字符不需要删去；

		若 skip 不为 0，则说明当前字符需要删去，我们让 skip 减 1。

*/
func backspaceCompare(s string, t string) bool {
	skipS,skipT := 0,0
	i,j := len(s)-1, len(t)-1

	for i>=0 || j>=0 {
		for i>=0 {
			//找到第一个普通字符
			if s[i]=='#'{
				skipS ++
				i--
			}else if skipS > 0{//普通字符
				skipS--
				i--
			}else{
				break
			}
		}
		for j>=0 {
			//找到第一个普通字符
			if t[j]=='#'{
				skipT ++
				j--
			}else if skipT > 0{//普通字符
				skipT--
				j--
			}else{
				break
			}
		}
		//比较
		if i>=0 && j>=0 {
			if s[i] != t[j] {
				return false
			}
		}else if i>=0 || j>=0 {
			return false
		}
		i--
		j--
	}
	return true	
}