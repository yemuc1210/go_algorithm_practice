class Solution {
    // 483. 最小好进制
    // n 的k进制数所有数位都是1，则k为一个好进制
    // *设n的最小好进制为k，转化为k进制后的字符串为"111…11"，
    // 进而n可以转化为 n==k^(m-1)+k^(m-2)…+k1+k0
    // (其中m待定，是一个常数，为转换为k进制的后的字符串的长度）
    // 这就是一个等比数列，根据等比数列的求和公式，
    // 可知 n == (k^m - 1) / (k - 1)，由于题目中要求进制k尽量的小，
    // 那么根据求和公式知需要m尽可能的大（转换k进制后1的个数尽量多）。
    // 因为（k>=2），所以m的最大值为 m = log2(num) + 1，
    // 并且由于n的取值范围是 [3, 10^18]，即m的最小值是2。

    public String smallestGoodBase(String n) {
        // 将字符串解析成long型数据
        long num = Long.parseLong(n);
        // 对进行m的大小进行穷举（m含义是转换为k进制后1的个数）
        for (int m = (int) (Math.log(num + 1) / Math.log(2)); m >= 2; m--) {
            // 用二分法搜索对应的k,(k的含义是k进制)
            long left = 2, right = (long) Math.pow(num, 1.0 / (m - 1)) + 1;
            while (left < right) {
                long mid = left + (right - left) / 2, sum = 0;

                // 等比数列求和
                for (int j = 0; j < m; j++)
                    sum = sum * mid + 1;

                if (sum == num)
                    return String.valueOf(mid);
                else if (sum < num)
                    left = mid + 1;
                else
                    right = mid;
            }
        }
        return String.valueOf(num - 1);
    }
}