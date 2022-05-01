package main

import(
    "fmt"
    "os"
    "bufio"
    "sort"
    "strings"
    "strconv"
)
const maxn = 1e7+5
var reader *bufio.Reader
var a [maxn]int

func HasJinShanShan(l,r int) bool{

    if r-l+1<3{
        return false
    }else if r-l+1>47{
        return true
    }else{
        weapon := []int{}
        for i:=l; i<=r;i++{
            weapon = append(weapon, a[i])
        }
        sort.Ints(weapon)
        for i:=0; i<len(weapon)-2;i++{
//             前两条边大于第三条，能构成
            if weapon[i]+weapon[i+1]>weapon[i+2]{
                return true
            }
        }
    }
    return false
}

func main(){
    reader = bufio.NewReader(os.Stdin)
    var n int

    nStr := splitStdString('\n')
    n,_ = strconv.Atoi(nStr)
    for i:=1; i<n; i++{
//         fmt.Print(buf, " ")
        buf := splitStdString(' ')
        a[i],_ = strconv.Atoi(buf)
    }
    buf := splitStdString('\n')

    a[n],_ = strconv.Atoi(buf)
//     fmt.Println(a[:n+1])

    var m int
    var l,r ,cnt int
    mStr := splitStdString('\n')
    m,_ = strconv.Atoi(mStr)
//     fmt.Println(m)
    for i:=0; i<m; i++{
        lrStr := splitStdString('\n')
        lr := strings.Split(lrStr, " ")

//         fmt.Println("lr",lr)
        l,_ = strconv.Atoi(lr[0])
        r,_ = strconv.Atoi(lr[1])
//         fmt.Println(l,r,HasJinShanShan(l,r))
        if HasJinShanShan(l,r){cnt++}
    }
    fmt.Println(cnt)
}

// 以split字符分割标准输入的字串
func splitStdString(split byte) string{
    text,_ := reader.ReadString(split)
    text = strings.Trim(text, string(split))
    return text
}