import java.util.ArrayDeque;
import java.util.Queue;

class Solution {
    // 1306. 跳跃游戏 III Java
    public boolean canReach(int[] arr, int start) {
        // 处于下标i时，可以和跳跃到 i+arr[i]  或者i-arr[i]
        // 判断能否到达元素值为0的任一下标处  则不能移动了
        // 不能越界
        int n = arr.length;
        if(start<0 || start>n){
            return false;
        }    
        boolean[] visited = new boolean[n];
        Queue<Integer> queue = new ArrayDeque<>();
        queue.offer(start);
        while(!queue.isEmpty()){
            int size = queue.size();
            for(int i=0;i<size;i++){
                int cur = queue.poll();
                // 越界检查
                if(cur<0 || cur>=n) {
                    continue;
                }
                if(visited[cur]){
                    continue;
                }
                if(arr[cur]==0){
                    return true;
                }
                visited[cur]=true;
                queue.offer(cur+arr[cur]);
                queue.offer(cur-arr[cur]);
            }

        }
        return false;

    }
}