package lt470

import "math/rand"

/*
已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，
试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。

不要使用系统的 Math.random() 方法。


*/
func rand7() int {
    return rand.Intn(8)  //[0,n)
}

func rand10() int {
    a, b := rand7(), rand7()

    // Generates [1, 6]
    // Even and odd numbers have same pick possibilities
    for a == 7 {
        a = rand7()
    }
    // Generates [1, 5]
    for b > 5 {
        b = rand7()
    }

    // Since even and odd numbers have same pick possibilities for `a`
    // We have 50% chance to get [1, 5]
    if a & 1 != 0 {
        return b
    }
    // And the other 50% chance to get [5+1, 5+5], which is [6, 10]
    return 5 + b
}