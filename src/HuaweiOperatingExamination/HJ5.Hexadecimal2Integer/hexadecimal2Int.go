package main
import (
    "bufio"
    "os"
    "fmt"
)

//十六进制转十进制

func main(){
    //接收十六进制  输出十进制
    in := bufio.NewScanner(os.Stdin)
//     var helper = "0123456789ABCDEF"
    for in.Scan() {
        par := in.Text()
        //首先取出前两个字符  0x 只是一个标记
        if par == "" {
            return 
        }
        if par[:2] == "0x" {
            //取出前两个
            par = par[2:]
        }
        //转换
        res := 0
        n := len(par)
        idx := 0
        x := 1
        for idx<n {
            if par[n-idx-1]>='0' && par[n-idx-1]<='9' {
                res += int(par[n-idx-1]-'0') * x
            }else if par[n-idx-1]>='A'&&par[n-idx-1]<='F'{
                res += (int(par[n-idx-1]-'A') + 10) * x
            }
            x *= 16
            idx++
        }
        fmt.Println(res)
    }
}