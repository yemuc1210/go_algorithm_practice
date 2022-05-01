package leetcode


/*
实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
*/

//牛顿迭代法
/*
函数表示为 f(x)=x^2 - C，问题转为成函数的零点
牛顿迭代本质是借助泰勒级数
首先，任取一点x0作为初始值
每一次迭代，找到函数图像上的点(xi,f(xi))，过该店作斜率为该点导数f'(xi)的直线，
	得到与横轴的交点xi+1。xi+1离xi距离零点更近，
多次迭代重复


选择x0=C作为初始点
每一次迭代，当前交点为xi，找到函数上的（xi,xi^2-C）,斜率2xi,
	直线方程为 yl=2xi(x-xi)+xi^2-C = 2xix-(xi^2+C)
	横轴交点为 2x_i*x-(xi^2+C)=0的解，得到新的x_i+1
		x_i+1 = 1/2 (xi+C/xi)

*/
func mySqrt(x int) int {
	if x==0{
		return 0
	}
	C,x0 := float64(x),float64(x)

	for {
		xi := 0.5*(x0 +C/x0)
		if myAbs(x0,xi) < 1e-7{  //会有个问题  负数也进入了 
			break
		}
		x0 = xi
	}
	return int(x0)
}
func myAbs(a,b float64) float64 {
	if a-b >= 0 {
		return a-b
	}
	return b-a
}
func mySqrt_1(x int) int {

	y,check := 1.0,0.0
	fx := float64(x)
	y = (fx/y + y)/2.0
	check = y*y - fx
	for check > 0.001 {
		y = (fx/y + y)/2.0
		check = y*y - fx
	}
	return int(y)
}