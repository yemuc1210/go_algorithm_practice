// import java.util.*;
package Solution;

class Solution {
    private int MOD = (int) 1e9 + 7;
    // 30 以内的质数
    private int[] p = new int[]{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}; //  10 个

    public int numberOfGoodSubsets(int[] nums) {
        // 统计数字出现的次数
        int[] cnts = new int[31];
        for (int num : nums) cnts[num]++;

        //dp[5]  5 二进制位 0000000101, 表示子集所有元素乘积为 p[0] * p[2] = 2 * 5 的个数
        // 动态转移方式为在已有的情况下再增加其他质数  (p[0] * p[2]) * p[2] = (2 * 5) * 3
        // 注意这里 dp 用 int 会溢出
        long[] dp = new long[1 << p.length];
        dp[0] = 1;

        // 尝试将 num 加入子集中，并累计结果数
        for (int num = 2; num < 31; num++) {
            if (cnts[num] < 1) continue;

            boolean ok = true;
            // t 用来你分解质因数， mask 二进制位用来记录分解出现的质因数位置
            int t = num, mask = 0;
            for (int j = 0; j < p.length; j++) {
                int cnt = 0;
                while (t % p[j] == 0) {
                    mask |= (1 << j);
                    t /= p[j];
                    cnt++;
                }

                // 质数出现两次,则不符合题意要求,例如 4 = 2＊２
                if (cnt > 1) {
                    ok = false;
                    break;
                }
            }

            if (!ok) continue;
            for (int i = 0; i < dp.length; i++) {
                if ((i & mask) != 0) continue; // 冲突，同一个质数多次被选中，因此不能累积
                // dp[i] （i 的二进制表的）子集可以添加 num , 追加以后的结果集为  dp[i | mask]
                dp[i | mask] = (dp[i | mask] + dp[i] * cnts[num]) % MOD;
            }
        }

        long ans = 0;
        for (int i = 1; i < dp.length; i++) {
            ans = (ans + dp[i]) % MOD;
        }

        // 数字1，特殊处理，因为 1 可以添加到任何子集中
        for (int i = 0; i < cnts[1]; i++) ans = (2 * ans) % MOD;

        return (int)ans;
    }
}