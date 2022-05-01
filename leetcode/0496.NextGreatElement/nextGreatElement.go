package lt496
/**
496. 下一个更大元素 I
给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。
如果不存在，对应位置输出 -1 。

我们可以使用单调栈来解决第 1 个子问题。倒序遍历nums2，
并用单调栈中维护当前位置右边的更大的元素列表，从栈底到栈顶的元素是单调递减的。

具体地，每次我们移动到数组中一个新的位置 i，就将当前单调栈中所有小于nums 2[i] 的元素弹出单调栈，
当前位置右边的第一个更大的元素即为栈顶元素，
如果栈为空则说明当前位置右边没有更大的元素。
随后我们将位置 i 的元素入栈。
*/
func nextGreaterElement(nums1, nums2 []int) []int {
    mp := map[int]int{}
    stack := []int{}
    for i := len(nums2) - 1; i >= 0; i-- {
        num := nums2[i]
        for len(stack) > 0 && num >= stack[len(stack)-1] { //大于栈顶元素
            stack = stack[:len(stack)-1]  //弹出
        }
		//此时栈顶小于num
        if len(stack) > 0 {
            mp[num] = stack[len(stack)-1]
        } else {
            mp[num] = -1
        }
        stack = append(stack, num)
    }
    res := make([]int, len(nums1))
    for i, num := range nums1 {
        res[i] = mp[num]
    }
    return res
}

