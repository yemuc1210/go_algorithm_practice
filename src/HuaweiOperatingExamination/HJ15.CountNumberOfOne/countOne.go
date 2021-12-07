package main
import "fmt"

//求int型中1的个数  二进制中1的个数  lt191
func main(){
	var N int32
	fmt.Scanln(&N)
	//计算
	cnt := 0
	for ;N>0; N&=N-1 {
		cnt++
	}

	fmt.Println(cnt)
}