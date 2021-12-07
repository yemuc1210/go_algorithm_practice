package tencent50
/*中
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。
*/
/*
动态规划
dp[i]  [0:i]的最大值   [i:j] (j-i)*min(h[i],h[j])
全局max
从两侧不断收缩，从两端开始是因为此时宽度是最大的
*/
func maxArea(height []int) int {
	max,start,end := 0,0, len(height)-1
	high,width := 0,0
	
	for start < end {
		width = end - start
		//更新high   min{}
		if height[start] < height[end] {
			high = height[start]
			start ++
		}else{
			high = height[end]
			end --
		}
		//计算
		if width*high > max {
			max = width * high
		}
	}
	return max
}