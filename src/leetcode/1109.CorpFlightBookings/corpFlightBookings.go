package lt1109
/*
每日一题
这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 
	意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。



差分数组+前缀和

一个预定记录代表一个区间的增量，增量叠加得到答案
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _,arr := range bookings {
		res[arr[0] - 1] += arr[2]  //因为编号是1-n
		if arr[1] < n {
			res[arr[1]] -= arr[2]
		}
	}
	for i:=1;i<n;i++{
		res[i] += res[i-1]  //求前缀和
	}
	return res

}