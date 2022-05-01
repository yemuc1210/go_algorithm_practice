package lt78

/*
算法计划  递归、回溯、

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

 
示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/
func subsets(nums []int) [][]int {
	res := [][]int{}
	set := []int{}
	var dfs func(int)
    dfs = func(cur int) {
        if cur == len(nums)-1 {
            res = append(res, append([]int(nil), set...))
            return
        }
        set = append(set, nums[cur])
		dfs(cur+1)
		set = set[:len(set)-1]
		dfs(cur+1)
    }
	dfs(0)
	return res
}

/*
记原序列中元素的总数为 nn。原序列中的每个数字ai的状态可能有两种，
即「在子集中」和「不在子集中」。
我们用 1 表示「在子集中」，0 表示不在子集中，
那么每一个子集可以对应一个长度为 n 的 0/1 序列，第 ii 位表示ai是否在子集中。
例如，n=3 a={5,2,9} 时：0/1序列子集	0/1 序列对应的二进制数
000000	\{ \}{}	00
001001	\{ 9 \}{9}	11
010010	\{ 2 \}{2}	22
011011	\{ 2, 9 \}{2,9}	33
100100	\{ 5 \}{5}	44
101101	\{ 5, 9 \}{5,9}	55
110110	\{ 5, 2 \}{5,2}	66
111111	\{ 5, 2, 9 \}{5,2,9}	77
可以发现 0/1 序列对应的二进制数正好从 0 到 2^n - 1。
我们可以枚举 mask∈[0,2^n−1]，mask 的二进制表示是一个 0/1 序列，
我们可以按照这个0/1 序列在原集合当中取数。当我们枚举完所有 2^n个mask，
我们也就能构造出所有的子集。
*/
func subsets1(nums []int) (ans [][]int) {
    n := len(nums)
    for mask := 0; mask < 1<<n; mask++ {
        set := []int{}
        for i, v := range nums {
            if mask>>i&1 > 0 {
                set = append(set, v)
            }
        }
        ans = append(ans, append([]int(nil), set...))
    }
    return
}


// class Solution(object):
//     def subsets(self, nums):
//         res = [[]]
//         for i in range(len(nums)-1, -1, -1):
//             for subres in res[:]: res.append(subres+[nums[i]])
    
//         return res