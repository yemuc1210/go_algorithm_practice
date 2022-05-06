import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.Map;
import java.util.Queue;
// 127. 单词接龙 Java?
class Solution {
    // 构建图结构，如果两个单词可以通过改变一个字母相互转换，则存在双向边
    Map<String,Integer> wordId = new HashMap<String,Integer>();
    List<List<Integer>> graph = new ArrayList<List<Integer>>();
    int nodeNum =0;
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        for(String word : wordList) {
            addEdge(word);
        }
        addEdge(beginWord);
        if(!wordId.containsKey(endWord)) {
            return 0;
        }
        int[] dis = new int[nodeNum];
        Arrays.fill(dis, Integer.MAX_VALUE);

        int beginId = wordId.get(beginWord),endId=wordId.get(endWord);
        dis[beginId]=0;

        //bfs
        Queue<Integer> queue = new LinkedList<>();
        queue.offer(beginId);

        while(!queue.isEmpty()) {
            int top = queue.poll();
            if(top == endId) {
                return dis[endId]/2 + 1;
            }
            for(int it:graph.get(top)) {
                if(dis[it] == Integer.MAX_VALUE) {
                    dis[it] = dis[top] +1;
                    queue.offer(it);
                }
            }
        }
        return 0;
    }
    void addEdge(String word) {
        addWord(word);
        int id1 = wordId.get(word);
        char[] arr = word.toCharArray();
        int len = arr.length;
        for(int i=0;i<len;i++){
            char tmp = arr[i];
            arr[i]='*';
            String newWord = new String(arr);
            addWord(newWord);

            int id2 = wordId.get(newWord);
            graph.get(id1).add(id2);
            graph.get(id2).add(id1);
            arr[i] = tmp;

        }
    }
    void addWord(String word) {
        if(!wordId.containsKey(word)) {
            wordId.put(word, nodeNum++);
            graph.add(new ArrayList<Integer>());
        }
    }
}