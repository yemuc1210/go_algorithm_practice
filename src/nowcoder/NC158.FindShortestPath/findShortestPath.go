package NC158

import "container/list"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 单源最短路
 * 返回1-n的最短路径值  n个点 m边 有向图
 * @param n int 顶点数
 * @param m int 边数
 * @param graph int二维数组 一维3个数据，表示顶点到另外一个顶点的边长度是多少​
 * @return int
 */
type node struct {
    pos int
    state int
    step int
}

func shortestPathLength(n,m int, graph [][]int) int {
    visited := make([]map[int]struct{}, len(graph)+1)
    queue := list.New()
    success := 0
    for i := 0; i < n; i++ {
        state := 1 << uint(i)
        success |= state
        visited[i] = make(map[int]struct{})
        visited[i][state] = struct{}{}
        queue.PushBack(node{i, state, 0})
    }
    
    for queue.Len() > 0 {
        e := queue.Front()
        now := e.Value.(node)
        if now.state == success {
            return now.step
        }
        queue.Remove(e)
        for _, g := range graph[now.pos] {
            t := now.state | 1 << uint(g)
            if _, ok := visited[g][t]; ok {
                continue
            }
            visited[g][t] = struct{}{}
            queue.PushBack(node{g, t, now.step+1})
        }
    }
    return -1
}



// public class Solution {
//     /**
//      * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
//      *
//      * 
//      * @param n int 顶点数
//      * @param m int 边数
//      * @param graph int二维数组 一维3个数据，表示顶点到另外一个顶点的边长度是多少​
//      * @return int
//      */
//         public int findShortestPath (int n, int m, int[][] graph) {
//         int[][] dp = new int[n][n];
//         for(int[] d : dp){
//             Arrays.fill(d , Integer.MAX_VALUE);
//         }
//         for(int[] edge : graph){
//             dp[edge[0] - 1][edge[1] - 1] = Math.min(edge[2] , dp[edge[0] - 1][edge[1] - 1]);
//         }
//         //优先队列按照花费进行从小到大的排序
//         PriorityQueue<CityInfo> queue = new PriorityQueue<CityInfo>((o1,o2) ->{return o1.cost - o2.cost;});
//         queue.offer(new CityInfo(0, 0));
//         while(!queue.isEmpty()){
//             CityInfo info = queue.poll();
//             //如果到达了终点，而queue.poll的就是花费最低的，那么就给其返回
//             if(info.dst == n - 1){
//                 return info.cost;
//             }
//             //到达目前的该点到其他点的路线是否存在，如果存在，就加入到队列中
//             for(int i = 0 ;i < n ;i++){
//                 if(dp[info.dst][i] != Integer.MAX_VALUE){
//                     queue.offer(new CityInfo(i , dp[info.dst][i] + info.cost));
//                 }
//             }
//         }
//         return -1;
//     }
//     class CityInfo{
//         int dst;
//         int cost;
//         public CityInfo(int dst ,int cost){
//             this.dst = dst;
//             this.cost = cost;
//         }
//     }
// }

