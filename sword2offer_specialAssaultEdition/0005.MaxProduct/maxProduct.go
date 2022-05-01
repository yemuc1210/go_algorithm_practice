package offerS5

import "fmt"

/*
剑指 Offer II 005. 单词长度的最大乘积
给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，
它们长度的乘积的最大值。
假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。


leetcode 318
*/
func maxProduct(words []string) int {
	matrix := make([][]int,len(words))
	for i := range matrix{
		matrix[i] = make([]int, len(words))
		matrix[i][i] = 0
	}
	//matrix n*n   记录两个单词之间长度的乘积
	//更新
	max := 0
	for i:=0;i<len(words);i++{
		for j:=i+1;j<len(words);j++{
			if !hasSameChar(words[i],words[j]) {
				
				matrix[i][j] = len(words[i])*len(words[j])
				if matrix[i][j] > max{
					max = matrix[i][j]
				}
				fmt.Printf("not have same s1=%s  s2=%s  product=%d\n",words[i],words[j],matrix[i][j])
			}
		}
	}
	return max
}

//判断两个字符串是否有相同字符
func hasSameChar(a,b string) bool {
	//排序O(nlog(n))    遍历异或O(n^2)
	for i:=0;i<len(a);i++{
		for j:=0;j<len(b);j++{
			if a[i] ^ b[j] == 0 {
				return true
			}
		}
	}
	return false
}