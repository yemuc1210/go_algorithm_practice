package NC115

/**
 * 栈排序
 * @param a int整型一维数组 描述入栈顺序
 * @return int整型一维数组
*/
func solve( a []int ) []int {
    // 只通过出栈和入栈操作得到降序序列    可以投机取巧
	if len(a) == 0 || len(a)==1{
		return a
	}
	//首先，新建一个栈，用于存放按顺序入栈的元素
	//然后建dp，dp[i]表示当前元素以及后面元素中的最大值。逆向遍历原序列，给dp赋值
	//正序遍历dp，将当前元素入栈，只要栈顶元素比后面元素都大，出栈，保证大的值尽量在序列前面
	//最后还在栈中的元素，输出
	n := len(a)
	res :=make([]int,n)
	stack := []int{}
	id :=0 
	dp := make([]int,n)
	dp[n-1] = a[n-1]
	for i:=n-2;i>=0;i--{
		dp[i] = Max(a[i],dp[i+1])
	}
	for i:=0;i<n;i++{
		//当前元素入栈
		stack = append(stack, a[i])
		//如果栈顶元素比后面元素都打，出栈
		for len(stack)>0 && i<n-1 && stack[len(stack)-1] > dp[i+1] {
			res[id] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			id ++
		}
	}
	//剩余按顺序
	for len(stack)>0 {
		res[id] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			id ++
	}
	return res
}
func Max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

/*
描述
给你一个 1 到 n 的排列和一个栈，并按照排列顺序入栈

你要在不打乱入栈顺序的情况下，仅利用入栈和出栈两种操作，输出字典序最大的出栈序列

排列：指 1 到 n 每个数字出现且仅出现一次

数据范围：  1≤n≤5×10^4，排列中的值都满足1≤val≤n
进阶：空间复杂度 O(n) ，时间复杂度 O(n)
*/