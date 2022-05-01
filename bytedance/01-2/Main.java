import java.util.*;

public class Main {
    public static void main(String[] args) {
        int n,k;
        Scanner in = new Scanner(System.in) ;
        String[] s = in.nextLine().split(" ");
        n = Integer.valueOf(s[0]);
        k = Integer.valueOf(s[1]);
        String s1 = in.nextLine();
        char[] chars = s1.toCharArray();
        int[] freq = new int[26];
        // 记录字符出现频次  a-存在位置0  b存在位置1....
        for(int i=0;i<n;i++) {
            freq[chars[i]-'a'] ++;
        }
        // 计算
        int score = 0;
        // 依次遍历每个字符 最多26个
        for(int i=0;i<26;i++) {
            // 每个字符需要的敲击次数是i+1
            // 因为i++，所以不用考虑去重的问题，每个字符只会选择一次
            if(k>=i+1 && freq[i]>0) {
                score ++;
                k -= i+1;
            }
        }
        System.out.println(score);
        in.close();
    }
}