package main

// 我的日程安排表III
// 当K个日程安排有一些时间上的交叉，例如k个日程安排在同一时间内，就会产生k次预订
// 给定一些日程安排[start,end]，请在每个日程安排添加后，返回一个整数k
// 表示所有先前日程安排会产生的最大k预订

// 实现MyCalendarThree类存放日程安排，可以一直添加新的日程安排

// 使用树存储，区间树，节点存储子树中节点的数量。
// 以时间为区间查找日程
// 线段树

type pair struct{
	// 懒标记lazy标记区间[l,r]进行累加的次数
	num,lazy int
}

type MyCalendarThree map[int]pair

// 初始化对象
func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}
func (t MyCalendarThree) update(start,end,l,r,idx int) {
	if r<start || end<l {
		return
	}
	if start<=l && r<=end {
		p:=t[idx]
		p.num++
		p.lazy++
		t[idx]=p
	}else{
		mid := (l+r)/2
		t.update(start,end,l,mid,idx*2)
		t.update(start,end,mid+1,r,idx*2+1)
		p:=t[idx]
		p.num=p.lazy+max(t[idx*2].num,t[idx*2+1].num)
		t[idx] =p
	}
}
func max(a,b int) int {
	if a>b {return a}
	return b
}
// 添加一个行程，返回一个整数k，表示日历中存在的k次预订的最大值
func (this *MyCalendarThree) Book(start int, end int) int {
	this.update(start,end-1,0,1e9,1)
	return (*this)[1].num
}