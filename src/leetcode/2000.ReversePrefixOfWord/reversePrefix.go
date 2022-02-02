package lt2000

import "fmt"

/*
2000. 反转单词前缀
给你一个下标从 0 开始的字符串 word 和一个字符 ch 。
找出 ch 第一次出现的下标 i ，反转 word 中从下标 0 开始、直到下标 i 结束（含下标 i ）的那段字符。
如果 word 中不存在字符 ch ，则无需进行任何操作。

例如，如果 word = "abcdefd" 且 ch = "d" ，那么你应该 反转 从下标 0 开始、直到下标 3 结束（含下标 3 ）。
结果字符串将会是 "dcbaefd" 。
返回 结果字符串 。
*/
func reversePrefix(word string, ch byte) string {
	if len(word) == 0 || len(word)==1{
		return word
	}
	// 找到ch
	idx := 0 
	for idx < len(word) && word[idx]!=ch {idx++}
	// fmt.Println(idx)
	// 此时idx 指向ch  进行逆置
	// 可能出现没有ch的情况
	if idx == len(word) {
        return word
    }
	sub := []byte(word[0:idx+1]) 
	
	word = word[idx+1:]
	fmt.Println(string(sub),word)
	// 对sub逆置
	n := len(sub)
	for i:=0;i<n/2;i++ {
		sub[i],sub[n-i-1] = sub[n-i-1],sub[i]
	}
	sub1 := string(sub)
	fmt.Println(sub1)
	word = sub1 + word
	return word
}