package lt400

import (
	"fmt"
	"math"
)

/*
400. 第 N 位数字
给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...]
中找出并返回第 n 位数字。
*/
/*
找规律
首先计算前缀部分，再计算尾部。
	小于10，直接返回。(step 1)
	否则，计算前缀部分，全部被占用部分总共有多少位即length。(step 2)
	计算尾部，其一、计算第一个数字即pow(10, i-1)。
		其二、推理出最后一个出现的多位数字即num，并计算出位于第几位即index。(step 3)
//1234567891011121314151617181920
//1 - [1,9]             9
//2 - [10,99]          90
//3 - [100,999]       900
//4 - [1000,9999]    9000
//.........2^31-1 = 2147483647
*/
func findNthDigit(n int) int {
	if n<10 {
		return n
	}
	//计算
	length,cnt,i := 0,9,1
	for length+cnt*i < n {
		length += cnt *i 
		cnt *=10
		i++
	}
	//得到最后一个多位数
	num := int(math.Pow10(i-1)) + (n-length-1)/i
	idx := (n-length-1)%i
	//返回num[idx]
	snum := fmt.Sprintf("%d",num)
	return int(snum[idx]-'0')
}