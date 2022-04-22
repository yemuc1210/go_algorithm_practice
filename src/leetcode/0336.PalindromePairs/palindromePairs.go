package main

import "fmt"

/*
给定一组 互不相同 的单词， \
找出所有 不同 的索引对 (i, j)，
使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。
*/
// 暴力搜索，尝试所有的组合，对每一种组合判断是否是回文
// 索引对(i,j)，只有两个
func palindromePairs(words []string) [][]int {
	// 计算排列    O(n^2)
	res := [][]int{}
	for i:=0;i<len(words);i++{
		for j:=0; j<len(words);j++{
			// 组合（i,j)
			if i==j {continue}

			s := words[i] + words[j]
			if isPalindrome1(s) {
				res = append(res, []int{i,j})
			}
		}
	}
	return res
}

func isPalindrome1(s string) bool {
	if len(s) == 1 {
		return true
	}
	// 使用栈判断是否是回文
	// 问题  aabbccdd这种没法识别
	// 解决：记录弹出次数，只能弹1次（连续弹出也算1次）
	// 双指针判断
	i,j := 0,len(s)-1
	for ;i<j && s[i]==s[j];i,j=i+1,j-1 {}
	// 判断i,j的位置关系
	// 偶数情况  i>j
	// 奇数回文： i==j
	return i>j ||i==j 
}
func palindromePairs1(words []string) [][]int {
    wordsRev := []string{}
    indices := map[string]int{}

    n := len(words)
    for _, word := range words {
        wordsRev = append(wordsRev, reverse(word))
    }
    for i := 0; i < n; i++ {
        indices[wordsRev[i]] = i
    }

    ret := [][]int{}
    for i := 0; i < n; i++ {
        word := words[i]
        m := len(word)
        if m == 0 {
            continue
        }
        for j := 0; j <= m; j++ {
            if isPalindrome(word, j, m - 1) {
                leftId := findWord(word, 0, j - 1, indices)
                if leftId != -1 && leftId != i {
                    ret = append(ret, []int{i, leftId})
                }
            }
            if j != 0 && isPalindrome(word, 0, j - 1) {
                rightId := findWord(word, j, m - 1, indices)
                if rightId != -1 && rightId != i {
                    ret = append(ret, []int{rightId, i})
                }
            }
        }
    }
    return ret
}

func findWord(s string, left, right int, indices map[string]int) int {
    if v, ok := indices[s[left:right+1]]; ok {
        return v
    }
    return -1
}

func isPalindrome(s string, left, right int) bool {
    for i := 0; i < (right - left + 1) / 2; i++ {
        if s[left + i] != s[right - i] {
            return false
        }
    }
    return true
}

func reverse(s string) string {
    n := len(s)
    b := []byte(s)
    for i := 0; i < n/2; i++ {
        b[i], b[n-i-1] = b[n-i-1], b[i]
    }
    return string(b)
}
func main() {
	words:=[]string{"a",""}
	res := palindromePairs1(words)
	fmt.Println(res)
}