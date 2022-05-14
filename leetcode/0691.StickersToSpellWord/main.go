package main

// 691. 贴纸拼词
// n种贴纸，每个贴纸上都有一个小写英文单词
// 拼出target，从贴纸上切割单个字母并重新排列。
// 贴纸数量无限  可以重复使用
// 拼出target的最小贴纸数量
// 若任务不可能，但会-1

func minStickers(stickers []string, target string) int {
    m := len(target)
    f := make([]int, 1<<m)
    for i := range f {
        f[i] = -1
    }
    f[0] = 0
    var dp func(int) int
    dp = func(mask int) int {
        if f[mask] != -1 {
            return f[mask]
        }
        f[mask] = m + 1
        for _, sticker := range stickers {
            left := mask
            cnt := [26]int{}
            for _, c := range sticker {
                cnt[c-'a']++
            }
            for i, c := range target {
                if mask>>i&1 == 1 && cnt[c-'a'] > 0 {
                    cnt[c-'a']--
                    left ^= 1 << i
                }
            }
            if left < mask {
                f[mask] = min(f[mask], dp(left)+1)
            }
        }
        return f[mask]
    }
    ans := dp(1<<m - 1)
    if ans <= m {
        return ans
    }
    return -1
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
