package main

import "fmt"

/**
正则匹配,'.'匹配任意字符，'*'表示0个或多个
 */

func main(){
    str := "abcda"
    pattern := "a.*.da"
    fmt.Println(match(str,pattern))

}

func match(str,pattern string) bool{
    strLen,patternLen := len(str),len(pattern)
    if strLen <= 0 || patternLen <= 0 {
        return false
    }
    return matchCore(0,0,strLen,patternLen,str,pattern)
}

func matchCore(strIndex,patternIndex,strLen,patternLen int,str,pattern string) bool {
    if strIndex == strLen && patternIndex == patternLen {
        return true
    }
    if strIndex < strLen && patternIndex >= patternLen {
        return false
    }
    if patternIndex + 1 < patternLen && pattern[patternIndex+1] == '*' {
        if strIndex != strLen && (str[strIndex] == pattern[patternIndex] || pattern[patternIndex] == '.') {
            //*前一个字符匹配上时，可以将*处理为匹配0次，匹配1次，匹配多次
            return matchCore(strIndex,patternIndex+2,strLen,patternLen,str,pattern) || matchCore(strIndex+1,patternIndex+2,strLen,patternLen,str,pattern) || matchCore(strIndex+1,patternIndex,strLen,patternLen,str,pattern)
        }else {
            //*前一个字符未匹配上时，直接将*处理为匹配0次
            return matchCore(strIndex,patternIndex+2,strLen,patternLen,str,pattern)
        }
    }
    if strIndex != strLen && (str[strIndex] == pattern[patternIndex] || pattern[patternIndex] == '.') {
        return matchCore(strIndex+1,patternIndex+1,strLen,patternLen,str,pattern)
    }
    return false
}
