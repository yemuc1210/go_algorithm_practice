package main

import (
    "fmt"
)
// 求解立方根
func main() {
    var N float64
    fmt.Scanln(&N)
    
    fmt.Printf("%0.1f\n",solve(N))
}


func solve(y float64) float64 {
    //f(x) = x^3-y   求f(x)=0  x^3=y
    var x float64
    //牛顿迭代法  xn+1 = xn - f'(xn)  x=(x^3-y)/(3x^2)=(2x+y/x/x)/3
    // 泰勒展开  f(x) = f(x0)+f'(x0)(x-x0) = x0^3-y + 3x0^2(x-x0)=0
    //   x = x0 - (x0^3-a)/(3x0^2)
    for x=1.0;abs(x*x*x-y)>1e-7;{
        x = x - ((x*x*x - y) / (3*x*x))
    }
    return x
}

func abs(x float64) float64 {
    if x<0 {
        return -x
    }
    return x
}