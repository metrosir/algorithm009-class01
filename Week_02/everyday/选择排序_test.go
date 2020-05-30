package Week_02_everyday_test

import (
    "fmt"
    "testing"
)

func xuanz(arr []int) []int {
    for i:=0; i<len(arr) ; i++ {
        min := i
        for j := i+1; j < len(arr) ; j++  {
            if arr[min] > arr[j] {
                min = j
            }
        }
        if min != i {
            tmp := arr[i]
            arr[i] = arr[min]
            arr[min] = tmp
        }
    }
    return arr
}

func TestXuanz(t *testing.T)  {
    arr := []int{23,15,43,25,54,2,6,82,11,5,21,32,65}
    fmt.Println(xuanz(arr))
}
