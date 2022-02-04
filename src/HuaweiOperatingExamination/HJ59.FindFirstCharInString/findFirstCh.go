package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main(){
    var params []string
    input := bufio.NewScanner(os.Stdin)
    for input.Scan(){
        params = append(params, input.Text())
    }
    for loop:=0 ;loop< len(params); loop++{
        res := paramNum(params[loop])
        fmt.Println(res)
    }
}

func paramNum(s string) string {
    for loop:=0 ;loop< len(s); loop++{
        if strings.Count(s, string(s[loop])) == 1{
            return string(s[loop])
        }
    }
    return "-1"
}