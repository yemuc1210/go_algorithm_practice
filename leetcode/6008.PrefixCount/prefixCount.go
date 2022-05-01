package main

import "fmt"

func prefixCount(words []string, pref string) int {
    // hasPrefix
    var res int
    for _,v:= range words {
        if hasPrefix(v, pref) {
            res++
        }
    }
    return res
}
func hasPrefix(s string, prefix string) bool {
    if len(s) < len(prefix) {
        return false
    }
    return s[:len(prefix)] == prefix
}

func main() {
	words := []string{"pay","attention","practice","attend"}
	pref := "at"
	res := prefixCount(words,pref)
	fmt.Println(res)
}