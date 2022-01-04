package main

import (
	"fmt"
	"strings"
)

/*html语法检查
标签闭合：双标签闭合    自闭合<div />
标签可以嵌套
	不可以交叉
多属性直接必须有空格
输入只会有a-z <> "=
return 1 0(没有错误)
*/

func main() {
	var line string
	fmt.Scanln(&line)

	//去除空格
	line = strings.Trim(line, " ")

	//判断是否符合语法
	//用栈
	type pair struct {
		tag   string //标签名称  不含尖括号的
		isEnd bool   //是否是闭合的标签
	}

	stack := []pair{}

	for i := 0; i < len(line); i++ {
		//处理一个标签
		var isEnd bool
		if line[i] == '<' {
			i++
		}
		if line[i] == '/' {
			isEnd = true //</  是闭合标签
			i++
		}
		//读取标签
		var tag string
		for i < len(line) && line[i] != '>' && line[i] != '/' {
			tag += string(line[i])
			i++
		}
		if i<len(line) && line[i] == '/' {
			//自闭合
			//跳转到下一个
			i += 2
			continue
		}
		//此时i应该指向一个>
		i++ //指向下一个<

		if isEnd {
			//当前的是闭合标签，和栈顶的比较
			top := stack[len(stack)-1]

			//如果相等，弹出
			if top.tag == tag {
				stack = stack[:len(stack)-1]
			} else {
				//错误

				fmt.Println(1)
				break
			}
		} else {
			//否则 压入栈
			stack = append(stack, pair{tag, isEnd})
		}
	}

	//判断栈是否空
	if len(stack) == 0 {
		fmt.Println(0)
	}
}
