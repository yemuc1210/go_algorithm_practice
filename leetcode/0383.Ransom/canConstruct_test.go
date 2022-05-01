package lt0383

import (
	"fmt"
	"testing"
)

type question struct{
	ransom,magazine string
	ans bool
}
func TestProm(t *testing.T){
	qs:=[]question{
		{
			ransom: "a",
			magazine: "b",
			ans: false,
		},
		{
			ransom: "aa",
			magazine: "ab",
			ans: false,
		},
		{
			ransom: "aa",
			magazine: "aab",
			ans: true,
		},
	}
	for _,ex:=range qs{
		ransomNote,magazineNote,as:=ex.ransom,ex.magazine,ex.ans
		res:=canConstruct(ransomNote,magazineNote)
		if res!=as{
			t.Errorf("error [input]:%v,%v   [output]:%v   [ans]:%v\n",ransomNote,magazineNote,res,as)
		}else{
			fmt.Printf("pass")
		}
	}
}