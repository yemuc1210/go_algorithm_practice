package lt488

import "fmt"

/*
488. 祖玛游戏
你正在参与祖玛游戏的一个变种。

在这个祖玛游戏变体中，桌面上有 一排 彩球，每个球的颜色可能是：红色 'R'、黄色 'Y'、蓝色 'B'、绿色 'G' 或白色 'W' 。你的手中也有一些彩球。

你的目标是 清空 桌面上所有的球。每一回合：

从你手上的彩球中选出 任意一颗 ，然后将其插入桌面上那一排球中：两球之间或这一排球的任一端。
接着，如果有出现 三个或者三个以上 且 颜色相同 的球相连的话，就把它们移除掉。
如果这种移除操作同样导致出现三个或者三个以上且颜色相同的球相连，则可以继续移除这些球，直到不再满足移除条件。
如果桌面上所有球都被移除，则认为你赢得本场游戏。
重复这个过程，直到你赢了游戏或者手中没有更多的球。
给你一个字符串 board ，表示桌面上最开始的那排球。另给你一个字符串 hand ，表示手里的彩球。请你按上述操作步骤移除掉桌上所有球，计算并返回所需的 最少 球数。如果不能移除桌上所有的球，返回 -1 。
*/
func findMinStep(board string, hand string) int {
	// BFS 暴力模拟 + 去重剪枝
	q := make([][]string, 0)
	q = append(q, []string{board, hand})
	depth := 0
	m := map[string]bool{} //用于去重剪枝
	for len(q) > 0 {
		depth++
		k := len(q)
		for k > 0 {
			k--
			cur := q[0]
			q = q[1:]
			// 把 h 中的每个元素暴力插入 b 中的每个位置，然后递归碰撞
			b, h := cur[0], cur[1]
			for i := 0; i < len(b); i++ {
				for j := 0; j < len(h); j++ {
					b2 := del3(b[0:i] + string(h[j]) + b[i:])
					h2 := h[0:j] + h[j+1:]
					if b2 == "" { // 递归碰撞后所有小球消失
						return depth
					}
					// 去重剪枝
					key := fmt.Sprintf("%v_%v", b2, h2)
					if m[key] {
						continue
					}
					m[key] = true
					q = append(q, []string{b2, h2})
				}
			}

		}
	}
	return -1
}

// 递归碰撞删除所有长度3及以上的子串
func del3(s string) string {
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			if count >= 3 {
				return del3(s[0:i-count] + s[i:])
			}
			count = 1
		}
	}
	if count >= 3 {
		return del3(s[0 : len(s)-count])
	}
	return s
}

