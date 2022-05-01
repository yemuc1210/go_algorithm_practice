import java.util.LinkedList;
import java.util.Queue;

class Solution {
    // 1091. 二进制矩阵中的最短路径
    public int shortestPathBinaryMatrix(int[][] grid) {
        if(grid == null || grid.length==0 || grid[0].length==0){
            return -1;
        }
        if(grid[0][0]==1) {
            return -1; // 起点阻塞
        }
        int m=grid.length,n=grid[0].length;
        int[][] dirs = {
            {-1,0},{1,0},{0,-1},{0,1},
            {-1,-1},{-1,1},{1,1},{1,-1},
        };
        Queue<int[]> queue = new LinkedList<>();
        queue.add(new int[]{0,0});  // 起点
        grid[0][0] = 1;  // 标记
        int levels = 1;  // 层数
        while(!queue.isEmpty()) {
            int size = queue.size();
            while(size-- > 0 ) {
                int[] cur = queue.poll();
                int x=cur[0],y=cur[1];
                if(x==m-1 &&y==n-1) {// 等于终点
                    return levels;
                }
                for (int[] dir : dirs) {
                    int x1 = x+dir[0],y1=y+dir[1];
                    if(x1<0 || x1>=m || y1<0 ||y1>=n || grid[x1][y1]==1) {
                        continue;
                    }
                    queue.add(new int[]{x1,y1});
                    grid[x1][y1] = 1;
                }
            }
            levels ++;
        }
        return -1;
    }
}