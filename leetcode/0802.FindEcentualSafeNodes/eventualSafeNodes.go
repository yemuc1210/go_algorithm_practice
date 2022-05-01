package lt802

/*
dfs+三色标记
有环则是不安全的
0 未访问
1 灰色 递归栈或某个环
2 黑色  安全
当我们首次访问一个节点时，将其标记为灰色，并继续搜索与其相连的节点。

如果在搜索过程中遇到了一个灰色节点，则说明找到了一个环，此时退出搜索，栈中的节点仍保持为灰色，这一做法可以将「找到了环」这一信息传递到栈中的所有节点上。

如果搜索过程中没有遇到灰色节点，则说明没有遇到环，那么递归返回前，我们将其标记由灰色改为黑色，即表示它是一个安全的节点。


*/
func eventualSafeNodes(graph [][]int) (ans []int) {
    n := len(graph)
    color := make([]int, n)
    var safe func(int) bool
    safe = func(x int) bool {
        if color[x] > 0 {
            return color[x] == 2
        }
        color[x] = 1
        for _, y := range graph[x] {
            if !safe(y) {
                return false
            }
        }
        color[x] = 2
        return true
    }
    for i := 0; i < n; i++ {
        if safe(i) {
            ans = append(ans, i)
        }
    }
    return
}

