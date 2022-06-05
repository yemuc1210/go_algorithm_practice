package main

import "math/rand"

// 在⚪内随机生成点
// 给定⚪的半径和圆心位置，在⚪内产生均匀随机点
type Solution struct {
	// 半径
	radius float64
	// 圆心位置
	x,y float64
}


func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius,x: x_center,y: y_center}
}

// 随机生成点
func (this *Solution) RandPoint() []float64 {
	for {
		x1 := rand.Float64()*2-1
		y1 := rand.Float64()*2-1   //[-1,1)
		if x1*x1+y1*y1 <1 {
			return []float64{this.x+x1*this.radius, this.y+y1*this.radius}
		}
	}                                                         
}