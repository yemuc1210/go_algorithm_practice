package lt911
/*
911. 在线选举
给你两个整数数组 persons 和 times 。
在选举中，第 i 张票是在时刻为 times[i] 时投给候选人 persons[i] 的。

对于发生在时刻 t 的每个查询，需要找出在 t 时刻在选举中 领先的候选人的编号。

在 t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。

实现 TopVotedCandidate 类：

TopVotedCandidate(int[] persons, int[] times) 使用 persons 和 times 数组初始化对象。
int q(int t) 根据前面描述的规则，返回在时刻 t 在选举中领先的候选人的编号。
*/


import (
	"sort"
)

// TopVotedCandidate define
type TopVotedCandidate struct {
	persons []int
	times   []int
}

// Constructor911 define
func Constructor911(persons []int, times []int) TopVotedCandidate {
	leaders, votes := make([]int, len(persons)), make([]int, len(persons))
	leader := persons[0]
	for i := 0; i < len(persons); i++ {
		p := persons[i]
		votes[p]++
		if votes[p] >= votes[leader] {
			leader = p
		}
		leaders[i] = leader
	}
	return TopVotedCandidate{persons: leaders, times: times}
}

// Q define
func (tvc *TopVotedCandidate) Q(t int) int {
	i := sort.Search(len(tvc.times), func(p int) bool { return tvc.times[p] > t })
	return tvc.persons[i-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
