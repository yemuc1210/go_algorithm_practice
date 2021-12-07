/*
请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

    数字 1-9 在每一行只能出现一次。
    数字 1-9 在每一列只能出现一次。
    数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

数独部分空格内已填入了数字，空白格用 '.' 表示。

注意：

    一个有效的数独（部分已被填充）不一定是可解的。
    只需要根据以上规则，验证已经填入的数字是否有效即可。
*/

package leetcode

//添加缓存   O(n^2)
func isValidSudoku(board [][]byte) bool {
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	// 遍历一次，添加缓存
	for r := 0; r < 9; r++ {    //r 行c 列
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				num := board[r][c] - '0' - byte(1)
				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}
	return true
}

func isValidSudoku1(board [][]byte) bool {
    var rows, columns [9][9]int
    var subboxes [3][3][9]int
    for i, row := range board {
        for j, c := range row {
            if c == '.' {
                continue
            }
            index := c - '1'
            rows[i][index]++
            columns[j][index]++
            subboxes[i/3][j/3][index]++
            if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
                return false
            }
        }
    }
    return true
}

