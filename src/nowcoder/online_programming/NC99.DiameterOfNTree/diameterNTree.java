/**
 * 思路
题目中有提到连通图，所以就想着将给的数据转换成无向图，再遍历无向图找到一条最长路径，使用邻接矩阵存储无向图。

具体步骤
首先将所给数据转换成图，由所给例子可以看出，其是个无向图，故，转换成无向图时，需要将(i,j)与(j,i)同时设置为边对应的权重；
以每一个节点为起点，去计算以当前起点为节点的最大路径
取所有路径中的最大值，即为最长路径。

 */

// import java.util.*;
// /*
// * public class Interval {
// *   int start;
// *   int end;
// * }
// */
// public class Solution {
//     /**
//     * 树的直径
//     * @param n int整型 树的节点个数
//     * @param Tree_edge Interval类一维数组 树的边
//     * @param Edge_value int整型一维数组 边的权值
//     * @return int整型
//     */
//     public int solve (int n, Interval[] Tree_edge, int[] Edge_value) {
//     // write code here
//     int[][] graph = new int[n][n];
//     // 将树的边以及权重转换成无向图
//     for (int i = 0;i < Tree_edge.length;++i){
//         graph[Tree_edge[i].start][Tree_edge[i].end] = Edge_value[i];
//         graph[Tree_edge[i].end][Tree_edge[i].start] = Edge_value[i];
//     }
//     int maxVal = 0;
//     // 计算每个点最长路径
//     for (int i = 0; i < n; ++i){
//         int temp = compute(graph,i,new boolean[n]);
//         maxVal = maxVal > temp ? maxVal : temp;
//     }
//     // 返回最长路径
//     return maxVal;
//     }
//     // 递归计算每一个点的最长路径
//     private int compute(int[][] graph,
//                     int index, boolean[] to){
//     int max = 0;
//     to[index] = true;
//     int temp = 0;
//     for (int i = 0; i < graph.length; ++i){
//         // 判断两点之间是否有路径
//         if (graph[index][i] != 0){
//             // 判断当前点是否已经走过
//             if (!to[i]){
//                 to[i] = true;
//                 // 计算下一个点的路径
//                 temp = compute(graph,i,to)+graph[index][i];
//                 max = max > temp ? max : temp;
//             }
//         }
//     }
//     return max;
//     }
// }

/**
 * 思路
方法一的时间复杂度与空间复杂度都是比较大的，无法完全通过牛客上的所有测试用例，所以需要对其进行优化，可以考虑使用邻接表来存储无向图，并且容易想到，在计算路径过程中其实是有很多冗余计算的，所以考虑使用来减少冗余计算，降低时间复杂度。
笔者又从网上查了查树的直径，发现其还是有更为简便的计算方法，在一个连通无向无环图中，以任意结点出发所能到达的最远结点，一定是该图直径的端点之一，故可以做两次DFS来计算树的直径。
具体步骤
首先，根据连通关系以及边的权重构建无向图；
随机找一顶点，利用深度优先搜索找到距离该点最远的顶点remote，这里默认就是0节点。
从remote顶点开始深度优先搜索找到最长路径，该路径即为直径。
 */

import java.util.*;
public class Solution {
// 图节点类
class Node{
   int start;
   Map<Integer,Integer> edges = new HashMap<>();
   public Node(int start){
       this.start = start;
   }
}
/**
* 树的直径
* @param n int整型 树的节点个数
* @param Tree_edge Interval类一维数组 树的边
* @param Edge_value int整型一维数组 边的权值
* @return int整型
*/
public int solve (int n, Interval[] Tree_edge, int[] Edge_value) {
   // write code here
   if(Tree_edge == null || Edge_value == null ||
      Tree_edge.length != Edge_value.length){return 0;}
   // 将树的边以及权重转换成无向图
   Map<Integer,Node> graph = trans(Tree_edge,Edge_value);
   // remote[0] 代表以0为起点的最长路径长度，
   // remote[1]代表最长路径的终点
   int[] remote = {0, 0};    
   dfs(graph, 0, -1, 0, remote);
   int[] res = {0, 0};
   dfs(graph, remote[1], -1, 0, res);
   return res[0];
}
// 转换函数
private Map<Integer,Node> trans(Interval[] Tree_edge, int[] Edge_value) {
   int len = Tree_edge.length;
   Map<Integer,Node> graph = new HashMap<>();
   for (int i = 0;i < len; ++i){
       Interval edge = Tree_edge[i];
       if (!graph.containsKey(edge.start)){
           Node node = new Node(edge.start);
           node.edges.put(edge.end,Edge_value[i]);
           graph.put(node.start,node);
       }
       graph.get(edge.start).edges.put(edge.end,Edge_value[i]);
       // 无向图所以需要建两次边
       if (!graph.containsKey(edge.end)){
           Node node = new Node(edge.end);
           node.edges.put(edge.start,Edge_value[i]);
           graph.put(node.start,node);
       }
       graph.get(edge.end).edges.put(edge.start,Edge_value[i]);
   }
   return graph;
}
// 深度优先搜索
private void dfs(Map<Integer,Node> graph, int from, int prev, int path, int[] res){
   Node node = graph.get(from);
   Map<Integer,Integer> edges = node.edges;
   for (Map.Entry<Integer,Integer> edge : edges.entrySet()){
       // 未曾经过
       if (edge.getKey() != prev){
           path += edge.getValue();
           if (path > res[0]){
               res[0] = path;
               res[1] = edge.getKey();
           }
           dfs(graph,edge.getKey(),from,path,res);
           // 回溯
           path -= edge.getValue();
       }
   }
}
}