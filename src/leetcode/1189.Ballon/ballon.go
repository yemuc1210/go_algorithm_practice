package lt1189

func maxNumberOfBalloons(text string) int {
	m := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		m[text[i]]++
	}
	// 遍历 balloon的最小次数
	var res int
	for {
		if m['b'] >= 1 {
			m['b']--
		} else {
			break
		}
		if m['a'] >= 1 {
			m['a']--
		} else {
			break
		}
		if m['l'] >= 2 {
			m['l'] -= 2
		} else {
			break
		}
		if m['o'] >= 2 {
			m['o'] -= 2
		} else {
			break
		}
		if m['n'] >= 1 {
			m['n']--
		} else {
			break
		}
		res++
	}
	return res
}
