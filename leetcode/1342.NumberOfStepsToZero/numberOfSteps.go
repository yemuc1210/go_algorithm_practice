package lt1342

import "fmt"

func numberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
			steps ++
		}else{
			num -= 1
			steps ++
		}
		fmt.Println(num, steps)
	}
	fmt.Println(num,steps)
	return steps
}