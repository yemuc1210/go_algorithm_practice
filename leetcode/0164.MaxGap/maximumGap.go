package lt164
/*
164. 最大间距
给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
如果数组元素个数小于 2，则返回 0。
*/
/*
有序元素的相邻元素的最大差值
思路1：排序
思路2：hash表
	假设一开始res = MaxInt32   
	从i=0开始  判断 nums[i] + res 在不在哈希表里 若在 则res ok  否则 缩小res  不咋行。。。
思路3：基数排序  桶排序
	按照个位数的值 分配到编号为0-9的桶
	将这数值串，然后根据十位数分配
	。。。知道根据最高位数排
	分为两种 LSD   MSD   即一种从低位开始  一种从高位作为基底
*/
func maximumGap(nums []int) int {
	if len(nums)<2 {
		return 0
	}
	m := nums[0]  //令m是数组最大值
	for i:=1;i<len(nums);i++{
		m  = max(m, nums[i])
	}
	exp := 1 // 1 10 100 1000
	R := 10  //
	aux := make([]int, len(nums))
	for (m/exp) > 0 {   //lsd
		cnt := make([]int, R)  //0-9 桶
		for i:=0;i<len(nums);i++{
			cnt[(nums[i]/exp)%10] ++   //每个桶计数
		}
		for i:=1;i<len(cnt);i++{
			cnt[i] += cnt[i-1]  //总计数
		}
		for i:=len(nums)-1;i>=0;i--{
			tmp := cnt[(nums[i]/exp)%10]
			tmp -- 
			aux[tmp] = nums[i]
			cnt[(nums[i]/exp)%10] = tmp
		}
		for i:=0;i<len(nums);i++{
			nums[i] = aux[i]
		}
		exp *=10
	}
	maxVal := 0
	for i:=1;i<len(aux);i++{
		maxVal = max(maxVal,  aux[i]-aux[i-1])
	}
	return maxVal
}
func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}

//快排
func maximumGap1(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	quickSort164(nums, 0, len(nums)-1)
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i+1] - nums[i]) > res {
			res = nums[i+1] - nums[i]
		}
	}
	return res
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort164(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition164(a, lo, hi)
	quickSort164(a, lo, p-1)
	quickSort164(a, p+1, hi)
}