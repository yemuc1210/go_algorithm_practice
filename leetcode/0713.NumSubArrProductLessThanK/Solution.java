class Solution {
    // 713. 乘积小于 K 的子数组 Java
    // 连续子数组的数目    滑动窗口
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        int left=0,right=0;
        int multi = 1;
        int res = 0;
        while(right < nums.length) {
            int num = nums[right++];
            if(num >= k) {
                left = right;
                multi = 1;
                continue;
            }
            multi = multi * num;
            while(multi>=k){
                multi = multi/nums[left];
                left++;
            }
            if(multi<k){
                res += right - left;
            }
        }
        return res;

    }
}