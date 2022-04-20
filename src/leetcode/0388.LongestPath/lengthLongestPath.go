package main

import (
	"fmt"

)

/*
input = "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
一个\t表示时一个次级路径
返回绝对路径最长的文件的路径长度，没有文件返回0
文件名规范：name.extension

直接找到 '.' 即可判断有几个文件
然后解析每个文件的路径
"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
得到文件 file1.ext   file2.ext
如何解析文件路径
如何存储多级路径?
map[int]map[string]int
 1: 存放一级路径名    v-map[string]int
 2: 存放二级路径名

 哈希表记录每个层级最新的文件夹路径，通过字符串拼接得到cur的完整路径path
*/

func lengthLongestPath(input string) int {

	m := make(map[int]string)
	n := len(input)
	var res string
	var i int
	for i=0;i<n; {
		level := 0 
		for i<n && input[i]=='\t' && level+1 >=0 {
			level++
			i++
		}
		// 得到具体的level
		j := i 
		var isDir bool = true
		for j<n && input[j]!='\n' {
			if input[j] == '.' {
				isDir = false
			}
			j++
		}
		// 发现文件，
		var cur string = input[i:j]   // 此时j指向 \n   i是第一个字符
		var pre string
		var ok bool
		pre,ok = m[level-1]
		if !ok {pre= ""}

		var path string 
		if pre=="" {
			path = cur
		}else{
			path = pre+"/"+cur
		}
		if isDir {
			m[level]=path
		}else if (len(path)>len(res)) {
			res = path
		}
		i = j+1   // 指向下一个路径首字母
	}
	return len(res)
}

func main() {
	fmt.Println("vim-go")
	input := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Println(lengthLongestPath(input))
}
