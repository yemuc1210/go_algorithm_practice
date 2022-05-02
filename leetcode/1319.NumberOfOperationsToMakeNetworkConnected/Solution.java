class Solution {
    // 1319. 连通网络的操作次数  Java
    // 并查集
    // 思路： 
    // 1、三个连通块要想联通至少得需要两条边（也就是两条线）
    // 那么不难看出最终结果就是连通块数量-1 
    // 2、注意一个前提也就是线要够（n个节点至少需要n-1条线）
    int[] parent;
    public int makeConnected(int n, int[][] connections) {
        if(connections.length<n-1){
            return -1;
        }
        parent = new int[n];
        for (int i = 0; i < n; i++) {
            parent[i] = i;   // 初始化并查集
        }
        // 合并
        for(int[] connection : connections) {
            union(connection[0],connection[1]);
        }
        int count = 0;
        for (int i = 0; i < n; i++) {
            if (parent[i] == i) {
                count++;
            }
        }
        return count - 1;
    }
    int find(int node){
        return parent[node]==node ?node:(parent[node]=find(parent[node]));
    }
    void union(int node1,int node2){
        int root1=find(node1);
        int root2=find(node2);
        if(root1==root2){
            return ;
        }
        parent[root1]=root2;
    }
}