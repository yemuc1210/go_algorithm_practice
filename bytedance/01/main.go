package main

import (
	"fmt"
	"sort"
)

func main() {
	var n,k int
	fmt.Scanln(&n,&k)
	var s = make([]byte,n) 
	for i:=0;i<n;i++{
		fmt.Scanf("%c",&s[i])
	}
	// n-字符串长度，k-按按钮次数
	// 按按钮一次 得到'a' 按两次得到'b'
	// 选择某些字符，通过按按钮生成得到分数，每个字符得一分，每个字符只能选择一次
	// 贪心的思路，优先选择最小的字符
	sort.Slice(s, func(i, j int) bool {
		return s[i]<s[j]
	})

	// edcda   ->  acdde : a(1) c(3) e(5)  ;   a(1)c(3)d(4)   3分
	// dfs 每次判断当前字符敲还是不敲
	leftCnt := k
	choosen := make(map[byte]int)
	score := 0
	for i:=0;i<n;i++{
		needCnt := int(s[i]-'a'+1) 
		// 判断能否加入
		// fmt.Printf("当前字符：%c, 需要次数%d, 剩余次数%d, 该字符是否选择过%d\n",s[i],needCnt,leftCnt,choosen[s[i]])
		if needCnt<=leftCnt && choosen[s[i]]==0 {
			// 敲当前字符
			choosen[s[i]] ++
			leftCnt -= needCnt
			score ++
		}		
	}
	fmt.Println(score)
}