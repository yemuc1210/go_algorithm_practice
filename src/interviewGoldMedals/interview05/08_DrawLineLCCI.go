package interview05

// 绘图（x1,y) (x2,y) 像素点为1
// 每行像素点 01序列 -> length
// w: 屏幕宽度   w/32表示一行有多少int至
// length/(w/32)表示一共有多少行
func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	var cnt uint = 0x80000000
	var a uint = 0 
	arr := []int{}
	// y行之前的int加入数组
	for i:=0;i<y;i++{
		n := w/32  // 一行有多少搁int
		for n!=0 {
			arr = append(arr, 0)
			n--
		}
	}
	// y行的x1到x2的0置为1
	for j:=0;j<=w;j++{  // 列索引习惯用j
		if j>=x1 && j<=x2 {
			// 置为1
			a = a | cnt
		}
		cnt = cnt>>1  //右移一位
		if (j+1)%32==0 {
			// 一个置换完成的int
			arr = append(arr, int(a))
			a = 0  // 恢复初值
			cnt = 0x80000000
		}
	}
	// 剩下的int加入数组
	if length>w/32*(y+1) {
		m := length - w/32*(y+1)
		for i:=0;i<m;i++{
			arr = append(arr, 0)
		}
	}
	return arr
}
