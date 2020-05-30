package Week_02_everyday

import (
    "fmt"
    "testing"
)

// https://leetcode-cn.com/problems/number-of-segments-in-a-string/

//统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
//
//请注意，你可以假定字符串里不包括任何不可打印的字符。
//
//示例:
//
//输入: "Hello, my name is John"
//输出: 5
//解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。

func countSegments(s string) int {
    if len(s) < 1 {
        return 0
    }
    arr := map[int]string{}
    tmp_nb := 0
    for i := 0; i < len(s); i++  {
        if s[i] != ' ' {
            arr[tmp_nb] = arr[tmp_nb] + string(s[i])
        } else {
            tmp_nb ++
        }
    }

    return len(arr)
}

func TestCountSegments(t *testing.T)  {

    fmt.Print(countSegments("Hello, my name is John     John"))
}

