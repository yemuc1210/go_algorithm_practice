/*
LCP 37. 最小矩形面积
二维平面上有 NN 条直线，形式为 y = kx + b，其中 k、b为整数 且 k > 0。所有直线以 [k,b] 的形式存于二维数组 lines 中，不存在重合的两条直线。两两直线之间可能存在一个交点，最多会有 C_N^2C 
N
2
​
  个交点。我们用一个平行于坐标轴的矩形覆盖所有的交点，请问这个矩形最小面积是多少。若直线之间无交点、仅有一个交点或所有交点均在同一条平行坐标轴的直线上，则返回0。

注意：返回结果是浮点数，与标准答案 绝对误差或相对误差 在 10^-4 以内的结果都被视为正确结果
*/


const int T = 50;
const double inf = 1e9;

class Solution {
public:
    double minRecSize(vector<vector<int>>& lines) {
        int n = lines.size();
        if (n == 1)
            return 0.0;
        
        double rl = -inf, rr = inf;
        int t = T;
        sort(lines.begin(), lines.end(), [](vector<int> &a, vector<int> &b){
            return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1]); 
        });
        while (t--) {
            double mid = (rl + rr) * 0.5;
            bool conflict = false;
            for (int i = 0; i < n - 1; ++i) {
                double a = (double)lines[i][0] * mid + lines[i][1];
                double b = (double)lines[i + 1][0] * mid + lines[i + 1][1];
                if (a - b > 1e-8) {
                    conflict = true;
                    break;
                }
            }
            if (conflict)
                rl = mid;
            else
                rr = mid;
        }
        
        t = T;
        double ll = -inf, lr = inf;
        sort(lines.begin(), lines.end(), [](vector<int> &a, vector<int> &b){
            return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1]); 
        });
        while (t--) {
            double mid = (ll + lr) * 0.5;
            bool conflict = false;
            for (int i = 0; i < n - 1; ++i) {
                double a = (double)lines[i][0] * mid + lines[i][1];
                double b = (double)lines[i + 1][0] * mid + lines[i + 1][1];
                if (a - b > 1e-8) {
                    conflict = true;
                    break;
                }
            }
            if (conflict)
                lr = mid;
            else
                ll = mid;
        }
        
        t = T;
        double ul = -inf, ur = inf;
        sort(lines.begin(), lines.end(), [](vector<int> &a, vector<int> &b){
            return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1]); 
        });
        while (t--) {
            double mid = (ul + ur) * 0.5;
            bool conflict = false;
            for (int i = 0; i < n - 1; ++i) {
                double a = (mid - lines[i][1]) / lines[i][0];
                double b = (mid - lines[i + 1][1]) / lines[i + 1][0];
                if (a - b > 1e-8) {
                    conflict = true;
                    break;
                }
            }
            if (conflict)
                ul = mid;
            else
                ur = mid;
        }
        
        t = T;
        double dl = -inf, dr = inf;
        sort(lines.begin(), lines.end(), [](vector<int> &a, vector<int> &b){
            return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1]); 
        });
        while (t--) {
            double mid = (dl + dr) * 0.5;
            bool conflict = false;
            for (int i = 0; i < n - 1; ++i) {
                double a = (mid - lines[i][1]) / lines[i][0];
                double b = (mid - lines[i + 1][1]) / lines[i + 1][0];
                if (a - b > 1e-8) {
                    conflict = true;
                    break;
                }
            }
            if (conflict)
                dr = mid;
            else
                dl = mid;
        }
        
        if (ll - rr > 1e-8)
            return 0.0;
        
        return (rr - ll) * (ur - dl);
    }
};

