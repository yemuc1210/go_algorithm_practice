package lt384

import "math/rand"

/*
算法计划21 最后一天  中
给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。

实现 Solution class:

	Solution(int[] nums) 使用整数数组 nums 初始化对象
	int[] reset() 重设数组到它的初始状态并返回
	int[] shuffle() 返回数组随机打乱后的结果


洗牌算法
*/
type Solution struct {
    arr []int
}


func Constructor(nums []int) Solution {
    
    return Solution{nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.arr
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    n := len(this.arr)
    res := make([]int, n)
    copy(res, this.arr)
    for i := n-1; i >= 0; i-- {
        rand := rand.Intn(i+1)    // math.rand中的Intn(i+1)返回[0, i]范围的整数，每次数组在下标index为[0, i]范围内随机找一个下标对应的元素与当前位置i处的元素进行交换
        res[i], res[rand] = res[rand], res[i]   // 对应位置元素交换，也可以使用如下代码
        // tmp := res[i]
        // res[i] = res[rand]
        // res[rand] = tmp
    } 
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
