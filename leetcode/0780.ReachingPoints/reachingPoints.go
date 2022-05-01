package main


/*
780. 到达终点
给定四个整数 sx , sy ，tx 和 ty，
如果通过一系列的转换可以从起点 (sx, sy) 到达终点 (tx, ty)，则返回 true，否则返回 false。

从点 (x, y) 可以转换到 (x, x+y)  或者 (x+y, y)。
*/
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx>=sx && ty>=sy {
		if tx==sx && ty==sy {
			break
		}
		if tx > ty {
			tx -= max(1,(tx-sx)/ty) * ty 
		}else{
			ty -= max(1, (ty-sy)/tx) * tx
		}
	}
	return tx==sx && ty==sy
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
// 思路分析：这道题确实非常妙！大部分的人一般都会惯性思维死劲想着怎么从(sx, sy)推到(tx, ty)，
// 蛋式由于可以变换的情况非常多，特别是当起点与终点的差距比较大的时候。如果我们逆向思考呢，
// 从(tx, ty)推到(sx, sy)，则时只能有一种操作，就是将tx、ty中较大值减去较小值（因为顺推的时候是(x, y)
//  可以转换到 (x, x+y) 或者 (x+y, y)，则逆推的时候只能将较大者减去较小者）

// if : tx > ty then : tx = tx % ty
// if : ty > tx then : ty = ty % tx