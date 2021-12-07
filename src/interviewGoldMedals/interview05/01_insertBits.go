package interview05

// import "fmt"

/*
面试题 05.01. 插入
给定两个整型数字 N 与 M，以及表示比特位置的 i 与 j（i <= j，且从 0 位开始计算）。

编写一种方法，使 M 对应的二进制数字插入 N 对应的二进制数字的第 i ~ j 位区域，
不足之处用 0 补齐。具体插入过程如图所示。

位运算 与运算
*/
func insertBits(N int, M int, i int, j int) int {
	//将M变为j-i+1位二进制数，不足位置补为0   左移/右移动
	//判断二进制M的位数
	// M <<= i
	// for k := i; k <= j; k++ {
	// 	if N&(1<<k) != 0 {
	// 		N^=1<<k
	// 	}
	// }
	// return N | M
	//left in (,j+1];middle in [i,j];right in [i-1,0];
    left := N>>j>>1; //把最左边的部分调整好了，即抛弃了替换部分和低位部分
    left = left<<j<<1;  //因此最后要进行或运算，所以把他再移到原来的高位上。
	middle := M<<i;  //替换N的j<-----i位，那么只需要将M左移i位即可
    right := N&((1<<i)-1);//只需要N的低位，将高位置零,(1<<2)-1 = (11)2
    return left | middle | right;

}