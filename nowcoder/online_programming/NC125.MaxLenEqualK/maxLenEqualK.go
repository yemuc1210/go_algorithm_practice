package NC125

/**
 * max length of the subarray sum = k
 * @param arr int整型一维数组 the array
 * @param k int整型 target
 * @return int整型
*/
func maxlenEqualK( arr []int ,  k int ) int {
    // 求累计k最长的连续子数组长度
	//offer 10   lt560
	return subarraySum(arr,k)

}

func subarraySum(nums []int, k int) int {
    count := 0
	maxLen := 0
    for start := 0; start < len(nums); start++ {
        sum := 0
        for end := start; end >= 0; end-- {
            sum += nums[end]
            if sum == k {
				// len := start-end 
				maxLen = max(maxLen, start-end)
                count++
            }
        }
    }
    return maxLen
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}



// public class Solution {
//     /**
//      * max length of the subarray sum = k
//      * @param arr int整型一维数组 the array
//      * @param k int整型 target
//      * @return int整型
//      */
//     public int maxlenEqualK (int[] arr, int k) {
//         // write code here
//         int len  = arr.length;
//         if(arr == null || len == 0){
//             return 0;
//         }
//         Map<Integer,Integer> map = new HashMap<Integer,Integer>();
//         map.put(0,-1);
//         int length = 0;
//         int sum = 0;
//         for(int i = 0;i<arr.length;i++){
//             sum += arr[i];
//             if(map.containsKey(sum-k)){
//                 length = Math.max(i-map.get(sum-k),length);
//             }
//             if(!map.containsKey(sum)){
//                 map.put(sum,i);
//             }
//         }
//         return length;
//     }
// }
