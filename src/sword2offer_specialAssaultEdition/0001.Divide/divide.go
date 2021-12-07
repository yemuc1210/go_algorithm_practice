package offerS1

import "math"
/*
注意：

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1

*/
func divide(a int, b int) int {
	//除法  不使用*  、   %
	//使用减法  a/b就是求a中有多少个b
	//健壮性
	if a==math.MinInt32 && b==-1{
		return math.MaxInt32    //越界
	}

	//正负号
	sign  := 1
	if (a>0 && b<0) || (a<0 && b>0){
		sign = -1
	}

	if a>0{
		a = -a
	}
	if b>0{
		b=-b
	}

	res := 0
	// for a <= b{   //此时a,b小于 0
	// 	a -= b
	// 	res ++
	// }   //超时
	for a <= b {
        value, k := b, 1
        // 0xc0000000 是十进制 -2^30 的十六进制的表示
        // 判断 value >= 0xc0000000 的原因：保证 value + value 不会溢出
        // 可以这样判断的原因是：0xc0000000 是最小值 -2^31 的一半，
        // 而 a 的值不可能比 -2^31 还要小，所以 value 不可能比 0xc0000000 小
        // -2^31 / 2 = -2^30
        for value >= 0xc0000000 && a <= value + value {
            value += value
            k += k
        }
        a -= value
        res += k
    }

	return res*sign
}

func divide1(a int, b int) int {
    if a == math.MinInt32 && b == -1 {
        return math.MaxInt32
    }

    sign := 1
    if (a > 0 && b < 0) || (a < 0 && b > 0) {
        sign = -1
    }

    a = abs(a)
    b = abs(b)

    res := 0
    for i := 31; i >= 0; i-- {
        if (a >> i) - b >= 0 {
            a = a - (b << i)
            res += 1 << i
        }
    }
    return sign * res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
