package lt735

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

