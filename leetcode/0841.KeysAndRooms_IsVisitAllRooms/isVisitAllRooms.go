package lt841

/*
n个房间，初始在0
房间里存在钥匙，可以进入下一个房间
房间i，存在钥匙列表room[i]

是否可以进入所有房间

dfs
bfs
*/
func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 1{
		return true
	}
	n := len(rooms)
	visited := make([]int,n)
	visited[0] = 1
	queue := []int{0}
	for len(queue)!=0 {
		cur := queue[0]
		queue = queue[1:]  //出队
		keys := rooms[cur]
		for _,key := range keys {
			if visited[key] == 0 {
				//房间key未访问
				queue = append(queue, key)
				visited[key] = 1
			}
		}
	}
	//遍历检查
	for _,i := range visited{
		if i==0{
			return false
		}
	}
	return true
}


















