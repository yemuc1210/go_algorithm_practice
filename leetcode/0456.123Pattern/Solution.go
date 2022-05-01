package main

import (
	"fmt"
	"math"
)
func find132pattern(nums []int) bool {
	n:=len(nums)
	if n < 3 {
		return false
	}
	last := math.MinInt32  // 2
	sk := []int{}  // 存3
	for i:=n-1;i>=0;i--{
		if nums[i]<last{
			return true
		}
		for len(sk)>0 && sk[len(sk)-1]<nums[i] {
			last = sk[len(sk)-1]
			sk = sk[:len(sk)-1]
		}
		sk = append(sk, nums[i])
	}
	return false
}
func main(){
	nums := []int{3,1,4,2}
	res := find132pattern(nums)
	fmt.Println(res)
}