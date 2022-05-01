package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    var numNodes int
	fmt.Scanln(&numNodes)
    // numNodes,_ := strconv.Atoi(nums)
	// fmt.Println(numNodes)
	//输入链表的值
	nodeSlice := make([]int,numNodes+1)
	nodes,_:= reader.ReadString('\n')
	// fmt.Println(nodes)
	
	nodes = strings.ReplaceAll(nodes,"\n"," ")
	s := strings.Split(nodes," ")
	// fmt.Println(s,s[len(s)-1],len(s[len(s)-1]))
	for i:=0;i<len(s);i++{
		nodeSlice[i],_ = strconv.Atoi(s[i])
	}
	// fmt.Println(nodeSlice)

	var k int
	fmt.Scanln(&k)  //要求k<=n
	if len(nodeSlice) < k {
		return 
	}
	n := len(nodeSlice)
	
	fmt.Println(nodeSlice[n-k-1])

    
}