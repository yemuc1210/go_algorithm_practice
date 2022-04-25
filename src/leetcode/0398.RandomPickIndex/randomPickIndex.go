package main

import "math/rand"

// 给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。
// 您可以假设给定的数字一定存在于数组中。

// 注意：
// 数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。

// 数字相同，则每个索引返回的概率应该相同
// 排除map
// 蓄水池采样，水塘抽样：从包含n个项目的集合中抽取k个样本。
// 适用于不能吧所有n个项目都存放到内存的情况

type Solution map[int][]int

func Constructor(nums []int) Solution {
    pos := map[int][]int{}
    for i, v := range nums {
        pos[v] = append(pos[v], i)
    }
    return pos
}

func (pos Solution) Pick(target int) int {
    indices := pos[target]
    return indices[rand.Intn(len(indices))]
}




/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

 func main() {

 }