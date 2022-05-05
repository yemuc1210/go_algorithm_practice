import java.util.Arrays;

class Solution {
    // 787. K 站中转内最便宜的航班 Java
    // flights[i] = [from,to,price] 带权图
    // 路径：最多不超过k个中转站 保证src-dst价格最便宜

    public int findCheapestPrice(int n, int[][] flights, int src, int dst, int k) {
        // 从终点逆推
        int INF = 10000 * 101 + 1;
        int[][] dp = new int[k + 2][n];
        // dp[t][i] 表示通过t次航班到达城市i的最小花费
        for (int i = 0; i < k + 2; i++) {
            Arrays.fill(dp[i], INF);
        }
        dp[0][src] = 0; // t=0不搭乘航班
        for (int t = 1; t <= k + 1; t++) {
            for (int[] flight : flights) {
                int j = flight[0], i = flight[1], cost = flight[2];
                // 转移方程
                dp[t][i] = Math.min(dp[t][i], dp[t - 1][j] + cost); // 通过j中转一次
            }
        }
        int ans = INF;
        for (int t = 1; t <= k + 1; t++) {
            ans = Math.min(ans, dp[t][dst]);
        }
        return ans==INF?-1:ans;
    }
}