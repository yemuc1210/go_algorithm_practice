package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

/*
HJ1 字符串最后一个单词的长度
较难  通过率：30.23%  时间限制：1秒  空间限制：32M
知识点	字符串
描述
计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000。
（注：字符串末尾不以空格为结尾）
输入描述：
输入一行，代表要计算的字符串，非空，长度小于5000。

输出描述：
输出一个整数，表示输入字符串最后一个单词的长度。
*/

func readLine() string{
	inputReader := bufio.NewReader(os.Stdin)
	input,err := inputReader.ReadString('\n')  //以换行为结束符 读取一行
	if err != nil {
		return ""
	}
	inputReader.ReadByte()
	return input
}
func lengthOfLastWord(line string) int {
	if len(line) == 0{
		return 0
	}
	
	//处理前导0
	idx,n := 0,len(line)
	for idx<n && line[idx]==' '{idx++}
	//得到第一个字符
	res := 0
	fmt.Println()
	for idx<n && line[idx]!=' '{
		res++
		idx ++
		fmt.Printf("%c %d\n",line[idx],res)
	}
	return res
}

func main(){
	input := readLine()
	fmt.Println(input)
	input = reverseString(input)
	fmt.Print(input)
	fmt.Println(lengthOfLastWord(input))
}

func reverseString(in string) string{
	//将字符串逆置，则最后一个变为第一个，当然是逆置的
	n := len(in)
	byteIn := []byte(in)
	for i:=0;i<n/2;i++ {
		byteIn[i],byteIn[n-i-1] = byteIn[n-i-1],byteIn[i]
	}
	return string(byteIn)
}