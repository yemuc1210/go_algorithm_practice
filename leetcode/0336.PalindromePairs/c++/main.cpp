#include<iostream>
#include<vector>
#include<bits/stdc++.h>
using namespace std;


class Solution {
public:
    vector<vector<int>> palindromePairs(vector<string>& words) {
        vector<vector<int>> res;
        int n = words.size();
        for(int i=0;i<n;i++) {
            for(int j=0;j<n;j++) {
                if(i==j) {
                    continue;
                }
                string s = words[i]+words[j];
                if(isPalindrome(s))  {
                    res.push_back(vector<int>{i,j});
                }
            }
        }
        return res;
    }
    bool isPalindrome(string s) {
        if(s.length()==1) {
            return true;
        }
        int i=0,j=s.length()-1;
        while(i<j && s[i]==s[j]) {
            i++;
            j--;
        }
        return i>j || i==j;
    }
};

int main() {
    Solution* s = new Solution();
    
    return 0;
}