import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.LinkedList;

class Solution {
    // 797. 所有可能的路径
    // 节点0到n-1的路径
    List<List<Integer>> allPaths = new ArrayList<List<Integer>>();
    Deque<Integer> stack = new ArrayDeque<>();

    public List<List<Integer>> allPathsSourceTarget(int[][] graph) {
        stack.offerLast(0);  // 起点
        dfs(graph.length-1, 0, graph);

        return allPaths;
    }
    public void dfs(int target, int curNode,int[][] graph) {
        if(curNode == target) {
            // 得到一条路径
            allPaths.add(new ArrayList<Integer>(stack));
            return ;
        }
        for(int node : graph[curNode]) {
            stack.offerLast(node);
            dfs(target,node,graph);
            stack.pollLast();
        }
    }
}