package lt346

/*
剑指 Offer II 041. 滑动窗口的平均值
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，
计算滑动窗口里所有数字的平均值。

实现 MovingAverage 类：
	MovingAverage(int size) 用窗口大小 size 初始化对象。
	double next(int val) 成员函数 next 每次调用的时候都会往滑动窗口增加一个整数，
	请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。
*/

type MovingAverage struct {
	size int   //滑动窗口尺寸
	curSum int 
	nums []int  // 数据量
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		nums: []int{},
		curSum: 0,
	}
}


func (this *MovingAverage) Next(val int) float64 {
	this.curSum += val  //数据流+一个数
	this.nums = append(this.nums, val)
	for len(this.nums) > this.size{
		//移除多余
		remove := this.nums[0]
		this.curSum -= remove
		this.nums = this.nums[1:]
		// this.nums[0] = nil //   看来不行 无法及时让GC回收
	}
	return float64(this.curSum)/float64(len(this.nums))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */