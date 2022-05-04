import java.util.ArrayDeque;
import java.util.Queue;

class Solution {
    // 1823. 找出游戏的获胜者 Java
    public int findTheWinner(int n, int k) {
        int winner = 0;
        for(int i=2;i<=n;i++){
            winner = (winner+k)%i;
        }
        return winner+1;
    }
    public int findTheWinner1(int n,int k) {
        Queue<Integer> queue = new ArrayDeque<>();
        for(int i=1;i<=n;i++){
            queue.offer(i);
        }
        while(!queue.isEmpty()){
            for(int i=1;i<k;i++){
                queue.offer(queue.poll());
            }
            queue.poll();
        }
        return queue.poll();
    }
}