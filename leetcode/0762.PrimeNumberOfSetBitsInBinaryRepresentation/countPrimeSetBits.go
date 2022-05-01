package main

import "math/bits"

// 762. 二进制表示中质数个计算置位
// 给你两个整数 left 和 right ，在闭区间 [left, right] 范围内，统计并返回 计算置位位数为质数 的整数个数。

// 计算置位位数 就是二进制表示中 1 的个数。

// 例如， 21 的二进制表示 10101 有 3 个计算置位。

func isPrime(x int) bool {
    if x < 2 {
        return false
    }
    for i := 2; i*i <= x; i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}
// 665772的二进制表示是0000 0000 0000 1010 0010 1000 1010 1100，所有的1都在质数位置。
func countPrimeSetBits(left, right int) (ans int) {
    for x := left; x <= right; x++ {
        if isPrime(bits.OnesCount(uint(x))) {
            ans++
        }
    }
    return
}
// 665772的二进制表示是0000 0000 0000 1010 0010 1000 1010 1100，所有的1都在质数位置。

func solve(left,right int) int {
	res := 0
	for i:=left;i<=right;i++{
		res += 665772 >>  bits.OnesCount(uint(i)) & 1
	}
	return res
}