package main

import "strings"

// 395. 至少有 K 个重复字符的最长子串
// 给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串，
// 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

// tag:滑动窗口、分治、哈希
//提示： s 仅由小写字母组成    1<=k<=10^5

func longestSubstring(s string, k int) int {
	if s=="" {
		return 0
	}
	cnt := [26]int{}
    for _, ch := range s {
        cnt[ch-'a']++   // 计数
    }
	var res int
    var split byte
    for i, c := range cnt[:] {
        if 0 < c && c < k {
            split = 'a' + byte(i)
            break
        }
    }
    if split == 0 {
        return len(s)
    }

    for _, subStr := range strings.Split(s, string(split)) {
        res = max(res, longestSubstring(subStr, k))
    }
    return res

}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}