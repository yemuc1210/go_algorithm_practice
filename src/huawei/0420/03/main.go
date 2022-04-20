package main

import "fmt"

type node struct {
	num   int
	left  *node
	right *node
}

func main() {
	n := 8
	id := []int{16, 12, 3, 10, 5, 2, 4, 1}
	pid := []int{0, 16, 16, 12, 3, 3, 2, 2}
	task := 7
	fmt.Println(n,id,pid,task)

	// 构建树
	root := &node{num:id[0]}
	i:=1
	cur := root
	for i<len(pid){
		if pid[i] == pid[i+1] {
			root.left = &node{num: id[i]}
			root.right = &node{num: id[i+1]}
			i+=2
			// 递归
		}else{
			root.left = &node{num: id[i]}
			i+=1
		}
	}
	// pathSum问题

}

func buildTree(root *node, i int, id,pid []int,cnt *int) int{
	if i<len(pid) {
		if pid[i] == pid[i+1] {
			root.left = &node{num: id[i]}
			root.right = &node{num: id[i+1]}
			// 递归
			k:=0
			buildTree(root.left,i+2,id,pid,&k)
			buildTree(root.right,i+2+k,id,pid,&k)
		}else{
			root.left = &node{num: id[i]}
			i+=1
		}
	}
}
