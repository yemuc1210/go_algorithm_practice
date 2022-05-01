class Solution {
    // 322 零钱兑换
    // 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
    // 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
    // 你可以认为每种硬币的数量是无限的。
    public int coinChange(int[] coins, int amount) {
        if(amount<1 && coins.length<1) {
            return -1;
        }
        // 背包问题
        // dp[i]: 金额i需要最少的硬币数
        int[] dp = new int[amount+1];

        for (int i = 0; i < dp.length; i++) {
            dp[i] = amount+1;  // 极大值
        }
        dp[0] = 0;
        for (int i = 1; i <=amount; i++) {
            for(int j=0;j<coins.length;j++){
                if(coins[j]<=i) {
                    dp[i] = Math.min(dp[i], dp[i-coins[j]]+1);
                }
                
            }
        }
        if(dp[amount]>amount){
            return -1;
        }
        return dp[amount];
    }
}