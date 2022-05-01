package lt0045
/*

刷题计划     DP
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。

 

示例 1:

输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

贪心法，从后往前找
0055题
我们的目标是到达数组的最后一个位置，
因此我们可以考虑最后一步跳跃前所在的位置，该位置通过跳跃能够到达最后一个位置。

如果有多个位置通过跳跃都能够到达最后一个位置，
那么我们应该如何进行选择呢？
直观上来看，我们可以「贪心」地选择距离最后一个位置最远的那个位置，
也就是对应下标最小的那个位置。
因此，我们可以从左到右遍历数组，选择第一个满足要求的位置。

找到最后一步跳跃前所在的位置之后，我们继续贪心地寻找倒数第二步跳跃前所在的位置，
以此类推，直到找到数组的开始位置。
*/
func jump(nums []int) int {
	index := len(nums)-1
	if index==0{
		return 1
	}
	steps:=0
	for index>0{
		for i:=0;i<index;i++{
			if i+nums[i] >= index{
				index = i
				steps++
				break
			}
		}
	}
	return steps
}

//贪心  正向找 每次找可到达的最远距离
func jump1(nums []int)int{
	length := len(nums)
	end:=0
	maxPosition := 0
	steps:=0
	for i:=0;i<length-1;i++{
		maxPosition = max(maxPosition,i+nums[i])
		if i==end{
			end = maxPosition
			steps++
		}
	}
	return steps
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}