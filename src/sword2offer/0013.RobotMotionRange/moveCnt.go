package offer13

// import "fmt"

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，
它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，
因为3+5+3+7=18。但它不能进入方格 [35, 38]，
因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：
输入：m = 2, n = 3, k = 1
输出：3
*/

func movingCount(m int, n int, k int) int {
	if k==0{
		return 1
	}
	res:=0
	visited := make([][]bool,m)
	for i:=0;i<len(visited);i++{
		visited[i]=make([]bool, n)
	}
	solveByDFS(m,n,&res,k,0,0,&visited)
	return res
}
func solveByDFS(m,n int, res *int,k int,cur_i,cur_j int,visited *[][]bool){
	if cur_i<0 || cur_i>m-1 || cur_j<0 || cur_j>n-1{
		return                //坐标越界
	}
	if (*visited)[cur_i][cur_j] {  //如果被访问过，也是直接返回，所以下面不考虑是否访问过
		return
	}
	(*visited)[cur_i][cur_j]=true   //不管怎么样，已经访问了
	if isMovable(cur_i,cur_j,m,n,k){
		// fmt.Println(cur_i,cur_j,*res)
		*res = *res + 1   
		// fmt.Println(*res)
		// solveByDFS(m,n,res,k,cur_i,cur_j-1,visited)     //左
	
		solveByDFS(m,n,res,k,cur_i,cur_j+1,visited)
	
		// solveByDFS(m,n,res,k,cur_i-1,cur_j,visited) //上
	
		solveByDFS(m,n,res,k,cur_i+1,cur_j,visited)
		
	}
}

//判断（i,j）方格能否进入
func isMovable(i,j int, m,n,k int)bool{
	//首先判断是否越界
	if i<0 || i>m-1 || j<0 || j>n-1{
		return false               //坐标越界
	}
	//分解数位
	return ((getBitSum(i)+getBitSum(j))<=k)
}
// func getBitSum(a int)int{
// 	if a>=0 && a<=9{
// 		return a
// 	}
// 	return (a%10)+getBitSum(a/10)
// }
//改为非递归，运行结果来看，效率没有多少提升
func getBitSum(a int)int{
	sum:=0
	for a !=0 {
		sum += a%10
		a = a/10
	}
	return sum 
}



/*
数位和增量公式： x的数位和是sx，x+1是s_x+1    逢进位突变一次
当（x+1）%10=0: s_x+1 = sx - 8         如 19  20
当（x+1）%10!=0      s_x+1 = sx+1         1   2
*/