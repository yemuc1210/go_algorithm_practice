package lt913

/*
轮流行动  
无向图graph[a]是列表   即graph[a]=b   ab是一条边  图的邻接表
老鼠从节点1开始，猫从节点2开始，节点0为洞
沿着边行动 ，猫无法移动到节点0
结束条件：
（1）猫win：相遇  return-2
（2）鼠win: arrive 0  return-1
（3）某一位置重复出现（位置和移动顺序和上一次相同） 平局  return-0
	这只能出现在猫鼠皆在【一个/两个节点】的连通分量中

视为状态方程  [mouse,cat,turn]
猫赢得状态[i,i,1]/[i,i,2]  1-表示老鼠走  2-表示猫走
老鼠赢【0，i，1/2】


用0代表未知状态，1代表老鼠赢，2代表猫赢
由最终的胜利状态，回推
假如当前父节点轮次是1（父节点轮次是2同样的道理）
父节点=1 if 子节点是1
或者
父节点=2 if 所有子节点是2
*/
const (
    draw     = 0
    mouseWin = 1
    catWin   = 2
)

func catMouseGame(graph [][]int) int {
    n := len(graph)
    dp := make([][][]int, n)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n*2)
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }

    var getResult, getNextResult func(int, int, int) int
    getResult = func(mouse, cat, turns int) int {
        if turns == n*2 {
            return draw
        }
        res := dp[mouse][cat][turns]
        if res != -1 {
            return res
        }
        if mouse == 0 {
            res = mouseWin
        } else if cat == mouse {
            res = catWin
        } else {
            res = getNextResult(mouse, cat, turns)
        }
        dp[mouse][cat][turns] = res
        return res
    }
    getNextResult = func(mouse, cat, turns int) int {
        curMove := mouse
        if turns%2 == 1 {
            curMove = cat
        }
        defaultRes := mouseWin
        if curMove == mouse {
            defaultRes = catWin
        }
        res := defaultRes
        for _, next := range graph[curMove] {
            if curMove == cat && next == 0 {
                continue
            }
            nextMouse, nextCat := mouse, cat
            if curMove == mouse {
                nextMouse = next
            } else if curMove == cat {
                nextCat = next
            }
            nextRes := getResult(nextMouse, nextCat, turns+1)
            if nextRes != defaultRes {
                res = nextRes
                if res != draw {
                    break
                }
            }
        }
        return res
    }
    return getResult(1, 2, 0)
}
