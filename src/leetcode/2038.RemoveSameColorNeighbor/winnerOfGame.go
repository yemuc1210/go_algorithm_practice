package main

import "fmt"

func winnerOfGame(colors string) bool {
	var sum int
	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == 'A' && colors[i-1] == 'A' && colors[i+1] == 'A' {
			sum++
		}
		if colors[i] == 'B' && colors[i-1] == 'B' && colors[i+1] == 'B' {
			sum--
		}
	}
	return sum > 0
}

func main() {
	colors := "AAABABB"
	res := winnerOfGame(colors)
	fmt.Println(res)
}
