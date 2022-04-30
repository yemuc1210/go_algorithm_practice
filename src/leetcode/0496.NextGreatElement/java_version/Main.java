import java.util.Deque;
import java.util.HashMap;
import java.util.LinkedList;

class Solution {
    // 496. 下一个更大元素 I
    // nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
    // nums1 nums2没有重复元素
    // 返回 ans[i] 
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        int n = nums1.length;
        int[] ans = new int[n];
        // 关键,找到nums2[j] == nums1[i] 然后找到右侧第一个大的
        // map: k elem nxtLarget
        // 右侧第一个比它大的   单调栈
        Deque<Integer> sk = new LinkedList<>();
        HashMap<Integer,Integer> map = new HashMap<>();
        // nums1 nums2 长度不一定相等
        for(int i=nums2.length-1;i>=0;i--){
            int num = nums2[i];
            // 维护单调栈,弹出笔当前元素小的
            while(!sk.isEmpty() && num>=sk.peek()){
                sk.pop();
            }
            // 不为空,则得到下一个元素
            if(!sk.isEmpty()){
                // map[num] = sk.peek();
                map.put(num, sk.peek());
            }else{
                map.put(num, -1);
            }
            // sk.add(num);
            sk.push(num);
        }
        // 遍历nums1
        for(int i=0;i<n;i++){
            ans[i] = map.getOrDefault(nums1[i], -1);
        }
        return ans;
    }
    
}

public class Main {
    public static void main(String[] args) {
        Solution s1 = new Solution();
        int[] nums1 = {2,4};
        int[] nums2 = {1,2,3,4};
        int[] res = s1.nextGreaterElement(nums1, nums2);
        for (int i : res) {
            System.out.println(i);
        }
    }
}