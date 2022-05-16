package main

import (
	"fmt"
	"reflect"
	"sort"
)

func removeAnagrams(words []string) []string {
	// 0<i<words.length
	// words[i-1]和words[i]是字母异位词
	bs := [][]byte{}
	for _,word := range words {
		b := []byte(word)
		sort.Slice(b,func(i, j int) bool {
			return b[i]<b[j]
		})
		bs = append(bs, b)
	} 
	res := []string{}
	res = append(res, words[0])
	for i:=1;i<len(bs);i++{
		if reflect.DeepEqual(bs[i-1],bs[i]) {
			// 删除i
			continue
		}
		res = append(res, words[i])
	}
	return res
}

func main() {
	words :=[]string{
		"abba","baba","bbaa","cd","cd",
	}
	res := removeAnagrams(words)
	fmt.Println(res)
}