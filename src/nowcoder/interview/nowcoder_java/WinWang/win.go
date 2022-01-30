package main

import "fmt"

/*
小明，小王，小李三人正在进行一个游戏。
游戏有个回合，每个人都有一个字符串，三人的字符串长度相等。
每个回合他们必须更改自己字符串中的一个字母。
最后每个人的分数是字自己的字符串中出现字符最多的字母的数量。
分数最高者获胜，输出获胜者名字，小明获胜输出xiaoming，小王获胜输出xiaowang，小李获胜输出xiaoli，
如果有两个或者两个以上相同的最高分输出draw。

首先分别获取三人手中字符串中字母出现最高次数，每个回合+1，
也就是把别的字符改成出现次数最高的字母，当回合没有结束时，
字符串全是相同字母，那最高分数就是字符串长度。

*/
func main() {
	var n int
	fmt.Scanln(&n)
	var s1,s2,s3 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)

	// 每个人必须修改一次字母
	m  := len(s1)  // 字符串长度
	if m <= n {
		// 那么每个人可以修改每一个字母 最后结果应该是相同的 
		fmt.Println("draw")
	}else{
		// 不能修改/重置字母串 需要讲究策略
		// 修改n 次 相同字符最大数量

	}
}