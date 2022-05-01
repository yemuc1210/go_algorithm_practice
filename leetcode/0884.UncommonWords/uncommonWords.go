package lt884
import(
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
    // map
    m := make(map[string]int)
	m1 := make(map[string]int)
    words1 := strings.Split(s1," ")
    words2 := strings.Split(s2," ")
    for _,word := range words1 {
        m[word] ++
    }
	for _,word := range words2 {
		m1[word]++
	}
    //删除非出现一次的
    for k,v := range m {
        if v != 1 {
            delete(m,k)
        }
    }
	for k,v := range m1 {
        if v != 1 {
            delete(m1,k)
        }
    }
	//查看另一个句子是否有 有则删除
	for _,word := range words2 {
		if _,exist:=m[word];exist {
			delete(m, word)
		}
	}
	for _,word := range words1 {
		if _,exist:=m1[word];exist {
			delete(m1, word)
		}
	}
	//返回
	res := []string{}
	for k := range m {
		res = append(res, k)
	}
	// 再来一次 2中出现 1没有的
	for k := range m1 {
		res = append(res, k)
	}
	return res
}