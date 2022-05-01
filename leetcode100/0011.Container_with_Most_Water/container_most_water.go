package leetcode


/*
给定非负整数，表示在坐标轴上的一个虚拟高度的墙，两堵墙构成容器，x轴之间的距离为底，min{ai,aj}为高
求最大的容器

Input: [1,8,6,2,5,4,8,3,7]
Output: 49     【8：7】  中间长度7
*/

func find_most_container(height []int)int{
	
	//只需要记录最大值即可，相对来说比较容易
	max, start, end := 0,0, len(height)-1
	high,width := 0,0

	/*
	从左往右，或者从右往左的视角
	每次都要寻找能个
	*/
	for start < end {
		width = end - start
		//high = min{h[s],h[end]}  //不过不返回具体下标啊
		if height[start] < height[end]{
			high = height[start]
			start ++
		}else{
			high = height[end]
			end --
		}

		if width * high > max {
			max = width * high
		}
	}
	return max
}