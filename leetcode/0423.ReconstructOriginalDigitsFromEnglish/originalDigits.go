package lt423

import (
	"strconv"
	"strings"
)

/*每日一题 1124
423. 从英文中重建数字
给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
*/
/*
这特么 既要找原来的数字（单词 one two...)还要升序返回
标签：哈希表、数字、字符串
zero one two three four five six seven eight nine
所以字符：e g f i h o n s r u t w v x z 15个字符
	e:one three five seven eight nine zero 9
	g:eight 1
	f:four five 2
	i:five six eight nine 4
	h:eight 1
	o:one two four zero 4
	n:one seven nine 4
	s:six seven 2
	r:three four zero 3
	u:four 1
	t:two three eight 3
	w:two 1
	v:seven 1
	x:six 1
	z: zero 1
*/
/*
从字符频次唯一的入手：出现g h 则必有eight  出现u必有four 出现w必有two 出现v必有seven
	出现x必有six 出现z 必有zero

*/
func originalDigits(s string) string {

	n0:=strings.Count(s,"z")   //zero   z唯一
	n2:=strings.Count(s,"w")   //two  w唯一
	n8:=strings.Count(s,"g")   //eight -g/h
	n6:=strings.Count(s,"x")   // six  x唯一
	n3:=strings.Count(s,"t")-n2-n8
	n4:=strings.Count(s,"r")-n3-n0
	n7:=strings.Count(s,"s")-n6
	n1:=strings.Count(s,"o")-n4-n2-n0
	n5:=strings.Count(s,"v")-n7
	n9:=strings.Count(s,"i")-n8-n6-n5

	ns := []int{n0,n1,n2,n3,n4,n5,n6,n7,n8,n9}
	res := ""
	for i,n := range ns {
		res += strings.Repeat(strconv.Itoa(i),n)
	}
	return res
}



// ns = (n0,n1,n2,n3,n4,n5,n6,n7,n8,n9)

// return "".join((str(i)*n for i, n in enumerate(ns)))