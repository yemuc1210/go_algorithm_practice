package NC11

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 最小的K个数  offer40
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution( input []int ,  k int ) []int {
    // write code here
	//不去重的最小k个数 思路1：排序  思路2；基于快排的
	sort.Ints(input)
    res:=[]int{}
    for i:=0;i<k;i++{
        res=append(res,input[i])
    }
    return res
}