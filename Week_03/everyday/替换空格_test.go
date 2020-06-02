package everyday_test

import (
	"fmt"
	"testing"
)

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
//示例 1：
//
//输入：s = "We are happy."
//输出："We%20are%20happy."
func replaceSpace(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			res += "%20"
		} else {
			res += string(s[i])
		}
	}
	return res
}

func TestReplace(t *testing.T) {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
