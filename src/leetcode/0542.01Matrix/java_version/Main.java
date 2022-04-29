import java.util.LinkedList;
import java.util.Queue;

class Solution {
    // 542. 01 矩阵
    // 给定一个由 0 和 1 组成的矩阵 mat ，
    // 请输出一个大小相同的矩阵，
    // 其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
    // 两个相邻元素间的距离为 1 。
    public int[][] updateMatrix(int[][] mat) {
        int m=mat.length,n=mat[0].length;
        Queue<int[]> queue = new LinkedList<>();
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(mat[i][j]==0){
                    queue.add(new int[]{i,j});
                }else{
                    mat[i][j] = -1;
                }
            }
        }
        int[][] dirs = {
            {0,1},{0,-1},{1,0},{-1,0},
        };
        while(!queue.isEmpty()){
            int[] point = queue.poll();
            for(int[] dir : dirs) {
                int x=point[0]+dir[0],y=point[1]+dir[1];
                if(x>=0&&x<m&&y>=0&&y<n&&mat[x][y]==-1){
                    mat[x][y] = mat[point[0]][point[1]] + 1;  // 一层一层的
                    queue.add(new int[]{x,y});
                }
            }
        }
        return mat;
    }
}