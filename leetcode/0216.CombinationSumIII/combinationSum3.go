package lt216

/*lt39  组合综述
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，
并且每种组合中不存在重复的数字。

说明：
	所有数字都是正整数。
	解集不能包含重复的组合。
示例 1:
输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
*/
/*
k个数 和为 n  限制了组合的长度为K
组合只允许出现1-9  不允许重复  所以组合的结果存在上界 1+2+...9=45
*/
func combinationSum3(k int, n int) [][]int {
	if k==0 || k>9 || n > 45{
		return [][]int{}
	}
	path, res := []int{}, [][]int{}
	findCombination(k,n,1,path,&res)
	return res
}

func findCombination(k,target,idx int, path []int, res *[][]int) {
	if target==0 {
		if len(path) == k {  //得到一组结果
			tmp := make([]int, k)
			copy(tmp, path)
			*res = append(*res, tmp)
		}
		return 
	}
	//正常遍历
	for i:=idx; i<10; i++{  //idx就是当前的数
		if target >= i {
			path = append(path, i)
			findCombination(k,target-i, i+1, path, res)
			path = path[:len(path)-1]  //跳过当前数 回溯
		}
	}
}