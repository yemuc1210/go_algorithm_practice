package lt1436

/*
1436. 旅行终点站
给你一份旅游线路图，该线路图中的旅行线路用数组 paths 表示，
	其中 paths[i] = [cityAi, cityBi] 表示该线路将会从 cityAi 直接前往 cityBi 。
请你找出这次旅行的终点站，即没有任何可以通往其他城市的线路的城市。

题目数据保证线路图会形成一条不存在循环的线路，因此恰有一个旅行终点站。

出度为0的节点罢了

根据终点站的定义，终点只会出现在cityB在，遍历cityB,返回不在cityA的城市

*/
func destCity(paths [][]string) string {
	//map
	cityA := map[string]bool{}
	//初始化
	for _,path := range paths {
		cityA[path[0]] = true   //记录出现在cityA中的
	}

	for _,path := range paths{
		if !cityA[path[1]] {
			return path[1]
		}
	}
	return ""
}