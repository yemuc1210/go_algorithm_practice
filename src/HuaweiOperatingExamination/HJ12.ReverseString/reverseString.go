package main
import "fmt"
func main(){
	var s string
	fmt.Scanln(&s)
	//只包含小写字母的字符串
	bytes := []byte(s)
	n := len(bytes)
	for i:=0;i<n/2;i++{
		bytes[i],bytes[n-i-1] = bytes[n-i-1],bytes[i]
	}
	//输出翻转后的
	fmt.Println(string(bytes))
}