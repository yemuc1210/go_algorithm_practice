package main

/**
强密码：
6-20个字符
至少一个小写 一个大写  一个数字
同一个字符不能连续出现三次

给定一个字符，但会修改成强字符需要的最少修改次数
**/

var lower string = "abcdefghijklmnopqrstuvwxyz"
var upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digit string = "0123456789"

func strongPasswordChecker1(password string) int {
    hasLower, hasUpper, hasDigit := 0, 0, 0
    for _, ch := range password {
        if unicode.IsLower(ch) {
            hasLower = 1
        } else if unicode.IsUpper(ch) {
            hasUpper = 1
        } else if unicode.IsDigit(ch) {
            hasDigit = 1
        }
    }
    categories := hasLower + hasUpper + hasDigit

    switch n := len(password); {
    case n < 6:
        return max(6-n, 3-categories)
    case n <= 20:
        replace, cnt, cur := 0, 0, '#'
        for _, ch := range password {
            if ch == cur {
                cnt++
            } else {
                replace += cnt / 3
                cnt = 1
                cur = ch
            }
        }
        replace += cnt / 3
        return max(replace, 3-categories)
    default:
        // 替换次数和删除次数
        replace, remove := 0, n-20
        // k mod 3 = 1 的组数，即删除 2 个字符可以减少 1 次替换操作
        rm2, cnt, cur := 0, 0, '#'
        for _, ch := range password {
            if ch == cur {
                cnt++
                continue
            }
            if remove > 0 && cnt >= 3 {
                if cnt%3 == 0 {
                    // 如果是 k % 3 = 0 的组，那么优先删除 1 个字符，减少 1 次替换操作
                    remove--
                    replace--
                } else if cnt%3 == 1 {
                    // 如果是 k % 3 = 1 的组，那么存下来备用
                    rm2++
                }
                // k % 3 = 2 的组无需显式考虑
            }
            replace += cnt / 3
            cnt = 1
            cur = ch
        }

        if remove > 0 && cnt >= 3 {
            if cnt%3 == 0 {
                remove--
                replace--
            } else if cnt%3 == 1 {
                rm2++
            }
        }

        replace += cnt / 3

        // 使用 k % 3 = 1 的组的数量，由剩余的替换次数、组数和剩余的删除次数共同决定
        use2 := min(min(replace, rm2), remove/2)
        replace -= use2
        remove -= use2 * 2

        // 由于每有一次替换次数就一定有 3 个连续相同的字符（k / 3 决定），因此这里可以直接计算出使用 k % 3 = 2 的组的数量
        use3 := min(replace, remove/3)
        replace -= use3
        remove -= use3 * 3
        return (n - 20) + max(replace, 3-categories)
    }
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}


func strongPasswordChecker(password string) int {
	// 返回修改成强密码的最少修改步数
	if isStrongPassed(password) {
		return 0
	}
	// 否则需要修改
	// 插入  删除  替换
	n := len(password)
	if n<6 {
		return 6-n
	}
	if n>20 {
		return 20-n
	}
	// 先得到必须补充的字符个数 lack，然后按密码长度 n 分类讨论：

// n<6，可以再按 n==5 和 n<5 分类考虑，结果都满足 max(6-n, lack)
// 6<=n<=20，对于任一连续段长度 a，用 a//3 步替换操作即可满足条件“同一字符不连续出现三次 ”，
// 最后取 max(替换操作, lack)即可。
// n>20，必须进行删除操作。
// 优先删除连续段的字符，而且优先删除长度 a%3==0 的连续段，从而减少之后的替换操作。
// 于是考虑用堆，每轮选 a%3 最小的 a 进行删除操作，直到没有连续段长度 >=3 或 n==20 为止。
// 删除完成后，就转为上一情况了。
	// 删除只适合连续3个的情况   一般删除还得添加---->所以用在两侧添加替代
	return 0
}

func isStrongPassed(passwd string) bool {
	// 1. 长度
	if len(passwd)<6 || len(passwd) > 20 {
		return false
	}
	// 2. 字符种类
	var hasLower,hasUpper,hasDigit bool
	for i:=0;i<len(passwd);i++{
		if isLowerCase(passwd[i]){
			hasLower = true
		} else if isUpperCase(passwd[i]) {
			hasUpper = true
		}else if isDigit(passwd[i]){
			hasDigit = true
		}
	}
	if !(hasLower && hasUpper && hasDigit) {
		return false
	}
	// 3. 同一字符不能连续出现三次
	// 三个三个遍历
	n := len(passwd)
	for i:=0;i<=n-3;i++ {
		// 查看第一个是否和后面两个都相等
		if passwd[i] == passwd[i+1] && passwd[i] == passwd[i+2] {
			return false
		}
	}
	return true
}	

func isLowerCase(ch byte) bool {
	num := int(ch-'a')
	return num>=0 && num<=25
}
func isUpperCase(ch byte) bool {
	num := int(ch-'A')
	return num>=0 && num<=25
}
func isDigit(ch byte) bool {
	num := int(ch-'0')
	return num>=0 && num<=9
}