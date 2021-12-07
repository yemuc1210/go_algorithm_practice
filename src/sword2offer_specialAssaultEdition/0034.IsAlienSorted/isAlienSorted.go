package offerS34

/*
外星语：英文小写字母，但是顺序不同
给定单词和字母表顺序，只有当给定的单词在按字典序排列时，返回true

*/
func isAlienSorted(words []string, order string) bool {
	rec := make([]int,26)
	for i,c := range order{ //字符相对大小
		rec[c-'a'] = i
	}
	for i:=0;i<len(words)-1;i++{
		w1,w2 := words[i],words[i+1]
		comm := true // 是否前缀串
		for j:=0;j<min(len(w1),len(w2));j++{
            if w1[j]!=w2[j]{ // 字符不等
                if rec[w1[j]-'a']>rec[w2[j]-'a']{ // 前面的串小于后面的串直接返回
                    return false
                }
                comm = false // 一个串不是另一个串的前缀串
                break
            }
        }
        if comm&&len(w1)>len(w2){ // 后一个字符串为前一个字符串的前缀串返回false
            return false  
        }
    }
    return true

}
func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}

