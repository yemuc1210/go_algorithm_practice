package main

import (
	"fmt"
	"strings"
)


func toGoatLatin(sentence string) string {
    // 分割
	words := strings.Split(sentence," ")
	res := []string{}
	for i,word := range words {
		// 判断首字母
		bs := []byte(word)
		c := bs[0]
		if c=='a' || c=='e' || c=='i' ||c=='o'||c=='u'||
		c=='A' || c=='E' || c=='I' ||c=='O'||c=='U'  {
			bs = append(bs, []byte("ma")...)
		}else{
			bs = append(bs[1:], append([]byte{c}, []byte("ma")...)...)
		}
		// 添加a
		for j:=0;j<=i;j++{
			bs = append(bs, 'a')
		}
		res = append(res, string(bs))
	}
	// 构建
	s := ""
	for i,w := range res {
		if i==len(res)-1{
			s+=w
			continue
		}
		s +=w+" "
	}
	return s
}

func main(){
	sentence:="I speak Goat Latin"
	s:=toGoatLatin(sentence)
	fmt.Println(s)
	fmt.Printf("%#v\n",s)
}