import java.util.ArrayList;
import java.util.HashSet;

class Solution {
    // 1042. 不邻接植花
    // 图 花园之间存在路径
    // 所有花园最多有三条路径可以加入或离开
    // 为每个花园种一种花，要就邻接的种类不同       图染色
    // 返回任意可行方案
    // 花的种类  1 2 3 4
    // n 表示花园个数
    public int[] gardenNoAdj(int n, int[][] paths) {
        int[] ans = new int[n];
        // 图
        List<Integer>[] adjGardens = new List[n];
        for(int i=0;i<n;i++){
            adjGardens[i] = new ArrayList<Integer>();
        }
        for(int[] path:paths){
            int garden0=path[0]-1,garden1=path[1]-1;
            if(garden0<garden1){
                adjGardens[garden1].add(garden0);
            }else{
                adjGardens[garden0].add(garden1);
            }
        }
        for(int i=0;i<n;i++){
            List<Integer> adj = adjGardens[i];
            boolean[] used = new boolean[5];
            for(int garden : adj) {  // 遍历得到相邻花园，并得到其种类
                int adjType = ans[garden];
                used[adjType] = true;
            }
            int type=1;
            while(used[type]){
                type++; // 相邻花园的花种类 决定当前花园
            }
            ans[i] = type; // 当前花的种类
        }
        return ans;
    }
    int[] bfs(int n,int[][] paths,boolean[][] visited,int nodeIdx) {
        int[] flowerTpe = {1,2,3,4};
        int[] ans = new int[n];
        // 构建图
        List<Integer>[] adjGraph = new List[n];
        Set<Integer> used = new Set[n];
        for(int i=0;i<n;i++){
            adj[i] = new ArrayList<>();
            used[i] = new HashSet<>();
        }
        for(int[] edge: paths) {
            adjGraph[edge[0]-1].add(edge[1]-1);
            adjGraph[edge[1]-1].add(edge[0]-1);
        }
        // 遍历图
        for(int i=0;i<n;i++){
            for(int f:flowerTpe) {
                if(!used[i].contains(f)) {
                    // 确定颜色
                    ans[i] = f;
                    for(int v:adjGraph[i]) {  // 邻接节点
                        if(v>i){
                            used[v].add(f);  // 更新邻接节点的信息，即让邻接节点知晓颜色
                        }
                    }
                    break;   // --->for i
                }
            }
        }

        return ans;
    }
}