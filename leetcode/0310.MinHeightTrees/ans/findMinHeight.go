package main



// 找到距离最远的距离，然后最小高度一定在其路径上
// 综合上述推理，设两个叶子节点的最长距离为maxdist，可以得到结论最小高度树的高度为 maxdist/2，
// 且最小高度树的根节点一定存在其最长路径上。
// 假设最长的路径的 m 个节点依次为p1→p2→⋯→pm，
// 最长路径的长度为 m−1，可以得到以下结论：
// 如果 m 为偶数，此时最小高度树的根节点为pm/2或者p_{m/2+1}，且此时最小的高度为m/2；
// 如果 m 为奇数，此时最小高度树的根节点为p_{(m+1/2},且此时最小的高度为(m-1/2}

// 因此我们只需要求出路径最长的两个叶子节点即可，并求出其路径的最中间的节点即为最小高度树的根节点。
// 可以利用以下算法找到图中距离最远的两个节点与它们之间的路径：

// 以任意节点 p 出现，利用广度优先搜索或者深度优先搜索找到以 p 为起点的最长路径的终点 x；

// 以节点 x 出发，找到以 x 为起点的最长路径的终点 y；

// x 到 y 之间的路径即为图中的最长路径，找到路径的中间节点即为根节点。


func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }

    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    parents := make([]int, n)
    bfs := func(start int) (x int) {
        vis := make([]bool, n)
        vis[start] = true
        q := []int{start}
        for len(q) > 0 {
            x, q = q[0], q[1:]
            for _, y := range g[x] {
                if !vis[y] {
                    vis[y] = true
                    parents[y] = x
                    q = append(q, y)
                }
            }
        }
        return
    }
    x := bfs(0) // 找到与节点 0 最远的节点 x
    y := bfs(x) // 找到与节点 x 最远的节点 y

    path := []int{}
    parents[x] = -1
    for y != -1 {
        path = append(path, y)
        y = parents[y]
    }
    m := len(path)
    if m%2 == 0 {
        return []int{path[m/2-1], path[m/2]}
    }
    return []int{path[m/2]}
}
 