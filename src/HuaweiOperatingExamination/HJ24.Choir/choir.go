package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//计算最少出列多少位同学 使得剩下同学排成合唱队形
//身高升序   最长增序子序列 lt300
func main() {
	reader := bufio.NewScanner(os.Stdin)
	var N int
	fmt.Scanln(&N)
	var height = make([]string, N)
	// var title = 0
	//输入
	for {
		reader.Scan()
		s := reader.Text()
		if s == "" {
			break
		}
		height = strings.Split(s," ")
		
		numList := []int{}
		for _,v := range height {
			num,_ := strconv.Atoi(v)
			numList = append(numList, num)
		}
		maxNum := lengthOfLIS(numList)
		fmt.Println(len(numList)-maxNum)
	}
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := []int{}
	maxDP := 0
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxDP {
			maxDP = dp[i]
		}
	}
	return maxDP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
