package offerS35


/*
剑指 Offer II 035. 最小时间差

给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。

将时间变为数字
*/
func findMinDifference(timePoints []string) int {
	const Max = 1440   //最大值
	minutesArr := [Max]int{}
	for _,v := range timePoints {
		// 计算时间
        t := int(v[0] - '0') * 10 * 60 +
        int(v[1] - '0') * 60 +
        int(v[3] - '0') * 10 +
        int(v[4] - '0')

		minutesArr[t] ++
	}
	
	time := []int{}
    for t, count := range minutesArr {  //k v
        if count > 1 {
            return 0
        }
        if count == 1 {
            time = append(time, t)
        }
    }

    // 计算每个时间段的间隔
    min := 0
    for i := 1; i < len(time); i++ {
        if min == 0 || min > (time[i] - time[i-1]) {
            min = time[i] - time[i-1]
        }
    }
    // 最后别漏了末尾时间到起始时间的间隔
    ring := (Max - time[len(time)-1]) + time[0]
    if ring < min {
        min = ring
    }

    return min
	
}

