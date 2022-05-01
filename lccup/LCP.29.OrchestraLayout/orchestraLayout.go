package lcp29
/*
某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。
乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。

为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。例如当 num = 5 时，站位如图所示

请返回位于场地坐标 [Xpos,Ypos] 的成员所持乐器编号。


模拟，初始化一个二维矩阵，然后在相应位置写入对应编号
剑指offer29

*/
func min(args... int)int{
    var m = args[0]
    for i := 1;i<len(args);i++{
        if args[i] < m{
            m = args[i]
        }
    }
    return m
}

func orchestraLayout(num int, xPos int, yPos int) int {
    // 遍历计算移动到该position需要的距离
    // var s int
    // for xPos != 0 && yPos != 0 && xPos != num - 1 && yPos != num - 1{
    //     s += 4 * (num - 1)
    //     xPos -= 1
    //     yPos -= 1
    //     num -= 2
    // }
    // 优化上面代码 转为O(1)
    min_distance := min(xPos, yPos, num - xPos - 1, num - yPos - 1)
    x1 := 4 * (num - 1)
    xn := 4 * (num - 1 - 2 * (min_distance - 1))
    s := (x1 + xn) * min_distance / 2
    num -= 2 * min_distance
    xPos -= min_distance
    yPos -= min_distance
    switch {
        case xPos == 0:
            s += yPos + 1
        case yPos == num -1:
            s += (num - 1) + xPos + 1
        case xPos == num - 1:
            s += (num - 1) * 3 - yPos + 1
        case yPos == 0:
            s += (num - 1) * 4 - xPos + 1
    }
    return (s - 1) % 9 + 1
}
