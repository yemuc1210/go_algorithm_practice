package main

func maximalNetworkRank(n int, roads [][]int) int {
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}
	cnt := make([]int, n)
	//构建图
	for _, road := range roads {
		cnt[road[0]]++
		cnt[road[1]]++
		graph[road[0]][road[1]] = true
		graph[road[1]][road[0]] = true
	}
	// 遍历
	max := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			total := cnt[i] + cnt[j]
			// 去重
			if graph[i][j] {
				total--
			}
			// 更新
			if total > max {
				max = total
			}
		}
	}
	return max
}
