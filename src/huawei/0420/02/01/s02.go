package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	root := []int{}
	for {
		// 输入一个数组
		var c byte
		var a int
		fmt.Scanf("%d%c",&a,&c)
		root = append(root, a)
		if c=='\n' {
			break
		}
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
	// fmt.Println(root,path,sub)
	// root :=[]int{1,1,2,0,0,4,5}
	// path:="/1/1"
	// sub:=[]int{5,3,0}
	// 求解
	root = solve(root,path,sub)
	for i:=0;i<len(root);i++{
		if root[i]!=0 {
			fmt.Printf("%d ",root[i])
		}
	}
	fmt.Println()
}
// root-line1 
func solve(root []int, path string, sub []int) []int{
	// 1. 处理path
	path = strings.ReplaceAll(path,"/","")
	pathNums := strings.Split(path,"")
	nums := []int{}
	for _,v := range pathNums {
		if v !=" " {
			n,_ := strconv.Atoi(v)
			nums = append(nums, n)
		}
	}
	// 2. 遍历
	n := len(root)
	j,i := 0,0
	for i<n && j<len(nums) {
		if root[i] == nums[j] {
			if (j+1) < len(nums) {
				j++
			}else{
				break
			}
			if root[2*i+1] == nums[j] {
				i = 2*i+1
			}else if root[2*i+2] == nums[j] {
				i = 2*i+2
			}
		}
	}
	// 将以i为根的替换
	// 清除i为根的部分
	clear(root,i,n)
	// 问题：子树的节点可能非常多
	// 对root重新申请空间
	tmp := make([]int,n)   
	copy(tmp, root)     // 备份
	// 重新申请空间
	root = make([]int, len(tmp)+len(sub))   // 此时空间绰绰有余，因为tmp里面包含需要被删除的子树的长度
	copy(root,tmp)
	replace(root,i,sub,0,n,len(sub))
	return root
}

func replace(root []int, i int, sub []int, j int,n int, m int) {
	if i<n && j< m{
		root[i] = sub[j]
		replace(root,i*2+1, sub, j*2+1, n,m)
		replace(root,i*2+2,sub,j*2+2,n,m)
	}
}

func clear(root []int, i int, n int) {
	if i<n{
		root[i] = 0
		// 递归
		clear(root,2*i+1,n)
		clear(root, 2*i+2,n)
	}
}