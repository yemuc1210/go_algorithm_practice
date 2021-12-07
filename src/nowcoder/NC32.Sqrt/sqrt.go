package NC32



/**
 * 求平方根 lt69  牛顿迭代法
 * @param x int整型
 * @return int整型
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
func sqrt( x int ) int {
    // write code here
	if x==0{
		return 0
	}
	C,x0 := float64(x),float64(x)  //x0=C作为初始点
	for { //迭代
		xi := 0.5*(x0 + C/x0)  //下一个迭代点
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