package Week_02_everyday_test

import (
    "fmt"
    "strings"
    "testing"
)

//删除最外层的括号，返回结果

//(()())(())

//暴力
//(


func removeOuterParentheses(S string) string {
    if len(S) == 0 {
        return ""
    }

    //先把左括号入栈
    stack := []byte{S[0]}
    //设置开始指针为1
    start := 1
    var res strings.Builder
    for i := 1; i < len(S) ; i++ {
        //栈的长度大于0 的情况下，判断当前循环的坐标是否是左括号以及栈中的倒数第二个括号是否是右括号
        if len(stack) > 0 && S[i] == ')' && stack[len(stack) - 1] == '(' {
            stack = stack[:len(stack) - 1]
            if len(stack) == 0 {
                res.WriteString(S[start:i])
                start = i + 2
            }
        } else {
            stack = append(stack, S[i])
        }
    }

    return res.String()
}

func TestRe(t *testing.T)  {
    str := "(()())(())"
    fmt.Print(removeOuterParentheses(str))
}