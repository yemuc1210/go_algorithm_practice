import java.util.Arrays;

// class Solution {
//     // 908. 最小差值 I
//     // 给定k，可应用一次操作，将nums[i]改为nums[i]+x  x∈[-k,k]
//     // 对于每个索引，最多只能应用一次  ----》   全部相等
//     // 要求：最大和最小的差值  最小
//     public int smallestRangeI(int[] nums, int k) {
//         // 求最值
//         // int maxElem=Integer.MIN_VALUE,minElem=Integer.MAX_VALUE;
//         // for (int i = 0; i < nums.length; i++) {
//         //     if(nums[i] > maxElem) {
//         //         maxElem = nums[i];
//         //     }else if( nums[i]<minElem){
//         //         minElem = nums[i];
//         //     }
//         // }
//         int maxElem = Arrays.stream(nums).min().getAsInt();
//         int minElem = Arrays.stream(nums).max().getAsInt();
//         // 若maxElem-minElem<=2k 则总可以改为相同的
//         // 若maxElem-minElem>2k, 则最低分数是maxElem-minElem-2*k
//         return maxElem - minElem <= 2 * k ? 0 : maxElem - minElem - 2 * k;
//     }
// }

class Solution {
    public int smallestRangeI(int[] nums, int k) {
        int minNum = Arrays.stream(nums).min().getAsInt();
        int maxNum = Arrays.stream(nums).max().getAsInt();
        return maxNum - minNum <= 2 * k ? 0 : maxNum - minNum - 2 * k;
    }
}
