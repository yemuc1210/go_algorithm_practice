import java.util.*;

public class Bonus {
    public int getMost(int[][] board) {
        int n=board.length; // 行数
        int m=board[0].length; // 列数
        int[] dp=new int[m+1]; //加一列虚拟列
        for (int i=0;i<n;i++){
            for (int j=1;j<m+1;j++){ // 列从第二列开始
                if(board[i][j-1]>100&&board[i][j-1]<1000){
                    dp[j]=Math.max(dp[j],dp[j-1])+board[i][j-1];
                }
                else{
                    dp[j]=0;
                }
            }
        }
        return dp[m];
    }
}