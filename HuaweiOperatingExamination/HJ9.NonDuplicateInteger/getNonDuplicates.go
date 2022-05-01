package main

import "fmt"

/*
HJ9 提取不重复的整数
知识点 数组 位运算 哈希
描述
输入一个 int 型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
保证输入的整数最后一位不是 0 。

数据范围：
输入描述：
输入一个int型整数

输出描述：
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数
*/
//逆序数字 去重重复元素
func main(){
	var input int64
	fmt.Scanln(&input)
	//分解每位数字 保存到数组
	nums := []int64{}
	var tmp int64 = 0
	for input != 0 {
		nums = append(nums, input%10)
		tmp = tmp*10 + input%10
		input /= 10
	}
	//nums 保存的逆序 数位
	//去重
	m := make(map[int64]int)
	for i:=0;i<len(nums);i++{
		m[nums[i]] ++
	}
	//重组
	var res int64 = 0
	for i:=0;i<len(nums);i++{
		if cnt,_ := m[nums[i]] ; cnt!=0{
			//第一次使用
			res = res*10 + nums[i]
			m[nums[i]] = 0
		}else{
			continue
		}
	}
	fmt.Println(res)
}