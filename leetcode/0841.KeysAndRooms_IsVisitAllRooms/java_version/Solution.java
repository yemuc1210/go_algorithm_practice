import java.util.ArrayDeque;
import java.util.Deque;
import java.util.LinkedList;
import java.util.Queue;

class Solution {
    // 841. 钥匙和房间 Java
    // 进入房间，获得钥匙
    // 如果能进入所有房间，返回true
    public boolean canVisitAllRooms(List<List<Integer>> rooms) {
        if(rooms.size()==1){
            return true;
        }
        int n = rooms.size();
        boolean[] visited = new boolean[n];
        visited[0] = true;
        // Deque<Integer> queue = new ArrayDeque<>();
        Queue<Integer> queue = new LinkedList<>();
        queue.offer(0);
        while(!queue.isEmpty()){
            int curNode = queue.poll();
            // 获得keys
            List<Integer> keys = rooms.get(curNode);
            for (Integer key : keys) {
                if(!visited[key]){
                    // queue.push(key);
                    queue.add(key);
                    visited[key] = true;
                }
            }
        }
        //check
        for(int i=0;i<n;i++){
            if(!visited[i]) {
                return false;
            }
        }
        return true;
    }
}