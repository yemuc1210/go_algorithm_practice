package leetcode

import "sort"

//给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//首先统计每个字符的频次，然后排序，最后按照频次从高到低进行输出即可。
func frequencySort(s string) string {
	if s == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	for _, b := range sb {
		sMap[b]++
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}