package lt517
/*
贪心法

517. 超级洗衣机
假设有 n 台超级洗衣机放在同一排上。开始的时候，每台洗衣机内可能有一定量的衣服，也可能是空的。

在每一步操作中，你可以选择任意 m (1 <= m <= n) 台洗衣机，与此同时将每台洗衣机的一件衣服送到相邻的一台洗衣机。

给定一个整数数组 machines 代表从左至右每台洗衣机中的衣物数量，请给出能让所有洗衣机中剩下的衣物的数量相等的 最少的操作步数 。如果不能使每台洗衣机中衣物的数量相等，则返回 -1 。
*/
func findMinMoves(machines []int) (ans int) {
    tot := 0
    for _, v := range machines {
        tot += v
    }
    n := len(machines)
    if tot%n > 0 {
        return -1
    }
    avg := tot / n
    sum := 0
    for _, num := range machines {
        num -= avg
        sum += num
        ans = max(ans, max(abs(sum), num))
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

