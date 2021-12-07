package lt447
/*

给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。

回旋镖 是由点 (i, j, k) 表示的元组 ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。

返回平面上所有回旋镖的数量。


*/


//枚举+哈希
/*
V型折线，锚点是point[i]    得到m个距离point[i]相等的点，从这其中选两个。由于考虑元组顺序，所以方案数是m中选两个的排列A_m^2=m*(m-1)

遍历数组，统计所有点到point[i]的距离，将每个距离出现次数记在哈希表；然后遍历哈希表
*/
func numberOfBoomerangs(points [][]int) int {
	res :=0
	for _,p := range points {
		//统计每个距离p的值
		cnt := map[int]int{}
		for _,q := range points{
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])   //a^2+b^2
			cnt[dis] ++
		}

		//根据公式计算
		for _,m := range cnt {
			res += m*(m-1)
		}
	}
	return res
}