import java.util.LinkedList;
import java.util.Queue;

public class Solution {
    // 从（sr,sc) 开始上色  row col
    // bfs
    public int[][] floodFill(int[][] image, int sr, int sc, int newColor) {
        int[][] dir = {{-1,0},{1,0},{0,-1},{0,1}};
        int srcColor = image[sr][sc];
        if(srcColor == newColor) {
            return image;
        }
        int m= image.length,n = image[0].length;
        Queue<int[]> queue = new LinkedList<int[]>();
        queue.offer(new int[]{sr,sc});
        image[sr][sc] = newColor;
        while(!queue.isEmpty()) {
            int[] cell = queue.poll();
            int x=cell[0],y=cell[1];
            for (int i=0;i<4;i++){
                int mx =  x+dir[i][0],my=y+dir[i][1];
                if(mx>=0 && mx<m && my>=0 && my<n) {
                    queue.offer(new int[]{mx,my});
                    image[mx][my] = newColor;
                }
            }
        }
        return image;
    }
}