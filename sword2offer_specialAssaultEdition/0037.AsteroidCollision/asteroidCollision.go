package offerS37

/*
剑指 Offer II 037. 小行星碰撞

给定一个整数数组 asteroids，表示在同一行的小行星。

对于数组中的每一个元素，其绝对值表示小行星的大小，
正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。
每一颗小行星以相同的速度移动。

找出碰撞后剩下的所有小行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。
如果两颗行星大小相同，则两颗行星都会爆炸。
两颗移动方向相同的行星，永远不会发生碰撞。

示例 3：

输入：asteroids = [10,2,-5]
输出：[10]
解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。

*/
func asteroidCollision(asteroids []int) []int {
    res := make([]int, 0)
    stack := make([]int, 0)

    for _, n := range(asteroids){
        if n > 0{   

            stack = append(stack, n)
        
        }else{
            for len(stack) > 0{
                end := len(stack) - 1
                t := n + stack[end]
                
                if t < 0{
                    stack = stack[:end]
                }else{
                    if t == 0{
                     stack = stack[:end]   
                    }
                    n = 0
                    break
                }
            }
            if n < 0{
                res = append(res, n)
            }
        }
    }
    return append(res, stack...)
}
