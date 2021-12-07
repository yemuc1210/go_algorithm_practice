package leetcode

/*
蛇梯旗
*/
// func id2rc(id, n int) (r, c int) {
//     r, c = (id-1)/n, (id-1)%n
//     if r%2 == 1 {
//         c = n - 1 - c
//     }
//     r = n - 1 - r
//     return
// }

// func snakesAndLadders(board [][]int) int {
//     n := len(board)
//     vis := make([]bool, n*n+1)
//     type pair struct{ id, step int }
//     q := []pair{{1, 0}}
//     for len(q) > 0 {
//         p := q[0]
//         q = q[1:]
//         for i := 1; i <= 6; i++ {
//             nxt := p.id + i
//             if nxt > n*n { // 超出边界
//                 break
//             }
//             r, c := id2rc(nxt, n) // 得到下一步的行列
//             if board[r][c] > 0 {  // 存在蛇或梯子
//                 nxt = board[r][c]
//             }
//             if nxt == n*n { // 到达终点
//                 return p.step + 1
//             }
//             if !vis[nxt] {
//                 vis[nxt] = true
//                 q = append(q, pair{nxt, p.step + 1}) // 扩展新状态
//             }
//         }
//     }
//     return -1
// }
func getRC(num,n int)(r,c int){
	// r = n-1 - (num-1)/n   //行是从小往上的
	r, c = (num-1)/n, (num-1)%n
	//列需要考虑奇偶行的影响
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return
}
func snakesAndLadders(board [][]int) int{
	n,m := len(board),len(board[0])   //n行m列
	visted := make([]bool,n*m)
	type cur_status struct{cur,step int}
	// queue := make([]cur_status,0)
	queue := []cur_status{{1,0}}   //切片
	for len(queue) > 0 {   //队列不为空
		//弹出队头
		cur := queue[0]
		queue = queue[1:]   //出队

		//下面访问队头的各种可能情况，共6个选项
		for i:=1;i<=6;i++{
			next := cur.cur + i   //存储下一个节点编号
			if next > n*n {
				break  //超出边界
			}
			//下面正常执行
			//首先，需要得到next的行列
			r,c := getRC(next,n)
			if board[r][c]>0{
				//存在蛇或梯子
				next = board[r][c]   //传递
			}
			if next==n*n{  //若直接到了终点
				return cur.step+1    //+1 是因为next这一步还未算
			}
			//否则，再判断当前节点是否入队
			if !visted[next]{
				visted[next] = true
				queue = append(queue, cur_status{next,cur.step+1})
			}
		}

	}
	return -1  //未在循环中返回，那就一定出错了

}
func snakesAndLadders_dfs(board [][]int) int {
    minReach := make([][]int, len(board))
    for i := 0; i < len(minReach); i++ {
		minReach[i] = make([]int, len(board[0]))
    }
    minReach[len(board)-1][0] = 1
    min := 2 << 31
    dfs(1, 0, len(board), &min, board, minReach)
    if min == 2 << 31 {
        return -1
    }
	return min
}

func dfs(idx, step, n int, min *int, board [][]int, minReach [][]int) {
    if step >= *min {
        return
    }
    if idx >= n * n {
		if *min > step {
			*min = step
		}
		return
	}
	x, y := getIdx(idx, n)
	if board[x][y] != -1 {
		idx = board[x][y]
	}
	if idx >= n * n {
		if *min > step {
			*min = step
		}
		return
	}
	x, y = getIdx(idx, n)
	if minReach[x][y] > 0 && minReach[x][y] <= step {
		return
	}
	minReach[x][y] = step
	for i := 1; i <= 6; i++ {
		dfs(idx + i, step + 1, n, min, board, minReach)
	}
	return
}

func getIdx(idx int, n int) (i int,j int) {
	i = n - 1 - (idx-1)/n
	if (n - 1 - i)%2 == 0 {
		j = (idx%n + n - 1)%n
	} else {
		j = (n - idx%n)%n
	}
	return
}