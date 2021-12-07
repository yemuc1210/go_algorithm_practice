/*
给定一个字符串 s 和一些 长度相同 的单词 words 。
找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，
但不需要考虑 words 中单词串联的顺序。
*/

package leetcode

/*
先找出所有单词的排列？这个复杂度。。
一个关键线索是单词长度相同
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。


这一题看似很难，但是有 2 个限定条件也导致这题不是特别难。
1. 字符串数组里面的字符串长度都是一样的。
2. 要求字符串数组中的字符串都要连续连在一起的，
前后顺序可以是任意排列组合。
解题思路，
先将字符串数组里面的所有字符串都存到 map 中，并累计出现的次数。
然后从源字符串从头开始扫，每次判断字符串数组里面的字符串是否全部都用完了(计数是否为 0)，
如果全部都用完了，并且长度正好是字符串数组任意排列组合的总长度，
就记录下这个组合的起始下标。如果不符合，就继续考察源字符串的下一个字符，
直到扫完整个源字符串。
*/
func findSubstring(s string, words []string) []int {
	if len(words)==0{
		return []int{}
	}
	res:=[]int{}
	counter_map := map[string]int{}
	for _,w := range words{
		counter_map[w]+=1
	}//得到字符计数
	//下面开始遍历字符串
	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter_map)
	for i,start:=0,0;i<len(s)-length+1 && start<len(s)-length+1;i++{
		if tmpCounter[s[i:i+length]] > 0 {  //得到字符计数
			tmpCounter[s[i:i+length]]--
			//fmt.Printf("******sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
			if checkWords(tmpCounter) && (i+length-start == totalLen) {
				res = append(res, start)
				continue
			}
			i = i + length - 1
		} else {
			start++
			i = start - 1
			tmpCounter = copyMap(counter_map)
		}
	}
	return res
}

func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {   //计数是否用完
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}


func findSubstring_1(s string, words []string) []int {
    wL := len(words)        // words 中单词个数
    oneWordL := len(words[0])   // 每个单词长度
    hm := make(map[string]int)  // 存words中单词的map
    Allcount := 0        // 记录不同单词的个数，用于在窗口滑动过程中判断是否满足答案条件
    for i:=0; i<wL; i++{
        if _ , ok := hm[words[i]]; !ok{
            Allcount++
        }
        hm[words[i]]++
    }

    ans := []int{}
    for i := 0; i < oneWordL; i++ {         // 这里参考各位大佬的思路，从0 - oneWordL的各位置开始向后滑
        Count := 0          // 和AllCount对应，这里记录滑动窗口中不同单词的个数
        tmpHm := map[string]int{}       // 滑动窗口中的单词记录map
        left , right := i , i + oneWordL    
        for  ; right <= len(s) ; right += oneWordL{
            str := s[right - oneWordL : right]
            if _ , ok := hm[str]; ok{
                tmpHm[str]++
                if tmpHm[str] == hm[str]{       // Count++的条件是该单词出现次数在窗口tmpHm中的出现次数和hm次数一样
                    Count++
                }
            }       // 如果这一步出现了不相关的单词怎么办？不影响，后面在比较Count和AllCount时，还会比较right - left == oneWordL * wL 。
            
            for right - left > oneWordL * wL{      
                str := s[left : left + oneWordL]
                if _ , ok := tmpHm[str] ; ok{
                    tmpHm[str]--
                }
                if tmpHm[str] + 1 == hm[str]{       // 这个条件重要，tmpHm[str] < hm[str]不对
                    Count--
                }
                left += oneWordL
            }

            if Count == Allcount && right - left == oneWordL * wL{      // 两个条件判断是否满足答案
                ans = append(ans , left)
            }  
        }
    }
    return ans
}
