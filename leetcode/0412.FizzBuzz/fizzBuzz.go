package lt412

import "strconv"

/*
412. Fizz Buzz
写一个程序，输出从 1 到 n 数字的字符串表示。

1. 如果 n 是3的倍数，输出“Fizz”；

2. 如果 n 是5的倍数，输出“Buzz”；

3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
n = 15,

返回:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]

*/
func fizzBuzz(n int) []string {
	res := []string{}
	for i:=1;i<=n;i++{
		if i%3 == 0 && i%5 != 0 {
			res = append(res, "Fizz")
		}else if i%5 == 0 && i%3 != 0{
			res = append(res, "Buzz")
		}else if i%5 ==0 && i%3 == 0{
			res = append(res, "FizzBuzz")
		}else{
			res = append(res, strconv.Itoa(i))
		}
	}
	return res

}