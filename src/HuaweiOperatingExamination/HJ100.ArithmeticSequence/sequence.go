package main

import (
    "fmt"
)

func main(){
    	//Sn = n*a1 +n(n-1)d/2   = n(a1+an)/2 an=a1+(n-1)d
	//a1=2  d=3
    var nums int 
    for {
        _, err := fmt.Scan(&nums)
        if err != nil {
            break
        }
        fmt.Println( 2*nums + nums*(nums -1)/2*3)
    }
}

