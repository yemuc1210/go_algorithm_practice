package main

import (
	"fmt"
	"sort"
	"strings"
)


func main() {
	var N int
	var name string
	var n int
	var err error
	for {
		n,err = fmt.Scanln(&N)
		if n==0 || err!=nil {
			break
		}
		for i:=0;i<N;i++ {
			fmt.Scanln(&name)
			//字母不区分大小写  先全部转化成小写  每个名字之间独立
			name = strings.ToLower(name)
			// //计算漂亮度  
			res := solve(name)
			fmt.Println(res)
		}
		//计算方法应该是，重复次数最高的分配较高值
		//zhangsan   an:26 25   51*2=102  102+24(z)+23(h)+22(g)+21(s) = 
		// lisi     26*2+25+24
	}
}

//计算字母的漂亮程度
func solve(s string) int {
	bytes := []byte(s)
	curBeauty := 26  //从最大的开始分配
	//先统计出现次数，然后根据次数排序
	type pair struct{
		ch byte
		count int
	}
	cnts := make([]pair, 26)  //计算26个字母出现的个数   
	for i:=0;i<len(bytes);i++{
		curChIdx := int(bytes[i]-'a')  //当前字符对应的位置
		cnts[curChIdx].ch = bytes[i]
		cnts[curChIdx].count ++    
	} 
	//然后 排序  根据count排序
	sort.Slice(cnts, func(i, j int) bool {
		return cnts[i].count > cnts[j].count  //降序
	})

	sum := 0   //结果
	//分配 先从次数最多的
	for i:=0;i<26;i++ {
		if cnts[i].count != 0 {
			sum += curBeauty*cnts[i].count
			curBeauty--
		}else{
			break  //后面的也都是0，没有计算必要
		}
	}
	return sum
}