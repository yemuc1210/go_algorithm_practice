import java.util.ArrayDeque;
import java.util.Queue;

// 933. 最近的请求次数 Java
class RecentCounter {
    int counter;
    // 队列
    Queue<Integer> queue;
    public RecentCounter() {
        counter = 0;  // 初始化计数器，请求数为0
        queue = new ArrayDeque<>();
    }
    
    public int ping(int t) {
        // 在时间t添加一个请求，t以毫秒为单位
        // 返回过去3000ms内发生的所有的请求数，包括新请求  [t-3000,t]
        // t加入队尾
        queue.offer(t);
        // 从队尾开始，移除小于t-3000的
        while(queue.peek() < t-3000) {
            queue.poll();
        }
        return queue.size();
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter obj = new RecentCounter();
 * int param_1 = obj.ping(t);
 */