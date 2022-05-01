package main

import (
	"bufio"
	"fmt"
	"os"
)

//句子逆序  单词不逆序
func main(){
	//输入英文句子 单词见用空格隔开
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {  //多个输入
		sentence := reader.Text()
		// fmt.Println(sentence)

		//处理前导0
		idx,n := 0,len(sentence)
		for idx<n && sentence[idx] == ' '{
			idx ++
		}
		//获得单词列表
		words := []string{}
		for idx < n {
			//获得第一个单词
			word := []byte{}
			for idx<n && sentence[idx]!=' '{
				word = append(word, sentence[idx])
				idx ++
			}
			//处理中间的零
			for idx<n && sentence[idx] == ' '{idx ++}
			words = append(words, string(word))
		}
		//逆序输出
		for i:=len(words)-1;i>=0;i--{
			if i==0 {
				fmt.Print(words[i])
			}else{
				fmt.Print(words[i]+" ")
			}
		}
	}
}