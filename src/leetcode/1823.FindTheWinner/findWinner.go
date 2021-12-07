package lt1823
//约瑟夫环问题
func findTheWinner(n int, k int) int {
	winner := 0 
	for i:=2;i<=n;i++{
		winner = (winner + k ) %i
	}
	return winner+1
}

//队列模拟
func findTheWinner1(n,k int)int{
	queue := []int{}
	for i:=0;i<n;i++{
		queue = append(queue, i+1)
	}
	cur:=0
	for i:=0;i<n-1;i++{
		cur += k
		cur %= len(queue)
		if cur==0{
			queue = queue[:len(queue)-1]
			continue
		}
		queue = append(queue[:cur-1], queue[cur:]...)
		cur--
	}
	return queue[0]

}