class Solution {
    // 279. 完全平方数
    // 返回和为n的完全平方数的最少数量
    public int numSquares(int n) {
        // dp[i] : 最少需要i个数的平方和
        // dp[i] = 1 + min{dp[i-j^2]}   j=[1,i**0.5]
        int[] dp = new int[n+1];
        dp[0]=0;
        for(int i=1;i<=n;i++){
            int minN = Integer.MAX_VALUE;
            for(int j=1;j*j<=i;j++){
                minN = Math.min(minN, dp[i-j*j]);
            }
            dp[i] = minN + 1;
        }
        return dp[n];
    }
}