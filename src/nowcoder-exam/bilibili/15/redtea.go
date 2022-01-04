package main

import (
	// "bufio"
	"fmt"
	// "os"
	"sort"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d%c", &nums[i])
	}
	var need int
	fmt.Scanln(&need)

	sort.Ints(nums)

	var flag bool
	l, r := 0, n-1
	for l < r {
		if nums[l]+nums[r] > need {
			r--
		} else if nums[l]+nums[r] < need {
			l++
		} else {
			fmt.Printf("%d %d\n", nums[l], nums[r])
			l++
			flag = true
		}
	}
	if !flag {
		fmt.Println("NO")
	}
}
