class Solution:
    def nearestPalindromic(self, n: str) -> str:
        if int(n)<10 or int(n[::-1])==1:
            return str(int(n)-1)
        if n=='11':
            return '9'
        if set(n)=={'9'}:
            return str(int(n)+2)
        a,b=n[:(len(n)+1)//2],n[(len(n)+1)//2:]
        temp=[str(int(a)-1),a,str(int(a)+1)]
        temp=[i+i[len(b)-1::-1] for i in temp]
        return min(temp,key=lambda x:abs(int(x)-int(n)) or float('inf'))
# 总结下特殊情况：

# 小于等于10的数，返回n-1
# 10的幂，返回n-1
# 若干个9，返回n+2
# 11，这个数字比较特殊，返回9
# 接下来是普通情况怎么求解

# 首先把n从中间分成a、b两部分，如果长度是奇数就给a多分点。

# 然后用a、a+1、a-1为左边分别构建一个回文数，注意n长度为奇数的时候左边的最后一个字符不能复制过去。

# 最后选取离n最近且不为n的结果即可m