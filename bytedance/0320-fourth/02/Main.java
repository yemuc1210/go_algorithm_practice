import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        int M,N;
        Scanner in = new Scanner(System.in);
        String[] line = in.nextLine().split(" ");
        M = Integer.parseInt(line[0]);
        N = Integer.parseInt(line[1]);
        // N行  N个补给站
        int[] A = new int[N];
        int[] B = new int[N];
        for(int i=0;i<N;i++) {
            int A1,B1;
            line = in.nextLine().split(" ");
            A1 = Integer.parseInt(line[0]); //第A天经过 
            B1 = Integer.parseInt(line[1]);  // 每份价格B元
            // 求最少花费金额，每天消耗一份补给
            A[i] = A1;
            B[i] = B1;
        }
        // dp[i] : 第i天最少花费的金额
        // dp[0]=B[0]
        int[] day = new int[M];
        for(int i=0;i<A.length;i++){
            day[A[i]] = B[i];
        }
        int[] dp = new int[M];
        dp[0] = B[0];
        // 一共M天
        int lastPrice = B[0];
        for(int i=1;i<M;i++){
            if(day[i]!=0){// 第i天可以考虑更换补给
                dp[i] = dp[i-1] + Math.min(day[i], lastPrice);
                lastPrice = Math.min(day[i],day[i-1]);
            }else{
                // 第i天没有补给，当天的成本来源于上一次补给的价格
                dp[i] = dp[i-1] + lastPrice;
            }
        }
        System.out.println(dp[M-1]);
        in.close();
    }
}