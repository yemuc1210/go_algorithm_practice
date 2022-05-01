#include<vector>
#include<iostream>
using namespace std;

const int N = 1e5;
struct node
{
    //区间[l,r]的和为sum
    int l, r, sum, lazytag;
};
vector<node> tree(N);
vector<int> nums;
//递归建树
/*
每次将区间[l,r]从中划分为两段
id的左孩子为2*id指向左半区间
id的右孩子为2*id+1指向右半区间
当l==r时到达叶子节点，返回即可
id：当前节点
l：区间左边界
r：区间右边界
*/
void build(int id, int l, int r)
{
    tree[id].l = l;
    tree[id].r = r;
    if (l == r)
    {
        tree[id].sum = nums[l];
        return;
    }
    int mid = (l + r) / 2;
    build(2 * id, l, mid);
    build(2 * id + 1, mid + 1, r);
    tree[id].sum = tree[2 * id].sum + tree[2 * id + 1].sum;
    return;
}
//区间查询：查询给定区间[l,r]的和
/*
从根节点开始查询，节点区间与给定区间之间存在四种情况
1.节点区间完全包含于给定区间，此时查询结果加上节点区间和，直接返回
2.节点区间与给定区间完全不相交，此时直接返回0
3.节点区间的左儿子与给定区间相交，此时递归查询左儿子
4.节点区间的右儿子与给定区间相交，此时递归查询右儿子
id：当前节点
l：区间左边界
r：区间右边界
*/
int search(int id, int l, int r)
{
    if (tree[id].l >= l && tree[id].r <= r)
        return tree[id].sum;
    if (tree[id].r < l || tree[id + 1].l > r)
        return 0;
    int s = 0;
    if (tree[2 * id].r >= l)
        s += search(2 * id, l, r);
    if (tree[2 * id + 1].l <= r)
        s += search(2 * id + 1, l, r);
    return s;
}
//单点修改
/*
一直搜索到叶子节点，首先修改叶子结点的sum，然后逐层返回，修改对应的sum
id：当前节点
i：要修改的节点
k：修改的值
*/
void change(int id, int i, int k)
{
    if (tree[id].l == i && tree[id].r == i)
    {
        tree[id].sum = k;
        return;
    }
    int mid = (tree[id].l + tree[id].r) / 2;
    if (i <= mid)
        // 左孩子 2*id
        change(2 * id, i, k);
    else
        change(2 * id + 1, i, k);
    tree[id].sum = tree[2 * id].sum + tree[2 * id + 1].sum;
    return;
}
//——————————————————————————————————————————————————————————————————————————————————————————
/*
lazytag：记录区间的修改情况
pushdown函数：下传懒标记，即将当前区间的修改情况下传到其左右孩子节点
*/
// x:下传x节点的懒标记
void pushdown(int x)
{
    if (tree[x].lazytag != 0)
    {
        int lt = tree[x].lazytag;
        tree[2 * x].lazytag += lt;
        tree[2 * x].sum += lt * (tree[2 * x].r - tree[2 * x].l + 1);
        tree[2 * x + 1].lazytag += lt;
        tree[2 * x + 1].lazytag += lt * (tree[2 * x + 1].r - tree[2 * x + 1].l + 1);
        tree[x].lazytag -= 0;
    }
    return;
}
//区间修改
/*
1.如果节点区间完全包含于目标区间，则将节点的sum加上区间长度*修改情况
2.如果节点区间没有完全包含于目标区间，则先下传懒标记
3.如果这个区间的左儿子和目标区间有交集，那么搜索左儿子
4.如果这个区间的右儿子和目标区间有交集，那么搜索右儿子
*/
void modify(int id, int l, int r, int k)
{
    if (tree[id].l >= l && tree[id].r <= r)
    {
        tree[id].sum += k * (tree[id].r - tree[id].l + 1);
        tree[id].lazytag += k;
        return;
    }
    pushdown(id);
    if (tree[2 * id].r >= l)
        modify(2 * id, l, r, k);
    if (tree[2 * id + 1].l <= r)
        modify(2 * id + 1, l, r, k);
    tree[id].sum = tree[2 * id].sum + tree[2 * id + 1].sum;
    return;
}
//区间查询
int query(int id, int l, int r)
{
    if (tree[id].l >= l && tree[id].r <= r)
        return tree[id].sum;
    if (tree[id].r < l || tree[id].l > r)
        return 0;
    pushdown(id);
    int s = 0;
    if (tree[2 * id].r >= l)
        s += query(2 * id, l, r);
    if (tree[2 * id + 1].l <= r)
        s += query(2 * id + 1, l, r);
    return s;
}
class NumArray
{
public:
    NumArray(vector<int> &num)
    {
        nums.assign(num.begin(), num.end());
        build(1, 0, nums.size() - 1);
    }

    void update(int index, int val)
    {
        change(1, index, val);
    }

    int sumRange(int left, int right)
    {
        return query(1, left, right);
    }
};