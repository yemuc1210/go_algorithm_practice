package lt575
/*
575. 分糖果
给定一个  偶数长度  的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。
你需要把这些糖果   平均  分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。

所有糖不重样，一半
中 每个糖都有两个  能得到所有种类，同上
一种糖，一个 得到一个
*/
func distributeCandies(candyType []int) int {
	// min{len(类型)，len(candies)//2}
	types := make(map[int]int)
	for _,v:= range candyType {
		types[v] ++
	}
	var min func(a,b int)int
	min = func(a, b int) int {
		if a<b{
			return a
		}
		return b
	}
	return min(len(types), len(candyType)/2)
}
