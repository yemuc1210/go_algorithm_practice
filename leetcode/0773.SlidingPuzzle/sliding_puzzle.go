/*
在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示.
一次移动定义为选择 0 与一个相邻的数字（上下左右）进行交换.
最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
给出一个谜板的初始状态，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。


十五谜问题。。。。
分支限界搜索
*/
package leetcode

import (
	"fmt"
	"strings"
)


type LiveListNode struct {
	val  [][]int
	next *LiveListNode
}
type LiveList struct {
	head   *LiveListNode
	rear   *LiveListNode
	length int
}

func (l *LiveList) add(val [][]int) {
	//创建
	add_arr := make([][]int, len(val)+1)
	for i := 0; i < len(val); i++ {
		add_arr[i] = make([]int, len(val)+1)
		for j:=0;j<len(val);j++{
			add_arr[i][j] =  val[i][j]
		}
	}
	add_node := LiveListNode{
		val: add_arr,
	}

	add_node.next = l.rear.next
	l.rear.next = &add_node
}
func (l *LiveList) printList(){
	p := l.head
	for p!=nil {
		fmt.Println(p.val)
		p = p.next
	}
	
}
func slidingPuzzle(board [][]int) int {
	res := 0
	n := len(board) //长宽尺寸

	return res+n

}

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func slidingPuzzle_1(board [][]int) int {
    const target = "123450"

    s := make([]byte, 0, 6)
    for _, r := range board {
        for _, v := range r {
            s = append(s, '0'+byte(v))
        }
    }
    start := string(s)
    if start == target {
        return 0
    }

    // 枚举 status 通过一次交换操作得到的状态
    get := func(status string) (ret []string) {
        s := []byte(status)
        x := strings.Index(status, "0")
        for _, y := range neighbors[x] {
            s[x], s[y] = s[y], s[x]
            ret = append(ret, string(s))
            s[x], s[y] = s[y], s[x]
        }
        return
    }

    type pair struct {
        status string
        step   int
    }
    q := []pair{{start, 0}}
    seen := map[string]bool{start: true}
    for len(q) > 0 {
        p := q[0]
        q = q[1:]
        for _, nxt := range get(p.status) {
            if !seen[nxt] {
                if nxt == target {
                    return p.step + 1
                }
                seen[nxt] = true
                q = append(q, pair{nxt, p.step + 1})
            }
        }
    }
    return -1
}
