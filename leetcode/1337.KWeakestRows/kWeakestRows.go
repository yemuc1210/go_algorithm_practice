package lt1337

import (
	"container/heap"
	"sort"
)

/*
矩阵 1表示军人，0表示平民  且1在每行最前面

*/
// func kWeakestRows(mat [][]int, k int) []int {
// 	res := []int{}

// 	//遍历每一行，得到1的个数
// 	sum := make(map[int]int,len(mat))  //key是值   val 是下标
// 	keys := []int{}
// 	for index,row := range mat{
// 		//row为一行的数组
// 		//求和
// 		tmpSum := 0
// 		for _,x := range row{
// 			if x!=0{
// 				tmpSum += x
// 			}else{
// 				break
// 			}
// 		}
// 		fmt.Printf("tmpSum=%v,index=%v\n",tmpSum,index)
// 		sum[tmpSum] = index
// 		keys = append(keys, tmpSum)
// 	}
// 	//从sum数组中找出最小的k个，需要排序
// 	sort.Ints(keys)
// 	fmt.Println(keys)
// 	for i:=0;i<k;i++ {
// 		res = append(res, sum[keys[i]])
// 	}
// 	return res
// }
func kWeakestRows(mat [][]int, k int) []int {
    h := hp{}
    for i, row := range mat {
        pow := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
        h = append(h, pair{pow, i})
    }
    heap.Init(&h)
    ans := make([]int, k)
    for i := range ans {
        ans[i] = heap.Pop(&h).(pair).idx
    }
    return ans
}

type pair struct{ pow, idx int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

