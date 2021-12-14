package lt49

import (
	// "fmt"
	"sort"
)

/*
将每个字符串都排序，排序完成以后，相同 Anagrams 的字符串必然排序结果一样。
把排序以后的字符串当做 key 存入到 map 中。
遍历数组以后，就能得到一个 map，
	key 是排序以后的字符串，
	value 对应的是这个排序字符串以后的 Anagrams 字符串集合。
最后再将这些 value 对应的字符串数组输出即可。
*/


type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
func groupAnagrams(strs []string) [][]string {
	record,res := map[string][]string{},[][]string{}
	for _,str := range strs{
		sByte := []rune(str)
		//Sort sorts data. It makes one call to data.Len to determine n and 
		//O(n*log(n)) calls to data.Less and data.Swap. 
		//The sort is not guaranteed to be stable.
		sort.Sort(sortRunes(sByte))
		// fmt.Printf("sByte sorted:%v\n",sByte)
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for _,v := range record{
		res = append(res, v)
	}
	return res
}