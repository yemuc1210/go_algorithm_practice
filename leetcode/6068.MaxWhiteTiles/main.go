package main

import (
	"fmt"
	"sort"
)

// 毯子覆盖的最多白色砖数
// tiles[i] = [li,ri] 表示li,ri之间瓷砖白色
// 毛毯长度
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	// 区块合并
	// 根据左区间排序
	sort.Slice(tiles,func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	// 合并
	fmt.Println(tiles)
	sk := [][]int{}
	sk = append(sk, []int{tiles[0][0],tiles[0][1],tiles[0][1]-tiles[0][0]+1})
	n := len(tiles)
	// 记录每个区间的情况
	// 计算当前白色瓷砖数量
	var curNum int = tiles[0][1]-tiles[0][0]+1
	nMap := make(map[int]int)
	nMap[tiles[0][0]] = 0
	nMap[tiles[0][1]] = curNum
	for i:=1;i<n;i++ {
		// 判断是否可以合并
		top := sk[len(sk)-1]
		l1,r1 := top[0],top[1]	
		l2,r2 := tiles[i][0],tiles[i][1]
		// fmt.Println(l1,r1,l2,r2)
		// 合并
		// l2 肯定大于等于l1
		if r2 <= r1 {
			// 第一个区间大
			// sk = append(sk, []int{l1,r1,curNum})
			nMap[l2] = nMap[l1]
			nMap[r2] = nMap[r1]
			continue
		}else if l2-1 == r1 {
			// 交叉  l2<r1 
			nMap[l2] = nMap[r1]+1
			nMap[r2] = curNum + r2-l2+1
			sk = sk[:len(sk)-1]
			sk = append(sk, []int{l1,r2,curNum+r2-l2+1})
			curNum += r2-l2+1
		}else{
			// 两个独立
			// sk =append(sk, []int{l1,r1,curNum})
			nMap[l2] = nMap[r1] +1
			nMap[r2] = curNum + r2-l2+1
			sk = append(sk, []int{l2,r2,curNum + r2-l2+1})
			curNum += r2-l2+1
		}
	}
	// 判断最长的一块
	maxLen := 0 
	for i:=0;i<len(sk);i++ {
		l,r,length := sk[i][0],sk[i][1], sk[i][1] - sk[i][0] + 1
		cnt := sk[i][2]
		if nMap[r+carpetLen-1] >0 {
			if nMap[r+carpetLen-1] > maxLen {
				maxLen = nMap[r+carpetLen-1]
			}
		}
	}
	if maxLen < carpetLen {
		return maxLen
	}
	return carpetLen
}
func solve(tiles [][]int, carpetLen int)  int{
	sort.Slice(tiles,func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	bitMap := make([]bool, tiles[len(tiles)-1][1]+1)
	for _,tile := range tiles {
		l,r := tile[0],tile[1]
		for j:=l;j<=r;j++{
			bitMap[j] = true
		}
	}
	var res int
	for i:=0;i<len(bitMap);i++{
		var cnt int
		for j:=i;j<i+carpetLen;j++{
			if bitMap[j] {
				cnt ++
			}
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}
func main() {
	// tiles := [][]int{
	// 	{1,5},{10,11},{12,18},{20,25},{30,32},
	// }
	// tiles := [][]int{
	// 	{10,11},{1,1},
	// }
	tiles := [][]int{{8051,8057},{8074,8089},{7994,7995},{7969,7987},{8013,8020},{8123,8139},{7930,7950},{8096,8104},{7917,7925},{8027,8035},{8003,8011}}
	// 9854
	carpetLen := 9854
	res := maximumWhiteTiles(tiles,carpetLen)
	fmt.Println(res)

	res1 :=solve(tiles,carpetLen)
	fmt.Println(res1)
}
// {{8051,8057},{8074,8089},{7994,7995},{7969,7987},{8013,8020},{8123,8139},{7930,7950},{8096,8104},{7917,7925},{8027,8035},{8003,8011}}
// 9854