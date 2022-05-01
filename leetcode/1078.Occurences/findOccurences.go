package lt1078

import "strings"


func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text," ")
	//三个三个遍历
	if len(words) < 3 {
		return []string{}
	}
	s1,s2,s3 := words[0],words[1],words[2]
	res := []string{}
	if s1==first && s2==second {
		res = append(res, s3)
	}
	for i:=3;i<len(words);i++{
		s1,s2,s3 = s2,s3,words[i]
		if s1==first && s2==second {
			res = append(res, s3)
		}
	}
	return res
}