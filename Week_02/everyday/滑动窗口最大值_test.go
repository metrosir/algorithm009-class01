package Week_02_everyday_test

import (
    "fmt"
    "testing"
)

//{1,3,-1,-3,5,3,6,7}

//len: 8

//range = len - ()
//2 = len - 0
//3 = len - 1
//4 = len - 2
//...
//8 = 1
// len - (k - 2) 层 / 个桶

//[0,1] [1,2] [2,3] [3,4] [4,5] [5,6] [6,7] [7,8]
func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
    if len(nums) < k {
        return res
    }
    if k == 1 || k == len(nums) {
        return nums
    }
    t := len(nums) - (k - 2)
    temp_ := make([][]int, t)


    fmt.Println(temp_)
return  res
}




//0,1
//1,2
//2,3
//3,4
//4,5
//5,6
//6,7
//7,8
//2-8

//0,1,2
//1,2,3
//2,3,4
//3,4,5
//4,5,6
//5,6,7
//6,7,8
//3-7

//0,1,2,3
//1,2,3,4
//2,3,4,5
//3,4,5,6
//4,5,6,7
//5,6,7,8
//4-6

//0,1,2,3,4
//1,2,3,4,5
//2,3,4,5,6
//3,4,5,6,7
//4,5,6,7,8
//5-5


//func maxSlidingWindow(nums []int, k int) []int {
//    res := []int{}
//    if len(nums) < k {
//        return res
//    }
//    if k == 1 || k == len(nums) {
//        return nums
//    }
//    temp_arr1 := [][]int{}
//    temp_arr2 := []int{}
//    for i:=0; i<k ; i++  {
//
//
//    }
//    fmt.Println(temp_arr1)
//    return res
//}

func TestMaxSl(t *testing.T)  {
    nums := []int{1,3,-1,-3,5,3,6,7}
    k := 3
    fmt.Println(maxSlidingWindow(nums, k))
}