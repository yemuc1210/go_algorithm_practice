class Solution {
    // 28. 实现 strStr() Java
    // 
    public int strStr(String haystack, String needle) {
        if(needle.length()==0){
            return 0;
        }
        int n1=haystack.length(),n2=needle.length();
        for(int i=0;i<n1-n2+1;i++){
            if(haystack.substring(i,i+n2).equals(needle.substring(0))){
                return i;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        String haystack = "hello";
        System.out.println(haystack.substring(2,2+2));
    }
}

