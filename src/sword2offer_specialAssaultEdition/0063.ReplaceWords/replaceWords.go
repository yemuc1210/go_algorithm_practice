package offerS63

import (
	"sort"
	"strings"
)

/*
将继承词换成词根
*/
func replaceWords(dictionary []string, sentence string) string {
        sort.Strings(dictionary)

        sentences := strings.Split(sentence, " ")

        for idx, s := range sentences {
                for _, s2 := range dictionary {
                        if strings.HasPrefix(s, s2) {
                                sentences[idx] = s2
                                break
                        }
                }

        }

        return strings.Join(sentences, " ")
}
