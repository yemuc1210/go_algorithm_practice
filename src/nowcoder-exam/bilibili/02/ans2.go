package main

import "fmt"

//超时   c++可以通过。。。

var fa [10001]int
func findParent(x int) int {
	if fa[x] != x {
		fa[x] = findParent(fa[x])
	}
	return fa[x]
}

func Main() {
	var n,ai,m int
	fmt.Scanln(&n)
	fmt.Scanln(&ai)
	fmt.Scanln(&m)

	//初始化
	for i:=0;i<n;i++{
		fa[i] = i
	}

	var cnt int
	//得到m组关系对
	for i:=0;i<m;i++{
		var a,b int
		fmt.Scanf("%d,%d",&a,&b)
		pa,pb := findParent(a),findParent(b)
		if a==ai || b==ai {
			cnt--   //已经认识的
		}
		if pa!=pb {
			//合并 简化路径
			fa[pa] = pb    //pa的父亲/祖先 变为pb
		}
	}

	pai := findParent(ai)
	for i:=0;i<n;i++{
		if findParent(i) == pai {
			cnt ++
		}
	}
	fmt.Println(cnt-1)   //减去自身
}