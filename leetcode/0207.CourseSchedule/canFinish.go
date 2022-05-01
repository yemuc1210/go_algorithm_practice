package lt207
/*中
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/
/*
图 dfs bfs 拓扑排序  AOV网
构建图 拓扑排序
循环执行，直到不存在入度为0的顶点：
	1. 选择一个入度为0的输出
	2. 删除该节点和所有出边（同时更新下一层节点的入度）
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int,numCourses)
	frees := make([][]int, numCourses)
	next := make([]int,0,numCourses)

	for _,v := range prerequisites { //[0,1] 要想学0 先学1 所以1指向0
		in[v[0]] ++ //入度++
		frees[v[1]] = append(frees[v[1]], v[0])  //节点的出边集合
	}
	for i:=0;i<numCourses;i++{
		if in[i]==0{
			next = append(next, i)  //找出入边为0的
		}
	}
	for i:=0;i!=len(next);i++{
		c := next[i]  //选择一个入度为0
		v := frees[c]  //入度为0的出边(点）集合 
		for _, v1 := range v{
			//删除出边 对应节点入度--
			in[v1] -- 
			if in[v1] ==0 {
				next = append(next, v1)  //更新入度为0的集合
			}
		}
	}
	return len(next) == numCourses  //是否都可以完成
}