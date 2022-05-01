package offer65
/*
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
*/
func add(a int, b int) int {
	for b!=0{
		c := (a&b) << 1    //进位c    &  进位加
		a ^= b   //a=非进位和    ^不进位加法
		b = c    //b=进位
	}
	return a
}