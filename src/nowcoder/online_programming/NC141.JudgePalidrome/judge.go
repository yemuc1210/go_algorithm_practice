package main

// import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 判断回文
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
func judge( str string ) bool {
    if len(str) == 0 {
		return false
	}
	if len(str) == 1 {
		return true
	}
	n := len(str)
	chs := []byte(str)
	//判断奇偶性
	if n%2!=0 {
		//奇数
		//删去中间 
		chs = append(chs[:n/2],  chs[n/2+1:]... )
	}
	// fmt.Println(chs)
	n=len(chs)
	for i:=0;i<n/2;i++{
		if chs[i] != chs[n-i-1] {
			return false
		}
	}
	return true
}

func main(){
	str := "acb"
	judge(str)
}