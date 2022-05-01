package main
import (
    "bufio"
    "os"
	"fmt"
)
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		res := input.Text()
		for len(res) >= 8 {
			fmt.Println(res[:8])
			res = res[8:]
		}
		if len(res) == 0 {
			continue
		}
		fmt.Print(res)
		for i := 0; i < 8-len(res); i++ {
			fmt.Print("0")
		}
		fmt.Println()
	}
}
