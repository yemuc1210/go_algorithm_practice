package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//坐标移动
func main(){
	reader := bufio.NewScanner(os.Stdin)
	dirs := make(map[byte]struct{x,y int},4)
	dirs['A'] = struct{x int; y int}{-1,0}
	dirs['D'] = struct{x int; y int}{1,0}
	dirs['W'] = struct{x int; y int}{0,1} //y轴
	dirs['S'] = struct{x int; y int}{0,-1}
	type point struct{x,y int}
	start := point{0,0}
	if reader.Scan() {  //输入一行字符串
		input := reader.Text()
		input = strings.TrimSpace(input)
		s := strings.Split(input,";")
		// fmt.Println(s,len(s))
		j:=0
		next: for j<len(s) {
			//得到一组  
			e := s[j]
			// 可能这组数前面是合法的 但是后面非法 如A1A 所以暂存这组输入的改变量
			tmp := start
			i:=0
			for i<len(e) {
				dir,exist := dirs[e[i]]
				if !exist {
					//非法字符
					j++
					goto next
				}
				// fmt.Println(dir)
				//剩下的一位或两位是数字
				num := 0
				i++
				for i<len(e) && (e[i]>='0' && e[i]<='9') {
					num = num*10 + int(e[i]-'0')
					i++
				}
				// fmt.Println(num)
				//数字最多两位
				if num == 0 {
					j++
					goto next
				}
				if num > 99{
					//非法数字  跳出本组e  下一组数
					j++
					goto next
				}
				//得到移动方向和移动数  更新坐标
				tmp.x += dir.x*num
				tmp.y += dir.y*num
				// fmt.Println(start)
			}
			//如果整个都合法，会进行到这里
			if i==len(e){
				start = tmp
			}
			j ++
		}
		//输出坐标
		fmt.Printf("%d,%d",start.x,start.y)
	}
}