package lt986

/*
给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。

返回这 两个区间列表的交集 。

形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。

两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。


双指针题

比较区间首尾，有交集的情况，则把交集加入到结果

然后判断first还是second该到下一个区间：区间结束位置，小的移动

*/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	lenFirst,lenSecond := len(firstList),len(secondList)

	if lenFirst * lenSecond == 0{
		//有空的
		return [][]int{}
	}
	res := make([][]int,0)

	i,j := 0,0

	for i<lenFirst && j<lenSecond {
		startFirst,endFirst := firstList[i][0], firstList[i][1]
		startSecond,endSecond := secondList[j][0],secondList[j][1]

		//判断是否有交集
		if endFirst >= startSecond && endSecond >= startFirst {
			res = append(res, []int{max(startFirst,startSecond),   min(endFirst,endSecond)})
		}
		//移动
		if endFirst > endSecond{
			j++
		}else{
			i++
		}
	}


	return res
}

func max(a,b int) int{
	if a > b {
        return a
    } else {
        return b
    }
}

func min(a,b int)int{
	if a < b {
        return a
    } else {
        return b
    }
}