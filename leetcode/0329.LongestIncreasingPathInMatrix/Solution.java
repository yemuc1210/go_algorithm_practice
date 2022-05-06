class Solution {
    // 记忆搜索
    int[][] dirs = {{-1,0},{1,0},{0,-1},{0,1}};
    public int longestIncreasingPath(int[][] matrix) {
        int m=matrix.length,n=matrix[0].length;
        int[][] cache = new int[m][n];
        int res = 0;
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                // dfs
                dfs(i,j,cache,matrix,Integer.MIN_VALUE);
                res = Math.max(res,cache[i][j]);
            }
        }
        return res;
    }

    int dfs(int i,int j,int[][] cache,int[][] board, int lastElem) {
        // 越界判断
        
        // 递增判断
        if(board[i][j] <= lastElem) {
            return 0;
        }
        // 查缓存
        if(cache[i][j] > 0 ) {
            return cache[i][j];
        }

        // 访问
        int count = 1;
        for(int[] dir:dirs) {
            int x = i+dir[0],y=j+dir[1];
            // 越界判断（二选一
            if(x>=0&&x<board.length && y>=0 && y<board[0].length){
                // 可以递归
                count = Math.max(count, dfs(x, y, cache, board, board[i][j])+1);
            }
        }
        cache[i][j] = count;   // (i,j)开始，可以得到的最长递增路径
        return count;
    }
}