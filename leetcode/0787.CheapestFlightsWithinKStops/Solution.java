class Solution {
    // 787. K 站中转内最便宜的航班  Java
    // flights[i] = [from,to,price]   带权图
    // 路径：最多不超过k个中转站  保证src-dst价格最便宜
    int totalPrice = 0;
    public int findCheapestPrice(int n, int[][] flights, int src, int dst, int k) {

    }
    void dfs(int[] path,int k,int curPrice, int curNode){
        if(path.length>k) {
            return ;
        }
        // 
    }
}