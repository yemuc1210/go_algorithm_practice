package lt68

import "strings"

/*
每日一题

68. 文本左右对齐
给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

说明:

	单词是指由非空格字符组成的字符序列。
	每个单词的长度大于 0，小于等于 maxWidth。
	输入单词数组 words 至少包含一个单词。



首先确定放置多少单词，然后得到空格数量，再确定每个单词之间的空格数
三种情况：
	当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格；
	当前行不是最后一行，且只有一个单词：该单词左对齐，在行末填充空格；
	当前行不是最后一行，且不只一个单词：设当前行单词数为 numWords，空格数为 numSpaces，
		我们需要将空格均匀分配在单词之间，则单词之间应至少有	avgspave = numspace / (numwords-1)  向下取整
		多出的空格extraspace = numspace mod (numwords-1)填充在前extraspace个单词之间，所以前
		extraspace个单词之间填充avgspace+1个
*/
func fullJustify(words []string, maxWidth int) (ans []string) {
	right, n := 0, len(words)
    for {
        left := right // 当前行的第一个单词在 words 的位置
        sumLen := 0   // 统计这一行单词长度之和
        // 循环确定当前行可以放多少单词，注意单词之间应至少有一个空格
        for right < n && sumLen+len(words[right])+right-left <= maxWidth {
            sumLen += len(words[right])
            right++
        }

        // 当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格
        if right == n {
            s := strings.Join(words[left:], " ")
            ans = append(ans, s+blank(maxWidth-len(s)))
            return
        }

        numWords := right - left
        numSpaces := maxWidth - sumLen

        // 当前行只有一个单词：该单词左对齐，在行末填充剩余空格
        if numWords == 1 {
            ans = append(ans, words[left]+blank(numSpaces))
            continue
        }

        // 当前行不只一个单词
        avgSpaces := numSpaces / (numWords - 1)
        extraSpaces := numSpaces % (numWords - 1)
        s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1)) // 拼接额外加一个空格的单词
        s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))  // 拼接其余单词
        ans = append(ans, s1+blank(avgSpaces)+s2)
    }

}

func blank(n int) string {
    return strings.Repeat(" ", n)
}