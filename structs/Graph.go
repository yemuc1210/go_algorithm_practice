package structs

//图的邻接表表示
type Node struct{
	Val int
	Neighbors []*Node
}
