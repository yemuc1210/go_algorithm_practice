package offer56
/*
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
请写程序找出这两个只出现一次的数字。
要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：
输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]

时间复杂度O(n)，则不能使用排序
空间复杂度O(1)，不能使用map


一个数和自身异或，等于自身

分组异或
*/
func singleNumbers(nums []int) []int {
    // 此题是找一个孤立值的变种，关键在于拆分问题，如何分组
    var a int
    for _,v := range nums{
        a ^= v
    }
    mask := a&(-a)

    ans := make([]int,2)
    for _,v := range nums{
        if v & mask == 0{// 此处进行分组
            ans[0] ^= v // 不用给每个组分配多余的空间，只要把这个组中所有的数异或起来，就能得到那个孤立的数
        }else{
            ans[1] ^= v
        }
    }
    return ans
}