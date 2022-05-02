class Solution {
    // 547. 省份数量 Java
    // 省份：连通子图
    public int findCircleNum(int[][] isConnected) {
        int res=0;
        boolean[] visited = new boolean[isConnected.length];
        for(int i=0;i<isConnected.length;i++){
            if(!visited[i]){
                dfs(isConnected, visited, i);
                res++;
            }
        }
        return res;
    }
    void dfs(int[][] isConnected, boolean[] visited, int i) {
        // i开始搜索
        for(int j=0;j<isConnected.length;j++){
            if(isConnected[i][j]==1 && !visited[j]){
                visited[j] = true;
                dfs(isConnected, visited, j);
            }
        }
    }
}