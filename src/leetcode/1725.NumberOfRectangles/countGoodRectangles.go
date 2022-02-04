package lt1725

    /*
        简单模拟
        遍历所有矩形，获取每个矩形所能切割的最大正方形边长
        此时存在三种情况： 1. 边长大于已经存在的最大边长 ， 则替换最大边长
        2. 边长等于已经存在的最大边长，答案+1
        3. 边长小于已经存在的最大边长，跳过

        执行用时：1 ms, 在所有 Java 提交中击败了100.00%的用户
    */
	func countGoodRectangles(rectangles [][]int) (ans int) {
		maxLen := 0
		for _, rect := range rectangles {
			k := min(rect[0], rect[1])
			if k == maxLen {
				ans++
			} else if k > maxLen {
				maxLen, ans = k, 1
			}
		}
		return
	}
	
	func min(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	
