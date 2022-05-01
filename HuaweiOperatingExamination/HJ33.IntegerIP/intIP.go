package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
	var ipInput string
	var intInput int32
	for {
		n,_ := fmt.Scanln(&ipInput)
		if n==0 {
			break
		}
		fmt.Scanln(&intInput)
		// 输出 ip-> integer
		fmt.Println(ip2int(ipInput))
		// 输出 integer->ip
		fmt.Println(int2ip(intInput))
	}
}

func ip2int(ip string) int32 {
	// ip 地址32位 所以整数最大也就int32了
	// 变换方法  ip每段变为二进制数 然后拼接
	var b []int
	// 1 分割
	ips := strings.Split(ip,".")
	// 2 拼接
	for i:=0;i<len(ips);i++{
		n,_ := strconv.Atoi(ips[i])
		b = append(b, int2binary(int32(n))...)
	}
	// 3 转换成整数
	factor := 1
	n := len(b)
	var res int32
	for i:=n-1;i>=0;i-- {
		res += int32(b[i]*factor)
		factor *= 2
	}
	return res
}
// 整数转二进制     不足位数要补全 保证32位
func int2binary(n int32) []int {
	var res []int
	for n!=0 {
		a := int(n&1)  //得到最后一位的0 1情况
		n = n>>1  //右移
		res = append(res, a)
	}
	// 逆置一下
	l := len(res)
	for i:=0;i<l/2;i++ {
		res[i],res[l-i-1] = res[l-i-1],res[i]
	}
	
	if l<8 {   // 对应用例：输入ip 计算每一小段的二进制表示
		//需要在前面补零 假设数字输入都是合法的情况 只会有位数不足的情况
		left := 8-l   // 差几位0
		for i:=0;i<left;i++{
			res = append([]int{0},res...)
		}
	}
	if l>8 && l< 32 {
		// 对应 输入一个整数转换成ip的用例
		left := 32-l
		for i:=0;i<left;i++{
			res = append([]int{0},res...)
		}
	}

	// fmt.Printf("int2binary: res=%v,len(res)=%v\n",res,len(res))   //
	return res
}

func int2ip(intInput int32) string {
	// 1 将整数转为二进制
	b := int2binary(intInput)   // b 需要确保是32位的
	// 2 分割二进制  8位一组
	ints := []int32{}
	for i:=31;i>=0;i-=8{
		// 8个一组
		factor := 1
		var tmp int32
		for j:=0;j<8;j++ {
			tmp = tmp + int32(factor*b[i-j])
			factor *= 2
		}
		ints = append([]int32{tmp},ints...) // 头插
	}
	// 3 得到字符串ip结果
	var res string
	for i:=0;i<len(ints);i++ {
		if i==len(ints)-1 {
			res += strconv.Itoa(int(ints[i])) 
		}else{
			res += strconv.Itoa(int(ints[i])) + "."
		}
	}
	return res
}