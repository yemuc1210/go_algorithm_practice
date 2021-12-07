/*
给你一个数组 points ，
其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。
求最多有多少个点在同一条直线上。

提示：

1 <= points.length <= 300
points[i].length == 2
-104 <= xi, yi <= 104
points 中的所有点 互不相同

*/
package leetcode



// import "fmt"

// type point struct{
// 	x int
// 	y int
// }
// type location struct{
// 	dis_x int
// 	dis_y int
// }
// func maxPoints(points [][]int)int{
// 	n := len(points)     //n组点
// 	if n==2{
// 		return 2
// 	}
// 	loc := make([][]location,n)
// 	for i := range loc{
// 		loc[i] = make([]location, n)
// 	}
// 	// fmt.Println(n,points[0][1],points[5][1])
// 	cnt := 0
// 	for i:=0;i<n;i++{
// 		for j:=i;j<n;j++{
// 			disx := points[j][0] - points[i][0]
// 			disy := points[j][1] - points[i][1]       //j点距离第一个点的距离  i为基准
// 			loc[j][i] = location{dis_x: disx,dis_y: disy}
// 			loc[i][j] = location{dis_x: -disx,dis_y: -disy}
// 		}
// 	}  //复杂度 n^2
// 	//统计半角中相同元素，或者成比例的，用map
// 	m := make(map[location]int)
// 	is_add := false
// 	var last location
// 	for i:=0;i<n;i++{
// 		for j:=0;j<i;j++{
// 			tmp := location{dis_x: loc[i][j].dis_x%2,dis_y: loc[i][j].dis_y%2}
// 			//现在就是正负可能相反了，那么统一第一个数字为整数
// 			if tmp.dis_x >= 0{
// 				if last.dis_x==tmp.dis_x && last.dis_y==tmp.dis_y{
// 					is_add = true
// 				}else{
// 					is_add = false
// 				}
// 				m[tmp]++
// 			}else{
// 				tmp.dis_x *=-1
// 				tmp.dis_y *=-1
// 				if last.dis_x!=tmp.dis_x && last.dis_y!=tmp.dis_y{
// 					is_add = true
// 				}else{
// 					is_add = false
// 				}
// 				m[tmp]++
// 			}
			
// 		}
// 	}
// 	//输出最大的哪一个

// 	for _,v := range m {
// 		if v > cnt {
// 			cnt = v
// 		}
// 	}
// 	if is_add{

// 	}
// 	return cnt
// }

// 直线上最长的点

// 思考：
// 暴力搜索比较，每两个点作为一条直线，判断其他点是否在这条直线上

// 暴力直观的思路
// 直线的两点表示法： (y-y1)/(x-x1) = (y2-y1)/(x2-x1)
// 直线的斜率表示法： y = kx + b
// 这两个表示法中，两点表示方法方便直观地判断一个点是否在一条直线上
// 但由 y = kx + b 可以使用 k,b两参数确定一条直线，计算好k,b之后，
// 可以用其表示一条独一无二的直线，这可以有效的去重!!
// 由于浮点数不能直接比较，所以k,b都必须使用两个整数来表示分子与分母(并且要化到最简)
// 判断某个点是否在直线上则是 y = k1/k2 * x + b1/b2  => y*k2*b2 ?= k1*b2*x + k2*b1
func maxPoints(points [][]int) int {
	n := len(points)	// 每个point是个长度为2的数组，就不检查长度是否为2了
	if n < 3 {return n}	// 特殊情况

	// 特殊情况：测例认为所有点相同，也算是直线，并返回n
	i := 0
	for ; i<n-1; i++ {
		if points[i][0] != points[i+1][0] || points[i][1] != points[i+1][1] {break}
	}
	if i==n-1 {return n}

	// 用于记录直线的哈希表，值为直线所拥有的点数量
	lines := make(map[Line]int)

	// 遍历points，统计所有line包含的点数
	var line Line
	count := 0
	for i:=0; i<n; i++ {	// 这两层循环保证将所有两点的组合都取出
		for j:=i+1; j<n; j++ {
			// 避免两个点相同
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {continue}

			// 固定了一条直线，当前这轮中就需要检查所有点是否在直线上
			line = NewLine(Point{points[i][0], points[i][1]}, Point{points[j][0], points[j][1]})

			// 检查line是否前面生成过，由于前面的检查的点位更多，所以只要
			// 重复出现了line，那后面重复出现的都没必要去查看line上有几个点
			if lines[line] != 0 {continue}

			// line不曾出现过，则去遍历这条line有多少个点
			count = 2	// 线原本就有两点
			for a:=0; a<n; a++ {
				if a==i || a==j {continue}
				if line.HasPoint(Point{points[a][0], points[a][1]}) {count++}
			}

			// 记录line对应的点数
			lines[line] = count
			//fmt.Println(line, count)
		}
	}

	// 再遍历一遍lines哈希表，得到最大值
	max := 0
	for _, v := range lines {
		if v > max {max = v}
	}

	return max
}

// 由于测试时需要多次调用，所以在函数外定义较好

// 定义Point，便于理解和实现（直接用数组也行，但是容易疏忽写错）
type Point struct {
	x, y int
}


// 定义斜率K
type K struct {
	k1, k2 int	// k1分子,k2分母
}

// 定义截距B
type B struct {
	b1, b2 int	// b1分子,b2分母
}

// 定义Line，便于理解实现
type Line struct {
	K
	B
	c int	// c默认为0(非竖直线情况)，若为垂直线，方程为 x=c，k1=k2=b1=b2=0（其实只需要k2=0就足够了）
			// 以作判断(注意不能直接用c==0判断，因为可能有x=0的竖直线)
}

func NewLine(p,q Point) Line {
	// k = (y2-y1)/(x2-x1); b = y1-k*x1

	// 特殊情况是垂直线，x=x2, (k=inf，b=inf)
	if p.x == q.x {
		return Line{c:p.x}
	}

	// 计算斜率
	k1, k2 := p.y - q.y, p.x - q.x
	g1 := gcd(k1, k2)
	k := K{k1/g1, k2/g1}

	// 计算截距
	b1, b2 := q.y * k.k2 - k.k1 * q.x, k.k2
	g2 := gcd(b1, b2)
	b := B{b1/g2, b2/g2}

	return Line{K:k,B:b}
}

// 判断线上是否有这么一点
// 判断某个点是否在直线上则是 y = k1/k2 * x + b1/b2  => y*k2*b2 ?= k1*b2*x + k2*b1
func (l Line) HasPoint(p Point) bool {
	if l.k2==0 {	// 竖直线
		return p.x == l.c
	}
	return p.y * l.k2 * l.b2 == l.k1 * l.b2 * p.x + l.k2 * l.b1
}
// 辗转相除， 计算公约数
func gcd(a,b int) int {
	for b!=0 {
		a, b = b, a%b
	}
	return a
}
