
import java.util.Stack;

class Solution { 
    // 121. 买卖股票的最佳时机  Java
    public int maxProfit(int[] prices) {
        // 某天买入，未来某天卖出，计算获得的最大利润
        // 买卖一次
        // 单调栈
        if(prices.length==0){
            return 0;
        }
        int res = 0;
        
        Stack<Integer> sk = new Stack<>();
        sk.push(prices[0]);
        for(int i=1;i<prices.length;i++){
            if(sk.peek()>prices[i]) {
                while(!sk.isEmpty() && sk.peek() > prices[i]) {
                    sk.pop();
                    sk.push(prices[i]);                    
                } // 保证栈顶始终是最小的
            }else{
                res = Math.max(res, prices[i] - sk.peek());
            }
        }
        return res;
    }
}