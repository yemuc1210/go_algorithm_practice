package offer40

import "sort"

/*
剑指 Offer 40. 最小的k个数
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

快排的思想
划分
*/
func getLeastNumbers(arr []int, k int) []int {
    sort.Ints(arr)
    res:=[]int{}
    for i:=0;i<k;i++{
        res=append(res,arr[i])
    }
    return res
}