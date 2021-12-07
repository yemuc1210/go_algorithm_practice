class Solution {
public:
    int maxGroupNumber(vector<int>& tiles) {
        map<int, int> count;  // [点数, 对应的牌数]，C++中的map是按照key(点数)从小到大排列的。
        for (auto tile : tiles)
            count[tile]++;

        // dp[x][y] 表示
        // 在预留x张 [tile-2] 和y张 [tile-1] 的前提下，
        // [tile] 之前的牌能组成的牌组数
        vector<vector<int>> dp(5, vector<int>(5, -1));
        dp[0][0] = 0;
        int prev_tile = 0; // 前一张牌的点数
        for (auto [tile, cnt] : count) { // [tile表示当前点数, cnt表示对应的牌数]
            // 如果上一张牌和这张牌没法连起来
            // 意味着无论之前留几张牌，都无法和tile一起组成顺子，因此，只保留dp[0][0]的情形。
            if (prev_tile != tile - 1) {     
                // dp[0][0] 表示，之前留下的 “tile-2点的牌数” 和 “tile-1点的牌数” 都为0
                int dp00 = dp[0][0];
                // dp[x][y] == -1 表示，之前没有 “留下x张[tile-2]点的和y张[tile-1]点” 的情况
                dp = vector<vector<int>>(5, vector<int>(5, -1));
                dp[0][0] = dp00;
            }
            // 新的dp数组
            vector<vector<int>> new_dp(5, vector<int>(5, -1));
            
            for (int cnt_2 = 0; cnt_2 < 5; ++cnt_2) // [tile-2] 的牌数
                for (int cnt_1 = 0; cnt_1 < 5; ++cnt_1) { // [tile-1] 的牌数
                    // 如果之前没有留下这么多张牌
                    if (dp[cnt_2][cnt_1] < 0)
                        continue;

                    // 顺子数量不能超过[tile-2]、[tile-1]、[tile]的牌数
                    for (int shunzi = 0; shunzi <= min(cnt_2, min(cnt_1, cnt)); ++shunzi) {
                        int new_2 = cnt_1 - shunzi; // 对于下一个点数 new_tile = tile + 1 而言，
                                                    // [new_tile - 2] 就是当前的 [tile - 1]
                                                    // new_2 代表预留的 [new_tile - 2] 的牌数
                                                    // 也就是当前的 [tile - 1] 的牌数 - 顺子数量
                        // 同理，对于下一个点数 [new_tile] 而言，new_1 代表预留的 [new_tile - 1] 的牌数，
                        // 也就是预留的 [tile] 的数量。
                        // 预留的数量不超过四张，也不超过 ( [tile]的牌数 - 顺子数量 )
                        for (int new_1 = 0; new_1 <= min(4, cnt - shunzi); ++new_1) {
                            // 新的牌组数等于以下三者相加：
                            // 1. dp数组保存的，留下 cnt_2 张 [tile-2] 和 cnt_1 张 [tile-1] 的前提下，tile-1 之前的牌面能凑出来的牌组数
                            // 2. 顺子数量
                            // 3. [tile] 组成的刻子数量 = ( [tile] - 顺子数量 - 留下备用的牌 ) / 3
                            int new_score = dp[cnt_2][cnt_1] + shunzi + (cnt - shunzi - new_1) / 3;
                            new_dp[new_2][new_1] = max(new_dp[new_2][new_1], new_score);
                        }
                    }
                }
            
            // 将new_dp数组赋值给dp数组
            dp = move(new_dp);
            // 将当前tile记录都上一个tile中
            prev_tile = tile;
        }
        
        // 找到并返回dp的最大值
        int ans = 0;
        for (int cnt_2 = 0; cnt_2 < 5; ++cnt_2)
            for (int cnt_1 = 0; cnt_1 < 5; ++cnt_1)
                ans = max(ans, dp[cnt_2][cnt_1]);
        
        return ans;
    }
};

