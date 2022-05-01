package lt794

import "strings"

/*
794. 有效的井字游戏
给你一个字符串数组 board 表示井字游戏的棋盘。当且仅当在井字游戏过程中，棋盘有可能达到 board 所显示的状态时，才返回 true 。

井字游戏的棋盘是一个 3 x 3 数组，由字符 ' '，'X' 和 'O' 组成。字符 ' ' 代表一个空位。

以下是井字游戏的规则：

玩家轮流将字符放入空位（' '）中。
玩家 1 总是放字符 'X' ，而玩家 2 总是放字符 'O' 。
'X' 和 'O' 只允许放置在空位中，不允许对已放有字符的位置进行填充。
当有 3 个相同（且非空）的字符填充任何行、列或对角线时，游戏结束。
当所有位置非空时，也算为游戏结束。
如果游戏结束，玩家不允许再放置字符。
*/

/*
回溯+剪枝  状态
十五谜问题
判断board状态是否可以达到
*/

func win(board []string, p byte) bool {
    for i := 0; i < 3; i++ {
        if board[i][0] == p && board[i][1] == p && board[i][2] == p ||
            board[0][i] == p && board[1][i] == p && board[2][i] == p {
            return true
        }
    }
    return board[0][0] == p && board[1][1] == p && board[2][2] == p ||
        board[0][2] == p && board[1][1] == p && board[2][0] == p
}

func validTicTacToe(board []string) bool {
    oCount, xCount := 0, 0
    for _, row := range board {
        oCount += strings.Count(row, "O")
        xCount += strings.Count(row, "X")
    }
    return !(oCount != xCount && oCount != xCount-1 ||
        oCount != xCount && win(board, 'O') ||
        oCount != xCount-1 && win(board, 'X'))
}

