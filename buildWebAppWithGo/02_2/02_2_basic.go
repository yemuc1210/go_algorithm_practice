package main

import "fmt"

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func define_group(){
	var (
		i,j int
		pi float32
		prefix string
	)

	i,j = 0,0
	pi = 3.14
	prefix = "02-2-basic"

	fmt.Println(pi,i,j,prefix)
}

func itoa_test(){
	//itoa关键字 在声明enum是采用
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}

func goto_test(){
	i:=0
	Here:
		println(i)
		i++
		if i>10 {
			goto Out
		}
		goto Here
	Out:
		return
}
func main(){
	//
	define_group()
	itoa_test()
	goto_test()
}