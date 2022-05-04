package main
// 1466. 重新规划路线

func minReorder(n int, connections [][]int) int {
    arrived := make([]bool, n)
    ret := 0
    stack := [][]int{}
    // 统计直达
    for _, conn := range connections {
        // 需要翻转
        if conn[0] == 0 || arrived[conn[0]] {
            ret++
            arrived[conn[1]] = true
            continue
        }
        if conn[1] == 0 || arrived[conn[1]] {
            arrived[conn[0]] = true
            continue
        }
        stack = append(stack, conn)
    }
    // 统计间接到达
    for len(stack) > 0 {
        nStack := [][]int{}
        for _, conn := range stack {
            // 需要翻转
            if conn[0] == 0 || arrived[conn[0]] {
                ret++
                arrived[conn[1]] = true
                continue
            }
            if conn[1] == 0 || arrived[conn[1]] {
                arrived[conn[0]] = true
                continue
            }
            nStack = append(nStack, conn)
        }
        stack = nStack
    }
    
    return ret
}

