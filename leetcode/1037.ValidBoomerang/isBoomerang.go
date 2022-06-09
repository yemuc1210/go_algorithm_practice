package main

// 如果所有点构成回旋镖则返回true
// 回旋镖：一组三个点，各不相同，且不在一条直线上
// 提示：points.length==3     points[i].length==2

func isBoomerang(points [][]int)bool{
	// 判断三个点是否相同
	// 计算面积是否是0
	// 面积公式 S=1/2 [(x1y2-x2y1)+(x2y3-x3y2)+(x3y1-x1y3)]
	// x1,x2,x3,y1,y2,y3 := points[0][0],points[1][0],points[2][0],points[0][1],points[1][1],points[2][1]
	// S := ((x1*y2-x2*y1)+(x2*y3-x3*y2)+(x3*y1-x1*y3))/2
	// return S==0

	// 斜率计算
	x1 := points[0][0] - points[1][0]
	y1 := points[0][1] - points[1][1]
	// 斜率 k1 = y1/x1

	x2 := points[0][0] - points[2][0]
	y2 := points[0][1] - points[2][1]
	// 斜率 k2 = y2/x2

	// 判断k1=!=k2  ->  k1=k2  x1*y2==x2*y1
	return x1*y2 != x2*y1
}