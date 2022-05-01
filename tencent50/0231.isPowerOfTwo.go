package tencent50
/*简单
231. 2 的幂
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。
*/
/*
n位数是2的幂，则二进制位数只能有一个1
取最低位的1 再判断剩下是否为0

n &(n-1)移除最低位1
n &(-n) 获取最低位的1
如补码： 3： 011   2： 010
 		-3 ： 101  -2  1 10
*/
func isPowerOfTwo(n int) bool {
	return n>0 && (n & -n)==n
}

// bool isPowerOfTwo(int n){
//     if(n==1) return true;
//     if(n==0) return false;
// 	if(n%2!=0) return false;
// 	return isPowerOfTwo(n/2);
// }