package offer45

import (
	"fmt"
	"sort"
	"strings"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，
打印能拼接出的所有数字中最小的一个。

示例 1:
输入: [10,2]
输出: "102"

示例 2:
输入: [3,30,34,5,9]
输出: "3033459"

若拼接字符串x+y > y+x   则"x">"y"
反之，“x”<“y”   意味着排序完成后，x在y左边
3  30
330 >303   所以"30"<"3"  30排在3前面
*/
func minNumber(nums []int) string {
	var res strings.Builder
	//定义排序规则
	sort.Slice(nums, func(i, j int) bool {
		return compareNum(nums[i],nums[j])
	})
	for i:=0;i<len(nums);i++{
		res.WriteString(fmt.Sprintf("%d",nums[i]))
	}
	return res.String()
}
func compareNum(a,b int)bool{
	str1 := fmt.Sprintf("%d%d",a,b)
	str2 := fmt.Sprintf("%d%d",b,a)
	if str1<str2{
		return true
	}
	return false
}