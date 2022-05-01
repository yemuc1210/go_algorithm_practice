package lt407
/*lt42
407. 接雨水 II   困难
给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，
请计算图中形状最多能接多少体积的雨水。
*/

/*
BFS
假设初始矩阵每个格子接满水，高度均为maxHeight（高度最高的格子）
方块接水后的高度为
	water[i][j]=max（heightMap[i],min(water[i-1])[j],water[i+1][j],
				water[i][j-1],water[i][j+1]))
方块能接住水；该方块不是外层，自身高度比四周接水后的高度还要低

方块（i，j）的实际接水容量  water[i][j] - heightMap[i][j]
我们首先假设每个方块 (i,j)的接水后的高度均为water[i][j]=maxHeight，
首先我们知道最外层的方块的肯定不能接水，所有的多余的水都会从最外层的方块溢出，
我们每次发现当前方块 (i,j)的接水高度water[i][j] 小于与它相邻的 4 个模块的接水高度时，
则我们将进行调整接水高度，
我们将其相邻的四个方块的接水高度调整与 (i,j) 的高度保持一致，
我们不断重复的进行调整，直到所有的方块的接水高度不再有调整时即为满足要求。

*/
func trapRainWater(heightMap [][]int) (ans int) {
    m, n := len(heightMap), len(heightMap[0])
    maxHeight := 0
    for _, row := range heightMap {
        for _, h := range row {
            maxHeight = max(maxHeight, h)
        }
    }  //得到所有格子中的最大高度

    water := make([][]int, m)   //接水高度
    for i := range water {
        water[i] = make([]int, n)   // water:m*n维
        for j := range water[i] {
            water[i][j] = maxHeight   //初始化
        }
    }
    type pair struct{ x, y int }
    q := []pair{}
    for i, row := range heightMap {  //row 1 4 3 1 3 2
        for j, h := range row {    //h 格子的高度
            if (i == 0 || i == m-1 || j == 0 || j == n-1) && h < water[i][j] {
                water[i][j] = h  //边缘的格子不接水
                q = append(q, pair{i, j})
            }
        }
    }

    dirs := []int{-1, 0, 1, 0, -1}
    for len(q) > 0 {
        p := q[0]
        q = q[1:] // BFS 出队操作
        x, y := p.x, p.y
        for i := 0; i < 4; i++ {
            nx, ny := x+dirs[i], y+dirs[i+1]
            if 0 <= nx && nx < m && 0 <= ny && ny < n && water[nx][ny] > water[x][y] && water[nx][ny] > heightMap[nx][ny] {
                water[nx][ny] = max(water[x][y], heightMap[nx][ny])
                q = append(q, pair{nx, ny})
            }
        }
    }

    for i, row := range heightMap {
        for j, h := range row {
            ans += water[i][j] - h
        }
    }
    return
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

