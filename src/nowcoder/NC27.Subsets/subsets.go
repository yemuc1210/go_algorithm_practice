package NC27

/**
 * 没有重复元素的整数集合 求子集
 * @param A int整型一维数组 
 * @return int整型二维数组
*/
func subsets( nums []int ) (ans [][]int) {
    // write code here
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

func subsets1(nums []int) [][]int {
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