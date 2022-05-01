package main

import "fmt"

// m次投掷的结果，均值mean，遗失的次数n，求n次投掷的情况
func missingRolls(rolls []int, mean int, n int) []int {
	// 均值的计算特点 average(rolls) + average(leftN) = mean ?
	// [3 2 4 3] mean=4 n=2  leftN=[6,6]
	// average(rolls) = 12/4=3    average(leftN)=12/2=6
	// sum(rolls)/(m+n) = 12/6=2     sum(leftN)/(m+n)=12/6=2

	totalSum := mean * (len(rolls) + n)
	sumOfRolls := 0
	for _, v := range rolls {
		sumOfRolls += v
	}

	missingSum := totalSum - sumOfRolls
	// 遗失的n次总和missingSum
	// 问题变为：n个1-6的数，其总和为missingSum
	// 组合题
	// 首先判断是否有解
	if missingSum > 6*n || missingSum < n {
		return []int{}
	}
	// 有解，计算
	res := make([]int, n)
	quotient := missingSum / n
	remainder := missingSum % n
	// 缺失的n个数中，有remainder个 quotient+1，其余都是quotient
	for i := 0; i < n; i++ {
		res[i] = quotient
		if i < remainder {
			res[i]++
		}
	}
	return res
}

func main() {
	rolls := []int{3, 2, 4, 3}
	mean, n := 4, 2

	res := missingRolls(rolls, mean, n)

	fmt.Println(res)
}
