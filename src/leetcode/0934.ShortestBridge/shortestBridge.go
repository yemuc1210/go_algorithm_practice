package main
// 存在两座岛,由四面1组成
// 现将0变1,使得两岛相连
// 返回必须翻转的0的最小个数
func shortestBridge(grid [][]int) int {
	var landArr [][]int
	var findLand func(int,int)
	var bridgeNum int
	var endNum int
	var queue [][]int
	findLand = func(i int, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid){
			return
		}
		if grid[i][j] == 1{
			grid[i][j] = 2
			landArr = append(landArr, []int{i,j})
			findLand(i+1,j)
			findLand(i-1,j)
			findLand(i,j+1)
			findLand(i,j-1)
		}else if grid[i][j] == 0{
			queue = append(queue, []int{i,j})
			return
		}
		return
	}
	for i := 0;i<len(grid);i ++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j] == 1{
				findLand(i,j)
				goto Loop
			}
		}
	}
Loop:
	var findNext func(int,int)
	findNext = func(i int,j int) {
		if i < 0 || j <0 || i >= len(grid) || j>= len(grid[0]){
			return
		}
		if grid[i][j] == 0 {
			grid[i][j] = 2
			queue = append(queue, []int{i,j})
		}else{
			if grid[i][j] == 1{
				endNum = bridgeNum
			}
		}
		return
	}
	for len(queue) > 0{
		if endNum > 0{
			break
		}
		bridgeNum ++
		for tmpLen := len(queue) - 1;tmpLen >= 0;tmpLen --{
			tmpi := queue[0][0]
			tmpj := queue[0][1]
			findNext(tmpi+1,tmpj)
			findNext(tmpi-1,tmpj)
			findNext(tmpi,tmpj+1)
			findNext(tmpi,tmpj-1)
			queue = queue[1:]
		}
	}

	return endNum
}

func main(){

}