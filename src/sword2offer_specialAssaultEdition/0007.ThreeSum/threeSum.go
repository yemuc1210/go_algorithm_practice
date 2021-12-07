package offerS7

import "sort"

/*
剑指 Offer II 007. 数组中和为 0 的三个数
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，
使得 a + b + c = 0 ？请找出所有和为 0 且 不重复 的三元组。

 leetcode 15

（1）暴力解法  三层循环
（2）排序  三元组 第一个数变为target  然后求两数之和
	可能存在多对，需要去重  去重一般使用map
	先排序 对去重也很有帮助

*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)  //利用排序 帮助去重
	ans := make([][]int, 0)
 
    // 枚举 a
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // c 对应的指针初始指向数组的最右端    twoSum思路
        third := n - 1
        target := -1 * nums[first]    //第一个数变为target 剩下两个数twoSum
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                ans = append(ans, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return ans


}