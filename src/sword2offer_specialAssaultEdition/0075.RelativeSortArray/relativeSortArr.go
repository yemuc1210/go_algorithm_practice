package offerS75
/*
剑指 Offer II 075. 数组相对排序   计数排序
给定两个数组，arr1 和 arr2，
	arr2 中的元素各不相同
	arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。
未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
提示：元素范围是【0，1000】

错误原因：
输入
[2,3,1,3,2,4,6,7,9,2,19]
[2,1,4,3,9,6]
输出
[2,2,2,1,4,3,3,9,6,2,19]
预期结果
[2,2,2,1,4,3,3,9,6,7,19]

需要考虑arr1原来的元素   如元素7

输入：
[28,6,22,8,44,17]
[22,28,8,6]
输出：
[22,28,8,6,44,17]
预期结果：
[22,28,8,6,17,44]
还需要考虑相对顺序
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
    upper := 0
	/*
	我们可以对空间复杂度进行一个小优化。
	实际上，我们不需要使用长度为 1001 的数组，
	而是可以找出数组arr1中的最大值upper，
	使用长度为1upper+1 的数组即可。
	*/
    for _, v := range arr1 {
        if v > upper {
            upper = v
        }
    }
    frequency := make([]int, upper+1)
    for _, v := range arr1 {  //计数
        frequency[v]++
    }

    ans := make([]int, 0, len(arr1))
    for _, v := range arr2 {  //首先取arr2中出现的元素
        for ; frequency[v] > 0; frequency[v]-- {
            ans = append(ans, v)
        }
    }
    for v, freq := range frequency {  //处理arr2中未出现的
        for ; freq > 0; freq-- {
            ans = append(ans, v)
        }
    }
    return ans
}

