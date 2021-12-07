package lt1583
/*
每日一题   2021.8.14

给你一份 n 位朋友的亲近程度列表，其中 n 总是 偶数 。

对每位朋友 i，preferences[i] 包含一份 按亲近程度从高到低排列 的朋友列表。
换句话说，排在列表前面的朋友与 i 的亲近程度比排在列表后面的朋友更高。每个列表中的朋友均以 0 到 n-1 之间的整数表示。

所有的朋友被分成几对，配对情况以列表 pairs 给出，其中 pairs[i] = [xi, yi] 表示 xi 与 yi 配对，且 yi 与 xi 配对。

但是，这样的配对情况可能会是其中部分朋友感到不开心。在 x 与 y 配对且 u 与 v 配对的情况下，如果同时满足下述两个条件，x 就会不开心：

x 与 u 的亲近程度胜过 x 与 y，且
u 与 x 的亲近程度胜过 u 与 v
返回 不开心的朋友的数目 。

 

模拟
（1）处理preference数组
	创建二维矩阵order   order[i][j]表示j在i中的亲近程度，下标序列（排序）
(2) 朋友match   match[x]=y && match[y]=x
(3) 统计
	找到与x配对的y
	找到y在x中的亲近程度  index
	x的朋友列表中的下标从0到index-1的朋友都是可能的u，遍历每个可能的u，找到与朋友u配对的v
	如果order[u][x] < order[u][v]  则x不开心，因为x亲近程度更高，但是没有却没有配对
*/
func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	order := make([][]int, n)
	var ans int
    for i, preference := range preferences {
        order[i] = make([]int, n)
        for j, p := range preference {
            order[i][p] = j
        }
    }
    match := make([]int, n)
    for _, p := range pairs {
        match[p[0]] = p[1]
        match[p[1]] = p[0]
    }

    for x, y := range match {
        index := order[x][y]
        for _, u := range preferences[x][:index] {
            v := match[u]
            if order[u][x] < order[u][v] {
                ans++
                break
            }
        }
    }
    return ans

}