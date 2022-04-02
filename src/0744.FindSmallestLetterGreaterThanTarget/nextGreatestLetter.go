package main

import (
	"fmt"
	"sort"
)

type params struct {
	letters []byte
	target  byte
}
type output struct {
	res byte
}

type prob744 struct {
	params
	output
}

func main() {
	fmt.Println("vim-go")
	probs := []prob744{
		{
			params: params{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			output: output{'c'},
		},
		{params: params{letters: []byte{'c', 'f', 'j'}, target: 'c'},output: output{'f'}},
		{params: params{letters: []byte{'c', 'f', 'j'}, target: 'd'},output: output{'f'}},
	}

	for _, v := range probs {
		p1, p2 := v.params.letters, v.params.target
		ans := v.output.res
		res := nextGreatestLetter(p1, p2)
		fmt.Println(string(p1),string(p2),string(res),string(ans))
		if res != ans {
			fmt.Println("error res", string(res))
		} else {
			fmt.Println("pass")
		}
	}
}

func nextGreatestLetter(letters []byte, target byte) byte {
    if target >= letters[len(letters)-1] {
        return letters[0]
    }
    i := sort.Search(len(letters)-1, func(i int) bool { return letters[i] > target })
    return letters[i]
}
// class Solution {
//     public char nextGreatestLetter(char[] letters, char target) {
//         int n = letters.length;
//         int l = 0, r = n - 1;
//         while (l < r) {
//             int mid = l + r >> 1;
//             if (letters[mid] > target) r = mid;
//             else l = mid + 1;
//         }
//         return letters[r] > target ? letters[r] : letters[0];
//     }
// }
 