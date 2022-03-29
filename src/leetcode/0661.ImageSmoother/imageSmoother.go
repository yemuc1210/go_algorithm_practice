package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	// 3*3
	m, n := len(img), len(img[0])

	// 求每个元素四周的和，然后再整体做除法
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 计算当前值周围的和
			cnt := 0
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if x >= 0 && x < n && y >= 0 && y < m {
					cnt++
					res[i][j] += img[x][y]
				}
			}
			res[i][j] /= cnt
		}
	}
	return res
}
func imageSmoother1(img [][]int) [][]int {
    m, n := len(img), len(img[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
        for j := range ans[i] {
            sum, num := 0, 0
            for _, row := range img[max(i-1, 0):min(i+2, m)] {
                for _, v := range row[max(j-1, 0):min(j+2, n)] {
                    sum += v
                    num++
                }
            }
            ans[i][j] = sum / num
        }
    }
    return ans
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
 
func main() {
	fmt.Println("vim-go")
	// img := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	img1 := [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10},{11,12,13},{14,15,16}}

	imageSmoother1(img1)
}
