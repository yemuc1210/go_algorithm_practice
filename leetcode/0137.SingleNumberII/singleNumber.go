package lt137
//lt 136   lt260
/*中
137. 只出现一次的数字 II
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。
请你找出并返回那个只出现了一次的元素。
进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？  
*/
/*
其余元素出现三次 则所有元素都是出现奇数次 异或行不通
出现三次 其余元素的和是奇数  如果最后所有元素的和是偶数 则结果是个单数  。。。没啥用

0001  0001   0001    0010 

设计电路 一个出现3次消为0
记录出现三次，需要两个二进制位
	ab^00 = ab   ab^ab=00   x&-x = 0     x& -0 = x
使用a的第k位与b的第k位组合成两位二进制，表示当前位出现几次
x = x[7] x[6] x[5] x[4] x[3] x[2] x[1] x[0]
x = (a[7]b[7])......(a[0]b[0])

Grandyang有举了个例子，希望能帮助大家理解。
	ab xor x
	00 (+) 1 = 01
	01 (+) 1 = 10
	10 (+) 1 = 00 ( mod 3)
转换为算法：
	b = b xor x & ~a;
	a = a xor x & ~b;
	说明：当a[i]b[i]为00时，遇到x[i]为1，这时候应该把a[i]更新为0，b[i]更新为1....etc

答案的第 i 个二进制位就是数组中所有元素的第 i 个二进制位之和除以 3 的余数。
这样一来，对于数组中的每一个元素 x，我们使用位运算 (x >> i) & 1 得到 x 的第 i 个二进制位，
并将它们相加再对 3 取余，得到的结果一定为 0 或 1，即为答案的第 i 个二进制位

*/
func singleNumber(nums []int) int {
	res := int32(0)
	for i:=0;i<32;i++{
		total := int32(0)
		for _,num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			res = res | 1<<i
		}
	}
	return int(res)
}