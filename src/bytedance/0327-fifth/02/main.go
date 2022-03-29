package main

import "fmt"

func main() {
	var N, M, T int
	fmt.Scanln(&N, &M, &T)
	// 需要判断各种做法
	// 1+2+....+M = (M+1)*M/2 = s1       M+10--s1   工作效率s1/(M+10)
	// 1+2+...+M-1 = M(M-1)/2 = s2      M-1+5-s2    工作效率s2/(M+4)
	// (M-1)/2 ,, ,,, M-1
	// 最终会越来越接近0 
	//
	
	if M < 10 {
		// 休息时间比工作时间长
		// 最优的做法是工作到M-1，然后摸鱼
		// (1+M-1)*(M-1)/2 = s2     M-1+5= M+4 -- s2
		s2 := (M - 1) * M / 2
		left := N % s2
		cnt := 0
		for left != 0 {
			cnt++
			left -= cnt
		}
		fmt.Println(N/s2*(M+4) + cnt)
	} else {
		// M > 10
		// 工作时间比休息时间长
	}
}
