package lt846

/*同1296
846. 一手顺子
Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，
使每一组的牌数都是 groupSize ，并且由 groupSize 张连续的牌组成。

给你一个整数数组 hand 其中 hand[i] 是写在第 i 张牌，和一个整数 groupSize 。
如果她可能重新排列这些牌，返回 true ；否则，返回 false 。

1296. 划分数组为连续数字的集合
给你一个整数数组 nums 和一个正整数 k，请你判断是否可以把这个数组划分成一些由 k 个连续数字组成的集合。
如果可以，请返回 true；否则，返回 false。
*/
func isNStraightHand(hand []int, groupSize int) bool {
	//map
	m := make(map[int]int)   //num-cnt
	for _,v := range hand{
		m[v] ++
	}

	//得到元素及其个数，然后遍历
	for _,v := range hand{
		if m[v] == 0 {
			continue
		}

		//找到v所在的左右端
		l,r := v,v
		for m[l-1] > 0 {
			l--
		}
		for m[r+1] > 0 {
			r++
		}

		//[l,r]连续，可能作为分割
		for l<=r {
			if m[l] != 0 {   //[l ,l+k-1]
				for i:=0;i<groupSize;i++{
					if m[l+i] == 0 {
						return false
					}  //不连续
					m[l+i]--
				}
			}else{
				l++
			}
		}
	}
	return true
}