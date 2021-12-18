package NC151
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 最大公约数
 * 求出a、b的最大公约数。
 * @param a int 
 * @param b int 
 * @return int
*/
func gcd( a int ,  b int ) int {
    //条件假设a,b>=1
	//令a为较大值
	if a<b {
		a,b = b,a
	}
	//辗转相除
	if a%b == 0 {
		return b
	}
	for b!=0 {
		a,b = b,a%b
	}
	return a
}