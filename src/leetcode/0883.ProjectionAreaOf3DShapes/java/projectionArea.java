class Solution {
    public int projectionArea(int[][] grid) {
        // 三个投影面积之和
        int xy=0;  // x0y平面，只要grid>1即+1
        int xz = 0;  // xoz，
        int yz=0;  //yoz平面，每一列最大值
        int m=grid.length,n=grid[0].length;
        int[] col = new int[n];  // 每一列的最大值
        for(int i=0;i<m;i++) {
            int rowMax = 0;  // 每一行得最大值
            for(int j=0;j<n;j++){
                if(grid[i][j]>0) {
                    xy ++;
                }
                // 记录最大值
                rowMax = Math.max(rowMax, grid[i][j]);
                col[j] = Math.max(col[j], grid[i][j]);
            }
            // 每一行
            xz += rowMax;
        }
        // 每一列得最大值，只有在遍历结束才能确定
        for(int j=0;j<n;j++){
            yz += col[j];
        }
        return xy+xz+yz;
    }
}