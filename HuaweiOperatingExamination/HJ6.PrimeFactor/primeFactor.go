package main

import "fmt"
/*
HJ6 质数因子
中等  通过率：26.24%  时间限制：2秒  空间限制：32M
知识点   排序
warning 校招时部分企业笔试将禁止编程题跳出页面，为提前适应，练习时请使用在线自测，而非本地IDE。
描述
功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）


数据范围： 
输入描述：
输入一个整数

输出描述：
按照从小到大的顺序输出它的所有质数的因子，以空格隔开。最后一个数后面也要有空格。
*/
func main() {
	var input int64
	fmt.Scan(&input)
	var i int64
	for i = 2; i * i  <= input; i++ {
		for input%i == 0 {  //
			fmt.Printf("%d ", i)
			input /= i
		}
	}
    if input >= 2{
       fmt.Println(input)
    }
}

