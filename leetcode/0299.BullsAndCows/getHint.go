package lt299

import "fmt"

/*
299. 猜数字游戏
你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：

写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：

	猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls", 公牛），
	有多少位属于数字猜对了但是位置不对（称为 "Cows", 奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字 secret 和朋友猜测的数字 guess ，请你返回对朋友这次猜测的提示。

提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B 表示奶牛。

请注意秘密数字和朋友猜测的数字都可能含有重复数字。
*/

/*
公牛数 遍历即可得到
奶牛数：数字对位置不对
	不匹配时，统计secret和guess各个字符的出现字数

*/

func getHint(secret string, guess string) string {
	bulls := 0
	var cntS,cntG [10]int
	for i:= range secret {
		if secret[i] == guess[i] {
			bulls ++
		}else{  //统计不在其位的数字
			cntS[secret[i]-'0'] ++
			cntG[guess[i]-'0'] ++
		}
	}

	cows := 0
	for i:=0;i<10;i++{
		cows += min(cntG[i],cntS[i])
	}
	return fmt.Sprintf("%dA%dB",bulls,cows)
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}