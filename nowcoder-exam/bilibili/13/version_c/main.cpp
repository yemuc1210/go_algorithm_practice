#include <cstdio>
#include <vector>
#include <algorithm>
using namespace std;

int main()
{
    long n, k;
    scanf("%ld%ld", &n, &k);
    vector<long> A(n);
    for (int i = 0; i < n; ++i)
        scanf("%ld", &A[i]);
    sort(A.begin(), A.end());

    long left = 1, right = A[n-1]*A[n-2];
    while (left<=right)
    {
        long mid = (left+right)/2;
        long count = 0;
        int i = 0, j = n-1;
        while (i < j && count < k)
        {
            while (i < j && A[i]*A[j]<mid) i++;
            count += j-i;
            j--;
        }
        if (count >= k)
            left = mid+1;
        else 
            right = mid-1;
    }
    printf("%ld\n", right);

    return 0;
}
