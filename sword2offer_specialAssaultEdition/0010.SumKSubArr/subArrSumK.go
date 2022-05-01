package offerS10

/*
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

我们可以枚举 [0..i]里所有的下标 j来判断是否符合条件，
可能有读者会认为假定我们确定了子数组的开头和结尾，还需要 O(n)的时间复杂度遍历子数组来求和，
那样复杂度就将达到 O(n^3) 从而无法通过所有测试用例。
但是如果我们知道 [j,i]子数组的和，就能 O(1) 推出 [j-1,i]的和，
因此这部分的遍历求和是不需要的，我们在枚举下标 j 的时候已经能 O(1) 
求出 [j,i] 的子数组之和。
*/

func subarraySum(k int, nums []int) int{
    count := 0
    for start := 0; start < len(nums); start++ {
        sum := 0
        for end := start; end >= 0; end-- {
            sum += nums[end]
            if sum == k {
                count++
            }
        }
    }
    return count
}