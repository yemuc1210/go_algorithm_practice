package main

import "fmt"

func countOperations(num1 int, num2 int) int {
    var res int
    for num1!=0 && num2!=0 {
        if num1>=num2 {
            num1 -= num2
			
        }else{
            num2 -= num1
        }
		fmt.Println(num1)
        res ++ 
    }
    return res
}

func main() {
	res := countOperations(2,3)
	fmt.Println(res)
}