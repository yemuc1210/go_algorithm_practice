package lt1020
func numEnclaves(grid [][]int) int {
    var dirs = [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    m,n := len(grid),len(grid[0])

    var dfs func(i,j int) 
    dfs = func(i,j int) {
        // 边界
        if i<0 || i>=m || j<0 || j>=n || grid[i][j]==0 || grid[i][j]==2{
            return 
        }
        // 访问
        grid[i][j] = 2
        for _,dir := range dirs {
            x,y := i+dir[0],j+dir[1]
            dfs(x,y)
        }
    }
    //  从边界访问
    for i:=range grid {
        dfs(i,0)
        dfs(i,n-1)
    }
    for j:=1;j<n-1;j++{
        dfs(0,j)
        dfs(m-1,j)
    }
    // 统计
    var res int
    for i:=1;i<m;i++{
        for j:=1;j<n-1;j++{
            // 判断是否访问过
            if grid[i][j] ==1 {
                res ++
            }
        }
    }
    return res
}