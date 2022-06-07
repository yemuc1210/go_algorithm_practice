package lt875

/*
剑指 Offer II 073. 狒狒吃香蕉
狒狒喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。
警卫已经离开了，将在 H 小时后回来。

狒狒可以决定她吃香蕉的速度 K （单位：根/小时）。
每个小时，她将会选择一堆香蕉，从中吃掉 K 根。
如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，
然后这一小时内不会再吃更多的香蕉，下一个小时才会开始吃另一堆的香蕉。

狒狒喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。
*/

//返回最小速度   二分查找最接近
func minEatingSpeed(piles []int, h int) int {
	//一次迭代代表一小时，从最多的那堆开始吃
	var eatingTime func(k int) int
	eatingTime = func(k int) int{   //速度k时耗时
		cnt := 0
		for _,pile := range piles{
			cnt += pile/k
			if pile%k!=0 {  //还有余量，一次吃不完
				cnt ++
			}
		}
		return cnt
	}
	//1 求最大值
	max :=0
	for _,pile := range piles{
		if pile>max{
			max = pile
		}
	}
	//2 二分  在速度范围内二分【1，max]
	left,right := 1,max
	result := -1
	for left <= right {
		mid := (left+right) /2
		hours := eatingTime(mid)
		if hours > h {// 二分查找最合适的速度
			left = mid+1
		}else{
			result = mid // <=h 肯定是符合的，但是不一定是最合理的数据，可以先存下来
            right = mid-1 // 继续从 mid-1之前找最小值
		}

	}
	return result
}
