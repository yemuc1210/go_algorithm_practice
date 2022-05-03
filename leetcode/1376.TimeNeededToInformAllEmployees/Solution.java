import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Queue;
class Solution {
    public int numOfMinutes(int n, int headID, int[] manager, int[] informTime) {
         Map<Integer, List<Integer>> graph = new HashMap<>();
        for (int i = 0; i < manager.length; i++) {
            graph.put(i,new ArrayList<>());
        }

        for (int i = 0; i < manager.length; i++) {
            if(manager[i]!=-1){
                graph.get(manager[i]).add(i);
            }
        }

        Queue<int[]> queue = new LinkedList<>();
        queue.offer(new int[]{headID,informTime[headID]});
        int ans  = 0;

        while (!queue.isEmpty()){

            int qsize = queue.size();
            for (int i = 0; i < qsize; i++) {
                int[] pair  = queue.poll();
                int nowhead = pair[0];
                int step = pair[1];

                ans = Math.max(ans,step);

                int size = graph.get(nowhead).size();
                for (int j = 0; j < size; j++) {
                    int idx = graph.get(nowhead).get(j);
                    queue.offer(new int[]{idx,step+informTime[idx]});
                }
            }
        }

        return ans;
    }
}
