// package bilibili
package  main
import "fmt"

/*
题目描述：
描述
小A参加了一个n人的活动，每个人都有一个唯一编号i(i>=0 & i<n)，
其中m对相互认识，在活动中两个人可以通过互相都认识的一个人介绍认识。现在问活动结束后，小A最多会认识多少人？

并查集：
find
union
*/

//初始化
func makeSet(size int, fa *[]int) {
	for i:=0;i<size;i++{
		(*fa)[i] = i    //i就在其本身的集合里
	}
}
// func find(x int,fa []int ) int {
// 	//寻找x的祖先/根
// 	if fa[x] == x {
// 		//x是祖先
// 		return x
// 	}else{
// 		return find(fa[x], fa)   //查x的祖先fa[x]  向上一层一层查
// 	}
// }

//路径压缩  把路径上每个节点都直接连到根上
func find1(x int, fa []int) int {
	if x!=fa[x] {
		//初始化时  fa[i] = i  在处理过程中 fa[i]=i的是根节点
		fa[x] = find1(fa[x],fa)  //查找x的祖先知道找到根/代表节点，并做路径压缩
	}
	return fa[x]
}

//合并
func union(x,y int, fa []int) {
	x = find1(x,fa)  //找到根 合并
	y = find1(y,fa)
	fa[x] = y   //x的祖先变成y祖先的儿子
}
//启发式合并方法  按秩合并  
//不同的连接方法会得到不同的复杂度    将点数与深度较小的集合树合并到较大的更优


func main(){
	var n int   //聚会人数
	fmt.Scanln(&n)
	fa := make([]int, n)
	// for i:=0;i<n;i++{
	// 	fa[i] = i
	// }
	makeSet(n, &fa)	
	var ai int  //小A的编号
	fmt.Scanln(&ai)
	// fmt.Println(ai)
	var m int   //互相认识的数目
	fmt.Scanln(&m)

	//下面是互相认识的对，以','分割
	alreadyKnown := []int{}
	for i:=0;i<m;i++{
		var a,b int 
		fmt.Scanf("%d,%d",&a,&b)
		// fmt.Println(a,b)
		if a==ai || b==ai {
			if a==ai {
				alreadyKnown = append(alreadyKnown, b)
			}else{
				alreadyKnown = append(alreadyKnown, a)
			}
		}
		union(a,b,fa)
	}
	//得到最多认识的人  不过问题问的是新认识的
	an := find1(ai,fa)   //ai的祖先
	// fmt.Println(fa)
	//遍历fa得到
	cnt := 0
	for i:=0;i<len(fa);i++{
		if i != ai && fa[i] == an {
			cnt ++
		}
	}
	fmt.Println(cnt - len(alreadyKnown))
}