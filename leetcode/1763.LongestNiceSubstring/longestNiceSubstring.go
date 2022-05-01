package lt1763

import (
	"math/bits"
	"unicode"
)

func longestNiceSubstring(s string) string {
	// 最长的美好子
	// 方法一  排序  没意思 略
	// 方法二  map
	m := make(map[byte]int)    // 字母-出现次数
	for i:= range s {
		m[s[i]]++
	}
	// foreach
	for k,_ := range m {
		// k- k'
		// 查看对应的另一个大小写字母在不在
		if isLower(k) {
			upper := k-32
			if _,exist := m[upper]; exist {

			}

		} 
	}
	return ""
}
func isLower(b byte) bool {
	return b>='a' && b<='z'
}
func isUpper(b byte) bool {
	return b>='A' && b<='Z'
}


func longestNiceSubstring1(s string) (ans string) {
    for i := range s {
        lower, upper := 0, 0  // 用二进制位表示  枚举全部的
        for j := i; j < len(s); j++ {
            if unicode.IsLower(rune(s[j])) {
                lower |= 1 << (s[j] - 'a')
            } else {
                upper |= 1 << (s[j] - 'A')
            }
            if lower == upper && j-i+1 > len(ans) {
                ans = s[i : j+1]
            }
        }
    }
    return
}


func longestNiceSubstring2(s string) (ans string) {
    mask := uint(0)
    for _, ch := range s {
        mask |= 1 << (unicode.ToLower(ch) - 'a')
    }
    maxTypeNum := bits.OnesCount(mask)

    for typeNum := 1; typeNum <= maxTypeNum; typeNum++ {
        var lowerCnt, upperCnt [26]int
        var total, cnt, l int
        for r, ch := range s {
            idx := unicode.ToLower(ch) - 'a'
            if unicode.IsLower(ch) {
                lowerCnt[idx]++
                if lowerCnt[idx] == 1 && upperCnt[idx] > 0 {
                    cnt++
                }
            } else {
                upperCnt[idx]++
                if upperCnt[idx] == 1 && lowerCnt[idx] > 0 {
                    cnt++
                }
            }
            if lowerCnt[idx]+upperCnt[idx] == 1 {
                total++
            }

            for total > typeNum {
                idx := unicode.ToLower(rune(s[l])) - 'a'
                if lowerCnt[idx]+upperCnt[idx] == 1 {
                    total--
                }
                if unicode.IsLower(rune(s[l])) {
                    lowerCnt[idx]--
                    if lowerCnt[idx] == 0 && upperCnt[idx] > 0 {
                        cnt--
                    }
                } else {
                    upperCnt[idx]--
                    if upperCnt[idx] == 0 && lowerCnt[idx] > 0 {
                        cnt--
                    }
                }
                l++
            }

            if cnt == typeNum && r-l+1 > len(ans) {
                ans = s[l : r+1]
            }
        }
    }
    return
}
