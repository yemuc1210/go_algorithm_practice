package NC43

/**
 * 没有重复项数字的所有排列
 * @param num int整型一维数组 
 * @return int整型二维数组
*/
func permute( nums []int ) [][]int {
    // write code here
	if len(nums)==1 || len(nums)==0{
		return [][]int{nums}
	}
	res := [][]int{}
	used := make([]bool, len(nums))
	
	var dfs func(idx int, path []int) 
	dfs = func(idx int, path []int) {
		if idx == len(nums) {
			res = append(res, path)
			return 
		}
		for i:=0;i<len(nums);i++{
			if !used[i] {
				used[i] = true
				path = append(path, nums[i])
				dfs(idx+1, path)
				path = path[:len(path)-1]
				used[i] =false
			}
		}
	}
	dfs(0,[]int{})
	return res
}