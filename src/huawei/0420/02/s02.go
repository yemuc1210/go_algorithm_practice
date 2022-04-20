package s02

import (
	"fmt"
	"strings"
)

func getInput() ([]int,string,[]int){
	root := []int{}
	for {
		// 输入一个数组
		var c byte
		var a int
		fmt.Scanf("%d%c",&a,&c)
		if c=='\n' {
			break
		}
		root = append(root, a)
	}
	// 输入第二行
	var path string
	fmt.Scanln(&path)
	// 输入第三行 一个数组
	sub := []int{}
	for {
		// 输入一个数组
		var c byte
		var a int
		fmt.Scanf("%d%c",&a,&c)
		if c=='\n' {
			break
		}
		sub = append(sub, a)
	}
	fmt.Println(sub)
	// 求解
	return root,path,sub
	// solve(root,path,sub)
}

func solve(root []int, path string, sub []int) []int{
	// 二叉树直接以数组形式表示
	// 根下标为i  则左孩为2*i+1  右孩子为2*i+2
	// 正常数字为1-9 若数字为0表示此处为nil
	// 根据path 定位子树的根节点，然后替换
	// 1. 处理path
	path = strings.ReplaceAll(path,"/","")
	fmt.Println(path)
	
	return nil
}