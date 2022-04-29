package main
type Node struct {
	Val bool  // 不为叶子节点，有值 1-true 0-false
	IsLeaf bool  // 若有四个子节点，则为false
	TopLeft *Node
	TopRight *Node
	BottomLeft *Node
	BottomRight *Node
}

// func construct(grid [][]int) *Node {
//     var dfs func([][]int, int,int) *Node
// 	dfs = func(rows [][]int, c0, c1 int) *Node {
// 		for _,row := range rows{
// 			for _,v := range row[c0:c1] {
// 				if v != rows[0][c0] { // 不是叶子节点
// 					rMid,cMid := len(row)/2, (c0+c1)/2
// 					return &Node{
// 						Val: true,
// 						IsLeaf: false,
// 						TopLeft: dfs(rows[:rMid], c0,cMid),
// 						TopRight: dfs(rows[:rMid],cMid,c1),
// 						BottomLeft: dfs(rows[rMid:], c0,cMid),
// 						BottomRight: dfs(rows[rMid:], cMid,c1),
// 					}
// 				}
// 			}
// 		}
// 		// 叶子节点
// 		return &Node{Val: rows[0][c0]==1 , IsLeaf: true}
// 	}
// 	return dfs(grid,0,len(grid))
// }
func construct(grid [][]int) *Node {
    var dfs func([][]int, int, int) *Node
    dfs = func(rows [][]int, c0, c1 int) *Node {
        for _, row := range rows {
            for _, v := range row[c0:c1] {
                if v != rows[0][c0] { // 不是叶节点
                    rMid, cMid := len(rows)/2, (c0+c1)/2
                    return &Node{
                        true,
                        false,
                        dfs(rows[:rMid], c0, cMid),
                        dfs(rows[:rMid], cMid, c1),
                        dfs(rows[rMid:], c0, cMid),
                        dfs(rows[rMid:], cMid, c1),
                    }
                }
            }
        }
        // 是叶节点
        return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
    }
    return dfs(grid, 0, len(grid))
}
