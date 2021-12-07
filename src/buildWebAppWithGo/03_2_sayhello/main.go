package main
/*
Go 搭建一个Web服务器  3.2
*/


import (
	"fmt"
	"log"
	"net/http"
	"strings"
)


func sayHelloName(w http.ResponseWriter,  r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Printf("path:%s\n", r.URL.Path)
	fmt.Printf("scheme:%s\n", r.URL.Scheme)
	fmt.Printf("r.Form[\"url_log\"]:%s",r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello web world!") //这个写入到w的是输出到客户端的
}
/*
func http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
handler   处理请求和生成返回信息的处理逻辑
 
*/
func main(){
	/*
	为模式"/"注册处理器函数
	*/
	http.HandleFunc("/", sayHelloName)   //设置访问的路由

	/*
	第二个参数是处理器参数，处理请求    一般为nil
	The handler is typically nil, in which case the DefaultServeMux is used.

	默认获取handler = DefaultServeMux   路由器 匹配url跳转到相应的handle函数
		即http.HandleFUnc(  , handle: syyHelloName)
		当请求为"/"时，路由会跳转到syaHelloName
	*/
	err := http.ListenAndServe(":9090", nil)  //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServer:",err)
	}
}