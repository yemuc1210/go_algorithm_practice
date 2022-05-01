package offerS82

import "sort"

/*
剑指 Offer II 082. 含有重复元素集合的组合
给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
*/
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	//为了避免重复使用元素 元素只能使用一次 排序
	sort.Ints(candidates)
	var freq [][2]int  //key-val 【0】：key  [1]: freq
	for _,num := range candidates{
		if freq == nil || num != freq[len(freq)-1][0] {  //和上一个元素不同，num非重复元素
			freq = append(freq, [2]int{num,1})
		}else{
			freq[len(freq)-1][1] ++ 
		}
	}
	//下面进行遍历  dfs
	var sequence []int  //单个序列 答案
	var dfs func(pos,rest int) 
	dfs  = func(pos, rest int) {//pos当前位置 rest：剩余量
		if rest == 0{//找到一组答案
			res = append(res, append([]int(nil), sequence...))
		}
		if pos == len(freq) || rest < freq[pos][0]{
			return  //越界和不可能情况
		}	
		dfs(pos+1,  rest)  //跳过当前元素
		
		most := min(rest/freq[pos][0],   freq[pos][1])
		for i:=1;i<=most;i++{ //相同的数放在一起处理
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1,  rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0,target)
	return res
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}