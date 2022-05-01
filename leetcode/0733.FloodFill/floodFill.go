package lt733


/*
二维数组表示图像   值表示像素值大小
先给定一个坐标，以及颜色值，将像素值与起始坐标相同的点都渲染成新的颜色

寻址问题  BFS  DFS

*/
var (
    dx = []int{1, 0, 0, -1}
    dy = []int{0, 1, -1, 0}
)//四个方向 delta变量
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	//DFS
	curColor := image[sr][sc]
	if curColor != newColor {
		dfs(image, sr,sc, curColor,newColor)
	}
	return image
}

func dfs(image [][]int,  x,y,color,newColor int){
	if image[x][y] == color {
		image[x][y] = newColor
		//上下左右遍历
		for i:=0;i<4;i++{
			mx,my := x+dx[i], y+dy[i]
			if mx>=0 && mx <len(image) && my>=0 && my<len(image[0]) {
				dfs(image, mx,my,color,newColor)
			}
		}
	}
}