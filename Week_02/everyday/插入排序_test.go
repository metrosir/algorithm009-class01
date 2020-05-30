package Week_02_everyday_test

import (
    "fmt"
    "testing"
)

func charu(arr []int) []int  {
    for i:=0; i<len(arr); i++ {
        index := i - 1
        val := arr[i]
        for index >= 0 && val < arr[index]  {
            arr[index + 1] = arr[index]
            index -= 1
        }
        arr[index+1] = val
    }
    return arr
}

func TestCharu(t *testing.T)  {
    arr := []int{23,15,43,25,54,2,6,82,11,5,21,32,65}
    //arr := []int{5,2}
    fmt.Println(charu(arr))
}