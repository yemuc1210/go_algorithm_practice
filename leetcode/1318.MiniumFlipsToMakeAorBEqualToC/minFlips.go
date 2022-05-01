package lt1318

/*
1318. 或运算的最小翻转次数
给你三个正整数 a、b 和 c。

你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c  成立的最小翻转次数。

「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1 。

x & (1 << k) 来判断 x 二进制表示的第 k 位（最低位为第 0 位）是否为 1
(x >> k) & 1 得到 x 二进制表示的第 k 位
x & 1 得到 x 的二进制表示的最低位，

设 a、b 和 c 二进制表示的第 i 位分别为 bit_a、bit_b 和 bit_c，根据 bit_c 的值，会有以下两种情况：

	若 bit_c 的值为 0，那么 bit_a 和 bit_b 必须都为 0，需要的翻转次数为 bit_a + bit_b；

	若 bit_c 的值为 1，那么 bit_a 和 bit_b 中至少有一个为 1，只有当它们都为 0 时，才需要 1 次翻转；

*/
func minFlips(a int, b int, c int) int {
	res :=0

	bit_a,bit_b,bit_c := 0,0,0

	for i:=0;i<31;i++{
		bit_a = (a>>i) & 1   //得到a二进制位第i位
		bit_b = (b>>i)&1
		bit_c = (c>>i)&1

		if bit_c == 0{
			//则需要bita b都为0
			res += bit_a + bit_b  //若bit_x不为0，则其为1
		}else{
			//bit_c=1，则需要其中一个为1
			if bit_a+bit_b==0{
				res += 1
			}else{
				res +=  0
			}
		}
	}

	return res
}