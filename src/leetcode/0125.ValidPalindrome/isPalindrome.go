package lt125

import "strings"

/*
验证回文串
*/
func isPalindrome(s string) bool {
    var sgood string
    for i := 0; i < len(s); i++ {
        if isalnum(s[i]) {
            sgood += string(s[i])
        }
    }

    n := len(sgood)
    sgood = strings.ToLower(sgood)
    for i := 0; i < n/2; i++ {
        if sgood[i] != sgood[n-1-i] {
            return false
        }
    }
    return true
}

func isalnum(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
