import java.util.HashSet;
import java.util.Queue;

class Solution {
    // 1654. 到家的最少跳跃次数 Java
    public int minimumJumps(int[] forbidden, int a, int b, int x) {
        // 确定数据范围
        int INF = 8000;
        boolean[][] visited = new boolean[8001][2];
        Set<Integer> forbid = new HashSet<>(forbidden.length);
        for (int i : forbidden) {
            forbid.add(i);
        }
        

        Queue<int[]> queue = new ArrayDeque<>();
        queue.offer(new int[] { 0, 0 });
        int layer = -1;
        while (!queue.isEmpty()) {
            int curWidth = queue.size(); // 当前层的宽度
            layer++;
            for (int i = 0; i < curWidth; i++) {
                //
                int[] p = queue.poll();
                int cur = p[0], backwardCnt = p[1];
                if (cur == x) {
                    return layer;
                }
                if (visited[cur][backwardCnt]) {
                    continue;
                }
                //
                visited[cur][backwardCnt] = true;
                if (cur + a <= INF && !forbid.contains(cur + a)) {
                    queue.offer(new int[] { cur + a, 0 });
                }
                if (cur - b >= 0 && backwardCnt < 1 && !forbid.contains(cur - b)) {
                    queue.offer(new int[] { cur - b, backwardCnt + 1 });
                }
            }

        }
        return -1;
    }
}