package NC157

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 单调栈
 * 给定一个长度为 n 的可能含有重复值的数组 arr ，找到每一个 i 位置左边和右边离 i 位置最近且值比 arri 小的位置。
 * @param nums int一维数组 
 * @return int二维数组
*/

/*
描述
给定一个长度为 n 的可能含有重复值的数组 arr ，
找到每一个 i 位置左边和右边离 i 位置最近且值比 arri 小的位置。

请设计算法，返回一个二维数组，表示所有位置相应的信息。
位置信息包括：两个数字 l 和 r，如果不存在，则值为 -1，下标从 0 开始。

进阶：空间复杂度 O(n) ，时间复杂度 O(n)
*/
func foundMonotoneStack( nums []int ) [][]int {
    // write code here
	n := len(nums)
	res := make([][]int,n)
	for i:=range res{
		res[i] = make([]int, 2)
	}

	stack := []int{}
	//从左到右 一次入栈 保存从左到右的升序只
	for i:=0;i<n;i++{
		//如果栈里面栈比其大 pop
		for len(stack)>0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		//栈空 说明nums[i]左边没有比其小的
		if len(stack) == 0 {
			res[i][0] = -1
		}else{
			//若存在比其小的 则栈中第一个即为结果
			res[i][0] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	//从右往左再来一次
	stack = []int{}
	for i:=n-1;i>=0;i--{
		//如果栈里面栈比其大 pop
		for len(stack)>0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		//栈空 说明nums[i]左边没有比其小的
		if len(stack) == 0 {
			res[i][1] = -1
		}else{
			//若存在比其小的 则栈中第一个即为结果
			res[i][1] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}