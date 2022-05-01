import java.util.*;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] value = new int[n];
        int[] num = new int[n];
        for (int i = 0; i < n; i++) {
            value[i] = sc.nextInt();
        }
        for (int i = 0; i < n; i++) {
            num[i] = sc.nextInt();
        }
        int max = 0;
        int[][] dp = new int[n][n];
        for (int i = 0; i < n; i++) {
            int r = i + 1;
            if (r >= n) {
                break;
            }
            dp[i][r] = value[i] * num[r] + value[r] * num[i] -
                    value[i] * num[i] - value[r] * num[r];

            max = Math.max(max, dp[i][r]);

        }

        for (int i = 2; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (j + i >= n) {
                    break;
                }
                int r = j + i;
                dp[j][r] = dp[j + 1][r - 1] + value[j] * num[r] + value[r] * num[j]
                        - value[j] * num[j] - value[r] * num[r];
                max = Math.max(max, dp[j][r]);
            }
        }
        int sumb = 0;
        for (int i = 0; i < n; i++) {
            sumb += value[i] * num[i];
        }
        System.out.println(sumb + max);
        sc.close();
    }
}
