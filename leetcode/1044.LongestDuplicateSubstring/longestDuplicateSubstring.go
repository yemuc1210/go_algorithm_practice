package lt1044

import (
	"index/suffixarray"
	"reflect"
	"unsafe"
)

/*
1044. 最长重复子串
给你一个字符串 s ，考虑其所有 重复子串 ：即，s 的连续子串，
在 s 中出现 2 次或更多次。这些出现之间可能存在重叠。

返回 任意一个 可能具有最长长度的重复子串。如果 s 不含重复子串，那么答案为 "" 。
*/
/*
后缀数组   后缀树
设字符串为S(1-n)由n个字符组成。则字符串有n个相同后缀的子串。分别为s(1-n),s(2-n),…,s(n,n)
构建一个SA数组，每个数组存储这些后缀的子串，存储后进行字典序排序
构建height数组，表示SA数组每个元素和前一个元素相同前缀的字符个数

那么，最长重复子串的长度就是height数组的最大值。
因为最长重复子串一定是两个不同后缀的公共前缀，而且这两个不同后缀的字典序排列后一定是相连的，
否则一定有比他更长的
所以height的最大值能够找到那两个后缀，然后提取公共前缀找到答案
*/
func longestDupSubstring(s string) string {
	// 求出后缀数组和高度数组
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}
	
	// 高度数组中的最大值对应的就是最长的重复子串
	idx, maxH := 0, 0
	for i, h := range height {
		if h > maxH {
			idx, maxH = i, h
		}
	}
	return s[sa[idx] : int(sa[idx])+maxH]
}



// class Solution {
//     int N = 30010;
//     int[] x = new int[N], y = new int[N], c = new int[N];
//     int[] sa = new int[N], rk = new int[N], height = new int[N];
//     void swap(int[] a, int[] b) {
//         int n = a.length;
//         int[] c = a.clone();
//         for (int i = 0; i < n; i++) a[i] = b[i];
//         for (int i = 0; i < n; i++) b[i] = c[i];
//     }
//     public String longestDupSubstring(String s) {
//         int n = s.length(), m = 128;
//         s = " " + s;
//         char[] cs = s.toCharArray();
//         // sa：两遍基数排序，板子背不下来也可以直接使用 sort，复杂度退化到 n \log^2 n
//         for (int i = 1; i <= n; i++) {
//             x[i] = cs[i]; c[x[i]]++;
//         }
//         for (int i = 2; i <= m; i++) c[i] += c[i - 1];
//         for (int i = n; i > 0; i--) sa[c[x[i]]--] = i;
//         for (int k = 1; k <= n; k <<= 1) {
//             int cur = 0;
//             for (int i = n - k + 1; i <= n; i++) y[++cur] = i;
//             for (int i = 1; i <= n; i ++) {
//                 if (sa[i] > k) y[++cur] = sa[i] - k;
//             }
//             for (int i = 1; i <= m; i++) c[i] = 0;
//             for (int i = 1; i <= n; i++) c[x[i]]++;
//             for (int i = 2; i <= m; i++) c[i] += c[i - 1];
//             for (int i = n; i > 0; i--) {
//                 sa[c[x[y[i]]]--] = y[i];
//                 y[i] = 0;
//             }
//             swap(x, y);
//             x[sa[1]] = cur = 1;
//             for (int i = 2; i <= n; i++) {
//                 if (y[sa[i]] == y[sa[i - 1]] && y[sa[i] + k] == y[sa[i - 1] + k]) x[sa[i]] = cur;
//                 else x[sa[i]] = ++cur;
//             }
//             if (cur == n) break;
//             m = cur;
//         }
//         // height
//         int k = 0;
//         for (int i = 1; i <= n; i ++) rk[sa[i]] = i;
//         for (int i = 1; i <= n; i++) {
//             if (rk[i] == 1) continue;
//             int j = sa[rk[i] - 1];
//             k = k > 0 ? k - 1 : k;
//             while (i + k <= n && j + k <= n && cs[i + k] == cs[j + k]) k++;
//             height[rk[i]] = k;
//         }
//         int idx = -1, max = 0;
//         for (int i = 1; i <= n; i++) {
//             if (height[i] > max) {
//                 max = height[i];
//                 idx = sa[i];
//             }
//         }
//         return max == 0 ? "" : s.substring(idx, idx + max);
//     }
// }
// 时间复杂度：O(n\log{n})O(nlogn)
// 空间复杂度：O(n)O(n)
// 下一篇：O(n) 复杂度的后缀数组解法 耗时 28 ms
// © 著作权归作者所有
