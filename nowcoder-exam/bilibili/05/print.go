package main

import(
    "fmt"
    "strings"
    "strconv"
    "os"
    "bufio"
)

var r *bufio.Reader
var rowLen, colLen int

func printM(m [][]int) {
    
//     矩阵外圈一层框架遍历
    rows, rowe, cols, cole := 0, rowLen-1, 0, colLen-1
    for rows<= rowe&&cols<=cole{
        for i:=cols; i<=cole; i++{
            if i == 0{fmt.Print(m[cols][i])}else{fmt.Print(",",m[cols][i])}
        }
        for i:=rows+1; i<=rowe; i++{
            fmt.Print(",",m[i][cole])
        }
        for i:=cole-1; i>=cols&&rowe!=rows; i--{
            fmt.Print(",",m[rowe][i])
        }
        for i:=rowe-1; i>rows&&cole!=cols; i--{
            fmt.Print(",",m[i][cols])
        }
        
        rows++
        rowe--
        cols++
        cole--
    }
    fmt.Println()
}

// 以split字符分割标准输入的字串
func splitStdString(split byte) string{
	text,_ := r.ReadString(split)
	text = strings.Trim(text, string(split))
	return text
}

func main() {
    r = bufio.NewReader(os.Stdin)

    for {
        matrixLen := splitStdString('\n')
        rc := strings.Split(matrixLen, " ")
        if matrixLen == "-1 -1"{
            break
        }
        rowLen, _ = strconv.Atoi(rc[0]) 
        colLen, _ = strconv.Atoi(rc[1]) 
//         开辟矩阵
        m := make([][]int, rowLen)
        for i:= 0; i<rowLen; i++{
            m[i] = make([]int, colLen)
        }
        var numStr string
        var numByte []string
        for i:=0; i<rowLen;i++{
            numStr = splitStdString('\n')
            numByte = strings.Split(numStr, " ")
            for j:=0; j<colLen; j++{
                m[i][j],_ = strconv.Atoi(numByte[j])
            }
        }
//         fmt.Println(m)
        printM(m)
    }
}