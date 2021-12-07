package main

import (
	// "bufio"
	"fmt"
	// "os"
)
func main(){
	// reader := bufio.NewReader(os.Stdin)
	var N,m int32
	fmt.Scanln(&N,&m)  //N总钱数  m希望购买物品个数
	//下面输入物品基本数据  v价格 p重要程度  q0主 >0则是所属逐渐的编号
	//构建结构体？
	type elem struct{
		v,p int
		subElem []int 
	}
	elems := make([]elem, m+1)
	var v,p,q int
	var i int32
	for i=1;i<=m;i++{
		fmt.Scanln(&v,&p,&q)
		if q == 0 {
			//主
			elems[i] = elem{v:v,p:p}
		}else{
			
		}
	}
}