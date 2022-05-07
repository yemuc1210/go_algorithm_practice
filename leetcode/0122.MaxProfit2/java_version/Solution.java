import java.util.Stack;

class Solution {
    // 122. 买卖股票的最佳时机 II Java
    public int maxProfit(int[] prices) {
        // 每一天，可以觉得是否购买或出售，最多只能持有一股，可以先购买，然后同一天出售
        // 上升趋势的总和
        if (prices == null || prices.length == 0) {
            return 0;
        }
        int res = 0;
        Stack<Integer> sk = new Stack<>();
        sk.push(prices[0]);
        for (int i = 1; i < prices.length; i++) {
            // 入栈元素大于栈顶，入栈
            if (prices[i] > sk.peek()) {
                sk.push(prices[i]);
            } else {
                // 否则 入栈元素小于栈顶，循环出栈，计算栈顶和栈底的差值
                int top = sk.peek(), tmpMax = -1;
                while (!sk.isEmpty()) {
                    tmpMax = Math.max(tmpMax, top - tmpMax);
                }
                //
                sk.push(prices[i]);
                res += tmpMax;
            }
        }
        int top = sk.peek(), tmpMax = -1;
        while (!sk.isEmpty()) {
            int t = sk.pop();
            tmpMax = Math.max(tmpMax, top - t);
        }
        return res + tmpMax;
    }
}