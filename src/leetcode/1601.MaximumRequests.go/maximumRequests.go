package lt1601

func maximumRequests(n int, requests [][]int) int {
	var res int
	var zero int = n
	var chosenCnt int
	var delta = make([]int,n)
	// 回溯
	var backtrack func(idx int)
	// 当前回溯的是下标为idx的请求
	backtrack = func(idx int) {
		// 确定终止条件
		if idx == len(requests) {
			if zero == n && chosenCnt > res {
				res = chosenCnt
			}
			return 
		}
		// 单层判断
		// 对于第idx个请求，使用delta记录请求变化 
		// 若选择 delta[x]-1 delta[y]+1   即出度和入度 通过自身度的变化角度来衡量
		// 不选该请求
		backtrack(idx+1) 

		// 选择该请求
		// 维护delta中0的个数  0说明没有变化 是可行的
		z := zero
		chosenCnt ++
		request := requests[idx]
		x,y := request[0],request[1]
		if delta[x] == 0 {
			zero --  // 因为原先是0，现在要变化了
		}
		delta[x] --
		if delta[x] == 0 {
			zero ++ // 原先不是0，现在可是
		}
		// y 通用操作
		if delta[y] == 0 {
			zero --  // 因为原先是0，现在要变化了
		}
		delta[y] ++
		if delta[y] == 0 {
			zero ++ // 原先不是0，现在可是
		}
		// 进入下一层
		backtrack(idx+1) 
		// 状态回溯
		delta[y] --
		delta[x] ++
		chosenCnt -- 
		zero = z
	} 
	backtrack(0)
	return res
}