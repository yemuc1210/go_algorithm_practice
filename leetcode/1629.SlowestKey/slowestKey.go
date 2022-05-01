package lt1629

/*
按键持续时间最长的键   0109  简单
*/
func slowestKey(releaseTimes []int, keysPressed string) byte {
	//一次遍历
	res := keysPressed[0]
	
	maxTime := releaseTimes[0]
	for i:=1;i<len(keysPressed);i++ {
		key := keysPressed[i]
		time := releaseTimes[i] - releaseTimes[i-1]

		if time>maxTime || time == maxTime && key>res {
			res = key
			maxTime = time
		}
	}
	return res
}