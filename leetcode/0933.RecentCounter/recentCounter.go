package lt933
/*
剑指 Offer II 042. 最近请求次数
写一个 RecentCounter 类来计算特定时间范围内最近的请求。

请实现 RecentCounter 类：

RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，
并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。
确切地说，返回在 [t-3000, t] 内发生的请求数。
保证 每次对 ping 的调用都使用比之前更大的 t 值。

关键利用：新加入的ping调用，时间上是最大的，所以每次向前查询即可
可以用双端队列实现。每次将 t 进入队尾，同时从队头开始依次移除小于 t-3000 的元素。
然后返回队列的大小 q.size() 即可。
*/
type RecentCounter struct {
	queue []int
}


func Constructor() RecentCounter {
	return RecentCounter{queue: []int{}}
}


func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	for this.queue[0] < t - 3000 {
		//出
		this.queue = this.queue[1:len(this.queue)]
	}
	return len(this.queue)
}

