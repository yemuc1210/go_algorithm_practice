package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	const maxWidth = 100
	lines, width := 1, 0
	for _, c := range s {
		need := widths[c-'a']
		width += need
		if width > maxWidth {
			lines++
			width = need
		}
	}
	return []int{lines, width}
}
func main() {
	fmt.Println("vim-go")
}
