package offer64

/*   中等题
剑指 Offer 64. 求1+2+…+n
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
*/


/*
乘除使用移位完成，左移×2  右移÷2

公式（1+n）*n/2
*/

// //全局变量
// var res int = 0
// func sumNums(n int) int {
// 	//递归   +终止判断
// 	x := n>1 && sumNums(n-1)>0       //if A&&B   如果A为false 则终止了  B不会执行   if A {B}  
// 	res += n
// 	return res
// }
// class Solution {
// 	public:
// 		int sumNums(int n) {
// 			bool arr[n][n+1];
// 			return sizeof(arr)>>1;
// 		}
// 	};
func sumNums(n int) int {
    ans, A, B := 0, n, n + 1
    addGreatZero := func() bool {
        ans += A
        return ans > 0
    }

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1

    _ = ((B & 1) > 0) && addGreatZero()
    A <<= 1
    B >>= 1
    return ans >> 1
}
