package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//找出第k个仅出现一次的字
//计数用set
func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		line := reader.Text()

		k, err := strconv.Atoi(string(line[0]))
		if err != nil {
			panic(err)
		}
		s := line[2:] //输入的字符串

		m := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			m[s[i]]++
		}

		for i := 0; i < len(s); i++ {
			cnt := m[s[i]]
			if cnt == 1 {
				k--
				if k == 0 {
					fmt.Printf("[%v]\n", string(s[i]))
					goto next
				}
			}
			
		}
		if k > 0 {
			fmt.Println("Myon~")
			goto next
		}
	next:
	}
}


// import java.io.BufferedReader;
// import java.io.InputStreamReader;

// public class Main {
//     public static void main(String[] args) throws Exception {
//         BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
//         String str;
//         while ((str = br.readLine()) != null) {
//             String num = str.split(" ")[0];
//             int k = Integer.parseInt(num);
//             String s = str.substring(num.length() + 1);
//             System.out.println(appearFirstK(k, s));
//         }
//     }

//     private static String appearFirstK(int k, String s) {
//         char[] chars = s.toCharArray();
//         int[] hash = new int[200];
//         for (char c : chars) {
//             hash[c]++;
//         }
//         for (char c : chars) {
//             if (hash[c] == 1 && k == 1) {
//                 return "[" + c + "]";
//             } else if (hash[c] == 1) {
//                 k--;
//             }
//         }
//         return "Myon~";
//     }
// }