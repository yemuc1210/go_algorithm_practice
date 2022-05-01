package main

import "fmt"

func maxConsecutiveAnswers1(answerKeys string, k int) int {
	// 最大化连续相同结果 结果的题数
	//最多操作k次
	// 贪心 dp 栈
	// 栈 TTFF  T2F2    TFFT  T1F2T1   TTFTTFTT   T2F1T2F1T2
	// 双指针
	// 滑动窗口：遇到相同的右指针++；遇到不同的，k-- 右指针++ k=0则左指针重新开始，从右指针-k
	maxLen := 1
	left, right := 0, 0
	n := len(answerKeys)
	cpyK := k
	for left < len(answerKeys) {
		cpyK = k
		for right < n {
			if answerKeys[right] == answerKeys[left] {
				right++
			} else {
				if cpyK > 0 {
					cpyK--
					right++
				} else {
					break
				}
			}
		}

		curLen := right - left
		if curLen > maxLen {
			maxLen = curLen
		}
		left = right
	}
	return maxLen
}

// 窗口内异种字符的个数，当大于k时，left++
func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	left, sum := 0, 0
	for right := range answerKey {
		if answerKey[right] != ch {
			sum++
		}
		for sum > k {
			if answerKey[left] != ch {
				sum--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'),
		maxConsecutiveChar(answerKey, k, 'F'))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	fmt.Println("vim-go")
	answerKeys := "TTFF"
	k := 2
	res := maxConsecutiveAnswers(answerKeys, k)
	fmt.Println(res)
}
