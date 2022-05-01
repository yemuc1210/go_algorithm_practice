package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
)

func main(){
    var params []string
    input := bufio.NewScanner(os.Stdin)
    for input.Scan(){
        params = append(params, input.Text())
    }
    for loop:= 0; loop< len(params); loop++{
        res := OutPrint(params[loop])
        fmt.Println(res)
    }
}

func OutPrint(s string) int{
    var res int
    nums, _ := strconv.Atoi(s)
    for loop:=1; loop<=nums; loop++{
        if isPrintPN(loop){
            res++
        }
    }
    return res
}

func isPrintPN(num int) bool {
    var tmp int
    for loop:=1 ;loop<num; loop++{
        if num% loop == 0{
            tmp += loop
        }
    }
    if tmp == num {
        return true
    }
    return false
}
