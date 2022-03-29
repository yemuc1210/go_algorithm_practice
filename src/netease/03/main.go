package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	// 节点权值是1-n的排列,意味着不可以重复
	// 除了根节点,每个节点权值和它父亲的权值乘积为偶数
	// 回溯咯
	//       2
	//  4          3
	// 1
	if n <= 2 {
		fmt.Println("12")
	} else if n == 3 {
		fmt.Println("213")
	} else {
		leftNum := make([]int, n)
		for i := 0; i < n; i++ {
			leftNum[i] = i + 1
		}
		res := [][]int{}
		solve(leftNum, &res, []int{}, n,0)
		fmt.Println(res)
	}

}
func solve(leftNum []int, res *[][]int, path []int, n int, curIdx int) {
	// return 条件设计
	if curIdx == n { // 剩余可用数字
		// 得到一组结果
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		// copy(*res, path)
		return
	}
	// 当层处理
	for i := 0; i < len(leftNum); i++ {
		// 当前数字对应的父亲节点
		// 完全二叉树中 左节点i的父亲下标为  i/2    右节点i的父亲结点为(i-1)/2
		// 先判断当前数字应该为左还是右
		var parentIdx int
		if curIdx == 0 {
			// 依次尝试每个数字
			path = append(path, leftNum[i])
			solve(append(leftNum[:i], leftNum[i+1:]...), res, path, n,curIdx+1)
			path = path[:len(path)-1]
		} else {
			if len(path)%2 != 0 {
				// 左结点
				parentIdx = len(path) / 2
			} else {
				// 右节点
				parentIdx = (len(path) - 1) / 2

			}
			if leftNum[i]*path[parentIdx]%2 == 0 {
				// 如果乘积为偶数,ok
				// 进入下一层
				path = append(path, leftNum[i])
				// 移除当前元素i
				// leftNum = append(leftNum[:i], leftNum[i+1:]...)
				solve(append(leftNum[:i], leftNum[i+1:]...), res, path, n,curIdx+1)
				// 回溯
				path = path[:len(path)-1]
			}
		}
	}
}
