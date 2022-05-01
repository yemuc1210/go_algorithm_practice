package leetcode

import (
	"fmt"
	"testing"
)

type question5 struct {
	para5
	ans5
}

// para 是参数
// one 代表第一个参数
type para5 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans5 struct {
	one []string
}

func Test_Problem5(t *testing.T) {

	qs := []question5{

		{
			para5{"babad"},
			ans5{[]string{"bab","aba"}},
		},

		{
			para5{"cbbd"},
			ans5{[]string{"bb"}},
		},

		{
			para5{"a"},
			ans5{[]string{"a"}},
		},

		{
			para5{"ac"},
			ans5{[]string{"a","c"}},
		},

		{
			para5{"aa"},
			ans5{[]string{"aa"}},
		},

		{
			para5{"ajgiljtperkvubjmdsefcylksrxtftqrehoitdgdtttswwttmfuvwgwrruuqmxttzsbmuhgfaoueumvbhajqsaxkkihjwevzzedizmrsmpxqavyryklbotwzngxscvyuqjkkaotitddlhhnutmotupwuwyltebtsdfssbwayuxrbgihmtphshdslktvsjadaykyjivbzhwujcdvzdxxfiixnzrmusqvwujjmxhbqbdpauacnzojnzxxgrkmupadfcsujkcwajsgintahwgbjnvjqubcxajdyyapposrkpqtpqfjcvbhlmwfutgognqxgaukpmdyaxghgoqkqnigcllachmwzrazwhpppmsodvxilrccfqgpkmdqhoorxpyjsrtbeeidsinpeyxxpsjnymxkouskyhenzgieybwkgzrhhrzgkwbyeigznehyksuokxmynjspxxyepnisdieebtrsjypxroohqdmkpgqfccrlixvdosmppphwzarzwmhcallcginqkqoghgxaydmpkuagxqngogtufwmlhbvcjfqptqpkrsoppayydjaxcbuqjvnjbgwhatnigsjawckjuscfdapumkrgxxznjozncauapdbqbhxmjjuwvqsumrznxiifxxdzvdcjuwhzbvijykyadajsvtklsdhshptmhigbrxuyawbssfdstbetlywuwputomtunhhlddtitoakkjquyvcsxgnzwtoblkyryvaqxpmsrmzidezzvewjhikkxasqjahbvmueuoafghumbszttxmquurrwgwvufmttwwstttdgdtioherqtftxrsklycfesdmjbuvkreptjligja"},
			ans5{[]string{"ajgiljtperkvubjmdsefcylksrxtftqrehoitdgdtttswwttmfuvwgwrruuqmxttzsbmuhgfaoueumvbhajqsaxkkihjwevzzedizmrsmpxqavyryklbotwzngxscvyuqjkkaotitddlhhnutmotupwuwyltebtsdfssbwayuxrbgihmtphshdslktvsjadaykyjivbzhwujcdvzdxxfiixnzrmusqvwujjmxhbqbdpauacnzojnzxxgrkmupadfcsujkcwajsgintahwgbjnvjqubcxajdyyapposrkpqtpqfjcvbhlmwfutgognqxgaukpmdyaxghgoqkqnigcllachmwzrazwhpppmsodvxilrccfqgpkmdqhoorxpyjsrtbeeidsinpeyxxpsjnymxkouskyhenzgieybwkgzrhhrzgkwbyeigznehyksuokxmynjspxxyepnisdieebtrsjypxroohqdmkpgqfccrlixvdosmppphwzarzwmhcallcginqkqoghgxaydmpkuagxqngogtufwmlhbvcjfqptqpkrsoppayydjaxcbuqjvnjbgwhatnigsjawckjuscfdapumkrgxxznjozncauapdbqbhxmjjuwvqsumrznxiifxxdzvdcjuwhzbvijykyadajsvtklsdhshptmhigbrxuyawbssfdstbetlywuwputomtunhhlddtitoakkjquyvcsxgnzwtoblkyryvaqxpmsrmzidezzvewjhikkxasqjahbvmueuoafghumbszttxmquurrwgwvufmttwwstttdgdtioherqtftxrsklycfesdmjbuvkreptjligja"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 5------------------------\n")

	for _, q := range qs {
		as, p := q.ans5, q.para5

		res := longestPalindrome(p.s)
		if !res_in_ans(res,as.one){
			t.Errorf("error [input]:%v      [output]:%v\n",p.s,res)
		}

		//test push
		fmt.Printf("【input】:%v       【output】:%v\n", p, res)
	}
	fmt.Printf("\n\n\n")
}

func res_in_ans(res string, ans []string) bool{
	for _,v := range ans {
		if res == v {
			// fmt.Println(res,v)
			return true
		}
	}
	return false
}