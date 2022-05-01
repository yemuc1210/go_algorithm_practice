package main

import "fmt"

func main(){
    var m int
    for {
		n,_ := fmt.Scanln(&m)
		if n==0 {
			break
		}
        var a,b,c =1,0,0 //a:一个月兔子数，b：两个月兔子数，c：三个月兔子个数
        for m-=1;m>0;m--{//每过一个月兔子数变化
            c+=b;
            b=a;
            a=c;
        }
        fmt.Println(a+b+c)
    }
}