package Week_02_everyday_test

import (
    "fmt"
    "testing"
)

func addDigits(num int) int {
    if num < 10 {
        return num
    }
    return (num - 1) % 9 + 1
    for num >= 10  {
        num = num / 10 + num % 10
        fmt.Println(num / 10, num % 10)
    }
    return num
}


func TestAddDigits(t *testing.T)  {
    num := 38
    fmt.Print(addDigits(num))
}