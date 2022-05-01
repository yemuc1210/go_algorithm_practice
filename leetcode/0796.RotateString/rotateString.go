package main

import "strings"

// s经过旋转变为goal  若能返回true
func rotateString(s string, goal string) bool {
	// 滑动窗口的比较思路，goal固定，s滑动   大概可行
	// 将字符串扩展两倍 s->2*s   然后判断goal是否在2*s里面
		// python: return len(s) == len(goal) && goal in s+s
	if len(s) != len(goal) {
		return false
	}
	// s+s 包含了所有通过旋转可以得到的字符串
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

// 手写KMP
// s-source p-pattern 判断p是否是s的子串
func KMP(s string, p string) bool {
	search(s,p)
	return true
}
func buildNext(p string) []int {
    next := []int{}
    next = append(next, 0)  //next[0] = 0
    i := 1
    now := 0  //next[x-1]
    
    //now 对应于next数组
    // i对应于p串
    
    for i < len(p) {
        if p[now] == p[i] {
            // 可以向后扩展一位
            now += 1
            i++
            next = append(next, now)
        }else if now!=0 {
            // 缩小now
            now = next[now-1] 
        }else{
            //now==0
            // 无法再缩小
            next = append(next, 0)
            i+=1
        }
    } 
    return next
}

func search(s,p string) {
    tar := 0  // 主串s中将要匹配的位置
    pos := 0 // 模式串中将要匹配的位置
    next := buildNext(p)
    for tar < len(s) {
        if s[tar] == p[pos] {
            tar++
            pos++
        }else if pos==0 {  // 失配，pos!=0
            //根据next移动
            pos = next[pos-1]
        }else{
            // 失配 且 pos==0
            tar +=1
        }
        if pos == len(p) {
            // 匹配成功
            println(tar-pos+1)  //内置函数
            pos = next[pos-1]
        }
    }
}