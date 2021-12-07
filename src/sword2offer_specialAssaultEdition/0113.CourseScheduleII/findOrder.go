package offerS113

/*lt210   图
剑指 Offer II 113. 课程顺序
现在总共有 numCourses 门课需要选，记为 0 到 numCourses-1。

给定一个数组 prerequisites ，它的每一个元素 prerequisites[i] 表示两门课程之间的先修顺序。 例如 prerequisites[i] = [ai, bi] 表示想要学习课程 ai ，需要先完成课程 bi 。

请根据给出的总课程数  numCourses 和表示先修顺序的 prerequisites 得出一个可行的修课序列。

可能会有多个正确的顺序，只要任意返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	frees := make([][]int, numCourses)
	next := make([]int, 0, numCourses)
	for _, v := range prerequisites {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}
	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]
		for _, vv := range v {
			in[vv]--
			if in[vv] == 0 {
				next = append(next, vv)
			}
		}
	}
	if len(next) == numCourses {
		return next
	}
	return []int{}
}
