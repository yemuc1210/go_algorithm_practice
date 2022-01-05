package main

import (
	"fmt"
	"math/rand"
	"time"
)

//替换所有的问号
func modifyString(s string) string {
	//将？替换成若干小写字母，让最终的字符串不含重复字符  任选其一
	// fmt.Println(string(chars))
	//不允许连续重复。。。。 直接三指针
	pre,mid,rear := 0,1,2
	if len(s) == 0 || len(s) == 1{
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return string(s[0]+1) + string(s[1])
		}else{
			return s
		}
	}
	bytes := []byte(s)
	//否则至少有三个
	//处理第一个就是问好的
	if bytes[pre] == '?' {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r1 := rand.Intn(27)
		c2 := int(bytes[mid]-'a')
        for r1 == int(bytes[mid]-'a') || r1+c2 >= 26{ 
			r1 = r.Intn(27)
		}
		bytes[pre] = byte(r1+'a')
	}
	for i:=1;i<len(bytes)-1;i++{
		
		if bytes[mid] == '?' {
			//与其左右不等
			c1,c3 := int(bytes[pre]-'a'),int(bytes[rear]-'a')
			c2 := int(bytes[mid]-'a')
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			r1 := rand.Intn(27)
			for (r1==c1 || r1==c3) && (r1+c2>=26) {
				r1 = r.Intn(27)
			}
			bytes[mid] = byte(r1+'a')
			// return string(bytes)
		}//所有问号都要处理
		pre,mid,rear = pre+1,mid+1,rear+1
	}
	//最后一个可能是
	if bytes[mid] == '?' {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r1 := rand.Intn(27)
		c2 := int(bytes[mid-1]-'a')
        for r1 == int(bytes[mid-1]-'a') || r1+c2 >= 26{ 
			r1 = r.Intn(27)
		}
		bytes[mid-1] = byte(r1+'a')
	}
	return string(bytes)
}
func modifyString1(s string) string {
    res := []byte(s)
    n := len(res)
    for i, ch := range res {
        if ch == '?' {
            for b := byte('a'); b <= 'c'; b++ {
                if !(i > 0 && res[i-1] == b || i < n-1 && res[i+1] == b) {
                    res[i] = b
                    break
                }
            }
        }
    }
    return string(res)
}
func main(){
	s := modifyString1("??yw?ipkj?")
	fmt.Println(s)
}