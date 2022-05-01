

const int dx[] = {-1, 0, 1, 0}, dy[] = {0, 1, 0, -1};
struct S{
    int curt, x, y, st;
    S(int a, int b, int c, int d): curt(a), x(b), y(c), st(d){}
};
class Solution {
    bool ok[105][55][55][6];
    queue<S> q;
public:
    bool escapeMaze(vector<vector<string>>& maze) {
        int T = maze.size();
        int n = maze[0].size(), m = maze[0][0].length();
        
        ok[0][0][0][3] = true;
        // 1: 能临时，不能停留; 0: 不能临时，不能停留
        // 3: 能临时，能停留; 2: 不能临时，能停留
        // 5: 能临时，在停留中; 4: 不能临时，在停留中
        q.push(S(0, 0, 0, 3));
        while (!q.empty()){
            S h = q.front();
            q.pop();
            int curt = h.curt;
            if (curt == T - 1) continue ;
            vector<string>& mz = maze[curt + 1];
            int x = h.x;
            int y = h.y;
            int st = h.st;

            if (st >= 4){
                // 继续停留
                if (!ok[curt + 1][x][y][st]){
                    ok[curt + 1][x][y][st] = true;
                    q.push(S(curt + 1, x, y, st));
                }
                // 不再停留
                st -= 4;
                for (int i = 0; i < 4; ++i){
                    int cx = x + dx[i], cy = y + dy[i];
                    if (cx < 0 || cx >= n || cy < 0 || cy >= m)
                        continue;
                    if (mz[cx][cy] == '#'){
                        // 需要临时经过
                        if (!st) continue;
                        if (!ok[curt + 1][cx][cy][0]){
                            ok[curt + 1][cx][cy][0] = true;
                            q.push(S(curt + 1, cx, cy, 0));
                        }
                    } else {
                        if (!ok[curt + 1][cx][cy][st]){
                            ok[curt + 1][cx][cy][st] = true;
                            q.push(S(curt + 1, cx, cy, st));
                        }
                    }
                }
            } else {
                if (st & 2){
                    // 尝试开始停留
                    if (!ok[curt + 1][x][y][st + 2]){
                        ok[curt + 1][x][y][st + 2] = true;
                        q.push(S(curt + 1, x, y, st + 2));
                    }
                }
                for (int i = 0; i < 4; ++i){
                    int cx = x + dx[i], cy = y + dy[i];
                    if (cx < 0 || cx >= n || cy < 0 || cy >= m)
                        continue;
                    if (mz[cx][cy] == '#'){
                        if (st & 1){
                            if (!ok[curt + 1][cx][cy][st - 1]){
                                ok[curt + 1][cx][cy][st - 1] = true;
                                q.push(S(curt + 1, cx, cy, st - 1));
                            }
                        }
                    } else {
                        if (!ok[curt + 1][cx][cy][st]){
                            ok[curt + 1][cx][cy][st] = true;
                            q.push(S(curt + 1, cx, cy, st));
                        }
                    }
                    if (st & 2){
                        // 尝试前去停留
                        if (!ok[curt + 1][cx][cy][st + 2]){
                            ok[curt + 1][cx][cy][st + 2] = true;
                            q.push(S(curt + 1, cx, cy, st + 2));
                        }
                    }
                }
                
                // 呆在原地，但不停留
                if (mz[x][y] == '#'){
                    if (st & 1){
                        if (!ok[curt + 1][x][y][st - 1]){
                            ok[curt + 1][x][y][st - 1] = true;
                            q.push(S(curt + 1, x, y, st - 1));
                        }
                    }
                } else {
                    if (!ok[curt + 1][x][y][st]){
                        ok[curt + 1][x][y][st] = true;
                        q.push(S(curt + 1, x, y, st));
                    }
                }
            }
        }
        
        for (int i = 0; i < 6; ++i)
            if (ok[T - 1][n - 1][m - 1][i]) return true;
        return false;
    }
};
