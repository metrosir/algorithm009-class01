package Week_02_everyday_test

import (
    "fmt"
    "testing"
)

func xier(arr []int) []int {
    length := len(arr)
    gap := 1
    //fmt.Println(gap/3)
    //for gap < gap/3 {
    //    gap = gap*3 + 1
    //}
    //fmt.Println(gap)
    for gap > 0 {
        for i := gap; i < length; i++ {
            temp := arr[i]
            j := i - gap
            for j >= 0 && arr[j] > temp {
                arr[j+gap] = arr[j]
                j -= gap
            }
            arr[j+gap] = temp
        }
        gap = gap / 3
    }
    return arr
}

func TestXier(t *testing.T)  {
    arr := []int{23,15,43,25,54,2,6,82,11,5,21,32,65}
    fmt.Println(xier(arr))
}
