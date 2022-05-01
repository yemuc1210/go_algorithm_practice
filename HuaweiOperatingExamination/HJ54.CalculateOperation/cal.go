package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
    "unicode"
)

func main(){
    var params []string
    input := bufio.NewScanner(os.Stdin)
    for input.Scan(){
        params = append(params, input.Text())
    }
    for loop:= 0; loop< len(params); loop++{
        tmp := strings.Replace(strings.Replace(params[loop], "]", ")", -1), "[", "(", -1)
        tmp = strings.Replace(strings.Replace(tmp, "}", ")", -1), "{", "(", -1)
        res := OutPrint(tmp)
        fmt.Println(res)
    }
}

func OutPrint(s string) int {
    bytes := make([]byte, len(s))
    copy(bytes, s)
    return subCal(&bytes)
}

func subCal(bytes *[]byte) int {
    stack := []int{}
    num := 0
    var sigh byte = '+'
    for len(*bytes) > 0 {
        item := (*bytes)[0]
        *bytes = (*bytes)[1:]
        if unicode.IsDigit(rune(item)) {
            val, _ := strconv.Atoi(string(item))
            num = num*10 + val
        }
        if item == '(' {
            num = subCal(bytes)
        }
        if (!unicode.IsDigit(rune(item)) && item != ' ') || len(*bytes) == 0 {
            switch sigh {
                case '+':
                    stack = append(stack, num)
                case '-':
                    stack = append(stack, -num)
                case '*':
                    num *= stack[len(stack)-1]
                    stack = stack[:len(stack)-1]
                    stack = append(stack, num)
                case '/':
                    num = stack[len(stack)-1]/num
                    stack = stack[:len(stack)-1]
                    stack = append(stack, num)
            }
        sigh = item
        num = 0
        }
        if item == ')' {
            break
        }
    }
    sum := 0
    for i := 0; i < len(stack); i++ {
        sum += stack[i]
    }
    return sum
}

