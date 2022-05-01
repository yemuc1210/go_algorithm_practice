package main

import(
	"fmt"
)

func main(){
	var input int64
	fmt.Scanln(&input)
	// fmt.Println(input)
	//整数以字符串形式逆序输出
	res := []byte{}
	if input == 0 {
		fmt.Print(input)
	}
	for input != 0 {
		// fmt.Println(byte(input%10))
		res = append(res, byte(input%10))
		input /= 10
	}
	for i:=0;i<len(res);i++{
		fmt.Print(res[i])
	}
}