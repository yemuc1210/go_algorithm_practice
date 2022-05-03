class Solution {
    // 3. 无重复字符的最长子串
    // 子串：连续的
    // 滑动窗口
    public int lengthOfLongestSubstring(String s) {
        if(s.length()==0){
            return 0;
        }
        int[] freq = new int[256];
        int res=0,left=0,right=-1;
        int n = s.length();
        while(left<n){
            if(right+1 < n && freq[s.charAt(right+1)]==0) {
                // 该字符首次出现
                freq[s.charAt(right+1)]++;
                right++;
            }else {
                freq[s.charAt(left)]--;
                left++;
            }
            res = Math.max(res, right-left+1);
        }
        return  res;
    }
    public static void main(String[] args){
        Solution s = new Solution();
        String s1 = " ";
        int res = s.lengthOfLongestSubstring(s1);
        System.out.println(res);
    }
}