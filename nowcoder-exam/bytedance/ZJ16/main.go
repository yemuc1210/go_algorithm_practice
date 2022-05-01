package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	var n, m int
	var num int
	var err error
	num, err = fmt.Scanln(&n, &m)
	for  num != 0 && err != errors.New(io.EOF.Error()) {
		sum := 0.0
		var tmp float64 = float64(n)
		//求tmp平方根
		for i := 0; i < m; i++ {
			sum += tmp
			tmp = mySqrt(tmp)
			if tmp < 0.0001 {
				goto next //精度要求哦i
			}
		}
		fmt.Printf("%.2f\n", sum)
	next:
		num, err = fmt.Scanln(&n, &m)
	}
}

//牛顿迭代法  f(x)=x*x - C   求零点
func mySqrt(x float64) float64 {

	y, check := 1.0, 0.0
	fx := float64(x)
	y = (fx/y + y) / 2.0
	check = y*y - fx
	for check > 0.001 {
		y = (fx/y + y) / 2.0
		check = y*y - fx
	}
	return y
}
