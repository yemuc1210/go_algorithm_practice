package main

import (
	"fmt"
)

func main() {
	var n,k int
	fmt.Scanln(&n,&k)
	var s = make([]byte,n) 
	freq := make([]int,26)
	for i:=0;i<n;i++{
		fmt.Scanf("%c",&s[i])
		freq[int(s[i]-'a')] ++
	}
	score := 0
	for i:=0;i<26;i++{
		if freq[i] > 0 && k>=i+1 {
			k -= i+1
			score ++
		}
	}
	fmt.Println(score)
}