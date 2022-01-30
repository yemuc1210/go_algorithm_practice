package NC97

import (
	"sort"
	"strconv"
)

/**
 * return topK string
 * @param strings string字符串一维数组 strings
 * @param k int整型 the k
 * @return string字符串二维数组
 */
func topKstrings( strings []string ,  k int ) [][]string {
    // 返回出现频次前k的字符串机器次数   map  答案按频次从高到低排序 相同频率按字典序排列（小优先）
	m := make(map[string]int)   
	res := [][]string{}
	//遍历记录
	for _,str := range strings {
		m[str] ++
	}
	//对map数据排序 得到前k个
	type pair struct {
		val string
		cnt int
	}
	pairs := []pair{}
	for k,v := range m {
		pairs = append(pairs, pair{k,v})
	}
	//定义排序规则
	sort.Slice(pairs,func(i, j int) bool {
		if pairs[i].cnt == pairs[j].cnt {  //根据字典序排列
			return pairs[i].val < pairs[j].val
		}
		return pairs[i].cnt > pairs[j].cnt   //降序排列
	})
	//得到结果  不用考虑k>len(pairs)的情况吧
	for i:=0;i<len(pairs);i++{
		res = append(res, []string{pairs[i].val,strconv.Itoa(pairs[i].cnt)})
		k -- 
		if k == 0 {
			break  //此时获得k个了
		}
	}
	return res
}