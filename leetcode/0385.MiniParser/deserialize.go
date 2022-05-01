package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() (b bool) {return }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() (a int) {return }

// Set this NestedInteger to hold a single integer.
func (n NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() (c []NestedInteger) {return }

func deserialize(s string) *NestedInteger {
	// s- 整数嵌套列表，实现语法分析
	// 栈模拟
	// dfs 【--待解析为列表  遇到',' 表示列表存在多个元素，继续     】----说明列表解析完毕
	//      包含整数
	if s[0] != '[' {
		num,_ := strconv.Atoi(s)
		ni := &NestedInteger{}
		ni.SetInteger(num)
		return ni
	}
	sk := []*NestedInteger{}
	num,negative := 0,false
	for i,ch := range s {
		if ch == '-' {
			// 负数
			negative = true
		}else if unicode.IsDigit(ch) {
			num = num*10 + int(ch - '0')
		}else if ch =='[' {
			// 列表
			sk = append(sk, &NestedInteger{})
		} else if ch==',' || ch==']' {
			// 处理完一段  
			// 数字的处理
			if unicode.IsDigit(rune(s[i-1])) {
				if negative {
					num = -num
				}
				ni := NestedInteger{}
				ni.SetInteger(num)
				sk[len(sk)-1].Add(ni)
			}  
			num,negative = 0,false
			// NestInteger的处理
			if ch==']' && len(sk)>1 {
				// 得到一段
				sk[len(sk)-2].Add(*sk[len(sk)-1])
				sk = sk[:len(sk)-1]
			}
		}

	}
	return sk[len(sk)-1]
}

func main() {
	fmt.Println("vim-go")
}
