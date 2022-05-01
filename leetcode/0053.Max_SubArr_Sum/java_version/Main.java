public class Main {
    // 最大子数组和
    // 具有最大和的连续子数组---滑动窗口
    public int maxSubArray(int[] nums) {
        int maxSum = Integer.MIN_VALUE;
        int curSum = 0;// 当前窗口zhi
        for (int i : nums) {
            curSum += i;
            if(curSum > maxSum) {
                maxSum = curSum;
            }
            if( curSum <= 0) {
                curSum = 0;   // 对应left移动得步骤
            }
        }
        return maxSum;
    }
    public static void main(String[] args) {
        int[] nums = {-2,1,-3,4,-1,2,1,-5,4};
        Main m = new Main();
        int res = m.maxSubArray(nums);
        System.out.println(res);
    }
}