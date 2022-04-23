package main

import (
	"fmt"

	"sort"
)

// 最短的绳子围所有的树
// 返回边界上树的坐标

// 根据函数名，找最外层的树？
// 最短：贪心   凸包算法

// 以最小y值的点p做原点，与其他点的连线的极角逐渐增大
// 按照与p的极角值进行排序
// 遍历排序后的n-1个点，只保留那些进行逆时针旋转的点
type point struct {
	x int
	y int
	angle float64
}
func outerTrees(trees [][]int) [][]int {
	n := len(trees)
	if n<4 {
		return trees
	}
	// 1 找到y值最小的
	bottom := 0 
	for i,tree := range trees {
		if tree[1] < trees[bottom][1] {
			bottom = i
		}
	}
	trees[bottom],trees[0] = trees[0],trees[bottom]

	// 2 求解剩余点和p的极角值
	tr := trees[1:]
	sort.Slice(tr, func(i, j int) bool {
		// 以bottom为原点
		a,b := tr[i],tr[j]
		diff := cross(trees[0],a,b) - cross(trees[0],b,a)
		// diff=0 在一条线，则根据距离排序
		return diff>0 || diff==0 && distance(trees[0],a) < distance(trees[0],b)
	})
	// 对于凸包最后且在同一条直线上的元素安卓距离从大到小排序
	r := n-1
	for r>=0 && cross(trees[0],trees[n-1],trees[r])==0 {
		r--
	}
	for l,h := r+1,n-1;l<h;l++{
		trees[l],trees[h] = trees[h],trees[l]
		h--
	}
	sk := []int{0,1}
	for i:=2;i<n;i++{
		// 如果当前元素与栈顶两个元素顺时针，则弹出
		for len(sk)>1 && cross(trees[sk[len(sk)-2]], trees[sk[len(sk)-1]], trees[i]) < 0 {
			sk = sk[:len(sk)-1]
		}
		// 逆时针 入栈
		sk = append(sk, i)
	}
	res := [][]int{}
	for _,idx := range sk {
		res = append(res, trees[idx])
	}
	return res
}
// 计算叉积 cross(p,q,r)=pq * qr
// 若叉积小于0，顺时针  大于0，逆时针，左拐   等于0，三点一条直线
// 若当前点与上一条线之间是右拐（顺时针）某则上一点不应该被保留
func cross(p, q, r []int) int {
    return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
}

func distance(p, q []int) int {
    return (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
}



func main(){
	trees := [][]int{
		{1,1},{2,2},{2,0},{2,4},{3,3},{4,2},
	}
	res := outerTrees(trees)
	fmt.Println(res)
}