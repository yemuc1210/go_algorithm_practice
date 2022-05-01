package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param a int整型一维数组
 * @return int整型
 */
func getNumber(a []int) int {
	// write code here
	// 将下标为非质数的元素删除
	// 删除完之后,重新拼接
	// 不断循环,知道数组a大小为1
	// 求最后剩余的元素
	n := len(a)+2
	// 小于等于n的
	nonPrimes := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if nonPrimes[i] {
			continue
		}
		for j := i * i; j < n; j = j + i {
			// j遍历i的倍数
			nonPrimes[j] = true // 是质数的倍数关系的,非质数
		}
	}
	nums := []int{}
	// nonPrime 中 不是质数的为true
	nonPrimes[1] = true
	for len(a) != 1 {
		for i := 0; i < len(a); i++ {
			if nonPrimes[i+1] { //
				continue
			} else {
				nums = append(nums, a[i])
			}
		}
		// 令a 指向
		a = make([]int, len(nums))
		copy(a, nums)
		nums = []int{}
	}
	return a[0]
}

// 2<=n<=2*10^5  长度
// 1<=ai<=10^5

func main() {
	// 求所有质数
	a := []int{1, 2, 3, 4}
	res := getNumber(a)


	fmt.Println(res)
}
