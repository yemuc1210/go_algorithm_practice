package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

var (
    code1 = "0123456789abcdef"
        // 3 --> 0011 --> 1100 --> 12 --> 0x0C
    code2 = "084C2A6E195D3B7F"
)

func main() {
    r := bufio.NewScanner(os.Stdin)
    for {
        r.Scan()
        t := r.Text()
        if t == "" {
            break
        }
        tList := strings.Split(t, " ")
        so, sj := process(strings.Join(tList, ""))
        s := ""
        for {
            if len(so) != 0 {
                s += so[0]
                so = append(so[0:0], so[1:]...)
            }
            if len(sj) != 0 {
                s += sj[0]
                sj = append(sj[0:0], sj[1:]...)
            }

            if len(so) == 0 && len(sj) == 0 {
                break
            }
        }
        str := ""
        for _, v := range s {
            n := strings.ToLower(fmt.Sprintf("%c", v))
            if !strings.Contains(code1, n) {
                str += fmt.Sprintf("%c", v)
                continue
            }
            str += fmt.Sprintf("%c", code2[strings.Index(code1, n)])
        }
        fmt.Println(str)
    }
}

// 分奇偶，病排序
func process(s string) ([]string, []string) {
    sj := make([]string, 0, len(s))
    so := make([]string, 0, len(s))
    for k, ss := range s {
        if (k+1)%2 == 0 {
            so = append(so, fmt.Sprintf("%c", ss))
        } else {
            sj = append(sj, fmt.Sprintf("%c", ss))
        }

    }
    sort.Strings(sj)
    sort.Strings(so)
    return sj, so

}
