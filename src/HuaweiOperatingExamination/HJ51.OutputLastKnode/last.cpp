#include<iostream>
using namespace std;
struct ListNode{
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(NULL){}
};

//头插
int main(){
    int nums;
    while(cin >> nums){
        ListNode *pHead = new ListNode(-1);
        ListNode *p = pHead;
        for(int i = 0; i < nums; ++i){
            int data;
            cin >> data;
            ListNode *q = new ListNode(data);
            p->next = q;
            p = p->next;
        }
        int Kth;
        cin >> Kth;
        p = pHead;
        if(Kth == 0)  // 边界测试
            cout << "0" << endl;
        else if(nums-Kth >= 0){
            for(int i = 0; i <= nums-Kth; ++i)
                p = p->next;
            cout << p->val << endl; 
        }
        else
        	cout << "NULL" << endl; 
    }
    return 0;
}