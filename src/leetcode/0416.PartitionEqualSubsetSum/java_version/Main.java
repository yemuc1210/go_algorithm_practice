class Solution {
    // 416. 分割等和子集
    // 数组只包含正整数，分割为两个子集，每个子集和相等
    public boolean canPartition(int[] nums) {
        int sum = 0;
        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];
        }
        if(sum%2==1) {
            return false;

        }
        int target = sum/2;
        // 背包问题
        // dp[j]:容量为j的背包，价值最大
        int[] dp = new int[target+1];  
        dp[0] = 0;
        for (int i = 0; i < nums.length; i++) {
            for (int j = target; j >=nums[i]; j--) {
                // 判断放入和不放入的价值
                dp[j] = Math.max(dp[j], dp[j-nums[i]]+nums[i]);
            }
        }
        return dp[target] == target;
    }
}