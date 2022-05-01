package offerS110

/* 剑指 Offer II 110. 所有路径    lt797
算法计划   dfs  bfs 

给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）

二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。

译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。

 

*/

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
    stk := []int{0}
    var dfs func(int)
    dfs = func(x int) {
        if x == len(graph)-1 {
            ans = append(ans, append([]int(nil), stk...))
            return
        }
        for _, y := range graph[x] {
            stk = append(stk, y)
            dfs(y)
            stk = stk[:len(stk)-1]
        }
    }
    dfs(0)
    return
}