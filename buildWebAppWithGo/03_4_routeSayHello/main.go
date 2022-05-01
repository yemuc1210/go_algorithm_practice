package main

import (
	"fmt"
	"net/http"
)

/*实现一个简易的路由器
即实现一个serveHTTP函数处理对应的pattern
*/

type SelfMutex struct{

} //serveMutex

func (mx *SelfMutex) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.URL.Path =="/"{  //获取r 中的URI field 然后获取其中的Path值
		//调用某函数处理
		sayHelloName(w,r)  
		return 
	}
	http.NotFound(w,r)   //否则没有匹配的handle函数
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello self route mutex")
}


func main(){
	mux := &SelfMutex{}  //
	//func http.ListenAndServe(addr string, handler http.Handler) error
	// mux.handler(r).ServeHTTP(w, r)    handler函数按照规则选择对应的ServeHTTP
	http.ListenAndServe(":9090", mux)
}