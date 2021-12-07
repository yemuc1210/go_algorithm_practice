package main
/*
HJ3 明明的随机数
较难  通过率：20.31%  时间限制：1秒  空间限制：32M
知识点	数组
warning 校招时部分企业笔试将禁止编程题跳出页面，为提前适应，练习时请使用在线自测，而非本地IDE。
描述
明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，
他先用计算机生成了 N 个 1 到 1000 之间的随机整数（ N≤1000 ），对
于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。
然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。
请你协助明明完成“去重”与“排序”的工作(同一个测试用例里可能会有多组数据(用于不同的调查)，
希望大家能正确处理)。

注：测试用例保证输入参数的正确性，答题者无需验证。测试用例不止一组。
当没有新的输入时，说明输入结束。
数据范围：  ，输入的数字大小满足
输入描述：
注意：输入可能有多组数据(用于不同的调查)。每组数据都包括多行，第一行先输入随机整数的个数 N ，接下来的 N 行再输入相应个数的整数。具体格式请看下面的"示例"。
输出描述：
返回多行，处理后的结果
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "sort"
)
func mysort_asc(mylist []int) ([]int) {
	// 升序排序
// 	for i:=0;i<len(mylist);i++ {
// 		for j:=i;j<len(mylist);j++ {
// 			if mylist[i] > mylist[j] {
// 				mylist[i],mylist[j] =mylist[j],mylist[i]
// 			}
// 		}
// 	}
    sort.Ints(mylist)
	//for i:= range mylist {
	//	fmt.Printf("%d ",mylist[i])
	//}
	return mylist
}

func deduplicate(mylist []int) ([]int){
    //去重
	var res []int
	for _,v := range mylist {
		// 标记是否存在
		var flag int = 0
		for _,x := range res {
			// 如果存在
			if v == x {
				flag = 1
				break
			}
		}
		if flag == 0 {
			// 不存在
			res = append(res, v)
		}
	}
	return res
}

func main()  {
	// 初始化一个对象
	scanner := bufio.NewScanner(os.Stdin)


  for {

  	// 从输入读取
  	scanner.Scan()
  	var str = scanner.Text() //fmt.Println("输入的是："+str)

  	// 如果读取到空行就退出
  	if str == "" {
		return
	}
	// 获取第一个数
	var n int
	n,_ = strconv.Atoi(str)

	// 定义切片，用于存放后面输入的具体数
	var nums []int
	//fmt.Println(n)
	for i:=0;i<n;i++{
		scanner.Scan()
		str = scanner.Text()
		x,_ := strconv.Atoi(str)
		nums = append(nums,x)
	}
	//fmt.Println(nums)

	// 升序排序
	nums = mysort_asc(nums)

	// 去冲
	nums = deduplicate(nums)

	// 输出
	for _,v := range nums{
		fmt.Println(v)
	}
  }
}