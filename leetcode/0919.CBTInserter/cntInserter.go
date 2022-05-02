package lt919
import "go_practice/structs"
type TreeNode = structs.TreeNode
//双端队列  完全二叉树插入     剑指offerII 43
type CBTInserter struct {
    p []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
    this := &CBTInserter{[]*TreeNode{root}}
    for i := 0; i < len(this.p); i++ {
        if this.p[i].Left != nil {
            this.p = append(this.p, this.p[i].Left)
        }
        if this.p[i].Right != nil {
            this.p = append(this.p, this.p[i].Right)
        }
    }
    return *this
}

func (this *CBTInserter) Insert(v int) int {
    n := len(this.p)
    father := this.p[(n - 1) >> 1]
    this.p = append(this.p, &TreeNode{Val: v})
    if n & 1 == 1 {
        father.Left = this.p[n]
    } else {
        father.Right = this.p[n]
    }
    return father.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
    return this.p[0]
}
