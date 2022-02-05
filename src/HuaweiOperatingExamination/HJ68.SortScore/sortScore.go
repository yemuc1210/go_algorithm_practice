package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type sm [][]string

func (s sm) Len() int {
    return len(s)
}

func (s sm) Less(i, j int) bool {
    a, _ := strconv.Atoi(s[i][1])
    b, _ := strconv.Atoi(s[j][1])
    return a < b
}

func (s sm) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {
    scan := bufio.NewScanner(os.Stdin)
    for {
        scan.Scan()
        if scan.Text() == "" {
            break
        }
        n, _ := strconv.Atoi(strings.Trim(scan.Text(), " "))
        s := ""
        scan.Scan()
        desc := scan.Text()
        nameList := make(sm, n)
        // 接收数据
        for i := 0; i < n; i++ {
            scan.Scan()
            nl := strings.Split(scan.Text(), " ")
            nls := []string{nl[0], nl[1]}
            nameList[i] = nls
            s += strings.Join(nls, "")
        }
        // 排序
        sort.Sort(nameList)

        newL := make([][]string, n)
        j := 0
        // 正逆调整
        if desc == "0" {
            for i := n - 1; i >= 0; i-- {
                newL[j] = nameList[i]
                j++
            }
        } else {
            for i := 0; i < n; i++ {
                newL[j] = nameList[i]
                j++
            }
        }
        // 相同成绩录入顺序对比
        for i := 0; i < n; i++ {
            for j := i; j < n; j++ {
                if newL[i][1] == newL[j][1] {
                    a := strings.Index(s, strings.Join(newL[i], ""))
                    b := strings.Index(s, strings.Join(newL[j], ""))
                    if a > b {
                        newL[i], newL[j] = newL[j], newL[i]
                    }
                }

            }

        }
        for _, v := range newL {
            fmt.Println(fmt.Sprintf("%s %s", v[0], v[1]))

        }
    }
}