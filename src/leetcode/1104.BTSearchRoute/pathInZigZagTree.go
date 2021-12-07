package lt1104

import "math/bits"

/*
在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。
如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；
而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。


层次遍历 偶数层逆着来
*/
func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	return append(pathInZigZagTree(1<<(bits.Len(uint(label))-2)*3-1-label/2), label)
}

func getReverse(label, row int) int {
	return 1<<(row-1) + 1<<row - 1 - label
}

func pathInZigZagTree1(label int) (path []int) {
	row, rowStart := 1, 1
	for rowStart*2 <= label {
		row++
		rowStart *= 2
	}
	if row%2 == 0 {
		label = getReverse(label, row)
	}
	for row > 0 {
		if row%2 == 0 {
			path = append(path, getReverse(label, row))
		} else {
			path = append(path, label)
		}
		row--
		label >>= 1
	}
	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}
	return
}
