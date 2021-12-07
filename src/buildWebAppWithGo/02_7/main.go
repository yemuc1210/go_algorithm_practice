package main

import (
	"fmt"
	"runtime"
)

//goroutine

func runtimeFunc(){
	fmt.Println(runtime.NumCPU())   //out:16 
}

func main(){
	runtimeFunc()
}