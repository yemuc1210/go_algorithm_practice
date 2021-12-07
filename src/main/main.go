package main
import "hello"
func main(){
	hello.Foo()  //之前GO111MODULE=on，所以编译器只会去goroot下搜索    gopath和gomod是不同的包管理方案
	/*
	gomod下查找包，解析go.mod文件查找包，mod 包名就是包的前缀，里面的目录是后续路径
	gomod不会自动到gopath下查找
	*/

}
