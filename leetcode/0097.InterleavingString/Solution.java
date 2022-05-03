class Solution {
    // 97. 交错字符串
    // 子串相对位置不变
    public boolean isInterleave(String s1, String s2, String s3) {
        int m=s1.length(),n=s2.length(),k=s3.length();
        if(m+n !=k){
            return false;
        }
        boolean[][] dp = new boolean[m+1][n+1];
        dp[0][0] = true;
        for(int i=1;i<m+1;i++){
            dp[i][0] = dp[i-1][0] && s1.charAt(i-1)==s3.charAt(i-1);//s2空
        }
        
        for(int i=1;i<n+1;i++){
            dp[0][i] = dp[0][i-1] && s2.charAt(i-1)==s3.charAt(i-1);//s1
        }
        for(int i=1;i<m+1;i++){
            for(int j=1;j<n+1;j++){
                //要么s1[i-1]==s3[i+j-1]且之前的都匹配上了要么s2[j-1]==s3[i+j-1]之前的也都匹配上了
                dp[i][j]=(dp[i-1][j] &&s1.charAt(i-1)==s3.charAt(i+j-1) || (dp[i][j-1] && s2.charAt(j-1)==s3.charAt(i+j-1)));

            }
        }
        return dp[m][n];
    }
}

