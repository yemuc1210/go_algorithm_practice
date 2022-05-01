package lt720

import "strings"

func longestWord(words []string) string {
	m := make(map[string]bool)
	for _,word := range words {
		m[word] = true
	}
	var res  string
	
	next:for _,word := range words {
		for i:=1;i<len(word);i++ {
			if !m[word[:i]] {
				continue next
			}
		}
		if l1,l2:=len(word),len(res);l1>l2 || (l1==l2 && strings.Compare(word,res)<0) {
			res = word
		}
	}
	return res
}