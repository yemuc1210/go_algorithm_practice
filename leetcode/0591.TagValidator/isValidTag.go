package main

import (
	"fmt"
	"regexp"
)

func isValid(code string) bool {
	re := regexp.MustCompile(`"<!\[CDATA\[.*?\]\]>|t"`)
	src := []byte(code)
	repl := []byte("-")
	src = re.ReplaceAll(src,repl)
	code = string(src)
	fmt.Println(code)
	var prev string
	ret := regexp.MustCompile(`"<([A-Z]{1,9})>[^<]*</\\1>"`)
	replt := []byte("t")
	for code != prev {
		prev = code
		code = string(ret.ReplaceAll([]byte(code), replt))
		fmt.Println(code)
	}
	return code == "t"
}

func main() {
	code :="<DIV>This is the first line <![CDATA[<div>]]></DIV>"
	res := isValid(code)
	fmt.Println(res)
}