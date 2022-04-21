package main

import "fmt"

// books[i] = [thickness,height]  厚度和高度
// 总宽度shelfWidth
// 可新建新层的书架
// 摆放顺序是固定的
// 每一层的高度=摆放的书的最大高度
// 求最小高度
func minHeightShelves(books [][]int, shelfWidth int) int {
	// 最小？  dfs   
	// dp[i] 第i本书摆放后，可以获得的最小高度
	dp := make([]int, len(books)+1)
	n:=len(books)
	dp[0] = 0
	for i:=1;i<=n;i++{
		dp[i] = 1000*1000  // 极大值
	}
	for i:=0;i<=n;i++{
		// 遍历每一本书，判断将这本书之前的书向后调整，是否可以减少之前的书架高度
		// dp[i] = min(dp[i],  dp[j-1]+h   )   j最后一层所能容下的书籍索引 h最后一层最大高度
		tmpWidth,j,h := 0,i,0
		for j>0 {
			tmpWidth += books[j-1][0]
			if tmpWidth>shelfWidth {
				break
			}
			h = max(h, books[j-1][1])
			dp[i] = min(dp[i], dp[j-1]+h)
			j--
		}
	}
	return dp[n]
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}

func main(){
	books := [][]int{
		{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2},
	}
	shelfWidth := 4
	res := minHeightShelves(books,shelfWidth)
	fmt.Println(res)
}