package lt391


/*每日一题 困难
391. 完美矩形
给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi]
	表示一个坐标轴平行的矩形。
这个矩形的左下顶点是 (xi, yi) ，右上顶点是 (ai, bi) 。

如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。
*/
/*
精确覆盖的意思：不能存在相交区域 ，不能有空隔
如果是完美矩形 那么一定满足两点：
（1）最左下 最左上 最右下 最右上 的四个点只出现一次 其他点成对出现
（2）四个点围城的矩形面积 = 小矩形的面积之和
*/
func isRectangleCover(rectangles [][]int) bool {
	type point struct{ x, y int }
	area, minX, minY, maxX, maxY := 0, rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	cnt := map[point]int{}
	for _, rect := range rectangles {
		x, y, a, b := rect[0], rect[1], rect[2], rect[3]
		area += (a - x) * (b - y)

		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, a)
		maxY = max(maxY, b)

		cnt[point{x, y}]++
		cnt[point{x, b}]++
		cnt[point{a, y}]++
		cnt[point{a, b}]++
	}

	if area != (maxX-minX)*(maxY-minY) || cnt[point{minX, minY}] != 1 || cnt[point{minX, maxY}] != 1 || cnt[point{maxX, minY}] != 1 || cnt[point{maxX, maxY}] != 1 {
		return false
	}

	delete(cnt, point{minX, minY})
	delete(cnt, point{minX, maxY})
	delete(cnt, point{maxX, minY})
	delete(cnt, point{maxX, maxY})

	for _, c := range cnt {
		if c != 2 && c != 4 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

