package lt1743
/*
存在一个n个元素的nums数组，但已记不清具体内容，只记得每一对相邻元素

给定一个adjacentPairs，大小为n-1，每个adjacentPairs[i]=[ui,vi],表示相邻
相邻元素可以是任意顺序的，【nums[i],nums[i+1]】或者是【nums[i+1],nums[i]】

返回原始数组

adjacentPairs:  (n-1) * 2

先构建一个长度为 2*n的  链表？   在左右添加，最后遍历，如果连续两个相同，则去除重复的一个

哈希表法，
	如果只有一个相邻的，则是头或尾。找到头尾则好办了
	使用哈希表白哦记录相邻元素，方便查找


*/
func restoreArray(adjacentPairs [][]int)[]int{
	n := len(adjacentPairs) + 1
	m := make(map[int][]int,n)
	for _,v := range adjacentPairs {
		v,w := v[0],v[1]
		m[v] = append(m[v], w)
		m[w] = append(m[v], v)
	}
	res := make([]int,n)
	for e,adj := range m{
		if len(adj) == 1{
			// front  rear
			res[0] = e
			break
		}
	}
	res[1] = m[res[0]][0]
	//继续
	for i:=2;i<n;i++{
		adj := m[res[i-1]]   //是一个数组
		if res[i-2] == adj[0] {
			res[i] = adj[1]
		}else{
			res[i] = adj[0]
		}
	}
	return res
}
func restoreArray1(adjacentPairs [][]int)[]int{
	n := len(adjacentPairs) + 1
    g := make(map[int][]int, n)
    for _, p := range adjacentPairs {
        v, w := p[0], p[1]
        g[v] = append(g[v], w)
        g[w] = append(g[w], v)
    }

    ans := make([]int, n)
    for e, adj := range g {
        if len(adj) == 1 {
            ans[0] = e
            break
        }
    }

    ans[1] = g[ans[0]][0]
    for i := 2; i < n; i++ {
        adj := g[ans[i-1]]
        if ans[i-2] == adj[0] {
            ans[i] = adj[1]
        } else {
            ans[i] = adj[0]
        }
    }
    return ans

}