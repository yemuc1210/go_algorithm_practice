package NC120

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 输入整数n 计算32位二进制中1的个数   offer15
 * @param n int整型 
 * @return int整型
*/

/*
首先判断n是不是负数，当n为负数的时候，直接用后面的while循环会导致死循环，因为负数
向左移位的话最高位补1 ！ 因此需要一点点特殊操作，可以将最高位的符号位1变成0，也就
是n & 0x7FFFFFFF，这样就把负数转化成正数了，唯一差别就是最高位由1变成0，因为少了
一个1，所以count加1。之后再按照while循环里处理正数的方法来操作就可以啦！
*/
func NumberOf1( n int ) int {
    // write code here
	cnt := 0
	if n<0 {
		n = n & 0x7FFFFFFF
		cnt ++
	}  //最高位变为0
	

	for n!=0 {
		cnt += n&1  //取最后一位的1
		n = n>>1  //右移一位
	}
	return cnt
}

// func hammingWeight(num uint32) int {
//     cnt := 0
// 	one := uint32(1)
// 	for num != 0{
// 		if num & one != 0{   //最后一位是1   
// 			cnt ++
// 		}
// 		num = num >> 1
// 	}
	
// 	return cnt
// }