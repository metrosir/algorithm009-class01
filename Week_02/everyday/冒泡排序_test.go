package Week_02_everyday_test

import (
    "fmt"
    "testing"
)

func maop(arr []int) []int {
    cn := len(arr)
    for i := 0; i < cn ; i++  {
        for j := i + 1; j < cn; j++ {
            if arr[i] < arr[j] {
                tmp := arr[i]
                arr[i] = arr[j]
                arr[j] = tmp
            }
        }
    }
    return arr
}

func TestMaop(t *testing.T)  {
    arr := []int{23,15,43,25,54,2,6,82,11,5,21,32,65}
    fmt.Println(maop(arr))
}