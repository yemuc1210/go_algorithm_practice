#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    vector<int> drawLine(int length, int w, int x1, int x2, int y) {
        unsigned int cnt=0x80000000;
        int a=0;
        vector<int> ans;
        for(int i=0;i<y;++i){//把y行之前的int加入数组；
            int n=w/32;
            while(n!=0){
                ans.push_back(0);
                n--;
            }
        }
        for(int j=0;j<=w;++j){
            if(j>=x1&&j<=x2){//如果j在x1和x2之间；
                a=a|cnt;//把j位置1；
            }
            cnt=cnt>>1;
            if((j+1)%32==0){//将一个置换完成的int加入数组；
                ans.push_back(a);
                a=0;//恢复初值，为下一个int准备；
                cnt=0x80000000;
            }
        }
        if(length>w/32*(y+1)){//把剩下的int加入数组；
            int m=length-w/32*(y+1);
            for(int i=0;i<m;++i){
                ans.push_back(0);
            }
        }
            return ans;
        }
};


int main() {
    Solution* s=new Solution;
    int lenght=1,w=32,x1=30,x2=31,y=0;
    vector<int> a;
    a = s->drawLine(lenght,w,x1,x2,y);
    for(int i=0;i<a.size();i++) {
        int m = a.front();
        cout<<m<<endl;
    }
    return 0;
}