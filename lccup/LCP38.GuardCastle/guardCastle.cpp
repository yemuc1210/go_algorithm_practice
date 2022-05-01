const int inf=1e9;
template<typename T>
struct dinic{
    struct edge {
        int u, rev;
        T cap, flow;
    };

    int n, s, t;
    T flow;
    vector<int> cur;
    vector<int> dep;
    vector<vector<edge>> to;

    T scaling_lim;
    bool scaling;

    dinic(int n, int s, int t, bool scaling = false) : n(n), s(s), t(t), scaling(scaling) {
        to.resize(n);
        for(auto& t:to)t.reserve(5);
        dep.resize(n);
        cur.resize(n);
        flow = 0;
    }

    void add_edge(int v, int u, T cap, bool directed = true) {
        to[v].push_back({ u, (int)to[u].size(), cap, 0 });
        to[u].push_back({ v, (int)to[v].size() - 1, directed ? 0 : cap, 0 });
    }

    T dfs(int v, T flow) {
        if (v == t)
            return flow;
        if (flow == 0)
            return 0;
        for (; cur[v] < to[v].size(); ++cur[v]) {
            edge& e = to[v][cur[v]];
            if (dep[e.u] != dep[v] + 1)
                continue;
            T add = dfs(e.u, min(flow, e.cap - e.flow));
            if (add > 0) {
                flow -= add;
                e.flow += add;
                to[e.u][e.rev].flow -= add;
                return add;
            }
            if (flow == 0)
                break;
        }
        return 0;
    }

    bool bfs() {
        fill(dep.begin(), dep.end(), -1);
        queue<int> q({ s });
        dep[s] = 0;
        while (!q.empty() && dep[t] == -1) {
            int v = q.front(); q.pop();
            for (auto& e : to[v]) {
                if (dep[e.u] == -1 && e.cap - e.flow >= scaling_lim) {
                    q.push(e.u);
                    dep[e.u] = dep[v] + 1;
                }
            }
        }
        return dep[t] != -1;
    }

    T maxflow() {
        T max_lim = numeric_limits<T>::max() / 2 + 1;
        for (scaling_lim = scaling ? max_lim : 1; scaling_lim > 0; scaling_lim >>= 1) {
            while (bfs()) {
                fill(cur.begin(), cur.end(), 0);
                T add;
                while ((add = dfs(s, numeric_limits<T>::max())) > 0)
                    flow += add;
            }
        }
        return flow;
    }
};
class Solution {
public:
    int guardCastle(vector<string>& mp) {
        // n为列数
        int n = mp[0].size();

        // 原图中的每个点拆成两个点
        vector<vector<pair<int, int>>> p(2, vector<pair<int, int>>(n));
        for(int i=0; i<2; i++){
            for(int j=0; j<n; j++){
                p[i][j] = {i*n+j+1, (i+2)*n+j+1};
            }
        }

        // 建图
        dinic<int> g(4*n+2, 0, 0);
        for(int i=0; i<2; i++){
            for(int j=0; j<n; j++){
                // 从起点向出生点连边
                if(mp[i][j]=='S')
                    g.add_edge(0, p[i][j].first, inf);

                // 从城堡向终点连边
                else if(mp[i][j]=='C')
                    g.t = p[i][j].second;
                
                // 将每个传送点连向一个特殊的虚拟点（4n+1）
                else if(mp[i][j]=='P')
                    g.add_edge(p[i][j].second, 4*n+1, inf), 
                    g.add_edge(4*n+1, p[i][j].first, inf);
                
                if(mp[i][j]!='#'){
                    // 除障碍物点外，每个点向周围连边
                    g.add_edge(p[i][j].second, p[i^1][j].first, inf);
                    if(j<n-1)
                        g.add_edge(p[i][j].second, p[i][j+1].first, inf);
                    if(j>0)
                        g.add_edge(p[i][j].second, p[i][j-1].first, inf);

                    // 自己拆的点之间连边
                    g.add_edge(p[i][j].first, p[i][j].second, mp[i][j]=='.'?1:inf);
                }
            }
        }
        int ans = g.maxflow();
        return ans<2*n?ans:-1;
    }
};
