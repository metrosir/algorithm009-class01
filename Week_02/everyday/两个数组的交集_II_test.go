package Week_02_everyday_test

import (
    "fmt"
    "testing"
)

//link: https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2,2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [4,9]
//说明：
//
//输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//我们可以不考虑输出结果的顺序。
//进阶:
//
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果 nums1 的大小比 nums2 小很多，哪种方法更优？
//如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？


func intersect(nums1 []int, nums2 []int) []int {

    res := []int{}
    hashMap := map[int]int{}
    if len(nums1) > len(nums2) {
        for i:=0; i < len(nums2); i++  {
            if _, ok := hashMap[nums2[i]]; !ok {
                hashMap[nums2[i]] = 1
            } else {
                hashMap[nums2[i]] ++
            }
        }
        for i := 0; i < len(nums1) ; i++ {
            if hashMap[nums1[i]] > 0 {
                res = append(res, nums1[i])
                hashMap[nums1[i]] --
            }
        }
    } else {
        for i:=0; i < len(nums1); i++  {
            if _, ok := hashMap[nums1[i]]; !ok {
                hashMap[nums1[i]] = 1
            } else {
                hashMap[nums1[i]] ++
            }
        }
        for i := 0; i < len(nums2) ; i++ {
            if hashMap[nums2[i]] > 0 {
                res = append(res, nums2[i])
                hashMap[nums2[i]] --
            }
        }
    }

    return res
}
func TestIntersect(t *testing.T) {
    num1 := []int{1,2,2,1}
    num2 := []int{2,2}
    fmt.Println(intersect(num1, num2))
}